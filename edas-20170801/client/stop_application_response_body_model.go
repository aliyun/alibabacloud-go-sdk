// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopApplicationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetChangeOrderId(v string) *StopApplicationResponseBody
	GetChangeOrderId() *string
	SetCode(v int32) *StopApplicationResponseBody
	GetCode() *int32
	SetMessage(v string) *StopApplicationResponseBody
	GetMessage() *string
	SetRequestId(v string) *StopApplicationResponseBody
	GetRequestId() *string
}

type StopApplicationResponseBody struct {
	// The ID of the change process. You can call the GetChangeOrderInfo operation to query the details about the change process. For more information, see [GetChangeOrderInfo](https://help.aliyun.com/document_detail/62072.html).
	//
	// example:
	//
	// a9557bac-ddd7-*********************
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
	// D16979DC-4D42-****************
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopApplicationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StopApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *StopApplicationResponseBody) GetChangeOrderId() *string {
	return s.ChangeOrderId
}

func (s *StopApplicationResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *StopApplicationResponseBody) GetMessage() *string {
	return s.Message
}

func (s *StopApplicationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StopApplicationResponseBody) SetChangeOrderId(v string) *StopApplicationResponseBody {
	s.ChangeOrderId = &v
	return s
}

func (s *StopApplicationResponseBody) SetCode(v int32) *StopApplicationResponseBody {
	s.Code = &v
	return s
}

func (s *StopApplicationResponseBody) SetMessage(v string) *StopApplicationResponseBody {
	s.Message = &v
	return s
}

func (s *StopApplicationResponseBody) SetRequestId(v string) *StopApplicationResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopApplicationResponseBody) Validate() error {
	return dara.Validate(s)
}
