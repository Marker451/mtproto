package mtproto

const (
	SESSION MEventType = "session"
	MCONN   MEventType = "mconn"
)

type MEventType string
type MEvent interface {
	Type() MEventType
}

// Session Events
type newsession struct {
	// If connId is zero, MManager makes new connection and assigns it the new session.
	// Otherwise, the new session is allocated to the connection of connId.
	connId      int32
	phonenumber string
	addr        string
	useIPv6     bool
	resp        chan sessionResponse
}

type loadsession struct {
	// If connId is zero, MManager makes new connection and assigns it the loaded session.
	// Otherwise, the loaded session is allocated to the connection of connId.
	connId        int32
	phonenumber   string
	preferredAddr string
	resp          chan sessionResponse
}

type sessionResponse struct {
	connId  int32
	session *MSession
	//mconn 	*MConn
	err error
}

// Established = made + bound
type SessionEstablished struct {
	//connId		int32
	session *MSession
	//sessionId 	int64
	//mconn 	*MConn
}

type discardSession struct {
	connId    int32
	sessionId int64
	resp      chan sessionResponse
}

type SessionDiscarded struct {
	//connId		int32
	//sessionId 	int64
	boundConnId                  int32
	discardedSessionId           int64
	discardedSessionUpdatesState TL_updates_state
}

// discardSession + newsession
type renewSession struct {
	//connId		int32
	sessionId   int64
	phonenumber string
	addr        string
	useIPv6     bool
	resp        chan sessionResponse
}

// discardSession + loadsession
type refreshSession struct {
	//connId		int32
	sessionId   int64
	phonenumber string
	resp        chan sessionResponse
}

// Connection Events
type ConnectionOpened struct {
	mconn *MConn
}
type sessionBound struct {
	mconn *MConn
}
type sessionUnbound struct {
	mconn            *MConn
	unboundSessionId int64
}
type closeConnection struct {
	connId int32
	resp   chan error
}
type connectionClosed struct {
	closedConnId int32
}

// Update Event
type updateReceived struct {
	update MUpdate
}

func (e newsession) Type() MEventType         { return SESSION }
func (e loadsession) Type() MEventType        { return SESSION }
func (e SessionEstablished) Type() MEventType { return SESSION }
func (e renewSession) Type() MEventType       { return SESSION }
func (e refreshSession) Type() MEventType     { return SESSION }
func (e discardSession) Type() MEventType     { return SESSION }
func (e SessionDiscarded) Type() MEventType   { return SESSION }
func (e ConnectionOpened) Type() MEventType   { return MCONN }
func (e sessionBound) Type() MEventType       { return MCONN }
func (e sessionUnbound) Type() MEventType     { return MCONN }
func (e closeConnection) Type() MEventType    { return MCONN }
func (e connectionClosed) Type() MEventType   { return MCONN }
func (e updateReceived) Type() MEventType     { return SESSION }

//func (e newsession) SessionId() (int64)          {return 0}
//func (e loadsession) SessionId() (int64)         {return 0}
//func (e SessionEstablished) SessionId() (int64)  {return e.session.sessionId}
//func (e renewSession) SessionId() (int64)        {return e.sessionId}
//func (e refreshSession) SessionId() (int64)      {return e.sessionId}
//func (e discardSession) SessionId() (int64)   {return e.sessionId}
//func (e SessionDiscarded) SessionId() (int64) {return e.discardedSessionId}
//func (e ConnectionOpened) SessionId() (int64) {return 0}
//func (e sessionBound) SessionId() (int64) {
//	session, err := e.mconn.Session()
//	if err != nil {
//		return 0
//	}
//	return session.sessionId
//}
//func (e sessionUnbound) SessionId() (int64) 		{return e.unboundSessionId}
//func (e closeConnection) SessionId() (int64) 	{return 0}
//func (e connectionClosed) SessionId() (int64) 	{return 0}
//
//func (e newsession) ConnectionId() (int32)         {return 0}
//func (e loadsession) ConnectionId() (int32)        {return 0}
//func (e SessionEstablished) ConnectionId() (int32) {return e.connId}
//func (e renewSession) ConnectionId() (int32)       {return e.connId}
//func (e refreshSession) ConnectionId() (int32)     {return e.connId}
//func (e discardSession) ConnectionId() (int32)     {return e.connId}
//func (e SessionDiscarded) ConnectionId() (int32)   {return e.boundConnId}
//func (e ConnectionOpened) ConnectionId() (int32)   {return e.mconn.connId}
//func (e sessionBound) ConnectionId() (int32)       {return e.mconn.connId}
//func (e sessionUnbound) ConnectionId() (int32)     {return e.mconn.connId}
//func (e closeConnection) ConnectionId() (int32)    {return e.connId}
//func (e connectionClosed) ConnectionId() (int32)   {return e.closedConnId}
