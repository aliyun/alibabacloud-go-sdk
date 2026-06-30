// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAnalyzeLabelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *AnalyzeLabelResponseBody
	GetCode() *string
	SetData(v string) *AnalyzeLabelResponseBody
	GetData() *string
	SetMessage(v string) *AnalyzeLabelResponseBody
	GetMessage() *string
	SetRequestId(v string) *AnalyzeLabelResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AnalyzeLabelResponseBody
	GetSuccess() *bool
}

type AnalyzeLabelResponseBody struct {
	// The result code. A value of **200*	- indicates success. Other values indicate failure. You can use this field to determine the cause of the failure.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The analysis task ID.
	//
	// example:
	//
	// 20260629-DCC646E7-BE7F-114E-9F32-0C928292FC7F
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The error message, if any.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 106C6CA0-282D-4AF7-85F0-D2D24***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful. true: The call was successful. false: The call failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AnalyzeLabelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AnalyzeLabelResponseBody) GoString() string {
	return s.String()
}

func (s *AnalyzeLabelResponseBody) GetCode() *string {
	return s.Code
}

func (s *AnalyzeLabelResponseBody) GetData() *string {
	return s.Data
}

func (s *AnalyzeLabelResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AnalyzeLabelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AnalyzeLabelResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AnalyzeLabelResponseBody) SetCode(v string) *AnalyzeLabelResponseBody {
	s.Code = &v
	return s
}

func (s *AnalyzeLabelResponseBody) SetData(v string) *AnalyzeLabelResponseBody {
	s.Data = &v
	return s
}

func (s *AnalyzeLabelResponseBody) SetMessage(v string) *AnalyzeLabelResponseBody {
	s.Message = &v
	return s
}

func (s *AnalyzeLabelResponseBody) SetRequestId(v string) *AnalyzeLabelResponseBody {
	s.RequestId = &v
	return s
}

func (s *AnalyzeLabelResponseBody) SetSuccess(v bool) *AnalyzeLabelResponseBody {
	s.Success = &v
	return s
}

func (s *AnalyzeLabelResponseBody) Validate() error {
	return dara.Validate(s)
}
