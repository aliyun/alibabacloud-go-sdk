// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFindBestMatchSecurityStrategyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *FindBestMatchSecurityStrategyResponseBodyData) *FindBestMatchSecurityStrategyResponseBody
	GetData() *FindBestMatchSecurityStrategyResponseBodyData
	SetRequestId(v string) *FindBestMatchSecurityStrategyResponseBody
	GetRequestId() *string
}

type FindBestMatchSecurityStrategyResponseBody struct {
	// The data object.
	Data *FindBestMatchSecurityStrategyResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The **API request ID**.
	//
	// example:
	//
	// 0bc5df3a17***903790e8e8a
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s FindBestMatchSecurityStrategyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s FindBestMatchSecurityStrategyResponseBody) GoString() string {
	return s.String()
}

func (s *FindBestMatchSecurityStrategyResponseBody) GetData() *FindBestMatchSecurityStrategyResponseBodyData {
	return s.Data
}

func (s *FindBestMatchSecurityStrategyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *FindBestMatchSecurityStrategyResponseBody) SetData(v *FindBestMatchSecurityStrategyResponseBodyData) *FindBestMatchSecurityStrategyResponseBody {
	s.Data = v
	return s
}

func (s *FindBestMatchSecurityStrategyResponseBody) SetRequestId(v string) *FindBestMatchSecurityStrategyResponseBody {
	s.RequestId = &v
	return s
}

func (s *FindBestMatchSecurityStrategyResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type FindBestMatchSecurityStrategyResponseBodyData struct {
	// The purchased DataWorks edition.
	//
	// example:
	//
	// STANDARD
	Edition *string `json:"Edition,omitempty" xml:"Edition,omitempty"`
	// The name of the purchased DataWorks edition.
	//
	// example:
	//
	// Standard Edition
	EditionDisplayName *string `json:"EditionDisplayName,omitempty" xml:"EditionDisplayName,omitempty"`
	// The security policy.
	SecurityStrategy *FindBestMatchSecurityStrategyResponseBodyDataSecurityStrategy `json:"SecurityStrategy,omitempty" xml:"SecurityStrategy,omitempty" type:"Struct"`
}

func (s FindBestMatchSecurityStrategyResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s FindBestMatchSecurityStrategyResponseBodyData) GoString() string {
	return s.String()
}

func (s *FindBestMatchSecurityStrategyResponseBodyData) GetEdition() *string {
	return s.Edition
}

func (s *FindBestMatchSecurityStrategyResponseBodyData) GetEditionDisplayName() *string {
	return s.EditionDisplayName
}

func (s *FindBestMatchSecurityStrategyResponseBodyData) GetSecurityStrategy() *FindBestMatchSecurityStrategyResponseBodyDataSecurityStrategy {
	return s.SecurityStrategy
}

func (s *FindBestMatchSecurityStrategyResponseBodyData) SetEdition(v string) *FindBestMatchSecurityStrategyResponseBodyData {
	s.Edition = &v
	return s
}

func (s *FindBestMatchSecurityStrategyResponseBodyData) SetEditionDisplayName(v string) *FindBestMatchSecurityStrategyResponseBodyData {
	s.EditionDisplayName = &v
	return s
}

func (s *FindBestMatchSecurityStrategyResponseBodyData) SetSecurityStrategy(v *FindBestMatchSecurityStrategyResponseBodyDataSecurityStrategy) *FindBestMatchSecurityStrategyResponseBodyData {
	s.SecurityStrategy = v
	return s
}

func (s *FindBestMatchSecurityStrategyResponseBodyData) Validate() error {
	if s.SecurityStrategy != nil {
		if err := s.SecurityStrategy.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type FindBestMatchSecurityStrategyResponseBodyDataSecurityStrategy struct {
	// The **policy content**. For more information, see the StrategyContent definition.
	Content *FindBestMatchSecurityStrategyResponseBodyDataSecurityStrategyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	// The **control scope*	- (Workspace/Tenant).
	//
	// example:
	//
	// Tenant
	ControlDwScope *string `json:"ControlDwScope,omitempty" xml:"ControlDwScope,omitempty"`
	// The **control module**.
	//
	// example:
	//
	// DataQuery
	ControlModule *string `json:"ControlModule,omitempty" xml:"ControlModule,omitempty"`
	// The **control submodule**.
	//
	// example:
	//
	// MyCatalog
	ControlSubModule *string `json:"ControlSubModule,omitempty" xml:"ControlSubModule,omitempty"`
	// The creation time in the YYYY-MM-DDTHH:mm:ss format. The time zone is not specified in the current document.
	//
	// example:
	//
	// 2026-05-25T20:46:19
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The user ID of the creator.
	//
	// example:
	//
	// 203322746501002787
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// The **policy description**.
	//
	// example:
	//
	// Controls the security behavior of query results in the Data Analysis module
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Indicates whether the policy is **enabled**.
	//
	// example:
	//
	// true
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The **policy ID**.
	//
	// example:
	//
	// 16
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The **policy name**.
	//
	// example:
	//
	// Default Data Analysis policy
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The source ID of the system default policy.
	//
	// example:
	//
	// 12
	OriginPolicyId *int64 `json:"OriginPolicyId,omitempty" xml:"OriginPolicyId,omitempty"`
	// The **schema template name**.
	//
	// example:
	//
	// DataQuerySecurityStrategySchema
	SchemaName *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
	// The update time in the YYYY-MM-DDTHH:mm:ss format. The time zone is not specified in the current document.
	//
	// example:
	//
	// 2026-05-25T20:46:19
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// The user ID of the last updater.
	//
	// example:
	//
	// 207947397706614297
	Updater *string `json:"Updater,omitempty" xml:"Updater,omitempty"`
	// The **list of associated workspace IDs**.
	Workspaces []*int64 `json:"Workspaces,omitempty" xml:"Workspaces,omitempty" type:"Repeated"`
}

func (s FindBestMatchSecurityStrategyResponseBodyDataSecurityStrategy) String() string {
	return dara.Prettify(s)
}

func (s FindBestMatchSecurityStrategyResponseBodyDataSecurityStrategy) GoString() string {
	return s.String()
}

func (s *FindBestMatchSecurityStrategyResponseBodyDataSecurityStrategy) GetContent() *FindBestMatchSecurityStrategyResponseBodyDataSecurityStrategyContent {
	return s.Content
}

func (s *FindBestMatchSecurityStrategyResponseBodyDataSecurityStrategy) GetControlDwScope() *string {
	return s.ControlDwScope
}

func (s *FindBestMatchSecurityStrategyResponseBodyDataSecurityStrategy) GetControlModule() *string {
	return s.ControlModule
}

func (s *FindBestMatchSecurityStrategyResponseBodyDataSecurityStrategy) GetControlSubModule() *string {
	return s.ControlSubModule
}

func (s *FindBestMatchSecurityStrategyResponseBodyDataSecurityStrategy) GetCreateTime() *string {
	return s.CreateTime
}

func (s *FindBestMatchSecurityStrategyResponseBodyDataSecurityStrategy) GetCreator() *string {
	return s.Creator
}

func (s *FindBestMatchSecurityStrategyResponseBodyDataSecurityStrategy) GetDescription() *string {
	return s.Description
}

func (s *FindBestMatchSecurityStrategyResponseBodyDataSecurityStrategy) GetEnabled() *bool {
	return s.Enabled
}

func (s *FindBestMatchSecurityStrategyResponseBodyDataSecurityStrategy) GetId() *int64 {
	return s.Id
}

func (s *FindBestMatchSecurityStrategyResponseBodyDataSecurityStrategy) GetName() *string {
	return s.Name
}

func (s *FindBestMatchSecurityStrategyResponseBodyDataSecurityStrategy) GetOriginPolicyId() *int64 {
	return s.OriginPolicyId
}

func (s *FindBestMatchSecurityStrategyResponseBodyDataSecurityStrategy) GetSchemaName() *string {
	return s.SchemaName
}

func (s *FindBestMatchSecurityStrategyResponseBodyDataSecurityStrategy) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *FindBestMatchSecurityStrategyResponseBodyDataSecurityStrategy) GetUpdater() *string {
	return s.Updater
}

func (s *FindBestMatchSecurityStrategyResponseBodyDataSecurityStrategy) GetWorkspaces() []*int64 {
	return s.Workspaces
}

func (s *FindBestMatchSecurityStrategyResponseBodyDataSecurityStrategy) SetContent(v *FindBestMatchSecurityStrategyResponseBodyDataSecurityStrategyContent) *FindBestMatchSecurityStrategyResponseBodyDataSecurityStrategy {
	s.Content = v
	return s
}

func (s *FindBestMatchSecurityStrategyResponseBodyDataSecurityStrategy) SetControlDwScope(v string) *FindBestMatchSecurityStrategyResponseBodyDataSecurityStrategy {
	s.ControlDwScope = &v
	return s
}

func (s *FindBestMatchSecurityStrategyResponseBodyDataSecurityStrategy) SetControlModule(v string) *FindBestMatchSecurityStrategyResponseBodyDataSecurityStrategy {
	s.ControlModule = &v
	return s
}

func (s *FindBestMatchSecurityStrategyResponseBodyDataSecurityStrategy) SetControlSubModule(v string) *FindBestMatchSecurityStrategyResponseBodyDataSecurityStrategy {
	s.ControlSubModule = &v
	return s
}

func (s *FindBestMatchSecurityStrategyResponseBodyDataSecurityStrategy) SetCreateTime(v string) *FindBestMatchSecurityStrategyResponseBodyDataSecurityStrategy {
	s.CreateTime = &v
	return s
}

func (s *FindBestMatchSecurityStrategyResponseBodyDataSecurityStrategy) SetCreator(v string) *FindBestMatchSecurityStrategyResponseBodyDataSecurityStrategy {
	s.Creator = &v
	return s
}

func (s *FindBestMatchSecurityStrategyResponseBodyDataSecurityStrategy) SetDescription(v string) *FindBestMatchSecurityStrategyResponseBodyDataSecurityStrategy {
	s.Description = &v
	return s
}

func (s *FindBestMatchSecurityStrategyResponseBodyDataSecurityStrategy) SetEnabled(v bool) *FindBestMatchSecurityStrategyResponseBodyDataSecurityStrategy {
	s.Enabled = &v
	return s
}

func (s *FindBestMatchSecurityStrategyResponseBodyDataSecurityStrategy) SetId(v int64) *FindBestMatchSecurityStrategyResponseBodyDataSecurityStrategy {
	s.Id = &v
	return s
}

func (s *FindBestMatchSecurityStrategyResponseBodyDataSecurityStrategy) SetName(v string) *FindBestMatchSecurityStrategyResponseBodyDataSecurityStrategy {
	s.Name = &v
	return s
}

func (s *FindBestMatchSecurityStrategyResponseBodyDataSecurityStrategy) SetOriginPolicyId(v int64) *FindBestMatchSecurityStrategyResponseBodyDataSecurityStrategy {
	s.OriginPolicyId = &v
	return s
}

func (s *FindBestMatchSecurityStrategyResponseBodyDataSecurityStrategy) SetSchemaName(v string) *FindBestMatchSecurityStrategyResponseBodyDataSecurityStrategy {
	s.SchemaName = &v
	return s
}

func (s *FindBestMatchSecurityStrategyResponseBodyDataSecurityStrategy) SetUpdateTime(v string) *FindBestMatchSecurityStrategyResponseBodyDataSecurityStrategy {
	s.UpdateTime = &v
	return s
}

func (s *FindBestMatchSecurityStrategyResponseBodyDataSecurityStrategy) SetUpdater(v string) *FindBestMatchSecurityStrategyResponseBodyDataSecurityStrategy {
	s.Updater = &v
	return s
}

func (s *FindBestMatchSecurityStrategyResponseBodyDataSecurityStrategy) SetWorkspaces(v []*int64) *FindBestMatchSecurityStrategyResponseBodyDataSecurityStrategy {
	s.Workspaces = v
	return s
}

func (s *FindBestMatchSecurityStrategyResponseBodyDataSecurityStrategy) Validate() error {
	if s.Content != nil {
		if err := s.Content.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type FindBestMatchSecurityStrategyResponseBodyDataSecurityStrategyContent struct {
	// The **control scope**.
	//
	// example:
	//
	// Tenant
	ControlDwScope *string `json:"ControlDwScope,omitempty" xml:"ControlDwScope,omitempty"`
	// The **control module**.
	//
	// example:
	//
	// DataQuery
	ControlModule *string `json:"ControlModule,omitempty" xml:"ControlModule,omitempty"`
	// The **control submodule**.
	//
	// example:
	//
	// MyCatalog
	ControlSubModule *string `json:"ControlSubModule,omitempty" xml:"ControlSubModule,omitempty"`
	// The list of control items.
	//
	// Note: Valid control items depend on the selected schema. For more information, see the Controller definition and the control item list for each schema.
	Controllers []*FindBestMatchSecurityStrategyResponseBodyDataSecurityStrategyContentControllers `json:"Controllers,omitempty" xml:"Controllers,omitempty" type:"Repeated"`
	// The **display name**.
	//
	// example:
	//
	// Data Analysis
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The **English display name**.
	//
	// example:
	//
	// Data Query
	DisplayNameEn *string `json:"DisplayNameEn,omitempty" xml:"DisplayNameEn,omitempty"`
	// The **schema name**.
	//
	// example:
	//
	// DataQuerySecurityStrategySchema
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The display name of the system policy.
	//
	// example:
	//
	// Default system generate data query policy
	SystemPolicyDisplayName *string `json:"SystemPolicyDisplayName,omitempty" xml:"SystemPolicyDisplayName,omitempty"`
	// The system policy name. If this parameter is not empty, automatic creation of a default policy is triggered.
	//
	// example:
	//
	// SYSTEM_GENERATE_DEFAULT_DATA_QUERY
	SystemPolicyName *string `json:"SystemPolicyName,omitempty" xml:"SystemPolicyName,omitempty"`
}

func (s FindBestMatchSecurityStrategyResponseBodyDataSecurityStrategyContent) String() string {
	return dara.Prettify(s)
}

func (s FindBestMatchSecurityStrategyResponseBodyDataSecurityStrategyContent) GoString() string {
	return s.String()
}

func (s *FindBestMatchSecurityStrategyResponseBodyDataSecurityStrategyContent) GetControlDwScope() *string {
	return s.ControlDwScope
}

func (s *FindBestMatchSecurityStrategyResponseBodyDataSecurityStrategyContent) GetControlModule() *string {
	return s.ControlModule
}

func (s *FindBestMatchSecurityStrategyResponseBodyDataSecurityStrategyContent) GetControlSubModule() *string {
	return s.ControlSubModule
}

func (s *FindBestMatchSecurityStrategyResponseBodyDataSecurityStrategyContent) GetControllers() []*FindBestMatchSecurityStrategyResponseBodyDataSecurityStrategyContentControllers {
	return s.Controllers
}

func (s *FindBestMatchSecurityStrategyResponseBodyDataSecurityStrategyContent) GetDisplayName() *string {
	return s.DisplayName
}

func (s *FindBestMatchSecurityStrategyResponseBodyDataSecurityStrategyContent) GetDisplayNameEn() *string {
	return s.DisplayNameEn
}

func (s *FindBestMatchSecurityStrategyResponseBodyDataSecurityStrategyContent) GetName() *string {
	return s.Name
}

func (s *FindBestMatchSecurityStrategyResponseBodyDataSecurityStrategyContent) GetSystemPolicyDisplayName() *string {
	return s.SystemPolicyDisplayName
}

func (s *FindBestMatchSecurityStrategyResponseBodyDataSecurityStrategyContent) GetSystemPolicyName() *string {
	return s.SystemPolicyName
}

func (s *FindBestMatchSecurityStrategyResponseBodyDataSecurityStrategyContent) SetControlDwScope(v string) *FindBestMatchSecurityStrategyResponseBodyDataSecurityStrategyContent {
	s.ControlDwScope = &v
	return s
}

func (s *FindBestMatchSecurityStrategyResponseBodyDataSecurityStrategyContent) SetControlModule(v string) *FindBestMatchSecurityStrategyResponseBodyDataSecurityStrategyContent {
	s.ControlModule = &v
	return s
}

func (s *FindBestMatchSecurityStrategyResponseBodyDataSecurityStrategyContent) SetControlSubModule(v string) *FindBestMatchSecurityStrategyResponseBodyDataSecurityStrategyContent {
	s.ControlSubModule = &v
	return s
}

func (s *FindBestMatchSecurityStrategyResponseBodyDataSecurityStrategyContent) SetControllers(v []*FindBestMatchSecurityStrategyResponseBodyDataSecurityStrategyContentControllers) *FindBestMatchSecurityStrategyResponseBodyDataSecurityStrategyContent {
	s.Controllers = v
	return s
}

func (s *FindBestMatchSecurityStrategyResponseBodyDataSecurityStrategyContent) SetDisplayName(v string) *FindBestMatchSecurityStrategyResponseBodyDataSecurityStrategyContent {
	s.DisplayName = &v
	return s
}

func (s *FindBestMatchSecurityStrategyResponseBodyDataSecurityStrategyContent) SetDisplayNameEn(v string) *FindBestMatchSecurityStrategyResponseBodyDataSecurityStrategyContent {
	s.DisplayNameEn = &v
	return s
}

func (s *FindBestMatchSecurityStrategyResponseBodyDataSecurityStrategyContent) SetName(v string) *FindBestMatchSecurityStrategyResponseBodyDataSecurityStrategyContent {
	s.Name = &v
	return s
}

func (s *FindBestMatchSecurityStrategyResponseBodyDataSecurityStrategyContent) SetSystemPolicyDisplayName(v string) *FindBestMatchSecurityStrategyResponseBodyDataSecurityStrategyContent {
	s.SystemPolicyDisplayName = &v
	return s
}

func (s *FindBestMatchSecurityStrategyResponseBodyDataSecurityStrategyContent) SetSystemPolicyName(v string) *FindBestMatchSecurityStrategyResponseBodyDataSecurityStrategyContent {
	s.SystemPolicyName = &v
	return s
}

func (s *FindBestMatchSecurityStrategyResponseBodyDataSecurityStrategyContent) Validate() error {
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

type FindBestMatchSecurityStrategyResponseBodyDataSecurityStrategyContentControllers struct {
	// The **default value for Basic Edition**.
	//
	// example:
	//
	// 0
	BasicEditionDefaultValue interface{} `json:"BasicEditionDefaultValue,omitempty" xml:"BasicEditionDefaultValue,omitempty"`
	// The valid value range [min, max] for Basic Edition.
	BasicEditionIntervalValue []*int32 `json:"BasicEditionIntervalValue,omitempty" xml:"BasicEditionIntervalValue,omitempty" type:"Repeated"`
	// The identifier of the control item. For valid values, see the control item list for each schema.
	//
	// example:
	//
	// viewCount
	Controller *string `json:"Controller,omitempty" xml:"Controller,omitempty"`
	// The **value type (Boolean/Integer/Long/String)**.
	//
	// example:
	//
	// Integer
	ControllerValueType *string `json:"ControllerValueType,omitempty" xml:"ControllerValueType,omitempty"`
	// The display name.
	//
	// example:
	//
	// Query results - Maximum number of records displayed at a time
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The **English display name**.
	//
	// example:
	//
	// Query Results - Single Display Record Limit
	DisplayNameEn *string `json:"DisplayNameEn,omitempty" xml:"DisplayNameEn,omitempty"`
	// Indicates whether the control item is **enabled**.
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
	// The valid value range [min, max] for Enterprise Edition.
	EnterpriseEditionIntervalValue []*int32 `json:"EnterpriseEditionIntervalValue,omitempty" xml:"EnterpriseEditionIntervalValue,omitempty" type:"Repeated"`
	// The default value for Professional Edition.
	//
	// example:
	//
	// 200000
	ProfessionalEditionDefaultValue interface{} `json:"ProfessionalEditionDefaultValue,omitempty" xml:"ProfessionalEditionDefaultValue,omitempty"`
	// The valid value range [min, max] for Professional Edition.
	ProfessionalEditionIntervalValue []*int32 `json:"ProfessionalEditionIntervalValue,omitempty" xml:"ProfessionalEditionIntervalValue,omitempty" type:"Repeated"`
	// The **default value for Standard Edition**.
	//
	// example:
	//
	// 10000
	StandardEditionDefaultValue interface{} `json:"StandardEditionDefaultValue,omitempty" xml:"StandardEditionDefaultValue,omitempty"`
	// The valid value range [min, max] for Standard Edition.
	StandardEditionIntervalValue []*int32 `json:"StandardEditionIntervalValue,omitempty" xml:"StandardEditionIntervalValue,omitempty" type:"Repeated"`
	// The **user-configured value. The data type depends on the ControllerValueType**.
	//
	// example:
	//
	// 20
	UserConfigValue interface{} `json:"UserConfigValue,omitempty" xml:"UserConfigValue,omitempty"`
}

func (s FindBestMatchSecurityStrategyResponseBodyDataSecurityStrategyContentControllers) String() string {
	return dara.Prettify(s)
}

func (s FindBestMatchSecurityStrategyResponseBodyDataSecurityStrategyContentControllers) GoString() string {
	return s.String()
}

func (s *FindBestMatchSecurityStrategyResponseBodyDataSecurityStrategyContentControllers) GetBasicEditionDefaultValue() interface{} {
	return s.BasicEditionDefaultValue
}

func (s *FindBestMatchSecurityStrategyResponseBodyDataSecurityStrategyContentControllers) GetBasicEditionIntervalValue() []*int32 {
	return s.BasicEditionIntervalValue
}

func (s *FindBestMatchSecurityStrategyResponseBodyDataSecurityStrategyContentControllers) GetController() *string {
	return s.Controller
}

func (s *FindBestMatchSecurityStrategyResponseBodyDataSecurityStrategyContentControllers) GetControllerValueType() *string {
	return s.ControllerValueType
}

func (s *FindBestMatchSecurityStrategyResponseBodyDataSecurityStrategyContentControllers) GetDisplayName() *string {
	return s.DisplayName
}

func (s *FindBestMatchSecurityStrategyResponseBodyDataSecurityStrategyContentControllers) GetDisplayNameEn() *string {
	return s.DisplayNameEn
}

func (s *FindBestMatchSecurityStrategyResponseBodyDataSecurityStrategyContentControllers) GetEnable() *bool {
	return s.Enable
}

func (s *FindBestMatchSecurityStrategyResponseBodyDataSecurityStrategyContentControllers) GetEnterpriseEditionDefaultValue() interface{} {
	return s.EnterpriseEditionDefaultValue
}

func (s *FindBestMatchSecurityStrategyResponseBodyDataSecurityStrategyContentControllers) GetEnterpriseEditionIntervalValue() []*int32 {
	return s.EnterpriseEditionIntervalValue
}

func (s *FindBestMatchSecurityStrategyResponseBodyDataSecurityStrategyContentControllers) GetProfessionalEditionDefaultValue() interface{} {
	return s.ProfessionalEditionDefaultValue
}

func (s *FindBestMatchSecurityStrategyResponseBodyDataSecurityStrategyContentControllers) GetProfessionalEditionIntervalValue() []*int32 {
	return s.ProfessionalEditionIntervalValue
}

func (s *FindBestMatchSecurityStrategyResponseBodyDataSecurityStrategyContentControllers) GetStandardEditionDefaultValue() interface{} {
	return s.StandardEditionDefaultValue
}

func (s *FindBestMatchSecurityStrategyResponseBodyDataSecurityStrategyContentControllers) GetStandardEditionIntervalValue() []*int32 {
	return s.StandardEditionIntervalValue
}

func (s *FindBestMatchSecurityStrategyResponseBodyDataSecurityStrategyContentControllers) GetUserConfigValue() interface{} {
	return s.UserConfigValue
}

func (s *FindBestMatchSecurityStrategyResponseBodyDataSecurityStrategyContentControllers) SetBasicEditionDefaultValue(v interface{}) *FindBestMatchSecurityStrategyResponseBodyDataSecurityStrategyContentControllers {
	s.BasicEditionDefaultValue = v
	return s
}

func (s *FindBestMatchSecurityStrategyResponseBodyDataSecurityStrategyContentControllers) SetBasicEditionIntervalValue(v []*int32) *FindBestMatchSecurityStrategyResponseBodyDataSecurityStrategyContentControllers {
	s.BasicEditionIntervalValue = v
	return s
}

func (s *FindBestMatchSecurityStrategyResponseBodyDataSecurityStrategyContentControllers) SetController(v string) *FindBestMatchSecurityStrategyResponseBodyDataSecurityStrategyContentControllers {
	s.Controller = &v
	return s
}

func (s *FindBestMatchSecurityStrategyResponseBodyDataSecurityStrategyContentControllers) SetControllerValueType(v string) *FindBestMatchSecurityStrategyResponseBodyDataSecurityStrategyContentControllers {
	s.ControllerValueType = &v
	return s
}

func (s *FindBestMatchSecurityStrategyResponseBodyDataSecurityStrategyContentControllers) SetDisplayName(v string) *FindBestMatchSecurityStrategyResponseBodyDataSecurityStrategyContentControllers {
	s.DisplayName = &v
	return s
}

func (s *FindBestMatchSecurityStrategyResponseBodyDataSecurityStrategyContentControllers) SetDisplayNameEn(v string) *FindBestMatchSecurityStrategyResponseBodyDataSecurityStrategyContentControllers {
	s.DisplayNameEn = &v
	return s
}

func (s *FindBestMatchSecurityStrategyResponseBodyDataSecurityStrategyContentControllers) SetEnable(v bool) *FindBestMatchSecurityStrategyResponseBodyDataSecurityStrategyContentControllers {
	s.Enable = &v
	return s
}

func (s *FindBestMatchSecurityStrategyResponseBodyDataSecurityStrategyContentControllers) SetEnterpriseEditionDefaultValue(v interface{}) *FindBestMatchSecurityStrategyResponseBodyDataSecurityStrategyContentControllers {
	s.EnterpriseEditionDefaultValue = v
	return s
}

func (s *FindBestMatchSecurityStrategyResponseBodyDataSecurityStrategyContentControllers) SetEnterpriseEditionIntervalValue(v []*int32) *FindBestMatchSecurityStrategyResponseBodyDataSecurityStrategyContentControllers {
	s.EnterpriseEditionIntervalValue = v
	return s
}

func (s *FindBestMatchSecurityStrategyResponseBodyDataSecurityStrategyContentControllers) SetProfessionalEditionDefaultValue(v interface{}) *FindBestMatchSecurityStrategyResponseBodyDataSecurityStrategyContentControllers {
	s.ProfessionalEditionDefaultValue = v
	return s
}

func (s *FindBestMatchSecurityStrategyResponseBodyDataSecurityStrategyContentControllers) SetProfessionalEditionIntervalValue(v []*int32) *FindBestMatchSecurityStrategyResponseBodyDataSecurityStrategyContentControllers {
	s.ProfessionalEditionIntervalValue = v
	return s
}

func (s *FindBestMatchSecurityStrategyResponseBodyDataSecurityStrategyContentControllers) SetStandardEditionDefaultValue(v interface{}) *FindBestMatchSecurityStrategyResponseBodyDataSecurityStrategyContentControllers {
	s.StandardEditionDefaultValue = v
	return s
}

func (s *FindBestMatchSecurityStrategyResponseBodyDataSecurityStrategyContentControllers) SetStandardEditionIntervalValue(v []*int32) *FindBestMatchSecurityStrategyResponseBodyDataSecurityStrategyContentControllers {
	s.StandardEditionIntervalValue = v
	return s
}

func (s *FindBestMatchSecurityStrategyResponseBodyDataSecurityStrategyContentControllers) SetUserConfigValue(v interface{}) *FindBestMatchSecurityStrategyResponseBodyDataSecurityStrategyContentControllers {
	s.UserConfigValue = v
	return s
}

func (s *FindBestMatchSecurityStrategyResponseBodyDataSecurityStrategyContentControllers) Validate() error {
	return dara.Validate(s)
}
