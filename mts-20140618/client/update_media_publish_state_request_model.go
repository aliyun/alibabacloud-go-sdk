// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMediaPublishStateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMediaId(v string) *UpdateMediaPublishStateRequest
	GetMediaId() *string
	SetOwnerAccount(v string) *UpdateMediaPublishStateRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *UpdateMediaPublishStateRequest
	GetOwnerId() *int64
	SetPublish(v bool) *UpdateMediaPublishStateRequest
	GetPublish() *bool
	SetResourceOwnerAccount(v string) *UpdateMediaPublishStateRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *UpdateMediaPublishStateRequest
	GetResourceOwnerId() *int64
}

type UpdateMediaPublishStateRequest struct {
	// The ID of the media file whose publishing status you want to update. You can obtain the ID of a media file from the response of the [AddMedia](https://help.aliyun.com/document_detail/44458.html) operation. Alternatively, perform the following operations in the ApsaraVideo Media Processing (MPS) console: In the left-side navigation pane, choose **Media Management*	- > **Media List**. Find the required video and click **Manage**. The ID of the video is displayed on the Basics tab.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3e6149d5a8c944c09b1a8d2dc3e4****
	MediaId      *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The publishing status. Default value: **Initialed**. Valid values:
	//
	// 	- **true**: published.
	//
	// 	- **false**: unpublished.
	//
	// example:
	//
	// true
	Publish              *bool   `json:"Publish,omitempty" xml:"Publish,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s UpdateMediaPublishStateRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateMediaPublishStateRequest) GoString() string {
	return s.String()
}

func (s *UpdateMediaPublishStateRequest) GetMediaId() *string {
	return s.MediaId
}

func (s *UpdateMediaPublishStateRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *UpdateMediaPublishStateRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UpdateMediaPublishStateRequest) GetPublish() *bool {
	return s.Publish
}

func (s *UpdateMediaPublishStateRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *UpdateMediaPublishStateRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *UpdateMediaPublishStateRequest) SetMediaId(v string) *UpdateMediaPublishStateRequest {
	s.MediaId = &v
	return s
}

func (s *UpdateMediaPublishStateRequest) SetOwnerAccount(v string) *UpdateMediaPublishStateRequest {
	s.OwnerAccount = &v
	return s
}

func (s *UpdateMediaPublishStateRequest) SetOwnerId(v int64) *UpdateMediaPublishStateRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateMediaPublishStateRequest) SetPublish(v bool) *UpdateMediaPublishStateRequest {
	s.Publish = &v
	return s
}

func (s *UpdateMediaPublishStateRequest) SetResourceOwnerAccount(v string) *UpdateMediaPublishStateRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpdateMediaPublishStateRequest) SetResourceOwnerId(v int64) *UpdateMediaPublishStateRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UpdateMediaPublishStateRequest) Validate() error {
	return dara.Validate(s)
}
