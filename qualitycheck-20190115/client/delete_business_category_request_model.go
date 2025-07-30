// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBusinessCategoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBaseMeAgentId(v int64) *DeleteBusinessCategoryRequest
	GetBaseMeAgentId() *int64
	SetJsonStr(v string) *DeleteBusinessCategoryRequest
	GetJsonStr() *string
}

type DeleteBusinessCategoryRequest struct {
	// baseMeAgentId
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// This parameter is required.
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s DeleteBusinessCategoryRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteBusinessCategoryRequest) GoString() string {
	return s.String()
}

func (s *DeleteBusinessCategoryRequest) GetBaseMeAgentId() *int64 {
	return s.BaseMeAgentId
}

func (s *DeleteBusinessCategoryRequest) GetJsonStr() *string {
	return s.JsonStr
}

func (s *DeleteBusinessCategoryRequest) SetBaseMeAgentId(v int64) *DeleteBusinessCategoryRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *DeleteBusinessCategoryRequest) SetJsonStr(v string) *DeleteBusinessCategoryRequest {
	s.JsonStr = &v
	return s
}

func (s *DeleteBusinessCategoryRequest) Validate() error {
	return dara.Validate(s)
}
