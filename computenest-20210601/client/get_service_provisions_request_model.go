// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetServiceProvisionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetParameters(v map[string]interface{}) *GetServiceProvisionsRequest
	GetParameters() map[string]interface{}
	SetRegionId(v string) *GetServiceProvisionsRequest
	GetRegionId() *string
	SetServiceId(v string) *GetServiceProvisionsRequest
	GetServiceId() *string
	SetServiceVersion(v string) *GetServiceProvisionsRequest
	GetServiceVersion() *string
	SetTemplateName(v string) *GetServiceProvisionsRequest
	GetTemplateName() *string
	SetTrialType(v string) *GetServiceProvisionsRequest
	GetTrialType() *string
}

type GetServiceProvisionsRequest struct {
	// The parameters configured for the service instance.
	//
	// example:
	//
	// {\\"RegionId\\":\\"cn-hangzhou\\",\\"ZoneId\\":\\"cn-hangzhou-g\\",\\"EcsInstanceType\\":\\"ecs.g5.large\\",\\"InstancePassword\\":\\"xxxxxxxx\\",\\"PayType\\":\\"PostPaid\\",\\"PayPeriodUnit\\":\\"Month\\",\\"PayPeriod\\":1}
	Parameters map[string]interface{} `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The service ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// service-0efc0db451794bxxxxxx
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// The version of the service.
	//
	// example:
	//
	// 1
	ServiceVersion *string `json:"ServiceVersion,omitempty" xml:"ServiceVersion,omitempty"`
	// The template name.
	//
	// example:
	//
	// ECS
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// The trial policy. Valid values:
	//
	// 	- Trial: Trials are supported.
	//
	// 	- NotTrial: Trials are not supported.
	//
	// example:
	//
	// Trial
	TrialType *string `json:"TrialType,omitempty" xml:"TrialType,omitempty"`
}

func (s GetServiceProvisionsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetServiceProvisionsRequest) GoString() string {
	return s.String()
}

func (s *GetServiceProvisionsRequest) GetParameters() map[string]interface{} {
	return s.Parameters
}

func (s *GetServiceProvisionsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetServiceProvisionsRequest) GetServiceId() *string {
	return s.ServiceId
}

func (s *GetServiceProvisionsRequest) GetServiceVersion() *string {
	return s.ServiceVersion
}

func (s *GetServiceProvisionsRequest) GetTemplateName() *string {
	return s.TemplateName
}

func (s *GetServiceProvisionsRequest) GetTrialType() *string {
	return s.TrialType
}

func (s *GetServiceProvisionsRequest) SetParameters(v map[string]interface{}) *GetServiceProvisionsRequest {
	s.Parameters = v
	return s
}

func (s *GetServiceProvisionsRequest) SetRegionId(v string) *GetServiceProvisionsRequest {
	s.RegionId = &v
	return s
}

func (s *GetServiceProvisionsRequest) SetServiceId(v string) *GetServiceProvisionsRequest {
	s.ServiceId = &v
	return s
}

func (s *GetServiceProvisionsRequest) SetServiceVersion(v string) *GetServiceProvisionsRequest {
	s.ServiceVersion = &v
	return s
}

func (s *GetServiceProvisionsRequest) SetTemplateName(v string) *GetServiceProvisionsRequest {
	s.TemplateName = &v
	return s
}

func (s *GetServiceProvisionsRequest) SetTrialType(v string) *GetServiceProvisionsRequest {
	s.TrialType = &v
	return s
}

func (s *GetServiceProvisionsRequest) Validate() error {
	return dara.Validate(s)
}
