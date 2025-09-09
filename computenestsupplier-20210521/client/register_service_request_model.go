// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRegisterServiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *RegisterServiceRequest
	GetClientToken() *string
	SetRegionId(v string) *RegisterServiceRequest
	GetRegionId() *string
	SetServiceId(v string) *RegisterServiceRequest
	GetServiceId() *string
}

type RegisterServiceRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 10CM943JP0EN9D51H
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The service ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// service-f7024a22ea5149xxxxxx
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
}

func (s RegisterServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s RegisterServiceRequest) GoString() string {
	return s.String()
}

func (s *RegisterServiceRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *RegisterServiceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *RegisterServiceRequest) GetServiceId() *string {
	return s.ServiceId
}

func (s *RegisterServiceRequest) SetClientToken(v string) *RegisterServiceRequest {
	s.ClientToken = &v
	return s
}

func (s *RegisterServiceRequest) SetRegionId(v string) *RegisterServiceRequest {
	s.RegionId = &v
	return s
}

func (s *RegisterServiceRequest) SetServiceId(v string) *RegisterServiceRequest {
	s.ServiceId = &v
	return s
}

func (s *RegisterServiceRequest) Validate() error {
	return dara.Validate(s)
}
