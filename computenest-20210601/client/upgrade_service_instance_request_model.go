// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeServiceInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *UpgradeServiceInstanceRequest
	GetClientToken() *string
	SetDryRun(v string) *UpgradeServiceInstanceRequest
	GetDryRun() *string
	SetParameters(v map[string]interface{}) *UpgradeServiceInstanceRequest
	GetParameters() map[string]interface{}
	SetRegionId(v string) *UpgradeServiceInstanceRequest
	GetRegionId() *string
	SetServiceInstanceId(v string) *UpgradeServiceInstanceRequest
	GetServiceInstanceId() *string
	SetServiceVersion(v string) *UpgradeServiceInstanceRequest
	GetServiceVersion() *string
}

type UpgradeServiceInstanceRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform only a dry run for the request to check information such as the permissions and instance status. Valid values:
	//
	// 	- **true**: performs a dry run for the request, but does not upgrade service instance.
	//
	// 	- **false**: performs a dry run for the request, and upgrade service instance if the request passes the dry run.
	//
	// example:
	//
	// true
	DryRun *string `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The parameters required for the upgrade. This parameter is required if the destination version of the service has new parameters.
	//
	// example:
	//
	// { \\"RegionId\\": \\"cn-hangzhou\\", \\"InstanceType\\": \\"ecs.g5.large\\"}
	Parameters map[string]interface{} `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the service instance.
	//
	// example:
	//
	// si-d6ab3a63ccbb4bxxxxxx
	ServiceInstanceId *string `json:"ServiceInstanceId,omitempty" xml:"ServiceInstanceId,omitempty"`
	// The destination version.
	//
	// example:
	//
	// 2
	ServiceVersion *string `json:"ServiceVersion,omitempty" xml:"ServiceVersion,omitempty"`
}

func (s UpgradeServiceInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s UpgradeServiceInstanceRequest) GoString() string {
	return s.String()
}

func (s *UpgradeServiceInstanceRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpgradeServiceInstanceRequest) GetDryRun() *string {
	return s.DryRun
}

func (s *UpgradeServiceInstanceRequest) GetParameters() map[string]interface{} {
	return s.Parameters
}

func (s *UpgradeServiceInstanceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpgradeServiceInstanceRequest) GetServiceInstanceId() *string {
	return s.ServiceInstanceId
}

func (s *UpgradeServiceInstanceRequest) GetServiceVersion() *string {
	return s.ServiceVersion
}

func (s *UpgradeServiceInstanceRequest) SetClientToken(v string) *UpgradeServiceInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *UpgradeServiceInstanceRequest) SetDryRun(v string) *UpgradeServiceInstanceRequest {
	s.DryRun = &v
	return s
}

func (s *UpgradeServiceInstanceRequest) SetParameters(v map[string]interface{}) *UpgradeServiceInstanceRequest {
	s.Parameters = v
	return s
}

func (s *UpgradeServiceInstanceRequest) SetRegionId(v string) *UpgradeServiceInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *UpgradeServiceInstanceRequest) SetServiceInstanceId(v string) *UpgradeServiceInstanceRequest {
	s.ServiceInstanceId = &v
	return s
}

func (s *UpgradeServiceInstanceRequest) SetServiceVersion(v string) *UpgradeServiceInstanceRequest {
	s.ServiceVersion = &v
	return s
}

func (s *UpgradeServiceInstanceRequest) Validate() error {
	return dara.Validate(s)
}
