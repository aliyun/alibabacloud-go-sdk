// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateIntegrationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoRecover(v bool) *CreateIntegrationRequest
	GetAutoRecover() *bool
	SetDescription(v string) *CreateIntegrationRequest
	GetDescription() *string
	SetIntegrationName(v string) *CreateIntegrationRequest
	GetIntegrationName() *string
	SetIntegrationProductType(v string) *CreateIntegrationRequest
	GetIntegrationProductType() *string
	SetRecoverTime(v int64) *CreateIntegrationRequest
	GetRecoverTime() *int64
	SetRegionId(v string) *CreateIntegrationRequest
	GetRegionId() *string
}

type CreateIntegrationRequest struct {
	// Specifies whether to automatically clear alert events. Default value: true. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	AutoRecover *bool `json:"AutoRecover,omitempty" xml:"AutoRecover,omitempty"`
	// The description of the alert integration.
	//
	// example:
	//
	// Test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the alert integration.
	//
	// This parameter is required.
	//
	// example:
	//
	// CloudMonitor integration
	IntegrationName *string `json:"IntegrationName,omitempty" xml:"IntegrationName,omitempty"`
	// The service of the alert integration. Valid values:
	//
	// 	- CLOUD_MONITOR: CloudMonitor
	//
	// 	- LOG_SERVICE: Log Service
	//
	// This parameter is required.
	//
	// example:
	//
	// CLOUD_MONITOR
	IntegrationProductType *string `json:"IntegrationProductType,omitempty" xml:"IntegrationProductType,omitempty"`
	// The period of time within which alert events are automatically cleared. Unit: seconds. Default value: 300.
	//
	// example:
	//
	// 300
	RecoverTime *int64 `json:"RecoverTime,omitempty" xml:"RecoverTime,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateIntegrationRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateIntegrationRequest) GoString() string {
	return s.String()
}

func (s *CreateIntegrationRequest) GetAutoRecover() *bool {
	return s.AutoRecover
}

func (s *CreateIntegrationRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateIntegrationRequest) GetIntegrationName() *string {
	return s.IntegrationName
}

func (s *CreateIntegrationRequest) GetIntegrationProductType() *string {
	return s.IntegrationProductType
}

func (s *CreateIntegrationRequest) GetRecoverTime() *int64 {
	return s.RecoverTime
}

func (s *CreateIntegrationRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateIntegrationRequest) SetAutoRecover(v bool) *CreateIntegrationRequest {
	s.AutoRecover = &v
	return s
}

func (s *CreateIntegrationRequest) SetDescription(v string) *CreateIntegrationRequest {
	s.Description = &v
	return s
}

func (s *CreateIntegrationRequest) SetIntegrationName(v string) *CreateIntegrationRequest {
	s.IntegrationName = &v
	return s
}

func (s *CreateIntegrationRequest) SetIntegrationProductType(v string) *CreateIntegrationRequest {
	s.IntegrationProductType = &v
	return s
}

func (s *CreateIntegrationRequest) SetRecoverTime(v int64) *CreateIntegrationRequest {
	s.RecoverTime = &v
	return s
}

func (s *CreateIntegrationRequest) SetRegionId(v string) *CreateIntegrationRequest {
	s.RegionId = &v
	return s
}

func (s *CreateIntegrationRequest) Validate() error {
	return dara.Validate(s)
}
