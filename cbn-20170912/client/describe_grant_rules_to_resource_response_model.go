// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGrantRulesToResourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeGrantRulesToResourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeGrantRulesToResourceResponse
	GetStatusCode() *int32
	SetBody(v *DescribeGrantRulesToResourceResponseBody) *DescribeGrantRulesToResourceResponse
	GetBody() *DescribeGrantRulesToResourceResponseBody
}

type DescribeGrantRulesToResourceResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeGrantRulesToResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeGrantRulesToResourceResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeGrantRulesToResourceResponse) GoString() string {
	return s.String()
}

func (s *DescribeGrantRulesToResourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeGrantRulesToResourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeGrantRulesToResourceResponse) GetBody() *DescribeGrantRulesToResourceResponseBody {
	return s.Body
}

func (s *DescribeGrantRulesToResourceResponse) SetHeaders(v map[string]*string) *DescribeGrantRulesToResourceResponse {
	s.Headers = v
	return s
}

func (s *DescribeGrantRulesToResourceResponse) SetStatusCode(v int32) *DescribeGrantRulesToResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeGrantRulesToResourceResponse) SetBody(v *DescribeGrantRulesToResourceResponseBody) *DescribeGrantRulesToResourceResponse {
	s.Body = v
	return s
}

func (s *DescribeGrantRulesToResourceResponse) Validate() error {
	return dara.Validate(s)
}
