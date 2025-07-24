// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteServiceInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DeleteServiceInstancesRequest
	GetClientToken() *string
	SetRegionId(v string) *DeleteServiceInstancesRequest
	GetRegionId() *string
	SetServiceInstanceId(v []*string) *DeleteServiceInstancesRequest
	GetServiceInstanceId() []*string
}

type DeleteServiceInstancesRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The IDs of the service instances.
	//
	// This parameter is required.
	ServiceInstanceId []*string `json:"ServiceInstanceId,omitempty" xml:"ServiceInstanceId,omitempty" type:"Repeated"`
}

func (s DeleteServiceInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteServiceInstancesRequest) GoString() string {
	return s.String()
}

func (s *DeleteServiceInstancesRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteServiceInstancesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteServiceInstancesRequest) GetServiceInstanceId() []*string {
	return s.ServiceInstanceId
}

func (s *DeleteServiceInstancesRequest) SetClientToken(v string) *DeleteServiceInstancesRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteServiceInstancesRequest) SetRegionId(v string) *DeleteServiceInstancesRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteServiceInstancesRequest) SetServiceInstanceId(v []*string) *DeleteServiceInstancesRequest {
	s.ServiceInstanceId = v
	return s
}

func (s *DeleteServiceInstancesRequest) Validate() error {
	return dara.Validate(s)
}
