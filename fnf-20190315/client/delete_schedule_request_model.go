// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteScheduleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFlowName(v string) *DeleteScheduleRequest
	GetFlowName() *string
	SetScheduleName(v string) *DeleteScheduleRequest
	GetScheduleName() *string
}

type DeleteScheduleRequest struct {
	// The name of the workflow with which the scheduling task that you want to delete is associated.
	//
	// This parameter is required.
	//
	// example:
	//
	// testFlowName
	FlowName *string `json:"FlowName,omitempty" xml:"FlowName,omitempty"`
	// The name of the scheduling task that you want to delete.
	//
	// This parameter is required.
	//
	// example:
	//
	// testScheduleName
	ScheduleName *string `json:"ScheduleName,omitempty" xml:"ScheduleName,omitempty"`
}

func (s DeleteScheduleRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteScheduleRequest) GoString() string {
	return s.String()
}

func (s *DeleteScheduleRequest) GetFlowName() *string {
	return s.FlowName
}

func (s *DeleteScheduleRequest) GetScheduleName() *string {
	return s.ScheduleName
}

func (s *DeleteScheduleRequest) SetFlowName(v string) *DeleteScheduleRequest {
	s.FlowName = &v
	return s
}

func (s *DeleteScheduleRequest) SetScheduleName(v string) *DeleteScheduleRequest {
	s.ScheduleName = &v
	return s
}

func (s *DeleteScheduleRequest) Validate() error {
	return dara.Validate(s)
}
