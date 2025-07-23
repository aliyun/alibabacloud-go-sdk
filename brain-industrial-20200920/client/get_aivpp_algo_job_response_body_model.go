// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAivppAlgoJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetAivppAlgoJobResponseBody
	GetCode() *string
	SetData(v *GetAivppAlgoJobResponseBodyData) *GetAivppAlgoJobResponseBody
	GetData() *GetAivppAlgoJobResponseBodyData
	SetMessage(v string) *GetAivppAlgoJobResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetAivppAlgoJobResponseBody
	GetRequestId() *string
	SetSuccess(v string) *GetAivppAlgoJobResponseBody
	GetSuccess() *string
}

type GetAivppAlgoJobResponseBody struct {
	// example:
	//
	// 200
	Code *string                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetAivppAlgoJobResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s GetAivppAlgoJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAivppAlgoJobResponseBody) GoString() string {
	return s.String()
}

func (s *GetAivppAlgoJobResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetAivppAlgoJobResponseBody) GetData() *GetAivppAlgoJobResponseBodyData {
	return s.Data
}

func (s *GetAivppAlgoJobResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetAivppAlgoJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAivppAlgoJobResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *GetAivppAlgoJobResponseBody) SetCode(v string) *GetAivppAlgoJobResponseBody {
	s.Code = &v
	return s
}

func (s *GetAivppAlgoJobResponseBody) SetData(v *GetAivppAlgoJobResponseBodyData) *GetAivppAlgoJobResponseBody {
	s.Data = v
	return s
}

func (s *GetAivppAlgoJobResponseBody) SetMessage(v string) *GetAivppAlgoJobResponseBody {
	s.Message = &v
	return s
}

func (s *GetAivppAlgoJobResponseBody) SetRequestId(v string) *GetAivppAlgoJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAivppAlgoJobResponseBody) SetSuccess(v string) *GetAivppAlgoJobResponseBody {
	s.Success = &v
	return s
}

func (s *GetAivppAlgoJobResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetAivppAlgoJobResponseBodyData struct {
	// example:
	//
	// False
	Completed *bool `json:"Completed,omitempty" xml:"Completed,omitempty"`
	// example:
	//
	// 2024-12-10 17:50:48
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// ""
	Error *string `json:"Error,omitempty" xml:"Error,omitempty"`
	// example:
	//
	// 5854bfa6-f002-43c2-8e1d-e9b2c28f9384
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// example:
	//
	// 100
	Progress *int32                                   `json:"Progress,omitempty" xml:"Progress,omitempty"`
	Response *GetAivppAlgoJobResponseBodyDataResponse `json:"Response,omitempty" xml:"Response,omitempty" type:"Struct"`
	// example:
	//
	// RUNNING
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetAivppAlgoJobResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetAivppAlgoJobResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAivppAlgoJobResponseBodyData) GetCompleted() *bool {
	return s.Completed
}

func (s *GetAivppAlgoJobResponseBodyData) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetAivppAlgoJobResponseBodyData) GetError() *string {
	return s.Error
}

func (s *GetAivppAlgoJobResponseBodyData) GetJobId() *string {
	return s.JobId
}

func (s *GetAivppAlgoJobResponseBodyData) GetProgress() *int32 {
	return s.Progress
}

func (s *GetAivppAlgoJobResponseBodyData) GetResponse() *GetAivppAlgoJobResponseBodyDataResponse {
	return s.Response
}

func (s *GetAivppAlgoJobResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *GetAivppAlgoJobResponseBodyData) SetCompleted(v bool) *GetAivppAlgoJobResponseBodyData {
	s.Completed = &v
	return s
}

func (s *GetAivppAlgoJobResponseBodyData) SetCreateTime(v string) *GetAivppAlgoJobResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *GetAivppAlgoJobResponseBodyData) SetError(v string) *GetAivppAlgoJobResponseBodyData {
	s.Error = &v
	return s
}

func (s *GetAivppAlgoJobResponseBodyData) SetJobId(v string) *GetAivppAlgoJobResponseBodyData {
	s.JobId = &v
	return s
}

func (s *GetAivppAlgoJobResponseBodyData) SetProgress(v int32) *GetAivppAlgoJobResponseBodyData {
	s.Progress = &v
	return s
}

func (s *GetAivppAlgoJobResponseBodyData) SetResponse(v *GetAivppAlgoJobResponseBodyDataResponse) *GetAivppAlgoJobResponseBodyData {
	s.Response = v
	return s
}

func (s *GetAivppAlgoJobResponseBodyData) SetStatus(v string) *GetAivppAlgoJobResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetAivppAlgoJobResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type GetAivppAlgoJobResponseBodyDataResponse struct {
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

func (s GetAivppAlgoJobResponseBodyDataResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAivppAlgoJobResponseBodyDataResponse) GoString() string {
	return s.String()
}

func (s *GetAivppAlgoJobResponseBodyDataResponse) GetDebugInfo() interface{} {
	return s.DebugInfo
}

func (s *GetAivppAlgoJobResponseBodyDataResponse) GetJobType() *string {
	return s.JobType
}

func (s *GetAivppAlgoJobResponseBodyDataResponse) GetResult() interface{} {
	return s.Result
}

func (s *GetAivppAlgoJobResponseBodyDataResponse) SetDebugInfo(v interface{}) *GetAivppAlgoJobResponseBodyDataResponse {
	s.DebugInfo = v
	return s
}

func (s *GetAivppAlgoJobResponseBodyDataResponse) SetJobType(v string) *GetAivppAlgoJobResponseBodyDataResponse {
	s.JobType = &v
	return s
}

func (s *GetAivppAlgoJobResponseBodyDataResponse) SetResult(v interface{}) *GetAivppAlgoJobResponseBodyDataResponse {
	s.Result = v
	return s
}

func (s *GetAivppAlgoJobResponseBodyDataResponse) Validate() error {
	return dara.Validate(s)
}
