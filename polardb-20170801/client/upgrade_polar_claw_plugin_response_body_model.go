// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradePolarClawPluginResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *UpgradePolarClawPluginResponseBody
	GetApplicationId() *string
	SetCode(v int32) *UpgradePolarClawPluginResponseBody
	GetCode() *int32
	SetMessage(v string) *UpgradePolarClawPluginResponseBody
	GetMessage() *string
	SetNpmPackage(v string) *UpgradePolarClawPluginResponseBody
	GetNpmPackage() *string
	SetOk(v bool) *UpgradePolarClawPluginResponseBody
	GetOk() *bool
	SetPluginId(v string) *UpgradePolarClawPluginResponseBody
	GetPluginId() *string
	SetRequestId(v string) *UpgradePolarClawPluginResponseBody
	GetRequestId() *string
	SetRestarted(v bool) *UpgradePolarClawPluginResponseBody
	GetRestarted() *bool
}

type UpgradePolarClawPluginResponseBody struct {
	// example:
	//
	// pa-**************
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// @larksuite/openclaw-lark@2026.4.7
	NpmPackage *string `json:"NpmPackage,omitempty" xml:"NpmPackage,omitempty"`
	// example:
	//
	// true
	Ok *bool `json:"Ok,omitempty" xml:"Ok,omitempty"`
	// example:
	//
	// openclaw-lark
	PluginId *string `json:"PluginId,omitempty" xml:"PluginId,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 7E2FE3BB-C677-5FF9-9FC5-9CF364BD6BE5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Restarted *bool `json:"Restarted,omitempty" xml:"Restarted,omitempty"`
}

func (s UpgradePolarClawPluginResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpgradePolarClawPluginResponseBody) GoString() string {
	return s.String()
}

func (s *UpgradePolarClawPluginResponseBody) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *UpgradePolarClawPluginResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *UpgradePolarClawPluginResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpgradePolarClawPluginResponseBody) GetNpmPackage() *string {
	return s.NpmPackage
}

func (s *UpgradePolarClawPluginResponseBody) GetOk() *bool {
	return s.Ok
}

func (s *UpgradePolarClawPluginResponseBody) GetPluginId() *string {
	return s.PluginId
}

func (s *UpgradePolarClawPluginResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpgradePolarClawPluginResponseBody) GetRestarted() *bool {
	return s.Restarted
}

func (s *UpgradePolarClawPluginResponseBody) SetApplicationId(v string) *UpgradePolarClawPluginResponseBody {
	s.ApplicationId = &v
	return s
}

func (s *UpgradePolarClawPluginResponseBody) SetCode(v int32) *UpgradePolarClawPluginResponseBody {
	s.Code = &v
	return s
}

func (s *UpgradePolarClawPluginResponseBody) SetMessage(v string) *UpgradePolarClawPluginResponseBody {
	s.Message = &v
	return s
}

func (s *UpgradePolarClawPluginResponseBody) SetNpmPackage(v string) *UpgradePolarClawPluginResponseBody {
	s.NpmPackage = &v
	return s
}

func (s *UpgradePolarClawPluginResponseBody) SetOk(v bool) *UpgradePolarClawPluginResponseBody {
	s.Ok = &v
	return s
}

func (s *UpgradePolarClawPluginResponseBody) SetPluginId(v string) *UpgradePolarClawPluginResponseBody {
	s.PluginId = &v
	return s
}

func (s *UpgradePolarClawPluginResponseBody) SetRequestId(v string) *UpgradePolarClawPluginResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpgradePolarClawPluginResponseBody) SetRestarted(v bool) *UpgradePolarClawPluginResponseBody {
	s.Restarted = &v
	return s
}

func (s *UpgradePolarClawPluginResponseBody) Validate() error {
	return dara.Validate(s)
}
