// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDocParserJobStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDocParserJobStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDocParserJobStatusResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDocParserJobStatusResponseBody) *DescribeDocParserJobStatusResponse
	GetBody() *DescribeDocParserJobStatusResponseBody
}

type DescribeDocParserJobStatusResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDocParserJobStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDocParserJobStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDocParserJobStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeDocParserJobStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDocParserJobStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDocParserJobStatusResponse) GetBody() *DescribeDocParserJobStatusResponseBody {
	return s.Body
}

func (s *DescribeDocParserJobStatusResponse) SetHeaders(v map[string]*string) *DescribeDocParserJobStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeDocParserJobStatusResponse) SetStatusCode(v int32) *DescribeDocParserJobStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDocParserJobStatusResponse) SetBody(v *DescribeDocParserJobStatusResponseBody) *DescribeDocParserJobStatusResponse {
	s.Body = v
	return s
}

func (s *DescribeDocParserJobStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
