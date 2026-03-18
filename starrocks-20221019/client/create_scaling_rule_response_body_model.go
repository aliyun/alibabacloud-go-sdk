// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateScalingRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *CreateScalingRuleResponseBody
	GetAccessDeniedDetail() *string
	SetData(v bool) *CreateScalingRuleResponseBody
	GetData() *bool
	SetErrCode(v string) *CreateScalingRuleResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *CreateScalingRuleResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *CreateScalingRuleResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *CreateScalingRuleResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateScalingRuleResponseBody
	GetSuccess() *bool
}

type CreateScalingRuleResponseBody struct {
	// AccessDeniedDetail
	//
	// example:
	//
	// {     "PolicyType": "AccountLevelIdentityBasedPolicy",     "AuthPrincipalOwnerId": "xxx",     "EncodedDiagnosticMessage": "xxx",     "AuthPrincipalType": "xxx",     "AuthPrincipalDisplayName": "xxx",     "NoPermissionType": "ImplicitDeny",     "AuthAction": "sr:xxx"   }
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// InvalidParams
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// example:
	//
	// Invalid params: [instance not exists].
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// 32A44F0D-BFF6-5664-999A-218BBDE7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateScalingRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateScalingRuleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateScalingRuleResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *CreateScalingRuleResponseBody) GetData() *bool {
	return s.Data
}

func (s *CreateScalingRuleResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *CreateScalingRuleResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *CreateScalingRuleResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateScalingRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateScalingRuleResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateScalingRuleResponseBody) SetAccessDeniedDetail(v string) *CreateScalingRuleResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CreateScalingRuleResponseBody) SetData(v bool) *CreateScalingRuleResponseBody {
	s.Data = &v
	return s
}

func (s *CreateScalingRuleResponseBody) SetErrCode(v string) *CreateScalingRuleResponseBody {
	s.ErrCode = &v
	return s
}

func (s *CreateScalingRuleResponseBody) SetErrMessage(v string) *CreateScalingRuleResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *CreateScalingRuleResponseBody) SetHttpStatusCode(v int32) *CreateScalingRuleResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateScalingRuleResponseBody) SetRequestId(v string) *CreateScalingRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateScalingRuleResponseBody) SetSuccess(v bool) *CreateScalingRuleResponseBody {
	s.Success = &v
	return s
}

func (s *CreateScalingRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
