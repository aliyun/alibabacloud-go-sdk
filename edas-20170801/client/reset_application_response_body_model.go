// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetApplicationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetChangeOrderId(v string) *ResetApplicationResponseBody
	GetChangeOrderId() *string
	SetCode(v int32) *ResetApplicationResponseBody
	GetCode() *int32
	SetMessage(v string) *ResetApplicationResponseBody
	GetMessage() *string
	SetRequestId(v string) *ResetApplicationResponseBody
	GetRequestId() *string
}

type ResetApplicationResponseBody struct {
	// The ID of the change process for this operation. You can call the GetChangeOrderInfo operation to query the progress of this operation. For more information, see [GetChangeOrderInfo](https://help.aliyun.com/document_detail/62072.html).
	//
	// example:
	//
	// 1c66548e-a082-40************
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
	// D16979DC-4D42-*********
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ResetApplicationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ResetApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *ResetApplicationResponseBody) GetChangeOrderId() *string {
	return s.ChangeOrderId
}

func (s *ResetApplicationResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ResetApplicationResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ResetApplicationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ResetApplicationResponseBody) SetChangeOrderId(v string) *ResetApplicationResponseBody {
	s.ChangeOrderId = &v
	return s
}

func (s *ResetApplicationResponseBody) SetCode(v int32) *ResetApplicationResponseBody {
	s.Code = &v
	return s
}

func (s *ResetApplicationResponseBody) SetMessage(v string) *ResetApplicationResponseBody {
	s.Message = &v
	return s
}

func (s *ResetApplicationResponseBody) SetRequestId(v string) *ResetApplicationResponseBody {
	s.RequestId = &v
	return s
}

func (s *ResetApplicationResponseBody) Validate() error {
	return dara.Validate(s)
}
