// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iScaleOutApplicationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetChangeOrderId(v string) *ScaleOutApplicationResponseBody
	GetChangeOrderId() *string
	SetCode(v int32) *ScaleOutApplicationResponseBody
	GetCode() *int32
	SetMessage(v string) *ScaleOutApplicationResponseBody
	GetMessage() *string
	SetRequestId(v string) *ScaleOutApplicationResponseBody
	GetRequestId() *string
}

type ScaleOutApplicationResponseBody struct {
	// The ID of the change process. You can call the GetChangeOrderInfo operation to query the progress of this scale-out. For more information, see [GetChangeOrderInfo](https://help.aliyun.com/document_detail/62072.html).
	//
	// example:
	//
	// f4208118-7171-4e20-92************
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
	// D16979DC-4D42-***********************
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ScaleOutApplicationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ScaleOutApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *ScaleOutApplicationResponseBody) GetChangeOrderId() *string {
	return s.ChangeOrderId
}

func (s *ScaleOutApplicationResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ScaleOutApplicationResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ScaleOutApplicationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ScaleOutApplicationResponseBody) SetChangeOrderId(v string) *ScaleOutApplicationResponseBody {
	s.ChangeOrderId = &v
	return s
}

func (s *ScaleOutApplicationResponseBody) SetCode(v int32) *ScaleOutApplicationResponseBody {
	s.Code = &v
	return s
}

func (s *ScaleOutApplicationResponseBody) SetMessage(v string) *ScaleOutApplicationResponseBody {
	s.Message = &v
	return s
}

func (s *ScaleOutApplicationResponseBody) SetRequestId(v string) *ScaleOutApplicationResponseBody {
	s.RequestId = &v
	return s
}

func (s *ScaleOutApplicationResponseBody) Validate() error {
	return dara.Validate(s)
}
