// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iScheduledPlanExecutedStatus interface {
	dara.Model
	String() string
	GoString() string
	SetRestartType(v string) *ScheduledPlanExecutedStatus
	GetRestartType() *string
	SetStatusState(v string) *ScheduledPlanExecutedStatus
	GetStatusState() *string
}

type ScheduledPlanExecutedStatus struct {
	// example:
	//
	// HOT_UPDATE
	RestartType *string `json:"restartType,omitempty" xml:"restartType,omitempty"`
	// example:
	//
	// UPGRADED
	StatusState *string `json:"statusState,omitempty" xml:"statusState,omitempty"`
}

func (s ScheduledPlanExecutedStatus) String() string {
	return dara.Prettify(s)
}

func (s ScheduledPlanExecutedStatus) GoString() string {
	return s.String()
}

func (s *ScheduledPlanExecutedStatus) GetRestartType() *string {
	return s.RestartType
}

func (s *ScheduledPlanExecutedStatus) GetStatusState() *string {
	return s.StatusState
}

func (s *ScheduledPlanExecutedStatus) SetRestartType(v string) *ScheduledPlanExecutedStatus {
	s.RestartType = &v
	return s
}

func (s *ScheduledPlanExecutedStatus) SetStatusState(v string) *ScheduledPlanExecutedStatus {
	s.StatusState = &v
	return s
}

func (s *ScheduledPlanExecutedStatus) Validate() error {
	return dara.Validate(s)
}
