// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchEnableJobsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BatchEnableJobsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BatchEnableJobsResponse
	GetStatusCode() *int32
	SetBody(v *BatchEnableJobsResponseBody) *BatchEnableJobsResponse
	GetBody() *BatchEnableJobsResponseBody
}

type BatchEnableJobsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchEnableJobsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchEnableJobsResponse) String() string {
	return dara.Prettify(s)
}

func (s BatchEnableJobsResponse) GoString() string {
	return s.String()
}

func (s *BatchEnableJobsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BatchEnableJobsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BatchEnableJobsResponse) GetBody() *BatchEnableJobsResponseBody {
	return s.Body
}

func (s *BatchEnableJobsResponse) SetHeaders(v map[string]*string) *BatchEnableJobsResponse {
	s.Headers = v
	return s
}

func (s *BatchEnableJobsResponse) SetStatusCode(v int32) *BatchEnableJobsResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchEnableJobsResponse) SetBody(v *BatchEnableJobsResponseBody) *BatchEnableJobsResponse {
	s.Body = v
	return s
}

func (s *BatchEnableJobsResponse) Validate() error {
	return dara.Validate(s)
}
