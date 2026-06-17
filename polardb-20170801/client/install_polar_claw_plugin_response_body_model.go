// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInstallPolarClawPluginResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *InstallPolarClawPluginResponseBody
	GetApplicationId() *string
	SetCode(v int32) *InstallPolarClawPluginResponseBody
	GetCode() *int32
	SetMessage(v string) *InstallPolarClawPluginResponseBody
	GetMessage() *string
	SetNpmPackage(v string) *InstallPolarClawPluginResponseBody
	GetNpmPackage() *string
	SetOk(v bool) *InstallPolarClawPluginResponseBody
	GetOk() *bool
	SetPluginId(v string) *InstallPolarClawPluginResponseBody
	GetPluginId() *string
	SetRequestId(v string) *InstallPolarClawPluginResponseBody
	GetRequestId() *string
	SetRestarted(v bool) *InstallPolarClawPluginResponseBody
	GetRestarted() *bool
}

type InstallPolarClawPluginResponseBody struct {
	// The application ID.
	//
	// example:
	//
	// pa-**************
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response message.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The name of the installed npm package.
	//
	// example:
	//
	// @larksuite/openclaw-lark@2026.4.7
	NpmPackage *string `json:"NpmPackage,omitempty" xml:"NpmPackage,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Ok *bool `json:"Ok,omitempty" xml:"Ok,omitempty"`
	// The ID of the installed plugin.
	//
	// example:
	//
	// openclaw-lark
	PluginId *string `json:"PluginId,omitempty" xml:"PluginId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// CDB3258F-B5DE-43C4-8935-CBA0CA******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the gateway restarted.
	//
	// example:
	//
	// true
	Restarted *bool `json:"Restarted,omitempty" xml:"Restarted,omitempty"`
}

func (s InstallPolarClawPluginResponseBody) String() string {
	return dara.Prettify(s)
}

func (s InstallPolarClawPluginResponseBody) GoString() string {
	return s.String()
}

func (s *InstallPolarClawPluginResponseBody) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *InstallPolarClawPluginResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *InstallPolarClawPluginResponseBody) GetMessage() *string {
	return s.Message
}

func (s *InstallPolarClawPluginResponseBody) GetNpmPackage() *string {
	return s.NpmPackage
}

func (s *InstallPolarClawPluginResponseBody) GetOk() *bool {
	return s.Ok
}

func (s *InstallPolarClawPluginResponseBody) GetPluginId() *string {
	return s.PluginId
}

func (s *InstallPolarClawPluginResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *InstallPolarClawPluginResponseBody) GetRestarted() *bool {
	return s.Restarted
}

func (s *InstallPolarClawPluginResponseBody) SetApplicationId(v string) *InstallPolarClawPluginResponseBody {
	s.ApplicationId = &v
	return s
}

func (s *InstallPolarClawPluginResponseBody) SetCode(v int32) *InstallPolarClawPluginResponseBody {
	s.Code = &v
	return s
}

func (s *InstallPolarClawPluginResponseBody) SetMessage(v string) *InstallPolarClawPluginResponseBody {
	s.Message = &v
	return s
}

func (s *InstallPolarClawPluginResponseBody) SetNpmPackage(v string) *InstallPolarClawPluginResponseBody {
	s.NpmPackage = &v
	return s
}

func (s *InstallPolarClawPluginResponseBody) SetOk(v bool) *InstallPolarClawPluginResponseBody {
	s.Ok = &v
	return s
}

func (s *InstallPolarClawPluginResponseBody) SetPluginId(v string) *InstallPolarClawPluginResponseBody {
	s.PluginId = &v
	return s
}

func (s *InstallPolarClawPluginResponseBody) SetRequestId(v string) *InstallPolarClawPluginResponseBody {
	s.RequestId = &v
	return s
}

func (s *InstallPolarClawPluginResponseBody) SetRestarted(v bool) *InstallPolarClawPluginResponseBody {
	s.Restarted = &v
	return s
}

func (s *InstallPolarClawPluginResponseBody) Validate() error {
	return dara.Validate(s)
}
