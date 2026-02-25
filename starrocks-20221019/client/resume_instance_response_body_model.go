// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResumeInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ResumeInstanceResponseBody
	GetAccessDeniedDetail() *string
	SetData(v bool) *ResumeInstanceResponseBody
	GetData() *bool
	SetErrCode(v string) *ResumeInstanceResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ResumeInstanceResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *ResumeInstanceResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *ResumeInstanceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ResumeInstanceResponseBody
	GetSuccess() *bool
}

type ResumeInstanceResponseBody struct {
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
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ResumeInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ResumeInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *ResumeInstanceResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ResumeInstanceResponseBody) GetData() *bool {
	return s.Data
}

func (s *ResumeInstanceResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ResumeInstanceResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ResumeInstanceResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ResumeInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ResumeInstanceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ResumeInstanceResponseBody) SetAccessDeniedDetail(v string) *ResumeInstanceResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ResumeInstanceResponseBody) SetData(v bool) *ResumeInstanceResponseBody {
	s.Data = &v
	return s
}

func (s *ResumeInstanceResponseBody) SetErrCode(v string) *ResumeInstanceResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ResumeInstanceResponseBody) SetErrMessage(v string) *ResumeInstanceResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ResumeInstanceResponseBody) SetHttpStatusCode(v int32) *ResumeInstanceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ResumeInstanceResponseBody) SetRequestId(v string) *ResumeInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ResumeInstanceResponseBody) SetSuccess(v bool) *ResumeInstanceResponseBody {
	s.Success = &v
	return s
}

func (s *ResumeInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
