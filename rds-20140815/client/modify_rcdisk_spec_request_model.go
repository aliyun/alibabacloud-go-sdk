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
	// Specifies whether to enable automatic payment. Valid values:
	//
	// 	- **true (default)**: automatically completes the payment. Make sure that your account balance is sufficient.
	//
	// 	- **false**: does not automatically complete the payment. An unpaid order is generated.
	//
	// >  If your account balance is insufficient, you can set the AutoPay parameter to false. In this case, an unpaid order is generated. You can complete the payment in the Expenses and Costs console.
	//
	// example:
	//
	// true
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// The new disk type. Valid values:
	//
	// 	- **cloud_essd**: ESSD.
	//
	// 	- **cloud_auto**: ESSD AutoPL disk
	//
	// This parameter is empty by default.
	//
	// example:
	//
	// cloud_essd
	DiskCategory *string `json:"DiskCategory,omitempty" xml:"DiskCategory,omitempty"`
	// The cloud disk ID.
	//
	// example:
	//
	// rcd-wz9f3peueu5npsl****
	DiskId *string `json:"DiskId,omitempty" xml:"DiskId,omitempty"`
	// Specifies whether to perform a dry run. Valid values: Valid values:
	//
	// 	- **true**: performs a dry run and does not perform the actual request. The system checks the request parameters, request syntax, limits, and available resources.
	//
	// 	- **false*	- (default): performs a dry run and performs the actual request. If the request passes the dry run, the operation is performed.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The PL of the disk. Valid values:
	//
	// 	- **PL1*	- (default): A single ESSD can deliver up to 50,000 random read/write IOPS.
	//
	// 	- **PL2**: A single ESSD delivers up to 100,000 random read/write IOPS.
	//
	// 	- **PL3**: A single ESSD delivers up to 1,000,000 random read/write IOPS.
	//
	// example:
	//
	// PL2
	PerformanceLevel *string `json:"PerformanceLevel,omitempty" xml:"PerformanceLevel,omitempty"`
	// The ID of the region in which the instance resides.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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
