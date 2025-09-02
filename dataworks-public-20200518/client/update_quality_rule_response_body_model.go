// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateQualityRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *UpdateQualityRuleResponseBody
	GetData() *bool
	SetErrorCode(v string) *UpdateQualityRuleResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *UpdateQualityRuleResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v int32) *UpdateQualityRuleResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *UpdateQualityRuleResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateQualityRuleResponseBody
	GetSuccess() *bool
}

type UpdateQualityRuleResponseBody struct {
	// Indicates whether the monitoring rule is updated.
	//
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The error code.
	//
	// example:
	//
	// 0
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// You have no permission.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The request ID. You can troubleshoot issues based on the ID.
	//
	// example:
	//
	// 576b9457-2cf5-4****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateQualityRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateQualityRuleResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateQualityRuleResponseBody) GetData() *bool {
	return s.Data
}

func (s *UpdateQualityRuleResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *UpdateQualityRuleResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *UpdateQualityRuleResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateQualityRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateQualityRuleResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateQualityRuleResponseBody) SetData(v bool) *UpdateQualityRuleResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateQualityRuleResponseBody) SetErrorCode(v string) *UpdateQualityRuleResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateQualityRuleResponseBody) SetErrorMessage(v string) *UpdateQualityRuleResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateQualityRuleResponseBody) SetHttpStatusCode(v int32) *UpdateQualityRuleResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateQualityRuleResponseBody) SetRequestId(v string) *UpdateQualityRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateQualityRuleResponseBody) SetSuccess(v bool) *UpdateQualityRuleResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateQualityRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
