// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateHealthCheckTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetHealthCheckTemplateId(v string) *CreateHealthCheckTemplateResponseBody
	GetHealthCheckTemplateId() *string
	SetRequestId(v string) *CreateHealthCheckTemplateResponseBody
	GetRequestId() *string
}

type CreateHealthCheckTemplateResponseBody struct {
	// The ID of the health check template.
	//
	// example:
	//
	// hct-1224334
	HealthCheckTemplateId *string `json:"HealthCheckTemplateId,omitempty" xml:"HealthCheckTemplateId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 365F4154-92F6-4AE4-92F8-7FF34B540710
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateHealthCheckTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateHealthCheckTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *CreateHealthCheckTemplateResponseBody) GetHealthCheckTemplateId() *string {
	return s.HealthCheckTemplateId
}

func (s *CreateHealthCheckTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateHealthCheckTemplateResponseBody) SetHealthCheckTemplateId(v string) *CreateHealthCheckTemplateResponseBody {
	s.HealthCheckTemplateId = &v
	return s
}

func (s *CreateHealthCheckTemplateResponseBody) SetRequestId(v string) *CreateHealthCheckTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateHealthCheckTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}
