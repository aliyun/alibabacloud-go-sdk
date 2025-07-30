// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBusinessCategoryListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBaseMeAgentId(v int64) *GetBusinessCategoryListRequest
	GetBaseMeAgentId() *int64
	SetJsonStr(v string) *GetBusinessCategoryListRequest
	GetJsonStr() *string
}

type GetBusinessCategoryListRequest struct {
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ""
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s GetBusinessCategoryListRequest) String() string {
	return dara.Prettify(s)
}

func (s GetBusinessCategoryListRequest) GoString() string {
	return s.String()
}

func (s *GetBusinessCategoryListRequest) GetBaseMeAgentId() *int64 {
	return s.BaseMeAgentId
}

func (s *GetBusinessCategoryListRequest) GetJsonStr() *string {
	return s.JsonStr
}

func (s *GetBusinessCategoryListRequest) SetBaseMeAgentId(v int64) *GetBusinessCategoryListRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *GetBusinessCategoryListRequest) SetJsonStr(v string) *GetBusinessCategoryListRequest {
	s.JsonStr = &v
	return s
}

func (s *GetBusinessCategoryListRequest) Validate() error {
	return dara.Validate(s)
}
