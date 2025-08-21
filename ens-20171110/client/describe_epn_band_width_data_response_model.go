// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEpnBandWidthDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeEpnBandWidthDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeEpnBandWidthDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeEpnBandWidthDataResponseBody) *DescribeEpnBandWidthDataResponse
	GetBody() *DescribeEpnBandWidthDataResponseBody
}

type DescribeEpnBandWidthDataResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeEpnBandWidthDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeEpnBandWidthDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeEpnBandWidthDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeEpnBandWidthDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeEpnBandWidthDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeEpnBandWidthDataResponse) GetBody() *DescribeEpnBandWidthDataResponseBody {
	return s.Body
}

func (s *DescribeEpnBandWidthDataResponse) SetHeaders(v map[string]*string) *DescribeEpnBandWidthDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeEpnBandWidthDataResponse) SetStatusCode(v int32) *DescribeEpnBandWidthDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeEpnBandWidthDataResponse) SetBody(v *DescribeEpnBandWidthDataResponseBody) *DescribeEpnBandWidthDataResponse {
	s.Body = v
	return s
}

func (s *DescribeEpnBandWidthDataResponse) Validate() error {
	return dara.Validate(s)
}
