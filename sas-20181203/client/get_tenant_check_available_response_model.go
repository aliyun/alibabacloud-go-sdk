// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTenantCheckAvailableResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetTenantCheckAvailableResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetTenantCheckAvailableResponse
	GetStatusCode() *int32
	SetBody(v *GetTenantCheckAvailableResponseBody) *GetTenantCheckAvailableResponse
	GetBody() *GetTenantCheckAvailableResponseBody
}

type GetTenantCheckAvailableResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTenantCheckAvailableResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTenantCheckAvailableResponse) String() string {
	return dara.Prettify(s)
}

func (s GetTenantCheckAvailableResponse) GoString() string {
	return s.String()
}

func (s *GetTenantCheckAvailableResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetTenantCheckAvailableResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetTenantCheckAvailableResponse) GetBody() *GetTenantCheckAvailableResponseBody {
	return s.Body
}

func (s *GetTenantCheckAvailableResponse) SetHeaders(v map[string]*string) *GetTenantCheckAvailableResponse {
	s.Headers = v
	return s
}

func (s *GetTenantCheckAvailableResponse) SetStatusCode(v int32) *GetTenantCheckAvailableResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTenantCheckAvailableResponse) SetBody(v *GetTenantCheckAvailableResponseBody) *GetTenantCheckAvailableResponse {
	s.Body = v
	return s
}

func (s *GetTenantCheckAvailableResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
