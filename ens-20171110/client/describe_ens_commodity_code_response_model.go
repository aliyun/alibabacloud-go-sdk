// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEnsCommodityCodeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeEnsCommodityCodeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeEnsCommodityCodeResponse
	GetStatusCode() *int32
	SetBody(v *DescribeEnsCommodityCodeResponseBody) *DescribeEnsCommodityCodeResponse
	GetBody() *DescribeEnsCommodityCodeResponseBody
}

type DescribeEnsCommodityCodeResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeEnsCommodityCodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeEnsCommodityCodeResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeEnsCommodityCodeResponse) GoString() string {
	return s.String()
}

func (s *DescribeEnsCommodityCodeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeEnsCommodityCodeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeEnsCommodityCodeResponse) GetBody() *DescribeEnsCommodityCodeResponseBody {
	return s.Body
}

func (s *DescribeEnsCommodityCodeResponse) SetHeaders(v map[string]*string) *DescribeEnsCommodityCodeResponse {
	s.Headers = v
	return s
}

func (s *DescribeEnsCommodityCodeResponse) SetStatusCode(v int32) *DescribeEnsCommodityCodeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeEnsCommodityCodeResponse) SetBody(v *DescribeEnsCommodityCodeResponseBody) *DescribeEnsCommodityCodeResponse {
	s.Body = v
	return s
}

func (s *DescribeEnsCommodityCodeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
