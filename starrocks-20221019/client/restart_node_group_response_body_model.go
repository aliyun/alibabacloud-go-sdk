// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestartNodeGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *RestartNodeGroupResponseBody
	GetAccessDeniedDetail() *string
	SetData(v bool) *RestartNodeGroupResponseBody
	GetData() *bool
	SetErrCode(v string) *RestartNodeGroupResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *RestartNodeGroupResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *RestartNodeGroupResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *RestartNodeGroupResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *RestartNodeGroupResponseBody
	GetSuccess() *bool
}

type RestartNodeGroupResponseBody struct {
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
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RestartNodeGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RestartNodeGroupResponseBody) GoString() string {
	return s.String()
}

func (s *RestartNodeGroupResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *RestartNodeGroupResponseBody) GetData() *bool {
	return s.Data
}

func (s *RestartNodeGroupResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *RestartNodeGroupResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *RestartNodeGroupResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *RestartNodeGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RestartNodeGroupResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *RestartNodeGroupResponseBody) SetAccessDeniedDetail(v string) *RestartNodeGroupResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *RestartNodeGroupResponseBody) SetData(v bool) *RestartNodeGroupResponseBody {
	s.Data = &v
	return s
}

func (s *RestartNodeGroupResponseBody) SetErrCode(v string) *RestartNodeGroupResponseBody {
	s.ErrCode = &v
	return s
}

func (s *RestartNodeGroupResponseBody) SetErrMessage(v string) *RestartNodeGroupResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *RestartNodeGroupResponseBody) SetHttpStatusCode(v int32) *RestartNodeGroupResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *RestartNodeGroupResponseBody) SetRequestId(v string) *RestartNodeGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *RestartNodeGroupResponseBody) SetSuccess(v bool) *RestartNodeGroupResponseBody {
	s.Success = &v
	return s
}

func (s *RestartNodeGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
