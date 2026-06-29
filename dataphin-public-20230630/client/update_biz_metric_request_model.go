// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateBizMetricRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *UpdateBizMetricRequest
	GetOpTenantId() *int64
	SetUpdateBizMetricCommand(v *UpdateBizMetricRequestUpdateBizMetricCommand) *UpdateBizMetricRequest
	GetUpdateBizMetricCommand() *UpdateBizMetricRequestUpdateBizMetricCommand
}

type UpdateBizMetricRequest struct {
	// Tenant ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// Update request
	//
	// This parameter is required.
	UpdateBizMetricCommand *UpdateBizMetricRequestUpdateBizMetricCommand `json:"UpdateBizMetricCommand,omitempty" xml:"UpdateBizMetricCommand,omitempty" type:"Struct"`
}

func (s UpdateBizMetricRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateBizMetricRequest) GoString() string {
	return s.String()
}

func (s *UpdateBizMetricRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *UpdateBizMetricRequest) GetUpdateBizMetricCommand() *UpdateBizMetricRequestUpdateBizMetricCommand {
	return s.UpdateBizMetricCommand
}

func (s *UpdateBizMetricRequest) SetOpTenantId(v int64) *UpdateBizMetricRequest {
	s.OpTenantId = &v
	return s
}

func (s *UpdateBizMetricRequest) SetUpdateBizMetricCommand(v *UpdateBizMetricRequestUpdateBizMetricCommand) *UpdateBizMetricRequest {
	s.UpdateBizMetricCommand = v
	return s
}

func (s *UpdateBizMetricRequest) Validate() error {
	if s.UpdateBizMetricCommand != nil {
		if err := s.UpdateBizMetricCommand.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateBizMetricRequestUpdateBizMetricCommand struct {
	// List of associated technical metrics. Enter the full name of the technical metric in the format of "TableFullName.MetricName", where "TableFullName" equals "AssetSource.TableName". A technical metric can only be associated with one business metric and cannot be associated repeatedly
	AssociatedTechMetricFullNames []*string `json:"AssociatedTechMetricFullNames,omitempty" xml:"AssociatedTechMetricFullNames,omitempty" type:"Repeated"`
	// Business owner. Enter the username of the owner account
	//
	// example:
	//
	// SuperAdmin
	BizOwnerName *string `json:"BizOwnerName,omitempty" xml:"BizOwnerName,omitempty"`
	// List of catalog IDs
	CatalogIds []*int64 `json:"CatalogIds,omitempty" xml:"CatalogIds,omitempty" type:"Repeated"`
	// List of custom attributes. Enter the attribute code and attribute values for each
	CustomAttribute []*UpdateBizMetricRequestUpdateBizMetricCommandCustomAttribute `json:"CustomAttribute,omitempty" xml:"CustomAttribute,omitempty" type:"Repeated"`
	// Description
	//
	// example:
	//
	// Metric Desc
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Display name
	//
	// example:
	//
	// Metric1_DisplayName
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// Asset labels
	Labels []*string `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	// Metric definition. To reference other business metrics, enclose the metric name in square brackets [ ]
	//
	// example:
	//
	// [Metric2]+[Metric3]
	MetricDefinition *string `json:"MetricDefinition,omitempty" xml:"MetricDefinition,omitempty"`
	// This parameter is read only when the metric relationship diagram is enabled. Only a calculation expression composed of metric names selected from related business metrics is supported. Supported operators include +, -, *, /, (), %, and ∑. Each metric name must be enclosed in square brackets [ ]. If no operator is specified between two metrics, the system automatically fills in a placeholder. If no metric relationship expression is configured, the metric relationship diagram switch is automatically disabled
	//
	// example:
	//
	// [Metric1]+[Metric2]
	MetricRelationDiagramExpression *string `json:"MetricRelationDiagramExpression,omitempty" xml:"MetricRelationDiagramExpression,omitempty"`
	// Metric relationship diagram switch. Valid values: true (enabled) and false (disabled). This can be enabled only when at least one related business metric exists. Otherwise, it is automatically disabled
	//
	// example:
	//
	// true
	MetricRelationDiagramSwitchOpen *bool `json:"MetricRelationDiagramSwitchOpen,omitempty" xml:"MetricRelationDiagramSwitchOpen,omitempty"`
	// Enter the name of the business metric to update
	//
	// This parameter is required.
	//
	// example:
	//
	// Metric1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The new name. Enter this when you need to modify the metric name
	//
	// example:
	//
	// Metric1_new
	NewName *string `json:"NewName,omitempty" xml:"NewName,omitempty"`
	// Content of the usage instructions. Only text format is supported
	//
	// example:
	//
	// content test
	OperateInstructionContent *string `json:"OperateInstructionContent,omitempty" xml:"OperateInstructionContent,omitempty"`
	// Specifies whether the usage instructions are enabled. Valid values: true (enabled) and false (disabled)
	//
	// example:
	//
	// true
	OperateInstructionEnabled *bool `json:"OperateInstructionEnabled,omitempty" xml:"OperateInstructionEnabled,omitempty"`
	// List of related business metrics
	RelatedBizMetrics []*UpdateBizMetricRequestUpdateBizMetricCommandRelatedBizMetrics `json:"RelatedBizMetrics,omitempty" xml:"RelatedBizMetrics,omitempty" type:"Repeated"`
	// Visibility scope
	ViewScope *UpdateBizMetricRequestUpdateBizMetricCommandViewScope `json:"ViewScope,omitempty" xml:"ViewScope,omitempty" type:"Struct"`
}

func (s UpdateBizMetricRequestUpdateBizMetricCommand) String() string {
	return dara.Prettify(s)
}

func (s UpdateBizMetricRequestUpdateBizMetricCommand) GoString() string {
	return s.String()
}

func (s *UpdateBizMetricRequestUpdateBizMetricCommand) GetAssociatedTechMetricFullNames() []*string {
	return s.AssociatedTechMetricFullNames
}

func (s *UpdateBizMetricRequestUpdateBizMetricCommand) GetBizOwnerName() *string {
	return s.BizOwnerName
}

func (s *UpdateBizMetricRequestUpdateBizMetricCommand) GetCatalogIds() []*int64 {
	return s.CatalogIds
}

func (s *UpdateBizMetricRequestUpdateBizMetricCommand) GetCustomAttribute() []*UpdateBizMetricRequestUpdateBizMetricCommandCustomAttribute {
	return s.CustomAttribute
}

func (s *UpdateBizMetricRequestUpdateBizMetricCommand) GetDescription() *string {
	return s.Description
}

func (s *UpdateBizMetricRequestUpdateBizMetricCommand) GetDisplayName() *string {
	return s.DisplayName
}

func (s *UpdateBizMetricRequestUpdateBizMetricCommand) GetLabels() []*string {
	return s.Labels
}

func (s *UpdateBizMetricRequestUpdateBizMetricCommand) GetMetricDefinition() *string {
	return s.MetricDefinition
}

func (s *UpdateBizMetricRequestUpdateBizMetricCommand) GetMetricRelationDiagramExpression() *string {
	return s.MetricRelationDiagramExpression
}

func (s *UpdateBizMetricRequestUpdateBizMetricCommand) GetMetricRelationDiagramSwitchOpen() *bool {
	return s.MetricRelationDiagramSwitchOpen
}

func (s *UpdateBizMetricRequestUpdateBizMetricCommand) GetName() *string {
	return s.Name
}

func (s *UpdateBizMetricRequestUpdateBizMetricCommand) GetNewName() *string {
	return s.NewName
}

func (s *UpdateBizMetricRequestUpdateBizMetricCommand) GetOperateInstructionContent() *string {
	return s.OperateInstructionContent
}

func (s *UpdateBizMetricRequestUpdateBizMetricCommand) GetOperateInstructionEnabled() *bool {
	return s.OperateInstructionEnabled
}

func (s *UpdateBizMetricRequestUpdateBizMetricCommand) GetRelatedBizMetrics() []*UpdateBizMetricRequestUpdateBizMetricCommandRelatedBizMetrics {
	return s.RelatedBizMetrics
}

func (s *UpdateBizMetricRequestUpdateBizMetricCommand) GetViewScope() *UpdateBizMetricRequestUpdateBizMetricCommandViewScope {
	return s.ViewScope
}

func (s *UpdateBizMetricRequestUpdateBizMetricCommand) SetAssociatedTechMetricFullNames(v []*string) *UpdateBizMetricRequestUpdateBizMetricCommand {
	s.AssociatedTechMetricFullNames = v
	return s
}

func (s *UpdateBizMetricRequestUpdateBizMetricCommand) SetBizOwnerName(v string) *UpdateBizMetricRequestUpdateBizMetricCommand {
	s.BizOwnerName = &v
	return s
}

func (s *UpdateBizMetricRequestUpdateBizMetricCommand) SetCatalogIds(v []*int64) *UpdateBizMetricRequestUpdateBizMetricCommand {
	s.CatalogIds = v
	return s
}

func (s *UpdateBizMetricRequestUpdateBizMetricCommand) SetCustomAttribute(v []*UpdateBizMetricRequestUpdateBizMetricCommandCustomAttribute) *UpdateBizMetricRequestUpdateBizMetricCommand {
	s.CustomAttribute = v
	return s
}

func (s *UpdateBizMetricRequestUpdateBizMetricCommand) SetDescription(v string) *UpdateBizMetricRequestUpdateBizMetricCommand {
	s.Description = &v
	return s
}

func (s *UpdateBizMetricRequestUpdateBizMetricCommand) SetDisplayName(v string) *UpdateBizMetricRequestUpdateBizMetricCommand {
	s.DisplayName = &v
	return s
}

func (s *UpdateBizMetricRequestUpdateBizMetricCommand) SetLabels(v []*string) *UpdateBizMetricRequestUpdateBizMetricCommand {
	s.Labels = v
	return s
}

func (s *UpdateBizMetricRequestUpdateBizMetricCommand) SetMetricDefinition(v string) *UpdateBizMetricRequestUpdateBizMetricCommand {
	s.MetricDefinition = &v
	return s
}

func (s *UpdateBizMetricRequestUpdateBizMetricCommand) SetMetricRelationDiagramExpression(v string) *UpdateBizMetricRequestUpdateBizMetricCommand {
	s.MetricRelationDiagramExpression = &v
	return s
}

func (s *UpdateBizMetricRequestUpdateBizMetricCommand) SetMetricRelationDiagramSwitchOpen(v bool) *UpdateBizMetricRequestUpdateBizMetricCommand {
	s.MetricRelationDiagramSwitchOpen = &v
	return s
}

func (s *UpdateBizMetricRequestUpdateBizMetricCommand) SetName(v string) *UpdateBizMetricRequestUpdateBizMetricCommand {
	s.Name = &v
	return s
}

func (s *UpdateBizMetricRequestUpdateBizMetricCommand) SetNewName(v string) *UpdateBizMetricRequestUpdateBizMetricCommand {
	s.NewName = &v
	return s
}

func (s *UpdateBizMetricRequestUpdateBizMetricCommand) SetOperateInstructionContent(v string) *UpdateBizMetricRequestUpdateBizMetricCommand {
	s.OperateInstructionContent = &v
	return s
}

func (s *UpdateBizMetricRequestUpdateBizMetricCommand) SetOperateInstructionEnabled(v bool) *UpdateBizMetricRequestUpdateBizMetricCommand {
	s.OperateInstructionEnabled = &v
	return s
}

func (s *UpdateBizMetricRequestUpdateBizMetricCommand) SetRelatedBizMetrics(v []*UpdateBizMetricRequestUpdateBizMetricCommandRelatedBizMetrics) *UpdateBizMetricRequestUpdateBizMetricCommand {
	s.RelatedBizMetrics = v
	return s
}

func (s *UpdateBizMetricRequestUpdateBizMetricCommand) SetViewScope(v *UpdateBizMetricRequestUpdateBizMetricCommandViewScope) *UpdateBizMetricRequestUpdateBizMetricCommand {
	s.ViewScope = v
	return s
}

func (s *UpdateBizMetricRequestUpdateBizMetricCommand) Validate() error {
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

type UpdateBizMetricRequestUpdateBizMetricCommandCustomAttribute struct {
	// Custom attribute code
	//
	// example:
	//
	// CustomAttributeCode
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// List of attribute values. 1. For custom input and single-select dropdown attributes, only the first value in the list is read. 2. For multi-select dropdown attributes, all values in the list are read. 3. For hyperlink attributes, the first value is used as the display text and the second value is used as the link URL.
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s UpdateBizMetricRequestUpdateBizMetricCommandCustomAttribute) String() string {
	return dara.Prettify(s)
}

func (s UpdateBizMetricRequestUpdateBizMetricCommandCustomAttribute) GoString() string {
	return s.String()
}

func (s *UpdateBizMetricRequestUpdateBizMetricCommandCustomAttribute) GetCode() *string {
	return s.Code
}

func (s *UpdateBizMetricRequestUpdateBizMetricCommandCustomAttribute) GetValues() []*string {
	return s.Values
}

func (s *UpdateBizMetricRequestUpdateBizMetricCommandCustomAttribute) SetCode(v string) *UpdateBizMetricRequestUpdateBizMetricCommandCustomAttribute {
	s.Code = &v
	return s
}

func (s *UpdateBizMetricRequestUpdateBizMetricCommandCustomAttribute) SetValues(v []*string) *UpdateBizMetricRequestUpdateBizMetricCommandCustomAttribute {
	s.Values = v
	return s
}

func (s *UpdateBizMetricRequestUpdateBizMetricCommandCustomAttribute) Validate() error {
	return dara.Validate(s)
}

type UpdateBizMetricRequestUpdateBizMetricCommandRelatedBizMetrics struct {
	// Business metric name
	//
	// example:
	//
	// Metric2
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Correlation type. Valid values: POSITIVE (positive correlation), NEGATIVE (negative correlation), and OTHER (other)
	//
	// example:
	//
	// POSITIVE
	RelationType *string `json:"RelationType,omitempty" xml:"RelationType,omitempty"`
}

func (s UpdateBizMetricRequestUpdateBizMetricCommandRelatedBizMetrics) String() string {
	return dara.Prettify(s)
}

func (s UpdateBizMetricRequestUpdateBizMetricCommandRelatedBizMetrics) GoString() string {
	return s.String()
}

func (s *UpdateBizMetricRequestUpdateBizMetricCommandRelatedBizMetrics) GetName() *string {
	return s.Name
}

func (s *UpdateBizMetricRequestUpdateBizMetricCommandRelatedBizMetrics) GetRelationType() *string {
	return s.RelationType
}

func (s *UpdateBizMetricRequestUpdateBizMetricCommandRelatedBizMetrics) SetName(v string) *UpdateBizMetricRequestUpdateBizMetricCommandRelatedBizMetrics {
	s.Name = &v
	return s
}

func (s *UpdateBizMetricRequestUpdateBizMetricCommandRelatedBizMetrics) SetRelationType(v string) *UpdateBizMetricRequestUpdateBizMetricCommandRelatedBizMetrics {
	s.RelationType = &v
	return s
}

func (s *UpdateBizMetricRequestUpdateBizMetricCommandRelatedBizMetrics) Validate() error {
	return dara.Validate(s)
}

type UpdateBizMetricRequestUpdateBizMetricCommandViewScope struct {
	// Visibility scope type. Valid values: ALL_USERS_CAN_VIEW (visible to all users), PART_USERS_CAN_VIEW (visible to specified users), and PART_USERS_CAN_NOT_VIEW (invisible to specified users)
	//
	// example:
	//
	// ALL_USERS_CAN_VIEW
	ScopeType *string `json:"ScopeType,omitempty" xml:"ScopeType,omitempty"`
	// Enter user group names. This parameter is read only when the visibility scope is set to PART_USERS_CAN_VIEW or PART_USERS_CAN_NOT_VIEW
	UserGroupNames []*string `json:"UserGroupNames,omitempty" xml:"UserGroupNames,omitempty" type:"Repeated"`
	// Enter the usernames of individual accounts. This parameter takes effect only when the visibility scope is set to PART_USERS_CAN_VIEW or PART_USERS_CAN_NOT_VIEW
	UserNames []*string `json:"UserNames,omitempty" xml:"UserNames,omitempty" type:"Repeated"`
}

func (s UpdateBizMetricRequestUpdateBizMetricCommandViewScope) String() string {
	return dara.Prettify(s)
}

func (s UpdateBizMetricRequestUpdateBizMetricCommandViewScope) GoString() string {
	return s.String()
}

func (s *UpdateBizMetricRequestUpdateBizMetricCommandViewScope) GetScopeType() *string {
	return s.ScopeType
}

func (s *UpdateBizMetricRequestUpdateBizMetricCommandViewScope) GetUserGroupNames() []*string {
	return s.UserGroupNames
}

func (s *UpdateBizMetricRequestUpdateBizMetricCommandViewScope) GetUserNames() []*string {
	return s.UserNames
}

func (s *UpdateBizMetricRequestUpdateBizMetricCommandViewScope) SetScopeType(v string) *UpdateBizMetricRequestUpdateBizMetricCommandViewScope {
	s.ScopeType = &v
	return s
}

func (s *UpdateBizMetricRequestUpdateBizMetricCommandViewScope) SetUserGroupNames(v []*string) *UpdateBizMetricRequestUpdateBizMetricCommandViewScope {
	s.UserGroupNames = v
	return s
}

func (s *UpdateBizMetricRequestUpdateBizMetricCommandViewScope) SetUserNames(v []*string) *UpdateBizMetricRequestUpdateBizMetricCommandViewScope {
	s.UserNames = v
	return s
}

func (s *UpdateBizMetricRequestUpdateBizMetricCommandViewScope) Validate() error {
	return dara.Validate(s)
}
