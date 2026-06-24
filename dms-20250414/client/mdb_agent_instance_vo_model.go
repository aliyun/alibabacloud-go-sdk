// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMdbAgentInstanceVo interface {
	dara.Model
	String() string
	GoString() string
	SetAccessUrl(v string) *MdbAgentInstanceVo
	GetAccessUrl() *string
	SetChargeType(v string) *MdbAgentInstanceVo
	GetChargeType() *string
	SetEngineInstanceId(v string) *MdbAgentInstanceVo
	GetEngineInstanceId() *string
	SetEngineRegion(v string) *MdbAgentInstanceVo
	GetEngineRegion() *string
	SetEngineType(v string) *MdbAgentInstanceVo
	GetEngineType() *string
	SetGmtCreate(v string) *MdbAgentInstanceVo
	GetGmtCreate() *string
	SetInstanceId(v string) *MdbAgentInstanceVo
	GetInstanceId() *string
	SetInstanceName(v string) *MdbAgentInstanceVo
	GetInstanceName() *string
	SetLastActiveTime(v string) *MdbAgentInstanceVo
	GetLastActiveTime() *string
	SetLockTime(v string) *MdbAgentInstanceVo
	GetLockTime() *string
	SetOrderId(v string) *MdbAgentInstanceVo
	GetOrderId() *string
	SetPublicDomain(v string) *MdbAgentInstanceVo
	GetPublicDomain() *string
	SetStatus(v int32) *MdbAgentInstanceVo
	GetStatus() *int32
	SetStatusDesc(v string) *MdbAgentInstanceVo
	GetStatusDesc() *string
	SetStatusMessage(v string) *MdbAgentInstanceVo
	GetStatusMessage() *string
}

type MdbAgentInstanceVo struct {
	AccessUrl        *string `json:"AccessUrl,omitempty" xml:"AccessUrl,omitempty"`
	ChargeType       *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	EngineInstanceId *string `json:"EngineInstanceId,omitempty" xml:"EngineInstanceId,omitempty"`
	EngineRegion     *string `json:"EngineRegion,omitempty" xml:"EngineRegion,omitempty"`
	EngineType       *string `json:"EngineType,omitempty" xml:"EngineType,omitempty"`
	GmtCreate        *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	InstanceId       *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceName     *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	LastActiveTime   *string `json:"LastActiveTime,omitempty" xml:"LastActiveTime,omitempty"`
	LockTime         *string `json:"LockTime,omitempty" xml:"LockTime,omitempty"`
	OrderId          *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	PublicDomain     *string `json:"PublicDomain,omitempty" xml:"PublicDomain,omitempty"`
	Status           *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	StatusDesc       *string `json:"StatusDesc,omitempty" xml:"StatusDesc,omitempty"`
	StatusMessage    *string `json:"StatusMessage,omitempty" xml:"StatusMessage,omitempty"`
}

func (s MdbAgentInstanceVo) String() string {
	return dara.Prettify(s)
}

func (s MdbAgentInstanceVo) GoString() string {
	return s.String()
}

func (s *MdbAgentInstanceVo) GetAccessUrl() *string {
	return s.AccessUrl
}

func (s *MdbAgentInstanceVo) GetChargeType() *string {
	return s.ChargeType
}

func (s *MdbAgentInstanceVo) GetEngineInstanceId() *string {
	return s.EngineInstanceId
}

func (s *MdbAgentInstanceVo) GetEngineRegion() *string {
	return s.EngineRegion
}

func (s *MdbAgentInstanceVo) GetEngineType() *string {
	return s.EngineType
}

func (s *MdbAgentInstanceVo) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *MdbAgentInstanceVo) GetInstanceId() *string {
	return s.InstanceId
}

func (s *MdbAgentInstanceVo) GetInstanceName() *string {
	return s.InstanceName
}

func (s *MdbAgentInstanceVo) GetLastActiveTime() *string {
	return s.LastActiveTime
}

func (s *MdbAgentInstanceVo) GetLockTime() *string {
	return s.LockTime
}

func (s *MdbAgentInstanceVo) GetOrderId() *string {
	return s.OrderId
}

func (s *MdbAgentInstanceVo) GetPublicDomain() *string {
	return s.PublicDomain
}

func (s *MdbAgentInstanceVo) GetStatus() *int32 {
	return s.Status
}

func (s *MdbAgentInstanceVo) GetStatusDesc() *string {
	return s.StatusDesc
}

func (s *MdbAgentInstanceVo) GetStatusMessage() *string {
	return s.StatusMessage
}

func (s *MdbAgentInstanceVo) SetAccessUrl(v string) *MdbAgentInstanceVo {
	s.AccessUrl = &v
	return s
}

func (s *MdbAgentInstanceVo) SetChargeType(v string) *MdbAgentInstanceVo {
	s.ChargeType = &v
	return s
}

func (s *MdbAgentInstanceVo) SetEngineInstanceId(v string) *MdbAgentInstanceVo {
	s.EngineInstanceId = &v
	return s
}

func (s *MdbAgentInstanceVo) SetEngineRegion(v string) *MdbAgentInstanceVo {
	s.EngineRegion = &v
	return s
}

func (s *MdbAgentInstanceVo) SetEngineType(v string) *MdbAgentInstanceVo {
	s.EngineType = &v
	return s
}

func (s *MdbAgentInstanceVo) SetGmtCreate(v string) *MdbAgentInstanceVo {
	s.GmtCreate = &v
	return s
}

func (s *MdbAgentInstanceVo) SetInstanceId(v string) *MdbAgentInstanceVo {
	s.InstanceId = &v
	return s
}

func (s *MdbAgentInstanceVo) SetInstanceName(v string) *MdbAgentInstanceVo {
	s.InstanceName = &v
	return s
}

func (s *MdbAgentInstanceVo) SetLastActiveTime(v string) *MdbAgentInstanceVo {
	s.LastActiveTime = &v
	return s
}

func (s *MdbAgentInstanceVo) SetLockTime(v string) *MdbAgentInstanceVo {
	s.LockTime = &v
	return s
}

func (s *MdbAgentInstanceVo) SetOrderId(v string) *MdbAgentInstanceVo {
	s.OrderId = &v
	return s
}

func (s *MdbAgentInstanceVo) SetPublicDomain(v string) *MdbAgentInstanceVo {
	s.PublicDomain = &v
	return s
}

func (s *MdbAgentInstanceVo) SetStatus(v int32) *MdbAgentInstanceVo {
	s.Status = &v
	return s
}

func (s *MdbAgentInstanceVo) SetStatusDesc(v string) *MdbAgentInstanceVo {
	s.StatusDesc = &v
	return s
}

func (s *MdbAgentInstanceVo) SetStatusMessage(v string) *MdbAgentInstanceVo {
	s.StatusMessage = &v
	return s
}

func (s *MdbAgentInstanceVo) Validate() error {
	return dara.Validate(s)
}
