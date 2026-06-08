// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMyApplicationsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListMyApplicationsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListMyApplicationsResponse
	GetStatusCode() *int32
	SetBody(v *ListMyApplicationsResponseBody) *ListMyApplicationsResponse
	GetBody() *ListMyApplicationsResponseBody
}

type ListMyApplicationsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListMyApplicationsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMyApplicationsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListMyApplicationsResponse) GoString() string {
	return s.String()
}

func (s *ListMyApplicationsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListMyApplicationsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListMyApplicationsResponse) GetBody() *ListMyApplicationsResponseBody {
	return s.Body
}

func (s *ListMyApplicationsResponse) SetHeaders(v map[string]*string) *ListMyApplicationsResponse {
	s.Headers = v
	return s
}

func (s *ListMyApplicationsResponse) SetStatusCode(v int32) *ListMyApplicationsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMyApplicationsResponse) SetBody(v *ListMyApplicationsResponseBody) *ListMyApplicationsResponse {
	s.Body = v
	return s
}

func (s *ListMyApplicationsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
