// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryIProductionJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFunctionName(v string) *QueryIProductionJobResponseBody
	GetFunctionName() *string
	SetInput(v string) *QueryIProductionJobResponseBody
	GetInput() *string
	SetJobId(v string) *QueryIProductionJobResponseBody
	GetJobId() *string
	SetJobParams(v string) *QueryIProductionJobResponseBody
	GetJobParams() *string
	SetOutput(v string) *QueryIProductionJobResponseBody
	GetOutput() *string
	SetPipelineId(v string) *QueryIProductionJobResponseBody
	GetPipelineId() *string
	SetRequestId(v string) *QueryIProductionJobResponseBody
	GetRequestId() *string
	SetResult(v string) *QueryIProductionJobResponseBody
	GetResult() *string
	SetState(v string) *QueryIProductionJobResponseBody
	GetState() *string
	SetUserData(v string) *QueryIProductionJobResponseBody
	GetUserData() *string
}

type QueryIProductionJobResponseBody struct {
	// example:
	//
	// ImageCartoonize
	FunctionName *string `json:"FunctionName,omitempty" xml:"FunctionName,omitempty"`
	// example:
	//
	// oss://example-****.oss-cn-hangzhou.aliyuncs.com/example.mp4
	Input *string `json:"Input,omitempty" xml:"Input,omitempty"`
	// example:
	//
	// 88c6ca184c0e432bbf5b665e2a15****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// example:
	//
	// {mode:"gif"}
	JobParams *string `json:"JobParams,omitempty" xml:"JobParams,omitempty"`
	// example:
	//
	// oss://example-****.oss-cn-hangzhou.aliyuncs.com/iproduction/{source}-{timestamp}-{sequenceId}.srt
	Output *string `json:"Output,omitempty" xml:"Output,omitempty"`
	// example:
	//
	// 39f8e0bc005e4f309379701645f4****
	PipelineId *string `json:"PipelineId,omitempty" xml:"PipelineId,omitempty"`
	// example:
	//
	// D127C68E-F1A1-4CE5-A874-8FF724881A12
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// {"Code":"Success","Data":"{\\"result\\":[{\\"file\\":\\"iproduction/test-result.jpg\\"},{\\"file\\":\\"iproduction/test-origin.jpg\\"}]}","Message":"Successful."}
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
	// example:
	//
	// Success
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// example:
	//
	// null
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s QueryIProductionJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryIProductionJobResponseBody) GoString() string {
	return s.String()
}

func (s *QueryIProductionJobResponseBody) GetFunctionName() *string {
	return s.FunctionName
}

func (s *QueryIProductionJobResponseBody) GetInput() *string {
	return s.Input
}

func (s *QueryIProductionJobResponseBody) GetJobId() *string {
	return s.JobId
}

func (s *QueryIProductionJobResponseBody) GetJobParams() *string {
	return s.JobParams
}

func (s *QueryIProductionJobResponseBody) GetOutput() *string {
	return s.Output
}

func (s *QueryIProductionJobResponseBody) GetPipelineId() *string {
	return s.PipelineId
}

func (s *QueryIProductionJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryIProductionJobResponseBody) GetResult() *string {
	return s.Result
}

func (s *QueryIProductionJobResponseBody) GetState() *string {
	return s.State
}

func (s *QueryIProductionJobResponseBody) GetUserData() *string {
	return s.UserData
}

func (s *QueryIProductionJobResponseBody) SetFunctionName(v string) *QueryIProductionJobResponseBody {
	s.FunctionName = &v
	return s
}

func (s *QueryIProductionJobResponseBody) SetInput(v string) *QueryIProductionJobResponseBody {
	s.Input = &v
	return s
}

func (s *QueryIProductionJobResponseBody) SetJobId(v string) *QueryIProductionJobResponseBody {
	s.JobId = &v
	return s
}

func (s *QueryIProductionJobResponseBody) SetJobParams(v string) *QueryIProductionJobResponseBody {
	s.JobParams = &v
	return s
}

func (s *QueryIProductionJobResponseBody) SetOutput(v string) *QueryIProductionJobResponseBody {
	s.Output = &v
	return s
}

func (s *QueryIProductionJobResponseBody) SetPipelineId(v string) *QueryIProductionJobResponseBody {
	s.PipelineId = &v
	return s
}

func (s *QueryIProductionJobResponseBody) SetRequestId(v string) *QueryIProductionJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryIProductionJobResponseBody) SetResult(v string) *QueryIProductionJobResponseBody {
	s.Result = &v
	return s
}

func (s *QueryIProductionJobResponseBody) SetState(v string) *QueryIProductionJobResponseBody {
	s.State = &v
	return s
}

func (s *QueryIProductionJobResponseBody) SetUserData(v string) *QueryIProductionJobResponseBody {
	s.UserData = &v
	return s
}

func (s *QueryIProductionJobResponseBody) Validate() error {
	return dara.Validate(s)
}
