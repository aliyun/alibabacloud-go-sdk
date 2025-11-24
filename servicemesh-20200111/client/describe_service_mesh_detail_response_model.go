// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeServiceMeshDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeServiceMeshDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeServiceMeshDetailResponse
	GetStatusCode() *int32
	SetBody(v *DescribeServiceMeshDetailResponseBody) *DescribeServiceMeshDetailResponse
	GetBody() *DescribeServiceMeshDetailResponseBody
}

type DescribeServiceMeshDetailResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeServiceMeshDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeServiceMeshDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeServiceMeshDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeServiceMeshDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeServiceMeshDetailResponse) GetBody() *DescribeServiceMeshDetailResponseBody {
	return s.Body
}

func (s *DescribeServiceMeshDetailResponse) SetHeaders(v map[string]*string) *DescribeServiceMeshDetailResponse {
	s.Headers = v
	return s
}

func (s *DescribeServiceMeshDetailResponse) SetStatusCode(v int32) *DescribeServiceMeshDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeServiceMeshDetailResponse) SetBody(v *DescribeServiceMeshDetailResponseBody) *DescribeServiceMeshDetailResponse {
	s.Body = v
	return s
}

func (s *DescribeServiceMeshDetailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
