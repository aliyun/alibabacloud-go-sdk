// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperateSwitchStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRuleId(v int64) *OperateSwitchStatusRequest
	GetRuleId() *int64
	SetStatus(v string) *OperateSwitchStatusRequest
	GetStatus() *string
}

type OperateSwitchStatusRequest struct {
	// The ID of the rule.
	//
	// >  You can call the ListContainerWebDefenseRule operation to query the IDs of rules.
	//
	// This parameter is required.
	//
	// example:
	//
	// 900001
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The status of the rule. Valid values: on and off.
	//
	// This parameter is required.
	//
	// example:
	//
	// on
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s OperateSwitchStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s OperateSwitchStatusRequest) GoString() string {
	return s.String()
}

func (s *OperateSwitchStatusRequest) GetRuleId() *int64 {
	return s.RuleId
}

func (s *OperateSwitchStatusRequest) GetStatus() *string {
	return s.Status
}

func (s *OperateSwitchStatusRequest) SetRuleId(v int64) *OperateSwitchStatusRequest {
	s.RuleId = &v
	return s
}

func (s *OperateSwitchStatusRequest) SetStatus(v string) *OperateSwitchStatusRequest {
	s.Status = &v
	return s
}

func (s *OperateSwitchStatusRequest) Validate() error {
	return dara.Validate(s)
}
