// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeListFaceVerifyInfosResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeListFaceVerifyInfosResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeListFaceVerifyInfosResponse
	GetStatusCode() *int32
	SetBody(v *DescribeListFaceVerifyInfosResponseBody) *DescribeListFaceVerifyInfosResponse
	GetBody() *DescribeListFaceVerifyInfosResponseBody
}

type DescribeListFaceVerifyInfosResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeListFaceVerifyInfosResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeListFaceVerifyInfosResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeListFaceVerifyInfosResponse) GoString() string {
	return s.String()
}

func (s *DescribeListFaceVerifyInfosResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeListFaceVerifyInfosResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeListFaceVerifyInfosResponse) GetBody() *DescribeListFaceVerifyInfosResponseBody {
	return s.Body
}

func (s *DescribeListFaceVerifyInfosResponse) SetHeaders(v map[string]*string) *DescribeListFaceVerifyInfosResponse {
	s.Headers = v
	return s
}

func (s *DescribeListFaceVerifyInfosResponse) SetStatusCode(v int32) *DescribeListFaceVerifyInfosResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeListFaceVerifyInfosResponse) SetBody(v *DescribeListFaceVerifyInfosResponseBody) *DescribeListFaceVerifyInfosResponse {
	s.Body = v
	return s
}

func (s *DescribeListFaceVerifyInfosResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
