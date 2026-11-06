// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAdaptorShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAdaptorId(v string) *UpdateAdaptorShrinkRequest
	GetAdaptorId() *string
	SetAppScript(v string) *UpdateAdaptorShrinkRequest
	GetAppScript() *string
	SetExecuteFileAddress(v string) *UpdateAdaptorShrinkRequest
	GetExecuteFileAddress() *string
	SetOptimizeParam(v string) *UpdateAdaptorShrinkRequest
	GetOptimizeParam() *string
	SetSilenceParam(v string) *UpdateAdaptorShrinkRequest
	GetSilenceParam() *string
	SetUserConfigShrink(v string) *UpdateAdaptorShrinkRequest
	GetUserConfigShrink() *string
}

type UpdateAdaptorShrinkRequest struct {
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
	UserConfigShrink *string `json:"UserConfig,omitempty" xml:"UserConfig,omitempty"`
}

func (s UpdateAdaptorShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAdaptorShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateAdaptorShrinkRequest) GetAdaptorId() *string {
	return s.AdaptorId
}

func (s *UpdateAdaptorShrinkRequest) GetAppScript() *string {
	return s.AppScript
}

func (s *UpdateAdaptorShrinkRequest) GetExecuteFileAddress() *string {
	return s.ExecuteFileAddress
}

func (s *UpdateAdaptorShrinkRequest) GetOptimizeParam() *string {
	return s.OptimizeParam
}

func (s *UpdateAdaptorShrinkRequest) GetSilenceParam() *string {
	return s.SilenceParam
}

func (s *UpdateAdaptorShrinkRequest) GetUserConfigShrink() *string {
	return s.UserConfigShrink
}

func (s *UpdateAdaptorShrinkRequest) SetAdaptorId(v string) *UpdateAdaptorShrinkRequest {
	s.AdaptorId = &v
	return s
}

func (s *UpdateAdaptorShrinkRequest) SetAppScript(v string) *UpdateAdaptorShrinkRequest {
	s.AppScript = &v
	return s
}

func (s *UpdateAdaptorShrinkRequest) SetExecuteFileAddress(v string) *UpdateAdaptorShrinkRequest {
	s.ExecuteFileAddress = &v
	return s
}

func (s *UpdateAdaptorShrinkRequest) SetOptimizeParam(v string) *UpdateAdaptorShrinkRequest {
	s.OptimizeParam = &v
	return s
}

func (s *UpdateAdaptorShrinkRequest) SetSilenceParam(v string) *UpdateAdaptorShrinkRequest {
	s.SilenceParam = &v
	return s
}

func (s *UpdateAdaptorShrinkRequest) SetUserConfigShrink(v string) *UpdateAdaptorShrinkRequest {
	s.UserConfigShrink = &v
	return s
}

func (s *UpdateAdaptorShrinkRequest) Validate() error {
	return dara.Validate(s)
}
