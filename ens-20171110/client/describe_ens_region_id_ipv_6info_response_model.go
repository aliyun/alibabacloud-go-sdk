// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEnsRegionIdIpv6InfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeEnsRegionIdIpv6InfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeEnsRegionIdIpv6InfoResponse
	GetStatusCode() *int32
	SetBody(v *DescribeEnsRegionIdIpv6InfoResponseBody) *DescribeEnsRegionIdIpv6InfoResponse
	GetBody() *DescribeEnsRegionIdIpv6InfoResponseBody
}

type DescribeEnsRegionIdIpv6InfoResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeEnsRegionIdIpv6InfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeEnsRegionIdIpv6InfoResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeEnsRegionIdIpv6InfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeEnsRegionIdIpv6InfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeEnsRegionIdIpv6InfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeEnsRegionIdIpv6InfoResponse) GetBody() *DescribeEnsRegionIdIpv6InfoResponseBody {
	return s.Body
}

func (s *DescribeEnsRegionIdIpv6InfoResponse) SetHeaders(v map[string]*string) *DescribeEnsRegionIdIpv6InfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeEnsRegionIdIpv6InfoResponse) SetStatusCode(v int32) *DescribeEnsRegionIdIpv6InfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeEnsRegionIdIpv6InfoResponse) SetBody(v *DescribeEnsRegionIdIpv6InfoResponseBody) *DescribeEnsRegionIdIpv6InfoResponse {
	s.Body = v
	return s
}

func (s *DescribeEnsRegionIdIpv6InfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
