// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSdnsRuleCategoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientIpRanges(v string) *UpdateSdnsRuleCategoryRequest
	GetClientIpRanges() *string
	SetId(v string) *UpdateSdnsRuleCategoryRequest
	GetId() *string
	SetRules(v []*UpdateSdnsRuleCategoryRequestRules) *UpdateSdnsRuleCategoryRequest
	GetRules() []*UpdateSdnsRuleCategoryRequestRules
}

type UpdateSdnsRuleCategoryRequest struct {
	// The JSON text string of the IP address range array. This parameter is used only to update IP_RANGE policies. If this parameter is omitted, the original IP address ranges are retained. If this parameter is provided, the original IP address ranges are replaced as a whole. You cannot pass an empty string, an empty array [], null, or a non-string array. Single IP addresses, CIDR blocks, and start-end ranges of IPv4 and IPv6 are supported. Multiple entries have an OR relationship. For CIDR blocks, the IPv4 prefix length ranges from 0 to 32, and the IPv6 prefix length ranges from 0 to 128. Start-end ranges are inclusive and must belong to the same address family, and the start address must not be greater than the end address. Overlapping, containing, and duplicate ranges are allowed within the same policy. When the policy is saved and read back, the array order, duplicate items, and the original text of each string are retained. An array can contain up to 2,000 array elements. The total number of parsed non-empty address expressions ranges from 1 to 2,000, and duplicate items are counted. We recommend that you specify one expression per element. It is compatible with line breaks within an element. Each expression can contain up to 128 characters, including leading and trailing whitespace. The serialized UTF-8 JSON sequence cannot exceed 65,535 bytes. The quantity and byte limits take effect at the same time. Overlapping is not allowed among different IP_RANGE policies of the same account, the same domain name, and the same address family. Disabled policies are also checked. Disabling a policy does not release its IP address ranges. This limit does not compare LINE policies and the independent configurations of the legacy SetIPRegionResult operation. The new ranges do not conflict with the old ranges of the same policy during an update. Send this parameter as a single form field, and do not add quotes to the JSON text again.
	//
	// example:
	//
	// ["198.51.100.0/24","2001:db8:1::/64"]
	ClientIpRanges *string `json:"ClientIpRanges,omitempty" xml:"ClientIpRanges,omitempty"`
	// The ID of the policy for the current account. This parameter is required. You can obtain the ID from the Id field in the response of CreateSdnsRuleCategory or from the list returned by SdnsListRuleCategories. Use a decimal string to avoid precision loss for large integers.
	//
	// example:
	//
	// 123456789012345678
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The complete list of rules. This parameter is required and must contain at least one rule. If you match only by IP address range, you can configure a default rule without ArgKeys and Args, and configure the A or AAAA destination addresses in the Records field of the rule. The update replaces the original rule list as a whole. Old rules that are not included in the new list are removed. When the rules are saved, the serialized UTF-8 JSON content must be less than 65,535 bytes.
	Rules []*UpdateSdnsRuleCategoryRequestRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
}

func (s UpdateSdnsRuleCategoryRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateSdnsRuleCategoryRequest) GoString() string {
	return s.String()
}

func (s *UpdateSdnsRuleCategoryRequest) GetClientIpRanges() *string {
	return s.ClientIpRanges
}

func (s *UpdateSdnsRuleCategoryRequest) GetId() *string {
	return s.Id
}

func (s *UpdateSdnsRuleCategoryRequest) GetRules() []*UpdateSdnsRuleCategoryRequestRules {
	return s.Rules
}

func (s *UpdateSdnsRuleCategoryRequest) SetClientIpRanges(v string) *UpdateSdnsRuleCategoryRequest {
	s.ClientIpRanges = &v
	return s
}

func (s *UpdateSdnsRuleCategoryRequest) SetId(v string) *UpdateSdnsRuleCategoryRequest {
	s.Id = &v
	return s
}

func (s *UpdateSdnsRuleCategoryRequest) SetRules(v []*UpdateSdnsRuleCategoryRequestRules) *UpdateSdnsRuleCategoryRequest {
	s.Rules = v
	return s
}

func (s *UpdateSdnsRuleCategoryRequest) Validate() error {
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

type UpdateSdnsRuleCategoryRequestRules struct {
	// The optional list of custom parameter names, which is used to record the display order. Omit this parameter if you match only by IP address range. If you use Args, the names must correspond to the parameter names in Args.
	ArgKeys []*string `json:"ArgKeys,omitempty" xml:"ArgKeys,omitempty" type:"Repeated"`
	// The optional custom parameter conditions. The key is the parameter name, and the value is the match value. Omit this parameter if you match only by IP address range. You do not need to add business parameters.
	Args map[string]*string `json:"Args,omitempty" xml:"Args,omitempty"`
	// The name of the rule. This parameter is required and cannot be null. We recommend that you use a non-empty name that has business meaning. The name cannot replace the policy ID.
	//
	// example:
	//
	// Default rule
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The list of configured DNS records. The list can include A and AAAA records. Actual DNS resolution returns only the records of the queried type. If the list is empty or does not contain records of the queried type, the system continues to fall back.
	Records []*UpdateSdnsRuleCategoryRequestRulesRecords `json:"Records,omitempty" xml:"Records,omitempty" type:"Repeated"`
	// The record return mode. Valid values:
	//
	// - 0: Returns all records of the queried type.
	//
	// - 1: Returns one record of the queried type based on the weight.
	//
	// example:
	//
	// 0
	ResultMode *int32 `json:"ResultMode,omitempty" xml:"ResultMode,omitempty"`
	// The cache validity period of the record, in seconds. The value must be a positive integer.
	//
	// example:
	//
	// 60
	Ttl *int32 `json:"Ttl,omitempty" xml:"Ttl,omitempty"`
}

func (s UpdateSdnsRuleCategoryRequestRules) String() string {
	return dara.Prettify(s)
}

func (s UpdateSdnsRuleCategoryRequestRules) GoString() string {
	return s.String()
}

func (s *UpdateSdnsRuleCategoryRequestRules) GetArgKeys() []*string {
	return s.ArgKeys
}

func (s *UpdateSdnsRuleCategoryRequestRules) GetArgs() map[string]*string {
	return s.Args
}

func (s *UpdateSdnsRuleCategoryRequestRules) GetName() *string {
	return s.Name
}

func (s *UpdateSdnsRuleCategoryRequestRules) GetRecords() []*UpdateSdnsRuleCategoryRequestRulesRecords {
	return s.Records
}

func (s *UpdateSdnsRuleCategoryRequestRules) GetResultMode() *int32 {
	return s.ResultMode
}

func (s *UpdateSdnsRuleCategoryRequestRules) GetTtl() *int32 {
	return s.Ttl
}

func (s *UpdateSdnsRuleCategoryRequestRules) SetArgKeys(v []*string) *UpdateSdnsRuleCategoryRequestRules {
	s.ArgKeys = v
	return s
}

func (s *UpdateSdnsRuleCategoryRequestRules) SetArgs(v map[string]*string) *UpdateSdnsRuleCategoryRequestRules {
	s.Args = v
	return s
}

func (s *UpdateSdnsRuleCategoryRequestRules) SetName(v string) *UpdateSdnsRuleCategoryRequestRules {
	s.Name = &v
	return s
}

func (s *UpdateSdnsRuleCategoryRequestRules) SetRecords(v []*UpdateSdnsRuleCategoryRequestRulesRecords) *UpdateSdnsRuleCategoryRequestRules {
	s.Records = v
	return s
}

func (s *UpdateSdnsRuleCategoryRequestRules) SetResultMode(v int32) *UpdateSdnsRuleCategoryRequestRules {
	s.ResultMode = &v
	return s
}

func (s *UpdateSdnsRuleCategoryRequestRules) SetTtl(v int32) *UpdateSdnsRuleCategoryRequestRules {
	s.Ttl = &v
	return s
}

func (s *UpdateSdnsRuleCategoryRequestRules) Validate() error {
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

type UpdateSdnsRuleCategoryRequestRulesRecords struct {
	// The single destination IP address corresponding to the record. When Type is set to 1, the IP address is an IPv4 address, such as 192.0.2.200. When Type is set to 28, the IP address is an IPv6 address, such as 2001:db8::200. Do not use CIDR blocks, start-end ranges, or domain names.
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
	// The auxiliary display text for the record type. This text is not used for matching or DNS resolution decisions. You can omit this parameter in the request.
	//
	// example:
	//
	// A
	TypeLabel *string `json:"TypeLabel,omitempty" xml:"TypeLabel,omitempty"`
	// The relative weight of the record. This parameter is used to select one record of the same queried type only when ResultMode is set to 1. We recommend that you set this parameter to a positive integer from 1 to 100. The sum of the weights of all records does not need to be 100. This parameter is not used for selection when ResultMode is set to 0.
	//
	// example:
	//
	// 100
	Weight *int32 `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s UpdateSdnsRuleCategoryRequestRulesRecords) String() string {
	return dara.Prettify(s)
}

func (s UpdateSdnsRuleCategoryRequestRulesRecords) GoString() string {
	return s.String()
}

func (s *UpdateSdnsRuleCategoryRequestRulesRecords) GetIp() *string {
	return s.Ip
}

func (s *UpdateSdnsRuleCategoryRequestRulesRecords) GetType() *int32 {
	return s.Type
}

func (s *UpdateSdnsRuleCategoryRequestRulesRecords) GetTypeLabel() *string {
	return s.TypeLabel
}

func (s *UpdateSdnsRuleCategoryRequestRulesRecords) GetWeight() *int32 {
	return s.Weight
}

func (s *UpdateSdnsRuleCategoryRequestRulesRecords) SetIp(v string) *UpdateSdnsRuleCategoryRequestRulesRecords {
	s.Ip = &v
	return s
}

func (s *UpdateSdnsRuleCategoryRequestRulesRecords) SetType(v int32) *UpdateSdnsRuleCategoryRequestRulesRecords {
	s.Type = &v
	return s
}

func (s *UpdateSdnsRuleCategoryRequestRulesRecords) SetTypeLabel(v string) *UpdateSdnsRuleCategoryRequestRulesRecords {
	s.TypeLabel = &v
	return s
}

func (s *UpdateSdnsRuleCategoryRequestRulesRecords) SetWeight(v int32) *UpdateSdnsRuleCategoryRequestRulesRecords {
	s.Weight = &v
	return s
}

func (s *UpdateSdnsRuleCategoryRequestRulesRecords) Validate() error {
	return dara.Validate(s)
}
