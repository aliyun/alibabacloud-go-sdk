// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateJobScriptResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *UpdateJobScriptResponseBody
	GetCode() *int32
	SetMessage(v string) *UpdateJobScriptResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateJobScriptResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateJobScriptResponseBody
	GetSuccess() *bool
}

type UpdateJobScriptResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// not support update script, please upgrade engine version to 2.2.2+
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 29ED6209-5DE6-5E1D-89B0-B7B1D823A1BF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateJobScriptResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateJobScriptResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateJobScriptResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *UpdateJobScriptResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateJobScriptResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateJobScriptResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateJobScriptResponseBody) SetCode(v int32) *UpdateJobScriptResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateJobScriptResponseBody) SetMessage(v string) *UpdateJobScriptResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateJobScriptResponseBody) SetRequestId(v string) *UpdateJobScriptResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateJobScriptResponseBody) SetSuccess(v bool) *UpdateJobScriptResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateJobScriptResponseBody) Validate() error {
	return dara.Validate(s)
}
