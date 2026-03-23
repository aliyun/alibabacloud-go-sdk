// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyRCDiskSpecRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoPay(v bool) *ModifyRCDiskSpecRequest
	GetAutoPay() *bool
	SetDiskCategory(v string) *ModifyRCDiskSpecRequest
	GetDiskCategory() *string
	SetDiskId(v string) *ModifyRCDiskSpecRequest
	GetDiskId() *string
	SetDryRun(v bool) *ModifyRCDiskSpecRequest
	GetDryRun() *bool
	SetPerformanceLevel(v string) *ModifyRCDiskSpecRequest
	GetPerformanceLevel() *string
	SetRegionId(v string) *ModifyRCDiskSpecRequest
	GetRegionId() *string
}

type ModifyRCDiskSpecRequest struct {
	AutoPay          *bool   `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	DiskCategory     *string `json:"DiskCategory,omitempty" xml:"DiskCategory,omitempty"`
	DiskId           *string `json:"DiskId,omitempty" xml:"DiskId,omitempty"`
	DryRun           *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	PerformanceLevel *string `json:"PerformanceLevel,omitempty" xml:"PerformanceLevel,omitempty"`
	RegionId         *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ModifyRCDiskSpecRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyRCDiskSpecRequest) GoString() string {
	return s.String()
}

func (s *ModifyRCDiskSpecRequest) GetAutoPay() *bool {
	return s.AutoPay
}

func (s *ModifyRCDiskSpecRequest) GetDiskCategory() *string {
	return s.DiskCategory
}

func (s *ModifyRCDiskSpecRequest) GetDiskId() *string {
	return s.DiskId
}

func (s *ModifyRCDiskSpecRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *ModifyRCDiskSpecRequest) GetPerformanceLevel() *string {
	return s.PerformanceLevel
}

func (s *ModifyRCDiskSpecRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyRCDiskSpecRequest) SetAutoPay(v bool) *ModifyRCDiskSpecRequest {
	s.AutoPay = &v
	return s
}

func (s *ModifyRCDiskSpecRequest) SetDiskCategory(v string) *ModifyRCDiskSpecRequest {
	s.DiskCategory = &v
	return s
}

func (s *ModifyRCDiskSpecRequest) SetDiskId(v string) *ModifyRCDiskSpecRequest {
	s.DiskId = &v
	return s
}

func (s *ModifyRCDiskSpecRequest) SetDryRun(v bool) *ModifyRCDiskSpecRequest {
	s.DryRun = &v
	return s
}

func (s *ModifyRCDiskSpecRequest) SetPerformanceLevel(v string) *ModifyRCDiskSpecRequest {
	s.PerformanceLevel = &v
	return s
}

func (s *ModifyRCDiskSpecRequest) SetRegionId(v string) *ModifyRCDiskSpecRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyRCDiskSpecRequest) Validate() error {
	return dara.Validate(s)
}
