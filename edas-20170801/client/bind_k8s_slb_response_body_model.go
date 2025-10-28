// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindK8sSlbResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetChangeOrderId(v string) *BindK8sSlbResponseBody
	GetChangeOrderId() *string
	SetCode(v int32) *BindK8sSlbResponseBody
	GetCode() *int32
	SetMessage(v string) *BindK8sSlbResponseBody
	GetMessage() *string
	SetRequestId(v string) *BindK8sSlbResponseBody
	GetRequestId() *string
}

type BindK8sSlbResponseBody struct {
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
	// 4823-bhjf-23u4-****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BindK8sSlbResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BindK8sSlbResponseBody) GoString() string {
	return s.String()
}

func (s *BindK8sSlbResponseBody) GetChangeOrderId() *string {
	return s.ChangeOrderId
}

func (s *BindK8sSlbResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *BindK8sSlbResponseBody) GetMessage() *string {
	return s.Message
}

func (s *BindK8sSlbResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BindK8sSlbResponseBody) SetChangeOrderId(v string) *BindK8sSlbResponseBody {
	s.ChangeOrderId = &v
	return s
}

func (s *BindK8sSlbResponseBody) SetCode(v int32) *BindK8sSlbResponseBody {
	s.Code = &v
	return s
}

func (s *BindK8sSlbResponseBody) SetMessage(v string) *BindK8sSlbResponseBody {
	s.Message = &v
	return s
}

func (s *BindK8sSlbResponseBody) SetRequestId(v string) *BindK8sSlbResponseBody {
	s.RequestId = &v
	return s
}

func (s *BindK8sSlbResponseBody) Validate() error {
	return dara.Validate(s)
}
