// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVideoModerationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *VideoModerationResponseBody
	GetCode() *int32
	SetData(v *VideoModerationResponseBodyData) *VideoModerationResponseBody
	GetData() *VideoModerationResponseBodyData
	SetMessage(v string) *VideoModerationResponseBody
	GetMessage() *string
	SetRequestId(v string) *VideoModerationResponseBody
	GetRequestId() *string
}

type VideoModerationResponseBody struct {
	// The returned HTTP status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	Data *VideoModerationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s VideoModerationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s VideoModerationResponseBody) GoString() string {
	return s.String()
}

func (s *VideoModerationResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *VideoModerationResponseBody) GetData() *VideoModerationResponseBodyData {
	return s.Data
}

func (s *VideoModerationResponseBody) GetMessage() *string {
	return s.Message
}

func (s *VideoModerationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *VideoModerationResponseBody) SetCode(v int32) *VideoModerationResponseBody {
	s.Code = &v
	return s
}

func (s *VideoModerationResponseBody) SetData(v *VideoModerationResponseBodyData) *VideoModerationResponseBody {
	s.Data = v
	return s
}

func (s *VideoModerationResponseBody) SetMessage(v string) *VideoModerationResponseBody {
	s.Message = &v
	return s
}

func (s *VideoModerationResponseBody) SetRequestId(v string) *VideoModerationResponseBody {
	s.RequestId = &v
	return s
}

func (s *VideoModerationResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type VideoModerationResponseBodyData struct {
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

func (s VideoModerationResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s VideoModerationResponseBodyData) GoString() string {
	return s.String()
}

func (s *VideoModerationResponseBodyData) GetDataId() *string {
	return s.DataId
}

func (s *VideoModerationResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *VideoModerationResponseBodyData) SetDataId(v string) *VideoModerationResponseBodyData {
	s.DataId = &v
	return s
}

func (s *VideoModerationResponseBodyData) SetTaskId(v string) *VideoModerationResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *VideoModerationResponseBodyData) Validate() error {
	return dara.Validate(s)
}
