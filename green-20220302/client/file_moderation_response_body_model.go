// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFileModerationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *FileModerationResponseBody
	GetCode() *int32
	SetData(v *FileModerationResponseBodyData) *FileModerationResponseBody
	GetData() *FileModerationResponseBodyData
	SetMessage(v string) *FileModerationResponseBody
	GetMessage() *string
	SetRequestId(v string) *FileModerationResponseBody
	GetRequestId() *string
}

type FileModerationResponseBody struct {
	// The returned HTTP status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	Data *FileModerationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s FileModerationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s FileModerationResponseBody) GoString() string {
	return s.String()
}

func (s *FileModerationResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *FileModerationResponseBody) GetData() *FileModerationResponseBodyData {
	return s.Data
}

func (s *FileModerationResponseBody) GetMessage() *string {
	return s.Message
}

func (s *FileModerationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *FileModerationResponseBody) SetCode(v int32) *FileModerationResponseBody {
	s.Code = &v
	return s
}

func (s *FileModerationResponseBody) SetData(v *FileModerationResponseBodyData) *FileModerationResponseBody {
	s.Data = v
	return s
}

func (s *FileModerationResponseBody) SetMessage(v string) *FileModerationResponseBody {
	s.Message = &v
	return s
}

func (s *FileModerationResponseBody) SetRequestId(v string) *FileModerationResponseBody {
	s.RequestId = &v
	return s
}

func (s *FileModerationResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type FileModerationResponseBodyData struct {
	// The task ID.
	//
	// example:
	//
	// xxxxx-xxxxx
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s FileModerationResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s FileModerationResponseBodyData) GoString() string {
	return s.String()
}

func (s *FileModerationResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *FileModerationResponseBodyData) SetTaskId(v string) *FileModerationResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *FileModerationResponseBodyData) Validate() error {
	return dara.Validate(s)
}
