// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDefenseTemplatesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDefenseTemplatesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDefenseTemplatesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDefenseTemplatesResponseBody) *DescribeDefenseTemplatesResponse
	GetBody() *DescribeDefenseTemplatesResponseBody
}

type DescribeDefenseTemplatesResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDefenseTemplatesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDefenseTemplatesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDefenseTemplatesResponse) GoString() string {
	return s.String()
}

func (s *DescribeDefenseTemplatesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDefenseTemplatesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDefenseTemplatesResponse) GetBody() *DescribeDefenseTemplatesResponseBody {
	return s.Body
}

func (s *DescribeDefenseTemplatesResponse) SetHeaders(v map[string]*string) *DescribeDefenseTemplatesResponse {
	s.Headers = v
	return s
}

func (s *DescribeDefenseTemplatesResponse) SetStatusCode(v int32) *DescribeDefenseTemplatesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDefenseTemplatesResponse) SetBody(v *DescribeDefenseTemplatesResponseBody) *DescribeDefenseTemplatesResponse {
	s.Body = v
	return s
}

func (s *DescribeDefenseTemplatesResponse) Validate() error {
	return dara.Validate(s)
}
