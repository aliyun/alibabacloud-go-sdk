// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRecordingsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeRecordingsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeRecordingsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeRecordingsResponseBody) *DescribeRecordingsResponse
	GetBody() *DescribeRecordingsResponseBody
}

type DescribeRecordingsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRecordingsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRecordingsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeRecordingsResponse) GoString() string {
	return s.String()
}

func (s *DescribeRecordingsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeRecordingsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeRecordingsResponse) GetBody() *DescribeRecordingsResponseBody {
	return s.Body
}

func (s *DescribeRecordingsResponse) SetHeaders(v map[string]*string) *DescribeRecordingsResponse {
	s.Headers = v
	return s
}

func (s *DescribeRecordingsResponse) SetStatusCode(v int32) *DescribeRecordingsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRecordingsResponse) SetBody(v *DescribeRecordingsResponseBody) *DescribeRecordingsResponse {
	s.Body = v
	return s
}

func (s *DescribeRecordingsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
