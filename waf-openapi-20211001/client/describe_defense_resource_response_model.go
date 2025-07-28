// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDefenseResourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDefenseResourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDefenseResourceResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDefenseResourceResponseBody) *DescribeDefenseResourceResponse
	GetBody() *DescribeDefenseResourceResponseBody
}

type DescribeDefenseResourceResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDefenseResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDefenseResourceResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDefenseResourceResponse) GoString() string {
	return s.String()
}

func (s *DescribeDefenseResourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDefenseResourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDefenseResourceResponse) GetBody() *DescribeDefenseResourceResponseBody {
	return s.Body
}

func (s *DescribeDefenseResourceResponse) SetHeaders(v map[string]*string) *DescribeDefenseResourceResponse {
	s.Headers = v
	return s
}

func (s *DescribeDefenseResourceResponse) SetStatusCode(v int32) *DescribeDefenseResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDefenseResourceResponse) SetBody(v *DescribeDefenseResourceResponseBody) *DescribeDefenseResourceResponse {
	s.Body = v
	return s
}

func (s *DescribeDefenseResourceResponse) Validate() error {
	return dara.Validate(s)
}
