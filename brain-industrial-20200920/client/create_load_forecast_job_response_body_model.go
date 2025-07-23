// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLoadForecastJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateLoadForecastJobResponseBody
	GetCode() *string
	SetData(v *CreateLoadForecastJobResponseBodyData) *CreateLoadForecastJobResponseBody
	GetData() *CreateLoadForecastJobResponseBodyData
	SetMessage(v string) *CreateLoadForecastJobResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateLoadForecastJobResponseBody
	GetRequestId() *string
	SetSuccess(v string) *CreateLoadForecastJobResponseBody
	GetSuccess() *string
}

type CreateLoadForecastJobResponseBody struct {
	// example:
	//
	// 200
	Code *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *CreateLoadForecastJobResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s CreateLoadForecastJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateLoadForecastJobResponseBody) GoString() string {
	return s.String()
}

func (s *CreateLoadForecastJobResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateLoadForecastJobResponseBody) GetData() *CreateLoadForecastJobResponseBodyData {
	return s.Data
}

func (s *CreateLoadForecastJobResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateLoadForecastJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateLoadForecastJobResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *CreateLoadForecastJobResponseBody) SetCode(v string) *CreateLoadForecastJobResponseBody {
	s.Code = &v
	return s
}

func (s *CreateLoadForecastJobResponseBody) SetData(v *CreateLoadForecastJobResponseBodyData) *CreateLoadForecastJobResponseBody {
	s.Data = v
	return s
}

func (s *CreateLoadForecastJobResponseBody) SetMessage(v string) *CreateLoadForecastJobResponseBody {
	s.Message = &v
	return s
}

func (s *CreateLoadForecastJobResponseBody) SetRequestId(v string) *CreateLoadForecastJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateLoadForecastJobResponseBody) SetSuccess(v string) *CreateLoadForecastJobResponseBody {
	s.Success = &v
	return s
}

func (s *CreateLoadForecastJobResponseBody) Validate() error {
	return dara.Validate(s)
}

type CreateLoadForecastJobResponseBodyData struct {
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
	Progress *int32                                         `json:"Progress,omitempty" xml:"Progress,omitempty"`
	Response *CreateLoadForecastJobResponseBodyDataResponse `json:"Response,omitempty" xml:"Response,omitempty" type:"Struct"`
	// example:
	//
	// RUNNING
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s CreateLoadForecastJobResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateLoadForecastJobResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateLoadForecastJobResponseBodyData) GetCompleted() *bool {
	return s.Completed
}

func (s *CreateLoadForecastJobResponseBodyData) GetCreateTime() *string {
	return s.CreateTime
}

func (s *CreateLoadForecastJobResponseBodyData) GetError() *string {
	return s.Error
}

func (s *CreateLoadForecastJobResponseBodyData) GetJobId() *string {
	return s.JobId
}

func (s *CreateLoadForecastJobResponseBodyData) GetProgress() *int32 {
	return s.Progress
}

func (s *CreateLoadForecastJobResponseBodyData) GetResponse() *CreateLoadForecastJobResponseBodyDataResponse {
	return s.Response
}

func (s *CreateLoadForecastJobResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *CreateLoadForecastJobResponseBodyData) SetCompleted(v bool) *CreateLoadForecastJobResponseBodyData {
	s.Completed = &v
	return s
}

func (s *CreateLoadForecastJobResponseBodyData) SetCreateTime(v string) *CreateLoadForecastJobResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *CreateLoadForecastJobResponseBodyData) SetError(v string) *CreateLoadForecastJobResponseBodyData {
	s.Error = &v
	return s
}

func (s *CreateLoadForecastJobResponseBodyData) SetJobId(v string) *CreateLoadForecastJobResponseBodyData {
	s.JobId = &v
	return s
}

func (s *CreateLoadForecastJobResponseBodyData) SetProgress(v int32) *CreateLoadForecastJobResponseBodyData {
	s.Progress = &v
	return s
}

func (s *CreateLoadForecastJobResponseBodyData) SetResponse(v *CreateLoadForecastJobResponseBodyDataResponse) *CreateLoadForecastJobResponseBodyData {
	s.Response = v
	return s
}

func (s *CreateLoadForecastJobResponseBodyData) SetStatus(v string) *CreateLoadForecastJobResponseBodyData {
	s.Status = &v
	return s
}

func (s *CreateLoadForecastJobResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type CreateLoadForecastJobResponseBodyDataResponse struct {
	// example:
	//
	// {}
	DebugInfo interface{} `json:"DebugInfo,omitempty" xml:"DebugInfo,omitempty"`
	// example:
	//
	// PowerForecast
	JobType *string `json:"JobType,omitempty" xml:"JobType,omitempty"`
	// example:
	//
	// {}
	Result interface{} `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s CreateLoadForecastJobResponseBodyDataResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateLoadForecastJobResponseBodyDataResponse) GoString() string {
	return s.String()
}

func (s *CreateLoadForecastJobResponseBodyDataResponse) GetDebugInfo() interface{} {
	return s.DebugInfo
}

func (s *CreateLoadForecastJobResponseBodyDataResponse) GetJobType() *string {
	return s.JobType
}

func (s *CreateLoadForecastJobResponseBodyDataResponse) GetResult() interface{} {
	return s.Result
}

func (s *CreateLoadForecastJobResponseBodyDataResponse) SetDebugInfo(v interface{}) *CreateLoadForecastJobResponseBodyDataResponse {
	s.DebugInfo = v
	return s
}

func (s *CreateLoadForecastJobResponseBodyDataResponse) SetJobType(v string) *CreateLoadForecastJobResponseBodyDataResponse {
	s.JobType = &v
	return s
}

func (s *CreateLoadForecastJobResponseBodyDataResponse) SetResult(v interface{}) *CreateLoadForecastJobResponseBodyDataResponse {
	s.Result = v
	return s
}

func (s *CreateLoadForecastJobResponseBodyDataResponse) Validate() error {
	return dara.Validate(s)
}
