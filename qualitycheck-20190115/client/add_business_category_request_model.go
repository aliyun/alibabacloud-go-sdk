// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddBusinessCategoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBaseMeAgentId(v int64) *AddBusinessCategoryRequest
	GetBaseMeAgentId() *int64
	SetJsonStr(v string) *AddBusinessCategoryRequest
	GetJsonStr() *string
}

type AddBusinessCategoryRequest struct {
	// baseMeAgentId
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// This parameter is required.
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s AddBusinessCategoryRequest) String() string {
	return dara.Prettify(s)
}

func (s AddBusinessCategoryRequest) GoString() string {
	return s.String()
}

func (s *AddBusinessCategoryRequest) GetBaseMeAgentId() *int64 {
	return s.BaseMeAgentId
}

func (s *AddBusinessCategoryRequest) GetJsonStr() *string {
	return s.JsonStr
}

func (s *AddBusinessCategoryRequest) SetBaseMeAgentId(v int64) *AddBusinessCategoryRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *AddBusinessCategoryRequest) SetJsonStr(v string) *AddBusinessCategoryRequest {
	s.JsonStr = &v
	return s
}

func (s *AddBusinessCategoryRequest) Validate() error {
	return dara.Validate(s)
}
