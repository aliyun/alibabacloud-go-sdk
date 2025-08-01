// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyLosslessRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ModifyLosslessRuleResponseBody
	GetCode() *int32
	SetData(v interface{}) *ModifyLosslessRuleResponseBody
	GetData() interface{}
	SetErrorCode(v string) *ModifyLosslessRuleResponseBody
	GetErrorCode() *string
	SetHttpStatusCode(v int32) *ModifyLosslessRuleResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ModifyLosslessRuleResponseBody
	GetMessage() *string
	SetRequestId(v string) *ModifyLosslessRuleResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModifyLosslessRuleResponseBody
	GetSuccess() *bool
}

type ModifyLosslessRuleResponseBody struct {
	// 响应码。
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// 返回数据。
	//
	// example:
	//
	// null
	Data interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	// The error code returned if the request failed.
	//
	// example:
	//
	// mse-100-000
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// HTTP状态码。
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The message returned.
	//
	// example:
	//
	// modifyLosslessRule success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 3369AD10-F1A6-4E6F-B99E-20F51826****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- `true`: The request was successful.
	//
	// 	- `false`: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyLosslessRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyLosslessRuleResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyLosslessRuleResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ModifyLosslessRuleResponseBody) GetData() interface{} {
	return s.Data
}

func (s *ModifyLosslessRuleResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ModifyLosslessRuleResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModifyLosslessRuleResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ModifyLosslessRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyLosslessRuleResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModifyLosslessRuleResponseBody) SetCode(v int32) *ModifyLosslessRuleResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyLosslessRuleResponseBody) SetData(v interface{}) *ModifyLosslessRuleResponseBody {
	s.Data = v
	return s
}

func (s *ModifyLosslessRuleResponseBody) SetErrorCode(v string) *ModifyLosslessRuleResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ModifyLosslessRuleResponseBody) SetHttpStatusCode(v int32) *ModifyLosslessRuleResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModifyLosslessRuleResponseBody) SetMessage(v string) *ModifyLosslessRuleResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyLosslessRuleResponseBody) SetRequestId(v string) *ModifyLosslessRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyLosslessRuleResponseBody) SetSuccess(v bool) *ModifyLosslessRuleResponseBody {
	s.Success = &v
	return s
}

func (s *ModifyLosslessRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
