// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iManagedDaOrderVO interface {
	dara.Model
	String() string
	GoString() string
	SetExpireTime(v string) *ManagedDaOrderVO
	GetExpireTime() *string
	SetGmtCreate(v string) *ManagedDaOrderVO
	GetGmtCreate() *string
	SetInstanceId(v string) *ManagedDaOrderVO
	GetInstanceId() *string
	SetOrderId(v int64) *ManagedDaOrderVO
	GetOrderId() *int64
	SetPayNum(v int32) *ManagedDaOrderVO
	GetPayNum() *int32
	SetRegion(v string) *ManagedDaOrderVO
	GetRegion() *string
	SetState(v string) *ManagedDaOrderVO
	GetState() *string
	SetSubscriptionPlan(v string) *ManagedDaOrderVO
	GetSubscriptionPlan() *string
}

type ManagedDaOrderVO struct {
	// Use the UTC time format: yyyy-MM-ddTHH:mm:ss.SSSZ
	ExpireTime *string `json:"expireTime,omitempty" xml:"expireTime,omitempty"`
	// Use the UTC time format: yyyy-MM-ddTHH:mm:ss.SSSZ
	GmtCreate        *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	InstanceId       *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	OrderId          *int64  `json:"orderId,omitempty" xml:"orderId,omitempty"`
	PayNum           *int32  `json:"payNum,omitempty" xml:"payNum,omitempty"`
	Region           *string `json:"region,omitempty" xml:"region,omitempty"`
	State            *string `json:"state,omitempty" xml:"state,omitempty"`
	SubscriptionPlan *string `json:"subscriptionPlan,omitempty" xml:"subscriptionPlan,omitempty"`
}

func (s ManagedDaOrderVO) String() string {
	return dara.Prettify(s)
}

func (s ManagedDaOrderVO) GoString() string {
	return s.String()
}

func (s *ManagedDaOrderVO) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *ManagedDaOrderVO) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ManagedDaOrderVO) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ManagedDaOrderVO) GetOrderId() *int64 {
	return s.OrderId
}

func (s *ManagedDaOrderVO) GetPayNum() *int32 {
	return s.PayNum
}

func (s *ManagedDaOrderVO) GetRegion() *string {
	return s.Region
}

func (s *ManagedDaOrderVO) GetState() *string {
	return s.State
}

func (s *ManagedDaOrderVO) GetSubscriptionPlan() *string {
	return s.SubscriptionPlan
}

func (s *ManagedDaOrderVO) SetExpireTime(v string) *ManagedDaOrderVO {
	s.ExpireTime = &v
	return s
}

func (s *ManagedDaOrderVO) SetGmtCreate(v string) *ManagedDaOrderVO {
	s.GmtCreate = &v
	return s
}

func (s *ManagedDaOrderVO) SetInstanceId(v string) *ManagedDaOrderVO {
	s.InstanceId = &v
	return s
}

func (s *ManagedDaOrderVO) SetOrderId(v int64) *ManagedDaOrderVO {
	s.OrderId = &v
	return s
}

func (s *ManagedDaOrderVO) SetPayNum(v int32) *ManagedDaOrderVO {
	s.PayNum = &v
	return s
}

func (s *ManagedDaOrderVO) SetRegion(v string) *ManagedDaOrderVO {
	s.Region = &v
	return s
}

func (s *ManagedDaOrderVO) SetState(v string) *ManagedDaOrderVO {
	s.State = &v
	return s
}

func (s *ManagedDaOrderVO) SetSubscriptionPlan(v string) *ManagedDaOrderVO {
	s.SubscriptionPlan = &v
	return s
}

func (s *ManagedDaOrderVO) Validate() error {
	return dara.Validate(s)
}
