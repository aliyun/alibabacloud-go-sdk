// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTenantAppResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListTenantAppResponseBody
	GetCode() *string
	SetData(v []*ListTenantAppResponseBodyData) *ListTenantAppResponseBody
	GetData() []*ListTenantAppResponseBodyData
	SetHttpStatusCode(v int32) *ListTenantAppResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListTenantAppResponseBody
	GetMessage() *string
	SetPageNumber(v int32) *ListTenantAppResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListTenantAppResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListTenantAppResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListTenantAppResponseBody
	GetSuccess() *bool
	SetTotalCount(v int64) *ListTenantAppResponseBody
	GetTotalCount() *int64
}

type ListTenantAppResponseBody struct {
	// The business result code, returned as a string. The value is typically "200" when the request is successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The list of applications on the current page. Each element represents an application.
	Data []*ListTenantAppResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The HTTP status code field in the business response. This field may be empty. The actual transmission status is determined by the HTTP response status.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The description of the request processing result.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The actual page number of the query.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The actual number of entries per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request tracking ID. Provide this value when reporting issues.
	//
	// example:
	//
	// 11111111-2222-4333-8444-555555555555
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was processed successfully. Valid values:
	//
	// - true: Succeeded.
	//
	// - false: Failed.
	//
	// Refer to the corresponding field descriptions for specific business meanings.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of applications that match the filter conditions. This value is not equal to the length of the array on the current page.
	//
	// example:
	//
	// 1
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListTenantAppResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTenantAppResponseBody) GoString() string {
	return s.String()
}

func (s *ListTenantAppResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListTenantAppResponseBody) GetData() []*ListTenantAppResponseBodyData {
	return s.Data
}

func (s *ListTenantAppResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListTenantAppResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListTenantAppResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListTenantAppResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListTenantAppResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTenantAppResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListTenantAppResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListTenantAppResponseBody) SetCode(v string) *ListTenantAppResponseBody {
	s.Code = &v
	return s
}

func (s *ListTenantAppResponseBody) SetData(v []*ListTenantAppResponseBodyData) *ListTenantAppResponseBody {
	s.Data = v
	return s
}

func (s *ListTenantAppResponseBody) SetHttpStatusCode(v int32) *ListTenantAppResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListTenantAppResponseBody) SetMessage(v string) *ListTenantAppResponseBody {
	s.Message = &v
	return s
}

func (s *ListTenantAppResponseBody) SetPageNumber(v int32) *ListTenantAppResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListTenantAppResponseBody) SetPageSize(v int32) *ListTenantAppResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListTenantAppResponseBody) SetRequestId(v string) *ListTenantAppResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTenantAppResponseBody) SetSuccess(v bool) *ListTenantAppResponseBody {
	s.Success = &v
	return s
}

func (s *ListTenantAppResponseBody) SetTotalCount(v int64) *ListTenantAppResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListTenantAppResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListTenantAppResponseBodyData struct {
	// An internal field. We do not recommend that you use this field.
	AdminTag []*string `json:"AdminTag,omitempty" xml:"AdminTag,omitempty" type:"Repeated"`
	// An internal field. We do not recommend that you use this field.
	//
	// example:
	//
	// Recommended
	AppAdminTag *string `json:"AppAdminTag,omitempty" xml:"AppAdminTag,omitempty"`
	// The application registry identification information, returned as a string. This value can be used to match the name of an application registry entry.
	//
	// example:
	//
	// SampleEditor
	AppRegInfo *string `json:"AppRegInfo,omitempty" xml:"AppRegInfo,omitempty"`
	// An internal field. We do not recommend that you use this field.
	AppTag []*string `json:"AppTag,omitempty" xml:"AppTag,omitempty" type:"Repeated"`
	// The application UID string, which is a different identifier from the numeric Id field.
	//
	// example:
	//
	// app-demo-001
	AppUid *string `json:"AppUid,omitempty" xml:"AppUid,omitempty"`
	// The authorization dimension of the application.
	//
	// Valid values:
	//
	// - auth_type_user: Authorized by user.
	//
	// - auth_type_resource_group: Authorized by resource group.
	//
	// example:
	//
	// auth_type_user
	AuthType *string `json:"AuthType,omitempty" xml:"AuthType,omitempty"`
	// **[Deprecated]**
	AutoDeleteFlag *bool `json:"AutoDeleteFlag,omitempty" xml:"AutoDeleteFlag,omitempty"`
	// **[Deprecated]**
	AutoInstallFlag *bool `json:"AutoInstallFlag,omitempty" xml:"AutoInstallFlag,omitempty"`
	// **[Deprecated]**
	//
	// example:
	//
	// SILENCE_INSTALL
	AutoInstallType *string `json:"AutoInstallType,omitempty" xml:"AutoInstallType,omitempty"`
	// The new automatic installation scope policy. Use this field together with AuthType to determine the authorization dimension. Valid values:
	//
	// - 0: Automatically install for all visible users or resource groups.
	//
	// - 1: Automatically install for some visible users or resource groups.
	//
	// - 2: Disable automatic installation.
	//
	// - 99: Unknown policy.
	//
	// This field describes the configuration scope and does not indicate that the installation has been completed on the endpoint.
	//
	// example:
	//
	// 1
	AutoInstallmentType *int32 `json:"AutoInstallmentType,omitempty" xml:"AutoInstallmentType,omitempty"`
	// The display category ID of the application. The category ID is a dynamic identifier and is not a fixed enumeration.
	//
	// example:
	//
	// 1001
	CateId *int32 `json:"CateId,omitempty" xml:"CateId,omitempty"`
	// The display category name of the application.
	//
	// example:
	//
	// Office
	CateName *string `json:"CateName,omitempty" xml:"CateName,omitempty"`
	// **[Deprecated]**
	//
	// example:
	//
	// cluster-demo-001
	ClusterUid *string `json:"ClusterUid,omitempty" xml:"ClusterUid,omitempty"`
	// The description of the application.
	//
	// example:
	//
	// Used for document editing
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the application developer.
	//
	// example:
	//
	// Sample Software Company
	Developer *string `json:"Developer,omitempty" xml:"Developer,omitempty"`
	// **[Deprecated]**
	//
	// example:
	//
	// 0
	DiscountPrice *float64 `json:"DiscountPrice,omitempty" xml:"DiscountPrice,omitempty"`
	// The authorization distribution scope of the application. This field must be interpreted together with AuthType. An empty value does not necessarily mean that the application is not distributed.
	//
	// Valid values:
	//
	// - ALL: Distributed to all.
	//
	// - DESIGNATED: Distributed to a specified scope.
	//
	// - NOTDISTRO: Not distributed.
	//
	// - UNKNOWN: Unknown scope.
	//
	// The distribution target is determined by AuthType.
	//
	// example:
	//
	// ALL
	DistributeType *string `json:"DistributeType,omitempty" xml:"DistributeType,omitempty"`
	// The authorization end time of the application. The value is returned as a string with a time zone, in the format of date, the letter T, hours-minutes-seconds, 3-digit milliseconds, and a time zone offset without colons. The +0000 in the example indicates UTC. This field may be empty or not returned if no value is available.
	//
	// example:
	//
	// 2026-10-07T00:00:00.000+0000
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The extended information of the application, returned as a string. There is no unified fixed field structure.
	//
	// example:
	//
	// {}
	Extend *string `json:"Extend,omitempty" xml:"Extend,omitempty"`
	// The storage file name of the installation package, which may differ from the original file name.
	//
	// example:
	//
	// editor_1.2.3.exe
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// The storage path of the installation package. This value is not a directly accessible download URL.
	//
	// example:
	//
	// packages/example/editor.exe
	FilePath *string `json:"FilePath,omitempty" xml:"FilePath,omitempty"`
	// The original file name of the installation package.
	//
	// example:
	//
	// editor.exe
	FileRealName *string `json:"FileRealName,omitempty" xml:"FileRealName,omitempty"`
	// The creation time of the application record. The value is returned as a string with a time zone, in the format of date, the letter T, hours-minutes-seconds, 3-digit milliseconds, and a time zone offset without colons. The +0000 in the example indicates UTC. This field may be empty or not returned if no value is available.
	//
	// example:
	//
	// 2026-09-07T09:04:38.000+0000
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The last modification time of the application record. The value is returned as a string with a time zone, in the format of date, the letter T, hours-minutes-seconds, 3-digit milliseconds, and a time zone offset without colons. The +0000 in the example indicates UTC. This field may be empty or not returned if no value is available.
	//
	// example:
	//
	// 2026-09-07T09:04:38.000+0000
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// **[Deprecated]**
	HasCert *bool `json:"HasCert,omitempty" xml:"HasCert,omitempty"`
	// The icon URL of the application.
	//
	// example:
	//
	// https://example.com/icons/editor.png
	IconUrl *string `json:"IconUrl,omitempty" xml:"IconUrl,omitempty"`
	// The internal network icon URL of the application. Use this URL only when the corresponding network access conditions are met. The domain name in the example is for illustrative purposes only.
	//
	// example:
	//
	// https://example.com/icons/editor.png
	IconUrlInternal *string `json:"IconUrlInternal,omitempty" xml:"IconUrlInternal,omitempty"`
	// The numeric ID of the application, used as the identity of the application and to associate what to do next.
	//
	// example:
	//
	// 10001
	Id *int32 `json:"Id,omitempty" xml:"Id,omitempty"`
	// **[Deprecated]**
	Install *bool `json:"Install,omitempty" xml:"Install,omitempty"`
	// **[Deprecated]**
	//
	// example:
	//
	// 0
	InstallMode *int32 `json:"InstallMode,omitempty" xml:"InstallMode,omitempty"`
	// Specifies whether elevated privilege installation is configured. This does not indicate the administrator identity of the caller.
	//
	// Valid values:
	//
	// - true: Elevated privilege installation is configured.
	//
	// - false: Elevated privilege installation is not configured.
	//
	// example:
	//
	// false
	IsAdmin *bool `json:"IsAdmin,omitempty" xml:"IsAdmin,omitempty"`
	// **[Deprecated]**
	//
	// example:
	//
	// TRUE
	IsFree *string `json:"IsFree,omitempty" xml:"IsFree,omitempty"`
	// **[Deprecated]**
	IsGame *bool `json:"IsGame,omitempty" xml:"IsGame,omitempty"`
	// An internal field. We do not recommend that you use this field.
	//
	// example:
	//
	// 0
	IsWhiteList *int32 `json:"IsWhiteList,omitempty" xml:"IsWhiteList,omitempty"`
	// **[Deprecated]**
	//
	// example:
	//
	// app-demo-001
	ItemCode *string `json:"ItemCode,omitempty" xml:"ItemCode,omitempty"`
	// An internal field. We do not recommend that you use this field.
	//
	// example:
	//
	// Office
	Labels *string `json:"Labels,omitempty" xml:"Labels,omitempty"`
	// An internal field. We do not recommend that you use this field.
	//
	// example:
	//
	// MANUAL
	LicenseType *string `json:"LicenseType,omitempty" xml:"LicenseType,omitempty"`
	// The English name of the management category of the application.
	//
	// example:
	//
	// Office
	ManageCateEnName *string `json:"ManageCateEnName,omitempty" xml:"ManageCateEnName,omitempty"`
	// The management category ID of the application, which may differ from the display category CateId.
	//
	// example:
	//
	// 1001
	ManageCateId *int32 `json:"ManageCateId,omitempty" xml:"ManageCateId,omitempty"`
	// The management category name of the application.
	//
	// example:
	//
	// Office
	ManageCateName *string `json:"ManageCateName,omitempty" xml:"ManageCateName,omitempty"`
	// The name of the application.
	//
	// example:
	//
	// SampleEditor
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The application type.
	//
	// Valid values:
	//
	// - ClientBase: Client-based application.
	//
	// - WebBase: Web-based application.
	//
	// example:
	//
	// ClientBase
	OriginAppType *string `json:"OriginAppType,omitempty" xml:"OriginAppType,omitempty"`
	// **[Deprecated]**
	//
	// example:
	//
	// 0
	OriginalPrice *float64 `json:"OriginalPrice,omitempty" xml:"OriginalPrice,omitempty"`
	// The operating system type of the application.
	//
	// Valid values:
	//
	// - WINDOWS: Windows.
	//
	// - LINUX: Linux.
	//
	// - ANDROID: Android.
	//
	// - UNKNOWN: Unknown operating system.
	//
	// example:
	//
	// WINDOWS
	OsType *string `json:"OsType,omitempty" xml:"OsType,omitempty"`
	// **[Deprecated]**
	//
	// example:
	//
	// WINDOWS
	OwnerOs *string `json:"OwnerOs,omitempty" xml:"OwnerOs,omitempty"`
	// **[Deprecated]**
	//
	// example:
	//
	// 0
	PaymentType *int32 `json:"PaymentType,omitempty" xml:"PaymentType,omitempty"`
	// **[Deprecated]**
	//
	// example:
	//
	// 0
	Price *string `json:"Price,omitempty" xml:"Price,omitempty"`
	// **[Deprecated]**
	//
	// example:
	//
	// 0
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The publish time of the application. The value is returned as a string with a time zone, in the format of date, the letter T, hours-minutes-seconds, 3-digit milliseconds, and a time zone offset without colons. The +0000 in the example indicates UTC. This field may be empty or not returned if no value is available.
	//
	// example:
	//
	// 2026-09-07T09:04:39.000+0000
	PublishDate *string `json:"PublishDate,omitempty" xml:"PublishDate,omitempty"`
	// The customer scope of the application.
	//
	// Valid values:
	//
	// - ENT: Enterprise.
	//
	// - PER: Individual.
	//
	// - BOTH: Enterprise and individual.
	//
	// example:
	//
	// ENT
	PublishType *string `json:"PublishType,omitempty" xml:"PublishType,omitempty"`
	// **[Deprecated]**
	//
	// example:
	//
	// 0
	SandboxMode *int32 `json:"SandboxMode,omitempty" xml:"SandboxMode,omitempty"`
	// An internal field. We do not recommend that you use this field.
	//
	// example:
	//
	// Edit Document
	SearchTag *string `json:"SearchTag,omitempty" xml:"SearchTag,omitempty"`
	// **[Deprecated]**
	//
	// example:
	//
	// 1
	SilenceDeleteFlag *int32 `json:"SilenceDeleteFlag,omitempty" xml:"SilenceDeleteFlag,omitempty"`
	// **[Deprecated]**
	//
	// example:
	//
	// /uninstall /quiet
	SilenceDeleteParam *string `json:"SilenceDeleteParam,omitempty" xml:"SilenceDeleteParam,omitempty"`
	// Indicates whether silent installtion is supported. Valid values:
	//
	// - 0: Not supported.
	//
	// - 1: Supported.
	//
	// This field indicates the application capability and does not represent the actual installation execute result.
	//
	// example:
	//
	// 1
	SilenceFlag *int32 `json:"SilenceFlag,omitempty" xml:"SilenceFlag,omitempty"`
	// The silent installtion parameters, used by the corresponding installation flow.
	//
	// example:
	//
	// /quiet
	SilenceParam *string `json:"SilenceParam,omitempty" xml:"SilenceParam,omitempty"`
	// The size of the installation package.
	//
	// example:
	//
	// 10485760
	Size *int64 `json:"Size,omitempty" xml:"Size,omitempty"`
	// The application source. Valid values:
	//
	// - MARKET: Alibaba Cloud Marketplace application.
	//
	// - TENANT: Tenant-uploaded application.
	//
	// - UNKNOWN: Unknown source.
	//
	// example:
	//
	// TENANT
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// The start time of the application authorization. The value is returned as a string with time zone information, in the format of date, the letter T, hours-minutes-seconds, 3-digit milliseconds, and a time zone offset without colons. The +0000 in the example indicates UTC. This field may be empty or not returned if no value is available.
	//
	// example:
	//
	// 2026-09-07T00:00:00.000+0000
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The application status. This field does not represent the installation status on the endpoint. Valid values:
	//
	// - NORMAL: Normal.
	//
	// - DELETE: Deleted.
	//
	// - UNCHECK: Not reviewed or not verified.
	//
	// - DISABLE: All versions are unavailable.
	//
	// - UNKNOWN: Unknown status.
	//
	// **The following historical statuses from the sandbox packaging and publishing process are deprecated. Do not use them: UNPACKED (not packaged), TESTING (packaged, pending testing), UNPUBLISHED (testing completed, not published), PUBLISHED (published).**
	//
	// example:
	//
	// NORMAL
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// **[Deprecated]**
	//
	// example:
	//
	// sandbox
	SubAppType *string `json:"SubAppType,omitempty" xml:"SubAppType,omitempty"`
	// The secondary source of the application. Valid values:
	//
	// - ALI_MARKET: Alibaba Cloud Marketplace.
	//
	// - ISV: Independent software vendor.
	//
	// - OPS: Operations channel.
	//
	// - UNKNOWN: Unknown source.
	//
	// example:
	//
	// ISV
	SubSourceType *string `json:"SubSourceType,omitempty" xml:"SubSourceType,omitempty"`
	// **[Deprecated]**
	//
	// example:
	//
	// 0
	SubscribeCount *int64 `json:"SubscribeCount,omitempty" xml:"SubscribeCount,omitempty"`
	// The account ID of the application supplier or uploader.
	//
	// example:
	//
	// 1234567890123456
	SupplierId *int64 `json:"SupplierId,omitempty" xml:"SupplierId,omitempty"`
	// An internal field. We do not recommend that you use this field.
	UserTag []*string `json:"UserTag,omitempty" xml:"UserTag,omitempty" type:"Repeated"`
	// The application version number.
	//
	// example:
	//
	// 1.2.3
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
	// The display name of the application version.
	//
	// example:
	//
	// 1.2.3 Release
	VersionName *string `json:"VersionName,omitempty" xml:"VersionName,omitempty"`
	// **[Deprecated]**
	//
	// example:
	//
	// editor_1.2.3.wam
	WamFileName *string `json:"WamFileName,omitempty" xml:"WamFileName,omitempty"`
	// **[Deprecated]**
	//
	// example:
	//
	// packages/example/editor.wam
	WamFilePath *string `json:"WamFilePath,omitempty" xml:"WamFilePath,omitempty"`
	// **[Deprecated]**
	//
	// example:
	//
	// editor.wam
	WamFileRealName *string `json:"WamFileRealName,omitempty" xml:"WamFileRealName,omitempty"`
	// **[Deprecated]**
	//
	// example:
	//
	// 1024
	WamFileSize *int64 `json:"WamFileSize,omitempty" xml:"WamFileSize,omitempty"`
}

func (s ListTenantAppResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListTenantAppResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListTenantAppResponseBodyData) GetAdminTag() []*string {
	return s.AdminTag
}

func (s *ListTenantAppResponseBodyData) GetAppAdminTag() *string {
	return s.AppAdminTag
}

func (s *ListTenantAppResponseBodyData) GetAppRegInfo() *string {
	return s.AppRegInfo
}

func (s *ListTenantAppResponseBodyData) GetAppTag() []*string {
	return s.AppTag
}

func (s *ListTenantAppResponseBodyData) GetAppUid() *string {
	return s.AppUid
}

func (s *ListTenantAppResponseBodyData) GetAuthType() *string {
	return s.AuthType
}

func (s *ListTenantAppResponseBodyData) GetAutoDeleteFlag() *bool {
	return s.AutoDeleteFlag
}

func (s *ListTenantAppResponseBodyData) GetAutoInstallFlag() *bool {
	return s.AutoInstallFlag
}

func (s *ListTenantAppResponseBodyData) GetAutoInstallType() *string {
	return s.AutoInstallType
}

func (s *ListTenantAppResponseBodyData) GetAutoInstallmentType() *int32 {
	return s.AutoInstallmentType
}

func (s *ListTenantAppResponseBodyData) GetCateId() *int32 {
	return s.CateId
}

func (s *ListTenantAppResponseBodyData) GetCateName() *string {
	return s.CateName
}

func (s *ListTenantAppResponseBodyData) GetClusterUid() *string {
	return s.ClusterUid
}

func (s *ListTenantAppResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *ListTenantAppResponseBodyData) GetDeveloper() *string {
	return s.Developer
}

func (s *ListTenantAppResponseBodyData) GetDiscountPrice() *float64 {
	return s.DiscountPrice
}

func (s *ListTenantAppResponseBodyData) GetDistributeType() *string {
	return s.DistributeType
}

func (s *ListTenantAppResponseBodyData) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *ListTenantAppResponseBodyData) GetExtend() *string {
	return s.Extend
}

func (s *ListTenantAppResponseBodyData) GetFileName() *string {
	return s.FileName
}

func (s *ListTenantAppResponseBodyData) GetFilePath() *string {
	return s.FilePath
}

func (s *ListTenantAppResponseBodyData) GetFileRealName() *string {
	return s.FileRealName
}

func (s *ListTenantAppResponseBodyData) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ListTenantAppResponseBodyData) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ListTenantAppResponseBodyData) GetHasCert() *bool {
	return s.HasCert
}

func (s *ListTenantAppResponseBodyData) GetIconUrl() *string {
	return s.IconUrl
}

func (s *ListTenantAppResponseBodyData) GetIconUrlInternal() *string {
	return s.IconUrlInternal
}

func (s *ListTenantAppResponseBodyData) GetId() *int32 {
	return s.Id
}

func (s *ListTenantAppResponseBodyData) GetInstall() *bool {
	return s.Install
}

func (s *ListTenantAppResponseBodyData) GetInstallMode() *int32 {
	return s.InstallMode
}

func (s *ListTenantAppResponseBodyData) GetIsAdmin() *bool {
	return s.IsAdmin
}

func (s *ListTenantAppResponseBodyData) GetIsFree() *string {
	return s.IsFree
}

func (s *ListTenantAppResponseBodyData) GetIsGame() *bool {
	return s.IsGame
}

func (s *ListTenantAppResponseBodyData) GetIsWhiteList() *int32 {
	return s.IsWhiteList
}

func (s *ListTenantAppResponseBodyData) GetItemCode() *string {
	return s.ItemCode
}

func (s *ListTenantAppResponseBodyData) GetLabels() *string {
	return s.Labels
}

func (s *ListTenantAppResponseBodyData) GetLicenseType() *string {
	return s.LicenseType
}

func (s *ListTenantAppResponseBodyData) GetManageCateEnName() *string {
	return s.ManageCateEnName
}

func (s *ListTenantAppResponseBodyData) GetManageCateId() *int32 {
	return s.ManageCateId
}

func (s *ListTenantAppResponseBodyData) GetManageCateName() *string {
	return s.ManageCateName
}

func (s *ListTenantAppResponseBodyData) GetName() *string {
	return s.Name
}

func (s *ListTenantAppResponseBodyData) GetOriginAppType() *string {
	return s.OriginAppType
}

func (s *ListTenantAppResponseBodyData) GetOriginalPrice() *float64 {
	return s.OriginalPrice
}

func (s *ListTenantAppResponseBodyData) GetOsType() *string {
	return s.OsType
}

func (s *ListTenantAppResponseBodyData) GetOwnerOs() *string {
	return s.OwnerOs
}

func (s *ListTenantAppResponseBodyData) GetPaymentType() *int32 {
	return s.PaymentType
}

func (s *ListTenantAppResponseBodyData) GetPrice() *string {
	return s.Price
}

func (s *ListTenantAppResponseBodyData) GetPriority() *int32 {
	return s.Priority
}

func (s *ListTenantAppResponseBodyData) GetPublishDate() *string {
	return s.PublishDate
}

func (s *ListTenantAppResponseBodyData) GetPublishType() *string {
	return s.PublishType
}

func (s *ListTenantAppResponseBodyData) GetSandboxMode() *int32 {
	return s.SandboxMode
}

func (s *ListTenantAppResponseBodyData) GetSearchTag() *string {
	return s.SearchTag
}

func (s *ListTenantAppResponseBodyData) GetSilenceDeleteFlag() *int32 {
	return s.SilenceDeleteFlag
}

func (s *ListTenantAppResponseBodyData) GetSilenceDeleteParam() *string {
	return s.SilenceDeleteParam
}

func (s *ListTenantAppResponseBodyData) GetSilenceFlag() *int32 {
	return s.SilenceFlag
}

func (s *ListTenantAppResponseBodyData) GetSilenceParam() *string {
	return s.SilenceParam
}

func (s *ListTenantAppResponseBodyData) GetSize() *int64 {
	return s.Size
}

func (s *ListTenantAppResponseBodyData) GetSourceType() *string {
	return s.SourceType
}

func (s *ListTenantAppResponseBodyData) GetStartTime() *string {
	return s.StartTime
}

func (s *ListTenantAppResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *ListTenantAppResponseBodyData) GetSubAppType() *string {
	return s.SubAppType
}

func (s *ListTenantAppResponseBodyData) GetSubSourceType() *string {
	return s.SubSourceType
}

func (s *ListTenantAppResponseBodyData) GetSubscribeCount() *int64 {
	return s.SubscribeCount
}

func (s *ListTenantAppResponseBodyData) GetSupplierId() *int64 {
	return s.SupplierId
}

func (s *ListTenantAppResponseBodyData) GetUserTag() []*string {
	return s.UserTag
}

func (s *ListTenantAppResponseBodyData) GetVersion() *string {
	return s.Version
}

func (s *ListTenantAppResponseBodyData) GetVersionName() *string {
	return s.VersionName
}

func (s *ListTenantAppResponseBodyData) GetWamFileName() *string {
	return s.WamFileName
}

func (s *ListTenantAppResponseBodyData) GetWamFilePath() *string {
	return s.WamFilePath
}

func (s *ListTenantAppResponseBodyData) GetWamFileRealName() *string {
	return s.WamFileRealName
}

func (s *ListTenantAppResponseBodyData) GetWamFileSize() *int64 {
	return s.WamFileSize
}

func (s *ListTenantAppResponseBodyData) SetAdminTag(v []*string) *ListTenantAppResponseBodyData {
	s.AdminTag = v
	return s
}

func (s *ListTenantAppResponseBodyData) SetAppAdminTag(v string) *ListTenantAppResponseBodyData {
	s.AppAdminTag = &v
	return s
}

func (s *ListTenantAppResponseBodyData) SetAppRegInfo(v string) *ListTenantAppResponseBodyData {
	s.AppRegInfo = &v
	return s
}

func (s *ListTenantAppResponseBodyData) SetAppTag(v []*string) *ListTenantAppResponseBodyData {
	s.AppTag = v
	return s
}

func (s *ListTenantAppResponseBodyData) SetAppUid(v string) *ListTenantAppResponseBodyData {
	s.AppUid = &v
	return s
}

func (s *ListTenantAppResponseBodyData) SetAuthType(v string) *ListTenantAppResponseBodyData {
	s.AuthType = &v
	return s
}

func (s *ListTenantAppResponseBodyData) SetAutoDeleteFlag(v bool) *ListTenantAppResponseBodyData {
	s.AutoDeleteFlag = &v
	return s
}

func (s *ListTenantAppResponseBodyData) SetAutoInstallFlag(v bool) *ListTenantAppResponseBodyData {
	s.AutoInstallFlag = &v
	return s
}

func (s *ListTenantAppResponseBodyData) SetAutoInstallType(v string) *ListTenantAppResponseBodyData {
	s.AutoInstallType = &v
	return s
}

func (s *ListTenantAppResponseBodyData) SetAutoInstallmentType(v int32) *ListTenantAppResponseBodyData {
	s.AutoInstallmentType = &v
	return s
}

func (s *ListTenantAppResponseBodyData) SetCateId(v int32) *ListTenantAppResponseBodyData {
	s.CateId = &v
	return s
}

func (s *ListTenantAppResponseBodyData) SetCateName(v string) *ListTenantAppResponseBodyData {
	s.CateName = &v
	return s
}

func (s *ListTenantAppResponseBodyData) SetClusterUid(v string) *ListTenantAppResponseBodyData {
	s.ClusterUid = &v
	return s
}

func (s *ListTenantAppResponseBodyData) SetDescription(v string) *ListTenantAppResponseBodyData {
	s.Description = &v
	return s
}

func (s *ListTenantAppResponseBodyData) SetDeveloper(v string) *ListTenantAppResponseBodyData {
	s.Developer = &v
	return s
}

func (s *ListTenantAppResponseBodyData) SetDiscountPrice(v float64) *ListTenantAppResponseBodyData {
	s.DiscountPrice = &v
	return s
}

func (s *ListTenantAppResponseBodyData) SetDistributeType(v string) *ListTenantAppResponseBodyData {
	s.DistributeType = &v
	return s
}

func (s *ListTenantAppResponseBodyData) SetExpireTime(v string) *ListTenantAppResponseBodyData {
	s.ExpireTime = &v
	return s
}

func (s *ListTenantAppResponseBodyData) SetExtend(v string) *ListTenantAppResponseBodyData {
	s.Extend = &v
	return s
}

func (s *ListTenantAppResponseBodyData) SetFileName(v string) *ListTenantAppResponseBodyData {
	s.FileName = &v
	return s
}

func (s *ListTenantAppResponseBodyData) SetFilePath(v string) *ListTenantAppResponseBodyData {
	s.FilePath = &v
	return s
}

func (s *ListTenantAppResponseBodyData) SetFileRealName(v string) *ListTenantAppResponseBodyData {
	s.FileRealName = &v
	return s
}

func (s *ListTenantAppResponseBodyData) SetGmtCreate(v string) *ListTenantAppResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *ListTenantAppResponseBodyData) SetGmtModified(v string) *ListTenantAppResponseBodyData {
	s.GmtModified = &v
	return s
}

func (s *ListTenantAppResponseBodyData) SetHasCert(v bool) *ListTenantAppResponseBodyData {
	s.HasCert = &v
	return s
}

func (s *ListTenantAppResponseBodyData) SetIconUrl(v string) *ListTenantAppResponseBodyData {
	s.IconUrl = &v
	return s
}

func (s *ListTenantAppResponseBodyData) SetIconUrlInternal(v string) *ListTenantAppResponseBodyData {
	s.IconUrlInternal = &v
	return s
}

func (s *ListTenantAppResponseBodyData) SetId(v int32) *ListTenantAppResponseBodyData {
	s.Id = &v
	return s
}

func (s *ListTenantAppResponseBodyData) SetInstall(v bool) *ListTenantAppResponseBodyData {
	s.Install = &v
	return s
}

func (s *ListTenantAppResponseBodyData) SetInstallMode(v int32) *ListTenantAppResponseBodyData {
	s.InstallMode = &v
	return s
}

func (s *ListTenantAppResponseBodyData) SetIsAdmin(v bool) *ListTenantAppResponseBodyData {
	s.IsAdmin = &v
	return s
}

func (s *ListTenantAppResponseBodyData) SetIsFree(v string) *ListTenantAppResponseBodyData {
	s.IsFree = &v
	return s
}

func (s *ListTenantAppResponseBodyData) SetIsGame(v bool) *ListTenantAppResponseBodyData {
	s.IsGame = &v
	return s
}

func (s *ListTenantAppResponseBodyData) SetIsWhiteList(v int32) *ListTenantAppResponseBodyData {
	s.IsWhiteList = &v
	return s
}

func (s *ListTenantAppResponseBodyData) SetItemCode(v string) *ListTenantAppResponseBodyData {
	s.ItemCode = &v
	return s
}

func (s *ListTenantAppResponseBodyData) SetLabels(v string) *ListTenantAppResponseBodyData {
	s.Labels = &v
	return s
}

func (s *ListTenantAppResponseBodyData) SetLicenseType(v string) *ListTenantAppResponseBodyData {
	s.LicenseType = &v
	return s
}

func (s *ListTenantAppResponseBodyData) SetManageCateEnName(v string) *ListTenantAppResponseBodyData {
	s.ManageCateEnName = &v
	return s
}

func (s *ListTenantAppResponseBodyData) SetManageCateId(v int32) *ListTenantAppResponseBodyData {
	s.ManageCateId = &v
	return s
}

func (s *ListTenantAppResponseBodyData) SetManageCateName(v string) *ListTenantAppResponseBodyData {
	s.ManageCateName = &v
	return s
}

func (s *ListTenantAppResponseBodyData) SetName(v string) *ListTenantAppResponseBodyData {
	s.Name = &v
	return s
}

func (s *ListTenantAppResponseBodyData) SetOriginAppType(v string) *ListTenantAppResponseBodyData {
	s.OriginAppType = &v
	return s
}

func (s *ListTenantAppResponseBodyData) SetOriginalPrice(v float64) *ListTenantAppResponseBodyData {
	s.OriginalPrice = &v
	return s
}

func (s *ListTenantAppResponseBodyData) SetOsType(v string) *ListTenantAppResponseBodyData {
	s.OsType = &v
	return s
}

func (s *ListTenantAppResponseBodyData) SetOwnerOs(v string) *ListTenantAppResponseBodyData {
	s.OwnerOs = &v
	return s
}

func (s *ListTenantAppResponseBodyData) SetPaymentType(v int32) *ListTenantAppResponseBodyData {
	s.PaymentType = &v
	return s
}

func (s *ListTenantAppResponseBodyData) SetPrice(v string) *ListTenantAppResponseBodyData {
	s.Price = &v
	return s
}

func (s *ListTenantAppResponseBodyData) SetPriority(v int32) *ListTenantAppResponseBodyData {
	s.Priority = &v
	return s
}

func (s *ListTenantAppResponseBodyData) SetPublishDate(v string) *ListTenantAppResponseBodyData {
	s.PublishDate = &v
	return s
}

func (s *ListTenantAppResponseBodyData) SetPublishType(v string) *ListTenantAppResponseBodyData {
	s.PublishType = &v
	return s
}

func (s *ListTenantAppResponseBodyData) SetSandboxMode(v int32) *ListTenantAppResponseBodyData {
	s.SandboxMode = &v
	return s
}

func (s *ListTenantAppResponseBodyData) SetSearchTag(v string) *ListTenantAppResponseBodyData {
	s.SearchTag = &v
	return s
}

func (s *ListTenantAppResponseBodyData) SetSilenceDeleteFlag(v int32) *ListTenantAppResponseBodyData {
	s.SilenceDeleteFlag = &v
	return s
}

func (s *ListTenantAppResponseBodyData) SetSilenceDeleteParam(v string) *ListTenantAppResponseBodyData {
	s.SilenceDeleteParam = &v
	return s
}

func (s *ListTenantAppResponseBodyData) SetSilenceFlag(v int32) *ListTenantAppResponseBodyData {
	s.SilenceFlag = &v
	return s
}

func (s *ListTenantAppResponseBodyData) SetSilenceParam(v string) *ListTenantAppResponseBodyData {
	s.SilenceParam = &v
	return s
}

func (s *ListTenantAppResponseBodyData) SetSize(v int64) *ListTenantAppResponseBodyData {
	s.Size = &v
	return s
}

func (s *ListTenantAppResponseBodyData) SetSourceType(v string) *ListTenantAppResponseBodyData {
	s.SourceType = &v
	return s
}

func (s *ListTenantAppResponseBodyData) SetStartTime(v string) *ListTenantAppResponseBodyData {
	s.StartTime = &v
	return s
}

func (s *ListTenantAppResponseBodyData) SetStatus(v string) *ListTenantAppResponseBodyData {
	s.Status = &v
	return s
}

func (s *ListTenantAppResponseBodyData) SetSubAppType(v string) *ListTenantAppResponseBodyData {
	s.SubAppType = &v
	return s
}

func (s *ListTenantAppResponseBodyData) SetSubSourceType(v string) *ListTenantAppResponseBodyData {
	s.SubSourceType = &v
	return s
}

func (s *ListTenantAppResponseBodyData) SetSubscribeCount(v int64) *ListTenantAppResponseBodyData {
	s.SubscribeCount = &v
	return s
}

func (s *ListTenantAppResponseBodyData) SetSupplierId(v int64) *ListTenantAppResponseBodyData {
	s.SupplierId = &v
	return s
}

func (s *ListTenantAppResponseBodyData) SetUserTag(v []*string) *ListTenantAppResponseBodyData {
	s.UserTag = v
	return s
}

func (s *ListTenantAppResponseBodyData) SetVersion(v string) *ListTenantAppResponseBodyData {
	s.Version = &v
	return s
}

func (s *ListTenantAppResponseBodyData) SetVersionName(v string) *ListTenantAppResponseBodyData {
	s.VersionName = &v
	return s
}

func (s *ListTenantAppResponseBodyData) SetWamFileName(v string) *ListTenantAppResponseBodyData {
	s.WamFileName = &v
	return s
}

func (s *ListTenantAppResponseBodyData) SetWamFilePath(v string) *ListTenantAppResponseBodyData {
	s.WamFilePath = &v
	return s
}

func (s *ListTenantAppResponseBodyData) SetWamFileRealName(v string) *ListTenantAppResponseBodyData {
	s.WamFileRealName = &v
	return s
}

func (s *ListTenantAppResponseBodyData) SetWamFileSize(v int64) *ListTenantAppResponseBodyData {
	s.WamFileSize = &v
	return s
}

func (s *ListTenantAppResponseBodyData) Validate() error {
	return dara.Validate(s)
}
