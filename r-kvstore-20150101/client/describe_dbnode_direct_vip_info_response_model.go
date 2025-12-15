// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBNodeDirectVipInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDBNodeDirectVipInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDBNodeDirectVipInfoResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDBNodeDirectVipInfoResponseBody) *DescribeDBNodeDirectVipInfoResponse
	GetBody() *DescribeDBNodeDirectVipInfoResponseBody
}

type DescribeDBNodeDirectVipInfoResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDBNodeDirectVipInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDBNodeDirectVipInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBNodeDirectVipInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBNodeDirectVipInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDBNodeDirectVipInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDBNodeDirectVipInfoResponse) GetBody() *DescribeDBNodeDirectVipInfoResponseBody {
	return s.Body
}

func (s *DescribeDBNodeDirectVipInfoResponse) SetHeaders(v map[string]*string) *DescribeDBNodeDirectVipInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBNodeDirectVipInfoResponse) SetStatusCode(v int32) *DescribeDBNodeDirectVipInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDBNodeDirectVipInfoResponse) SetBody(v *DescribeDBNodeDirectVipInfoResponseBody) *DescribeDBNodeDirectVipInfoResponse {
	s.Body = v
	return s
}

func (s *DescribeDBNodeDirectVipInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
