// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetAppsAuthoritiesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetAppsAuthoritiesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetAppsAuthoritiesResponse
	GetStatusCode() *int32
	SetBody(v *SetAppsAuthoritiesResponseBody) *SetAppsAuthoritiesResponse
	GetBody() *SetAppsAuthoritiesResponseBody
}

type SetAppsAuthoritiesResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetAppsAuthoritiesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetAppsAuthoritiesResponse) String() string {
	return dara.Prettify(s)
}

func (s SetAppsAuthoritiesResponse) GoString() string {
	return s.String()
}

func (s *SetAppsAuthoritiesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetAppsAuthoritiesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetAppsAuthoritiesResponse) GetBody() *SetAppsAuthoritiesResponseBody {
	return s.Body
}

func (s *SetAppsAuthoritiesResponse) SetHeaders(v map[string]*string) *SetAppsAuthoritiesResponse {
	s.Headers = v
	return s
}

func (s *SetAppsAuthoritiesResponse) SetStatusCode(v int32) *SetAppsAuthoritiesResponse {
	s.StatusCode = &v
	return s
}

func (s *SetAppsAuthoritiesResponse) SetBody(v *SetAppsAuthoritiesResponseBody) *SetAppsAuthoritiesResponse {
	s.Body = v
	return s
}

func (s *SetAppsAuthoritiesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
