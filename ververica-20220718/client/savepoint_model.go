// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSavepoint interface {
	dara.Model
	String() string
	GoString() string
	SetCreatedAt(v int64) *Savepoint
	GetCreatedAt() *int64
	SetDeploymentId(v string) *Savepoint
	GetDeploymentId() *string
	SetDescription(v string) *Savepoint
	GetDescription() *string
	SetJobId(v string) *Savepoint
	GetJobId() *string
	SetModifiedAt(v int64) *Savepoint
	GetModifiedAt() *int64
	SetNamespace(v string) *Savepoint
	GetNamespace() *string
	SetNativeFormat(v bool) *Savepoint
	GetNativeFormat() *bool
	SetSavepointId(v string) *Savepoint
	GetSavepointId() *string
	SetSavepointLocation(v string) *Savepoint
	GetSavepointLocation() *string
	SetSavepointOrigin(v string) *Savepoint
	GetSavepointOrigin() *string
	SetStatus(v *SavepointStatus) *Savepoint
	GetStatus() *SavepointStatus
	SetStopWithDrainEnabled(v bool) *Savepoint
	GetStopWithDrainEnabled() *bool
}

type Savepoint struct {
	// The time when the savepoint was created.
	//
	// example:
	//
	// 1659066711
	CreatedAt *int64 `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// The deployment ID.
	//
	// example:
	//
	// 1d716b22-6aad-4be2-85c2-50cfc757****
	DeploymentId *string `json:"deploymentId,omitempty" xml:"deploymentId,omitempty"`
	// The description of the savepoint.
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The job ID.
	//
	// example:
	//
	// 5af678c0-7db0-4650-94c2-d2604f0a****
	JobId *string `json:"jobId,omitempty" xml:"jobId,omitempty"`
	// The time when the savepoint was last modified.
	//
	// example:
	//
	// 1659069473
	ModifiedAt *int64 `json:"modifiedAt,omitempty" xml:"modifiedAt,omitempty"`
	// The name of the namespace.
	//
	// example:
	//
	// namespacetest
	Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
	// Specifies whether the savepoint is in native format.
	//
	// example:
	//
	// TRUE
	NativeFormat *bool `json:"nativeFormat,omitempty" xml:"nativeFormat,omitempty"`
	// The savepoint ID.
	//
	// example:
	//
	// 354dde66-a3ae-463e-967a-0b4107fd****
	SavepointId *string `json:"savepointId,omitempty" xml:"savepointId,omitempty"`
	// The storage URL of the savepoint.
	//
	// example:
	//
	// https://oss/bucket/flink/flink-jobs/namespaces/vvp-team/deployments/5a19a71b-1c42-4f34-94fd-86cf60782c81/checkpoints/sp-3285
	SavepointLocation *string `json:"savepointLocation,omitempty" xml:"savepointLocation,omitempty"`
	// The method that is used to create a savepoint.
	//
	// 	- `USER_REQUEST`: The savepoint is manually created.
	//
	// 	- `STOP_WITH_SAVEPOINT`: The savepoint is created when you cancel the deployment.
	//
	// 	- `RETAINED_CHECKPOINT`: The savepoint is created based on the returned checkpoint.
	//
	// example:
	//
	// USER_REQUEST
	SavepointOrigin *string `json:"savepointOrigin,omitempty" xml:"savepointOrigin,omitempty"`
	// The status of the savepoint.
	Status *SavepointStatus `json:"status,omitempty" xml:"status,omitempty"`
	// Specifies whether to use the stop-with-drain mode.
	//
	// example:
	//
	// TRUE
	StopWithDrainEnabled *bool `json:"stopWithDrainEnabled,omitempty" xml:"stopWithDrainEnabled,omitempty"`
}

func (s Savepoint) String() string {
	return dara.Prettify(s)
}

func (s Savepoint) GoString() string {
	return s.String()
}

func (s *Savepoint) GetCreatedAt() *int64 {
	return s.CreatedAt
}

func (s *Savepoint) GetDeploymentId() *string {
	return s.DeploymentId
}

func (s *Savepoint) GetDescription() *string {
	return s.Description
}

func (s *Savepoint) GetJobId() *string {
	return s.JobId
}

func (s *Savepoint) GetModifiedAt() *int64 {
	return s.ModifiedAt
}

func (s *Savepoint) GetNamespace() *string {
	return s.Namespace
}

func (s *Savepoint) GetNativeFormat() *bool {
	return s.NativeFormat
}

func (s *Savepoint) GetSavepointId() *string {
	return s.SavepointId
}

func (s *Savepoint) GetSavepointLocation() *string {
	return s.SavepointLocation
}

func (s *Savepoint) GetSavepointOrigin() *string {
	return s.SavepointOrigin
}

func (s *Savepoint) GetStatus() *SavepointStatus {
	return s.Status
}

func (s *Savepoint) GetStopWithDrainEnabled() *bool {
	return s.StopWithDrainEnabled
}

func (s *Savepoint) SetCreatedAt(v int64) *Savepoint {
	s.CreatedAt = &v
	return s
}

func (s *Savepoint) SetDeploymentId(v string) *Savepoint {
	s.DeploymentId = &v
	return s
}

func (s *Savepoint) SetDescription(v string) *Savepoint {
	s.Description = &v
	return s
}

func (s *Savepoint) SetJobId(v string) *Savepoint {
	s.JobId = &v
	return s
}

func (s *Savepoint) SetModifiedAt(v int64) *Savepoint {
	s.ModifiedAt = &v
	return s
}

func (s *Savepoint) SetNamespace(v string) *Savepoint {
	s.Namespace = &v
	return s
}

func (s *Savepoint) SetNativeFormat(v bool) *Savepoint {
	s.NativeFormat = &v
	return s
}

func (s *Savepoint) SetSavepointId(v string) *Savepoint {
	s.SavepointId = &v
	return s
}

func (s *Savepoint) SetSavepointLocation(v string) *Savepoint {
	s.SavepointLocation = &v
	return s
}

func (s *Savepoint) SetSavepointOrigin(v string) *Savepoint {
	s.SavepointOrigin = &v
	return s
}

func (s *Savepoint) SetStatus(v *SavepointStatus) *Savepoint {
	s.Status = v
	return s
}

func (s *Savepoint) SetStopWithDrainEnabled(v bool) *Savepoint {
	s.StopWithDrainEnabled = &v
	return s
}

func (s *Savepoint) Validate() error {
	if s.Status != nil {
		if err := s.Status.Validate(); err != nil {
			return err
		}
	}
	return nil
}
