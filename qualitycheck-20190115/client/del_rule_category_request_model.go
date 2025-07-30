// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDelRuleCategoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBaseMeAgentId(v int64) *DelRuleCategoryRequest
	GetBaseMeAgentId() *int64
	SetJsonStr(v string) *DelRuleCategoryRequest
	GetJsonStr() *string
}

type DelRuleCategoryRequest struct {
	// baseMeAgentId
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// This parameter is required.
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s DelRuleCategoryRequest) String() string {
	return dara.Prettify(s)
}

func (s DelRuleCategoryRequest) GoString() string {
	return s.String()
}

func (s *DelRuleCategoryRequest) GetBaseMeAgentId() *int64 {
	return s.BaseMeAgentId
}

func (s *DelRuleCategoryRequest) GetJsonStr() *string {
	return s.JsonStr
}

func (s *DelRuleCategoryRequest) SetBaseMeAgentId(v int64) *DelRuleCategoryRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *DelRuleCategoryRequest) SetJsonStr(v string) *DelRuleCategoryRequest {
	s.JsonStr = &v
	return s
}

func (s *DelRuleCategoryRequest) Validate() error {
	return dara.Validate(s)
}
