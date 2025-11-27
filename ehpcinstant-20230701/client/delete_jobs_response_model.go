// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteJobsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteJobsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteJobsResponse
	GetStatusCode() *int32
	SetBody(v *DeleteJobsResponseBody) *DeleteJobsResponse
	GetBody() *DeleteJobsResponseBody
}

type DeleteJobsResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteJobsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteJobsResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteJobsResponse) GoString() string {
	return s.String()
}

func (s *DeleteJobsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteJobsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteJobsResponse) GetBody() *DeleteJobsResponseBody {
	return s.Body
}

func (s *DeleteJobsResponse) SetHeaders(v map[string]*string) *DeleteJobsResponse {
	s.Headers = v
	return s
}

func (s *DeleteJobsResponse) SetStatusCode(v int32) *DeleteJobsResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteJobsResponse) SetBody(v *DeleteJobsResponseBody) *DeleteJobsResponse {
	s.Body = v
	return s
}

func (s *DeleteJobsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
