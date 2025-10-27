// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchDisableJobsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BatchDisableJobsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BatchDisableJobsResponse
	GetStatusCode() *int32
	SetBody(v *BatchDisableJobsResponseBody) *BatchDisableJobsResponse
	GetBody() *BatchDisableJobsResponseBody
}

type BatchDisableJobsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchDisableJobsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchDisableJobsResponse) String() string {
	return dara.Prettify(s)
}

func (s BatchDisableJobsResponse) GoString() string {
	return s.String()
}

func (s *BatchDisableJobsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BatchDisableJobsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BatchDisableJobsResponse) GetBody() *BatchDisableJobsResponseBody {
	return s.Body
}

func (s *BatchDisableJobsResponse) SetHeaders(v map[string]*string) *BatchDisableJobsResponse {
	s.Headers = v
	return s
}

func (s *BatchDisableJobsResponse) SetStatusCode(v int32) *BatchDisableJobsResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchDisableJobsResponse) SetBody(v *BatchDisableJobsResponseBody) *BatchDisableJobsResponse {
	s.Body = v
	return s
}

func (s *BatchDisableJobsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
