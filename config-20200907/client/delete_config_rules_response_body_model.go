// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteConfigRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOperateRuleResult(v *DeleteConfigRulesResponseBodyOperateRuleResult) *DeleteConfigRulesResponseBody
	GetOperateRuleResult() *DeleteConfigRulesResponseBodyOperateRuleResult
	SetRequestId(v string) *DeleteConfigRulesResponseBody
	GetRequestId() *string
}

type DeleteConfigRulesResponseBody struct {
	// The results of the delete operations.
	OperateRuleResult *DeleteConfigRulesResponseBodyOperateRuleResult `json:"OperateRuleResult,omitempty" xml:"OperateRuleResult,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 6721BBD3-F2A6-5349-9051-EE0111036D3F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteConfigRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteConfigRulesResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteConfigRulesResponseBody) GetOperateRuleResult() *DeleteConfigRulesResponseBodyOperateRuleResult {
	return s.OperateRuleResult
}

func (s *DeleteConfigRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteConfigRulesResponseBody) SetOperateRuleResult(v *DeleteConfigRulesResponseBodyOperateRuleResult) *DeleteConfigRulesResponseBody {
	s.OperateRuleResult = v
	return s
}

func (s *DeleteConfigRulesResponseBody) SetRequestId(v string) *DeleteConfigRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteConfigRulesResponseBody) Validate() error {
	if s.OperateRuleResult != nil {
		if err := s.OperateRuleResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteConfigRulesResponseBodyOperateRuleResult struct {
	// The result of the delete operation.
	OperateRuleItemList []*DeleteConfigRulesResponseBodyOperateRuleResultOperateRuleItemList `json:"OperateRuleItemList,omitempty" xml:"OperateRuleItemList,omitempty" type:"Repeated"`
}

func (s DeleteConfigRulesResponseBodyOperateRuleResult) String() string {
	return dara.Prettify(s)
}

func (s DeleteConfigRulesResponseBodyOperateRuleResult) GoString() string {
	return s.String()
}

func (s *DeleteConfigRulesResponseBodyOperateRuleResult) GetOperateRuleItemList() []*DeleteConfigRulesResponseBodyOperateRuleResultOperateRuleItemList {
	return s.OperateRuleItemList
}

func (s *DeleteConfigRulesResponseBodyOperateRuleResult) SetOperateRuleItemList(v []*DeleteConfigRulesResponseBodyOperateRuleResultOperateRuleItemList) *DeleteConfigRulesResponseBodyOperateRuleResult {
	s.OperateRuleItemList = v
	return s
}

func (s *DeleteConfigRulesResponseBodyOperateRuleResult) Validate() error {
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

type DeleteConfigRulesResponseBodyOperateRuleResultOperateRuleItemList struct {
	// The rule ID.
	//
	// example:
	//
	// cr-9908626622af0035****
	ConfigRuleId *string `json:"ConfigRuleId,omitempty" xml:"ConfigRuleId,omitempty"`
	// The error code.
	//
	// 	- If the rule is disabled, no error code is returned.
	//
	// 	- If the rule fails to be disabled, an error code is returned. For more information about error codes, see [Error codes](https://api.alibabacloud.com/document/Config/2020-09-07/errorCode).
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

func (s DeleteConfigRulesResponseBodyOperateRuleResultOperateRuleItemList) String() string {
	return dara.Prettify(s)
}

func (s DeleteConfigRulesResponseBodyOperateRuleResultOperateRuleItemList) GoString() string {
	return s.String()
}

func (s *DeleteConfigRulesResponseBodyOperateRuleResultOperateRuleItemList) GetConfigRuleId() *string {
	return s.ConfigRuleId
}

func (s *DeleteConfigRulesResponseBodyOperateRuleResultOperateRuleItemList) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DeleteConfigRulesResponseBodyOperateRuleResultOperateRuleItemList) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteConfigRulesResponseBodyOperateRuleResultOperateRuleItemList) SetConfigRuleId(v string) *DeleteConfigRulesResponseBodyOperateRuleResultOperateRuleItemList {
	s.ConfigRuleId = &v
	return s
}

func (s *DeleteConfigRulesResponseBodyOperateRuleResultOperateRuleItemList) SetErrorCode(v string) *DeleteConfigRulesResponseBodyOperateRuleResultOperateRuleItemList {
	s.ErrorCode = &v
	return s
}

func (s *DeleteConfigRulesResponseBodyOperateRuleResultOperateRuleItemList) SetSuccess(v bool) *DeleteConfigRulesResponseBodyOperateRuleResultOperateRuleItemList {
	s.Success = &v
	return s
}

func (s *DeleteConfigRulesResponseBodyOperateRuleResultOperateRuleItemList) Validate() error {
	return dara.Validate(s)
}
