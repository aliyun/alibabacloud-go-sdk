// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMmsDataSourcesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListMmsDataSourcesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListMmsDataSourcesResponse
	GetStatusCode() *int32
	SetBody(v *ListMmsDataSourcesResponseBody) *ListMmsDataSourcesResponse
	GetBody() *ListMmsDataSourcesResponseBody
}

type ListMmsDataSourcesResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListMmsDataSourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMmsDataSourcesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListMmsDataSourcesResponse) GoString() string {
	return s.String()
}

func (s *ListMmsDataSourcesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListMmsDataSourcesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListMmsDataSourcesResponse) GetBody() *ListMmsDataSourcesResponseBody {
	return s.Body
}

func (s *ListMmsDataSourcesResponse) SetHeaders(v map[string]*string) *ListMmsDataSourcesResponse {
	s.Headers = v
	return s
}

func (s *ListMmsDataSourcesResponse) SetStatusCode(v int32) *ListMmsDataSourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMmsDataSourcesResponse) SetBody(v *ListMmsDataSourcesResponseBody) *ListMmsDataSourcesResponse {
	s.Body = v
	return s
}

func (s *ListMmsDataSourcesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
