// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSecurityStrategyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *UpdateSecurityStrategyRequest
	GetClientToken() *string
	SetContent(v *UpdateSecurityStrategyRequestContent) *UpdateSecurityStrategyRequest
	GetContent() *UpdateSecurityStrategyRequestContent
	SetDescription(v string) *UpdateSecurityStrategyRequest
	GetDescription() *string
	SetId(v int64) *UpdateSecurityStrategyRequest
	GetId() *int64
	SetName(v string) *UpdateSecurityStrategyRequest
	GetName() *string
	SetWorkspaces(v []*int64) *UpdateSecurityStrategyRequest
	GetWorkspaces() []*int64
}

type UpdateSecurityStrategyRequest struct {
	// example:
	//
	// 1AFAE64E-D1BE-432B-A9*****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	Content     *UpdateSecurityStrategyRequestContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	Description *string                               `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 13
	Id         *int64   `json:"Id,omitempty" xml:"Id,omitempty"`
	Name       *string  `json:"Name,omitempty" xml:"Name,omitempty"`
	Workspaces []*int64 `json:"Workspaces,omitempty" xml:"Workspaces,omitempty" type:"Repeated"`
}

func (s UpdateSecurityStrategyRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateSecurityStrategyRequest) GoString() string {
	return s.String()
}

func (s *UpdateSecurityStrategyRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateSecurityStrategyRequest) GetContent() *UpdateSecurityStrategyRequestContent {
	return s.Content
}

func (s *UpdateSecurityStrategyRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateSecurityStrategyRequest) GetId() *int64 {
	return s.Id
}

func (s *UpdateSecurityStrategyRequest) GetName() *string {
	return s.Name
}

func (s *UpdateSecurityStrategyRequest) GetWorkspaces() []*int64 {
	return s.Workspaces
}

func (s *UpdateSecurityStrategyRequest) SetClientToken(v string) *UpdateSecurityStrategyRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateSecurityStrategyRequest) SetContent(v *UpdateSecurityStrategyRequestContent) *UpdateSecurityStrategyRequest {
	s.Content = v
	return s
}

func (s *UpdateSecurityStrategyRequest) SetDescription(v string) *UpdateSecurityStrategyRequest {
	s.Description = &v
	return s
}

func (s *UpdateSecurityStrategyRequest) SetId(v int64) *UpdateSecurityStrategyRequest {
	s.Id = &v
	return s
}

func (s *UpdateSecurityStrategyRequest) SetName(v string) *UpdateSecurityStrategyRequest {
	s.Name = &v
	return s
}

func (s *UpdateSecurityStrategyRequest) SetWorkspaces(v []*int64) *UpdateSecurityStrategyRequest {
	s.Workspaces = v
	return s
}

func (s *UpdateSecurityStrategyRequest) Validate() error {
	if s.Content != nil {
		if err := s.Content.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateSecurityStrategyRequestContent struct {
	// This parameter is required.
	Controllers []*UpdateSecurityStrategyRequestContentControllers `json:"Controllers,omitempty" xml:"Controllers,omitempty" type:"Repeated"`
}

func (s UpdateSecurityStrategyRequestContent) String() string {
	return dara.Prettify(s)
}

func (s UpdateSecurityStrategyRequestContent) GoString() string {
	return s.String()
}

func (s *UpdateSecurityStrategyRequestContent) GetControllers() []*UpdateSecurityStrategyRequestContentControllers {
	return s.Controllers
}

func (s *UpdateSecurityStrategyRequestContent) SetControllers(v []*UpdateSecurityStrategyRequestContentControllers) *UpdateSecurityStrategyRequestContent {
	s.Controllers = v
	return s
}

func (s *UpdateSecurityStrategyRequestContent) Validate() error {
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

type UpdateSecurityStrategyRequestContentControllers struct {
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
	// 20
	UserConfigValue interface{} `json:"UserConfigValue,omitempty" xml:"UserConfigValue,omitempty"`
}

func (s UpdateSecurityStrategyRequestContentControllers) String() string {
	return dara.Prettify(s)
}

func (s UpdateSecurityStrategyRequestContentControllers) GoString() string {
	return s.String()
}

func (s *UpdateSecurityStrategyRequestContentControllers) GetBasicEditionDefaultValue() interface{} {
	return s.BasicEditionDefaultValue
}

func (s *UpdateSecurityStrategyRequestContentControllers) GetBasicEditionIntervalValue() []*int32 {
	return s.BasicEditionIntervalValue
}

func (s *UpdateSecurityStrategyRequestContentControllers) GetController() *string {
	return s.Controller
}

func (s *UpdateSecurityStrategyRequestContentControllers) GetControllerValueType() *string {
	return s.ControllerValueType
}

func (s *UpdateSecurityStrategyRequestContentControllers) GetDisplayName() *string {
	return s.DisplayName
}

func (s *UpdateSecurityStrategyRequestContentControllers) GetDisplayNameEn() *string {
	return s.DisplayNameEn
}

func (s *UpdateSecurityStrategyRequestContentControllers) GetEnable() *bool {
	return s.Enable
}

func (s *UpdateSecurityStrategyRequestContentControllers) GetEnterpriseEditionDefaultValue() interface{} {
	return s.EnterpriseEditionDefaultValue
}

func (s *UpdateSecurityStrategyRequestContentControllers) GetEnterpriseEditionIntervalValue() []*int32 {
	return s.EnterpriseEditionIntervalValue
}

func (s *UpdateSecurityStrategyRequestContentControllers) GetProfessionalEditionDefaultValue() interface{} {
	return s.ProfessionalEditionDefaultValue
}

func (s *UpdateSecurityStrategyRequestContentControllers) GetProfessionalEditionIntervalValue() []*int32 {
	return s.ProfessionalEditionIntervalValue
}

func (s *UpdateSecurityStrategyRequestContentControllers) GetStandardEditionDefaultValue() interface{} {
	return s.StandardEditionDefaultValue
}

func (s *UpdateSecurityStrategyRequestContentControllers) GetStandardEditionIntervalValue() []*int32 {
	return s.StandardEditionIntervalValue
}

func (s *UpdateSecurityStrategyRequestContentControllers) GetUserConfigValue() interface{} {
	return s.UserConfigValue
}

func (s *UpdateSecurityStrategyRequestContentControllers) SetBasicEditionDefaultValue(v interface{}) *UpdateSecurityStrategyRequestContentControllers {
	s.BasicEditionDefaultValue = v
	return s
}

func (s *UpdateSecurityStrategyRequestContentControllers) SetBasicEditionIntervalValue(v []*int32) *UpdateSecurityStrategyRequestContentControllers {
	s.BasicEditionIntervalValue = v
	return s
}

func (s *UpdateSecurityStrategyRequestContentControllers) SetController(v string) *UpdateSecurityStrategyRequestContentControllers {
	s.Controller = &v
	return s
}

func (s *UpdateSecurityStrategyRequestContentControllers) SetControllerValueType(v string) *UpdateSecurityStrategyRequestContentControllers {
	s.ControllerValueType = &v
	return s
}

func (s *UpdateSecurityStrategyRequestContentControllers) SetDisplayName(v string) *UpdateSecurityStrategyRequestContentControllers {
	s.DisplayName = &v
	return s
}

func (s *UpdateSecurityStrategyRequestContentControllers) SetDisplayNameEn(v string) *UpdateSecurityStrategyRequestContentControllers {
	s.DisplayNameEn = &v
	return s
}

func (s *UpdateSecurityStrategyRequestContentControllers) SetEnable(v bool) *UpdateSecurityStrategyRequestContentControllers {
	s.Enable = &v
	return s
}

func (s *UpdateSecurityStrategyRequestContentControllers) SetEnterpriseEditionDefaultValue(v interface{}) *UpdateSecurityStrategyRequestContentControllers {
	s.EnterpriseEditionDefaultValue = v
	return s
}

func (s *UpdateSecurityStrategyRequestContentControllers) SetEnterpriseEditionIntervalValue(v []*int32) *UpdateSecurityStrategyRequestContentControllers {
	s.EnterpriseEditionIntervalValue = v
	return s
}

func (s *UpdateSecurityStrategyRequestContentControllers) SetProfessionalEditionDefaultValue(v interface{}) *UpdateSecurityStrategyRequestContentControllers {
	s.ProfessionalEditionDefaultValue = v
	return s
}

func (s *UpdateSecurityStrategyRequestContentControllers) SetProfessionalEditionIntervalValue(v []*int32) *UpdateSecurityStrategyRequestContentControllers {
	s.ProfessionalEditionIntervalValue = v
	return s
}

func (s *UpdateSecurityStrategyRequestContentControllers) SetStandardEditionDefaultValue(v interface{}) *UpdateSecurityStrategyRequestContentControllers {
	s.StandardEditionDefaultValue = v
	return s
}

func (s *UpdateSecurityStrategyRequestContentControllers) SetStandardEditionIntervalValue(v []*int32) *UpdateSecurityStrategyRequestContentControllers {
	s.StandardEditionIntervalValue = v
	return s
}

func (s *UpdateSecurityStrategyRequestContentControllers) SetUserConfigValue(v interface{}) *UpdateSecurityStrategyRequestContentControllers {
	s.UserConfigValue = v
	return s
}

func (s *UpdateSecurityStrategyRequestContentControllers) Validate() error {
	return dara.Validate(s)
}
