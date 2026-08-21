// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAIVideoTagResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMediaId(v string) *GetAIVideoTagResultRequest
	GetMediaId() *string
	SetOwnerAccount(v string) *GetAIVideoTagResultRequest
	GetOwnerAccount() *string
	SetOwnerId(v string) *GetAIVideoTagResultRequest
	GetOwnerId() *string
	SetResourceOwnerAccount(v string) *GetAIVideoTagResultRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v string) *GetAIVideoTagResultRequest
	GetResourceOwnerId() *string
}

type GetAIVideoTagResultRequest struct {
	// The media asset ID, which is the video ID. You can obtain the video ID by using the following methods:
	//
	// - For videos uploaded through the console, logon to the [ApsaraVideo VOD console](https://vod.console.aliyun.com) and choose **Media Files*	- > **Audio/Video*	- to view the video ID.
	//
	// - When you call the [CreateUploadVideo](https://help.aliyun.com/document_detail/55407.html) operation to obtain the upload URL and credential, the video ID is the value of the VideoId parameter in the response.
	//
	// - After the video is uploaded, call the [SearchMedia](https://help.aliyun.com/document_detail/86044.html) operation to query the video ID. The video ID is the value of the VideoId parameter in the response.
	//
	// This parameter is required.
	//
	// example:
	//
	// 595d020bad37421f37433451720****
	MediaId              *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *string `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s GetAIVideoTagResultRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAIVideoTagResultRequest) GoString() string {
	return s.String()
}

func (s *GetAIVideoTagResultRequest) GetMediaId() *string {
	return s.MediaId
}

func (s *GetAIVideoTagResultRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *GetAIVideoTagResultRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *GetAIVideoTagResultRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *GetAIVideoTagResultRequest) GetResourceOwnerId() *string {
	return s.ResourceOwnerId
}

func (s *GetAIVideoTagResultRequest) SetMediaId(v string) *GetAIVideoTagResultRequest {
	s.MediaId = &v
	return s
}

func (s *GetAIVideoTagResultRequest) SetOwnerAccount(v string) *GetAIVideoTagResultRequest {
	s.OwnerAccount = &v
	return s
}

func (s *GetAIVideoTagResultRequest) SetOwnerId(v string) *GetAIVideoTagResultRequest {
	s.OwnerId = &v
	return s
}

func (s *GetAIVideoTagResultRequest) SetResourceOwnerAccount(v string) *GetAIVideoTagResultRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetAIVideoTagResultRequest) SetResourceOwnerId(v string) *GetAIVideoTagResultRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetAIVideoTagResultRequest) Validate() error {
	return dara.Validate(s)
}
