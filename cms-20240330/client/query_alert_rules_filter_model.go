// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAlertRulesFilter interface {
	dara.Model
	String() string
	GoString() string
	SetDisplayName(v *DisplayNameFilter) *QueryAlertRulesFilter
	GetDisplayName() *DisplayNameFilter
	SetEnabled(v *EnabledFilter) *QueryAlertRulesFilter
	GetEnabled() *EnabledFilter
	SetLabels(v *LabelsFilter) *QueryAlertRulesFilter
	GetLabels() *LabelsFilter
	SetStatus(v *StatusFilter) *QueryAlertRulesFilter
	GetStatus() *StatusFilter
	SetUuid(v *UuidFilter) *QueryAlertRulesFilter
	GetUuid() *UuidFilter
}

type QueryAlertRulesFilter struct {
	DisplayName *DisplayNameFilter `json:"displayName,omitempty" xml:"displayName,omitempty"`
	Enabled     *EnabledFilter     `json:"enabled,omitempty" xml:"enabled,omitempty"`
	Labels      *LabelsFilter      `json:"labels,omitempty" xml:"labels,omitempty"`
	Status      *StatusFilter      `json:"status,omitempty" xml:"status,omitempty"`
	Uuid        *UuidFilter        `json:"uuid,omitempty" xml:"uuid,omitempty"`
}

func (s QueryAlertRulesFilter) String() string {
	return dara.Prettify(s)
}

func (s QueryAlertRulesFilter) GoString() string {
	return s.String()
}

func (s *QueryAlertRulesFilter) GetDisplayName() *DisplayNameFilter {
	return s.DisplayName
}

func (s *QueryAlertRulesFilter) GetEnabled() *EnabledFilter {
	return s.Enabled
}

func (s *QueryAlertRulesFilter) GetLabels() *LabelsFilter {
	return s.Labels
}

func (s *QueryAlertRulesFilter) GetStatus() *StatusFilter {
	return s.Status
}

func (s *QueryAlertRulesFilter) GetUuid() *UuidFilter {
	return s.Uuid
}

func (s *QueryAlertRulesFilter) SetDisplayName(v *DisplayNameFilter) *QueryAlertRulesFilter {
	s.DisplayName = v
	return s
}

func (s *QueryAlertRulesFilter) SetEnabled(v *EnabledFilter) *QueryAlertRulesFilter {
	s.Enabled = v
	return s
}

func (s *QueryAlertRulesFilter) SetLabels(v *LabelsFilter) *QueryAlertRulesFilter {
	s.Labels = v
	return s
}

func (s *QueryAlertRulesFilter) SetStatus(v *StatusFilter) *QueryAlertRulesFilter {
	s.Status = v
	return s
}

func (s *QueryAlertRulesFilter) SetUuid(v *UuidFilter) *QueryAlertRulesFilter {
	s.Uuid = v
	return s
}

func (s *QueryAlertRulesFilter) Validate() error {
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
