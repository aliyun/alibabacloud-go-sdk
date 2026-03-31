// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachAggregateConfigRuleToCompliancePackResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOperateRuleResult(v *AttachAggregateConfigRuleToCompliancePackResponseBodyOperateRuleResult) *AttachAggregateConfigRuleToCompliancePackResponseBody
	GetOperateRuleResult() *AttachAggregateConfigRuleToCompliancePackResponseBodyOperateRuleResult
	SetRequestId(v string) *AttachAggregateConfigRuleToCompliancePackResponseBody
	GetRequestId() *string
}

type AttachAggregateConfigRuleToCompliancePackResponseBody struct {
	// The results of the operations to add one or more rules.
	OperateRuleResult *AttachAggregateConfigRuleToCompliancePackResponseBodyOperateRuleResult `json:"OperateRuleResult,omitempty" xml:"OperateRuleResult,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// DE72B7B5-D0EA-15E6-A359-EDECBB9BDFA3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AttachAggregateConfigRuleToCompliancePackResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AttachAggregateConfigRuleToCompliancePackResponseBody) GoString() string {
	return s.String()
}

func (s *AttachAggregateConfigRuleToCompliancePackResponseBody) GetOperateRuleResult() *AttachAggregateConfigRuleToCompliancePackResponseBodyOperateRuleResult {
	return s.OperateRuleResult
}

func (s *AttachAggregateConfigRuleToCompliancePackResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AttachAggregateConfigRuleToCompliancePackResponseBody) SetOperateRuleResult(v *AttachAggregateConfigRuleToCompliancePackResponseBodyOperateRuleResult) *AttachAggregateConfigRuleToCompliancePackResponseBody {
	s.OperateRuleResult = v
	return s
}

func (s *AttachAggregateConfigRuleToCompliancePackResponseBody) SetRequestId(v string) *AttachAggregateConfigRuleToCompliancePackResponseBody {
	s.RequestId = &v
	return s
}

func (s *AttachAggregateConfigRuleToCompliancePackResponseBody) Validate() error {
	if s.OperateRuleResult != nil {
		if err := s.OperateRuleResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AttachAggregateConfigRuleToCompliancePackResponseBodyOperateRuleResult struct {
	// The result of the operation to add the rule.
	OperateRuleItemList []*AttachAggregateConfigRuleToCompliancePackResponseBodyOperateRuleResultOperateRuleItemList `json:"OperateRuleItemList,omitempty" xml:"OperateRuleItemList,omitempty" type:"Repeated"`
}

func (s AttachAggregateConfigRuleToCompliancePackResponseBodyOperateRuleResult) String() string {
	return dara.Prettify(s)
}

func (s AttachAggregateConfigRuleToCompliancePackResponseBodyOperateRuleResult) GoString() string {
	return s.String()
}

func (s *AttachAggregateConfigRuleToCompliancePackResponseBodyOperateRuleResult) GetOperateRuleItemList() []*AttachAggregateConfigRuleToCompliancePackResponseBodyOperateRuleResultOperateRuleItemList {
	return s.OperateRuleItemList
}

func (s *AttachAggregateConfigRuleToCompliancePackResponseBodyOperateRuleResult) SetOperateRuleItemList(v []*AttachAggregateConfigRuleToCompliancePackResponseBodyOperateRuleResultOperateRuleItemList) *AttachAggregateConfigRuleToCompliancePackResponseBodyOperateRuleResult {
	s.OperateRuleItemList = v
	return s
}

func (s *AttachAggregateConfigRuleToCompliancePackResponseBodyOperateRuleResult) Validate() error {
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

type AttachAggregateConfigRuleToCompliancePackResponseBodyOperateRuleResultOperateRuleItemList struct {
	// The rule ID.
	//
	// example:
	//
	// cr-a124626622af00e7****
	ConfigRuleId *string `json:"ConfigRuleId,omitempty" xml:"ConfigRuleId,omitempty"`
	// The error code.
	//
	// 	- This parameter is empty if the rule is added to the compliance package.
	//
	// 	- An error code is returned if the rule fails to be added to the compliance package. For more information about error codes, see [Error codes](https://error-center.alibabacloud.com/status/product/Config).
	//
	// example:
	//
	// ConfigRuleNotExists
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AttachAggregateConfigRuleToCompliancePackResponseBodyOperateRuleResultOperateRuleItemList) String() string {
	return dara.Prettify(s)
}

func (s AttachAggregateConfigRuleToCompliancePackResponseBodyOperateRuleResultOperateRuleItemList) GoString() string {
	return s.String()
}

func (s *AttachAggregateConfigRuleToCompliancePackResponseBodyOperateRuleResultOperateRuleItemList) GetConfigRuleId() *string {
	return s.ConfigRuleId
}

func (s *AttachAggregateConfigRuleToCompliancePackResponseBodyOperateRuleResultOperateRuleItemList) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *AttachAggregateConfigRuleToCompliancePackResponseBodyOperateRuleResultOperateRuleItemList) GetSuccess() *bool {
	return s.Success
}

func (s *AttachAggregateConfigRuleToCompliancePackResponseBodyOperateRuleResultOperateRuleItemList) SetConfigRuleId(v string) *AttachAggregateConfigRuleToCompliancePackResponseBodyOperateRuleResultOperateRuleItemList {
	s.ConfigRuleId = &v
	return s
}

func (s *AttachAggregateConfigRuleToCompliancePackResponseBodyOperateRuleResultOperateRuleItemList) SetErrorCode(v string) *AttachAggregateConfigRuleToCompliancePackResponseBodyOperateRuleResultOperateRuleItemList {
	s.ErrorCode = &v
	return s
}

func (s *AttachAggregateConfigRuleToCompliancePackResponseBodyOperateRuleResultOperateRuleItemList) SetSuccess(v bool) *AttachAggregateConfigRuleToCompliancePackResponseBodyOperateRuleResultOperateRuleItemList {
	s.Success = &v
	return s
}

func (s *AttachAggregateConfigRuleToCompliancePackResponseBodyOperateRuleResultOperateRuleItemList) Validate() error {
	return dara.Validate(s)
}
