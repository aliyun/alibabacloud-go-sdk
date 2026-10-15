// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRewriteChannelTitleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *RewriteChannelTitleResponseBody
	GetAccessDeniedDetail() *string
	SetAllowRetry(v bool) *RewriteChannelTitleResponseBody
	GetAllowRetry() *bool
	SetAppName(v string) *RewriteChannelTitleResponseBody
	GetAppName() *string
	SetDynamicCode(v string) *RewriteChannelTitleResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *RewriteChannelTitleResponseBody
	GetDynamicMessage() *string
	SetErrorArgs(v []interface{}) *RewriteChannelTitleResponseBody
	GetErrorArgs() []interface{}
	SetModule(v *RewriteChannelTitleResponseBodyModule) *RewriteChannelTitleResponseBody
	GetModule() *RewriteChannelTitleResponseBodyModule
	SetRequestId(v string) *RewriteChannelTitleResponseBody
	GetRequestId() *string
	SetRootErrorCode(v string) *RewriteChannelTitleResponseBody
	GetRootErrorCode() *string
	SetRootErrorMsg(v string) *RewriteChannelTitleResponseBody
	GetRootErrorMsg() *string
	SetSynchro(v bool) *RewriteChannelTitleResponseBody
	GetSynchro() *bool
}

type RewriteChannelTitleResponseBody struct {
	// The details of the permission verification failure.
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
	// The application name.
	//
	// example:
	//
	// ish-intelligence-store-platform-admin-web
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The dynamic error code.
	//
	// example:
	//
	// ERROR-oo1
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// The dynamic error message, which is used to replace the `%s` placeholder in the **ErrMessage*	- response parameter.
	//
	// > If **ErrMessage*	- returns **The Value of Input Parameter %s is not valid*	- and **DynamicMessage*	- returns **DtsJobId**, the value of the request parameter **DtsJobId*	- is invalid.
	//
	// example:
	//
	// abc
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// The error parameters.
	ErrorArgs []interface{} `json:"ErrorArgs,omitempty" xml:"ErrorArgs,omitempty" type:"Repeated"`
	// Indicates whether the operation is successful.
	Module *RewriteChannelTitleResponseBodyModule `json:"Module,omitempty" xml:"Module,omitempty" type:"Struct"`
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
	// The error message.
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

func (s RewriteChannelTitleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RewriteChannelTitleResponseBody) GoString() string {
	return s.String()
}

func (s *RewriteChannelTitleResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *RewriteChannelTitleResponseBody) GetAllowRetry() *bool {
	return s.AllowRetry
}

func (s *RewriteChannelTitleResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *RewriteChannelTitleResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *RewriteChannelTitleResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *RewriteChannelTitleResponseBody) GetErrorArgs() []interface{} {
	return s.ErrorArgs
}

func (s *RewriteChannelTitleResponseBody) GetModule() *RewriteChannelTitleResponseBodyModule {
	return s.Module
}

func (s *RewriteChannelTitleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RewriteChannelTitleResponseBody) GetRootErrorCode() *string {
	return s.RootErrorCode
}

func (s *RewriteChannelTitleResponseBody) GetRootErrorMsg() *string {
	return s.RootErrorMsg
}

func (s *RewriteChannelTitleResponseBody) GetSynchro() *bool {
	return s.Synchro
}

func (s *RewriteChannelTitleResponseBody) SetAccessDeniedDetail(v string) *RewriteChannelTitleResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *RewriteChannelTitleResponseBody) SetAllowRetry(v bool) *RewriteChannelTitleResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *RewriteChannelTitleResponseBody) SetAppName(v string) *RewriteChannelTitleResponseBody {
	s.AppName = &v
	return s
}

func (s *RewriteChannelTitleResponseBody) SetDynamicCode(v string) *RewriteChannelTitleResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *RewriteChannelTitleResponseBody) SetDynamicMessage(v string) *RewriteChannelTitleResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *RewriteChannelTitleResponseBody) SetErrorArgs(v []interface{}) *RewriteChannelTitleResponseBody {
	s.ErrorArgs = v
	return s
}

func (s *RewriteChannelTitleResponseBody) SetModule(v *RewriteChannelTitleResponseBodyModule) *RewriteChannelTitleResponseBody {
	s.Module = v
	return s
}

func (s *RewriteChannelTitleResponseBody) SetRequestId(v string) *RewriteChannelTitleResponseBody {
	s.RequestId = &v
	return s
}

func (s *RewriteChannelTitleResponseBody) SetRootErrorCode(v string) *RewriteChannelTitleResponseBody {
	s.RootErrorCode = &v
	return s
}

func (s *RewriteChannelTitleResponseBody) SetRootErrorMsg(v string) *RewriteChannelTitleResponseBody {
	s.RootErrorMsg = &v
	return s
}

func (s *RewriteChannelTitleResponseBody) SetSynchro(v bool) *RewriteChannelTitleResponseBody {
	s.Synchro = &v
	return s
}

func (s *RewriteChannelTitleResponseBody) Validate() error {
	if s.Module != nil {
		if err := s.Module.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RewriteChannelTitleResponseBodyModule struct {
	// The AI adaptation status. Valid values: NONE, ADAPTING, DONE, and FAILED.
	//
	// example:
	//
	// xxx
	AdaptStatus *string `json:"AdaptStatus,omitempty" xml:"AdaptStatus,omitempty"`
	// The channel-adapted content.
	//
	// example:
	//
	// xxx
	AdaptedContent *string `json:"AdaptedContent,omitempty" xml:"AdaptedContent,omitempty"`
	// The channel-adapted title.
	//
	// example:
	//
	// xxx
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
	// xxx
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
	// sipmplm
	ChannelName *string `json:"ChannelName,omitempty" xml:"ChannelName,omitempty"`
	// The channel type. Valid values: DOMESTIC, OVERSEA, and INTERNAL.
	//
	// example:
	//
	// wechat
	ChannelType *string `json:"ChannelType,omitempty" xml:"ChannelType,omitempty"`
	// The list of channel cover images.
	CoverImages []*RewriteChannelTitleResponseBodyModuleCoverImages `json:"CoverImages,omitempty" xml:"CoverImages,omitempty" type:"Repeated"`
	// The channel draft ID.
	//
	// example:
	//
	// draft-789
	DraftId *string `json:"DraftId,omitempty" xml:"DraftId,omitempty"`
	// The ID returned by the platform.
	//
	// example:
	//
	// a2b36d4b1dab614623187f01f01d138e809d
	ExternalId *string `json:"ExternalId,omitempty" xml:"ExternalId,omitempty"`
	// The redirect URL of the platform.
	//
	// example:
	//
	// https://net-video.mistong.com:10160/api/preview/video/2025/11/03/53a90dd9bdfa44bda53aa2534bb8f463_20251103100032.mp4
	ExternalUrl *string `json:"ExternalUrl,omitempty" xml:"ExternalUrl,omitempty"`
	// The failure reason.
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
	// The publishing time, in millisecond timestamp.
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

func (s RewriteChannelTitleResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s RewriteChannelTitleResponseBodyModule) GoString() string {
	return s.String()
}

func (s *RewriteChannelTitleResponseBodyModule) GetAdaptStatus() *string {
	return s.AdaptStatus
}

func (s *RewriteChannelTitleResponseBodyModule) GetAdaptedContent() *string {
	return s.AdaptedContent
}

func (s *RewriteChannelTitleResponseBodyModule) GetAdaptedTitle() *string {
	return s.AdaptedTitle
}

func (s *RewriteChannelTitleResponseBodyModule) GetChannel() *string {
	return s.Channel
}

func (s *RewriteChannelTitleResponseBodyModule) GetChannelAccount() *string {
	return s.ChannelAccount
}

func (s *RewriteChannelTitleResponseBodyModule) GetChannelAccountName() *string {
	return s.ChannelAccountName
}

func (s *RewriteChannelTitleResponseBodyModule) GetChannelName() *string {
	return s.ChannelName
}

func (s *RewriteChannelTitleResponseBodyModule) GetChannelType() *string {
	return s.ChannelType
}

func (s *RewriteChannelTitleResponseBodyModule) GetCoverImages() []*RewriteChannelTitleResponseBodyModuleCoverImages {
	return s.CoverImages
}

func (s *RewriteChannelTitleResponseBodyModule) GetDraftId() *string {
	return s.DraftId
}

func (s *RewriteChannelTitleResponseBodyModule) GetExternalId() *string {
	return s.ExternalId
}

func (s *RewriteChannelTitleResponseBodyModule) GetExternalUrl() *string {
	return s.ExternalUrl
}

func (s *RewriteChannelTitleResponseBodyModule) GetFailReason() *string {
	return s.FailReason
}

func (s *RewriteChannelTitleResponseBodyModule) GetPublishConfig() *string {
	return s.PublishConfig
}

func (s *RewriteChannelTitleResponseBodyModule) GetPublishedAt() *int64 {
	return s.PublishedAt
}

func (s *RewriteChannelTitleResponseBodyModule) GetStatus() *string {
	return s.Status
}

func (s *RewriteChannelTitleResponseBodyModule) SetAdaptStatus(v string) *RewriteChannelTitleResponseBodyModule {
	s.AdaptStatus = &v
	return s
}

func (s *RewriteChannelTitleResponseBodyModule) SetAdaptedContent(v string) *RewriteChannelTitleResponseBodyModule {
	s.AdaptedContent = &v
	return s
}

func (s *RewriteChannelTitleResponseBodyModule) SetAdaptedTitle(v string) *RewriteChannelTitleResponseBodyModule {
	s.AdaptedTitle = &v
	return s
}

func (s *RewriteChannelTitleResponseBodyModule) SetChannel(v string) *RewriteChannelTitleResponseBodyModule {
	s.Channel = &v
	return s
}

func (s *RewriteChannelTitleResponseBodyModule) SetChannelAccount(v string) *RewriteChannelTitleResponseBodyModule {
	s.ChannelAccount = &v
	return s
}

func (s *RewriteChannelTitleResponseBodyModule) SetChannelAccountName(v string) *RewriteChannelTitleResponseBodyModule {
	s.ChannelAccountName = &v
	return s
}

func (s *RewriteChannelTitleResponseBodyModule) SetChannelName(v string) *RewriteChannelTitleResponseBodyModule {
	s.ChannelName = &v
	return s
}

func (s *RewriteChannelTitleResponseBodyModule) SetChannelType(v string) *RewriteChannelTitleResponseBodyModule {
	s.ChannelType = &v
	return s
}

func (s *RewriteChannelTitleResponseBodyModule) SetCoverImages(v []*RewriteChannelTitleResponseBodyModuleCoverImages) *RewriteChannelTitleResponseBodyModule {
	s.CoverImages = v
	return s
}

func (s *RewriteChannelTitleResponseBodyModule) SetDraftId(v string) *RewriteChannelTitleResponseBodyModule {
	s.DraftId = &v
	return s
}

func (s *RewriteChannelTitleResponseBodyModule) SetExternalId(v string) *RewriteChannelTitleResponseBodyModule {
	s.ExternalId = &v
	return s
}

func (s *RewriteChannelTitleResponseBodyModule) SetExternalUrl(v string) *RewriteChannelTitleResponseBodyModule {
	s.ExternalUrl = &v
	return s
}

func (s *RewriteChannelTitleResponseBodyModule) SetFailReason(v string) *RewriteChannelTitleResponseBodyModule {
	s.FailReason = &v
	return s
}

func (s *RewriteChannelTitleResponseBodyModule) SetPublishConfig(v string) *RewriteChannelTitleResponseBodyModule {
	s.PublishConfig = &v
	return s
}

func (s *RewriteChannelTitleResponseBodyModule) SetPublishedAt(v int64) *RewriteChannelTitleResponseBodyModule {
	s.PublishedAt = &v
	return s
}

func (s *RewriteChannelTitleResponseBodyModule) SetStatus(v string) *RewriteChannelTitleResponseBodyModule {
	s.Status = &v
	return s
}

func (s *RewriteChannelTitleResponseBodyModule) Validate() error {
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

type RewriteChannelTitleResponseBodyModuleCoverImages struct {
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
	// https://embedding-pic.oss-cn-beijing-internal.aliyuncs.com/30516570
	OssUrl *string `json:"OssUrl,omitempty" xml:"OssUrl,omitempty"`
	// The sort order number.
	//
	// example:
	//
	// Ascending
	SortOrder *int32 `json:"SortOrder,omitempty" xml:"SortOrder,omitempty"`
}

func (s RewriteChannelTitleResponseBodyModuleCoverImages) String() string {
	return dara.Prettify(s)
}

func (s RewriteChannelTitleResponseBodyModuleCoverImages) GoString() string {
	return s.String()
}

func (s *RewriteChannelTitleResponseBodyModuleCoverImages) GetMaterialFileId() *string {
	return s.MaterialFileId
}

func (s *RewriteChannelTitleResponseBodyModuleCoverImages) GetOssUrl() *string {
	return s.OssUrl
}

func (s *RewriteChannelTitleResponseBodyModuleCoverImages) GetSortOrder() *int32 {
	return s.SortOrder
}

func (s *RewriteChannelTitleResponseBodyModuleCoverImages) SetMaterialFileId(v string) *RewriteChannelTitleResponseBodyModuleCoverImages {
	s.MaterialFileId = &v
	return s
}

func (s *RewriteChannelTitleResponseBodyModuleCoverImages) SetOssUrl(v string) *RewriteChannelTitleResponseBodyModuleCoverImages {
	s.OssUrl = &v
	return s
}

func (s *RewriteChannelTitleResponseBodyModuleCoverImages) SetSortOrder(v int32) *RewriteChannelTitleResponseBodyModuleCoverImages {
	s.SortOrder = &v
	return s
}

func (s *RewriteChannelTitleResponseBodyModuleCoverImages) Validate() error {
	return dara.Validate(s)
}
