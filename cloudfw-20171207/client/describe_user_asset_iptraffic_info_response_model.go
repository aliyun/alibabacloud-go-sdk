// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserAssetIPTrafficInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeUserAssetIPTrafficInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeUserAssetIPTrafficInfoResponse
	GetStatusCode() *int32
	SetBody(v *DescribeUserAssetIPTrafficInfoResponseBody) *DescribeUserAssetIPTrafficInfoResponse
	GetBody() *DescribeUserAssetIPTrafficInfoResponseBody
}

type DescribeUserAssetIPTrafficInfoResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeUserAssetIPTrafficInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeUserAssetIPTrafficInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserAssetIPTrafficInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeUserAssetIPTrafficInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeUserAssetIPTrafficInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeUserAssetIPTrafficInfoResponse) GetBody() *DescribeUserAssetIPTrafficInfoResponseBody {
	return s.Body
}

func (s *DescribeUserAssetIPTrafficInfoResponse) SetHeaders(v map[string]*string) *DescribeUserAssetIPTrafficInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeUserAssetIPTrafficInfoResponse) SetStatusCode(v int32) *DescribeUserAssetIPTrafficInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeUserAssetIPTrafficInfoResponse) SetBody(v *DescribeUserAssetIPTrafficInfoResponseBody) *DescribeUserAssetIPTrafficInfoResponse {
	s.Body = v
	return s
}

func (s *DescribeUserAssetIPTrafficInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
