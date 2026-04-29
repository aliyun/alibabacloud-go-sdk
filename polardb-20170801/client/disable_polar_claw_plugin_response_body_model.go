// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisablePolarClawPluginResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *DisablePolarClawPluginResponseBody
	GetApplicationId() *string
	SetCode(v int32) *DisablePolarClawPluginResponseBody
	GetCode() *int32
	SetMessage(v string) *DisablePolarClawPluginResponseBody
	GetMessage() *string
	SetOk(v bool) *DisablePolarClawPluginResponseBody
	GetOk() *bool
	SetPluginId(v string) *DisablePolarClawPluginResponseBody
	GetPluginId() *string
	SetRequestId(v string) *DisablePolarClawPluginResponseBody
	GetRequestId() *string
	SetRestarted(v bool) *DisablePolarClawPluginResponseBody
	GetRestarted() *bool
}

type DisablePolarClawPluginResponseBody struct {
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
	// 580EF224-9647-59E7-9950-D9EBFD6A2921
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Restarted *bool `json:"Restarted,omitempty" xml:"Restarted,omitempty"`
}

func (s DisablePolarClawPluginResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DisablePolarClawPluginResponseBody) GoString() string {
	return s.String()
}

func (s *DisablePolarClawPluginResponseBody) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *DisablePolarClawPluginResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DisablePolarClawPluginResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DisablePolarClawPluginResponseBody) GetOk() *bool {
	return s.Ok
}

func (s *DisablePolarClawPluginResponseBody) GetPluginId() *string {
	return s.PluginId
}

func (s *DisablePolarClawPluginResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DisablePolarClawPluginResponseBody) GetRestarted() *bool {
	return s.Restarted
}

func (s *DisablePolarClawPluginResponseBody) SetApplicationId(v string) *DisablePolarClawPluginResponseBody {
	s.ApplicationId = &v
	return s
}

func (s *DisablePolarClawPluginResponseBody) SetCode(v int32) *DisablePolarClawPluginResponseBody {
	s.Code = &v
	return s
}

func (s *DisablePolarClawPluginResponseBody) SetMessage(v string) *DisablePolarClawPluginResponseBody {
	s.Message = &v
	return s
}

func (s *DisablePolarClawPluginResponseBody) SetOk(v bool) *DisablePolarClawPluginResponseBody {
	s.Ok = &v
	return s
}

func (s *DisablePolarClawPluginResponseBody) SetPluginId(v string) *DisablePolarClawPluginResponseBody {
	s.PluginId = &v
	return s
}

func (s *DisablePolarClawPluginResponseBody) SetRequestId(v string) *DisablePolarClawPluginResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisablePolarClawPluginResponseBody) SetRestarted(v bool) *DisablePolarClawPluginResponseBody {
	s.Restarted = &v
	return s
}

func (s *DisablePolarClawPluginResponseBody) Validate() error {
	return dara.Validate(s)
}
