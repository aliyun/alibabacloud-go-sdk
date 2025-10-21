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
	// example:
	//
	// TRUE
	AllowNonRestoredState *bool `json:"allowNonRestoredState,omitempty" xml:"allowNonRestoredState,omitempty"`
	// example:
	//
	// 1660293803155
	JobStartTimeInMs *int64 `json:"jobStartTimeInMs,omitempty" xml:"jobStartTimeInMs,omitempty"`
	// example:
	//
	// LATEST_STATE
	Kind *string `json:"kind,omitempty" xml:"kind,omitempty"`
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
