// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWebApplicationsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListWebApplicationsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListWebApplicationsResponse
	GetStatusCode() *int32
	SetBody(v *ListWebApplicationsBody) *ListWebApplicationsResponse
	GetBody() *ListWebApplicationsBody
}

type ListWebApplicationsResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListWebApplicationsBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListWebApplicationsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListWebApplicationsResponse) GoString() string {
	return s.String()
}

func (s *ListWebApplicationsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListWebApplicationsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListWebApplicationsResponse) GetBody() *ListWebApplicationsBody {
	return s.Body
}

func (s *ListWebApplicationsResponse) SetHeaders(v map[string]*string) *ListWebApplicationsResponse {
	s.Headers = v
	return s
}

func (s *ListWebApplicationsResponse) SetStatusCode(v int32) *ListWebApplicationsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListWebApplicationsResponse) SetBody(v *ListWebApplicationsBody) *ListWebApplicationsResponse {
	s.Body = v
	return s
}

func (s *ListWebApplicationsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
