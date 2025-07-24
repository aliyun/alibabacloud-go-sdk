// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateServiceInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateServiceInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateServiceInstanceResponse
	GetStatusCode() *int32
	SetBody(v *CreateServiceInstanceResponseBody) *CreateServiceInstanceResponse
	GetBody() *CreateServiceInstanceResponseBody
}

type CreateServiceInstanceResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateServiceInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateServiceInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateServiceInstanceResponse) GoString() string {
	return s.String()
}

func (s *CreateServiceInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateServiceInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateServiceInstanceResponse) GetBody() *CreateServiceInstanceResponseBody {
	return s.Body
}

func (s *CreateServiceInstanceResponse) SetHeaders(v map[string]*string) *CreateServiceInstanceResponse {
	s.Headers = v
	return s
}

func (s *CreateServiceInstanceResponse) SetStatusCode(v int32) *CreateServiceInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateServiceInstanceResponse) SetBody(v *CreateServiceInstanceResponseBody) *CreateServiceInstanceResponse {
	s.Body = v
	return s
}

func (s *CreateServiceInstanceResponse) Validate() error {
	return dara.Validate(s)
}
