// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetDesktopsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDesktopGroupId(v string) *ResetDesktopsRequest
	GetDesktopGroupId() *string
	SetDesktopGroupIds(v []*string) *ResetDesktopsRequest
	GetDesktopGroupIds() []*string
	SetDesktopId(v []*string) *ResetDesktopsRequest
	GetDesktopId() []*string
	SetImageId(v string) *ResetDesktopsRequest
	GetImageId() *string
	SetPayType(v string) *ResetDesktopsRequest
	GetPayType() *string
	SetRegionId(v string) *ResetDesktopsRequest
	GetRegionId() *string
	SetResetScope(v string) *ResetDesktopsRequest
	GetResetScope() *string
	SetResetType(v string) *ResetDesktopsRequest
	GetResetType() *string
}

type ResetDesktopsRequest struct {
	// The ID of the cloud computer share.
	//
	// 	- If you specify `DesktopId`, ignore `DesktopGroupId`.
	//
	// 	- If you leave `DesktopId` empty, the system obtains the IDs of all cloud computers within the share specified by `DesktopGroupId`.``
	//
	// example:
	//
	// dg-07if7qsxoxkb6****
	DesktopGroupId *string `json:"DesktopGroupId,omitempty" xml:"DesktopGroupId,omitempty"`
	// The IDs of the cloud computer shares.
	DesktopGroupIds []*string `json:"DesktopGroupIds,omitempty" xml:"DesktopGroupIds,omitempty" type:"Repeated"`
	// The IDs of the cloud computers. You can specify the IDs of 1 to 100 cloud computers.
	DesktopId []*string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty" type:"Repeated"`
	// The ID of the image.
	//
	// example:
	//
	// m-4zfb6zj728hhr****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The billing method of the cloud computer share.
	//
	// >  This parameter takes effect when you reset a cloud computer share. If you leave this parameter empty, all cloud computers in that share are reset.
	//
	// Valid values:
	//
	// 	- PostPaid: pay-as-you-go.
	//
	// 	- PrePaid: subscription.
	//
	// example:
	//
	// PrePaid
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/436773.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The reset scope. You can configure this parameter to reset the image or cloud computer.
	//
	// Valid values:
	//
	// 	- ALL (default): resets the image and cloud computer.
	//
	// 	- IMAGE: resets only the image.
	//
	// example:
	//
	// ALL
	ResetScope *string `json:"ResetScope,omitempty" xml:"ResetScope,omitempty"`
	// The disk reset type.
	//
	// Valid values:
	//
	// 	- 0: does not reset disks.
	//
	// 	- 1: resets only the system disk.
	//
	// 	- 2: resets only the user disk.
	//
	// 	- 3: resets the system disk and the user disk.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	ResetType *string `json:"ResetType,omitempty" xml:"ResetType,omitempty"`
}

func (s ResetDesktopsRequest) String() string {
	return dara.Prettify(s)
}

func (s ResetDesktopsRequest) GoString() string {
	return s.String()
}

func (s *ResetDesktopsRequest) GetDesktopGroupId() *string {
	return s.DesktopGroupId
}

func (s *ResetDesktopsRequest) GetDesktopGroupIds() []*string {
	return s.DesktopGroupIds
}

func (s *ResetDesktopsRequest) GetDesktopId() []*string {
	return s.DesktopId
}

func (s *ResetDesktopsRequest) GetImageId() *string {
	return s.ImageId
}

func (s *ResetDesktopsRequest) GetPayType() *string {
	return s.PayType
}

func (s *ResetDesktopsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ResetDesktopsRequest) GetResetScope() *string {
	return s.ResetScope
}

func (s *ResetDesktopsRequest) GetResetType() *string {
	return s.ResetType
}

func (s *ResetDesktopsRequest) SetDesktopGroupId(v string) *ResetDesktopsRequest {
	s.DesktopGroupId = &v
	return s
}

func (s *ResetDesktopsRequest) SetDesktopGroupIds(v []*string) *ResetDesktopsRequest {
	s.DesktopGroupIds = v
	return s
}

func (s *ResetDesktopsRequest) SetDesktopId(v []*string) *ResetDesktopsRequest {
	s.DesktopId = v
	return s
}

func (s *ResetDesktopsRequest) SetImageId(v string) *ResetDesktopsRequest {
	s.ImageId = &v
	return s
}

func (s *ResetDesktopsRequest) SetPayType(v string) *ResetDesktopsRequest {
	s.PayType = &v
	return s
}

func (s *ResetDesktopsRequest) SetRegionId(v string) *ResetDesktopsRequest {
	s.RegionId = &v
	return s
}

func (s *ResetDesktopsRequest) SetResetScope(v string) *ResetDesktopsRequest {
	s.ResetScope = &v
	return s
}

func (s *ResetDesktopsRequest) SetResetType(v string) *ResetDesktopsRequest {
	s.ResetType = &v
	return s
}

func (s *ResetDesktopsRequest) Validate() error {
	return dara.Validate(s)
}
