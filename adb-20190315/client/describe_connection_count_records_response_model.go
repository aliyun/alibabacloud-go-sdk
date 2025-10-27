// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeConnectionCountRecordsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeConnectionCountRecordsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeConnectionCountRecordsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeConnectionCountRecordsResponseBody) *DescribeConnectionCountRecordsResponse
	GetBody() *DescribeConnectionCountRecordsResponseBody
}

type DescribeConnectionCountRecordsResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeConnectionCountRecordsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeConnectionCountRecordsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeConnectionCountRecordsResponse) GoString() string {
	return s.String()
}

func (s *DescribeConnectionCountRecordsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeConnectionCountRecordsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeConnectionCountRecordsResponse) GetBody() *DescribeConnectionCountRecordsResponseBody {
	return s.Body
}

func (s *DescribeConnectionCountRecordsResponse) SetHeaders(v map[string]*string) *DescribeConnectionCountRecordsResponse {
	s.Headers = v
	return s
}

func (s *DescribeConnectionCountRecordsResponse) SetStatusCode(v int32) *DescribeConnectionCountRecordsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeConnectionCountRecordsResponse) SetBody(v *DescribeConnectionCountRecordsResponseBody) *DescribeConnectionCountRecordsResponse {
	s.Body = v
	return s
}

func (s *DescribeConnectionCountRecordsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
