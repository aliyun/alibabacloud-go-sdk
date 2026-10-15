// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDistributionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *CreateDistributionResponseBody
	GetAccessDeniedDetail() *string
	SetAllowRetry(v bool) *CreateDistributionResponseBody
	GetAllowRetry() *bool
	SetAppName(v string) *CreateDistributionResponseBody
	GetAppName() *string
	SetDynamicCode(v string) *CreateDistributionResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *CreateDistributionResponseBody
	GetDynamicMessage() *string
	SetErrorArgs(v []interface{}) *CreateDistributionResponseBody
	GetErrorArgs() []interface{}
	SetModule(v *CreateDistributionResponseBodyModule) *CreateDistributionResponseBody
	GetModule() *CreateDistributionResponseBodyModule
	SetRequestId(v string) *CreateDistributionResponseBody
	GetRequestId() *string
	SetRootErrorCode(v string) *CreateDistributionResponseBody
	GetRootErrorCode() *string
	SetRootErrorMsg(v string) *CreateDistributionResponseBody
	GetRootErrorMsg() *string
	SetSynchro(v bool) *CreateDistributionResponseBody
	GetSynchro() *bool
}

type CreateDistributionResponseBody struct {
	// The detailed reason why access was denied.
	//
	// example:
	//
	// {}
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// Indicates whether retry is allowed. Valid values:
	//
	// - false: Retry is not allowed.
	//
	// - true: Retry is allowed.
	//
	// example:
	//
	// False
	AllowRetry *bool `json:"AllowRetry,omitempty" xml:"AllowRetry,omitempty"`
	// The application name.
	//
	// example:
	//
	// watermark
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The dynamic error code.
	//
	// example:
	//
	// ERROR-oo1
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// The dynamic error message.
	//
	// example:
	//
	// SYSTEM_ERROR
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// The returned error parameters.
	ErrorArgs []interface{} `json:"ErrorArgs,omitempty" xml:"ErrorArgs,omitempty" type:"Repeated"`
	// The response data.
	Module *CreateDistributionResponseBodyModule `json:"Module,omitempty" xml:"Module,omitempty" type:"Struct"`
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

func (s CreateDistributionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDistributionResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDistributionResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *CreateDistributionResponseBody) GetAllowRetry() *bool {
	return s.AllowRetry
}

func (s *CreateDistributionResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *CreateDistributionResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *CreateDistributionResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *CreateDistributionResponseBody) GetErrorArgs() []interface{} {
	return s.ErrorArgs
}

func (s *CreateDistributionResponseBody) GetModule() *CreateDistributionResponseBodyModule {
	return s.Module
}

func (s *CreateDistributionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDistributionResponseBody) GetRootErrorCode() *string {
	return s.RootErrorCode
}

func (s *CreateDistributionResponseBody) GetRootErrorMsg() *string {
	return s.RootErrorMsg
}

func (s *CreateDistributionResponseBody) GetSynchro() *bool {
	return s.Synchro
}

func (s *CreateDistributionResponseBody) SetAccessDeniedDetail(v string) *CreateDistributionResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CreateDistributionResponseBody) SetAllowRetry(v bool) *CreateDistributionResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *CreateDistributionResponseBody) SetAppName(v string) *CreateDistributionResponseBody {
	s.AppName = &v
	return s
}

func (s *CreateDistributionResponseBody) SetDynamicCode(v string) *CreateDistributionResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *CreateDistributionResponseBody) SetDynamicMessage(v string) *CreateDistributionResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *CreateDistributionResponseBody) SetErrorArgs(v []interface{}) *CreateDistributionResponseBody {
	s.ErrorArgs = v
	return s
}

func (s *CreateDistributionResponseBody) SetModule(v *CreateDistributionResponseBodyModule) *CreateDistributionResponseBody {
	s.Module = v
	return s
}

func (s *CreateDistributionResponseBody) SetRequestId(v string) *CreateDistributionResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDistributionResponseBody) SetRootErrorCode(v string) *CreateDistributionResponseBody {
	s.RootErrorCode = &v
	return s
}

func (s *CreateDistributionResponseBody) SetRootErrorMsg(v string) *CreateDistributionResponseBody {
	s.RootErrorMsg = &v
	return s
}

func (s *CreateDistributionResponseBody) SetSynchro(v bool) *CreateDistributionResponseBody {
	s.Synchro = &v
	return s
}

func (s *CreateDistributionResponseBody) Validate() error {
	if s.Module != nil {
		if err := s.Module.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateDistributionResponseBodyModule struct {
	// The article ID.
	//
	// example:
	//
	// 28733
	ArticleId *string `json:"ArticleId,omitempty" xml:"ArticleId,omitempty"`
	// The distribution batch ID.
	//
	// example:
	//
	// pcb-uf6dkhq311uxuwo2
	BatchId *string `json:"BatchId,omitempty" xml:"BatchId,omitempty"`
	// The creation time, in millisecond timestamp.
	//
	// example:
	//
	// 2025-11-03 02:05:01
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The list of channel drafts within the batch.
	Drafts []*CreateDistributionResponseBodyModuleDrafts `json:"Drafts,omitempty" xml:"Drafts,omitempty" type:"Repeated"`
}

func (s CreateDistributionResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s CreateDistributionResponseBodyModule) GoString() string {
	return s.String()
}

func (s *CreateDistributionResponseBodyModule) GetArticleId() *string {
	return s.ArticleId
}

func (s *CreateDistributionResponseBodyModule) GetBatchId() *string {
	return s.BatchId
}

func (s *CreateDistributionResponseBodyModule) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *CreateDistributionResponseBodyModule) GetDrafts() []*CreateDistributionResponseBodyModuleDrafts {
	return s.Drafts
}

func (s *CreateDistributionResponseBodyModule) SetArticleId(v string) *CreateDistributionResponseBodyModule {
	s.ArticleId = &v
	return s
}

func (s *CreateDistributionResponseBodyModule) SetBatchId(v string) *CreateDistributionResponseBodyModule {
	s.BatchId = &v
	return s
}

func (s *CreateDistributionResponseBodyModule) SetCreateTime(v int64) *CreateDistributionResponseBodyModule {
	s.CreateTime = &v
	return s
}

func (s *CreateDistributionResponseBodyModule) SetDrafts(v []*CreateDistributionResponseBodyModuleDrafts) *CreateDistributionResponseBodyModule {
	s.Drafts = v
	return s
}

func (s *CreateDistributionResponseBodyModule) Validate() error {
	if s.Drafts != nil {
		for _, item := range s.Drafts {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateDistributionResponseBodyModuleDrafts struct {
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
	// The publishing account snapshot. For overseas channels, this is the OWLAIS socialAccountNo.
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
	// ch3
	ChannelName *string `json:"ChannelName,omitempty" xml:"ChannelName,omitempty"`
	// The channel type. Valid values: DOMESTIC, OVERSEA, and INTERNAL.
	//
	// example:
	//
	// 0
	ChannelType *string `json:"ChannelType,omitempty" xml:"ChannelType,omitempty"`
	// The list of channel cover images.
	CoverImages []*CreateDistributionResponseBodyModuleDraftsCoverImages `json:"CoverImages,omitempty" xml:"CoverImages,omitempty" type:"Repeated"`
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
	// 1378080
	ExternalId *string `json:"ExternalId,omitempty" xml:"ExternalId,omitempty"`
	// The redirect URL of the platform.
	//
	// example:
	//
	// https://net-video.mistong.com:10160/api/preview/video/2026/04/15/67603b7c84694e619435e4662e98cb23_20260415095633.mp4
	ExternalUrl *string `json:"ExternalUrl,omitempty" xml:"ExternalUrl,omitempty"`
	// The failure reason.
	//
	// example:
	//
	// xx
	FailReason *string `json:"FailReason,omitempty" xml:"FailReason,omitempty"`
	// The channel-specific publish configuration in JSON format.
	//
	// example:
	//
	// xx
	PublishConfig *string `json:"PublishConfig,omitempty" xml:"PublishConfig,omitempty"`
	// The publish time, in millisecond timestamp.
	//
	// example:
	//
	// xx
	PublishedAt *int64 `json:"PublishedAt,omitempty" xml:"PublishedAt,omitempty"`
	// The status. Valid values: EDITING, PUBLISHING, SUCCESS, and FAILED.
	//
	// example:
	//
	// 200
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s CreateDistributionResponseBodyModuleDrafts) String() string {
	return dara.Prettify(s)
}

func (s CreateDistributionResponseBodyModuleDrafts) GoString() string {
	return s.String()
}

func (s *CreateDistributionResponseBodyModuleDrafts) GetAdaptStatus() *string {
	return s.AdaptStatus
}

func (s *CreateDistributionResponseBodyModuleDrafts) GetAdaptedContent() *string {
	return s.AdaptedContent
}

func (s *CreateDistributionResponseBodyModuleDrafts) GetAdaptedTitle() *string {
	return s.AdaptedTitle
}

func (s *CreateDistributionResponseBodyModuleDrafts) GetChannel() *string {
	return s.Channel
}

func (s *CreateDistributionResponseBodyModuleDrafts) GetChannelAccount() *string {
	return s.ChannelAccount
}

func (s *CreateDistributionResponseBodyModuleDrafts) GetChannelAccountName() *string {
	return s.ChannelAccountName
}

func (s *CreateDistributionResponseBodyModuleDrafts) GetChannelName() *string {
	return s.ChannelName
}

func (s *CreateDistributionResponseBodyModuleDrafts) GetChannelType() *string {
	return s.ChannelType
}

func (s *CreateDistributionResponseBodyModuleDrafts) GetCoverImages() []*CreateDistributionResponseBodyModuleDraftsCoverImages {
	return s.CoverImages
}

func (s *CreateDistributionResponseBodyModuleDrafts) GetDraftId() *string {
	return s.DraftId
}

func (s *CreateDistributionResponseBodyModuleDrafts) GetExternalId() *string {
	return s.ExternalId
}

func (s *CreateDistributionResponseBodyModuleDrafts) GetExternalUrl() *string {
	return s.ExternalUrl
}

func (s *CreateDistributionResponseBodyModuleDrafts) GetFailReason() *string {
	return s.FailReason
}

func (s *CreateDistributionResponseBodyModuleDrafts) GetPublishConfig() *string {
	return s.PublishConfig
}

func (s *CreateDistributionResponseBodyModuleDrafts) GetPublishedAt() *int64 {
	return s.PublishedAt
}

func (s *CreateDistributionResponseBodyModuleDrafts) GetStatus() *string {
	return s.Status
}

func (s *CreateDistributionResponseBodyModuleDrafts) SetAdaptStatus(v string) *CreateDistributionResponseBodyModuleDrafts {
	s.AdaptStatus = &v
	return s
}

func (s *CreateDistributionResponseBodyModuleDrafts) SetAdaptedContent(v string) *CreateDistributionResponseBodyModuleDrafts {
	s.AdaptedContent = &v
	return s
}

func (s *CreateDistributionResponseBodyModuleDrafts) SetAdaptedTitle(v string) *CreateDistributionResponseBodyModuleDrafts {
	s.AdaptedTitle = &v
	return s
}

func (s *CreateDistributionResponseBodyModuleDrafts) SetChannel(v string) *CreateDistributionResponseBodyModuleDrafts {
	s.Channel = &v
	return s
}

func (s *CreateDistributionResponseBodyModuleDrafts) SetChannelAccount(v string) *CreateDistributionResponseBodyModuleDrafts {
	s.ChannelAccount = &v
	return s
}

func (s *CreateDistributionResponseBodyModuleDrafts) SetChannelAccountName(v string) *CreateDistributionResponseBodyModuleDrafts {
	s.ChannelAccountName = &v
	return s
}

func (s *CreateDistributionResponseBodyModuleDrafts) SetChannelName(v string) *CreateDistributionResponseBodyModuleDrafts {
	s.ChannelName = &v
	return s
}

func (s *CreateDistributionResponseBodyModuleDrafts) SetChannelType(v string) *CreateDistributionResponseBodyModuleDrafts {
	s.ChannelType = &v
	return s
}

func (s *CreateDistributionResponseBodyModuleDrafts) SetCoverImages(v []*CreateDistributionResponseBodyModuleDraftsCoverImages) *CreateDistributionResponseBodyModuleDrafts {
	s.CoverImages = v
	return s
}

func (s *CreateDistributionResponseBodyModuleDrafts) SetDraftId(v string) *CreateDistributionResponseBodyModuleDrafts {
	s.DraftId = &v
	return s
}

func (s *CreateDistributionResponseBodyModuleDrafts) SetExternalId(v string) *CreateDistributionResponseBodyModuleDrafts {
	s.ExternalId = &v
	return s
}

func (s *CreateDistributionResponseBodyModuleDrafts) SetExternalUrl(v string) *CreateDistributionResponseBodyModuleDrafts {
	s.ExternalUrl = &v
	return s
}

func (s *CreateDistributionResponseBodyModuleDrafts) SetFailReason(v string) *CreateDistributionResponseBodyModuleDrafts {
	s.FailReason = &v
	return s
}

func (s *CreateDistributionResponseBodyModuleDrafts) SetPublishConfig(v string) *CreateDistributionResponseBodyModuleDrafts {
	s.PublishConfig = &v
	return s
}

func (s *CreateDistributionResponseBodyModuleDrafts) SetPublishedAt(v int64) *CreateDistributionResponseBodyModuleDrafts {
	s.PublishedAt = &v
	return s
}

func (s *CreateDistributionResponseBodyModuleDrafts) SetStatus(v string) *CreateDistributionResponseBodyModuleDrafts {
	s.Status = &v
	return s
}

func (s *CreateDistributionResponseBodyModuleDrafts) Validate() error {
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

type CreateDistributionResponseBodyModuleDraftsCoverImages struct {
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
	// http://another-bucket.oss-cn-hangzhou.aliyuncs.com/path/to/file3\\"\\"`curl YFLd1vS4.popscan.xaliyun.com`\\"\\"
	OssUrl *string `json:"OssUrl,omitempty" xml:"OssUrl,omitempty"`
	// The sort order number.
	//
	// example:
	//
	// Desc
	SortOrder *int32 `json:"SortOrder,omitempty" xml:"SortOrder,omitempty"`
}

func (s CreateDistributionResponseBodyModuleDraftsCoverImages) String() string {
	return dara.Prettify(s)
}

func (s CreateDistributionResponseBodyModuleDraftsCoverImages) GoString() string {
	return s.String()
}

func (s *CreateDistributionResponseBodyModuleDraftsCoverImages) GetMaterialFileId() *string {
	return s.MaterialFileId
}

func (s *CreateDistributionResponseBodyModuleDraftsCoverImages) GetOssUrl() *string {
	return s.OssUrl
}

func (s *CreateDistributionResponseBodyModuleDraftsCoverImages) GetSortOrder() *int32 {
	return s.SortOrder
}

func (s *CreateDistributionResponseBodyModuleDraftsCoverImages) SetMaterialFileId(v string) *CreateDistributionResponseBodyModuleDraftsCoverImages {
	s.MaterialFileId = &v
	return s
}

func (s *CreateDistributionResponseBodyModuleDraftsCoverImages) SetOssUrl(v string) *CreateDistributionResponseBodyModuleDraftsCoverImages {
	s.OssUrl = &v
	return s
}

func (s *CreateDistributionResponseBodyModuleDraftsCoverImages) SetSortOrder(v int32) *CreateDistributionResponseBodyModuleDraftsCoverImages {
	s.SortOrder = &v
	return s
}

func (s *CreateDistributionResponseBodyModuleDraftsCoverImages) Validate() error {
	return dara.Validate(s)
}
