// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWmExtractTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *CreateWmExtractTaskResponseBodyData) *CreateWmExtractTaskResponseBody
	GetData() *CreateWmExtractTaskResponseBodyData
	SetRequestId(v string) *CreateWmExtractTaskResponseBody
	GetRequestId() *string
}

type CreateWmExtractTaskResponseBody struct {
	// The information about the watermark extraction task.
	Data *CreateWmExtractTaskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// D6707286-A50E-57B1-B2CF-EFAC59E850D8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateWmExtractTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateWmExtractTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateWmExtractTaskResponseBody) GetData() *CreateWmExtractTaskResponseBodyData {
	return s.Data
}

func (s *CreateWmExtractTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateWmExtractTaskResponseBody) SetData(v *CreateWmExtractTaskResponseBodyData) *CreateWmExtractTaskResponseBody {
	s.Data = v
	return s
}

func (s *CreateWmExtractTaskResponseBody) SetRequestId(v string) *CreateWmExtractTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateWmExtractTaskResponseBody) Validate() error {
	return dara.Validate(s)
}

type CreateWmExtractTaskResponseBodyData struct {
	// The task ID. You can use task IDs to query task results.
	//
	// example:
	//
	// wmt-9648c22d2eb2cb57bb855dcae7898464********
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CreateWmExtractTaskResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateWmExtractTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateWmExtractTaskResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *CreateWmExtractTaskResponseBodyData) SetTaskId(v string) *CreateWmExtractTaskResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *CreateWmExtractTaskResponseBodyData) Validate() error {
	return dara.Validate(s)
}
