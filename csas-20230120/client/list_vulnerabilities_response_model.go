// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVulnerabilitiesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListVulnerabilitiesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListVulnerabilitiesResponse
	GetStatusCode() *int32
	SetBody(v *ListVulnerabilitiesResponseBody) *ListVulnerabilitiesResponse
	GetBody() *ListVulnerabilitiesResponseBody
}

type ListVulnerabilitiesResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListVulnerabilitiesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListVulnerabilitiesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListVulnerabilitiesResponse) GoString() string {
	return s.String()
}

func (s *ListVulnerabilitiesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListVulnerabilitiesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListVulnerabilitiesResponse) GetBody() *ListVulnerabilitiesResponseBody {
	return s.Body
}

func (s *ListVulnerabilitiesResponse) SetHeaders(v map[string]*string) *ListVulnerabilitiesResponse {
	s.Headers = v
	return s
}

func (s *ListVulnerabilitiesResponse) SetStatusCode(v int32) *ListVulnerabilitiesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListVulnerabilitiesResponse) SetBody(v *ListVulnerabilitiesResponseBody) *ListVulnerabilitiesResponse {
	s.Body = v
	return s
}

func (s *ListVulnerabilitiesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
