// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateHealthCheckTemplateAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateHealthCheckTemplateAttributeResponseBody
	GetRequestId() *string
}

type UpdateHealthCheckTemplateAttributeResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 365F4154-92F6-4AE4-92F8-7FF34B540710
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateHealthCheckTemplateAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateHealthCheckTemplateAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateHealthCheckTemplateAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateHealthCheckTemplateAttributeResponseBody) SetRequestId(v string) *UpdateHealthCheckTemplateAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateHealthCheckTemplateAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}
