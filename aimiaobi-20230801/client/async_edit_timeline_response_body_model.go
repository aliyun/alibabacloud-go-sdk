// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAsyncEditTimelineResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *AsyncEditTimelineResponseBody
	GetCode() *string
	SetData(v *AsyncEditTimelineResponseBodyData) *AsyncEditTimelineResponseBody
	GetData() *AsyncEditTimelineResponseBodyData
	SetHttpStatusCode(v int32) *AsyncEditTimelineResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *AsyncEditTimelineResponseBody
	GetMessage() *string
	SetRequestId(v string) *AsyncEditTimelineResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AsyncEditTimelineResponseBody
	GetSuccess() *bool
}

type AsyncEditTimelineResponseBody struct {
	// example:
	//
	// successful
	Code           *string                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *AsyncEditTimelineResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	HttpStatusCode *int32                             `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                            `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 1813ceee-7fe5-41b4-87e5-982a4d18cca5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AsyncEditTimelineResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AsyncEditTimelineResponseBody) GoString() string {
	return s.String()
}

func (s *AsyncEditTimelineResponseBody) GetCode() *string {
	return s.Code
}

func (s *AsyncEditTimelineResponseBody) GetData() *AsyncEditTimelineResponseBodyData {
	return s.Data
}

func (s *AsyncEditTimelineResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *AsyncEditTimelineResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AsyncEditTimelineResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AsyncEditTimelineResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AsyncEditTimelineResponseBody) SetCode(v string) *AsyncEditTimelineResponseBody {
	s.Code = &v
	return s
}

func (s *AsyncEditTimelineResponseBody) SetData(v *AsyncEditTimelineResponseBodyData) *AsyncEditTimelineResponseBody {
	s.Data = v
	return s
}

func (s *AsyncEditTimelineResponseBody) SetHttpStatusCode(v int32) *AsyncEditTimelineResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *AsyncEditTimelineResponseBody) SetMessage(v string) *AsyncEditTimelineResponseBody {
	s.Message = &v
	return s
}

func (s *AsyncEditTimelineResponseBody) SetRequestId(v string) *AsyncEditTimelineResponseBody {
	s.RequestId = &v
	return s
}

func (s *AsyncEditTimelineResponseBody) SetSuccess(v bool) *AsyncEditTimelineResponseBody {
	s.Success = &v
	return s
}

func (s *AsyncEditTimelineResponseBody) Validate() error {
	return dara.Validate(s)
}

type AsyncEditTimelineResponseBodyData struct {
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 51e4efd1908242eb93ca9bbb7fc4359d
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s AsyncEditTimelineResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s AsyncEditTimelineResponseBodyData) GoString() string {
	return s.String()
}

func (s *AsyncEditTimelineResponseBodyData) GetMessage() *string {
	return s.Message
}

func (s *AsyncEditTimelineResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *AsyncEditTimelineResponseBodyData) SetMessage(v string) *AsyncEditTimelineResponseBodyData {
	s.Message = &v
	return s
}

func (s *AsyncEditTimelineResponseBodyData) SetTaskId(v string) *AsyncEditTimelineResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *AsyncEditTimelineResponseBodyData) Validate() error {
	return dara.Validate(s)
}
