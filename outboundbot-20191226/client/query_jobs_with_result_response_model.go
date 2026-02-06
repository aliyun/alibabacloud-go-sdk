// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryJobsWithResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryJobsWithResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryJobsWithResultResponse
	GetStatusCode() *int32
	SetBody(v *QueryJobsWithResultResponseBody) *QueryJobsWithResultResponse
	GetBody() *QueryJobsWithResultResponseBody
}

type QueryJobsWithResultResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryJobsWithResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryJobsWithResultResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryJobsWithResultResponse) GoString() string {
	return s.String()
}

func (s *QueryJobsWithResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryJobsWithResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryJobsWithResultResponse) GetBody() *QueryJobsWithResultResponseBody {
	return s.Body
}

func (s *QueryJobsWithResultResponse) SetHeaders(v map[string]*string) *QueryJobsWithResultResponse {
	s.Headers = v
	return s
}

func (s *QueryJobsWithResultResponse) SetStatusCode(v int32) *QueryJobsWithResultResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryJobsWithResultResponse) SetBody(v *QueryJobsWithResultResponseBody) *QueryJobsWithResultResponse {
	s.Body = v
	return s
}

func (s *QueryJobsWithResultResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
