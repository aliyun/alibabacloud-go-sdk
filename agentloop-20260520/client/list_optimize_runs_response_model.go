// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOptimizeRunsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListOptimizeRunsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListOptimizeRunsResponse
	GetStatusCode() *int32
	SetBody(v *ListOptimizeRunsResponseBody) *ListOptimizeRunsResponse
	GetBody() *ListOptimizeRunsResponseBody
}

type ListOptimizeRunsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListOptimizeRunsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListOptimizeRunsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListOptimizeRunsResponse) GoString() string {
	return s.String()
}

func (s *ListOptimizeRunsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListOptimizeRunsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListOptimizeRunsResponse) GetBody() *ListOptimizeRunsResponseBody {
	return s.Body
}

func (s *ListOptimizeRunsResponse) SetHeaders(v map[string]*string) *ListOptimizeRunsResponse {
	s.Headers = v
	return s
}

func (s *ListOptimizeRunsResponse) SetStatusCode(v int32) *ListOptimizeRunsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListOptimizeRunsResponse) SetBody(v *ListOptimizeRunsResponseBody) *ListOptimizeRunsResponse {
	s.Body = v
	return s
}

func (s *ListOptimizeRunsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
