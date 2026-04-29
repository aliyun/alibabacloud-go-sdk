// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUninstallPolarClawPluginResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *UninstallPolarClawPluginResponseBody
	GetApplicationId() *string
	SetCode(v int32) *UninstallPolarClawPluginResponseBody
	GetCode() *int32
	SetMessage(v string) *UninstallPolarClawPluginResponseBody
	GetMessage() *string
	SetOk(v bool) *UninstallPolarClawPluginResponseBody
	GetOk() *bool
	SetPluginId(v string) *UninstallPolarClawPluginResponseBody
	GetPluginId() *string
	SetRequestId(v string) *UninstallPolarClawPluginResponseBody
	GetRequestId() *string
	SetRestarted(v bool) *UninstallPolarClawPluginResponseBody
	GetRestarted() *bool
}

type UninstallPolarClawPluginResponseBody struct {
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
	// 3E5CD764-FCCA-5C9C-838E-20E0DE84B2AF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Restarted *bool `json:"Restarted,omitempty" xml:"Restarted,omitempty"`
}

func (s UninstallPolarClawPluginResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UninstallPolarClawPluginResponseBody) GoString() string {
	return s.String()
}

func (s *UninstallPolarClawPluginResponseBody) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *UninstallPolarClawPluginResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *UninstallPolarClawPluginResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UninstallPolarClawPluginResponseBody) GetOk() *bool {
	return s.Ok
}

func (s *UninstallPolarClawPluginResponseBody) GetPluginId() *string {
	return s.PluginId
}

func (s *UninstallPolarClawPluginResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UninstallPolarClawPluginResponseBody) GetRestarted() *bool {
	return s.Restarted
}

func (s *UninstallPolarClawPluginResponseBody) SetApplicationId(v string) *UninstallPolarClawPluginResponseBody {
	s.ApplicationId = &v
	return s
}

func (s *UninstallPolarClawPluginResponseBody) SetCode(v int32) *UninstallPolarClawPluginResponseBody {
	s.Code = &v
	return s
}

func (s *UninstallPolarClawPluginResponseBody) SetMessage(v string) *UninstallPolarClawPluginResponseBody {
	s.Message = &v
	return s
}

func (s *UninstallPolarClawPluginResponseBody) SetOk(v bool) *UninstallPolarClawPluginResponseBody {
	s.Ok = &v
	return s
}

func (s *UninstallPolarClawPluginResponseBody) SetPluginId(v string) *UninstallPolarClawPluginResponseBody {
	s.PluginId = &v
	return s
}

func (s *UninstallPolarClawPluginResponseBody) SetRequestId(v string) *UninstallPolarClawPluginResponseBody {
	s.RequestId = &v
	return s
}

func (s *UninstallPolarClawPluginResponseBody) SetRestarted(v bool) *UninstallPolarClawPluginResponseBody {
	s.Restarted = &v
	return s
}

func (s *UninstallPolarClawPluginResponseBody) Validate() error {
	return dara.Validate(s)
}
