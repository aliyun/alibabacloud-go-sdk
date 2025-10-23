// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitDocExtractionTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *SubmitDocExtractionTaskResponseBodyData) *SubmitDocExtractionTaskResponseBody
	GetData() *SubmitDocExtractionTaskResponseBodyData
	SetRequestId(v string) *SubmitDocExtractionTaskResponseBody
	GetRequestId() *string
}

type SubmitDocExtractionTaskResponseBody struct {
	// Returned data
	Data *SubmitDocExtractionTaskResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// Request ID.
	//
	// example:
	//
	// 83A5A7DD-8974-5769-952E-590A97BEA34E
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s SubmitDocExtractionTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitDocExtractionTaskResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitDocExtractionTaskResponseBody) GetData() *SubmitDocExtractionTaskResponseBodyData {
	return s.Data
}

func (s *SubmitDocExtractionTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitDocExtractionTaskResponseBody) SetData(v *SubmitDocExtractionTaskResponseBodyData) *SubmitDocExtractionTaskResponseBody {
	s.Data = v
	return s
}

func (s *SubmitDocExtractionTaskResponseBody) SetRequestId(v string) *SubmitDocExtractionTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitDocExtractionTaskResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SubmitDocExtractionTaskResponseBodyData struct {
	// Task ID.
	//
	// example:
	//
	// 864773ec-d35b-4c36-8871-52d07fbe806d
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s SubmitDocExtractionTaskResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s SubmitDocExtractionTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *SubmitDocExtractionTaskResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *SubmitDocExtractionTaskResponseBodyData) SetTaskId(v string) *SubmitDocExtractionTaskResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *SubmitDocExtractionTaskResponseBodyData) Validate() error {
	return dara.Validate(s)
}
