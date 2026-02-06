// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryJobsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryJobsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryJobsResponse
	GetStatusCode() *int32
	SetBody(v *QueryJobsResponseBody) *QueryJobsResponse
	GetBody() *QueryJobsResponseBody
}

type QueryJobsResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryJobsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryJobsResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryJobsResponse) GoString() string {
	return s.String()
}

func (s *QueryJobsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryJobsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryJobsResponse) GetBody() *QueryJobsResponseBody {
	return s.Body
}

func (s *QueryJobsResponse) SetHeaders(v map[string]*string) *QueryJobsResponse {
	s.Headers = v
	return s
}

func (s *QueryJobsResponse) SetStatusCode(v int32) *QueryJobsResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryJobsResponse) SetBody(v *QueryJobsResponseBody) *QueryJobsResponse {
	s.Body = v
	return s
}

func (s *QueryJobsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
