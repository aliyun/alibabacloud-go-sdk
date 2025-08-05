// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTrFirewallsV2ListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeTrFirewallsV2ListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeTrFirewallsV2ListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeTrFirewallsV2ListResponseBody) *DescribeTrFirewallsV2ListResponse
	GetBody() *DescribeTrFirewallsV2ListResponseBody
}

type DescribeTrFirewallsV2ListResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeTrFirewallsV2ListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeTrFirewallsV2ListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeTrFirewallsV2ListResponse) GoString() string {
	return s.String()
}

func (s *DescribeTrFirewallsV2ListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeTrFirewallsV2ListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeTrFirewallsV2ListResponse) GetBody() *DescribeTrFirewallsV2ListResponseBody {
	return s.Body
}

func (s *DescribeTrFirewallsV2ListResponse) SetHeaders(v map[string]*string) *DescribeTrFirewallsV2ListResponse {
	s.Headers = v
	return s
}

func (s *DescribeTrFirewallsV2ListResponse) SetStatusCode(v int32) *DescribeTrFirewallsV2ListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeTrFirewallsV2ListResponse) SetBody(v *DescribeTrFirewallsV2ListResponseBody) *DescribeTrFirewallsV2ListResponse {
	s.Body = v
	return s
}

func (s *DescribeTrFirewallsV2ListResponse) Validate() error {
	return dara.Validate(s)
}
