// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAggregateConfigRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOperateRuleResult(v *DeleteAggregateConfigRulesResponseBodyOperateRuleResult) *DeleteAggregateConfigRulesResponseBody
	GetOperateRuleResult() *DeleteAggregateConfigRulesResponseBodyOperateRuleResult
	SetRequestId(v string) *DeleteAggregateConfigRulesResponseBody
	GetRequestId() *string
}

type DeleteAggregateConfigRulesResponseBody struct {
	// The results of the delete operations.
	OperateRuleResult *DeleteAggregateConfigRulesResponseBodyOperateRuleResult `json:"OperateRuleResult,omitempty" xml:"OperateRuleResult,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 6EC7AED1-172F-42AE-9C12-295BC2ADB751
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteAggregateConfigRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteAggregateConfigRulesResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAggregateConfigRulesResponseBody) GetOperateRuleResult() *DeleteAggregateConfigRulesResponseBodyOperateRuleResult {
	return s.OperateRuleResult
}

func (s *DeleteAggregateConfigRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteAggregateConfigRulesResponseBody) SetOperateRuleResult(v *DeleteAggregateConfigRulesResponseBodyOperateRuleResult) *DeleteAggregateConfigRulesResponseBody {
	s.OperateRuleResult = v
	return s
}

func (s *DeleteAggregateConfigRulesResponseBody) SetRequestId(v string) *DeleteAggregateConfigRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAggregateConfigRulesResponseBody) Validate() error {
	if s.OperateRuleResult != nil {
		if err := s.OperateRuleResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteAggregateConfigRulesResponseBodyOperateRuleResult struct {
	// The details of the result.
	OperateRuleItemList []*DeleteAggregateConfigRulesResponseBodyOperateRuleResultOperateRuleItemList `json:"OperateRuleItemList,omitempty" xml:"OperateRuleItemList,omitempty" type:"Repeated"`
}

func (s DeleteAggregateConfigRulesResponseBodyOperateRuleResult) String() string {
	return dara.Prettify(s)
}

func (s DeleteAggregateConfigRulesResponseBodyOperateRuleResult) GoString() string {
	return s.String()
}

func (s *DeleteAggregateConfigRulesResponseBodyOperateRuleResult) GetOperateRuleItemList() []*DeleteAggregateConfigRulesResponseBodyOperateRuleResultOperateRuleItemList {
	return s.OperateRuleItemList
}

func (s *DeleteAggregateConfigRulesResponseBodyOperateRuleResult) SetOperateRuleItemList(v []*DeleteAggregateConfigRulesResponseBodyOperateRuleResultOperateRuleItemList) *DeleteAggregateConfigRulesResponseBodyOperateRuleResult {
	s.OperateRuleItemList = v
	return s
}

func (s *DeleteAggregateConfigRulesResponseBodyOperateRuleResult) Validate() error {
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

type DeleteAggregateConfigRulesResponseBodyOperateRuleResultOperateRuleItemList struct {
	// The rule ID.
	//
	// example:
	//
	// cr-4e3d626622af0080****
	ConfigRuleId *string `json:"ConfigRuleId,omitempty" xml:"ConfigRuleId,omitempty"`
	// The error code returned.
	//
	// 	- If the rule is deleted, no error code is returned.
	//
	// 	- If the rule fails to be deleted, an error code is returned. For more information about error codes, see [Error codes](https://error-center.alibabacloud.com/status/product/Config).
	//
	// example:
	//
	// ConfigRuleCanNotDelete
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true: The request was successful.
	//
	// 	- false: The request failed.
	//
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteAggregateConfigRulesResponseBodyOperateRuleResultOperateRuleItemList) String() string {
	return dara.Prettify(s)
}

func (s DeleteAggregateConfigRulesResponseBodyOperateRuleResultOperateRuleItemList) GoString() string {
	return s.String()
}

func (s *DeleteAggregateConfigRulesResponseBodyOperateRuleResultOperateRuleItemList) GetConfigRuleId() *string {
	return s.ConfigRuleId
}

func (s *DeleteAggregateConfigRulesResponseBodyOperateRuleResultOperateRuleItemList) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DeleteAggregateConfigRulesResponseBodyOperateRuleResultOperateRuleItemList) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteAggregateConfigRulesResponseBodyOperateRuleResultOperateRuleItemList) SetConfigRuleId(v string) *DeleteAggregateConfigRulesResponseBodyOperateRuleResultOperateRuleItemList {
	s.ConfigRuleId = &v
	return s
}

func (s *DeleteAggregateConfigRulesResponseBodyOperateRuleResultOperateRuleItemList) SetErrorCode(v string) *DeleteAggregateConfigRulesResponseBodyOperateRuleResultOperateRuleItemList {
	s.ErrorCode = &v
	return s
}

func (s *DeleteAggregateConfigRulesResponseBodyOperateRuleResultOperateRuleItemList) SetSuccess(v bool) *DeleteAggregateConfigRulesResponseBodyOperateRuleResultOperateRuleItemList {
	s.Success = &v
	return s
}

func (s *DeleteAggregateConfigRulesResponseBodyOperateRuleResultOperateRuleItemList) Validate() error {
	return dara.Validate(s)
}
