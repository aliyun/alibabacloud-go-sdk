// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAsyncCreateClipsTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *AsyncCreateClipsTaskResponseBody
	GetCode() *string
	SetData(v *AsyncCreateClipsTaskResponseBodyData) *AsyncCreateClipsTaskResponseBody
	GetData() *AsyncCreateClipsTaskResponseBodyData
	SetHttpStatusCode(v int32) *AsyncCreateClipsTaskResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *AsyncCreateClipsTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *AsyncCreateClipsTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AsyncCreateClipsTaskResponseBody
	GetSuccess() *bool
}

type AsyncCreateClipsTaskResponseBody struct {
	// example:
	//
	// successful
	Code           *string                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *AsyncCreateClipsTaskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	HttpStatusCode *int32                                `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                               `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 1813ceee-7fe5-41b4-87e5-982a4d18cca5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AsyncCreateClipsTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AsyncCreateClipsTaskResponseBody) GoString() string {
	return s.String()
}

func (s *AsyncCreateClipsTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *AsyncCreateClipsTaskResponseBody) GetData() *AsyncCreateClipsTaskResponseBodyData {
	return s.Data
}

func (s *AsyncCreateClipsTaskResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *AsyncCreateClipsTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AsyncCreateClipsTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AsyncCreateClipsTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AsyncCreateClipsTaskResponseBody) SetCode(v string) *AsyncCreateClipsTaskResponseBody {
	s.Code = &v
	return s
}

func (s *AsyncCreateClipsTaskResponseBody) SetData(v *AsyncCreateClipsTaskResponseBodyData) *AsyncCreateClipsTaskResponseBody {
	s.Data = v
	return s
}

func (s *AsyncCreateClipsTaskResponseBody) SetHttpStatusCode(v int32) *AsyncCreateClipsTaskResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *AsyncCreateClipsTaskResponseBody) SetMessage(v string) *AsyncCreateClipsTaskResponseBody {
	s.Message = &v
	return s
}

func (s *AsyncCreateClipsTaskResponseBody) SetRequestId(v string) *AsyncCreateClipsTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *AsyncCreateClipsTaskResponseBody) SetSuccess(v bool) *AsyncCreateClipsTaskResponseBody {
	s.Success = &v
	return s
}

func (s *AsyncCreateClipsTaskResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AsyncCreateClipsTaskResponseBodyData struct {
	// example:
	//
	// 3f7045e099474ba28ceca1b4eb6d6e21
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s AsyncCreateClipsTaskResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s AsyncCreateClipsTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *AsyncCreateClipsTaskResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *AsyncCreateClipsTaskResponseBodyData) SetTaskId(v string) *AsyncCreateClipsTaskResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *AsyncCreateClipsTaskResponseBodyData) Validate() error {
	return dara.Validate(s)
}
