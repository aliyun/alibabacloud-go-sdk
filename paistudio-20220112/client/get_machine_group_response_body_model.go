// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMachineGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCount(v int64) *GetMachineGroupResponseBody
	GetCount() *int64
	SetDefaultDriver(v string) *GetMachineGroupResponseBody
	GetDefaultDriver() *string
	SetDuration(v string) *GetMachineGroupResponseBody
	GetDuration() *string
	SetEcsType(v string) *GetMachineGroupResponseBody
	GetEcsType() *string
	SetGmtCreated(v string) *GetMachineGroupResponseBody
	GetGmtCreated() *string
	SetGmtExpired(v string) *GetMachineGroupResponseBody
	GetGmtExpired() *string
	SetGmtModified(v string) *GetMachineGroupResponseBody
	GetGmtModified() *string
	SetGmtStarted(v string) *GetMachineGroupResponseBody
	GetGmtStarted() *string
	SetMachineGroupID(v string) *GetMachineGroupResponseBody
	GetMachineGroupID() *string
	SetOrderID(v string) *GetMachineGroupResponseBody
	GetOrderID() *string
	SetOrderInstanceId(v string) *GetMachineGroupResponseBody
	GetOrderInstanceId() *string
	SetPAIResourceID(v string) *GetMachineGroupResponseBody
	GetPAIResourceID() *string
	SetPayType(v string) *GetMachineGroupResponseBody
	GetPayType() *string
	SetPricingCycle(v string) *GetMachineGroupResponseBody
	GetPricingCycle() *string
	SetRegionID(v string) *GetMachineGroupResponseBody
	GetRegionID() *string
	SetRequestId(v string) *GetMachineGroupResponseBody
	GetRequestId() *string
	SetStatus(v string) *GetMachineGroupResponseBody
	GetStatus() *string
	SetSupportedDrivers(v []*string) *GetMachineGroupResponseBody
	GetSupportedDrivers() []*string
}

type GetMachineGroupResponseBody struct {
	// The quantity.
	//
	// example:
	//
	// 10
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The default driver version.
	//
	// example:
	//
	// 535
	DefaultDriver *string `json:"DefaultDriver,omitempty" xml:"DefaultDriver,omitempty"`
	// The purchase duration.
	//
	// example:
	//
	// 2
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The ECS instance type.
	//
	// example:
	//
	// ecs.c6.large
	EcsType *string `json:"EcsType,omitempty" xml:"EcsType,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2026-08-10T03:17:31Z
	GmtCreated *string `json:"GmtCreated,omitempty" xml:"GmtCreated,omitempty"`
	// The expiration time.
	//
	// example:
	//
	// 2023-08-22T00:00:00Z
	GmtExpired *string `json:"GmtExpired,omitempty" xml:"GmtExpired,omitempty"`
	// The modification time.
	//
	// example:
	//
	// 2023-06-22T00:00:00Z
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The start time.
	//
	// example:
	//
	// 2023-06-22T00:00:00Z
	GmtStarted *string `json:"GmtStarted,omitempty" xml:"GmtStarted,omitempty"`
	// The machine group ID.
	//
	// example:
	//
	// mg1
	MachineGroupID *string `json:"MachineGroupID,omitempty" xml:"MachineGroupID,omitempty"`
	// The order ID.
	OrderID *string `json:"OrderID,omitempty" xml:"OrderID,omitempty"`
	// The order instance ID.
	//
	// example:
	//
	// 236553689400333
	OrderInstanceId *string `json:"OrderInstanceId,omitempty" xml:"OrderInstanceId,omitempty"`
	// The PAI resource ID.
	PAIResourceID *string `json:"PAIResourceID,omitempty" xml:"PAIResourceID,omitempty"`
	// The billing type.
	//
	// example:
	//
	// PrePay
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// The pricing cycle.
	//
	// example:
	//
	// Monthly
	PricingCycle *string `json:"PricingCycle,omitempty" xml:"PricingCycle,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionID *string `json:"RegionID,omitempty" xml:"RegionID,omitempty"`
	// The request ID.
	//
	// example:
	//
	// F2D0392B-D749-5C48-A98A-3FAE5C9444A6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status.
	//
	// example:
	//
	// Success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The supported driver versions.
	SupportedDrivers []*string `json:"SupportedDrivers,omitempty" xml:"SupportedDrivers,omitempty" type:"Repeated"`
}

func (s GetMachineGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetMachineGroupResponseBody) GoString() string {
	return s.String()
}

func (s *GetMachineGroupResponseBody) GetCount() *int64 {
	return s.Count
}

func (s *GetMachineGroupResponseBody) GetDefaultDriver() *string {
	return s.DefaultDriver
}

func (s *GetMachineGroupResponseBody) GetDuration() *string {
	return s.Duration
}

func (s *GetMachineGroupResponseBody) GetEcsType() *string {
	return s.EcsType
}

func (s *GetMachineGroupResponseBody) GetGmtCreated() *string {
	return s.GmtCreated
}

func (s *GetMachineGroupResponseBody) GetGmtExpired() *string {
	return s.GmtExpired
}

func (s *GetMachineGroupResponseBody) GetGmtModified() *string {
	return s.GmtModified
}

func (s *GetMachineGroupResponseBody) GetGmtStarted() *string {
	return s.GmtStarted
}

func (s *GetMachineGroupResponseBody) GetMachineGroupID() *string {
	return s.MachineGroupID
}

func (s *GetMachineGroupResponseBody) GetOrderID() *string {
	return s.OrderID
}

func (s *GetMachineGroupResponseBody) GetOrderInstanceId() *string {
	return s.OrderInstanceId
}

func (s *GetMachineGroupResponseBody) GetPAIResourceID() *string {
	return s.PAIResourceID
}

func (s *GetMachineGroupResponseBody) GetPayType() *string {
	return s.PayType
}

func (s *GetMachineGroupResponseBody) GetPricingCycle() *string {
	return s.PricingCycle
}

func (s *GetMachineGroupResponseBody) GetRegionID() *string {
	return s.RegionID
}

func (s *GetMachineGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetMachineGroupResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetMachineGroupResponseBody) GetSupportedDrivers() []*string {
	return s.SupportedDrivers
}

func (s *GetMachineGroupResponseBody) SetCount(v int64) *GetMachineGroupResponseBody {
	s.Count = &v
	return s
}

func (s *GetMachineGroupResponseBody) SetDefaultDriver(v string) *GetMachineGroupResponseBody {
	s.DefaultDriver = &v
	return s
}

func (s *GetMachineGroupResponseBody) SetDuration(v string) *GetMachineGroupResponseBody {
	s.Duration = &v
	return s
}

func (s *GetMachineGroupResponseBody) SetEcsType(v string) *GetMachineGroupResponseBody {
	s.EcsType = &v
	return s
}

func (s *GetMachineGroupResponseBody) SetGmtCreated(v string) *GetMachineGroupResponseBody {
	s.GmtCreated = &v
	return s
}

func (s *GetMachineGroupResponseBody) SetGmtExpired(v string) *GetMachineGroupResponseBody {
	s.GmtExpired = &v
	return s
}

func (s *GetMachineGroupResponseBody) SetGmtModified(v string) *GetMachineGroupResponseBody {
	s.GmtModified = &v
	return s
}

func (s *GetMachineGroupResponseBody) SetGmtStarted(v string) *GetMachineGroupResponseBody {
	s.GmtStarted = &v
	return s
}

func (s *GetMachineGroupResponseBody) SetMachineGroupID(v string) *GetMachineGroupResponseBody {
	s.MachineGroupID = &v
	return s
}

func (s *GetMachineGroupResponseBody) SetOrderID(v string) *GetMachineGroupResponseBody {
	s.OrderID = &v
	return s
}

func (s *GetMachineGroupResponseBody) SetOrderInstanceId(v string) *GetMachineGroupResponseBody {
	s.OrderInstanceId = &v
	return s
}

func (s *GetMachineGroupResponseBody) SetPAIResourceID(v string) *GetMachineGroupResponseBody {
	s.PAIResourceID = &v
	return s
}

func (s *GetMachineGroupResponseBody) SetPayType(v string) *GetMachineGroupResponseBody {
	s.PayType = &v
	return s
}

func (s *GetMachineGroupResponseBody) SetPricingCycle(v string) *GetMachineGroupResponseBody {
	s.PricingCycle = &v
	return s
}

func (s *GetMachineGroupResponseBody) SetRegionID(v string) *GetMachineGroupResponseBody {
	s.RegionID = &v
	return s
}

func (s *GetMachineGroupResponseBody) SetRequestId(v string) *GetMachineGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMachineGroupResponseBody) SetStatus(v string) *GetMachineGroupResponseBody {
	s.Status = &v
	return s
}

func (s *GetMachineGroupResponseBody) SetSupportedDrivers(v []*string) *GetMachineGroupResponseBody {
	s.SupportedDrivers = v
	return s
}

func (s *GetMachineGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
