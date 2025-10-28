// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartApplicationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetChangeOrderId(v string) *StartApplicationResponseBody
	GetChangeOrderId() *string
	SetCode(v int32) *StartApplicationResponseBody
	GetCode() *int32
	SetMessage(v string) *StartApplicationResponseBody
	GetMessage() *string
	SetRequestId(v string) *StartApplicationResponseBody
	GetRequestId() *string
}

type StartApplicationResponseBody struct {
	// The ID of the change process for this operation. You can call the GetChangeOrderInfo operation to query the progress of this startup. For more information, see [GetChangeOrderInfo](https://help.aliyun.com/document_detail/62072.html).
	//
	// example:
	//
	// 426d3328-11a***************
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
	// D16979DC-4D42-**************
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StartApplicationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StartApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *StartApplicationResponseBody) GetChangeOrderId() *string {
	return s.ChangeOrderId
}

func (s *StartApplicationResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *StartApplicationResponseBody) GetMessage() *string {
	return s.Message
}

func (s *StartApplicationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StartApplicationResponseBody) SetChangeOrderId(v string) *StartApplicationResponseBody {
	s.ChangeOrderId = &v
	return s
}

func (s *StartApplicationResponseBody) SetCode(v int32) *StartApplicationResponseBody {
	s.Code = &v
	return s
}

func (s *StartApplicationResponseBody) SetMessage(v string) *StartApplicationResponseBody {
	s.Message = &v
	return s
}

func (s *StartApplicationResponseBody) SetRequestId(v string) *StartApplicationResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartApplicationResponseBody) Validate() error {
	return dara.Validate(s)
}
