// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCustomerInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *GetCustomerInfoRequest
	GetInstanceId() *string
	SetMemberId(v int64) *GetCustomerInfoRequest
	GetMemberId() *int64
}

type GetCustomerInfoRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ccc_xp_pre-cn-***
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 82345678****
	MemberId *int64 `json:"MemberId,omitempty" xml:"MemberId,omitempty"`
}

func (s GetCustomerInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s GetCustomerInfoRequest) GoString() string {
	return s.String()
}

func (s *GetCustomerInfoRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetCustomerInfoRequest) GetMemberId() *int64 {
	return s.MemberId
}

func (s *GetCustomerInfoRequest) SetInstanceId(v string) *GetCustomerInfoRequest {
	s.InstanceId = &v
	return s
}

func (s *GetCustomerInfoRequest) SetMemberId(v int64) *GetCustomerInfoRequest {
	s.MemberId = &v
	return s
}

func (s *GetCustomerInfoRequest) Validate() error {
	return dara.Validate(s)
}
