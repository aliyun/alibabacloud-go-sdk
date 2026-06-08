// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSecurityStrategyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetSecurityStrategyResponseBody
	GetRequestId() *string
	SetSecurityStrategy(v *GetSecurityStrategyResponseBodySecurityStrategy) *GetSecurityStrategyResponseBody
	GetSecurityStrategy() *GetSecurityStrategyResponseBodySecurityStrategy
}

type GetSecurityStrategyResponseBody struct {
	// example:
	//
	// 0bc5df3a17****903790e8e8a
	RequestId        *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SecurityStrategy *GetSecurityStrategyResponseBodySecurityStrategy `json:"SecurityStrategy,omitempty" xml:"SecurityStrategy,omitempty" type:"Struct"`
}

func (s GetSecurityStrategyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSecurityStrategyResponseBody) GoString() string {
	return s.String()
}

func (s *GetSecurityStrategyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSecurityStrategyResponseBody) GetSecurityStrategy() *GetSecurityStrategyResponseBodySecurityStrategy {
	return s.SecurityStrategy
}

func (s *GetSecurityStrategyResponseBody) SetRequestId(v string) *GetSecurityStrategyResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSecurityStrategyResponseBody) SetSecurityStrategy(v *GetSecurityStrategyResponseBodySecurityStrategy) *GetSecurityStrategyResponseBody {
	s.SecurityStrategy = v
	return s
}

func (s *GetSecurityStrategyResponseBody) Validate() error {
	if s.SecurityStrategy != nil {
		if err := s.SecurityStrategy.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetSecurityStrategyResponseBodySecurityStrategy struct {
	Content *GetSecurityStrategyResponseBodySecurityStrategyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	// example:
	//
	// Tenant
	ControlDwScope *string `json:"ControlDwScope,omitempty" xml:"ControlDwScope,omitempty"`
	// example:
	//
	// DataQuery
	ControlModule *string `json:"ControlModule,omitempty" xml:"ControlModule,omitempty"`
	// example:
	//
	// MyCatalog
	ControlSubModule *string `json:"ControlSubModule,omitempty" xml:"ControlSubModule,omitempty"`
	// example:
	//
	// 2026-05-25T20:46:19
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 203322746501002787
	Creator     *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// true
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// example:
	//
	// 13
	Id   *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 12
	OriginPolicyId *int64 `json:"OriginPolicyId,omitempty" xml:"OriginPolicyId,omitempty"`
	// example:
	//
	// DataQuerySecurityStrategySchema
	SchemaName *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
	// example:
	//
	// 2026-05-25T20:46:19
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// example:
	//
	// 203322746501002787
	Updater    *string  `json:"Updater,omitempty" xml:"Updater,omitempty"`
	Workspaces []*int64 `json:"Workspaces,omitempty" xml:"Workspaces,omitempty" type:"Repeated"`
}

func (s GetSecurityStrategyResponseBodySecurityStrategy) String() string {
	return dara.Prettify(s)
}

func (s GetSecurityStrategyResponseBodySecurityStrategy) GoString() string {
	return s.String()
}

func (s *GetSecurityStrategyResponseBodySecurityStrategy) GetContent() *GetSecurityStrategyResponseBodySecurityStrategyContent {
	return s.Content
}

func (s *GetSecurityStrategyResponseBodySecurityStrategy) GetControlDwScope() *string {
	return s.ControlDwScope
}

func (s *GetSecurityStrategyResponseBodySecurityStrategy) GetControlModule() *string {
	return s.ControlModule
}

func (s *GetSecurityStrategyResponseBodySecurityStrategy) GetControlSubModule() *string {
	return s.ControlSubModule
}

func (s *GetSecurityStrategyResponseBodySecurityStrategy) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetSecurityStrategyResponseBodySecurityStrategy) GetCreator() *string {
	return s.Creator
}

func (s *GetSecurityStrategyResponseBodySecurityStrategy) GetDescription() *string {
	return s.Description
}

func (s *GetSecurityStrategyResponseBodySecurityStrategy) GetEnabled() *bool {
	return s.Enabled
}

func (s *GetSecurityStrategyResponseBodySecurityStrategy) GetId() *int64 {
	return s.Id
}

func (s *GetSecurityStrategyResponseBodySecurityStrategy) GetName() *string {
	return s.Name
}

func (s *GetSecurityStrategyResponseBodySecurityStrategy) GetOriginPolicyId() *int64 {
	return s.OriginPolicyId
}

func (s *GetSecurityStrategyResponseBodySecurityStrategy) GetSchemaName() *string {
	return s.SchemaName
}

func (s *GetSecurityStrategyResponseBodySecurityStrategy) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *GetSecurityStrategyResponseBodySecurityStrategy) GetUpdater() *string {
	return s.Updater
}

func (s *GetSecurityStrategyResponseBodySecurityStrategy) GetWorkspaces() []*int64 {
	return s.Workspaces
}

func (s *GetSecurityStrategyResponseBodySecurityStrategy) SetContent(v *GetSecurityStrategyResponseBodySecurityStrategyContent) *GetSecurityStrategyResponseBodySecurityStrategy {
	s.Content = v
	return s
}

func (s *GetSecurityStrategyResponseBodySecurityStrategy) SetControlDwScope(v string) *GetSecurityStrategyResponseBodySecurityStrategy {
	s.ControlDwScope = &v
	return s
}

func (s *GetSecurityStrategyResponseBodySecurityStrategy) SetControlModule(v string) *GetSecurityStrategyResponseBodySecurityStrategy {
	s.ControlModule = &v
	return s
}

func (s *GetSecurityStrategyResponseBodySecurityStrategy) SetControlSubModule(v string) *GetSecurityStrategyResponseBodySecurityStrategy {
	s.ControlSubModule = &v
	return s
}

func (s *GetSecurityStrategyResponseBodySecurityStrategy) SetCreateTime(v string) *GetSecurityStrategyResponseBodySecurityStrategy {
	s.CreateTime = &v
	return s
}

func (s *GetSecurityStrategyResponseBodySecurityStrategy) SetCreator(v string) *GetSecurityStrategyResponseBodySecurityStrategy {
	s.Creator = &v
	return s
}

func (s *GetSecurityStrategyResponseBodySecurityStrategy) SetDescription(v string) *GetSecurityStrategyResponseBodySecurityStrategy {
	s.Description = &v
	return s
}

func (s *GetSecurityStrategyResponseBodySecurityStrategy) SetEnabled(v bool) *GetSecurityStrategyResponseBodySecurityStrategy {
	s.Enabled = &v
	return s
}

func (s *GetSecurityStrategyResponseBodySecurityStrategy) SetId(v int64) *GetSecurityStrategyResponseBodySecurityStrategy {
	s.Id = &v
	return s
}

func (s *GetSecurityStrategyResponseBodySecurityStrategy) SetName(v string) *GetSecurityStrategyResponseBodySecurityStrategy {
	s.Name = &v
	return s
}

func (s *GetSecurityStrategyResponseBodySecurityStrategy) SetOriginPolicyId(v int64) *GetSecurityStrategyResponseBodySecurityStrategy {
	s.OriginPolicyId = &v
	return s
}

func (s *GetSecurityStrategyResponseBodySecurityStrategy) SetSchemaName(v string) *GetSecurityStrategyResponseBodySecurityStrategy {
	s.SchemaName = &v
	return s
}

func (s *GetSecurityStrategyResponseBodySecurityStrategy) SetUpdateTime(v string) *GetSecurityStrategyResponseBodySecurityStrategy {
	s.UpdateTime = &v
	return s
}

func (s *GetSecurityStrategyResponseBodySecurityStrategy) SetUpdater(v string) *GetSecurityStrategyResponseBodySecurityStrategy {
	s.Updater = &v
	return s
}

func (s *GetSecurityStrategyResponseBodySecurityStrategy) SetWorkspaces(v []*int64) *GetSecurityStrategyResponseBodySecurityStrategy {
	s.Workspaces = v
	return s
}

func (s *GetSecurityStrategyResponseBodySecurityStrategy) Validate() error {
	if s.Content != nil {
		if err := s.Content.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetSecurityStrategyResponseBodySecurityStrategyContent struct {
	// example:
	//
	// Tenant
	ControlDwScope *string `json:"ControlDwScope,omitempty" xml:"ControlDwScope,omitempty"`
	// example:
	//
	// DataQuery
	ControlModule *string `json:"ControlModule,omitempty" xml:"ControlModule,omitempty"`
	// example:
	//
	// MyCatalog
	ControlSubModule *string                                                              `json:"ControlSubModule,omitempty" xml:"ControlSubModule,omitempty"`
	Controllers      []*GetSecurityStrategyResponseBodySecurityStrategyContentControllers `json:"Controllers,omitempty" xml:"Controllers,omitempty" type:"Repeated"`
	DisplayName      *string                                                              `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// example:
	//
	// Data Query
	DisplayNameEn *string `json:"DisplayNameEn,omitempty" xml:"DisplayNameEn,omitempty"`
	// example:
	//
	// DataQuerySecurityStrategySchema
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// Default system generate data query policy
	SystemPolicyDisplayName *string `json:"SystemPolicyDisplayName,omitempty" xml:"SystemPolicyDisplayName,omitempty"`
	// example:
	//
	// SYSTEM_GENERATE_DEFAULT_DATA_QUERY
	SystemPolicyName *string `json:"SystemPolicyName,omitempty" xml:"SystemPolicyName,omitempty"`
}

func (s GetSecurityStrategyResponseBodySecurityStrategyContent) String() string {
	return dara.Prettify(s)
}

func (s GetSecurityStrategyResponseBodySecurityStrategyContent) GoString() string {
	return s.String()
}

func (s *GetSecurityStrategyResponseBodySecurityStrategyContent) GetControlDwScope() *string {
	return s.ControlDwScope
}

func (s *GetSecurityStrategyResponseBodySecurityStrategyContent) GetControlModule() *string {
	return s.ControlModule
}

func (s *GetSecurityStrategyResponseBodySecurityStrategyContent) GetControlSubModule() *string {
	return s.ControlSubModule
}

func (s *GetSecurityStrategyResponseBodySecurityStrategyContent) GetControllers() []*GetSecurityStrategyResponseBodySecurityStrategyContentControllers {
	return s.Controllers
}

func (s *GetSecurityStrategyResponseBodySecurityStrategyContent) GetDisplayName() *string {
	return s.DisplayName
}

func (s *GetSecurityStrategyResponseBodySecurityStrategyContent) GetDisplayNameEn() *string {
	return s.DisplayNameEn
}

func (s *GetSecurityStrategyResponseBodySecurityStrategyContent) GetName() *string {
	return s.Name
}

func (s *GetSecurityStrategyResponseBodySecurityStrategyContent) GetSystemPolicyDisplayName() *string {
	return s.SystemPolicyDisplayName
}

func (s *GetSecurityStrategyResponseBodySecurityStrategyContent) GetSystemPolicyName() *string {
	return s.SystemPolicyName
}

func (s *GetSecurityStrategyResponseBodySecurityStrategyContent) SetControlDwScope(v string) *GetSecurityStrategyResponseBodySecurityStrategyContent {
	s.ControlDwScope = &v
	return s
}

func (s *GetSecurityStrategyResponseBodySecurityStrategyContent) SetControlModule(v string) *GetSecurityStrategyResponseBodySecurityStrategyContent {
	s.ControlModule = &v
	return s
}

func (s *GetSecurityStrategyResponseBodySecurityStrategyContent) SetControlSubModule(v string) *GetSecurityStrategyResponseBodySecurityStrategyContent {
	s.ControlSubModule = &v
	return s
}

func (s *GetSecurityStrategyResponseBodySecurityStrategyContent) SetControllers(v []*GetSecurityStrategyResponseBodySecurityStrategyContentControllers) *GetSecurityStrategyResponseBodySecurityStrategyContent {
	s.Controllers = v
	return s
}

func (s *GetSecurityStrategyResponseBodySecurityStrategyContent) SetDisplayName(v string) *GetSecurityStrategyResponseBodySecurityStrategyContent {
	s.DisplayName = &v
	return s
}

func (s *GetSecurityStrategyResponseBodySecurityStrategyContent) SetDisplayNameEn(v string) *GetSecurityStrategyResponseBodySecurityStrategyContent {
	s.DisplayNameEn = &v
	return s
}

func (s *GetSecurityStrategyResponseBodySecurityStrategyContent) SetName(v string) *GetSecurityStrategyResponseBodySecurityStrategyContent {
	s.Name = &v
	return s
}

func (s *GetSecurityStrategyResponseBodySecurityStrategyContent) SetSystemPolicyDisplayName(v string) *GetSecurityStrategyResponseBodySecurityStrategyContent {
	s.SystemPolicyDisplayName = &v
	return s
}

func (s *GetSecurityStrategyResponseBodySecurityStrategyContent) SetSystemPolicyName(v string) *GetSecurityStrategyResponseBodySecurityStrategyContent {
	s.SystemPolicyName = &v
	return s
}

func (s *GetSecurityStrategyResponseBodySecurityStrategyContent) Validate() error {
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

type GetSecurityStrategyResponseBodySecurityStrategyContentControllers struct {
	// example:
	//
	// 10000
	BasicEditionDefaultValue  interface{} `json:"BasicEditionDefaultValue,omitempty" xml:"BasicEditionDefaultValue,omitempty"`
	BasicEditionIntervalValue []*int32    `json:"BasicEditionIntervalValue,omitempty" xml:"BasicEditionIntervalValue,omitempty" type:"Repeated"`
	// example:
	//
	// viewCount
	Controller *string `json:"Controller,omitempty" xml:"Controller,omitempty"`
	// example:
	//
	// Integer
	ControllerValueType *string `json:"ControllerValueType,omitempty" xml:"ControllerValueType,omitempty"`
	DisplayName         *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// example:
	//
	// Query Results - Single Display Record Limit
	DisplayNameEn *string `json:"DisplayNameEn,omitempty" xml:"DisplayNameEn,omitempty"`
	// example:
	//
	// true
	Enable *bool `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// example:
	//
	// 10000
	EnterpriseEditionDefaultValue  interface{} `json:"EnterpriseEditionDefaultValue,omitempty" xml:"EnterpriseEditionDefaultValue,omitempty"`
	EnterpriseEditionIntervalValue []*int32    `json:"EnterpriseEditionIntervalValue,omitempty" xml:"EnterpriseEditionIntervalValue,omitempty" type:"Repeated"`
	// example:
	//
	// 10000
	ProfessionalEditionDefaultValue  interface{} `json:"ProfessionalEditionDefaultValue,omitempty" xml:"ProfessionalEditionDefaultValue,omitempty"`
	ProfessionalEditionIntervalValue []*int32    `json:"ProfessionalEditionIntervalValue,omitempty" xml:"ProfessionalEditionIntervalValue,omitempty" type:"Repeated"`
	// example:
	//
	// 10000
	StandardEditionDefaultValue  interface{} `json:"StandardEditionDefaultValue,omitempty" xml:"StandardEditionDefaultValue,omitempty"`
	StandardEditionIntervalValue []*int32    `json:"StandardEditionIntervalValue,omitempty" xml:"StandardEditionIntervalValue,omitempty" type:"Repeated"`
	// example:
	//
	// 10
	UserConfigValue interface{} `json:"UserConfigValue,omitempty" xml:"UserConfigValue,omitempty"`
}

func (s GetSecurityStrategyResponseBodySecurityStrategyContentControllers) String() string {
	return dara.Prettify(s)
}

func (s GetSecurityStrategyResponseBodySecurityStrategyContentControllers) GoString() string {
	return s.String()
}

func (s *GetSecurityStrategyResponseBodySecurityStrategyContentControllers) GetBasicEditionDefaultValue() interface{} {
	return s.BasicEditionDefaultValue
}

func (s *GetSecurityStrategyResponseBodySecurityStrategyContentControllers) GetBasicEditionIntervalValue() []*int32 {
	return s.BasicEditionIntervalValue
}

func (s *GetSecurityStrategyResponseBodySecurityStrategyContentControllers) GetController() *string {
	return s.Controller
}

func (s *GetSecurityStrategyResponseBodySecurityStrategyContentControllers) GetControllerValueType() *string {
	return s.ControllerValueType
}

func (s *GetSecurityStrategyResponseBodySecurityStrategyContentControllers) GetDisplayName() *string {
	return s.DisplayName
}

func (s *GetSecurityStrategyResponseBodySecurityStrategyContentControllers) GetDisplayNameEn() *string {
	return s.DisplayNameEn
}

func (s *GetSecurityStrategyResponseBodySecurityStrategyContentControllers) GetEnable() *bool {
	return s.Enable
}

func (s *GetSecurityStrategyResponseBodySecurityStrategyContentControllers) GetEnterpriseEditionDefaultValue() interface{} {
	return s.EnterpriseEditionDefaultValue
}

func (s *GetSecurityStrategyResponseBodySecurityStrategyContentControllers) GetEnterpriseEditionIntervalValue() []*int32 {
	return s.EnterpriseEditionIntervalValue
}

func (s *GetSecurityStrategyResponseBodySecurityStrategyContentControllers) GetProfessionalEditionDefaultValue() interface{} {
	return s.ProfessionalEditionDefaultValue
}

func (s *GetSecurityStrategyResponseBodySecurityStrategyContentControllers) GetProfessionalEditionIntervalValue() []*int32 {
	return s.ProfessionalEditionIntervalValue
}

func (s *GetSecurityStrategyResponseBodySecurityStrategyContentControllers) GetStandardEditionDefaultValue() interface{} {
	return s.StandardEditionDefaultValue
}

func (s *GetSecurityStrategyResponseBodySecurityStrategyContentControllers) GetStandardEditionIntervalValue() []*int32 {
	return s.StandardEditionIntervalValue
}

func (s *GetSecurityStrategyResponseBodySecurityStrategyContentControllers) GetUserConfigValue() interface{} {
	return s.UserConfigValue
}

func (s *GetSecurityStrategyResponseBodySecurityStrategyContentControllers) SetBasicEditionDefaultValue(v interface{}) *GetSecurityStrategyResponseBodySecurityStrategyContentControllers {
	s.BasicEditionDefaultValue = v
	return s
}

func (s *GetSecurityStrategyResponseBodySecurityStrategyContentControllers) SetBasicEditionIntervalValue(v []*int32) *GetSecurityStrategyResponseBodySecurityStrategyContentControllers {
	s.BasicEditionIntervalValue = v
	return s
}

func (s *GetSecurityStrategyResponseBodySecurityStrategyContentControllers) SetController(v string) *GetSecurityStrategyResponseBodySecurityStrategyContentControllers {
	s.Controller = &v
	return s
}

func (s *GetSecurityStrategyResponseBodySecurityStrategyContentControllers) SetControllerValueType(v string) *GetSecurityStrategyResponseBodySecurityStrategyContentControllers {
	s.ControllerValueType = &v
	return s
}

func (s *GetSecurityStrategyResponseBodySecurityStrategyContentControllers) SetDisplayName(v string) *GetSecurityStrategyResponseBodySecurityStrategyContentControllers {
	s.DisplayName = &v
	return s
}

func (s *GetSecurityStrategyResponseBodySecurityStrategyContentControllers) SetDisplayNameEn(v string) *GetSecurityStrategyResponseBodySecurityStrategyContentControllers {
	s.DisplayNameEn = &v
	return s
}

func (s *GetSecurityStrategyResponseBodySecurityStrategyContentControllers) SetEnable(v bool) *GetSecurityStrategyResponseBodySecurityStrategyContentControllers {
	s.Enable = &v
	return s
}

func (s *GetSecurityStrategyResponseBodySecurityStrategyContentControllers) SetEnterpriseEditionDefaultValue(v interface{}) *GetSecurityStrategyResponseBodySecurityStrategyContentControllers {
	s.EnterpriseEditionDefaultValue = v
	return s
}

func (s *GetSecurityStrategyResponseBodySecurityStrategyContentControllers) SetEnterpriseEditionIntervalValue(v []*int32) *GetSecurityStrategyResponseBodySecurityStrategyContentControllers {
	s.EnterpriseEditionIntervalValue = v
	return s
}

func (s *GetSecurityStrategyResponseBodySecurityStrategyContentControllers) SetProfessionalEditionDefaultValue(v interface{}) *GetSecurityStrategyResponseBodySecurityStrategyContentControllers {
	s.ProfessionalEditionDefaultValue = v
	return s
}

func (s *GetSecurityStrategyResponseBodySecurityStrategyContentControllers) SetProfessionalEditionIntervalValue(v []*int32) *GetSecurityStrategyResponseBodySecurityStrategyContentControllers {
	s.ProfessionalEditionIntervalValue = v
	return s
}

func (s *GetSecurityStrategyResponseBodySecurityStrategyContentControllers) SetStandardEditionDefaultValue(v interface{}) *GetSecurityStrategyResponseBodySecurityStrategyContentControllers {
	s.StandardEditionDefaultValue = v
	return s
}

func (s *GetSecurityStrategyResponseBodySecurityStrategyContentControllers) SetStandardEditionIntervalValue(v []*int32) *GetSecurityStrategyResponseBodySecurityStrategyContentControllers {
	s.StandardEditionIntervalValue = v
	return s
}

func (s *GetSecurityStrategyResponseBodySecurityStrategyContentControllers) SetUserConfigValue(v interface{}) *GetSecurityStrategyResponseBodySecurityStrategyContentControllers {
	s.UserConfigValue = v
	return s
}

func (s *GetSecurityStrategyResponseBodySecurityStrategyContentControllers) Validate() error {
	return dara.Validate(s)
}
