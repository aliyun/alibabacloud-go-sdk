// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInstallHotfixResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *InstallHotfixResponseBody
	GetCode() *string
	SetData(v *InstallHotfixResponseBodyData) *InstallHotfixResponseBody
	GetData() *InstallHotfixResponseBodyData
	SetMessage(v string) *InstallHotfixResponseBody
	GetMessage() *string
	SetRequestId(v string) *InstallHotfixResponseBody
	GetRequestId() *string
}

type InstallHotfixResponseBody struct {
	// The status code.
	//
	// - If `code == Success`, the authorization is successful.
	//
	// - Other status codes indicate that the authorization failed. If the authorization fails, check the `message` field for the detailed fault message.
	//
	// example:
	//
	// Success
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The returned data.
	Data *InstallHotfixResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The error message.
	//
	// - If `code == Success`, this field is empty.
	//
	// - Otherwise, this field contains the request error message.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID, which can be used for end-to-end diagnostics.
	//
	// example:
	//
	// 2D693121-C925-5154-8DF6-C09A8B369822
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s InstallHotfixResponseBody) String() string {
	return dara.Prettify(s)
}

func (s InstallHotfixResponseBody) GoString() string {
	return s.String()
}

func (s *InstallHotfixResponseBody) GetCode() *string {
	return s.Code
}

func (s *InstallHotfixResponseBody) GetData() *InstallHotfixResponseBodyData {
	return s.Data
}

func (s *InstallHotfixResponseBody) GetMessage() *string {
	return s.Message
}

func (s *InstallHotfixResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *InstallHotfixResponseBody) SetCode(v string) *InstallHotfixResponseBody {
	s.Code = &v
	return s
}

func (s *InstallHotfixResponseBody) SetData(v *InstallHotfixResponseBodyData) *InstallHotfixResponseBody {
	s.Data = v
	return s
}

func (s *InstallHotfixResponseBody) SetMessage(v string) *InstallHotfixResponseBody {
	s.Message = &v
	return s
}

func (s *InstallHotfixResponseBody) SetRequestId(v string) *InstallHotfixResponseBody {
	s.RequestId = &v
	return s
}

func (s *InstallHotfixResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type InstallHotfixResponseBodyData struct {
	// The asynchronous task ID.
	//
	// example:
	//
	// hotfix-da3dcc0b-5019-4b59-b12f-2e500da9083c
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s InstallHotfixResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s InstallHotfixResponseBodyData) GoString() string {
	return s.String()
}

func (s *InstallHotfixResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *InstallHotfixResponseBodyData) SetTaskId(v string) *InstallHotfixResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *InstallHotfixResponseBodyData) Validate() error {
	return dara.Validate(s)
}
