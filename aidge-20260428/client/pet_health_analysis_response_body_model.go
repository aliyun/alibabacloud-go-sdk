// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPetHealthAnalysisResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *PetHealthAnalysisResponseBody
	GetCode() *string
	SetData(v *PetHealthAnalysisResponseBodyData) *PetHealthAnalysisResponseBody
	GetData() *PetHealthAnalysisResponseBodyData
	SetMessage(v string) *PetHealthAnalysisResponseBody
	GetMessage() *string
	SetRequestId(v string) *PetHealthAnalysisResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *PetHealthAnalysisResponseBody
	GetSuccess() *bool
}

type PetHealthAnalysisResponseBody struct {
	// The error code.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The asynchronous task submit status.
	Data *PetHealthAnalysisResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The response message or failure description.
	//
	// example:
	//
	// Task submitted
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 70CBEFDF-BB17-1EB3-8A21-569F3124738F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call is successful.
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s PetHealthAnalysisResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PetHealthAnalysisResponseBody) GoString() string {
	return s.String()
}

func (s *PetHealthAnalysisResponseBody) GetCode() *string {
	return s.Code
}

func (s *PetHealthAnalysisResponseBody) GetData() *PetHealthAnalysisResponseBodyData {
	return s.Data
}

func (s *PetHealthAnalysisResponseBody) GetMessage() *string {
	return s.Message
}

func (s *PetHealthAnalysisResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PetHealthAnalysisResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *PetHealthAnalysisResponseBody) SetCode(v string) *PetHealthAnalysisResponseBody {
	s.Code = &v
	return s
}

func (s *PetHealthAnalysisResponseBody) SetData(v *PetHealthAnalysisResponseBodyData) *PetHealthAnalysisResponseBody {
	s.Data = v
	return s
}

func (s *PetHealthAnalysisResponseBody) SetMessage(v string) *PetHealthAnalysisResponseBody {
	s.Message = &v
	return s
}

func (s *PetHealthAnalysisResponseBody) SetRequestId(v string) *PetHealthAnalysisResponseBody {
	s.RequestId = &v
	return s
}

func (s *PetHealthAnalysisResponseBody) SetSuccess(v bool) *PetHealthAnalysisResponseBody {
	s.Success = &v
	return s
}

func (s *PetHealthAnalysisResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type PetHealthAnalysisResponseBodyData struct {
	// The asynchronous task ID used for QueryAsyncTaskResult queries.
	//
	// example:
	//
	// task_778fa8bd21804828a5d147050e30edac
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s PetHealthAnalysisResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s PetHealthAnalysisResponseBodyData) GoString() string {
	return s.String()
}

func (s *PetHealthAnalysisResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *PetHealthAnalysisResponseBodyData) SetTaskId(v string) *PetHealthAnalysisResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *PetHealthAnalysisResponseBodyData) Validate() error {
	return dara.Validate(s)
}
