// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyImagePermissionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddAccount(v []*int64) *ModifyImagePermissionRequest
	GetAddAccount() []*int64
	SetImageId(v string) *ModifyImagePermissionRequest
	GetImageId() *string
	SetRegionId(v string) *ModifyImagePermissionRequest
	GetRegionId() *string
	SetRemoveAccount(v []*int64) *ModifyImagePermissionRequest
	GetRemoveAccount() []*int64
}

type ModifyImagePermissionRequest struct {
	// The IDs of Alibaba Cloud accounts to which to share the image that will be created based on the image template. You can specify up to 20 account IDs.
	AddAccount []*int64 `json:"AddAccount,omitempty" xml:"AddAccount,omitempty" type:"Repeated"`
	// The IDs of the images.
	//
	// This parameter is required.
	//
	// example:
	//
	// m-gx2x1dhsmusr2****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/196646.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of Alibaba Cloud account N from which you want to unshare the custom image. Valid values of N: 1 to 10. If the value of N is greater than 10, this parameter is ignored.
	RemoveAccount []*int64 `json:"RemoveAccount,omitempty" xml:"RemoveAccount,omitempty" type:"Repeated"`
}

func (s ModifyImagePermissionRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyImagePermissionRequest) GoString() string {
	return s.String()
}

func (s *ModifyImagePermissionRequest) GetAddAccount() []*int64 {
	return s.AddAccount
}

func (s *ModifyImagePermissionRequest) GetImageId() *string {
	return s.ImageId
}

func (s *ModifyImagePermissionRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyImagePermissionRequest) GetRemoveAccount() []*int64 {
	return s.RemoveAccount
}

func (s *ModifyImagePermissionRequest) SetAddAccount(v []*int64) *ModifyImagePermissionRequest {
	s.AddAccount = v
	return s
}

func (s *ModifyImagePermissionRequest) SetImageId(v string) *ModifyImagePermissionRequest {
	s.ImageId = &v
	return s
}

func (s *ModifyImagePermissionRequest) SetRegionId(v string) *ModifyImagePermissionRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyImagePermissionRequest) SetRemoveAccount(v []*int64) *ModifyImagePermissionRequest {
	s.RemoveAccount = v
	return s
}

func (s *ModifyImagePermissionRequest) Validate() error {
	return dara.Validate(s)
}
