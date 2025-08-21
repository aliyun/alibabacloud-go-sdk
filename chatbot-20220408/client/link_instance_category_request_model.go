// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLinkInstanceCategoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAbilityType(v string) *LinkInstanceCategoryRequest
	GetAbilityType() *string
	SetAgentKey(v string) *LinkInstanceCategoryRequest
	GetAgentKey() *string
	SetCategoryIds(v string) *LinkInstanceCategoryRequest
	GetCategoryIds() *string
	SetInstanceId(v string) *LinkInstanceCategoryRequest
	GetInstanceId() *string
}

type LinkInstanceCategoryRequest struct {
	// example:
	//
	// FAQ,MRC
	AbilityType *string `json:"AbilityType,omitempty" xml:"AbilityType,omitempty"`
	// example:
	//
	// ac627989eb4f8a98ed05fd098bbae5_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// example:
	//
	// [\\"30000065789\\"]
	CategoryIds *string `json:"CategoryIds,omitempty" xml:"CategoryIds,omitempty"`
	// example:
	//
	// chatbot-cn-mp90s2lrk00050
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s LinkInstanceCategoryRequest) String() string {
	return dara.Prettify(s)
}

func (s LinkInstanceCategoryRequest) GoString() string {
	return s.String()
}

func (s *LinkInstanceCategoryRequest) GetAbilityType() *string {
	return s.AbilityType
}

func (s *LinkInstanceCategoryRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *LinkInstanceCategoryRequest) GetCategoryIds() *string {
	return s.CategoryIds
}

func (s *LinkInstanceCategoryRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *LinkInstanceCategoryRequest) SetAbilityType(v string) *LinkInstanceCategoryRequest {
	s.AbilityType = &v
	return s
}

func (s *LinkInstanceCategoryRequest) SetAgentKey(v string) *LinkInstanceCategoryRequest {
	s.AgentKey = &v
	return s
}

func (s *LinkInstanceCategoryRequest) SetCategoryIds(v string) *LinkInstanceCategoryRequest {
	s.CategoryIds = &v
	return s
}

func (s *LinkInstanceCategoryRequest) SetInstanceId(v string) *LinkInstanceCategoryRequest {
	s.InstanceId = &v
	return s
}

func (s *LinkInstanceCategoryRequest) Validate() error {
	return dara.Validate(s)
}
