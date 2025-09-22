// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitVLExtractionTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *SubmitVLExtractionTaskResponseBodyData) *SubmitVLExtractionTaskResponseBody
	GetData() *SubmitVLExtractionTaskResponseBodyData
	SetRequestId(v string) *SubmitVLExtractionTaskResponseBody
	GetRequestId() *string
}

type SubmitVLExtractionTaskResponseBody struct {
	// Returned data structure.
	Data *SubmitVLExtractionTaskResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// Request ID.
	//
	// example:
	//
	// 83A5A7DD-8974-5769-952E-590A97BEA34E
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s SubmitVLExtractionTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitVLExtractionTaskResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitVLExtractionTaskResponseBody) GetData() *SubmitVLExtractionTaskResponseBodyData {
	return s.Data
}

func (s *SubmitVLExtractionTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitVLExtractionTaskResponseBody) SetData(v *SubmitVLExtractionTaskResponseBodyData) *SubmitVLExtractionTaskResponseBody {
	s.Data = v
	return s
}

func (s *SubmitVLExtractionTaskResponseBody) SetRequestId(v string) *SubmitVLExtractionTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitVLExtractionTaskResponseBody) Validate() error {
	return dara.Validate(s)
}

type SubmitVLExtractionTaskResponseBodyData struct {
	// Task ID.
	//
	// example:
	//
	// 411ce93a-7eb5-40cf-836a-53c32f097663
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s SubmitVLExtractionTaskResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s SubmitVLExtractionTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *SubmitVLExtractionTaskResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *SubmitVLExtractionTaskResponseBodyData) SetTaskId(v string) *SubmitVLExtractionTaskResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *SubmitVLExtractionTaskResponseBodyData) Validate() error {
	return dara.Validate(s)
}
