// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePhysicalConnectionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPhysicalConnectionId(v string) *CreatePhysicalConnectionResponseBody
	GetPhysicalConnectionId() *string
	SetRequestId(v string) *CreatePhysicalConnectionResponseBody
	GetRequestId() *string
}

type CreatePhysicalConnectionResponseBody struct {
	// The ID of the Express Connect circuit.
	//
	// example:
	//
	// pc-bp1ciz7ekd2grn1as****
	PhysicalConnectionId *string `json:"PhysicalConnectionId,omitempty" xml:"PhysicalConnectionId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 8A6A5EC5-6F6C-4906-9689-56ACE58A13E0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreatePhysicalConnectionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreatePhysicalConnectionResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePhysicalConnectionResponseBody) GetPhysicalConnectionId() *string {
	return s.PhysicalConnectionId
}

func (s *CreatePhysicalConnectionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreatePhysicalConnectionResponseBody) SetPhysicalConnectionId(v string) *CreatePhysicalConnectionResponseBody {
	s.PhysicalConnectionId = &v
	return s
}

func (s *CreatePhysicalConnectionResponseBody) SetRequestId(v string) *CreatePhysicalConnectionResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreatePhysicalConnectionResponseBody) Validate() error {
	return dara.Validate(s)
}
