// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPrivateAccessApplicationsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListPrivateAccessApplicationsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListPrivateAccessApplicationsResponse
	GetStatusCode() *int32
	SetBody(v *ListPrivateAccessApplicationsResponseBody) *ListPrivateAccessApplicationsResponse
	GetBody() *ListPrivateAccessApplicationsResponseBody
}

type ListPrivateAccessApplicationsResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPrivateAccessApplicationsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPrivateAccessApplicationsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListPrivateAccessApplicationsResponse) GoString() string {
	return s.String()
}

func (s *ListPrivateAccessApplicationsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListPrivateAccessApplicationsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListPrivateAccessApplicationsResponse) GetBody() *ListPrivateAccessApplicationsResponseBody {
	return s.Body
}

func (s *ListPrivateAccessApplicationsResponse) SetHeaders(v map[string]*string) *ListPrivateAccessApplicationsResponse {
	s.Headers = v
	return s
}

func (s *ListPrivateAccessApplicationsResponse) SetStatusCode(v int32) *ListPrivateAccessApplicationsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPrivateAccessApplicationsResponse) SetBody(v *ListPrivateAccessApplicationsResponseBody) *ListPrivateAccessApplicationsResponse {
	s.Body = v
	return s
}

func (s *ListPrivateAccessApplicationsResponse) Validate() error {
	return dara.Validate(s)
}
