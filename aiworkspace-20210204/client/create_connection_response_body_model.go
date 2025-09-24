// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateConnectionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConnectionId(v string) *CreateConnectionResponseBody
	GetConnectionId() *string
	SetRequestId(v string) *CreateConnectionResponseBody
	GetRequestId() *string
}

type CreateConnectionResponseBody struct {
	// The connection ID.
	//
	// example:
	//
	// conn-pai9m***mi47
	ConnectionId *string `json:"ConnectionId,omitempty" xml:"ConnectionId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 5A14FA81-DD4E-******-6343FE44B941
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateConnectionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateConnectionResponseBody) GoString() string {
	return s.String()
}

func (s *CreateConnectionResponseBody) GetConnectionId() *string {
	return s.ConnectionId
}

func (s *CreateConnectionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateConnectionResponseBody) SetConnectionId(v string) *CreateConnectionResponseBody {
	s.ConnectionId = &v
	return s
}

func (s *CreateConnectionResponseBody) SetRequestId(v string) *CreateConnectionResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateConnectionResponseBody) Validate() error {
	return dara.Validate(s)
}
