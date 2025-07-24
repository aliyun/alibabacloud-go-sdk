// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iContinueDeployServiceInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *ContinueDeployServiceInstanceRequest
	GetClientToken() *string
	SetDryRun(v bool) *ContinueDeployServiceInstanceRequest
	GetDryRun() *bool
	SetOption(v []*string) *ContinueDeployServiceInstanceRequest
	GetOption() []*string
	SetParameters(v string) *ContinueDeployServiceInstanceRequest
	GetParameters() *string
	SetRegionId(v string) *ContinueDeployServiceInstanceRequest
	GetRegionId() *string
	SetServiceInstanceId(v string) *ContinueDeployServiceInstanceRequest
	GetServiceInstanceId() *string
}

type ContinueDeployServiceInstanceRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform only a dry run for the request to check information such as the permissions and instance status. Valid values:
	//
	// 	- true: performs a dry run for the request, but does not create a service instance.
	//
	// 	- false: performs a dry run for the request, and creates a service instance if the request passes the dry run.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The options that the system adopts when the system continues to create the service instance.
	Option []*string `json:"Option,omitempty" xml:"Option,omitempty" type:"Repeated"`
	// The parameters configured for the service instance.
	//
	// example:
	//
	// {"NodeCount": 3, "SystemDiskSize": 40, "InstancePassword": "******"}
	Parameters *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
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

func (s ContinueDeployServiceInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s ContinueDeployServiceInstanceRequest) GoString() string {
	return s.String()
}

func (s *ContinueDeployServiceInstanceRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ContinueDeployServiceInstanceRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *ContinueDeployServiceInstanceRequest) GetOption() []*string {
	return s.Option
}

func (s *ContinueDeployServiceInstanceRequest) GetParameters() *string {
	return s.Parameters
}

func (s *ContinueDeployServiceInstanceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ContinueDeployServiceInstanceRequest) GetServiceInstanceId() *string {
	return s.ServiceInstanceId
}

func (s *ContinueDeployServiceInstanceRequest) SetClientToken(v string) *ContinueDeployServiceInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *ContinueDeployServiceInstanceRequest) SetDryRun(v bool) *ContinueDeployServiceInstanceRequest {
	s.DryRun = &v
	return s
}

func (s *ContinueDeployServiceInstanceRequest) SetOption(v []*string) *ContinueDeployServiceInstanceRequest {
	s.Option = v
	return s
}

func (s *ContinueDeployServiceInstanceRequest) SetParameters(v string) *ContinueDeployServiceInstanceRequest {
	s.Parameters = &v
	return s
}

func (s *ContinueDeployServiceInstanceRequest) SetRegionId(v string) *ContinueDeployServiceInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *ContinueDeployServiceInstanceRequest) SetServiceInstanceId(v string) *ContinueDeployServiceInstanceRequest {
	s.ServiceInstanceId = &v
	return s
}

func (s *ContinueDeployServiceInstanceRequest) Validate() error {
	return dara.Validate(s)
}
