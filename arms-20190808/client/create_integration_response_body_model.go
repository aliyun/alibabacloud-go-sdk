// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateIntegrationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetIntegration(v *CreateIntegrationResponseBodyIntegration) *CreateIntegrationResponseBody
	GetIntegration() *CreateIntegrationResponseBodyIntegration
	SetRequestId(v string) *CreateIntegrationResponseBody
	GetRequestId() *string
}

type CreateIntegrationResponseBody struct {
	// The returned information about the alert integration.
	Integration *CreateIntegrationResponseBodyIntegration `json:"Integration,omitempty" xml:"Integration,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 34ED024E-9E31-434A-9E4E-D9D15C3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateIntegrationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateIntegrationResponseBody) GoString() string {
	return s.String()
}

func (s *CreateIntegrationResponseBody) GetIntegration() *CreateIntegrationResponseBodyIntegration {
	return s.Integration
}

func (s *CreateIntegrationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateIntegrationResponseBody) SetIntegration(v *CreateIntegrationResponseBodyIntegration) *CreateIntegrationResponseBody {
	s.Integration = v
	return s
}

func (s *CreateIntegrationResponseBody) SetRequestId(v string) *CreateIntegrationResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateIntegrationResponseBody) Validate() error {
	return dara.Validate(s)
}

type CreateIntegrationResponseBodyIntegration struct {
	// Indicates whether alert events are automatically cleared. Default value: true. Valid values:
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
	// The ID of the alert integration.
	//
	// example:
	//
	// 1234
	IntegrationId *int64 `json:"IntegrationId,omitempty" xml:"IntegrationId,omitempty"`
	// The name of the alert integration.
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
}

func (s CreateIntegrationResponseBodyIntegration) String() string {
	return dara.Prettify(s)
}

func (s CreateIntegrationResponseBodyIntegration) GoString() string {
	return s.String()
}

func (s *CreateIntegrationResponseBodyIntegration) GetAutoRecover() *bool {
	return s.AutoRecover
}

func (s *CreateIntegrationResponseBodyIntegration) GetDescription() *string {
	return s.Description
}

func (s *CreateIntegrationResponseBodyIntegration) GetIntegrationId() *int64 {
	return s.IntegrationId
}

func (s *CreateIntegrationResponseBodyIntegration) GetIntegrationName() *string {
	return s.IntegrationName
}

func (s *CreateIntegrationResponseBodyIntegration) GetIntegrationProductType() *string {
	return s.IntegrationProductType
}

func (s *CreateIntegrationResponseBodyIntegration) GetRecoverTime() *int64 {
	return s.RecoverTime
}

func (s *CreateIntegrationResponseBodyIntegration) SetAutoRecover(v bool) *CreateIntegrationResponseBodyIntegration {
	s.AutoRecover = &v
	return s
}

func (s *CreateIntegrationResponseBodyIntegration) SetDescription(v string) *CreateIntegrationResponseBodyIntegration {
	s.Description = &v
	return s
}

func (s *CreateIntegrationResponseBodyIntegration) SetIntegrationId(v int64) *CreateIntegrationResponseBodyIntegration {
	s.IntegrationId = &v
	return s
}

func (s *CreateIntegrationResponseBodyIntegration) SetIntegrationName(v string) *CreateIntegrationResponseBodyIntegration {
	s.IntegrationName = &v
	return s
}

func (s *CreateIntegrationResponseBodyIntegration) SetIntegrationProductType(v string) *CreateIntegrationResponseBodyIntegration {
	s.IntegrationProductType = &v
	return s
}

func (s *CreateIntegrationResponseBodyIntegration) SetRecoverTime(v int64) *CreateIntegrationResponseBodyIntegration {
	s.RecoverTime = &v
	return s
}

func (s *CreateIntegrationResponseBodyIntegration) Validate() error {
	return dara.Validate(s)
}
