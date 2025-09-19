// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddCheckInstanceResultWhiteListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCheckGroupId(v string) *AddCheckInstanceResultWhiteListRequest
	GetCheckGroupId() *string
	SetCheckId(v int64) *AddCheckInstanceResultWhiteListRequest
	GetCheckId() *int64
	SetInstanceIds(v []*string) *AddCheckInstanceResultWhiteListRequest
	GetInstanceIds() []*string
	SetInstanceList(v []*AddCheckInstanceResultWhiteListRequestInstanceList) *AddCheckInstanceResultWhiteListRequest
	GetInstanceList() []*AddCheckInstanceResultWhiteListRequestInstanceList
	SetRemark(v string) *AddCheckInstanceResultWhiteListRequest
	GetRemark() *string
	SetRuleType(v string) *AddCheckInstanceResultWhiteListRequest
	GetRuleType() *string
}

type AddCheckInstanceResultWhiteListRequest struct {
	// The ID of the group to which the check item belongs.
	//
	// example:
	//
	// cQFq20UzZ49K6gRSJD1301****
	CheckGroupId *string `json:"CheckGroupId,omitempty" xml:"CheckGroupId,omitempty"`
	// The ID of the check item.
	//
	// >  You can call the [ListCheckResult](~~ListCheckResult~~) operation to query the IDs of check items.
	//
	// example:
	//
	// 132
	CheckId *int64 `json:"CheckId,omitempty" xml:"CheckId,omitempty"`
	// The instance IDs of the assets.
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	// The asset instances.
	InstanceList []*AddCheckInstanceResultWhiteListRequestInstanceList `json:"InstanceList,omitempty" xml:"InstanceList,omitempty" type:"Repeated"`
	// The description. The value of this parameter can be up to 65,535 bytes in length.
	//
	// example:
	//
	// test
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The type of the rule. Default value: **WHITE**. Valid value:
	//
	// 	- WHITE: adds check items to the whitelist.
	//
	// example:
	//
	// WHITE
	RuleType *string `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
}

func (s AddCheckInstanceResultWhiteListRequest) String() string {
	return dara.Prettify(s)
}

func (s AddCheckInstanceResultWhiteListRequest) GoString() string {
	return s.String()
}

func (s *AddCheckInstanceResultWhiteListRequest) GetCheckGroupId() *string {
	return s.CheckGroupId
}

func (s *AddCheckInstanceResultWhiteListRequest) GetCheckId() *int64 {
	return s.CheckId
}

func (s *AddCheckInstanceResultWhiteListRequest) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *AddCheckInstanceResultWhiteListRequest) GetInstanceList() []*AddCheckInstanceResultWhiteListRequestInstanceList {
	return s.InstanceList
}

func (s *AddCheckInstanceResultWhiteListRequest) GetRemark() *string {
	return s.Remark
}

func (s *AddCheckInstanceResultWhiteListRequest) GetRuleType() *string {
	return s.RuleType
}

func (s *AddCheckInstanceResultWhiteListRequest) SetCheckGroupId(v string) *AddCheckInstanceResultWhiteListRequest {
	s.CheckGroupId = &v
	return s
}

func (s *AddCheckInstanceResultWhiteListRequest) SetCheckId(v int64) *AddCheckInstanceResultWhiteListRequest {
	s.CheckId = &v
	return s
}

func (s *AddCheckInstanceResultWhiteListRequest) SetInstanceIds(v []*string) *AddCheckInstanceResultWhiteListRequest {
	s.InstanceIds = v
	return s
}

func (s *AddCheckInstanceResultWhiteListRequest) SetInstanceList(v []*AddCheckInstanceResultWhiteListRequestInstanceList) *AddCheckInstanceResultWhiteListRequest {
	s.InstanceList = v
	return s
}

func (s *AddCheckInstanceResultWhiteListRequest) SetRemark(v string) *AddCheckInstanceResultWhiteListRequest {
	s.Remark = &v
	return s
}

func (s *AddCheckInstanceResultWhiteListRequest) SetRuleType(v string) *AddCheckInstanceResultWhiteListRequest {
	s.RuleType = &v
	return s
}

func (s *AddCheckInstanceResultWhiteListRequest) Validate() error {
	return dara.Validate(s)
}

type AddCheckInstanceResultWhiteListRequestInstanceList struct {
	// The instance ID of the asset.
	//
	// >  You can call the [ListCheckInstanceResult](~~ListCheckInstanceResult~~) operation to query the instance IDs of assets.
	//
	// example:
	//
	// i-wz9fdluqx20mp2x7****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the asset.
	//
	// >  You can call the [ListCheckInstanceResult](~~ListCheckInstanceResult~~) operation to query the region ID of the asset.
	//
	// example:
	//
	// cn-hongkong
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s AddCheckInstanceResultWhiteListRequestInstanceList) String() string {
	return dara.Prettify(s)
}

func (s AddCheckInstanceResultWhiteListRequestInstanceList) GoString() string {
	return s.String()
}

func (s *AddCheckInstanceResultWhiteListRequestInstanceList) GetInstanceId() *string {
	return s.InstanceId
}

func (s *AddCheckInstanceResultWhiteListRequestInstanceList) GetRegionId() *string {
	return s.RegionId
}

func (s *AddCheckInstanceResultWhiteListRequestInstanceList) SetInstanceId(v string) *AddCheckInstanceResultWhiteListRequestInstanceList {
	s.InstanceId = &v
	return s
}

func (s *AddCheckInstanceResultWhiteListRequestInstanceList) SetRegionId(v string) *AddCheckInstanceResultWhiteListRequestInstanceList {
	s.RegionId = &v
	return s
}

func (s *AddCheckInstanceResultWhiteListRequestInstanceList) Validate() error {
	return dara.Validate(s)
}
