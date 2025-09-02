// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDsgDesensPlanAddOrUpdateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDesensRules(v []*DsgDesensPlanAddOrUpdateRequestDesensRules) *DsgDesensPlanAddOrUpdateRequest
	GetDesensRules() []*DsgDesensPlanAddOrUpdateRequestDesensRules
}

type DsgDesensPlanAddOrUpdateRequest struct {
	// A collection of data masking rules that you want to add or modify.
	//
	// This parameter is required.
	DesensRules []*DsgDesensPlanAddOrUpdateRequestDesensRules `json:"DesensRules,omitempty" xml:"DesensRules,omitempty" type:"Repeated"`
}

func (s DsgDesensPlanAddOrUpdateRequest) String() string {
	return dara.Prettify(s)
}

func (s DsgDesensPlanAddOrUpdateRequest) GoString() string {
	return s.String()
}

func (s *DsgDesensPlanAddOrUpdateRequest) GetDesensRules() []*DsgDesensPlanAddOrUpdateRequestDesensRules {
	return s.DesensRules
}

func (s *DsgDesensPlanAddOrUpdateRequest) SetDesensRules(v []*DsgDesensPlanAddOrUpdateRequestDesensRules) *DsgDesensPlanAddOrUpdateRequest {
	s.DesensRules = v
	return s
}

func (s *DsgDesensPlanAddOrUpdateRequest) Validate() error {
	return dara.Validate(s)
}

type DsgDesensPlanAddOrUpdateRequestDesensRules struct {
	// Specifies whether to add a watermark. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	CheckWatermark *bool `json:"CheckWatermark,omitempty" xml:"CheckWatermark,omitempty"`
	// The sensitive field type.
	//
	// This parameter is required.
	//
	// example:
	//
	// phone
	DataType *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	// The data masking rule.
	//
	// This parameter is required.
	DesensPlan *DsgDesensPlanAddOrUpdateRequestDesensRulesDesensPlan `json:"DesensPlan,omitempty" xml:"DesensPlan,omitempty" type:"Struct"`
	// The ID of the data masking rule. You can call the [DsgDesensPlanQueryList](https://help.aliyun.com/document_detail/2786578.html) operation to query the ID of the data masking rule.
	//
	// example:
	//
	// 123
	Id *int32 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The owner of the data masking rule.
	//
	// This parameter is required.
	//
	// example:
	//
	// test_user
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The name of the data masking rule.
	//
	// This parameter is required.
	//
	// example:
	//
	// phone_hash
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The level-2 data masking scenario.
	//
	// This parameter is required.
	SceneIds []*int32 `json:"SceneIds,omitempty" xml:"SceneIds,omitempty" type:"Repeated"`
	// The status of the data masking rule. Valid values:
	//
	// 	- 0: expired
	//
	// 	- 1: effective
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DsgDesensPlanAddOrUpdateRequestDesensRules) String() string {
	return dara.Prettify(s)
}

func (s DsgDesensPlanAddOrUpdateRequestDesensRules) GoString() string {
	return s.String()
}

func (s *DsgDesensPlanAddOrUpdateRequestDesensRules) GetCheckWatermark() *bool {
	return s.CheckWatermark
}

func (s *DsgDesensPlanAddOrUpdateRequestDesensRules) GetDataType() *string {
	return s.DataType
}

func (s *DsgDesensPlanAddOrUpdateRequestDesensRules) GetDesensPlan() *DsgDesensPlanAddOrUpdateRequestDesensRulesDesensPlan {
	return s.DesensPlan
}

func (s *DsgDesensPlanAddOrUpdateRequestDesensRules) GetId() *int32 {
	return s.Id
}

func (s *DsgDesensPlanAddOrUpdateRequestDesensRules) GetOwner() *string {
	return s.Owner
}

func (s *DsgDesensPlanAddOrUpdateRequestDesensRules) GetRuleName() *string {
	return s.RuleName
}

func (s *DsgDesensPlanAddOrUpdateRequestDesensRules) GetSceneIds() []*int32 {
	return s.SceneIds
}

func (s *DsgDesensPlanAddOrUpdateRequestDesensRules) GetStatus() *int32 {
	return s.Status
}

func (s *DsgDesensPlanAddOrUpdateRequestDesensRules) SetCheckWatermark(v bool) *DsgDesensPlanAddOrUpdateRequestDesensRules {
	s.CheckWatermark = &v
	return s
}

func (s *DsgDesensPlanAddOrUpdateRequestDesensRules) SetDataType(v string) *DsgDesensPlanAddOrUpdateRequestDesensRules {
	s.DataType = &v
	return s
}

func (s *DsgDesensPlanAddOrUpdateRequestDesensRules) SetDesensPlan(v *DsgDesensPlanAddOrUpdateRequestDesensRulesDesensPlan) *DsgDesensPlanAddOrUpdateRequestDesensRules {
	s.DesensPlan = v
	return s
}

func (s *DsgDesensPlanAddOrUpdateRequestDesensRules) SetId(v int32) *DsgDesensPlanAddOrUpdateRequestDesensRules {
	s.Id = &v
	return s
}

func (s *DsgDesensPlanAddOrUpdateRequestDesensRules) SetOwner(v string) *DsgDesensPlanAddOrUpdateRequestDesensRules {
	s.Owner = &v
	return s
}

func (s *DsgDesensPlanAddOrUpdateRequestDesensRules) SetRuleName(v string) *DsgDesensPlanAddOrUpdateRequestDesensRules {
	s.RuleName = &v
	return s
}

func (s *DsgDesensPlanAddOrUpdateRequestDesensRules) SetSceneIds(v []*int32) *DsgDesensPlanAddOrUpdateRequestDesensRules {
	s.SceneIds = v
	return s
}

func (s *DsgDesensPlanAddOrUpdateRequestDesensRules) SetStatus(v int32) *DsgDesensPlanAddOrUpdateRequestDesensRules {
	s.Status = &v
	return s
}

func (s *DsgDesensPlanAddOrUpdateRequestDesensRules) Validate() error {
	return dara.Validate(s)
}

type DsgDesensPlanAddOrUpdateRequestDesensRulesDesensPlan struct {
	// The masking method configured in the data masking rule. Valid values:
	//
	// 	- hash
	//
	// 	- mapping
	//
	// 	- mask
	//
	// 	- charreplacement
	//
	// 	- intervalselect
	//
	// 	- decimalpoint
	//
	// 	- emptydesens
	//
	// This parameter is required.
	//
	// example:
	//
	// hash
	DesensPlanType *string `json:"DesensPlanType,omitempty" xml:"DesensPlanType,omitempty"`
	// The parameters for the data masking rule.
	ExtParam map[string]interface{} `json:"ExtParam,omitempty" xml:"ExtParam,omitempty"`
}

func (s DsgDesensPlanAddOrUpdateRequestDesensRulesDesensPlan) String() string {
	return dara.Prettify(s)
}

func (s DsgDesensPlanAddOrUpdateRequestDesensRulesDesensPlan) GoString() string {
	return s.String()
}

func (s *DsgDesensPlanAddOrUpdateRequestDesensRulesDesensPlan) GetDesensPlanType() *string {
	return s.DesensPlanType
}

func (s *DsgDesensPlanAddOrUpdateRequestDesensRulesDesensPlan) GetExtParam() map[string]interface{} {
	return s.ExtParam
}

func (s *DsgDesensPlanAddOrUpdateRequestDesensRulesDesensPlan) SetDesensPlanType(v string) *DsgDesensPlanAddOrUpdateRequestDesensRulesDesensPlan {
	s.DesensPlanType = &v
	return s
}

func (s *DsgDesensPlanAddOrUpdateRequestDesensRulesDesensPlan) SetExtParam(v map[string]interface{}) *DsgDesensPlanAddOrUpdateRequestDesensRulesDesensPlan {
	s.ExtParam = v
	return s
}

func (s *DsgDesensPlanAddOrUpdateRequestDesensRulesDesensPlan) Validate() error {
	return dara.Validate(s)
}
