// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePowerForecastByFileUrlJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreatePowerForecastByFileUrlJobResponseBody
	GetCode() *string
	SetData(v *CreatePowerForecastByFileUrlJobResponseBodyData) *CreatePowerForecastByFileUrlJobResponseBody
	GetData() *CreatePowerForecastByFileUrlJobResponseBodyData
	SetMessage(v string) *CreatePowerForecastByFileUrlJobResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreatePowerForecastByFileUrlJobResponseBody
	GetRequestId() *string
	SetSuccess(v string) *CreatePowerForecastByFileUrlJobResponseBody
	GetSuccess() *string
}

type CreatePowerForecastByFileUrlJobResponseBody struct {
	// example:
	//
	// 200
	Code *string                                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *CreatePowerForecastByFileUrlJobResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 68738E75-43C1-5AE5-9F3A-AFEF576D7B5F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreatePowerForecastByFileUrlJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreatePowerForecastByFileUrlJobResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePowerForecastByFileUrlJobResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreatePowerForecastByFileUrlJobResponseBody) GetData() *CreatePowerForecastByFileUrlJobResponseBodyData {
	return s.Data
}

func (s *CreatePowerForecastByFileUrlJobResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreatePowerForecastByFileUrlJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreatePowerForecastByFileUrlJobResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *CreatePowerForecastByFileUrlJobResponseBody) SetCode(v string) *CreatePowerForecastByFileUrlJobResponseBody {
	s.Code = &v
	return s
}

func (s *CreatePowerForecastByFileUrlJobResponseBody) SetData(v *CreatePowerForecastByFileUrlJobResponseBodyData) *CreatePowerForecastByFileUrlJobResponseBody {
	s.Data = v
	return s
}

func (s *CreatePowerForecastByFileUrlJobResponseBody) SetMessage(v string) *CreatePowerForecastByFileUrlJobResponseBody {
	s.Message = &v
	return s
}

func (s *CreatePowerForecastByFileUrlJobResponseBody) SetRequestId(v string) *CreatePowerForecastByFileUrlJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreatePowerForecastByFileUrlJobResponseBody) SetSuccess(v string) *CreatePowerForecastByFileUrlJobResponseBody {
	s.Success = &v
	return s
}

func (s *CreatePowerForecastByFileUrlJobResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreatePowerForecastByFileUrlJobResponseBodyData struct {
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
	// 100
	Progress *int32                                                   `json:"Progress,omitempty" xml:"Progress,omitempty"`
	Response *CreatePowerForecastByFileUrlJobResponseBodyDataResponse `json:"Response,omitempty" xml:"Response,omitempty" type:"Struct"`
	// example:
	//
	// RUNNING
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s CreatePowerForecastByFileUrlJobResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreatePowerForecastByFileUrlJobResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreatePowerForecastByFileUrlJobResponseBodyData) GetCompleted() *bool {
	return s.Completed
}

func (s *CreatePowerForecastByFileUrlJobResponseBodyData) GetCreateTime() *string {
	return s.CreateTime
}

func (s *CreatePowerForecastByFileUrlJobResponseBodyData) GetError() *string {
	return s.Error
}

func (s *CreatePowerForecastByFileUrlJobResponseBodyData) GetJobId() *string {
	return s.JobId
}

func (s *CreatePowerForecastByFileUrlJobResponseBodyData) GetProgress() *int32 {
	return s.Progress
}

func (s *CreatePowerForecastByFileUrlJobResponseBodyData) GetResponse() *CreatePowerForecastByFileUrlJobResponseBodyDataResponse {
	return s.Response
}

func (s *CreatePowerForecastByFileUrlJobResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *CreatePowerForecastByFileUrlJobResponseBodyData) SetCompleted(v bool) *CreatePowerForecastByFileUrlJobResponseBodyData {
	s.Completed = &v
	return s
}

func (s *CreatePowerForecastByFileUrlJobResponseBodyData) SetCreateTime(v string) *CreatePowerForecastByFileUrlJobResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *CreatePowerForecastByFileUrlJobResponseBodyData) SetError(v string) *CreatePowerForecastByFileUrlJobResponseBodyData {
	s.Error = &v
	return s
}

func (s *CreatePowerForecastByFileUrlJobResponseBodyData) SetJobId(v string) *CreatePowerForecastByFileUrlJobResponseBodyData {
	s.JobId = &v
	return s
}

func (s *CreatePowerForecastByFileUrlJobResponseBodyData) SetProgress(v int32) *CreatePowerForecastByFileUrlJobResponseBodyData {
	s.Progress = &v
	return s
}

func (s *CreatePowerForecastByFileUrlJobResponseBodyData) SetResponse(v *CreatePowerForecastByFileUrlJobResponseBodyDataResponse) *CreatePowerForecastByFileUrlJobResponseBodyData {
	s.Response = v
	return s
}

func (s *CreatePowerForecastByFileUrlJobResponseBodyData) SetStatus(v string) *CreatePowerForecastByFileUrlJobResponseBodyData {
	s.Status = &v
	return s
}

func (s *CreatePowerForecastByFileUrlJobResponseBodyData) Validate() error {
	if s.Response != nil {
		if err := s.Response.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreatePowerForecastByFileUrlJobResponseBodyDataResponse struct {
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
	// {"runTime": \\"2025-01-01 00:00:00\\", "value": 0.0}
	Result interface{} `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s CreatePowerForecastByFileUrlJobResponseBodyDataResponse) String() string {
	return dara.Prettify(s)
}

func (s CreatePowerForecastByFileUrlJobResponseBodyDataResponse) GoString() string {
	return s.String()
}

func (s *CreatePowerForecastByFileUrlJobResponseBodyDataResponse) GetDebugInfo() interface{} {
	return s.DebugInfo
}

func (s *CreatePowerForecastByFileUrlJobResponseBodyDataResponse) GetJobType() *string {
	return s.JobType
}

func (s *CreatePowerForecastByFileUrlJobResponseBodyDataResponse) GetResult() interface{} {
	return s.Result
}

func (s *CreatePowerForecastByFileUrlJobResponseBodyDataResponse) SetDebugInfo(v interface{}) *CreatePowerForecastByFileUrlJobResponseBodyDataResponse {
	s.DebugInfo = v
	return s
}

func (s *CreatePowerForecastByFileUrlJobResponseBodyDataResponse) SetJobType(v string) *CreatePowerForecastByFileUrlJobResponseBodyDataResponse {
	s.JobType = &v
	return s
}

func (s *CreatePowerForecastByFileUrlJobResponseBodyDataResponse) SetResult(v interface{}) *CreatePowerForecastByFileUrlJobResponseBodyDataResponse {
	s.Result = v
	return s
}

func (s *CreatePowerForecastByFileUrlJobResponseBodyDataResponse) Validate() error {
	return dara.Validate(s)
}
