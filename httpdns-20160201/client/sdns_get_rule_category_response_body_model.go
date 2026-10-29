// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSdnsGetRuleCategoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *SdnsGetRuleCategoryResponseBodyData) *SdnsGetRuleCategoryResponseBody
	GetData() *SdnsGetRuleCategoryResponseBodyData
	SetRequestId(v string) *SdnsGetRuleCategoryResponseBody
	GetRequestId() *string
}

type SdnsGetRuleCategoryResponseBody struct {
	Data *SdnsGetRuleCategoryResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 50F9C40E-188D-4208-BE2C-74271337****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SdnsGetRuleCategoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SdnsGetRuleCategoryResponseBody) GoString() string {
	return s.String()
}

func (s *SdnsGetRuleCategoryResponseBody) GetData() *SdnsGetRuleCategoryResponseBodyData {
	return s.Data
}

func (s *SdnsGetRuleCategoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SdnsGetRuleCategoryResponseBody) SetData(v *SdnsGetRuleCategoryResponseBodyData) *SdnsGetRuleCategoryResponseBody {
	s.Data = v
	return s
}

func (s *SdnsGetRuleCategoryResponseBody) SetRequestId(v string) *SdnsGetRuleCategoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *SdnsGetRuleCategoryResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SdnsGetRuleCategoryResponseBodyData struct {
	ClientIpRanges []*string `json:"ClientIpRanges,omitempty" xml:"ClientIpRanges,omitempty" type:"Repeated"`
	// example:
	//
	// api.example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// example:
	//
	// 123456789012345678
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// IP_RANGE
	MatchType  *string                                        `json:"MatchType,omitempty" xml:"MatchType,omitempty"`
	RecordLine *SdnsGetRuleCategoryResponseBodyDataRecordLine `json:"RecordLine,omitempty" xml:"RecordLine,omitempty" type:"Struct"`
	Region     *SdnsGetRuleCategoryResponseBodyDataRegion     `json:"Region,omitempty" xml:"Region,omitempty" type:"Struct"`
	Rules      []*SdnsGetRuleCategoryResponseBodyDataRules    `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
	// example:
	//
	// 0
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s SdnsGetRuleCategoryResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s SdnsGetRuleCategoryResponseBodyData) GoString() string {
	return s.String()
}

func (s *SdnsGetRuleCategoryResponseBodyData) GetClientIpRanges() []*string {
	return s.ClientIpRanges
}

func (s *SdnsGetRuleCategoryResponseBodyData) GetDomainName() *string {
	return s.DomainName
}

func (s *SdnsGetRuleCategoryResponseBodyData) GetId() *string {
	return s.Id
}

func (s *SdnsGetRuleCategoryResponseBodyData) GetMatchType() *string {
	return s.MatchType
}

func (s *SdnsGetRuleCategoryResponseBodyData) GetRecordLine() *SdnsGetRuleCategoryResponseBodyDataRecordLine {
	return s.RecordLine
}

func (s *SdnsGetRuleCategoryResponseBodyData) GetRegion() *SdnsGetRuleCategoryResponseBodyDataRegion {
	return s.Region
}

func (s *SdnsGetRuleCategoryResponseBodyData) GetRules() []*SdnsGetRuleCategoryResponseBodyDataRules {
	return s.Rules
}

func (s *SdnsGetRuleCategoryResponseBodyData) GetStatus() *int32 {
	return s.Status
}

func (s *SdnsGetRuleCategoryResponseBodyData) SetClientIpRanges(v []*string) *SdnsGetRuleCategoryResponseBodyData {
	s.ClientIpRanges = v
	return s
}

func (s *SdnsGetRuleCategoryResponseBodyData) SetDomainName(v string) *SdnsGetRuleCategoryResponseBodyData {
	s.DomainName = &v
	return s
}

func (s *SdnsGetRuleCategoryResponseBodyData) SetId(v string) *SdnsGetRuleCategoryResponseBodyData {
	s.Id = &v
	return s
}

func (s *SdnsGetRuleCategoryResponseBodyData) SetMatchType(v string) *SdnsGetRuleCategoryResponseBodyData {
	s.MatchType = &v
	return s
}

func (s *SdnsGetRuleCategoryResponseBodyData) SetRecordLine(v *SdnsGetRuleCategoryResponseBodyDataRecordLine) *SdnsGetRuleCategoryResponseBodyData {
	s.RecordLine = v
	return s
}

func (s *SdnsGetRuleCategoryResponseBodyData) SetRegion(v *SdnsGetRuleCategoryResponseBodyDataRegion) *SdnsGetRuleCategoryResponseBodyData {
	s.Region = v
	return s
}

func (s *SdnsGetRuleCategoryResponseBodyData) SetRules(v []*SdnsGetRuleCategoryResponseBodyDataRules) *SdnsGetRuleCategoryResponseBodyData {
	s.Rules = v
	return s
}

func (s *SdnsGetRuleCategoryResponseBodyData) SetStatus(v int32) *SdnsGetRuleCategoryResponseBodyData {
	s.Status = &v
	return s
}

func (s *SdnsGetRuleCategoryResponseBodyData) Validate() error {
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

type SdnsGetRuleCategoryResponseBodyDataRecordLine struct {
	// example:
	//
	// internal
	LineId *string `json:"LineId,omitempty" xml:"LineId,omitempty"`
	// example:
	//
	// 中国地区
	LineLabel *string                                               `json:"LineLabel,omitempty" xml:"LineLabel,omitempty"`
	Paths     []*SdnsGetRuleCategoryResponseBodyDataRecordLinePaths `json:"Paths,omitempty" xml:"Paths,omitempty" type:"Repeated"`
}

func (s SdnsGetRuleCategoryResponseBodyDataRecordLine) String() string {
	return dara.Prettify(s)
}

func (s SdnsGetRuleCategoryResponseBodyDataRecordLine) GoString() string {
	return s.String()
}

func (s *SdnsGetRuleCategoryResponseBodyDataRecordLine) GetLineId() *string {
	return s.LineId
}

func (s *SdnsGetRuleCategoryResponseBodyDataRecordLine) GetLineLabel() *string {
	return s.LineLabel
}

func (s *SdnsGetRuleCategoryResponseBodyDataRecordLine) GetPaths() []*SdnsGetRuleCategoryResponseBodyDataRecordLinePaths {
	return s.Paths
}

func (s *SdnsGetRuleCategoryResponseBodyDataRecordLine) SetLineId(v string) *SdnsGetRuleCategoryResponseBodyDataRecordLine {
	s.LineId = &v
	return s
}

func (s *SdnsGetRuleCategoryResponseBodyDataRecordLine) SetLineLabel(v string) *SdnsGetRuleCategoryResponseBodyDataRecordLine {
	s.LineLabel = &v
	return s
}

func (s *SdnsGetRuleCategoryResponseBodyDataRecordLine) SetPaths(v []*SdnsGetRuleCategoryResponseBodyDataRecordLinePaths) *SdnsGetRuleCategoryResponseBodyDataRecordLine {
	s.Paths = v
	return s
}

func (s *SdnsGetRuleCategoryResponseBodyDataRecordLine) Validate() error {
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

type SdnsGetRuleCategoryResponseBodyDataRecordLinePaths struct {
	// example:
	//
	// chinanet
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// isps
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// 电信
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// example:
	//
	// 1
	Level *int32 `json:"Level,omitempty" xml:"Level,omitempty"`
}

func (s SdnsGetRuleCategoryResponseBodyDataRecordLinePaths) String() string {
	return dara.Prettify(s)
}

func (s SdnsGetRuleCategoryResponseBodyDataRecordLinePaths) GoString() string {
	return s.String()
}

func (s *SdnsGetRuleCategoryResponseBodyDataRecordLinePaths) GetId() *string {
	return s.Id
}

func (s *SdnsGetRuleCategoryResponseBodyDataRecordLinePaths) GetKey() *string {
	return s.Key
}

func (s *SdnsGetRuleCategoryResponseBodyDataRecordLinePaths) GetLabel() *string {
	return s.Label
}

func (s *SdnsGetRuleCategoryResponseBodyDataRecordLinePaths) GetLevel() *int32 {
	return s.Level
}

func (s *SdnsGetRuleCategoryResponseBodyDataRecordLinePaths) SetId(v string) *SdnsGetRuleCategoryResponseBodyDataRecordLinePaths {
	s.Id = &v
	return s
}

func (s *SdnsGetRuleCategoryResponseBodyDataRecordLinePaths) SetKey(v string) *SdnsGetRuleCategoryResponseBodyDataRecordLinePaths {
	s.Key = &v
	return s
}

func (s *SdnsGetRuleCategoryResponseBodyDataRecordLinePaths) SetLabel(v string) *SdnsGetRuleCategoryResponseBodyDataRecordLinePaths {
	s.Label = &v
	return s
}

func (s *SdnsGetRuleCategoryResponseBodyDataRecordLinePaths) SetLevel(v int32) *SdnsGetRuleCategoryResponseBodyDataRecordLinePaths {
	s.Level = &v
	return s
}

func (s *SdnsGetRuleCategoryResponseBodyDataRecordLinePaths) Validate() error {
	return dara.Validate(s)
}

type SdnsGetRuleCategoryResponseBodyDataRegion struct {
	Area      *SdnsGetRuleCategoryResponseBodyDataRegionArea      `json:"Area,omitempty" xml:"Area,omitempty" type:"Struct"`
	City      *SdnsGetRuleCategoryResponseBodyDataRegionCity      `json:"City,omitempty" xml:"City,omitempty" type:"Struct"`
	Continent *SdnsGetRuleCategoryResponseBodyDataRegionContinent `json:"Continent,omitempty" xml:"Continent,omitempty" type:"Struct"`
	Country   *SdnsGetRuleCategoryResponseBodyDataRegionCountry   `json:"Country,omitempty" xml:"Country,omitempty" type:"Struct"`
	Isp       *SdnsGetRuleCategoryResponseBodyDataRegionIsp       `json:"Isp,omitempty" xml:"Isp,omitempty" type:"Struct"`
	Province  *SdnsGetRuleCategoryResponseBodyDataRegionProvince  `json:"Province,omitempty" xml:"Province,omitempty" type:"Struct"`
}

func (s SdnsGetRuleCategoryResponseBodyDataRegion) String() string {
	return dara.Prettify(s)
}

func (s SdnsGetRuleCategoryResponseBodyDataRegion) GoString() string {
	return s.String()
}

func (s *SdnsGetRuleCategoryResponseBodyDataRegion) GetArea() *SdnsGetRuleCategoryResponseBodyDataRegionArea {
	return s.Area
}

func (s *SdnsGetRuleCategoryResponseBodyDataRegion) GetCity() *SdnsGetRuleCategoryResponseBodyDataRegionCity {
	return s.City
}

func (s *SdnsGetRuleCategoryResponseBodyDataRegion) GetContinent() *SdnsGetRuleCategoryResponseBodyDataRegionContinent {
	return s.Continent
}

func (s *SdnsGetRuleCategoryResponseBodyDataRegion) GetCountry() *SdnsGetRuleCategoryResponseBodyDataRegionCountry {
	return s.Country
}

func (s *SdnsGetRuleCategoryResponseBodyDataRegion) GetIsp() *SdnsGetRuleCategoryResponseBodyDataRegionIsp {
	return s.Isp
}

func (s *SdnsGetRuleCategoryResponseBodyDataRegion) GetProvince() *SdnsGetRuleCategoryResponseBodyDataRegionProvince {
	return s.Province
}

func (s *SdnsGetRuleCategoryResponseBodyDataRegion) SetArea(v *SdnsGetRuleCategoryResponseBodyDataRegionArea) *SdnsGetRuleCategoryResponseBodyDataRegion {
	s.Area = v
	return s
}

func (s *SdnsGetRuleCategoryResponseBodyDataRegion) SetCity(v *SdnsGetRuleCategoryResponseBodyDataRegionCity) *SdnsGetRuleCategoryResponseBodyDataRegion {
	s.City = v
	return s
}

func (s *SdnsGetRuleCategoryResponseBodyDataRegion) SetContinent(v *SdnsGetRuleCategoryResponseBodyDataRegionContinent) *SdnsGetRuleCategoryResponseBodyDataRegion {
	s.Continent = v
	return s
}

func (s *SdnsGetRuleCategoryResponseBodyDataRegion) SetCountry(v *SdnsGetRuleCategoryResponseBodyDataRegionCountry) *SdnsGetRuleCategoryResponseBodyDataRegion {
	s.Country = v
	return s
}

func (s *SdnsGetRuleCategoryResponseBodyDataRegion) SetIsp(v *SdnsGetRuleCategoryResponseBodyDataRegionIsp) *SdnsGetRuleCategoryResponseBodyDataRegion {
	s.Isp = v
	return s
}

func (s *SdnsGetRuleCategoryResponseBodyDataRegion) SetProvince(v *SdnsGetRuleCategoryResponseBodyDataRegionProvince) *SdnsGetRuleCategoryResponseBodyDataRegion {
	s.Province = v
	return s
}

func (s *SdnsGetRuleCategoryResponseBodyDataRegion) Validate() error {
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

type SdnsGetRuleCategoryResponseBodyDataRegionArea struct {
	// example:
	//
	// huadong
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 华东
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
}

func (s SdnsGetRuleCategoryResponseBodyDataRegionArea) String() string {
	return dara.Prettify(s)
}

func (s SdnsGetRuleCategoryResponseBodyDataRegionArea) GoString() string {
	return s.String()
}

func (s *SdnsGetRuleCategoryResponseBodyDataRegionArea) GetId() *string {
	return s.Id
}

func (s *SdnsGetRuleCategoryResponseBodyDataRegionArea) GetLabel() *string {
	return s.Label
}

func (s *SdnsGetRuleCategoryResponseBodyDataRegionArea) SetId(v string) *SdnsGetRuleCategoryResponseBodyDataRegionArea {
	s.Id = &v
	return s
}

func (s *SdnsGetRuleCategoryResponseBodyDataRegionArea) SetLabel(v string) *SdnsGetRuleCategoryResponseBodyDataRegionArea {
	s.Label = &v
	return s
}

func (s *SdnsGetRuleCategoryResponseBodyDataRegionArea) Validate() error {
	return dara.Validate(s)
}

type SdnsGetRuleCategoryResponseBodyDataRegionCity struct {
	// example:
	//
	// "x"
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// ""
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
}

func (s SdnsGetRuleCategoryResponseBodyDataRegionCity) String() string {
	return dara.Prettify(s)
}

func (s SdnsGetRuleCategoryResponseBodyDataRegionCity) GoString() string {
	return s.String()
}

func (s *SdnsGetRuleCategoryResponseBodyDataRegionCity) GetId() *string {
	return s.Id
}

func (s *SdnsGetRuleCategoryResponseBodyDataRegionCity) GetLabel() *string {
	return s.Label
}

func (s *SdnsGetRuleCategoryResponseBodyDataRegionCity) SetId(v string) *SdnsGetRuleCategoryResponseBodyDataRegionCity {
	s.Id = &v
	return s
}

func (s *SdnsGetRuleCategoryResponseBodyDataRegionCity) SetLabel(v string) *SdnsGetRuleCategoryResponseBodyDataRegionCity {
	s.Label = &v
	return s
}

func (s *SdnsGetRuleCategoryResponseBodyDataRegionCity) Validate() error {
	return dara.Validate(s)
}

type SdnsGetRuleCategoryResponseBodyDataRegionContinent struct {
	// example:
	//
	// cn
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 亚洲
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
}

func (s SdnsGetRuleCategoryResponseBodyDataRegionContinent) String() string {
	return dara.Prettify(s)
}

func (s SdnsGetRuleCategoryResponseBodyDataRegionContinent) GoString() string {
	return s.String()
}

func (s *SdnsGetRuleCategoryResponseBodyDataRegionContinent) GetId() *string {
	return s.Id
}

func (s *SdnsGetRuleCategoryResponseBodyDataRegionContinent) GetLabel() *string {
	return s.Label
}

func (s *SdnsGetRuleCategoryResponseBodyDataRegionContinent) SetId(v string) *SdnsGetRuleCategoryResponseBodyDataRegionContinent {
	s.Id = &v
	return s
}

func (s *SdnsGetRuleCategoryResponseBodyDataRegionContinent) SetLabel(v string) *SdnsGetRuleCategoryResponseBodyDataRegionContinent {
	s.Label = &v
	return s
}

func (s *SdnsGetRuleCategoryResponseBodyDataRegionContinent) Validate() error {
	return dara.Validate(s)
}

type SdnsGetRuleCategoryResponseBodyDataRegionCountry struct {
	// example:
	//
	// cn
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 中国大陆
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
}

func (s SdnsGetRuleCategoryResponseBodyDataRegionCountry) String() string {
	return dara.Prettify(s)
}

func (s SdnsGetRuleCategoryResponseBodyDataRegionCountry) GoString() string {
	return s.String()
}

func (s *SdnsGetRuleCategoryResponseBodyDataRegionCountry) GetId() *string {
	return s.Id
}

func (s *SdnsGetRuleCategoryResponseBodyDataRegionCountry) GetLabel() *string {
	return s.Label
}

func (s *SdnsGetRuleCategoryResponseBodyDataRegionCountry) SetId(v string) *SdnsGetRuleCategoryResponseBodyDataRegionCountry {
	s.Id = &v
	return s
}

func (s *SdnsGetRuleCategoryResponseBodyDataRegionCountry) SetLabel(v string) *SdnsGetRuleCategoryResponseBodyDataRegionCountry {
	s.Label = &v
	return s
}

func (s *SdnsGetRuleCategoryResponseBodyDataRegionCountry) Validate() error {
	return dara.Validate(s)
}

type SdnsGetRuleCategoryResponseBodyDataRegionIsp struct {
	// example:
	//
	// chinanet
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 电信
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
}

func (s SdnsGetRuleCategoryResponseBodyDataRegionIsp) String() string {
	return dara.Prettify(s)
}

func (s SdnsGetRuleCategoryResponseBodyDataRegionIsp) GoString() string {
	return s.String()
}

func (s *SdnsGetRuleCategoryResponseBodyDataRegionIsp) GetId() *string {
	return s.Id
}

func (s *SdnsGetRuleCategoryResponseBodyDataRegionIsp) GetLabel() *string {
	return s.Label
}

func (s *SdnsGetRuleCategoryResponseBodyDataRegionIsp) SetId(v string) *SdnsGetRuleCategoryResponseBodyDataRegionIsp {
	s.Id = &v
	return s
}

func (s *SdnsGetRuleCategoryResponseBodyDataRegionIsp) SetLabel(v string) *SdnsGetRuleCategoryResponseBodyDataRegionIsp {
	s.Label = &v
	return s
}

func (s *SdnsGetRuleCategoryResponseBodyDataRegionIsp) Validate() error {
	return dara.Validate(s)
}

type SdnsGetRuleCategoryResponseBodyDataRegionProvince struct {
	// example:
	//
	// zhejiang
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 浙江
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
}

func (s SdnsGetRuleCategoryResponseBodyDataRegionProvince) String() string {
	return dara.Prettify(s)
}

func (s SdnsGetRuleCategoryResponseBodyDataRegionProvince) GoString() string {
	return s.String()
}

func (s *SdnsGetRuleCategoryResponseBodyDataRegionProvince) GetId() *string {
	return s.Id
}

func (s *SdnsGetRuleCategoryResponseBodyDataRegionProvince) GetLabel() *string {
	return s.Label
}

func (s *SdnsGetRuleCategoryResponseBodyDataRegionProvince) SetId(v string) *SdnsGetRuleCategoryResponseBodyDataRegionProvince {
	s.Id = &v
	return s
}

func (s *SdnsGetRuleCategoryResponseBodyDataRegionProvince) SetLabel(v string) *SdnsGetRuleCategoryResponseBodyDataRegionProvince {
	s.Label = &v
	return s
}

func (s *SdnsGetRuleCategoryResponseBodyDataRegionProvince) Validate() error {
	return dara.Validate(s)
}

type SdnsGetRuleCategoryResponseBodyDataRules struct {
	ArgKeys []*string          `json:"ArgKeys,omitempty" xml:"ArgKeys,omitempty" type:"Repeated"`
	Args    map[string]*string `json:"Args,omitempty" xml:"Args,omitempty"`
	// example:
	//
	// 默认规则
	Name    *string                                            `json:"Name,omitempty" xml:"Name,omitempty"`
	Records []*SdnsGetRuleCategoryResponseBodyDataRulesRecords `json:"Records,omitempty" xml:"Records,omitempty" type:"Repeated"`
	// example:
	//
	// 0
	ResultMode *int32 `json:"ResultMode,omitempty" xml:"ResultMode,omitempty"`
	// example:
	//
	// 60
	Ttl *int32 `json:"Ttl,omitempty" xml:"Ttl,omitempty"`
}

func (s SdnsGetRuleCategoryResponseBodyDataRules) String() string {
	return dara.Prettify(s)
}

func (s SdnsGetRuleCategoryResponseBodyDataRules) GoString() string {
	return s.String()
}

func (s *SdnsGetRuleCategoryResponseBodyDataRules) GetArgKeys() []*string {
	return s.ArgKeys
}

func (s *SdnsGetRuleCategoryResponseBodyDataRules) GetArgs() map[string]*string {
	return s.Args
}

func (s *SdnsGetRuleCategoryResponseBodyDataRules) GetName() *string {
	return s.Name
}

func (s *SdnsGetRuleCategoryResponseBodyDataRules) GetRecords() []*SdnsGetRuleCategoryResponseBodyDataRulesRecords {
	return s.Records
}

func (s *SdnsGetRuleCategoryResponseBodyDataRules) GetResultMode() *int32 {
	return s.ResultMode
}

func (s *SdnsGetRuleCategoryResponseBodyDataRules) GetTtl() *int32 {
	return s.Ttl
}

func (s *SdnsGetRuleCategoryResponseBodyDataRules) SetArgKeys(v []*string) *SdnsGetRuleCategoryResponseBodyDataRules {
	s.ArgKeys = v
	return s
}

func (s *SdnsGetRuleCategoryResponseBodyDataRules) SetArgs(v map[string]*string) *SdnsGetRuleCategoryResponseBodyDataRules {
	s.Args = v
	return s
}

func (s *SdnsGetRuleCategoryResponseBodyDataRules) SetName(v string) *SdnsGetRuleCategoryResponseBodyDataRules {
	s.Name = &v
	return s
}

func (s *SdnsGetRuleCategoryResponseBodyDataRules) SetRecords(v []*SdnsGetRuleCategoryResponseBodyDataRulesRecords) *SdnsGetRuleCategoryResponseBodyDataRules {
	s.Records = v
	return s
}

func (s *SdnsGetRuleCategoryResponseBodyDataRules) SetResultMode(v int32) *SdnsGetRuleCategoryResponseBodyDataRules {
	s.ResultMode = &v
	return s
}

func (s *SdnsGetRuleCategoryResponseBodyDataRules) SetTtl(v int32) *SdnsGetRuleCategoryResponseBodyDataRules {
	s.Ttl = &v
	return s
}

func (s *SdnsGetRuleCategoryResponseBodyDataRules) Validate() error {
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

type SdnsGetRuleCategoryResponseBodyDataRulesRecords struct {
	// example:
	//
	// 192.0.2.200
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// example:
	//
	// 1
	Type *int32 `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// A
	TypeLabel *string `json:"TypeLabel,omitempty" xml:"TypeLabel,omitempty"`
	// example:
	//
	// 100
	Weight *int32 `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s SdnsGetRuleCategoryResponseBodyDataRulesRecords) String() string {
	return dara.Prettify(s)
}

func (s SdnsGetRuleCategoryResponseBodyDataRulesRecords) GoString() string {
	return s.String()
}

func (s *SdnsGetRuleCategoryResponseBodyDataRulesRecords) GetIp() *string {
	return s.Ip
}

func (s *SdnsGetRuleCategoryResponseBodyDataRulesRecords) GetType() *int32 {
	return s.Type
}

func (s *SdnsGetRuleCategoryResponseBodyDataRulesRecords) GetTypeLabel() *string {
	return s.TypeLabel
}

func (s *SdnsGetRuleCategoryResponseBodyDataRulesRecords) GetWeight() *int32 {
	return s.Weight
}

func (s *SdnsGetRuleCategoryResponseBodyDataRulesRecords) SetIp(v string) *SdnsGetRuleCategoryResponseBodyDataRulesRecords {
	s.Ip = &v
	return s
}

func (s *SdnsGetRuleCategoryResponseBodyDataRulesRecords) SetType(v int32) *SdnsGetRuleCategoryResponseBodyDataRulesRecords {
	s.Type = &v
	return s
}

func (s *SdnsGetRuleCategoryResponseBodyDataRulesRecords) SetTypeLabel(v string) *SdnsGetRuleCategoryResponseBodyDataRulesRecords {
	s.TypeLabel = &v
	return s
}

func (s *SdnsGetRuleCategoryResponseBodyDataRulesRecords) SetWeight(v int32) *SdnsGetRuleCategoryResponseBodyDataRulesRecords {
	s.Weight = &v
	return s
}

func (s *SdnsGetRuleCategoryResponseBodyDataRulesRecords) Validate() error {
	return dara.Validate(s)
}
