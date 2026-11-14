// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublicWarmPool interface {
	dara.Model
	String() string
	GoString() string
	SetCapacityObservedTime(v string) *PublicWarmPool
	GetCapacityObservedTime() *string
	SetCapacityPolicy(v *WarmPoolCapacityPolicy) *PublicWarmPool
	GetCapacityPolicy() *WarmPoolCapacityPolicy
	SetCurrentError(v string) *PublicWarmPool
	GetCurrentError() *string
	SetDeleting(v bool) *PublicWarmPool
	GetDeleting() *bool
	SetPreparing(v int64) *PublicWarmPool
	GetPreparing() *int64
	SetReady(v int64) *PublicWarmPool
	GetReady() *int64
	SetReleasing(v int64) *PublicWarmPool
	GetReleasing() *int64
	SetSourceUnavailable(v bool) *PublicWarmPool
	GetSourceUnavailable() *bool
	SetTargetCapacity(v int64) *PublicWarmPool
	GetTargetCapacity() *int64
	SetTeamID(v string) *PublicWarmPool
	GetTeamID() *string
	SetTemplateID(v string) *PublicWarmPool
	GetTemplateID() *string
}

type PublicWarmPool struct {
	// The most recent observation time of the capacity status.
	//
	// example:
	//
	// 2026-12-01T00:00:00Z
	CapacityObservedTime *string `json:"capacityObservedTime,omitempty" xml:"capacityObservedTime,omitempty"`
	// The capacity policy of the warm pool, which includes the base target capacity and optional scheduled rules.
	CapacityPolicy *WarmPoolCapacityPolicy `json:"capacityPolicy,omitempty" xml:"capacityPolicy,omitempty"`
	// The masked text of the current background error. This value is null if no error occurs.
	//
	// example:
	//
	// SourceNotFound: source template does not exist
	CurrentError *string `json:"currentError,omitempty" xml:"currentError,omitempty"`
	// Indicates whether the warm pool is being deleted.
	Deleting *bool `json:"deleting,omitempty" xml:"deleting,omitempty"`
	// The number of sandboxes that are currently being prepared.
	//
	// example:
	//
	// 0
	Preparing *int64 `json:"preparing,omitempty" xml:"preparing,omitempty"`
	// The number of sandboxes that are currently ready.
	//
	// example:
	//
	// 0
	Ready *int64 `json:"ready,omitempty" xml:"ready,omitempty"`
	// The number of sandboxes that are currently being released.
	//
	// example:
	//
	// 0
	Releasing *int64 `json:"releasing,omitempty" xml:"releasing,omitempty"`
	// Indicates whether the source of the associated template is unavailable.
	SourceUnavailable *bool `json:"sourceUnavailable,omitempty" xml:"sourceUnavailable,omitempty"`
	// The current target capacity.
	//
	// example:
	//
	// 0
	TargetCapacity *int64 `json:"targetCapacity,omitempty" xml:"targetCapacity,omitempty"`
	// The team ID.
	//
	// example:
	//
	// team-example1234567890
	TeamID *string `json:"teamID,omitempty" xml:"teamID,omitempty"`
	// The template ID, which specifies the template associated with the warm pool.
	//
	// example:
	//
	// tpl-example1234567890
	TemplateID *string `json:"templateID,omitempty" xml:"templateID,omitempty"`
}

func (s PublicWarmPool) String() string {
	return dara.Prettify(s)
}

func (s PublicWarmPool) GoString() string {
	return s.String()
}

func (s *PublicWarmPool) GetCapacityObservedTime() *string {
	return s.CapacityObservedTime
}

func (s *PublicWarmPool) GetCapacityPolicy() *WarmPoolCapacityPolicy {
	return s.CapacityPolicy
}

func (s *PublicWarmPool) GetCurrentError() *string {
	return s.CurrentError
}

func (s *PublicWarmPool) GetDeleting() *bool {
	return s.Deleting
}

func (s *PublicWarmPool) GetPreparing() *int64 {
	return s.Preparing
}

func (s *PublicWarmPool) GetReady() *int64 {
	return s.Ready
}

func (s *PublicWarmPool) GetReleasing() *int64 {
	return s.Releasing
}

func (s *PublicWarmPool) GetSourceUnavailable() *bool {
	return s.SourceUnavailable
}

func (s *PublicWarmPool) GetTargetCapacity() *int64 {
	return s.TargetCapacity
}

func (s *PublicWarmPool) GetTeamID() *string {
	return s.TeamID
}

func (s *PublicWarmPool) GetTemplateID() *string {
	return s.TemplateID
}

func (s *PublicWarmPool) SetCapacityObservedTime(v string) *PublicWarmPool {
	s.CapacityObservedTime = &v
	return s
}

func (s *PublicWarmPool) SetCapacityPolicy(v *WarmPoolCapacityPolicy) *PublicWarmPool {
	s.CapacityPolicy = v
	return s
}

func (s *PublicWarmPool) SetCurrentError(v string) *PublicWarmPool {
	s.CurrentError = &v
	return s
}

func (s *PublicWarmPool) SetDeleting(v bool) *PublicWarmPool {
	s.Deleting = &v
	return s
}

func (s *PublicWarmPool) SetPreparing(v int64) *PublicWarmPool {
	s.Preparing = &v
	return s
}

func (s *PublicWarmPool) SetReady(v int64) *PublicWarmPool {
	s.Ready = &v
	return s
}

func (s *PublicWarmPool) SetReleasing(v int64) *PublicWarmPool {
	s.Releasing = &v
	return s
}

func (s *PublicWarmPool) SetSourceUnavailable(v bool) *PublicWarmPool {
	s.SourceUnavailable = &v
	return s
}

func (s *PublicWarmPool) SetTargetCapacity(v int64) *PublicWarmPool {
	s.TargetCapacity = &v
	return s
}

func (s *PublicWarmPool) SetTeamID(v string) *PublicWarmPool {
	s.TeamID = &v
	return s
}

func (s *PublicWarmPool) SetTemplateID(v string) *PublicWarmPool {
	s.TemplateID = &v
	return s
}

func (s *PublicWarmPool) Validate() error {
	if s.CapacityPolicy != nil {
		if err := s.CapacityPolicy.Validate(); err != nil {
			return err
		}
	}
	return nil
}
