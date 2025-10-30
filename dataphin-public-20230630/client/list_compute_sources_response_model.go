// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListComputeSourcesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListComputeSourcesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListComputeSourcesResponse
	GetStatusCode() *int32
	SetBody(v *ListComputeSourcesResponseBody) *ListComputeSourcesResponse
	GetBody() *ListComputeSourcesResponseBody
}

type ListComputeSourcesResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListComputeSourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListComputeSourcesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListComputeSourcesResponse) GoString() string {
	return s.String()
}

func (s *ListComputeSourcesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListComputeSourcesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListComputeSourcesResponse) GetBody() *ListComputeSourcesResponseBody {
	return s.Body
}

func (s *ListComputeSourcesResponse) SetHeaders(v map[string]*string) *ListComputeSourcesResponse {
	s.Headers = v
	return s
}

func (s *ListComputeSourcesResponse) SetStatusCode(v int32) *ListComputeSourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListComputeSourcesResponse) SetBody(v *ListComputeSourcesResponseBody) *ListComputeSourcesResponse {
	s.Body = v
	return s
}

func (s *ListComputeSourcesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
