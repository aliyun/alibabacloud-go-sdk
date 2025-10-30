// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNodePoolVulsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeNodePoolVulsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeNodePoolVulsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeNodePoolVulsResponseBody) *DescribeNodePoolVulsResponse
	GetBody() *DescribeNodePoolVulsResponseBody
}

type DescribeNodePoolVulsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeNodePoolVulsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeNodePoolVulsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeNodePoolVulsResponse) GoString() string {
	return s.String()
}

func (s *DescribeNodePoolVulsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeNodePoolVulsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeNodePoolVulsResponse) GetBody() *DescribeNodePoolVulsResponseBody {
	return s.Body
}

func (s *DescribeNodePoolVulsResponse) SetHeaders(v map[string]*string) *DescribeNodePoolVulsResponse {
	s.Headers = v
	return s
}

func (s *DescribeNodePoolVulsResponse) SetStatusCode(v int32) *DescribeNodePoolVulsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeNodePoolVulsResponse) SetBody(v *DescribeNodePoolVulsResponseBody) *DescribeNodePoolVulsResponse {
	s.Body = v
	return s
}

func (s *DescribeNodePoolVulsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
