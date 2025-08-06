// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitAIJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfig(v string) *SubmitAIJobRequest
	GetConfig() *string
	SetMediaId(v string) *SubmitAIJobRequest
	GetMediaId() *string
	SetOwnerAccount(v string) *SubmitAIJobRequest
	GetOwnerAccount() *string
	SetOwnerId(v string) *SubmitAIJobRequest
	GetOwnerId() *string
	SetResourceOwnerAccount(v string) *SubmitAIJobRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v string) *SubmitAIJobRequest
	GetResourceOwnerId() *string
	SetTypes(v string) *SubmitAIJobRequest
	GetTypes() *string
	SetUserData(v string) *SubmitAIJobRequest
	GetUserData() *string
}

type SubmitAIJobRequest struct {
	// The configurations of the AI job. The value is a JSON string.
	//
	// 	- If you set `Types` to `AIVideoTag`, you can specify `AnalyseTypes` for `Config` to set the analysis algorithm of a smart tagging job. Valid values:
	//
	//     	- ASR: automatic speech recognition (ASR)
	//
	//     	- OCR: image optical character recognition (OCR)
	//
	// 	- If you set `Types` to `AIMediaDNA`, you can specify `DNADBId` for `Config` to set the ID of the media fingerprint library for video fingerprinting jobs.
	//
	// example:
	//
	// {"AIVideoTag": {"AnalyseTypes": "Face,ASR"} }
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// The ID of the video. You can use one of the following methods to obtain the ID:
	//
	// 	- Log on to the [ApsaraVideo VOD](https://vod.console.aliyun.com) console. In the left-side navigation pane, choose **Media Files*	- > **Audio/Video**. On the Video and Audio page, view the ID of the audio or video file. This method is applicable to files that are uploaded by using the ApsaraVideo VOD console.
	//
	// 	- Obtain the value of VideoId from the response to the [CreateUploadVideo](https://help.aliyun.com/document_detail/55407.html) operation that you call to upload media files.
	//
	// 	- Obtain the value of VideoId from the response to the [SearchMedia](https://help.aliyun.com/document_detail/86044.html) operation after you upload media files.
	//
	// example:
	//
	// 3D3D12340d9401fab46a0b847****
	MediaId              *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *string `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The type of the AI job. Separate multiple types with commas (,). Valid values:
	//
	// 	- **AIMediaDNA**: The media fingerprinting job.
	//
	// 	- **AIVideoTag**: The smart tagging job.
	//
	// example:
	//
	// AIVideoTag
	Types *string `json:"Types,omitempty" xml:"Types,omitempty"`
	// The custom settings. The value is a JSON string. For more information, see [Request parameters](~~86952#h2--userdata-div-id-userdata-div-3~~).
	//
	// example:
	//
	// {"Extend":{"localId":"***","test":"www"}}
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s SubmitAIJobRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitAIJobRequest) GoString() string {
	return s.String()
}

func (s *SubmitAIJobRequest) GetConfig() *string {
	return s.Config
}

func (s *SubmitAIJobRequest) GetMediaId() *string {
	return s.MediaId
}

func (s *SubmitAIJobRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *SubmitAIJobRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *SubmitAIJobRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *SubmitAIJobRequest) GetResourceOwnerId() *string {
	return s.ResourceOwnerId
}

func (s *SubmitAIJobRequest) GetTypes() *string {
	return s.Types
}

func (s *SubmitAIJobRequest) GetUserData() *string {
	return s.UserData
}

func (s *SubmitAIJobRequest) SetConfig(v string) *SubmitAIJobRequest {
	s.Config = &v
	return s
}

func (s *SubmitAIJobRequest) SetMediaId(v string) *SubmitAIJobRequest {
	s.MediaId = &v
	return s
}

func (s *SubmitAIJobRequest) SetOwnerAccount(v string) *SubmitAIJobRequest {
	s.OwnerAccount = &v
	return s
}

func (s *SubmitAIJobRequest) SetOwnerId(v string) *SubmitAIJobRequest {
	s.OwnerId = &v
	return s
}

func (s *SubmitAIJobRequest) SetResourceOwnerAccount(v string) *SubmitAIJobRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SubmitAIJobRequest) SetResourceOwnerId(v string) *SubmitAIJobRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SubmitAIJobRequest) SetTypes(v string) *SubmitAIJobRequest {
	s.Types = &v
	return s
}

func (s *SubmitAIJobRequest) SetUserData(v string) *SubmitAIJobRequest {
	s.UserData = &v
	return s
}

func (s *SubmitAIJobRequest) Validate() error {
	return dara.Validate(s)
}
