// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTranslateChannelDraftResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *TranslateChannelDraftResponseBody
	GetAccessDeniedDetail() *string
	SetAllowRetry(v bool) *TranslateChannelDraftResponseBody
	GetAllowRetry() *bool
	SetAppName(v string) *TranslateChannelDraftResponseBody
	GetAppName() *string
	SetDynamicCode(v string) *TranslateChannelDraftResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *TranslateChannelDraftResponseBody
	GetDynamicMessage() *string
	SetErrorArgs(v []interface{}) *TranslateChannelDraftResponseBody
	GetErrorArgs() []interface{}
	SetModule(v *TranslateChannelDraftResponseBodyModule) *TranslateChannelDraftResponseBody
	GetModule() *TranslateChannelDraftResponseBodyModule
	SetRequestId(v string) *TranslateChannelDraftResponseBody
	GetRequestId() *string
	SetRootErrorCode(v string) *TranslateChannelDraftResponseBody
	GetRootErrorCode() *string
	SetRootErrorMsg(v string) *TranslateChannelDraftResponseBody
	GetRootErrorMsg() *string
	SetSynchro(v bool) *TranslateChannelDraftResponseBody
	GetSynchro() *bool
}

type TranslateChannelDraftResponseBody struct {
	// The detailed reason why access was denied.
	//
	// example:
	//
	// {}
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// Indicates whether retry is allowed.
	//
	// example:
	//
	// False
	AllowRetry *bool `json:"AllowRetry,omitempty" xml:"AllowRetry,omitempty"`
	// The application name used to query the application.
	//
	// example:
	//
	// ish-intelligence-store-platform-admin-web
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The dynamic code. This parameter is not in use. Ignore this parameter.
	//
	// example:
	//
	// ERROR-oo1
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// The dynamic error message that is used to replace the `%s` variable in the **ErrMessage*	- return parameter.
	//
	// > If **ErrMessage*	- returns **The Value of Input Parameter %s is not valid*	- and **DynamicMessage*	- returns **DtsJobId**, the value of the request parameter **DtsJobId*	- is invalid.
	//
	// example:
	//
	// abc
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// The error parameters.
	ErrorArgs []interface{} `json:"ErrorArgs,omitempty" xml:"ErrorArgs,omitempty" type:"Repeated"`
	// Indicates whether the deletion was successful.
	Module *TranslateChannelDraftResponseBodyModule `json:"Module,omitempty" xml:"Module,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 6C6B99AC-39EC-5350-874C-204128C905E6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The error code.
	//
	// example:
	//
	// SYSTEM.ERROR
	RootErrorCode *string `json:"RootErrorCode,omitempty" xml:"RootErrorCode,omitempty"`
	// The exception message.
	//
	// example:
	//
	// System exception
	RootErrorMsg *string `json:"RootErrorMsg,omitempty" xml:"RootErrorMsg,omitempty"`
	// Indicates whether the request is processed synchronously.
	//
	// example:
	//
	// True
	Synchro *bool `json:"Synchro,omitempty" xml:"Synchro,omitempty"`
}

func (s TranslateChannelDraftResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TranslateChannelDraftResponseBody) GoString() string {
	return s.String()
}

func (s *TranslateChannelDraftResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *TranslateChannelDraftResponseBody) GetAllowRetry() *bool {
	return s.AllowRetry
}

func (s *TranslateChannelDraftResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *TranslateChannelDraftResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *TranslateChannelDraftResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *TranslateChannelDraftResponseBody) GetErrorArgs() []interface{} {
	return s.ErrorArgs
}

func (s *TranslateChannelDraftResponseBody) GetModule() *TranslateChannelDraftResponseBodyModule {
	return s.Module
}

func (s *TranslateChannelDraftResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TranslateChannelDraftResponseBody) GetRootErrorCode() *string {
	return s.RootErrorCode
}

func (s *TranslateChannelDraftResponseBody) GetRootErrorMsg() *string {
	return s.RootErrorMsg
}

func (s *TranslateChannelDraftResponseBody) GetSynchro() *bool {
	return s.Synchro
}

func (s *TranslateChannelDraftResponseBody) SetAccessDeniedDetail(v string) *TranslateChannelDraftResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *TranslateChannelDraftResponseBody) SetAllowRetry(v bool) *TranslateChannelDraftResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *TranslateChannelDraftResponseBody) SetAppName(v string) *TranslateChannelDraftResponseBody {
	s.AppName = &v
	return s
}

func (s *TranslateChannelDraftResponseBody) SetDynamicCode(v string) *TranslateChannelDraftResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *TranslateChannelDraftResponseBody) SetDynamicMessage(v string) *TranslateChannelDraftResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *TranslateChannelDraftResponseBody) SetErrorArgs(v []interface{}) *TranslateChannelDraftResponseBody {
	s.ErrorArgs = v
	return s
}

func (s *TranslateChannelDraftResponseBody) SetModule(v *TranslateChannelDraftResponseBodyModule) *TranslateChannelDraftResponseBody {
	s.Module = v
	return s
}

func (s *TranslateChannelDraftResponseBody) SetRequestId(v string) *TranslateChannelDraftResponseBody {
	s.RequestId = &v
	return s
}

func (s *TranslateChannelDraftResponseBody) SetRootErrorCode(v string) *TranslateChannelDraftResponseBody {
	s.RootErrorCode = &v
	return s
}

func (s *TranslateChannelDraftResponseBody) SetRootErrorMsg(v string) *TranslateChannelDraftResponseBody {
	s.RootErrorMsg = &v
	return s
}

func (s *TranslateChannelDraftResponseBody) SetSynchro(v bool) *TranslateChannelDraftResponseBody {
	s.Synchro = &v
	return s
}

func (s *TranslateChannelDraftResponseBody) Validate() error {
	if s.Module != nil {
		if err := s.Module.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type TranslateChannelDraftResponseBodyModule struct {
	// The AI adaptation status. Valid values: NONE, ADAPTING, DONE, and FAILED.
	//
	// example:
	//
	// xx
	AdaptStatus *string `json:"AdaptStatus,omitempty" xml:"AdaptStatus,omitempty"`
	// The channel-adapted content body.
	//
	// example:
	//
	// xx
	AdaptedContent *string `json:"AdaptedContent,omitempty" xml:"AdaptedContent,omitempty"`
	// The channel-adapted title.
	//
	// example:
	//
	// xx
	AdaptedTitle *string `json:"AdaptedTitle,omitempty" xml:"AdaptedTitle,omitempty"`
	// The channel enumeration value.
	//
	// example:
	//
	// WECHAT
	Channel *string `json:"Channel,omitempty" xml:"Channel,omitempty"`
	// The publishing account snapshot.
	//
	// example:
	//
	// xx
	ChannelAccount *string `json:"ChannelAccount,omitempty" xml:"ChannelAccount,omitempty"`
	// The display name of the publishing account.
	//
	// example:
	//
	// xx
	ChannelAccountName *string `json:"ChannelAccountName,omitempty" xml:"ChannelAccountName,omitempty"`
	// The display name of the channel.
	//
	// example:
	//
	// channel_905
	ChannelName *string `json:"ChannelName,omitempty" xml:"ChannelName,omitempty"`
	// The channel type. Valid values: DOMESTIC, OVERSEA, and INTERNAL.
	//
	// example:
	//
	// wechat
	ChannelType *string `json:"ChannelType,omitempty" xml:"ChannelType,omitempty"`
	// The list of channel cover images.
	CoverImages []*TranslateChannelDraftResponseBodyModuleCoverImages `json:"CoverImages,omitempty" xml:"CoverImages,omitempty" type:"Repeated"`
	// The ID of the channel draft.
	//
	// example:
	//
	// 85e219ef-c458-4a0d-a249-2776864143bb
	DraftId *string `json:"DraftId,omitempty" xml:"DraftId,omitempty"`
	// The ID returned by the platform.
	//
	// example:
	//
	// a367d19c02ff004ac20a7b2078066fe1f370
	ExternalId *string `json:"ExternalId,omitempty" xml:"ExternalId,omitempty"`
	// The redirect URL of the platform.
	//
	// example:
	//
	// https://net-video.mistong.com:10160/api/preview/video/2026/02/02/7181627703f6437480a7bffe8ed1e7b0_20260202100338.mp4
	ExternalUrl *string `json:"ExternalUrl,omitempty" xml:"ExternalUrl,omitempty"`
	// The reason for the failure.
	//
	// example:
	//
	// Invalid parameter
	FailReason *string `json:"FailReason,omitempty" xml:"FailReason,omitempty"`
	// The channel-specific publishing configuration in JSON format.
	//
	// example:
	//
	// xx
	PublishConfig *string `json:"PublishConfig,omitempty" xml:"PublishConfig,omitempty"`
	// The publishing time, in millisecond timestamp format.
	//
	// example:
	//
	// xx
	PublishedAt *int64 `json:"PublishedAt,omitempty" xml:"PublishedAt,omitempty"`
	// The status. Valid values: EDITING, PUBLISHING, SUCCESS, and FAILED.
	//
	// example:
	//
	// NORMAL
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s TranslateChannelDraftResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s TranslateChannelDraftResponseBodyModule) GoString() string {
	return s.String()
}

func (s *TranslateChannelDraftResponseBodyModule) GetAdaptStatus() *string {
	return s.AdaptStatus
}

func (s *TranslateChannelDraftResponseBodyModule) GetAdaptedContent() *string {
	return s.AdaptedContent
}

func (s *TranslateChannelDraftResponseBodyModule) GetAdaptedTitle() *string {
	return s.AdaptedTitle
}

func (s *TranslateChannelDraftResponseBodyModule) GetChannel() *string {
	return s.Channel
}

func (s *TranslateChannelDraftResponseBodyModule) GetChannelAccount() *string {
	return s.ChannelAccount
}

func (s *TranslateChannelDraftResponseBodyModule) GetChannelAccountName() *string {
	return s.ChannelAccountName
}

func (s *TranslateChannelDraftResponseBodyModule) GetChannelName() *string {
	return s.ChannelName
}

func (s *TranslateChannelDraftResponseBodyModule) GetChannelType() *string {
	return s.ChannelType
}

func (s *TranslateChannelDraftResponseBodyModule) GetCoverImages() []*TranslateChannelDraftResponseBodyModuleCoverImages {
	return s.CoverImages
}

func (s *TranslateChannelDraftResponseBodyModule) GetDraftId() *string {
	return s.DraftId
}

func (s *TranslateChannelDraftResponseBodyModule) GetExternalId() *string {
	return s.ExternalId
}

func (s *TranslateChannelDraftResponseBodyModule) GetExternalUrl() *string {
	return s.ExternalUrl
}

func (s *TranslateChannelDraftResponseBodyModule) GetFailReason() *string {
	return s.FailReason
}

func (s *TranslateChannelDraftResponseBodyModule) GetPublishConfig() *string {
	return s.PublishConfig
}

func (s *TranslateChannelDraftResponseBodyModule) GetPublishedAt() *int64 {
	return s.PublishedAt
}

func (s *TranslateChannelDraftResponseBodyModule) GetStatus() *string {
	return s.Status
}

func (s *TranslateChannelDraftResponseBodyModule) SetAdaptStatus(v string) *TranslateChannelDraftResponseBodyModule {
	s.AdaptStatus = &v
	return s
}

func (s *TranslateChannelDraftResponseBodyModule) SetAdaptedContent(v string) *TranslateChannelDraftResponseBodyModule {
	s.AdaptedContent = &v
	return s
}

func (s *TranslateChannelDraftResponseBodyModule) SetAdaptedTitle(v string) *TranslateChannelDraftResponseBodyModule {
	s.AdaptedTitle = &v
	return s
}

func (s *TranslateChannelDraftResponseBodyModule) SetChannel(v string) *TranslateChannelDraftResponseBodyModule {
	s.Channel = &v
	return s
}

func (s *TranslateChannelDraftResponseBodyModule) SetChannelAccount(v string) *TranslateChannelDraftResponseBodyModule {
	s.ChannelAccount = &v
	return s
}

func (s *TranslateChannelDraftResponseBodyModule) SetChannelAccountName(v string) *TranslateChannelDraftResponseBodyModule {
	s.ChannelAccountName = &v
	return s
}

func (s *TranslateChannelDraftResponseBodyModule) SetChannelName(v string) *TranslateChannelDraftResponseBodyModule {
	s.ChannelName = &v
	return s
}

func (s *TranslateChannelDraftResponseBodyModule) SetChannelType(v string) *TranslateChannelDraftResponseBodyModule {
	s.ChannelType = &v
	return s
}

func (s *TranslateChannelDraftResponseBodyModule) SetCoverImages(v []*TranslateChannelDraftResponseBodyModuleCoverImages) *TranslateChannelDraftResponseBodyModule {
	s.CoverImages = v
	return s
}

func (s *TranslateChannelDraftResponseBodyModule) SetDraftId(v string) *TranslateChannelDraftResponseBodyModule {
	s.DraftId = &v
	return s
}

func (s *TranslateChannelDraftResponseBodyModule) SetExternalId(v string) *TranslateChannelDraftResponseBodyModule {
	s.ExternalId = &v
	return s
}

func (s *TranslateChannelDraftResponseBodyModule) SetExternalUrl(v string) *TranslateChannelDraftResponseBodyModule {
	s.ExternalUrl = &v
	return s
}

func (s *TranslateChannelDraftResponseBodyModule) SetFailReason(v string) *TranslateChannelDraftResponseBodyModule {
	s.FailReason = &v
	return s
}

func (s *TranslateChannelDraftResponseBodyModule) SetPublishConfig(v string) *TranslateChannelDraftResponseBodyModule {
	s.PublishConfig = &v
	return s
}

func (s *TranslateChannelDraftResponseBodyModule) SetPublishedAt(v int64) *TranslateChannelDraftResponseBodyModule {
	s.PublishedAt = &v
	return s
}

func (s *TranslateChannelDraftResponseBodyModule) SetStatus(v string) *TranslateChannelDraftResponseBodyModule {
	s.Status = &v
	return s
}

func (s *TranslateChannelDraftResponseBodyModule) Validate() error {
	if s.CoverImages != nil {
		for _, item := range s.CoverImages {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type TranslateChannelDraftResponseBodyModuleCoverImages struct {
	// The file ID in the material center.
	//
	// example:
	//
	// xx
	MaterialFileId *string `json:"MaterialFileId,omitempty" xml:"MaterialFileId,omitempty"`
	// The CDN URL of the image.
	//
	// example:
	//
	// oss://anchashiscan.oss-cn-hangzhou.aliyuncs.com/prodpopscanBnlK054f
	OssUrl *string `json:"OssUrl,omitempty" xml:"OssUrl,omitempty"`
	// The sort order number.
	//
	// example:
	//
	// Desc
	SortOrder *int32 `json:"SortOrder,omitempty" xml:"SortOrder,omitempty"`
}

func (s TranslateChannelDraftResponseBodyModuleCoverImages) String() string {
	return dara.Prettify(s)
}

func (s TranslateChannelDraftResponseBodyModuleCoverImages) GoString() string {
	return s.String()
}

func (s *TranslateChannelDraftResponseBodyModuleCoverImages) GetMaterialFileId() *string {
	return s.MaterialFileId
}

func (s *TranslateChannelDraftResponseBodyModuleCoverImages) GetOssUrl() *string {
	return s.OssUrl
}

func (s *TranslateChannelDraftResponseBodyModuleCoverImages) GetSortOrder() *int32 {
	return s.SortOrder
}

func (s *TranslateChannelDraftResponseBodyModuleCoverImages) SetMaterialFileId(v string) *TranslateChannelDraftResponseBodyModuleCoverImages {
	s.MaterialFileId = &v
	return s
}

func (s *TranslateChannelDraftResponseBodyModuleCoverImages) SetOssUrl(v string) *TranslateChannelDraftResponseBodyModuleCoverImages {
	s.OssUrl = &v
	return s
}

func (s *TranslateChannelDraftResponseBodyModuleCoverImages) SetSortOrder(v int32) *TranslateChannelDraftResponseBodyModuleCoverImages {
	s.SortOrder = &v
	return s
}

func (s *TranslateChannelDraftResponseBodyModuleCoverImages) Validate() error {
	return dara.Validate(s)
}
