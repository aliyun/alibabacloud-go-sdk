// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeListFaceVerifyDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeListFaceVerifyDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeListFaceVerifyDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeListFaceVerifyDataResponseBody) *DescribeListFaceVerifyDataResponse
	GetBody() *DescribeListFaceVerifyDataResponseBody
}

type DescribeListFaceVerifyDataResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeListFaceVerifyDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeListFaceVerifyDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeListFaceVerifyDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeListFaceVerifyDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeListFaceVerifyDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeListFaceVerifyDataResponse) GetBody() *DescribeListFaceVerifyDataResponseBody {
	return s.Body
}

func (s *DescribeListFaceVerifyDataResponse) SetHeaders(v map[string]*string) *DescribeListFaceVerifyDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeListFaceVerifyDataResponse) SetStatusCode(v int32) *DescribeListFaceVerifyDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeListFaceVerifyDataResponse) SetBody(v *DescribeListFaceVerifyDataResponseBody) *DescribeListFaceVerifyDataResponse {
	s.Body = v
	return s
}

func (s *DescribeListFaceVerifyDataResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
