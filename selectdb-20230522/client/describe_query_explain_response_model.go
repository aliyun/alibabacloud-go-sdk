// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeQueryExplainResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeQueryExplainResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeQueryExplainResponse
	GetStatusCode() *int32
	SetBody(v *DescribeQueryExplainResponseBody) *DescribeQueryExplainResponse
	GetBody() *DescribeQueryExplainResponseBody
}

type DescribeQueryExplainResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeQueryExplainResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeQueryExplainResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeQueryExplainResponse) GoString() string {
	return s.String()
}

func (s *DescribeQueryExplainResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeQueryExplainResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeQueryExplainResponse) GetBody() *DescribeQueryExplainResponseBody {
	return s.Body
}

func (s *DescribeQueryExplainResponse) SetHeaders(v map[string]*string) *DescribeQueryExplainResponse {
	s.Headers = v
	return s
}

func (s *DescribeQueryExplainResponse) SetStatusCode(v int32) *DescribeQueryExplainResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeQueryExplainResponse) SetBody(v *DescribeQueryExplainResponseBody) *DescribeQueryExplainResponse {
	s.Body = v
	return s
}

func (s *DescribeQueryExplainResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
