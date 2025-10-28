// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateK8sSlbResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetChangeOrderId(v string) *UpdateK8sSlbResponseBody
	GetChangeOrderId() *string
	SetCode(v int32) *UpdateK8sSlbResponseBody
	GetCode() *int32
	SetMessage(v string) *UpdateK8sSlbResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateK8sSlbResponseBody
	GetRequestId() *string
}

type UpdateK8sSlbResponseBody struct {
	// The ID of the change process.
	//
	// example:
	//
	// 9a1dcdee-****-****-ad37-cbf9dc91fba9
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
	// 4823-bhjf-23u4-eiufh
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateK8sSlbResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateK8sSlbResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateK8sSlbResponseBody) GetChangeOrderId() *string {
	return s.ChangeOrderId
}

func (s *UpdateK8sSlbResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *UpdateK8sSlbResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateK8sSlbResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateK8sSlbResponseBody) SetChangeOrderId(v string) *UpdateK8sSlbResponseBody {
	s.ChangeOrderId = &v
	return s
}

func (s *UpdateK8sSlbResponseBody) SetCode(v int32) *UpdateK8sSlbResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateK8sSlbResponseBody) SetMessage(v string) *UpdateK8sSlbResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateK8sSlbResponseBody) SetRequestId(v string) *UpdateK8sSlbResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateK8sSlbResponseBody) Validate() error {
	return dara.Validate(s)
}
