// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMultiModalGuardAsyncResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *MultiModalGuardAsyncResponseBody
	GetCode() *int32
	SetData(v *MultiModalGuardAsyncResponseBodyData) *MultiModalGuardAsyncResponseBody
	GetData() *MultiModalGuardAsyncResponseBodyData
	SetMessage(v string) *MultiModalGuardAsyncResponseBody
	GetMessage() *string
	SetRequestId(v string) *MultiModalGuardAsyncResponseBody
	GetRequestId() *string
}

type MultiModalGuardAsyncResponseBody struct {
	// The response code. A value of 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response data.
	Data *MultiModalGuardAsyncResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The response message.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s MultiModalGuardAsyncResponseBody) String() string {
	return dara.Prettify(s)
}

func (s MultiModalGuardAsyncResponseBody) GoString() string {
	return s.String()
}

func (s *MultiModalGuardAsyncResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *MultiModalGuardAsyncResponseBody) GetData() *MultiModalGuardAsyncResponseBodyData {
	return s.Data
}

func (s *MultiModalGuardAsyncResponseBody) GetMessage() *string {
	return s.Message
}

func (s *MultiModalGuardAsyncResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *MultiModalGuardAsyncResponseBody) SetCode(v int32) *MultiModalGuardAsyncResponseBody {
	s.Code = &v
	return s
}

func (s *MultiModalGuardAsyncResponseBody) SetData(v *MultiModalGuardAsyncResponseBodyData) *MultiModalGuardAsyncResponseBody {
	s.Data = v
	return s
}

func (s *MultiModalGuardAsyncResponseBody) SetMessage(v string) *MultiModalGuardAsyncResponseBody {
	s.Message = &v
	return s
}

func (s *MultiModalGuardAsyncResponseBody) SetRequestId(v string) *MultiModalGuardAsyncResponseBody {
	s.RequestId = &v
	return s
}

func (s *MultiModalGuardAsyncResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type MultiModalGuardAsyncResponseBodyData struct {
	// The custom data ID.
	//
	// example:
	//
	// dataIdxxx
	DataId *string `json:"DataId,omitempty" xml:"DataId,omitempty"`
	// The ID of the asynchronous task.
	//
	// example:
	//
	// au_f_xxxxx
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s MultiModalGuardAsyncResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s MultiModalGuardAsyncResponseBodyData) GoString() string {
	return s.String()
}

func (s *MultiModalGuardAsyncResponseBodyData) GetDataId() *string {
	return s.DataId
}

func (s *MultiModalGuardAsyncResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *MultiModalGuardAsyncResponseBodyData) SetDataId(v string) *MultiModalGuardAsyncResponseBodyData {
	s.DataId = &v
	return s
}

func (s *MultiModalGuardAsyncResponseBodyData) SetTaskId(v string) *MultiModalGuardAsyncResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *MultiModalGuardAsyncResponseBodyData) Validate() error {
	return dara.Validate(s)
}
