// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iProductVideoGenerationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ProductVideoGenerationResponseBody
	GetCode() *string
	SetData(v *ProductVideoGenerationResponseBodyData) *ProductVideoGenerationResponseBody
	GetData() *ProductVideoGenerationResponseBodyData
	SetMessage(v string) *ProductVideoGenerationResponseBody
	GetMessage() *string
	SetRequestId(v string) *ProductVideoGenerationResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ProductVideoGenerationResponseBody
	GetSuccess() *bool
}

type ProductVideoGenerationResponseBody struct {
	// The error code.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The task acceptance result.
	Data *ProductVideoGenerationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message or failure description.
	//
	// example:
	//
	// Task submitted
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 70CBEFDF-BB17-1EB3-8A21-569F3124738F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful.
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ProductVideoGenerationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ProductVideoGenerationResponseBody) GoString() string {
	return s.String()
}

func (s *ProductVideoGenerationResponseBody) GetCode() *string {
	return s.Code
}

func (s *ProductVideoGenerationResponseBody) GetData() *ProductVideoGenerationResponseBodyData {
	return s.Data
}

func (s *ProductVideoGenerationResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ProductVideoGenerationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ProductVideoGenerationResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ProductVideoGenerationResponseBody) SetCode(v string) *ProductVideoGenerationResponseBody {
	s.Code = &v
	return s
}

func (s *ProductVideoGenerationResponseBody) SetData(v *ProductVideoGenerationResponseBodyData) *ProductVideoGenerationResponseBody {
	s.Data = v
	return s
}

func (s *ProductVideoGenerationResponseBody) SetMessage(v string) *ProductVideoGenerationResponseBody {
	s.Message = &v
	return s
}

func (s *ProductVideoGenerationResponseBody) SetRequestId(v string) *ProductVideoGenerationResponseBody {
	s.RequestId = &v
	return s
}

func (s *ProductVideoGenerationResponseBody) SetSuccess(v bool) *ProductVideoGenerationResponseBody {
	s.Success = &v
	return s
}

func (s *ProductVideoGenerationResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ProductVideoGenerationResponseBodyData struct {
	// The planning route written upon acceptance. Valid values: DIRECT (with script) and AUTO_PLAN (automatic planning).
	//
	// example:
	//
	// AUTO_PLAN
	Route *string `json:"route,omitempty" xml:"route,omitempty"`
	// The current status of the task.
	//
	// example:
	//
	// PENDING
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The task ID of the service.
	//
	// example:
	//
	// task_778fa8bd21804828a5d147050e30edac
	TaskId *string `json:"task_id,omitempty" xml:"task_id,omitempty"`
}

func (s ProductVideoGenerationResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ProductVideoGenerationResponseBodyData) GoString() string {
	return s.String()
}

func (s *ProductVideoGenerationResponseBodyData) GetRoute() *string {
	return s.Route
}

func (s *ProductVideoGenerationResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *ProductVideoGenerationResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *ProductVideoGenerationResponseBodyData) SetRoute(v string) *ProductVideoGenerationResponseBodyData {
	s.Route = &v
	return s
}

func (s *ProductVideoGenerationResponseBodyData) SetStatus(v string) *ProductVideoGenerationResponseBodyData {
	s.Status = &v
	return s
}

func (s *ProductVideoGenerationResponseBodyData) SetTaskId(v string) *ProductVideoGenerationResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *ProductVideoGenerationResponseBodyData) Validate() error {
	return dara.Validate(s)
}
