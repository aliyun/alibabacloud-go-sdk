// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReadRevisionHistoryListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *ReadRevisionHistoryListRequest
	GetAcceptLanguage() *string
	SetAppName(v string) *ReadRevisionHistoryListRequest
	GetAppName() *string
	SetBizName(v string) *ReadRevisionHistoryListRequest
	GetBizName() *string
	SetCallerProtocol(v string) *ReadRevisionHistoryListRequest
	GetCallerProtocol() *string
	SetCategoryCode(v string) *ReadRevisionHistoryListRequest
	GetCategoryCode() *string
	SetChannelGroupCode(v string) *ReadRevisionHistoryListRequest
	GetChannelGroupCode() *string
	SetClientSource(v string) *ReadRevisionHistoryListRequest
	GetClientSource() *string
	SetCookies(v string) *ReadRevisionHistoryListRequest
	GetCookies() *string
	SetPageInfo(v *ReadRevisionHistoryListRequestPageInfo) *ReadRevisionHistoryListRequest
	GetPageInfo() *ReadRevisionHistoryListRequestPageInfo
	SetSrcUrl(v string) *ReadRevisionHistoryListRequest
	GetSrcUrl() *string
	SetTenantCode(v string) *ReadRevisionHistoryListRequest
	GetTenantCode() *string
	SetUidType(v string) *ReadRevisionHistoryListRequest
	GetUidType() *string
}

type ReadRevisionHistoryListRequest struct {
	// The language. Automatically passed through by the browser. You can manually override this value.
	//
	// example:
	//
	// zh-CN
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// Ignored. No need to pass this parameter. The application name of the caller.
	//
	// example:
	//
	// /
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// Ignored. No need to pass this parameter. The business line of the caller.
	//
	// example:
	//
	// /
	BizName *string `json:"BizName,omitempty" xml:"BizName,omitempty"`
	// Ignored. No need to pass this parameter. The request protocol type.
	//
	// example:
	//
	// /
	CallerProtocol *string `json:"CallerProtocol,omitempty" xml:"CallerProtocol,omitempty"`
	// The category code.
	//
	// example:
	//
	// prod_edu_content
	CategoryCode *string `json:"CategoryCode,omitempty" xml:"CategoryCode,omitempty"`
	// The channel group.
	//
	// example:
	//
	// base
	ChannelGroupCode *string `json:"ChannelGroupCode,omitempty" xml:"ChannelGroupCode,omitempty"`
	// Ignored. No need to pass this parameter. The source of the operation terminal.
	//
	// example:
	//
	// /
	ClientSource *string `json:"ClientSource,omitempty" xml:"ClientSource,omitempty"`
	// Ignored. No need to pass this parameter. The user cookies.
	//
	// example:
	//
	// /
	Cookies *string `json:"Cookies,omitempty" xml:"Cookies,omitempty"`
	// The pagination information.
	PageInfo *ReadRevisionHistoryListRequestPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// Ignored. No need to pass this parameter. The source page URL.
	//
	// example:
	//
	// /
	SrcUrl *string `json:"SrcUrl,omitempty" xml:"SrcUrl,omitempty"`
	// Ignored. No need to pass this parameter. The tenant information.
	//
	// example:
	//
	// /
	TenantCode *string `json:"TenantCode,omitempty" xml:"TenantCode,omitempty"`
	// Ignored. No need to pass this parameter. The user type.
	//
	// example:
	//
	// /
	UidType *string `json:"UidType,omitempty" xml:"UidType,omitempty"`
}

func (s ReadRevisionHistoryListRequest) String() string {
	return dara.Prettify(s)
}

func (s ReadRevisionHistoryListRequest) GoString() string {
	return s.String()
}

func (s *ReadRevisionHistoryListRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *ReadRevisionHistoryListRequest) GetAppName() *string {
	return s.AppName
}

func (s *ReadRevisionHistoryListRequest) GetBizName() *string {
	return s.BizName
}

func (s *ReadRevisionHistoryListRequest) GetCallerProtocol() *string {
	return s.CallerProtocol
}

func (s *ReadRevisionHistoryListRequest) GetCategoryCode() *string {
	return s.CategoryCode
}

func (s *ReadRevisionHistoryListRequest) GetChannelGroupCode() *string {
	return s.ChannelGroupCode
}

func (s *ReadRevisionHistoryListRequest) GetClientSource() *string {
	return s.ClientSource
}

func (s *ReadRevisionHistoryListRequest) GetCookies() *string {
	return s.Cookies
}

func (s *ReadRevisionHistoryListRequest) GetPageInfo() *ReadRevisionHistoryListRequestPageInfo {
	return s.PageInfo
}

func (s *ReadRevisionHistoryListRequest) GetSrcUrl() *string {
	return s.SrcUrl
}

func (s *ReadRevisionHistoryListRequest) GetTenantCode() *string {
	return s.TenantCode
}

func (s *ReadRevisionHistoryListRequest) GetUidType() *string {
	return s.UidType
}

func (s *ReadRevisionHistoryListRequest) SetAcceptLanguage(v string) *ReadRevisionHistoryListRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *ReadRevisionHistoryListRequest) SetAppName(v string) *ReadRevisionHistoryListRequest {
	s.AppName = &v
	return s
}

func (s *ReadRevisionHistoryListRequest) SetBizName(v string) *ReadRevisionHistoryListRequest {
	s.BizName = &v
	return s
}

func (s *ReadRevisionHistoryListRequest) SetCallerProtocol(v string) *ReadRevisionHistoryListRequest {
	s.CallerProtocol = &v
	return s
}

func (s *ReadRevisionHistoryListRequest) SetCategoryCode(v string) *ReadRevisionHistoryListRequest {
	s.CategoryCode = &v
	return s
}

func (s *ReadRevisionHistoryListRequest) SetChannelGroupCode(v string) *ReadRevisionHistoryListRequest {
	s.ChannelGroupCode = &v
	return s
}

func (s *ReadRevisionHistoryListRequest) SetClientSource(v string) *ReadRevisionHistoryListRequest {
	s.ClientSource = &v
	return s
}

func (s *ReadRevisionHistoryListRequest) SetCookies(v string) *ReadRevisionHistoryListRequest {
	s.Cookies = &v
	return s
}

func (s *ReadRevisionHistoryListRequest) SetPageInfo(v *ReadRevisionHistoryListRequestPageInfo) *ReadRevisionHistoryListRequest {
	s.PageInfo = v
	return s
}

func (s *ReadRevisionHistoryListRequest) SetSrcUrl(v string) *ReadRevisionHistoryListRequest {
	s.SrcUrl = &v
	return s
}

func (s *ReadRevisionHistoryListRequest) SetTenantCode(v string) *ReadRevisionHistoryListRequest {
	s.TenantCode = &v
	return s
}

func (s *ReadRevisionHistoryListRequest) SetUidType(v string) *ReadRevisionHistoryListRequest {
	s.UidType = &v
	return s
}

func (s *ReadRevisionHistoryListRequest) Validate() error {
	if s.PageInfo != nil {
		if err := s.PageInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ReadRevisionHistoryListRequestPageInfo struct {
	// The maximum number of entries to return.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token for the next page of data.
	//
	// example:
	//
	// e2b5170336162251e0b5e8707e68181f
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// Specifies whether to return the total count.
	//
	// example:
	//
	// false
	ReturnTotalCount *bool `json:"ReturnTotalCount,omitempty" xml:"ReturnTotalCount,omitempty"`
}

func (s ReadRevisionHistoryListRequestPageInfo) String() string {
	return dara.Prettify(s)
}

func (s ReadRevisionHistoryListRequestPageInfo) GoString() string {
	return s.String()
}

func (s *ReadRevisionHistoryListRequestPageInfo) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ReadRevisionHistoryListRequestPageInfo) GetNextToken() *string {
	return s.NextToken
}

func (s *ReadRevisionHistoryListRequestPageInfo) GetReturnTotalCount() *bool {
	return s.ReturnTotalCount
}

func (s *ReadRevisionHistoryListRequestPageInfo) SetMaxResults(v int32) *ReadRevisionHistoryListRequestPageInfo {
	s.MaxResults = &v
	return s
}

func (s *ReadRevisionHistoryListRequestPageInfo) SetNextToken(v string) *ReadRevisionHistoryListRequestPageInfo {
	s.NextToken = &v
	return s
}

func (s *ReadRevisionHistoryListRequestPageInfo) SetReturnTotalCount(v bool) *ReadRevisionHistoryListRequestPageInfo {
	s.ReturnTotalCount = &v
	return s
}

func (s *ReadRevisionHistoryListRequestPageInfo) Validate() error {
	return dara.Validate(s)
}
