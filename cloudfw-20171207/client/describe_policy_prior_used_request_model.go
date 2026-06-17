// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePolicyPriorUsedRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDirection(v string) *DescribePolicyPriorUsedRequest
	GetDirection() *string
	SetIpVersion(v string) *DescribePolicyPriorUsedRequest
	GetIpVersion() *string
	SetLang(v string) *DescribePolicyPriorUsedRequest
	GetLang() *string
	SetSourceIp(v string) *DescribePolicyPriorUsedRequest
	GetSourceIp() *string
}

type DescribePolicyPriorUsedRequest struct {
	// The traffic direction of the access control policy.
	//
	// Valid values:
	//
	// - **in**: Inbound traffic
	//
	// - **out**: Outbound traffic
	//
	// This parameter is required.
	//
	// example:
	//
	// in
	Direction *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	// The IP version for assets protected by Cloud Firewall.
	//
	// Valid values:
	//
	// - **4*	- (Default): IPv4
	//
	// - **6**: IPv6
	//
	// example:
	//
	// 6
	IpVersion *string `json:"IpVersion,omitempty" xml:"IpVersion,omitempty"`
	// The language of the request and response.
	//
	// Valid values:
	//
	// - **zh*	- (Default): Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Deprecated
	//
	// The source IP address of the request.
	//
	// example:
	//
	// 192.0.XX.XX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DescribePolicyPriorUsedRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePolicyPriorUsedRequest) GoString() string {
	return s.String()
}

func (s *DescribePolicyPriorUsedRequest) GetDirection() *string {
	return s.Direction
}

func (s *DescribePolicyPriorUsedRequest) GetIpVersion() *string {
	return s.IpVersion
}

func (s *DescribePolicyPriorUsedRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribePolicyPriorUsedRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribePolicyPriorUsedRequest) SetDirection(v string) *DescribePolicyPriorUsedRequest {
	s.Direction = &v
	return s
}

func (s *DescribePolicyPriorUsedRequest) SetIpVersion(v string) *DescribePolicyPriorUsedRequest {
	s.IpVersion = &v
	return s
}

func (s *DescribePolicyPriorUsedRequest) SetLang(v string) *DescribePolicyPriorUsedRequest {
	s.Lang = &v
	return s
}

func (s *DescribePolicyPriorUsedRequest) SetSourceIp(v string) *DescribePolicyPriorUsedRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribePolicyPriorUsedRequest) Validate() error {
	return dara.Validate(s)
}
