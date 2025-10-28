// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iScaleK8sApplicationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetChangeOrderId(v string) *ScaleK8sApplicationResponseBody
	GetChangeOrderId() *string
	SetCode(v int32) *ScaleK8sApplicationResponseBody
	GetCode() *int32
	SetMessage(v string) *ScaleK8sApplicationResponseBody
	GetMessage() *string
	SetRequestId(v string) *ScaleK8sApplicationResponseBody
	GetRequestId() *string
}

type ScaleK8sApplicationResponseBody struct {
	// The ID of the change process. You can call the GetChangeOrderInfo operation to query the progress of this scaling operation. For more information, see [GetChangeOrderInfo](https://help.aliyun.com/document_detail/62072.html).
	//
	// example:
	//
	// 9d7232b2-****-****-b9d9-7e17695779ab
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
	// a5281053-08e4-47a5-b2ab-5c0323de7b5a
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ScaleK8sApplicationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ScaleK8sApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *ScaleK8sApplicationResponseBody) GetChangeOrderId() *string {
	return s.ChangeOrderId
}

func (s *ScaleK8sApplicationResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ScaleK8sApplicationResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ScaleK8sApplicationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ScaleK8sApplicationResponseBody) SetChangeOrderId(v string) *ScaleK8sApplicationResponseBody {
	s.ChangeOrderId = &v
	return s
}

func (s *ScaleK8sApplicationResponseBody) SetCode(v int32) *ScaleK8sApplicationResponseBody {
	s.Code = &v
	return s
}

func (s *ScaleK8sApplicationResponseBody) SetMessage(v string) *ScaleK8sApplicationResponseBody {
	s.Message = &v
	return s
}

func (s *ScaleK8sApplicationResponseBody) SetRequestId(v string) *ScaleK8sApplicationResponseBody {
	s.RequestId = &v
	return s
}

func (s *ScaleK8sApplicationResponseBody) Validate() error {
	return dara.Validate(s)
}
