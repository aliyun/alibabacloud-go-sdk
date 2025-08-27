// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMealOrderListQueryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetUserId(v string) *MealOrderListQueryRequest
	GetUserId() *string
}

type MealOrderListQueryRequest struct {
	// example:
	//
	// 1000
	UserId *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
}

func (s MealOrderListQueryRequest) String() string {
	return dara.Prettify(s)
}

func (s MealOrderListQueryRequest) GoString() string {
	return s.String()
}

func (s *MealOrderListQueryRequest) GetUserId() *string {
	return s.UserId
}

func (s *MealOrderListQueryRequest) SetUserId(v string) *MealOrderListQueryRequest {
	s.UserId = &v
	return s
}

func (s *MealOrderListQueryRequest) Validate() error {
	return dara.Validate(s)
}
