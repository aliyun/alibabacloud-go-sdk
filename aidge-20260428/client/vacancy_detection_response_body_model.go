// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVacancyDetectionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *VacancyDetectionResponseBody
	GetCode() *string
	SetData(v *VacancyDetectionResponseBodyData) *VacancyDetectionResponseBody
	GetData() *VacancyDetectionResponseBodyData
	SetMessage(v string) *VacancyDetectionResponseBody
	GetMessage() *string
	SetRequestId(v string) *VacancyDetectionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *VacancyDetectionResponseBody
	GetSuccess() *bool
}

type VacancyDetectionResponseBody struct {
	// The error code.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The submission result of the asynchronous task.
	//
	// example:
	//
	// {\\"TaskId\\": \\"TID_2f567233-78dd-4226-8b69-ac9288a66da0\\"}
	Data *VacancyDetectionResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message or failure description.
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The unique ID of the request.
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

func (s VacancyDetectionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s VacancyDetectionResponseBody) GoString() string {
	return s.String()
}

func (s *VacancyDetectionResponseBody) GetCode() *string {
	return s.Code
}

func (s *VacancyDetectionResponseBody) GetData() *VacancyDetectionResponseBodyData {
	return s.Data
}

func (s *VacancyDetectionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *VacancyDetectionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *VacancyDetectionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *VacancyDetectionResponseBody) SetCode(v string) *VacancyDetectionResponseBody {
	s.Code = &v
	return s
}

func (s *VacancyDetectionResponseBody) SetData(v *VacancyDetectionResponseBodyData) *VacancyDetectionResponseBody {
	s.Data = v
	return s
}

func (s *VacancyDetectionResponseBody) SetMessage(v string) *VacancyDetectionResponseBody {
	s.Message = &v
	return s
}

func (s *VacancyDetectionResponseBody) SetRequestId(v string) *VacancyDetectionResponseBody {
	s.RequestId = &v
	return s
}

func (s *VacancyDetectionResponseBody) SetSuccess(v bool) *VacancyDetectionResponseBody {
	s.Success = &v
	return s
}

func (s *VacancyDetectionResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type VacancyDetectionResponseBodyData struct {
	// The ID of the asynchronous task. You can use this ID to call QueryAsyncTaskResult to query the result later.
	//
	// example:
	//
	// 518c7318185d4db694236f79d672c156
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The usage details. This parameter is left empty during asynchronous commit and is populated by scheduling after the node reaches the desired state.
	//
	// example:
	//
	// {"ProcessingCount":1}
	UsageMap map[string]*int64 `json:"UsageMap,omitempty" xml:"UsageMap,omitempty"`
}

func (s VacancyDetectionResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s VacancyDetectionResponseBodyData) GoString() string {
	return s.String()
}

func (s *VacancyDetectionResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *VacancyDetectionResponseBodyData) GetUsageMap() map[string]*int64 {
	return s.UsageMap
}

func (s *VacancyDetectionResponseBodyData) SetTaskId(v string) *VacancyDetectionResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *VacancyDetectionResponseBodyData) SetUsageMap(v map[string]*int64) *VacancyDetectionResponseBodyData {
	s.UsageMap = v
	return s
}

func (s *VacancyDetectionResponseBodyData) Validate() error {
	return dara.Validate(s)
}
