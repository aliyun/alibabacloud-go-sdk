// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateServiceUsageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateServiceUsageRequest
	GetClientToken() *string
	SetRegionId(v string) *CreateServiceUsageRequest
	GetRegionId() *string
	SetServiceId(v string) *CreateServiceUsageRequest
	GetServiceId() *string
}

type CreateServiceUsageRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// mRdxWuW2ts
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The region ID.
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
	// service-c2d118c9193e49xxxxxx
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
}

func (s CreateServiceUsageRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateServiceUsageRequest) GoString() string {
	return s.String()
}

func (s *CreateServiceUsageRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateServiceUsageRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateServiceUsageRequest) GetServiceId() *string {
	return s.ServiceId
}

func (s *CreateServiceUsageRequest) SetClientToken(v string) *CreateServiceUsageRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateServiceUsageRequest) SetRegionId(v string) *CreateServiceUsageRequest {
	s.RegionId = &v
	return s
}

func (s *CreateServiceUsageRequest) SetServiceId(v string) *CreateServiceUsageRequest {
	s.ServiceId = &v
	return s
}

func (s *CreateServiceUsageRequest) Validate() error {
	return dara.Validate(s)
}
