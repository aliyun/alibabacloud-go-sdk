// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAppliedAdvicesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAppliedAdvicesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAppliedAdvicesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAppliedAdvicesResponseBody) *DescribeAppliedAdvicesResponse
	GetBody() *DescribeAppliedAdvicesResponseBody
}

type DescribeAppliedAdvicesResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAppliedAdvicesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAppliedAdvicesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppliedAdvicesResponse) GoString() string {
	return s.String()
}

func (s *DescribeAppliedAdvicesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAppliedAdvicesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAppliedAdvicesResponse) GetBody() *DescribeAppliedAdvicesResponseBody {
	return s.Body
}

func (s *DescribeAppliedAdvicesResponse) SetHeaders(v map[string]*string) *DescribeAppliedAdvicesResponse {
	s.Headers = v
	return s
}

func (s *DescribeAppliedAdvicesResponse) SetStatusCode(v int32) *DescribeAppliedAdvicesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAppliedAdvicesResponse) SetBody(v *DescribeAppliedAdvicesResponseBody) *DescribeAppliedAdvicesResponse {
	s.Body = v
	return s
}

func (s *DescribeAppliedAdvicesResponse) Validate() error {
	return dara.Validate(s)
}
