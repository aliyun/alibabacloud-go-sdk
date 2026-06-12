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
	// A client-generated token that ensures the idempotence of the request. The token must be unique for each request. It can contain only ASCII characters and must be no more than 64 characters long.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform a dry run. A dry run checks for issues such as permissions and the instance status. Valid values:
	//
	// - true: Sends the request without upgrading the service instance.
	//
	// - false: Sends the request and upgrades the service instance after the check is passed.
	//
	// example:
	//
	// true
	DryRun *string `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The parameters required for the upgrade. This is used when new parameters are added to the new service version.
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
	// The service instance ID.
	//
	// example:
	//
	// si-d6ab3a63ccbb4bxxxxxx
	ServiceInstanceId *string `json:"ServiceInstanceId,omitempty" xml:"ServiceInstanceId,omitempty"`
	// The service version to upgrade to.
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
