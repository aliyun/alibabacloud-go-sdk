// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSdnsVerifyRuleCategoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetArgs(v map[string]*string) *SdnsVerifyRuleCategoryRequest
	GetArgs() map[string]*string
	SetCip(v string) *SdnsVerifyRuleCategoryRequest
	GetCip() *string
	SetClientIpRanges(v []*string) *SdnsVerifyRuleCategoryRequest
	GetClientIpRanges() []*string
	SetMatchType(v string) *SdnsVerifyRuleCategoryRequest
	GetMatchType() *string
	SetRules(v []*SdnsVerifyRuleCategoryRequestRules) *SdnsVerifyRuleCategoryRequest
	GetRules() []*SdnsVerifyRuleCategoryRequestRules
}

type SdnsVerifyRuleCategoryRequest struct {
	// The parameter object used for the current validation. If no custom parameters are used, pass an empty JSON object {}. Do not pass an empty string or null.
	Args map[string]*string `json:"Args,omitempty" xml:"Args,omitempty"`
	// The simulated client IP address. This parameter is required in IP_RANGE mode. Supports a single IPv4 or IPv6 address. CIDR blocks and start-to-end ranges are not supported. This parameter is used for CIDR block matching in the current request and does not change the source of the actual DNS resolution request.
	//
	// example:
	//
	// 198.51.100.10
	Cip *string `json:"Cip,omitempty" xml:"Cip,omitempty"`
	// The array of client CIDR blocks that must be provided in the IP_RANGE pattern. Send this parameter as a JSON form field. Do not pass this parameter in the LINE pattern. Supports single IPv4 and IPv6 addresses, CIDR blocks, and start-to-end ranges. Multiple entries have an OR relationship. For CIDR blocks, the IPv4 prefix length ranges from 0 to 32, and the IPv6 prefix length ranges from 0 to 128. Start-to-end ranges are inclusive and must belong to the same address family, with the start address not greater than the end address. Overlapping, containing, and duplicate ranges are allowed within the same policy. When the policy is saved and read back, the array sequence, duplicate items, and the original text of each string are preserved. The array can contain up to 2,000 array elements. The total number of parsed non-empty address expressions ranges from 1 to 2,000, including duplicates. Specify one expression per array element. It is compatible with line breaks within an array element. Each expression can contain up to 128 characters, including leading and trailing whitespace. The serialized UTF-8 JSON content cannot exceed 65,535 bytes. Both the quantity and byte limits apply simultaneously. To authenticate, the system only checks the ranges that you commit and does not check for conflicts with saved policies.
	ClientIpRanges []*string `json:"ClientIpRanges,omitempty" xml:"ClientIpRanges,omitempty" type:"Repeated"`
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
	Rules []*SdnsVerifyRuleCategoryRequestRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
}

func (s SdnsVerifyRuleCategoryRequest) String() string {
	return dara.Prettify(s)
}

func (s SdnsVerifyRuleCategoryRequest) GoString() string {
	return s.String()
}

func (s *SdnsVerifyRuleCategoryRequest) GetArgs() map[string]*string {
	return s.Args
}

func (s *SdnsVerifyRuleCategoryRequest) GetCip() *string {
	return s.Cip
}

func (s *SdnsVerifyRuleCategoryRequest) GetClientIpRanges() []*string {
	return s.ClientIpRanges
}

func (s *SdnsVerifyRuleCategoryRequest) GetMatchType() *string {
	return s.MatchType
}

func (s *SdnsVerifyRuleCategoryRequest) GetRules() []*SdnsVerifyRuleCategoryRequestRules {
	return s.Rules
}

func (s *SdnsVerifyRuleCategoryRequest) SetArgs(v map[string]*string) *SdnsVerifyRuleCategoryRequest {
	s.Args = v
	return s
}

func (s *SdnsVerifyRuleCategoryRequest) SetCip(v string) *SdnsVerifyRuleCategoryRequest {
	s.Cip = &v
	return s
}

func (s *SdnsVerifyRuleCategoryRequest) SetClientIpRanges(v []*string) *SdnsVerifyRuleCategoryRequest {
	s.ClientIpRanges = v
	return s
}

func (s *SdnsVerifyRuleCategoryRequest) SetMatchType(v string) *SdnsVerifyRuleCategoryRequest {
	s.MatchType = &v
	return s
}

func (s *SdnsVerifyRuleCategoryRequest) SetRules(v []*SdnsVerifyRuleCategoryRequestRules) *SdnsVerifyRuleCategoryRequest {
	s.Rules = v
	return s
}

func (s *SdnsVerifyRuleCategoryRequest) Validate() error {
	if s.Rules != nil {
		for _, item := range s.Rules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SdnsVerifyRuleCategoryRequestRules struct {
	// The optional list of custom parameter names, used to record the display order. Omit this parameter if you only match CIDR blocks. If you use Args, the names must correspond to the parameter names in Args.
	ArgKeys []*string `json:"ArgKeys,omitempty" xml:"ArgKeys,omitempty" type:"Repeated"`
	// The optional custom parameter conditions. The key is the parameter name and the value is the match value. Omit this parameter if you only match CIDR blocks. You do not need to add business parameters.
	Args map[string]*string `json:"Args,omitempty" xml:"Args,omitempty"`
	// The rule name. This parameter is required and cannot be null. Use a non-empty name with business meaning. This parameter cannot replace the policy ID.
	//
	// example:
	//
	// Default rule
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The list of configured DNS records, which can include A and AAAA records. Actual DNS resolution returns only records of the queried type. If the list is empty or does not contain records of the queried type, the system continues to fall back. This operation returns the records configured in the rule and does not filter them by query type.
	Records []*SdnsVerifyRuleCategoryRequestRulesRecords `json:"Records,omitempty" xml:"Records,omitempty" type:"Repeated"`
	// The record return mode. Valid values:
	//
	// - 0: Returns all records of the queried type.
	//
	// - 1: Returns one record of the queried type based on weight.
	//
	// Policy validation only returns the rule configuration and does not perform weight sampling.
	//
	// example:
	//
	// 0
	ResultMode *int32 `json:"ResultMode,omitempty" xml:"ResultMode,omitempty"`
	// The cache validity period of the record, in seconds. This parameter must be a positive integer.
	//
	// example:
	//
	// 60
	Ttl *int32 `json:"Ttl,omitempty" xml:"Ttl,omitempty"`
}

func (s SdnsVerifyRuleCategoryRequestRules) String() string {
	return dara.Prettify(s)
}

func (s SdnsVerifyRuleCategoryRequestRules) GoString() string {
	return s.String()
}

func (s *SdnsVerifyRuleCategoryRequestRules) GetArgKeys() []*string {
	return s.ArgKeys
}

func (s *SdnsVerifyRuleCategoryRequestRules) GetArgs() map[string]*string {
	return s.Args
}

func (s *SdnsVerifyRuleCategoryRequestRules) GetName() *string {
	return s.Name
}

func (s *SdnsVerifyRuleCategoryRequestRules) GetRecords() []*SdnsVerifyRuleCategoryRequestRulesRecords {
	return s.Records
}

func (s *SdnsVerifyRuleCategoryRequestRules) GetResultMode() *int32 {
	return s.ResultMode
}

func (s *SdnsVerifyRuleCategoryRequestRules) GetTtl() *int32 {
	return s.Ttl
}

func (s *SdnsVerifyRuleCategoryRequestRules) SetArgKeys(v []*string) *SdnsVerifyRuleCategoryRequestRules {
	s.ArgKeys = v
	return s
}

func (s *SdnsVerifyRuleCategoryRequestRules) SetArgs(v map[string]*string) *SdnsVerifyRuleCategoryRequestRules {
	s.Args = v
	return s
}

func (s *SdnsVerifyRuleCategoryRequestRules) SetName(v string) *SdnsVerifyRuleCategoryRequestRules {
	s.Name = &v
	return s
}

func (s *SdnsVerifyRuleCategoryRequestRules) SetRecords(v []*SdnsVerifyRuleCategoryRequestRulesRecords) *SdnsVerifyRuleCategoryRequestRules {
	s.Records = v
	return s
}

func (s *SdnsVerifyRuleCategoryRequestRules) SetResultMode(v int32) *SdnsVerifyRuleCategoryRequestRules {
	s.ResultMode = &v
	return s
}

func (s *SdnsVerifyRuleCategoryRequestRules) SetTtl(v int32) *SdnsVerifyRuleCategoryRequestRules {
	s.Ttl = &v
	return s
}

func (s *SdnsVerifyRuleCategoryRequestRules) Validate() error {
	if s.Records != nil {
		for _, item := range s.Records {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SdnsVerifyRuleCategoryRequestRulesRecords struct {
	// The single destination IP address for the record. If Type is set to 1, specify an IPv4 address, such as 192.0.2.200. If Type is set to 28, specify an IPv6 address, such as 2001:db8::200. Do not use CIDR blocks, start-to-end ranges, or domain names.
	//
	// example:
	//
	// 192.0.2.200
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The DNS record type. Valid values:
	//
	// - 1: A record
	//
	// - 28: AAAA record
	//
	// Specify an integer. Do not use the strings A or AAAA.
	//
	// example:
	//
	// 1
	Type *int32 `json:"Type,omitempty" xml:"Type,omitempty"`
	// The auxiliary display text for the record type. This parameter does not participate in matching or resolution decisions and can be omitted in the request.
	//
	// example:
	//
	// A
	TypeLabel *string `json:"TypeLabel,omitempty" xml:"TypeLabel,omitempty"`
	// The relative weight of the record. This parameter is used to select one record of the same query type only when ResultMode is set to 1. Set this parameter to a positive integer from 1 to 100. The sum of the weights of all records does not need to be 100. This parameter does not participate in selection when ResultMode is set to 0.
	//
	// example:
	//
	// 100
	Weight *int32 `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s SdnsVerifyRuleCategoryRequestRulesRecords) String() string {
	return dara.Prettify(s)
}

func (s SdnsVerifyRuleCategoryRequestRulesRecords) GoString() string {
	return s.String()
}

func (s *SdnsVerifyRuleCategoryRequestRulesRecords) GetIp() *string {
	return s.Ip
}

func (s *SdnsVerifyRuleCategoryRequestRulesRecords) GetType() *int32 {
	return s.Type
}

func (s *SdnsVerifyRuleCategoryRequestRulesRecords) GetTypeLabel() *string {
	return s.TypeLabel
}

func (s *SdnsVerifyRuleCategoryRequestRulesRecords) GetWeight() *int32 {
	return s.Weight
}

func (s *SdnsVerifyRuleCategoryRequestRulesRecords) SetIp(v string) *SdnsVerifyRuleCategoryRequestRulesRecords {
	s.Ip = &v
	return s
}

func (s *SdnsVerifyRuleCategoryRequestRulesRecords) SetType(v int32) *SdnsVerifyRuleCategoryRequestRulesRecords {
	s.Type = &v
	return s
}

func (s *SdnsVerifyRuleCategoryRequestRulesRecords) SetTypeLabel(v string) *SdnsVerifyRuleCategoryRequestRulesRecords {
	s.TypeLabel = &v
	return s
}

func (s *SdnsVerifyRuleCategoryRequestRulesRecords) SetWeight(v int32) *SdnsVerifyRuleCategoryRequestRulesRecords {
	s.Weight = &v
	return s
}

func (s *SdnsVerifyRuleCategoryRequestRulesRecords) Validate() error {
	return dara.Validate(s)
}
