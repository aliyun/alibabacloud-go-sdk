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
	SetRequestId(v string) *AsyncCreateClipsTaskResponseBody
	GetRequestId() *string
}

type AsyncCreateClipsTaskResponseBody struct {
	// example:
	//
	// successful
	Code *string                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *AsyncCreateClipsTaskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 1813ceee-7fe5-41b4-87e5-982a4d18cca5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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

func (s *AsyncCreateClipsTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AsyncCreateClipsTaskResponseBody) SetCode(v string) *AsyncCreateClipsTaskResponseBody {
	s.Code = &v
	return s
}

func (s *AsyncCreateClipsTaskResponseBody) SetData(v *AsyncCreateClipsTaskResponseBodyData) *AsyncCreateClipsTaskResponseBody {
	s.Data = v
	return s
}

func (s *AsyncCreateClipsTaskResponseBody) SetRequestId(v string) *AsyncCreateClipsTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *AsyncCreateClipsTaskResponseBody) Validate() error {
	return dara.Validate(s)
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
