// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateFuncSwitchRecordResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateFuncSwitchRecordResponseBody
	GetCode() *string
	SetData(v *UpdateFuncSwitchRecordResponseBodyData) *UpdateFuncSwitchRecordResponseBody
	GetData() *UpdateFuncSwitchRecordResponseBodyData
	SetMessage(v string) *UpdateFuncSwitchRecordResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateFuncSwitchRecordResponseBody
	GetRequestId() *string
}

type UpdateFuncSwitchRecordResponseBody struct {
	// example:
	//
	// Success
	Code *string                                 `json:"code,omitempty" xml:"code,omitempty"`
	Data *UpdateFuncSwitchRecordResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// result: code=1 msg=(Request failed, status_code != 200)
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 2D693121-C925-5154-8DF6-C09A8B369822
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateFuncSwitchRecordResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateFuncSwitchRecordResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateFuncSwitchRecordResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateFuncSwitchRecordResponseBody) GetData() *UpdateFuncSwitchRecordResponseBodyData {
	return s.Data
}

func (s *UpdateFuncSwitchRecordResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateFuncSwitchRecordResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateFuncSwitchRecordResponseBody) SetCode(v string) *UpdateFuncSwitchRecordResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateFuncSwitchRecordResponseBody) SetData(v *UpdateFuncSwitchRecordResponseBodyData) *UpdateFuncSwitchRecordResponseBody {
	s.Data = v
	return s
}

func (s *UpdateFuncSwitchRecordResponseBody) SetMessage(v string) *UpdateFuncSwitchRecordResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateFuncSwitchRecordResponseBody) SetRequestId(v string) *UpdateFuncSwitchRecordResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateFuncSwitchRecordResponseBody) Validate() error {
	return dara.Validate(s)
}

type UpdateFuncSwitchRecordResponseBodyData struct {
	// example:
	//
	// 63fc5acb99e642d793f42912612e8001
	TaskId *string `json:"task_id,omitempty" xml:"task_id,omitempty"`
}

func (s UpdateFuncSwitchRecordResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UpdateFuncSwitchRecordResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateFuncSwitchRecordResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *UpdateFuncSwitchRecordResponseBodyData) SetTaskId(v string) *UpdateFuncSwitchRecordResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *UpdateFuncSwitchRecordResponseBodyData) Validate() error {
	return dara.Validate(s)
}
