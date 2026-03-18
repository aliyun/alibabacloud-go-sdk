// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyScalingRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ModifyScalingRuleResponseBody
	GetAccessDeniedDetail() *string
	SetData(v bool) *ModifyScalingRuleResponseBody
	GetData() *bool
	SetErrCode(v string) *ModifyScalingRuleResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ModifyScalingRuleResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *ModifyScalingRuleResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *ModifyScalingRuleResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModifyScalingRuleResponseBody
	GetSuccess() *bool
}

type ModifyScalingRuleResponseBody struct {
	// AccessDeniedDetail
	//
	// example:
	//
	// {     "PolicyType": "AccountLevelIdentityBasedPolicy",     "AuthPrincipalOwnerId": "xxx",     "EncodedDiagnosticMessage": "xxx",     "AuthPrincipalType": "xxx",     "AuthPrincipalDisplayName": "xxx",     "NoPermissionType": "ImplicitDeny",     "AuthAction": "sr:xxx"   }
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// 24151320976****
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// InvalidParams
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// example:
	//
	// Invalid params: [Region id should be select from set [cn-beijing, cn-hangzhou]]
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
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyScalingRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyScalingRuleResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyScalingRuleResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ModifyScalingRuleResponseBody) GetData() *bool {
	return s.Data
}

func (s *ModifyScalingRuleResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ModifyScalingRuleResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ModifyScalingRuleResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModifyScalingRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyScalingRuleResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModifyScalingRuleResponseBody) SetAccessDeniedDetail(v string) *ModifyScalingRuleResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ModifyScalingRuleResponseBody) SetData(v bool) *ModifyScalingRuleResponseBody {
	s.Data = &v
	return s
}

func (s *ModifyScalingRuleResponseBody) SetErrCode(v string) *ModifyScalingRuleResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ModifyScalingRuleResponseBody) SetErrMessage(v string) *ModifyScalingRuleResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ModifyScalingRuleResponseBody) SetHttpStatusCode(v int32) *ModifyScalingRuleResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModifyScalingRuleResponseBody) SetRequestId(v string) *ModifyScalingRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyScalingRuleResponseBody) SetSuccess(v bool) *ModifyScalingRuleResponseBody {
	s.Success = &v
	return s
}

func (s *ModifyScalingRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
