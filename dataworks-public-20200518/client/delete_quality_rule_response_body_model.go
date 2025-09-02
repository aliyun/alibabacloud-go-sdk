// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteQualityRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *DeleteQualityRuleResponseBody
	GetData() *bool
	SetErrorCode(v string) *DeleteQualityRuleResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *DeleteQualityRuleResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v int32) *DeleteQualityRuleResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *DeleteQualityRuleResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteQualityRuleResponseBody
	GetSuccess() *bool
}

type DeleteQualityRuleResponseBody struct {
	// Indicates whether the monitoring rule was deleted.
	//
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
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
	// 6d739ef6-098a-47****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteQualityRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteQualityRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteQualityRuleResponseBody) GetData() *bool {
	return s.Data
}

func (s *DeleteQualityRuleResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DeleteQualityRuleResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DeleteQualityRuleResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteQualityRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteQualityRuleResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteQualityRuleResponseBody) SetData(v bool) *DeleteQualityRuleResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteQualityRuleResponseBody) SetErrorCode(v string) *DeleteQualityRuleResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteQualityRuleResponseBody) SetErrorMessage(v string) *DeleteQualityRuleResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteQualityRuleResponseBody) SetHttpStatusCode(v int32) *DeleteQualityRuleResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteQualityRuleResponseBody) SetRequestId(v string) *DeleteQualityRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteQualityRuleResponseBody) SetSuccess(v bool) *DeleteQualityRuleResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteQualityRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
