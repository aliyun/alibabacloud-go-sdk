// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRefreshUploadVideoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *RefreshUploadVideoRequest
	GetOwnerId() *int64
	SetReferenceId(v string) *RefreshUploadVideoRequest
	GetReferenceId() *string
	SetResourceOwnerAccount(v string) *RefreshUploadVideoRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *RefreshUploadVideoRequest
	GetResourceOwnerId() *int64
	SetVideoId(v string) *RefreshUploadVideoRequest
	GetVideoId() *string
}

type RefreshUploadVideoRequest struct {
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// 123-123
	ReferenceId          *string `json:"ReferenceId,omitempty" xml:"ReferenceId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the audio or video file. You can use one of the following methods to obtain the ID:
	//
	// 	- Log on to the [ApsaraVideo VOD console](https://vod.console.aliyun.com) and choose **Media Files*	- > **Audio/Video*	- in the left-side navigation pane to view the ID.
	//
	// 	- View the value of the VideoId parameter returned by the [CreateUploadVideo](https://help.aliyun.com/document_detail/55407.html) operation that you called to upload the audio or video file.
	//
	// 	- After an audio or video file is uploaded, obtain the value of VideoId from the response to the [SearchMedia](https://help.aliyun.com/document_detail/86044.html) operation that you call to query the audio or video ID.
	//
	// example:
	//
	// c6a23a870c8c4ffcd40cbd381333****
	VideoId *string `json:"VideoId,omitempty" xml:"VideoId,omitempty"`
}

func (s RefreshUploadVideoRequest) String() string {
	return dara.Prettify(s)
}

func (s RefreshUploadVideoRequest) GoString() string {
	return s.String()
}

func (s *RefreshUploadVideoRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *RefreshUploadVideoRequest) GetReferenceId() *string {
	return s.ReferenceId
}

func (s *RefreshUploadVideoRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *RefreshUploadVideoRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *RefreshUploadVideoRequest) GetVideoId() *string {
	return s.VideoId
}

func (s *RefreshUploadVideoRequest) SetOwnerId(v int64) *RefreshUploadVideoRequest {
	s.OwnerId = &v
	return s
}

func (s *RefreshUploadVideoRequest) SetReferenceId(v string) *RefreshUploadVideoRequest {
	s.ReferenceId = &v
	return s
}

func (s *RefreshUploadVideoRequest) SetResourceOwnerAccount(v string) *RefreshUploadVideoRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *RefreshUploadVideoRequest) SetResourceOwnerId(v int64) *RefreshUploadVideoRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *RefreshUploadVideoRequest) SetVideoId(v string) *RefreshUploadVideoRequest {
	s.VideoId = &v
	return s
}

func (s *RefreshUploadVideoRequest) Validate() error {
	return dara.Validate(s)
}
