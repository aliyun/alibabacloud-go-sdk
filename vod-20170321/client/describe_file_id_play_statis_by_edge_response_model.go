// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFileIdPlayStatisByEdgeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeFileIdPlayStatisByEdgeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeFileIdPlayStatisByEdgeResponse
	GetStatusCode() *int32
	SetBody(v *DescribeFileIdPlayStatisByEdgeResponseBody) *DescribeFileIdPlayStatisByEdgeResponse
	GetBody() *DescribeFileIdPlayStatisByEdgeResponseBody
}

type DescribeFileIdPlayStatisByEdgeResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeFileIdPlayStatisByEdgeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeFileIdPlayStatisByEdgeResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeFileIdPlayStatisByEdgeResponse) GoString() string {
	return s.String()
}

func (s *DescribeFileIdPlayStatisByEdgeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeFileIdPlayStatisByEdgeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeFileIdPlayStatisByEdgeResponse) GetBody() *DescribeFileIdPlayStatisByEdgeResponseBody {
	return s.Body
}

func (s *DescribeFileIdPlayStatisByEdgeResponse) SetHeaders(v map[string]*string) *DescribeFileIdPlayStatisByEdgeResponse {
	s.Headers = v
	return s
}

func (s *DescribeFileIdPlayStatisByEdgeResponse) SetStatusCode(v int32) *DescribeFileIdPlayStatisByEdgeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeFileIdPlayStatisByEdgeResponse) SetBody(v *DescribeFileIdPlayStatisByEdgeResponseBody) *DescribeFileIdPlayStatisByEdgeResponse {
	s.Body = v
	return s
}

func (s *DescribeFileIdPlayStatisByEdgeResponse) Validate() error {
	return dara.Validate(s)
}
