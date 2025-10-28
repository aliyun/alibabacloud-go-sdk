// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnbindK8sSlbResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetChangeOrderId(v string) *UnbindK8sSlbResponseBody
	GetChangeOrderId() *string
	SetCode(v int32) *UnbindK8sSlbResponseBody
	GetCode() *int32
	SetMessage(v string) *UnbindK8sSlbResponseBody
	GetMessage() *string
	SetRequestId(v string) *UnbindK8sSlbResponseBody
	GetRequestId() *string
}

type UnbindK8sSlbResponseBody struct {
	// The ID of the change process.
	//
	// example:
	//
	// b0a8441e-****-4e8e-9874-b56dea02952f
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
	// 1234-1sda-321d-1234
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UnbindK8sSlbResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UnbindK8sSlbResponseBody) GoString() string {
	return s.String()
}

func (s *UnbindK8sSlbResponseBody) GetChangeOrderId() *string {
	return s.ChangeOrderId
}

func (s *UnbindK8sSlbResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *UnbindK8sSlbResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UnbindK8sSlbResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UnbindK8sSlbResponseBody) SetChangeOrderId(v string) *UnbindK8sSlbResponseBody {
	s.ChangeOrderId = &v
	return s
}

func (s *UnbindK8sSlbResponseBody) SetCode(v int32) *UnbindK8sSlbResponseBody {
	s.Code = &v
	return s
}

func (s *UnbindK8sSlbResponseBody) SetMessage(v string) *UnbindK8sSlbResponseBody {
	s.Message = &v
	return s
}

func (s *UnbindK8sSlbResponseBody) SetRequestId(v string) *UnbindK8sSlbResponseBody {
	s.RequestId = &v
	return s
}

func (s *UnbindK8sSlbResponseBody) Validate() error {
	return dara.Validate(s)
}
