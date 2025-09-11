// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySmsAuthorizationLetterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *QuerySmsAuthorizationLetterResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *QuerySmsAuthorizationLetterResponseBody
	GetCode() *string
	SetData(v []*QuerySmsAuthorizationLetterResponseBodyData) *QuerySmsAuthorizationLetterResponseBody
	GetData() []*QuerySmsAuthorizationLetterResponseBodyData
	SetMessage(v string) *QuerySmsAuthorizationLetterResponseBody
	GetMessage() *string
	SetRequestId(v string) *QuerySmsAuthorizationLetterResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QuerySmsAuthorizationLetterResponseBody
	GetSuccess() *bool
}

type QuerySmsAuthorizationLetterResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string                                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Data []*QuerySmsAuthorizationLetterResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 25D5AFDE-8EBC-132E-8909-1FDC071D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QuerySmsAuthorizationLetterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QuerySmsAuthorizationLetterResponseBody) GoString() string {
	return s.String()
}

func (s *QuerySmsAuthorizationLetterResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *QuerySmsAuthorizationLetterResponseBody) GetCode() *string {
	return s.Code
}

func (s *QuerySmsAuthorizationLetterResponseBody) GetData() []*QuerySmsAuthorizationLetterResponseBodyData {
	return s.Data
}

func (s *QuerySmsAuthorizationLetterResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QuerySmsAuthorizationLetterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QuerySmsAuthorizationLetterResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QuerySmsAuthorizationLetterResponseBody) SetAccessDeniedDetail(v string) *QuerySmsAuthorizationLetterResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *QuerySmsAuthorizationLetterResponseBody) SetCode(v string) *QuerySmsAuthorizationLetterResponseBody {
	s.Code = &v
	return s
}

func (s *QuerySmsAuthorizationLetterResponseBody) SetData(v []*QuerySmsAuthorizationLetterResponseBodyData) *QuerySmsAuthorizationLetterResponseBody {
	s.Data = v
	return s
}

func (s *QuerySmsAuthorizationLetterResponseBody) SetMessage(v string) *QuerySmsAuthorizationLetterResponseBody {
	s.Message = &v
	return s
}

func (s *QuerySmsAuthorizationLetterResponseBody) SetRequestId(v string) *QuerySmsAuthorizationLetterResponseBody {
	s.RequestId = &v
	return s
}

func (s *QuerySmsAuthorizationLetterResponseBody) SetSuccess(v bool) *QuerySmsAuthorizationLetterResponseBody {
	s.Success = &v
	return s
}

func (s *QuerySmsAuthorizationLetterResponseBody) Validate() error {
	return dara.Validate(s)
}

type QuerySmsAuthorizationLetterResponseBodyData struct {
	// 委托授权方
	//
	// example:
	//
	// 示例值示例值
	Authorization *string `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
	// 委托授权书有效期
	//
	// example:
	//
	// 2023-01-01~2026-01-01
	AuthorizationLetterExpDate *string `json:"AuthorizationLetterExpDate,omitempty" xml:"AuthorizationLetterExpDate,omitempty"`
	// 委托授权书id
	//
	// example:
	//
	// 10000******
	AuthorizationLetterId *int64 `json:"AuthorizationLetterId,omitempty" xml:"AuthorizationLetterId,omitempty"`
	// 委托授权书命名
	//
	// example:
	//
	// 示例值示例值
	AuthorizationLetterName *string `json:"AuthorizationLetterName,omitempty" xml:"AuthorizationLetterName,omitempty"`
	// 委托授权书图片地址
	AuthorizationLetterPic *string `json:"AuthorizationLetterPic,omitempty" xml:"AuthorizationLetterPic,omitempty"`
	// 授权方统一社会信用代码
	//
	// example:
	//
	// 9****************A
	OrganizationCode *string `json:"OrganizationCode,omitempty" xml:"OrganizationCode,omitempty"`
	// 被委托授权方
	//
	// example:
	//
	// 示例值示例值示例值
	ProxyAuthorization *string `json:"ProxyAuthorization,omitempty" xml:"ProxyAuthorization,omitempty"`
	// 委托授权签名范围
	//
	// example:
	//
	// 示例值
	SignScope *string `json:"SignScope,omitempty" xml:"SignScope,omitempty"`
	// 委托授权书审核状态（初始化INT/审核通过PASSED）
	//
	// example:
	//
	// PASSED
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// 委托授权书可用状态（可用VALID/不可用INVALID）
	//
	// example:
	//
	// VALID
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s QuerySmsAuthorizationLetterResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QuerySmsAuthorizationLetterResponseBodyData) GoString() string {
	return s.String()
}

func (s *QuerySmsAuthorizationLetterResponseBodyData) GetAuthorization() *string {
	return s.Authorization
}

func (s *QuerySmsAuthorizationLetterResponseBodyData) GetAuthorizationLetterExpDate() *string {
	return s.AuthorizationLetterExpDate
}

func (s *QuerySmsAuthorizationLetterResponseBodyData) GetAuthorizationLetterId() *int64 {
	return s.AuthorizationLetterId
}

func (s *QuerySmsAuthorizationLetterResponseBodyData) GetAuthorizationLetterName() *string {
	return s.AuthorizationLetterName
}

func (s *QuerySmsAuthorizationLetterResponseBodyData) GetAuthorizationLetterPic() *string {
	return s.AuthorizationLetterPic
}

func (s *QuerySmsAuthorizationLetterResponseBodyData) GetOrganizationCode() *string {
	return s.OrganizationCode
}

func (s *QuerySmsAuthorizationLetterResponseBodyData) GetProxyAuthorization() *string {
	return s.ProxyAuthorization
}

func (s *QuerySmsAuthorizationLetterResponseBodyData) GetSignScope() *string {
	return s.SignScope
}

func (s *QuerySmsAuthorizationLetterResponseBodyData) GetState() *string {
	return s.State
}

func (s *QuerySmsAuthorizationLetterResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *QuerySmsAuthorizationLetterResponseBodyData) SetAuthorization(v string) *QuerySmsAuthorizationLetterResponseBodyData {
	s.Authorization = &v
	return s
}

func (s *QuerySmsAuthorizationLetterResponseBodyData) SetAuthorizationLetterExpDate(v string) *QuerySmsAuthorizationLetterResponseBodyData {
	s.AuthorizationLetterExpDate = &v
	return s
}

func (s *QuerySmsAuthorizationLetterResponseBodyData) SetAuthorizationLetterId(v int64) *QuerySmsAuthorizationLetterResponseBodyData {
	s.AuthorizationLetterId = &v
	return s
}

func (s *QuerySmsAuthorizationLetterResponseBodyData) SetAuthorizationLetterName(v string) *QuerySmsAuthorizationLetterResponseBodyData {
	s.AuthorizationLetterName = &v
	return s
}

func (s *QuerySmsAuthorizationLetterResponseBodyData) SetAuthorizationLetterPic(v string) *QuerySmsAuthorizationLetterResponseBodyData {
	s.AuthorizationLetterPic = &v
	return s
}

func (s *QuerySmsAuthorizationLetterResponseBodyData) SetOrganizationCode(v string) *QuerySmsAuthorizationLetterResponseBodyData {
	s.OrganizationCode = &v
	return s
}

func (s *QuerySmsAuthorizationLetterResponseBodyData) SetProxyAuthorization(v string) *QuerySmsAuthorizationLetterResponseBodyData {
	s.ProxyAuthorization = &v
	return s
}

func (s *QuerySmsAuthorizationLetterResponseBodyData) SetSignScope(v string) *QuerySmsAuthorizationLetterResponseBodyData {
	s.SignScope = &v
	return s
}

func (s *QuerySmsAuthorizationLetterResponseBodyData) SetState(v string) *QuerySmsAuthorizationLetterResponseBodyData {
	s.State = &v
	return s
}

func (s *QuerySmsAuthorizationLetterResponseBodyData) SetStatus(v string) *QuerySmsAuthorizationLetterResponseBodyData {
	s.Status = &v
	return s
}

func (s *QuerySmsAuthorizationLetterResponseBodyData) Validate() error {
	return dara.Validate(s)
}
