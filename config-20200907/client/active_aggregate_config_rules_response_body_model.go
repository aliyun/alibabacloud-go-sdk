// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iActiveAggregateConfigRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOperateRuleResult(v *ActiveAggregateConfigRulesResponseBodyOperateRuleResult) *ActiveAggregateConfigRulesResponseBody
	GetOperateRuleResult() *ActiveAggregateConfigRulesResponseBodyOperateRuleResult
	SetRequestId(v string) *ActiveAggregateConfigRulesResponseBody
	GetRequestId() *string
}

type ActiveAggregateConfigRulesResponseBody struct {
	OperateRuleResult *ActiveAggregateConfigRulesResponseBodyOperateRuleResult `json:"OperateRuleResult,omitempty" xml:"OperateRuleResult,omitempty" type:"Struct"`
	RequestId         *string                                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ActiveAggregateConfigRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ActiveAggregateConfigRulesResponseBody) GoString() string {
	return s.String()
}

func (s *ActiveAggregateConfigRulesResponseBody) GetOperateRuleResult() *ActiveAggregateConfigRulesResponseBodyOperateRuleResult {
	return s.OperateRuleResult
}

func (s *ActiveAggregateConfigRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ActiveAggregateConfigRulesResponseBody) SetOperateRuleResult(v *ActiveAggregateConfigRulesResponseBodyOperateRuleResult) *ActiveAggregateConfigRulesResponseBody {
	s.OperateRuleResult = v
	return s
}

func (s *ActiveAggregateConfigRulesResponseBody) SetRequestId(v string) *ActiveAggregateConfigRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ActiveAggregateConfigRulesResponseBody) Validate() error {
	if s.OperateRuleResult != nil {
		if err := s.OperateRuleResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ActiveAggregateConfigRulesResponseBodyOperateRuleResult struct {
	OperateRuleItemList []*ActiveAggregateConfigRulesResponseBodyOperateRuleResultOperateRuleItemList `json:"OperateRuleItemList,omitempty" xml:"OperateRuleItemList,omitempty" type:"Repeated"`
}

func (s ActiveAggregateConfigRulesResponseBodyOperateRuleResult) String() string {
	return dara.Prettify(s)
}

func (s ActiveAggregateConfigRulesResponseBodyOperateRuleResult) GoString() string {
	return s.String()
}

func (s *ActiveAggregateConfigRulesResponseBodyOperateRuleResult) GetOperateRuleItemList() []*ActiveAggregateConfigRulesResponseBodyOperateRuleResultOperateRuleItemList {
	return s.OperateRuleItemList
}

func (s *ActiveAggregateConfigRulesResponseBodyOperateRuleResult) SetOperateRuleItemList(v []*ActiveAggregateConfigRulesResponseBodyOperateRuleResultOperateRuleItemList) *ActiveAggregateConfigRulesResponseBodyOperateRuleResult {
	s.OperateRuleItemList = v
	return s
}

func (s *ActiveAggregateConfigRulesResponseBodyOperateRuleResult) Validate() error {
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

type ActiveAggregateConfigRulesResponseBodyOperateRuleResultOperateRuleItemList struct {
	ConfigRuleId *string `json:"ConfigRuleId,omitempty" xml:"ConfigRuleId,omitempty"`
	ErrorCode    *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ActiveAggregateConfigRulesResponseBodyOperateRuleResultOperateRuleItemList) String() string {
	return dara.Prettify(s)
}

func (s ActiveAggregateConfigRulesResponseBodyOperateRuleResultOperateRuleItemList) GoString() string {
	return s.String()
}

func (s *ActiveAggregateConfigRulesResponseBodyOperateRuleResultOperateRuleItemList) GetConfigRuleId() *string {
	return s.ConfigRuleId
}

func (s *ActiveAggregateConfigRulesResponseBodyOperateRuleResultOperateRuleItemList) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ActiveAggregateConfigRulesResponseBodyOperateRuleResultOperateRuleItemList) GetSuccess() *bool {
	return s.Success
}

func (s *ActiveAggregateConfigRulesResponseBodyOperateRuleResultOperateRuleItemList) SetConfigRuleId(v string) *ActiveAggregateConfigRulesResponseBodyOperateRuleResultOperateRuleItemList {
	s.ConfigRuleId = &v
	return s
}

func (s *ActiveAggregateConfigRulesResponseBodyOperateRuleResultOperateRuleItemList) SetErrorCode(v string) *ActiveAggregateConfigRulesResponseBodyOperateRuleResultOperateRuleItemList {
	s.ErrorCode = &v
	return s
}

func (s *ActiveAggregateConfigRulesResponseBodyOperateRuleResultOperateRuleItemList) SetSuccess(v bool) *ActiveAggregateConfigRulesResponseBodyOperateRuleResultOperateRuleItemList {
	s.Success = &v
	return s
}

func (s *ActiveAggregateConfigRulesResponseBodyOperateRuleResultOperateRuleItemList) Validate() error {
	return dara.Validate(s)
}
