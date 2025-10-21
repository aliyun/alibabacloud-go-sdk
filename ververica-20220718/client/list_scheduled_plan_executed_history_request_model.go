// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListScheduledPlanExecutedHistoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeploymentId(v string) *ListScheduledPlanExecutedHistoryRequest
	GetDeploymentId() *string
	SetOrigin(v string) *ListScheduledPlanExecutedHistoryRequest
	GetOrigin() *string
}

type ListScheduledPlanExecutedHistoryRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 58718c99-3b29-4c5e-93bb-c9fc4ec6****
	DeploymentId *string `json:"deploymentId,omitempty" xml:"deploymentId,omitempty"`
	// example:
	//
	// SCHEDULED_PLAN
	Origin *string `json:"origin,omitempty" xml:"origin,omitempty"`
}

func (s ListScheduledPlanExecutedHistoryRequest) String() string {
	return dara.Prettify(s)
}

func (s ListScheduledPlanExecutedHistoryRequest) GoString() string {
	return s.String()
}

func (s *ListScheduledPlanExecutedHistoryRequest) GetDeploymentId() *string {
	return s.DeploymentId
}

func (s *ListScheduledPlanExecutedHistoryRequest) GetOrigin() *string {
	return s.Origin
}

func (s *ListScheduledPlanExecutedHistoryRequest) SetDeploymentId(v string) *ListScheduledPlanExecutedHistoryRequest {
	s.DeploymentId = &v
	return s
}

func (s *ListScheduledPlanExecutedHistoryRequest) SetOrigin(v string) *ListScheduledPlanExecutedHistoryRequest {
	s.Origin = &v
	return s
}

func (s *ListScheduledPlanExecutedHistoryRequest) Validate() error {
	return dara.Validate(s)
}
