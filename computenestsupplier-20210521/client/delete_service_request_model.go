// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteServiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DeleteServiceRequest
	GetClientToken() *string
	SetRegionId(v string) *DeleteServiceRequest
	GetRegionId() *string
	SetServiceId(v string) *DeleteServiceRequest
	GetServiceId() *string
	SetServiceVersion(v string) *DeleteServiceRequest
	GetServiceVersion() *string
}

type DeleteServiceRequest struct {
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
	// service-0e6fca6a51a54420****
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// The service version.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2
	ServiceVersion *string `json:"ServiceVersion,omitempty" xml:"ServiceVersion,omitempty"`
}

func (s DeleteServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteServiceRequest) GoString() string {
	return s.String()
}

func (s *DeleteServiceRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteServiceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteServiceRequest) GetServiceId() *string {
	return s.ServiceId
}

func (s *DeleteServiceRequest) GetServiceVersion() *string {
	return s.ServiceVersion
}

func (s *DeleteServiceRequest) SetClientToken(v string) *DeleteServiceRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteServiceRequest) SetRegionId(v string) *DeleteServiceRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteServiceRequest) SetServiceId(v string) *DeleteServiceRequest {
	s.ServiceId = &v
	return s
}

func (s *DeleteServiceRequest) SetServiceVersion(v string) *DeleteServiceRequest {
	s.ServiceVersion = &v
	return s
}

func (s *DeleteServiceRequest) Validate() error {
	return dara.Validate(s)
}
