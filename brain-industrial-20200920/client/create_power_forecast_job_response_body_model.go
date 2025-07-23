// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePowerForecastJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreatePowerForecastJobResponseBody
	GetCode() *string
	SetData(v *CreatePowerForecastJobResponseBodyData) *CreatePowerForecastJobResponseBody
	GetData() *CreatePowerForecastJobResponseBodyData
	SetMessage(v string) *CreatePowerForecastJobResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreatePowerForecastJobResponseBody
	GetRequestId() *string
	SetSuccess(v string) *CreatePowerForecastJobResponseBody
	GetSuccess() *string
}

type CreatePowerForecastJobResponseBody struct {
	// example:
	//
	// 200
	Code *string                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *CreatePowerForecastJobResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s CreatePowerForecastJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreatePowerForecastJobResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePowerForecastJobResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreatePowerForecastJobResponseBody) GetData() *CreatePowerForecastJobResponseBodyData {
	return s.Data
}

func (s *CreatePowerForecastJobResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreatePowerForecastJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreatePowerForecastJobResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *CreatePowerForecastJobResponseBody) SetCode(v string) *CreatePowerForecastJobResponseBody {
	s.Code = &v
	return s
}

func (s *CreatePowerForecastJobResponseBody) SetData(v *CreatePowerForecastJobResponseBodyData) *CreatePowerForecastJobResponseBody {
	s.Data = v
	return s
}

func (s *CreatePowerForecastJobResponseBody) SetMessage(v string) *CreatePowerForecastJobResponseBody {
	s.Message = &v
	return s
}

func (s *CreatePowerForecastJobResponseBody) SetRequestId(v string) *CreatePowerForecastJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreatePowerForecastJobResponseBody) SetSuccess(v string) *CreatePowerForecastJobResponseBody {
	s.Success = &v
	return s
}

func (s *CreatePowerForecastJobResponseBody) Validate() error {
	return dara.Validate(s)
}

type CreatePowerForecastJobResponseBodyData struct {
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
	Progress *int32                                          `json:"Progress,omitempty" xml:"Progress,omitempty"`
	Response *CreatePowerForecastJobResponseBodyDataResponse `json:"Response,omitempty" xml:"Response,omitempty" type:"Struct"`
	// example:
	//
	// RUNNING
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s CreatePowerForecastJobResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreatePowerForecastJobResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreatePowerForecastJobResponseBodyData) GetCompleted() *bool {
	return s.Completed
}

func (s *CreatePowerForecastJobResponseBodyData) GetCreateTime() *string {
	return s.CreateTime
}

func (s *CreatePowerForecastJobResponseBodyData) GetError() *string {
	return s.Error
}

func (s *CreatePowerForecastJobResponseBodyData) GetJobId() *string {
	return s.JobId
}

func (s *CreatePowerForecastJobResponseBodyData) GetProgress() *int32 {
	return s.Progress
}

func (s *CreatePowerForecastJobResponseBodyData) GetResponse() *CreatePowerForecastJobResponseBodyDataResponse {
	return s.Response
}

func (s *CreatePowerForecastJobResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *CreatePowerForecastJobResponseBodyData) SetCompleted(v bool) *CreatePowerForecastJobResponseBodyData {
	s.Completed = &v
	return s
}

func (s *CreatePowerForecastJobResponseBodyData) SetCreateTime(v string) *CreatePowerForecastJobResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *CreatePowerForecastJobResponseBodyData) SetError(v string) *CreatePowerForecastJobResponseBodyData {
	s.Error = &v
	return s
}

func (s *CreatePowerForecastJobResponseBodyData) SetJobId(v string) *CreatePowerForecastJobResponseBodyData {
	s.JobId = &v
	return s
}

func (s *CreatePowerForecastJobResponseBodyData) SetProgress(v int32) *CreatePowerForecastJobResponseBodyData {
	s.Progress = &v
	return s
}

func (s *CreatePowerForecastJobResponseBodyData) SetResponse(v *CreatePowerForecastJobResponseBodyDataResponse) *CreatePowerForecastJobResponseBodyData {
	s.Response = v
	return s
}

func (s *CreatePowerForecastJobResponseBodyData) SetStatus(v string) *CreatePowerForecastJobResponseBodyData {
	s.Status = &v
	return s
}

func (s *CreatePowerForecastJobResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type CreatePowerForecastJobResponseBodyDataResponse struct {
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

func (s CreatePowerForecastJobResponseBodyDataResponse) String() string {
	return dara.Prettify(s)
}

func (s CreatePowerForecastJobResponseBodyDataResponse) GoString() string {
	return s.String()
}

func (s *CreatePowerForecastJobResponseBodyDataResponse) GetDebugInfo() interface{} {
	return s.DebugInfo
}

func (s *CreatePowerForecastJobResponseBodyDataResponse) GetJobType() *string {
	return s.JobType
}

func (s *CreatePowerForecastJobResponseBodyDataResponse) GetResult() interface{} {
	return s.Result
}

func (s *CreatePowerForecastJobResponseBodyDataResponse) SetDebugInfo(v interface{}) *CreatePowerForecastJobResponseBodyDataResponse {
	s.DebugInfo = v
	return s
}

func (s *CreatePowerForecastJobResponseBodyDataResponse) SetJobType(v string) *CreatePowerForecastJobResponseBodyDataResponse {
	s.JobType = &v
	return s
}

func (s *CreatePowerForecastJobResponseBodyDataResponse) SetResult(v interface{}) *CreatePowerForecastJobResponseBodyDataResponse {
	s.Result = v
	return s
}

func (s *CreatePowerForecastJobResponseBodyDataResponse) Validate() error {
	return dara.Validate(s)
}
