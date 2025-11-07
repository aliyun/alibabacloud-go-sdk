// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePageFaceVerifyDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePageFaceVerifyDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePageFaceVerifyDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribePageFaceVerifyDataResponseBody) *DescribePageFaceVerifyDataResponse
	GetBody() *DescribePageFaceVerifyDataResponseBody
}

type DescribePageFaceVerifyDataResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePageFaceVerifyDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePageFaceVerifyDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePageFaceVerifyDataResponse) GoString() string {
	return s.String()
}

func (s *DescribePageFaceVerifyDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePageFaceVerifyDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePageFaceVerifyDataResponse) GetBody() *DescribePageFaceVerifyDataResponseBody {
	return s.Body
}

func (s *DescribePageFaceVerifyDataResponse) SetHeaders(v map[string]*string) *DescribePageFaceVerifyDataResponse {
	s.Headers = v
	return s
}

func (s *DescribePageFaceVerifyDataResponse) SetStatusCode(v int32) *DescribePageFaceVerifyDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePageFaceVerifyDataResponse) SetBody(v *DescribePageFaceVerifyDataResponseBody) *DescribePageFaceVerifyDataResponse {
	s.Body = v
	return s
}

func (s *DescribePageFaceVerifyDataResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
