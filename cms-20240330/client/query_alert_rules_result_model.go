// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAlertRulesResult interface {
	dara.Model
	String() string
	GoString() string
	SetAlertRules(v []*AlertRuleV2) *QueryAlertRulesResult
	GetAlertRules() []*AlertRuleV2
	SetTotalCount(v int64) *QueryAlertRulesResult
	GetTotalCount() *int64
}

type QueryAlertRulesResult struct {
	// The list of alert rules returned by the query. Each element contains the complete configuration information of an alert rule.
	//
	// example:
	//
	// [{"uuid":"a1b2c3d4-e5f6-7890-abcd-ef1234567890","displayName":"CPU usage alert","status":"OK"}]
	AlertRules []*AlertRuleV2 `json:"alertRules,omitempty" xml:"alertRules,omitempty" type:"Repeated"`
	// The total number of alert rules that match the query conditions.
	//
	// example:
	//
	// 5
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s QueryAlertRulesResult) String() string {
	return dara.Prettify(s)
}

func (s QueryAlertRulesResult) GoString() string {
	return s.String()
}

func (s *QueryAlertRulesResult) GetAlertRules() []*AlertRuleV2 {
	return s.AlertRules
}

func (s *QueryAlertRulesResult) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *QueryAlertRulesResult) SetAlertRules(v []*AlertRuleV2) *QueryAlertRulesResult {
	s.AlertRules = v
	return s
}

func (s *QueryAlertRulesResult) SetTotalCount(v int64) *QueryAlertRulesResult {
	s.TotalCount = &v
	return s
}

func (s *QueryAlertRulesResult) Validate() error {
	if s.AlertRules != nil {
		for _, item := range s.AlertRules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
