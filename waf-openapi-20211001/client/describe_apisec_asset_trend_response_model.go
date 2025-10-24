// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApisecAssetTrendResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeApisecAssetTrendResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeApisecAssetTrendResponse
	GetStatusCode() *int32
	SetBody(v *DescribeApisecAssetTrendResponseBody) *DescribeApisecAssetTrendResponse
	GetBody() *DescribeApisecAssetTrendResponseBody
}

type DescribeApisecAssetTrendResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeApisecAssetTrendResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeApisecAssetTrendResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeApisecAssetTrendResponse) GoString() string {
	return s.String()
}

func (s *DescribeApisecAssetTrendResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeApisecAssetTrendResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeApisecAssetTrendResponse) GetBody() *DescribeApisecAssetTrendResponseBody {
	return s.Body
}

func (s *DescribeApisecAssetTrendResponse) SetHeaders(v map[string]*string) *DescribeApisecAssetTrendResponse {
	s.Headers = v
	return s
}

func (s *DescribeApisecAssetTrendResponse) SetStatusCode(v int32) *DescribeApisecAssetTrendResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeApisecAssetTrendResponse) SetBody(v *DescribeApisecAssetTrendResponseBody) *DescribeApisecAssetTrendResponse {
	s.Body = v
	return s
}

func (s *DescribeApisecAssetTrendResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
