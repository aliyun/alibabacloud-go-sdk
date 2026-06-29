// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBizMetricByNameResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetBizMetricByNameResponseBody
	GetCode() *string
	SetData(v *GetBizMetricByNameResponseBodyData) *GetBizMetricByNameResponseBody
	GetData() *GetBizMetricByNameResponseBodyData
	SetHttpStatusCode(v int32) *GetBizMetricByNameResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetBizMetricByNameResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetBizMetricByNameResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetBizMetricByNameResponseBody
	GetSuccess() *bool
}

type GetBizMetricByNameResponseBody struct {
	// The response code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The details of the business metric.
	Data *GetBizMetricByNameResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// SUCCESS
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 75DD06F8-1661-5A6E-B0A6-7E23133BDC60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetBizMetricByNameResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetBizMetricByNameResponseBody) GoString() string {
	return s.String()
}

func (s *GetBizMetricByNameResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetBizMetricByNameResponseBody) GetData() *GetBizMetricByNameResponseBodyData {
	return s.Data
}

func (s *GetBizMetricByNameResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetBizMetricByNameResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetBizMetricByNameResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetBizMetricByNameResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetBizMetricByNameResponseBody) SetCode(v string) *GetBizMetricByNameResponseBody {
	s.Code = &v
	return s
}

func (s *GetBizMetricByNameResponseBody) SetData(v *GetBizMetricByNameResponseBodyData) *GetBizMetricByNameResponseBody {
	s.Data = v
	return s
}

func (s *GetBizMetricByNameResponseBody) SetHttpStatusCode(v int32) *GetBizMetricByNameResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetBizMetricByNameResponseBody) SetMessage(v string) *GetBizMetricByNameResponseBody {
	s.Message = &v
	return s
}

func (s *GetBizMetricByNameResponseBody) SetRequestId(v string) *GetBizMetricByNameResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetBizMetricByNameResponseBody) SetSuccess(v bool) *GetBizMetricByNameResponseBody {
	s.Success = &v
	return s
}

func (s *GetBizMetricByNameResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetBizMetricByNameResponseBodyData struct {
	// The list of associated technical metrics.
	AssociatedTechMetrics []*GetBizMetricByNameResponseBodyDataAssociatedTechMetrics `json:"AssociatedTechMetrics,omitempty" xml:"AssociatedTechMetrics,omitempty" type:"Repeated"`
	// The name of the business owner.
	//
	// example:
	//
	// SuperAdmin
	BizOwnerName *string `json:"BizOwnerName,omitempty" xml:"BizOwnerName,omitempty"`
	// The list of affiliated catalogs.
	Catalogs []*GetBizMetricByNameResponseBodyDataCatalogs `json:"Catalogs,omitempty" xml:"Catalogs,omitempty" type:"Repeated"`
	// The list of custom attributes.
	CustomAttribute []*GetBizMetricByNameResponseBodyDataCustomAttribute `json:"CustomAttribute,omitempty" xml:"CustomAttribute,omitempty" type:"Repeated"`
	// The description.
	//
	// example:
	//
	// Metric Desc
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The display name.
	//
	// example:
	//
	// Metric Display Name
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The GUID of the business metric.
	//
	// example:
	//
	// guid
	Guid *string `json:"Guid,omitempty" xml:"Guid,omitempty"`
	// The list of labels.
	Labels []*string `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	// The metric definition.
	//
	// example:
	//
	// [Metric1]+[Metric2]
	MetricDefinition *string `json:"MetricDefinition,omitempty" xml:"MetricDefinition,omitempty"`
	// The expression of the metric relation diagram.
	//
	// example:
	//
	// [Metric2]*10
	MetricRelationDiagramExpression *string `json:"MetricRelationDiagramExpression,omitempty" xml:"MetricRelationDiagramExpression,omitempty"`
	// Indicates whether the metric relation diagram is enabled. A value of true indicates that the diagram is enabled. A value of false indicates that the diagram is disabled.
	//
	// example:
	//
	// true
	MetricRelationDiagramSwitchOpen *bool `json:"MetricRelationDiagramSwitchOpen,omitempty" xml:"MetricRelationDiagramSwitchOpen,omitempty"`
	// The name of the business metric.
	//
	// example:
	//
	// Metric1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The text content of the operation instruction.
	//
	// example:
	//
	// content
	OperateInstructionContent *string `json:"OperateInstructionContent,omitempty" xml:"OperateInstructionContent,omitempty"`
	// Indicates whether the operation instruction is enabled. A value of true indicates that the operation instruction is enabled. A value of false indicates that the operation instruction is disabled.
	//
	// example:
	//
	// true
	OperateInstructionEnabled *bool `json:"OperateInstructionEnabled,omitempty" xml:"OperateInstructionEnabled,omitempty"`
	// The list of related business metrics.
	RelatedBizMetrics []*GetBizMetricByNameResponseBodyDataRelatedBizMetrics `json:"RelatedBizMetrics,omitempty" xml:"RelatedBizMetrics,omitempty" type:"Repeated"`
	// The tenant ID.
	//
	// example:
	//
	// 30001011
	TenantId *int64 `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// The view scope.
	ViewScope *GetBizMetricByNameResponseBodyDataViewScope `json:"ViewScope,omitempty" xml:"ViewScope,omitempty" type:"Struct"`
}

func (s GetBizMetricByNameResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetBizMetricByNameResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetBizMetricByNameResponseBodyData) GetAssociatedTechMetrics() []*GetBizMetricByNameResponseBodyDataAssociatedTechMetrics {
	return s.AssociatedTechMetrics
}

func (s *GetBizMetricByNameResponseBodyData) GetBizOwnerName() *string {
	return s.BizOwnerName
}

func (s *GetBizMetricByNameResponseBodyData) GetCatalogs() []*GetBizMetricByNameResponseBodyDataCatalogs {
	return s.Catalogs
}

func (s *GetBizMetricByNameResponseBodyData) GetCustomAttribute() []*GetBizMetricByNameResponseBodyDataCustomAttribute {
	return s.CustomAttribute
}

func (s *GetBizMetricByNameResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *GetBizMetricByNameResponseBodyData) GetDisplayName() *string {
	return s.DisplayName
}

func (s *GetBizMetricByNameResponseBodyData) GetGuid() *string {
	return s.Guid
}

func (s *GetBizMetricByNameResponseBodyData) GetLabels() []*string {
	return s.Labels
}

func (s *GetBizMetricByNameResponseBodyData) GetMetricDefinition() *string {
	return s.MetricDefinition
}

func (s *GetBizMetricByNameResponseBodyData) GetMetricRelationDiagramExpression() *string {
	return s.MetricRelationDiagramExpression
}

func (s *GetBizMetricByNameResponseBodyData) GetMetricRelationDiagramSwitchOpen() *bool {
	return s.MetricRelationDiagramSwitchOpen
}

func (s *GetBizMetricByNameResponseBodyData) GetName() *string {
	return s.Name
}

func (s *GetBizMetricByNameResponseBodyData) GetOperateInstructionContent() *string {
	return s.OperateInstructionContent
}

func (s *GetBizMetricByNameResponseBodyData) GetOperateInstructionEnabled() *bool {
	return s.OperateInstructionEnabled
}

func (s *GetBizMetricByNameResponseBodyData) GetRelatedBizMetrics() []*GetBizMetricByNameResponseBodyDataRelatedBizMetrics {
	return s.RelatedBizMetrics
}

func (s *GetBizMetricByNameResponseBodyData) GetTenantId() *int64 {
	return s.TenantId
}

func (s *GetBizMetricByNameResponseBodyData) GetViewScope() *GetBizMetricByNameResponseBodyDataViewScope {
	return s.ViewScope
}

func (s *GetBizMetricByNameResponseBodyData) SetAssociatedTechMetrics(v []*GetBizMetricByNameResponseBodyDataAssociatedTechMetrics) *GetBizMetricByNameResponseBodyData {
	s.AssociatedTechMetrics = v
	return s
}

func (s *GetBizMetricByNameResponseBodyData) SetBizOwnerName(v string) *GetBizMetricByNameResponseBodyData {
	s.BizOwnerName = &v
	return s
}

func (s *GetBizMetricByNameResponseBodyData) SetCatalogs(v []*GetBizMetricByNameResponseBodyDataCatalogs) *GetBizMetricByNameResponseBodyData {
	s.Catalogs = v
	return s
}

func (s *GetBizMetricByNameResponseBodyData) SetCustomAttribute(v []*GetBizMetricByNameResponseBodyDataCustomAttribute) *GetBizMetricByNameResponseBodyData {
	s.CustomAttribute = v
	return s
}

func (s *GetBizMetricByNameResponseBodyData) SetDescription(v string) *GetBizMetricByNameResponseBodyData {
	s.Description = &v
	return s
}

func (s *GetBizMetricByNameResponseBodyData) SetDisplayName(v string) *GetBizMetricByNameResponseBodyData {
	s.DisplayName = &v
	return s
}

func (s *GetBizMetricByNameResponseBodyData) SetGuid(v string) *GetBizMetricByNameResponseBodyData {
	s.Guid = &v
	return s
}

func (s *GetBizMetricByNameResponseBodyData) SetLabels(v []*string) *GetBizMetricByNameResponseBodyData {
	s.Labels = v
	return s
}

func (s *GetBizMetricByNameResponseBodyData) SetMetricDefinition(v string) *GetBizMetricByNameResponseBodyData {
	s.MetricDefinition = &v
	return s
}

func (s *GetBizMetricByNameResponseBodyData) SetMetricRelationDiagramExpression(v string) *GetBizMetricByNameResponseBodyData {
	s.MetricRelationDiagramExpression = &v
	return s
}

func (s *GetBizMetricByNameResponseBodyData) SetMetricRelationDiagramSwitchOpen(v bool) *GetBizMetricByNameResponseBodyData {
	s.MetricRelationDiagramSwitchOpen = &v
	return s
}

func (s *GetBizMetricByNameResponseBodyData) SetName(v string) *GetBizMetricByNameResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetBizMetricByNameResponseBodyData) SetOperateInstructionContent(v string) *GetBizMetricByNameResponseBodyData {
	s.OperateInstructionContent = &v
	return s
}

func (s *GetBizMetricByNameResponseBodyData) SetOperateInstructionEnabled(v bool) *GetBizMetricByNameResponseBodyData {
	s.OperateInstructionEnabled = &v
	return s
}

func (s *GetBizMetricByNameResponseBodyData) SetRelatedBizMetrics(v []*GetBizMetricByNameResponseBodyDataRelatedBizMetrics) *GetBizMetricByNameResponseBodyData {
	s.RelatedBizMetrics = v
	return s
}

func (s *GetBizMetricByNameResponseBodyData) SetTenantId(v int64) *GetBizMetricByNameResponseBodyData {
	s.TenantId = &v
	return s
}

func (s *GetBizMetricByNameResponseBodyData) SetViewScope(v *GetBizMetricByNameResponseBodyDataViewScope) *GetBizMetricByNameResponseBodyData {
	s.ViewScope = v
	return s
}

func (s *GetBizMetricByNameResponseBodyData) Validate() error {
	if s.AssociatedTechMetrics != nil {
		for _, item := range s.AssociatedTechMetrics {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Catalogs != nil {
		for _, item := range s.Catalogs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.CustomAttribute != nil {
		for _, item := range s.CustomAttribute {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.RelatedBizMetrics != nil {
		for _, item := range s.RelatedBizMetrics {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ViewScope != nil {
		if err := s.ViewScope.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetBizMetricByNameResponseBodyDataAssociatedTechMetrics struct {
	// The description.
	//
	// example:
	//
	// desc
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The display name.
	//
	// example:
	//
	// display name
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The GUID.
	//
	// example:
	//
	// table1.a.b
	Guid *string `json:"Guid,omitempty" xml:"Guid,omitempty"`
	// The name.
	//
	// example:
	//
	// metric3
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The type of the technical metric. Valid values: INDEX (modeling metric) and CUSTOM_INDEX (custom metric).
	//
	// example:
	//
	// CUSTOM_INDEX
	SubType *string `json:"SubType,omitempty" xml:"SubType,omitempty"`
}

func (s GetBizMetricByNameResponseBodyDataAssociatedTechMetrics) String() string {
	return dara.Prettify(s)
}

func (s GetBizMetricByNameResponseBodyDataAssociatedTechMetrics) GoString() string {
	return s.String()
}

func (s *GetBizMetricByNameResponseBodyDataAssociatedTechMetrics) GetDescription() *string {
	return s.Description
}

func (s *GetBizMetricByNameResponseBodyDataAssociatedTechMetrics) GetDisplayName() *string {
	return s.DisplayName
}

func (s *GetBizMetricByNameResponseBodyDataAssociatedTechMetrics) GetGuid() *string {
	return s.Guid
}

func (s *GetBizMetricByNameResponseBodyDataAssociatedTechMetrics) GetName() *string {
	return s.Name
}

func (s *GetBizMetricByNameResponseBodyDataAssociatedTechMetrics) GetSubType() *string {
	return s.SubType
}

func (s *GetBizMetricByNameResponseBodyDataAssociatedTechMetrics) SetDescription(v string) *GetBizMetricByNameResponseBodyDataAssociatedTechMetrics {
	s.Description = &v
	return s
}

func (s *GetBizMetricByNameResponseBodyDataAssociatedTechMetrics) SetDisplayName(v string) *GetBizMetricByNameResponseBodyDataAssociatedTechMetrics {
	s.DisplayName = &v
	return s
}

func (s *GetBizMetricByNameResponseBodyDataAssociatedTechMetrics) SetGuid(v string) *GetBizMetricByNameResponseBodyDataAssociatedTechMetrics {
	s.Guid = &v
	return s
}

func (s *GetBizMetricByNameResponseBodyDataAssociatedTechMetrics) SetName(v string) *GetBizMetricByNameResponseBodyDataAssociatedTechMetrics {
	s.Name = &v
	return s
}

func (s *GetBizMetricByNameResponseBodyDataAssociatedTechMetrics) SetSubType(v string) *GetBizMetricByNameResponseBodyDataAssociatedTechMetrics {
	s.SubType = &v
	return s
}

func (s *GetBizMetricByNameResponseBodyDataAssociatedTechMetrics) Validate() error {
	return dara.Validate(s)
}

type GetBizMetricByNameResponseBodyDataCatalogs struct {
	// The catalog description.
	//
	// example:
	//
	// catalog desc
	CatalogDesc *string `json:"CatalogDesc,omitempty" xml:"CatalogDesc,omitempty"`
	// The catalog ID.
	//
	// example:
	//
	// 1561740764851842
	CatalogId *int64 `json:"CatalogId,omitempty" xml:"CatalogId,omitempty"`
	// The catalog name.
	//
	// example:
	//
	// test catalog
	CatalogName *string `json:"CatalogName,omitempty" xml:"CatalogName,omitempty"`
	// The parent catalog ID.
	//
	// example:
	//
	// 1561740764851841
	ParentCatalogId *int64 `json:"ParentCatalogId,omitempty" xml:"ParentCatalogId,omitempty"`
	// The parent path of the catalog.
	//
	// example:
	//
	// /catalog1/
	ParentPath *string `json:"ParentPath,omitempty" xml:"ParentPath,omitempty"`
	// The topic ID to which the catalog belongs.
	//
	// example:
	//
	// 43297700
	TopicId *int64 `json:"TopicId,omitempty" xml:"TopicId,omitempty"`
	// The topic name to which the catalog belongs.
	//
	// example:
	//
	// test topic
	TopicName *string `json:"TopicName,omitempty" xml:"TopicName,omitempty"`
}

func (s GetBizMetricByNameResponseBodyDataCatalogs) String() string {
	return dara.Prettify(s)
}

func (s GetBizMetricByNameResponseBodyDataCatalogs) GoString() string {
	return s.String()
}

func (s *GetBizMetricByNameResponseBodyDataCatalogs) GetCatalogDesc() *string {
	return s.CatalogDesc
}

func (s *GetBizMetricByNameResponseBodyDataCatalogs) GetCatalogId() *int64 {
	return s.CatalogId
}

func (s *GetBizMetricByNameResponseBodyDataCatalogs) GetCatalogName() *string {
	return s.CatalogName
}

func (s *GetBizMetricByNameResponseBodyDataCatalogs) GetParentCatalogId() *int64 {
	return s.ParentCatalogId
}

func (s *GetBizMetricByNameResponseBodyDataCatalogs) GetParentPath() *string {
	return s.ParentPath
}

func (s *GetBizMetricByNameResponseBodyDataCatalogs) GetTopicId() *int64 {
	return s.TopicId
}

func (s *GetBizMetricByNameResponseBodyDataCatalogs) GetTopicName() *string {
	return s.TopicName
}

func (s *GetBizMetricByNameResponseBodyDataCatalogs) SetCatalogDesc(v string) *GetBizMetricByNameResponseBodyDataCatalogs {
	s.CatalogDesc = &v
	return s
}

func (s *GetBizMetricByNameResponseBodyDataCatalogs) SetCatalogId(v int64) *GetBizMetricByNameResponseBodyDataCatalogs {
	s.CatalogId = &v
	return s
}

func (s *GetBizMetricByNameResponseBodyDataCatalogs) SetCatalogName(v string) *GetBizMetricByNameResponseBodyDataCatalogs {
	s.CatalogName = &v
	return s
}

func (s *GetBizMetricByNameResponseBodyDataCatalogs) SetParentCatalogId(v int64) *GetBizMetricByNameResponseBodyDataCatalogs {
	s.ParentCatalogId = &v
	return s
}

func (s *GetBizMetricByNameResponseBodyDataCatalogs) SetParentPath(v string) *GetBizMetricByNameResponseBodyDataCatalogs {
	s.ParentPath = &v
	return s
}

func (s *GetBizMetricByNameResponseBodyDataCatalogs) SetTopicId(v int64) *GetBizMetricByNameResponseBodyDataCatalogs {
	s.TopicId = &v
	return s
}

func (s *GetBizMetricByNameResponseBodyDataCatalogs) SetTopicName(v string) *GetBizMetricByNameResponseBodyDataCatalogs {
	s.TopicName = &v
	return s
}

func (s *GetBizMetricByNameResponseBodyDataCatalogs) Validate() error {
	return dara.Validate(s)
}

type GetBizMetricByNameResponseBodyDataCustomAttribute struct {
	// The code of the custom attribute.
	//
	// example:
	//
	// CustomAttributeCode
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The list of attribute values. 1. For custom input and single-select dropdown attributes, the first value in the list is used. 2. For multi-select dropdown attributes, all values in the list are used. 3. For hyperlink attributes, the first value is the display text and the second value is the link URL.
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s GetBizMetricByNameResponseBodyDataCustomAttribute) String() string {
	return dara.Prettify(s)
}

func (s GetBizMetricByNameResponseBodyDataCustomAttribute) GoString() string {
	return s.String()
}

func (s *GetBizMetricByNameResponseBodyDataCustomAttribute) GetCode() *string {
	return s.Code
}

func (s *GetBizMetricByNameResponseBodyDataCustomAttribute) GetValues() []*string {
	return s.Values
}

func (s *GetBizMetricByNameResponseBodyDataCustomAttribute) SetCode(v string) *GetBizMetricByNameResponseBodyDataCustomAttribute {
	s.Code = &v
	return s
}

func (s *GetBizMetricByNameResponseBodyDataCustomAttribute) SetValues(v []*string) *GetBizMetricByNameResponseBodyDataCustomAttribute {
	s.Values = v
	return s
}

func (s *GetBizMetricByNameResponseBodyDataCustomAttribute) Validate() error {
	return dara.Validate(s)
}

type GetBizMetricByNameResponseBodyDataRelatedBizMetrics struct {
	// The description.
	//
	// example:
	//
	// desc
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The display name.
	//
	// example:
	//
	// display name
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The GUID.
	//
	// example:
	//
	// test
	Guid *string `json:"Guid,omitempty" xml:"Guid,omitempty"`
	// The name.
	//
	// example:
	//
	// Metric2
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The relation type. Valid values: POSITIVE (positive correlation), NEGATIVE (negative correlation), and OTHER (other).
	//
	// example:
	//
	// POSITIVE
	RelationType *string `json:"RelationType,omitempty" xml:"RelationType,omitempty"`
}

func (s GetBizMetricByNameResponseBodyDataRelatedBizMetrics) String() string {
	return dara.Prettify(s)
}

func (s GetBizMetricByNameResponseBodyDataRelatedBizMetrics) GoString() string {
	return s.String()
}

func (s *GetBizMetricByNameResponseBodyDataRelatedBizMetrics) GetDescription() *string {
	return s.Description
}

func (s *GetBizMetricByNameResponseBodyDataRelatedBizMetrics) GetDisplayName() *string {
	return s.DisplayName
}

func (s *GetBizMetricByNameResponseBodyDataRelatedBizMetrics) GetGuid() *string {
	return s.Guid
}

func (s *GetBizMetricByNameResponseBodyDataRelatedBizMetrics) GetName() *string {
	return s.Name
}

func (s *GetBizMetricByNameResponseBodyDataRelatedBizMetrics) GetRelationType() *string {
	return s.RelationType
}

func (s *GetBizMetricByNameResponseBodyDataRelatedBizMetrics) SetDescription(v string) *GetBizMetricByNameResponseBodyDataRelatedBizMetrics {
	s.Description = &v
	return s
}

func (s *GetBizMetricByNameResponseBodyDataRelatedBizMetrics) SetDisplayName(v string) *GetBizMetricByNameResponseBodyDataRelatedBizMetrics {
	s.DisplayName = &v
	return s
}

func (s *GetBizMetricByNameResponseBodyDataRelatedBizMetrics) SetGuid(v string) *GetBizMetricByNameResponseBodyDataRelatedBizMetrics {
	s.Guid = &v
	return s
}

func (s *GetBizMetricByNameResponseBodyDataRelatedBizMetrics) SetName(v string) *GetBizMetricByNameResponseBodyDataRelatedBizMetrics {
	s.Name = &v
	return s
}

func (s *GetBizMetricByNameResponseBodyDataRelatedBizMetrics) SetRelationType(v string) *GetBizMetricByNameResponseBodyDataRelatedBizMetrics {
	s.RelationType = &v
	return s
}

func (s *GetBizMetricByNameResponseBodyDataRelatedBizMetrics) Validate() error {
	return dara.Validate(s)
}

type GetBizMetricByNameResponseBodyDataViewScope struct {
	// The type of view scope. Valid values: ALL_USERS_CAN_VIEW (visible to all users), PART_USERS_CAN_VIEW (visible to specific users), and PART_USERS_CAN_NOT_VIEW (invisible to specific users).
	//
	// example:
	//
	// PART_USERS_CAN_NOT_VIEW
	ScopeType *string `json:"ScopeType,omitempty" xml:"ScopeType,omitempty"`
	// The names of user groups. This parameter takes effect only when the view scope is set to PART_USERS_CAN_VIEW or PART_USERS_CAN_NOT_VIEW.
	UserGroupNames []*string `json:"UserGroupNames,omitempty" xml:"UserGroupNames,omitempty" type:"Repeated"`
	// The usernames of individual accounts. This parameter is valid only when the view scope is set to PART_USERS_CAN_VIEW or PART_USERS_CAN_NOT_VIEW.
	UserNames []*string `json:"UserNames,omitempty" xml:"UserNames,omitempty" type:"Repeated"`
}

func (s GetBizMetricByNameResponseBodyDataViewScope) String() string {
	return dara.Prettify(s)
}

func (s GetBizMetricByNameResponseBodyDataViewScope) GoString() string {
	return s.String()
}

func (s *GetBizMetricByNameResponseBodyDataViewScope) GetScopeType() *string {
	return s.ScopeType
}

func (s *GetBizMetricByNameResponseBodyDataViewScope) GetUserGroupNames() []*string {
	return s.UserGroupNames
}

func (s *GetBizMetricByNameResponseBodyDataViewScope) GetUserNames() []*string {
	return s.UserNames
}

func (s *GetBizMetricByNameResponseBodyDataViewScope) SetScopeType(v string) *GetBizMetricByNameResponseBodyDataViewScope {
	s.ScopeType = &v
	return s
}

func (s *GetBizMetricByNameResponseBodyDataViewScope) SetUserGroupNames(v []*string) *GetBizMetricByNameResponseBodyDataViewScope {
	s.UserGroupNames = v
	return s
}

func (s *GetBizMetricByNameResponseBodyDataViewScope) SetUserNames(v []*string) *GetBizMetricByNameResponseBodyDataViewScope {
	s.UserNames = v
	return s
}

func (s *GetBizMetricByNameResponseBodyDataViewScope) Validate() error {
	return dara.Validate(s)
}
