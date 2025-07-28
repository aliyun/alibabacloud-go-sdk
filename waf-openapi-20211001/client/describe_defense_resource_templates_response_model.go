// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDefenseResourceTemplatesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDefenseResourceTemplatesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDefenseResourceTemplatesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDefenseResourceTemplatesResponseBody) *DescribeDefenseResourceTemplatesResponse
	GetBody() *DescribeDefenseResourceTemplatesResponseBody
}

type DescribeDefenseResourceTemplatesResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDefenseResourceTemplatesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDefenseResourceTemplatesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDefenseResourceTemplatesResponse) GoString() string {
	return s.String()
}

func (s *DescribeDefenseResourceTemplatesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDefenseResourceTemplatesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDefenseResourceTemplatesResponse) GetBody() *DescribeDefenseResourceTemplatesResponseBody {
	return s.Body
}

func (s *DescribeDefenseResourceTemplatesResponse) SetHeaders(v map[string]*string) *DescribeDefenseResourceTemplatesResponse {
	s.Headers = v
	return s
}

func (s *DescribeDefenseResourceTemplatesResponse) SetStatusCode(v int32) *DescribeDefenseResourceTemplatesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDefenseResourceTemplatesResponse) SetBody(v *DescribeDefenseResourceTemplatesResponseBody) *DescribeDefenseResourceTemplatesResponse {
	s.Body = v
	return s
}

func (s *DescribeDefenseResourceTemplatesResponse) Validate() error {
	return dara.Validate(s)
}
