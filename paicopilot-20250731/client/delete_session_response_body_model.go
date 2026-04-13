// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSessionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteSessionResponseBody
	GetRequestId() *string
	SetSessionId(v string) *DeleteSessionResponseBody
	GetSessionId() *string
}

type DeleteSessionResponseBody struct {
	// example:
	//
	// 44553E9A-******-37ADC33FE2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// se-p4k******vx2y
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
}

func (s DeleteSessionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteSessionResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSessionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteSessionResponseBody) GetSessionId() *string {
	return s.SessionId
}

func (s *DeleteSessionResponseBody) SetRequestId(v string) *DeleteSessionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteSessionResponseBody) SetSessionId(v string) *DeleteSessionResponseBody {
	s.SessionId = &v
	return s
}

func (s *DeleteSessionResponseBody) Validate() error {
	return dara.Validate(s)
}
