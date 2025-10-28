// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateK8sServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetChangeOrderId(v string) *CreateK8sServiceResponseBody
	GetChangeOrderId() *string
	SetCode(v int32) *CreateK8sServiceResponseBody
	GetCode() *int32
	SetMessage(v string) *CreateK8sServiceResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateK8sServiceResponseBody
	GetRequestId() *string
}

type CreateK8sServiceResponseBody struct {
	// The change process ID.
	//
	// example:
	//
	// b4b37bde-a125-****-****-741f7f4a9ae3
	ChangeOrderId *string `json:"ChangeOrderId,omitempty" xml:"ChangeOrderId,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The additional information returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 4823-bhjf-23u4-eiufh
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateK8sServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateK8sServiceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateK8sServiceResponseBody) GetChangeOrderId() *string {
	return s.ChangeOrderId
}

func (s *CreateK8sServiceResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *CreateK8sServiceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateK8sServiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateK8sServiceResponseBody) SetChangeOrderId(v string) *CreateK8sServiceResponseBody {
	s.ChangeOrderId = &v
	return s
}

func (s *CreateK8sServiceResponseBody) SetCode(v int32) *CreateK8sServiceResponseBody {
	s.Code = &v
	return s
}

func (s *CreateK8sServiceResponseBody) SetMessage(v string) *CreateK8sServiceResponseBody {
	s.Message = &v
	return s
}

func (s *CreateK8sServiceResponseBody) SetRequestId(v string) *CreateK8sServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateK8sServiceResponseBody) Validate() error {
	return dara.Validate(s)
}
