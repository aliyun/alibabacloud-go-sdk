// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRemoteADBDataSourcesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListRemoteADBDataSourcesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListRemoteADBDataSourcesResponse
	GetStatusCode() *int32
	SetBody(v *ListRemoteADBDataSourcesResponseBody) *ListRemoteADBDataSourcesResponse
	GetBody() *ListRemoteADBDataSourcesResponseBody
}

type ListRemoteADBDataSourcesResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListRemoteADBDataSourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListRemoteADBDataSourcesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListRemoteADBDataSourcesResponse) GoString() string {
	return s.String()
}

func (s *ListRemoteADBDataSourcesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListRemoteADBDataSourcesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListRemoteADBDataSourcesResponse) GetBody() *ListRemoteADBDataSourcesResponseBody {
	return s.Body
}

func (s *ListRemoteADBDataSourcesResponse) SetHeaders(v map[string]*string) *ListRemoteADBDataSourcesResponse {
	s.Headers = v
	return s
}

func (s *ListRemoteADBDataSourcesResponse) SetStatusCode(v int32) *ListRemoteADBDataSourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRemoteADBDataSourcesResponse) SetBody(v *ListRemoteADBDataSourcesResponseBody) *ListRemoteADBDataSourcesResponse {
	s.Body = v
	return s
}

func (s *ListRemoteADBDataSourcesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
