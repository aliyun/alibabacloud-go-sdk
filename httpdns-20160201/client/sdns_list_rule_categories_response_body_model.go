// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSdnsListRuleCategoriesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetList(v []*SdnsListRuleCategoriesResponseBodyList) *SdnsListRuleCategoriesResponseBody
	GetList() []*SdnsListRuleCategoriesResponseBodyList
	SetRequestId(v string) *SdnsListRuleCategoriesResponseBody
	GetRequestId() *string
}

type SdnsListRuleCategoriesResponseBody struct {
	// The list of all custom resolution policies for the domain name, including LINE and IP_RANGE policies, as well as enabled and disabled policies. The order of the list is not guaranteed.
	List []*SdnsListRuleCategoriesResponseBodyList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	// The ID of the request, which is used to troubleshoot call issues and is not the policy ID.
	//
	// example:
	//
	// 50F9C40E-188D-4208-BE2C-74271337****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SdnsListRuleCategoriesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SdnsListRuleCategoriesResponseBody) GoString() string {
	return s.String()
}

func (s *SdnsListRuleCategoriesResponseBody) GetList() []*SdnsListRuleCategoriesResponseBodyList {
	return s.List
}

func (s *SdnsListRuleCategoriesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SdnsListRuleCategoriesResponseBody) SetList(v []*SdnsListRuleCategoriesResponseBodyList) *SdnsListRuleCategoriesResponseBody {
	s.List = v
	return s
}

func (s *SdnsListRuleCategoriesResponseBody) SetRequestId(v string) *SdnsListRuleCategoriesResponseBody {
	s.RequestId = &v
	return s
}

func (s *SdnsListRuleCategoriesResponseBody) Validate() error {
	if s.List != nil {
		for _, item := range s.List {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SdnsListRuleCategoriesResponseBodyList struct {
	// The array of original CIDR blocks from the most recent successful submission for an IP_RANGE policy. The order, duplicates, and content of each string are preserved. The response is a JSON array, not a text string of an array. This field is not used for LINE policies.
	ClientIpRanges []*string `json:"ClientIpRanges,omitempty" xml:"ClientIpRanges,omitempty" type:"Repeated"`
	// The domain name to which the policy is bound.
	//
	// example:
	//
	// api.example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The ID of the policy, which can be used to query, update, enable, disable, or delete the policy. Treat it as a string to avoid precision loss of large integers.
	//
	// example:
	//
	// 123456789012345678
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The match type of the policy. Valid values: LINE: line policy. IP_RANGE: client CIDR block policy.
	//
	// example:
	//
	// IP_RANGE
	MatchType *string `json:"MatchType,omitempty" xml:"MatchType,omitempty"`
	// The line display information for a LINE policy. This field can be empty or omitted for an IP_RANGE policy.
	RecordLine *SdnsListRuleCategoriesResponseBodyListRecordLine `json:"RecordLine,omitempty" xml:"RecordLine,omitempty" type:"Struct"`
	// The line and region display information for compatibility with existing clients. This field applies only to LINE policies and can be empty or omitted for IP_RANGE policies.
	Region *SdnsListRuleCategoriesResponseBodyListRegion `json:"Region,omitempty" xml:"Region,omitempty" type:"Struct"`
	// The list of custom resolution rules included in the policy. If multiple rules have no parameter conditions, only the first rule is prioritized. Multiple return addresses should be configured in a single rule.
	Rules []*SdnsListRuleCategoriesResponseBodyListRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
	// The status of the policy. Valid values: 0: disabled. 1: enabled. After a policy is enabled, actual resolution requires waiting for configuration propagation.
	//
	// example:
	//
	// 0
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s SdnsListRuleCategoriesResponseBodyList) String() string {
	return dara.Prettify(s)
}

func (s SdnsListRuleCategoriesResponseBodyList) GoString() string {
	return s.String()
}

func (s *SdnsListRuleCategoriesResponseBodyList) GetClientIpRanges() []*string {
	return s.ClientIpRanges
}

func (s *SdnsListRuleCategoriesResponseBodyList) GetDomainName() *string {
	return s.DomainName
}

func (s *SdnsListRuleCategoriesResponseBodyList) GetId() *string {
	return s.Id
}

func (s *SdnsListRuleCategoriesResponseBodyList) GetMatchType() *string {
	return s.MatchType
}

func (s *SdnsListRuleCategoriesResponseBodyList) GetRecordLine() *SdnsListRuleCategoriesResponseBodyListRecordLine {
	return s.RecordLine
}

func (s *SdnsListRuleCategoriesResponseBodyList) GetRegion() *SdnsListRuleCategoriesResponseBodyListRegion {
	return s.Region
}

func (s *SdnsListRuleCategoriesResponseBodyList) GetRules() []*SdnsListRuleCategoriesResponseBodyListRules {
	return s.Rules
}

func (s *SdnsListRuleCategoriesResponseBodyList) GetStatus() *int32 {
	return s.Status
}

func (s *SdnsListRuleCategoriesResponseBodyList) SetClientIpRanges(v []*string) *SdnsListRuleCategoriesResponseBodyList {
	s.ClientIpRanges = v
	return s
}

func (s *SdnsListRuleCategoriesResponseBodyList) SetDomainName(v string) *SdnsListRuleCategoriesResponseBodyList {
	s.DomainName = &v
	return s
}

func (s *SdnsListRuleCategoriesResponseBodyList) SetId(v string) *SdnsListRuleCategoriesResponseBodyList {
	s.Id = &v
	return s
}

func (s *SdnsListRuleCategoriesResponseBodyList) SetMatchType(v string) *SdnsListRuleCategoriesResponseBodyList {
	s.MatchType = &v
	return s
}

func (s *SdnsListRuleCategoriesResponseBodyList) SetRecordLine(v *SdnsListRuleCategoriesResponseBodyListRecordLine) *SdnsListRuleCategoriesResponseBodyList {
	s.RecordLine = v
	return s
}

func (s *SdnsListRuleCategoriesResponseBodyList) SetRegion(v *SdnsListRuleCategoriesResponseBodyListRegion) *SdnsListRuleCategoriesResponseBodyList {
	s.Region = v
	return s
}

func (s *SdnsListRuleCategoriesResponseBodyList) SetRules(v []*SdnsListRuleCategoriesResponseBodyListRules) *SdnsListRuleCategoriesResponseBodyList {
	s.Rules = v
	return s
}

func (s *SdnsListRuleCategoriesResponseBodyList) SetStatus(v int32) *SdnsListRuleCategoriesResponseBodyList {
	s.Status = &v
	return s
}

func (s *SdnsListRuleCategoriesResponseBodyList) Validate() error {
	if s.RecordLine != nil {
		if err := s.RecordLine.Validate(); err != nil {
			return err
		}
	}
	if s.Region != nil {
		if err := s.Region.Validate(); err != nil {
			return err
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

type SdnsListRuleCategoriesResponseBodyListRecordLine struct {
	// The line identifier. Valid values: internal: the Chinese mainland. overseas: outside the Chinese mainland.
	//
	// example:
	//
	// internal
	LineId *string `json:"LineId,omitempty" xml:"LineId,omitempty"`
	// The display name of the line.
	//
	// example:
	//
	// The Chinese mainland
	LineLabel *string `json:"LineLabel,omitempty" xml:"LineLabel,omitempty"`
	// The line selection path, which describes the line conditions at each level. This field applies only to LINE policies. IP_RANGE policies do not use line paths.
	Paths []*SdnsListRuleCategoriesResponseBodyListRecordLinePaths `json:"Paths,omitempty" xml:"Paths,omitempty" type:"Repeated"`
}

func (s SdnsListRuleCategoriesResponseBodyListRecordLine) String() string {
	return dara.Prettify(s)
}

func (s SdnsListRuleCategoriesResponseBodyListRecordLine) GoString() string {
	return s.String()
}

func (s *SdnsListRuleCategoriesResponseBodyListRecordLine) GetLineId() *string {
	return s.LineId
}

func (s *SdnsListRuleCategoriesResponseBodyListRecordLine) GetLineLabel() *string {
	return s.LineLabel
}

func (s *SdnsListRuleCategoriesResponseBodyListRecordLine) GetPaths() []*SdnsListRuleCategoriesResponseBodyListRecordLinePaths {
	return s.Paths
}

func (s *SdnsListRuleCategoriesResponseBodyListRecordLine) SetLineId(v string) *SdnsListRuleCategoriesResponseBodyListRecordLine {
	s.LineId = &v
	return s
}

func (s *SdnsListRuleCategoriesResponseBodyListRecordLine) SetLineLabel(v string) *SdnsListRuleCategoriesResponseBodyListRecordLine {
	s.LineLabel = &v
	return s
}

func (s *SdnsListRuleCategoriesResponseBodyListRecordLine) SetPaths(v []*SdnsListRuleCategoriesResponseBodyListRecordLinePaths) *SdnsListRuleCategoriesResponseBodyListRecordLine {
	s.Paths = v
	return s
}

func (s *SdnsListRuleCategoriesResponseBodyListRecordLine) Validate() error {
	if s.Paths != nil {
		for _, item := range s.Paths {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SdnsListRuleCategoriesResponseBodyListRecordLinePaths struct {
	// The code of the option at this level. For example, chinanet at the isps level indicates China Telecom. Use valid options for the associated line and maintain consistent hierarchical relationships.
	//
	// example:
	//
	// chinanet
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The level identifier. For the Chinese mainland, valid values are isps, subRegions, and provinces. For regions outside the Chinese mainland, valid values are continents and geoRegions. This corresponds to the Level field.
	//
	// example:
	//
	// isps
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The display name of the line option. The display name of the default option can be empty.
	//
	// example:
	//
	// China Telecom
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The sequence number of the line level. For the Chinese mainland, the values are 1 (ISP), 2 (region), and 3 (province). For regions outside the Chinese mainland, the values are 1 (continent) and 2 (geographical region).
	//
	// example:
	//
	// 1
	Level *int32 `json:"Level,omitempty" xml:"Level,omitempty"`
}

func (s SdnsListRuleCategoriesResponseBodyListRecordLinePaths) String() string {
	return dara.Prettify(s)
}

func (s SdnsListRuleCategoriesResponseBodyListRecordLinePaths) GoString() string {
	return s.String()
}

func (s *SdnsListRuleCategoriesResponseBodyListRecordLinePaths) GetId() *string {
	return s.Id
}

func (s *SdnsListRuleCategoriesResponseBodyListRecordLinePaths) GetKey() *string {
	return s.Key
}

func (s *SdnsListRuleCategoriesResponseBodyListRecordLinePaths) GetLabel() *string {
	return s.Label
}

func (s *SdnsListRuleCategoriesResponseBodyListRecordLinePaths) GetLevel() *int32 {
	return s.Level
}

func (s *SdnsListRuleCategoriesResponseBodyListRecordLinePaths) SetId(v string) *SdnsListRuleCategoriesResponseBodyListRecordLinePaths {
	s.Id = &v
	return s
}

func (s *SdnsListRuleCategoriesResponseBodyListRecordLinePaths) SetKey(v string) *SdnsListRuleCategoriesResponseBodyListRecordLinePaths {
	s.Key = &v
	return s
}

func (s *SdnsListRuleCategoriesResponseBodyListRecordLinePaths) SetLabel(v string) *SdnsListRuleCategoriesResponseBodyListRecordLinePaths {
	s.Label = &v
	return s
}

func (s *SdnsListRuleCategoriesResponseBodyListRecordLinePaths) SetLevel(v int32) *SdnsListRuleCategoriesResponseBodyListRecordLinePaths {
	s.Level = &v
	return s
}

func (s *SdnsListRuleCategoriesResponseBodyListRecordLinePaths) Validate() error {
	return dara.Validate(s)
}

type SdnsListRuleCategoriesResponseBodyListRegion struct {
	// The code and display name of the region.
	Area *SdnsListRuleCategoriesResponseBodyListRegionArea `json:"Area,omitempty" xml:"Area,omitempty" type:"Struct"`
	// The code and display name of the city. This is a reserved field for compatibility. City-level granularity is not currently provided.
	City *SdnsListRuleCategoriesResponseBodyListRegionCity `json:"City,omitempty" xml:"City,omitempty" type:"Struct"`
	// The code and display name of the continent.
	Continent *SdnsListRuleCategoriesResponseBodyListRegionContinent `json:"Continent,omitempty" xml:"Continent,omitempty" type:"Struct"`
	// The code and display name of the country or region.
	Country *SdnsListRuleCategoriesResponseBodyListRegionCountry `json:"Country,omitempty" xml:"Country,omitempty" type:"Struct"`
	// The code and display name of the ISP.
	Isp *SdnsListRuleCategoriesResponseBodyListRegionIsp `json:"Isp,omitempty" xml:"Isp,omitempty" type:"Struct"`
	// The code and display name of the province.
	Province *SdnsListRuleCategoriesResponseBodyListRegionProvince `json:"Province,omitempty" xml:"Province,omitempty" type:"Struct"`
}

func (s SdnsListRuleCategoriesResponseBodyListRegion) String() string {
	return dara.Prettify(s)
}

func (s SdnsListRuleCategoriesResponseBodyListRegion) GoString() string {
	return s.String()
}

func (s *SdnsListRuleCategoriesResponseBodyListRegion) GetArea() *SdnsListRuleCategoriesResponseBodyListRegionArea {
	return s.Area
}

func (s *SdnsListRuleCategoriesResponseBodyListRegion) GetCity() *SdnsListRuleCategoriesResponseBodyListRegionCity {
	return s.City
}

func (s *SdnsListRuleCategoriesResponseBodyListRegion) GetContinent() *SdnsListRuleCategoriesResponseBodyListRegionContinent {
	return s.Continent
}

func (s *SdnsListRuleCategoriesResponseBodyListRegion) GetCountry() *SdnsListRuleCategoriesResponseBodyListRegionCountry {
	return s.Country
}

func (s *SdnsListRuleCategoriesResponseBodyListRegion) GetIsp() *SdnsListRuleCategoriesResponseBodyListRegionIsp {
	return s.Isp
}

func (s *SdnsListRuleCategoriesResponseBodyListRegion) GetProvince() *SdnsListRuleCategoriesResponseBodyListRegionProvince {
	return s.Province
}

func (s *SdnsListRuleCategoriesResponseBodyListRegion) SetArea(v *SdnsListRuleCategoriesResponseBodyListRegionArea) *SdnsListRuleCategoriesResponseBodyListRegion {
	s.Area = v
	return s
}

func (s *SdnsListRuleCategoriesResponseBodyListRegion) SetCity(v *SdnsListRuleCategoriesResponseBodyListRegionCity) *SdnsListRuleCategoriesResponseBodyListRegion {
	s.City = v
	return s
}

func (s *SdnsListRuleCategoriesResponseBodyListRegion) SetContinent(v *SdnsListRuleCategoriesResponseBodyListRegionContinent) *SdnsListRuleCategoriesResponseBodyListRegion {
	s.Continent = v
	return s
}

func (s *SdnsListRuleCategoriesResponseBodyListRegion) SetCountry(v *SdnsListRuleCategoriesResponseBodyListRegionCountry) *SdnsListRuleCategoriesResponseBodyListRegion {
	s.Country = v
	return s
}

func (s *SdnsListRuleCategoriesResponseBodyListRegion) SetIsp(v *SdnsListRuleCategoriesResponseBodyListRegionIsp) *SdnsListRuleCategoriesResponseBodyListRegion {
	s.Isp = v
	return s
}

func (s *SdnsListRuleCategoriesResponseBodyListRegion) SetProvince(v *SdnsListRuleCategoriesResponseBodyListRegionProvince) *SdnsListRuleCategoriesResponseBodyListRegion {
	s.Province = v
	return s
}

func (s *SdnsListRuleCategoriesResponseBodyListRegion) Validate() error {
	if s.Area != nil {
		if err := s.Area.Validate(); err != nil {
			return err
		}
	}
	if s.City != nil {
		if err := s.City.Validate(); err != nil {
			return err
		}
	}
	if s.Continent != nil {
		if err := s.Continent.Validate(); err != nil {
			return err
		}
	}
	if s.Country != nil {
		if err := s.Country.Validate(); err != nil {
			return err
		}
	}
	if s.Isp != nil {
		if err := s.Isp.Validate(); err != nil {
			return err
		}
	}
	if s.Province != nil {
		if err := s.Province.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SdnsListRuleCategoriesResponseBodyListRegionArea struct {
	// The code of the region. The value x indicates the default item.
	//
	// example:
	//
	// huadong
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The display name of the region. An empty string can be returned for the default item.
	//
	// example:
	//
	// East China
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
}

func (s SdnsListRuleCategoriesResponseBodyListRegionArea) String() string {
	return dara.Prettify(s)
}

func (s SdnsListRuleCategoriesResponseBodyListRegionArea) GoString() string {
	return s.String()
}

func (s *SdnsListRuleCategoriesResponseBodyListRegionArea) GetId() *string {
	return s.Id
}

func (s *SdnsListRuleCategoriesResponseBodyListRegionArea) GetLabel() *string {
	return s.Label
}

func (s *SdnsListRuleCategoriesResponseBodyListRegionArea) SetId(v string) *SdnsListRuleCategoriesResponseBodyListRegionArea {
	s.Id = &v
	return s
}

func (s *SdnsListRuleCategoriesResponseBodyListRegionArea) SetLabel(v string) *SdnsListRuleCategoriesResponseBodyListRegionArea {
	s.Label = &v
	return s
}

func (s *SdnsListRuleCategoriesResponseBodyListRegionArea) Validate() error {
	return dara.Validate(s)
}

type SdnsListRuleCategoriesResponseBodyListRegionCity struct {
	// The code of the city. This is a reserved field for compatibility. The default value is x, which indicates that no city is specified. The example uses the JSON string literal "x", where the double quotes are not part of the actual field content.
	//
	// example:
	//
	// "x"
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The display name of the city. This is a reserved field for compatibility, and an empty string is returned. The example uses the JSON literal "" to represent an empty string, where the double quotes are not part of the actual field content.
	//
	// example:
	//
	// ""
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
}

func (s SdnsListRuleCategoriesResponseBodyListRegionCity) String() string {
	return dara.Prettify(s)
}

func (s SdnsListRuleCategoriesResponseBodyListRegionCity) GoString() string {
	return s.String()
}

func (s *SdnsListRuleCategoriesResponseBodyListRegionCity) GetId() *string {
	return s.Id
}

func (s *SdnsListRuleCategoriesResponseBodyListRegionCity) GetLabel() *string {
	return s.Label
}

func (s *SdnsListRuleCategoriesResponseBodyListRegionCity) SetId(v string) *SdnsListRuleCategoriesResponseBodyListRegionCity {
	s.Id = &v
	return s
}

func (s *SdnsListRuleCategoriesResponseBodyListRegionCity) SetLabel(v string) *SdnsListRuleCategoriesResponseBodyListRegionCity {
	s.Label = &v
	return s
}

func (s *SdnsListRuleCategoriesResponseBodyListRegionCity) Validate() error {
	return dara.Validate(s)
}

type SdnsListRuleCategoriesResponseBodyListRegionContinent struct {
	// The code of the continent. The value x indicates the default item.
	//
	// example:
	//
	// cn
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The display name of the continent. An empty string can be returned for the default item.
	//
	// example:
	//
	// Asia
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
}

func (s SdnsListRuleCategoriesResponseBodyListRegionContinent) String() string {
	return dara.Prettify(s)
}

func (s SdnsListRuleCategoriesResponseBodyListRegionContinent) GoString() string {
	return s.String()
}

func (s *SdnsListRuleCategoriesResponseBodyListRegionContinent) GetId() *string {
	return s.Id
}

func (s *SdnsListRuleCategoriesResponseBodyListRegionContinent) GetLabel() *string {
	return s.Label
}

func (s *SdnsListRuleCategoriesResponseBodyListRegionContinent) SetId(v string) *SdnsListRuleCategoriesResponseBodyListRegionContinent {
	s.Id = &v
	return s
}

func (s *SdnsListRuleCategoriesResponseBodyListRegionContinent) SetLabel(v string) *SdnsListRuleCategoriesResponseBodyListRegionContinent {
	s.Label = &v
	return s
}

func (s *SdnsListRuleCategoriesResponseBodyListRegionContinent) Validate() error {
	return dara.Validate(s)
}

type SdnsListRuleCategoriesResponseBodyListRegionCountry struct {
	// The code of the country or region. The value x indicates the default item.
	//
	// example:
	//
	// cn
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The display name of the country or region. An empty string can be returned for the default item.
	//
	// example:
	//
	// The Chinese mainland
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
}

func (s SdnsListRuleCategoriesResponseBodyListRegionCountry) String() string {
	return dara.Prettify(s)
}

func (s SdnsListRuleCategoriesResponseBodyListRegionCountry) GoString() string {
	return s.String()
}

func (s *SdnsListRuleCategoriesResponseBodyListRegionCountry) GetId() *string {
	return s.Id
}

func (s *SdnsListRuleCategoriesResponseBodyListRegionCountry) GetLabel() *string {
	return s.Label
}

func (s *SdnsListRuleCategoriesResponseBodyListRegionCountry) SetId(v string) *SdnsListRuleCategoriesResponseBodyListRegionCountry {
	s.Id = &v
	return s
}

func (s *SdnsListRuleCategoriesResponseBodyListRegionCountry) SetLabel(v string) *SdnsListRuleCategoriesResponseBodyListRegionCountry {
	s.Label = &v
	return s
}

func (s *SdnsListRuleCategoriesResponseBodyListRegionCountry) Validate() error {
	return dara.Validate(s)
}

type SdnsListRuleCategoriesResponseBodyListRegionIsp struct {
	// The code of the ISP. The value x indicates the default item.
	//
	// example:
	//
	// chinanet
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The display name of the ISP. An empty string can be returned for the default item.
	//
	// example:
	//
	// China Telecom
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
}

func (s SdnsListRuleCategoriesResponseBodyListRegionIsp) String() string {
	return dara.Prettify(s)
}

func (s SdnsListRuleCategoriesResponseBodyListRegionIsp) GoString() string {
	return s.String()
}

func (s *SdnsListRuleCategoriesResponseBodyListRegionIsp) GetId() *string {
	return s.Id
}

func (s *SdnsListRuleCategoriesResponseBodyListRegionIsp) GetLabel() *string {
	return s.Label
}

func (s *SdnsListRuleCategoriesResponseBodyListRegionIsp) SetId(v string) *SdnsListRuleCategoriesResponseBodyListRegionIsp {
	s.Id = &v
	return s
}

func (s *SdnsListRuleCategoriesResponseBodyListRegionIsp) SetLabel(v string) *SdnsListRuleCategoriesResponseBodyListRegionIsp {
	s.Label = &v
	return s
}

func (s *SdnsListRuleCategoriesResponseBodyListRegionIsp) Validate() error {
	return dara.Validate(s)
}

type SdnsListRuleCategoriesResponseBodyListRegionProvince struct {
	// The code of the province. The value x indicates the default item.
	//
	// example:
	//
	// zhejiang
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The display name of the province. An empty string can be returned for the default item.
	//
	// example:
	//
	// Zhejiang
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
}

func (s SdnsListRuleCategoriesResponseBodyListRegionProvince) String() string {
	return dara.Prettify(s)
}

func (s SdnsListRuleCategoriesResponseBodyListRegionProvince) GoString() string {
	return s.String()
}

func (s *SdnsListRuleCategoriesResponseBodyListRegionProvince) GetId() *string {
	return s.Id
}

func (s *SdnsListRuleCategoriesResponseBodyListRegionProvince) GetLabel() *string {
	return s.Label
}

func (s *SdnsListRuleCategoriesResponseBodyListRegionProvince) SetId(v string) *SdnsListRuleCategoriesResponseBodyListRegionProvince {
	s.Id = &v
	return s
}

func (s *SdnsListRuleCategoriesResponseBodyListRegionProvince) SetLabel(v string) *SdnsListRuleCategoriesResponseBodyListRegionProvince {
	s.Label = &v
	return s
}

func (s *SdnsListRuleCategoriesResponseBodyListRegionProvince) Validate() error {
	return dara.Validate(s)
}

type SdnsListRuleCategoriesResponseBodyListRules struct {
	// The names and display order of custom parameters. This field can be empty if no custom parameters are used.
	ArgKeys []*string `json:"ArgKeys,omitempty" xml:"ArgKeys,omitempty" type:"Repeated"`
	// The custom parameter matching conditions attached to the rule. If not configured, no parameter restrictions are applied.
	Args map[string]*string `json:"Args,omitempty" xml:"Args,omitempty"`
	// The name of the rule, which is used to identify the rule and cannot replace the policy ID.
	//
	// example:
	//
	// Default rule
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The list of configured DNS records, which can include A and AAAA records. Actual resolution returns only records of the queried type. If the list is empty or contains no records of the queried type, the system continues to fall back.
	Records []*SdnsListRuleCategoriesResponseBodyListRulesRecords `json:"Records,omitempty" xml:"Records,omitempty" type:"Repeated"`
	// The record return mode. Valid values: 0: returns all records of the queried type. 1: returns one record of the queried type based on weight.
	//
	// example:
	//
	// 0
	ResultMode *int32 `json:"ResultMode,omitempty" xml:"ResultMode,omitempty"`
	// The cache validity period of the record, in seconds. This value must be a positive integer.
	//
	// example:
	//
	// 60
	Ttl *int32 `json:"Ttl,omitempty" xml:"Ttl,omitempty"`
}

func (s SdnsListRuleCategoriesResponseBodyListRules) String() string {
	return dara.Prettify(s)
}

func (s SdnsListRuleCategoriesResponseBodyListRules) GoString() string {
	return s.String()
}

func (s *SdnsListRuleCategoriesResponseBodyListRules) GetArgKeys() []*string {
	return s.ArgKeys
}

func (s *SdnsListRuleCategoriesResponseBodyListRules) GetArgs() map[string]*string {
	return s.Args
}

func (s *SdnsListRuleCategoriesResponseBodyListRules) GetName() *string {
	return s.Name
}

func (s *SdnsListRuleCategoriesResponseBodyListRules) GetRecords() []*SdnsListRuleCategoriesResponseBodyListRulesRecords {
	return s.Records
}

func (s *SdnsListRuleCategoriesResponseBodyListRules) GetResultMode() *int32 {
	return s.ResultMode
}

func (s *SdnsListRuleCategoriesResponseBodyListRules) GetTtl() *int32 {
	return s.Ttl
}

func (s *SdnsListRuleCategoriesResponseBodyListRules) SetArgKeys(v []*string) *SdnsListRuleCategoriesResponseBodyListRules {
	s.ArgKeys = v
	return s
}

func (s *SdnsListRuleCategoriesResponseBodyListRules) SetArgs(v map[string]*string) *SdnsListRuleCategoriesResponseBodyListRules {
	s.Args = v
	return s
}

func (s *SdnsListRuleCategoriesResponseBodyListRules) SetName(v string) *SdnsListRuleCategoriesResponseBodyListRules {
	s.Name = &v
	return s
}

func (s *SdnsListRuleCategoriesResponseBodyListRules) SetRecords(v []*SdnsListRuleCategoriesResponseBodyListRulesRecords) *SdnsListRuleCategoriesResponseBodyListRules {
	s.Records = v
	return s
}

func (s *SdnsListRuleCategoriesResponseBodyListRules) SetResultMode(v int32) *SdnsListRuleCategoriesResponseBodyListRules {
	s.ResultMode = &v
	return s
}

func (s *SdnsListRuleCategoriesResponseBodyListRules) SetTtl(v int32) *SdnsListRuleCategoriesResponseBodyListRules {
	s.Ttl = &v
	return s
}

func (s *SdnsListRuleCategoriesResponseBodyListRules) Validate() error {
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

type SdnsListRuleCategoriesResponseBodyListRulesRecords struct {
	// The single destination IP address corresponding to the record. When Type is 1, it is an IPv4 address, such as 192.0.2.200. When Type is 28, it is an IPv6 address, such as 2001:db8::200. CIDR blocks, IP ranges, or domain names are not used.
	//
	// example:
	//
	// 192.0.2.200
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The DNS record type. Valid values: 1: A record. 28: AAAA record. Enter an integer instead of the string A or AAAA.
	//
	// example:
	//
	// 1
	Type *int32 `json:"Type,omitempty" xml:"Type,omitempty"`
	// The auxiliary display text for the record type, which is not used for matching or resolution decisions. This field can be omitted in requests.
	//
	// example:
	//
	// A
	TypeLabel *string `json:"TypeLabel,omitempty" xml:"TypeLabel,omitempty"`
	// The relative weight of the record, which is used to select one record of the same query type only when ResultMode is 1. Set this to a positive integer from 1 to 100. The sum of the weights of all records does not need to be 100. This field is not used for selection when ResultMode is 0.
	//
	// example:
	//
	// 100
	Weight *int32 `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s SdnsListRuleCategoriesResponseBodyListRulesRecords) String() string {
	return dara.Prettify(s)
}

func (s SdnsListRuleCategoriesResponseBodyListRulesRecords) GoString() string {
	return s.String()
}

func (s *SdnsListRuleCategoriesResponseBodyListRulesRecords) GetIp() *string {
	return s.Ip
}

func (s *SdnsListRuleCategoriesResponseBodyListRulesRecords) GetType() *int32 {
	return s.Type
}

func (s *SdnsListRuleCategoriesResponseBodyListRulesRecords) GetTypeLabel() *string {
	return s.TypeLabel
}

func (s *SdnsListRuleCategoriesResponseBodyListRulesRecords) GetWeight() *int32 {
	return s.Weight
}

func (s *SdnsListRuleCategoriesResponseBodyListRulesRecords) SetIp(v string) *SdnsListRuleCategoriesResponseBodyListRulesRecords {
	s.Ip = &v
	return s
}

func (s *SdnsListRuleCategoriesResponseBodyListRulesRecords) SetType(v int32) *SdnsListRuleCategoriesResponseBodyListRulesRecords {
	s.Type = &v
	return s
}

func (s *SdnsListRuleCategoriesResponseBodyListRulesRecords) SetTypeLabel(v string) *SdnsListRuleCategoriesResponseBodyListRulesRecords {
	s.TypeLabel = &v
	return s
}

func (s *SdnsListRuleCategoriesResponseBodyListRulesRecords) SetWeight(v int32) *SdnsListRuleCategoriesResponseBodyListRulesRecords {
	s.Weight = &v
	return s
}

func (s *SdnsListRuleCategoriesResponseBodyListRulesRecords) Validate() error {
	return dara.Validate(s)
}
