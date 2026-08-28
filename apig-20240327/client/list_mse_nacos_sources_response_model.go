// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMseNacosSourcesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListMseNacosSourcesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListMseNacosSourcesResponse
	GetStatusCode() *int32
	SetBody(v *ListMseNacosSourcesResponseBody) *ListMseNacosSourcesResponse
	GetBody() *ListMseNacosSourcesResponseBody
}

type ListMseNacosSourcesResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListMseNacosSourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMseNacosSourcesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListMseNacosSourcesResponse) GoString() string {
	return s.String()
}

func (s *ListMseNacosSourcesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListMseNacosSourcesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListMseNacosSourcesResponse) GetBody() *ListMseNacosSourcesResponseBody {
	return s.Body
}

func (s *ListMseNacosSourcesResponse) SetHeaders(v map[string]*string) *ListMseNacosSourcesResponse {
	s.Headers = v
	return s
}

func (s *ListMseNacosSourcesResponse) SetStatusCode(v int32) *ListMseNacosSourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMseNacosSourcesResponse) SetBody(v *ListMseNacosSourcesResponseBody) *ListMseNacosSourcesResponse {
	s.Body = v
	return s
}

func (s *ListMseNacosSourcesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
