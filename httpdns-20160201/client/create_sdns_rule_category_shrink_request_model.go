// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSdnsRuleCategoryShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientIpRangesShrink(v string) *CreateSdnsRuleCategoryShrinkRequest
	GetClientIpRangesShrink() *string
	SetDomainName(v string) *CreateSdnsRuleCategoryShrinkRequest
	GetDomainName() *string
	SetEnable(v bool) *CreateSdnsRuleCategoryShrinkRequest
	GetEnable() *bool
	SetLineId(v string) *CreateSdnsRuleCategoryShrinkRequest
	GetLineId() *string
	SetMatchType(v string) *CreateSdnsRuleCategoryShrinkRequest
	GetMatchType() *string
	SetPathsShrink(v string) *CreateSdnsRuleCategoryShrinkRequest
	GetPathsShrink() *string
	SetRulesShrink(v string) *CreateSdnsRuleCategoryShrinkRequest
	GetRulesShrink() *string
}

type CreateSdnsRuleCategoryShrinkRequest struct {
	// The array of client IP address ranges required for the IP_RANGE pattern. Send this parameter as a JSON form field. Do not specify this parameter for the LINE pattern. This parameter supports single IPv4 and IPv6 addresses, CIDR blocks, and start-end ranges. Multiple entries have an OR relationship. For CIDR blocks, the IPv4 prefix length ranges from 0 to 32, and the IPv6 prefix length ranges from 0 to 128. Start-end ranges are inclusive and must belong to the same address family, and the start address cannot be greater than the end address. Duplicates, inclusions, and overlaps are allowed within the same policy. When the policy is saved and read back, the array sequence, duplicates, and original text of each string are preserved. The array can contain up to 2,000 array elements. The total number of parsed non-empty address expressions ranges from 1 to 2,000, and duplicates are counted. We recommend that you specify one expression per array element. It is compatible with line breaks within an array element. Each expression can contain up to 128 characters, including leading and trailing spaces. The serialized UTF-8 JSON content cannot exceed 65,535 bytes. The quantity and byte limits take effect simultaneously. Overlaps are not allowed between different IP_RANGE policies for the same account, domain name, and address family. Disabled policies are also checked. Disabling a policy does not release the ranges. This limit does not compare independent configurations of LINE policies and the legacy SetIPRegionResult operation.
	ClientIpRangesShrink *string `json:"ClientIpRanges,omitempty" xml:"ClientIpRanges,omitempty"`
	// The domain name connected to HTTPDNS for the current account. This parameter is required.
	//
	// example:
	//
	// api.example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// Specifies whether to enable the policy upon creation. If this parameter is omitted or set to false, the policy is disabled (Status=0). If this parameter is set to true, the policy is enabled (Status=1), and you do not need to call the enable operation. This parameter applies to the LINE and IP_RANGE modes. Actual resolution still requires waiting for configuration propagation.
	Enable *bool `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// The line identifier required for the LINE mode. Valid values:
	//
	// - internal: the Chinese mainland.
	//
	// - overseas: outside the Chinese mainland.
	//
	// Do not specify this parameter for the IP_RANGE mode.
	//
	// example:
	//
	// internal
	LineId *string `json:"LineId,omitempty" xml:"LineId,omitempty"`
	// The match type. Valid values:
	//
	// - LINE: the line mode.
	//
	// - IP_RANGE: the client IP address range mode.
	//
	// If this parameter is omitted, it is processed as LINE. To create an IP address range policy, you must explicitly set this parameter to IP_RANGE and provide ClientIpRanges. You cannot provide LineId or Paths at the same time. The match type cannot be changed after creation.
	//
	// example:
	//
	// IP_RANGE
	MatchType *string `json:"MatchType,omitempty" xml:"MatchType,omitempty"`
	// The line selection path that describes the line conditions at each level. This parameter applies only to the LINE mode. The IP_RANGE mode does not use line paths. When you create a line policy, provide valid options at each level. Use x for the default option. For the Chinese mainland, the levels include ISP, region, and province. For regions outside the Chinese mainland, the levels include continent and region.
	PathsShrink *string `json:"Paths,omitempty" xml:"Paths,omitempty"`
	// The complete list of rules. This parameter is required and must contain at least one rule. When matching only by IP address range, configure a default rule without ArgKeys and Args, and configure the A or AAAA destination addresses in the Records of this rule. When saved, the serialized UTF-8 JSON content must be less than 65,535 bytes.
	RulesShrink *string `json:"Rules,omitempty" xml:"Rules,omitempty"`
}

func (s CreateSdnsRuleCategoryShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSdnsRuleCategoryShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateSdnsRuleCategoryShrinkRequest) GetClientIpRangesShrink() *string {
	return s.ClientIpRangesShrink
}

func (s *CreateSdnsRuleCategoryShrinkRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *CreateSdnsRuleCategoryShrinkRequest) GetEnable() *bool {
	return s.Enable
}

func (s *CreateSdnsRuleCategoryShrinkRequest) GetLineId() *string {
	return s.LineId
}

func (s *CreateSdnsRuleCategoryShrinkRequest) GetMatchType() *string {
	return s.MatchType
}

func (s *CreateSdnsRuleCategoryShrinkRequest) GetPathsShrink() *string {
	return s.PathsShrink
}

func (s *CreateSdnsRuleCategoryShrinkRequest) GetRulesShrink() *string {
	return s.RulesShrink
}

func (s *CreateSdnsRuleCategoryShrinkRequest) SetClientIpRangesShrink(v string) *CreateSdnsRuleCategoryShrinkRequest {
	s.ClientIpRangesShrink = &v
	return s
}

func (s *CreateSdnsRuleCategoryShrinkRequest) SetDomainName(v string) *CreateSdnsRuleCategoryShrinkRequest {
	s.DomainName = &v
	return s
}

func (s *CreateSdnsRuleCategoryShrinkRequest) SetEnable(v bool) *CreateSdnsRuleCategoryShrinkRequest {
	s.Enable = &v
	return s
}

func (s *CreateSdnsRuleCategoryShrinkRequest) SetLineId(v string) *CreateSdnsRuleCategoryShrinkRequest {
	s.LineId = &v
	return s
}

func (s *CreateSdnsRuleCategoryShrinkRequest) SetMatchType(v string) *CreateSdnsRuleCategoryShrinkRequest {
	s.MatchType = &v
	return s
}

func (s *CreateSdnsRuleCategoryShrinkRequest) SetPathsShrink(v string) *CreateSdnsRuleCategoryShrinkRequest {
	s.PathsShrink = &v
	return s
}

func (s *CreateSdnsRuleCategoryShrinkRequest) SetRulesShrink(v string) *CreateSdnsRuleCategoryShrinkRequest {
	s.RulesShrink = &v
	return s
}

func (s *CreateSdnsRuleCategoryShrinkRequest) Validate() error {
	return dara.Validate(s)
}
