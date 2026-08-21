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
	// The AI job configuration in JSON format.
	//
	// - If `Types` is set to `AIVideoTag`, `Config` supports the `AnalyseTypes` parameter to specify the analysis algorithm types for the intelligent tagging job. Valid values:
	//
	//   - ASR: speech recognition. Identifies tags from the audio speech in the video.
	//
	//   - OCR: optical character recognition. Identifies tags from the text in the video images.
	//
	// - If `Types` is set to `AIMediaDNA`, `Config` supports the `DNADBId` parameter to specify the fingerprint library ID for the media fingerprint job.
	//
	// example:
	//
	// {"AIVideoTag": {"AnalyseTypes": "ASR"} }
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// The video ID. You can obtain the video ID by using one of the following methods:
	//
	// - For videos uploaded in the console, log on to the [ApsaraVideo VOD console](https://vod.console.aliyun.com) and choose **Media Files*	- > **Audio/Video*	- to view the video ID.
	//
	// - When you call the [CreateUploadVideo](https://help.aliyun.com/document_detail/55407.html) operation to obtain the upload URL and credential, the video ID is the value of the VideoId response parameter.
	//
	// - After the video is uploaded, you can call the [SearchMedia](https://help.aliyun.com/document_detail/86044.html) operation to query the video ID, which is the value of the VideoId response parameter.
	//
	// example:
	//
	// 3D3D12340d9401fab46a0b847****
	MediaId              *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *string `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The AI job type. Separate multiple job types with commas (,). Valid values:
	//
	// - **AIMediaDNA**: media fingerprint.
	//
	// - **AIVideoTag**: intelligent tagging.
	//
	// example:
	//
	// AIVideoTag
	Types *string `json:"Types,omitempty" xml:"Types,omitempty"`
	// The custom settings in JSON format. For more information about the parameter structure, see [UserData](~~86952#h2--userdata-div-id-userdata-div-3~~).
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
