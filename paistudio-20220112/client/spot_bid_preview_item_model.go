// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSpotBidPreviewItem interface {
	dara.Model
	String() string
	GoString() string
	SetActive(v bool) *SpotBidPreviewItem
	GetActive() *bool
	SetAllowCrossHpnZone(v bool) *SpotBidPreviewItem
	GetAllowCrossHpnZone() *bool
	SetClusterId(v string) *SpotBidPreviewItem
	GetClusterId() *string
	SetGcLevel(v string) *SpotBidPreviewItem
	GetGcLevel() *string
	SetInstanceType(v string) *SpotBidPreviewItem
	GetInstanceType() *string
	SetJobName(v string) *SpotBidPreviewItem
	GetJobName() *string
	SetMaxDiscount(v float64) *SpotBidPreviewItem
	GetMaxDiscount() *float64
	SetMessage(v string) *SpotBidPreviewItem
	GetMessage() *string
	SetName(v string) *SpotBidPreviewItem
	GetName() *string
	SetPhase(v string) *SpotBidPreviewItem
	GetPhase() *string
	SetReplicas(v int32) *SpotBidPreviewItem
	GetReplicas() *int32
}

type SpotBidPreviewItem struct {
	// Indicates whether the spot bid is active. If set to `false`, the bid is paused.
	Active *bool `json:"active,omitempty" xml:"active,omitempty"`
	// Determines whether instances can be deployed across different High-Performance Network (HPN) zones. Defaults to `false`.
	AllowCrossHpnZone *bool `json:"allowCrossHpnZone,omitempty" xml:"allowCrossHpnZone,omitempty"`
	// The ID of the cluster where resources are provisioned.
	ClusterId *string `json:"clusterId,omitempty" xml:"clusterId,omitempty"`
	// The GC level for the spot instance.
	GcLevel *string `json:"gcLevel,omitempty" xml:"gcLevel,omitempty"`
	// The type of compute instance.
	InstanceType *string `json:"instanceType,omitempty" xml:"instanceType,omitempty"`
	// The name of the associated job.
	JobName *string `json:"jobName,omitempty" xml:"jobName,omitempty"`
	// The maximum discount percentage from the on-demand price.
	MaxDiscount *float64 `json:"maxDiscount,omitempty" xml:"maxDiscount,omitempty"`
	// A message that provides additional details about the current phase.
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The name of the spot bid preview.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The current phase of the spot bid preview. Valid values are `Pending`, `Active`, and `Failed`.
	Phase *string `json:"phase,omitempty" xml:"phase,omitempty"`
	// The number of instance replicas.
	Replicas *int32 `json:"replicas,omitempty" xml:"replicas,omitempty"`
}

func (s SpotBidPreviewItem) String() string {
	return dara.Prettify(s)
}

func (s SpotBidPreviewItem) GoString() string {
	return s.String()
}

func (s *SpotBidPreviewItem) GetActive() *bool {
	return s.Active
}

func (s *SpotBidPreviewItem) GetAllowCrossHpnZone() *bool {
	return s.AllowCrossHpnZone
}

func (s *SpotBidPreviewItem) GetClusterId() *string {
	return s.ClusterId
}

func (s *SpotBidPreviewItem) GetGcLevel() *string {
	return s.GcLevel
}

func (s *SpotBidPreviewItem) GetInstanceType() *string {
	return s.InstanceType
}

func (s *SpotBidPreviewItem) GetJobName() *string {
	return s.JobName
}

func (s *SpotBidPreviewItem) GetMaxDiscount() *float64 {
	return s.MaxDiscount
}

func (s *SpotBidPreviewItem) GetMessage() *string {
	return s.Message
}

func (s *SpotBidPreviewItem) GetName() *string {
	return s.Name
}

func (s *SpotBidPreviewItem) GetPhase() *string {
	return s.Phase
}

func (s *SpotBidPreviewItem) GetReplicas() *int32 {
	return s.Replicas
}

func (s *SpotBidPreviewItem) SetActive(v bool) *SpotBidPreviewItem {
	s.Active = &v
	return s
}

func (s *SpotBidPreviewItem) SetAllowCrossHpnZone(v bool) *SpotBidPreviewItem {
	s.AllowCrossHpnZone = &v
	return s
}

func (s *SpotBidPreviewItem) SetClusterId(v string) *SpotBidPreviewItem {
	s.ClusterId = &v
	return s
}

func (s *SpotBidPreviewItem) SetGcLevel(v string) *SpotBidPreviewItem {
	s.GcLevel = &v
	return s
}

func (s *SpotBidPreviewItem) SetInstanceType(v string) *SpotBidPreviewItem {
	s.InstanceType = &v
	return s
}

func (s *SpotBidPreviewItem) SetJobName(v string) *SpotBidPreviewItem {
	s.JobName = &v
	return s
}

func (s *SpotBidPreviewItem) SetMaxDiscount(v float64) *SpotBidPreviewItem {
	s.MaxDiscount = &v
	return s
}

func (s *SpotBidPreviewItem) SetMessage(v string) *SpotBidPreviewItem {
	s.Message = &v
	return s
}

func (s *SpotBidPreviewItem) SetName(v string) *SpotBidPreviewItem {
	s.Name = &v
	return s
}

func (s *SpotBidPreviewItem) SetPhase(v string) *SpotBidPreviewItem {
	s.Phase = &v
	return s
}

func (s *SpotBidPreviewItem) SetReplicas(v int32) *SpotBidPreviewItem {
	s.Replicas = &v
	return s
}

func (s *SpotBidPreviewItem) Validate() error {
	return dara.Validate(s)
}
