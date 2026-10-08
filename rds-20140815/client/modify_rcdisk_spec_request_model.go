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
	// - **false**: Only an order is generated. No payment is made.
	//
	// > If your payment method has an insufficient balance, set AutoPay to false. An unpaid order is generated. You can log on to the ApsaraDB RDS console to complete the payment.
	//
	// >
	//
	// example:
	//
	// true
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// The type of the cloud disk. Valid values:
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
	// The cloud disk ID.
	//
	// example:
	//
	// rcd-wz9f3peueu5npsl****
	DiskId *string `json:"DiskId,omitempty" xml:"DiskId,omitempty"`
	// Specifies whether to perform a dry run for this operation. Valid values:
	//
	// 	- **true**: A dry run is performed without executing the change. The check items include request parameters, request format, business limits, and inventory.
	//
	// 	- **false*	- (default): A normal request is sent. After the check is passed, the change is directly executed.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The performance level (PL) of the ESSD cloud disk. Valid values:
	//
	// - **PL1*	- (default): A maximum of 50,000 random read/write IOPS per disk.
	//
	// - **PL2**: A maximum of 100,000 random read/write IOPS per disk.
	//
	// - **PL3**: A maximum of 1,000,000 random read/write IOPS per disk.
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
