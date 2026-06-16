// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDocParserJobResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDocParserJobResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDocParserJobResultResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDocParserJobResultResponseBody) *DescribeDocParserJobResultResponse
	GetBody() *DescribeDocParserJobResultResponseBody
}

type DescribeDocParserJobResultResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDocParserJobResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDocParserJobResultResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDocParserJobResultResponse) GoString() string {
	return s.String()
}

func (s *DescribeDocParserJobResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDocParserJobResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDocParserJobResultResponse) GetBody() *DescribeDocParserJobResultResponseBody {
	return s.Body
}

func (s *DescribeDocParserJobResultResponse) SetHeaders(v map[string]*string) *DescribeDocParserJobResultResponse {
	s.Headers = v
	return s
}

func (s *DescribeDocParserJobResultResponse) SetStatusCode(v int32) *DescribeDocParserJobResultResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDocParserJobResultResponse) SetBody(v *DescribeDocParserJobResultResponseBody) *DescribeDocParserJobResultResponse {
	s.Body = v
	return s
}

func (s *DescribeDocParserJobResultResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
