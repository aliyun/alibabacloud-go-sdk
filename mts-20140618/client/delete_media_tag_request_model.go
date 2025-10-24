// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMediaTagRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMediaId(v string) *DeleteMediaTagRequest
	GetMediaId() *string
	SetOwnerAccount(v string) *DeleteMediaTagRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteMediaTagRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DeleteMediaTagRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteMediaTagRequest
	GetResourceOwnerId() *int64
	SetTag(v string) *DeleteMediaTagRequest
	GetTag() *string
}

type DeleteMediaTagRequest struct {
	// The ID of the media file for which you want to remove a tag. To obtain the ID of a media file, you can call the [AddMedia](https://help.aliyun.com/document_detail/44458.html) operation. Alternatively, perform the following operations in the ApsaraVideo Media Processing (MPS) console: In the left-side navigation pane, choose **Media Management*	- > **Media List**. Find the required video and click **Manage*	- in the Actions column. The ID of the video is displayed on the Basics tab.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3e6149d5a8c944c09b1a8d2dc3e4****
	MediaId              *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The media tag that you want to remove. The value is encoded in UTF-8 and can be up to 32 bytes in length.
	//
	// > You can remove only one tag at a time.
	//
	// example:
	//
	// tag1
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
}

func (s DeleteMediaTagRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteMediaTagRequest) GoString() string {
	return s.String()
}

func (s *DeleteMediaTagRequest) GetMediaId() *string {
	return s.MediaId
}

func (s *DeleteMediaTagRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteMediaTagRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteMediaTagRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteMediaTagRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteMediaTagRequest) GetTag() *string {
	return s.Tag
}

func (s *DeleteMediaTagRequest) SetMediaId(v string) *DeleteMediaTagRequest {
	s.MediaId = &v
	return s
}

func (s *DeleteMediaTagRequest) SetOwnerAccount(v string) *DeleteMediaTagRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteMediaTagRequest) SetOwnerId(v int64) *DeleteMediaTagRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteMediaTagRequest) SetResourceOwnerAccount(v string) *DeleteMediaTagRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteMediaTagRequest) SetResourceOwnerId(v int64) *DeleteMediaTagRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteMediaTagRequest) SetTag(v string) *DeleteMediaTagRequest {
	s.Tag = &v
	return s
}

func (s *DeleteMediaTagRequest) Validate() error {
	return dara.Validate(s)
}
