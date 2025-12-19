// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRecordFileResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeRecordFileResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeRecordFileResponse
	GetStatusCode() *int32
	SetBody(v *DescribeRecordFileResponseBody) *DescribeRecordFileResponse
	GetBody() *DescribeRecordFileResponseBody
}

type DescribeRecordFileResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRecordFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRecordFileResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeRecordFileResponse) GoString() string {
	return s.String()
}

func (s *DescribeRecordFileResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeRecordFileResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeRecordFileResponse) GetBody() *DescribeRecordFileResponseBody {
	return s.Body
}

func (s *DescribeRecordFileResponse) SetHeaders(v map[string]*string) *DescribeRecordFileResponse {
	s.Headers = v
	return s
}

func (s *DescribeRecordFileResponse) SetStatusCode(v int32) *DescribeRecordFileResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRecordFileResponse) SetBody(v *DescribeRecordFileResponseBody) *DescribeRecordFileResponse {
	s.Body = v
	return s
}

func (s *DescribeRecordFileResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
