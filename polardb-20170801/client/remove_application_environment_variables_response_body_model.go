// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveApplicationEnvironmentVariablesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *RemoveApplicationEnvironmentVariablesResponseBody
	GetApplicationId() *string
	SetCode(v int32) *RemoveApplicationEnvironmentVariablesResponseBody
	GetCode() *int32
	SetMessage(v string) *RemoveApplicationEnvironmentVariablesResponseBody
	GetMessage() *string
	SetOk(v bool) *RemoveApplicationEnvironmentVariablesResponseBody
	GetOk() *bool
	SetRemovedCount(v int32) *RemoveApplicationEnvironmentVariablesResponseBody
	GetRemovedCount() *int32
	SetRequestId(v string) *RemoveApplicationEnvironmentVariablesResponseBody
	GetRequestId() *string
	SetRestarted(v bool) *RemoveApplicationEnvironmentVariablesResponseBody
	GetRestarted() *bool
	SetTotalVariables(v int32) *RemoveApplicationEnvironmentVariablesResponseBody
	GetTotalVariables() *int32
}

type RemoveApplicationEnvironmentVariablesResponseBody struct {
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
	// 1
	RemovedCount *int32 `json:"RemovedCount,omitempty" xml:"RemovedCount,omitempty"`
	// Id of the request
	//
	// example:
	//
	// CDB3258F-B5DE-43C4-8935-CBA0CA******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Restarted *bool `json:"Restarted,omitempty" xml:"Restarted,omitempty"`
	// example:
	//
	// 0
	TotalVariables *int32 `json:"TotalVariables,omitempty" xml:"TotalVariables,omitempty"`
}

func (s RemoveApplicationEnvironmentVariablesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RemoveApplicationEnvironmentVariablesResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveApplicationEnvironmentVariablesResponseBody) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *RemoveApplicationEnvironmentVariablesResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *RemoveApplicationEnvironmentVariablesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *RemoveApplicationEnvironmentVariablesResponseBody) GetOk() *bool {
	return s.Ok
}

func (s *RemoveApplicationEnvironmentVariablesResponseBody) GetRemovedCount() *int32 {
	return s.RemovedCount
}

func (s *RemoveApplicationEnvironmentVariablesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RemoveApplicationEnvironmentVariablesResponseBody) GetRestarted() *bool {
	return s.Restarted
}

func (s *RemoveApplicationEnvironmentVariablesResponseBody) GetTotalVariables() *int32 {
	return s.TotalVariables
}

func (s *RemoveApplicationEnvironmentVariablesResponseBody) SetApplicationId(v string) *RemoveApplicationEnvironmentVariablesResponseBody {
	s.ApplicationId = &v
	return s
}

func (s *RemoveApplicationEnvironmentVariablesResponseBody) SetCode(v int32) *RemoveApplicationEnvironmentVariablesResponseBody {
	s.Code = &v
	return s
}

func (s *RemoveApplicationEnvironmentVariablesResponseBody) SetMessage(v string) *RemoveApplicationEnvironmentVariablesResponseBody {
	s.Message = &v
	return s
}

func (s *RemoveApplicationEnvironmentVariablesResponseBody) SetOk(v bool) *RemoveApplicationEnvironmentVariablesResponseBody {
	s.Ok = &v
	return s
}

func (s *RemoveApplicationEnvironmentVariablesResponseBody) SetRemovedCount(v int32) *RemoveApplicationEnvironmentVariablesResponseBody {
	s.RemovedCount = &v
	return s
}

func (s *RemoveApplicationEnvironmentVariablesResponseBody) SetRequestId(v string) *RemoveApplicationEnvironmentVariablesResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveApplicationEnvironmentVariablesResponseBody) SetRestarted(v bool) *RemoveApplicationEnvironmentVariablesResponseBody {
	s.Restarted = &v
	return s
}

func (s *RemoveApplicationEnvironmentVariablesResponseBody) SetTotalVariables(v int32) *RemoveApplicationEnvironmentVariablesResponseBody {
	s.TotalVariables = &v
	return s
}

func (s *RemoveApplicationEnvironmentVariablesResponseBody) Validate() error {
	return dara.Validate(s)
}
