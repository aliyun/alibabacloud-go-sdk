// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iActiveConfigRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOperateRuleResult(v *ActiveConfigRulesResponseBodyOperateRuleResult) *ActiveConfigRulesResponseBody
	GetOperateRuleResult() *ActiveConfigRulesResponseBodyOperateRuleResult
	SetRequestId(v string) *ActiveConfigRulesResponseBody
	GetRequestId() *string
}

type ActiveConfigRulesResponseBody struct {
	// The returned results.
	OperateRuleResult *ActiveConfigRulesResponseBodyOperateRuleResult `json:"OperateRuleResult,omitempty" xml:"OperateRuleResult,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 61C1A88F-D163-40DF-84A6-F200229F37B2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ActiveConfigRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ActiveConfigRulesResponseBody) GoString() string {
	return s.String()
}

func (s *ActiveConfigRulesResponseBody) GetOperateRuleResult() *ActiveConfigRulesResponseBodyOperateRuleResult {
	return s.OperateRuleResult
}

func (s *ActiveConfigRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ActiveConfigRulesResponseBody) SetOperateRuleResult(v *ActiveConfigRulesResponseBodyOperateRuleResult) *ActiveConfigRulesResponseBody {
	s.OperateRuleResult = v
	return s
}

func (s *ActiveConfigRulesResponseBody) SetRequestId(v string) *ActiveConfigRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ActiveConfigRulesResponseBody) Validate() error {
	if s.OperateRuleResult != nil {
		if err := s.OperateRuleResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ActiveConfigRulesResponseBodyOperateRuleResult struct {
	// The returned results.
	OperateRuleItemList []*ActiveConfigRulesResponseBodyOperateRuleResultOperateRuleItemList `json:"OperateRuleItemList,omitempty" xml:"OperateRuleItemList,omitempty" type:"Repeated"`
}

func (s ActiveConfigRulesResponseBodyOperateRuleResult) String() string {
	return dara.Prettify(s)
}

func (s ActiveConfigRulesResponseBodyOperateRuleResult) GoString() string {
	return s.String()
}

func (s *ActiveConfigRulesResponseBodyOperateRuleResult) GetOperateRuleItemList() []*ActiveConfigRulesResponseBodyOperateRuleResultOperateRuleItemList {
	return s.OperateRuleItemList
}

func (s *ActiveConfigRulesResponseBodyOperateRuleResult) SetOperateRuleItemList(v []*ActiveConfigRulesResponseBodyOperateRuleResultOperateRuleItemList) *ActiveConfigRulesResponseBodyOperateRuleResult {
	s.OperateRuleItemList = v
	return s
}

func (s *ActiveConfigRulesResponseBodyOperateRuleResult) Validate() error {
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

type ActiveConfigRulesResponseBodyOperateRuleResultOperateRuleItemList struct {
	// The rule ID.
	//
	// example:
	//
	// cr-2da35180a8d1008e****
	ConfigRuleId *string `json:"ConfigRuleId,omitempty" xml:"ConfigRuleId,omitempty"`
	// The error code.
	//
	// 	- If the rule is enabled, no error code is returned.
	//
	// 	- If the rule fails to be enabled, an error code is returned. For more information about error codes, see [Error codes](https://next.api.aliyun.com/document/Config/2020-09-07/errorCode).
	//
	// example:
	//
	// ConfigRuleStatusNotInActive
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// Indicates whether the rule is enabled. Valid values:
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

func (s ActiveConfigRulesResponseBodyOperateRuleResultOperateRuleItemList) String() string {
	return dara.Prettify(s)
}

func (s ActiveConfigRulesResponseBodyOperateRuleResultOperateRuleItemList) GoString() string {
	return s.String()
}

func (s *ActiveConfigRulesResponseBodyOperateRuleResultOperateRuleItemList) GetConfigRuleId() *string {
	return s.ConfigRuleId
}

func (s *ActiveConfigRulesResponseBodyOperateRuleResultOperateRuleItemList) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ActiveConfigRulesResponseBodyOperateRuleResultOperateRuleItemList) GetSuccess() *bool {
	return s.Success
}

func (s *ActiveConfigRulesResponseBodyOperateRuleResultOperateRuleItemList) SetConfigRuleId(v string) *ActiveConfigRulesResponseBodyOperateRuleResultOperateRuleItemList {
	s.ConfigRuleId = &v
	return s
}

func (s *ActiveConfigRulesResponseBodyOperateRuleResultOperateRuleItemList) SetErrorCode(v string) *ActiveConfigRulesResponseBodyOperateRuleResultOperateRuleItemList {
	s.ErrorCode = &v
	return s
}

func (s *ActiveConfigRulesResponseBodyOperateRuleResultOperateRuleItemList) SetSuccess(v bool) *ActiveConfigRulesResponseBodyOperateRuleResultOperateRuleItemList {
	s.Success = &v
	return s
}

func (s *ActiveConfigRulesResponseBodyOperateRuleResultOperateRuleItemList) Validate() error {
	return dara.Validate(s)
}
