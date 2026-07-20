// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMealApplyQueryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetThirdPartApplyId(v string) *MealApplyQueryRequest
	GetThirdPartApplyId() *string
}

type MealApplyQueryRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1234
	ThirdPartApplyId *string `json:"third_part_apply_id,omitempty" xml:"third_part_apply_id,omitempty"`
}

func (s MealApplyQueryRequest) String() string {
	return dara.Prettify(s)
}

func (s MealApplyQueryRequest) GoString() string {
	return s.String()
}

func (s *MealApplyQueryRequest) GetThirdPartApplyId() *string {
	return s.ThirdPartApplyId
}

func (s *MealApplyQueryRequest) SetThirdPartApplyId(v string) *MealApplyQueryRequest {
	s.ThirdPartApplyId = &v
	return s
}

func (s *MealApplyQueryRequest) Validate() error {
	return dara.Validate(s)
}
