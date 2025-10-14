// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEnsNetSaleDistrictResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeEnsNetSaleDistrictResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeEnsNetSaleDistrictResponse
	GetStatusCode() *int32
	SetBody(v *DescribeEnsNetSaleDistrictResponseBody) *DescribeEnsNetSaleDistrictResponse
	GetBody() *DescribeEnsNetSaleDistrictResponseBody
}

type DescribeEnsNetSaleDistrictResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeEnsNetSaleDistrictResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeEnsNetSaleDistrictResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeEnsNetSaleDistrictResponse) GoString() string {
	return s.String()
}

func (s *DescribeEnsNetSaleDistrictResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeEnsNetSaleDistrictResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeEnsNetSaleDistrictResponse) GetBody() *DescribeEnsNetSaleDistrictResponseBody {
	return s.Body
}

func (s *DescribeEnsNetSaleDistrictResponse) SetHeaders(v map[string]*string) *DescribeEnsNetSaleDistrictResponse {
	s.Headers = v
	return s
}

func (s *DescribeEnsNetSaleDistrictResponse) SetStatusCode(v int32) *DescribeEnsNetSaleDistrictResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeEnsNetSaleDistrictResponse) SetBody(v *DescribeEnsNetSaleDistrictResponseBody) *DescribeEnsNetSaleDistrictResponse {
	s.Body = v
	return s
}

func (s *DescribeEnsNetSaleDistrictResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
