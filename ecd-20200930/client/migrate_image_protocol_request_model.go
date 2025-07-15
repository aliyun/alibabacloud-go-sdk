// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMigrateImageProtocolRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageId(v []*string) *MigrateImageProtocolRequest
	GetImageId() []*string
	SetRegionId(v string) *MigrateImageProtocolRequest
	GetRegionId() *string
	SetTargetProtocolType(v string) *MigrateImageProtocolRequest
	GetTargetProtocolType() *string
}

type MigrateImageProtocolRequest struct {
	// The image IDs.
	//
	// This parameter is required.
	ImageId []*string `json:"ImageId,omitempty" xml:"ImageId,omitempty" type:"Repeated"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/196646.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The protocol to which you want to update the image protocols. Set the value to ASP.
	//
	// example:
	//
	// ASP
	TargetProtocolType *string `json:"TargetProtocolType,omitempty" xml:"TargetProtocolType,omitempty"`
}

func (s MigrateImageProtocolRequest) String() string {
	return dara.Prettify(s)
}

func (s MigrateImageProtocolRequest) GoString() string {
	return s.String()
}

func (s *MigrateImageProtocolRequest) GetImageId() []*string {
	return s.ImageId
}

func (s *MigrateImageProtocolRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *MigrateImageProtocolRequest) GetTargetProtocolType() *string {
	return s.TargetProtocolType
}

func (s *MigrateImageProtocolRequest) SetImageId(v []*string) *MigrateImageProtocolRequest {
	s.ImageId = v
	return s
}

func (s *MigrateImageProtocolRequest) SetRegionId(v string) *MigrateImageProtocolRequest {
	s.RegionId = &v
	return s
}

func (s *MigrateImageProtocolRequest) SetTargetProtocolType(v string) *MigrateImageProtocolRequest {
	s.TargetProtocolType = &v
	return s
}

func (s *MigrateImageProtocolRequest) Validate() error {
	return dara.Validate(s)
}
