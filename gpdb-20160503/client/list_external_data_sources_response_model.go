// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListExternalDataSourcesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListExternalDataSourcesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListExternalDataSourcesResponse
	GetStatusCode() *int32
	SetBody(v *ListExternalDataSourcesResponseBody) *ListExternalDataSourcesResponse
	GetBody() *ListExternalDataSourcesResponseBody
}

type ListExternalDataSourcesResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListExternalDataSourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListExternalDataSourcesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListExternalDataSourcesResponse) GoString() string {
	return s.String()
}

func (s *ListExternalDataSourcesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListExternalDataSourcesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListExternalDataSourcesResponse) GetBody() *ListExternalDataSourcesResponseBody {
	return s.Body
}

func (s *ListExternalDataSourcesResponse) SetHeaders(v map[string]*string) *ListExternalDataSourcesResponse {
	s.Headers = v
	return s
}

func (s *ListExternalDataSourcesResponse) SetStatusCode(v int32) *ListExternalDataSourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListExternalDataSourcesResponse) SetBody(v *ListExternalDataSourcesResponseBody) *ListExternalDataSourcesResponse {
	s.Body = v
	return s
}

func (s *ListExternalDataSourcesResponse) Validate() error {
	return dara.Validate(s)
}
