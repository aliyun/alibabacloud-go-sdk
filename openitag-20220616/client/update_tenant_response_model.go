// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTenantResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateTenantResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateTenantResponse
	GetStatusCode() *int32
	SetBody(v *UpdateTenantResponseBody) *UpdateTenantResponse
	GetBody() *UpdateTenantResponseBody
}

type UpdateTenantResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateTenantResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateTenantResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateTenantResponse) GoString() string {
	return s.String()
}

func (s *UpdateTenantResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateTenantResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateTenantResponse) GetBody() *UpdateTenantResponseBody {
	return s.Body
}

func (s *UpdateTenantResponse) SetHeaders(v map[string]*string) *UpdateTenantResponse {
	s.Headers = v
	return s
}

func (s *UpdateTenantResponse) SetStatusCode(v int32) *UpdateTenantResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateTenantResponse) SetBody(v *UpdateTenantResponseBody) *UpdateTenantResponse {
	s.Body = v
	return s
}

func (s *UpdateTenantResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
