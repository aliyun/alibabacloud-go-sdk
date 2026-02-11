// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCloudSiemAssetsCounterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCloudSiemAssetsCounterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCloudSiemAssetsCounterResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCloudSiemAssetsCounterResponseBody) *DescribeCloudSiemAssetsCounterResponse
	GetBody() *DescribeCloudSiemAssetsCounterResponseBody
}

type DescribeCloudSiemAssetsCounterResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCloudSiemAssetsCounterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCloudSiemAssetsCounterResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudSiemAssetsCounterResponse) GoString() string {
	return s.String()
}

func (s *DescribeCloudSiemAssetsCounterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCloudSiemAssetsCounterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCloudSiemAssetsCounterResponse) GetBody() *DescribeCloudSiemAssetsCounterResponseBody {
	return s.Body
}

func (s *DescribeCloudSiemAssetsCounterResponse) SetHeaders(v map[string]*string) *DescribeCloudSiemAssetsCounterResponse {
	s.Headers = v
	return s
}

func (s *DescribeCloudSiemAssetsCounterResponse) SetStatusCode(v int32) *DescribeCloudSiemAssetsCounterResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCloudSiemAssetsCounterResponse) SetBody(v *DescribeCloudSiemAssetsCounterResponseBody) *DescribeCloudSiemAssetsCounterResponse {
	s.Body = v
	return s
}

func (s *DescribeCloudSiemAssetsCounterResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
