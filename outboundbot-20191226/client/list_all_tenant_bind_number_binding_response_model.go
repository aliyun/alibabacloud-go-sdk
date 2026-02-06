// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAllTenantBindNumberBindingResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAllTenantBindNumberBindingResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAllTenantBindNumberBindingResponse
	GetStatusCode() *int32
	SetBody(v *ListAllTenantBindNumberBindingResponseBody) *ListAllTenantBindNumberBindingResponse
	GetBody() *ListAllTenantBindNumberBindingResponseBody
}

type ListAllTenantBindNumberBindingResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAllTenantBindNumberBindingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAllTenantBindNumberBindingResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAllTenantBindNumberBindingResponse) GoString() string {
	return s.String()
}

func (s *ListAllTenantBindNumberBindingResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAllTenantBindNumberBindingResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAllTenantBindNumberBindingResponse) GetBody() *ListAllTenantBindNumberBindingResponseBody {
	return s.Body
}

func (s *ListAllTenantBindNumberBindingResponse) SetHeaders(v map[string]*string) *ListAllTenantBindNumberBindingResponse {
	s.Headers = v
	return s
}

func (s *ListAllTenantBindNumberBindingResponse) SetStatusCode(v int32) *ListAllTenantBindNumberBindingResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAllTenantBindNumberBindingResponse) SetBody(v *ListAllTenantBindNumberBindingResponseBody) *ListAllTenantBindNumberBindingResponse {
	s.Body = v
	return s
}

func (s *ListAllTenantBindNumberBindingResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
