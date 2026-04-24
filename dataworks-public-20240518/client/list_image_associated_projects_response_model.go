// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListImageAssociatedProjectsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListImageAssociatedProjectsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListImageAssociatedProjectsResponse
	GetStatusCode() *int32
	SetBody(v *ListImageAssociatedProjectsResponseBody) *ListImageAssociatedProjectsResponse
	GetBody() *ListImageAssociatedProjectsResponseBody
}

type ListImageAssociatedProjectsResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListImageAssociatedProjectsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListImageAssociatedProjectsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListImageAssociatedProjectsResponse) GoString() string {
	return s.String()
}

func (s *ListImageAssociatedProjectsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListImageAssociatedProjectsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListImageAssociatedProjectsResponse) GetBody() *ListImageAssociatedProjectsResponseBody {
	return s.Body
}

func (s *ListImageAssociatedProjectsResponse) SetHeaders(v map[string]*string) *ListImageAssociatedProjectsResponse {
	s.Headers = v
	return s
}

func (s *ListImageAssociatedProjectsResponse) SetStatusCode(v int32) *ListImageAssociatedProjectsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListImageAssociatedProjectsResponse) SetBody(v *ListImageAssociatedProjectsResponseBody) *ListImageAssociatedProjectsResponse {
	s.Body = v
	return s
}

func (s *ListImageAssociatedProjectsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
