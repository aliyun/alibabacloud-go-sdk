// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMachineGroup interface {
	dara.Model
	String() string
	GoString() string
	SetCreatorID(v string) *MachineGroup
	GetCreatorID() *string
	SetDefaultDriver(v string) *MachineGroup
	GetDefaultDriver() *string
	SetEcsCount(v int64) *MachineGroup
	GetEcsCount() *int64
	SetEcsSpec(v string) *MachineGroup
	GetEcsSpec() *string
	SetGmtCreatedTime(v string) *MachineGroup
	GetGmtCreatedTime() *string
	SetGmtExpiredTime(v string) *MachineGroup
	GetGmtExpiredTime() *string
	SetGmtModifiedTime(v string) *MachineGroup
	GetGmtModifiedTime() *string
	SetGmtStartedTime(v string) *MachineGroup
	GetGmtStartedTime() *string
	SetMachineGroupID(v string) *MachineGroup
	GetMachineGroupID() *string
	SetOrderInstanceId(v string) *MachineGroup
	GetOrderInstanceId() *string
	SetPaymentDuration(v string) *MachineGroup
	GetPaymentDuration() *string
	SetPaymentDurationUnit(v string) *MachineGroup
	GetPaymentDurationUnit() *string
	SetPaymentType(v string) *MachineGroup
	GetPaymentType() *string
	SetReasonCode(v string) *MachineGroup
	GetReasonCode() *string
	SetReasonMessage(v string) *MachineGroup
	GetReasonMessage() *string
	SetResourceGroupID(v string) *MachineGroup
	GetResourceGroupID() *string
	SetStatus(v string) *MachineGroup
	GetStatus() *string
	SetSupportedDrivers(v []*string) *MachineGroup
	GetSupportedDrivers() []*string
}

type MachineGroup struct {
	CreatorID *string `json:"CreatorID,omitempty" xml:"CreatorID,omitempty"`
	// example:
	//
	// 470.199.02
	DefaultDriver   *string `json:"DefaultDriver,omitempty" xml:"DefaultDriver,omitempty"`
	EcsCount        *int64  `json:"EcsCount,omitempty" xml:"EcsCount,omitempty"`
	EcsSpec         *string `json:"EcsSpec,omitempty" xml:"EcsSpec,omitempty"`
	GmtCreatedTime  *string `json:"GmtCreatedTime,omitempty" xml:"GmtCreatedTime,omitempty"`
	GmtExpiredTime  *string `json:"GmtExpiredTime,omitempty" xml:"GmtExpiredTime,omitempty"`
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	GmtStartedTime  *string `json:"GmtStartedTime,omitempty" xml:"GmtStartedTime,omitempty"`
	// example:
	//
	// mg1
	MachineGroupID      *string   `json:"MachineGroupID,omitempty" xml:"MachineGroupID,omitempty"`
	OrderInstanceId     *string   `json:"OrderInstanceId,omitempty" xml:"OrderInstanceId,omitempty"`
	PaymentDuration     *string   `json:"PaymentDuration,omitempty" xml:"PaymentDuration,omitempty"`
	PaymentDurationUnit *string   `json:"PaymentDurationUnit,omitempty" xml:"PaymentDurationUnit,omitempty"`
	PaymentType         *string   `json:"PaymentType,omitempty" xml:"PaymentType,omitempty"`
	ReasonCode          *string   `json:"ReasonCode,omitempty" xml:"ReasonCode,omitempty"`
	ReasonMessage       *string   `json:"ReasonMessage,omitempty" xml:"ReasonMessage,omitempty"`
	ResourceGroupID     *string   `json:"ResourceGroupID,omitempty" xml:"ResourceGroupID,omitempty"`
	Status              *string   `json:"Status,omitempty" xml:"Status,omitempty"`
	SupportedDrivers    []*string `json:"SupportedDrivers,omitempty" xml:"SupportedDrivers,omitempty" type:"Repeated"`
}

func (s MachineGroup) String() string {
	return dara.Prettify(s)
}

func (s MachineGroup) GoString() string {
	return s.String()
}

func (s *MachineGroup) GetCreatorID() *string {
	return s.CreatorID
}

func (s *MachineGroup) GetDefaultDriver() *string {
	return s.DefaultDriver
}

func (s *MachineGroup) GetEcsCount() *int64 {
	return s.EcsCount
}

func (s *MachineGroup) GetEcsSpec() *string {
	return s.EcsSpec
}

func (s *MachineGroup) GetGmtCreatedTime() *string {
	return s.GmtCreatedTime
}

func (s *MachineGroup) GetGmtExpiredTime() *string {
	return s.GmtExpiredTime
}

func (s *MachineGroup) GetGmtModifiedTime() *string {
	return s.GmtModifiedTime
}

func (s *MachineGroup) GetGmtStartedTime() *string {
	return s.GmtStartedTime
}

func (s *MachineGroup) GetMachineGroupID() *string {
	return s.MachineGroupID
}

func (s *MachineGroup) GetOrderInstanceId() *string {
	return s.OrderInstanceId
}

func (s *MachineGroup) GetPaymentDuration() *string {
	return s.PaymentDuration
}

func (s *MachineGroup) GetPaymentDurationUnit() *string {
	return s.PaymentDurationUnit
}

func (s *MachineGroup) GetPaymentType() *string {
	return s.PaymentType
}

func (s *MachineGroup) GetReasonCode() *string {
	return s.ReasonCode
}

func (s *MachineGroup) GetReasonMessage() *string {
	return s.ReasonMessage
}

func (s *MachineGroup) GetResourceGroupID() *string {
	return s.ResourceGroupID
}

func (s *MachineGroup) GetStatus() *string {
	return s.Status
}

func (s *MachineGroup) GetSupportedDrivers() []*string {
	return s.SupportedDrivers
}

func (s *MachineGroup) SetCreatorID(v string) *MachineGroup {
	s.CreatorID = &v
	return s
}

func (s *MachineGroup) SetDefaultDriver(v string) *MachineGroup {
	s.DefaultDriver = &v
	return s
}

func (s *MachineGroup) SetEcsCount(v int64) *MachineGroup {
	s.EcsCount = &v
	return s
}

func (s *MachineGroup) SetEcsSpec(v string) *MachineGroup {
	s.EcsSpec = &v
	return s
}

func (s *MachineGroup) SetGmtCreatedTime(v string) *MachineGroup {
	s.GmtCreatedTime = &v
	return s
}

func (s *MachineGroup) SetGmtExpiredTime(v string) *MachineGroup {
	s.GmtExpiredTime = &v
	return s
}

func (s *MachineGroup) SetGmtModifiedTime(v string) *MachineGroup {
	s.GmtModifiedTime = &v
	return s
}

func (s *MachineGroup) SetGmtStartedTime(v string) *MachineGroup {
	s.GmtStartedTime = &v
	return s
}

func (s *MachineGroup) SetMachineGroupID(v string) *MachineGroup {
	s.MachineGroupID = &v
	return s
}

func (s *MachineGroup) SetOrderInstanceId(v string) *MachineGroup {
	s.OrderInstanceId = &v
	return s
}

func (s *MachineGroup) SetPaymentDuration(v string) *MachineGroup {
	s.PaymentDuration = &v
	return s
}

func (s *MachineGroup) SetPaymentDurationUnit(v string) *MachineGroup {
	s.PaymentDurationUnit = &v
	return s
}

func (s *MachineGroup) SetPaymentType(v string) *MachineGroup {
	s.PaymentType = &v
	return s
}

func (s *MachineGroup) SetReasonCode(v string) *MachineGroup {
	s.ReasonCode = &v
	return s
}

func (s *MachineGroup) SetReasonMessage(v string) *MachineGroup {
	s.ReasonMessage = &v
	return s
}

func (s *MachineGroup) SetResourceGroupID(v string) *MachineGroup {
	s.ResourceGroupID = &v
	return s
}

func (s *MachineGroup) SetStatus(v string) *MachineGroup {
	s.Status = &v
	return s
}

func (s *MachineGroup) SetSupportedDrivers(v []*string) *MachineGroup {
	s.SupportedDrivers = v
	return s
}

func (s *MachineGroup) Validate() error {
	return dara.Validate(s)
}
