// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGlobalQuestionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeGlobalQuestionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeGlobalQuestionResponse
	GetStatusCode() *int32
	SetBody(v *DescribeGlobalQuestionResponseBody) *DescribeGlobalQuestionResponse
	GetBody() *DescribeGlobalQuestionResponseBody
}

type DescribeGlobalQuestionResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeGlobalQuestionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeGlobalQuestionResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeGlobalQuestionResponse) GoString() string {
	return s.String()
}

func (s *DescribeGlobalQuestionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeGlobalQuestionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeGlobalQuestionResponse) GetBody() *DescribeGlobalQuestionResponseBody {
	return s.Body
}

func (s *DescribeGlobalQuestionResponse) SetHeaders(v map[string]*string) *DescribeGlobalQuestionResponse {
	s.Headers = v
	return s
}

func (s *DescribeGlobalQuestionResponse) SetStatusCode(v int32) *DescribeGlobalQuestionResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeGlobalQuestionResponse) SetBody(v *DescribeGlobalQuestionResponseBody) *DescribeGlobalQuestionResponse {
	s.Body = v
	return s
}

func (s *DescribeGlobalQuestionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
