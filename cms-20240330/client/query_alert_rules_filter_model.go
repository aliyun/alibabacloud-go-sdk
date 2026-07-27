// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAlertRulesFilter interface {
	dara.Model
	String() string
	GoString() string
	SetBizSource(v *BizSourceFilter) *QueryAlertRulesFilter
	GetBizSource() *BizSourceFilter
	SetDatasourceType(v *DatasourceTypeFilter) *QueryAlertRulesFilter
	GetDatasourceType() *DatasourceTypeFilter
	SetDisplayName(v *DisplayNameFilter) *QueryAlertRulesFilter
	GetDisplayName() *DisplayNameFilter
	SetEnabled(v *EnabledFilter) *QueryAlertRulesFilter
	GetEnabled() *EnabledFilter
	SetLabels(v *LabelsFilter) *QueryAlertRulesFilter
	GetLabels() *LabelsFilter
	SetNotificationChannels(v *NotificationChannelsFilter) *QueryAlertRulesFilter
	GetNotificationChannels() *NotificationChannelsFilter
	SetNotifyStrategyId(v *NotifyStrategyIdFilter) *QueryAlertRulesFilter
	GetNotifyStrategyId() *NotifyStrategyIdFilter
	SetObserveResourceConfig(v *ObserveResourceConfigFilter) *QueryAlertRulesFilter
	GetObserveResourceConfig() *ObserveResourceConfigFilter
	SetObserveResourceGlobalScope(v *ObserveResourceGlobalScopeFilter) *QueryAlertRulesFilter
	GetObserveResourceGlobalScope() *ObserveResourceGlobalScopeFilter
	SetObserveResourceInstanceId(v string) *QueryAlertRulesFilter
	GetObserveResourceInstanceId() *string
	SetObserveResourceList(v *ObserveResourceListFilter) *QueryAlertRulesFilter
	GetObserveResourceList() *ObserveResourceListFilter
	SetObserveResourceType(v *ObserveResourceTypeFilter) *QueryAlertRulesFilter
	GetObserveResourceType() *ObserveResourceTypeFilter
	SetPartitionKey(v *PartitionKeyFilter) *QueryAlertRulesFilter
	GetPartitionKey() *PartitionKeyFilter
	SetSeverityLevels(v *SeverityLevelsFilter) *QueryAlertRulesFilter
	GetSeverityLevels() *SeverityLevelsFilter
	SetStatus(v *StatusFilter) *QueryAlertRulesFilter
	GetStatus() *StatusFilter
	SetUuid(v *UuidFilter) *QueryAlertRulesFilter
	GetUuid() *UuidFilter
}

type QueryAlertRulesFilter struct {
	BizSource                  *BizSourceFilter                  `json:"bizSource,omitempty" xml:"bizSource,omitempty"`
	DatasourceType             *DatasourceTypeFilter             `json:"datasourceType,omitempty" xml:"datasourceType,omitempty"`
	DisplayName                *DisplayNameFilter                `json:"displayName,omitempty" xml:"displayName,omitempty"`
	Enabled                    *EnabledFilter                    `json:"enabled,omitempty" xml:"enabled,omitempty"`
	Labels                     *LabelsFilter                     `json:"labels,omitempty" xml:"labels,omitempty"`
	NotificationChannels       *NotificationChannelsFilter       `json:"notificationChannels,omitempty" xml:"notificationChannels,omitempty"`
	NotifyStrategyId           *NotifyStrategyIdFilter           `json:"notifyStrategyId,omitempty" xml:"notifyStrategyId,omitempty"`
	ObserveResourceConfig      *ObserveResourceConfigFilter      `json:"observeResourceConfig,omitempty" xml:"observeResourceConfig,omitempty"`
	ObserveResourceGlobalScope *ObserveResourceGlobalScopeFilter `json:"observeResourceGlobalScope,omitempty" xml:"observeResourceGlobalScope,omitempty"`
	// Deprecated
	//
	// example:
	//
	// i-bp1abcxxxxxxxx
	ObserveResourceInstanceId *string                    `json:"observeResourceInstanceId,omitempty" xml:"observeResourceInstanceId,omitempty"`
	ObserveResourceList       *ObserveResourceListFilter `json:"observeResourceList,omitempty" xml:"observeResourceList,omitempty"`
	// Deprecated
	ObserveResourceType *ObserveResourceTypeFilter `json:"observeResourceType,omitempty" xml:"observeResourceType,omitempty"`
	PartitionKey        *PartitionKeyFilter        `json:"partitionKey,omitempty" xml:"partitionKey,omitempty"`
	SeverityLevels      *SeverityLevelsFilter      `json:"severityLevels,omitempty" xml:"severityLevels,omitempty"`
	Status              *StatusFilter              `json:"status,omitempty" xml:"status,omitempty"`
	Uuid                *UuidFilter                `json:"uuid,omitempty" xml:"uuid,omitempty"`
}

func (s QueryAlertRulesFilter) String() string {
	return dara.Prettify(s)
}

func (s QueryAlertRulesFilter) GoString() string {
	return s.String()
}

func (s *QueryAlertRulesFilter) GetBizSource() *BizSourceFilter {
	return s.BizSource
}

func (s *QueryAlertRulesFilter) GetDatasourceType() *DatasourceTypeFilter {
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

func (s *QueryAlertRulesFilter) GetNotificationChannels() *NotificationChannelsFilter {
	return s.NotificationChannels
}

func (s *QueryAlertRulesFilter) GetNotifyStrategyId() *NotifyStrategyIdFilter {
	return s.NotifyStrategyId
}

func (s *QueryAlertRulesFilter) GetObserveResourceConfig() *ObserveResourceConfigFilter {
	return s.ObserveResourceConfig
}

func (s *QueryAlertRulesFilter) GetObserveResourceGlobalScope() *ObserveResourceGlobalScopeFilter {
	return s.ObserveResourceGlobalScope
}

func (s *QueryAlertRulesFilter) GetObserveResourceInstanceId() *string {
	return s.ObserveResourceInstanceId
}

func (s *QueryAlertRulesFilter) GetObserveResourceList() *ObserveResourceListFilter {
	return s.ObserveResourceList
}

func (s *QueryAlertRulesFilter) GetObserveResourceType() *ObserveResourceTypeFilter {
	return s.ObserveResourceType
}

func (s *QueryAlertRulesFilter) GetPartitionKey() *PartitionKeyFilter {
	return s.PartitionKey
}

func (s *QueryAlertRulesFilter) GetSeverityLevels() *SeverityLevelsFilter {
	return s.SeverityLevels
}

func (s *QueryAlertRulesFilter) GetStatus() *StatusFilter {
	return s.Status
}

func (s *QueryAlertRulesFilter) GetUuid() *UuidFilter {
	return s.Uuid
}

func (s *QueryAlertRulesFilter) SetBizSource(v *BizSourceFilter) *QueryAlertRulesFilter {
	s.BizSource = v
	return s
}

func (s *QueryAlertRulesFilter) SetDatasourceType(v *DatasourceTypeFilter) *QueryAlertRulesFilter {
	s.DatasourceType = v
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

func (s *QueryAlertRulesFilter) SetNotificationChannels(v *NotificationChannelsFilter) *QueryAlertRulesFilter {
	s.NotificationChannels = v
	return s
}

func (s *QueryAlertRulesFilter) SetNotifyStrategyId(v *NotifyStrategyIdFilter) *QueryAlertRulesFilter {
	s.NotifyStrategyId = v
	return s
}

func (s *QueryAlertRulesFilter) SetObserveResourceConfig(v *ObserveResourceConfigFilter) *QueryAlertRulesFilter {
	s.ObserveResourceConfig = v
	return s
}

func (s *QueryAlertRulesFilter) SetObserveResourceGlobalScope(v *ObserveResourceGlobalScopeFilter) *QueryAlertRulesFilter {
	s.ObserveResourceGlobalScope = v
	return s
}

func (s *QueryAlertRulesFilter) SetObserveResourceInstanceId(v string) *QueryAlertRulesFilter {
	s.ObserveResourceInstanceId = &v
	return s
}

func (s *QueryAlertRulesFilter) SetObserveResourceList(v *ObserveResourceListFilter) *QueryAlertRulesFilter {
	s.ObserveResourceList = v
	return s
}

func (s *QueryAlertRulesFilter) SetObserveResourceType(v *ObserveResourceTypeFilter) *QueryAlertRulesFilter {
	s.ObserveResourceType = v
	return s
}

func (s *QueryAlertRulesFilter) SetPartitionKey(v *PartitionKeyFilter) *QueryAlertRulesFilter {
	s.PartitionKey = v
	return s
}

func (s *QueryAlertRulesFilter) SetSeverityLevels(v *SeverityLevelsFilter) *QueryAlertRulesFilter {
	s.SeverityLevels = v
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
	if s.BizSource != nil {
		if err := s.BizSource.Validate(); err != nil {
			return err
		}
	}
	if s.DatasourceType != nil {
		if err := s.DatasourceType.Validate(); err != nil {
			return err
		}
	}
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
	if s.NotificationChannels != nil {
		if err := s.NotificationChannels.Validate(); err != nil {
			return err
		}
	}
	if s.NotifyStrategyId != nil {
		if err := s.NotifyStrategyId.Validate(); err != nil {
			return err
		}
	}
	if s.ObserveResourceConfig != nil {
		if err := s.ObserveResourceConfig.Validate(); err != nil {
			return err
		}
	}
	if s.ObserveResourceGlobalScope != nil {
		if err := s.ObserveResourceGlobalScope.Validate(); err != nil {
			return err
		}
	}
	if s.ObserveResourceList != nil {
		if err := s.ObserveResourceList.Validate(); err != nil {
			return err
		}
	}
	if s.ObserveResourceType != nil {
		if err := s.ObserveResourceType.Validate(); err != nil {
			return err
		}
	}
	if s.PartitionKey != nil {
		if err := s.PartitionKey.Validate(); err != nil {
			return err
		}
	}
	if s.SeverityLevels != nil {
		if err := s.SeverityLevels.Validate(); err != nil {
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
