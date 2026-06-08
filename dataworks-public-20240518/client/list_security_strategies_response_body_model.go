// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSecurityStrategiesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPagingInfo(v *ListSecurityStrategiesResponseBodyPagingInfo) *ListSecurityStrategiesResponseBody
	GetPagingInfo() *ListSecurityStrategiesResponseBodyPagingInfo
	SetRequestId(v string) *ListSecurityStrategiesResponseBody
	GetRequestId() *string
}

type ListSecurityStrategiesResponseBody struct {
	PagingInfo *ListSecurityStrategiesResponseBodyPagingInfo `json:"PagingInfo,omitempty" xml:"PagingInfo,omitempty" type:"Struct"`
	// example:
	//
	// 0bc5df3a17****903790e8e8a
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListSecurityStrategiesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSecurityStrategiesResponseBody) GoString() string {
	return s.String()
}

func (s *ListSecurityStrategiesResponseBody) GetPagingInfo() *ListSecurityStrategiesResponseBodyPagingInfo {
	return s.PagingInfo
}

func (s *ListSecurityStrategiesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSecurityStrategiesResponseBody) SetPagingInfo(v *ListSecurityStrategiesResponseBodyPagingInfo) *ListSecurityStrategiesResponseBody {
	s.PagingInfo = v
	return s
}

func (s *ListSecurityStrategiesResponseBody) SetRequestId(v string) *ListSecurityStrategiesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSecurityStrategiesResponseBody) Validate() error {
	if s.PagingInfo != nil {
		if err := s.PagingInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListSecurityStrategiesResponseBodyPagingInfo struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize           *int32                                                            `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SecurityStrategies []*ListSecurityStrategiesResponseBodyPagingInfoSecurityStrategies `json:"SecurityStrategies,omitempty" xml:"SecurityStrategies,omitempty" type:"Repeated"`
	// example:
	//
	// 12
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListSecurityStrategiesResponseBodyPagingInfo) String() string {
	return dara.Prettify(s)
}

func (s ListSecurityStrategiesResponseBodyPagingInfo) GoString() string {
	return s.String()
}

func (s *ListSecurityStrategiesResponseBodyPagingInfo) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListSecurityStrategiesResponseBodyPagingInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListSecurityStrategiesResponseBodyPagingInfo) GetSecurityStrategies() []*ListSecurityStrategiesResponseBodyPagingInfoSecurityStrategies {
	return s.SecurityStrategies
}

func (s *ListSecurityStrategiesResponseBodyPagingInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListSecurityStrategiesResponseBodyPagingInfo) SetPageNumber(v int32) *ListSecurityStrategiesResponseBodyPagingInfo {
	s.PageNumber = &v
	return s
}

func (s *ListSecurityStrategiesResponseBodyPagingInfo) SetPageSize(v int32) *ListSecurityStrategiesResponseBodyPagingInfo {
	s.PageSize = &v
	return s
}

func (s *ListSecurityStrategiesResponseBodyPagingInfo) SetSecurityStrategies(v []*ListSecurityStrategiesResponseBodyPagingInfoSecurityStrategies) *ListSecurityStrategiesResponseBodyPagingInfo {
	s.SecurityStrategies = v
	return s
}

func (s *ListSecurityStrategiesResponseBodyPagingInfo) SetTotalCount(v int32) *ListSecurityStrategiesResponseBodyPagingInfo {
	s.TotalCount = &v
	return s
}

func (s *ListSecurityStrategiesResponseBodyPagingInfo) Validate() error {
	if s.SecurityStrategies != nil {
		for _, item := range s.SecurityStrategies {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListSecurityStrategiesResponseBodyPagingInfoSecurityStrategies struct {
	Content *ListSecurityStrategiesResponseBodyPagingInfoSecurityStrategiesContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
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
	// false
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// example:
	//
	// 12345
	Id   *string `json:"Id,omitempty" xml:"Id,omitempty"`
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

func (s ListSecurityStrategiesResponseBodyPagingInfoSecurityStrategies) String() string {
	return dara.Prettify(s)
}

func (s ListSecurityStrategiesResponseBodyPagingInfoSecurityStrategies) GoString() string {
	return s.String()
}

func (s *ListSecurityStrategiesResponseBodyPagingInfoSecurityStrategies) GetContent() *ListSecurityStrategiesResponseBodyPagingInfoSecurityStrategiesContent {
	return s.Content
}

func (s *ListSecurityStrategiesResponseBodyPagingInfoSecurityStrategies) GetControlDwScope() *string {
	return s.ControlDwScope
}

func (s *ListSecurityStrategiesResponseBodyPagingInfoSecurityStrategies) GetControlModule() *string {
	return s.ControlModule
}

func (s *ListSecurityStrategiesResponseBodyPagingInfoSecurityStrategies) GetControlSubModule() *string {
	return s.ControlSubModule
}

func (s *ListSecurityStrategiesResponseBodyPagingInfoSecurityStrategies) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListSecurityStrategiesResponseBodyPagingInfoSecurityStrategies) GetCreator() *string {
	return s.Creator
}

func (s *ListSecurityStrategiesResponseBodyPagingInfoSecurityStrategies) GetDescription() *string {
	return s.Description
}

func (s *ListSecurityStrategiesResponseBodyPagingInfoSecurityStrategies) GetEnabled() *bool {
	return s.Enabled
}

func (s *ListSecurityStrategiesResponseBodyPagingInfoSecurityStrategies) GetId() *string {
	return s.Id
}

func (s *ListSecurityStrategiesResponseBodyPagingInfoSecurityStrategies) GetName() *string {
	return s.Name
}

func (s *ListSecurityStrategiesResponseBodyPagingInfoSecurityStrategies) GetOriginPolicyId() *int64 {
	return s.OriginPolicyId
}

func (s *ListSecurityStrategiesResponseBodyPagingInfoSecurityStrategies) GetSchemaName() *string {
	return s.SchemaName
}

func (s *ListSecurityStrategiesResponseBodyPagingInfoSecurityStrategies) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *ListSecurityStrategiesResponseBodyPagingInfoSecurityStrategies) GetUpdater() *string {
	return s.Updater
}

func (s *ListSecurityStrategiesResponseBodyPagingInfoSecurityStrategies) GetWorkspaces() []*int64 {
	return s.Workspaces
}

func (s *ListSecurityStrategiesResponseBodyPagingInfoSecurityStrategies) SetContent(v *ListSecurityStrategiesResponseBodyPagingInfoSecurityStrategiesContent) *ListSecurityStrategiesResponseBodyPagingInfoSecurityStrategies {
	s.Content = v
	return s
}

func (s *ListSecurityStrategiesResponseBodyPagingInfoSecurityStrategies) SetControlDwScope(v string) *ListSecurityStrategiesResponseBodyPagingInfoSecurityStrategies {
	s.ControlDwScope = &v
	return s
}

func (s *ListSecurityStrategiesResponseBodyPagingInfoSecurityStrategies) SetControlModule(v string) *ListSecurityStrategiesResponseBodyPagingInfoSecurityStrategies {
	s.ControlModule = &v
	return s
}

func (s *ListSecurityStrategiesResponseBodyPagingInfoSecurityStrategies) SetControlSubModule(v string) *ListSecurityStrategiesResponseBodyPagingInfoSecurityStrategies {
	s.ControlSubModule = &v
	return s
}

func (s *ListSecurityStrategiesResponseBodyPagingInfoSecurityStrategies) SetCreateTime(v string) *ListSecurityStrategiesResponseBodyPagingInfoSecurityStrategies {
	s.CreateTime = &v
	return s
}

func (s *ListSecurityStrategiesResponseBodyPagingInfoSecurityStrategies) SetCreator(v string) *ListSecurityStrategiesResponseBodyPagingInfoSecurityStrategies {
	s.Creator = &v
	return s
}

func (s *ListSecurityStrategiesResponseBodyPagingInfoSecurityStrategies) SetDescription(v string) *ListSecurityStrategiesResponseBodyPagingInfoSecurityStrategies {
	s.Description = &v
	return s
}

func (s *ListSecurityStrategiesResponseBodyPagingInfoSecurityStrategies) SetEnabled(v bool) *ListSecurityStrategiesResponseBodyPagingInfoSecurityStrategies {
	s.Enabled = &v
	return s
}

func (s *ListSecurityStrategiesResponseBodyPagingInfoSecurityStrategies) SetId(v string) *ListSecurityStrategiesResponseBodyPagingInfoSecurityStrategies {
	s.Id = &v
	return s
}

func (s *ListSecurityStrategiesResponseBodyPagingInfoSecurityStrategies) SetName(v string) *ListSecurityStrategiesResponseBodyPagingInfoSecurityStrategies {
	s.Name = &v
	return s
}

func (s *ListSecurityStrategiesResponseBodyPagingInfoSecurityStrategies) SetOriginPolicyId(v int64) *ListSecurityStrategiesResponseBodyPagingInfoSecurityStrategies {
	s.OriginPolicyId = &v
	return s
}

func (s *ListSecurityStrategiesResponseBodyPagingInfoSecurityStrategies) SetSchemaName(v string) *ListSecurityStrategiesResponseBodyPagingInfoSecurityStrategies {
	s.SchemaName = &v
	return s
}

func (s *ListSecurityStrategiesResponseBodyPagingInfoSecurityStrategies) SetUpdateTime(v string) *ListSecurityStrategiesResponseBodyPagingInfoSecurityStrategies {
	s.UpdateTime = &v
	return s
}

func (s *ListSecurityStrategiesResponseBodyPagingInfoSecurityStrategies) SetUpdater(v string) *ListSecurityStrategiesResponseBodyPagingInfoSecurityStrategies {
	s.Updater = &v
	return s
}

func (s *ListSecurityStrategiesResponseBodyPagingInfoSecurityStrategies) SetWorkspaces(v []*int64) *ListSecurityStrategiesResponseBodyPagingInfoSecurityStrategies {
	s.Workspaces = v
	return s
}

func (s *ListSecurityStrategiesResponseBodyPagingInfoSecurityStrategies) Validate() error {
	if s.Content != nil {
		if err := s.Content.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListSecurityStrategiesResponseBodyPagingInfoSecurityStrategiesContent struct {
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
	ControlSubModule *string                                                                             `json:"ControlSubModule,omitempty" xml:"ControlSubModule,omitempty"`
	Controllers      []*ListSecurityStrategiesResponseBodyPagingInfoSecurityStrategiesContentControllers `json:"Controllers,omitempty" xml:"Controllers,omitempty" type:"Repeated"`
	DisplayName      *string                                                                             `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
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

func (s ListSecurityStrategiesResponseBodyPagingInfoSecurityStrategiesContent) String() string {
	return dara.Prettify(s)
}

func (s ListSecurityStrategiesResponseBodyPagingInfoSecurityStrategiesContent) GoString() string {
	return s.String()
}

func (s *ListSecurityStrategiesResponseBodyPagingInfoSecurityStrategiesContent) GetControlDwScope() *string {
	return s.ControlDwScope
}

func (s *ListSecurityStrategiesResponseBodyPagingInfoSecurityStrategiesContent) GetControlModule() *string {
	return s.ControlModule
}

func (s *ListSecurityStrategiesResponseBodyPagingInfoSecurityStrategiesContent) GetControlSubModule() *string {
	return s.ControlSubModule
}

func (s *ListSecurityStrategiesResponseBodyPagingInfoSecurityStrategiesContent) GetControllers() []*ListSecurityStrategiesResponseBodyPagingInfoSecurityStrategiesContentControllers {
	return s.Controllers
}

func (s *ListSecurityStrategiesResponseBodyPagingInfoSecurityStrategiesContent) GetDisplayName() *string {
	return s.DisplayName
}

func (s *ListSecurityStrategiesResponseBodyPagingInfoSecurityStrategiesContent) GetDisplayNameEn() *string {
	return s.DisplayNameEn
}

func (s *ListSecurityStrategiesResponseBodyPagingInfoSecurityStrategiesContent) GetName() *string {
	return s.Name
}

func (s *ListSecurityStrategiesResponseBodyPagingInfoSecurityStrategiesContent) GetSystemPolicyDisplayName() *string {
	return s.SystemPolicyDisplayName
}

func (s *ListSecurityStrategiesResponseBodyPagingInfoSecurityStrategiesContent) GetSystemPolicyName() *string {
	return s.SystemPolicyName
}

func (s *ListSecurityStrategiesResponseBodyPagingInfoSecurityStrategiesContent) SetControlDwScope(v string) *ListSecurityStrategiesResponseBodyPagingInfoSecurityStrategiesContent {
	s.ControlDwScope = &v
	return s
}

func (s *ListSecurityStrategiesResponseBodyPagingInfoSecurityStrategiesContent) SetControlModule(v string) *ListSecurityStrategiesResponseBodyPagingInfoSecurityStrategiesContent {
	s.ControlModule = &v
	return s
}

func (s *ListSecurityStrategiesResponseBodyPagingInfoSecurityStrategiesContent) SetControlSubModule(v string) *ListSecurityStrategiesResponseBodyPagingInfoSecurityStrategiesContent {
	s.ControlSubModule = &v
	return s
}

func (s *ListSecurityStrategiesResponseBodyPagingInfoSecurityStrategiesContent) SetControllers(v []*ListSecurityStrategiesResponseBodyPagingInfoSecurityStrategiesContentControllers) *ListSecurityStrategiesResponseBodyPagingInfoSecurityStrategiesContent {
	s.Controllers = v
	return s
}

func (s *ListSecurityStrategiesResponseBodyPagingInfoSecurityStrategiesContent) SetDisplayName(v string) *ListSecurityStrategiesResponseBodyPagingInfoSecurityStrategiesContent {
	s.DisplayName = &v
	return s
}

func (s *ListSecurityStrategiesResponseBodyPagingInfoSecurityStrategiesContent) SetDisplayNameEn(v string) *ListSecurityStrategiesResponseBodyPagingInfoSecurityStrategiesContent {
	s.DisplayNameEn = &v
	return s
}

func (s *ListSecurityStrategiesResponseBodyPagingInfoSecurityStrategiesContent) SetName(v string) *ListSecurityStrategiesResponseBodyPagingInfoSecurityStrategiesContent {
	s.Name = &v
	return s
}

func (s *ListSecurityStrategiesResponseBodyPagingInfoSecurityStrategiesContent) SetSystemPolicyDisplayName(v string) *ListSecurityStrategiesResponseBodyPagingInfoSecurityStrategiesContent {
	s.SystemPolicyDisplayName = &v
	return s
}

func (s *ListSecurityStrategiesResponseBodyPagingInfoSecurityStrategiesContent) SetSystemPolicyName(v string) *ListSecurityStrategiesResponseBodyPagingInfoSecurityStrategiesContent {
	s.SystemPolicyName = &v
	return s
}

func (s *ListSecurityStrategiesResponseBodyPagingInfoSecurityStrategiesContent) Validate() error {
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

type ListSecurityStrategiesResponseBodyPagingInfoSecurityStrategiesContentControllers struct {
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
	Enable        *bool   `json:"Enable,omitempty" xml:"Enable,omitempty"`
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
	// 20
	UserConfigValue interface{} `json:"UserConfigValue,omitempty" xml:"UserConfigValue,omitempty"`
}

func (s ListSecurityStrategiesResponseBodyPagingInfoSecurityStrategiesContentControllers) String() string {
	return dara.Prettify(s)
}

func (s ListSecurityStrategiesResponseBodyPagingInfoSecurityStrategiesContentControllers) GoString() string {
	return s.String()
}

func (s *ListSecurityStrategiesResponseBodyPagingInfoSecurityStrategiesContentControllers) GetBasicEditionDefaultValue() interface{} {
	return s.BasicEditionDefaultValue
}

func (s *ListSecurityStrategiesResponseBodyPagingInfoSecurityStrategiesContentControllers) GetBasicEditionIntervalValue() []*int32 {
	return s.BasicEditionIntervalValue
}

func (s *ListSecurityStrategiesResponseBodyPagingInfoSecurityStrategiesContentControllers) GetController() *string {
	return s.Controller
}

func (s *ListSecurityStrategiesResponseBodyPagingInfoSecurityStrategiesContentControllers) GetControllerValueType() *string {
	return s.ControllerValueType
}

func (s *ListSecurityStrategiesResponseBodyPagingInfoSecurityStrategiesContentControllers) GetDisplayName() *string {
	return s.DisplayName
}

func (s *ListSecurityStrategiesResponseBodyPagingInfoSecurityStrategiesContentControllers) GetDisplayNameEn() *string {
	return s.DisplayNameEn
}

func (s *ListSecurityStrategiesResponseBodyPagingInfoSecurityStrategiesContentControllers) GetEnable() *bool {
	return s.Enable
}

func (s *ListSecurityStrategiesResponseBodyPagingInfoSecurityStrategiesContentControllers) GetEnterpriseEditionDefaultValue() interface{} {
	return s.EnterpriseEditionDefaultValue
}

func (s *ListSecurityStrategiesResponseBodyPagingInfoSecurityStrategiesContentControllers) GetEnterpriseEditionIntervalValue() []*int32 {
	return s.EnterpriseEditionIntervalValue
}

func (s *ListSecurityStrategiesResponseBodyPagingInfoSecurityStrategiesContentControllers) GetProfessionalEditionDefaultValue() interface{} {
	return s.ProfessionalEditionDefaultValue
}

func (s *ListSecurityStrategiesResponseBodyPagingInfoSecurityStrategiesContentControllers) GetProfessionalEditionIntervalValue() []*int32 {
	return s.ProfessionalEditionIntervalValue
}

func (s *ListSecurityStrategiesResponseBodyPagingInfoSecurityStrategiesContentControllers) GetStandardEditionDefaultValue() interface{} {
	return s.StandardEditionDefaultValue
}

func (s *ListSecurityStrategiesResponseBodyPagingInfoSecurityStrategiesContentControllers) GetStandardEditionIntervalValue() []*int32 {
	return s.StandardEditionIntervalValue
}

func (s *ListSecurityStrategiesResponseBodyPagingInfoSecurityStrategiesContentControllers) GetUserConfigValue() interface{} {
	return s.UserConfigValue
}

func (s *ListSecurityStrategiesResponseBodyPagingInfoSecurityStrategiesContentControllers) SetBasicEditionDefaultValue(v interface{}) *ListSecurityStrategiesResponseBodyPagingInfoSecurityStrategiesContentControllers {
	s.BasicEditionDefaultValue = v
	return s
}

func (s *ListSecurityStrategiesResponseBodyPagingInfoSecurityStrategiesContentControllers) SetBasicEditionIntervalValue(v []*int32) *ListSecurityStrategiesResponseBodyPagingInfoSecurityStrategiesContentControllers {
	s.BasicEditionIntervalValue = v
	return s
}

func (s *ListSecurityStrategiesResponseBodyPagingInfoSecurityStrategiesContentControllers) SetController(v string) *ListSecurityStrategiesResponseBodyPagingInfoSecurityStrategiesContentControllers {
	s.Controller = &v
	return s
}

func (s *ListSecurityStrategiesResponseBodyPagingInfoSecurityStrategiesContentControllers) SetControllerValueType(v string) *ListSecurityStrategiesResponseBodyPagingInfoSecurityStrategiesContentControllers {
	s.ControllerValueType = &v
	return s
}

func (s *ListSecurityStrategiesResponseBodyPagingInfoSecurityStrategiesContentControllers) SetDisplayName(v string) *ListSecurityStrategiesResponseBodyPagingInfoSecurityStrategiesContentControllers {
	s.DisplayName = &v
	return s
}

func (s *ListSecurityStrategiesResponseBodyPagingInfoSecurityStrategiesContentControllers) SetDisplayNameEn(v string) *ListSecurityStrategiesResponseBodyPagingInfoSecurityStrategiesContentControllers {
	s.DisplayNameEn = &v
	return s
}

func (s *ListSecurityStrategiesResponseBodyPagingInfoSecurityStrategiesContentControllers) SetEnable(v bool) *ListSecurityStrategiesResponseBodyPagingInfoSecurityStrategiesContentControllers {
	s.Enable = &v
	return s
}

func (s *ListSecurityStrategiesResponseBodyPagingInfoSecurityStrategiesContentControllers) SetEnterpriseEditionDefaultValue(v interface{}) *ListSecurityStrategiesResponseBodyPagingInfoSecurityStrategiesContentControllers {
	s.EnterpriseEditionDefaultValue = v
	return s
}

func (s *ListSecurityStrategiesResponseBodyPagingInfoSecurityStrategiesContentControllers) SetEnterpriseEditionIntervalValue(v []*int32) *ListSecurityStrategiesResponseBodyPagingInfoSecurityStrategiesContentControllers {
	s.EnterpriseEditionIntervalValue = v
	return s
}

func (s *ListSecurityStrategiesResponseBodyPagingInfoSecurityStrategiesContentControllers) SetProfessionalEditionDefaultValue(v interface{}) *ListSecurityStrategiesResponseBodyPagingInfoSecurityStrategiesContentControllers {
	s.ProfessionalEditionDefaultValue = v
	return s
}

func (s *ListSecurityStrategiesResponseBodyPagingInfoSecurityStrategiesContentControllers) SetProfessionalEditionIntervalValue(v []*int32) *ListSecurityStrategiesResponseBodyPagingInfoSecurityStrategiesContentControllers {
	s.ProfessionalEditionIntervalValue = v
	return s
}

func (s *ListSecurityStrategiesResponseBodyPagingInfoSecurityStrategiesContentControllers) SetStandardEditionDefaultValue(v interface{}) *ListSecurityStrategiesResponseBodyPagingInfoSecurityStrategiesContentControllers {
	s.StandardEditionDefaultValue = v
	return s
}

func (s *ListSecurityStrategiesResponseBodyPagingInfoSecurityStrategiesContentControllers) SetStandardEditionIntervalValue(v []*int32) *ListSecurityStrategiesResponseBodyPagingInfoSecurityStrategiesContentControllers {
	s.StandardEditionIntervalValue = v
	return s
}

func (s *ListSecurityStrategiesResponseBodyPagingInfoSecurityStrategiesContentControllers) SetUserConfigValue(v interface{}) *ListSecurityStrategiesResponseBodyPagingInfoSecurityStrategiesContentControllers {
	s.UserConfigValue = v
	return s
}

func (s *ListSecurityStrategiesResponseBodyPagingInfoSecurityStrategiesContentControllers) Validate() error {
	return dara.Validate(s)
}
