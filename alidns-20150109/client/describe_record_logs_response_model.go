// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRecordLogsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeRecordLogsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeRecordLogsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeRecordLogsResponseBody) *DescribeRecordLogsResponse
	GetBody() *DescribeRecordLogsResponseBody
}

type DescribeRecordLogsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRecordLogsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRecordLogsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeRecordLogsResponse) GoString() string {
	return s.String()
}

func (s *DescribeRecordLogsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeRecordLogsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeRecordLogsResponse) GetBody() *DescribeRecordLogsResponseBody {
	return s.Body
}

func (s *DescribeRecordLogsResponse) SetHeaders(v map[string]*string) *DescribeRecordLogsResponse {
	s.Headers = v
	return s
}

func (s *DescribeRecordLogsResponse) SetStatusCode(v int32) *DescribeRecordLogsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRecordLogsResponse) SetBody(v *DescribeRecordLogsResponseBody) *DescribeRecordLogsResponse {
	s.Body = v
	return s
}

func (s *DescribeRecordLogsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
