// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListStreamingDataSourcesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListStreamingDataSourcesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListStreamingDataSourcesResponse
	GetStatusCode() *int32
	SetBody(v *ListStreamingDataSourcesResponseBody) *ListStreamingDataSourcesResponse
	GetBody() *ListStreamingDataSourcesResponseBody
}

type ListStreamingDataSourcesResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListStreamingDataSourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListStreamingDataSourcesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListStreamingDataSourcesResponse) GoString() string {
	return s.String()
}

func (s *ListStreamingDataSourcesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListStreamingDataSourcesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListStreamingDataSourcesResponse) GetBody() *ListStreamingDataSourcesResponseBody {
	return s.Body
}

func (s *ListStreamingDataSourcesResponse) SetHeaders(v map[string]*string) *ListStreamingDataSourcesResponse {
	s.Headers = v
	return s
}

func (s *ListStreamingDataSourcesResponse) SetStatusCode(v int32) *ListStreamingDataSourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListStreamingDataSourcesResponse) SetBody(v *ListStreamingDataSourcesResponseBody) *ListStreamingDataSourcesResponse {
	s.Body = v
	return s
}

func (s *ListStreamingDataSourcesResponse) Validate() error {
	return dara.Validate(s)
}
