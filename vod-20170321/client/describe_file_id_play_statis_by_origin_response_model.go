// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFileIdPlayStatisByOriginResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeFileIdPlayStatisByOriginResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeFileIdPlayStatisByOriginResponse
	GetStatusCode() *int32
	SetBody(v *DescribeFileIdPlayStatisByOriginResponseBody) *DescribeFileIdPlayStatisByOriginResponse
	GetBody() *DescribeFileIdPlayStatisByOriginResponseBody
}

type DescribeFileIdPlayStatisByOriginResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeFileIdPlayStatisByOriginResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeFileIdPlayStatisByOriginResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeFileIdPlayStatisByOriginResponse) GoString() string {
	return s.String()
}

func (s *DescribeFileIdPlayStatisByOriginResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeFileIdPlayStatisByOriginResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeFileIdPlayStatisByOriginResponse) GetBody() *DescribeFileIdPlayStatisByOriginResponseBody {
	return s.Body
}

func (s *DescribeFileIdPlayStatisByOriginResponse) SetHeaders(v map[string]*string) *DescribeFileIdPlayStatisByOriginResponse {
	s.Headers = v
	return s
}

func (s *DescribeFileIdPlayStatisByOriginResponse) SetStatusCode(v int32) *DescribeFileIdPlayStatisByOriginResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeFileIdPlayStatisByOriginResponse) SetBody(v *DescribeFileIdPlayStatisByOriginResponseBody) *DescribeFileIdPlayStatisByOriginResponse {
	s.Body = v
	return s
}

func (s *DescribeFileIdPlayStatisByOriginResponse) Validate() error {
	return dara.Validate(s)
}
