// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteContainerDefenseRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRuleIds(v []*int64) *DeleteContainerDefenseRuleRequest
	GetRuleIds() []*int64
}

type DeleteContainerDefenseRuleRequest struct {
	// The IDs of the rules that you want to delete.
	//
	// >  You can call the [ListContainerDefenseRule](https://help.aliyun.com/document_detail/2590599.html) operation to query the rule IDs.
	RuleIds []*int64 `json:"RuleIds,omitempty" xml:"RuleIds,omitempty" type:"Repeated"`
}

func (s DeleteContainerDefenseRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteContainerDefenseRuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteContainerDefenseRuleRequest) GetRuleIds() []*int64 {
	return s.RuleIds
}

func (s *DeleteContainerDefenseRuleRequest) SetRuleIds(v []*int64) *DeleteContainerDefenseRuleRequest {
	s.RuleIds = v
	return s
}

func (s *DeleteContainerDefenseRuleRequest) Validate() error {
	return dara.Validate(s)
}
