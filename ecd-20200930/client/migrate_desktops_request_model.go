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
	SetTargetOfficeSiteId(v string) *MigrateDesktopsRequest
	GetTargetOfficeSiteId() *string
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
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the destination office network.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shenzhen+dir-388505****
	TargetOfficeSiteId *string `json:"TargetOfficeSiteId,omitempty" xml:"TargetOfficeSiteId,omitempty"`
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

func (s *MigrateDesktopsRequest) GetTargetOfficeSiteId() *string {
	return s.TargetOfficeSiteId
}

func (s *MigrateDesktopsRequest) SetDesktopId(v []*string) *MigrateDesktopsRequest {
	s.DesktopId = v
	return s
}

func (s *MigrateDesktopsRequest) SetRegionId(v string) *MigrateDesktopsRequest {
	s.RegionId = &v
	return s
}

func (s *MigrateDesktopsRequest) SetTargetOfficeSiteId(v string) *MigrateDesktopsRequest {
	s.TargetOfficeSiteId = &v
	return s
}

func (s *MigrateDesktopsRequest) Validate() error {
	return dara.Validate(s)
}
