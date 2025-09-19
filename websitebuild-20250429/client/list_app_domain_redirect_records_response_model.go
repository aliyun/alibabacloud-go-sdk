// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAppDomainRedirectRecordsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAppDomainRedirectRecordsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAppDomainRedirectRecordsResponse
	GetStatusCode() *int32
	SetBody(v *ListAppDomainRedirectRecordsResponseBody) *ListAppDomainRedirectRecordsResponse
	GetBody() *ListAppDomainRedirectRecordsResponseBody
}

type ListAppDomainRedirectRecordsResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAppDomainRedirectRecordsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAppDomainRedirectRecordsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAppDomainRedirectRecordsResponse) GoString() string {
	return s.String()
}

func (s *ListAppDomainRedirectRecordsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAppDomainRedirectRecordsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAppDomainRedirectRecordsResponse) GetBody() *ListAppDomainRedirectRecordsResponseBody {
	return s.Body
}

func (s *ListAppDomainRedirectRecordsResponse) SetHeaders(v map[string]*string) *ListAppDomainRedirectRecordsResponse {
	s.Headers = v
	return s
}

func (s *ListAppDomainRedirectRecordsResponse) SetStatusCode(v int32) *ListAppDomainRedirectRecordsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAppDomainRedirectRecordsResponse) SetBody(v *ListAppDomainRedirectRecordsResponseBody) *ListAppDomainRedirectRecordsResponse {
	s.Body = v
	return s
}

func (s *ListAppDomainRedirectRecordsResponse) Validate() error {
	return dara.Validate(s)
}
