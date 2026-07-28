// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAutopilotPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetAutopilotPolicyResponseBodyData) *GetAutopilotPolicyResponseBody
	GetData() *GetAutopilotPolicyResponseBodyData
	SetErrorCode(v string) *GetAutopilotPolicyResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetAutopilotPolicyResponseBody
	GetErrorMessage() *string
	SetHttpCode(v int32) *GetAutopilotPolicyResponseBody
	GetHttpCode() *int32
	SetRequestId(v string) *GetAutopilotPolicyResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetAutopilotPolicyResponseBody
	GetSuccess() *bool
}

type GetAutopilotPolicyResponseBody struct {
	// The Autopilot tuning policy response data.
	Data *GetAutopilotPolicyResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The business error code. This field is not empty when success is false. This field is empty when success is true.
	//
	// example:
	//
	// ""
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// The business error message. This field is not empty when success is false. This field is empty when success is true.
	//
	// example:
	//
	// ""
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// The business status code, which is always 200. Use the success field to determine whether the request was successful.
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
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetAutopilotPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAutopilotPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *GetAutopilotPolicyResponseBody) GetData() *GetAutopilotPolicyResponseBodyData {
	return s.Data
}

func (s *GetAutopilotPolicyResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetAutopilotPolicyResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetAutopilotPolicyResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *GetAutopilotPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAutopilotPolicyResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetAutopilotPolicyResponseBody) SetData(v *GetAutopilotPolicyResponseBodyData) *GetAutopilotPolicyResponseBody {
	s.Data = v
	return s
}

func (s *GetAutopilotPolicyResponseBody) SetErrorCode(v string) *GetAutopilotPolicyResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetAutopilotPolicyResponseBody) SetErrorMessage(v string) *GetAutopilotPolicyResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetAutopilotPolicyResponseBody) SetHttpCode(v int32) *GetAutopilotPolicyResponseBody {
	s.HttpCode = &v
	return s
}

func (s *GetAutopilotPolicyResponseBody) SetRequestId(v string) *GetAutopilotPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAutopilotPolicyResponseBody) SetSuccess(v bool) *GetAutopilotPolicyResponseBody {
	s.Success = &v
	return s
}

func (s *GetAutopilotPolicyResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAutopilotPolicyResponseBodyData struct {
	// Indicates whether automatic tuning is enabled. A value of true indicates that automatic tuning is in the ACTIVE state. A value of false indicates that tuning is not enabled.
	//
	// example:
	//
	// false
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// The tuning policy configuration.
	PolicyConfig *AutopilotPolicy `json:"policyConfig,omitempty" xml:"policyConfig,omitempty"`
}

func (s GetAutopilotPolicyResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetAutopilotPolicyResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAutopilotPolicyResponseBodyData) GetEnabled() *bool {
	return s.Enabled
}

func (s *GetAutopilotPolicyResponseBodyData) GetPolicyConfig() *AutopilotPolicy {
	return s.PolicyConfig
}

func (s *GetAutopilotPolicyResponseBodyData) SetEnabled(v bool) *GetAutopilotPolicyResponseBodyData {
	s.Enabled = &v
	return s
}

func (s *GetAutopilotPolicyResponseBodyData) SetPolicyConfig(v *AutopilotPolicy) *GetAutopilotPolicyResponseBodyData {
	s.PolicyConfig = v
	return s
}

func (s *GetAutopilotPolicyResponseBodyData) Validate() error {
	if s.PolicyConfig != nil {
		if err := s.PolicyConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}
