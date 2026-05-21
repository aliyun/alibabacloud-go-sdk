// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateOperationEventScheduleTimeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *UpdateOperationEventScheduleTimeRequest
	GetId() *string
	SetScheduleTime(v string) *UpdateOperationEventScheduleTimeRequest
	GetScheduleTime() *string
}

type UpdateOperationEventScheduleTimeRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 2034449120420339713
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2024-08-23 14:55:00
	ScheduleTime *string `json:"scheduleTime,omitempty" xml:"scheduleTime,omitempty"`
}

func (s UpdateOperationEventScheduleTimeRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateOperationEventScheduleTimeRequest) GoString() string {
	return s.String()
}

func (s *UpdateOperationEventScheduleTimeRequest) GetId() *string {
	return s.Id
}

func (s *UpdateOperationEventScheduleTimeRequest) GetScheduleTime() *string {
	return s.ScheduleTime
}

func (s *UpdateOperationEventScheduleTimeRequest) SetId(v string) *UpdateOperationEventScheduleTimeRequest {
	s.Id = &v
	return s
}

func (s *UpdateOperationEventScheduleTimeRequest) SetScheduleTime(v string) *UpdateOperationEventScheduleTimeRequest {
	s.ScheduleTime = &v
	return s
}

func (s *UpdateOperationEventScheduleTimeRequest) Validate() error {
	return dara.Validate(s)
}
