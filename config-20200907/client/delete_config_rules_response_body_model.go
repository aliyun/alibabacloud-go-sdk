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
	OperateRuleResult *DeleteConfigRulesResponseBodyOperateRuleResult `json:"OperateRuleResult,omitempty" xml:"OperateRuleResult,omitempty" type:"Struct"`
	RequestId         *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	ConfigRuleId *string `json:"ConfigRuleId,omitempty" xml:"ConfigRuleId,omitempty"`
	ErrorCode    *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
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
