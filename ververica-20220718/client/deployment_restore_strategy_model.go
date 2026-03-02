// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeploymentRestoreStrategy interface {
	dara.Model
	String() string
	GoString() string
	SetAllowNonRestoredState(v bool) *DeploymentRestoreStrategy
	GetAllowNonRestoredState() *bool
	SetJobStartTimeInMs(v int64) *DeploymentRestoreStrategy
	GetJobStartTimeInMs() *int64
	SetKind(v string) *DeploymentRestoreStrategy
	GetKind() *string
	SetSavepointId(v string) *DeploymentRestoreStrategy
	GetSavepointId() *string
}

type DeploymentRestoreStrategy struct {
	// Specifies whether to allow specific operator states to be skipped. This parameter is required only when a Python deployment or a JAR deployment is resumed with state data.
	//
	// example:
	//
	// TRUE
	AllowNonRestoredState *bool `json:"allowNonRestoredState,omitempty" xml:"allowNonRestoredState,omitempty"`
	// The time point at which the deployment is started without states. You must enter a 13-digit timestamp. If you set the kind parameter to NONE, you can configure this parameter to allow all source tables for which the startTime parameter is configured to read data from the specified time point.
	//
	// example:
	//
	// 1660293803155
	JobStartTimeInMs *int64 `json:"jobStartTimeInMs,omitempty" xml:"jobStartTimeInMs,omitempty"`
	// The type of the start offset. Valid values:
	//
	// 	- NONE: The deployment is started without states.
	//
	// 	- LATEST_SAVEPOINT: The deployment is started from the latest savepoint.
	//
	// 	- FROM_SAVEPOINT: The deployment is started from the specified savepoint.
	//
	// 	- LATEST_STATE: The deployment is started from the latest state of the deployment.
	//
	// example:
	//
	// LATEST_STATE
	Kind *string `json:"kind,omitempty" xml:"kind,omitempty"`
	// The ID of the savepoint for starting the deployment. This parameter is required when the kind parameter is set to FROM_SAVEPOINT.
	//
	// example:
	//
	// 354dde66-a3ae-463e-967a-0b4107fd****
	SavepointId *string `json:"savepointId,omitempty" xml:"savepointId,omitempty"`
}

func (s DeploymentRestoreStrategy) String() string {
	return dara.Prettify(s)
}

func (s DeploymentRestoreStrategy) GoString() string {
	return s.String()
}

func (s *DeploymentRestoreStrategy) GetAllowNonRestoredState() *bool {
	return s.AllowNonRestoredState
}

func (s *DeploymentRestoreStrategy) GetJobStartTimeInMs() *int64 {
	return s.JobStartTimeInMs
}

func (s *DeploymentRestoreStrategy) GetKind() *string {
	return s.Kind
}

func (s *DeploymentRestoreStrategy) GetSavepointId() *string {
	return s.SavepointId
}

func (s *DeploymentRestoreStrategy) SetAllowNonRestoredState(v bool) *DeploymentRestoreStrategy {
	s.AllowNonRestoredState = &v
	return s
}

func (s *DeploymentRestoreStrategy) SetJobStartTimeInMs(v int64) *DeploymentRestoreStrategy {
	s.JobStartTimeInMs = &v
	return s
}

func (s *DeploymentRestoreStrategy) SetKind(v string) *DeploymentRestoreStrategy {
	s.Kind = &v
	return s
}

func (s *DeploymentRestoreStrategy) SetSavepointId(v string) *DeploymentRestoreStrategy {
	s.SavepointId = &v
	return s
}

func (s *DeploymentRestoreStrategy) Validate() error {
	return dara.Validate(s)
}
