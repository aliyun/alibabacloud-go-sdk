// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeployServiceInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DeployServiceInstanceRequest
	GetClientToken() *string
	SetRegionId(v string) *DeployServiceInstanceRequest
	GetRegionId() *string
	SetServiceInstanceId(v string) *DeployServiceInstanceRequest
	GetServiceInstanceId() *string
}

type DeployServiceInstanceRequest struct {
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
	// The ID of the service instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// si-0e6fca6a51a54420****
	ServiceInstanceId *string `json:"ServiceInstanceId,omitempty" xml:"ServiceInstanceId,omitempty"`
}

func (s DeployServiceInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s DeployServiceInstanceRequest) GoString() string {
	return s.String()
}

func (s *DeployServiceInstanceRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeployServiceInstanceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeployServiceInstanceRequest) GetServiceInstanceId() *string {
	return s.ServiceInstanceId
}

func (s *DeployServiceInstanceRequest) SetClientToken(v string) *DeployServiceInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *DeployServiceInstanceRequest) SetRegionId(v string) *DeployServiceInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *DeployServiceInstanceRequest) SetServiceInstanceId(v string) *DeployServiceInstanceRequest {
	s.ServiceInstanceId = &v
	return s
}

func (s *DeployServiceInstanceRequest) Validate() error {
	return dara.Validate(s)
}
