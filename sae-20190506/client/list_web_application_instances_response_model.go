// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWebApplicationInstancesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListWebApplicationInstancesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListWebApplicationInstancesResponse
	GetStatusCode() *int32
	SetBody(v *ListWebApplicationInstancesBody) *ListWebApplicationInstancesResponse
	GetBody() *ListWebApplicationInstancesBody
}

type ListWebApplicationInstancesResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListWebApplicationInstancesBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListWebApplicationInstancesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListWebApplicationInstancesResponse) GoString() string {
	return s.String()
}

func (s *ListWebApplicationInstancesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListWebApplicationInstancesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListWebApplicationInstancesResponse) GetBody() *ListWebApplicationInstancesBody {
	return s.Body
}

func (s *ListWebApplicationInstancesResponse) SetHeaders(v map[string]*string) *ListWebApplicationInstancesResponse {
	s.Headers = v
	return s
}

func (s *ListWebApplicationInstancesResponse) SetStatusCode(v int32) *ListWebApplicationInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListWebApplicationInstancesResponse) SetBody(v *ListWebApplicationInstancesBody) *ListWebApplicationInstancesResponse {
	s.Body = v
	return s
}

func (s *ListWebApplicationInstancesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
