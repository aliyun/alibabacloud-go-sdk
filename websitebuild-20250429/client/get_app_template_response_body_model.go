// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAppTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *GetAppTemplateResponseBody
	GetAccessDeniedDetail() *string
	SetAllowRetry(v bool) *GetAppTemplateResponseBody
	GetAllowRetry() *bool
	SetAppName(v string) *GetAppTemplateResponseBody
	GetAppName() *string
	SetDynamicCode(v string) *GetAppTemplateResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *GetAppTemplateResponseBody
	GetDynamicMessage() *string
	SetErrorArgs(v []interface{}) *GetAppTemplateResponseBody
	GetErrorArgs() []interface{}
	SetModule(v *GetAppTemplateResponseBodyModule) *GetAppTemplateResponseBody
	GetModule() *GetAppTemplateResponseBodyModule
	SetRequestId(v string) *GetAppTemplateResponseBody
	GetRequestId() *string
	SetRootErrorCode(v string) *GetAppTemplateResponseBody
	GetRootErrorCode() *string
	SetRootErrorMsg(v string) *GetAppTemplateResponseBody
	GetRootErrorMsg() *string
	SetSynchro(v bool) *GetAppTemplateResponseBody
	GetSynchro() *bool
}

type GetAppTemplateResponseBody struct {
	// permission denied information
	//
	// example:
	//
	// {}
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// is retry allowed
	//
	// example:
	//
	// False
	AllowRetry *bool `json:"AllowRetry,omitempty" xml:"AllowRetry,omitempty"`
	// App Name.
	//
	// example:
	//
	// ish-intelligence-store-platform-admin-web
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// dynamic error Code
	//
	// example:
	//
	// ERROR-oo1
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// dynamic error message, used to replace the `%s` placeholder in the **ErrMessage*	- error message.
	//
	// > If **ErrMessage*	- returns **The Value of Input Parameter %s is not valid*	- and **DynamicMessage*	- returns **DtsJobId**, it indicates that the request parameter **DtsJobId*	- is invalid.
	//
	// example:
	//
	// SYSTEM_ERROR
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// returned error parameters
	ErrorArgs []interface{} `json:"ErrorArgs,omitempty" xml:"ErrorArgs,omitempty" type:"Repeated"`
	// whether deletion succeeded
	Module *GetAppTemplateResponseBodyModule `json:"Module,omitempty" xml:"Module,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 6C6B99AC-39EC-5350-874C-204128C905E6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Error code
	//
	// example:
	//
	// SYSTEM.ERROR
	RootErrorCode *string `json:"RootErrorCode,omitempty" xml:"RootErrorCode,omitempty"`
	// abnormal message
	//
	// example:
	//
	// 系统异常
	RootErrorMsg *string `json:"RootErrorMsg,omitempty" xml:"RootErrorMsg,omitempty"`
	// Reserved parameter.
	//
	// example:
	//
	// True
	Synchro *bool `json:"Synchro,omitempty" xml:"Synchro,omitempty"`
}

func (s GetAppTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAppTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *GetAppTemplateResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *GetAppTemplateResponseBody) GetAllowRetry() *bool {
	return s.AllowRetry
}

func (s *GetAppTemplateResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *GetAppTemplateResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *GetAppTemplateResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *GetAppTemplateResponseBody) GetErrorArgs() []interface{} {
	return s.ErrorArgs
}

func (s *GetAppTemplateResponseBody) GetModule() *GetAppTemplateResponseBodyModule {
	return s.Module
}

func (s *GetAppTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAppTemplateResponseBody) GetRootErrorCode() *string {
	return s.RootErrorCode
}

func (s *GetAppTemplateResponseBody) GetRootErrorMsg() *string {
	return s.RootErrorMsg
}

func (s *GetAppTemplateResponseBody) GetSynchro() *bool {
	return s.Synchro
}

func (s *GetAppTemplateResponseBody) SetAccessDeniedDetail(v string) *GetAppTemplateResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *GetAppTemplateResponseBody) SetAllowRetry(v bool) *GetAppTemplateResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *GetAppTemplateResponseBody) SetAppName(v string) *GetAppTemplateResponseBody {
	s.AppName = &v
	return s
}

func (s *GetAppTemplateResponseBody) SetDynamicCode(v string) *GetAppTemplateResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *GetAppTemplateResponseBody) SetDynamicMessage(v string) *GetAppTemplateResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *GetAppTemplateResponseBody) SetErrorArgs(v []interface{}) *GetAppTemplateResponseBody {
	s.ErrorArgs = v
	return s
}

func (s *GetAppTemplateResponseBody) SetModule(v *GetAppTemplateResponseBodyModule) *GetAppTemplateResponseBody {
	s.Module = v
	return s
}

func (s *GetAppTemplateResponseBody) SetRequestId(v string) *GetAppTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAppTemplateResponseBody) SetRootErrorCode(v string) *GetAppTemplateResponseBody {
	s.RootErrorCode = &v
	return s
}

func (s *GetAppTemplateResponseBody) SetRootErrorMsg(v string) *GetAppTemplateResponseBody {
	s.RootErrorMsg = &v
	return s
}

func (s *GetAppTemplateResponseBody) SetSynchro(v bool) *GetAppTemplateResponseBody {
	s.Synchro = &v
	return s
}

func (s *GetAppTemplateResponseBody) Validate() error {
	if s.Module != nil {
		if err := s.Module.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAppTemplateResponseBodyModule struct {
	// Application type
	//
	// example:
	//
	// TRACE
	AppType *string `json:"AppType,omitempty" xml:"AppType,omitempty"`
	// Business ID
	//
	// example:
	//
	// WD20250703155602000001
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// color
	//
	// example:
	//
	// red
	ColorScheme *string `json:"ColorScheme,omitempty" xml:"ColorScheme,omitempty"`
	// Color name
	//
	// example:
	//
	// red
	ColorSchemeName *string `json:"ColorSchemeName,omitempty" xml:"ColorSchemeName,omitempty"`
	// The actual data volume of the original table copied during DMS lockless change.
	//
	// example:
	//
	// 100
	CopyCount *int32 `json:"CopyCount,omitempty" xml:"CopyCount,omitempty"`
	// User ID of the creator.
	//
	// example:
	//
	// 208614160512124381
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// Creation UTC time, in ISO 8601 date format.
	//
	// Use the UTC time format: yyyy-MM-ddTHH:mm:ss.SSSZ
	//
	// example:
	//
	// 2025-07-04T01:40:25Z
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// Modify time
	//
	// Use the UTC time format: yyyy-MM-ddTHH:mm:ss.SSSZ
	//
	// example:
	//
	// 2025-06-05T11:16:57.785Z
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// Primary key
	//
	// example:
	//
	// 16257
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// Industry. Default value: **common*	- (general industry). Valid values:
	//
	// - **microVideo**: Short video industry.
	//
	// - **common**: General industry.
	//
	// example:
	//
	// 通用
	Industry *string `json:"Industry,omitempty" xml:"Industry,omitempty"`
	// Industry name.
	//
	// example:
	//
	// indextry
	IndustryName *string `json:"IndustryName,omitempty" xml:"IndustryName,omitempty"`
	// Last modifier
	//
	// example:
	//
	// admin
	LastModifier *string `json:"LastModifier,omitempty" xml:"LastModifier,omitempty"`
	// Like count.
	//
	// example:
	//
	// 100
	LikeCount *int32 `json:"LikeCount,omitempty" xml:"LikeCount,omitempty"`
	// Whether the current user has liked it
	Liked *bool `json:"Liked,omitempty" xml:"Liked,omitempty"`
	// <props="china">
	//
	// Map of metadata for text segmentation.
	//
	// > In document-search-type knowledge bases, the `file_path` field in the metadata Map is meaningless. Do not use it in your business code.
	//
	// > When retrieving a document-search-type knowledge base, if a segment contains an image, the image URL will be exposed through the `image_url` field in the metadata Map, along with a time-to-live (TTL).
	//
	//
	// <props="intl">
	//
	// Map of metadata for text segmentation.
	//
	// > In document-search-type knowledge bases, the `file_path` field in the metadata Map is meaningless. Do not use it in your business code.
	//
	// > When retrieving a document-search-type knowledge base, if a segment contains an image, the image URL will be exposed through the `image_url` field in the metadata Map, along with a time-to-live (TTL).
	//
	// example:
	//
	// {\\"operations\\":[{\\"operator\\":\\"replace\\",\\"oldMetaName\\":\\"character2_add\\",\\"newMetaType\\":\\"varchar(2)\\"}]}
	Metadata *string `json:"Metadata,omitempty" xml:"Metadata,omitempty"`
	// Preview URL
	//
	// example:
	//
	// https://preview-lyj.aliyuncs.com/preview/4825a3849c2e49e1b48804c7f90b766f?subSceneIds=807648
	PreviewUrl *string `json:"PreviewUrl,omitempty" xml:"PreviewUrl,omitempty"`
	// Edition. Valid values:
	//
	// - **BasicVersion**: Basic Edition.
	//
	// - **EnterpriseVersion**: Enterprise Edition.
	//
	// > This parameter is required only when ProductForm is set to IntegrationForm.
	//
	// example:
	//
	// V2
	ProductVersion *string `json:"ProductVersion,omitempty" xml:"ProductVersion,omitempty"`
	// Edition Name.
	//
	// example:
	//
	// 1.0.2
	ProductVersionName *string `json:"ProductVersionName,omitempty" xml:"ProductVersionName,omitempty"`
	// Share count
	//
	// example:
	//
	// 100
	ShareCount *int32 `json:"ShareCount,omitempty" xml:"ShareCount,omitempty"`
	// trial, draft, live, refunded, expired, released
	//
	// example:
	//
	// NORMAL
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Template ID
	//
	// example:
	//
	// 7cc17da1-b670-4be7-a6b4-0b3cdf7bf5f7
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// Template Name
	//
	// example:
	//
	// siemProtect_All_
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// application thumbnail
	//
	// example:
	//
	// placeHolder
	ThumbnailUrl *string `json:"ThumbnailUrl,omitempty" xml:"ThumbnailUrl,omitempty"`
	// View count.
	//
	// example:
	//
	// 100
	ViewCount *int32 `json:"ViewCount,omitempty" xml:"ViewCount,omitempty"`
	// Hot word weight, an integer in the range [1, 5].
	//
	// Common value: 4.
	//
	// If the effect is not obvious, you can moderately increase the weight. However, if the weight is too high, it may cause negative effects and reduce the accuracy of detecting other words.
	//
	// example:
	//
	// 100
	Weight *int32 `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s GetAppTemplateResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s GetAppTemplateResponseBodyModule) GoString() string {
	return s.String()
}

func (s *GetAppTemplateResponseBodyModule) GetAppType() *string {
	return s.AppType
}

func (s *GetAppTemplateResponseBodyModule) GetBizId() *string {
	return s.BizId
}

func (s *GetAppTemplateResponseBodyModule) GetColorScheme() *string {
	return s.ColorScheme
}

func (s *GetAppTemplateResponseBodyModule) GetColorSchemeName() *string {
	return s.ColorSchemeName
}

func (s *GetAppTemplateResponseBodyModule) GetCopyCount() *int32 {
	return s.CopyCount
}

func (s *GetAppTemplateResponseBodyModule) GetCreator() *string {
	return s.Creator
}

func (s *GetAppTemplateResponseBodyModule) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *GetAppTemplateResponseBodyModule) GetGmtModifiedTime() *string {
	return s.GmtModifiedTime
}

func (s *GetAppTemplateResponseBodyModule) GetId() *int64 {
	return s.Id
}

func (s *GetAppTemplateResponseBodyModule) GetIndustry() *string {
	return s.Industry
}

func (s *GetAppTemplateResponseBodyModule) GetIndustryName() *string {
	return s.IndustryName
}

func (s *GetAppTemplateResponseBodyModule) GetLastModifier() *string {
	return s.LastModifier
}

func (s *GetAppTemplateResponseBodyModule) GetLikeCount() *int32 {
	return s.LikeCount
}

func (s *GetAppTemplateResponseBodyModule) GetLiked() *bool {
	return s.Liked
}

func (s *GetAppTemplateResponseBodyModule) GetMetadata() *string {
	return s.Metadata
}

func (s *GetAppTemplateResponseBodyModule) GetPreviewUrl() *string {
	return s.PreviewUrl
}

func (s *GetAppTemplateResponseBodyModule) GetProductVersion() *string {
	return s.ProductVersion
}

func (s *GetAppTemplateResponseBodyModule) GetProductVersionName() *string {
	return s.ProductVersionName
}

func (s *GetAppTemplateResponseBodyModule) GetShareCount() *int32 {
	return s.ShareCount
}

func (s *GetAppTemplateResponseBodyModule) GetStatus() *string {
	return s.Status
}

func (s *GetAppTemplateResponseBodyModule) GetTemplateId() *string {
	return s.TemplateId
}

func (s *GetAppTemplateResponseBodyModule) GetTemplateName() *string {
	return s.TemplateName
}

func (s *GetAppTemplateResponseBodyModule) GetThumbnailUrl() *string {
	return s.ThumbnailUrl
}

func (s *GetAppTemplateResponseBodyModule) GetViewCount() *int32 {
	return s.ViewCount
}

func (s *GetAppTemplateResponseBodyModule) GetWeight() *int32 {
	return s.Weight
}

func (s *GetAppTemplateResponseBodyModule) SetAppType(v string) *GetAppTemplateResponseBodyModule {
	s.AppType = &v
	return s
}

func (s *GetAppTemplateResponseBodyModule) SetBizId(v string) *GetAppTemplateResponseBodyModule {
	s.BizId = &v
	return s
}

func (s *GetAppTemplateResponseBodyModule) SetColorScheme(v string) *GetAppTemplateResponseBodyModule {
	s.ColorScheme = &v
	return s
}

func (s *GetAppTemplateResponseBodyModule) SetColorSchemeName(v string) *GetAppTemplateResponseBodyModule {
	s.ColorSchemeName = &v
	return s
}

func (s *GetAppTemplateResponseBodyModule) SetCopyCount(v int32) *GetAppTemplateResponseBodyModule {
	s.CopyCount = &v
	return s
}

func (s *GetAppTemplateResponseBodyModule) SetCreator(v string) *GetAppTemplateResponseBodyModule {
	s.Creator = &v
	return s
}

func (s *GetAppTemplateResponseBodyModule) SetGmtCreateTime(v string) *GetAppTemplateResponseBodyModule {
	s.GmtCreateTime = &v
	return s
}

func (s *GetAppTemplateResponseBodyModule) SetGmtModifiedTime(v string) *GetAppTemplateResponseBodyModule {
	s.GmtModifiedTime = &v
	return s
}

func (s *GetAppTemplateResponseBodyModule) SetId(v int64) *GetAppTemplateResponseBodyModule {
	s.Id = &v
	return s
}

func (s *GetAppTemplateResponseBodyModule) SetIndustry(v string) *GetAppTemplateResponseBodyModule {
	s.Industry = &v
	return s
}

func (s *GetAppTemplateResponseBodyModule) SetIndustryName(v string) *GetAppTemplateResponseBodyModule {
	s.IndustryName = &v
	return s
}

func (s *GetAppTemplateResponseBodyModule) SetLastModifier(v string) *GetAppTemplateResponseBodyModule {
	s.LastModifier = &v
	return s
}

func (s *GetAppTemplateResponseBodyModule) SetLikeCount(v int32) *GetAppTemplateResponseBodyModule {
	s.LikeCount = &v
	return s
}

func (s *GetAppTemplateResponseBodyModule) SetLiked(v bool) *GetAppTemplateResponseBodyModule {
	s.Liked = &v
	return s
}

func (s *GetAppTemplateResponseBodyModule) SetMetadata(v string) *GetAppTemplateResponseBodyModule {
	s.Metadata = &v
	return s
}

func (s *GetAppTemplateResponseBodyModule) SetPreviewUrl(v string) *GetAppTemplateResponseBodyModule {
	s.PreviewUrl = &v
	return s
}

func (s *GetAppTemplateResponseBodyModule) SetProductVersion(v string) *GetAppTemplateResponseBodyModule {
	s.ProductVersion = &v
	return s
}

func (s *GetAppTemplateResponseBodyModule) SetProductVersionName(v string) *GetAppTemplateResponseBodyModule {
	s.ProductVersionName = &v
	return s
}

func (s *GetAppTemplateResponseBodyModule) SetShareCount(v int32) *GetAppTemplateResponseBodyModule {
	s.ShareCount = &v
	return s
}

func (s *GetAppTemplateResponseBodyModule) SetStatus(v string) *GetAppTemplateResponseBodyModule {
	s.Status = &v
	return s
}

func (s *GetAppTemplateResponseBodyModule) SetTemplateId(v string) *GetAppTemplateResponseBodyModule {
	s.TemplateId = &v
	return s
}

func (s *GetAppTemplateResponseBodyModule) SetTemplateName(v string) *GetAppTemplateResponseBodyModule {
	s.TemplateName = &v
	return s
}

func (s *GetAppTemplateResponseBodyModule) SetThumbnailUrl(v string) *GetAppTemplateResponseBodyModule {
	s.ThumbnailUrl = &v
	return s
}

func (s *GetAppTemplateResponseBodyModule) SetViewCount(v int32) *GetAppTemplateResponseBodyModule {
	s.ViewCount = &v
	return s
}

func (s *GetAppTemplateResponseBodyModule) SetWeight(v int32) *GetAppTemplateResponseBodyModule {
	s.Weight = &v
	return s
}

func (s *GetAppTemplateResponseBodyModule) Validate() error {
	return dara.Validate(s)
}
