// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateHookConfigurationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *UpdateHookConfigurationRequest
	GetAppId() *string
	SetGroupId(v string) *UpdateHookConfigurationRequest
	GetGroupId() *string
	SetHooks(v string) *UpdateHookConfigurationRequest
	GetHooks() *string
}

type UpdateHookConfigurationRequest struct {
	// The ID of the application.
	//
	// This parameter is required.
	//
	// example:
	//
	// d498****-1dd8ec229862
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The ID of the application instance group.
	//
	// example:
	//
	// d498****-1dd8ec229862
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The script to mount. Set the value in the JSON format. Example: `[{"ignoreFail":false,"name":"postprepareInstanceEnvironmentOnScaleOut","script":"ls"},{"ignoreFail":true,"name":"postdeleteInstanceDataOnScaleIn","script":""},{"ignoreFail":true,"name":"prestartInstance","script":""},{"ignoreFail":true,"name":"poststartInstance","script":""},{"ignoreFail":true,"name":"prestopInstance","script":""},{"ignoreFail":true,"name":"poststopInstance","script":""}]`
	//
	// This parameter is required.
	//
	// example:
	//
	// [{"ignoreFail":false,"name":"postprepareInstanceEnvironmentOnScaleOut","script":"ls"}]
	Hooks *string `json:"Hooks,omitempty" xml:"Hooks,omitempty"`
}

func (s UpdateHookConfigurationRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateHookConfigurationRequest) GoString() string {
	return s.String()
}

func (s *UpdateHookConfigurationRequest) GetAppId() *string {
	return s.AppId
}

func (s *UpdateHookConfigurationRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *UpdateHookConfigurationRequest) GetHooks() *string {
	return s.Hooks
}

func (s *UpdateHookConfigurationRequest) SetAppId(v string) *UpdateHookConfigurationRequest {
	s.AppId = &v
	return s
}

func (s *UpdateHookConfigurationRequest) SetGroupId(v string) *UpdateHookConfigurationRequest {
	s.GroupId = &v
	return s
}

func (s *UpdateHookConfigurationRequest) SetHooks(v string) *UpdateHookConfigurationRequest {
	s.Hooks = &v
	return s
}

func (s *UpdateHookConfigurationRequest) Validate() error {
	return dara.Validate(s)
}
