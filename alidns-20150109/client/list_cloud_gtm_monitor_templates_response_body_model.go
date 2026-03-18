// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCloudGtmMonitorTemplatesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListCloudGtmMonitorTemplatesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListCloudGtmMonitorTemplatesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListCloudGtmMonitorTemplatesResponseBody
	GetRequestId() *string
	SetTemplates(v *ListCloudGtmMonitorTemplatesResponseBodyTemplates) *ListCloudGtmMonitorTemplatesResponseBody
	GetTemplates() *ListCloudGtmMonitorTemplatesResponseBodyTemplates
	SetTotalItems(v int32) *ListCloudGtmMonitorTemplatesResponseBody
	GetTotalItems() *int32
	SetTotalPages(v int32) *ListCloudGtmMonitorTemplatesResponseBody
	GetTotalPages() *int32
}

type ListCloudGtmMonitorTemplatesResponseBody struct {
	// Current page number, starting from 1, default is 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of rows per page when paginating queries, with a maximum value of 100 and a default of 20.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Unique request identification code.
	//
	// example:
	//
	// 75446CC1-FC9A-4595-8D96-089D73D7A63D
	RequestId *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Templates *ListCloudGtmMonitorTemplatesResponseBodyTemplates `json:"Templates,omitempty" xml:"Templates,omitempty" type:"Struct"`
	// Total number of health check template entries retrieved.
	//
	// example:
	//
	// 30
	TotalItems *int32 `json:"TotalItems,omitempty" xml:"TotalItems,omitempty"`
	// Total number of pages after data pagination.
	//
	// example:
	//
	// 2
	TotalPages *int32 `json:"TotalPages,omitempty" xml:"TotalPages,omitempty"`
}

func (s ListCloudGtmMonitorTemplatesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListCloudGtmMonitorTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *ListCloudGtmMonitorTemplatesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListCloudGtmMonitorTemplatesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListCloudGtmMonitorTemplatesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListCloudGtmMonitorTemplatesResponseBody) GetTemplates() *ListCloudGtmMonitorTemplatesResponseBodyTemplates {
	return s.Templates
}

func (s *ListCloudGtmMonitorTemplatesResponseBody) GetTotalItems() *int32 {
	return s.TotalItems
}

func (s *ListCloudGtmMonitorTemplatesResponseBody) GetTotalPages() *int32 {
	return s.TotalPages
}

func (s *ListCloudGtmMonitorTemplatesResponseBody) SetPageNumber(v int32) *ListCloudGtmMonitorTemplatesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListCloudGtmMonitorTemplatesResponseBody) SetPageSize(v int32) *ListCloudGtmMonitorTemplatesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListCloudGtmMonitorTemplatesResponseBody) SetRequestId(v string) *ListCloudGtmMonitorTemplatesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCloudGtmMonitorTemplatesResponseBody) SetTemplates(v *ListCloudGtmMonitorTemplatesResponseBodyTemplates) *ListCloudGtmMonitorTemplatesResponseBody {
	s.Templates = v
	return s
}

func (s *ListCloudGtmMonitorTemplatesResponseBody) SetTotalItems(v int32) *ListCloudGtmMonitorTemplatesResponseBody {
	s.TotalItems = &v
	return s
}

func (s *ListCloudGtmMonitorTemplatesResponseBody) SetTotalPages(v int32) *ListCloudGtmMonitorTemplatesResponseBody {
	s.TotalPages = &v
	return s
}

func (s *ListCloudGtmMonitorTemplatesResponseBody) Validate() error {
	if s.Templates != nil {
		if err := s.Templates.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListCloudGtmMonitorTemplatesResponseBodyTemplates struct {
	Template []*ListCloudGtmMonitorTemplatesResponseBodyTemplatesTemplate `json:"Template,omitempty" xml:"Template,omitempty" type:"Repeated"`
}

func (s ListCloudGtmMonitorTemplatesResponseBodyTemplates) String() string {
	return dara.Prettify(s)
}

func (s ListCloudGtmMonitorTemplatesResponseBodyTemplates) GoString() string {
	return s.String()
}

func (s *ListCloudGtmMonitorTemplatesResponseBodyTemplates) GetTemplate() []*ListCloudGtmMonitorTemplatesResponseBodyTemplatesTemplate {
	return s.Template
}

func (s *ListCloudGtmMonitorTemplatesResponseBodyTemplates) SetTemplate(v []*ListCloudGtmMonitorTemplatesResponseBodyTemplatesTemplate) *ListCloudGtmMonitorTemplatesResponseBodyTemplates {
	s.Template = v
	return s
}

func (s *ListCloudGtmMonitorTemplatesResponseBodyTemplates) Validate() error {
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

type ListCloudGtmMonitorTemplatesResponseBodyTemplatesTemplate struct {
	CreateTime      *string                                                                `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CreateTimestamp *int64                                                                 `json:"CreateTimestamp,omitempty" xml:"CreateTimestamp,omitempty"`
	EvaluationCount *int32                                                                 `json:"EvaluationCount,omitempty" xml:"EvaluationCount,omitempty"`
	ExtendInfo      *string                                                                `json:"ExtendInfo,omitempty" xml:"ExtendInfo,omitempty"`
	FailureRate     *int32                                                                 `json:"FailureRate,omitempty" xml:"FailureRate,omitempty"`
	Interval        *int32                                                                 `json:"Interval,omitempty" xml:"Interval,omitempty"`
	IpVersion       *string                                                                `json:"IpVersion,omitempty" xml:"IpVersion,omitempty"`
	IspCityNodes    *ListCloudGtmMonitorTemplatesResponseBodyTemplatesTemplateIspCityNodes `json:"IspCityNodes,omitempty" xml:"IspCityNodes,omitempty" type:"Struct"`
	Name            *string                                                                `json:"Name,omitempty" xml:"Name,omitempty"`
	Protocol        *string                                                                `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	Remark          *string                                                                `json:"Remark,omitempty" xml:"Remark,omitempty"`
	TemplateId      *string                                                                `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	Timeout         *int32                                                                 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
	UpdateTime      *string                                                                `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	UpdateTimestamp *int64                                                                 `json:"UpdateTimestamp,omitempty" xml:"UpdateTimestamp,omitempty"`
}

func (s ListCloudGtmMonitorTemplatesResponseBodyTemplatesTemplate) String() string {
	return dara.Prettify(s)
}

func (s ListCloudGtmMonitorTemplatesResponseBodyTemplatesTemplate) GoString() string {
	return s.String()
}

func (s *ListCloudGtmMonitorTemplatesResponseBodyTemplatesTemplate) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListCloudGtmMonitorTemplatesResponseBodyTemplatesTemplate) GetCreateTimestamp() *int64 {
	return s.CreateTimestamp
}

func (s *ListCloudGtmMonitorTemplatesResponseBodyTemplatesTemplate) GetEvaluationCount() *int32 {
	return s.EvaluationCount
}

func (s *ListCloudGtmMonitorTemplatesResponseBodyTemplatesTemplate) GetExtendInfo() *string {
	return s.ExtendInfo
}

func (s *ListCloudGtmMonitorTemplatesResponseBodyTemplatesTemplate) GetFailureRate() *int32 {
	return s.FailureRate
}

func (s *ListCloudGtmMonitorTemplatesResponseBodyTemplatesTemplate) GetInterval() *int32 {
	return s.Interval
}

func (s *ListCloudGtmMonitorTemplatesResponseBodyTemplatesTemplate) GetIpVersion() *string {
	return s.IpVersion
}

func (s *ListCloudGtmMonitorTemplatesResponseBodyTemplatesTemplate) GetIspCityNodes() *ListCloudGtmMonitorTemplatesResponseBodyTemplatesTemplateIspCityNodes {
	return s.IspCityNodes
}

func (s *ListCloudGtmMonitorTemplatesResponseBodyTemplatesTemplate) GetName() *string {
	return s.Name
}

func (s *ListCloudGtmMonitorTemplatesResponseBodyTemplatesTemplate) GetProtocol() *string {
	return s.Protocol
}

func (s *ListCloudGtmMonitorTemplatesResponseBodyTemplatesTemplate) GetRemark() *string {
	return s.Remark
}

func (s *ListCloudGtmMonitorTemplatesResponseBodyTemplatesTemplate) GetTemplateId() *string {
	return s.TemplateId
}

func (s *ListCloudGtmMonitorTemplatesResponseBodyTemplatesTemplate) GetTimeout() *int32 {
	return s.Timeout
}

func (s *ListCloudGtmMonitorTemplatesResponseBodyTemplatesTemplate) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *ListCloudGtmMonitorTemplatesResponseBodyTemplatesTemplate) GetUpdateTimestamp() *int64 {
	return s.UpdateTimestamp
}

func (s *ListCloudGtmMonitorTemplatesResponseBodyTemplatesTemplate) SetCreateTime(v string) *ListCloudGtmMonitorTemplatesResponseBodyTemplatesTemplate {
	s.CreateTime = &v
	return s
}

func (s *ListCloudGtmMonitorTemplatesResponseBodyTemplatesTemplate) SetCreateTimestamp(v int64) *ListCloudGtmMonitorTemplatesResponseBodyTemplatesTemplate {
	s.CreateTimestamp = &v
	return s
}

func (s *ListCloudGtmMonitorTemplatesResponseBodyTemplatesTemplate) SetEvaluationCount(v int32) *ListCloudGtmMonitorTemplatesResponseBodyTemplatesTemplate {
	s.EvaluationCount = &v
	return s
}

func (s *ListCloudGtmMonitorTemplatesResponseBodyTemplatesTemplate) SetExtendInfo(v string) *ListCloudGtmMonitorTemplatesResponseBodyTemplatesTemplate {
	s.ExtendInfo = &v
	return s
}

func (s *ListCloudGtmMonitorTemplatesResponseBodyTemplatesTemplate) SetFailureRate(v int32) *ListCloudGtmMonitorTemplatesResponseBodyTemplatesTemplate {
	s.FailureRate = &v
	return s
}

func (s *ListCloudGtmMonitorTemplatesResponseBodyTemplatesTemplate) SetInterval(v int32) *ListCloudGtmMonitorTemplatesResponseBodyTemplatesTemplate {
	s.Interval = &v
	return s
}

func (s *ListCloudGtmMonitorTemplatesResponseBodyTemplatesTemplate) SetIpVersion(v string) *ListCloudGtmMonitorTemplatesResponseBodyTemplatesTemplate {
	s.IpVersion = &v
	return s
}

func (s *ListCloudGtmMonitorTemplatesResponseBodyTemplatesTemplate) SetIspCityNodes(v *ListCloudGtmMonitorTemplatesResponseBodyTemplatesTemplateIspCityNodes) *ListCloudGtmMonitorTemplatesResponseBodyTemplatesTemplate {
	s.IspCityNodes = v
	return s
}

func (s *ListCloudGtmMonitorTemplatesResponseBodyTemplatesTemplate) SetName(v string) *ListCloudGtmMonitorTemplatesResponseBodyTemplatesTemplate {
	s.Name = &v
	return s
}

func (s *ListCloudGtmMonitorTemplatesResponseBodyTemplatesTemplate) SetProtocol(v string) *ListCloudGtmMonitorTemplatesResponseBodyTemplatesTemplate {
	s.Protocol = &v
	return s
}

func (s *ListCloudGtmMonitorTemplatesResponseBodyTemplatesTemplate) SetRemark(v string) *ListCloudGtmMonitorTemplatesResponseBodyTemplatesTemplate {
	s.Remark = &v
	return s
}

func (s *ListCloudGtmMonitorTemplatesResponseBodyTemplatesTemplate) SetTemplateId(v string) *ListCloudGtmMonitorTemplatesResponseBodyTemplatesTemplate {
	s.TemplateId = &v
	return s
}

func (s *ListCloudGtmMonitorTemplatesResponseBodyTemplatesTemplate) SetTimeout(v int32) *ListCloudGtmMonitorTemplatesResponseBodyTemplatesTemplate {
	s.Timeout = &v
	return s
}

func (s *ListCloudGtmMonitorTemplatesResponseBodyTemplatesTemplate) SetUpdateTime(v string) *ListCloudGtmMonitorTemplatesResponseBodyTemplatesTemplate {
	s.UpdateTime = &v
	return s
}

func (s *ListCloudGtmMonitorTemplatesResponseBodyTemplatesTemplate) SetUpdateTimestamp(v int64) *ListCloudGtmMonitorTemplatesResponseBodyTemplatesTemplate {
	s.UpdateTimestamp = &v
	return s
}

func (s *ListCloudGtmMonitorTemplatesResponseBodyTemplatesTemplate) Validate() error {
	if s.IspCityNodes != nil {
		if err := s.IspCityNodes.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListCloudGtmMonitorTemplatesResponseBodyTemplatesTemplateIspCityNodes struct {
	IspCityNode []*ListCloudGtmMonitorTemplatesResponseBodyTemplatesTemplateIspCityNodesIspCityNode `json:"IspCityNode,omitempty" xml:"IspCityNode,omitempty" type:"Repeated"`
}

func (s ListCloudGtmMonitorTemplatesResponseBodyTemplatesTemplateIspCityNodes) String() string {
	return dara.Prettify(s)
}

func (s ListCloudGtmMonitorTemplatesResponseBodyTemplatesTemplateIspCityNodes) GoString() string {
	return s.String()
}

func (s *ListCloudGtmMonitorTemplatesResponseBodyTemplatesTemplateIspCityNodes) GetIspCityNode() []*ListCloudGtmMonitorTemplatesResponseBodyTemplatesTemplateIspCityNodesIspCityNode {
	return s.IspCityNode
}

func (s *ListCloudGtmMonitorTemplatesResponseBodyTemplatesTemplateIspCityNodes) SetIspCityNode(v []*ListCloudGtmMonitorTemplatesResponseBodyTemplatesTemplateIspCityNodesIspCityNode) *ListCloudGtmMonitorTemplatesResponseBodyTemplatesTemplateIspCityNodes {
	s.IspCityNode = v
	return s
}

func (s *ListCloudGtmMonitorTemplatesResponseBodyTemplatesTemplateIspCityNodes) Validate() error {
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

type ListCloudGtmMonitorTemplatesResponseBodyTemplatesTemplateIspCityNodesIspCityNode struct {
	CityCode    *string `json:"CityCode,omitempty" xml:"CityCode,omitempty"`
	CityName    *string `json:"CityName,omitempty" xml:"CityName,omitempty"`
	CountryCode *string `json:"CountryCode,omitempty" xml:"CountryCode,omitempty"`
	CountryName *string `json:"CountryName,omitempty" xml:"CountryName,omitempty"`
	GroupName   *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	GroupType   *string `json:"GroupType,omitempty" xml:"GroupType,omitempty"`
	IspCode     *string `json:"IspCode,omitempty" xml:"IspCode,omitempty"`
	IspName     *string `json:"IspName,omitempty" xml:"IspName,omitempty"`
}

func (s ListCloudGtmMonitorTemplatesResponseBodyTemplatesTemplateIspCityNodesIspCityNode) String() string {
	return dara.Prettify(s)
}

func (s ListCloudGtmMonitorTemplatesResponseBodyTemplatesTemplateIspCityNodesIspCityNode) GoString() string {
	return s.String()
}

func (s *ListCloudGtmMonitorTemplatesResponseBodyTemplatesTemplateIspCityNodesIspCityNode) GetCityCode() *string {
	return s.CityCode
}

func (s *ListCloudGtmMonitorTemplatesResponseBodyTemplatesTemplateIspCityNodesIspCityNode) GetCityName() *string {
	return s.CityName
}

func (s *ListCloudGtmMonitorTemplatesResponseBodyTemplatesTemplateIspCityNodesIspCityNode) GetCountryCode() *string {
	return s.CountryCode
}

func (s *ListCloudGtmMonitorTemplatesResponseBodyTemplatesTemplateIspCityNodesIspCityNode) GetCountryName() *string {
	return s.CountryName
}

func (s *ListCloudGtmMonitorTemplatesResponseBodyTemplatesTemplateIspCityNodesIspCityNode) GetGroupName() *string {
	return s.GroupName
}

func (s *ListCloudGtmMonitorTemplatesResponseBodyTemplatesTemplateIspCityNodesIspCityNode) GetGroupType() *string {
	return s.GroupType
}

func (s *ListCloudGtmMonitorTemplatesResponseBodyTemplatesTemplateIspCityNodesIspCityNode) GetIspCode() *string {
	return s.IspCode
}

func (s *ListCloudGtmMonitorTemplatesResponseBodyTemplatesTemplateIspCityNodesIspCityNode) GetIspName() *string {
	return s.IspName
}

func (s *ListCloudGtmMonitorTemplatesResponseBodyTemplatesTemplateIspCityNodesIspCityNode) SetCityCode(v string) *ListCloudGtmMonitorTemplatesResponseBodyTemplatesTemplateIspCityNodesIspCityNode {
	s.CityCode = &v
	return s
}

func (s *ListCloudGtmMonitorTemplatesResponseBodyTemplatesTemplateIspCityNodesIspCityNode) SetCityName(v string) *ListCloudGtmMonitorTemplatesResponseBodyTemplatesTemplateIspCityNodesIspCityNode {
	s.CityName = &v
	return s
}

func (s *ListCloudGtmMonitorTemplatesResponseBodyTemplatesTemplateIspCityNodesIspCityNode) SetCountryCode(v string) *ListCloudGtmMonitorTemplatesResponseBodyTemplatesTemplateIspCityNodesIspCityNode {
	s.CountryCode = &v
	return s
}

func (s *ListCloudGtmMonitorTemplatesResponseBodyTemplatesTemplateIspCityNodesIspCityNode) SetCountryName(v string) *ListCloudGtmMonitorTemplatesResponseBodyTemplatesTemplateIspCityNodesIspCityNode {
	s.CountryName = &v
	return s
}

func (s *ListCloudGtmMonitorTemplatesResponseBodyTemplatesTemplateIspCityNodesIspCityNode) SetGroupName(v string) *ListCloudGtmMonitorTemplatesResponseBodyTemplatesTemplateIspCityNodesIspCityNode {
	s.GroupName = &v
	return s
}

func (s *ListCloudGtmMonitorTemplatesResponseBodyTemplatesTemplateIspCityNodesIspCityNode) SetGroupType(v string) *ListCloudGtmMonitorTemplatesResponseBodyTemplatesTemplateIspCityNodesIspCityNode {
	s.GroupType = &v
	return s
}

func (s *ListCloudGtmMonitorTemplatesResponseBodyTemplatesTemplateIspCityNodesIspCityNode) SetIspCode(v string) *ListCloudGtmMonitorTemplatesResponseBodyTemplatesTemplateIspCityNodesIspCityNode {
	s.IspCode = &v
	return s
}

func (s *ListCloudGtmMonitorTemplatesResponseBodyTemplatesTemplateIspCityNodesIspCityNode) SetIspName(v string) *ListCloudGtmMonitorTemplatesResponseBodyTemplatesTemplateIspCityNodesIspCityNode {
	s.IspName = &v
	return s
}

func (s *ListCloudGtmMonitorTemplatesResponseBodyTemplatesTemplateIspCityNodesIspCityNode) Validate() error {
	return dara.Validate(s)
}
