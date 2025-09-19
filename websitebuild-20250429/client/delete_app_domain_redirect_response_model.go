// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAppDomainRedirectResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteAppDomainRedirectResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteAppDomainRedirectResponse
	GetStatusCode() *int32
	SetBody(v *DeleteAppDomainRedirectResponseBody) *DeleteAppDomainRedirectResponse
	GetBody() *DeleteAppDomainRedirectResponseBody
}

type DeleteAppDomainRedirectResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAppDomainRedirectResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAppDomainRedirectResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteAppDomainRedirectResponse) GoString() string {
	return s.String()
}

func (s *DeleteAppDomainRedirectResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteAppDomainRedirectResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteAppDomainRedirectResponse) GetBody() *DeleteAppDomainRedirectResponseBody {
	return s.Body
}

func (s *DeleteAppDomainRedirectResponse) SetHeaders(v map[string]*string) *DeleteAppDomainRedirectResponse {
	s.Headers = v
	return s
}

func (s *DeleteAppDomainRedirectResponse) SetStatusCode(v int32) *DeleteAppDomainRedirectResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAppDomainRedirectResponse) SetBody(v *DeleteAppDomainRedirectResponseBody) *DeleteAppDomainRedirectResponse {
	s.Body = v
	return s
}

func (s *DeleteAppDomainRedirectResponse) Validate() error {
	return dara.Validate(s)
}
