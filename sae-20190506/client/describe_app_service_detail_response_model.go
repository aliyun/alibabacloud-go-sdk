// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAppServiceDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAppServiceDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAppServiceDetailResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAppServiceDetailResponseBody) *DescribeAppServiceDetailResponse
	GetBody() *DescribeAppServiceDetailResponseBody
}

type DescribeAppServiceDetailResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAppServiceDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAppServiceDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppServiceDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribeAppServiceDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAppServiceDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAppServiceDetailResponse) GetBody() *DescribeAppServiceDetailResponseBody {
	return s.Body
}

func (s *DescribeAppServiceDetailResponse) SetHeaders(v map[string]*string) *DescribeAppServiceDetailResponse {
	s.Headers = v
	return s
}

func (s *DescribeAppServiceDetailResponse) SetStatusCode(v int32) *DescribeAppServiceDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAppServiceDetailResponse) SetBody(v *DescribeAppServiceDetailResponseBody) *DescribeAppServiceDetailResponse {
	s.Body = v
	return s
}

func (s *DescribeAppServiceDetailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
