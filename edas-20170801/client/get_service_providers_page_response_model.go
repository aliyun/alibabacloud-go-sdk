// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetServiceProvidersPageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetServiceProvidersPageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetServiceProvidersPageResponse
	GetStatusCode() *int32
	SetBody(v *GetServiceProvidersPageResponseBody) *GetServiceProvidersPageResponse
	GetBody() *GetServiceProvidersPageResponseBody
}

type GetServiceProvidersPageResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetServiceProvidersPageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetServiceProvidersPageResponse) String() string {
	return dara.Prettify(s)
}

func (s GetServiceProvidersPageResponse) GoString() string {
	return s.String()
}

func (s *GetServiceProvidersPageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetServiceProvidersPageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetServiceProvidersPageResponse) GetBody() *GetServiceProvidersPageResponseBody {
	return s.Body
}

func (s *GetServiceProvidersPageResponse) SetHeaders(v map[string]*string) *GetServiceProvidersPageResponse {
	s.Headers = v
	return s
}

func (s *GetServiceProvidersPageResponse) SetStatusCode(v int32) *GetServiceProvidersPageResponse {
	s.StatusCode = &v
	return s
}

func (s *GetServiceProvidersPageResponse) SetBody(v *GetServiceProvidersPageResponseBody) *GetServiceProvidersPageResponse {
	s.Body = v
	return s
}

func (s *GetServiceProvidersPageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
