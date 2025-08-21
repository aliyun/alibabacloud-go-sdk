// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePublishStreamStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePublishStreamStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePublishStreamStatusResponse
	GetStatusCode() *int32
	SetBody(v *DescribePublishStreamStatusResponseBody) *DescribePublishStreamStatusResponse
	GetBody() *DescribePublishStreamStatusResponseBody
}

type DescribePublishStreamStatusResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePublishStreamStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePublishStreamStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePublishStreamStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribePublishStreamStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePublishStreamStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePublishStreamStatusResponse) GetBody() *DescribePublishStreamStatusResponseBody {
	return s.Body
}

func (s *DescribePublishStreamStatusResponse) SetHeaders(v map[string]*string) *DescribePublishStreamStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribePublishStreamStatusResponse) SetStatusCode(v int32) *DescribePublishStreamStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePublishStreamStatusResponse) SetBody(v *DescribePublishStreamStatusResponseBody) *DescribePublishStreamStatusResponse {
	s.Body = v
	return s
}

func (s *DescribePublishStreamStatusResponse) Validate() error {
	return dara.Validate(s)
}
