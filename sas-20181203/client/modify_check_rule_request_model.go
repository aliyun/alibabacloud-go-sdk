// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCheckRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddInstanceList(v []*ModifyCheckRuleRequestAddInstanceList) *ModifyCheckRuleRequest
	GetAddInstanceList() []*ModifyCheckRuleRequestAddInstanceList
	SetDeleteInstanceList(v []*ModifyCheckRuleRequestDeleteInstanceList) *ModifyCheckRuleRequest
	GetDeleteInstanceList() []*ModifyCheckRuleRequestDeleteInstanceList
	SetRemark(v string) *ModifyCheckRuleRequest
	GetRemark() *string
	SetRuleId(v int64) *ModifyCheckRuleRequest
	GetRuleId() *int64
	SetRuleType(v string) *ModifyCheckRuleRequest
	GetRuleType() *string
	SetScopeType(v string) *ModifyCheckRuleRequest
	GetScopeType() *string
}

type ModifyCheckRuleRequest struct {
	// The list of instances to be added in this rule update. This parameter does not need to be passed if there are no instances to add.
	AddInstanceList []*ModifyCheckRuleRequestAddInstanceList `json:"AddInstanceList,omitempty" xml:"AddInstanceList,omitempty" type:"Repeated"`
	// The list of instances to be removed in this rule update. This parameter does not need to be passed if there are no instances to remove.
	DeleteInstanceList []*ModifyCheckRuleRequestDeleteInstanceList `json:"DeleteInstanceList,omitempty" xml:"DeleteInstanceList,omitempty" type:"Repeated"`
	// Remarks.
	//
	// example:
	//
	// testRemark
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// Rule ID.
	//
	// > You can obtain this parameter by calling the [ListCheckRule](https://help.aliyun.com/document_detail/2590599.html) API.
	//
	// This parameter is required.
	//
	// example:
	//
	// 9000**
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// Rule type. Default is **WHITE**. Values:
	//
	// - **WHITE**: Add to whitelist
	//
	// example:
	//
	// WHITE
	RuleType *string `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
	// The scope of effect for modifying the rule:
	//
	// - **INSTANCE**: Instance
	//
	// - **ITEM**: Check item
	//
	// example:
	//
	// INSTANCE
	ScopeType *string `json:"ScopeType,omitempty" xml:"ScopeType,omitempty"`
}

func (s ModifyCheckRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyCheckRuleRequest) GoString() string {
	return s.String()
}

func (s *ModifyCheckRuleRequest) GetAddInstanceList() []*ModifyCheckRuleRequestAddInstanceList {
	return s.AddInstanceList
}

func (s *ModifyCheckRuleRequest) GetDeleteInstanceList() []*ModifyCheckRuleRequestDeleteInstanceList {
	return s.DeleteInstanceList
}

func (s *ModifyCheckRuleRequest) GetRemark() *string {
	return s.Remark
}

func (s *ModifyCheckRuleRequest) GetRuleId() *int64 {
	return s.RuleId
}

func (s *ModifyCheckRuleRequest) GetRuleType() *string {
	return s.RuleType
}

func (s *ModifyCheckRuleRequest) GetScopeType() *string {
	return s.ScopeType
}

func (s *ModifyCheckRuleRequest) SetAddInstanceList(v []*ModifyCheckRuleRequestAddInstanceList) *ModifyCheckRuleRequest {
	s.AddInstanceList = v
	return s
}

func (s *ModifyCheckRuleRequest) SetDeleteInstanceList(v []*ModifyCheckRuleRequestDeleteInstanceList) *ModifyCheckRuleRequest {
	s.DeleteInstanceList = v
	return s
}

func (s *ModifyCheckRuleRequest) SetRemark(v string) *ModifyCheckRuleRequest {
	s.Remark = &v
	return s
}

func (s *ModifyCheckRuleRequest) SetRuleId(v int64) *ModifyCheckRuleRequest {
	s.RuleId = &v
	return s
}

func (s *ModifyCheckRuleRequest) SetRuleType(v string) *ModifyCheckRuleRequest {
	s.RuleType = &v
	return s
}

func (s *ModifyCheckRuleRequest) SetScopeType(v string) *ModifyCheckRuleRequest {
	s.ScopeType = &v
	return s
}

func (s *ModifyCheckRuleRequest) Validate() error {
	if s.AddInstanceList != nil {
		for _, item := range s.AddInstanceList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.DeleteInstanceList != nil {
		for _, item := range s.DeleteInstanceList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ModifyCheckRuleRequestAddInstanceList struct {
	// The instance ID of the asset.
	//
	// example:
	//
	// i-wz9g8ljygfqs1ez3****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the asset.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ModifyCheckRuleRequestAddInstanceList) String() string {
	return dara.Prettify(s)
}

func (s ModifyCheckRuleRequestAddInstanceList) GoString() string {
	return s.String()
}

func (s *ModifyCheckRuleRequestAddInstanceList) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyCheckRuleRequestAddInstanceList) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyCheckRuleRequestAddInstanceList) SetInstanceId(v string) *ModifyCheckRuleRequestAddInstanceList {
	s.InstanceId = &v
	return s
}

func (s *ModifyCheckRuleRequestAddInstanceList) SetRegionId(v string) *ModifyCheckRuleRequestAddInstanceList {
	s.RegionId = &v
	return s
}

func (s *ModifyCheckRuleRequestAddInstanceList) Validate() error {
	return dara.Validate(s)
}

type ModifyCheckRuleRequestDeleteInstanceList struct {
	// The instance ID of the asset.
	//
	// example:
	//
	// i-8vb0e8qdaj0yyxjo****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the asset.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ModifyCheckRuleRequestDeleteInstanceList) String() string {
	return dara.Prettify(s)
}

func (s ModifyCheckRuleRequestDeleteInstanceList) GoString() string {
	return s.String()
}

func (s *ModifyCheckRuleRequestDeleteInstanceList) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyCheckRuleRequestDeleteInstanceList) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyCheckRuleRequestDeleteInstanceList) SetInstanceId(v string) *ModifyCheckRuleRequestDeleteInstanceList {
	s.InstanceId = &v
	return s
}

func (s *ModifyCheckRuleRequestDeleteInstanceList) SetRegionId(v string) *ModifyCheckRuleRequestDeleteInstanceList {
	s.RegionId = &v
	return s
}

func (s *ModifyCheckRuleRequestDeleteInstanceList) Validate() error {
	return dara.Validate(s)
}
