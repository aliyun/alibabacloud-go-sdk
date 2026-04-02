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
}

type ManageAlertRulesResult struct {
	AlertRule *AlertRuleV2 `json:"alertRule,omitempty" xml:"alertRule,omitempty"`
	// 成功删除的规则数量
	DeletedCount *int32 `json:"deletedCount,omitempty" xml:"deletedCount,omitempty"`
	// 成功删除的规则 UUID 列表
	DeletedUuidList []*string `json:"deletedUuidList,omitempty" xml:"deletedUuidList,omitempty" type:"Repeated"`
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

func (s *ManageAlertRulesResult) Validate() error {
	if s.AlertRule != nil {
		if err := s.AlertRule.Validate(); err != nil {
			return err
		}
	}
	return nil
}
