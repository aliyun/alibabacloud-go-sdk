// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iScaleInApplicationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetChangeOrderId(v string) *ScaleInApplicationResponseBody
	GetChangeOrderId() *string
	SetCode(v int32) *ScaleInApplicationResponseBody
	GetCode() *int32
	SetMessage(v string) *ScaleInApplicationResponseBody
	GetMessage() *string
}

type ScaleInApplicationResponseBody struct {
	// The ID of the change process for this operation. You can call the GetChangeOrderInfo operation to query the progress of this scale-in. For more information, see [GetChangeOrderInfo](https://help.aliyun.com/document_detail/62072.html). No ID is generated if the ForceStatus parameter is set to true. You can check whether the request is successful based on the value of the Code parameter.
	//
	// example:
	//
	// ddf5a4c7-a507-4a6e****************
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
}

func (s ScaleInApplicationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ScaleInApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *ScaleInApplicationResponseBody) GetChangeOrderId() *string {
	return s.ChangeOrderId
}

func (s *ScaleInApplicationResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ScaleInApplicationResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ScaleInApplicationResponseBody) SetChangeOrderId(v string) *ScaleInApplicationResponseBody {
	s.ChangeOrderId = &v
	return s
}

func (s *ScaleInApplicationResponseBody) SetCode(v int32) *ScaleInApplicationResponseBody {
	s.Code = &v
	return s
}

func (s *ScaleInApplicationResponseBody) SetMessage(v string) *ScaleInApplicationResponseBody {
	s.Message = &v
	return s
}

func (s *ScaleInApplicationResponseBody) Validate() error {
	return dara.Validate(s)
}
