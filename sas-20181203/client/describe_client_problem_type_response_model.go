// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeClientProblemTypeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeClientProblemTypeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeClientProblemTypeResponse
	GetStatusCode() *int32
	SetBody(v *DescribeClientProblemTypeResponseBody) *DescribeClientProblemTypeResponse
	GetBody() *DescribeClientProblemTypeResponseBody
}

type DescribeClientProblemTypeResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeClientProblemTypeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeClientProblemTypeResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeClientProblemTypeResponse) GoString() string {
	return s.String()
}

func (s *DescribeClientProblemTypeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeClientProblemTypeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeClientProblemTypeResponse) GetBody() *DescribeClientProblemTypeResponseBody {
	return s.Body
}

func (s *DescribeClientProblemTypeResponse) SetHeaders(v map[string]*string) *DescribeClientProblemTypeResponse {
	s.Headers = v
	return s
}

func (s *DescribeClientProblemTypeResponse) SetStatusCode(v int32) *DescribeClientProblemTypeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeClientProblemTypeResponse) SetBody(v *DescribeClientProblemTypeResponseBody) *DescribeClientProblemTypeResponse {
	s.Body = v
	return s
}

func (s *DescribeClientProblemTypeResponse) Validate() error {
	return dara.Validate(s)
}
