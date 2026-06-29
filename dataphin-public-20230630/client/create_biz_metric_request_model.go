// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBizMetricRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCreateBizMetricCommand(v *CreateBizMetricRequestCreateBizMetricCommand) *CreateBizMetricRequest
	GetCreateBizMetricCommand() *CreateBizMetricRequestCreateBizMetricCommand
	SetOpTenantId(v int64) *CreateBizMetricRequest
	GetOpTenantId() *int64
}

type CreateBizMetricRequest struct {
	// The create request.
	//
	// This parameter is required.
	CreateBizMetricCommand *CreateBizMetricRequestCreateBizMetricCommand `json:"CreateBizMetricCommand,omitempty" xml:"CreateBizMetricCommand,omitempty" type:"Struct"`
	// The tenant ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s CreateBizMetricRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateBizMetricRequest) GoString() string {
	return s.String()
}

func (s *CreateBizMetricRequest) GetCreateBizMetricCommand() *CreateBizMetricRequestCreateBizMetricCommand {
	return s.CreateBizMetricCommand
}

func (s *CreateBizMetricRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *CreateBizMetricRequest) SetCreateBizMetricCommand(v *CreateBizMetricRequestCreateBizMetricCommand) *CreateBizMetricRequest {
	s.CreateBizMetricCommand = v
	return s
}

func (s *CreateBizMetricRequest) SetOpTenantId(v int64) *CreateBizMetricRequest {
	s.OpTenantId = &v
	return s
}

func (s *CreateBizMetricRequest) Validate() error {
	if s.CreateBizMetricCommand != nil {
		if err := s.CreateBizMetricCommand.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateBizMetricRequestCreateBizMetricCommand struct {
	// The list of full names of associated technical metrics. Enter the full name in the format of "OwnerTableFullName.MetricName", where "OwnerTableFullName" equals "AssetSource.OwnerTableName". A technical metric can be associated with only one business metric and cannot be associated repeatedly.
	AssociatedTechMetricFullNames []*string `json:"AssociatedTechMetricFullNames,omitempty" xml:"AssociatedTechMetricFullNames,omitempty" type:"Repeated"`
	// The name of the business owner. Enter the username of the owner account.
	//
	// example:
	//
	// SuperAdmin
	BizOwnerName *string `json:"BizOwnerName,omitempty" xml:"BizOwnerName,omitempty"`
	// The IDs of the folders to which the metric belongs.
	CatalogIds []*int64 `json:"CatalogIds,omitempty" xml:"CatalogIds,omitempty" type:"Repeated"`
	// The custom attributes.
	CustomAttribute []*CreateBizMetricRequestCreateBizMetricCommandCustomAttribute `json:"CustomAttribute,omitempty" xml:"CustomAttribute,omitempty" type:"Repeated"`
	// The description.
	//
	// example:
	//
	// MetricDesc
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The display name.
	//
	// example:
	//
	// MetricDisplayName
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The asset labels.
	Labels []*string `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	// The metric definition. To reference other business metrics, enclose the metric name in square brackets [ ].
	//
	// example:
	//
	// [Metric1]+[Metric2]
	MetricDefinition *string `json:"MetricDefinition,omitempty" xml:"MetricDefinition,omitempty"`
	// This parameter is read only when the metric relationship diagram is enabled. Enter a calculation expression composed of metric names selected from the related business metrics. Supported operators include +, -, *, /, (, ), %, and ∑. Each metric name must be enclosed in square brackets [ ]. If no operator is specified between two metrics, the system automatically inserts a padding placeholder. If no metric relational expression is configured, the metric relationship diagram switch is automatically shutdown.
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
	// This parameter can be set to true only when at least one related business metric exists. Otherwise, the diagram is automatically disabled.
	//
	// example:
	//
	// true
	MetricRelationDiagramSwitchOpen *bool `json:"MetricRelationDiagramSwitchOpen,omitempty" xml:"MetricRelationDiagramSwitchOpen,omitempty"`
	// The name of the business metric. The name must be unique within the tenant. The name can contain letters, digits, and the following special characters: -_/\\·#$^&*()%+=. The name can be up to 256 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// Metric1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The content of the operation instructions. Only text format is supported.
	//
	// example:
	//
	// content
	OperateInstructionContent *string `json:"OperateInstructionContent,omitempty" xml:"OperateInstructionContent,omitempty"`
	// Specifies whether to enable the operation instructions. Valid values:
	//
	// - true: Enabled.
	//
	// - false: Disabled.
	//
	// example:
	//
	// true
	OperateInstructionEnabled *bool `json:"OperateInstructionEnabled,omitempty" xml:"OperateInstructionEnabled,omitempty"`
	// The list of related business metrics. Enter the names of related business metrics and their relationships.
	RelatedBizMetrics []*CreateBizMetricRequestCreateBizMetricCommandRelatedBizMetrics `json:"RelatedBizMetrics,omitempty" xml:"RelatedBizMetrics,omitempty" type:"Repeated"`
	// The visibility scope.
	ViewScope *CreateBizMetricRequestCreateBizMetricCommandViewScope `json:"ViewScope,omitempty" xml:"ViewScope,omitempty" type:"Struct"`
}

func (s CreateBizMetricRequestCreateBizMetricCommand) String() string {
	return dara.Prettify(s)
}

func (s CreateBizMetricRequestCreateBizMetricCommand) GoString() string {
	return s.String()
}

func (s *CreateBizMetricRequestCreateBizMetricCommand) GetAssociatedTechMetricFullNames() []*string {
	return s.AssociatedTechMetricFullNames
}

func (s *CreateBizMetricRequestCreateBizMetricCommand) GetBizOwnerName() *string {
	return s.BizOwnerName
}

func (s *CreateBizMetricRequestCreateBizMetricCommand) GetCatalogIds() []*int64 {
	return s.CatalogIds
}

func (s *CreateBizMetricRequestCreateBizMetricCommand) GetCustomAttribute() []*CreateBizMetricRequestCreateBizMetricCommandCustomAttribute {
	return s.CustomAttribute
}

func (s *CreateBizMetricRequestCreateBizMetricCommand) GetDescription() *string {
	return s.Description
}

func (s *CreateBizMetricRequestCreateBizMetricCommand) GetDisplayName() *string {
	return s.DisplayName
}

func (s *CreateBizMetricRequestCreateBizMetricCommand) GetLabels() []*string {
	return s.Labels
}

func (s *CreateBizMetricRequestCreateBizMetricCommand) GetMetricDefinition() *string {
	return s.MetricDefinition
}

func (s *CreateBizMetricRequestCreateBizMetricCommand) GetMetricRelationDiagramExpression() *string {
	return s.MetricRelationDiagramExpression
}

func (s *CreateBizMetricRequestCreateBizMetricCommand) GetMetricRelationDiagramSwitchOpen() *bool {
	return s.MetricRelationDiagramSwitchOpen
}

func (s *CreateBizMetricRequestCreateBizMetricCommand) GetName() *string {
	return s.Name
}

func (s *CreateBizMetricRequestCreateBizMetricCommand) GetOperateInstructionContent() *string {
	return s.OperateInstructionContent
}

func (s *CreateBizMetricRequestCreateBizMetricCommand) GetOperateInstructionEnabled() *bool {
	return s.OperateInstructionEnabled
}

func (s *CreateBizMetricRequestCreateBizMetricCommand) GetRelatedBizMetrics() []*CreateBizMetricRequestCreateBizMetricCommandRelatedBizMetrics {
	return s.RelatedBizMetrics
}

func (s *CreateBizMetricRequestCreateBizMetricCommand) GetViewScope() *CreateBizMetricRequestCreateBizMetricCommandViewScope {
	return s.ViewScope
}

func (s *CreateBizMetricRequestCreateBizMetricCommand) SetAssociatedTechMetricFullNames(v []*string) *CreateBizMetricRequestCreateBizMetricCommand {
	s.AssociatedTechMetricFullNames = v
	return s
}

func (s *CreateBizMetricRequestCreateBizMetricCommand) SetBizOwnerName(v string) *CreateBizMetricRequestCreateBizMetricCommand {
	s.BizOwnerName = &v
	return s
}

func (s *CreateBizMetricRequestCreateBizMetricCommand) SetCatalogIds(v []*int64) *CreateBizMetricRequestCreateBizMetricCommand {
	s.CatalogIds = v
	return s
}

func (s *CreateBizMetricRequestCreateBizMetricCommand) SetCustomAttribute(v []*CreateBizMetricRequestCreateBizMetricCommandCustomAttribute) *CreateBizMetricRequestCreateBizMetricCommand {
	s.CustomAttribute = v
	return s
}

func (s *CreateBizMetricRequestCreateBizMetricCommand) SetDescription(v string) *CreateBizMetricRequestCreateBizMetricCommand {
	s.Description = &v
	return s
}

func (s *CreateBizMetricRequestCreateBizMetricCommand) SetDisplayName(v string) *CreateBizMetricRequestCreateBizMetricCommand {
	s.DisplayName = &v
	return s
}

func (s *CreateBizMetricRequestCreateBizMetricCommand) SetLabels(v []*string) *CreateBizMetricRequestCreateBizMetricCommand {
	s.Labels = v
	return s
}

func (s *CreateBizMetricRequestCreateBizMetricCommand) SetMetricDefinition(v string) *CreateBizMetricRequestCreateBizMetricCommand {
	s.MetricDefinition = &v
	return s
}

func (s *CreateBizMetricRequestCreateBizMetricCommand) SetMetricRelationDiagramExpression(v string) *CreateBizMetricRequestCreateBizMetricCommand {
	s.MetricRelationDiagramExpression = &v
	return s
}

func (s *CreateBizMetricRequestCreateBizMetricCommand) SetMetricRelationDiagramSwitchOpen(v bool) *CreateBizMetricRequestCreateBizMetricCommand {
	s.MetricRelationDiagramSwitchOpen = &v
	return s
}

func (s *CreateBizMetricRequestCreateBizMetricCommand) SetName(v string) *CreateBizMetricRequestCreateBizMetricCommand {
	s.Name = &v
	return s
}

func (s *CreateBizMetricRequestCreateBizMetricCommand) SetOperateInstructionContent(v string) *CreateBizMetricRequestCreateBizMetricCommand {
	s.OperateInstructionContent = &v
	return s
}

func (s *CreateBizMetricRequestCreateBizMetricCommand) SetOperateInstructionEnabled(v bool) *CreateBizMetricRequestCreateBizMetricCommand {
	s.OperateInstructionEnabled = &v
	return s
}

func (s *CreateBizMetricRequestCreateBizMetricCommand) SetRelatedBizMetrics(v []*CreateBizMetricRequestCreateBizMetricCommandRelatedBizMetrics) *CreateBizMetricRequestCreateBizMetricCommand {
	s.RelatedBizMetrics = v
	return s
}

func (s *CreateBizMetricRequestCreateBizMetricCommand) SetViewScope(v *CreateBizMetricRequestCreateBizMetricCommandViewScope) *CreateBizMetricRequestCreateBizMetricCommand {
	s.ViewScope = v
	return s
}

func (s *CreateBizMetricRequestCreateBizMetricCommand) Validate() error {
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

type CreateBizMetricRequestCreateBizMetricCommandCustomAttribute struct {
	// The code of the custom attribute.
	//
	// example:
	//
	// CustomAttributeCode
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The list of attribute values. 1. For custom input and single-select dropdown attributes, only the first value in the list is read. 2. For multi-select dropdown attributes, all values in the list are read. 3. For hyperlink attributes, the first value is used as the display text and the second value is used as the redirect link.
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s CreateBizMetricRequestCreateBizMetricCommandCustomAttribute) String() string {
	return dara.Prettify(s)
}

func (s CreateBizMetricRequestCreateBizMetricCommandCustomAttribute) GoString() string {
	return s.String()
}

func (s *CreateBizMetricRequestCreateBizMetricCommandCustomAttribute) GetCode() *string {
	return s.Code
}

func (s *CreateBizMetricRequestCreateBizMetricCommandCustomAttribute) GetValues() []*string {
	return s.Values
}

func (s *CreateBizMetricRequestCreateBizMetricCommandCustomAttribute) SetCode(v string) *CreateBizMetricRequestCreateBizMetricCommandCustomAttribute {
	s.Code = &v
	return s
}

func (s *CreateBizMetricRequestCreateBizMetricCommandCustomAttribute) SetValues(v []*string) *CreateBizMetricRequestCreateBizMetricCommandCustomAttribute {
	s.Values = v
	return s
}

func (s *CreateBizMetricRequestCreateBizMetricCommandCustomAttribute) Validate() error {
	return dara.Validate(s)
}

type CreateBizMetricRequestCreateBizMetricCommandRelatedBizMetrics struct {
	// The metric name.
	//
	// example:
	//
	// Metric2
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The relationship type. Valid values:
	//
	// - POSITIVE: positive correlation.
	//
	// - NEGATIVE: negative correlation.
	//
	// - OTHER: other.
	//
	// example:
	//
	// POSITIVE
	RelationType *string `json:"RelationType,omitempty" xml:"RelationType,omitempty"`
}

func (s CreateBizMetricRequestCreateBizMetricCommandRelatedBizMetrics) String() string {
	return dara.Prettify(s)
}

func (s CreateBizMetricRequestCreateBizMetricCommandRelatedBizMetrics) GoString() string {
	return s.String()
}

func (s *CreateBizMetricRequestCreateBizMetricCommandRelatedBizMetrics) GetName() *string {
	return s.Name
}

func (s *CreateBizMetricRequestCreateBizMetricCommandRelatedBizMetrics) GetRelationType() *string {
	return s.RelationType
}

func (s *CreateBizMetricRequestCreateBizMetricCommandRelatedBizMetrics) SetName(v string) *CreateBizMetricRequestCreateBizMetricCommandRelatedBizMetrics {
	s.Name = &v
	return s
}

func (s *CreateBizMetricRequestCreateBizMetricCommandRelatedBizMetrics) SetRelationType(v string) *CreateBizMetricRequestCreateBizMetricCommandRelatedBizMetrics {
	s.RelationType = &v
	return s
}

func (s *CreateBizMetricRequestCreateBizMetricCommandRelatedBizMetrics) Validate() error {
	return dara.Validate(s)
}

type CreateBizMetricRequestCreateBizMetricCommandViewScope struct {
	// The visibility scope type. Valid values:
	//
	// - ALL_USERS_CAN_VIEW: visible to all users.
	//
	// - PART_USERS_CAN_VIEW: visible to specified users.
	//
	// - PART_USERS_CAN_NOT_VIEW: not visible to specified users.
	//
	// example:
	//
	// PART_USERS_CAN_VIEW
	ScopeType *string `json:"ScopeType,omitempty" xml:"ScopeType,omitempty"`
	// The names of user groups. This parameter is read only when the visibility scope is set to PART_USERS_CAN_VIEW or PART_USERS_CAN_NOT_VIEW.
	UserGroupNames []*string `json:"UserGroupNames,omitempty" xml:"UserGroupNames,omitempty" type:"Repeated"`
	// The usernames of individual accounts. This parameter takes effect only when the visibility scope is set to PART_USERS_CAN_VIEW or PART_USERS_CAN_NOT_VIEW.
	UserNames []*string `json:"UserNames,omitempty" xml:"UserNames,omitempty" type:"Repeated"`
}

func (s CreateBizMetricRequestCreateBizMetricCommandViewScope) String() string {
	return dara.Prettify(s)
}

func (s CreateBizMetricRequestCreateBizMetricCommandViewScope) GoString() string {
	return s.String()
}

func (s *CreateBizMetricRequestCreateBizMetricCommandViewScope) GetScopeType() *string {
	return s.ScopeType
}

func (s *CreateBizMetricRequestCreateBizMetricCommandViewScope) GetUserGroupNames() []*string {
	return s.UserGroupNames
}

func (s *CreateBizMetricRequestCreateBizMetricCommandViewScope) GetUserNames() []*string {
	return s.UserNames
}

func (s *CreateBizMetricRequestCreateBizMetricCommandViewScope) SetScopeType(v string) *CreateBizMetricRequestCreateBizMetricCommandViewScope {
	s.ScopeType = &v
	return s
}

func (s *CreateBizMetricRequestCreateBizMetricCommandViewScope) SetUserGroupNames(v []*string) *CreateBizMetricRequestCreateBizMetricCommandViewScope {
	s.UserGroupNames = v
	return s
}

func (s *CreateBizMetricRequestCreateBizMetricCommandViewScope) SetUserNames(v []*string) *CreateBizMetricRequestCreateBizMetricCommandViewScope {
	s.UserNames = v
	return s
}

func (s *CreateBizMetricRequestCreateBizMetricCommandViewScope) Validate() error {
	return dara.Validate(s)
}
