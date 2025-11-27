// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRCInvocationResultsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeRCInvocationResultsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeRCInvocationResultsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeRCInvocationResultsResponseBody) *DescribeRCInvocationResultsResponse
	GetBody() *DescribeRCInvocationResultsResponseBody
}

type DescribeRCInvocationResultsResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRCInvocationResultsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRCInvocationResultsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeRCInvocationResultsResponse) GoString() string {
	return s.String()
}

func (s *DescribeRCInvocationResultsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeRCInvocationResultsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeRCInvocationResultsResponse) GetBody() *DescribeRCInvocationResultsResponseBody {
	return s.Body
}

func (s *DescribeRCInvocationResultsResponse) SetHeaders(v map[string]*string) *DescribeRCInvocationResultsResponse {
	s.Headers = v
	return s
}

func (s *DescribeRCInvocationResultsResponse) SetStatusCode(v int32) *DescribeRCInvocationResultsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRCInvocationResultsResponse) SetBody(v *DescribeRCInvocationResultsResponseBody) *DescribeRCInvocationResultsResponse {
	s.Body = v
	return s
}

func (s *DescribeRCInvocationResultsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
