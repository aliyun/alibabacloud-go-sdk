// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSdnsRuleCategoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientIpRanges(v []*string) *CreateSdnsRuleCategoryRequest
	GetClientIpRanges() []*string
	SetDomainName(v string) *CreateSdnsRuleCategoryRequest
	GetDomainName() *string
	SetEnable(v bool) *CreateSdnsRuleCategoryRequest
	GetEnable() *bool
	SetLineId(v string) *CreateSdnsRuleCategoryRequest
	GetLineId() *string
	SetMatchType(v string) *CreateSdnsRuleCategoryRequest
	GetMatchType() *string
	SetPaths(v []*CreateSdnsRuleCategoryRequestPaths) *CreateSdnsRuleCategoryRequest
	GetPaths() []*CreateSdnsRuleCategoryRequestPaths
	SetRules(v []*CreateSdnsRuleCategoryRequestRules) *CreateSdnsRuleCategoryRequest
	GetRules() []*CreateSdnsRuleCategoryRequestRules
}

type CreateSdnsRuleCategoryRequest struct {
	// The array of client IP address ranges required for the IP_RANGE pattern. Send this parameter as a JSON form field. Do not specify this parameter for the LINE pattern. This parameter supports single IPv4 and IPv6 addresses, CIDR blocks, and start-end ranges. Multiple entries have an OR relationship. For CIDR blocks, the IPv4 prefix length ranges from 0 to 32, and the IPv6 prefix length ranges from 0 to 128. Start-end ranges are inclusive and must belong to the same address family, and the start address cannot be greater than the end address. Duplicates, inclusions, and overlaps are allowed within the same policy. When the policy is saved and read back, the array sequence, duplicates, and original text of each string are preserved. The array can contain up to 2,000 array elements. The total number of parsed non-empty address expressions ranges from 1 to 2,000, and duplicates are counted. We recommend that you specify one expression per array element. It is compatible with line breaks within an array element. Each expression can contain up to 128 characters, including leading and trailing spaces. The serialized UTF-8 JSON content cannot exceed 65,535 bytes. The quantity and byte limits take effect simultaneously. Overlaps are not allowed between different IP_RANGE policies for the same account, domain name, and address family. Disabled policies are also checked. Disabling a policy does not release the ranges. This limit does not compare independent configurations of LINE policies and the legacy SetIPRegionResult operation.
	ClientIpRanges []*string `json:"ClientIpRanges,omitempty" xml:"ClientIpRanges,omitempty" type:"Repeated"`
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
	Paths []*CreateSdnsRuleCategoryRequestPaths `json:"Paths,omitempty" xml:"Paths,omitempty" type:"Repeated"`
	// The complete list of rules. This parameter is required and must contain at least one rule. When matching only by IP address range, configure a default rule without ArgKeys and Args, and configure the A or AAAA destination addresses in the Records of this rule. When saved, the serialized UTF-8 JSON content must be less than 65,535 bytes.
	Rules []*CreateSdnsRuleCategoryRequestRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
}

func (s CreateSdnsRuleCategoryRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSdnsRuleCategoryRequest) GoString() string {
	return s.String()
}

func (s *CreateSdnsRuleCategoryRequest) GetClientIpRanges() []*string {
	return s.ClientIpRanges
}

func (s *CreateSdnsRuleCategoryRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *CreateSdnsRuleCategoryRequest) GetEnable() *bool {
	return s.Enable
}

func (s *CreateSdnsRuleCategoryRequest) GetLineId() *string {
	return s.LineId
}

func (s *CreateSdnsRuleCategoryRequest) GetMatchType() *string {
	return s.MatchType
}

func (s *CreateSdnsRuleCategoryRequest) GetPaths() []*CreateSdnsRuleCategoryRequestPaths {
	return s.Paths
}

func (s *CreateSdnsRuleCategoryRequest) GetRules() []*CreateSdnsRuleCategoryRequestRules {
	return s.Rules
}

func (s *CreateSdnsRuleCategoryRequest) SetClientIpRanges(v []*string) *CreateSdnsRuleCategoryRequest {
	s.ClientIpRanges = v
	return s
}

func (s *CreateSdnsRuleCategoryRequest) SetDomainName(v string) *CreateSdnsRuleCategoryRequest {
	s.DomainName = &v
	return s
}

func (s *CreateSdnsRuleCategoryRequest) SetEnable(v bool) *CreateSdnsRuleCategoryRequest {
	s.Enable = &v
	return s
}

func (s *CreateSdnsRuleCategoryRequest) SetLineId(v string) *CreateSdnsRuleCategoryRequest {
	s.LineId = &v
	return s
}

func (s *CreateSdnsRuleCategoryRequest) SetMatchType(v string) *CreateSdnsRuleCategoryRequest {
	s.MatchType = &v
	return s
}

func (s *CreateSdnsRuleCategoryRequest) SetPaths(v []*CreateSdnsRuleCategoryRequestPaths) *CreateSdnsRuleCategoryRequest {
	s.Paths = v
	return s
}

func (s *CreateSdnsRuleCategoryRequest) SetRules(v []*CreateSdnsRuleCategoryRequestRules) *CreateSdnsRuleCategoryRequest {
	s.Rules = v
	return s
}

func (s *CreateSdnsRuleCategoryRequest) Validate() error {
	if s.Paths != nil {
		for _, item := range s.Paths {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
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

type CreateSdnsRuleCategoryRequestPaths struct {
	// The code of the option at this level. For example, chinanet at the isps level indicates China Telecom. Use valid options for the associated line and maintain consistent hierarchical subordination.
	//
	// example:
	//
	// chinanet
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The level identifier. For the Chinese mainland, use isps, subRegions, and provinces. For regions outside the Chinese mainland, use continents and geoRegions. This parameter corresponds to Level.
	//
	// example:
	//
	// isps
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The sequence number of the line level. For the Chinese mainland, the values are 1 (ISP), 2 (region), and 3 (province). For regions outside the Chinese mainland, the values are 1 (continent) and 2 (region).
	//
	// example:
	//
	// 1
	Level *int32 `json:"Level,omitempty" xml:"Level,omitempty"`
}

func (s CreateSdnsRuleCategoryRequestPaths) String() string {
	return dara.Prettify(s)
}

func (s CreateSdnsRuleCategoryRequestPaths) GoString() string {
	return s.String()
}

func (s *CreateSdnsRuleCategoryRequestPaths) GetId() *string {
	return s.Id
}

func (s *CreateSdnsRuleCategoryRequestPaths) GetKey() *string {
	return s.Key
}

func (s *CreateSdnsRuleCategoryRequestPaths) GetLevel() *int32 {
	return s.Level
}

func (s *CreateSdnsRuleCategoryRequestPaths) SetId(v string) *CreateSdnsRuleCategoryRequestPaths {
	s.Id = &v
	return s
}

func (s *CreateSdnsRuleCategoryRequestPaths) SetKey(v string) *CreateSdnsRuleCategoryRequestPaths {
	s.Key = &v
	return s
}

func (s *CreateSdnsRuleCategoryRequestPaths) SetLevel(v int32) *CreateSdnsRuleCategoryRequestPaths {
	s.Level = &v
	return s
}

func (s *CreateSdnsRuleCategoryRequestPaths) Validate() error {
	return dara.Validate(s)
}

type CreateSdnsRuleCategoryRequestRules struct {
	// The optional list of custom parameter names used to record the display order. Omit this parameter when matching only by IP address range. If Args is used, the names must correspond to the parameter names in Args.
	ArgKeys []*string `json:"ArgKeys,omitempty" xml:"ArgKeys,omitempty" type:"Repeated"`
	// The optional custom parameter conditions. The key is the parameter name, and the value is the match value. Omit this parameter when matching only by IP address range. You do not need to add business parameters.
	Args map[string]*string `json:"Args,omitempty" xml:"Args,omitempty"`
	// The rule name. This parameter is required and cannot be null. We recommend that you use a non-empty name with business meaning. This parameter cannot replace the policy Id.
	//
	// example:
	//
	// Default rule
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The list of configured resolution records, which can include A and AAAA records. Actual resolution returns only records of the queried type. If the list is empty or does not contain records of the queried type, the system continues to fall back.
	Records []*CreateSdnsRuleCategoryRequestRulesRecords `json:"Records,omitempty" xml:"Records,omitempty" type:"Repeated"`
	// The record return mode. Valid values:
	//
	// - 0: returns all records of the queried type.
	//
	// - 1: returns one record of the queried type based on the weight.
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

func (s CreateSdnsRuleCategoryRequestRules) String() string {
	return dara.Prettify(s)
}

func (s CreateSdnsRuleCategoryRequestRules) GoString() string {
	return s.String()
}

func (s *CreateSdnsRuleCategoryRequestRules) GetArgKeys() []*string {
	return s.ArgKeys
}

func (s *CreateSdnsRuleCategoryRequestRules) GetArgs() map[string]*string {
	return s.Args
}

func (s *CreateSdnsRuleCategoryRequestRules) GetName() *string {
	return s.Name
}

func (s *CreateSdnsRuleCategoryRequestRules) GetRecords() []*CreateSdnsRuleCategoryRequestRulesRecords {
	return s.Records
}

func (s *CreateSdnsRuleCategoryRequestRules) GetResultMode() *int32 {
	return s.ResultMode
}

func (s *CreateSdnsRuleCategoryRequestRules) GetTtl() *int32 {
	return s.Ttl
}

func (s *CreateSdnsRuleCategoryRequestRules) SetArgKeys(v []*string) *CreateSdnsRuleCategoryRequestRules {
	s.ArgKeys = v
	return s
}

func (s *CreateSdnsRuleCategoryRequestRules) SetArgs(v map[string]*string) *CreateSdnsRuleCategoryRequestRules {
	s.Args = v
	return s
}

func (s *CreateSdnsRuleCategoryRequestRules) SetName(v string) *CreateSdnsRuleCategoryRequestRules {
	s.Name = &v
	return s
}

func (s *CreateSdnsRuleCategoryRequestRules) SetRecords(v []*CreateSdnsRuleCategoryRequestRulesRecords) *CreateSdnsRuleCategoryRequestRules {
	s.Records = v
	return s
}

func (s *CreateSdnsRuleCategoryRequestRules) SetResultMode(v int32) *CreateSdnsRuleCategoryRequestRules {
	s.ResultMode = &v
	return s
}

func (s *CreateSdnsRuleCategoryRequestRules) SetTtl(v int32) *CreateSdnsRuleCategoryRequestRules {
	s.Ttl = &v
	return s
}

func (s *CreateSdnsRuleCategoryRequestRules) Validate() error {
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

type CreateSdnsRuleCategoryRequestRulesRecords struct {
	// The single destination IP address corresponding to the record. If Type is set to 1, the value is an IPv4 address, such as 192.0.2.200. If Type is set to 28, the value is an IPv6 address, such as 2001:db8::200. Do not use CIDR blocks, start-end ranges, or domain names.
	//
	// example:
	//
	// 192.0.2.200
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The DNS record type. Valid values:
	//
	// - 1: A record.
	//
	// - 28: AAAA record.
	//
	// Specify an integer. Do not use the string A or AAAA.
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
	// The relative weight of the record. This parameter is used to select one record of the same queried type only when ResultMode is set to 1. We recommend that you set this parameter to a positive integer from 1 to 100. The sum of the weights of all records does not need to be 100. This parameter does not participate in selection when ResultMode is set to 0.
	//
	// example:
	//
	// 100
	Weight *int32 `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s CreateSdnsRuleCategoryRequestRulesRecords) String() string {
	return dara.Prettify(s)
}

func (s CreateSdnsRuleCategoryRequestRulesRecords) GoString() string {
	return s.String()
}

func (s *CreateSdnsRuleCategoryRequestRulesRecords) GetIp() *string {
	return s.Ip
}

func (s *CreateSdnsRuleCategoryRequestRulesRecords) GetType() *int32 {
	return s.Type
}

func (s *CreateSdnsRuleCategoryRequestRulesRecords) GetTypeLabel() *string {
	return s.TypeLabel
}

func (s *CreateSdnsRuleCategoryRequestRulesRecords) GetWeight() *int32 {
	return s.Weight
}

func (s *CreateSdnsRuleCategoryRequestRulesRecords) SetIp(v string) *CreateSdnsRuleCategoryRequestRulesRecords {
	s.Ip = &v
	return s
}

func (s *CreateSdnsRuleCategoryRequestRulesRecords) SetType(v int32) *CreateSdnsRuleCategoryRequestRulesRecords {
	s.Type = &v
	return s
}

func (s *CreateSdnsRuleCategoryRequestRulesRecords) SetTypeLabel(v string) *CreateSdnsRuleCategoryRequestRulesRecords {
	s.TypeLabel = &v
	return s
}

func (s *CreateSdnsRuleCategoryRequestRulesRecords) SetWeight(v int32) *CreateSdnsRuleCategoryRequestRulesRecords {
	s.Weight = &v
	return s
}

func (s *CreateSdnsRuleCategoryRequestRulesRecords) Validate() error {
	return dara.Validate(s)
}
