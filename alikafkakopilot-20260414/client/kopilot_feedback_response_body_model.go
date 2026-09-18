// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iKopilotFeedbackResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *KopilotFeedbackResponseBody
	GetCode() *int64
	SetData(v map[string]*int64) *KopilotFeedbackResponseBody
	GetData() map[string]*int64
	SetRequestId(v string) *KopilotFeedbackResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *KopilotFeedbackResponseBody
	GetSuccess() *bool
}

type KopilotFeedbackResponseBody struct {
	// The return code. A value of 200 indicates success.
	//
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned when the call is successful.
	Data map[string]*int64 `json:"Data,omitempty" xml:"Data,omitempty"`
	// The request ID.
	//
	// example:
	//
	// E522C7AB-AA99-53B1-BEE7-8BA83ED678F7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s KopilotFeedbackResponseBody) String() string {
	return dara.Prettify(s)
}

func (s KopilotFeedbackResponseBody) GoString() string {
	return s.String()
}

func (s *KopilotFeedbackResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *KopilotFeedbackResponseBody) GetData() map[string]*int64 {
	return s.Data
}

func (s *KopilotFeedbackResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *KopilotFeedbackResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *KopilotFeedbackResponseBody) SetCode(v int64) *KopilotFeedbackResponseBody {
	s.Code = &v
	return s
}

func (s *KopilotFeedbackResponseBody) SetData(v map[string]*int64) *KopilotFeedbackResponseBody {
	s.Data = v
	return s
}

func (s *KopilotFeedbackResponseBody) SetRequestId(v string) *KopilotFeedbackResponseBody {
	s.RequestId = &v
	return s
}

func (s *KopilotFeedbackResponseBody) SetSuccess(v bool) *KopilotFeedbackResponseBody {
	s.Success = &v
	return s
}

func (s *KopilotFeedbackResponseBody) Validate() error {
	return dara.Validate(s)
}
