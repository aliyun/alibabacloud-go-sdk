// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDefenseResourceNamesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDefenseResourceNamesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDefenseResourceNamesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDefenseResourceNamesResponseBody) *DescribeDefenseResourceNamesResponse
	GetBody() *DescribeDefenseResourceNamesResponseBody
}

type DescribeDefenseResourceNamesResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDefenseResourceNamesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDefenseResourceNamesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDefenseResourceNamesResponse) GoString() string {
	return s.String()
}

func (s *DescribeDefenseResourceNamesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDefenseResourceNamesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDefenseResourceNamesResponse) GetBody() *DescribeDefenseResourceNamesResponseBody {
	return s.Body
}

func (s *DescribeDefenseResourceNamesResponse) SetHeaders(v map[string]*string) *DescribeDefenseResourceNamesResponse {
	s.Headers = v
	return s
}

func (s *DescribeDefenseResourceNamesResponse) SetStatusCode(v int32) *DescribeDefenseResourceNamesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDefenseResourceNamesResponse) SetBody(v *DescribeDefenseResourceNamesResponseBody) *DescribeDefenseResourceNamesResponse {
	s.Body = v
	return s
}

func (s *DescribeDefenseResourceNamesResponse) Validate() error {
	return dara.Validate(s)
}
