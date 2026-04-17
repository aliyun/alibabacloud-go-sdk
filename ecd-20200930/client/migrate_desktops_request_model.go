// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMigrateDesktopsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDesktopId(v []*string) *MigrateDesktopsRequest
	GetDesktopId() []*string
	SetRegionId(v string) *MigrateDesktopsRequest
	GetRegionId() *string
	SetTargetMemberIp(v string) *MigrateDesktopsRequest
	GetTargetMemberIp() *string
	SetTargetOfficeSiteId(v string) *MigrateDesktopsRequest
	GetTargetOfficeSiteId() *string
	SetTargetSubnetId(v string) *MigrateDesktopsRequest
	GetTargetSubnetId() *string
}

type MigrateDesktopsRequest struct {
	// The IDs of the cloud computers. You can specify 1 to 100 IDs.
	//
	// This parameter is required.
	DesktopId []*string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty" type:"Repeated"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/196646.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	TargetMemberIp *string `json:"TargetMemberIp,omitempty" xml:"TargetMemberIp,omitempty"`
	// The ID of the destination office network.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shenzhen+dir-388505****
	TargetOfficeSiteId *string `json:"TargetOfficeSiteId,omitempty" xml:"TargetOfficeSiteId,omitempty"`
	// > This parameter is for internal use only.
	//
	// example:
	//
	// null
	TargetSubnetId *string `json:"TargetSubnetId,omitempty" xml:"TargetSubnetId,omitempty"`
}

func (s MigrateDesktopsRequest) String() string {
	return dara.Prettify(s)
}

func (s MigrateDesktopsRequest) GoString() string {
	return s.String()
}

func (s *MigrateDesktopsRequest) GetDesktopId() []*string {
	return s.DesktopId
}

func (s *MigrateDesktopsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *MigrateDesktopsRequest) GetTargetMemberIp() *string {
	return s.TargetMemberIp
}

func (s *MigrateDesktopsRequest) GetTargetOfficeSiteId() *string {
	return s.TargetOfficeSiteId
}

func (s *MigrateDesktopsRequest) GetTargetSubnetId() *string {
	return s.TargetSubnetId
}

func (s *MigrateDesktopsRequest) SetDesktopId(v []*string) *MigrateDesktopsRequest {
	s.DesktopId = v
	return s
}

func (s *MigrateDesktopsRequest) SetRegionId(v string) *MigrateDesktopsRequest {
	s.RegionId = &v
	return s
}

func (s *MigrateDesktopsRequest) SetTargetMemberIp(v string) *MigrateDesktopsRequest {
	s.TargetMemberIp = &v
	return s
}

func (s *MigrateDesktopsRequest) SetTargetOfficeSiteId(v string) *MigrateDesktopsRequest {
	s.TargetOfficeSiteId = &v
	return s
}

func (s *MigrateDesktopsRequest) SetTargetSubnetId(v string) *MigrateDesktopsRequest {
	s.TargetSubnetId = &v
	return s
}

func (s *MigrateDesktopsRequest) Validate() error {
	return dara.Validate(s)
}
