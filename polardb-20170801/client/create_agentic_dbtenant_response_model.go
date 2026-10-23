// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAgenticDBTenantResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateAgenticDBTenantResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateAgenticDBTenantResponse
	GetStatusCode() *int32
	SetBody(v *CreateAgenticDBTenantResponseBody) *CreateAgenticDBTenantResponse
	GetBody() *CreateAgenticDBTenantResponseBody
}

type CreateAgenticDBTenantResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAgenticDBTenantResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAgenticDBTenantResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateAgenticDBTenantResponse) GoString() string {
	return s.String()
}

func (s *CreateAgenticDBTenantResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateAgenticDBTenantResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateAgenticDBTenantResponse) GetBody() *CreateAgenticDBTenantResponseBody {
	return s.Body
}

func (s *CreateAgenticDBTenantResponse) SetHeaders(v map[string]*string) *CreateAgenticDBTenantResponse {
	s.Headers = v
	return s
}

func (s *CreateAgenticDBTenantResponse) SetStatusCode(v int32) *CreateAgenticDBTenantResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAgenticDBTenantResponse) SetBody(v *CreateAgenticDBTenantResponseBody) *CreateAgenticDBTenantResponse {
	s.Body = v
	return s
}

func (s *CreateAgenticDBTenantResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
