// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReadMessageListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *ReadMessageListRequest
	GetAcceptLanguage() *string
	SetAppName(v string) *ReadMessageListRequest
	GetAppName() *string
	SetBizName(v string) *ReadMessageListRequest
	GetBizName() *string
	SetCallerProtocol(v string) *ReadMessageListRequest
	GetCallerProtocol() *string
	SetClassId(v int64) *ReadMessageListRequest
	GetClassId() *int64
	SetClientSource(v string) *ReadMessageListRequest
	GetClientSource() *string
	SetContent(v string) *ReadMessageListRequest
	GetContent() *string
	SetCookies(v string) *ReadMessageListRequest
	GetCookies() *string
	SetLoc(v string) *ReadMessageListRequest
	GetLoc() *string
	SetMaxResults(v int32) *ReadMessageListRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ReadMessageListRequest
	GetNextToken() *string
	SetPage(v int32) *ReadMessageListRequest
	GetPage() *int32
	SetPageSize(v int32) *ReadMessageListRequest
	GetPageSize() *int32
	SetSrcUrl(v string) *ReadMessageListRequest
	GetSrcUrl() *string
	SetStatus(v int32) *ReadMessageListRequest
	GetStatus() *int32
	SetTenantCode(v string) *ReadMessageListRequest
	GetTenantCode() *string
	SetTitle(v string) *ReadMessageListRequest
	GetTitle() *string
	SetUidType(v string) *ReadMessageListRequest
	GetUidType() *string
}

type ReadMessageListRequest struct {
	// 语言，默认为简体中文
	//
	// example:
	//
	// zh-CN
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// 系统参数，无需填写
	//
	// example:
	//
	// /
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// 系统参数，无需填写
	//
	// example:
	//
	// /
	BizName *string `json:"BizName,omitempty" xml:"BizName,omitempty"`
	// 系统参数，无需填写
	//
	// example:
	//
	// /
	CallerProtocol *string `json:"CallerProtocol,omitempty" xml:"CallerProtocol,omitempty"`
	// 消息类目ID
	//
	// example:
	//
	// 1
	ClassId *int64 `json:"ClassId,omitempty" xml:"ClassId,omitempty"`
	// 系统参数，无需填写
	//
	// example:
	//
	// /
	ClientSource *string `json:"ClientSource,omitempty" xml:"ClientSource,omitempty"`
	// 消息内容，用于模糊搜索
	//
	// example:
	//
	// "消息内容示例“
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// 系统参数，无需填写
	//
	// example:
	//
	// /
	Cookies *string `json:"Cookies,omitempty" xml:"Cookies,omitempty"`
	// 栏位 nav代表控制台topbar
	//
	// example:
	//
	// /
	Loc *string `json:"Loc,omitempty" xml:"Loc,omitempty"`
	// 系统参数，无需填写
	//
	// example:
	//
	// /
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// 系统参数，无需填写
	//
	// example:
	//
	// /
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// 分页查询页码
	//
	// example:
	//
	// 2
	Page *int32 `json:"Page,omitempty" xml:"Page,omitempty"`
	// 分页查询大小
	//
	// example:
	//
	// 5
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// 系统参数，无需填写
	//
	// example:
	//
	// /
	SrcUrl *string `json:"SrcUrl,omitempty" xml:"SrcUrl,omitempty"`
	// 消息状态，已读为1，未读为0
	//
	// example:
	//
	// 0
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// 系统参数，无需填写
	//
	// example:
	//
	// /
	TenantCode *string `json:"TenantCode,omitempty" xml:"TenantCode,omitempty"`
	// 消息标题，用于模糊搜索
	//
	// example:
	//
	// "标题示例“
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// 系统参数，无需填写
	//
	// example:
	//
	// /
	UidType *string `json:"UidType,omitempty" xml:"UidType,omitempty"`
}

func (s ReadMessageListRequest) String() string {
	return dara.Prettify(s)
}

func (s ReadMessageListRequest) GoString() string {
	return s.String()
}

func (s *ReadMessageListRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *ReadMessageListRequest) GetAppName() *string {
	return s.AppName
}

func (s *ReadMessageListRequest) GetBizName() *string {
	return s.BizName
}

func (s *ReadMessageListRequest) GetCallerProtocol() *string {
	return s.CallerProtocol
}

func (s *ReadMessageListRequest) GetClassId() *int64 {
	return s.ClassId
}

func (s *ReadMessageListRequest) GetClientSource() *string {
	return s.ClientSource
}

func (s *ReadMessageListRequest) GetContent() *string {
	return s.Content
}

func (s *ReadMessageListRequest) GetCookies() *string {
	return s.Cookies
}

func (s *ReadMessageListRequest) GetLoc() *string {
	return s.Loc
}

func (s *ReadMessageListRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ReadMessageListRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ReadMessageListRequest) GetPage() *int32 {
	return s.Page
}

func (s *ReadMessageListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ReadMessageListRequest) GetSrcUrl() *string {
	return s.SrcUrl
}

func (s *ReadMessageListRequest) GetStatus() *int32 {
	return s.Status
}

func (s *ReadMessageListRequest) GetTenantCode() *string {
	return s.TenantCode
}

func (s *ReadMessageListRequest) GetTitle() *string {
	return s.Title
}

func (s *ReadMessageListRequest) GetUidType() *string {
	return s.UidType
}

func (s *ReadMessageListRequest) SetAcceptLanguage(v string) *ReadMessageListRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *ReadMessageListRequest) SetAppName(v string) *ReadMessageListRequest {
	s.AppName = &v
	return s
}

func (s *ReadMessageListRequest) SetBizName(v string) *ReadMessageListRequest {
	s.BizName = &v
	return s
}

func (s *ReadMessageListRequest) SetCallerProtocol(v string) *ReadMessageListRequest {
	s.CallerProtocol = &v
	return s
}

func (s *ReadMessageListRequest) SetClassId(v int64) *ReadMessageListRequest {
	s.ClassId = &v
	return s
}

func (s *ReadMessageListRequest) SetClientSource(v string) *ReadMessageListRequest {
	s.ClientSource = &v
	return s
}

func (s *ReadMessageListRequest) SetContent(v string) *ReadMessageListRequest {
	s.Content = &v
	return s
}

func (s *ReadMessageListRequest) SetCookies(v string) *ReadMessageListRequest {
	s.Cookies = &v
	return s
}

func (s *ReadMessageListRequest) SetLoc(v string) *ReadMessageListRequest {
	s.Loc = &v
	return s
}

func (s *ReadMessageListRequest) SetMaxResults(v int32) *ReadMessageListRequest {
	s.MaxResults = &v
	return s
}

func (s *ReadMessageListRequest) SetNextToken(v string) *ReadMessageListRequest {
	s.NextToken = &v
	return s
}

func (s *ReadMessageListRequest) SetPage(v int32) *ReadMessageListRequest {
	s.Page = &v
	return s
}

func (s *ReadMessageListRequest) SetPageSize(v int32) *ReadMessageListRequest {
	s.PageSize = &v
	return s
}

func (s *ReadMessageListRequest) SetSrcUrl(v string) *ReadMessageListRequest {
	s.SrcUrl = &v
	return s
}

func (s *ReadMessageListRequest) SetStatus(v int32) *ReadMessageListRequest {
	s.Status = &v
	return s
}

func (s *ReadMessageListRequest) SetTenantCode(v string) *ReadMessageListRequest {
	s.TenantCode = &v
	return s
}

func (s *ReadMessageListRequest) SetTitle(v string) *ReadMessageListRequest {
	s.Title = &v
	return s
}

func (s *ReadMessageListRequest) SetUidType(v string) *ReadMessageListRequest {
	s.UidType = &v
	return s
}

func (s *ReadMessageListRequest) Validate() error {
	return dara.Validate(s)
}
