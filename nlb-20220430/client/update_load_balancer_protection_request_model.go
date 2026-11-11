// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLoadBalancerProtectionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *UpdateLoadBalancerProtectionRequest
	GetClientToken() *string
	SetDeletionProtectionEnabled(v bool) *UpdateLoadBalancerProtectionRequest
	GetDeletionProtectionEnabled() *bool
	SetDeletionProtectionReason(v string) *UpdateLoadBalancerProtectionRequest
	GetDeletionProtectionReason() *string
	SetDryRun(v bool) *UpdateLoadBalancerProtectionRequest
	GetDryRun() *bool
	SetLoadBalancerId(v string) *UpdateLoadBalancerProtectionRequest
	GetLoadBalancerId() *string
	SetModificationProtectionReason(v string) *UpdateLoadBalancerProtectionRequest
	GetModificationProtectionReason() *string
	SetModificationProtectionStatus(v string) *UpdateLoadBalancerProtectionRequest
	GetModificationProtectionStatus() *string
	SetRegionId(v string) *UpdateLoadBalancerProtectionRequest
	GetRegionId() *string
}

type UpdateLoadBalancerProtectionRequest struct {
	ClientToken               *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DeletionProtectionEnabled *bool   `json:"DeletionProtectionEnabled,omitempty" xml:"DeletionProtectionEnabled,omitempty"`
	DeletionProtectionReason  *string `json:"DeletionProtectionReason,omitempty" xml:"DeletionProtectionReason,omitempty"`
	DryRun                    *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// This parameter is required.
	LoadBalancerId               *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	ModificationProtectionReason *string `json:"ModificationProtectionReason,omitempty" xml:"ModificationProtectionReason,omitempty"`
	ModificationProtectionStatus *string `json:"ModificationProtectionStatus,omitempty" xml:"ModificationProtectionStatus,omitempty"`
	RegionId                     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s UpdateLoadBalancerProtectionRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateLoadBalancerProtectionRequest) GoString() string {
	return s.String()
}

func (s *UpdateLoadBalancerProtectionRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateLoadBalancerProtectionRequest) GetDeletionProtectionEnabled() *bool {
	return s.DeletionProtectionEnabled
}

func (s *UpdateLoadBalancerProtectionRequest) GetDeletionProtectionReason() *string {
	return s.DeletionProtectionReason
}

func (s *UpdateLoadBalancerProtectionRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *UpdateLoadBalancerProtectionRequest) GetLoadBalancerId() *string {
	return s.LoadBalancerId
}

func (s *UpdateLoadBalancerProtectionRequest) GetModificationProtectionReason() *string {
	return s.ModificationProtectionReason
}

func (s *UpdateLoadBalancerProtectionRequest) GetModificationProtectionStatus() *string {
	return s.ModificationProtectionStatus
}

func (s *UpdateLoadBalancerProtectionRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateLoadBalancerProtectionRequest) SetClientToken(v string) *UpdateLoadBalancerProtectionRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateLoadBalancerProtectionRequest) SetDeletionProtectionEnabled(v bool) *UpdateLoadBalancerProtectionRequest {
	s.DeletionProtectionEnabled = &v
	return s
}

func (s *UpdateLoadBalancerProtectionRequest) SetDeletionProtectionReason(v string) *UpdateLoadBalancerProtectionRequest {
	s.DeletionProtectionReason = &v
	return s
}

func (s *UpdateLoadBalancerProtectionRequest) SetDryRun(v bool) *UpdateLoadBalancerProtectionRequest {
	s.DryRun = &v
	return s
}

func (s *UpdateLoadBalancerProtectionRequest) SetLoadBalancerId(v string) *UpdateLoadBalancerProtectionRequest {
	s.LoadBalancerId = &v
	return s
}

func (s *UpdateLoadBalancerProtectionRequest) SetModificationProtectionReason(v string) *UpdateLoadBalancerProtectionRequest {
	s.ModificationProtectionReason = &v
	return s
}

func (s *UpdateLoadBalancerProtectionRequest) SetModificationProtectionStatus(v string) *UpdateLoadBalancerProtectionRequest {
	s.ModificationProtectionStatus = &v
	return s
}

func (s *UpdateLoadBalancerProtectionRequest) SetRegionId(v string) *UpdateLoadBalancerProtectionRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateLoadBalancerProtectionRequest) Validate() error {
	return dara.Validate(s)
}
