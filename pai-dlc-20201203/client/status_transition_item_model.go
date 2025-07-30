// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStatusTransitionItem interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *StatusTransitionItem
	GetEndTime() *string
	SetReasonCode(v string) *StatusTransitionItem
	GetReasonCode() *string
	SetReasonMessage(v string) *StatusTransitionItem
	GetReasonMessage() *string
	SetStartTime(v string) *StatusTransitionItem
	GetStartTime() *string
	SetStatus(v string) *StatusTransitionItem
	GetStatus() *string
}

type StatusTransitionItem struct {
	EndTime       *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	ReasonCode    *string `json:"ReasonCode,omitempty" xml:"ReasonCode,omitempty"`
	ReasonMessage *string `json:"ReasonMessage,omitempty" xml:"ReasonMessage,omitempty"`
	StartTime     *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Status        *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s StatusTransitionItem) String() string {
	return dara.Prettify(s)
}

func (s StatusTransitionItem) GoString() string {
	return s.String()
}

func (s *StatusTransitionItem) GetEndTime() *string {
	return s.EndTime
}

func (s *StatusTransitionItem) GetReasonCode() *string {
	return s.ReasonCode
}

func (s *StatusTransitionItem) GetReasonMessage() *string {
	return s.ReasonMessage
}

func (s *StatusTransitionItem) GetStartTime() *string {
	return s.StartTime
}

func (s *StatusTransitionItem) GetStatus() *string {
	return s.Status
}

func (s *StatusTransitionItem) SetEndTime(v string) *StatusTransitionItem {
	s.EndTime = &v
	return s
}

func (s *StatusTransitionItem) SetReasonCode(v string) *StatusTransitionItem {
	s.ReasonCode = &v
	return s
}

func (s *StatusTransitionItem) SetReasonMessage(v string) *StatusTransitionItem {
	s.ReasonMessage = &v
	return s
}

func (s *StatusTransitionItem) SetStartTime(v string) *StatusTransitionItem {
	s.StartTime = &v
	return s
}

func (s *StatusTransitionItem) SetStatus(v string) *StatusTransitionItem {
	s.Status = &v
	return s
}

func (s *StatusTransitionItem) Validate() error {
	return dara.Validate(s)
}
