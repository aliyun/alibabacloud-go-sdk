// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMediaCategoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCateId(v int64) *UpdateMediaCategoryRequest
	GetCateId() *int64
	SetMediaId(v string) *UpdateMediaCategoryRequest
	GetMediaId() *string
	SetOwnerAccount(v string) *UpdateMediaCategoryRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *UpdateMediaCategoryRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *UpdateMediaCategoryRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *UpdateMediaCategoryRequest
	GetResourceOwnerId() *int64
}

type UpdateMediaCategoryRequest struct {
	// The ID of the category. The value cannot be negative.
	//
	// example:
	//
	// 1
	CateId *int64 `json:"CateId,omitempty" xml:"CateId,omitempty"`
	// The ID of the media file whose category you want to update.
	//
	// > To obtain the ID of a media file, you can call the [AddMedia](https://help.aliyun.com/document_detail/44458.html) operation. Alternatively, perform the following operations in the ApsaraVideo Media Processing (MPS) console: In the left-side navigation pane, choose **Media Management > Media List**. Find the required video and click **Manage*	- in the Actions column. The ID of the video is displayed on the Basics tab.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3e1cd21131a94525be55acf65888****
	MediaId              *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s UpdateMediaCategoryRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateMediaCategoryRequest) GoString() string {
	return s.String()
}

func (s *UpdateMediaCategoryRequest) GetCateId() *int64 {
	return s.CateId
}

func (s *UpdateMediaCategoryRequest) GetMediaId() *string {
	return s.MediaId
}

func (s *UpdateMediaCategoryRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *UpdateMediaCategoryRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UpdateMediaCategoryRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *UpdateMediaCategoryRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *UpdateMediaCategoryRequest) SetCateId(v int64) *UpdateMediaCategoryRequest {
	s.CateId = &v
	return s
}

func (s *UpdateMediaCategoryRequest) SetMediaId(v string) *UpdateMediaCategoryRequest {
	s.MediaId = &v
	return s
}

func (s *UpdateMediaCategoryRequest) SetOwnerAccount(v string) *UpdateMediaCategoryRequest {
	s.OwnerAccount = &v
	return s
}

func (s *UpdateMediaCategoryRequest) SetOwnerId(v int64) *UpdateMediaCategoryRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateMediaCategoryRequest) SetResourceOwnerAccount(v string) *UpdateMediaCategoryRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpdateMediaCategoryRequest) SetResourceOwnerId(v int64) *UpdateMediaCategoryRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UpdateMediaCategoryRequest) Validate() error {
	return dara.Validate(s)
}
