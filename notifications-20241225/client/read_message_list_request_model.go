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
	SetGroupCode(v string) *ReadMessageListRequest
	GetGroupCode() *string
	SetHistory(v string) *ReadMessageListRequest
	GetHistory() *string
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
	// The language. Default value: Simplified Chinese.
	//
	// example:
	//
	// zh-CN
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// A system parameter. You do not need to specify this parameter.
	//
	// example:
	//
	// /
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// A system parameter. You do not need to specify this parameter.
	//
	// example:
	//
	// /
	BizName *string `json:"BizName,omitempty" xml:"BizName,omitempty"`
	// A system parameter. You do not need to specify this parameter.
	//
	// example:
	//
	// /
	CallerProtocol *string `json:"CallerProtocol,omitempty" xml:"CallerProtocol,omitempty"`
	// Deprecated.
	//
	// example:
	//
	// 1
	ClassId *int64 `json:"ClassId,omitempty" xml:"ClassId,omitempty"`
	// A system parameter. You do not need to specify this parameter.
	//
	// example:
	//
	// /
	ClientSource *string `json:"ClientSource,omitempty" xml:"ClientSource,omitempty"`
	// The message content. This parameter is used for fuzzy match.
	//
	// example:
	//
	// "消息内容示例“
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// A system parameter. You do not need to specify this parameter.
	//
	// example:
	//
	// /
	Cookies *string `json:"Cookies,omitempty" xml:"Cookies,omitempty"`
	// The group code.
	//
	// example:
	//
	// test
	GroupCode *string `json:"GroupCode,omitempty" xml:"GroupCode,omitempty"`
	// Specifies whether the messages are historical messages.
	//
	// example:
	//
	// true
	History *string `json:"History,omitempty" xml:"History,omitempty"`
	// The location.
	//
	// example:
	//
	// nav
	Loc *string `json:"Loc,omitempty" xml:"Loc,omitempty"`
	// A system parameter. You do not need to specify this parameter.
	//
	// example:
	//
	// /
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// A system parameter. You do not need to specify this parameter.
	//
	// example:
	//
	// /
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The page number for the paged query.
	//
	// example:
	//
	// 2
	Page *int32 `json:"Page,omitempty" xml:"Page,omitempty"`
	// The page size for the paged query.
	//
	// example:
	//
	// 5
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// A system parameter. You do not need to specify this parameter.
	//
	// example:
	//
	// /
	SrcUrl *string `json:"SrcUrl,omitempty" xml:"SrcUrl,omitempty"`
	// The message status. A value of 1 indicates read. A value of 0 indicates unread. A value of -1 indicates all. Default value: -1.
	//
	// example:
	//
	// 0
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// A system parameter. You do not need to specify this parameter.
	//
	// example:
	//
	// /
	TenantCode *string `json:"TenantCode,omitempty" xml:"TenantCode,omitempty"`
	// The message title. This parameter is used for fuzzy match.
	//
	// example:
	//
	// "标题示例“
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// A system parameter. You do not need to specify this parameter.
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

func (s *ReadMessageListRequest) GetGroupCode() *string {
	return s.GroupCode
}

func (s *ReadMessageListRequest) GetHistory() *string {
	return s.History
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

func (s *ReadMessageListRequest) SetGroupCode(v string) *ReadMessageListRequest {
	s.GroupCode = &v
	return s
}

func (s *ReadMessageListRequest) SetHistory(v string) *ReadMessageListRequest {
	s.History = &v
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
