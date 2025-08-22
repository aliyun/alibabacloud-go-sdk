// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRegistryModuleAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateRegistryModuleAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateRegistryModuleAttributeResponse
	GetStatusCode() *int32
	SetBody(v *UpdateRegistryModuleAttributeResponseBody) *UpdateRegistryModuleAttributeResponse
	GetBody() *UpdateRegistryModuleAttributeResponseBody
}

type UpdateRegistryModuleAttributeResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateRegistryModuleAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateRegistryModuleAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateRegistryModuleAttributeResponse) GoString() string {
	return s.String()
}

func (s *UpdateRegistryModuleAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateRegistryModuleAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateRegistryModuleAttributeResponse) GetBody() *UpdateRegistryModuleAttributeResponseBody {
	return s.Body
}

func (s *UpdateRegistryModuleAttributeResponse) SetHeaders(v map[string]*string) *UpdateRegistryModuleAttributeResponse {
	s.Headers = v
	return s
}

func (s *UpdateRegistryModuleAttributeResponse) SetStatusCode(v int32) *UpdateRegistryModuleAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateRegistryModuleAttributeResponse) SetBody(v *UpdateRegistryModuleAttributeResponseBody) *UpdateRegistryModuleAttributeResponse {
	s.Body = v
	return s
}

func (s *UpdateRegistryModuleAttributeResponse) Validate() error {
	return dara.Validate(s)
}
