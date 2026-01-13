// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyRCDiskAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBurstingEnabled(v bool) *ModifyRCDiskAttributeRequest
	GetBurstingEnabled() *bool
	SetDeleteWithInstance(v bool) *ModifyRCDiskAttributeRequest
	GetDeleteWithInstance() *bool
	SetDescription(v string) *ModifyRCDiskAttributeRequest
	GetDescription() *string
	SetDiskId(v string) *ModifyRCDiskAttributeRequest
	GetDiskId() *string
	SetDiskName(v string) *ModifyRCDiskAttributeRequest
	GetDiskName() *string
	SetRegionId(v string) *ModifyRCDiskAttributeRequest
	GetRegionId() *string
}

type ModifyRCDiskAttributeRequest struct {
	// example:
	//
	// false
	BurstingEnabled *bool `json:"BurstingEnabled,omitempty" xml:"BurstingEnabled,omitempty"`
	// example:
	//
	// false
	DeleteWithInstance *bool `json:"DeleteWithInstance,omitempty" xml:"DeleteWithInstance,omitempty"`
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// rcd-wz9c8isqly8637zw****
	DiskId *string `json:"DiskId,omitempty" xml:"DiskId,omitempty"`
	// example:
	//
	// testDisk
	DiskName *string `json:"DiskName,omitempty" xml:"DiskName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ModifyRCDiskAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyRCDiskAttributeRequest) GoString() string {
	return s.String()
}

func (s *ModifyRCDiskAttributeRequest) GetBurstingEnabled() *bool {
	return s.BurstingEnabled
}

func (s *ModifyRCDiskAttributeRequest) GetDeleteWithInstance() *bool {
	return s.DeleteWithInstance
}

func (s *ModifyRCDiskAttributeRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifyRCDiskAttributeRequest) GetDiskId() *string {
	return s.DiskId
}

func (s *ModifyRCDiskAttributeRequest) GetDiskName() *string {
	return s.DiskName
}

func (s *ModifyRCDiskAttributeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyRCDiskAttributeRequest) SetBurstingEnabled(v bool) *ModifyRCDiskAttributeRequest {
	s.BurstingEnabled = &v
	return s
}

func (s *ModifyRCDiskAttributeRequest) SetDeleteWithInstance(v bool) *ModifyRCDiskAttributeRequest {
	s.DeleteWithInstance = &v
	return s
}

func (s *ModifyRCDiskAttributeRequest) SetDescription(v string) *ModifyRCDiskAttributeRequest {
	s.Description = &v
	return s
}

func (s *ModifyRCDiskAttributeRequest) SetDiskId(v string) *ModifyRCDiskAttributeRequest {
	s.DiskId = &v
	return s
}

func (s *ModifyRCDiskAttributeRequest) SetDiskName(v string) *ModifyRCDiskAttributeRequest {
	s.DiskName = &v
	return s
}

func (s *ModifyRCDiskAttributeRequest) SetRegionId(v string) *ModifyRCDiskAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyRCDiskAttributeRequest) Validate() error {
	return dara.Validate(s)
}
