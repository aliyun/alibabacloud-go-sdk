// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPreLaunchServiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *PreLaunchServiceRequest
	GetClientToken() *string
	SetRegionId(v string) *PreLaunchServiceRequest
	GetRegionId() *string
	SetServiceId(v string) *PreLaunchServiceRequest
	GetServiceId() *string
}

type PreLaunchServiceRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
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

func (s PreLaunchServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s PreLaunchServiceRequest) GoString() string {
	return s.String()
}

func (s *PreLaunchServiceRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *PreLaunchServiceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *PreLaunchServiceRequest) GetServiceId() *string {
	return s.ServiceId
}

func (s *PreLaunchServiceRequest) SetClientToken(v string) *PreLaunchServiceRequest {
	s.ClientToken = &v
	return s
}

func (s *PreLaunchServiceRequest) SetRegionId(v string) *PreLaunchServiceRequest {
	s.RegionId = &v
	return s
}

func (s *PreLaunchServiceRequest) SetServiceId(v string) *PreLaunchServiceRequest {
	s.ServiceId = &v
	return s
}

func (s *PreLaunchServiceRequest) Validate() error {
	return dara.Validate(s)
}
