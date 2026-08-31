// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyImageSharePermissionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddAccount(v []*string) *ModifyImageSharePermissionRequest
	GetAddAccount() []*string
	SetDryRun(v bool) *ModifyImageSharePermissionRequest
	GetDryRun() *bool
	SetImageId(v string) *ModifyImageSharePermissionRequest
	GetImageId() *string
	SetIsPublic(v bool) *ModifyImageSharePermissionRequest
	GetIsPublic() *bool
	SetLaunchPermission(v string) *ModifyImageSharePermissionRequest
	GetLaunchPermission() *string
	SetOwnerAccount(v string) *ModifyImageSharePermissionRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyImageSharePermissionRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ModifyImageSharePermissionRequest
	GetRegionId() *string
	SetRemoveAccount(v []*string) *ModifyImageSharePermissionRequest
	GetRemoveAccount() []*string
	SetResourceOwnerAccount(v string) *ModifyImageSharePermissionRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyImageSharePermissionRequest
	GetResourceOwnerId() *int64
}

type ModifyImageSharePermissionRequest struct {
	// The Alibaba Cloud account ID to which you want to grant authorization for the shared image. Valid values of N: 1 to 10. If you submit more than 10 Alibaba Cloud accounts in a single request, the system processes only the first 10 and ignores the rest.
	//
	// example:
	//
	// 1234567890
	AddAccount []*string `json:"AddAccount,omitempty" xml:"AddAccount,omitempty" type:"Repeated"`
	DryRun     *bool     `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ID of the custom image to be shared.
	//
	// 	Notice: Images encrypted with a service key can no longer be shared. Only images encrypted with a customer master key (CMK) can be shared. An error is returned if you attempt to share an image that is encrypted with a service key.
	//
	// This parameter is required.
	//
	// example:
	//
	// m-bp18ygjuqnwhechc****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// Specifies whether to publish or delist the image as a community image. Valid values:
	//
	// - true: Publishes the image as a community image.
	//
	// - false: Delists the image to a regular image. If the image is already a regular image, no change is made.
	//
	// Default value: false.
	//
	// example:
	//
	// false
	IsPublic *bool `json:"IsPublic,omitempty" xml:"IsPublic,omitempty"`
	// >This parameter is in invitational preview and is not available for use.
	//
	// example:
	//
	// hide
	LaunchPermission *string `json:"LaunchPermission,omitempty" xml:"LaunchPermission,omitempty"`
	OwnerAccount     *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId          *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the custom image. You can call [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The Alibaba Cloud account ID from which you want to delete image sharing. Valid values of N: 1 to 10. If you submit more than 10 Alibaba Cloud accounts in a single request, the system processes only the first 10 and ignores the rest.
	//
	// example:
	//
	// 1234567890
	RemoveAccount        []*string `json:"RemoveAccount,omitempty" xml:"RemoveAccount,omitempty" type:"Repeated"`
	ResourceOwnerAccount *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64    `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ModifyImageSharePermissionRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyImageSharePermissionRequest) GoString() string {
	return s.String()
}

func (s *ModifyImageSharePermissionRequest) GetAddAccount() []*string {
	return s.AddAccount
}

func (s *ModifyImageSharePermissionRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *ModifyImageSharePermissionRequest) GetImageId() *string {
	return s.ImageId
}

func (s *ModifyImageSharePermissionRequest) GetIsPublic() *bool {
	return s.IsPublic
}

func (s *ModifyImageSharePermissionRequest) GetLaunchPermission() *string {
	return s.LaunchPermission
}

func (s *ModifyImageSharePermissionRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyImageSharePermissionRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyImageSharePermissionRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyImageSharePermissionRequest) GetRemoveAccount() []*string {
	return s.RemoveAccount
}

func (s *ModifyImageSharePermissionRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyImageSharePermissionRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyImageSharePermissionRequest) SetAddAccount(v []*string) *ModifyImageSharePermissionRequest {
	s.AddAccount = v
	return s
}

func (s *ModifyImageSharePermissionRequest) SetDryRun(v bool) *ModifyImageSharePermissionRequest {
	s.DryRun = &v
	return s
}

func (s *ModifyImageSharePermissionRequest) SetImageId(v string) *ModifyImageSharePermissionRequest {
	s.ImageId = &v
	return s
}

func (s *ModifyImageSharePermissionRequest) SetIsPublic(v bool) *ModifyImageSharePermissionRequest {
	s.IsPublic = &v
	return s
}

func (s *ModifyImageSharePermissionRequest) SetLaunchPermission(v string) *ModifyImageSharePermissionRequest {
	s.LaunchPermission = &v
	return s
}

func (s *ModifyImageSharePermissionRequest) SetOwnerAccount(v string) *ModifyImageSharePermissionRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyImageSharePermissionRequest) SetOwnerId(v int64) *ModifyImageSharePermissionRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyImageSharePermissionRequest) SetRegionId(v string) *ModifyImageSharePermissionRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyImageSharePermissionRequest) SetRemoveAccount(v []*string) *ModifyImageSharePermissionRequest {
	s.RemoveAccount = v
	return s
}

func (s *ModifyImageSharePermissionRequest) SetResourceOwnerAccount(v string) *ModifyImageSharePermissionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyImageSharePermissionRequest) SetResourceOwnerId(v int64) *ModifyImageSharePermissionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyImageSharePermissionRequest) Validate() error {
	return dara.Validate(s)
}
