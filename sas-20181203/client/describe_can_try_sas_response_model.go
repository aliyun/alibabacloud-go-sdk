// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCanTrySasResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCanTrySasResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCanTrySasResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCanTrySasResponseBody) *DescribeCanTrySasResponse
	GetBody() *DescribeCanTrySasResponseBody
}

type DescribeCanTrySasResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCanTrySasResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCanTrySasResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCanTrySasResponse) GoString() string {
	return s.String()
}

func (s *DescribeCanTrySasResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCanTrySasResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCanTrySasResponse) GetBody() *DescribeCanTrySasResponseBody {
	return s.Body
}

func (s *DescribeCanTrySasResponse) SetHeaders(v map[string]*string) *DescribeCanTrySasResponse {
	s.Headers = v
	return s
}

func (s *DescribeCanTrySasResponse) SetStatusCode(v int32) *DescribeCanTrySasResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCanTrySasResponse) SetBody(v *DescribeCanTrySasResponseBody) *DescribeCanTrySasResponse {
	s.Body = v
	return s
}

func (s *DescribeCanTrySasResponse) Validate() error {
	return dara.Validate(s)
}
