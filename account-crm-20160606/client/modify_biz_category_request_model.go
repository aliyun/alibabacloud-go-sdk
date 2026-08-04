// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyBizCategoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetParamList(v string) *ModifyBizCategoryRequest
	GetParamList() *string
	SetUserId(v int64) *ModifyBizCategoryRequest
	GetUserId() *int64
}

type ModifyBizCategoryRequest struct {
	// This parameter is required.
	ParamList *string `json:"ParamList,omitempty" xml:"ParamList,omitempty"`
	// This parameter is required.
	UserId *int64 `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ModifyBizCategoryRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyBizCategoryRequest) GoString() string {
	return s.String()
}

func (s *ModifyBizCategoryRequest) GetParamList() *string {
	return s.ParamList
}

func (s *ModifyBizCategoryRequest) GetUserId() *int64 {
	return s.UserId
}

func (s *ModifyBizCategoryRequest) SetParamList(v string) *ModifyBizCategoryRequest {
	s.ParamList = &v
	return s
}

func (s *ModifyBizCategoryRequest) SetUserId(v int64) *ModifyBizCategoryRequest {
	s.UserId = &v
	return s
}

func (s *ModifyBizCategoryRequest) Validate() error {
	return dara.Validate(s)
}
