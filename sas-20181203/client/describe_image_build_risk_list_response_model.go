// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeImageBuildRiskListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeImageBuildRiskListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeImageBuildRiskListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeImageBuildRiskListResponseBody) *DescribeImageBuildRiskListResponse
	GetBody() *DescribeImageBuildRiskListResponseBody
}

type DescribeImageBuildRiskListResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeImageBuildRiskListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeImageBuildRiskListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageBuildRiskListResponse) GoString() string {
	return s.String()
}

func (s *DescribeImageBuildRiskListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeImageBuildRiskListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeImageBuildRiskListResponse) GetBody() *DescribeImageBuildRiskListResponseBody {
	return s.Body
}

func (s *DescribeImageBuildRiskListResponse) SetHeaders(v map[string]*string) *DescribeImageBuildRiskListResponse {
	s.Headers = v
	return s
}

func (s *DescribeImageBuildRiskListResponse) SetStatusCode(v int32) *DescribeImageBuildRiskListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeImageBuildRiskListResponse) SetBody(v *DescribeImageBuildRiskListResponseBody) *DescribeImageBuildRiskListResponse {
	s.Body = v
	return s
}

func (s *DescribeImageBuildRiskListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
