// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAlertRulesFilter interface {
	dara.Model
	String() string
	GoString() string
	SetDatasourceType(v string) *QueryAlertRulesFilter
	GetDatasourceType() *string
	SetDisplayName(v *DisplayNameFilter) *QueryAlertRulesFilter
	GetDisplayName() *DisplayNameFilter
	SetEnabled(v *EnabledFilter) *QueryAlertRulesFilter
	GetEnabled() *EnabledFilter
	SetLabels(v *LabelsFilter) *QueryAlertRulesFilter
	GetLabels() *LabelsFilter
	SetObserveResourceGlobalScope(v bool) *QueryAlertRulesFilter
	GetObserveResourceGlobalScope() *bool
	SetObserveResourceInstanceId(v string) *QueryAlertRulesFilter
	GetObserveResourceInstanceId() *string
	SetObserveResourceType(v string) *QueryAlertRulesFilter
	GetObserveResourceType() *string
	SetSeverityLevels(v string) *QueryAlertRulesFilter
	GetSeverityLevels() *string
	SetStatus(v *StatusFilter) *QueryAlertRulesFilter
	GetStatus() *StatusFilter
	SetUuid(v *UuidFilter) *QueryAlertRulesFilter
	GetUuid() *UuidFilter
}

type QueryAlertRulesFilter struct {
	DatasourceType *string `json:"datasourceType,omitempty" xml:"datasourceType,omitempty"`
	// Filters alert rules by display name.
	DisplayName *DisplayNameFilter `json:"displayName,omitempty" xml:"displayName,omitempty"`
	// Filters alert rules by enabled status.
	Enabled *EnabledFilter `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// Filters alert rules by label.
	Labels                     *LabelsFilter `json:"labels,omitempty" xml:"labels,omitempty"`
	ObserveResourceGlobalScope *bool         `json:"observeResourceGlobalScope,omitempty" xml:"observeResourceGlobalScope,omitempty"`
	// example:
	//
	// i-bp1abcxxxxxxxx
	ObserveResourceInstanceId *string `json:"observeResourceInstanceId,omitempty" xml:"observeResourceInstanceId,omitempty"`
	ObserveResourceType       *string `json:"observeResourceType,omitempty" xml:"observeResourceType,omitempty"`
	SeverityLevels            *string `json:"severityLevels,omitempty" xml:"severityLevels,omitempty"`
	// Filters alert rules by status.
	Status *StatusFilter `json:"status,omitempty" xml:"status,omitempty"`
	// Filters alert rules by UUID.
	Uuid *UuidFilter `json:"uuid,omitempty" xml:"uuid,omitempty"`
}

func (s QueryAlertRulesFilter) String() string {
	return dara.Prettify(s)
}

func (s QueryAlertRulesFilter) GoString() string {
	return s.String()
}

func (s *QueryAlertRulesFilter) GetDatasourceType() *string {
	return s.DatasourceType
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

func (s *QueryAlertRulesFilter) GetObserveResourceGlobalScope() *bool {
	return s.ObserveResourceGlobalScope
}

func (s *QueryAlertRulesFilter) GetObserveResourceInstanceId() *string {
	return s.ObserveResourceInstanceId
}

func (s *QueryAlertRulesFilter) GetObserveResourceType() *string {
	return s.ObserveResourceType
}

func (s *QueryAlertRulesFilter) GetSeverityLevels() *string {
	return s.SeverityLevels
}

func (s *QueryAlertRulesFilter) GetStatus() *StatusFilter {
	return s.Status
}

func (s *QueryAlertRulesFilter) GetUuid() *UuidFilter {
	return s.Uuid
}

func (s *QueryAlertRulesFilter) SetDatasourceType(v string) *QueryAlertRulesFilter {
	s.DatasourceType = &v
	return s
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

func (s *QueryAlertRulesFilter) SetObserveResourceGlobalScope(v bool) *QueryAlertRulesFilter {
	s.ObserveResourceGlobalScope = &v
	return s
}

func (s *QueryAlertRulesFilter) SetObserveResourceInstanceId(v string) *QueryAlertRulesFilter {
	s.ObserveResourceInstanceId = &v
	return s
}

func (s *QueryAlertRulesFilter) SetObserveResourceType(v string) *QueryAlertRulesFilter {
	s.ObserveResourceType = &v
	return s
}

func (s *QueryAlertRulesFilter) SetSeverityLevels(v string) *QueryAlertRulesFilter {
	s.SeverityLevels = &v
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
