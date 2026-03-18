// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchCloudGtmMonitorTemplatesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *SearchCloudGtmMonitorTemplatesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *SearchCloudGtmMonitorTemplatesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *SearchCloudGtmMonitorTemplatesResponseBody
	GetRequestId() *string
	SetTemplates(v *SearchCloudGtmMonitorTemplatesResponseBodyTemplates) *SearchCloudGtmMonitorTemplatesResponseBody
	GetTemplates() *SearchCloudGtmMonitorTemplatesResponseBodyTemplates
	SetTotalItems(v int32) *SearchCloudGtmMonitorTemplatesResponseBody
	GetTotalItems() *int32
	SetTotalPages(v int32) *SearchCloudGtmMonitorTemplatesResponseBody
	GetTotalPages() *int32
}

type SearchCloudGtmMonitorTemplatesResponseBody struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 536E9CAD-DB30-4647-AC87-AA5CC38C5382
	RequestId *string                                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Templates *SearchCloudGtmMonitorTemplatesResponseBodyTemplates `json:"Templates,omitempty" xml:"Templates,omitempty" type:"Struct"`
	// example:
	//
	// 1
	TotalItems *int32 `json:"TotalItems,omitempty" xml:"TotalItems,omitempty"`
	// example:
	//
	// 1
	TotalPages *int32 `json:"TotalPages,omitempty" xml:"TotalPages,omitempty"`
}

func (s SearchCloudGtmMonitorTemplatesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SearchCloudGtmMonitorTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *SearchCloudGtmMonitorTemplatesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *SearchCloudGtmMonitorTemplatesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *SearchCloudGtmMonitorTemplatesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SearchCloudGtmMonitorTemplatesResponseBody) GetTemplates() *SearchCloudGtmMonitorTemplatesResponseBodyTemplates {
	return s.Templates
}

func (s *SearchCloudGtmMonitorTemplatesResponseBody) GetTotalItems() *int32 {
	return s.TotalItems
}

func (s *SearchCloudGtmMonitorTemplatesResponseBody) GetTotalPages() *int32 {
	return s.TotalPages
}

func (s *SearchCloudGtmMonitorTemplatesResponseBody) SetPageNumber(v int32) *SearchCloudGtmMonitorTemplatesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *SearchCloudGtmMonitorTemplatesResponseBody) SetPageSize(v int32) *SearchCloudGtmMonitorTemplatesResponseBody {
	s.PageSize = &v
	return s
}

func (s *SearchCloudGtmMonitorTemplatesResponseBody) SetRequestId(v string) *SearchCloudGtmMonitorTemplatesResponseBody {
	s.RequestId = &v
	return s
}

func (s *SearchCloudGtmMonitorTemplatesResponseBody) SetTemplates(v *SearchCloudGtmMonitorTemplatesResponseBodyTemplates) *SearchCloudGtmMonitorTemplatesResponseBody {
	s.Templates = v
	return s
}

func (s *SearchCloudGtmMonitorTemplatesResponseBody) SetTotalItems(v int32) *SearchCloudGtmMonitorTemplatesResponseBody {
	s.TotalItems = &v
	return s
}

func (s *SearchCloudGtmMonitorTemplatesResponseBody) SetTotalPages(v int32) *SearchCloudGtmMonitorTemplatesResponseBody {
	s.TotalPages = &v
	return s
}

func (s *SearchCloudGtmMonitorTemplatesResponseBody) Validate() error {
	if s.Templates != nil {
		if err := s.Templates.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SearchCloudGtmMonitorTemplatesResponseBodyTemplates struct {
	Template []*SearchCloudGtmMonitorTemplatesResponseBodyTemplatesTemplate `json:"Template,omitempty" xml:"Template,omitempty" type:"Repeated"`
}

func (s SearchCloudGtmMonitorTemplatesResponseBodyTemplates) String() string {
	return dara.Prettify(s)
}

func (s SearchCloudGtmMonitorTemplatesResponseBodyTemplates) GoString() string {
	return s.String()
}

func (s *SearchCloudGtmMonitorTemplatesResponseBodyTemplates) GetTemplate() []*SearchCloudGtmMonitorTemplatesResponseBodyTemplatesTemplate {
	return s.Template
}

func (s *SearchCloudGtmMonitorTemplatesResponseBodyTemplates) SetTemplate(v []*SearchCloudGtmMonitorTemplatesResponseBodyTemplatesTemplate) *SearchCloudGtmMonitorTemplatesResponseBodyTemplates {
	s.Template = v
	return s
}

func (s *SearchCloudGtmMonitorTemplatesResponseBodyTemplates) Validate() error {
	if s.Template != nil {
		for _, item := range s.Template {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SearchCloudGtmMonitorTemplatesResponseBodyTemplatesTemplate struct {
	CreateTime      *string                                                                  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CreateTimestamp *int64                                                                   `json:"CreateTimestamp,omitempty" xml:"CreateTimestamp,omitempty"`
	EvaluationCount *int32                                                                   `json:"EvaluationCount,omitempty" xml:"EvaluationCount,omitempty"`
	ExtendInfo      *string                                                                  `json:"ExtendInfo,omitempty" xml:"ExtendInfo,omitempty"`
	FailureRate     *int32                                                                   `json:"FailureRate,omitempty" xml:"FailureRate,omitempty"`
	Interval        *int32                                                                   `json:"Interval,omitempty" xml:"Interval,omitempty"`
	IpVersion       *string                                                                  `json:"IpVersion,omitempty" xml:"IpVersion,omitempty"`
	IspCityNodes    *SearchCloudGtmMonitorTemplatesResponseBodyTemplatesTemplateIspCityNodes `json:"IspCityNodes,omitempty" xml:"IspCityNodes,omitempty" type:"Struct"`
	Name            *string                                                                  `json:"Name,omitempty" xml:"Name,omitempty"`
	Protocol        *string                                                                  `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	Remark          *string                                                                  `json:"Remark,omitempty" xml:"Remark,omitempty"`
	TemplateId      *string                                                                  `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	Timeout         *int32                                                                   `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
	UpdateTime      *string                                                                  `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	UpdateTimestamp *int64                                                                   `json:"UpdateTimestamp,omitempty" xml:"UpdateTimestamp,omitempty"`
}

func (s SearchCloudGtmMonitorTemplatesResponseBodyTemplatesTemplate) String() string {
	return dara.Prettify(s)
}

func (s SearchCloudGtmMonitorTemplatesResponseBodyTemplatesTemplate) GoString() string {
	return s.String()
}

func (s *SearchCloudGtmMonitorTemplatesResponseBodyTemplatesTemplate) GetCreateTime() *string {
	return s.CreateTime
}

func (s *SearchCloudGtmMonitorTemplatesResponseBodyTemplatesTemplate) GetCreateTimestamp() *int64 {
	return s.CreateTimestamp
}

func (s *SearchCloudGtmMonitorTemplatesResponseBodyTemplatesTemplate) GetEvaluationCount() *int32 {
	return s.EvaluationCount
}

func (s *SearchCloudGtmMonitorTemplatesResponseBodyTemplatesTemplate) GetExtendInfo() *string {
	return s.ExtendInfo
}

func (s *SearchCloudGtmMonitorTemplatesResponseBodyTemplatesTemplate) GetFailureRate() *int32 {
	return s.FailureRate
}

func (s *SearchCloudGtmMonitorTemplatesResponseBodyTemplatesTemplate) GetInterval() *int32 {
	return s.Interval
}

func (s *SearchCloudGtmMonitorTemplatesResponseBodyTemplatesTemplate) GetIpVersion() *string {
	return s.IpVersion
}

func (s *SearchCloudGtmMonitorTemplatesResponseBodyTemplatesTemplate) GetIspCityNodes() *SearchCloudGtmMonitorTemplatesResponseBodyTemplatesTemplateIspCityNodes {
	return s.IspCityNodes
}

func (s *SearchCloudGtmMonitorTemplatesResponseBodyTemplatesTemplate) GetName() *string {
	return s.Name
}

func (s *SearchCloudGtmMonitorTemplatesResponseBodyTemplatesTemplate) GetProtocol() *string {
	return s.Protocol
}

func (s *SearchCloudGtmMonitorTemplatesResponseBodyTemplatesTemplate) GetRemark() *string {
	return s.Remark
}

func (s *SearchCloudGtmMonitorTemplatesResponseBodyTemplatesTemplate) GetTemplateId() *string {
	return s.TemplateId
}

func (s *SearchCloudGtmMonitorTemplatesResponseBodyTemplatesTemplate) GetTimeout() *int32 {
	return s.Timeout
}

func (s *SearchCloudGtmMonitorTemplatesResponseBodyTemplatesTemplate) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *SearchCloudGtmMonitorTemplatesResponseBodyTemplatesTemplate) GetUpdateTimestamp() *int64 {
	return s.UpdateTimestamp
}

func (s *SearchCloudGtmMonitorTemplatesResponseBodyTemplatesTemplate) SetCreateTime(v string) *SearchCloudGtmMonitorTemplatesResponseBodyTemplatesTemplate {
	s.CreateTime = &v
	return s
}

func (s *SearchCloudGtmMonitorTemplatesResponseBodyTemplatesTemplate) SetCreateTimestamp(v int64) *SearchCloudGtmMonitorTemplatesResponseBodyTemplatesTemplate {
	s.CreateTimestamp = &v
	return s
}

func (s *SearchCloudGtmMonitorTemplatesResponseBodyTemplatesTemplate) SetEvaluationCount(v int32) *SearchCloudGtmMonitorTemplatesResponseBodyTemplatesTemplate {
	s.EvaluationCount = &v
	return s
}

func (s *SearchCloudGtmMonitorTemplatesResponseBodyTemplatesTemplate) SetExtendInfo(v string) *SearchCloudGtmMonitorTemplatesResponseBodyTemplatesTemplate {
	s.ExtendInfo = &v
	return s
}

func (s *SearchCloudGtmMonitorTemplatesResponseBodyTemplatesTemplate) SetFailureRate(v int32) *SearchCloudGtmMonitorTemplatesResponseBodyTemplatesTemplate {
	s.FailureRate = &v
	return s
}

func (s *SearchCloudGtmMonitorTemplatesResponseBodyTemplatesTemplate) SetInterval(v int32) *SearchCloudGtmMonitorTemplatesResponseBodyTemplatesTemplate {
	s.Interval = &v
	return s
}

func (s *SearchCloudGtmMonitorTemplatesResponseBodyTemplatesTemplate) SetIpVersion(v string) *SearchCloudGtmMonitorTemplatesResponseBodyTemplatesTemplate {
	s.IpVersion = &v
	return s
}

func (s *SearchCloudGtmMonitorTemplatesResponseBodyTemplatesTemplate) SetIspCityNodes(v *SearchCloudGtmMonitorTemplatesResponseBodyTemplatesTemplateIspCityNodes) *SearchCloudGtmMonitorTemplatesResponseBodyTemplatesTemplate {
	s.IspCityNodes = v
	return s
}

func (s *SearchCloudGtmMonitorTemplatesResponseBodyTemplatesTemplate) SetName(v string) *SearchCloudGtmMonitorTemplatesResponseBodyTemplatesTemplate {
	s.Name = &v
	return s
}

func (s *SearchCloudGtmMonitorTemplatesResponseBodyTemplatesTemplate) SetProtocol(v string) *SearchCloudGtmMonitorTemplatesResponseBodyTemplatesTemplate {
	s.Protocol = &v
	return s
}

func (s *SearchCloudGtmMonitorTemplatesResponseBodyTemplatesTemplate) SetRemark(v string) *SearchCloudGtmMonitorTemplatesResponseBodyTemplatesTemplate {
	s.Remark = &v
	return s
}

func (s *SearchCloudGtmMonitorTemplatesResponseBodyTemplatesTemplate) SetTemplateId(v string) *SearchCloudGtmMonitorTemplatesResponseBodyTemplatesTemplate {
	s.TemplateId = &v
	return s
}

func (s *SearchCloudGtmMonitorTemplatesResponseBodyTemplatesTemplate) SetTimeout(v int32) *SearchCloudGtmMonitorTemplatesResponseBodyTemplatesTemplate {
	s.Timeout = &v
	return s
}

func (s *SearchCloudGtmMonitorTemplatesResponseBodyTemplatesTemplate) SetUpdateTime(v string) *SearchCloudGtmMonitorTemplatesResponseBodyTemplatesTemplate {
	s.UpdateTime = &v
	return s
}

func (s *SearchCloudGtmMonitorTemplatesResponseBodyTemplatesTemplate) SetUpdateTimestamp(v int64) *SearchCloudGtmMonitorTemplatesResponseBodyTemplatesTemplate {
	s.UpdateTimestamp = &v
	return s
}

func (s *SearchCloudGtmMonitorTemplatesResponseBodyTemplatesTemplate) Validate() error {
	if s.IspCityNodes != nil {
		if err := s.IspCityNodes.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SearchCloudGtmMonitorTemplatesResponseBodyTemplatesTemplateIspCityNodes struct {
	IspCityNode []*SearchCloudGtmMonitorTemplatesResponseBodyTemplatesTemplateIspCityNodesIspCityNode `json:"IspCityNode,omitempty" xml:"IspCityNode,omitempty" type:"Repeated"`
}

func (s SearchCloudGtmMonitorTemplatesResponseBodyTemplatesTemplateIspCityNodes) String() string {
	return dara.Prettify(s)
}

func (s SearchCloudGtmMonitorTemplatesResponseBodyTemplatesTemplateIspCityNodes) GoString() string {
	return s.String()
}

func (s *SearchCloudGtmMonitorTemplatesResponseBodyTemplatesTemplateIspCityNodes) GetIspCityNode() []*SearchCloudGtmMonitorTemplatesResponseBodyTemplatesTemplateIspCityNodesIspCityNode {
	return s.IspCityNode
}

func (s *SearchCloudGtmMonitorTemplatesResponseBodyTemplatesTemplateIspCityNodes) SetIspCityNode(v []*SearchCloudGtmMonitorTemplatesResponseBodyTemplatesTemplateIspCityNodesIspCityNode) *SearchCloudGtmMonitorTemplatesResponseBodyTemplatesTemplateIspCityNodes {
	s.IspCityNode = v
	return s
}

func (s *SearchCloudGtmMonitorTemplatesResponseBodyTemplatesTemplateIspCityNodes) Validate() error {
	if s.IspCityNode != nil {
		for _, item := range s.IspCityNode {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SearchCloudGtmMonitorTemplatesResponseBodyTemplatesTemplateIspCityNodesIspCityNode struct {
	CityCode    *string `json:"CityCode,omitempty" xml:"CityCode,omitempty"`
	CityName    *string `json:"CityName,omitempty" xml:"CityName,omitempty"`
	CountryCode *string `json:"CountryCode,omitempty" xml:"CountryCode,omitempty"`
	CountryName *string `json:"CountryName,omitempty" xml:"CountryName,omitempty"`
	GroupName   *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	GroupType   *string `json:"GroupType,omitempty" xml:"GroupType,omitempty"`
	IspCode     *string `json:"IspCode,omitempty" xml:"IspCode,omitempty"`
	IspName     *string `json:"IspName,omitempty" xml:"IspName,omitempty"`
}

func (s SearchCloudGtmMonitorTemplatesResponseBodyTemplatesTemplateIspCityNodesIspCityNode) String() string {
	return dara.Prettify(s)
}

func (s SearchCloudGtmMonitorTemplatesResponseBodyTemplatesTemplateIspCityNodesIspCityNode) GoString() string {
	return s.String()
}

func (s *SearchCloudGtmMonitorTemplatesResponseBodyTemplatesTemplateIspCityNodesIspCityNode) GetCityCode() *string {
	return s.CityCode
}

func (s *SearchCloudGtmMonitorTemplatesResponseBodyTemplatesTemplateIspCityNodesIspCityNode) GetCityName() *string {
	return s.CityName
}

func (s *SearchCloudGtmMonitorTemplatesResponseBodyTemplatesTemplateIspCityNodesIspCityNode) GetCountryCode() *string {
	return s.CountryCode
}

func (s *SearchCloudGtmMonitorTemplatesResponseBodyTemplatesTemplateIspCityNodesIspCityNode) GetCountryName() *string {
	return s.CountryName
}

func (s *SearchCloudGtmMonitorTemplatesResponseBodyTemplatesTemplateIspCityNodesIspCityNode) GetGroupName() *string {
	return s.GroupName
}

func (s *SearchCloudGtmMonitorTemplatesResponseBodyTemplatesTemplateIspCityNodesIspCityNode) GetGroupType() *string {
	return s.GroupType
}

func (s *SearchCloudGtmMonitorTemplatesResponseBodyTemplatesTemplateIspCityNodesIspCityNode) GetIspCode() *string {
	return s.IspCode
}

func (s *SearchCloudGtmMonitorTemplatesResponseBodyTemplatesTemplateIspCityNodesIspCityNode) GetIspName() *string {
	return s.IspName
}

func (s *SearchCloudGtmMonitorTemplatesResponseBodyTemplatesTemplateIspCityNodesIspCityNode) SetCityCode(v string) *SearchCloudGtmMonitorTemplatesResponseBodyTemplatesTemplateIspCityNodesIspCityNode {
	s.CityCode = &v
	return s
}

func (s *SearchCloudGtmMonitorTemplatesResponseBodyTemplatesTemplateIspCityNodesIspCityNode) SetCityName(v string) *SearchCloudGtmMonitorTemplatesResponseBodyTemplatesTemplateIspCityNodesIspCityNode {
	s.CityName = &v
	return s
}

func (s *SearchCloudGtmMonitorTemplatesResponseBodyTemplatesTemplateIspCityNodesIspCityNode) SetCountryCode(v string) *SearchCloudGtmMonitorTemplatesResponseBodyTemplatesTemplateIspCityNodesIspCityNode {
	s.CountryCode = &v
	return s
}

func (s *SearchCloudGtmMonitorTemplatesResponseBodyTemplatesTemplateIspCityNodesIspCityNode) SetCountryName(v string) *SearchCloudGtmMonitorTemplatesResponseBodyTemplatesTemplateIspCityNodesIspCityNode {
	s.CountryName = &v
	return s
}

func (s *SearchCloudGtmMonitorTemplatesResponseBodyTemplatesTemplateIspCityNodesIspCityNode) SetGroupName(v string) *SearchCloudGtmMonitorTemplatesResponseBodyTemplatesTemplateIspCityNodesIspCityNode {
	s.GroupName = &v
	return s
}

func (s *SearchCloudGtmMonitorTemplatesResponseBodyTemplatesTemplateIspCityNodesIspCityNode) SetGroupType(v string) *SearchCloudGtmMonitorTemplatesResponseBodyTemplatesTemplateIspCityNodesIspCityNode {
	s.GroupType = &v
	return s
}

func (s *SearchCloudGtmMonitorTemplatesResponseBodyTemplatesTemplateIspCityNodesIspCityNode) SetIspCode(v string) *SearchCloudGtmMonitorTemplatesResponseBodyTemplatesTemplateIspCityNodesIspCityNode {
	s.IspCode = &v
	return s
}

func (s *SearchCloudGtmMonitorTemplatesResponseBodyTemplatesTemplateIspCityNodesIspCityNode) SetIspName(v string) *SearchCloudGtmMonitorTemplatesResponseBodyTemplatesTemplateIspCityNodesIspCityNode {
	s.IspName = &v
	return s
}

func (s *SearchCloudGtmMonitorTemplatesResponseBodyTemplatesTemplateIspCityNodesIspCityNode) Validate() error {
	return dara.Validate(s)
}
