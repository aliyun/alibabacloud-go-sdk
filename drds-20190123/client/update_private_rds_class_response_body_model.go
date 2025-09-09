// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePrivateRdsClassResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *UpdatePrivateRdsClassResponseBody
	GetData() *string
	SetRequestId(v string) *UpdatePrivateRdsClassResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdatePrivateRdsClassResponseBody
	GetSuccess() *bool
}

type UpdatePrivateRdsClassResponseBody struct {
	// The ID of the order.
	//
	// example:
	//
	// {     "orderId": "209136011******"   }
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 57D86AB4-8703-4DF4-BAB6-F7DE44******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdatePrivateRdsClassResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdatePrivateRdsClassResponseBody) GoString() string {
	return s.String()
}

func (s *UpdatePrivateRdsClassResponseBody) GetData() *string {
	return s.Data
}

func (s *UpdatePrivateRdsClassResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdatePrivateRdsClassResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdatePrivateRdsClassResponseBody) SetData(v string) *UpdatePrivateRdsClassResponseBody {
	s.Data = &v
	return s
}

func (s *UpdatePrivateRdsClassResponseBody) SetRequestId(v string) *UpdatePrivateRdsClassResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdatePrivateRdsClassResponseBody) SetSuccess(v bool) *UpdatePrivateRdsClassResponseBody {
	s.Success = &v
	return s
}

func (s *UpdatePrivateRdsClassResponseBody) Validate() error {
	return dara.Validate(s)
}
