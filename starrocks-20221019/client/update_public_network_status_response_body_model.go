// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePublicNetworkStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *UpdatePublicNetworkStatusResponseBody
	GetAccessDeniedDetail() *string
	SetData(v bool) *UpdatePublicNetworkStatusResponseBody
	GetData() *bool
	SetErrCode(v string) *UpdatePublicNetworkStatusResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *UpdatePublicNetworkStatusResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *UpdatePublicNetworkStatusResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *UpdatePublicNetworkStatusResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdatePublicNetworkStatusResponseBody
	GetSuccess() *bool
	SetTotal(v int32) *UpdatePublicNetworkStatusResponseBody
	GetTotal() *int32
}

type UpdatePublicNetworkStatusResponseBody struct {
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
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 100
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s UpdatePublicNetworkStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdatePublicNetworkStatusResponseBody) GoString() string {
	return s.String()
}

func (s *UpdatePublicNetworkStatusResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *UpdatePublicNetworkStatusResponseBody) GetData() *bool {
	return s.Data
}

func (s *UpdatePublicNetworkStatusResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *UpdatePublicNetworkStatusResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *UpdatePublicNetworkStatusResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdatePublicNetworkStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdatePublicNetworkStatusResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdatePublicNetworkStatusResponseBody) GetTotal() *int32 {
	return s.Total
}

func (s *UpdatePublicNetworkStatusResponseBody) SetAccessDeniedDetail(v string) *UpdatePublicNetworkStatusResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *UpdatePublicNetworkStatusResponseBody) SetData(v bool) *UpdatePublicNetworkStatusResponseBody {
	s.Data = &v
	return s
}

func (s *UpdatePublicNetworkStatusResponseBody) SetErrCode(v string) *UpdatePublicNetworkStatusResponseBody {
	s.ErrCode = &v
	return s
}

func (s *UpdatePublicNetworkStatusResponseBody) SetErrMessage(v string) *UpdatePublicNetworkStatusResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *UpdatePublicNetworkStatusResponseBody) SetHttpStatusCode(v int32) *UpdatePublicNetworkStatusResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdatePublicNetworkStatusResponseBody) SetRequestId(v string) *UpdatePublicNetworkStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdatePublicNetworkStatusResponseBody) SetSuccess(v bool) *UpdatePublicNetworkStatusResponseBody {
	s.Success = &v
	return s
}

func (s *UpdatePublicNetworkStatusResponseBody) SetTotal(v int32) *UpdatePublicNetworkStatusResponseBody {
	s.Total = &v
	return s
}

func (s *UpdatePublicNetworkStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
