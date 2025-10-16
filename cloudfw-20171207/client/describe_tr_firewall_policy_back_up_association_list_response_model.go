// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTrFirewallPolicyBackUpAssociationListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeTrFirewallPolicyBackUpAssociationListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeTrFirewallPolicyBackUpAssociationListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeTrFirewallPolicyBackUpAssociationListResponseBody) *DescribeTrFirewallPolicyBackUpAssociationListResponse
	GetBody() *DescribeTrFirewallPolicyBackUpAssociationListResponseBody
}

type DescribeTrFirewallPolicyBackUpAssociationListResponse struct {
	Headers    map[string]*string                                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeTrFirewallPolicyBackUpAssociationListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeTrFirewallPolicyBackUpAssociationListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeTrFirewallPolicyBackUpAssociationListResponse) GoString() string {
	return s.String()
}

func (s *DescribeTrFirewallPolicyBackUpAssociationListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeTrFirewallPolicyBackUpAssociationListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeTrFirewallPolicyBackUpAssociationListResponse) GetBody() *DescribeTrFirewallPolicyBackUpAssociationListResponseBody {
	return s.Body
}

func (s *DescribeTrFirewallPolicyBackUpAssociationListResponse) SetHeaders(v map[string]*string) *DescribeTrFirewallPolicyBackUpAssociationListResponse {
	s.Headers = v
	return s
}

func (s *DescribeTrFirewallPolicyBackUpAssociationListResponse) SetStatusCode(v int32) *DescribeTrFirewallPolicyBackUpAssociationListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeTrFirewallPolicyBackUpAssociationListResponse) SetBody(v *DescribeTrFirewallPolicyBackUpAssociationListResponseBody) *DescribeTrFirewallPolicyBackUpAssociationListResponse {
	s.Body = v
	return s
}

func (s *DescribeTrFirewallPolicyBackUpAssociationListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
