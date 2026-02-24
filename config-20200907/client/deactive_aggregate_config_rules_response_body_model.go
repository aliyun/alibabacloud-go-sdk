// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeactiveAggregateConfigRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOperateRuleResult(v *DeactiveAggregateConfigRulesResponseBodyOperateRuleResult) *DeactiveAggregateConfigRulesResponseBody
	GetOperateRuleResult() *DeactiveAggregateConfigRulesResponseBodyOperateRuleResult
	SetRequestId(v string) *DeactiveAggregateConfigRulesResponseBody
	GetRequestId() *string
}

type DeactiveAggregateConfigRulesResponseBody struct {
	OperateRuleResult *DeactiveAggregateConfigRulesResponseBodyOperateRuleResult `json:"OperateRuleResult,omitempty" xml:"OperateRuleResult,omitempty" type:"Struct"`
	RequestId         *string                                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeactiveAggregateConfigRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeactiveAggregateConfigRulesResponseBody) GoString() string {
	return s.String()
}

func (s *DeactiveAggregateConfigRulesResponseBody) GetOperateRuleResult() *DeactiveAggregateConfigRulesResponseBodyOperateRuleResult {
	return s.OperateRuleResult
}

func (s *DeactiveAggregateConfigRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeactiveAggregateConfigRulesResponseBody) SetOperateRuleResult(v *DeactiveAggregateConfigRulesResponseBodyOperateRuleResult) *DeactiveAggregateConfigRulesResponseBody {
	s.OperateRuleResult = v
	return s
}

func (s *DeactiveAggregateConfigRulesResponseBody) SetRequestId(v string) *DeactiveAggregateConfigRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeactiveAggregateConfigRulesResponseBody) Validate() error {
	if s.OperateRuleResult != nil {
		if err := s.OperateRuleResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeactiveAggregateConfigRulesResponseBodyOperateRuleResult struct {
	OperateRuleItemList []*DeactiveAggregateConfigRulesResponseBodyOperateRuleResultOperateRuleItemList `json:"OperateRuleItemList,omitempty" xml:"OperateRuleItemList,omitempty" type:"Repeated"`
}

func (s DeactiveAggregateConfigRulesResponseBodyOperateRuleResult) String() string {
	return dara.Prettify(s)
}

func (s DeactiveAggregateConfigRulesResponseBodyOperateRuleResult) GoString() string {
	return s.String()
}

func (s *DeactiveAggregateConfigRulesResponseBodyOperateRuleResult) GetOperateRuleItemList() []*DeactiveAggregateConfigRulesResponseBodyOperateRuleResultOperateRuleItemList {
	return s.OperateRuleItemList
}

func (s *DeactiveAggregateConfigRulesResponseBodyOperateRuleResult) SetOperateRuleItemList(v []*DeactiveAggregateConfigRulesResponseBodyOperateRuleResultOperateRuleItemList) *DeactiveAggregateConfigRulesResponseBodyOperateRuleResult {
	s.OperateRuleItemList = v
	return s
}

func (s *DeactiveAggregateConfigRulesResponseBodyOperateRuleResult) Validate() error {
	if s.OperateRuleItemList != nil {
		for _, item := range s.OperateRuleItemList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DeactiveAggregateConfigRulesResponseBodyOperateRuleResultOperateRuleItemList struct {
	ConfigRuleId *string `json:"ConfigRuleId,omitempty" xml:"ConfigRuleId,omitempty"`
	ErrorCode    *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeactiveAggregateConfigRulesResponseBodyOperateRuleResultOperateRuleItemList) String() string {
	return dara.Prettify(s)
}

func (s DeactiveAggregateConfigRulesResponseBodyOperateRuleResultOperateRuleItemList) GoString() string {
	return s.String()
}

func (s *DeactiveAggregateConfigRulesResponseBodyOperateRuleResultOperateRuleItemList) GetConfigRuleId() *string {
	return s.ConfigRuleId
}

func (s *DeactiveAggregateConfigRulesResponseBodyOperateRuleResultOperateRuleItemList) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DeactiveAggregateConfigRulesResponseBodyOperateRuleResultOperateRuleItemList) GetSuccess() *bool {
	return s.Success
}

func (s *DeactiveAggregateConfigRulesResponseBodyOperateRuleResultOperateRuleItemList) SetConfigRuleId(v string) *DeactiveAggregateConfigRulesResponseBodyOperateRuleResultOperateRuleItemList {
	s.ConfigRuleId = &v
	return s
}

func (s *DeactiveAggregateConfigRulesResponseBodyOperateRuleResultOperateRuleItemList) SetErrorCode(v string) *DeactiveAggregateConfigRulesResponseBodyOperateRuleResultOperateRuleItemList {
	s.ErrorCode = &v
	return s
}

func (s *DeactiveAggregateConfigRulesResponseBodyOperateRuleResultOperateRuleItemList) SetSuccess(v bool) *DeactiveAggregateConfigRulesResponseBodyOperateRuleResultOperateRuleItemList {
	s.Success = &v
	return s
}

func (s *DeactiveAggregateConfigRulesResponseBodyOperateRuleResultOperateRuleItemList) Validate() error {
	return dara.Validate(s)
}
