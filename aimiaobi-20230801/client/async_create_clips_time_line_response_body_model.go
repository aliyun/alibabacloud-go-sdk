// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAsyncCreateClipsTimeLineResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *AsyncCreateClipsTimeLineResponseBody
	GetCode() *string
	SetData(v *AsyncCreateClipsTimeLineResponseBodyData) *AsyncCreateClipsTimeLineResponseBody
	GetData() *AsyncCreateClipsTimeLineResponseBodyData
	SetHttpStatusCode(v int32) *AsyncCreateClipsTimeLineResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *AsyncCreateClipsTimeLineResponseBody
	GetMessage() *string
	SetRequestId(v string) *AsyncCreateClipsTimeLineResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AsyncCreateClipsTimeLineResponseBody
	GetSuccess() *bool
}

type AsyncCreateClipsTimeLineResponseBody struct {
	// example:
	//
	// successful
	Code           *string                                   `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *AsyncCreateClipsTimeLineResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	HttpStatusCode *int32                                    `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                                   `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 1813ceee-7fe5-41b4-87e5-982a4d18cca5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AsyncCreateClipsTimeLineResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AsyncCreateClipsTimeLineResponseBody) GoString() string {
	return s.String()
}

func (s *AsyncCreateClipsTimeLineResponseBody) GetCode() *string {
	return s.Code
}

func (s *AsyncCreateClipsTimeLineResponseBody) GetData() *AsyncCreateClipsTimeLineResponseBodyData {
	return s.Data
}

func (s *AsyncCreateClipsTimeLineResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *AsyncCreateClipsTimeLineResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AsyncCreateClipsTimeLineResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AsyncCreateClipsTimeLineResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AsyncCreateClipsTimeLineResponseBody) SetCode(v string) *AsyncCreateClipsTimeLineResponseBody {
	s.Code = &v
	return s
}

func (s *AsyncCreateClipsTimeLineResponseBody) SetData(v *AsyncCreateClipsTimeLineResponseBodyData) *AsyncCreateClipsTimeLineResponseBody {
	s.Data = v
	return s
}

func (s *AsyncCreateClipsTimeLineResponseBody) SetHttpStatusCode(v int32) *AsyncCreateClipsTimeLineResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *AsyncCreateClipsTimeLineResponseBody) SetMessage(v string) *AsyncCreateClipsTimeLineResponseBody {
	s.Message = &v
	return s
}

func (s *AsyncCreateClipsTimeLineResponseBody) SetRequestId(v string) *AsyncCreateClipsTimeLineResponseBody {
	s.RequestId = &v
	return s
}

func (s *AsyncCreateClipsTimeLineResponseBody) SetSuccess(v bool) *AsyncCreateClipsTimeLineResponseBody {
	s.Success = &v
	return s
}

func (s *AsyncCreateClipsTimeLineResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AsyncCreateClipsTimeLineResponseBodyData struct {
	// example:
	//
	// 3f7045e099474ba28ceca1b4eb6d6e21
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s AsyncCreateClipsTimeLineResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s AsyncCreateClipsTimeLineResponseBodyData) GoString() string {
	return s.String()
}

func (s *AsyncCreateClipsTimeLineResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *AsyncCreateClipsTimeLineResponseBodyData) SetTaskId(v string) *AsyncCreateClipsTimeLineResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *AsyncCreateClipsTimeLineResponseBodyData) Validate() error {
	return dara.Validate(s)
}
