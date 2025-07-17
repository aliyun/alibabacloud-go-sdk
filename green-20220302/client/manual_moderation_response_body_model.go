// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iManualModerationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ManualModerationResponseBody
	GetCode() *int32
	SetData(v *ManualModerationResponseBodyData) *ManualModerationResponseBody
	GetData() *ManualModerationResponseBodyData
	SetMessage(v string) *ManualModerationResponseBody
	GetMessage() *string
	SetRequestId(v string) *ManualModerationResponseBody
	GetRequestId() *string
}

type ManualModerationResponseBody struct {
	// Status code
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// Returned data.
	Data *ManualModerationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Error message
	//
	// example:
	//
	// SUCCESS
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// ID of the request
	//
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ManualModerationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ManualModerationResponseBody) GoString() string {
	return s.String()
}

func (s *ManualModerationResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ManualModerationResponseBody) GetData() *ManualModerationResponseBodyData {
	return s.Data
}

func (s *ManualModerationResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ManualModerationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ManualModerationResponseBody) SetCode(v int32) *ManualModerationResponseBody {
	s.Code = &v
	return s
}

func (s *ManualModerationResponseBody) SetData(v *ManualModerationResponseBodyData) *ManualModerationResponseBody {
	s.Data = v
	return s
}

func (s *ManualModerationResponseBody) SetMessage(v string) *ManualModerationResponseBody {
	s.Message = &v
	return s
}

func (s *ManualModerationResponseBody) SetRequestId(v string) *ManualModerationResponseBody {
	s.RequestId = &v
	return s
}

func (s *ManualModerationResponseBody) Validate() error {
	return dara.Validate(s)
}

type ManualModerationResponseBodyData struct {
	// The value of dataId passed during the API request. This field will not be present if it was not provided during the request.
	//
	// example:
	//
	// 2a5389eb-4ff8-4584-ac99-644e2a539aa1
	DataId *string `json:"DataId,omitempty" xml:"DataId,omitempty"`
	// Task ID
	//
	// example:
	//
	// xxxxx-xxxxx
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s ManualModerationResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ManualModerationResponseBodyData) GoString() string {
	return s.String()
}

func (s *ManualModerationResponseBodyData) GetDataId() *string {
	return s.DataId
}

func (s *ManualModerationResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *ManualModerationResponseBodyData) SetDataId(v string) *ManualModerationResponseBodyData {
	s.DataId = &v
	return s
}

func (s *ManualModerationResponseBodyData) SetTaskId(v string) *ManualModerationResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *ManualModerationResponseBodyData) Validate() error {
	return dara.Validate(s)
}
