// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubscriptionForModify interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *SubscriptionForModify
	GetDescription() *string
	SetFilterSetting(v *FilterSetting) *SubscriptionForModify
	GetFilterSetting() *FilterSetting
	SetNotifyStrategyId(v string) *SubscriptionForModify
	GetNotifyStrategyId() *string
	SetPushingSetting(v *SubscriptionForModifyPushingSetting) *SubscriptionForModify
	GetPushingSetting() *SubscriptionForModifyPushingSetting
	SetSubscriptionName(v string) *SubscriptionForModify
	GetSubscriptionName() *string
}

type SubscriptionForModify struct {
	Description      *string                              `json:"description,omitempty" xml:"description,omitempty"`
	FilterSetting    *FilterSetting                       `json:"filterSetting,omitempty" xml:"filterSetting,omitempty"`
	NotifyStrategyId *string                              `json:"notifyStrategyId,omitempty" xml:"notifyStrategyId,omitempty"`
	PushingSetting   *SubscriptionForModifyPushingSetting `json:"pushingSetting,omitempty" xml:"pushingSetting,omitempty" type:"Struct"`
	// This parameter is required.
	SubscriptionName *string `json:"subscriptionName,omitempty" xml:"subscriptionName,omitempty"`
}

func (s SubscriptionForModify) String() string {
	return dara.Prettify(s)
}

func (s SubscriptionForModify) GoString() string {
	return s.String()
}

func (s *SubscriptionForModify) GetDescription() *string {
	return s.Description
}

func (s *SubscriptionForModify) GetFilterSetting() *FilterSetting {
	return s.FilterSetting
}

func (s *SubscriptionForModify) GetNotifyStrategyId() *string {
	return s.NotifyStrategyId
}

func (s *SubscriptionForModify) GetPushingSetting() *SubscriptionForModifyPushingSetting {
	return s.PushingSetting
}

func (s *SubscriptionForModify) GetSubscriptionName() *string {
	return s.SubscriptionName
}

func (s *SubscriptionForModify) SetDescription(v string) *SubscriptionForModify {
	s.Description = &v
	return s
}

func (s *SubscriptionForModify) SetFilterSetting(v *FilterSetting) *SubscriptionForModify {
	s.FilterSetting = v
	return s
}

func (s *SubscriptionForModify) SetNotifyStrategyId(v string) *SubscriptionForModify {
	s.NotifyStrategyId = &v
	return s
}

func (s *SubscriptionForModify) SetPushingSetting(v *SubscriptionForModifyPushingSetting) *SubscriptionForModify {
	s.PushingSetting = v
	return s
}

func (s *SubscriptionForModify) SetSubscriptionName(v string) *SubscriptionForModify {
	s.SubscriptionName = &v
	return s
}

func (s *SubscriptionForModify) Validate() error {
	return dara.Validate(s)
}

type SubscriptionForModifyPushingSetting struct {
	AlertActionIds   []*string `json:"alertActionIds,omitempty" xml:"alertActionIds,omitempty" type:"Repeated"`
	ResponsePlanId   *string   `json:"responsePlanId,omitempty" xml:"responsePlanId,omitempty"`
	RestoreActionIds []*string `json:"restoreActionIds,omitempty" xml:"restoreActionIds,omitempty" type:"Repeated"`
	TemplateUuid     *string   `json:"templateUuid,omitempty" xml:"templateUuid,omitempty"`
}

func (s SubscriptionForModifyPushingSetting) String() string {
	return dara.Prettify(s)
}

func (s SubscriptionForModifyPushingSetting) GoString() string {
	return s.String()
}

func (s *SubscriptionForModifyPushingSetting) GetAlertActionIds() []*string {
	return s.AlertActionIds
}

func (s *SubscriptionForModifyPushingSetting) GetResponsePlanId() *string {
	return s.ResponsePlanId
}

func (s *SubscriptionForModifyPushingSetting) GetRestoreActionIds() []*string {
	return s.RestoreActionIds
}

func (s *SubscriptionForModifyPushingSetting) GetTemplateUuid() *string {
	return s.TemplateUuid
}

func (s *SubscriptionForModifyPushingSetting) SetAlertActionIds(v []*string) *SubscriptionForModifyPushingSetting {
	s.AlertActionIds = v
	return s
}

func (s *SubscriptionForModifyPushingSetting) SetResponsePlanId(v string) *SubscriptionForModifyPushingSetting {
	s.ResponsePlanId = &v
	return s
}

func (s *SubscriptionForModifyPushingSetting) SetRestoreActionIds(v []*string) *SubscriptionForModifyPushingSetting {
	s.RestoreActionIds = v
	return s
}

func (s *SubscriptionForModifyPushingSetting) SetTemplateUuid(v string) *SubscriptionForModifyPushingSetting {
	s.TemplateUuid = &v
	return s
}

func (s *SubscriptionForModifyPushingSetting) Validate() error {
	return dara.Validate(s)
}
