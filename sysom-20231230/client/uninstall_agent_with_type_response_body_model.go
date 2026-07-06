// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUninstallAgentWithTypeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UninstallAgentWithTypeResponseBody
	GetCode() *string
	SetData(v *UninstallAgentWithTypeResponseBodyData) *UninstallAgentWithTypeResponseBody
	GetData() *UninstallAgentWithTypeResponseBodyData
	SetMessage(v string) *UninstallAgentWithTypeResponseBody
	GetMessage() *string
	SetRequestId(v string) *UninstallAgentWithTypeResponseBody
	GetRequestId() *string
}

type UninstallAgentWithTypeResponseBody struct {
	// example:
	//
	// Success
	Code *string                                 `json:"code,omitempty" xml:"code,omitempty"`
	Data *UninstallAgentWithTypeResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// “”
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 13772206-1162-5A0F-81F0-79A10C249A5E
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UninstallAgentWithTypeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UninstallAgentWithTypeResponseBody) GoString() string {
	return s.String()
}

func (s *UninstallAgentWithTypeResponseBody) GetCode() *string {
	return s.Code
}

func (s *UninstallAgentWithTypeResponseBody) GetData() *UninstallAgentWithTypeResponseBodyData {
	return s.Data
}

func (s *UninstallAgentWithTypeResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UninstallAgentWithTypeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UninstallAgentWithTypeResponseBody) SetCode(v string) *UninstallAgentWithTypeResponseBody {
	s.Code = &v
	return s
}

func (s *UninstallAgentWithTypeResponseBody) SetData(v *UninstallAgentWithTypeResponseBodyData) *UninstallAgentWithTypeResponseBody {
	s.Data = v
	return s
}

func (s *UninstallAgentWithTypeResponseBody) SetMessage(v string) *UninstallAgentWithTypeResponseBody {
	s.Message = &v
	return s
}

func (s *UninstallAgentWithTypeResponseBody) SetRequestId(v string) *UninstallAgentWithTypeResponseBody {
	s.RequestId = &v
	return s
}

func (s *UninstallAgentWithTypeResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UninstallAgentWithTypeResponseBodyData struct {
	// example:
	//
	// 63fc5acb99e642d793f42912612e8001
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s UninstallAgentWithTypeResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UninstallAgentWithTypeResponseBodyData) GoString() string {
	return s.String()
}

func (s *UninstallAgentWithTypeResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *UninstallAgentWithTypeResponseBodyData) SetTaskId(v string) *UninstallAgentWithTypeResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *UninstallAgentWithTypeResponseBodyData) Validate() error {
	return dara.Validate(s)
}
