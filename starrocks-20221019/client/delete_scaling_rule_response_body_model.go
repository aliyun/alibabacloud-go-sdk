// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteScalingRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *DeleteScalingRuleResponseBody
	GetAccessDeniedDetail() *string
	SetData(v bool) *DeleteScalingRuleResponseBody
	GetData() *bool
	SetErrCode(v string) *DeleteScalingRuleResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *DeleteScalingRuleResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *DeleteScalingRuleResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *DeleteScalingRuleResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteScalingRuleResponseBody
	GetSuccess() *bool
}

type DeleteScalingRuleResponseBody struct {
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
	// null
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// 32A44F0D-BFF6-5664-999A-218BBDE74XXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteScalingRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteScalingRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteScalingRuleResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *DeleteScalingRuleResponseBody) GetData() *bool {
	return s.Data
}

func (s *DeleteScalingRuleResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *DeleteScalingRuleResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *DeleteScalingRuleResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteScalingRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteScalingRuleResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteScalingRuleResponseBody) SetAccessDeniedDetail(v string) *DeleteScalingRuleResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *DeleteScalingRuleResponseBody) SetData(v bool) *DeleteScalingRuleResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteScalingRuleResponseBody) SetErrCode(v string) *DeleteScalingRuleResponseBody {
	s.ErrCode = &v
	return s
}

func (s *DeleteScalingRuleResponseBody) SetErrMessage(v string) *DeleteScalingRuleResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *DeleteScalingRuleResponseBody) SetHttpStatusCode(v int32) *DeleteScalingRuleResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteScalingRuleResponseBody) SetRequestId(v string) *DeleteScalingRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteScalingRuleResponseBody) SetSuccess(v bool) *DeleteScalingRuleResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteScalingRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
