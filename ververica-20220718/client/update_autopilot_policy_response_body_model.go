// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAutopilotPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *UpdateAutopilotPolicyResponseBodyData) *UpdateAutopilotPolicyResponseBody
	GetData() *UpdateAutopilotPolicyResponseBodyData
	SetErrorCode(v string) *UpdateAutopilotPolicyResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *UpdateAutopilotPolicyResponseBody
	GetErrorMessage() *string
	SetHttpCode(v int32) *UpdateAutopilotPolicyResponseBody
	GetHttpCode() *int32
	SetRequestId(v string) *UpdateAutopilotPolicyResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateAutopilotPolicyResponseBody
	GetSuccess() *bool
}

type UpdateAutopilotPolicyResponseBody struct {
	// The Autopilot tuning policy response data.
	Data *UpdateAutopilotPolicyResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The error code. This parameter is not empty when success is false, indicating a business error code. This parameter is empty when success is true.
	//
	// example:
	//
	// ""
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// The error message. This parameter is not empty when success is false, indicating a business error message. This parameter is empty when success is true.
	//
	// example:
	//
	// ""
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// The HTTP status code, which is always 200. Use the success parameter to determine whether the request was successful.
	//
	// example:
	//
	// 200
	HttpCode *int32 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// CBC799F0-AS7S-1D30-8A4F-882ED4DD****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// True
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdateAutopilotPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateAutopilotPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAutopilotPolicyResponseBody) GetData() *UpdateAutopilotPolicyResponseBodyData {
	return s.Data
}

func (s *UpdateAutopilotPolicyResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *UpdateAutopilotPolicyResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *UpdateAutopilotPolicyResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *UpdateAutopilotPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateAutopilotPolicyResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateAutopilotPolicyResponseBody) SetData(v *UpdateAutopilotPolicyResponseBodyData) *UpdateAutopilotPolicyResponseBody {
	s.Data = v
	return s
}

func (s *UpdateAutopilotPolicyResponseBody) SetErrorCode(v string) *UpdateAutopilotPolicyResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateAutopilotPolicyResponseBody) SetErrorMessage(v string) *UpdateAutopilotPolicyResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateAutopilotPolicyResponseBody) SetHttpCode(v int32) *UpdateAutopilotPolicyResponseBody {
	s.HttpCode = &v
	return s
}

func (s *UpdateAutopilotPolicyResponseBody) SetRequestId(v string) *UpdateAutopilotPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAutopilotPolicyResponseBody) SetSuccess(v bool) *UpdateAutopilotPolicyResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateAutopilotPolicyResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateAutopilotPolicyResponseBodyData struct {
	// Indicates whether automatic tuning is enabled. A value of true indicates that automatic tuning is active (ACTIVE), and a value of false indicates that tuning is not enabled.
	//
	// example:
	//
	// true
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// The tuning policy configuration. This parameter uses full PUT mode: when specified, the complete policy object replaces the existing configuration entirely (fields not included are cleared). If this parameter is not specified, the existing configuration is retained.
	PolicyConfig *AutopilotPolicy `json:"policyConfig,omitempty" xml:"policyConfig,omitempty"`
}

func (s UpdateAutopilotPolicyResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UpdateAutopilotPolicyResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateAutopilotPolicyResponseBodyData) GetEnabled() *bool {
	return s.Enabled
}

func (s *UpdateAutopilotPolicyResponseBodyData) GetPolicyConfig() *AutopilotPolicy {
	return s.PolicyConfig
}

func (s *UpdateAutopilotPolicyResponseBodyData) SetEnabled(v bool) *UpdateAutopilotPolicyResponseBodyData {
	s.Enabled = &v
	return s
}

func (s *UpdateAutopilotPolicyResponseBodyData) SetPolicyConfig(v *AutopilotPolicy) *UpdateAutopilotPolicyResponseBodyData {
	s.PolicyConfig = v
	return s
}

func (s *UpdateAutopilotPolicyResponseBodyData) Validate() error {
	if s.PolicyConfig != nil {
		if err := s.PolicyConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}
