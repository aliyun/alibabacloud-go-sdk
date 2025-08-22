// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRegistryModuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateRegistryModuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateRegistryModuleResponse
	GetStatusCode() *int32
	SetBody(v *CreateRegistryModuleResponseBody) *CreateRegistryModuleResponse
	GetBody() *CreateRegistryModuleResponseBody
}

type CreateRegistryModuleResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateRegistryModuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateRegistryModuleResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateRegistryModuleResponse) GoString() string {
	return s.String()
}

func (s *CreateRegistryModuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateRegistryModuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateRegistryModuleResponse) GetBody() *CreateRegistryModuleResponseBody {
	return s.Body
}

func (s *CreateRegistryModuleResponse) SetHeaders(v map[string]*string) *CreateRegistryModuleResponse {
	s.Headers = v
	return s
}

func (s *CreateRegistryModuleResponse) SetStatusCode(v int32) *CreateRegistryModuleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateRegistryModuleResponse) SetBody(v *CreateRegistryModuleResponseBody) *CreateRegistryModuleResponse {
	s.Body = v
	return s
}

func (s *CreateRegistryModuleResponse) Validate() error {
	return dara.Validate(s)
}
