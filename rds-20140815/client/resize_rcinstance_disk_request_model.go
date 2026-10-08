// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResizeRCInstanceDiskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoPay(v bool) *ResizeRCInstanceDiskRequest
	GetAutoPay() *bool
	SetDiskId(v string) *ResizeRCInstanceDiskRequest
	GetDiskId() *string
	SetDryRun(v bool) *ResizeRCInstanceDiskRequest
	GetDryRun() *bool
	SetInstanceId(v string) *ResizeRCInstanceDiskRequest
	GetInstanceId() *string
	SetNewSize(v int64) *ResizeRCInstanceDiskRequest
	GetNewSize() *int64
	SetRegionId(v string) *ResizeRCInstanceDiskRequest
	GetRegionId() *string
	SetType(v string) *ResizeRCInstanceDiskRequest
	GetType() *string
}

type ResizeRCInstanceDiskRequest struct {
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
	// false
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// The cloud disk ID.
	//
	// example:
	//
	// rcd-x4462840nwinu6rr61m5o
	DiskId *string `json:"DiskId,omitempty" xml:"DiskId,omitempty"`
	// Specifies whether to perform a dry run. Valid values:
	//
	// 	- **true**: performs a dry run without creating the instance. The system checks items such as the request parameters, request format, service limits, and available resources.
	//
	// 	- **false*	- (default): sends the request. If the request passes the check, the instance is created.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// rm-uf62br2491p5l****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The size of the disk after expansion. Unit: GiB.
	//
	// example:
	//
	// 100
	NewSize *int64 `json:"NewSize,omitempty" xml:"NewSize,omitempty"`
	// The region ID of the instance.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The method used to expand the disk. Valid values:
	//
	// - **offline*	- (default): Offline expansion. You must restart the instance for the expansion to take effect.
	//
	// - **online**: Online expansion. The expansion takes effect without restarting the instance.
	//
	// example:
	//
	// online
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ResizeRCInstanceDiskRequest) String() string {
	return dara.Prettify(s)
}

func (s ResizeRCInstanceDiskRequest) GoString() string {
	return s.String()
}

func (s *ResizeRCInstanceDiskRequest) GetAutoPay() *bool {
	return s.AutoPay
}

func (s *ResizeRCInstanceDiskRequest) GetDiskId() *string {
	return s.DiskId
}

func (s *ResizeRCInstanceDiskRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *ResizeRCInstanceDiskRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ResizeRCInstanceDiskRequest) GetNewSize() *int64 {
	return s.NewSize
}

func (s *ResizeRCInstanceDiskRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ResizeRCInstanceDiskRequest) GetType() *string {
	return s.Type
}

func (s *ResizeRCInstanceDiskRequest) SetAutoPay(v bool) *ResizeRCInstanceDiskRequest {
	s.AutoPay = &v
	return s
}

func (s *ResizeRCInstanceDiskRequest) SetDiskId(v string) *ResizeRCInstanceDiskRequest {
	s.DiskId = &v
	return s
}

func (s *ResizeRCInstanceDiskRequest) SetDryRun(v bool) *ResizeRCInstanceDiskRequest {
	s.DryRun = &v
	return s
}

func (s *ResizeRCInstanceDiskRequest) SetInstanceId(v string) *ResizeRCInstanceDiskRequest {
	s.InstanceId = &v
	return s
}

func (s *ResizeRCInstanceDiskRequest) SetNewSize(v int64) *ResizeRCInstanceDiskRequest {
	s.NewSize = &v
	return s
}

func (s *ResizeRCInstanceDiskRequest) SetRegionId(v string) *ResizeRCInstanceDiskRequest {
	s.RegionId = &v
	return s
}

func (s *ResizeRCInstanceDiskRequest) SetType(v string) *ResizeRCInstanceDiskRequest {
	s.Type = &v
	return s
}

func (s *ResizeRCInstanceDiskRequest) Validate() error {
	return dara.Validate(s)
}
