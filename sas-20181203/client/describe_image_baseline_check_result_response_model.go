// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeImageBaselineCheckResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeImageBaselineCheckResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeImageBaselineCheckResultResponse
	GetStatusCode() *int32
	SetBody(v *DescribeImageBaselineCheckResultResponseBody) *DescribeImageBaselineCheckResultResponse
	GetBody() *DescribeImageBaselineCheckResultResponseBody
}

type DescribeImageBaselineCheckResultResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeImageBaselineCheckResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeImageBaselineCheckResultResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageBaselineCheckResultResponse) GoString() string {
	return s.String()
}

func (s *DescribeImageBaselineCheckResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeImageBaselineCheckResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeImageBaselineCheckResultResponse) GetBody() *DescribeImageBaselineCheckResultResponseBody {
	return s.Body
}

func (s *DescribeImageBaselineCheckResultResponse) SetHeaders(v map[string]*string) *DescribeImageBaselineCheckResultResponse {
	s.Headers = v
	return s
}

func (s *DescribeImageBaselineCheckResultResponse) SetStatusCode(v int32) *DescribeImageBaselineCheckResultResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeImageBaselineCheckResultResponse) SetBody(v *DescribeImageBaselineCheckResultResponseBody) *DescribeImageBaselineCheckResultResponse {
	s.Body = v
	return s
}

func (s *DescribeImageBaselineCheckResultResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
