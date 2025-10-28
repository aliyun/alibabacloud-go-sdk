// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateK8sServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetChangeOrderId(v string) *UpdateK8sServiceResponseBody
	GetChangeOrderId() *string
	SetCode(v int32) *UpdateK8sServiceResponseBody
	GetCode() *int32
	SetMessage(v string) *UpdateK8sServiceResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateK8sServiceResponseBody
	GetRequestId() *string
}

type UpdateK8sServiceResponseBody struct {
	// The ID of the change process.
	//
	// example:
	//
	// b4b37bde-a125-****-****-741f7f4a9ae3
	ChangeOrderId *string `json:"ChangeOrderId,omitempty" xml:"ChangeOrderId,omitempty"`
	// The HTTP status code that is returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The additional information that is returned.
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

func (s UpdateK8sServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateK8sServiceResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateK8sServiceResponseBody) GetChangeOrderId() *string {
	return s.ChangeOrderId
}

func (s *UpdateK8sServiceResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *UpdateK8sServiceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateK8sServiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateK8sServiceResponseBody) SetChangeOrderId(v string) *UpdateK8sServiceResponseBody {
	s.ChangeOrderId = &v
	return s
}

func (s *UpdateK8sServiceResponseBody) SetCode(v int32) *UpdateK8sServiceResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateK8sServiceResponseBody) SetMessage(v string) *UpdateK8sServiceResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateK8sServiceResponseBody) SetRequestId(v string) *UpdateK8sServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateK8sServiceResponseBody) Validate() error {
	return dara.Validate(s)
}
