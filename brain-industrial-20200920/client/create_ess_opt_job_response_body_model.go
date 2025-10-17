// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEssOptJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateEssOptJobResponseBody
	GetCode() *string
	SetData(v *CreateEssOptJobResponseBodyData) *CreateEssOptJobResponseBody
	GetData() *CreateEssOptJobResponseBodyData
	SetMessage(v string) *CreateEssOptJobResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateEssOptJobResponseBody
	GetRequestId() *string
	SetSuccess(v string) *CreateEssOptJobResponseBody
	GetSuccess() *string
}

type CreateEssOptJobResponseBody struct {
	// example:
	//
	// 200
	Code *string                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *CreateEssOptJobResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 68738E75-43C1-5AE5-9F3A-AFEF576D7B5F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateEssOptJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateEssOptJobResponseBody) GoString() string {
	return s.String()
}

func (s *CreateEssOptJobResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateEssOptJobResponseBody) GetData() *CreateEssOptJobResponseBodyData {
	return s.Data
}

func (s *CreateEssOptJobResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateEssOptJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateEssOptJobResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *CreateEssOptJobResponseBody) SetCode(v string) *CreateEssOptJobResponseBody {
	s.Code = &v
	return s
}

func (s *CreateEssOptJobResponseBody) SetData(v *CreateEssOptJobResponseBodyData) *CreateEssOptJobResponseBody {
	s.Data = v
	return s
}

func (s *CreateEssOptJobResponseBody) SetMessage(v string) *CreateEssOptJobResponseBody {
	s.Message = &v
	return s
}

func (s *CreateEssOptJobResponseBody) SetRequestId(v string) *CreateEssOptJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateEssOptJobResponseBody) SetSuccess(v string) *CreateEssOptJobResponseBody {
	s.Success = &v
	return s
}

func (s *CreateEssOptJobResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateEssOptJobResponseBodyData struct {
	// example:
	//
	// True
	Completed *bool `json:"Completed,omitempty" xml:"Completed,omitempty"`
	// example:
	//
	// 2024-12-22 00:00:21
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// ""
	Error *string `json:"Error,omitempty" xml:"Error,omitempty"`
	// example:
	//
	// 8c0ca18a-246a-4acd-80ca-e16d8ff5ef33
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// example:
	//
	// 25
	Progress *int32                                   `json:"Progress,omitempty" xml:"Progress,omitempty"`
	Response *CreateEssOptJobResponseBodyDataResponse `json:"Response,omitempty" xml:"Response,omitempty" type:"Struct"`
	// example:
	//
	// RUNNING
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s CreateEssOptJobResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateEssOptJobResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateEssOptJobResponseBodyData) GetCompleted() *bool {
	return s.Completed
}

func (s *CreateEssOptJobResponseBodyData) GetCreateTime() *string {
	return s.CreateTime
}

func (s *CreateEssOptJobResponseBodyData) GetError() *string {
	return s.Error
}

func (s *CreateEssOptJobResponseBodyData) GetJobId() *string {
	return s.JobId
}

func (s *CreateEssOptJobResponseBodyData) GetProgress() *int32 {
	return s.Progress
}

func (s *CreateEssOptJobResponseBodyData) GetResponse() *CreateEssOptJobResponseBodyDataResponse {
	return s.Response
}

func (s *CreateEssOptJobResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *CreateEssOptJobResponseBodyData) SetCompleted(v bool) *CreateEssOptJobResponseBodyData {
	s.Completed = &v
	return s
}

func (s *CreateEssOptJobResponseBodyData) SetCreateTime(v string) *CreateEssOptJobResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *CreateEssOptJobResponseBodyData) SetError(v string) *CreateEssOptJobResponseBodyData {
	s.Error = &v
	return s
}

func (s *CreateEssOptJobResponseBodyData) SetJobId(v string) *CreateEssOptJobResponseBodyData {
	s.JobId = &v
	return s
}

func (s *CreateEssOptJobResponseBodyData) SetProgress(v int32) *CreateEssOptJobResponseBodyData {
	s.Progress = &v
	return s
}

func (s *CreateEssOptJobResponseBodyData) SetResponse(v *CreateEssOptJobResponseBodyDataResponse) *CreateEssOptJobResponseBodyData {
	s.Response = v
	return s
}

func (s *CreateEssOptJobResponseBodyData) SetStatus(v string) *CreateEssOptJobResponseBodyData {
	s.Status = &v
	return s
}

func (s *CreateEssOptJobResponseBodyData) Validate() error {
	if s.Response != nil {
		if err := s.Response.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateEssOptJobResponseBodyDataResponse struct {
	// example:
	//
	// {}
	DebugInfo interface{} `json:"DebugInfo,omitempty" xml:"DebugInfo,omitempty"`
	// example:
	//
	// LoadForecast
	JobType *string `json:"JobType,omitempty" xml:"JobType,omitempty"`
	// example:
	//
	// {}
	Result interface{} `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s CreateEssOptJobResponseBodyDataResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateEssOptJobResponseBodyDataResponse) GoString() string {
	return s.String()
}

func (s *CreateEssOptJobResponseBodyDataResponse) GetDebugInfo() interface{} {
	return s.DebugInfo
}

func (s *CreateEssOptJobResponseBodyDataResponse) GetJobType() *string {
	return s.JobType
}

func (s *CreateEssOptJobResponseBodyDataResponse) GetResult() interface{} {
	return s.Result
}

func (s *CreateEssOptJobResponseBodyDataResponse) SetDebugInfo(v interface{}) *CreateEssOptJobResponseBodyDataResponse {
	s.DebugInfo = v
	return s
}

func (s *CreateEssOptJobResponseBodyDataResponse) SetJobType(v string) *CreateEssOptJobResponseBodyDataResponse {
	s.JobType = &v
	return s
}

func (s *CreateEssOptJobResponseBodyDataResponse) SetResult(v interface{}) *CreateEssOptJobResponseBodyDataResponse {
	s.Result = v
	return s
}

func (s *CreateEssOptJobResponseBodyDataResponse) Validate() error {
	return dara.Validate(s)
}
