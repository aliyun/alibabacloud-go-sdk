// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGadInstanceDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeGadInstanceDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeGadInstanceDetailResponse
	GetStatusCode() *int32
	SetBody(v *DescribeGadInstanceDetailResponseBody) *DescribeGadInstanceDetailResponse
	GetBody() *DescribeGadInstanceDetailResponseBody
}

type DescribeGadInstanceDetailResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeGadInstanceDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeGadInstanceDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeGadInstanceDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribeGadInstanceDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeGadInstanceDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeGadInstanceDetailResponse) GetBody() *DescribeGadInstanceDetailResponseBody {
	return s.Body
}

func (s *DescribeGadInstanceDetailResponse) SetHeaders(v map[string]*string) *DescribeGadInstanceDetailResponse {
	s.Headers = v
	return s
}

func (s *DescribeGadInstanceDetailResponse) SetStatusCode(v int32) *DescribeGadInstanceDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeGadInstanceDetailResponse) SetBody(v *DescribeGadInstanceDetailResponseBody) *DescribeGadInstanceDetailResponse {
	s.Body = v
	return s
}

func (s *DescribeGadInstanceDetailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
