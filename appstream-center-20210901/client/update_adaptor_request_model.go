// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAdaptorRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAdaptorId(v string) *UpdateAdaptorRequest
	GetAdaptorId() *string
	SetAppScript(v string) *UpdateAdaptorRequest
	GetAppScript() *string
	SetExecuteFileAddress(v string) *UpdateAdaptorRequest
	GetExecuteFileAddress() *string
	SetOptimizeParam(v string) *UpdateAdaptorRequest
	GetOptimizeParam() *string
	SetSilenceParam(v string) *UpdateAdaptorRequest
	GetSilenceParam() *string
	SetUserConfig(v *UpdateAdaptorRequestUserConfig) *UpdateAdaptorRequest
	GetUserConfig() *UpdateAdaptorRequestUserConfig
}

type UpdateAdaptorRequest struct {
	// The ID of the application adapter to be modified. This parameter is required and must belong to an application adapter manageable by the current account.
	//
	// example:
	//
	// adt-bp1example****
	AdaptorId *string `json:"AdaptorId,omitempty" xml:"AdaptorId,omitempty"`
	// **Not enabled. Do not pass this parameter.*	- This operation does not support modifying application scripts through this parameter.
	//
	// The example "Not applicable (not enabled)" is only a description of the field status and is not a valid parameter value.
	//
	// example:
	//
	// Not applicable (not enabled)
	AppScript *string `json:"AppScript,omitempty" xml:"AppScript,omitempty"`
	// **Not enabled. Do not pass this parameter.*	- This operation does not support modifying executable file paths through this parameter.
	//
	// The example "Not applicable (not enabled)" is only a description of the field status and is not a valid parameter value.
	//
	// example:
	//
	// Not applicable (not enabled)
	ExecuteFileAddress *string `json:"ExecuteFileAddress,omitempty" xml:"ExecuteFileAddress,omitempty"`
	// The application optimization configuration, represented as a JSON object string. Pass this parameter only when specific optimization configuration requirements exist. The content must be a valid JSON object.
	//
	// example:
	//
	// {}
	OptimizeParam *string `json:"OptimizeParam,omitempty" xml:"OptimizeParam,omitempty"`
	// The command-line parameters used when the application starts. If not empty, this parameter is used to update the startup parameters of the application adapter. Specify the value based on the startup requirements of the application.
	//
	// example:
	//
	// --lang=zh-CN
	SilenceParam *string `json:"SilenceParam,omitempty" xml:"SilenceParam,omitempty"`
	// The application runtime configuration object. If this parameter is omitted, the existing user configuration is not modified.
	//
	// **If passed, the original UserConfig is entirely replaced and not merged item by item.*	- Submit the complete configuration and include the configuration items that need to be retained. Subfields that are not provided do not automatically inherit the original values.
	//
	// example:
	//
	// {"FixedResolution":"1440x2560","DisconnectAction":"KEEP_CURRENT","ShowWindow":"HIDE"}
	UserConfig *UpdateAdaptorRequestUserConfig `json:"UserConfig,omitempty" xml:"UserConfig,omitempty" type:"Struct"`
}

func (s UpdateAdaptorRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAdaptorRequest) GoString() string {
	return s.String()
}

func (s *UpdateAdaptorRequest) GetAdaptorId() *string {
	return s.AdaptorId
}

func (s *UpdateAdaptorRequest) GetAppScript() *string {
	return s.AppScript
}

func (s *UpdateAdaptorRequest) GetExecuteFileAddress() *string {
	return s.ExecuteFileAddress
}

func (s *UpdateAdaptorRequest) GetOptimizeParam() *string {
	return s.OptimizeParam
}

func (s *UpdateAdaptorRequest) GetSilenceParam() *string {
	return s.SilenceParam
}

func (s *UpdateAdaptorRequest) GetUserConfig() *UpdateAdaptorRequestUserConfig {
	return s.UserConfig
}

func (s *UpdateAdaptorRequest) SetAdaptorId(v string) *UpdateAdaptorRequest {
	s.AdaptorId = &v
	return s
}

func (s *UpdateAdaptorRequest) SetAppScript(v string) *UpdateAdaptorRequest {
	s.AppScript = &v
	return s
}

func (s *UpdateAdaptorRequest) SetExecuteFileAddress(v string) *UpdateAdaptorRequest {
	s.ExecuteFileAddress = &v
	return s
}

func (s *UpdateAdaptorRequest) SetOptimizeParam(v string) *UpdateAdaptorRequest {
	s.OptimizeParam = &v
	return s
}

func (s *UpdateAdaptorRequest) SetSilenceParam(v string) *UpdateAdaptorRequest {
	s.SilenceParam = &v
	return s
}

func (s *UpdateAdaptorRequest) SetUserConfig(v *UpdateAdaptorRequestUserConfig) *UpdateAdaptorRequest {
	s.UserConfig = v
	return s
}

func (s *UpdateAdaptorRequest) Validate() error {
	if s.UserConfig != nil {
		if err := s.UserConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateAdaptorRequestUserConfig struct {
	// The processing policy after a connection is disconnected. When modifying other UserConfig configurations, if you need to retain the original policy, pass the original value in the complete configuration object. This operation does not automatically populate default values.
	//
	// example:
	//
	// KEEP_CURRENT
	DisconnectAction *string `json:"DisconnectAction,omitempty" xml:"DisconnectAction,omitempty"`
	// The fixed resolution used when the application runs, in the format of width × height, in pixels.
	//
	// example:
	//
	// 1440x2560
	FixedResolution *string `json:"FixedResolution,omitempty" xml:"FixedResolution,omitempty"`
	// The application window display configuration. When modifying other UserConfig configurations, if you need to retain the original configuration, pass the original value in the complete configuration object. This operation does not automatically populate default values.
	//
	// example:
	//
	// HIDE
	ShowWindow *string `json:"ShowWindow,omitempty" xml:"ShowWindow,omitempty"`
}

func (s UpdateAdaptorRequestUserConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateAdaptorRequestUserConfig) GoString() string {
	return s.String()
}

func (s *UpdateAdaptorRequestUserConfig) GetDisconnectAction() *string {
	return s.DisconnectAction
}

func (s *UpdateAdaptorRequestUserConfig) GetFixedResolution() *string {
	return s.FixedResolution
}

func (s *UpdateAdaptorRequestUserConfig) GetShowWindow() *string {
	return s.ShowWindow
}

func (s *UpdateAdaptorRequestUserConfig) SetDisconnectAction(v string) *UpdateAdaptorRequestUserConfig {
	s.DisconnectAction = &v
	return s
}

func (s *UpdateAdaptorRequestUserConfig) SetFixedResolution(v string) *UpdateAdaptorRequestUserConfig {
	s.FixedResolution = &v
	return s
}

func (s *UpdateAdaptorRequestUserConfig) SetShowWindow(v string) *UpdateAdaptorRequestUserConfig {
	s.ShowWindow = &v
	return s
}

func (s *UpdateAdaptorRequestUserConfig) Validate() error {
	return dara.Validate(s)
}
