// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeActiveSQLRecordsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeActiveSQLRecordsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeActiveSQLRecordsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeActiveSQLRecordsResponseBody) *DescribeActiveSQLRecordsResponse
	GetBody() *DescribeActiveSQLRecordsResponseBody
}

type DescribeActiveSQLRecordsResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeActiveSQLRecordsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeActiveSQLRecordsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeActiveSQLRecordsResponse) GoString() string {
	return s.String()
}

func (s *DescribeActiveSQLRecordsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeActiveSQLRecordsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeActiveSQLRecordsResponse) GetBody() *DescribeActiveSQLRecordsResponseBody {
	return s.Body
}

func (s *DescribeActiveSQLRecordsResponse) SetHeaders(v map[string]*string) *DescribeActiveSQLRecordsResponse {
	s.Headers = v
	return s
}

func (s *DescribeActiveSQLRecordsResponse) SetStatusCode(v int32) *DescribeActiveSQLRecordsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeActiveSQLRecordsResponse) SetBody(v *DescribeActiveSQLRecordsResponseBody) *DescribeActiveSQLRecordsResponse {
	s.Body = v
	return s
}

func (s *DescribeActiveSQLRecordsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
