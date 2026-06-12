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
	// A client token to ensure the idempotence of a request. Generate a unique value from the client for this parameter. The token can contain only ASCII characters and must be no more than 64 characters in length.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The region ID. Possible values:
	//
	// - cn-hangzhou: China (Hangzhou).
	//
	// - ap-southeast-1: Singapore.
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
	// si-b58c874912fc4294****
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
