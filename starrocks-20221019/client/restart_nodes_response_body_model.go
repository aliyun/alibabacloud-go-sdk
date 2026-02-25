// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestartNodesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *RestartNodesResponseBody
	GetAccessDeniedDetail() *string
	SetData(v bool) *RestartNodesResponseBody
	GetData() *bool
	SetErrCode(v string) *RestartNodesResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *RestartNodesResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *RestartNodesResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *RestartNodesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *RestartNodesResponseBody
	GetSuccess() *bool
}

type RestartNodesResponseBody struct {
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
	// Success
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
	// 32A44F0D-BFF6-5664-999A-218BBDE74XXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RestartNodesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RestartNodesResponseBody) GoString() string {
	return s.String()
}

func (s *RestartNodesResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *RestartNodesResponseBody) GetData() *bool {
	return s.Data
}

func (s *RestartNodesResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *RestartNodesResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *RestartNodesResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *RestartNodesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RestartNodesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *RestartNodesResponseBody) SetAccessDeniedDetail(v string) *RestartNodesResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *RestartNodesResponseBody) SetData(v bool) *RestartNodesResponseBody {
	s.Data = &v
	return s
}

func (s *RestartNodesResponseBody) SetErrCode(v string) *RestartNodesResponseBody {
	s.ErrCode = &v
	return s
}

func (s *RestartNodesResponseBody) SetErrMessage(v string) *RestartNodesResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *RestartNodesResponseBody) SetHttpStatusCode(v int32) *RestartNodesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *RestartNodesResponseBody) SetRequestId(v string) *RestartNodesResponseBody {
	s.RequestId = &v
	return s
}

func (s *RestartNodesResponseBody) SetSuccess(v bool) *RestartNodesResponseBody {
	s.Success = &v
	return s
}

func (s *RestartNodesResponseBody) Validate() error {
	return dara.Validate(s)
}
