// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyImageSharePermissionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddAccounts(v string) *ModifyImageSharePermissionRequest
	GetAddAccounts() *string
	SetImageId(v string) *ModifyImageSharePermissionRequest
	GetImageId() *string
	SetRemoveAccounts(v string) *ModifyImageSharePermissionRequest
	GetRemoveAccounts() *string
}

type ModifyImageSharePermissionRequest struct {
	// The ID of the Alibaba Cloud account with which you want to share the image. You can specify multiple Alibaba Cloud IDs. Separate multiple IDs with commas (,).
	//
	// example:
	//
	// 1122334455**
	AddAccounts *string `json:"AddAccounts,omitempty" xml:"AddAccounts,omitempty"`
	// The ID of the image. You can specify only one image ID. Custom images and public images are supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// m-5s7qotzavwbrnzaqh4unm****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The ID of the Alibaba Cloud account from which you want to unshare the image. You can specify only one Alibaba Cloud account ID.
	//
	// example:
	//
	// 113355**
	RemoveAccounts *string `json:"RemoveAccounts,omitempty" xml:"RemoveAccounts,omitempty"`
}

func (s ModifyImageSharePermissionRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyImageSharePermissionRequest) GoString() string {
	return s.String()
}

func (s *ModifyImageSharePermissionRequest) GetAddAccounts() *string {
	return s.AddAccounts
}

func (s *ModifyImageSharePermissionRequest) GetImageId() *string {
	return s.ImageId
}

func (s *ModifyImageSharePermissionRequest) GetRemoveAccounts() *string {
	return s.RemoveAccounts
}

func (s *ModifyImageSharePermissionRequest) SetAddAccounts(v string) *ModifyImageSharePermissionRequest {
	s.AddAccounts = &v
	return s
}

func (s *ModifyImageSharePermissionRequest) SetImageId(v string) *ModifyImageSharePermissionRequest {
	s.ImageId = &v
	return s
}

func (s *ModifyImageSharePermissionRequest) SetRemoveAccounts(v string) *ModifyImageSharePermissionRequest {
	s.RemoveAccounts = &v
	return s
}

func (s *ModifyImageSharePermissionRequest) Validate() error {
	return dara.Validate(s)
}
