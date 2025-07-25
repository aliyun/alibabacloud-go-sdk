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
	Count            *int64    `json:"Count,omitempty" xml:"Count,omitempty"`
	DefaultDriver    *string   `json:"DefaultDriver,omitempty" xml:"DefaultDriver,omitempty"`
	Duration         *string   `json:"Duration,omitempty" xml:"Duration,omitempty"`
	EcsType          *string   `json:"EcsType,omitempty" xml:"EcsType,omitempty"`
	GmtCreated       *string   `json:"GmtCreated,omitempty" xml:"GmtCreated,omitempty"`
	GmtExpired       *string   `json:"GmtExpired,omitempty" xml:"GmtExpired,omitempty"`
	GmtModified      *string   `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	GmtStarted       *string   `json:"GmtStarted,omitempty" xml:"GmtStarted,omitempty"`
	MachineGroupID   *string   `json:"MachineGroupID,omitempty" xml:"MachineGroupID,omitempty"`
	OrderID          *string   `json:"OrderID,omitempty" xml:"OrderID,omitempty"`
	OrderInstanceId  *string   `json:"OrderInstanceId,omitempty" xml:"OrderInstanceId,omitempty"`
	PAIResourceID    *string   `json:"PAIResourceID,omitempty" xml:"PAIResourceID,omitempty"`
	PayType          *string   `json:"PayType,omitempty" xml:"PayType,omitempty"`
	PricingCycle     *string   `json:"PricingCycle,omitempty" xml:"PricingCycle,omitempty"`
	RegionID         *string   `json:"RegionID,omitempty" xml:"RegionID,omitempty"`
	RequestId        *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Status           *string   `json:"Status,omitempty" xml:"Status,omitempty"`
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
