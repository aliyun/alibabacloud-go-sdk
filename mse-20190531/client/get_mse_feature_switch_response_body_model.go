// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMseFeatureSwitchResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *GetMseFeatureSwitchResponseBody
	GetErrorCode() *string
	SetMessage(v string) *GetMseFeatureSwitchResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetMseFeatureSwitchResponseBody
	GetRequestId() *string
	SetResult(v map[string]interface{}) *GetMseFeatureSwitchResponseBody
	GetResult() map[string]interface{}
	SetSuccess(v bool) *GetMseFeatureSwitchResponseBody
	GetSuccess() *bool
}

type GetMseFeatureSwitchResponseBody struct {
	// The error code returned if the request failed.
	//
	// example:
	//
	// NoPermission
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The message returned.
	//
	// example:
	//
	// The request is successfully processed.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 6B4653A3-8D9C-5FDC-BB0C-936D40E9794B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The returned result.
	//
	// example:
	//
	// {\\"nacos_config_encrypt\\": False}
	Result map[string]interface{} `json:"Result,omitempty" xml:"Result,omitempty"`
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

func (s GetMseFeatureSwitchResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetMseFeatureSwitchResponseBody) GoString() string {
	return s.String()
}

func (s *GetMseFeatureSwitchResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetMseFeatureSwitchResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetMseFeatureSwitchResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetMseFeatureSwitchResponseBody) GetResult() map[string]interface{} {
	return s.Result
}

func (s *GetMseFeatureSwitchResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetMseFeatureSwitchResponseBody) SetErrorCode(v string) *GetMseFeatureSwitchResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetMseFeatureSwitchResponseBody) SetMessage(v string) *GetMseFeatureSwitchResponseBody {
	s.Message = &v
	return s
}

func (s *GetMseFeatureSwitchResponseBody) SetRequestId(v string) *GetMseFeatureSwitchResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMseFeatureSwitchResponseBody) SetResult(v map[string]interface{}) *GetMseFeatureSwitchResponseBody {
	s.Result = v
	return s
}

func (s *GetMseFeatureSwitchResponseBody) SetSuccess(v bool) *GetMseFeatureSwitchResponseBody {
	s.Success = &v
	return s
}

func (s *GetMseFeatureSwitchResponseBody) Validate() error {
	return dara.Validate(s)
}
