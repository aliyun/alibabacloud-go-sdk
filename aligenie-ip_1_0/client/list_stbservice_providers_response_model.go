// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSTBServiceProvidersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListSTBServiceProvidersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListSTBServiceProvidersResponse
	GetStatusCode() *int32
	SetBody(v *ListSTBServiceProvidersResponseBody) *ListSTBServiceProvidersResponse
	GetBody() *ListSTBServiceProvidersResponseBody
}

type ListSTBServiceProvidersResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSTBServiceProvidersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSTBServiceProvidersResponse) String() string {
	return dara.Prettify(s)
}

func (s ListSTBServiceProvidersResponse) GoString() string {
	return s.String()
}

func (s *ListSTBServiceProvidersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListSTBServiceProvidersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListSTBServiceProvidersResponse) GetBody() *ListSTBServiceProvidersResponseBody {
	return s.Body
}

func (s *ListSTBServiceProvidersResponse) SetHeaders(v map[string]*string) *ListSTBServiceProvidersResponse {
	s.Headers = v
	return s
}

func (s *ListSTBServiceProvidersResponse) SetStatusCode(v int32) *ListSTBServiceProvidersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSTBServiceProvidersResponse) SetBody(v *ListSTBServiceProvidersResponseBody) *ListSTBServiceProvidersResponse {
	s.Body = v
	return s
}

func (s *ListSTBServiceProvidersResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
