// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTransportLayerApplicationsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApplications(v []*ListTransportLayerApplicationsResponseBodyApplications) *ListTransportLayerApplicationsResponseBody
	GetApplications() []*ListTransportLayerApplicationsResponseBodyApplications
	SetPageNumber(v int32) *ListTransportLayerApplicationsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListTransportLayerApplicationsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListTransportLayerApplicationsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListTransportLayerApplicationsResponseBody
	GetTotalCount() *int32
}

type ListTransportLayerApplicationsResponseBody struct {
	Applications []*ListTransportLayerApplicationsResponseBodyApplications `json:"Applications,omitempty" xml:"Applications,omitempty" type:"Repeated"`
	PageNumber   *int32                                                    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize     *int32                                                    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId    *string                                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount   *int32                                                    `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListTransportLayerApplicationsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTransportLayerApplicationsResponseBody) GoString() string {
	return s.String()
}

func (s *ListTransportLayerApplicationsResponseBody) GetApplications() []*ListTransportLayerApplicationsResponseBodyApplications {
	return s.Applications
}

func (s *ListTransportLayerApplicationsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListTransportLayerApplicationsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListTransportLayerApplicationsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTransportLayerApplicationsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListTransportLayerApplicationsResponseBody) SetApplications(v []*ListTransportLayerApplicationsResponseBodyApplications) *ListTransportLayerApplicationsResponseBody {
	s.Applications = v
	return s
}

func (s *ListTransportLayerApplicationsResponseBody) SetPageNumber(v int32) *ListTransportLayerApplicationsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListTransportLayerApplicationsResponseBody) SetPageSize(v int32) *ListTransportLayerApplicationsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListTransportLayerApplicationsResponseBody) SetRequestId(v string) *ListTransportLayerApplicationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTransportLayerApplicationsResponseBody) SetTotalCount(v int32) *ListTransportLayerApplicationsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListTransportLayerApplicationsResponseBody) Validate() error {
	if s.Applications != nil {
		for _, item := range s.Applications {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListTransportLayerApplicationsResponseBodyApplications struct {
	ApplicationId           *int64                                                         `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	Cname                   *string                                                        `json:"Cname,omitempty" xml:"Cname,omitempty"`
	CrossBorderOptimization *string                                                        `json:"CrossBorderOptimization,omitempty" xml:"CrossBorderOptimization,omitempty"`
	IpAccessRule            *string                                                        `json:"IpAccessRule,omitempty" xml:"IpAccessRule,omitempty"`
	Ipv6                    *string                                                        `json:"Ipv6,omitempty" xml:"Ipv6,omitempty"`
	KeepAliveProtection     *string                                                        `json:"KeepAliveProtection,omitempty" xml:"KeepAliveProtection,omitempty"`
	RecordName              *string                                                        `json:"RecordName,omitempty" xml:"RecordName,omitempty"`
	Rules                   []*ListTransportLayerApplicationsResponseBodyApplicationsRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
	RulesCount              *int32                                                         `json:"RulesCount,omitempty" xml:"RulesCount,omitempty"`
	SiteId                  *int64                                                         `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	StaticIp                *string                                                        `json:"StaticIp,omitempty" xml:"StaticIp,omitempty"`
	// This parameter is required.
	StaticIpV4List []*ListTransportLayerApplicationsResponseBodyApplicationsStaticIpV4List `json:"StaticIpV4List,omitempty" xml:"StaticIpV4List,omitempty" type:"Repeated"`
	Status         *string                                                                 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListTransportLayerApplicationsResponseBodyApplications) String() string {
	return dara.Prettify(s)
}

func (s ListTransportLayerApplicationsResponseBodyApplications) GoString() string {
	return s.String()
}

func (s *ListTransportLayerApplicationsResponseBodyApplications) GetApplicationId() *int64 {
	return s.ApplicationId
}

func (s *ListTransportLayerApplicationsResponseBodyApplications) GetCname() *string {
	return s.Cname
}

func (s *ListTransportLayerApplicationsResponseBodyApplications) GetCrossBorderOptimization() *string {
	return s.CrossBorderOptimization
}

func (s *ListTransportLayerApplicationsResponseBodyApplications) GetIpAccessRule() *string {
	return s.IpAccessRule
}

func (s *ListTransportLayerApplicationsResponseBodyApplications) GetIpv6() *string {
	return s.Ipv6
}

func (s *ListTransportLayerApplicationsResponseBodyApplications) GetKeepAliveProtection() *string {
	return s.KeepAliveProtection
}

func (s *ListTransportLayerApplicationsResponseBodyApplications) GetRecordName() *string {
	return s.RecordName
}

func (s *ListTransportLayerApplicationsResponseBodyApplications) GetRules() []*ListTransportLayerApplicationsResponseBodyApplicationsRules {
	return s.Rules
}

func (s *ListTransportLayerApplicationsResponseBodyApplications) GetRulesCount() *int32 {
	return s.RulesCount
}

func (s *ListTransportLayerApplicationsResponseBodyApplications) GetSiteId() *int64 {
	return s.SiteId
}

func (s *ListTransportLayerApplicationsResponseBodyApplications) GetStaticIp() *string {
	return s.StaticIp
}

func (s *ListTransportLayerApplicationsResponseBodyApplications) GetStaticIpV4List() []*ListTransportLayerApplicationsResponseBodyApplicationsStaticIpV4List {
	return s.StaticIpV4List
}

func (s *ListTransportLayerApplicationsResponseBodyApplications) GetStatus() *string {
	return s.Status
}

func (s *ListTransportLayerApplicationsResponseBodyApplications) SetApplicationId(v int64) *ListTransportLayerApplicationsResponseBodyApplications {
	s.ApplicationId = &v
	return s
}

func (s *ListTransportLayerApplicationsResponseBodyApplications) SetCname(v string) *ListTransportLayerApplicationsResponseBodyApplications {
	s.Cname = &v
	return s
}

func (s *ListTransportLayerApplicationsResponseBodyApplications) SetCrossBorderOptimization(v string) *ListTransportLayerApplicationsResponseBodyApplications {
	s.CrossBorderOptimization = &v
	return s
}

func (s *ListTransportLayerApplicationsResponseBodyApplications) SetIpAccessRule(v string) *ListTransportLayerApplicationsResponseBodyApplications {
	s.IpAccessRule = &v
	return s
}

func (s *ListTransportLayerApplicationsResponseBodyApplications) SetIpv6(v string) *ListTransportLayerApplicationsResponseBodyApplications {
	s.Ipv6 = &v
	return s
}

func (s *ListTransportLayerApplicationsResponseBodyApplications) SetKeepAliveProtection(v string) *ListTransportLayerApplicationsResponseBodyApplications {
	s.KeepAliveProtection = &v
	return s
}

func (s *ListTransportLayerApplicationsResponseBodyApplications) SetRecordName(v string) *ListTransportLayerApplicationsResponseBodyApplications {
	s.RecordName = &v
	return s
}

func (s *ListTransportLayerApplicationsResponseBodyApplications) SetRules(v []*ListTransportLayerApplicationsResponseBodyApplicationsRules) *ListTransportLayerApplicationsResponseBodyApplications {
	s.Rules = v
	return s
}

func (s *ListTransportLayerApplicationsResponseBodyApplications) SetRulesCount(v int32) *ListTransportLayerApplicationsResponseBodyApplications {
	s.RulesCount = &v
	return s
}

func (s *ListTransportLayerApplicationsResponseBodyApplications) SetSiteId(v int64) *ListTransportLayerApplicationsResponseBodyApplications {
	s.SiteId = &v
	return s
}

func (s *ListTransportLayerApplicationsResponseBodyApplications) SetStaticIp(v string) *ListTransportLayerApplicationsResponseBodyApplications {
	s.StaticIp = &v
	return s
}

func (s *ListTransportLayerApplicationsResponseBodyApplications) SetStaticIpV4List(v []*ListTransportLayerApplicationsResponseBodyApplicationsStaticIpV4List) *ListTransportLayerApplicationsResponseBodyApplications {
	s.StaticIpV4List = v
	return s
}

func (s *ListTransportLayerApplicationsResponseBodyApplications) SetStatus(v string) *ListTransportLayerApplicationsResponseBodyApplications {
	s.Status = &v
	return s
}

func (s *ListTransportLayerApplicationsResponseBodyApplications) Validate() error {
	if s.Rules != nil {
		for _, item := range s.Rules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.StaticIpV4List != nil {
		for _, item := range s.StaticIpV4List {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListTransportLayerApplicationsResponseBodyApplicationsRules struct {
	ClientIPPassThroughMode *string `json:"ClientIPPassThroughMode,omitempty" xml:"ClientIPPassThroughMode,omitempty"`
	Comment                 *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	EdgePort                *string `json:"EdgePort,omitempty" xml:"EdgePort,omitempty"`
	Protocol                *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	RuleId                  *int64  `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	Source                  *string `json:"Source,omitempty" xml:"Source,omitempty"`
	SourcePort              *string `json:"SourcePort,omitempty" xml:"SourcePort,omitempty"`
	SourceType              *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
}

func (s ListTransportLayerApplicationsResponseBodyApplicationsRules) String() string {
	return dara.Prettify(s)
}

func (s ListTransportLayerApplicationsResponseBodyApplicationsRules) GoString() string {
	return s.String()
}

func (s *ListTransportLayerApplicationsResponseBodyApplicationsRules) GetClientIPPassThroughMode() *string {
	return s.ClientIPPassThroughMode
}

func (s *ListTransportLayerApplicationsResponseBodyApplicationsRules) GetComment() *string {
	return s.Comment
}

func (s *ListTransportLayerApplicationsResponseBodyApplicationsRules) GetEdgePort() *string {
	return s.EdgePort
}

func (s *ListTransportLayerApplicationsResponseBodyApplicationsRules) GetProtocol() *string {
	return s.Protocol
}

func (s *ListTransportLayerApplicationsResponseBodyApplicationsRules) GetRuleId() *int64 {
	return s.RuleId
}

func (s *ListTransportLayerApplicationsResponseBodyApplicationsRules) GetSource() *string {
	return s.Source
}

func (s *ListTransportLayerApplicationsResponseBodyApplicationsRules) GetSourcePort() *string {
	return s.SourcePort
}

func (s *ListTransportLayerApplicationsResponseBodyApplicationsRules) GetSourceType() *string {
	return s.SourceType
}

func (s *ListTransportLayerApplicationsResponseBodyApplicationsRules) SetClientIPPassThroughMode(v string) *ListTransportLayerApplicationsResponseBodyApplicationsRules {
	s.ClientIPPassThroughMode = &v
	return s
}

func (s *ListTransportLayerApplicationsResponseBodyApplicationsRules) SetComment(v string) *ListTransportLayerApplicationsResponseBodyApplicationsRules {
	s.Comment = &v
	return s
}

func (s *ListTransportLayerApplicationsResponseBodyApplicationsRules) SetEdgePort(v string) *ListTransportLayerApplicationsResponseBodyApplicationsRules {
	s.EdgePort = &v
	return s
}

func (s *ListTransportLayerApplicationsResponseBodyApplicationsRules) SetProtocol(v string) *ListTransportLayerApplicationsResponseBodyApplicationsRules {
	s.Protocol = &v
	return s
}

func (s *ListTransportLayerApplicationsResponseBodyApplicationsRules) SetRuleId(v int64) *ListTransportLayerApplicationsResponseBodyApplicationsRules {
	s.RuleId = &v
	return s
}

func (s *ListTransportLayerApplicationsResponseBodyApplicationsRules) SetSource(v string) *ListTransportLayerApplicationsResponseBodyApplicationsRules {
	s.Source = &v
	return s
}

func (s *ListTransportLayerApplicationsResponseBodyApplicationsRules) SetSourcePort(v string) *ListTransportLayerApplicationsResponseBodyApplicationsRules {
	s.SourcePort = &v
	return s
}

func (s *ListTransportLayerApplicationsResponseBodyApplicationsRules) SetSourceType(v string) *ListTransportLayerApplicationsResponseBodyApplicationsRules {
	s.SourceType = &v
	return s
}

func (s *ListTransportLayerApplicationsResponseBodyApplicationsRules) Validate() error {
	return dara.Validate(s)
}

type ListTransportLayerApplicationsResponseBodyApplicationsStaticIpV4List struct {
	Address *string `json:"Address,omitempty" xml:"Address,omitempty"`
	Status  *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListTransportLayerApplicationsResponseBodyApplicationsStaticIpV4List) String() string {
	return dara.Prettify(s)
}

func (s ListTransportLayerApplicationsResponseBodyApplicationsStaticIpV4List) GoString() string {
	return s.String()
}

func (s *ListTransportLayerApplicationsResponseBodyApplicationsStaticIpV4List) GetAddress() *string {
	return s.Address
}

func (s *ListTransportLayerApplicationsResponseBodyApplicationsStaticIpV4List) GetStatus() *string {
	return s.Status
}

func (s *ListTransportLayerApplicationsResponseBodyApplicationsStaticIpV4List) SetAddress(v string) *ListTransportLayerApplicationsResponseBodyApplicationsStaticIpV4List {
	s.Address = &v
	return s
}

func (s *ListTransportLayerApplicationsResponseBodyApplicationsStaticIpV4List) SetStatus(v string) *ListTransportLayerApplicationsResponseBodyApplicationsStaticIpV4List {
	s.Status = &v
	return s
}

func (s *ListTransportLayerApplicationsResponseBodyApplicationsStaticIpV4List) Validate() error {
	return dara.Validate(s)
}
