// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCheckWarningDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCheckWarningDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCheckWarningDetailResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCheckWarningDetailResponseBody) *DescribeCheckWarningDetailResponse
	GetBody() *DescribeCheckWarningDetailResponseBody
}

type DescribeCheckWarningDetailResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCheckWarningDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCheckWarningDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCheckWarningDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribeCheckWarningDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCheckWarningDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCheckWarningDetailResponse) GetBody() *DescribeCheckWarningDetailResponseBody {
	return s.Body
}

func (s *DescribeCheckWarningDetailResponse) SetHeaders(v map[string]*string) *DescribeCheckWarningDetailResponse {
	s.Headers = v
	return s
}

func (s *DescribeCheckWarningDetailResponse) SetStatusCode(v int32) *DescribeCheckWarningDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCheckWarningDetailResponse) SetBody(v *DescribeCheckWarningDetailResponseBody) *DescribeCheckWarningDetailResponse {
	s.Body = v
	return s
}

func (s *DescribeCheckWarningDetailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
