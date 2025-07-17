// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVoiceModerationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *VoiceModerationResponseBody
	GetCode() *int32
	SetData(v *VoiceModerationResponseBodyData) *VoiceModerationResponseBody
	GetData() *VoiceModerationResponseBodyData
	SetMessage(v string) *VoiceModerationResponseBody
	GetMessage() *string
	SetRequestId(v string) *VoiceModerationResponseBody
	GetRequestId() *string
}

type VoiceModerationResponseBody struct {
	// The returned HTTP status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	Data *VoiceModerationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The message that is returned in response to the request.
	//
	// example:
	//
	// SUCCESS
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s VoiceModerationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s VoiceModerationResponseBody) GoString() string {
	return s.String()
}

func (s *VoiceModerationResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *VoiceModerationResponseBody) GetData() *VoiceModerationResponseBodyData {
	return s.Data
}

func (s *VoiceModerationResponseBody) GetMessage() *string {
	return s.Message
}

func (s *VoiceModerationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *VoiceModerationResponseBody) SetCode(v int32) *VoiceModerationResponseBody {
	s.Code = &v
	return s
}

func (s *VoiceModerationResponseBody) SetData(v *VoiceModerationResponseBodyData) *VoiceModerationResponseBody {
	s.Data = v
	return s
}

func (s *VoiceModerationResponseBody) SetMessage(v string) *VoiceModerationResponseBody {
	s.Message = &v
	return s
}

func (s *VoiceModerationResponseBody) SetRequestId(v string) *VoiceModerationResponseBody {
	s.RequestId = &v
	return s
}

func (s *VoiceModerationResponseBody) Validate() error {
	return dara.Validate(s)
}

type VoiceModerationResponseBodyData struct {
	// The ID of the moderated object.
	//
	// example:
	//
	// data1234
	DataId *string `json:"DataId,omitempty" xml:"DataId,omitempty"`
	// The task ID.
	//
	// example:
	//
	// xxxxx-xxxxx
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s VoiceModerationResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s VoiceModerationResponseBodyData) GoString() string {
	return s.String()
}

func (s *VoiceModerationResponseBodyData) GetDataId() *string {
	return s.DataId
}

func (s *VoiceModerationResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *VoiceModerationResponseBodyData) SetDataId(v string) *VoiceModerationResponseBodyData {
	s.DataId = &v
	return s
}

func (s *VoiceModerationResponseBodyData) SetTaskId(v string) *VoiceModerationResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *VoiceModerationResponseBodyData) Validate() error {
	return dara.Validate(s)
}
