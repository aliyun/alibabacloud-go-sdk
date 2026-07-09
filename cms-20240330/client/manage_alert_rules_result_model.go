// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iManageAlertRulesResult interface {
	dara.Model
	String() string
	GoString() string
	SetAlertRule(v *AlertRuleV2) *ManageAlertRulesResult
	GetAlertRule() *AlertRuleV2
	SetDeletedCount(v int32) *ManageAlertRulesResult
	GetDeletedCount() *int32
	SetDeletedUuidList(v []*string) *ManageAlertRulesResult
	GetDeletedUuidList() []*string
	SetUpdatedCount(v int32) *ManageAlertRulesResult
	GetUpdatedCount() *int32
	SetUpdatedUuidList(v []*string) *ManageAlertRulesResult
	GetUpdatedUuidList() []*string
}

type ManageAlertRulesResult struct {
	// The details of the created or updated alert rule.
	AlertRule *AlertRuleV2 `json:"alertRule,omitempty" xml:"alertRule,omitempty"`
	// The number of deleted alert rules.
	DeletedCount *int32 `json:"deletedCount,omitempty" xml:"deletedCount,omitempty"`
	// A list of UUIDs of deleted alert rules.
	DeletedUuidList []*string `json:"deletedUuidList,omitempty" xml:"deletedUuidList,omitempty" type:"Repeated"`
	// The number of updated alert rules.
	UpdatedCount *int32 `json:"updatedCount,omitempty" xml:"updatedCount,omitempty"`
	// A list of UUIDs of updated alert rules.
	UpdatedUuidList []*string `json:"updatedUuidList,omitempty" xml:"updatedUuidList,omitempty" type:"Repeated"`
}

func (s ManageAlertRulesResult) String() string {
	return dara.Prettify(s)
}

func (s ManageAlertRulesResult) GoString() string {
	return s.String()
}

func (s *ManageAlertRulesResult) GetAlertRule() *AlertRuleV2 {
	return s.AlertRule
}

func (s *ManageAlertRulesResult) GetDeletedCount() *int32 {
	return s.DeletedCount
}

func (s *ManageAlertRulesResult) GetDeletedUuidList() []*string {
	return s.DeletedUuidList
}

func (s *ManageAlertRulesResult) GetUpdatedCount() *int32 {
	return s.UpdatedCount
}

func (s *ManageAlertRulesResult) GetUpdatedUuidList() []*string {
	return s.UpdatedUuidList
}

func (s *ManageAlertRulesResult) SetAlertRule(v *AlertRuleV2) *ManageAlertRulesResult {
	s.AlertRule = v
	return s
}

func (s *ManageAlertRulesResult) SetDeletedCount(v int32) *ManageAlertRulesResult {
	s.DeletedCount = &v
	return s
}

func (s *ManageAlertRulesResult) SetDeletedUuidList(v []*string) *ManageAlertRulesResult {
	s.DeletedUuidList = v
	return s
}

func (s *ManageAlertRulesResult) SetUpdatedCount(v int32) *ManageAlertRulesResult {
	s.UpdatedCount = &v
	return s
}

func (s *ManageAlertRulesResult) SetUpdatedUuidList(v []*string) *ManageAlertRulesResult {
	s.UpdatedUuidList = v
	return s
}

func (s *ManageAlertRulesResult) Validate() error {
	if s.AlertRule != nil {
		if err := s.AlertRule.Validate(); err != nil {
			return err
		}
	}
	return nil
}
