// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAppInstanceForPartnerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *GetAppInstanceForPartnerResponseBody
	GetAccessDeniedDetail() *string
	SetAllowRetry(v bool) *GetAppInstanceForPartnerResponseBody
	GetAllowRetry() *bool
	SetAppName(v string) *GetAppInstanceForPartnerResponseBody
	GetAppName() *string
	SetDynamicCode(v string) *GetAppInstanceForPartnerResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *GetAppInstanceForPartnerResponseBody
	GetDynamicMessage() *string
	SetErrorArgs(v []interface{}) *GetAppInstanceForPartnerResponseBody
	GetErrorArgs() []interface{}
	SetModule(v *GetAppInstanceForPartnerResponseBodyModule) *GetAppInstanceForPartnerResponseBody
	GetModule() *GetAppInstanceForPartnerResponseBodyModule
	SetRequestId(v string) *GetAppInstanceForPartnerResponseBody
	GetRequestId() *string
	SetRootErrorCode(v string) *GetAppInstanceForPartnerResponseBody
	GetRootErrorCode() *string
	SetRootErrorMsg(v string) *GetAppInstanceForPartnerResponseBody
	GetRootErrorMsg() *string
	SetSynchro(v bool) *GetAppInstanceForPartnerResponseBody
	GetSynchro() *bool
}

type GetAppInstanceForPartnerResponseBody struct {
	// Detailed reason why access was denied.
	//
	// example:
	//
	// {}
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// Indicates whether retry is allowed
	//
	// example:
	//
	// False
	AllowRetry *bool `json:"AllowRetry,omitempty" xml:"AllowRetry,omitempty"`
	// AppName.
	//
	// example:
	//
	// dewuApp
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// Dynamic error code.
	//
	// example:
	//
	// ERROR-oo1
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// Dynamic error message, used to replace the `%s` placeholder in the **ErrMessage*	- error message.
	//
	// > If **ErrMessage*	- returns **The Value of Input Parameter %s is not valid*	- and **DynamicMessage*	- returns **DtsJobId**, it indicates that the provided request parameter **DtsJobId*	- is invalid.
	//
	// example:
	//
	// https://check-result-file-sh.oss-cn-shanghai.aliyuncs.com/gl3d6l3817id8p1/gl3d6l3817id8p1.diff.zip?Expires=1750392068&OSSAccessKeyId=LTAI5tKUErVCETM4ev9SELNb&Signature=Bcj3eohy8nmlSQ7AAGdq7JZoLjM%3D
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// Returned error parameters
	ErrorArgs []interface{} `json:"ErrorArgs,omitempty" xml:"ErrorArgs,omitempty" type:"Repeated"`
	// Response data
	Module *GetAppInstanceForPartnerResponseBodyModule `json:"Module,omitempty" xml:"Module,omitempty" type:"Struct"`
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
	// Abnormal message
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

func (s GetAppInstanceForPartnerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAppInstanceForPartnerResponseBody) GoString() string {
	return s.String()
}

func (s *GetAppInstanceForPartnerResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *GetAppInstanceForPartnerResponseBody) GetAllowRetry() *bool {
	return s.AllowRetry
}

func (s *GetAppInstanceForPartnerResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *GetAppInstanceForPartnerResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *GetAppInstanceForPartnerResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *GetAppInstanceForPartnerResponseBody) GetErrorArgs() []interface{} {
	return s.ErrorArgs
}

func (s *GetAppInstanceForPartnerResponseBody) GetModule() *GetAppInstanceForPartnerResponseBodyModule {
	return s.Module
}

func (s *GetAppInstanceForPartnerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAppInstanceForPartnerResponseBody) GetRootErrorCode() *string {
	return s.RootErrorCode
}

func (s *GetAppInstanceForPartnerResponseBody) GetRootErrorMsg() *string {
	return s.RootErrorMsg
}

func (s *GetAppInstanceForPartnerResponseBody) GetSynchro() *bool {
	return s.Synchro
}

func (s *GetAppInstanceForPartnerResponseBody) SetAccessDeniedDetail(v string) *GetAppInstanceForPartnerResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBody) SetAllowRetry(v bool) *GetAppInstanceForPartnerResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBody) SetAppName(v string) *GetAppInstanceForPartnerResponseBody {
	s.AppName = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBody) SetDynamicCode(v string) *GetAppInstanceForPartnerResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBody) SetDynamicMessage(v string) *GetAppInstanceForPartnerResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBody) SetErrorArgs(v []interface{}) *GetAppInstanceForPartnerResponseBody {
	s.ErrorArgs = v
	return s
}

func (s *GetAppInstanceForPartnerResponseBody) SetModule(v *GetAppInstanceForPartnerResponseBodyModule) *GetAppInstanceForPartnerResponseBody {
	s.Module = v
	return s
}

func (s *GetAppInstanceForPartnerResponseBody) SetRequestId(v string) *GetAppInstanceForPartnerResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBody) SetRootErrorCode(v string) *GetAppInstanceForPartnerResponseBody {
	s.RootErrorCode = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBody) SetRootErrorMsg(v string) *GetAppInstanceForPartnerResponseBody {
	s.RootErrorMsg = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBody) SetSynchro(v bool) *GetAppInstanceForPartnerResponseBody {
	s.Synchro = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBody) Validate() error {
	if s.Module != nil {
		if err := s.Module.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAppInstanceForPartnerResponseBodyModule struct {
	// List of AI staff associated with the application
	AiStaffList []*GetAppInstanceForPartnerResponseBodyModuleAiStaffList `json:"AiStaffList,omitempty" xml:"AiStaffList,omitempty" type:"Repeated"`
	// Design order associated with the application
	AppDesignSpec *GetAppInstanceForPartnerResponseBodyModuleAppDesignSpec `json:"AppDesignSpec,omitempty" xml:"AppDesignSpec,omitempty" type:"Struct"`
	// Application operation address
	AppOperationAddress *GetAppInstanceForPartnerResponseBodyModuleAppOperationAddress `json:"AppOperationAddress,omitempty" xml:"AppOperationAddress,omitempty" type:"Struct"`
	// List of service instances associated with the application
	AppServiceList []*GetAppInstanceForPartnerResponseBodyModuleAppServiceList `json:"AppServiceList,omitempty" xml:"AppServiceList,omitempty" type:"Repeated"`
	// Application subtype
	//
	// example:
	//
	// placeHolder
	AppSubType *string `json:"AppSubType,omitempty" xml:"AppSubType,omitempty"`
	// Application Type
	//
	// example:
	//
	// TRACE
	AppType *string `json:"AppType,omitempty" xml:"AppType,omitempty"`
	// Application instance business ID
	//
	// example:
	//
	// WS20250915163734000001
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// Build Type
	//
	// example:
	//
	// placeHolder
	BuildType *string `json:"BuildType,omitempty" xml:"BuildType,omitempty"`
	// Is deleted
	//
	// example:
	//
	// fase
	Deleted *int32 `json:"Deleted,omitempty" xml:"Deleted,omitempty"`
	// Application Description
	//
	// example:
	//
	// /bak->serverless.handler(2020091300200279)
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// placeHolder
	//
	// example:
	//
	// placeHolder
	DesignSpecBizId *string `json:"DesignSpecBizId,omitempty" xml:"DesignSpecBizId,omitempty"`
	// Design Configuration ID
	//
	// example:
	//
	// placeHolder
	DesignSpecId *string `json:"DesignSpecId,omitempty" xml:"DesignSpecId,omitempty"`
	// Primary domain name
	//
	// example:
	//
	// stxycw.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// End time
	//
	// example:
	//
	// 2026-01-05T15:59:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// Environment
	//
	// example:
	//
	// pre
	Env *string `json:"Env,omitempty" xml:"Env,omitempty"`
	// Associated delivery order ID
	//
	// example:
	//
	// placeHolder
	EspBizId *string `json:"EspBizId,omitempty" xml:"EspBizId,omitempty"`
	// Creation Time
	//
	// example:
	//
	// 1740479834
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// Deletion time
	//
	// example:
	//
	// placeHolder
	GmtDelete *string `json:"GmtDelete,omitempty" xml:"GmtDelete,omitempty"`
	// Updated At
	//
	// example:
	//
	// 2025-08-28T02:25:41Z
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// Published At
	//
	// example:
	//
	// placeHolder
	GmtPublish *string `json:"GmtPublish,omitempty" xml:"GmtPublish,omitempty"`
	// Application icon
	//
	// example:
	//
	// https://yzm.zshud.com:34901/cloud-phone/养生_1763631730804.png
	IconUrl *string `json:"IconUrl,omitempty" xml:"IconUrl,omitempty"`
	// Application Name
	//
	// example:
	//
	// 文件名
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Order ID
	//
	// example:
	//
	// 250822465990301
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// Details
	PartnerDetail *GetAppInstanceForPartnerResponseBodyModulePartnerDetail `json:"PartnerDetail,omitempty" xml:"PartnerDetail,omitempty" type:"Struct"`
	// Instance Configuration
	Profile *GetAppInstanceForPartnerResponseBodyModuleProfile `json:"Profile,omitempty" xml:"Profile,omitempty" type:"Struct"`
	// List of instances associated with the application
	RelatedInstanceList []*GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceList `json:"RelatedInstanceList,omitempty" xml:"RelatedInstanceList,omitempty" type:"Repeated"`
	// Domain name
	//
	// example:
	//
	// abcd.scd.wanwang.xin
	SiteHost *string `json:"SiteHost,omitempty" xml:"SiteHost,omitempty"`
	// Application identifier
	//
	// example:
	//
	// placeHolder
	Slug *string `json:"Slug,omitempty" xml:"Slug,omitempty"`
	// Source MARKET_CLOUD_DREAM
	//
	// example:
	//
	// MARKET_CLOUD_DREAM
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// Start Time
	//
	// example:
	//
	// 2023-03-24T10:10Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// Status
	//
	// example:
	//
	// NORMAL
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Task Status text.
	//
	// example:
	//
	// FILE_DOWNLOAD_FAILED
	StatusText *string `json:"StatusText,omitempty" xml:"StatusText,omitempty"`
	// Application Thumbnail
	//
	// example:
	//
	// placeHolder
	ThumbnailUrl *string `json:"ThumbnailUrl,omitempty" xml:"ThumbnailUrl,omitempty"`
	// User ID
	//
	// example:
	//
	// placeHolder
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// Application instance Version
	//
	// example:
	//
	// 2019-04-02
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s GetAppInstanceForPartnerResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s GetAppInstanceForPartnerResponseBodyModule) GoString() string {
	return s.String()
}

func (s *GetAppInstanceForPartnerResponseBodyModule) GetAiStaffList() []*GetAppInstanceForPartnerResponseBodyModuleAiStaffList {
	return s.AiStaffList
}

func (s *GetAppInstanceForPartnerResponseBodyModule) GetAppDesignSpec() *GetAppInstanceForPartnerResponseBodyModuleAppDesignSpec {
	return s.AppDesignSpec
}

func (s *GetAppInstanceForPartnerResponseBodyModule) GetAppOperationAddress() *GetAppInstanceForPartnerResponseBodyModuleAppOperationAddress {
	return s.AppOperationAddress
}

func (s *GetAppInstanceForPartnerResponseBodyModule) GetAppServiceList() []*GetAppInstanceForPartnerResponseBodyModuleAppServiceList {
	return s.AppServiceList
}

func (s *GetAppInstanceForPartnerResponseBodyModule) GetAppSubType() *string {
	return s.AppSubType
}

func (s *GetAppInstanceForPartnerResponseBodyModule) GetAppType() *string {
	return s.AppType
}

func (s *GetAppInstanceForPartnerResponseBodyModule) GetBizId() *string {
	return s.BizId
}

func (s *GetAppInstanceForPartnerResponseBodyModule) GetBuildType() *string {
	return s.BuildType
}

func (s *GetAppInstanceForPartnerResponseBodyModule) GetDeleted() *int32 {
	return s.Deleted
}

func (s *GetAppInstanceForPartnerResponseBodyModule) GetDescription() *string {
	return s.Description
}

func (s *GetAppInstanceForPartnerResponseBodyModule) GetDesignSpecBizId() *string {
	return s.DesignSpecBizId
}

func (s *GetAppInstanceForPartnerResponseBodyModule) GetDesignSpecId() *string {
	return s.DesignSpecId
}

func (s *GetAppInstanceForPartnerResponseBodyModule) GetDomain() *string {
	return s.Domain
}

func (s *GetAppInstanceForPartnerResponseBodyModule) GetEndTime() *string {
	return s.EndTime
}

func (s *GetAppInstanceForPartnerResponseBodyModule) GetEnv() *string {
	return s.Env
}

func (s *GetAppInstanceForPartnerResponseBodyModule) GetEspBizId() *string {
	return s.EspBizId
}

func (s *GetAppInstanceForPartnerResponseBodyModule) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *GetAppInstanceForPartnerResponseBodyModule) GetGmtDelete() *string {
	return s.GmtDelete
}

func (s *GetAppInstanceForPartnerResponseBodyModule) GetGmtModified() *string {
	return s.GmtModified
}

func (s *GetAppInstanceForPartnerResponseBodyModule) GetGmtPublish() *string {
	return s.GmtPublish
}

func (s *GetAppInstanceForPartnerResponseBodyModule) GetIconUrl() *string {
	return s.IconUrl
}

func (s *GetAppInstanceForPartnerResponseBodyModule) GetName() *string {
	return s.Name
}

func (s *GetAppInstanceForPartnerResponseBodyModule) GetOrderId() *string {
	return s.OrderId
}

func (s *GetAppInstanceForPartnerResponseBodyModule) GetPartnerDetail() *GetAppInstanceForPartnerResponseBodyModulePartnerDetail {
	return s.PartnerDetail
}

func (s *GetAppInstanceForPartnerResponseBodyModule) GetProfile() *GetAppInstanceForPartnerResponseBodyModuleProfile {
	return s.Profile
}

func (s *GetAppInstanceForPartnerResponseBodyModule) GetRelatedInstanceList() []*GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceList {
	return s.RelatedInstanceList
}

func (s *GetAppInstanceForPartnerResponseBodyModule) GetSiteHost() *string {
	return s.SiteHost
}

func (s *GetAppInstanceForPartnerResponseBodyModule) GetSlug() *string {
	return s.Slug
}

func (s *GetAppInstanceForPartnerResponseBodyModule) GetSourceType() *string {
	return s.SourceType
}

func (s *GetAppInstanceForPartnerResponseBodyModule) GetStartTime() *string {
	return s.StartTime
}

func (s *GetAppInstanceForPartnerResponseBodyModule) GetStatus() *string {
	return s.Status
}

func (s *GetAppInstanceForPartnerResponseBodyModule) GetStatusText() *string {
	return s.StatusText
}

func (s *GetAppInstanceForPartnerResponseBodyModule) GetThumbnailUrl() *string {
	return s.ThumbnailUrl
}

func (s *GetAppInstanceForPartnerResponseBodyModule) GetUserId() *string {
	return s.UserId
}

func (s *GetAppInstanceForPartnerResponseBodyModule) GetVersion() *string {
	return s.Version
}

func (s *GetAppInstanceForPartnerResponseBodyModule) SetAiStaffList(v []*GetAppInstanceForPartnerResponseBodyModuleAiStaffList) *GetAppInstanceForPartnerResponseBodyModule {
	s.AiStaffList = v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModule) SetAppDesignSpec(v *GetAppInstanceForPartnerResponseBodyModuleAppDesignSpec) *GetAppInstanceForPartnerResponseBodyModule {
	s.AppDesignSpec = v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModule) SetAppOperationAddress(v *GetAppInstanceForPartnerResponseBodyModuleAppOperationAddress) *GetAppInstanceForPartnerResponseBodyModule {
	s.AppOperationAddress = v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModule) SetAppServiceList(v []*GetAppInstanceForPartnerResponseBodyModuleAppServiceList) *GetAppInstanceForPartnerResponseBodyModule {
	s.AppServiceList = v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModule) SetAppSubType(v string) *GetAppInstanceForPartnerResponseBodyModule {
	s.AppSubType = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModule) SetAppType(v string) *GetAppInstanceForPartnerResponseBodyModule {
	s.AppType = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModule) SetBizId(v string) *GetAppInstanceForPartnerResponseBodyModule {
	s.BizId = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModule) SetBuildType(v string) *GetAppInstanceForPartnerResponseBodyModule {
	s.BuildType = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModule) SetDeleted(v int32) *GetAppInstanceForPartnerResponseBodyModule {
	s.Deleted = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModule) SetDescription(v string) *GetAppInstanceForPartnerResponseBodyModule {
	s.Description = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModule) SetDesignSpecBizId(v string) *GetAppInstanceForPartnerResponseBodyModule {
	s.DesignSpecBizId = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModule) SetDesignSpecId(v string) *GetAppInstanceForPartnerResponseBodyModule {
	s.DesignSpecId = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModule) SetDomain(v string) *GetAppInstanceForPartnerResponseBodyModule {
	s.Domain = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModule) SetEndTime(v string) *GetAppInstanceForPartnerResponseBodyModule {
	s.EndTime = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModule) SetEnv(v string) *GetAppInstanceForPartnerResponseBodyModule {
	s.Env = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModule) SetEspBizId(v string) *GetAppInstanceForPartnerResponseBodyModule {
	s.EspBizId = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModule) SetGmtCreate(v string) *GetAppInstanceForPartnerResponseBodyModule {
	s.GmtCreate = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModule) SetGmtDelete(v string) *GetAppInstanceForPartnerResponseBodyModule {
	s.GmtDelete = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModule) SetGmtModified(v string) *GetAppInstanceForPartnerResponseBodyModule {
	s.GmtModified = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModule) SetGmtPublish(v string) *GetAppInstanceForPartnerResponseBodyModule {
	s.GmtPublish = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModule) SetIconUrl(v string) *GetAppInstanceForPartnerResponseBodyModule {
	s.IconUrl = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModule) SetName(v string) *GetAppInstanceForPartnerResponseBodyModule {
	s.Name = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModule) SetOrderId(v string) *GetAppInstanceForPartnerResponseBodyModule {
	s.OrderId = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModule) SetPartnerDetail(v *GetAppInstanceForPartnerResponseBodyModulePartnerDetail) *GetAppInstanceForPartnerResponseBodyModule {
	s.PartnerDetail = v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModule) SetProfile(v *GetAppInstanceForPartnerResponseBodyModuleProfile) *GetAppInstanceForPartnerResponseBodyModule {
	s.Profile = v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModule) SetRelatedInstanceList(v []*GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceList) *GetAppInstanceForPartnerResponseBodyModule {
	s.RelatedInstanceList = v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModule) SetSiteHost(v string) *GetAppInstanceForPartnerResponseBodyModule {
	s.SiteHost = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModule) SetSlug(v string) *GetAppInstanceForPartnerResponseBodyModule {
	s.Slug = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModule) SetSourceType(v string) *GetAppInstanceForPartnerResponseBodyModule {
	s.SourceType = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModule) SetStartTime(v string) *GetAppInstanceForPartnerResponseBodyModule {
	s.StartTime = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModule) SetStatus(v string) *GetAppInstanceForPartnerResponseBodyModule {
	s.Status = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModule) SetStatusText(v string) *GetAppInstanceForPartnerResponseBodyModule {
	s.StatusText = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModule) SetThumbnailUrl(v string) *GetAppInstanceForPartnerResponseBodyModule {
	s.ThumbnailUrl = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModule) SetUserId(v string) *GetAppInstanceForPartnerResponseBodyModule {
	s.UserId = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModule) SetVersion(v string) *GetAppInstanceForPartnerResponseBodyModule {
	s.Version = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModule) Validate() error {
	if s.AiStaffList != nil {
		for _, item := range s.AiStaffList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.AppDesignSpec != nil {
		if err := s.AppDesignSpec.Validate(); err != nil {
			return err
		}
	}
	if s.AppOperationAddress != nil {
		if err := s.AppOperationAddress.Validate(); err != nil {
			return err
		}
	}
	if s.AppServiceList != nil {
		for _, item := range s.AppServiceList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.PartnerDetail != nil {
		if err := s.PartnerDetail.Validate(); err != nil {
			return err
		}
	}
	if s.Profile != nil {
		if err := s.Profile.Validate(); err != nil {
			return err
		}
	}
	if s.RelatedInstanceList != nil {
		for _, item := range s.RelatedInstanceList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetAppInstanceForPartnerResponseBodyModuleAiStaffList struct {
	// Staff ID
	//
	// example:
	//
	// 370196
	StaffId *string `json:"StaffId,omitempty" xml:"StaffId,omitempty"`
	// Staff Name
	//
	// example:
	//
	// StaffName
	StaffName *string `json:"StaffName,omitempty" xml:"StaffName,omitempty"`
	// Staff Type
	//
	// example:
	//
	// StaffType
	StaffType *string `json:"StaffType,omitempty" xml:"StaffType,omitempty"`
	// Staff Status
	//
	// example:
	//
	// 200
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetAppInstanceForPartnerResponseBodyModuleAiStaffList) String() string {
	return dara.Prettify(s)
}

func (s GetAppInstanceForPartnerResponseBodyModuleAiStaffList) GoString() string {
	return s.String()
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAiStaffList) GetStaffId() *string {
	return s.StaffId
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAiStaffList) GetStaffName() *string {
	return s.StaffName
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAiStaffList) GetStaffType() *string {
	return s.StaffType
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAiStaffList) GetStatus() *string {
	return s.Status
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAiStaffList) SetStaffId(v string) *GetAppInstanceForPartnerResponseBodyModuleAiStaffList {
	s.StaffId = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAiStaffList) SetStaffName(v string) *GetAppInstanceForPartnerResponseBodyModuleAiStaffList {
	s.StaffName = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAiStaffList) SetStaffType(v string) *GetAppInstanceForPartnerResponseBodyModuleAiStaffList {
	s.StaffType = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAiStaffList) SetStatus(v string) *GetAppInstanceForPartnerResponseBodyModuleAiStaffList {
	s.Status = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAiStaffList) Validate() error {
	return dara.Validate(s)
}

type GetAppInstanceForPartnerResponseBodyModuleAppDesignSpec struct {
	// bill
	//
	// example:
	//
	// bilingual
	Bilingual *bool `json:"Bilingual,omitempty" xml:"Bilingual,omitempty"`
	// Customer Business ID
	//
	// example:
	//
	// WS20250801151731000007
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// busincessScope
	//
	// example:
	//
	// scopre
	BusinessScope *string `json:"BusinessScope,omitempty" xml:"BusinessScope,omitempty"`
	// Specifies the color. Valid values: - **black**: black - **white**: white - **red**: red - **blue**: blue
	//
	// example:
	//
	// style
	ColorStyle *string `json:"ColorStyle,omitempty" xml:"ColorStyle,omitempty"`
	// Company Name
	//
	// example:
	//
	// bvt_test
	CompanyName *string `json:"CompanyName,omitempty" xml:"CompanyName,omitempty"`
	// deployment region
	//
	// example:
	//
	// HongKong
	DeployArea *string `json:"DeployArea,omitempty" xml:"DeployArea,omitempty"`
	// Creation Time
	//
	// example:
	//
	// 2022-12-21T08:27:03Z
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// Updated At
	//
	// example:
	//
	// Wed Sep 10 09:53:35 CST 2025
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// primary key
	//
	// example:
	//
	// m1zumwgy6u
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// business
	//
	// example:
	//
	// business
	MainBusiness *string `json:"MainBusiness,omitempty" xml:"MainBusiness,omitempty"`
	// website
	//
	// example:
	//
	// website
	ReferenceWebsite *string `json:"ReferenceWebsite,omitempty" xml:"ReferenceWebsite,omitempty"`
	// sitegolas
	//
	// example:
	//
	// goals
	SiteGoals *string `json:"SiteGoals,omitempty" xml:"SiteGoals,omitempty"`
	// language
	//
	// example:
	//
	// sitelanguage
	SiteLanguage *string `json:"SiteLanguage,omitempty" xml:"SiteLanguage,omitempty"`
	// sitelogo
	//
	// example:
	//
	// logo
	SiteLogo *string `json:"SiteLogo,omitempty" xml:"SiteLogo,omitempty"`
	// siteslogan
	//
	// example:
	//
	// slogan
	SiteSlogan *string `json:"SiteSlogan,omitempty" xml:"SiteSlogan,omitempty"`
	// sitestyle
	//
	// example:
	//
	// style
	SiteStyle *string `json:"SiteStyle,omitempty" xml:"SiteStyle,omitempty"`
	// sitetitle
	//
	// example:
	//
	// title
	SiteTitle *string `json:"SiteTitle,omitempty" xml:"SiteTitle,omitempty"`
	// Application Type. Web & H5: web, miniapp: miniapp, iOS: ios, Android: android.
	//
	// example:
	//
	// web
	SiteType *string `json:"SiteType,omitempty" xml:"SiteType,omitempty"`
	// userid
	//
	// example:
	//
	// userid
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s GetAppInstanceForPartnerResponseBodyModuleAppDesignSpec) String() string {
	return dara.Prettify(s)
}

func (s GetAppInstanceForPartnerResponseBodyModuleAppDesignSpec) GoString() string {
	return s.String()
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppDesignSpec) GetBilingual() *bool {
	return s.Bilingual
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppDesignSpec) GetBizId() *string {
	return s.BizId
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppDesignSpec) GetBusinessScope() *string {
	return s.BusinessScope
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppDesignSpec) GetColorStyle() *string {
	return s.ColorStyle
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppDesignSpec) GetCompanyName() *string {
	return s.CompanyName
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppDesignSpec) GetDeployArea() *string {
	return s.DeployArea
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppDesignSpec) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppDesignSpec) GetGmtModified() *string {
	return s.GmtModified
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppDesignSpec) GetId() *int64 {
	return s.Id
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppDesignSpec) GetMainBusiness() *string {
	return s.MainBusiness
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppDesignSpec) GetReferenceWebsite() *string {
	return s.ReferenceWebsite
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppDesignSpec) GetSiteGoals() *string {
	return s.SiteGoals
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppDesignSpec) GetSiteLanguage() *string {
	return s.SiteLanguage
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppDesignSpec) GetSiteLogo() *string {
	return s.SiteLogo
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppDesignSpec) GetSiteSlogan() *string {
	return s.SiteSlogan
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppDesignSpec) GetSiteStyle() *string {
	return s.SiteStyle
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppDesignSpec) GetSiteTitle() *string {
	return s.SiteTitle
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppDesignSpec) GetSiteType() *string {
	return s.SiteType
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppDesignSpec) GetUserId() *string {
	return s.UserId
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppDesignSpec) SetBilingual(v bool) *GetAppInstanceForPartnerResponseBodyModuleAppDesignSpec {
	s.Bilingual = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppDesignSpec) SetBizId(v string) *GetAppInstanceForPartnerResponseBodyModuleAppDesignSpec {
	s.BizId = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppDesignSpec) SetBusinessScope(v string) *GetAppInstanceForPartnerResponseBodyModuleAppDesignSpec {
	s.BusinessScope = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppDesignSpec) SetColorStyle(v string) *GetAppInstanceForPartnerResponseBodyModuleAppDesignSpec {
	s.ColorStyle = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppDesignSpec) SetCompanyName(v string) *GetAppInstanceForPartnerResponseBodyModuleAppDesignSpec {
	s.CompanyName = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppDesignSpec) SetDeployArea(v string) *GetAppInstanceForPartnerResponseBodyModuleAppDesignSpec {
	s.DeployArea = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppDesignSpec) SetGmtCreate(v string) *GetAppInstanceForPartnerResponseBodyModuleAppDesignSpec {
	s.GmtCreate = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppDesignSpec) SetGmtModified(v string) *GetAppInstanceForPartnerResponseBodyModuleAppDesignSpec {
	s.GmtModified = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppDesignSpec) SetId(v int64) *GetAppInstanceForPartnerResponseBodyModuleAppDesignSpec {
	s.Id = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppDesignSpec) SetMainBusiness(v string) *GetAppInstanceForPartnerResponseBodyModuleAppDesignSpec {
	s.MainBusiness = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppDesignSpec) SetReferenceWebsite(v string) *GetAppInstanceForPartnerResponseBodyModuleAppDesignSpec {
	s.ReferenceWebsite = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppDesignSpec) SetSiteGoals(v string) *GetAppInstanceForPartnerResponseBodyModuleAppDesignSpec {
	s.SiteGoals = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppDesignSpec) SetSiteLanguage(v string) *GetAppInstanceForPartnerResponseBodyModuleAppDesignSpec {
	s.SiteLanguage = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppDesignSpec) SetSiteLogo(v string) *GetAppInstanceForPartnerResponseBodyModuleAppDesignSpec {
	s.SiteLogo = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppDesignSpec) SetSiteSlogan(v string) *GetAppInstanceForPartnerResponseBodyModuleAppDesignSpec {
	s.SiteSlogan = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppDesignSpec) SetSiteStyle(v string) *GetAppInstanceForPartnerResponseBodyModuleAppDesignSpec {
	s.SiteStyle = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppDesignSpec) SetSiteTitle(v string) *GetAppInstanceForPartnerResponseBodyModuleAppDesignSpec {
	s.SiteTitle = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppDesignSpec) SetSiteType(v string) *GetAppInstanceForPartnerResponseBodyModuleAppDesignSpec {
	s.SiteType = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppDesignSpec) SetUserId(v string) *GetAppInstanceForPartnerResponseBodyModuleAppDesignSpec {
	s.UserId = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppDesignSpec) Validate() error {
	return dara.Validate(s)
}

type GetAppInstanceForPartnerResponseBodyModuleAppOperationAddress struct {
	// List of console operation buttons
	Actions []*GetAppInstanceForPartnerResponseBodyModuleAppOperationAddressActions `json:"Actions,omitempty" xml:"Actions,omitempty" type:"Repeated"`
	// Customer service design URL
	//
	// example:
	//
	// placeHolder
	AiCustomerConfigUrl *string `json:"AiCustomerConfigUrl,omitempty" xml:"AiCustomerConfigUrl,omitempty"`
	// AI Design URL
	//
	// example:
	//
	// placeHolder
	AiDesignUrl *string `json:"AiDesignUrl,omitempty" xml:"AiDesignUrl,omitempty"`
	// Application publish URL
	//
	// example:
	//
	// placeHolder
	AppPublishUrl *string `json:"AppPublishUrl,omitempty" xml:"AppPublishUrl,omitempty"`
	// List of action buttons on the overview page
	DashboardActions []*GetAppInstanceForPartnerResponseBodyModuleAppOperationAddressDashboardActions `json:"DashboardActions,omitempty" xml:"DashboardActions,omitempty" type:"Repeated"`
	// Professional design URL
	//
	// example:
	//
	// placeHolder
	DesignUrl *string `json:"DesignUrl,omitempty" xml:"DesignUrl,omitempty"`
	// Application instance logon URL
	//
	// example:
	//
	// placeHolder
	InstanceLoginUrl *string `json:"InstanceLoginUrl,omitempty" xml:"InstanceLoginUrl,omitempty"`
	// Renewal and purchase URL
	//
	// example:
	//
	// placeHolder
	RenewBuyUrl *string `json:"RenewBuyUrl,omitempty" xml:"RenewBuyUrl,omitempty"`
	// Service delivery URL
	//
	// example:
	//
	// placeHolder
	ServerDeliveryUrl *string `json:"ServerDeliveryUrl,omitempty" xml:"ServerDeliveryUrl,omitempty"`
	// Upgrade purchase URL
	//
	// example:
	//
	// placeHolder
	UpgradeBuyUrl *string `json:"UpgradeBuyUrl,omitempty" xml:"UpgradeBuyUrl,omitempty"`
}

func (s GetAppInstanceForPartnerResponseBodyModuleAppOperationAddress) String() string {
	return dara.Prettify(s)
}

func (s GetAppInstanceForPartnerResponseBodyModuleAppOperationAddress) GoString() string {
	return s.String()
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppOperationAddress) GetActions() []*GetAppInstanceForPartnerResponseBodyModuleAppOperationAddressActions {
	return s.Actions
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppOperationAddress) GetAiCustomerConfigUrl() *string {
	return s.AiCustomerConfigUrl
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppOperationAddress) GetAiDesignUrl() *string {
	return s.AiDesignUrl
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppOperationAddress) GetAppPublishUrl() *string {
	return s.AppPublishUrl
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppOperationAddress) GetDashboardActions() []*GetAppInstanceForPartnerResponseBodyModuleAppOperationAddressDashboardActions {
	return s.DashboardActions
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppOperationAddress) GetDesignUrl() *string {
	return s.DesignUrl
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppOperationAddress) GetInstanceLoginUrl() *string {
	return s.InstanceLoginUrl
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppOperationAddress) GetRenewBuyUrl() *string {
	return s.RenewBuyUrl
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppOperationAddress) GetServerDeliveryUrl() *string {
	return s.ServerDeliveryUrl
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppOperationAddress) GetUpgradeBuyUrl() *string {
	return s.UpgradeBuyUrl
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppOperationAddress) SetActions(v []*GetAppInstanceForPartnerResponseBodyModuleAppOperationAddressActions) *GetAppInstanceForPartnerResponseBodyModuleAppOperationAddress {
	s.Actions = v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppOperationAddress) SetAiCustomerConfigUrl(v string) *GetAppInstanceForPartnerResponseBodyModuleAppOperationAddress {
	s.AiCustomerConfigUrl = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppOperationAddress) SetAiDesignUrl(v string) *GetAppInstanceForPartnerResponseBodyModuleAppOperationAddress {
	s.AiDesignUrl = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppOperationAddress) SetAppPublishUrl(v string) *GetAppInstanceForPartnerResponseBodyModuleAppOperationAddress {
	s.AppPublishUrl = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppOperationAddress) SetDashboardActions(v []*GetAppInstanceForPartnerResponseBodyModuleAppOperationAddressDashboardActions) *GetAppInstanceForPartnerResponseBodyModuleAppOperationAddress {
	s.DashboardActions = v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppOperationAddress) SetDesignUrl(v string) *GetAppInstanceForPartnerResponseBodyModuleAppOperationAddress {
	s.DesignUrl = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppOperationAddress) SetInstanceLoginUrl(v string) *GetAppInstanceForPartnerResponseBodyModuleAppOperationAddress {
	s.InstanceLoginUrl = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppOperationAddress) SetRenewBuyUrl(v string) *GetAppInstanceForPartnerResponseBodyModuleAppOperationAddress {
	s.RenewBuyUrl = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppOperationAddress) SetServerDeliveryUrl(v string) *GetAppInstanceForPartnerResponseBodyModuleAppOperationAddress {
	s.ServerDeliveryUrl = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppOperationAddress) SetUpgradeBuyUrl(v string) *GetAppInstanceForPartnerResponseBodyModuleAppOperationAddress {
	s.UpgradeBuyUrl = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppOperationAddress) Validate() error {
	if s.Actions != nil {
		for _, item := range s.Actions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.DashboardActions != nil {
		for _, item := range s.DashboardActions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetAppInstanceForPartnerResponseBodyModuleAppOperationAddressActions struct {
	// Operation key value, used as a UUID for an operation behavior.
	//
	// example:
	//
	// BeginDialogue
	ActionKey *string `json:"ActionKey,omitempty" xml:"ActionKey,omitempty"`
	// Action display text, used to show the action name in the interface
	//
	// example:
	//
	// placeHolder
	ActionText *string `json:"ActionText,omitempty" xml:"ActionText,omitempty"`
	// Enable status, indicating whether the operation is active.
	//
	// example:
	//
	// false
	Enable *bool `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// Link address, the URL to which the user is redirected when clicking the action
	//
	// example:
	//
	// placeHolder
	Href *string `json:"Href,omitempty" xml:"Href,omitempty"`
}

func (s GetAppInstanceForPartnerResponseBodyModuleAppOperationAddressActions) String() string {
	return dara.Prettify(s)
}

func (s GetAppInstanceForPartnerResponseBodyModuleAppOperationAddressActions) GoString() string {
	return s.String()
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppOperationAddressActions) GetActionKey() *string {
	return s.ActionKey
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppOperationAddressActions) GetActionText() *string {
	return s.ActionText
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppOperationAddressActions) GetEnable() *bool {
	return s.Enable
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppOperationAddressActions) GetHref() *string {
	return s.Href
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppOperationAddressActions) SetActionKey(v string) *GetAppInstanceForPartnerResponseBodyModuleAppOperationAddressActions {
	s.ActionKey = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppOperationAddressActions) SetActionText(v string) *GetAppInstanceForPartnerResponseBodyModuleAppOperationAddressActions {
	s.ActionText = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppOperationAddressActions) SetEnable(v bool) *GetAppInstanceForPartnerResponseBodyModuleAppOperationAddressActions {
	s.Enable = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppOperationAddressActions) SetHref(v string) *GetAppInstanceForPartnerResponseBodyModuleAppOperationAddressActions {
	s.Href = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppOperationAddressActions) Validate() error {
	return dara.Validate(s)
}

type GetAppInstanceForPartnerResponseBodyModuleAppOperationAddressDashboardActions struct {
	// Operation key used to uniquely identify an operation behavior.
	//
	// example:
	//
	// CollectedNumber
	ActionKey *string `json:"ActionKey,omitempty" xml:"ActionKey,omitempty"`
	// Action display text used to show the action name in the interface
	//
	// example:
	//
	// placeHolder
	ActionText *string `json:"ActionText,omitempty" xml:"ActionText,omitempty"`
	// Enable status indicating whether the action is active
	//
	// example:
	//
	// true
	Enable *bool `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// Link URL to which the user is redirected when clicking the action
	//
	// example:
	//
	// placeHolder
	Href *string `json:"Href,omitempty" xml:"Href,omitempty"`
}

func (s GetAppInstanceForPartnerResponseBodyModuleAppOperationAddressDashboardActions) String() string {
	return dara.Prettify(s)
}

func (s GetAppInstanceForPartnerResponseBodyModuleAppOperationAddressDashboardActions) GoString() string {
	return s.String()
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppOperationAddressDashboardActions) GetActionKey() *string {
	return s.ActionKey
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppOperationAddressDashboardActions) GetActionText() *string {
	return s.ActionText
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppOperationAddressDashboardActions) GetEnable() *bool {
	return s.Enable
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppOperationAddressDashboardActions) GetHref() *string {
	return s.Href
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppOperationAddressDashboardActions) SetActionKey(v string) *GetAppInstanceForPartnerResponseBodyModuleAppOperationAddressDashboardActions {
	s.ActionKey = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppOperationAddressDashboardActions) SetActionText(v string) *GetAppInstanceForPartnerResponseBodyModuleAppOperationAddressDashboardActions {
	s.ActionText = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppOperationAddressDashboardActions) SetEnable(v bool) *GetAppInstanceForPartnerResponseBodyModuleAppOperationAddressDashboardActions {
	s.Enable = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppOperationAddressDashboardActions) SetHref(v string) *GetAppInstanceForPartnerResponseBodyModuleAppOperationAddressDashboardActions {
	s.Href = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppOperationAddressDashboardActions) Validate() error {
	return dara.Validate(s)
}

type GetAppInstanceForPartnerResponseBodyModuleAppServiceList struct {
	// Application instance business ID
	//
	// example:
	//
	// WS20251211153330000001
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// Deletion flag
	//
	// example:
	//
	// fase
	Deleted *int32 `json:"Deleted,omitempty" xml:"Deleted,omitempty"`
	// End time
	//
	// example:
	//
	// 2025-07-18T02:23:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// esp bizId
	//
	// example:
	//
	// EspBizId
	EspBizId *string `json:"EspBizId,omitempty" xml:"EspBizId,omitempty"`
	// Creation Time
	//
	// example:
	//
	// 2025-07-30T02:08:40Z
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// Updated At
	//
	// example:
	//
	// Tue Sep 09 10:27:49 CST 2025
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// Application operation address
	Group *GetAppInstanceForPartnerResponseBodyModuleAppServiceListGroup `json:"Group,omitempty" xml:"Group,omitempty" type:"Struct"`
	// Associated business ID
	//
	// example:
	//
	// InstanceBizId
	InstanceBizId *string `json:"InstanceBizId,omitempty" xml:"InstanceBizId,omitempty"`
	// Service name
	//
	// example:
	//
	// 19609820.png
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Delivery flow
	NodeList []*GetAppInstanceForPartnerResponseBodyModuleAppServiceListNodeList `json:"NodeList,omitempty" xml:"NodeList,omitempty" type:"Repeated"`
	// Used to uniquely identify an operation behavior
	OperationAddress *GetAppInstanceForPartnerResponseBodyModuleAppServiceListOperationAddress `json:"OperationAddress,omitempty" xml:"OperationAddress,omitempty" type:"Struct"`
	// Order ID
	//
	// example:
	//
	// 253790948890026
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// Instance service configuration
	Profile *GetAppInstanceForPartnerResponseBodyModuleAppServiceListProfile `json:"Profile,omitempty" xml:"Profile,omitempty" type:"Struct"`
	// Service Type
	//
	// example:
	//
	// WEBSITE_DESIGN
	ServiceType *string `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
	// Service Type Description
	//
	// example:
	//
	// ServiceTypeText
	ServiceTypeText *string `json:"ServiceTypeText,omitempty" xml:"ServiceTypeText,omitempty"`
	// Application identity
	//
	// example:
	//
	// Slug
	Slug *string `json:"Slug,omitempty" xml:"Slug,omitempty"`
	// Start Time
	//
	// example:
	//
	// 1754447102000
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// Status
	//
	// example:
	//
	// success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// User ID
	//
	// example:
	//
	// UserId
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s GetAppInstanceForPartnerResponseBodyModuleAppServiceList) String() string {
	return dara.Prettify(s)
}

func (s GetAppInstanceForPartnerResponseBodyModuleAppServiceList) GoString() string {
	return s.String()
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppServiceList) GetBizId() *string {
	return s.BizId
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppServiceList) GetDeleted() *int32 {
	return s.Deleted
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppServiceList) GetEndTime() *string {
	return s.EndTime
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppServiceList) GetEspBizId() *string {
	return s.EspBizId
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppServiceList) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppServiceList) GetGmtModified() *string {
	return s.GmtModified
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppServiceList) GetGroup() *GetAppInstanceForPartnerResponseBodyModuleAppServiceListGroup {
	return s.Group
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppServiceList) GetInstanceBizId() *string {
	return s.InstanceBizId
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppServiceList) GetName() *string {
	return s.Name
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppServiceList) GetNodeList() []*GetAppInstanceForPartnerResponseBodyModuleAppServiceListNodeList {
	return s.NodeList
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppServiceList) GetOperationAddress() *GetAppInstanceForPartnerResponseBodyModuleAppServiceListOperationAddress {
	return s.OperationAddress
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppServiceList) GetOrderId() *string {
	return s.OrderId
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppServiceList) GetProfile() *GetAppInstanceForPartnerResponseBodyModuleAppServiceListProfile {
	return s.Profile
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppServiceList) GetServiceType() *string {
	return s.ServiceType
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppServiceList) GetServiceTypeText() *string {
	return s.ServiceTypeText
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppServiceList) GetSlug() *string {
	return s.Slug
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppServiceList) GetStartTime() *string {
	return s.StartTime
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppServiceList) GetStatus() *string {
	return s.Status
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppServiceList) GetUserId() *string {
	return s.UserId
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppServiceList) SetBizId(v string) *GetAppInstanceForPartnerResponseBodyModuleAppServiceList {
	s.BizId = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppServiceList) SetDeleted(v int32) *GetAppInstanceForPartnerResponseBodyModuleAppServiceList {
	s.Deleted = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppServiceList) SetEndTime(v string) *GetAppInstanceForPartnerResponseBodyModuleAppServiceList {
	s.EndTime = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppServiceList) SetEspBizId(v string) *GetAppInstanceForPartnerResponseBodyModuleAppServiceList {
	s.EspBizId = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppServiceList) SetGmtCreate(v string) *GetAppInstanceForPartnerResponseBodyModuleAppServiceList {
	s.GmtCreate = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppServiceList) SetGmtModified(v string) *GetAppInstanceForPartnerResponseBodyModuleAppServiceList {
	s.GmtModified = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppServiceList) SetGroup(v *GetAppInstanceForPartnerResponseBodyModuleAppServiceListGroup) *GetAppInstanceForPartnerResponseBodyModuleAppServiceList {
	s.Group = v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppServiceList) SetInstanceBizId(v string) *GetAppInstanceForPartnerResponseBodyModuleAppServiceList {
	s.InstanceBizId = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppServiceList) SetName(v string) *GetAppInstanceForPartnerResponseBodyModuleAppServiceList {
	s.Name = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppServiceList) SetNodeList(v []*GetAppInstanceForPartnerResponseBodyModuleAppServiceListNodeList) *GetAppInstanceForPartnerResponseBodyModuleAppServiceList {
	s.NodeList = v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppServiceList) SetOperationAddress(v *GetAppInstanceForPartnerResponseBodyModuleAppServiceListOperationAddress) *GetAppInstanceForPartnerResponseBodyModuleAppServiceList {
	s.OperationAddress = v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppServiceList) SetOrderId(v string) *GetAppInstanceForPartnerResponseBodyModuleAppServiceList {
	s.OrderId = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppServiceList) SetProfile(v *GetAppInstanceForPartnerResponseBodyModuleAppServiceListProfile) *GetAppInstanceForPartnerResponseBodyModuleAppServiceList {
	s.Profile = v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppServiceList) SetServiceType(v string) *GetAppInstanceForPartnerResponseBodyModuleAppServiceList {
	s.ServiceType = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppServiceList) SetServiceTypeText(v string) *GetAppInstanceForPartnerResponseBodyModuleAppServiceList {
	s.ServiceTypeText = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppServiceList) SetSlug(v string) *GetAppInstanceForPartnerResponseBodyModuleAppServiceList {
	s.Slug = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppServiceList) SetStartTime(v string) *GetAppInstanceForPartnerResponseBodyModuleAppServiceList {
	s.StartTime = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppServiceList) SetStatus(v string) *GetAppInstanceForPartnerResponseBodyModuleAppServiceList {
	s.Status = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppServiceList) SetUserId(v string) *GetAppInstanceForPartnerResponseBodyModuleAppServiceList {
	s.UserId = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppServiceList) Validate() error {
	if s.Group != nil {
		if err := s.Group.Validate(); err != nil {
			return err
		}
	}
	if s.NodeList != nil {
		for _, item := range s.NodeList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.OperationAddress != nil {
		if err := s.OperationAddress.Validate(); err != nil {
			return err
		}
	}
	if s.Profile != nil {
		if err := s.Profile.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAppInstanceForPartnerResponseBodyModuleAppServiceListGroup struct {
	// primary key
	//
	// example:
	//
	// d75fvq3ctk
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// Website Name
	//
	// example:
	//
	// mda-sb037wmidshg3w9v.mp4
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// QrCode
	//
	// example:
	//
	// qrcode
	QrCode *string `json:"QrCode,omitempty" xml:"QrCode,omitempty"`
	// dingtalk wx...
	//
	// example:
	//
	// hive
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// Temporary access URL for image
	//
	// example:
	//
	// https://static.yipigai.cn/timuocr/tmp_57bc9cb3be1075f4e5cdae87f5cbb86abc54a694aaf10965.jpg
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s GetAppInstanceForPartnerResponseBodyModuleAppServiceListGroup) String() string {
	return dara.Prettify(s)
}

func (s GetAppInstanceForPartnerResponseBodyModuleAppServiceListGroup) GoString() string {
	return s.String()
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppServiceListGroup) GetId() *string {
	return s.Id
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppServiceListGroup) GetName() *string {
	return s.Name
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppServiceListGroup) GetQrCode() *string {
	return s.QrCode
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppServiceListGroup) GetType() *string {
	return s.Type
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppServiceListGroup) GetUrl() *string {
	return s.Url
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppServiceListGroup) SetId(v string) *GetAppInstanceForPartnerResponseBodyModuleAppServiceListGroup {
	s.Id = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppServiceListGroup) SetName(v string) *GetAppInstanceForPartnerResponseBodyModuleAppServiceListGroup {
	s.Name = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppServiceListGroup) SetQrCode(v string) *GetAppInstanceForPartnerResponseBodyModuleAppServiceListGroup {
	s.QrCode = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppServiceListGroup) SetType(v string) *GetAppInstanceForPartnerResponseBodyModuleAppServiceListGroup {
	s.Type = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppServiceListGroup) SetUrl(v string) *GetAppInstanceForPartnerResponseBodyModuleAppServiceListGroup {
	s.Url = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppServiceListGroup) Validate() error {
	return dara.Validate(s)
}

type GetAppInstanceForPartnerResponseBodyModuleAppServiceListNodeList struct {
	// Child folder object
	Children []interface{} `json:"Children,omitempty" xml:"Children,omitempty" type:"Repeated"`
	// The ordinal number of the last node.
	//
	// example:
	//
	// FinalStepNo
	FinalStepNo *int32 `json:"FinalStepNo,omitempty" xml:"FinalStepNo,omitempty"`
	// When FinishStatus (baseline instance completion status) is FINISH (completed), the UNIX timestamp of the baseline instance completion time is returned.
	//
	// example:
	//
	// 2023-01-13T07:58:12Z
	FinishTime *int64 `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	// IsContainerNode
	//
	// example:
	//
	// IsContainerNode
	IsContainerNode *bool `json:"IsContainerNode,omitempty" xml:"IsContainerNode,omitempty"`
	// The ID of the OBServer edge zone where the resource node resides.
	//
	// example:
	//
	// 2927b500-c4e2-4241-bacf-0a2991c4be12
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// <props="china">The name of the matrix.
	//
	// <props="intl">This parameter is not publicly available.
	//
	// example:
	//
	// 单路输出
	NodeName *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	// Node running status. Valid values:
	//
	// - **UP**: The node is running.
	//
	// - **DOWN**: The node is faulty.
	//
	// example:
	//
	// NodeStatus
	NodeStatus *string `json:"NodeStatus,omitempty" xml:"NodeStatus,omitempty"`
	// The role of the operator.
	//
	// example:
	//
	// OperatorRole
	OperatorRole *string `json:"OperatorRole,omitempty" xml:"OperatorRole,omitempty"`
	// Parent node ID
	//
	// example:
	//
	// dpYLaezmVNRMGX56Cg4gLjrrVrMqPxX6
	ParentNodeId *string `json:"ParentNodeId,omitempty" xml:"ParentNodeId,omitempty"`
	// Ordinal number of the flow node
	//
	// example:
	//
	// StepNo
	StepNo *int32 `json:"StepNo,omitempty" xml:"StepNo,omitempty"`
}

func (s GetAppInstanceForPartnerResponseBodyModuleAppServiceListNodeList) String() string {
	return dara.Prettify(s)
}

func (s GetAppInstanceForPartnerResponseBodyModuleAppServiceListNodeList) GoString() string {
	return s.String()
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppServiceListNodeList) GetChildren() []interface{} {
	return s.Children
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppServiceListNodeList) GetFinalStepNo() *int32 {
	return s.FinalStepNo
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppServiceListNodeList) GetFinishTime() *int64 {
	return s.FinishTime
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppServiceListNodeList) GetIsContainerNode() *bool {
	return s.IsContainerNode
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppServiceListNodeList) GetNodeId() *string {
	return s.NodeId
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppServiceListNodeList) GetNodeName() *string {
	return s.NodeName
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppServiceListNodeList) GetNodeStatus() *string {
	return s.NodeStatus
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppServiceListNodeList) GetOperatorRole() *string {
	return s.OperatorRole
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppServiceListNodeList) GetParentNodeId() *string {
	return s.ParentNodeId
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppServiceListNodeList) GetStepNo() *int32 {
	return s.StepNo
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppServiceListNodeList) SetChildren(v []interface{}) *GetAppInstanceForPartnerResponseBodyModuleAppServiceListNodeList {
	s.Children = v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppServiceListNodeList) SetFinalStepNo(v int32) *GetAppInstanceForPartnerResponseBodyModuleAppServiceListNodeList {
	s.FinalStepNo = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppServiceListNodeList) SetFinishTime(v int64) *GetAppInstanceForPartnerResponseBodyModuleAppServiceListNodeList {
	s.FinishTime = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppServiceListNodeList) SetIsContainerNode(v bool) *GetAppInstanceForPartnerResponseBodyModuleAppServiceListNodeList {
	s.IsContainerNode = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppServiceListNodeList) SetNodeId(v string) *GetAppInstanceForPartnerResponseBodyModuleAppServiceListNodeList {
	s.NodeId = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppServiceListNodeList) SetNodeName(v string) *GetAppInstanceForPartnerResponseBodyModuleAppServiceListNodeList {
	s.NodeName = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppServiceListNodeList) SetNodeStatus(v string) *GetAppInstanceForPartnerResponseBodyModuleAppServiceListNodeList {
	s.NodeStatus = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppServiceListNodeList) SetOperatorRole(v string) *GetAppInstanceForPartnerResponseBodyModuleAppServiceListNodeList {
	s.OperatorRole = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppServiceListNodeList) SetParentNodeId(v string) *GetAppInstanceForPartnerResponseBodyModuleAppServiceListNodeList {
	s.ParentNodeId = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppServiceListNodeList) SetStepNo(v int32) *GetAppInstanceForPartnerResponseBodyModuleAppServiceListNodeList {
	s.StepNo = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppServiceListNodeList) Validate() error {
	return dara.Validate(s)
}

type GetAppInstanceForPartnerResponseBodyModuleAppServiceListOperationAddress struct {
	// List of console operation buttons
	Actions []*GetAppInstanceForPartnerResponseBodyModuleAppServiceListOperationAddressActions `json:"Actions,omitempty" xml:"Actions,omitempty" type:"Repeated"`
	// AI customer service design URL
	//
	// example:
	//
	// AiCustomerConfigUrl
	AiCustomerConfigUrl *string `json:"AiCustomerConfigUrl,omitempty" xml:"AiCustomerConfigUrl,omitempty"`
	// AI design URL
	//
	// example:
	//
	// AiDesignUrl
	AiDesignUrl *string `json:"AiDesignUrl,omitempty" xml:"AiDesignUrl,omitempty"`
	// Application publish URL
	//
	// example:
	//
	// AppPublishUrl
	AppPublishUrl *string `json:"AppPublishUrl,omitempty" xml:"AppPublishUrl,omitempty"`
	// List of overview page operation buttons
	DashboardActions []*GetAppInstanceForPartnerResponseBodyModuleAppServiceListOperationAddressDashboardActions `json:"DashboardActions,omitempty" xml:"DashboardActions,omitempty" type:"Repeated"`
	// Professional design URL
	//
	// example:
	//
	// DesignUrl
	DesignUrl *string `json:"DesignUrl,omitempty" xml:"DesignUrl,omitempty"`
	// Application instance logon URL
	//
	// example:
	//
	// InstanceLoginUrl
	InstanceLoginUrl *string `json:"InstanceLoginUrl,omitempty" xml:"InstanceLoginUrl,omitempty"`
	// Renewal and purchase URL
	//
	// example:
	//
	// renewBuyUrl
	RenewBuyUrl *string `json:"RenewBuyUrl,omitempty" xml:"RenewBuyUrl,omitempty"`
	// Service delivery URL
	//
	// example:
	//
	// ServerDeliveryUrl
	ServerDeliveryUrl *string `json:"ServerDeliveryUrl,omitempty" xml:"ServerDeliveryUrl,omitempty"`
	// Upgrade purchase URL
	//
	// example:
	//
	// UpgradeBuyUrl
	UpgradeBuyUrl *string `json:"UpgradeBuyUrl,omitempty" xml:"UpgradeBuyUrl,omitempty"`
}

func (s GetAppInstanceForPartnerResponseBodyModuleAppServiceListOperationAddress) String() string {
	return dara.Prettify(s)
}

func (s GetAppInstanceForPartnerResponseBodyModuleAppServiceListOperationAddress) GoString() string {
	return s.String()
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppServiceListOperationAddress) GetActions() []*GetAppInstanceForPartnerResponseBodyModuleAppServiceListOperationAddressActions {
	return s.Actions
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppServiceListOperationAddress) GetAiCustomerConfigUrl() *string {
	return s.AiCustomerConfigUrl
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppServiceListOperationAddress) GetAiDesignUrl() *string {
	return s.AiDesignUrl
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppServiceListOperationAddress) GetAppPublishUrl() *string {
	return s.AppPublishUrl
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppServiceListOperationAddress) GetDashboardActions() []*GetAppInstanceForPartnerResponseBodyModuleAppServiceListOperationAddressDashboardActions {
	return s.DashboardActions
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppServiceListOperationAddress) GetDesignUrl() *string {
	return s.DesignUrl
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppServiceListOperationAddress) GetInstanceLoginUrl() *string {
	return s.InstanceLoginUrl
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppServiceListOperationAddress) GetRenewBuyUrl() *string {
	return s.RenewBuyUrl
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppServiceListOperationAddress) GetServerDeliveryUrl() *string {
	return s.ServerDeliveryUrl
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppServiceListOperationAddress) GetUpgradeBuyUrl() *string {
	return s.UpgradeBuyUrl
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppServiceListOperationAddress) SetActions(v []*GetAppInstanceForPartnerResponseBodyModuleAppServiceListOperationAddressActions) *GetAppInstanceForPartnerResponseBodyModuleAppServiceListOperationAddress {
	s.Actions = v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppServiceListOperationAddress) SetAiCustomerConfigUrl(v string) *GetAppInstanceForPartnerResponseBodyModuleAppServiceListOperationAddress {
	s.AiCustomerConfigUrl = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppServiceListOperationAddress) SetAiDesignUrl(v string) *GetAppInstanceForPartnerResponseBodyModuleAppServiceListOperationAddress {
	s.AiDesignUrl = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppServiceListOperationAddress) SetAppPublishUrl(v string) *GetAppInstanceForPartnerResponseBodyModuleAppServiceListOperationAddress {
	s.AppPublishUrl = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppServiceListOperationAddress) SetDashboardActions(v []*GetAppInstanceForPartnerResponseBodyModuleAppServiceListOperationAddressDashboardActions) *GetAppInstanceForPartnerResponseBodyModuleAppServiceListOperationAddress {
	s.DashboardActions = v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppServiceListOperationAddress) SetDesignUrl(v string) *GetAppInstanceForPartnerResponseBodyModuleAppServiceListOperationAddress {
	s.DesignUrl = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppServiceListOperationAddress) SetInstanceLoginUrl(v string) *GetAppInstanceForPartnerResponseBodyModuleAppServiceListOperationAddress {
	s.InstanceLoginUrl = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppServiceListOperationAddress) SetRenewBuyUrl(v string) *GetAppInstanceForPartnerResponseBodyModuleAppServiceListOperationAddress {
	s.RenewBuyUrl = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppServiceListOperationAddress) SetServerDeliveryUrl(v string) *GetAppInstanceForPartnerResponseBodyModuleAppServiceListOperationAddress {
	s.ServerDeliveryUrl = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppServiceListOperationAddress) SetUpgradeBuyUrl(v string) *GetAppInstanceForPartnerResponseBodyModuleAppServiceListOperationAddress {
	s.UpgradeBuyUrl = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppServiceListOperationAddress) Validate() error {
	if s.Actions != nil {
		for _, item := range s.Actions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.DashboardActions != nil {
		for _, item := range s.DashboardActions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetAppInstanceForPartnerResponseBodyModuleAppServiceListOperationAddressActions struct {
	// Action key, used as a UUID for an action behavior.
	//
	// example:
	//
	// CollectedNumber
	ActionKey *string `json:"ActionKey,omitempty" xml:"ActionKey,omitempty"`
	// Display text for the action, used to show the action name in the interface.
	//
	// example:
	//
	// ActionText
	ActionText *string `json:"ActionText,omitempty" xml:"ActionText,omitempty"`
	// Enable status, indicating whether the action is active.
	//
	// example:
	//
	// false
	Enable *bool `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// Link address, the URL to which the user is redirected when clicking the action.
	//
	// example:
	//
	// Href
	Href *string `json:"Href,omitempty" xml:"Href,omitempty"`
}

func (s GetAppInstanceForPartnerResponseBodyModuleAppServiceListOperationAddressActions) String() string {
	return dara.Prettify(s)
}

func (s GetAppInstanceForPartnerResponseBodyModuleAppServiceListOperationAddressActions) GoString() string {
	return s.String()
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppServiceListOperationAddressActions) GetActionKey() *string {
	return s.ActionKey
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppServiceListOperationAddressActions) GetActionText() *string {
	return s.ActionText
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppServiceListOperationAddressActions) GetEnable() *bool {
	return s.Enable
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppServiceListOperationAddressActions) GetHref() *string {
	return s.Href
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppServiceListOperationAddressActions) SetActionKey(v string) *GetAppInstanceForPartnerResponseBodyModuleAppServiceListOperationAddressActions {
	s.ActionKey = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppServiceListOperationAddressActions) SetActionText(v string) *GetAppInstanceForPartnerResponseBodyModuleAppServiceListOperationAddressActions {
	s.ActionText = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppServiceListOperationAddressActions) SetEnable(v bool) *GetAppInstanceForPartnerResponseBodyModuleAppServiceListOperationAddressActions {
	s.Enable = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppServiceListOperationAddressActions) SetHref(v string) *GetAppInstanceForPartnerResponseBodyModuleAppServiceListOperationAddressActions {
	s.Href = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppServiceListOperationAddressActions) Validate() error {
	return dara.Validate(s)
}

type GetAppInstanceForPartnerResponseBodyModuleAppServiceListOperationAddressDashboardActions struct {
	// Operation key used as a UUID for an operation behavior
	//
	// example:
	//
	// AbortDialogue
	ActionKey *string `json:"ActionKey,omitempty" xml:"ActionKey,omitempty"`
	// The display text of the action, used to show the action name in the interface.
	//
	// example:
	//
	// ActionText
	ActionText *string `json:"ActionText,omitempty" xml:"ActionText,omitempty"`
	// The enable status, indicating whether the action is active.
	//
	// example:
	//
	// true
	Enable *bool `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// The URL address to which the hyperlink redirects when the action is clicked.
	//
	// example:
	//
	// Href
	Href *string `json:"Href,omitempty" xml:"Href,omitempty"`
}

func (s GetAppInstanceForPartnerResponseBodyModuleAppServiceListOperationAddressDashboardActions) String() string {
	return dara.Prettify(s)
}

func (s GetAppInstanceForPartnerResponseBodyModuleAppServiceListOperationAddressDashboardActions) GoString() string {
	return s.String()
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppServiceListOperationAddressDashboardActions) GetActionKey() *string {
	return s.ActionKey
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppServiceListOperationAddressDashboardActions) GetActionText() *string {
	return s.ActionText
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppServiceListOperationAddressDashboardActions) GetEnable() *bool {
	return s.Enable
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppServiceListOperationAddressDashboardActions) GetHref() *string {
	return s.Href
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppServiceListOperationAddressDashboardActions) SetActionKey(v string) *GetAppInstanceForPartnerResponseBodyModuleAppServiceListOperationAddressDashboardActions {
	s.ActionKey = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppServiceListOperationAddressDashboardActions) SetActionText(v string) *GetAppInstanceForPartnerResponseBodyModuleAppServiceListOperationAddressDashboardActions {
	s.ActionText = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppServiceListOperationAddressDashboardActions) SetEnable(v bool) *GetAppInstanceForPartnerResponseBodyModuleAppServiceListOperationAddressDashboardActions {
	s.Enable = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppServiceListOperationAddressDashboardActions) SetHref(v string) *GetAppInstanceForPartnerResponseBodyModuleAppServiceListOperationAddressDashboardActions {
	s.Href = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppServiceListOperationAddressDashboardActions) Validate() error {
	return dara.Validate(s)
}

type GetAppInstanceForPartnerResponseBodyModuleAppServiceListProfile struct {
	// Business ID
	//
	// example:
	//
	// WS20250801003834000003
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// Design Type
	//
	// example:
	//
	// DesignType
	DesignType *string `json:"DesignType,omitempty" xml:"DesignType,omitempty"`
	// Design type description
	//
	// example:
	//
	// DesignTypeText
	DesignTypeText *string `json:"DesignTypeText,omitempty" xml:"DesignTypeText,omitempty"`
	// Creation Time
	//
	// example:
	//
	// 1620711265000
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// Updated At
	//
	// example:
	//
	// Mon Sep 25 14:48:49 CST 2023
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// primary key
	//
	// example:
	//
	// 108232
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// Lingxiao instance ID
	//
	// example:
	//
	// LxInstanceId
	LxInstanceId *string `json:"LxInstanceId,omitempty" xml:"LxInstanceId,omitempty"`
	// order ID
	//
	// example:
	//
	// 256146659280026
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// Service Specification
	//
	// example:
	//
	// 8
	ServiceSpec *string `json:"ServiceSpec,omitempty" xml:"ServiceSpec,omitempty"`
	// Service specification description
	//
	// example:
	//
	// ServiceSpecText
	ServiceSpecText *string `json:"ServiceSpecText,omitempty" xml:"ServiceSpecText,omitempty"`
	UpgradeStatus   *string `json:"UpgradeStatus,omitempty" xml:"UpgradeStatus,omitempty"`
}

func (s GetAppInstanceForPartnerResponseBodyModuleAppServiceListProfile) String() string {
	return dara.Prettify(s)
}

func (s GetAppInstanceForPartnerResponseBodyModuleAppServiceListProfile) GoString() string {
	return s.String()
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppServiceListProfile) GetBizId() *string {
	return s.BizId
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppServiceListProfile) GetDesignType() *string {
	return s.DesignType
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppServiceListProfile) GetDesignTypeText() *string {
	return s.DesignTypeText
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppServiceListProfile) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppServiceListProfile) GetGmtModified() *string {
	return s.GmtModified
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppServiceListProfile) GetId() *string {
	return s.Id
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppServiceListProfile) GetLxInstanceId() *string {
	return s.LxInstanceId
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppServiceListProfile) GetOrderId() *string {
	return s.OrderId
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppServiceListProfile) GetServiceSpec() *string {
	return s.ServiceSpec
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppServiceListProfile) GetServiceSpecText() *string {
	return s.ServiceSpecText
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppServiceListProfile) GetUpgradeStatus() *string {
	return s.UpgradeStatus
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppServiceListProfile) SetBizId(v string) *GetAppInstanceForPartnerResponseBodyModuleAppServiceListProfile {
	s.BizId = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppServiceListProfile) SetDesignType(v string) *GetAppInstanceForPartnerResponseBodyModuleAppServiceListProfile {
	s.DesignType = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppServiceListProfile) SetDesignTypeText(v string) *GetAppInstanceForPartnerResponseBodyModuleAppServiceListProfile {
	s.DesignTypeText = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppServiceListProfile) SetGmtCreate(v string) *GetAppInstanceForPartnerResponseBodyModuleAppServiceListProfile {
	s.GmtCreate = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppServiceListProfile) SetGmtModified(v string) *GetAppInstanceForPartnerResponseBodyModuleAppServiceListProfile {
	s.GmtModified = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppServiceListProfile) SetId(v string) *GetAppInstanceForPartnerResponseBodyModuleAppServiceListProfile {
	s.Id = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppServiceListProfile) SetLxInstanceId(v string) *GetAppInstanceForPartnerResponseBodyModuleAppServiceListProfile {
	s.LxInstanceId = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppServiceListProfile) SetOrderId(v string) *GetAppInstanceForPartnerResponseBodyModuleAppServiceListProfile {
	s.OrderId = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppServiceListProfile) SetServiceSpec(v string) *GetAppInstanceForPartnerResponseBodyModuleAppServiceListProfile {
	s.ServiceSpec = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppServiceListProfile) SetServiceSpecText(v string) *GetAppInstanceForPartnerResponseBodyModuleAppServiceListProfile {
	s.ServiceSpecText = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppServiceListProfile) SetUpgradeStatus(v string) *GetAppInstanceForPartnerResponseBodyModuleAppServiceListProfile {
	s.UpgradeStatus = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleAppServiceListProfile) Validate() error {
	return dara.Validate(s)
}

type GetAppInstanceForPartnerResponseBodyModulePartnerDetail struct {
	// data
	BindData *GetAppInstanceForPartnerResponseBodyModulePartnerDetailBindData `json:"BindData,omitempty" xml:"BindData,omitempty" type:"Struct"`
	// Partner ID, extension parameter
	//
	// example:
	//
	// 10001
	PartnerId *string `json:"PartnerId,omitempty" xml:"PartnerId,omitempty"`
	// Possible values: unknown, init, testing, online
	//
	// example:
	//
	// 200
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetAppInstanceForPartnerResponseBodyModulePartnerDetail) String() string {
	return dara.Prettify(s)
}

func (s GetAppInstanceForPartnerResponseBodyModulePartnerDetail) GoString() string {
	return s.String()
}

func (s *GetAppInstanceForPartnerResponseBodyModulePartnerDetail) GetBindData() *GetAppInstanceForPartnerResponseBodyModulePartnerDetailBindData {
	return s.BindData
}

func (s *GetAppInstanceForPartnerResponseBodyModulePartnerDetail) GetPartnerId() *string {
	return s.PartnerId
}

func (s *GetAppInstanceForPartnerResponseBodyModulePartnerDetail) GetStatus() *string {
	return s.Status
}

func (s *GetAppInstanceForPartnerResponseBodyModulePartnerDetail) SetBindData(v *GetAppInstanceForPartnerResponseBodyModulePartnerDetailBindData) *GetAppInstanceForPartnerResponseBodyModulePartnerDetail {
	s.BindData = v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModulePartnerDetail) SetPartnerId(v string) *GetAppInstanceForPartnerResponseBodyModulePartnerDetail {
	s.PartnerId = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModulePartnerDetail) SetStatus(v string) *GetAppInstanceForPartnerResponseBodyModulePartnerDetail {
	s.Status = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModulePartnerDetail) Validate() error {
	if s.BindData != nil {
		if err := s.BindData.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAppInstanceForPartnerResponseBodyModulePartnerDetailBindData struct {
	// aliyun_pk
	//
	// example:
	//
	// ***
	AliyunPk *string `json:"AliyunPk,omitempty" xml:"AliyunPk,omitempty"`
	// External bizId
	//
	// example:
	//
	// WS20250801153127000002
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// Creation Time
	//
	// example:
	//
	// 1672123722000
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// Updated At
	//
	// example:
	//
	// Wed Nov 26 10:15:28 CST 2025
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// User who owns the ICP filing service code
	//
	// example:
	//
	// grantAliyunPk
	GrantAliyunPk *string `json:"GrantAliyunPk,omitempty" xml:"GrantAliyunPk,omitempty"`
	// Phone number
	//
	// example:
	//
	// ***********
	Mobile *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	// parent_pk
	//
	// example:
	//
	// 1123213
	ParentPk *string `json:"ParentPk,omitempty" xml:"ParentPk,omitempty"`
	// Externally associated partnerId
	//
	// example:
	//
	// 10001
	PartnerId *string `json:"PartnerId,omitempty" xml:"PartnerId,omitempty"`
	// Display name of the attached user
	//
	// example:
	//
	// diaplayName
	UserDisplayName *string `json:"UserDisplayName,omitempty" xml:"UserDisplayName,omitempty"`
}

func (s GetAppInstanceForPartnerResponseBodyModulePartnerDetailBindData) String() string {
	return dara.Prettify(s)
}

func (s GetAppInstanceForPartnerResponseBodyModulePartnerDetailBindData) GoString() string {
	return s.String()
}

func (s *GetAppInstanceForPartnerResponseBodyModulePartnerDetailBindData) GetAliyunPk() *string {
	return s.AliyunPk
}

func (s *GetAppInstanceForPartnerResponseBodyModulePartnerDetailBindData) GetBizId() *string {
	return s.BizId
}

func (s *GetAppInstanceForPartnerResponseBodyModulePartnerDetailBindData) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *GetAppInstanceForPartnerResponseBodyModulePartnerDetailBindData) GetGmtModified() *string {
	return s.GmtModified
}

func (s *GetAppInstanceForPartnerResponseBodyModulePartnerDetailBindData) GetGrantAliyunPk() *string {
	return s.GrantAliyunPk
}

func (s *GetAppInstanceForPartnerResponseBodyModulePartnerDetailBindData) GetMobile() *string {
	return s.Mobile
}

func (s *GetAppInstanceForPartnerResponseBodyModulePartnerDetailBindData) GetParentPk() *string {
	return s.ParentPk
}

func (s *GetAppInstanceForPartnerResponseBodyModulePartnerDetailBindData) GetPartnerId() *string {
	return s.PartnerId
}

func (s *GetAppInstanceForPartnerResponseBodyModulePartnerDetailBindData) GetUserDisplayName() *string {
	return s.UserDisplayName
}

func (s *GetAppInstanceForPartnerResponseBodyModulePartnerDetailBindData) SetAliyunPk(v string) *GetAppInstanceForPartnerResponseBodyModulePartnerDetailBindData {
	s.AliyunPk = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModulePartnerDetailBindData) SetBizId(v string) *GetAppInstanceForPartnerResponseBodyModulePartnerDetailBindData {
	s.BizId = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModulePartnerDetailBindData) SetGmtCreate(v string) *GetAppInstanceForPartnerResponseBodyModulePartnerDetailBindData {
	s.GmtCreate = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModulePartnerDetailBindData) SetGmtModified(v string) *GetAppInstanceForPartnerResponseBodyModulePartnerDetailBindData {
	s.GmtModified = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModulePartnerDetailBindData) SetGrantAliyunPk(v string) *GetAppInstanceForPartnerResponseBodyModulePartnerDetailBindData {
	s.GrantAliyunPk = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModulePartnerDetailBindData) SetMobile(v string) *GetAppInstanceForPartnerResponseBodyModulePartnerDetailBindData {
	s.Mobile = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModulePartnerDetailBindData) SetParentPk(v string) *GetAppInstanceForPartnerResponseBodyModulePartnerDetailBindData {
	s.ParentPk = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModulePartnerDetailBindData) SetPartnerId(v string) *GetAppInstanceForPartnerResponseBodyModulePartnerDetailBindData {
	s.PartnerId = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModulePartnerDetailBindData) SetUserDisplayName(v string) *GetAppInstanceForPartnerResponseBodyModulePartnerDetailBindData {
	s.UserDisplayName = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModulePartnerDetailBindData) Validate() error {
	return dara.Validate(s)
}

type GetAppInstanceForPartnerResponseBodyModuleProfile struct {
	// Management console URL
	//
	// example:
	//
	// placeHolder
	AdminUrl *string `json:"AdminUrl,omitempty" xml:"AdminUrl,omitempty"`
	// Application type code
	//
	// example:
	//
	// PC_WebSite
	ApplicationType *string `json:"ApplicationType,omitempty" xml:"ApplicationType,omitempty"`
	// Application type text description
	//
	// example:
	//
	// placeHolder
	ApplicationTypeText *string `json:"ApplicationTypeText,omitempty" xml:"ApplicationTypeText,omitempty"`
	// Attached CNAME record
	//
	// example:
	//
	// placeHolder
	BindCname *string `json:"BindCname,omitempty" xml:"BindCname,omitempty"`
	// Business ID
	//
	// example:
	//
	// WS20250801152128000005
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// Product encoding
	//
	// example:
	//
	// rds
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	// Customer Service Information
	//
	// example:
	//
	// placeHolder
	CustomerService *string `json:"CustomerService,omitempty" xml:"CustomerService,omitempty"`
	// Deployment area
	//
	// example:
	//
	// ChineseMainland
	DeployArea *string `json:"DeployArea,omitempty" xml:"DeployArea,omitempty"`
	// Domain name list
	//
	// example:
	//
	// [white:],*.eduresource.cn,*.dingtalk.com,*.aliyun.com,*.aliyuncs.com,euser.edu-aliyun.com,s-gm.mmstat.com
	DomainList *string `json:"DomainList,omitempty" xml:"DomainList,omitempty"`
	// Editor URL
	//
	// example:
	//
	// placeHolder
	EditorUrl *string `json:"EditorUrl,omitempty" xml:"EditorUrl,omitempty"`
	// Creation time
	//
	// example:
	//
	// 1605280632000
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// Updated At
	//
	// example:
	//
	// 2024-05-06T03:07:45Z
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// ICP number
	//
	// example:
	//
	// placeHolder
	IcpbaNo *string `json:"IcpbaNo,omitempty" xml:"IcpbaNo,omitempty"`
	// Image list
	//
	// example:
	//
	// {\\"Image\\": []}
	ImageList *string `json:"ImageList,omitempty" xml:"ImageList,omitempty"`
	// Lenovo instance ID
	//
	// example:
	//
	// placeHolder
	LxInstanceId *string `json:"LxInstanceId,omitempty" xml:"LxInstanceId,omitempty"`
	// Order time
	//
	// example:
	//
	// placeHolder
	OrdTime *string `json:"OrdTime,omitempty" xml:"OrdTime,omitempty"`
	// Order ID
	//
	// example:
	//
	// 256146659280026
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// Order quantity
	//
	// example:
	//
	// 1
	OrderNum *int32 `json:"OrderNum,omitempty" xml:"OrderNum,omitempty"`
	// Channel partner business ID
	//
	// example:
	//
	// 100086
	PartnerId *string `json:"PartnerId,omitempty" xml:"PartnerId,omitempty"`
	// Payment time
	//
	// example:
	//
	// placeHolder
	PayTime *string `json:"PayTime,omitempty" xml:"PayTime,omitempty"`
	// Preview URL
	//
	// example:
	//
	// https://preview-lyj.aliyuncs.com/preview/1daacb3ebbb8435d9091fb950c528d0f?subSceneIds=682185
	PreviewUrl *string `json:"PreviewUrl,omitempty" xml:"PreviewUrl,omitempty"`
	// SEO site information
	//
	// example:
	//
	// placeHolder
	SeoSite *string `json:"SeoSite,omitempty" xml:"SeoSite,omitempty"`
	// Site logo
	//
	// example:
	//
	// placeHolder
	SiteLogo *string `json:"SiteLogo,omitempty" xml:"SiteLogo,omitempty"`
	// Site version encoding
	//
	// example:
	//
	// Basic_Edition
	SiteVersion *string `json:"SiteVersion,omitempty" xml:"SiteVersion,omitempty"`
	// Site version text description
	//
	// example:
	//
	// placeHolder
	SiteVersionText *string `json:"SiteVersionText,omitempty" xml:"SiteVersionText,omitempty"`
	// Source information
	//
	// example:
	//
	// edasmsc
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// Template ETag identity
	//
	// example:
	//
	// placeHolder
	TemplateEtag *string `json:"TemplateEtag,omitempty" xml:"TemplateEtag,omitempty"`
	// Template ID
	//
	// example:
	//
	// 4644
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// Text list
	//
	// example:
	//
	// placeHolder
	TextList *string `json:"TextList,omitempty" xml:"TextList,omitempty"`
	// Thumbnail
	//
	// example:
	//
	// https://aloss-recruit-aiinterview.oss-cn-beijing.aliyuncs.com/thumbnail-1753495551714-10000.png
	Thumbnail     *string `json:"Thumbnail,omitempty" xml:"Thumbnail,omitempty"`
	UpgradeStatus *string `json:"UpgradeStatus,omitempty" xml:"UpgradeStatus,omitempty"`
}

func (s GetAppInstanceForPartnerResponseBodyModuleProfile) String() string {
	return dara.Prettify(s)
}

func (s GetAppInstanceForPartnerResponseBodyModuleProfile) GoString() string {
	return s.String()
}

func (s *GetAppInstanceForPartnerResponseBodyModuleProfile) GetAdminUrl() *string {
	return s.AdminUrl
}

func (s *GetAppInstanceForPartnerResponseBodyModuleProfile) GetApplicationType() *string {
	return s.ApplicationType
}

func (s *GetAppInstanceForPartnerResponseBodyModuleProfile) GetApplicationTypeText() *string {
	return s.ApplicationTypeText
}

func (s *GetAppInstanceForPartnerResponseBodyModuleProfile) GetBindCname() *string {
	return s.BindCname
}

func (s *GetAppInstanceForPartnerResponseBodyModuleProfile) GetBizId() *string {
	return s.BizId
}

func (s *GetAppInstanceForPartnerResponseBodyModuleProfile) GetCommodityCode() *string {
	return s.CommodityCode
}

func (s *GetAppInstanceForPartnerResponseBodyModuleProfile) GetCustomerService() *string {
	return s.CustomerService
}

func (s *GetAppInstanceForPartnerResponseBodyModuleProfile) GetDeployArea() *string {
	return s.DeployArea
}

func (s *GetAppInstanceForPartnerResponseBodyModuleProfile) GetDomainList() *string {
	return s.DomainList
}

func (s *GetAppInstanceForPartnerResponseBodyModuleProfile) GetEditorUrl() *string {
	return s.EditorUrl
}

func (s *GetAppInstanceForPartnerResponseBodyModuleProfile) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *GetAppInstanceForPartnerResponseBodyModuleProfile) GetGmtModified() *string {
	return s.GmtModified
}

func (s *GetAppInstanceForPartnerResponseBodyModuleProfile) GetIcpbaNo() *string {
	return s.IcpbaNo
}

func (s *GetAppInstanceForPartnerResponseBodyModuleProfile) GetImageList() *string {
	return s.ImageList
}

func (s *GetAppInstanceForPartnerResponseBodyModuleProfile) GetLxInstanceId() *string {
	return s.LxInstanceId
}

func (s *GetAppInstanceForPartnerResponseBodyModuleProfile) GetOrdTime() *string {
	return s.OrdTime
}

func (s *GetAppInstanceForPartnerResponseBodyModuleProfile) GetOrderId() *string {
	return s.OrderId
}

func (s *GetAppInstanceForPartnerResponseBodyModuleProfile) GetOrderNum() *int32 {
	return s.OrderNum
}

func (s *GetAppInstanceForPartnerResponseBodyModuleProfile) GetPartnerId() *string {
	return s.PartnerId
}

func (s *GetAppInstanceForPartnerResponseBodyModuleProfile) GetPayTime() *string {
	return s.PayTime
}

func (s *GetAppInstanceForPartnerResponseBodyModuleProfile) GetPreviewUrl() *string {
	return s.PreviewUrl
}

func (s *GetAppInstanceForPartnerResponseBodyModuleProfile) GetSeoSite() *string {
	return s.SeoSite
}

func (s *GetAppInstanceForPartnerResponseBodyModuleProfile) GetSiteLogo() *string {
	return s.SiteLogo
}

func (s *GetAppInstanceForPartnerResponseBodyModuleProfile) GetSiteVersion() *string {
	return s.SiteVersion
}

func (s *GetAppInstanceForPartnerResponseBodyModuleProfile) GetSiteVersionText() *string {
	return s.SiteVersionText
}

func (s *GetAppInstanceForPartnerResponseBodyModuleProfile) GetSource() *string {
	return s.Source
}

func (s *GetAppInstanceForPartnerResponseBodyModuleProfile) GetTemplateEtag() *string {
	return s.TemplateEtag
}

func (s *GetAppInstanceForPartnerResponseBodyModuleProfile) GetTemplateId() *string {
	return s.TemplateId
}

func (s *GetAppInstanceForPartnerResponseBodyModuleProfile) GetTextList() *string {
	return s.TextList
}

func (s *GetAppInstanceForPartnerResponseBodyModuleProfile) GetThumbnail() *string {
	return s.Thumbnail
}

func (s *GetAppInstanceForPartnerResponseBodyModuleProfile) GetUpgradeStatus() *string {
	return s.UpgradeStatus
}

func (s *GetAppInstanceForPartnerResponseBodyModuleProfile) SetAdminUrl(v string) *GetAppInstanceForPartnerResponseBodyModuleProfile {
	s.AdminUrl = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleProfile) SetApplicationType(v string) *GetAppInstanceForPartnerResponseBodyModuleProfile {
	s.ApplicationType = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleProfile) SetApplicationTypeText(v string) *GetAppInstanceForPartnerResponseBodyModuleProfile {
	s.ApplicationTypeText = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleProfile) SetBindCname(v string) *GetAppInstanceForPartnerResponseBodyModuleProfile {
	s.BindCname = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleProfile) SetBizId(v string) *GetAppInstanceForPartnerResponseBodyModuleProfile {
	s.BizId = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleProfile) SetCommodityCode(v string) *GetAppInstanceForPartnerResponseBodyModuleProfile {
	s.CommodityCode = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleProfile) SetCustomerService(v string) *GetAppInstanceForPartnerResponseBodyModuleProfile {
	s.CustomerService = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleProfile) SetDeployArea(v string) *GetAppInstanceForPartnerResponseBodyModuleProfile {
	s.DeployArea = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleProfile) SetDomainList(v string) *GetAppInstanceForPartnerResponseBodyModuleProfile {
	s.DomainList = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleProfile) SetEditorUrl(v string) *GetAppInstanceForPartnerResponseBodyModuleProfile {
	s.EditorUrl = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleProfile) SetGmtCreate(v string) *GetAppInstanceForPartnerResponseBodyModuleProfile {
	s.GmtCreate = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleProfile) SetGmtModified(v string) *GetAppInstanceForPartnerResponseBodyModuleProfile {
	s.GmtModified = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleProfile) SetIcpbaNo(v string) *GetAppInstanceForPartnerResponseBodyModuleProfile {
	s.IcpbaNo = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleProfile) SetImageList(v string) *GetAppInstanceForPartnerResponseBodyModuleProfile {
	s.ImageList = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleProfile) SetLxInstanceId(v string) *GetAppInstanceForPartnerResponseBodyModuleProfile {
	s.LxInstanceId = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleProfile) SetOrdTime(v string) *GetAppInstanceForPartnerResponseBodyModuleProfile {
	s.OrdTime = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleProfile) SetOrderId(v string) *GetAppInstanceForPartnerResponseBodyModuleProfile {
	s.OrderId = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleProfile) SetOrderNum(v int32) *GetAppInstanceForPartnerResponseBodyModuleProfile {
	s.OrderNum = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleProfile) SetPartnerId(v string) *GetAppInstanceForPartnerResponseBodyModuleProfile {
	s.PartnerId = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleProfile) SetPayTime(v string) *GetAppInstanceForPartnerResponseBodyModuleProfile {
	s.PayTime = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleProfile) SetPreviewUrl(v string) *GetAppInstanceForPartnerResponseBodyModuleProfile {
	s.PreviewUrl = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleProfile) SetSeoSite(v string) *GetAppInstanceForPartnerResponseBodyModuleProfile {
	s.SeoSite = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleProfile) SetSiteLogo(v string) *GetAppInstanceForPartnerResponseBodyModuleProfile {
	s.SiteLogo = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleProfile) SetSiteVersion(v string) *GetAppInstanceForPartnerResponseBodyModuleProfile {
	s.SiteVersion = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleProfile) SetSiteVersionText(v string) *GetAppInstanceForPartnerResponseBodyModuleProfile {
	s.SiteVersionText = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleProfile) SetSource(v string) *GetAppInstanceForPartnerResponseBodyModuleProfile {
	s.Source = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleProfile) SetTemplateEtag(v string) *GetAppInstanceForPartnerResponseBodyModuleProfile {
	s.TemplateEtag = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleProfile) SetTemplateId(v string) *GetAppInstanceForPartnerResponseBodyModuleProfile {
	s.TemplateId = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleProfile) SetTextList(v string) *GetAppInstanceForPartnerResponseBodyModuleProfile {
	s.TextList = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleProfile) SetThumbnail(v string) *GetAppInstanceForPartnerResponseBodyModuleProfile {
	s.Thumbnail = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleProfile) SetUpgradeStatus(v string) *GetAppInstanceForPartnerResponseBodyModuleProfile {
	s.UpgradeStatus = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleProfile) Validate() error {
	return dara.Validate(s)
}

type GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceList struct {
	// application sub-type
	//
	// example:
	//
	// placeHolder
	AppSubType *string `json:"AppSubType,omitempty" xml:"AppSubType,omitempty"`
	// application Type
	//
	// example:
	//
	// 1
	AppType *string `json:"AppType,omitempty" xml:"AppType,omitempty"`
	// application instance business ID
	//
	// example:
	//
	// WS20250929173805000001
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// Build Type
	//
	// example:
	//
	// placeHolder
	BuildType *string `json:"BuildType,omitempty" xml:"BuildType,omitempty"`
	// Is Deleted
	//
	// example:
	//
	// fase
	Deleted *int32 `json:"Deleted,omitempty" xml:"Deleted,omitempty"`
	// Application description
	//
	// example:
	//
	// base
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// placeHolder
	//
	// example:
	//
	// placeHolder
	DesignSpecBizId *string `json:"DesignSpecBizId,omitempty" xml:"DesignSpecBizId,omitempty"`
	// Design configuration ID
	//
	// example:
	//
	// placeHolder
	DesignSpecId *string `json:"DesignSpecId,omitempty" xml:"DesignSpecId,omitempty"`
	// primary domain name
	//
	// example:
	//
	// shikuntech.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// End time
	//
	// example:
	//
	// 2025-05-23T15:01:25.891Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// Environment
	//
	// example:
	//
	// pre
	Env *string `json:"Env,omitempty" xml:"Env,omitempty"`
	// Associated Delivery Order ID
	//
	// example:
	//
	// placeHolder
	EspBizId *string `json:"EspBizId,omitempty" xml:"EspBizId,omitempty"`
	// Creation Time
	//
	// example:
	//
	// 2025-06-19T07:39:55Z
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// Deletion time
	//
	// example:
	//
	// placeHolder
	GmtDelete *string `json:"GmtDelete,omitempty" xml:"GmtDelete,omitempty"`
	// Updated At
	//
	// example:
	//
	// 2025-01-14T09:09:57Z
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// Published At
	//
	// example:
	//
	// placeHolder
	GmtPublish *string `json:"GmtPublish,omitempty" xml:"GmtPublish,omitempty"`
	// application icon
	//
	// example:
	//
	// https://app-center-icon-prod-shanghai.oss-cn-shanghai.aliyuncs.com/tenant/1617863868712071/1749090558651_%E9%AB%98%E5%BE%B7.png
	IconUrl *string `json:"IconUrl,omitempty" xml:"IconUrl,omitempty"`
	// Application Name
	//
	// example:
	//
	// 19609820.png
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// order ID
	//
	// example:
	//
	// 264987642530867,264988827010867,264982589530867,264985058640867,264982290930867,264982387520867,264987147460867,264985752990867,264988432850867
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// Instance configuration
	Profile *GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceListProfile `json:"Profile,omitempty" xml:"Profile,omitempty" type:"Struct"`
	// domain name
	//
	// example:
	//
	// placeHolder
	SiteHost *string `json:"SiteHost,omitempty" xml:"SiteHost,omitempty"`
	// application identity
	//
	// example:
	//
	// placeHolder
	Slug *string `json:"Slug,omitempty" xml:"Slug,omitempty"`
	// Source MARKET_CLOUD_DREAM
	//
	// example:
	//
	// MARKET_CLOUD_DREAM
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// Start Time
	//
	// example:
	//
	// 2025-07-22T02:23:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// Status
	//
	// example:
	//
	// success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Task Status text.
	//
	// example:
	//
	// SUCCESS
	StatusText *string `json:"StatusText,omitempty" xml:"StatusText,omitempty"`
	// Application thumbnail
	//
	// example:
	//
	// placeHolder
	ThumbnailUrl *string `json:"ThumbnailUrl,omitempty" xml:"ThumbnailUrl,omitempty"`
	// User ID
	//
	// example:
	//
	// placeHolder
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// Application instance Version
	//
	// example:
	//
	// 2019-04-02
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceList) String() string {
	return dara.Prettify(s)
}

func (s GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceList) GoString() string {
	return s.String()
}

func (s *GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceList) GetAppSubType() *string {
	return s.AppSubType
}

func (s *GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceList) GetAppType() *string {
	return s.AppType
}

func (s *GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceList) GetBizId() *string {
	return s.BizId
}

func (s *GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceList) GetBuildType() *string {
	return s.BuildType
}

func (s *GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceList) GetDeleted() *int32 {
	return s.Deleted
}

func (s *GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceList) GetDescription() *string {
	return s.Description
}

func (s *GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceList) GetDesignSpecBizId() *string {
	return s.DesignSpecBizId
}

func (s *GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceList) GetDesignSpecId() *string {
	return s.DesignSpecId
}

func (s *GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceList) GetDomain() *string {
	return s.Domain
}

func (s *GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceList) GetEndTime() *string {
	return s.EndTime
}

func (s *GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceList) GetEnv() *string {
	return s.Env
}

func (s *GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceList) GetEspBizId() *string {
	return s.EspBizId
}

func (s *GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceList) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceList) GetGmtDelete() *string {
	return s.GmtDelete
}

func (s *GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceList) GetGmtModified() *string {
	return s.GmtModified
}

func (s *GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceList) GetGmtPublish() *string {
	return s.GmtPublish
}

func (s *GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceList) GetIconUrl() *string {
	return s.IconUrl
}

func (s *GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceList) GetName() *string {
	return s.Name
}

func (s *GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceList) GetOrderId() *string {
	return s.OrderId
}

func (s *GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceList) GetProfile() *GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceListProfile {
	return s.Profile
}

func (s *GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceList) GetSiteHost() *string {
	return s.SiteHost
}

func (s *GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceList) GetSlug() *string {
	return s.Slug
}

func (s *GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceList) GetSourceType() *string {
	return s.SourceType
}

func (s *GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceList) GetStartTime() *string {
	return s.StartTime
}

func (s *GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceList) GetStatus() *string {
	return s.Status
}

func (s *GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceList) GetStatusText() *string {
	return s.StatusText
}

func (s *GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceList) GetThumbnailUrl() *string {
	return s.ThumbnailUrl
}

func (s *GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceList) GetUserId() *string {
	return s.UserId
}

func (s *GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceList) GetVersion() *string {
	return s.Version
}

func (s *GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceList) SetAppSubType(v string) *GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceList {
	s.AppSubType = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceList) SetAppType(v string) *GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceList {
	s.AppType = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceList) SetBizId(v string) *GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceList {
	s.BizId = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceList) SetBuildType(v string) *GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceList {
	s.BuildType = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceList) SetDeleted(v int32) *GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceList {
	s.Deleted = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceList) SetDescription(v string) *GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceList {
	s.Description = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceList) SetDesignSpecBizId(v string) *GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceList {
	s.DesignSpecBizId = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceList) SetDesignSpecId(v string) *GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceList {
	s.DesignSpecId = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceList) SetDomain(v string) *GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceList {
	s.Domain = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceList) SetEndTime(v string) *GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceList {
	s.EndTime = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceList) SetEnv(v string) *GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceList {
	s.Env = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceList) SetEspBizId(v string) *GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceList {
	s.EspBizId = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceList) SetGmtCreate(v string) *GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceList {
	s.GmtCreate = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceList) SetGmtDelete(v string) *GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceList {
	s.GmtDelete = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceList) SetGmtModified(v string) *GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceList {
	s.GmtModified = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceList) SetGmtPublish(v string) *GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceList {
	s.GmtPublish = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceList) SetIconUrl(v string) *GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceList {
	s.IconUrl = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceList) SetName(v string) *GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceList {
	s.Name = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceList) SetOrderId(v string) *GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceList {
	s.OrderId = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceList) SetProfile(v *GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceListProfile) *GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceList {
	s.Profile = v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceList) SetSiteHost(v string) *GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceList {
	s.SiteHost = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceList) SetSlug(v string) *GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceList {
	s.Slug = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceList) SetSourceType(v string) *GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceList {
	s.SourceType = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceList) SetStartTime(v string) *GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceList {
	s.StartTime = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceList) SetStatus(v string) *GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceList {
	s.Status = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceList) SetStatusText(v string) *GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceList {
	s.StatusText = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceList) SetThumbnailUrl(v string) *GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceList {
	s.ThumbnailUrl = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceList) SetUserId(v string) *GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceList {
	s.UserId = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceList) SetVersion(v string) *GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceList {
	s.Version = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceList) Validate() error {
	if s.Profile != nil {
		if err := s.Profile.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceListProfile struct {
	// Management Console URL
	//
	// example:
	//
	// placeHolder
	AdminUrl *string `json:"AdminUrl,omitempty" xml:"AdminUrl,omitempty"`
	// Application Type Code
	//
	// example:
	//
	// PC_WebSite
	ApplicationType *string `json:"ApplicationType,omitempty" xml:"ApplicationType,omitempty"`
	// Application type text description
	//
	// example:
	//
	// placeHolder
	ApplicationTypeText *string `json:"ApplicationTypeText,omitempty" xml:"ApplicationTypeText,omitempty"`
	// Attached CNAME record
	//
	// example:
	//
	// placeHolder
	BindCname *string `json:"BindCname,omitempty" xml:"BindCname,omitempty"`
	// Business ID
	//
	// example:
	//
	// WS20250801153127000002
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// Product encoding
	//
	// example:
	//
	// sas
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	// [responses_200_schema_properties_Module_properties_RelatedInstanceList_items_properties_CustomerService_description]Customer service information
	//
	// example:
	//
	// placeHolder
	CustomerService *string `json:"CustomerService,omitempty" xml:"CustomerService,omitempty"`
	// Deployment area
	//
	// example:
	//
	// ChineseMainland
	DeployArea *string `json:"DeployArea,omitempty" xml:"DeployArea,omitempty"`
	// Domain Name List
	//
	// example:
	//
	// [\\"activity.syruijia.top\\", \\"admin.syruijia.top\\", \\"api.syruijia.top\\", \\"game-admin.syruijia.top\\", \\"h5game.syruijia.top\\", \\"home.syruijia.top\\", \\"invite.syruijia.top\\"]
	DomainList *string `json:"DomainList,omitempty" xml:"DomainList,omitempty"`
	// Editor URL
	//
	// example:
	//
	// placeHolder
	EditorUrl *string `json:"EditorUrl,omitempty" xml:"EditorUrl,omitempty"`
	// Creation Time
	//
	// example:
	//
	// 1621734214000
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// Updated At
	//
	// example:
	//
	// Thu Oct 24 09:12:31 CST 2024
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// ICP number
	//
	// example:
	//
	// placeHolder
	IcpbaNo *string `json:"IcpbaNo,omitempty" xml:"IcpbaNo,omitempty"`
	// Image list
	//
	// example:
	//
	// {\\"Image\\": []}
	ImageList *string `json:"ImageList,omitempty" xml:"ImageList,omitempty"`
	// Lenovo instance ID
	//
	// example:
	//
	// placeHolder
	LxInstanceId *string `json:"LxInstanceId,omitempty" xml:"LxInstanceId,omitempty"`
	// Order Time
	//
	// example:
	//
	// placeHolder
	OrdTime *string `json:"OrdTime,omitempty" xml:"OrdTime,omitempty"`
	// Order ID
	//
	// example:
	//
	// 248808934190692
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// Order Quantity
	//
	// example:
	//
	// 1
	OrderNum *int32 `json:"OrderNum,omitempty" xml:"OrderNum,omitempty"`
	// Channel Partner Business ID
	//
	// example:
	//
	// 100086
	PartnerId *string `json:"PartnerId,omitempty" xml:"PartnerId,omitempty"`
	// Payment time
	//
	// example:
	//
	// placeHolder
	PayTime *string `json:"PayTime,omitempty" xml:"PayTime,omitempty"`
	// Preview URL
	//
	// example:
	//
	// https://preview-lyj.aliyuncs.com/preview/temp/9fb36fc05d0a48cdb92d7397336c214f?subSceneIds=734285&type=interim
	PreviewUrl *string `json:"PreviewUrl,omitempty" xml:"PreviewUrl,omitempty"`
	// SEO site information
	//
	// example:
	//
	// placeHolder
	SeoSite *string `json:"SeoSite,omitempty" xml:"SeoSite,omitempty"`
	// Site logo
	//
	// example:
	//
	// placeHolder
	SiteLogo *string `json:"SiteLogo,omitempty" xml:"SiteLogo,omitempty"`
	// Site version encoding
	//
	// example:
	//
	// Trial_Edition
	SiteVersion *string `json:"SiteVersion,omitempty" xml:"SiteVersion,omitempty"`
	// Site Version Text Description
	//
	// example:
	//
	// placeHolder
	SiteVersionText *string `json:"SiteVersionText,omitempty" xml:"SiteVersionText,omitempty"`
	// Source Information
	//
	// example:
	//
	// 10.71.130.205
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// Template ETag identity
	//
	// example:
	//
	// placeHolder
	TemplateEtag *string `json:"TemplateEtag,omitempty" xml:"TemplateEtag,omitempty"`
	// Template ID
	//
	// example:
	//
	// 4644
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// Text List
	//
	// example:
	//
	// placeHolder
	TextList *string `json:"TextList,omitempty" xml:"TextList,omitempty"`
	// Thumbnail
	//
	// example:
	//
	// https://aloss-recruit-aiinterview.oss-cn-beijing.aliyuncs.com/thumbnail-1766456326256-0.png
	Thumbnail     *string `json:"Thumbnail,omitempty" xml:"Thumbnail,omitempty"`
	UpgradeStatus *string `json:"UpgradeStatus,omitempty" xml:"UpgradeStatus,omitempty"`
}

func (s GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceListProfile) String() string {
	return dara.Prettify(s)
}

func (s GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceListProfile) GoString() string {
	return s.String()
}

func (s *GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceListProfile) GetAdminUrl() *string {
	return s.AdminUrl
}

func (s *GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceListProfile) GetApplicationType() *string {
	return s.ApplicationType
}

func (s *GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceListProfile) GetApplicationTypeText() *string {
	return s.ApplicationTypeText
}

func (s *GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceListProfile) GetBindCname() *string {
	return s.BindCname
}

func (s *GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceListProfile) GetBizId() *string {
	return s.BizId
}

func (s *GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceListProfile) GetCommodityCode() *string {
	return s.CommodityCode
}

func (s *GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceListProfile) GetCustomerService() *string {
	return s.CustomerService
}

func (s *GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceListProfile) GetDeployArea() *string {
	return s.DeployArea
}

func (s *GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceListProfile) GetDomainList() *string {
	return s.DomainList
}

func (s *GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceListProfile) GetEditorUrl() *string {
	return s.EditorUrl
}

func (s *GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceListProfile) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceListProfile) GetGmtModified() *string {
	return s.GmtModified
}

func (s *GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceListProfile) GetIcpbaNo() *string {
	return s.IcpbaNo
}

func (s *GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceListProfile) GetImageList() *string {
	return s.ImageList
}

func (s *GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceListProfile) GetLxInstanceId() *string {
	return s.LxInstanceId
}

func (s *GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceListProfile) GetOrdTime() *string {
	return s.OrdTime
}

func (s *GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceListProfile) GetOrderId() *string {
	return s.OrderId
}

func (s *GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceListProfile) GetOrderNum() *int32 {
	return s.OrderNum
}

func (s *GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceListProfile) GetPartnerId() *string {
	return s.PartnerId
}

func (s *GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceListProfile) GetPayTime() *string {
	return s.PayTime
}

func (s *GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceListProfile) GetPreviewUrl() *string {
	return s.PreviewUrl
}

func (s *GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceListProfile) GetSeoSite() *string {
	return s.SeoSite
}

func (s *GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceListProfile) GetSiteLogo() *string {
	return s.SiteLogo
}

func (s *GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceListProfile) GetSiteVersion() *string {
	return s.SiteVersion
}

func (s *GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceListProfile) GetSiteVersionText() *string {
	return s.SiteVersionText
}

func (s *GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceListProfile) GetSource() *string {
	return s.Source
}

func (s *GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceListProfile) GetTemplateEtag() *string {
	return s.TemplateEtag
}

func (s *GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceListProfile) GetTemplateId() *string {
	return s.TemplateId
}

func (s *GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceListProfile) GetTextList() *string {
	return s.TextList
}

func (s *GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceListProfile) GetThumbnail() *string {
	return s.Thumbnail
}

func (s *GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceListProfile) GetUpgradeStatus() *string {
	return s.UpgradeStatus
}

func (s *GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceListProfile) SetAdminUrl(v string) *GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceListProfile {
	s.AdminUrl = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceListProfile) SetApplicationType(v string) *GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceListProfile {
	s.ApplicationType = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceListProfile) SetApplicationTypeText(v string) *GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceListProfile {
	s.ApplicationTypeText = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceListProfile) SetBindCname(v string) *GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceListProfile {
	s.BindCname = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceListProfile) SetBizId(v string) *GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceListProfile {
	s.BizId = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceListProfile) SetCommodityCode(v string) *GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceListProfile {
	s.CommodityCode = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceListProfile) SetCustomerService(v string) *GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceListProfile {
	s.CustomerService = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceListProfile) SetDeployArea(v string) *GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceListProfile {
	s.DeployArea = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceListProfile) SetDomainList(v string) *GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceListProfile {
	s.DomainList = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceListProfile) SetEditorUrl(v string) *GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceListProfile {
	s.EditorUrl = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceListProfile) SetGmtCreate(v string) *GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceListProfile {
	s.GmtCreate = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceListProfile) SetGmtModified(v string) *GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceListProfile {
	s.GmtModified = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceListProfile) SetIcpbaNo(v string) *GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceListProfile {
	s.IcpbaNo = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceListProfile) SetImageList(v string) *GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceListProfile {
	s.ImageList = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceListProfile) SetLxInstanceId(v string) *GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceListProfile {
	s.LxInstanceId = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceListProfile) SetOrdTime(v string) *GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceListProfile {
	s.OrdTime = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceListProfile) SetOrderId(v string) *GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceListProfile {
	s.OrderId = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceListProfile) SetOrderNum(v int32) *GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceListProfile {
	s.OrderNum = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceListProfile) SetPartnerId(v string) *GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceListProfile {
	s.PartnerId = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceListProfile) SetPayTime(v string) *GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceListProfile {
	s.PayTime = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceListProfile) SetPreviewUrl(v string) *GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceListProfile {
	s.PreviewUrl = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceListProfile) SetSeoSite(v string) *GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceListProfile {
	s.SeoSite = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceListProfile) SetSiteLogo(v string) *GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceListProfile {
	s.SiteLogo = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceListProfile) SetSiteVersion(v string) *GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceListProfile {
	s.SiteVersion = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceListProfile) SetSiteVersionText(v string) *GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceListProfile {
	s.SiteVersionText = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceListProfile) SetSource(v string) *GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceListProfile {
	s.Source = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceListProfile) SetTemplateEtag(v string) *GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceListProfile {
	s.TemplateEtag = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceListProfile) SetTemplateId(v string) *GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceListProfile {
	s.TemplateId = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceListProfile) SetTextList(v string) *GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceListProfile {
	s.TextList = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceListProfile) SetThumbnail(v string) *GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceListProfile {
	s.Thumbnail = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceListProfile) SetUpgradeStatus(v string) *GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceListProfile {
	s.UpgradeStatus = &v
	return s
}

func (s *GetAppInstanceForPartnerResponseBodyModuleRelatedInstanceListProfile) Validate() error {
	return dara.Validate(s)
}
