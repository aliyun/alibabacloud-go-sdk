// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTripBusinessInstanceQueryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBusinessInstanceId(v string) *TripBusinessInstanceQueryRequest
	GetBusinessInstanceId() *string
	SetThirdBusinessId(v string) *TripBusinessInstanceQueryRequest
	GetThirdBusinessId() *string
	SetUserId(v string) *TripBusinessInstanceQueryRequest
	GetUserId() *string
	SetUserName(v string) *TripBusinessInstanceQueryRequest
	GetUserName() *string
}

type TripBusinessInstanceQueryRequest struct {
	// example:
	//
	// 12345
	BusinessInstanceId *string `json:"business_instance_id,omitempty" xml:"business_instance_id,omitempty"`
	// example:
	//
	// 12345
	ThirdBusinessId *string `json:"third_business_id,omitempty" xml:"third_business_id,omitempty"`
	// example:
	//
	// thirdpart12138
	UserId   *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
	UserName *string `json:"user_name,omitempty" xml:"user_name,omitempty"`
}

func (s TripBusinessInstanceQueryRequest) String() string {
	return dara.Prettify(s)
}

func (s TripBusinessInstanceQueryRequest) GoString() string {
	return s.String()
}

func (s *TripBusinessInstanceQueryRequest) GetBusinessInstanceId() *string {
	return s.BusinessInstanceId
}

func (s *TripBusinessInstanceQueryRequest) GetThirdBusinessId() *string {
	return s.ThirdBusinessId
}

func (s *TripBusinessInstanceQueryRequest) GetUserId() *string {
	return s.UserId
}

func (s *TripBusinessInstanceQueryRequest) GetUserName() *string {
	return s.UserName
}

func (s *TripBusinessInstanceQueryRequest) SetBusinessInstanceId(v string) *TripBusinessInstanceQueryRequest {
	s.BusinessInstanceId = &v
	return s
}

func (s *TripBusinessInstanceQueryRequest) SetThirdBusinessId(v string) *TripBusinessInstanceQueryRequest {
	s.ThirdBusinessId = &v
	return s
}

func (s *TripBusinessInstanceQueryRequest) SetUserId(v string) *TripBusinessInstanceQueryRequest {
	s.UserId = &v
	return s
}

func (s *TripBusinessInstanceQueryRequest) SetUserName(v string) *TripBusinessInstanceQueryRequest {
	s.UserName = &v
	return s
}

func (s *TripBusinessInstanceQueryRequest) Validate() error {
	return dara.Validate(s)
}
