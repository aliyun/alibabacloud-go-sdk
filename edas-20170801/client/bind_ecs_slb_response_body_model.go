// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindEcsSlbResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetChangeOrderId(v string) *BindEcsSlbResponseBody
	GetChangeOrderId() *string
	SetCode(v int32) *BindEcsSlbResponseBody
	GetCode() *int32
	SetMessage(v string) *BindEcsSlbResponseBody
	GetMessage() *string
	SetRequestId(v string) *BindEcsSlbResponseBody
	GetRequestId() *string
}

type BindEcsSlbResponseBody struct {
	// The change process ID for this operation.
	//
	// example:
	//
	// cd65b247-****-475b-ad4b-7039040d625c
	ChangeOrderId *string `json:"ChangeOrderId,omitempty" xml:"ChangeOrderId,omitempty"`
	// The HTTP status code that is returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The message that is returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 03FD1520-0FD6-436A-****-265318D7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BindEcsSlbResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BindEcsSlbResponseBody) GoString() string {
	return s.String()
}

func (s *BindEcsSlbResponseBody) GetChangeOrderId() *string {
	return s.ChangeOrderId
}

func (s *BindEcsSlbResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *BindEcsSlbResponseBody) GetMessage() *string {
	return s.Message
}

func (s *BindEcsSlbResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BindEcsSlbResponseBody) SetChangeOrderId(v string) *BindEcsSlbResponseBody {
	s.ChangeOrderId = &v
	return s
}

func (s *BindEcsSlbResponseBody) SetCode(v int32) *BindEcsSlbResponseBody {
	s.Code = &v
	return s
}

func (s *BindEcsSlbResponseBody) SetMessage(v string) *BindEcsSlbResponseBody {
	s.Message = &v
	return s
}

func (s *BindEcsSlbResponseBody) SetRequestId(v string) *BindEcsSlbResponseBody {
	s.RequestId = &v
	return s
}

func (s *BindEcsSlbResponseBody) Validate() error {
	return dara.Validate(s)
}
