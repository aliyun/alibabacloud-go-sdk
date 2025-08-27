// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFlightExceedApplyQueryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplyId(v int64) *FlightExceedApplyQueryRequest
	GetApplyId() *int64
	SetBusinessInstanceId(v string) *FlightExceedApplyQueryRequest
	GetBusinessInstanceId() *string
}

type FlightExceedApplyQueryRequest struct {
	// example:
	//
	// 175634
	ApplyId            *int64  `json:"apply_id,omitempty" xml:"apply_id,omitempty"`
	BusinessInstanceId *string `json:"business_instance_id,omitempty" xml:"business_instance_id,omitempty"`
}

func (s FlightExceedApplyQueryRequest) String() string {
	return dara.Prettify(s)
}

func (s FlightExceedApplyQueryRequest) GoString() string {
	return s.String()
}

func (s *FlightExceedApplyQueryRequest) GetApplyId() *int64 {
	return s.ApplyId
}

func (s *FlightExceedApplyQueryRequest) GetBusinessInstanceId() *string {
	return s.BusinessInstanceId
}

func (s *FlightExceedApplyQueryRequest) SetApplyId(v int64) *FlightExceedApplyQueryRequest {
	s.ApplyId = &v
	return s
}

func (s *FlightExceedApplyQueryRequest) SetBusinessInstanceId(v string) *FlightExceedApplyQueryRequest {
	s.BusinessInstanceId = &v
	return s
}

func (s *FlightExceedApplyQueryRequest) Validate() error {
	return dara.Validate(s)
}
