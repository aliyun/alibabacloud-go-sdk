// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddRuleCategoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBaseMeAgentId(v int64) *AddRuleCategoryRequest
	GetBaseMeAgentId() *int64
	SetJsonStr(v string) *AddRuleCategoryRequest
	GetJsonStr() *string
}

type AddRuleCategoryRequest struct {
	// baseMeAgentId
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// This parameter is required.
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s AddRuleCategoryRequest) String() string {
	return dara.Prettify(s)
}

func (s AddRuleCategoryRequest) GoString() string {
	return s.String()
}

func (s *AddRuleCategoryRequest) GetBaseMeAgentId() *int64 {
	return s.BaseMeAgentId
}

func (s *AddRuleCategoryRequest) GetJsonStr() *string {
	return s.JsonStr
}

func (s *AddRuleCategoryRequest) SetBaseMeAgentId(v int64) *AddRuleCategoryRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *AddRuleCategoryRequest) SetJsonStr(v string) *AddRuleCategoryRequest {
	s.JsonStr = &v
	return s
}

func (s *AddRuleCategoryRequest) Validate() error {
	return dara.Validate(s)
}
