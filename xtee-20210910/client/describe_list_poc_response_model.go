// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeListPocResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeListPocResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeListPocResponse
	GetStatusCode() *int32
	SetBody(v *DescribeListPocResponseBody) *DescribeListPocResponse
	GetBody() *DescribeListPocResponseBody
}

type DescribeListPocResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeListPocResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeListPocResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeListPocResponse) GoString() string {
	return s.String()
}

func (s *DescribeListPocResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeListPocResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeListPocResponse) GetBody() *DescribeListPocResponseBody {
	return s.Body
}

func (s *DescribeListPocResponse) SetHeaders(v map[string]*string) *DescribeListPocResponse {
	s.Headers = v
	return s
}

func (s *DescribeListPocResponse) SetStatusCode(v int32) *DescribeListPocResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeListPocResponse) SetBody(v *DescribeListPocResponseBody) *DescribeListPocResponse {
	s.Body = v
	return s
}

func (s *DescribeListPocResponse) Validate() error {
	return dara.Validate(s)
}
