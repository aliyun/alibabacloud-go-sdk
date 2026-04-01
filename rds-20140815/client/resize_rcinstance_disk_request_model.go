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
	// Specifies whether to enable the automatic payment feature for the instance. Valid values:
	//
	// 	- **true*	- (default): enables the feature. Make sure that your account balance is sufficient.
	//
	// 	- **false**: disables the feature. An unpaid order is generated.
	//
	// >  If your account balance is insufficient, you can set AutoPay to false. In this case, an unpaid order is generated. You can complete the payment in the Expenses and Costs console.
	//
	// example:
	//
	// false
	AutoPay *bool   `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	DiskId  *string `json:"DiskId,omitempty" xml:"DiskId,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// 	- **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, service limits, and insufficient inventory errors.
	//
	// 	- **false**: performs a dry run and performs the actual request. If the request passes the dry run, the instance is created.
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
	// The new disk size. Unit: GiB.
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
	// The method that you want to use to resize the disk. Valid values:
	//
	// 	- **offline*	- (default): resizes disks offline. After you resize a disk offline, you must restart the instance for the resizing operation to take effect.
	//
	// 	- **online**: resizes disks online. After you resize a disk online, the resizing operation takes effect immediately and you do not need to restart the instance.
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
