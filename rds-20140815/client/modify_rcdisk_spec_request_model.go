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
	// - **true*	- (default): Automatic payment is enabled. Make sure that your account balance is sufficient.
	//
	// - **false**: Only an order is generated, and no payment is made.
	//
	// > If the balance of your payment method is insufficient, you can change the Settings of the parameter AutoPay to false. An unpaid order is generated. You can logon to the ApsaraDB RDS console to complete the payment.
	//
	// >
	//
	// example:
	//
	// true
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// The new category of the cloud disk. Valid values:
	//
	// - **cloud_essd*	- (default): ESSD cloud disk.
	//
	// - **cloud_auto**: ESSD AutoPL cloud disk.
	//
	// - **cloud_ssd**: standard SSD.
	//
	// example:
	//
	// cloud_essd
	DiskCategory *string `json:"DiskCategory,omitempty" xml:"DiskCategory,omitempty"`
	// The ID of the cloud disk.
	//
	// example:
	//
	// rcd-wz9f3peueu5npsl****
	DiskId *string `json:"DiskId,omitempty" xml:"DiskId,omitempty"`
	// Specifies whether to perform a dry run for this operation. Valid values:
	//
	// 	- **true**: A dry run is performed, and the modification is not executed. The check items include request parameters, request format, business limits, and inventory.
	//
	// 	- **false*	- (default): A normal request is sent, and the modification is executed after the check is passed.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The new performance level (PL) of the ESSD cloud disk. Valid values:
	//
	// - **PL1*	- (default): A single disk delivers up to 50,000 random read/write IOPS.
	//
	// - **PL2**: A single disk delivers up to 100,000 random read/write IOPS.
	//
	// - **PL3**: A single disk delivers up to 1,000,000 random read/write IOPS.
	//
	// example:
	//
	// PL2
	PerformanceLevel *string `json:"PerformanceLevel,omitempty" xml:"PerformanceLevel,omitempty"`
	// The region ID of the instance.
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
