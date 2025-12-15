// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRunningLogRecordsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeRunningLogRecordsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeRunningLogRecordsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeRunningLogRecordsResponseBody) *DescribeRunningLogRecordsResponse
	GetBody() *DescribeRunningLogRecordsResponseBody
}

type DescribeRunningLogRecordsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRunningLogRecordsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRunningLogRecordsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeRunningLogRecordsResponse) GoString() string {
	return s.String()
}

func (s *DescribeRunningLogRecordsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeRunningLogRecordsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeRunningLogRecordsResponse) GetBody() *DescribeRunningLogRecordsResponseBody {
	return s.Body
}

func (s *DescribeRunningLogRecordsResponse) SetHeaders(v map[string]*string) *DescribeRunningLogRecordsResponse {
	s.Headers = v
	return s
}

func (s *DescribeRunningLogRecordsResponse) SetStatusCode(v int32) *DescribeRunningLogRecordsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRunningLogRecordsResponse) SetBody(v *DescribeRunningLogRecordsResponseBody) *DescribeRunningLogRecordsResponse {
	s.Body = v
	return s
}

func (s *DescribeRunningLogRecordsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
