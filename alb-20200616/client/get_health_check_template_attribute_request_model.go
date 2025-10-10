// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHealthCheckTemplateAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHealthCheckTemplateId(v string) *GetHealthCheckTemplateAttributeRequest
	GetHealthCheckTemplateId() *string
}

type GetHealthCheckTemplateAttributeRequest struct {
	// The ID of the health check template.
	//
	// This parameter is required.
	//
	// example:
	//
	// hct-x4jazoyi6tvsq9****
	HealthCheckTemplateId *string `json:"HealthCheckTemplateId,omitempty" xml:"HealthCheckTemplateId,omitempty"`
}

func (s GetHealthCheckTemplateAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s GetHealthCheckTemplateAttributeRequest) GoString() string {
	return s.String()
}

func (s *GetHealthCheckTemplateAttributeRequest) GetHealthCheckTemplateId() *string {
	return s.HealthCheckTemplateId
}

func (s *GetHealthCheckTemplateAttributeRequest) SetHealthCheckTemplateId(v string) *GetHealthCheckTemplateAttributeRequest {
	s.HealthCheckTemplateId = &v
	return s
}

func (s *GetHealthCheckTemplateAttributeRequest) Validate() error {
	return dara.Validate(s)
}
