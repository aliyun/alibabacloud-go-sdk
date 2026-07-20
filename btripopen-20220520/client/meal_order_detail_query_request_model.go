// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMealOrderDetailQueryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetUserId(v string) *MealOrderDetailQueryRequest
	GetUserId() *string
}

type MealOrderDetailQueryRequest struct {
	// example:
	//
	// 1000
	UserId *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
}

func (s MealOrderDetailQueryRequest) String() string {
	return dara.Prettify(s)
}

func (s MealOrderDetailQueryRequest) GoString() string {
	return s.String()
}

func (s *MealOrderDetailQueryRequest) GetUserId() *string {
	return s.UserId
}

func (s *MealOrderDetailQueryRequest) SetUserId(v string) *MealOrderDetailQueryRequest {
	s.UserId = &v
	return s
}

func (s *MealOrderDetailQueryRequest) Validate() error {
	return dara.Validate(s)
}
