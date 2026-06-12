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
	// The client token that is used to ensure the idempotence of the request. Generate a unique value for this parameter from your client. The **ClientToken*	- value can contain only ASCII characters and must be no more than 64 characters in length.
	//
	// example:
	//
	// 10CM943JP0EN9D51H
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform a dry run. A dry run checks for issues such as permission errors and instance status. Valid values:
	//
	// - true: Sends a dry run request to check whether the request is valid. The service instance is not upgraded.
	//
	// - false: Sends a regular request. The service instance is upgraded after the request passes the check.
	//
	// example:
	//
	// false
	DryRun *string `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The configuration parameters of the service instance.
	//
	// example:
	//
	// {
	//
	//       "param": "value"
	//
	// }
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
	// si-5cbae874da0e47xxxxxx
	ServiceInstanceId *string `json:"ServiceInstanceId,omitempty" xml:"ServiceInstanceId,omitempty"`
	// The service version.
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
