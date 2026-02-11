// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeModelingProjectDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeModelingProjectDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeModelingProjectDetailResponse
	GetStatusCode() *int32
	SetBody(v *DescribeModelingProjectDetailResponseBody) *DescribeModelingProjectDetailResponse
	GetBody() *DescribeModelingProjectDetailResponseBody
}

type DescribeModelingProjectDetailResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeModelingProjectDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeModelingProjectDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeModelingProjectDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribeModelingProjectDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeModelingProjectDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeModelingProjectDetailResponse) GetBody() *DescribeModelingProjectDetailResponseBody {
	return s.Body
}

func (s *DescribeModelingProjectDetailResponse) SetHeaders(v map[string]*string) *DescribeModelingProjectDetailResponse {
	s.Headers = v
	return s
}

func (s *DescribeModelingProjectDetailResponse) SetStatusCode(v int32) *DescribeModelingProjectDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeModelingProjectDetailResponse) SetBody(v *DescribeModelingProjectDetailResponseBody) *DescribeModelingProjectDetailResponse {
	s.Body = v
	return s
}

func (s *DescribeModelingProjectDetailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
