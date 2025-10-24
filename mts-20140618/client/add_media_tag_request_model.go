// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddMediaTagRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMediaId(v string) *AddMediaTagRequest
	GetMediaId() *string
	SetOwnerAccount(v string) *AddMediaTagRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *AddMediaTagRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *AddMediaTagRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *AddMediaTagRequest
	GetResourceOwnerId() *int64
	SetTag(v string) *AddMediaTagRequest
	GetTag() *string
}

type AddMediaTagRequest struct {
	// The ID of the media file to which you want to add tags.
	//
	// > To obtain the ID of a media file, you can call the [AddMedia](https://help.aliyun.com/document_detail/44458.html) operation. Alternatively, perform the following operations in the ApsaraVideo Media Processing (MPS) console: In the left-side navigation pane, choose **Media Management*	- > **Media List**. Find the file that you want to manage and click **Manage*	- in the Actions column. The ID of the file is displayed on the Basics tab.
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
	// The tag that you want to add to the medial file. The value is encoded in UTF-8 and can be up to 32 bytes in length.
	//
	// example:
	//
	// tag1
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
}

func (s AddMediaTagRequest) String() string {
	return dara.Prettify(s)
}

func (s AddMediaTagRequest) GoString() string {
	return s.String()
}

func (s *AddMediaTagRequest) GetMediaId() *string {
	return s.MediaId
}

func (s *AddMediaTagRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *AddMediaTagRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AddMediaTagRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *AddMediaTagRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *AddMediaTagRequest) GetTag() *string {
	return s.Tag
}

func (s *AddMediaTagRequest) SetMediaId(v string) *AddMediaTagRequest {
	s.MediaId = &v
	return s
}

func (s *AddMediaTagRequest) SetOwnerAccount(v string) *AddMediaTagRequest {
	s.OwnerAccount = &v
	return s
}

func (s *AddMediaTagRequest) SetOwnerId(v int64) *AddMediaTagRequest {
	s.OwnerId = &v
	return s
}

func (s *AddMediaTagRequest) SetResourceOwnerAccount(v string) *AddMediaTagRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AddMediaTagRequest) SetResourceOwnerId(v int64) *AddMediaTagRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AddMediaTagRequest) SetTag(v string) *AddMediaTagRequest {
	s.Tag = &v
	return s
}

func (s *AddMediaTagRequest) Validate() error {
	return dara.Validate(s)
}
