// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSdnsVerifyRuleCategoryShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetArgsShrink(v string) *SdnsVerifyRuleCategoryShrinkRequest
	GetArgsShrink() *string
	SetCip(v string) *SdnsVerifyRuleCategoryShrinkRequest
	GetCip() *string
	SetClientIpRangesShrink(v string) *SdnsVerifyRuleCategoryShrinkRequest
	GetClientIpRangesShrink() *string
	SetMatchType(v string) *SdnsVerifyRuleCategoryShrinkRequest
	GetMatchType() *string
	SetRulesShrink(v string) *SdnsVerifyRuleCategoryShrinkRequest
	GetRulesShrink() *string
}

type SdnsVerifyRuleCategoryShrinkRequest struct {
	// The parameter object used for the current validation. If no custom parameters are used, pass an empty JSON object {}. Do not pass an empty string or null.
	ArgsShrink *string `json:"Args,omitempty" xml:"Args,omitempty"`
	// The simulated client IP address. This parameter is required in IP_RANGE mode. Supports a single IPv4 or IPv6 address. CIDR blocks and start-to-end ranges are not supported. This parameter is used for CIDR block matching in the current request and does not change the source of the actual DNS resolution request.
	//
	// example:
	//
	// 198.51.100.10
	Cip *string `json:"Cip,omitempty" xml:"Cip,omitempty"`
	// The array of client CIDR blocks that must be provided in the IP_RANGE pattern. Send this parameter as a JSON form field. Do not pass this parameter in the LINE pattern. Supports single IPv4 and IPv6 addresses, CIDR blocks, and start-to-end ranges. Multiple entries have an OR relationship. For CIDR blocks, the IPv4 prefix length ranges from 0 to 32, and the IPv6 prefix length ranges from 0 to 128. Start-to-end ranges are inclusive and must belong to the same address family, with the start address not greater than the end address. Overlapping, containing, and duplicate ranges are allowed within the same policy. When the policy is saved and read back, the array sequence, duplicate items, and the original text of each string are preserved. The array can contain up to 2,000 array elements. The total number of parsed non-empty address expressions ranges from 1 to 2,000, including duplicates. Specify one expression per array element. It is compatible with line breaks within an array element. Each expression can contain up to 128 characters, including leading and trailing whitespace. The serialized UTF-8 JSON content cannot exceed 65,535 bytes. Both the quantity and byte limits apply simultaneously. To authenticate, the system only checks the ranges that you commit and does not check for conflicts with saved policies.
	ClientIpRangesShrink *string `json:"ClientIpRanges,omitempty" xml:"ClientIpRanges,omitempty"`
	// The match type. Valid values:
	//
	// - LINE: line mode
	//
	// - IP_RANGE: client CIDR block mode
	//
	// If this parameter is omitted, LINE is used. In IP_RANGE mode, you must provide both ClientIpRanges and Cip.
	//
	// example:
	//
	// IP_RANGE
	MatchType *string `json:"MatchType,omitempty" xml:"MatchType,omitempty"`
	// The complete list of rules. This parameter is required and must contain at least one rule. If you only need to match CIDR blocks, configure a default rule without ArgKeys and Args, and configure the A/AAAA destination addresses in the Records field of this rule. This operation is used for match validation and does not save the rules.
	RulesShrink *string `json:"Rules,omitempty" xml:"Rules,omitempty"`
}

func (s SdnsVerifyRuleCategoryShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s SdnsVerifyRuleCategoryShrinkRequest) GoString() string {
	return s.String()
}

func (s *SdnsVerifyRuleCategoryShrinkRequest) GetArgsShrink() *string {
	return s.ArgsShrink
}

func (s *SdnsVerifyRuleCategoryShrinkRequest) GetCip() *string {
	return s.Cip
}

func (s *SdnsVerifyRuleCategoryShrinkRequest) GetClientIpRangesShrink() *string {
	return s.ClientIpRangesShrink
}

func (s *SdnsVerifyRuleCategoryShrinkRequest) GetMatchType() *string {
	return s.MatchType
}

func (s *SdnsVerifyRuleCategoryShrinkRequest) GetRulesShrink() *string {
	return s.RulesShrink
}

func (s *SdnsVerifyRuleCategoryShrinkRequest) SetArgsShrink(v string) *SdnsVerifyRuleCategoryShrinkRequest {
	s.ArgsShrink = &v
	return s
}

func (s *SdnsVerifyRuleCategoryShrinkRequest) SetCip(v string) *SdnsVerifyRuleCategoryShrinkRequest {
	s.Cip = &v
	return s
}

func (s *SdnsVerifyRuleCategoryShrinkRequest) SetClientIpRangesShrink(v string) *SdnsVerifyRuleCategoryShrinkRequest {
	s.ClientIpRangesShrink = &v
	return s
}

func (s *SdnsVerifyRuleCategoryShrinkRequest) SetMatchType(v string) *SdnsVerifyRuleCategoryShrinkRequest {
	s.MatchType = &v
	return s
}

func (s *SdnsVerifyRuleCategoryShrinkRequest) SetRulesShrink(v string) *SdnsVerifyRuleCategoryShrinkRequest {
	s.RulesShrink = &v
	return s
}

func (s *SdnsVerifyRuleCategoryShrinkRequest) Validate() error {
	return dara.Validate(s)
}
