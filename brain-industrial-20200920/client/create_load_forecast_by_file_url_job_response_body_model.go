// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLoadForecastByFileUrlJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateLoadForecastByFileUrlJobResponseBody
	GetCode() *string
	SetData(v *CreateLoadForecastByFileUrlJobResponseBodyData) *CreateLoadForecastByFileUrlJobResponseBody
	GetData() *CreateLoadForecastByFileUrlJobResponseBodyData
	SetMessage(v string) *CreateLoadForecastByFileUrlJobResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateLoadForecastByFileUrlJobResponseBody
	GetRequestId() *string
	SetSuccess(v string) *CreateLoadForecastByFileUrlJobResponseBody
	GetSuccess() *string
}

type CreateLoadForecastByFileUrlJobResponseBody struct {
	// example:
	//
	// 200
	Code *string                                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *CreateLoadForecastByFileUrlJobResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s CreateLoadForecastByFileUrlJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateLoadForecastByFileUrlJobResponseBody) GoString() string {
	return s.String()
}

func (s *CreateLoadForecastByFileUrlJobResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateLoadForecastByFileUrlJobResponseBody) GetData() *CreateLoadForecastByFileUrlJobResponseBodyData {
	return s.Data
}

func (s *CreateLoadForecastByFileUrlJobResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateLoadForecastByFileUrlJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateLoadForecastByFileUrlJobResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *CreateLoadForecastByFileUrlJobResponseBody) SetCode(v string) *CreateLoadForecastByFileUrlJobResponseBody {
	s.Code = &v
	return s
}

func (s *CreateLoadForecastByFileUrlJobResponseBody) SetData(v *CreateLoadForecastByFileUrlJobResponseBodyData) *CreateLoadForecastByFileUrlJobResponseBody {
	s.Data = v
	return s
}

func (s *CreateLoadForecastByFileUrlJobResponseBody) SetMessage(v string) *CreateLoadForecastByFileUrlJobResponseBody {
	s.Message = &v
	return s
}

func (s *CreateLoadForecastByFileUrlJobResponseBody) SetRequestId(v string) *CreateLoadForecastByFileUrlJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateLoadForecastByFileUrlJobResponseBody) SetSuccess(v string) *CreateLoadForecastByFileUrlJobResponseBody {
	s.Success = &v
	return s
}

func (s *CreateLoadForecastByFileUrlJobResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateLoadForecastByFileUrlJobResponseBodyData struct {
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
	Progress *int32                                                  `json:"Progress,omitempty" xml:"Progress,omitempty"`
	Response *CreateLoadForecastByFileUrlJobResponseBodyDataResponse `json:"Response,omitempty" xml:"Response,omitempty" type:"Struct"`
	// example:
	//
	// RUNNING
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s CreateLoadForecastByFileUrlJobResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateLoadForecastByFileUrlJobResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateLoadForecastByFileUrlJobResponseBodyData) GetCompleted() *bool {
	return s.Completed
}

func (s *CreateLoadForecastByFileUrlJobResponseBodyData) GetCreateTime() *string {
	return s.CreateTime
}

func (s *CreateLoadForecastByFileUrlJobResponseBodyData) GetError() *string {
	return s.Error
}

func (s *CreateLoadForecastByFileUrlJobResponseBodyData) GetJobId() *string {
	return s.JobId
}

func (s *CreateLoadForecastByFileUrlJobResponseBodyData) GetProgress() *int32 {
	return s.Progress
}

func (s *CreateLoadForecastByFileUrlJobResponseBodyData) GetResponse() *CreateLoadForecastByFileUrlJobResponseBodyDataResponse {
	return s.Response
}

func (s *CreateLoadForecastByFileUrlJobResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *CreateLoadForecastByFileUrlJobResponseBodyData) SetCompleted(v bool) *CreateLoadForecastByFileUrlJobResponseBodyData {
	s.Completed = &v
	return s
}

func (s *CreateLoadForecastByFileUrlJobResponseBodyData) SetCreateTime(v string) *CreateLoadForecastByFileUrlJobResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *CreateLoadForecastByFileUrlJobResponseBodyData) SetError(v string) *CreateLoadForecastByFileUrlJobResponseBodyData {
	s.Error = &v
	return s
}

func (s *CreateLoadForecastByFileUrlJobResponseBodyData) SetJobId(v string) *CreateLoadForecastByFileUrlJobResponseBodyData {
	s.JobId = &v
	return s
}

func (s *CreateLoadForecastByFileUrlJobResponseBodyData) SetProgress(v int32) *CreateLoadForecastByFileUrlJobResponseBodyData {
	s.Progress = &v
	return s
}

func (s *CreateLoadForecastByFileUrlJobResponseBodyData) SetResponse(v *CreateLoadForecastByFileUrlJobResponseBodyDataResponse) *CreateLoadForecastByFileUrlJobResponseBodyData {
	s.Response = v
	return s
}

func (s *CreateLoadForecastByFileUrlJobResponseBodyData) SetStatus(v string) *CreateLoadForecastByFileUrlJobResponseBodyData {
	s.Status = &v
	return s
}

func (s *CreateLoadForecastByFileUrlJobResponseBodyData) Validate() error {
	if s.Response != nil {
		if err := s.Response.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateLoadForecastByFileUrlJobResponseBodyDataResponse struct {
	// example:
	//
	// null
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

func (s CreateLoadForecastByFileUrlJobResponseBodyDataResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateLoadForecastByFileUrlJobResponseBodyDataResponse) GoString() string {
	return s.String()
}

func (s *CreateLoadForecastByFileUrlJobResponseBodyDataResponse) GetDebugInfo() interface{} {
	return s.DebugInfo
}

func (s *CreateLoadForecastByFileUrlJobResponseBodyDataResponse) GetJobType() *string {
	return s.JobType
}

func (s *CreateLoadForecastByFileUrlJobResponseBodyDataResponse) GetResult() interface{} {
	return s.Result
}

func (s *CreateLoadForecastByFileUrlJobResponseBodyDataResponse) SetDebugInfo(v interface{}) *CreateLoadForecastByFileUrlJobResponseBodyDataResponse {
	s.DebugInfo = v
	return s
}

func (s *CreateLoadForecastByFileUrlJobResponseBodyDataResponse) SetJobType(v string) *CreateLoadForecastByFileUrlJobResponseBodyDataResponse {
	s.JobType = &v
	return s
}

func (s *CreateLoadForecastByFileUrlJobResponseBodyDataResponse) SetResult(v interface{}) *CreateLoadForecastByFileUrlJobResponseBodyDataResponse {
	s.Result = v
	return s
}

func (s *CreateLoadForecastByFileUrlJobResponseBodyDataResponse) Validate() error {
	return dara.Validate(s)
}
