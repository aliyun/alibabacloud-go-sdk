// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeIdentifyTaskStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeIdentifyTaskStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeIdentifyTaskStatusResponse
	GetStatusCode() *int32
	SetBody(v *DescribeIdentifyTaskStatusResponseBody) *DescribeIdentifyTaskStatusResponse
	GetBody() *DescribeIdentifyTaskStatusResponseBody
}

type DescribeIdentifyTaskStatusResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeIdentifyTaskStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeIdentifyTaskStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeIdentifyTaskStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeIdentifyTaskStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeIdentifyTaskStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeIdentifyTaskStatusResponse) GetBody() *DescribeIdentifyTaskStatusResponseBody {
	return s.Body
}

func (s *DescribeIdentifyTaskStatusResponse) SetHeaders(v map[string]*string) *DescribeIdentifyTaskStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeIdentifyTaskStatusResponse) SetStatusCode(v int32) *DescribeIdentifyTaskStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeIdentifyTaskStatusResponse) SetBody(v *DescribeIdentifyTaskStatusResponseBody) *DescribeIdentifyTaskStatusResponse {
	s.Body = v
	return s
}

func (s *DescribeIdentifyTaskStatusResponse) Validate() error {
	return dara.Validate(s)
}
