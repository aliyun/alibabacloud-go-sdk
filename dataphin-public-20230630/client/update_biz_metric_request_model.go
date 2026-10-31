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
	SetOpUserId(v string) *UpdateBizMetricRequest
	GetOpUserId() *string
	SetUpdateBizMetricCommand(v *UpdateBizMetricRequestUpdateBizMetricCommand) *UpdateBizMetricRequest
	GetUpdateBizMetricCommand() *UpdateBizMetricRequestUpdateBizMetricCommand
}

type UpdateBizMetricRequest struct {
	// The tenant ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// The ID of the operator.
	//
	// example:
	//
	// 30001011
	OpUserId *string `json:"OpUserId,omitempty" xml:"OpUserId,omitempty"`
	// The update request.
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

func (s *UpdateBizMetricRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *UpdateBizMetricRequest) GetUpdateBizMetricCommand() *UpdateBizMetricRequestUpdateBizMetricCommand {
	return s.UpdateBizMetricCommand
}

func (s *UpdateBizMetricRequest) SetOpTenantId(v int64) *UpdateBizMetricRequest {
	s.OpTenantId = &v
	return s
}

func (s *UpdateBizMetricRequest) SetOpUserId(v string) *UpdateBizMetricRequest {
	s.OpUserId = &v
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
	// The associated technical metrics. Enter the technical metric in the "Full table name.Metric name" format, where the "Full table name" is equal to "Asset source.Table name". A technical metric can be associated with only one business metric and cannot be repeatedly associated.
	AssociatedTechMetricFullNames []*string `json:"AssociatedTechMetricFullNames,omitempty" xml:"AssociatedTechMetricFullNames,omitempty" type:"Repeated"`
	// The business owner. Enter the username of the owner\\"s account.
	//
	// example:
	//
	// SuperAdmin
	BizOwnerName *string `json:"BizOwnerName,omitempty" xml:"BizOwnerName,omitempty"`
	// The IDs of the parent folders.
	CatalogIds []*int64 `json:"CatalogIds,omitempty" xml:"CatalogIds,omitempty" type:"Repeated"`
	// The custom attributes. Enter the attribute code and attribute value separately.
	CustomAttribute []*UpdateBizMetricRequestUpdateBizMetricCommandCustomAttribute `json:"CustomAttribute,omitempty" xml:"CustomAttribute,omitempty" type:"Repeated"`
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
	// Metric1_DisplayName
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The asset labels.
	Labels []*string `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	// The metric definition. To reference other business metrics, enclose the metric names in half-width square brackets [].
	//
	// example:
	//
	// [Metric2]+[Metric3]
	MetricDefinition *string `json:"MetricDefinition,omitempty" xml:"MetricDefinition,omitempty"`
	// The expression for the metric relationship diagram. This parameter is read only when the metric relationship diagram is enabled. You can only enter a calculation expression composed of metric names selected from the related business metrics. The supported operators are +, -, *, /, (), %, and ∑. Each metric name must be enclosed in half-width square brackets []. If no operator is specified between two metrics, the system automatically adds padding as a placeholder. If the metric relational expression is not configured, a shutdown of the metric relationship diagram switch is automatically triggered.
	//
	// example:
	//
	// [Metric1]+[Metric2]
	MetricRelationDiagramExpression *string `json:"MetricRelationDiagramExpression,omitempty" xml:"MetricRelationDiagramExpression,omitempty"`
	// Specifies whether to enable the metric relationship diagram. Valid values:
	//
	// - true: Enabled.
	//
	// - false: Disabled.
	//
	// This switch can be enabled only when at least one related business metric exists. Otherwise, it is automatically disabled.
	//
	// example:
	//
	// true
	MetricRelationDiagramSwitchOpen *bool `json:"MetricRelationDiagramSwitchOpen,omitempty" xml:"MetricRelationDiagramSwitchOpen,omitempty"`
	// The name of the business metric to update.
	//
	// This parameter is required.
	//
	// example:
	//
	// Metric1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The new name. Specify this parameter when you need to modify the metric name.
	//
	// example:
	//
	// Metric1_new
	NewName *string `json:"NewName,omitempty" xml:"NewName,omitempty"`
	// The content of the usage instructions. Only text format is supported.
	//
	// example:
	//
	// content test
	OperateInstructionContent *string `json:"OperateInstructionContent,omitempty" xml:"OperateInstructionContent,omitempty"`
	// Specifies whether to enable the usage instructions. Valid values:
	//
	// - true: Enabled.
	//
	// - false: Disabled.
	//
	// example:
	//
	// true
	OperateInstructionEnabled *bool `json:"OperateInstructionEnabled,omitempty" xml:"OperateInstructionEnabled,omitempty"`
	// The related business metrics.
	RelatedBizMetrics []*UpdateBizMetricRequestUpdateBizMetricCommandRelatedBizMetrics `json:"RelatedBizMetrics,omitempty" xml:"RelatedBizMetrics,omitempty" type:"Repeated"`
	// The visibility scope.
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
	// The code of the custom attribute.
	//
	// example:
	//
	// CustomAttributeCode
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The attribute values. 1. For custom input and drop-down single-select attributes, only the first attribute value in the list is read. 2. For drop-down multi-select attributes, all attribute values in the list are read. 3. For hyperlink attributes, the first attribute value in the list is read as the display text, and the second attribute value is read as the redirect link.
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
	// The name of the business metric.
	//
	// example:
	//
	// Metric2
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The relationship type. Valid values:
	//
	// - POSITIVE: Positive correlation.
	//
	// - NEGATIVE: Negative correlation.
	//
	// - OTHER: Other.
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
	// The visibility scope type. Valid values:
	//
	// - ALL_USERS_CAN_VIEW: Visible to all users.
	//
	// - PART_USERS_CAN_VIEW: Visible to specific users.
	//
	// - PART_USERS_CAN_NOT_VIEW: Invisible to specific users.
	//
	// example:
	//
	// ALL_USERS_CAN_VIEW
	ScopeType *string `json:"ScopeType,omitempty" xml:"ScopeType,omitempty"`
	// The user group names. This parameter is read only when the visibility scope of the asset is set to Visible to specific users or Invisible to specific users.
	UserGroupNames []*string `json:"UserGroupNames,omitempty" xml:"UserGroupNames,omitempty" type:"Repeated"`
	// The usernames of individual accounts. This parameter is valid only when the visibility scope of the asset is set to Visible to specific users or Invisible to specific users.
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
