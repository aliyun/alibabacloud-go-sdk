// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iProduceEditingProjectVideoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *ProduceEditingProjectVideoRequest
	GetAppId() *string
	SetCoverURL(v string) *ProduceEditingProjectVideoRequest
	GetCoverURL() *string
	SetDescription(v string) *ProduceEditingProjectVideoRequest
	GetDescription() *string
	SetMediaMetadata(v string) *ProduceEditingProjectVideoRequest
	GetMediaMetadata() *string
	SetOwnerId(v int64) *ProduceEditingProjectVideoRequest
	GetOwnerId() *int64
	SetProduceConfig(v string) *ProduceEditingProjectVideoRequest
	GetProduceConfig() *string
	SetProjectId(v string) *ProduceEditingProjectVideoRequest
	GetProjectId() *string
	SetResourceOwnerAccount(v string) *ProduceEditingProjectVideoRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ProduceEditingProjectVideoRequest
	GetResourceOwnerId() *int64
	SetTimeline(v string) *ProduceEditingProjectVideoRequest
	GetTimeline() *string
	SetTitle(v string) *ProduceEditingProjectVideoRequest
	GetTitle() *string
	SetUserData(v string) *ProduceEditingProjectVideoRequest
	GetUserData() *string
}

type ProduceEditingProjectVideoRequest struct {
	// The ID of the application. Default value: **app-1000000**. For more information, see [Multi-application service](https://help.aliyun.com/document_detail/113600.html).
	//
	// example:
	//
	// app-****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The thumbnail URL of the online editing project.
	//
	// example:
	//
	// https://example.aliyundoc.com/6AB4D0E1E1C7446888351****.png
	CoverURL *string `json:"CoverURL,omitempty" xml:"CoverURL,omitempty"`
	// The description of the online editing project.
	//
	// example:
	//
	// description test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The video metadata. The value must be in JSON format. For more information about the parameter structure, see [MediaMetadata](~~52839#title_rtf_ry5_gjp~~).
	//
	// example:
	//
	// {"Description":"video description","Title":"userData test"}
	MediaMetadata *string `json:"MediaMetadata,omitempty" xml:"MediaMetadata,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The configuration of video production. The value must be in the JSON format. For more information about the parameter structure, see [ProduceConfig](~~52839#title-ybl-7cs-y7d~~).
	//
	// >  StorageLocation is required if you produce videos in a region other than China (Shanghai).
	//
	// example:
	//
	// {"TemplateGroupId":"6d11e25ea30a4c465435c74****"}
	ProduceConfig *string `json:"ProduceConfig,omitempty" xml:"ProduceConfig,omitempty"`
	// The ID of the online editing project. You can use one of the following methods to obtain the ID of the online editing project:
	//
	// 	- Log on to the [ApsaraVideo VOD console](https://vod.console.aliyun.com). In the left-side navigation pane, choose **Production Center*	- > **Video Editing*	- to view the ID of the online editing project.
	//
	// 	- Obtain the value of ProjectId from the response to the [AddEditingProject](https://help.aliyun.com/document_detail/69048.html) operation.
	//
	// example:
	//
	// fb2101bf24b4cb318787dc****
	ProjectId            *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The timeline of the online editing project. The value must be in JSON format. For more information about the parameter structure, see [Timeline](~~52839#07bc7fe0f2xuh~~).
	//
	// example:
	//
	// {"VideoTracks":[{"VideoTrackClips":[{"MediaId":"cc3308ac59615a54328bc3443****"},{"MediaId":"da87a9cff645cd88bc6d8326e4****"}]}]}
	Timeline *string `json:"Timeline,omitempty" xml:"Timeline,omitempty"`
	// The title of the online editing project.
	//
	// example:
	//
	// editing project test
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// The custom configurations, such as the callback configuration. The value must be a JSON string. For more information about the parameter structure, see [UserData](~~86952#title_vz7_xzs_0c5~~).
	//
	// > The callback configurations take effect only after you specify an HTTP URL for receiving callback notifications and select the event types in the ApsaraVideo VOD console.
	//
	// example:
	//
	// {"Extend":{"width":1280,"id":"028a8e56b1ebf6bb7afc74****","height":720},"MessageCallback":{"CallbackURL":"https://example.aliyundoc.com/2016-08-15/proxy/httpcallback/testcallback/","CallbackType":"http"}}
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s ProduceEditingProjectVideoRequest) String() string {
	return dara.Prettify(s)
}

func (s ProduceEditingProjectVideoRequest) GoString() string {
	return s.String()
}

func (s *ProduceEditingProjectVideoRequest) GetAppId() *string {
	return s.AppId
}

func (s *ProduceEditingProjectVideoRequest) GetCoverURL() *string {
	return s.CoverURL
}

func (s *ProduceEditingProjectVideoRequest) GetDescription() *string {
	return s.Description
}

func (s *ProduceEditingProjectVideoRequest) GetMediaMetadata() *string {
	return s.MediaMetadata
}

func (s *ProduceEditingProjectVideoRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ProduceEditingProjectVideoRequest) GetProduceConfig() *string {
	return s.ProduceConfig
}

func (s *ProduceEditingProjectVideoRequest) GetProjectId() *string {
	return s.ProjectId
}

func (s *ProduceEditingProjectVideoRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ProduceEditingProjectVideoRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ProduceEditingProjectVideoRequest) GetTimeline() *string {
	return s.Timeline
}

func (s *ProduceEditingProjectVideoRequest) GetTitle() *string {
	return s.Title
}

func (s *ProduceEditingProjectVideoRequest) GetUserData() *string {
	return s.UserData
}

func (s *ProduceEditingProjectVideoRequest) SetAppId(v string) *ProduceEditingProjectVideoRequest {
	s.AppId = &v
	return s
}

func (s *ProduceEditingProjectVideoRequest) SetCoverURL(v string) *ProduceEditingProjectVideoRequest {
	s.CoverURL = &v
	return s
}

func (s *ProduceEditingProjectVideoRequest) SetDescription(v string) *ProduceEditingProjectVideoRequest {
	s.Description = &v
	return s
}

func (s *ProduceEditingProjectVideoRequest) SetMediaMetadata(v string) *ProduceEditingProjectVideoRequest {
	s.MediaMetadata = &v
	return s
}

func (s *ProduceEditingProjectVideoRequest) SetOwnerId(v int64) *ProduceEditingProjectVideoRequest {
	s.OwnerId = &v
	return s
}

func (s *ProduceEditingProjectVideoRequest) SetProduceConfig(v string) *ProduceEditingProjectVideoRequest {
	s.ProduceConfig = &v
	return s
}

func (s *ProduceEditingProjectVideoRequest) SetProjectId(v string) *ProduceEditingProjectVideoRequest {
	s.ProjectId = &v
	return s
}

func (s *ProduceEditingProjectVideoRequest) SetResourceOwnerAccount(v string) *ProduceEditingProjectVideoRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ProduceEditingProjectVideoRequest) SetResourceOwnerId(v int64) *ProduceEditingProjectVideoRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ProduceEditingProjectVideoRequest) SetTimeline(v string) *ProduceEditingProjectVideoRequest {
	s.Timeline = &v
	return s
}

func (s *ProduceEditingProjectVideoRequest) SetTitle(v string) *ProduceEditingProjectVideoRequest {
	s.Title = &v
	return s
}

func (s *ProduceEditingProjectVideoRequest) SetUserData(v string) *ProduceEditingProjectVideoRequest {
	s.UserData = &v
	return s
}

func (s *ProduceEditingProjectVideoRequest) Validate() error {
	return dara.Validate(s)
}
