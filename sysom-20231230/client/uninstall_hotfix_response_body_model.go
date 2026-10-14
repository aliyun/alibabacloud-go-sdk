// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUninstallHotfixResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UninstallHotfixResponseBody
	GetCode() *string
	SetData(v *UninstallHotfixResponseBodyData) *UninstallHotfixResponseBody
	GetData() *UninstallHotfixResponseBodyData
	SetMessage(v string) *UninstallHotfixResponseBody
	GetMessage() *string
	SetRequestId(v string) *UninstallHotfixResponseBody
	GetRequestId() *string
}

type UninstallHotfixResponseBody struct {
	// The status code.
	//
	// - If `code == Success`, the authorization is successful.
	//
	// - Other status codes indicate that the authorization failed. If the authorization fails, check the `message` field for the detailed fault information.
	//
	// example:
	//
	// Success
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The returned data.
	Data *UninstallHotfixResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The error message.
	//
	// - If `code == Success`, this field is empty.
	//
	// - Otherwise, this field contains the error message of the request.
	//
	// example:
	//
	// ""
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 2D693121-C925-5154-8DF6-C09A8B369822
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UninstallHotfixResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UninstallHotfixResponseBody) GoString() string {
	return s.String()
}

func (s *UninstallHotfixResponseBody) GetCode() *string {
	return s.Code
}

func (s *UninstallHotfixResponseBody) GetData() *UninstallHotfixResponseBodyData {
	return s.Data
}

func (s *UninstallHotfixResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UninstallHotfixResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UninstallHotfixResponseBody) SetCode(v string) *UninstallHotfixResponseBody {
	s.Code = &v
	return s
}

func (s *UninstallHotfixResponseBody) SetData(v *UninstallHotfixResponseBodyData) *UninstallHotfixResponseBody {
	s.Data = v
	return s
}

func (s *UninstallHotfixResponseBody) SetMessage(v string) *UninstallHotfixResponseBody {
	s.Message = &v
	return s
}

func (s *UninstallHotfixResponseBody) SetRequestId(v string) *UninstallHotfixResponseBody {
	s.RequestId = &v
	return s
}

func (s *UninstallHotfixResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UninstallHotfixResponseBodyData struct {
	// The asynchronous task ID.
	//
	// example:
	//
	// hotfix-da3dcc0b-5019-4b59-b12f-2e500da9083c
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s UninstallHotfixResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UninstallHotfixResponseBodyData) GoString() string {
	return s.String()
}

func (s *UninstallHotfixResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *UninstallHotfixResponseBodyData) SetTaskId(v string) *UninstallHotfixResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *UninstallHotfixResponseBodyData) Validate() error {
	return dara.Validate(s)
}
