// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTripTaskQueryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBusinessInstanceId(v string) *TripTaskQueryRequest
	GetBusinessInstanceId() *string
	SetThirdBusinessId(v string) *TripTaskQueryRequest
	GetThirdBusinessId() *string
	SetUserId(v string) *TripTaskQueryRequest
	GetUserId() *string
	SetUserName(v string) *TripTaskQueryRequest
	GetUserName() *string
}

type TripTaskQueryRequest struct {
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

func (s TripTaskQueryRequest) String() string {
	return dara.Prettify(s)
}

func (s TripTaskQueryRequest) GoString() string {
	return s.String()
}

func (s *TripTaskQueryRequest) GetBusinessInstanceId() *string {
	return s.BusinessInstanceId
}

func (s *TripTaskQueryRequest) GetThirdBusinessId() *string {
	return s.ThirdBusinessId
}

func (s *TripTaskQueryRequest) GetUserId() *string {
	return s.UserId
}

func (s *TripTaskQueryRequest) GetUserName() *string {
	return s.UserName
}

func (s *TripTaskQueryRequest) SetBusinessInstanceId(v string) *TripTaskQueryRequest {
	s.BusinessInstanceId = &v
	return s
}

func (s *TripTaskQueryRequest) SetThirdBusinessId(v string) *TripTaskQueryRequest {
	s.ThirdBusinessId = &v
	return s
}

func (s *TripTaskQueryRequest) SetUserId(v string) *TripTaskQueryRequest {
	s.UserId = &v
	return s
}

func (s *TripTaskQueryRequest) SetUserName(v string) *TripTaskQueryRequest {
	s.UserName = &v
	return s
}

func (s *TripTaskQueryRequest) Validate() error {
	return dara.Validate(s)
}
