// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachConfigRuleToCompliancePackResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOperateRuleResult(v *DetachConfigRuleToCompliancePackResponseBodyOperateRuleResult) *DetachConfigRuleToCompliancePackResponseBody
	GetOperateRuleResult() *DetachConfigRuleToCompliancePackResponseBodyOperateRuleResult
	SetRequestId(v string) *DetachConfigRuleToCompliancePackResponseBody
	GetRequestId() *string
}

type DetachConfigRuleToCompliancePackResponseBody struct {
	// The results of the operations to remove one or more rules.
	OperateRuleResult *DetachConfigRuleToCompliancePackResponseBodyOperateRuleResult `json:"OperateRuleResult,omitempty" xml:"OperateRuleResult,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 1AC7AED1-172F-42AE-9C12-295BC2ADB12A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DetachConfigRuleToCompliancePackResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DetachConfigRuleToCompliancePackResponseBody) GoString() string {
	return s.String()
}

func (s *DetachConfigRuleToCompliancePackResponseBody) GetOperateRuleResult() *DetachConfigRuleToCompliancePackResponseBodyOperateRuleResult {
	return s.OperateRuleResult
}

func (s *DetachConfigRuleToCompliancePackResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DetachConfigRuleToCompliancePackResponseBody) SetOperateRuleResult(v *DetachConfigRuleToCompliancePackResponseBodyOperateRuleResult) *DetachConfigRuleToCompliancePackResponseBody {
	s.OperateRuleResult = v
	return s
}

func (s *DetachConfigRuleToCompliancePackResponseBody) SetRequestId(v string) *DetachConfigRuleToCompliancePackResponseBody {
	s.RequestId = &v
	return s
}

func (s *DetachConfigRuleToCompliancePackResponseBody) Validate() error {
	if s.OperateRuleResult != nil {
		if err := s.OperateRuleResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DetachConfigRuleToCompliancePackResponseBodyOperateRuleResult struct {
	// The result of the operation to remove the rule.
	OperateRuleItemList []*DetachConfigRuleToCompliancePackResponseBodyOperateRuleResultOperateRuleItemList `json:"OperateRuleItemList,omitempty" xml:"OperateRuleItemList,omitempty" type:"Repeated"`
}

func (s DetachConfigRuleToCompliancePackResponseBodyOperateRuleResult) String() string {
	return dara.Prettify(s)
}

func (s DetachConfigRuleToCompliancePackResponseBodyOperateRuleResult) GoString() string {
	return s.String()
}

func (s *DetachConfigRuleToCompliancePackResponseBodyOperateRuleResult) GetOperateRuleItemList() []*DetachConfigRuleToCompliancePackResponseBodyOperateRuleResultOperateRuleItemList {
	return s.OperateRuleItemList
}

func (s *DetachConfigRuleToCompliancePackResponseBodyOperateRuleResult) SetOperateRuleItemList(v []*DetachConfigRuleToCompliancePackResponseBodyOperateRuleResultOperateRuleItemList) *DetachConfigRuleToCompliancePackResponseBodyOperateRuleResult {
	s.OperateRuleItemList = v
	return s
}

func (s *DetachConfigRuleToCompliancePackResponseBodyOperateRuleResult) Validate() error {
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

type DetachConfigRuleToCompliancePackResponseBodyOperateRuleResultOperateRuleItemList struct {
	// The rule ID.
	//
	// example:
	//
	// cr-6cc4626622af00e7****
	ConfigRuleId *string `json:"ConfigRuleId,omitempty" xml:"ConfigRuleId,omitempty"`
	// The error code returned.
	//
	// 	- This parameter is empty if the rule is removed from the compliance package.
	//
	// 	- An error code is returned if the rule fails to be removed from the compliance package. For more information about error codes, see [Error codes](https://error-center.alibabacloud.com/status/product/Config).
	//
	// example:
	//
	// ConfigRuleNotExists
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true: The request was successful.
	//
	// 	- false: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DetachConfigRuleToCompliancePackResponseBodyOperateRuleResultOperateRuleItemList) String() string {
	return dara.Prettify(s)
}

func (s DetachConfigRuleToCompliancePackResponseBodyOperateRuleResultOperateRuleItemList) GoString() string {
	return s.String()
}

func (s *DetachConfigRuleToCompliancePackResponseBodyOperateRuleResultOperateRuleItemList) GetConfigRuleId() *string {
	return s.ConfigRuleId
}

func (s *DetachConfigRuleToCompliancePackResponseBodyOperateRuleResultOperateRuleItemList) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DetachConfigRuleToCompliancePackResponseBodyOperateRuleResultOperateRuleItemList) GetSuccess() *bool {
	return s.Success
}

func (s *DetachConfigRuleToCompliancePackResponseBodyOperateRuleResultOperateRuleItemList) SetConfigRuleId(v string) *DetachConfigRuleToCompliancePackResponseBodyOperateRuleResultOperateRuleItemList {
	s.ConfigRuleId = &v
	return s
}

func (s *DetachConfigRuleToCompliancePackResponseBodyOperateRuleResultOperateRuleItemList) SetErrorCode(v string) *DetachConfigRuleToCompliancePackResponseBodyOperateRuleResultOperateRuleItemList {
	s.ErrorCode = &v
	return s
}

func (s *DetachConfigRuleToCompliancePackResponseBodyOperateRuleResultOperateRuleItemList) SetSuccess(v bool) *DetachConfigRuleToCompliancePackResponseBodyOperateRuleResultOperateRuleItemList {
	s.Success = &v
	return s
}

func (s *DetachConfigRuleToCompliancePackResponseBodyOperateRuleResultOperateRuleItemList) Validate() error {
	return dara.Validate(s)
}
