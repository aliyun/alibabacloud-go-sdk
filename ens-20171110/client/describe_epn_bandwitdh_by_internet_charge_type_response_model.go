// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEpnBandwitdhByInternetChargeTypeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeEpnBandwitdhByInternetChargeTypeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeEpnBandwitdhByInternetChargeTypeResponse
	GetStatusCode() *int32
	SetBody(v *DescribeEpnBandwitdhByInternetChargeTypeResponseBody) *DescribeEpnBandwitdhByInternetChargeTypeResponse
	GetBody() *DescribeEpnBandwitdhByInternetChargeTypeResponseBody
}

type DescribeEpnBandwitdhByInternetChargeTypeResponse struct {
	Headers    map[string]*string                                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeEpnBandwitdhByInternetChargeTypeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeEpnBandwitdhByInternetChargeTypeResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeEpnBandwitdhByInternetChargeTypeResponse) GoString() string {
	return s.String()
}

func (s *DescribeEpnBandwitdhByInternetChargeTypeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeEpnBandwitdhByInternetChargeTypeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeEpnBandwitdhByInternetChargeTypeResponse) GetBody() *DescribeEpnBandwitdhByInternetChargeTypeResponseBody {
	return s.Body
}

func (s *DescribeEpnBandwitdhByInternetChargeTypeResponse) SetHeaders(v map[string]*string) *DescribeEpnBandwitdhByInternetChargeTypeResponse {
	s.Headers = v
	return s
}

func (s *DescribeEpnBandwitdhByInternetChargeTypeResponse) SetStatusCode(v int32) *DescribeEpnBandwitdhByInternetChargeTypeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeEpnBandwitdhByInternetChargeTypeResponse) SetBody(v *DescribeEpnBandwitdhByInternetChargeTypeResponseBody) *DescribeEpnBandwitdhByInternetChargeTypeResponse {
	s.Body = v
	return s
}

func (s *DescribeEpnBandwitdhByInternetChargeTypeResponse) Validate() error {
	return dara.Validate(s)
}
