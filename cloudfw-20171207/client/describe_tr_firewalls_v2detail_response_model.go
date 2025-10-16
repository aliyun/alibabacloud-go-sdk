// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTrFirewallsV2DetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeTrFirewallsV2DetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeTrFirewallsV2DetailResponse
	GetStatusCode() *int32
	SetBody(v *DescribeTrFirewallsV2DetailResponseBody) *DescribeTrFirewallsV2DetailResponse
	GetBody() *DescribeTrFirewallsV2DetailResponseBody
}

type DescribeTrFirewallsV2DetailResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeTrFirewallsV2DetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeTrFirewallsV2DetailResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeTrFirewallsV2DetailResponse) GoString() string {
	return s.String()
}

func (s *DescribeTrFirewallsV2DetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeTrFirewallsV2DetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeTrFirewallsV2DetailResponse) GetBody() *DescribeTrFirewallsV2DetailResponseBody {
	return s.Body
}

func (s *DescribeTrFirewallsV2DetailResponse) SetHeaders(v map[string]*string) *DescribeTrFirewallsV2DetailResponse {
	s.Headers = v
	return s
}

func (s *DescribeTrFirewallsV2DetailResponse) SetStatusCode(v int32) *DescribeTrFirewallsV2DetailResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeTrFirewallsV2DetailResponse) SetBody(v *DescribeTrFirewallsV2DetailResponseBody) *DescribeTrFirewallsV2DetailResponse {
	s.Body = v
	return s
}

func (s *DescribeTrFirewallsV2DetailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
