// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTenantResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetTenantResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetTenantResponse
	GetStatusCode() *int32
	SetBody(v *GetTenantResponseBody) *GetTenantResponse
	GetBody() *GetTenantResponseBody
}

type GetTenantResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTenantResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTenantResponse) String() string {
	return dara.Prettify(s)
}

func (s GetTenantResponse) GoString() string {
	return s.String()
}

func (s *GetTenantResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetTenantResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetTenantResponse) GetBody() *GetTenantResponseBody {
	return s.Body
}

func (s *GetTenantResponse) SetHeaders(v map[string]*string) *GetTenantResponse {
	s.Headers = v
	return s
}

func (s *GetTenantResponse) SetStatusCode(v int32) *GetTenantResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTenantResponse) SetBody(v *GetTenantResponseBody) *GetTenantResponse {
	s.Body = v
	return s
}

func (s *GetTenantResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
