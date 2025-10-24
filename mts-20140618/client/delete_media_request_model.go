// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMediaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMediaIds(v string) *DeleteMediaRequest
	GetMediaIds() *string
	SetOwnerAccount(v string) *DeleteMediaRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteMediaRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DeleteMediaRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteMediaRequest
	GetResourceOwnerId() *int64
}

type DeleteMediaRequest struct {
	// The IDs of the media files that you want to remove. Separate multiple IDs with commas (,). You can remove up to 10 media files at a time.
	//
	// > You can obtain the ID of the media file from the response parameters of the [AddMedia](https://help.aliyun.com/document_detail/44458.html) operation. Alternatively, you can log on to the MPS console. In the left-side navigation pane, choose **Media Management*	- > **Media List**. Find the required video and click **Manage*	- in the Actions column. The ID of the video is displayed on the Basics tab.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3e1cd21131a94525be55acf65888****,3e6149d5a8c944c09b1a8d2dc3e4****
	MediaIds             *string `json:"MediaIds,omitempty" xml:"MediaIds,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DeleteMediaRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteMediaRequest) GoString() string {
	return s.String()
}

func (s *DeleteMediaRequest) GetMediaIds() *string {
	return s.MediaIds
}

func (s *DeleteMediaRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteMediaRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteMediaRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteMediaRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteMediaRequest) SetMediaIds(v string) *DeleteMediaRequest {
	s.MediaIds = &v
	return s
}

func (s *DeleteMediaRequest) SetOwnerAccount(v string) *DeleteMediaRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteMediaRequest) SetOwnerId(v int64) *DeleteMediaRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteMediaRequest) SetResourceOwnerAccount(v string) *DeleteMediaRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteMediaRequest) SetResourceOwnerId(v int64) *DeleteMediaRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteMediaRequest) Validate() error {
	return dara.Validate(s)
}
