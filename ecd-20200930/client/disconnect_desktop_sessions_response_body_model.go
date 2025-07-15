// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisconnectDesktopSessionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInvalidSessions(v []*DisconnectDesktopSessionsResponseBodyInvalidSessions) *DisconnectDesktopSessionsResponseBody
	GetInvalidSessions() []*DisconnectDesktopSessionsResponseBodyInvalidSessions
	SetRequestId(v string) *DisconnectDesktopSessionsResponseBody
	GetRequestId() *string
}

type DisconnectDesktopSessionsResponseBody struct {
	// The list of invalid sessions.
	InvalidSessions []*DisconnectDesktopSessionsResponseBodyInvalidSessions `json:"InvalidSessions,omitempty" xml:"InvalidSessions,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 2507CFA8-FEAB-5208-98F5-5E028C50XXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisconnectDesktopSessionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DisconnectDesktopSessionsResponseBody) GoString() string {
	return s.String()
}

func (s *DisconnectDesktopSessionsResponseBody) GetInvalidSessions() []*DisconnectDesktopSessionsResponseBodyInvalidSessions {
	return s.InvalidSessions
}

func (s *DisconnectDesktopSessionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DisconnectDesktopSessionsResponseBody) SetInvalidSessions(v []*DisconnectDesktopSessionsResponseBodyInvalidSessions) *DisconnectDesktopSessionsResponseBody {
	s.InvalidSessions = v
	return s
}

func (s *DisconnectDesktopSessionsResponseBody) SetRequestId(v string) *DisconnectDesktopSessionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisconnectDesktopSessionsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DisconnectDesktopSessionsResponseBodyInvalidSessions struct {
	// The cloud desktop ID.
	//
	// example:
	//
	// ecd-2jv6wugbkp65pxxxx
	DesktopId *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
	// The end user ID.
	//
	// example:
	//
	// wy01
	EndUserId *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
}

func (s DisconnectDesktopSessionsResponseBodyInvalidSessions) String() string {
	return dara.Prettify(s)
}

func (s DisconnectDesktopSessionsResponseBodyInvalidSessions) GoString() string {
	return s.String()
}

func (s *DisconnectDesktopSessionsResponseBodyInvalidSessions) GetDesktopId() *string {
	return s.DesktopId
}

func (s *DisconnectDesktopSessionsResponseBodyInvalidSessions) GetEndUserId() *string {
	return s.EndUserId
}

func (s *DisconnectDesktopSessionsResponseBodyInvalidSessions) SetDesktopId(v string) *DisconnectDesktopSessionsResponseBodyInvalidSessions {
	s.DesktopId = &v
	return s
}

func (s *DisconnectDesktopSessionsResponseBodyInvalidSessions) SetEndUserId(v string) *DisconnectDesktopSessionsResponseBodyInvalidSessions {
	s.EndUserId = &v
	return s
}

func (s *DisconnectDesktopSessionsResponseBodyInvalidSessions) Validate() error {
	return dara.Validate(s)
}
