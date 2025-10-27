// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeIdcProbeScanResultListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeIdcProbeScanResultListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeIdcProbeScanResultListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeIdcProbeScanResultListResponseBody) *DescribeIdcProbeScanResultListResponse
	GetBody() *DescribeIdcProbeScanResultListResponseBody
}

type DescribeIdcProbeScanResultListResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeIdcProbeScanResultListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeIdcProbeScanResultListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeIdcProbeScanResultListResponse) GoString() string {
	return s.String()
}

func (s *DescribeIdcProbeScanResultListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeIdcProbeScanResultListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeIdcProbeScanResultListResponse) GetBody() *DescribeIdcProbeScanResultListResponseBody {
	return s.Body
}

func (s *DescribeIdcProbeScanResultListResponse) SetHeaders(v map[string]*string) *DescribeIdcProbeScanResultListResponse {
	s.Headers = v
	return s
}

func (s *DescribeIdcProbeScanResultListResponse) SetStatusCode(v int32) *DescribeIdcProbeScanResultListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeIdcProbeScanResultListResponse) SetBody(v *DescribeIdcProbeScanResultListResponseBody) *DescribeIdcProbeScanResultListResponse {
	s.Body = v
	return s
}

func (s *DescribeIdcProbeScanResultListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
