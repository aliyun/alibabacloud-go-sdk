// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateQualityRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *CreateQualityRuleResponseBody
	GetData() *string
	SetErrorCode(v string) *CreateQualityRuleResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *CreateQualityRuleResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v int32) *CreateQualityRuleResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *CreateQualityRuleResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateQualityRuleResponseBody
	GetSuccess() *bool
}

type CreateQualityRuleResponseBody struct {
	// The ID of the monitoring rule that you created.
	//
	// example:
	//
	// 12345
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The error code.
	//
	// example:
	//
	// 401
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
	// The request ID.
	//
	// example:
	//
	// ecb967ec-c137-48****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateQualityRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateQualityRuleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateQualityRuleResponseBody) GetData() *string {
	return s.Data
}

func (s *CreateQualityRuleResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CreateQualityRuleResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *CreateQualityRuleResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateQualityRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateQualityRuleResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateQualityRuleResponseBody) SetData(v string) *CreateQualityRuleResponseBody {
	s.Data = &v
	return s
}

func (s *CreateQualityRuleResponseBody) SetErrorCode(v string) *CreateQualityRuleResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateQualityRuleResponseBody) SetErrorMessage(v string) *CreateQualityRuleResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateQualityRuleResponseBody) SetHttpStatusCode(v int32) *CreateQualityRuleResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateQualityRuleResponseBody) SetRequestId(v string) *CreateQualityRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateQualityRuleResponseBody) SetSuccess(v bool) *CreateQualityRuleResponseBody {
	s.Success = &v
	return s
}

func (s *CreateQualityRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
