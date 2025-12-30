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
	// example:
	//
	// OK
	Code *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetBizMetricByNameResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
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
	AssociatedTechMetrics []*GetBizMetricByNameResponseBodyDataAssociatedTechMetrics `json:"AssociatedTechMetrics,omitempty" xml:"AssociatedTechMetrics,omitempty" type:"Repeated"`
	// example:
	//
	// SuperAdmin
	BizOwnerName    *string                                              `json:"BizOwnerName,omitempty" xml:"BizOwnerName,omitempty"`
	Catalogs        []*GetBizMetricByNameResponseBodyDataCatalogs        `json:"Catalogs,omitempty" xml:"Catalogs,omitempty" type:"Repeated"`
	CustomAttribute []*GetBizMetricByNameResponseBodyDataCustomAttribute `json:"CustomAttribute,omitempty" xml:"CustomAttribute,omitempty" type:"Repeated"`
	// example:
	//
	// Metric Desc
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// Metric Display Name
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// example:
	//
	// guid
	Guid   *string   `json:"Guid,omitempty" xml:"Guid,omitempty"`
	Labels []*string `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	// example:
	//
	// [Metric1]+[Metric2]
	MetricDefinition *string `json:"MetricDefinition,omitempty" xml:"MetricDefinition,omitempty"`
	// example:
	//
	// [Metric2]*10
	MetricRelationDiagramExpression *string `json:"MetricRelationDiagramExpression,omitempty" xml:"MetricRelationDiagramExpression,omitempty"`
	// example:
	//
	// true
	MetricRelationDiagramSwitchOpen *bool `json:"MetricRelationDiagramSwitchOpen,omitempty" xml:"MetricRelationDiagramSwitchOpen,omitempty"`
	// example:
	//
	// Metric1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// content
	OperateInstructionContent *string `json:"OperateInstructionContent,omitempty" xml:"OperateInstructionContent,omitempty"`
	// example:
	//
	// true
	OperateInstructionEnabled *bool                                                  `json:"OperateInstructionEnabled,omitempty" xml:"OperateInstructionEnabled,omitempty"`
	RelatedBizMetrics         []*GetBizMetricByNameResponseBodyDataRelatedBizMetrics `json:"RelatedBizMetrics,omitempty" xml:"RelatedBizMetrics,omitempty" type:"Repeated"`
	// example:
	//
	// 30001011
	TenantId  *int64                                       `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
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
	// example:
	//
	// desc
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// display name
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// guid
	//
	// example:
	//
	// table1.a.b
	Guid *string `json:"Guid,omitempty" xml:"Guid,omitempty"`
	// example:
	//
	// metric3
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
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
	// example:
	//
	// catalog desc
	CatalogDesc *string `json:"CatalogDesc,omitempty" xml:"CatalogDesc,omitempty"`
	// example:
	//
	// 1561740764851842
	CatalogId *int64 `json:"CatalogId,omitempty" xml:"CatalogId,omitempty"`
	// example:
	//
	// test catalog
	CatalogName *string `json:"CatalogName,omitempty" xml:"CatalogName,omitempty"`
	// example:
	//
	// 1561740764851841
	ParentCatalogId *int64 `json:"ParentCatalogId,omitempty" xml:"ParentCatalogId,omitempty"`
	// example:
	//
	// /catalog1/
	ParentPath *string `json:"ParentPath,omitempty" xml:"ParentPath,omitempty"`
	// example:
	//
	// 43297700
	TopicId *int64 `json:"TopicId,omitempty" xml:"TopicId,omitempty"`
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
	// example:
	//
	// CustomAttributeCode
	Code   *string   `json:"Code,omitempty" xml:"Code,omitempty"`
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
	// example:
	//
	// desc
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// display name
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// guid
	//
	// example:
	//
	// test
	Guid *string `json:"Guid,omitempty" xml:"Guid,omitempty"`
	// example:
	//
	// Metric2
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
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
	// example:
	//
	// PART_USERS_CAN_NOT_VIEW
	ScopeType      *string   `json:"ScopeType,omitempty" xml:"ScopeType,omitempty"`
	UserGroupNames []*string `json:"UserGroupNames,omitempty" xml:"UserGroupNames,omitempty" type:"Repeated"`
	UserNames      []*string `json:"UserNames,omitempty" xml:"UserNames,omitempty" type:"Repeated"`
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
