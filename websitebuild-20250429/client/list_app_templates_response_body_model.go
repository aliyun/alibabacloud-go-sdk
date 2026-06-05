// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAppTemplatesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ListAppTemplatesResponseBody
	GetAccessDeniedDetail() *string
	SetAllowRetry(v bool) *ListAppTemplatesResponseBody
	GetAllowRetry() *bool
	SetAppName(v string) *ListAppTemplatesResponseBody
	GetAppName() *string
	SetDynamicCode(v string) *ListAppTemplatesResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *ListAppTemplatesResponseBody
	GetDynamicMessage() *string
	SetErrorArgs(v []interface{}) *ListAppTemplatesResponseBody
	GetErrorArgs() []interface{}
	SetMaxResults(v int32) *ListAppTemplatesResponseBody
	GetMaxResults() *int32
	SetModule(v *ListAppTemplatesResponseBodyModule) *ListAppTemplatesResponseBody
	GetModule() *ListAppTemplatesResponseBodyModule
	SetNextToken(v string) *ListAppTemplatesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListAppTemplatesResponseBody
	GetRequestId() *string
	SetRootErrorCode(v string) *ListAppTemplatesResponseBody
	GetRootErrorCode() *string
	SetRootErrorMsg(v string) *ListAppTemplatesResponseBody
	GetRootErrorMsg() *string
	SetSynchro(v bool) *ListAppTemplatesResponseBody
	GetSynchro() *bool
}

type ListAppTemplatesResponseBody struct {
	// example:
	//
	// {}
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// False
	AllowRetry *bool `json:"AllowRetry,omitempty" xml:"AllowRetry,omitempty"`
	// example:
	//
	// or
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// example:
	//
	// ERROR-oo1
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// example:
	//
	// SYSTEM_ERROR
	DynamicMessage *string       `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	ErrorArgs      []interface{} `json:"ErrorArgs,omitempty" xml:"ErrorArgs,omitempty" type:"Repeated"`
	// example:
	//
	// 10
	MaxResults *int32                              `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	Module     *ListAppTemplatesResponseBodyModule `json:"Module,omitempty" xml:"Module,omitempty" type:"Struct"`
	// example:
	//
	// AAAAARbaCuN6hiD08qrLdwJ9Fh3BFw8paIJ7ylB6A7Qn9JjM
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 6C6B99AC-39EC-5350-874C-204128C905E6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// SYSTEM.ERROR
	RootErrorCode *string `json:"RootErrorCode,omitempty" xml:"RootErrorCode,omitempty"`
	// example:
	//
	// 系统异常
	RootErrorMsg *string `json:"RootErrorMsg,omitempty" xml:"RootErrorMsg,omitempty"`
	// example:
	//
	// True
	Synchro *bool `json:"Synchro,omitempty" xml:"Synchro,omitempty"`
}

func (s ListAppTemplatesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAppTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *ListAppTemplatesResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ListAppTemplatesResponseBody) GetAllowRetry() *bool {
	return s.AllowRetry
}

func (s *ListAppTemplatesResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *ListAppTemplatesResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *ListAppTemplatesResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *ListAppTemplatesResponseBody) GetErrorArgs() []interface{} {
	return s.ErrorArgs
}

func (s *ListAppTemplatesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListAppTemplatesResponseBody) GetModule() *ListAppTemplatesResponseBodyModule {
	return s.Module
}

func (s *ListAppTemplatesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAppTemplatesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAppTemplatesResponseBody) GetRootErrorCode() *string {
	return s.RootErrorCode
}

func (s *ListAppTemplatesResponseBody) GetRootErrorMsg() *string {
	return s.RootErrorMsg
}

func (s *ListAppTemplatesResponseBody) GetSynchro() *bool {
	return s.Synchro
}

func (s *ListAppTemplatesResponseBody) SetAccessDeniedDetail(v string) *ListAppTemplatesResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ListAppTemplatesResponseBody) SetAllowRetry(v bool) *ListAppTemplatesResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *ListAppTemplatesResponseBody) SetAppName(v string) *ListAppTemplatesResponseBody {
	s.AppName = &v
	return s
}

func (s *ListAppTemplatesResponseBody) SetDynamicCode(v string) *ListAppTemplatesResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *ListAppTemplatesResponseBody) SetDynamicMessage(v string) *ListAppTemplatesResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *ListAppTemplatesResponseBody) SetErrorArgs(v []interface{}) *ListAppTemplatesResponseBody {
	s.ErrorArgs = v
	return s
}

func (s *ListAppTemplatesResponseBody) SetMaxResults(v int32) *ListAppTemplatesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListAppTemplatesResponseBody) SetModule(v *ListAppTemplatesResponseBodyModule) *ListAppTemplatesResponseBody {
	s.Module = v
	return s
}

func (s *ListAppTemplatesResponseBody) SetNextToken(v string) *ListAppTemplatesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListAppTemplatesResponseBody) SetRequestId(v string) *ListAppTemplatesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAppTemplatesResponseBody) SetRootErrorCode(v string) *ListAppTemplatesResponseBody {
	s.RootErrorCode = &v
	return s
}

func (s *ListAppTemplatesResponseBody) SetRootErrorMsg(v string) *ListAppTemplatesResponseBody {
	s.RootErrorMsg = &v
	return s
}

func (s *ListAppTemplatesResponseBody) SetSynchro(v bool) *ListAppTemplatesResponseBody {
	s.Synchro = &v
	return s
}

func (s *ListAppTemplatesResponseBody) Validate() error {
	if s.Module != nil {
		if err := s.Module.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListAppTemplatesResponseBodyModule struct {
	// example:
	//
	// 1
	CurrentPageNum *int32                                    `json:"CurrentPageNum,omitempty" xml:"CurrentPageNum,omitempty"`
	Data           []*ListAppTemplatesResponseBodyModuleData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Next           *ListAppTemplatesResponseBodyModuleNext   `json:"Next,omitempty" xml:"Next,omitempty" type:"Struct"`
	// example:
	//
	// False
	NextPage *bool `json:"NextPage,omitempty" xml:"NextPage,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// False
	PrePage *bool `json:"PrePage,omitempty" xml:"PrePage,omitempty"`
	// example:
	//
	// False
	ResultLimit *bool `json:"ResultLimit,omitempty" xml:"ResultLimit,omitempty"`
	// example:
	//
	// 1
	TotalItemNum *int32 `json:"TotalItemNum,omitempty" xml:"TotalItemNum,omitempty"`
	// example:
	//
	// 1
	TotalPageNum *int32 `json:"TotalPageNum,omitempty" xml:"TotalPageNum,omitempty"`
}

func (s ListAppTemplatesResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s ListAppTemplatesResponseBodyModule) GoString() string {
	return s.String()
}

func (s *ListAppTemplatesResponseBodyModule) GetCurrentPageNum() *int32 {
	return s.CurrentPageNum
}

func (s *ListAppTemplatesResponseBodyModule) GetData() []*ListAppTemplatesResponseBodyModuleData {
	return s.Data
}

func (s *ListAppTemplatesResponseBodyModule) GetNext() *ListAppTemplatesResponseBodyModuleNext {
	return s.Next
}

func (s *ListAppTemplatesResponseBodyModule) GetNextPage() *bool {
	return s.NextPage
}

func (s *ListAppTemplatesResponseBodyModule) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAppTemplatesResponseBodyModule) GetPrePage() *bool {
	return s.PrePage
}

func (s *ListAppTemplatesResponseBodyModule) GetResultLimit() *bool {
	return s.ResultLimit
}

func (s *ListAppTemplatesResponseBodyModule) GetTotalItemNum() *int32 {
	return s.TotalItemNum
}

func (s *ListAppTemplatesResponseBodyModule) GetTotalPageNum() *int32 {
	return s.TotalPageNum
}

func (s *ListAppTemplatesResponseBodyModule) SetCurrentPageNum(v int32) *ListAppTemplatesResponseBodyModule {
	s.CurrentPageNum = &v
	return s
}

func (s *ListAppTemplatesResponseBodyModule) SetData(v []*ListAppTemplatesResponseBodyModuleData) *ListAppTemplatesResponseBodyModule {
	s.Data = v
	return s
}

func (s *ListAppTemplatesResponseBodyModule) SetNext(v *ListAppTemplatesResponseBodyModuleNext) *ListAppTemplatesResponseBodyModule {
	s.Next = v
	return s
}

func (s *ListAppTemplatesResponseBodyModule) SetNextPage(v bool) *ListAppTemplatesResponseBodyModule {
	s.NextPage = &v
	return s
}

func (s *ListAppTemplatesResponseBodyModule) SetPageSize(v int32) *ListAppTemplatesResponseBodyModule {
	s.PageSize = &v
	return s
}

func (s *ListAppTemplatesResponseBodyModule) SetPrePage(v bool) *ListAppTemplatesResponseBodyModule {
	s.PrePage = &v
	return s
}

func (s *ListAppTemplatesResponseBodyModule) SetResultLimit(v bool) *ListAppTemplatesResponseBodyModule {
	s.ResultLimit = &v
	return s
}

func (s *ListAppTemplatesResponseBodyModule) SetTotalItemNum(v int32) *ListAppTemplatesResponseBodyModule {
	s.TotalItemNum = &v
	return s
}

func (s *ListAppTemplatesResponseBodyModule) SetTotalPageNum(v int32) *ListAppTemplatesResponseBodyModule {
	s.TotalPageNum = &v
	return s
}

func (s *ListAppTemplatesResponseBodyModule) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Next != nil {
		if err := s.Next.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListAppTemplatesResponseBodyModuleData struct {
	// example:
	//
	// supabase
	AppType *string `json:"AppType,omitempty" xml:"AppType,omitempty"`
	// example:
	//
	// WD20250703155602000001
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// example:
	//
	// Red
	ColorScheme *string `json:"ColorScheme,omitempty" xml:"ColorScheme,omitempty"`
	// example:
	//
	// red
	ColorSchemeName *string `json:"ColorSchemeName,omitempty" xml:"ColorSchemeName,omitempty"`
	// example:
	//
	// 100
	CopyCount *int32 `json:"CopyCount,omitempty" xml:"CopyCount,omitempty"`
	// example:
	//
	// SOAR
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// Use the UTC time format: yyyy-MM-ddTHH:mm:ss.SSSZ
	//
	// example:
	//
	// 2025-11-20T02:26:38Z
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// Use the UTC time format: yyyy-MM-ddTHH:mm:ss.SSSZ
	//
	// example:
	//
	// 1627545952000
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// example:
	//
	// 9953352
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// NOUSE
	Industry *string `json:"Industry,omitempty" xml:"Industry,omitempty"`
	// example:
	//
	// abc
	IndustryName *string `json:"IndustryName,omitempty" xml:"IndustryName,omitempty"`
	// example:
	//
	// admin
	LastModifier *string `json:"LastModifier,omitempty" xml:"LastModifier,omitempty"`
	// example:
	//
	// 100
	LikeCount *int32 `json:"LikeCount,omitempty" xml:"LikeCount,omitempty"`
	Liked     *bool  `json:"Liked,omitempty" xml:"Liked,omitempty"`
	// example:
	//
	// {\\"memFrom\\": \\"text\\", \\"uid\\": \\"text\\", \\"labels\\": \\"jsonb\\", \\"content\\": \\"text\\", \\"fromId\\": \\"text\\", \\"uuid\\": \\"text\\"}
	Metadata *string `json:"Metadata,omitempty" xml:"Metadata,omitempty"`
	// example:
	//
	// https://preview-lyj.aliyuncs.com/preview/b2c5a245c44946b99cf5435210bbb8b8?subSceneIds=728166
	PreviewUrl *string `json:"PreviewUrl,omitempty" xml:"PreviewUrl,omitempty"`
	// example:
	//
	// EnterpriseVersion
	ProductVersion *string `json:"ProductVersion,omitempty" xml:"ProductVersion,omitempty"`
	// example:
	//
	// 1.0.1
	ProductVersionName *string `json:"ProductVersionName,omitempty" xml:"ProductVersionName,omitempty"`
	// example:
	//
	// 100
	ShareCount *int32 `json:"ShareCount,omitempty" xml:"ShareCount,omitempty"`
	// example:
	//
	// RUNNING
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// de10bf7d782392a70f293a3b1f7bb8fc
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// example:
	//
	// cfdna6
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// example:
	//
	// http://www.aliyun.com
	ThumbnailUrl *string `json:"ThumbnailUrl,omitempty" xml:"ThumbnailUrl,omitempty"`
	// example:
	//
	// 100
	ViewCount *int32 `json:"ViewCount,omitempty" xml:"ViewCount,omitempty"`
	// example:
	//
	// 255
	Weight *int32 `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s ListAppTemplatesResponseBodyModuleData) String() string {
	return dara.Prettify(s)
}

func (s ListAppTemplatesResponseBodyModuleData) GoString() string {
	return s.String()
}

func (s *ListAppTemplatesResponseBodyModuleData) GetAppType() *string {
	return s.AppType
}

func (s *ListAppTemplatesResponseBodyModuleData) GetBizId() *string {
	return s.BizId
}

func (s *ListAppTemplatesResponseBodyModuleData) GetColorScheme() *string {
	return s.ColorScheme
}

func (s *ListAppTemplatesResponseBodyModuleData) GetColorSchemeName() *string {
	return s.ColorSchemeName
}

func (s *ListAppTemplatesResponseBodyModuleData) GetCopyCount() *int32 {
	return s.CopyCount
}

func (s *ListAppTemplatesResponseBodyModuleData) GetCreator() *string {
	return s.Creator
}

func (s *ListAppTemplatesResponseBodyModuleData) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *ListAppTemplatesResponseBodyModuleData) GetGmtModifiedTime() *string {
	return s.GmtModifiedTime
}

func (s *ListAppTemplatesResponseBodyModuleData) GetId() *int64 {
	return s.Id
}

func (s *ListAppTemplatesResponseBodyModuleData) GetIndustry() *string {
	return s.Industry
}

func (s *ListAppTemplatesResponseBodyModuleData) GetIndustryName() *string {
	return s.IndustryName
}

func (s *ListAppTemplatesResponseBodyModuleData) GetLastModifier() *string {
	return s.LastModifier
}

func (s *ListAppTemplatesResponseBodyModuleData) GetLikeCount() *int32 {
	return s.LikeCount
}

func (s *ListAppTemplatesResponseBodyModuleData) GetLiked() *bool {
	return s.Liked
}

func (s *ListAppTemplatesResponseBodyModuleData) GetMetadata() *string {
	return s.Metadata
}

func (s *ListAppTemplatesResponseBodyModuleData) GetPreviewUrl() *string {
	return s.PreviewUrl
}

func (s *ListAppTemplatesResponseBodyModuleData) GetProductVersion() *string {
	return s.ProductVersion
}

func (s *ListAppTemplatesResponseBodyModuleData) GetProductVersionName() *string {
	return s.ProductVersionName
}

func (s *ListAppTemplatesResponseBodyModuleData) GetShareCount() *int32 {
	return s.ShareCount
}

func (s *ListAppTemplatesResponseBodyModuleData) GetStatus() *string {
	return s.Status
}

func (s *ListAppTemplatesResponseBodyModuleData) GetTemplateId() *string {
	return s.TemplateId
}

func (s *ListAppTemplatesResponseBodyModuleData) GetTemplateName() *string {
	return s.TemplateName
}

func (s *ListAppTemplatesResponseBodyModuleData) GetThumbnailUrl() *string {
	return s.ThumbnailUrl
}

func (s *ListAppTemplatesResponseBodyModuleData) GetViewCount() *int32 {
	return s.ViewCount
}

func (s *ListAppTemplatesResponseBodyModuleData) GetWeight() *int32 {
	return s.Weight
}

func (s *ListAppTemplatesResponseBodyModuleData) SetAppType(v string) *ListAppTemplatesResponseBodyModuleData {
	s.AppType = &v
	return s
}

func (s *ListAppTemplatesResponseBodyModuleData) SetBizId(v string) *ListAppTemplatesResponseBodyModuleData {
	s.BizId = &v
	return s
}

func (s *ListAppTemplatesResponseBodyModuleData) SetColorScheme(v string) *ListAppTemplatesResponseBodyModuleData {
	s.ColorScheme = &v
	return s
}

func (s *ListAppTemplatesResponseBodyModuleData) SetColorSchemeName(v string) *ListAppTemplatesResponseBodyModuleData {
	s.ColorSchemeName = &v
	return s
}

func (s *ListAppTemplatesResponseBodyModuleData) SetCopyCount(v int32) *ListAppTemplatesResponseBodyModuleData {
	s.CopyCount = &v
	return s
}

func (s *ListAppTemplatesResponseBodyModuleData) SetCreator(v string) *ListAppTemplatesResponseBodyModuleData {
	s.Creator = &v
	return s
}

func (s *ListAppTemplatesResponseBodyModuleData) SetGmtCreateTime(v string) *ListAppTemplatesResponseBodyModuleData {
	s.GmtCreateTime = &v
	return s
}

func (s *ListAppTemplatesResponseBodyModuleData) SetGmtModifiedTime(v string) *ListAppTemplatesResponseBodyModuleData {
	s.GmtModifiedTime = &v
	return s
}

func (s *ListAppTemplatesResponseBodyModuleData) SetId(v int64) *ListAppTemplatesResponseBodyModuleData {
	s.Id = &v
	return s
}

func (s *ListAppTemplatesResponseBodyModuleData) SetIndustry(v string) *ListAppTemplatesResponseBodyModuleData {
	s.Industry = &v
	return s
}

func (s *ListAppTemplatesResponseBodyModuleData) SetIndustryName(v string) *ListAppTemplatesResponseBodyModuleData {
	s.IndustryName = &v
	return s
}

func (s *ListAppTemplatesResponseBodyModuleData) SetLastModifier(v string) *ListAppTemplatesResponseBodyModuleData {
	s.LastModifier = &v
	return s
}

func (s *ListAppTemplatesResponseBodyModuleData) SetLikeCount(v int32) *ListAppTemplatesResponseBodyModuleData {
	s.LikeCount = &v
	return s
}

func (s *ListAppTemplatesResponseBodyModuleData) SetLiked(v bool) *ListAppTemplatesResponseBodyModuleData {
	s.Liked = &v
	return s
}

func (s *ListAppTemplatesResponseBodyModuleData) SetMetadata(v string) *ListAppTemplatesResponseBodyModuleData {
	s.Metadata = &v
	return s
}

func (s *ListAppTemplatesResponseBodyModuleData) SetPreviewUrl(v string) *ListAppTemplatesResponseBodyModuleData {
	s.PreviewUrl = &v
	return s
}

func (s *ListAppTemplatesResponseBodyModuleData) SetProductVersion(v string) *ListAppTemplatesResponseBodyModuleData {
	s.ProductVersion = &v
	return s
}

func (s *ListAppTemplatesResponseBodyModuleData) SetProductVersionName(v string) *ListAppTemplatesResponseBodyModuleData {
	s.ProductVersionName = &v
	return s
}

func (s *ListAppTemplatesResponseBodyModuleData) SetShareCount(v int32) *ListAppTemplatesResponseBodyModuleData {
	s.ShareCount = &v
	return s
}

func (s *ListAppTemplatesResponseBodyModuleData) SetStatus(v string) *ListAppTemplatesResponseBodyModuleData {
	s.Status = &v
	return s
}

func (s *ListAppTemplatesResponseBodyModuleData) SetTemplateId(v string) *ListAppTemplatesResponseBodyModuleData {
	s.TemplateId = &v
	return s
}

func (s *ListAppTemplatesResponseBodyModuleData) SetTemplateName(v string) *ListAppTemplatesResponseBodyModuleData {
	s.TemplateName = &v
	return s
}

func (s *ListAppTemplatesResponseBodyModuleData) SetThumbnailUrl(v string) *ListAppTemplatesResponseBodyModuleData {
	s.ThumbnailUrl = &v
	return s
}

func (s *ListAppTemplatesResponseBodyModuleData) SetViewCount(v int32) *ListAppTemplatesResponseBodyModuleData {
	s.ViewCount = &v
	return s
}

func (s *ListAppTemplatesResponseBodyModuleData) SetWeight(v int32) *ListAppTemplatesResponseBodyModuleData {
	s.Weight = &v
	return s
}

func (s *ListAppTemplatesResponseBodyModuleData) Validate() error {
	return dara.Validate(s)
}

type ListAppTemplatesResponseBodyModuleNext struct {
	// example:
	//
	// memory
	AppType *string `json:"AppType,omitempty" xml:"AppType,omitempty"`
	// example:
	//
	// WS20260502160409000001
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// example:
	//
	// red
	ColorScheme *string `json:"ColorScheme,omitempty" xml:"ColorScheme,omitempty"`
	// example:
	//
	// red
	ColorSchemeName *string `json:"ColorSchemeName,omitempty" xml:"ColorSchemeName,omitempty"`
	// example:
	//
	// 100
	CopyCount *int32 `json:"CopyCount,omitempty" xml:"CopyCount,omitempty"`
	// example:
	//
	// 208116853206125255
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// example:
	//
	// 2025-09-17 20:43:21
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// example:
	//
	// 2026-03-18T10:03:56+08:00
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// example:
	//
	// 1000039405002
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// NOUSE
	Industry *string `json:"Industry,omitempty" xml:"Industry,omitempty"`
	// example:
	//
	// abc
	IndustryName *string `json:"IndustryName,omitempty" xml:"IndustryName,omitempty"`
	// example:
	//
	// admin
	LastModifier *string `json:"LastModifier,omitempty" xml:"LastModifier,omitempty"`
	// example:
	//
	// 100
	LikeCount *int32 `json:"LikeCount,omitempty" xml:"LikeCount,omitempty"`
	Liked     *bool  `json:"Liked,omitempty" xml:"Liked,omitempty"`
	// example:
	//
	// {\\"knowledge_point\\":\\"text\\",\\"trunk_id\\":\\"text\\",\\"doc_name\\":\\"text\\",\\"knowledge_point_id\\":\\"text\\",\\"doc_id\\":\\"text\\",\\"trunk_content\\":\\"text\\"}
	Metadata *string `json:"Metadata,omitempty" xml:"Metadata,omitempty"`
	// example:
	//
	// https://preview-lyj.aliyuncs.com/preview/b989c9ac526e4fb48e018805f43d5fb1?subSceneIds=816576
	PreviewUrl *string `json:"PreviewUrl,omitempty" xml:"PreviewUrl,omitempty"`
	// example:
	//
	// BasicVersion
	ProductVersion *string `json:"ProductVersion,omitempty" xml:"ProductVersion,omitempty"`
	// example:
	//
	// 1.0.2
	ProductVersionName *string `json:"ProductVersionName,omitempty" xml:"ProductVersionName,omitempty"`
	// example:
	//
	// 100
	ShareCount *int32 `json:"ShareCount,omitempty" xml:"ShareCount,omitempty"`
	// example:
	//
	// Succeeded
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 3802aefc8cb0003b71286c47afc83624
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// example:
	//
	// iem
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// example:
	//
	// http://www.aliyun.com
	ThumbnailUrl *string `json:"ThumbnailUrl,omitempty" xml:"ThumbnailUrl,omitempty"`
	// example:
	//
	// 100
	ViewCount *int32 `json:"ViewCount,omitempty" xml:"ViewCount,omitempty"`
	// example:
	//
	// 1
	Weight *int32 `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s ListAppTemplatesResponseBodyModuleNext) String() string {
	return dara.Prettify(s)
}

func (s ListAppTemplatesResponseBodyModuleNext) GoString() string {
	return s.String()
}

func (s *ListAppTemplatesResponseBodyModuleNext) GetAppType() *string {
	return s.AppType
}

func (s *ListAppTemplatesResponseBodyModuleNext) GetBizId() *string {
	return s.BizId
}

func (s *ListAppTemplatesResponseBodyModuleNext) GetColorScheme() *string {
	return s.ColorScheme
}

func (s *ListAppTemplatesResponseBodyModuleNext) GetColorSchemeName() *string {
	return s.ColorSchemeName
}

func (s *ListAppTemplatesResponseBodyModuleNext) GetCopyCount() *int32 {
	return s.CopyCount
}

func (s *ListAppTemplatesResponseBodyModuleNext) GetCreator() *string {
	return s.Creator
}

func (s *ListAppTemplatesResponseBodyModuleNext) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ListAppTemplatesResponseBodyModuleNext) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ListAppTemplatesResponseBodyModuleNext) GetId() *int64 {
	return s.Id
}

func (s *ListAppTemplatesResponseBodyModuleNext) GetIndustry() *string {
	return s.Industry
}

func (s *ListAppTemplatesResponseBodyModuleNext) GetIndustryName() *string {
	return s.IndustryName
}

func (s *ListAppTemplatesResponseBodyModuleNext) GetLastModifier() *string {
	return s.LastModifier
}

func (s *ListAppTemplatesResponseBodyModuleNext) GetLikeCount() *int32 {
	return s.LikeCount
}

func (s *ListAppTemplatesResponseBodyModuleNext) GetLiked() *bool {
	return s.Liked
}

func (s *ListAppTemplatesResponseBodyModuleNext) GetMetadata() *string {
	return s.Metadata
}

func (s *ListAppTemplatesResponseBodyModuleNext) GetPreviewUrl() *string {
	return s.PreviewUrl
}

func (s *ListAppTemplatesResponseBodyModuleNext) GetProductVersion() *string {
	return s.ProductVersion
}

func (s *ListAppTemplatesResponseBodyModuleNext) GetProductVersionName() *string {
	return s.ProductVersionName
}

func (s *ListAppTemplatesResponseBodyModuleNext) GetShareCount() *int32 {
	return s.ShareCount
}

func (s *ListAppTemplatesResponseBodyModuleNext) GetStatus() *string {
	return s.Status
}

func (s *ListAppTemplatesResponseBodyModuleNext) GetTemplateId() *string {
	return s.TemplateId
}

func (s *ListAppTemplatesResponseBodyModuleNext) GetTemplateName() *string {
	return s.TemplateName
}

func (s *ListAppTemplatesResponseBodyModuleNext) GetThumbnailUrl() *string {
	return s.ThumbnailUrl
}

func (s *ListAppTemplatesResponseBodyModuleNext) GetViewCount() *int32 {
	return s.ViewCount
}

func (s *ListAppTemplatesResponseBodyModuleNext) GetWeight() *int32 {
	return s.Weight
}

func (s *ListAppTemplatesResponseBodyModuleNext) SetAppType(v string) *ListAppTemplatesResponseBodyModuleNext {
	s.AppType = &v
	return s
}

func (s *ListAppTemplatesResponseBodyModuleNext) SetBizId(v string) *ListAppTemplatesResponseBodyModuleNext {
	s.BizId = &v
	return s
}

func (s *ListAppTemplatesResponseBodyModuleNext) SetColorScheme(v string) *ListAppTemplatesResponseBodyModuleNext {
	s.ColorScheme = &v
	return s
}

func (s *ListAppTemplatesResponseBodyModuleNext) SetColorSchemeName(v string) *ListAppTemplatesResponseBodyModuleNext {
	s.ColorSchemeName = &v
	return s
}

func (s *ListAppTemplatesResponseBodyModuleNext) SetCopyCount(v int32) *ListAppTemplatesResponseBodyModuleNext {
	s.CopyCount = &v
	return s
}

func (s *ListAppTemplatesResponseBodyModuleNext) SetCreator(v string) *ListAppTemplatesResponseBodyModuleNext {
	s.Creator = &v
	return s
}

func (s *ListAppTemplatesResponseBodyModuleNext) SetGmtCreate(v string) *ListAppTemplatesResponseBodyModuleNext {
	s.GmtCreate = &v
	return s
}

func (s *ListAppTemplatesResponseBodyModuleNext) SetGmtModified(v string) *ListAppTemplatesResponseBodyModuleNext {
	s.GmtModified = &v
	return s
}

func (s *ListAppTemplatesResponseBodyModuleNext) SetId(v int64) *ListAppTemplatesResponseBodyModuleNext {
	s.Id = &v
	return s
}

func (s *ListAppTemplatesResponseBodyModuleNext) SetIndustry(v string) *ListAppTemplatesResponseBodyModuleNext {
	s.Industry = &v
	return s
}

func (s *ListAppTemplatesResponseBodyModuleNext) SetIndustryName(v string) *ListAppTemplatesResponseBodyModuleNext {
	s.IndustryName = &v
	return s
}

func (s *ListAppTemplatesResponseBodyModuleNext) SetLastModifier(v string) *ListAppTemplatesResponseBodyModuleNext {
	s.LastModifier = &v
	return s
}

func (s *ListAppTemplatesResponseBodyModuleNext) SetLikeCount(v int32) *ListAppTemplatesResponseBodyModuleNext {
	s.LikeCount = &v
	return s
}

func (s *ListAppTemplatesResponseBodyModuleNext) SetLiked(v bool) *ListAppTemplatesResponseBodyModuleNext {
	s.Liked = &v
	return s
}

func (s *ListAppTemplatesResponseBodyModuleNext) SetMetadata(v string) *ListAppTemplatesResponseBodyModuleNext {
	s.Metadata = &v
	return s
}

func (s *ListAppTemplatesResponseBodyModuleNext) SetPreviewUrl(v string) *ListAppTemplatesResponseBodyModuleNext {
	s.PreviewUrl = &v
	return s
}

func (s *ListAppTemplatesResponseBodyModuleNext) SetProductVersion(v string) *ListAppTemplatesResponseBodyModuleNext {
	s.ProductVersion = &v
	return s
}

func (s *ListAppTemplatesResponseBodyModuleNext) SetProductVersionName(v string) *ListAppTemplatesResponseBodyModuleNext {
	s.ProductVersionName = &v
	return s
}

func (s *ListAppTemplatesResponseBodyModuleNext) SetShareCount(v int32) *ListAppTemplatesResponseBodyModuleNext {
	s.ShareCount = &v
	return s
}

func (s *ListAppTemplatesResponseBodyModuleNext) SetStatus(v string) *ListAppTemplatesResponseBodyModuleNext {
	s.Status = &v
	return s
}

func (s *ListAppTemplatesResponseBodyModuleNext) SetTemplateId(v string) *ListAppTemplatesResponseBodyModuleNext {
	s.TemplateId = &v
	return s
}

func (s *ListAppTemplatesResponseBodyModuleNext) SetTemplateName(v string) *ListAppTemplatesResponseBodyModuleNext {
	s.TemplateName = &v
	return s
}

func (s *ListAppTemplatesResponseBodyModuleNext) SetThumbnailUrl(v string) *ListAppTemplatesResponseBodyModuleNext {
	s.ThumbnailUrl = &v
	return s
}

func (s *ListAppTemplatesResponseBodyModuleNext) SetViewCount(v int32) *ListAppTemplatesResponseBodyModuleNext {
	s.ViewCount = &v
	return s
}

func (s *ListAppTemplatesResponseBodyModuleNext) SetWeight(v int32) *ListAppTemplatesResponseBodyModuleNext {
	s.Weight = &v
	return s
}

func (s *ListAppTemplatesResponseBodyModuleNext) Validate() error {
	return dara.Validate(s)
}
