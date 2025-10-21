// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAppliedScheduledPlanRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeploymentId(v string) *GetAppliedScheduledPlanRequest
	GetDeploymentId() *string
}

type GetAppliedScheduledPlanRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 5737ef81-d2f1-49cf-8752-30910809****
	DeploymentId *string `json:"deploymentId,omitempty" xml:"deploymentId,omitempty"`
}

func (s GetAppliedScheduledPlanRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAppliedScheduledPlanRequest) GoString() string {
	return s.String()
}

func (s *GetAppliedScheduledPlanRequest) GetDeploymentId() *string {
	return s.DeploymentId
}

func (s *GetAppliedScheduledPlanRequest) SetDeploymentId(v string) *GetAppliedScheduledPlanRequest {
	s.DeploymentId = &v
	return s
}

func (s *GetAppliedScheduledPlanRequest) Validate() error {
	return dara.Validate(s)
}
