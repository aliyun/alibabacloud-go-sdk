// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSessionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateSessionResponseBody
	GetRequestId() *string
	SetSessionId(v string) *CreateSessionResponseBody
	GetSessionId() *string
}

type CreateSessionResponseBody struct {
	// example:
	//
	// 44553E9A-******-37ADC33FE2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// se-weg******iqwcw
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
}

func (s CreateSessionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateSessionResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSessionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateSessionResponseBody) GetSessionId() *string {
	return s.SessionId
}

func (s *CreateSessionResponseBody) SetRequestId(v string) *CreateSessionResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSessionResponseBody) SetSessionId(v string) *CreateSessionResponseBody {
	s.SessionId = &v
	return s
}

func (s *CreateSessionResponseBody) Validate() error {
	return dara.Validate(s)
}
