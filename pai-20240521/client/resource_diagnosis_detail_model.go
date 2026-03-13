// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResourceDiagnosisDetail interface {
	dara.Model
	String() string
	GoString() string
	SetExceedResources(v []*string) *ResourceDiagnosisDetail
	GetExceedResources() []*string
	SetLimit(v *ResourceAmount) *ResourceDiagnosisDetail
	GetLimit() *ResourceAmount
	SetStatus(v string) *ResourceDiagnosisDetail
	GetStatus() *string
	SetUsed(v *ResourceAmount) *ResourceDiagnosisDetail
	GetUsed() *ResourceAmount
	SetWorkloadIds(v []*string) *ResourceDiagnosisDetail
	GetWorkloadIds() []*string
}

type ResourceDiagnosisDetail struct {
	ExceedResources []*string       `json:"ExceedResources,omitempty" xml:"ExceedResources,omitempty" type:"Repeated"`
	Limit           *ResourceAmount `json:"Limit,omitempty" xml:"Limit,omitempty"`
	Status          *string         `json:"Status,omitempty" xml:"Status,omitempty"`
	Used            *ResourceAmount `json:"Used,omitempty" xml:"Used,omitempty"`
	WorkloadIds     []*string       `json:"WorkloadIds,omitempty" xml:"WorkloadIds,omitempty" type:"Repeated"`
}

func (s ResourceDiagnosisDetail) String() string {
	return dara.Prettify(s)
}

func (s ResourceDiagnosisDetail) GoString() string {
	return s.String()
}

func (s *ResourceDiagnosisDetail) GetExceedResources() []*string {
	return s.ExceedResources
}

func (s *ResourceDiagnosisDetail) GetLimit() *ResourceAmount {
	return s.Limit
}

func (s *ResourceDiagnosisDetail) GetStatus() *string {
	return s.Status
}

func (s *ResourceDiagnosisDetail) GetUsed() *ResourceAmount {
	return s.Used
}

func (s *ResourceDiagnosisDetail) GetWorkloadIds() []*string {
	return s.WorkloadIds
}

func (s *ResourceDiagnosisDetail) SetExceedResources(v []*string) *ResourceDiagnosisDetail {
	s.ExceedResources = v
	return s
}

func (s *ResourceDiagnosisDetail) SetLimit(v *ResourceAmount) *ResourceDiagnosisDetail {
	s.Limit = v
	return s
}

func (s *ResourceDiagnosisDetail) SetStatus(v string) *ResourceDiagnosisDetail {
	s.Status = &v
	return s
}

func (s *ResourceDiagnosisDetail) SetUsed(v *ResourceAmount) *ResourceDiagnosisDetail {
	s.Used = v
	return s
}

func (s *ResourceDiagnosisDetail) SetWorkloadIds(v []*string) *ResourceDiagnosisDetail {
	s.WorkloadIds = v
	return s
}

func (s *ResourceDiagnosisDetail) Validate() error {
	if s.Limit != nil {
		if err := s.Limit.Validate(); err != nil {
			return err
		}
	}
	if s.Used != nil {
		if err := s.Used.Validate(); err != nil {
			return err
		}
	}
	return nil
}
