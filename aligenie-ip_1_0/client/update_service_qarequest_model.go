// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateServiceQARequest interface {
	dara.Model
	String() string
	GoString() string
	SetAnswer(v string) *UpdateServiceQARequest
	GetAnswer() *string
	SetHotelId(v string) *UpdateServiceQARequest
	GetHotelId() *string
	SetServiceQAId(v int64) *UpdateServiceQARequest
	GetServiceQAId() *int64
	SetIsActive(v bool) *UpdateServiceQARequest
	GetIsActive() *bool
}

type UpdateServiceQARequest struct {
	Answer *string `json:"Answer,omitempty" xml:"Answer,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// af7***536
	HotelId *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
	// example:
	//
	// 1
	ServiceQAId *int64 `json:"ServiceQAId,omitempty" xml:"ServiceQAId,omitempty"`
	// example:
	//
	// true
	IsActive *bool `json:"isActive,omitempty" xml:"isActive,omitempty"`
}

func (s UpdateServiceQARequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateServiceQARequest) GoString() string {
	return s.String()
}

func (s *UpdateServiceQARequest) GetAnswer() *string {
	return s.Answer
}

func (s *UpdateServiceQARequest) GetHotelId() *string {
	return s.HotelId
}

func (s *UpdateServiceQARequest) GetServiceQAId() *int64 {
	return s.ServiceQAId
}

func (s *UpdateServiceQARequest) GetIsActive() *bool {
	return s.IsActive
}

func (s *UpdateServiceQARequest) SetAnswer(v string) *UpdateServiceQARequest {
	s.Answer = &v
	return s
}

func (s *UpdateServiceQARequest) SetHotelId(v string) *UpdateServiceQARequest {
	s.HotelId = &v
	return s
}

func (s *UpdateServiceQARequest) SetServiceQAId(v int64) *UpdateServiceQARequest {
	s.ServiceQAId = &v
	return s
}

func (s *UpdateServiceQARequest) SetIsActive(v bool) *UpdateServiceQARequest {
	s.IsActive = &v
	return s
}

func (s *UpdateServiceQARequest) Validate() error {
	return dara.Validate(s)
}
