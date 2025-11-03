// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVSwitchRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDryRun(v bool) *DeleteVSwitchRequest
	GetDryRun() *bool
	SetOwnerAccount(v string) *DeleteVSwitchRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteVSwitchRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DeleteVSwitchRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DeleteVSwitchRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteVSwitchRequest
	GetResourceOwnerId() *int64
	SetVSwitchId(v string) *DeleteVSwitchRequest
	GetVSwitchId() *string
}

type DeleteVSwitchRequest struct {
	// Specifies whether to perform a dry run. Valid values:
	//
	// - **true**: performs a dry run. The system checks the required parameters, request syntax, and limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// - **false*	- (default): performs a dry run and sends the request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
	//
	// example:
	//
	// false
	DryRun       *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the vSwitch.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the vSwitch that you want to delete.
	//
	// This parameter is required.
	//
	// example:
	//
	// vsw-asdfjlna****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
}

func (s DeleteVSwitchRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteVSwitchRequest) GoString() string {
	return s.String()
}

func (s *DeleteVSwitchRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *DeleteVSwitchRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteVSwitchRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteVSwitchRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteVSwitchRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteVSwitchRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteVSwitchRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DeleteVSwitchRequest) SetDryRun(v bool) *DeleteVSwitchRequest {
	s.DryRun = &v
	return s
}

func (s *DeleteVSwitchRequest) SetOwnerAccount(v string) *DeleteVSwitchRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteVSwitchRequest) SetOwnerId(v int64) *DeleteVSwitchRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteVSwitchRequest) SetRegionId(v string) *DeleteVSwitchRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteVSwitchRequest) SetResourceOwnerAccount(v string) *DeleteVSwitchRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteVSwitchRequest) SetResourceOwnerId(v int64) *DeleteVSwitchRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteVSwitchRequest) SetVSwitchId(v string) *DeleteVSwitchRequest {
	s.VSwitchId = &v
	return s
}

func (s *DeleteVSwitchRequest) Validate() error {
	return dara.Validate(s)
}
