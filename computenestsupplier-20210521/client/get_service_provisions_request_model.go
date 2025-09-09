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
}

type GetServiceProvisionsRequest struct {
	// The parameters that are specified to deploy the service instance.
	//
	// example:
	//
	// {\\"RegionId\\":\\"cn-hangzhou\\"}
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
	// service-20b8a396048346xxxxxx
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// The service version.
	//
	// example:
	//
	// 1
	ServiceVersion *string `json:"ServiceVersion,omitempty" xml:"ServiceVersion,omitempty"`
	// The template name.
	//
	// example:
	//
	// 模板1
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
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

func (s *GetServiceProvisionsRequest) Validate() error {
	return dara.Validate(s)
}
