// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFilter interface {
	dara.Model
	String() string
	GoString() string
	SetDisplayName(v *DisplayNameFilter) *Filter
	GetDisplayName() *DisplayNameFilter
	SetEnabled(v *EnabledFilter) *Filter
	GetEnabled() *EnabledFilter
	SetLabels(v *LabelsFilter) *Filter
	GetLabels() *LabelsFilter
	SetStatus(v *StatusFilter) *Filter
	GetStatus() *StatusFilter
	SetUuid(v *UuidFilter) *Filter
	GetUuid() *UuidFilter
}

type Filter struct {
	DisplayName *DisplayNameFilter `json:"displayName,omitempty" xml:"displayName,omitempty"`
	Enabled     *EnabledFilter     `json:"enabled,omitempty" xml:"enabled,omitempty"`
	Labels      *LabelsFilter      `json:"labels,omitempty" xml:"labels,omitempty"`
	Status      *StatusFilter      `json:"status,omitempty" xml:"status,omitempty"`
	Uuid        *UuidFilter        `json:"uuid,omitempty" xml:"uuid,omitempty"`
}

func (s Filter) String() string {
	return dara.Prettify(s)
}

func (s Filter) GoString() string {
	return s.String()
}

func (s *Filter) GetDisplayName() *DisplayNameFilter {
	return s.DisplayName
}

func (s *Filter) GetEnabled() *EnabledFilter {
	return s.Enabled
}

func (s *Filter) GetLabels() *LabelsFilter {
	return s.Labels
}

func (s *Filter) GetStatus() *StatusFilter {
	return s.Status
}

func (s *Filter) GetUuid() *UuidFilter {
	return s.Uuid
}

func (s *Filter) SetDisplayName(v *DisplayNameFilter) *Filter {
	s.DisplayName = v
	return s
}

func (s *Filter) SetEnabled(v *EnabledFilter) *Filter {
	s.Enabled = v
	return s
}

func (s *Filter) SetLabels(v *LabelsFilter) *Filter {
	s.Labels = v
	return s
}

func (s *Filter) SetStatus(v *StatusFilter) *Filter {
	s.Status = v
	return s
}

func (s *Filter) SetUuid(v *UuidFilter) *Filter {
	s.Uuid = v
	return s
}

func (s *Filter) Validate() error {
	if s.DisplayName != nil {
		if err := s.DisplayName.Validate(); err != nil {
			return err
		}
	}
	if s.Enabled != nil {
		if err := s.Enabled.Validate(); err != nil {
			return err
		}
	}
	if s.Labels != nil {
		if err := s.Labels.Validate(); err != nil {
			return err
		}
	}
	if s.Status != nil {
		if err := s.Status.Validate(); err != nil {
			return err
		}
	}
	if s.Uuid != nil {
		if err := s.Uuid.Validate(); err != nil {
			return err
		}
	}
	return nil
}
