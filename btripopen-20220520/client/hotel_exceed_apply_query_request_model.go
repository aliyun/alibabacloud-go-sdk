// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHotelExceedApplyQueryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplyId(v int64) *HotelExceedApplyQueryRequest
	GetApplyId() *int64
	SetBusinessInstanceId(v string) *HotelExceedApplyQueryRequest
	GetBusinessInstanceId() *string
}

type HotelExceedApplyQueryRequest struct {
	// example:
	//
	// 1287123
	ApplyId            *int64  `json:"apply_id,omitempty" xml:"apply_id,omitempty"`
	BusinessInstanceId *string `json:"business_instance_id,omitempty" xml:"business_instance_id,omitempty"`
}

func (s HotelExceedApplyQueryRequest) String() string {
	return dara.Prettify(s)
}

func (s HotelExceedApplyQueryRequest) GoString() string {
	return s.String()
}

func (s *HotelExceedApplyQueryRequest) GetApplyId() *int64 {
	return s.ApplyId
}

func (s *HotelExceedApplyQueryRequest) GetBusinessInstanceId() *string {
	return s.BusinessInstanceId
}

func (s *HotelExceedApplyQueryRequest) SetApplyId(v int64) *HotelExceedApplyQueryRequest {
	s.ApplyId = &v
	return s
}

func (s *HotelExceedApplyQueryRequest) SetBusinessInstanceId(v string) *HotelExceedApplyQueryRequest {
	s.BusinessInstanceId = &v
	return s
}

func (s *HotelExceedApplyQueryRequest) Validate() error {
	return dara.Validate(s)
}
