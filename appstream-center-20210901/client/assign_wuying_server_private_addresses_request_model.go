// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssignWuyingServerPrivateAddressesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSecondaryPrivateIpAddressCount(v int32) *AssignWuyingServerPrivateAddressesRequest
	GetSecondaryPrivateIpAddressCount() *int32
	SetWuyingServerId(v string) *AssignWuyingServerPrivateAddressesRequest
	GetWuyingServerId() *string
}

type AssignWuyingServerPrivateAddressesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1
	SecondaryPrivateIpAddressCount *int32 `json:"SecondaryPrivateIpAddressCount,omitempty" xml:"SecondaryPrivateIpAddressCount,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ws-bp1234567890abcde
	WuyingServerId *string `json:"WuyingServerId,omitempty" xml:"WuyingServerId,omitempty"`
}

func (s AssignWuyingServerPrivateAddressesRequest) String() string {
	return dara.Prettify(s)
}

func (s AssignWuyingServerPrivateAddressesRequest) GoString() string {
	return s.String()
}

func (s *AssignWuyingServerPrivateAddressesRequest) GetSecondaryPrivateIpAddressCount() *int32 {
	return s.SecondaryPrivateIpAddressCount
}

func (s *AssignWuyingServerPrivateAddressesRequest) GetWuyingServerId() *string {
	return s.WuyingServerId
}

func (s *AssignWuyingServerPrivateAddressesRequest) SetSecondaryPrivateIpAddressCount(v int32) *AssignWuyingServerPrivateAddressesRequest {
	s.SecondaryPrivateIpAddressCount = &v
	return s
}

func (s *AssignWuyingServerPrivateAddressesRequest) SetWuyingServerId(v string) *AssignWuyingServerPrivateAddressesRequest {
	s.WuyingServerId = &v
	return s
}

func (s *AssignWuyingServerPrivateAddressesRequest) Validate() error {
	return dara.Validate(s)
}
