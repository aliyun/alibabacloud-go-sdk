// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddRecordingRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *AddRecordingRuleResponseBody
	GetCode() *int32
	SetData(v string) *AddRecordingRuleResponseBody
	GetData() *string
	SetMessage(v string) *AddRecordingRuleResponseBody
	GetMessage() *string
	SetRequestId(v string) *AddRecordingRuleResponseBody
	GetRequestId() *string
}

type AddRecordingRuleResponseBody struct {
	// The status code. The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The status of the response.
	//
	// example:
	//
	// success
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9FEA6D00-317F-45E3-9004-7FB8B0B7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddRecordingRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddRecordingRuleResponseBody) GoString() string {
	return s.String()
}

func (s *AddRecordingRuleResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *AddRecordingRuleResponseBody) GetData() *string {
	return s.Data
}

func (s *AddRecordingRuleResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AddRecordingRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddRecordingRuleResponseBody) SetCode(v int32) *AddRecordingRuleResponseBody {
	s.Code = &v
	return s
}

func (s *AddRecordingRuleResponseBody) SetData(v string) *AddRecordingRuleResponseBody {
	s.Data = &v
	return s
}

func (s *AddRecordingRuleResponseBody) SetMessage(v string) *AddRecordingRuleResponseBody {
	s.Message = &v
	return s
}

func (s *AddRecordingRuleResponseBody) SetRequestId(v string) *AddRecordingRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddRecordingRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
