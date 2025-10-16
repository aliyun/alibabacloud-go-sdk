// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTrFirewallsV2RouteListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeTrFirewallsV2RouteListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeTrFirewallsV2RouteListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeTrFirewallsV2RouteListResponseBody) *DescribeTrFirewallsV2RouteListResponse
	GetBody() *DescribeTrFirewallsV2RouteListResponseBody
}

type DescribeTrFirewallsV2RouteListResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeTrFirewallsV2RouteListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeTrFirewallsV2RouteListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeTrFirewallsV2RouteListResponse) GoString() string {
	return s.String()
}

func (s *DescribeTrFirewallsV2RouteListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeTrFirewallsV2RouteListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeTrFirewallsV2RouteListResponse) GetBody() *DescribeTrFirewallsV2RouteListResponseBody {
	return s.Body
}

func (s *DescribeTrFirewallsV2RouteListResponse) SetHeaders(v map[string]*string) *DescribeTrFirewallsV2RouteListResponse {
	s.Headers = v
	return s
}

func (s *DescribeTrFirewallsV2RouteListResponse) SetStatusCode(v int32) *DescribeTrFirewallsV2RouteListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeTrFirewallsV2RouteListResponse) SetBody(v *DescribeTrFirewallsV2RouteListResponseBody) *DescribeTrFirewallsV2RouteListResponse {
	s.Body = v
	return s
}

func (s *DescribeTrFirewallsV2RouteListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
