// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCtripOrderPushRequest interface {
	dara.Model
	String() string
	GoString() string
	SetJourneyId(v string) *CtripOrderPushRequest
	GetJourneyId() *string
	SetScheduleItem(v string) *CtripOrderPushRequest
	GetScheduleItem() *string
	SetUid(v string) *CtripOrderPushRequest
	GetUid() *string
}

type CtripOrderPushRequest struct {
	JourneyId *string `json:"JourneyId,omitempty" xml:"JourneyId,omitempty"`
	// example:
	//
	// "scheduleItem"
	ScheduleItem *string `json:"ScheduleItem,omitempty" xml:"ScheduleItem,omitempty"`
	// example:
	//
	// "123456789"
	Uid *string `json:"Uid,omitempty" xml:"Uid,omitempty"`
}

func (s CtripOrderPushRequest) String() string {
	return dara.Prettify(s)
}

func (s CtripOrderPushRequest) GoString() string {
	return s.String()
}

func (s *CtripOrderPushRequest) GetJourneyId() *string {
	return s.JourneyId
}

func (s *CtripOrderPushRequest) GetScheduleItem() *string {
	return s.ScheduleItem
}

func (s *CtripOrderPushRequest) GetUid() *string {
	return s.Uid
}

func (s *CtripOrderPushRequest) SetJourneyId(v string) *CtripOrderPushRequest {
	s.JourneyId = &v
	return s
}

func (s *CtripOrderPushRequest) SetScheduleItem(v string) *CtripOrderPushRequest {
	s.ScheduleItem = &v
	return s
}

func (s *CtripOrderPushRequest) SetUid(v string) *CtripOrderPushRequest {
	s.Uid = &v
	return s
}

func (s *CtripOrderPushRequest) Validate() error {
	return dara.Validate(s)
}
