// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBLogFilesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDBLogFilesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDBLogFilesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDBLogFilesResponseBody) *DescribeDBLogFilesResponse
	GetBody() *DescribeDBLogFilesResponseBody
}

type DescribeDBLogFilesResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDBLogFilesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDBLogFilesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBLogFilesResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBLogFilesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDBLogFilesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDBLogFilesResponse) GetBody() *DescribeDBLogFilesResponseBody {
	return s.Body
}

func (s *DescribeDBLogFilesResponse) SetHeaders(v map[string]*string) *DescribeDBLogFilesResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBLogFilesResponse) SetStatusCode(v int32) *DescribeDBLogFilesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDBLogFilesResponse) SetBody(v *DescribeDBLogFilesResponseBody) *DescribeDBLogFilesResponse {
	s.Body = v
	return s
}

func (s *DescribeDBLogFilesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
