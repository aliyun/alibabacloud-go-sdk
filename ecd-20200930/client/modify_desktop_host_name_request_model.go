// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDesktopHostNameRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDesktopId(v string) *ModifyDesktopHostNameRequest
	GetDesktopId() *string
	SetDesktopIds(v []*string) *ModifyDesktopHostNameRequest
	GetDesktopIds() []*string
	SetNewHostName(v string) *ModifyDesktopHostNameRequest
	GetNewHostName() *string
	SetRegionId(v string) *ModifyDesktopHostNameRequest
	GetRegionId() *string
}

type ModifyDesktopHostNameRequest struct {
	// The ID of the cloud desktop.
	//
	// example:
	//
	// ecd-gx2x1dhsmucyy****
	DesktopId  *string   `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
	DesktopIds []*string `json:"DesktopIds,omitempty" xml:"DesktopIds,omitempty" type:"Repeated"`
	// The new hostname of the cloud desktop. The hostname must meet the following requirements:
	//
	// - The hostname must be 2 to 15 characters in length.
	//
	// - The hostname can contain letters, digits, and hyphens (-). It cannot start or end with a hyphen, contain consecutive hyphens, or consist of only digits.
	//
	// This parameter is required.
	//
	// example:
	//
	// NewName
	NewHostName *string `json:"NewHostName,omitempty" xml:"NewHostName,omitempty"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/196646.html) operation to query the available regions for Elastic Desktop Service.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ModifyDesktopHostNameRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDesktopHostNameRequest) GoString() string {
	return s.String()
}

func (s *ModifyDesktopHostNameRequest) GetDesktopId() *string {
	return s.DesktopId
}

func (s *ModifyDesktopHostNameRequest) GetDesktopIds() []*string {
	return s.DesktopIds
}

func (s *ModifyDesktopHostNameRequest) GetNewHostName() *string {
	return s.NewHostName
}

func (s *ModifyDesktopHostNameRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyDesktopHostNameRequest) SetDesktopId(v string) *ModifyDesktopHostNameRequest {
	s.DesktopId = &v
	return s
}

func (s *ModifyDesktopHostNameRequest) SetDesktopIds(v []*string) *ModifyDesktopHostNameRequest {
	s.DesktopIds = v
	return s
}

func (s *ModifyDesktopHostNameRequest) SetNewHostName(v string) *ModifyDesktopHostNameRequest {
	s.NewHostName = &v
	return s
}

func (s *ModifyDesktopHostNameRequest) SetRegionId(v string) *ModifyDesktopHostNameRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyDesktopHostNameRequest) Validate() error {
	return dara.Validate(s)
}
