// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachConfigRuleToCompliancePackResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOperateRuleResult(v *AttachConfigRuleToCompliancePackResponseBodyOperateRuleResult) *AttachConfigRuleToCompliancePackResponseBody
	GetOperateRuleResult() *AttachConfigRuleToCompliancePackResponseBodyOperateRuleResult
	SetRequestId(v string) *AttachConfigRuleToCompliancePackResponseBody
	GetRequestId() *string
}

type AttachConfigRuleToCompliancePackResponseBody struct {
	OperateRuleResult *AttachConfigRuleToCompliancePackResponseBodyOperateRuleResult `json:"OperateRuleResult,omitempty" xml:"OperateRuleResult,omitempty" type:"Struct"`
	RequestId         *string                                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AttachConfigRuleToCompliancePackResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AttachConfigRuleToCompliancePackResponseBody) GoString() string {
	return s.String()
}

func (s *AttachConfigRuleToCompliancePackResponseBody) GetOperateRuleResult() *AttachConfigRuleToCompliancePackResponseBodyOperateRuleResult {
	return s.OperateRuleResult
}

func (s *AttachConfigRuleToCompliancePackResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AttachConfigRuleToCompliancePackResponseBody) SetOperateRuleResult(v *AttachConfigRuleToCompliancePackResponseBodyOperateRuleResult) *AttachConfigRuleToCompliancePackResponseBody {
	s.OperateRuleResult = v
	return s
}

func (s *AttachConfigRuleToCompliancePackResponseBody) SetRequestId(v string) *AttachConfigRuleToCompliancePackResponseBody {
	s.RequestId = &v
	return s
}

func (s *AttachConfigRuleToCompliancePackResponseBody) Validate() error {
	if s.OperateRuleResult != nil {
		if err := s.OperateRuleResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AttachConfigRuleToCompliancePackResponseBodyOperateRuleResult struct {
	OperateRuleItemList []*AttachConfigRuleToCompliancePackResponseBodyOperateRuleResultOperateRuleItemList `json:"OperateRuleItemList,omitempty" xml:"OperateRuleItemList,omitempty" type:"Repeated"`
}

func (s AttachConfigRuleToCompliancePackResponseBodyOperateRuleResult) String() string {
	return dara.Prettify(s)
}

func (s AttachConfigRuleToCompliancePackResponseBodyOperateRuleResult) GoString() string {
	return s.String()
}

func (s *AttachConfigRuleToCompliancePackResponseBodyOperateRuleResult) GetOperateRuleItemList() []*AttachConfigRuleToCompliancePackResponseBodyOperateRuleResultOperateRuleItemList {
	return s.OperateRuleItemList
}

func (s *AttachConfigRuleToCompliancePackResponseBodyOperateRuleResult) SetOperateRuleItemList(v []*AttachConfigRuleToCompliancePackResponseBodyOperateRuleResultOperateRuleItemList) *AttachConfigRuleToCompliancePackResponseBodyOperateRuleResult {
	s.OperateRuleItemList = v
	return s
}

func (s *AttachConfigRuleToCompliancePackResponseBodyOperateRuleResult) Validate() error {
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

type AttachConfigRuleToCompliancePackResponseBodyOperateRuleResultOperateRuleItemList struct {
	ConfigRuleId *string `json:"ConfigRuleId,omitempty" xml:"ConfigRuleId,omitempty"`
	ErrorCode    *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AttachConfigRuleToCompliancePackResponseBodyOperateRuleResultOperateRuleItemList) String() string {
	return dara.Prettify(s)
}

func (s AttachConfigRuleToCompliancePackResponseBodyOperateRuleResultOperateRuleItemList) GoString() string {
	return s.String()
}

func (s *AttachConfigRuleToCompliancePackResponseBodyOperateRuleResultOperateRuleItemList) GetConfigRuleId() *string {
	return s.ConfigRuleId
}

func (s *AttachConfigRuleToCompliancePackResponseBodyOperateRuleResultOperateRuleItemList) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *AttachConfigRuleToCompliancePackResponseBodyOperateRuleResultOperateRuleItemList) GetSuccess() *bool {
	return s.Success
}

func (s *AttachConfigRuleToCompliancePackResponseBodyOperateRuleResultOperateRuleItemList) SetConfigRuleId(v string) *AttachConfigRuleToCompliancePackResponseBodyOperateRuleResultOperateRuleItemList {
	s.ConfigRuleId = &v
	return s
}

func (s *AttachConfigRuleToCompliancePackResponseBodyOperateRuleResultOperateRuleItemList) SetErrorCode(v string) *AttachConfigRuleToCompliancePackResponseBodyOperateRuleResultOperateRuleItemList {
	s.ErrorCode = &v
	return s
}

func (s *AttachConfigRuleToCompliancePackResponseBodyOperateRuleResultOperateRuleItemList) SetSuccess(v bool) *AttachConfigRuleToCompliancePackResponseBodyOperateRuleResultOperateRuleItemList {
	s.Success = &v
	return s
}

func (s *AttachConfigRuleToCompliancePackResponseBodyOperateRuleResultOperateRuleItemList) Validate() error {
	return dara.Validate(s)
}
