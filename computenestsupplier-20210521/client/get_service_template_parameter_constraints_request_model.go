// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetServiceTemplateParameterConstraintsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *GetServiceTemplateParameterConstraintsRequest
	GetClientToken() *string
	SetDeployRegionId(v string) *GetServiceTemplateParameterConstraintsRequest
	GetDeployRegionId() *string
	SetEnablePrivateVpcConnection(v bool) *GetServiceTemplateParameterConstraintsRequest
	GetEnablePrivateVpcConnection() *bool
	SetParameters(v []*GetServiceTemplateParameterConstraintsRequestParameters) *GetServiceTemplateParameterConstraintsRequest
	GetParameters() []*GetServiceTemplateParameterConstraintsRequestParameters
	SetRegionId(v string) *GetServiceTemplateParameterConstraintsRequest
	GetRegionId() *string
	SetServiceId(v string) *GetServiceTemplateParameterConstraintsRequest
	GetServiceId() *string
	SetServiceInstanceId(v string) *GetServiceTemplateParameterConstraintsRequest
	GetServiceInstanceId() *string
	SetServiceVersion(v string) *GetServiceTemplateParameterConstraintsRequest
	GetServiceVersion() *string
	SetTemplateName(v string) *GetServiceTemplateParameterConstraintsRequest
	GetTemplateName() *string
}

type GetServiceTemplateParameterConstraintsRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// example:
	//
	// 10CM943JP0EN9D51H
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the region in which the service instance is deployed.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-huhehaote
	DeployRegionId *string `json:"DeployRegionId,omitempty" xml:"DeployRegionId,omitempty"`
	// Specifies whether to enable the private connection feature. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	EnablePrivateVpcConnection *bool `json:"EnablePrivateVpcConnection,omitempty" xml:"EnablePrivateVpcConnection,omitempty"`
	// The parameters in the template.
	Parameters []*GetServiceTemplateParameterConstraintsRequestParameters `json:"Parameters,omitempty" xml:"Parameters,omitempty" type:"Repeated"`
	// The region ID.
	//
	// This parameter is required.
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
	// service-1c11f365190c44xxxxxx
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// The ID of the service instance.
	//
	// example:
	//
	// si-d6ab3a63ccbb4b17****
	ServiceInstanceId *string `json:"ServiceInstanceId,omitempty" xml:"ServiceInstanceId,omitempty"`
	// The service version.
	//
	// example:
	//
	// draft
	ServiceVersion *string `json:"ServiceVersion,omitempty" xml:"ServiceVersion,omitempty"`
	// The template name.
	//
	// This parameter is required.
	//
	// example:
	//
	// 模板1
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
}

func (s GetServiceTemplateParameterConstraintsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetServiceTemplateParameterConstraintsRequest) GoString() string {
	return s.String()
}

func (s *GetServiceTemplateParameterConstraintsRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *GetServiceTemplateParameterConstraintsRequest) GetDeployRegionId() *string {
	return s.DeployRegionId
}

func (s *GetServiceTemplateParameterConstraintsRequest) GetEnablePrivateVpcConnection() *bool {
	return s.EnablePrivateVpcConnection
}

func (s *GetServiceTemplateParameterConstraintsRequest) GetParameters() []*GetServiceTemplateParameterConstraintsRequestParameters {
	return s.Parameters
}

func (s *GetServiceTemplateParameterConstraintsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetServiceTemplateParameterConstraintsRequest) GetServiceId() *string {
	return s.ServiceId
}

func (s *GetServiceTemplateParameterConstraintsRequest) GetServiceInstanceId() *string {
	return s.ServiceInstanceId
}

func (s *GetServiceTemplateParameterConstraintsRequest) GetServiceVersion() *string {
	return s.ServiceVersion
}

func (s *GetServiceTemplateParameterConstraintsRequest) GetTemplateName() *string {
	return s.TemplateName
}

func (s *GetServiceTemplateParameterConstraintsRequest) SetClientToken(v string) *GetServiceTemplateParameterConstraintsRequest {
	s.ClientToken = &v
	return s
}

func (s *GetServiceTemplateParameterConstraintsRequest) SetDeployRegionId(v string) *GetServiceTemplateParameterConstraintsRequest {
	s.DeployRegionId = &v
	return s
}

func (s *GetServiceTemplateParameterConstraintsRequest) SetEnablePrivateVpcConnection(v bool) *GetServiceTemplateParameterConstraintsRequest {
	s.EnablePrivateVpcConnection = &v
	return s
}

func (s *GetServiceTemplateParameterConstraintsRequest) SetParameters(v []*GetServiceTemplateParameterConstraintsRequestParameters) *GetServiceTemplateParameterConstraintsRequest {
	s.Parameters = v
	return s
}

func (s *GetServiceTemplateParameterConstraintsRequest) SetRegionId(v string) *GetServiceTemplateParameterConstraintsRequest {
	s.RegionId = &v
	return s
}

func (s *GetServiceTemplateParameterConstraintsRequest) SetServiceId(v string) *GetServiceTemplateParameterConstraintsRequest {
	s.ServiceId = &v
	return s
}

func (s *GetServiceTemplateParameterConstraintsRequest) SetServiceInstanceId(v string) *GetServiceTemplateParameterConstraintsRequest {
	s.ServiceInstanceId = &v
	return s
}

func (s *GetServiceTemplateParameterConstraintsRequest) SetServiceVersion(v string) *GetServiceTemplateParameterConstraintsRequest {
	s.ServiceVersion = &v
	return s
}

func (s *GetServiceTemplateParameterConstraintsRequest) SetTemplateName(v string) *GetServiceTemplateParameterConstraintsRequest {
	s.TemplateName = &v
	return s
}

func (s *GetServiceTemplateParameterConstraintsRequest) Validate() error {
	return dara.Validate(s)
}

type GetServiceTemplateParameterConstraintsRequestParameters struct {
	// The parameter name.
	//
	// example:
	//
	// PayType
	ParameterKey *string `json:"ParameterKey,omitempty" xml:"ParameterKey,omitempty"`
	// The parameter value.
	//
	// example:
	//
	// PostPaid
	ParameterValue *string `json:"ParameterValue,omitempty" xml:"ParameterValue,omitempty"`
}

func (s GetServiceTemplateParameterConstraintsRequestParameters) String() string {
	return dara.Prettify(s)
}

func (s GetServiceTemplateParameterConstraintsRequestParameters) GoString() string {
	return s.String()
}

func (s *GetServiceTemplateParameterConstraintsRequestParameters) GetParameterKey() *string {
	return s.ParameterKey
}

func (s *GetServiceTemplateParameterConstraintsRequestParameters) GetParameterValue() *string {
	return s.ParameterValue
}

func (s *GetServiceTemplateParameterConstraintsRequestParameters) SetParameterKey(v string) *GetServiceTemplateParameterConstraintsRequestParameters {
	s.ParameterKey = &v
	return s
}

func (s *GetServiceTemplateParameterConstraintsRequestParameters) SetParameterValue(v string) *GetServiceTemplateParameterConstraintsRequestParameters {
	s.ParameterValue = &v
	return s
}

func (s *GetServiceTemplateParameterConstraintsRequestParameters) Validate() error {
	return dara.Validate(s)
}
