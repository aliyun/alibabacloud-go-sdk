// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangeApplyShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChangePassengerListShrink(v string) *ChangeApplyShrinkRequest
	GetChangePassengerListShrink() *string
	SetChangedJourneysShrink(v string) *ChangeApplyShrinkRequest
	GetChangedJourneysShrink() *string
	SetContactShrink(v string) *ChangeApplyShrinkRequest
	GetContactShrink() *string
	SetOrderNum(v int64) *ChangeApplyShrinkRequest
	GetOrderNum() *int64
	SetRemark(v string) *ChangeApplyShrinkRequest
	GetRemark() *string
	SetType(v int32) *ChangeApplyShrinkRequest
	GetType() *int32
}

type ChangeApplyShrinkRequest struct {
	// This parameter is required.
	ChangePassengerListShrink *string `json:"change_passenger_list,omitempty" xml:"change_passenger_list,omitempty"`
	// This parameter is required.
	ChangedJourneysShrink *string `json:"changed_journeys,omitempty" xml:"changed_journeys,omitempty"`
	// This parameter is required.
	ContactShrink *string `json:"contact,omitempty" xml:"contact,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 4988430***950
	OrderNum *int64 `json:"order_num,omitempty" xml:"order_num,omitempty"`
	// example:
	//
	// remark desc
	Remark *string `json:"remark,omitempty" xml:"remark,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0
	Type *int32 `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ChangeApplyShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ChangeApplyShrinkRequest) GoString() string {
	return s.String()
}

func (s *ChangeApplyShrinkRequest) GetChangePassengerListShrink() *string {
	return s.ChangePassengerListShrink
}

func (s *ChangeApplyShrinkRequest) GetChangedJourneysShrink() *string {
	return s.ChangedJourneysShrink
}

func (s *ChangeApplyShrinkRequest) GetContactShrink() *string {
	return s.ContactShrink
}

func (s *ChangeApplyShrinkRequest) GetOrderNum() *int64 {
	return s.OrderNum
}

func (s *ChangeApplyShrinkRequest) GetRemark() *string {
	return s.Remark
}

func (s *ChangeApplyShrinkRequest) GetType() *int32 {
	return s.Type
}

func (s *ChangeApplyShrinkRequest) SetChangePassengerListShrink(v string) *ChangeApplyShrinkRequest {
	s.ChangePassengerListShrink = &v
	return s
}

func (s *ChangeApplyShrinkRequest) SetChangedJourneysShrink(v string) *ChangeApplyShrinkRequest {
	s.ChangedJourneysShrink = &v
	return s
}

func (s *ChangeApplyShrinkRequest) SetContactShrink(v string) *ChangeApplyShrinkRequest {
	s.ContactShrink = &v
	return s
}

func (s *ChangeApplyShrinkRequest) SetOrderNum(v int64) *ChangeApplyShrinkRequest {
	s.OrderNum = &v
	return s
}

func (s *ChangeApplyShrinkRequest) SetRemark(v string) *ChangeApplyShrinkRequest {
	s.Remark = &v
	return s
}

func (s *ChangeApplyShrinkRequest) SetType(v int32) *ChangeApplyShrinkRequest {
	s.Type = &v
	return s
}

func (s *ChangeApplyShrinkRequest) Validate() error {
	return dara.Validate(s)
}
