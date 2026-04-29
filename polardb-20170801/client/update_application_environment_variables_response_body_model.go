// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateApplicationEnvironmentVariablesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *UpdateApplicationEnvironmentVariablesResponseBody
	GetApplicationId() *string
	SetCode(v int32) *UpdateApplicationEnvironmentVariablesResponseBody
	GetCode() *int32
	SetMessage(v string) *UpdateApplicationEnvironmentVariablesResponseBody
	GetMessage() *string
	SetOk(v bool) *UpdateApplicationEnvironmentVariablesResponseBody
	GetOk() *bool
	SetRequestId(v string) *UpdateApplicationEnvironmentVariablesResponseBody
	GetRequestId() *string
	SetRestarted(v bool) *UpdateApplicationEnvironmentVariablesResponseBody
	GetRestarted() *bool
	SetTotalVariables(v int32) *UpdateApplicationEnvironmentVariablesResponseBody
	GetTotalVariables() *int32
	SetUpdatedKeys(v []*string) *UpdateApplicationEnvironmentVariablesResponseBody
	GetUpdatedKeys() []*string
}

type UpdateApplicationEnvironmentVariablesResponseBody struct {
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
	// Id of the request
	//
	// example:
	//
	// 6A2EE5B4-CC9F-46E1-A747-E43BC9******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Restarted *bool `json:"Restarted,omitempty" xml:"Restarted,omitempty"`
	// example:
	//
	// 1
	TotalVariables *int32    `json:"TotalVariables,omitempty" xml:"TotalVariables,omitempty"`
	UpdatedKeys    []*string `json:"UpdatedKeys,omitempty" xml:"UpdatedKeys,omitempty" type:"Repeated"`
}

func (s UpdateApplicationEnvironmentVariablesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateApplicationEnvironmentVariablesResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateApplicationEnvironmentVariablesResponseBody) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *UpdateApplicationEnvironmentVariablesResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *UpdateApplicationEnvironmentVariablesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateApplicationEnvironmentVariablesResponseBody) GetOk() *bool {
	return s.Ok
}

func (s *UpdateApplicationEnvironmentVariablesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateApplicationEnvironmentVariablesResponseBody) GetRestarted() *bool {
	return s.Restarted
}

func (s *UpdateApplicationEnvironmentVariablesResponseBody) GetTotalVariables() *int32 {
	return s.TotalVariables
}

func (s *UpdateApplicationEnvironmentVariablesResponseBody) GetUpdatedKeys() []*string {
	return s.UpdatedKeys
}

func (s *UpdateApplicationEnvironmentVariablesResponseBody) SetApplicationId(v string) *UpdateApplicationEnvironmentVariablesResponseBody {
	s.ApplicationId = &v
	return s
}

func (s *UpdateApplicationEnvironmentVariablesResponseBody) SetCode(v int32) *UpdateApplicationEnvironmentVariablesResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateApplicationEnvironmentVariablesResponseBody) SetMessage(v string) *UpdateApplicationEnvironmentVariablesResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateApplicationEnvironmentVariablesResponseBody) SetOk(v bool) *UpdateApplicationEnvironmentVariablesResponseBody {
	s.Ok = &v
	return s
}

func (s *UpdateApplicationEnvironmentVariablesResponseBody) SetRequestId(v string) *UpdateApplicationEnvironmentVariablesResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateApplicationEnvironmentVariablesResponseBody) SetRestarted(v bool) *UpdateApplicationEnvironmentVariablesResponseBody {
	s.Restarted = &v
	return s
}

func (s *UpdateApplicationEnvironmentVariablesResponseBody) SetTotalVariables(v int32) *UpdateApplicationEnvironmentVariablesResponseBody {
	s.TotalVariables = &v
	return s
}

func (s *UpdateApplicationEnvironmentVariablesResponseBody) SetUpdatedKeys(v []*string) *UpdateApplicationEnvironmentVariablesResponseBody {
	s.UpdatedKeys = v
	return s
}

func (s *UpdateApplicationEnvironmentVariablesResponseBody) Validate() error {
	return dara.Validate(s)
}
