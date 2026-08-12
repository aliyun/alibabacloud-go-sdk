// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitApplyRecordResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SubmitApplyRecordResponseBody
	GetCode() *string
	SetData(v bool) *SubmitApplyRecordResponseBody
	GetData() *bool
	SetMessage(v string) *SubmitApplyRecordResponseBody
	GetMessage() *string
	SetRequestId(v string) *SubmitApplyRecordResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SubmitApplyRecordResponseBody
	GetSuccess() *bool
}

type SubmitApplyRecordResponseBody struct {
	// The status code.
	//
	// - **200**: Succeeded.
	//
	// - **Other (400, 500)**: Failed.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The submission result.
	//
	// - **true**: Submitted.
	//
	// - **false**: Submission failed.
	//
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The message content.
	//
	// example:
	//
	// successful‌
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 6B57D35D-9DAC-5393-AE39-07697E37C2E7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The call status.
	//
	// - **true**: Succeeded.
	//
	// - **false**: Failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SubmitApplyRecordResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitApplyRecordResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitApplyRecordResponseBody) GetCode() *string {
	return s.Code
}

func (s *SubmitApplyRecordResponseBody) GetData() *bool {
	return s.Data
}

func (s *SubmitApplyRecordResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SubmitApplyRecordResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitApplyRecordResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SubmitApplyRecordResponseBody) SetCode(v string) *SubmitApplyRecordResponseBody {
	s.Code = &v
	return s
}

func (s *SubmitApplyRecordResponseBody) SetData(v bool) *SubmitApplyRecordResponseBody {
	s.Data = &v
	return s
}

func (s *SubmitApplyRecordResponseBody) SetMessage(v string) *SubmitApplyRecordResponseBody {
	s.Message = &v
	return s
}

func (s *SubmitApplyRecordResponseBody) SetRequestId(v string) *SubmitApplyRecordResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitApplyRecordResponseBody) SetSuccess(v bool) *SubmitApplyRecordResponseBody {
	s.Success = &v
	return s
}

func (s *SubmitApplyRecordResponseBody) Validate() error {
	return dara.Validate(s)
}
