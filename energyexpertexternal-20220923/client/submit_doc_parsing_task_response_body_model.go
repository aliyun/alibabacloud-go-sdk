// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitDocParsingTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *SubmitDocParsingTaskResponseBodyData) *SubmitDocParsingTaskResponseBody
	GetData() *SubmitDocParsingTaskResponseBodyData
	SetRequestId(v string) *SubmitDocParsingTaskResponseBody
	GetRequestId() *string
}

type SubmitDocParsingTaskResponseBody struct {
	// Return result.
	Data *SubmitDocParsingTaskResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// Request ID.
	//
	// example:
	//
	// 83A5A7DD-8974-5769-952E-590A97BEA34E
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s SubmitDocParsingTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitDocParsingTaskResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitDocParsingTaskResponseBody) GetData() *SubmitDocParsingTaskResponseBodyData {
	return s.Data
}

func (s *SubmitDocParsingTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitDocParsingTaskResponseBody) SetData(v *SubmitDocParsingTaskResponseBodyData) *SubmitDocParsingTaskResponseBody {
	s.Data = v
	return s
}

func (s *SubmitDocParsingTaskResponseBody) SetRequestId(v string) *SubmitDocParsingTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitDocParsingTaskResponseBody) Validate() error {
	return dara.Validate(s)
}

type SubmitDocParsingTaskResponseBodyData struct {
	// TaskID
	//
	// example:
	//
	// ae9d07be-1a11-4d30-be75-cc962b98279c
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s SubmitDocParsingTaskResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s SubmitDocParsingTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *SubmitDocParsingTaskResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *SubmitDocParsingTaskResponseBodyData) SetTaskId(v string) *SubmitDocParsingTaskResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *SubmitDocParsingTaskResponseBodyData) Validate() error {
	return dara.Validate(s)
}
