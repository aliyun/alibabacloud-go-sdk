// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSessionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListSessionsResponseBody
	GetRequestId() *string
	SetSessions(v []*Session) *ListSessionsResponseBody
	GetSessions() []*Session
	SetTotalCount(v int64) *ListSessionsResponseBody
	GetTotalCount() *int64
}

type ListSessionsResponseBody struct {
	// example:
	//
	// 44553E9A-******-37ADC33FE2
	RequestId *string    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Sessions  []*Session `json:"Sessions,omitempty" xml:"Sessions,omitempty" type:"Repeated"`
	// example:
	//
	// 2
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListSessionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSessionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListSessionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSessionsResponseBody) GetSessions() []*Session {
	return s.Sessions
}

func (s *ListSessionsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListSessionsResponseBody) SetRequestId(v string) *ListSessionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSessionsResponseBody) SetSessions(v []*Session) *ListSessionsResponseBody {
	s.Sessions = v
	return s
}

func (s *ListSessionsResponseBody) SetTotalCount(v int64) *ListSessionsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListSessionsResponseBody) Validate() error {
	if s.Sessions != nil {
		for _, item := range s.Sessions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
