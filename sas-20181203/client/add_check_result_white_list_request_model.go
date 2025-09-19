// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddCheckResultWhiteListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCheckIds(v []*int64) *AddCheckResultWhiteListRequest
	GetCheckIds() []*int64
	SetRemark(v string) *AddCheckResultWhiteListRequest
	GetRemark() *string
	SetRuleType(v string) *AddCheckResultWhiteListRequest
	GetRuleType() *string
}

type AddCheckResultWhiteListRequest struct {
	// The IDs of the check items.
	//
	// >  You can call the [ListCheckResult](~~ListCheckResult~~) operation to query the IDs of the check items.
	CheckIds []*int64 `json:"CheckIds,omitempty" xml:"CheckIds,omitempty" type:"Repeated"`
	// The description. The value of this parameter can be up to 65,535 bytes in length.
	//
	// example:
	//
	// test
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The type of the rule. Default value: **WHITE**. Valid value:
	//
	// 	- **WHITE**: Add check items to the whitelist.
	//
	// example:
	//
	// WHITE
	RuleType *string `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
}

func (s AddCheckResultWhiteListRequest) String() string {
	return dara.Prettify(s)
}

func (s AddCheckResultWhiteListRequest) GoString() string {
	return s.String()
}

func (s *AddCheckResultWhiteListRequest) GetCheckIds() []*int64 {
	return s.CheckIds
}

func (s *AddCheckResultWhiteListRequest) GetRemark() *string {
	return s.Remark
}

func (s *AddCheckResultWhiteListRequest) GetRuleType() *string {
	return s.RuleType
}

func (s *AddCheckResultWhiteListRequest) SetCheckIds(v []*int64) *AddCheckResultWhiteListRequest {
	s.CheckIds = v
	return s
}

func (s *AddCheckResultWhiteListRequest) SetRemark(v string) *AddCheckResultWhiteListRequest {
	s.Remark = &v
	return s
}

func (s *AddCheckResultWhiteListRequest) SetRuleType(v string) *AddCheckResultWhiteListRequest {
	s.RuleType = &v
	return s
}

func (s *AddCheckResultWhiteListRequest) Validate() error {
	return dara.Validate(s)
}
