// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSdnsVerifyRuleCategoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *SdnsVerifyRuleCategoryResponseBodyData) *SdnsVerifyRuleCategoryResponseBody
	GetData() *SdnsVerifyRuleCategoryResponseBodyData
	SetRequestId(v string) *SdnsVerifyRuleCategoryResponseBody
	GetRequestId() *string
}

type SdnsVerifyRuleCategoryResponseBody struct {
	// The configuration of the matched rule. If no rules are matched, this field is empty or omitted. This field does not represent the final response of the actual resolution chain.
	Data *SdnsVerifyRuleCategoryResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The ID of the request, used to troubleshoot call issues. This field is not the policy ID.
	//
	// example:
	//
	// 50F9C40E-188D-4208-BE2C-74271337****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SdnsVerifyRuleCategoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SdnsVerifyRuleCategoryResponseBody) GoString() string {
	return s.String()
}

func (s *SdnsVerifyRuleCategoryResponseBody) GetData() *SdnsVerifyRuleCategoryResponseBodyData {
	return s.Data
}

func (s *SdnsVerifyRuleCategoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SdnsVerifyRuleCategoryResponseBody) SetData(v *SdnsVerifyRuleCategoryResponseBodyData) *SdnsVerifyRuleCategoryResponseBody {
	s.Data = v
	return s
}

func (s *SdnsVerifyRuleCategoryResponseBody) SetRequestId(v string) *SdnsVerifyRuleCategoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *SdnsVerifyRuleCategoryResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SdnsVerifyRuleCategoryResponseBodyData struct {
	// The custom parameter names and their display order. This field can be empty if no custom parameters are used.
	ArgKeys []*string `json:"ArgKeys,omitempty" xml:"ArgKeys,omitempty" type:"Repeated"`
	// The custom parameter match conditions attached to the rule. If not configured, no parameter restrictions are attached.
	Args map[string]*string `json:"Args,omitempty" xml:"Args,omitempty"`
	// The rule name, used to identify the rule. This field cannot replace the policy ID.
	//
	// example:
	//
	// Default rule
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The list of configured DNS records, which can include A and AAAA records. Actual DNS resolution returns only records of the queried type. If the list is empty or does not contain records of the queried type, the system continues to fall back. This operation returns the records configured in the rule and does not filter them by query type.
	Records []*SdnsVerifyRuleCategoryResponseBodyDataRecords `json:"Records,omitempty" xml:"Records,omitempty" type:"Repeated"`
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

func (s SdnsVerifyRuleCategoryResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s SdnsVerifyRuleCategoryResponseBodyData) GoString() string {
	return s.String()
}

func (s *SdnsVerifyRuleCategoryResponseBodyData) GetArgKeys() []*string {
	return s.ArgKeys
}

func (s *SdnsVerifyRuleCategoryResponseBodyData) GetArgs() map[string]*string {
	return s.Args
}

func (s *SdnsVerifyRuleCategoryResponseBodyData) GetName() *string {
	return s.Name
}

func (s *SdnsVerifyRuleCategoryResponseBodyData) GetRecords() []*SdnsVerifyRuleCategoryResponseBodyDataRecords {
	return s.Records
}

func (s *SdnsVerifyRuleCategoryResponseBodyData) GetResultMode() *int32 {
	return s.ResultMode
}

func (s *SdnsVerifyRuleCategoryResponseBodyData) GetTtl() *int32 {
	return s.Ttl
}

func (s *SdnsVerifyRuleCategoryResponseBodyData) SetArgKeys(v []*string) *SdnsVerifyRuleCategoryResponseBodyData {
	s.ArgKeys = v
	return s
}

func (s *SdnsVerifyRuleCategoryResponseBodyData) SetArgs(v map[string]*string) *SdnsVerifyRuleCategoryResponseBodyData {
	s.Args = v
	return s
}

func (s *SdnsVerifyRuleCategoryResponseBodyData) SetName(v string) *SdnsVerifyRuleCategoryResponseBodyData {
	s.Name = &v
	return s
}

func (s *SdnsVerifyRuleCategoryResponseBodyData) SetRecords(v []*SdnsVerifyRuleCategoryResponseBodyDataRecords) *SdnsVerifyRuleCategoryResponseBodyData {
	s.Records = v
	return s
}

func (s *SdnsVerifyRuleCategoryResponseBodyData) SetResultMode(v int32) *SdnsVerifyRuleCategoryResponseBodyData {
	s.ResultMode = &v
	return s
}

func (s *SdnsVerifyRuleCategoryResponseBodyData) SetTtl(v int32) *SdnsVerifyRuleCategoryResponseBodyData {
	s.Ttl = &v
	return s
}

func (s *SdnsVerifyRuleCategoryResponseBodyData) Validate() error {
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

type SdnsVerifyRuleCategoryResponseBodyDataRecords struct {
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

func (s SdnsVerifyRuleCategoryResponseBodyDataRecords) String() string {
	return dara.Prettify(s)
}

func (s SdnsVerifyRuleCategoryResponseBodyDataRecords) GoString() string {
	return s.String()
}

func (s *SdnsVerifyRuleCategoryResponseBodyDataRecords) GetIp() *string {
	return s.Ip
}

func (s *SdnsVerifyRuleCategoryResponseBodyDataRecords) GetType() *int32 {
	return s.Type
}

func (s *SdnsVerifyRuleCategoryResponseBodyDataRecords) GetTypeLabel() *string {
	return s.TypeLabel
}

func (s *SdnsVerifyRuleCategoryResponseBodyDataRecords) GetWeight() *int32 {
	return s.Weight
}

func (s *SdnsVerifyRuleCategoryResponseBodyDataRecords) SetIp(v string) *SdnsVerifyRuleCategoryResponseBodyDataRecords {
	s.Ip = &v
	return s
}

func (s *SdnsVerifyRuleCategoryResponseBodyDataRecords) SetType(v int32) *SdnsVerifyRuleCategoryResponseBodyDataRecords {
	s.Type = &v
	return s
}

func (s *SdnsVerifyRuleCategoryResponseBodyDataRecords) SetTypeLabel(v string) *SdnsVerifyRuleCategoryResponseBodyDataRecords {
	s.TypeLabel = &v
	return s
}

func (s *SdnsVerifyRuleCategoryResponseBodyDataRecords) SetWeight(v int32) *SdnsVerifyRuleCategoryResponseBodyDataRecords {
	s.Weight = &v
	return s
}

func (s *SdnsVerifyRuleCategoryResponseBodyDataRecords) Validate() error {
	return dara.Validate(s)
}
