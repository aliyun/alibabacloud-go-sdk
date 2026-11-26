// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitQualityCheckTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SubmitQualityCheckTaskResponseBody
	GetCode() *string
	SetData(v string) *SubmitQualityCheckTaskResponseBody
	GetData() *string
	SetMessage(v string) *SubmitQualityCheckTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *SubmitQualityCheckTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SubmitQualityCheckTaskResponseBody
	GetSuccess() *bool
}

type SubmitQualityCheckTaskResponseBody struct {
	// The result code. A value of **200*	- indicates success, and other values indicate failure. You can determine the cause of the failure based on this field.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The ID of the task.
	//
	// example:
	//
	// F6C2B68F-2311-4495-82AC-DAE86C9****
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The error details when an error occurs, or successful when the request is successful.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 00A044A2-D59B-4104-96BA-84060AE8345F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful. You can determine whether the request is successful based on this field. A value of true indicates success, and a value of false or null indicates failure.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SubmitQualityCheckTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitQualityCheckTaskResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitQualityCheckTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *SubmitQualityCheckTaskResponseBody) GetData() *string {
	return s.Data
}

func (s *SubmitQualityCheckTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SubmitQualityCheckTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitQualityCheckTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SubmitQualityCheckTaskResponseBody) SetCode(v string) *SubmitQualityCheckTaskResponseBody {
	s.Code = &v
	return s
}

func (s *SubmitQualityCheckTaskResponseBody) SetData(v string) *SubmitQualityCheckTaskResponseBody {
	s.Data = &v
	return s
}

func (s *SubmitQualityCheckTaskResponseBody) SetMessage(v string) *SubmitQualityCheckTaskResponseBody {
	s.Message = &v
	return s
}

func (s *SubmitQualityCheckTaskResponseBody) SetRequestId(v string) *SubmitQualityCheckTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitQualityCheckTaskResponseBody) SetSuccess(v bool) *SubmitQualityCheckTaskResponseBody {
	s.Success = &v
	return s
}

func (s *SubmitQualityCheckTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
