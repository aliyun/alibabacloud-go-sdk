// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeQuickSaleConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeQuickSaleConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeQuickSaleConfigResponse
	GetStatusCode() *int32
	SetBody(v *DescribeQuickSaleConfigResponseBody) *DescribeQuickSaleConfigResponse
	GetBody() *DescribeQuickSaleConfigResponseBody
}

type DescribeQuickSaleConfigResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeQuickSaleConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeQuickSaleConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeQuickSaleConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeQuickSaleConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeQuickSaleConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeQuickSaleConfigResponse) GetBody() *DescribeQuickSaleConfigResponseBody {
	return s.Body
}

func (s *DescribeQuickSaleConfigResponse) SetHeaders(v map[string]*string) *DescribeQuickSaleConfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeQuickSaleConfigResponse) SetStatusCode(v int32) *DescribeQuickSaleConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeQuickSaleConfigResponse) SetBody(v *DescribeQuickSaleConfigResponseBody) *DescribeQuickSaleConfigResponse {
	s.Body = v
	return s
}

func (s *DescribeQuickSaleConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
