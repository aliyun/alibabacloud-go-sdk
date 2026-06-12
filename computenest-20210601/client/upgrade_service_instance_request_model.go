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
	Parameters map[string]interface{} `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
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
