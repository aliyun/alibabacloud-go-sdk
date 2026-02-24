// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachAggregateConfigRuleToCompliancePackResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOperateRuleResult(v *DetachAggregateConfigRuleToCompliancePackResponseBodyOperateRuleResult) *DetachAggregateConfigRuleToCompliancePackResponseBody
	GetOperateRuleResult() *DetachAggregateConfigRuleToCompliancePackResponseBodyOperateRuleResult
	SetRequestId(v string) *DetachAggregateConfigRuleToCompliancePackResponseBody
	GetRequestId() *string
}

type DetachAggregateConfigRuleToCompliancePackResponseBody struct {
	OperateRuleResult *DetachAggregateConfigRuleToCompliancePackResponseBodyOperateRuleResult `json:"OperateRuleResult,omitempty" xml:"OperateRuleResult,omitempty" type:"Struct"`
	RequestId         *string                                                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DetachAggregateConfigRuleToCompliancePackResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DetachAggregateConfigRuleToCompliancePackResponseBody) GoString() string {
	return s.String()
}

func (s *DetachAggregateConfigRuleToCompliancePackResponseBody) GetOperateRuleResult() *DetachAggregateConfigRuleToCompliancePackResponseBodyOperateRuleResult {
	return s.OperateRuleResult
}

func (s *DetachAggregateConfigRuleToCompliancePackResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DetachAggregateConfigRuleToCompliancePackResponseBody) SetOperateRuleResult(v *DetachAggregateConfigRuleToCompliancePackResponseBodyOperateRuleResult) *DetachAggregateConfigRuleToCompliancePackResponseBody {
	s.OperateRuleResult = v
	return s
}

func (s *DetachAggregateConfigRuleToCompliancePackResponseBody) SetRequestId(v string) *DetachAggregateConfigRuleToCompliancePackResponseBody {
	s.RequestId = &v
	return s
}

func (s *DetachAggregateConfigRuleToCompliancePackResponseBody) Validate() error {
	if s.OperateRuleResult != nil {
		if err := s.OperateRuleResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DetachAggregateConfigRuleToCompliancePackResponseBodyOperateRuleResult struct {
	OperateRuleItemList []*DetachAggregateConfigRuleToCompliancePackResponseBodyOperateRuleResultOperateRuleItemList `json:"OperateRuleItemList,omitempty" xml:"OperateRuleItemList,omitempty" type:"Repeated"`
}

func (s DetachAggregateConfigRuleToCompliancePackResponseBodyOperateRuleResult) String() string {
	return dara.Prettify(s)
}

func (s DetachAggregateConfigRuleToCompliancePackResponseBodyOperateRuleResult) GoString() string {
	return s.String()
}

func (s *DetachAggregateConfigRuleToCompliancePackResponseBodyOperateRuleResult) GetOperateRuleItemList() []*DetachAggregateConfigRuleToCompliancePackResponseBodyOperateRuleResultOperateRuleItemList {
	return s.OperateRuleItemList
}

func (s *DetachAggregateConfigRuleToCompliancePackResponseBodyOperateRuleResult) SetOperateRuleItemList(v []*DetachAggregateConfigRuleToCompliancePackResponseBodyOperateRuleResultOperateRuleItemList) *DetachAggregateConfigRuleToCompliancePackResponseBodyOperateRuleResult {
	s.OperateRuleItemList = v
	return s
}

func (s *DetachAggregateConfigRuleToCompliancePackResponseBodyOperateRuleResult) Validate() error {
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

type DetachAggregateConfigRuleToCompliancePackResponseBodyOperateRuleResultOperateRuleItemList struct {
	ConfigRuleId *string `json:"ConfigRuleId,omitempty" xml:"ConfigRuleId,omitempty"`
	ErrorCode    *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DetachAggregateConfigRuleToCompliancePackResponseBodyOperateRuleResultOperateRuleItemList) String() string {
	return dara.Prettify(s)
}

func (s DetachAggregateConfigRuleToCompliancePackResponseBodyOperateRuleResultOperateRuleItemList) GoString() string {
	return s.String()
}

func (s *DetachAggregateConfigRuleToCompliancePackResponseBodyOperateRuleResultOperateRuleItemList) GetConfigRuleId() *string {
	return s.ConfigRuleId
}

func (s *DetachAggregateConfigRuleToCompliancePackResponseBodyOperateRuleResultOperateRuleItemList) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DetachAggregateConfigRuleToCompliancePackResponseBodyOperateRuleResultOperateRuleItemList) GetSuccess() *bool {
	return s.Success
}

func (s *DetachAggregateConfigRuleToCompliancePackResponseBodyOperateRuleResultOperateRuleItemList) SetConfigRuleId(v string) *DetachAggregateConfigRuleToCompliancePackResponseBodyOperateRuleResultOperateRuleItemList {
	s.ConfigRuleId = &v
	return s
}

func (s *DetachAggregateConfigRuleToCompliancePackResponseBodyOperateRuleResultOperateRuleItemList) SetErrorCode(v string) *DetachAggregateConfigRuleToCompliancePackResponseBodyOperateRuleResultOperateRuleItemList {
	s.ErrorCode = &v
	return s
}

func (s *DetachAggregateConfigRuleToCompliancePackResponseBodyOperateRuleResultOperateRuleItemList) SetSuccess(v bool) *DetachAggregateConfigRuleToCompliancePackResponseBodyOperateRuleResultOperateRuleItemList {
	s.Success = &v
	return s
}

func (s *DetachAggregateConfigRuleToCompliancePackResponseBodyOperateRuleResultOperateRuleItemList) Validate() error {
	return dara.Validate(s)
}
