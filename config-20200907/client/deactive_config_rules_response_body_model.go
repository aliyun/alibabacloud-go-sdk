// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeactiveConfigRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOperateRuleResult(v *DeactiveConfigRulesResponseBodyOperateRuleResult) *DeactiveConfigRulesResponseBody
	GetOperateRuleResult() *DeactiveConfigRulesResponseBodyOperateRuleResult
	SetRequestId(v string) *DeactiveConfigRulesResponseBody
	GetRequestId() *string
}

type DeactiveConfigRulesResponseBody struct {
	// The results of the operations that are performed to disable the specified rules.
	OperateRuleResult *DeactiveConfigRulesResponseBodyOperateRuleResult `json:"OperateRuleResult,omitempty" xml:"OperateRuleResult,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 54FA74D9-45D4-4CA5-9BE1-97F6EA19AF5B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeactiveConfigRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeactiveConfigRulesResponseBody) GoString() string {
	return s.String()
}

func (s *DeactiveConfigRulesResponseBody) GetOperateRuleResult() *DeactiveConfigRulesResponseBodyOperateRuleResult {
	return s.OperateRuleResult
}

func (s *DeactiveConfigRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeactiveConfigRulesResponseBody) SetOperateRuleResult(v *DeactiveConfigRulesResponseBodyOperateRuleResult) *DeactiveConfigRulesResponseBody {
	s.OperateRuleResult = v
	return s
}

func (s *DeactiveConfigRulesResponseBody) SetRequestId(v string) *DeactiveConfigRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeactiveConfigRulesResponseBody) Validate() error {
	if s.OperateRuleResult != nil {
		if err := s.OperateRuleResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeactiveConfigRulesResponseBodyOperateRuleResult struct {
	// The operations that are performed to disable the rule.
	OperateRuleItemList []*DeactiveConfigRulesResponseBodyOperateRuleResultOperateRuleItemList `json:"OperateRuleItemList,omitempty" xml:"OperateRuleItemList,omitempty" type:"Repeated"`
}

func (s DeactiveConfigRulesResponseBodyOperateRuleResult) String() string {
	return dara.Prettify(s)
}

func (s DeactiveConfigRulesResponseBodyOperateRuleResult) GoString() string {
	return s.String()
}

func (s *DeactiveConfigRulesResponseBodyOperateRuleResult) GetOperateRuleItemList() []*DeactiveConfigRulesResponseBodyOperateRuleResultOperateRuleItemList {
	return s.OperateRuleItemList
}

func (s *DeactiveConfigRulesResponseBodyOperateRuleResult) SetOperateRuleItemList(v []*DeactiveConfigRulesResponseBodyOperateRuleResultOperateRuleItemList) *DeactiveConfigRulesResponseBodyOperateRuleResult {
	s.OperateRuleItemList = v
	return s
}

func (s *DeactiveConfigRulesResponseBodyOperateRuleResult) Validate() error {
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

type DeactiveConfigRulesResponseBodyOperateRuleResultOperateRuleItemList struct {
	// The rule ID.
	//
	// example:
	//
	// cr-19a56457e0d90058****
	ConfigRuleId *string `json:"ConfigRuleId,omitempty" xml:"ConfigRuleId,omitempty"`
	// The error code.
	//
	// 	- If the rule is disabled, no error code is returned.
	//
	// 	- If the rule fails to be disabled, an error code is returned. For more information about error codes, see [Error codes](https://error-center.alibabacloud.com/status/product/Config).
	//
	// example:
	//
	// ConfigRuleCanNotDelete
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeactiveConfigRulesResponseBodyOperateRuleResultOperateRuleItemList) String() string {
	return dara.Prettify(s)
}

func (s DeactiveConfigRulesResponseBodyOperateRuleResultOperateRuleItemList) GoString() string {
	return s.String()
}

func (s *DeactiveConfigRulesResponseBodyOperateRuleResultOperateRuleItemList) GetConfigRuleId() *string {
	return s.ConfigRuleId
}

func (s *DeactiveConfigRulesResponseBodyOperateRuleResultOperateRuleItemList) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DeactiveConfigRulesResponseBodyOperateRuleResultOperateRuleItemList) GetSuccess() *bool {
	return s.Success
}

func (s *DeactiveConfigRulesResponseBodyOperateRuleResultOperateRuleItemList) SetConfigRuleId(v string) *DeactiveConfigRulesResponseBodyOperateRuleResultOperateRuleItemList {
	s.ConfigRuleId = &v
	return s
}

func (s *DeactiveConfigRulesResponseBodyOperateRuleResultOperateRuleItemList) SetErrorCode(v string) *DeactiveConfigRulesResponseBodyOperateRuleResultOperateRuleItemList {
	s.ErrorCode = &v
	return s
}

func (s *DeactiveConfigRulesResponseBodyOperateRuleResultOperateRuleItemList) SetSuccess(v bool) *DeactiveConfigRulesResponseBodyOperateRuleResultOperateRuleItemList {
	s.Success = &v
	return s
}

func (s *DeactiveConfigRulesResponseBodyOperateRuleResultOperateRuleItemList) Validate() error {
	return dara.Validate(s)
}
