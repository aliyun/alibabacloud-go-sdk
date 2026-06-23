// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSecurityStrategyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateSecurityStrategyRequest
	GetClientToken() *string
	SetContent(v *CreateSecurityStrategyRequestContent) *CreateSecurityStrategyRequest
	GetContent() *CreateSecurityStrategyRequestContent
	SetControlDwScope(v string) *CreateSecurityStrategyRequest
	GetControlDwScope() *string
	SetControlModule(v string) *CreateSecurityStrategyRequest
	GetControlModule() *string
	SetControlSubModule(v string) *CreateSecurityStrategyRequest
	GetControlSubModule() *string
	SetDescription(v string) *CreateSecurityStrategyRequest
	GetDescription() *string
	SetName(v string) *CreateSecurityStrategyRequest
	GetName() *string
	SetSchemaName(v string) *CreateSecurityStrategyRequest
	GetSchemaName() *string
	SetWorkspaces(v []*int64) *CreateSecurityStrategyRequest
	GetWorkspaces() []*int64
}

type CreateSecurityStrategyRequest struct {
	// A client-generated token that ensures request idempotency, preventing duplicate operations if you retry the request.
	//
	// example:
	//
	// ABFUOEUOTRTRJKE
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The content of the strategy. This value is constrained by the `SecurityStrategySchema`.
	//
	// This parameter is required.
	Content *CreateSecurityStrategyRequestContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	// **The control scope. Valid values: Workspace and Tenant.**
	//
	// This parameter is required.
	//
	// example:
	//
	// Tenant
	ControlDwScope *string `json:"ControlDwScope,omitempty" xml:"ControlDwScope,omitempty"`
	// **Control module**
	//
	// This parameter is required.
	//
	// example:
	//
	// DataQuery
	ControlModule *string `json:"ControlModule,omitempty" xml:"ControlModule,omitempty"`
	// **Control submodule**
	//
	// example:
	//
	// MyCatalog
	ControlSubModule *string `json:"ControlSubModule,omitempty" xml:"ControlSubModule,omitempty"`
	// **Strategy description**
	//
	// example:
	//
	// 控制数据分析模块的查询结果安全行为
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// **Strategy name**
	//
	// This parameter is required.
	//
	// example:
	//
	// 默认数据分析策略
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// **Schema template name**
	//
	// This parameter is required.
	//
	// example:
	//
	// DataQuerySecurityStrategySchema
	SchemaName *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
	// A list of associated workspace IDs.
	Workspaces []*int64 `json:"Workspaces,omitempty" xml:"Workspaces,omitempty" type:"Repeated"`
}

func (s CreateSecurityStrategyRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSecurityStrategyRequest) GoString() string {
	return s.String()
}

func (s *CreateSecurityStrategyRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateSecurityStrategyRequest) GetContent() *CreateSecurityStrategyRequestContent {
	return s.Content
}

func (s *CreateSecurityStrategyRequest) GetControlDwScope() *string {
	return s.ControlDwScope
}

func (s *CreateSecurityStrategyRequest) GetControlModule() *string {
	return s.ControlModule
}

func (s *CreateSecurityStrategyRequest) GetControlSubModule() *string {
	return s.ControlSubModule
}

func (s *CreateSecurityStrategyRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateSecurityStrategyRequest) GetName() *string {
	return s.Name
}

func (s *CreateSecurityStrategyRequest) GetSchemaName() *string {
	return s.SchemaName
}

func (s *CreateSecurityStrategyRequest) GetWorkspaces() []*int64 {
	return s.Workspaces
}

func (s *CreateSecurityStrategyRequest) SetClientToken(v string) *CreateSecurityStrategyRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateSecurityStrategyRequest) SetContent(v *CreateSecurityStrategyRequestContent) *CreateSecurityStrategyRequest {
	s.Content = v
	return s
}

func (s *CreateSecurityStrategyRequest) SetControlDwScope(v string) *CreateSecurityStrategyRequest {
	s.ControlDwScope = &v
	return s
}

func (s *CreateSecurityStrategyRequest) SetControlModule(v string) *CreateSecurityStrategyRequest {
	s.ControlModule = &v
	return s
}

func (s *CreateSecurityStrategyRequest) SetControlSubModule(v string) *CreateSecurityStrategyRequest {
	s.ControlSubModule = &v
	return s
}

func (s *CreateSecurityStrategyRequest) SetDescription(v string) *CreateSecurityStrategyRequest {
	s.Description = &v
	return s
}

func (s *CreateSecurityStrategyRequest) SetName(v string) *CreateSecurityStrategyRequest {
	s.Name = &v
	return s
}

func (s *CreateSecurityStrategyRequest) SetSchemaName(v string) *CreateSecurityStrategyRequest {
	s.SchemaName = &v
	return s
}

func (s *CreateSecurityStrategyRequest) SetWorkspaces(v []*int64) *CreateSecurityStrategyRequest {
	s.Workspaces = v
	return s
}

func (s *CreateSecurityStrategyRequest) Validate() error {
	if s.Content != nil {
		if err := s.Content.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateSecurityStrategyRequestContent struct {
	// The control scope. This corresponds to the `controlDwScope` property of the `SecurityStrategySchema` associated with the current strategy.
	//
	// example:
	//
	// Tenant
	ControlDwScope *string `json:"ControlDwScope,omitempty" xml:"ControlDwScope,omitempty"`
	// The control module. This corresponds to the `controlModule` property of the `SecurityStrategySchema` associated with the current strategy.
	//
	// This parameter is required.
	//
	// example:
	//
	// DataStudio
	ControlModule *string `json:"ControlModule,omitempty" xml:"ControlModule,omitempty"`
	// The control submodule. This corresponds to the `controlSubModule` property of the `SecurityStrategySchema` associated with the current strategy.
	//
	// example:
	//
	// MyCatalog
	ControlSubModule *string `json:"ControlSubModule,omitempty" xml:"ControlSubModule,omitempty"`
	// A list of controllers.
	//
	// **Note:*	- Valid controllers depend on the selected schema. For more information, see the controller definitions and the list of controllers for each schema.
	Controllers []*CreateSecurityStrategyRequestContentControllers `json:"Controllers,omitempty" xml:"Controllers,omitempty" type:"Repeated"`
	// The `displayName` property of the `SecurityStrategySchema` associated with the current strategy.
	//
	// example:
	//
	// 数据分析
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The `displayNameEn` property of the `SecurityStrategySchema` associated with the current strategy.
	//
	// example:
	//
	// Data Analysis
	DisplayNameEn *string `json:"DisplayNameEn,omitempty" xml:"DisplayNameEn,omitempty"`
	// The `name` property of the `SecurityStrategySchema` associated with the current strategy.
	//
	// This parameter is required.
	//
	// example:
	//
	// DataQuerySecurityStrategySchema
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The `systemPolicyDisplayName` property of the `SecurityStrategySchema` associated with the current strategy.
	//
	// example:
	//
	// Default system generate data query policy
	SystemPolicyDisplayName *string `json:"SystemPolicyDisplayName,omitempty" xml:"SystemPolicyDisplayName,omitempty"`
	// The `systemPolicyName` property of the `SecurityStrategySchema` associated with the current strategy.
	//
	// example:
	//
	// SYSTEM_GENERATE_DEFAULT_DATA_QUERY
	SystemPolicyName *string `json:"SystemPolicyName,omitempty" xml:"SystemPolicyName,omitempty"`
}

func (s CreateSecurityStrategyRequestContent) String() string {
	return dara.Prettify(s)
}

func (s CreateSecurityStrategyRequestContent) GoString() string {
	return s.String()
}

func (s *CreateSecurityStrategyRequestContent) GetControlDwScope() *string {
	return s.ControlDwScope
}

func (s *CreateSecurityStrategyRequestContent) GetControlModule() *string {
	return s.ControlModule
}

func (s *CreateSecurityStrategyRequestContent) GetControlSubModule() *string {
	return s.ControlSubModule
}

func (s *CreateSecurityStrategyRequestContent) GetControllers() []*CreateSecurityStrategyRequestContentControllers {
	return s.Controllers
}

func (s *CreateSecurityStrategyRequestContent) GetDisplayName() *string {
	return s.DisplayName
}

func (s *CreateSecurityStrategyRequestContent) GetDisplayNameEn() *string {
	return s.DisplayNameEn
}

func (s *CreateSecurityStrategyRequestContent) GetName() *string {
	return s.Name
}

func (s *CreateSecurityStrategyRequestContent) GetSystemPolicyDisplayName() *string {
	return s.SystemPolicyDisplayName
}

func (s *CreateSecurityStrategyRequestContent) GetSystemPolicyName() *string {
	return s.SystemPolicyName
}

func (s *CreateSecurityStrategyRequestContent) SetControlDwScope(v string) *CreateSecurityStrategyRequestContent {
	s.ControlDwScope = &v
	return s
}

func (s *CreateSecurityStrategyRequestContent) SetControlModule(v string) *CreateSecurityStrategyRequestContent {
	s.ControlModule = &v
	return s
}

func (s *CreateSecurityStrategyRequestContent) SetControlSubModule(v string) *CreateSecurityStrategyRequestContent {
	s.ControlSubModule = &v
	return s
}

func (s *CreateSecurityStrategyRequestContent) SetControllers(v []*CreateSecurityStrategyRequestContentControllers) *CreateSecurityStrategyRequestContent {
	s.Controllers = v
	return s
}

func (s *CreateSecurityStrategyRequestContent) SetDisplayName(v string) *CreateSecurityStrategyRequestContent {
	s.DisplayName = &v
	return s
}

func (s *CreateSecurityStrategyRequestContent) SetDisplayNameEn(v string) *CreateSecurityStrategyRequestContent {
	s.DisplayNameEn = &v
	return s
}

func (s *CreateSecurityStrategyRequestContent) SetName(v string) *CreateSecurityStrategyRequestContent {
	s.Name = &v
	return s
}

func (s *CreateSecurityStrategyRequestContent) SetSystemPolicyDisplayName(v string) *CreateSecurityStrategyRequestContent {
	s.SystemPolicyDisplayName = &v
	return s
}

func (s *CreateSecurityStrategyRequestContent) SetSystemPolicyName(v string) *CreateSecurityStrategyRequestContent {
	s.SystemPolicyName = &v
	return s
}

func (s *CreateSecurityStrategyRequestContent) Validate() error {
	if s.Controllers != nil {
		for _, item := range s.Controllers {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateSecurityStrategyRequestContentControllers struct {
	// The default value for Basic Edition.
	//
	// example:
	//
	// 0
	BasicEditionDefaultValue interface{} `json:"BasicEditionDefaultValue,omitempty" xml:"BasicEditionDefaultValue,omitempty"`
	// The valid value interval for Basic Edition, in the format `[min, max]`.
	BasicEditionIntervalValue []*int32 `json:"BasicEditionIntervalValue,omitempty" xml:"BasicEditionIntervalValue,omitempty" type:"Repeated"`
	// The controller identifier. For valid values, see the list of controllers for each schema.
	//
	// example:
	//
	// viewCount
	Controller *string `json:"Controller,omitempty" xml:"Controller,omitempty"`
	// The value type. Valid values: `Boolean`, `Integer`, `Long`, and `String`.
	//
	// example:
	//
	// Integer
	ControllerValueType *string `json:"ControllerValueType,omitempty" xml:"ControllerValueType,omitempty"`
	// The display name.
	//
	// example:
	//
	// 查询结果-单次展示记录值上限
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The English display name.
	//
	// example:
	//
	// Query Results - Single Display Record Limit
	DisplayNameEn *string `json:"DisplayNameEn,omitempty" xml:"DisplayNameEn,omitempty"`
	// Specifies whether to enable this controller.
	//
	// example:
	//
	// true
	Enable *bool `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// The default value for Enterprise Edition.
	//
	// example:
	//
	// 500000
	EnterpriseEditionDefaultValue interface{} `json:"EnterpriseEditionDefaultValue,omitempty" xml:"EnterpriseEditionDefaultValue,omitempty"`
	// The valid value interval for Enterprise Edition, in the format `[min, max]`.
	EnterpriseEditionIntervalValue []*int32 `json:"EnterpriseEditionIntervalValue,omitempty" xml:"EnterpriseEditionIntervalValue,omitempty" type:"Repeated"`
	// The default value for Professional Edition.
	//
	// example:
	//
	// 200000
	ProfessionalEditionDefaultValue interface{} `json:"ProfessionalEditionDefaultValue,omitempty" xml:"ProfessionalEditionDefaultValue,omitempty"`
	// The valid value interval for Professional Edition, in the format `[min, max]`.
	ProfessionalEditionIntervalValue []*int32 `json:"ProfessionalEditionIntervalValue,omitempty" xml:"ProfessionalEditionIntervalValue,omitempty" type:"Repeated"`
	// The default value for Standard Edition.
	//
	// example:
	//
	// 100000
	StandardEditionDefaultValue interface{} `json:"StandardEditionDefaultValue,omitempty" xml:"StandardEditionDefaultValue,omitempty"`
	// The valid value interval for Standard Edition, in the format `[min, max]`.
	StandardEditionIntervalValue []*int32 `json:"StandardEditionIntervalValue,omitempty" xml:"StandardEditionIntervalValue,omitempty" type:"Repeated"`
	// The user-configured value. The type of this value depends on the `ControllerValueType` parameter.
	//
	// example:
	//
	// 10
	UserConfigValue interface{} `json:"UserConfigValue,omitempty" xml:"UserConfigValue,omitempty"`
}

func (s CreateSecurityStrategyRequestContentControllers) String() string {
	return dara.Prettify(s)
}

func (s CreateSecurityStrategyRequestContentControllers) GoString() string {
	return s.String()
}

func (s *CreateSecurityStrategyRequestContentControllers) GetBasicEditionDefaultValue() interface{} {
	return s.BasicEditionDefaultValue
}

func (s *CreateSecurityStrategyRequestContentControllers) GetBasicEditionIntervalValue() []*int32 {
	return s.BasicEditionIntervalValue
}

func (s *CreateSecurityStrategyRequestContentControllers) GetController() *string {
	return s.Controller
}

func (s *CreateSecurityStrategyRequestContentControllers) GetControllerValueType() *string {
	return s.ControllerValueType
}

func (s *CreateSecurityStrategyRequestContentControllers) GetDisplayName() *string {
	return s.DisplayName
}

func (s *CreateSecurityStrategyRequestContentControllers) GetDisplayNameEn() *string {
	return s.DisplayNameEn
}

func (s *CreateSecurityStrategyRequestContentControllers) GetEnable() *bool {
	return s.Enable
}

func (s *CreateSecurityStrategyRequestContentControllers) GetEnterpriseEditionDefaultValue() interface{} {
	return s.EnterpriseEditionDefaultValue
}

func (s *CreateSecurityStrategyRequestContentControllers) GetEnterpriseEditionIntervalValue() []*int32 {
	return s.EnterpriseEditionIntervalValue
}

func (s *CreateSecurityStrategyRequestContentControllers) GetProfessionalEditionDefaultValue() interface{} {
	return s.ProfessionalEditionDefaultValue
}

func (s *CreateSecurityStrategyRequestContentControllers) GetProfessionalEditionIntervalValue() []*int32 {
	return s.ProfessionalEditionIntervalValue
}

func (s *CreateSecurityStrategyRequestContentControllers) GetStandardEditionDefaultValue() interface{} {
	return s.StandardEditionDefaultValue
}

func (s *CreateSecurityStrategyRequestContentControllers) GetStandardEditionIntervalValue() []*int32 {
	return s.StandardEditionIntervalValue
}

func (s *CreateSecurityStrategyRequestContentControllers) GetUserConfigValue() interface{} {
	return s.UserConfigValue
}

func (s *CreateSecurityStrategyRequestContentControllers) SetBasicEditionDefaultValue(v interface{}) *CreateSecurityStrategyRequestContentControllers {
	s.BasicEditionDefaultValue = v
	return s
}

func (s *CreateSecurityStrategyRequestContentControllers) SetBasicEditionIntervalValue(v []*int32) *CreateSecurityStrategyRequestContentControllers {
	s.BasicEditionIntervalValue = v
	return s
}

func (s *CreateSecurityStrategyRequestContentControllers) SetController(v string) *CreateSecurityStrategyRequestContentControllers {
	s.Controller = &v
	return s
}

func (s *CreateSecurityStrategyRequestContentControllers) SetControllerValueType(v string) *CreateSecurityStrategyRequestContentControllers {
	s.ControllerValueType = &v
	return s
}

func (s *CreateSecurityStrategyRequestContentControllers) SetDisplayName(v string) *CreateSecurityStrategyRequestContentControllers {
	s.DisplayName = &v
	return s
}

func (s *CreateSecurityStrategyRequestContentControllers) SetDisplayNameEn(v string) *CreateSecurityStrategyRequestContentControllers {
	s.DisplayNameEn = &v
	return s
}

func (s *CreateSecurityStrategyRequestContentControllers) SetEnable(v bool) *CreateSecurityStrategyRequestContentControllers {
	s.Enable = &v
	return s
}

func (s *CreateSecurityStrategyRequestContentControllers) SetEnterpriseEditionDefaultValue(v interface{}) *CreateSecurityStrategyRequestContentControllers {
	s.EnterpriseEditionDefaultValue = v
	return s
}

func (s *CreateSecurityStrategyRequestContentControllers) SetEnterpriseEditionIntervalValue(v []*int32) *CreateSecurityStrategyRequestContentControllers {
	s.EnterpriseEditionIntervalValue = v
	return s
}

func (s *CreateSecurityStrategyRequestContentControllers) SetProfessionalEditionDefaultValue(v interface{}) *CreateSecurityStrategyRequestContentControllers {
	s.ProfessionalEditionDefaultValue = v
	return s
}

func (s *CreateSecurityStrategyRequestContentControllers) SetProfessionalEditionIntervalValue(v []*int32) *CreateSecurityStrategyRequestContentControllers {
	s.ProfessionalEditionIntervalValue = v
	return s
}

func (s *CreateSecurityStrategyRequestContentControllers) SetStandardEditionDefaultValue(v interface{}) *CreateSecurityStrategyRequestContentControllers {
	s.StandardEditionDefaultValue = v
	return s
}

func (s *CreateSecurityStrategyRequestContentControllers) SetStandardEditionIntervalValue(v []*int32) *CreateSecurityStrategyRequestContentControllers {
	s.StandardEditionIntervalValue = v
	return s
}

func (s *CreateSecurityStrategyRequestContentControllers) SetUserConfigValue(v interface{}) *CreateSecurityStrategyRequestContentControllers {
	s.UserConfigValue = v
	return s
}

func (s *CreateSecurityStrategyRequestContentControllers) Validate() error {
	return dara.Validate(s)
}
