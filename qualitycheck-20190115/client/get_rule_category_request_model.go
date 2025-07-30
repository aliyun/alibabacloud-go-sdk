// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRuleCategoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBaseMeAgentId(v int64) *GetRuleCategoryRequest
	GetBaseMeAgentId() *int64
	SetJsonStr(v string) *GetRuleCategoryRequest
	GetJsonStr() *string
}

type GetRuleCategoryRequest struct {
	// baseMeAgentId
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ""
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s GetRuleCategoryRequest) String() string {
	return dara.Prettify(s)
}

func (s GetRuleCategoryRequest) GoString() string {
	return s.String()
}

func (s *GetRuleCategoryRequest) GetBaseMeAgentId() *int64 {
	return s.BaseMeAgentId
}

func (s *GetRuleCategoryRequest) GetJsonStr() *string {
	return s.JsonStr
}

func (s *GetRuleCategoryRequest) SetBaseMeAgentId(v int64) *GetRuleCategoryRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *GetRuleCategoryRequest) SetJsonStr(v string) *GetRuleCategoryRequest {
	s.JsonStr = &v
	return s
}

func (s *GetRuleCategoryRequest) Validate() error {
	return dara.Validate(s)
}
