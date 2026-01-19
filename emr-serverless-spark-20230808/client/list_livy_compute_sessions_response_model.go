// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLivyComputeSessionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListLivyComputeSessionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListLivyComputeSessionsResponse
	GetStatusCode() *int32
	SetBody(v *ListLivyComputeSessionsResponseBody) *ListLivyComputeSessionsResponse
	GetBody() *ListLivyComputeSessionsResponseBody
}

type ListLivyComputeSessionsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListLivyComputeSessionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListLivyComputeSessionsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListLivyComputeSessionsResponse) GoString() string {
	return s.String()
}

func (s *ListLivyComputeSessionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListLivyComputeSessionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListLivyComputeSessionsResponse) GetBody() *ListLivyComputeSessionsResponseBody {
	return s.Body
}

func (s *ListLivyComputeSessionsResponse) SetHeaders(v map[string]*string) *ListLivyComputeSessionsResponse {
	s.Headers = v
	return s
}

func (s *ListLivyComputeSessionsResponse) SetStatusCode(v int32) *ListLivyComputeSessionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListLivyComputeSessionsResponse) SetBody(v *ListLivyComputeSessionsResponseBody) *ListLivyComputeSessionsResponse {
	s.Body = v
	return s
}

func (s *ListLivyComputeSessionsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
