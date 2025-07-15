// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetApisAuthoritiesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetApisAuthoritiesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetApisAuthoritiesResponse
	GetStatusCode() *int32
	SetBody(v *SetApisAuthoritiesResponseBody) *SetApisAuthoritiesResponse
	GetBody() *SetApisAuthoritiesResponseBody
}

type SetApisAuthoritiesResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetApisAuthoritiesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetApisAuthoritiesResponse) String() string {
	return dara.Prettify(s)
}

func (s SetApisAuthoritiesResponse) GoString() string {
	return s.String()
}

func (s *SetApisAuthoritiesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetApisAuthoritiesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetApisAuthoritiesResponse) GetBody() *SetApisAuthoritiesResponseBody {
	return s.Body
}

func (s *SetApisAuthoritiesResponse) SetHeaders(v map[string]*string) *SetApisAuthoritiesResponse {
	s.Headers = v
	return s
}

func (s *SetApisAuthoritiesResponse) SetStatusCode(v int32) *SetApisAuthoritiesResponse {
	s.StatusCode = &v
	return s
}

func (s *SetApisAuthoritiesResponse) SetBody(v *SetApisAuthoritiesResponseBody) *SetApisAuthoritiesResponse {
	s.Body = v
	return s
}

func (s *SetApisAuthoritiesResponse) Validate() error {
	return dara.Validate(s)
}
