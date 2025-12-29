// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSceneCategoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHotelId(v string) *ListSceneCategoryRequest
	GetHotelId() *string
	SetType(v string) *ListSceneCategoryRequest
	GetType() *string
}

type ListSceneCategoryRequest struct {
	// hotelId
	//
	// This parameter is required.
	//
	// example:
	//
	// 80d84ea8ed9e422fbad52715c8fc56f1
	HotelId *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// REPAIR
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListSceneCategoryRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSceneCategoryRequest) GoString() string {
	return s.String()
}

func (s *ListSceneCategoryRequest) GetHotelId() *string {
	return s.HotelId
}

func (s *ListSceneCategoryRequest) GetType() *string {
	return s.Type
}

func (s *ListSceneCategoryRequest) SetHotelId(v string) *ListSceneCategoryRequest {
	s.HotelId = &v
	return s
}

func (s *ListSceneCategoryRequest) SetType(v string) *ListSceneCategoryRequest {
	s.Type = &v
	return s
}

func (s *ListSceneCategoryRequest) Validate() error {
	return dara.Validate(s)
}
