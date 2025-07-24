// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeServiceInstanceShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *UpgradeServiceInstanceShrinkRequest
	GetClientToken() *string
	SetDryRun(v string) *UpgradeServiceInstanceShrinkRequest
	GetDryRun() *string
	SetParametersShrink(v string) *UpgradeServiceInstanceShrinkRequest
	GetParametersShrink() *string
	SetRegionId(v string) *UpgradeServiceInstanceShrinkRequest
	GetRegionId() *string
	SetServiceInstanceId(v string) *UpgradeServiceInstanceShrinkRequest
	GetServiceInstanceId() *string
	SetServiceVersion(v string) *UpgradeServiceInstanceShrinkRequest
	GetServiceVersion() *string
}

type UpgradeServiceInstanceShrinkRequest struct {
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
	ParametersShrink *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
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

func (s UpgradeServiceInstanceShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpgradeServiceInstanceShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpgradeServiceInstanceShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpgradeServiceInstanceShrinkRequest) GetDryRun() *string {
	return s.DryRun
}

func (s *UpgradeServiceInstanceShrinkRequest) GetParametersShrink() *string {
	return s.ParametersShrink
}

func (s *UpgradeServiceInstanceShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpgradeServiceInstanceShrinkRequest) GetServiceInstanceId() *string {
	return s.ServiceInstanceId
}

func (s *UpgradeServiceInstanceShrinkRequest) GetServiceVersion() *string {
	return s.ServiceVersion
}

func (s *UpgradeServiceInstanceShrinkRequest) SetClientToken(v string) *UpgradeServiceInstanceShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *UpgradeServiceInstanceShrinkRequest) SetDryRun(v string) *UpgradeServiceInstanceShrinkRequest {
	s.DryRun = &v
	return s
}

func (s *UpgradeServiceInstanceShrinkRequest) SetParametersShrink(v string) *UpgradeServiceInstanceShrinkRequest {
	s.ParametersShrink = &v
	return s
}

func (s *UpgradeServiceInstanceShrinkRequest) SetRegionId(v string) *UpgradeServiceInstanceShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *UpgradeServiceInstanceShrinkRequest) SetServiceInstanceId(v string) *UpgradeServiceInstanceShrinkRequest {
	s.ServiceInstanceId = &v
	return s
}

func (s *UpgradeServiceInstanceShrinkRequest) SetServiceVersion(v string) *UpgradeServiceInstanceShrinkRequest {
	s.ServiceVersion = &v
	return s
}

func (s *UpgradeServiceInstanceShrinkRequest) Validate() error {
	return dara.Validate(s)
}
