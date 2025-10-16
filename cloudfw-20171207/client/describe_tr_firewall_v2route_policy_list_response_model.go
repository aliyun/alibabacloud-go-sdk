// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTrFirewallV2RoutePolicyListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeTrFirewallV2RoutePolicyListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeTrFirewallV2RoutePolicyListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeTrFirewallV2RoutePolicyListResponseBody) *DescribeTrFirewallV2RoutePolicyListResponse
	GetBody() *DescribeTrFirewallV2RoutePolicyListResponseBody
}

type DescribeTrFirewallV2RoutePolicyListResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeTrFirewallV2RoutePolicyListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeTrFirewallV2RoutePolicyListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeTrFirewallV2RoutePolicyListResponse) GoString() string {
	return s.String()
}

func (s *DescribeTrFirewallV2RoutePolicyListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeTrFirewallV2RoutePolicyListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeTrFirewallV2RoutePolicyListResponse) GetBody() *DescribeTrFirewallV2RoutePolicyListResponseBody {
	return s.Body
}

func (s *DescribeTrFirewallV2RoutePolicyListResponse) SetHeaders(v map[string]*string) *DescribeTrFirewallV2RoutePolicyListResponse {
	s.Headers = v
	return s
}

func (s *DescribeTrFirewallV2RoutePolicyListResponse) SetStatusCode(v int32) *DescribeTrFirewallV2RoutePolicyListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeTrFirewallV2RoutePolicyListResponse) SetBody(v *DescribeTrFirewallV2RoutePolicyListResponseBody) *DescribeTrFirewallV2RoutePolicyListResponse {
	s.Body = v
	return s
}

func (s *DescribeTrFirewallV2RoutePolicyListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
