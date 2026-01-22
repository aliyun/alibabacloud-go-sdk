// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListExternalApplicationsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListExternalApplicationsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListExternalApplicationsResponse
	GetStatusCode() *int32
	SetBody(v *ListExternalApplicationsResponseBody) *ListExternalApplicationsResponse
	GetBody() *ListExternalApplicationsResponseBody
}

type ListExternalApplicationsResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListExternalApplicationsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListExternalApplicationsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListExternalApplicationsResponse) GoString() string {
	return s.String()
}

func (s *ListExternalApplicationsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListExternalApplicationsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListExternalApplicationsResponse) GetBody() *ListExternalApplicationsResponseBody {
	return s.Body
}

func (s *ListExternalApplicationsResponse) SetHeaders(v map[string]*string) *ListExternalApplicationsResponse {
	s.Headers = v
	return s
}

func (s *ListExternalApplicationsResponse) SetStatusCode(v int32) *ListExternalApplicationsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListExternalApplicationsResponse) SetBody(v *ListExternalApplicationsResponseBody) *ListExternalApplicationsResponse {
	s.Body = v
	return s
}

func (s *ListExternalApplicationsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
