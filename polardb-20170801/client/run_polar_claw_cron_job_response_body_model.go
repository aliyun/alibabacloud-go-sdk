// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunPolarClawCronJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *RunPolarClawCronJobResponseBody
	GetApplicationId() *string
	SetCode(v int32) *RunPolarClawCronJobResponseBody
	GetCode() *int32
	SetJobId(v string) *RunPolarClawCronJobResponseBody
	GetJobId() *string
	SetMessage(v string) *RunPolarClawCronJobResponseBody
	GetMessage() *string
	SetOk(v bool) *RunPolarClawCronJobResponseBody
	GetOk() *bool
	SetRan(v bool) *RunPolarClawCronJobResponseBody
	GetRan() *bool
	SetRequestId(v string) *RunPolarClawCronJobResponseBody
	GetRequestId() *string
}

type RunPolarClawCronJobResponseBody struct {
	// example:
	//
	// pa-**************
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 8006e51c-dab3-4602-bc69-4f728002c6ce
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// true
	Ok *bool `json:"Ok,omitempty" xml:"Ok,omitempty"`
	// example:
	//
	// true
	Ran *bool `json:"Ran,omitempty" xml:"Ran,omitempty"`
	// Id of the request
	//
	// example:
	//
	// C61892A4-0850-4516-9E26-44D96C1782DE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RunPolarClawCronJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RunPolarClawCronJobResponseBody) GoString() string {
	return s.String()
}

func (s *RunPolarClawCronJobResponseBody) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *RunPolarClawCronJobResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *RunPolarClawCronJobResponseBody) GetJobId() *string {
	return s.JobId
}

func (s *RunPolarClawCronJobResponseBody) GetMessage() *string {
	return s.Message
}

func (s *RunPolarClawCronJobResponseBody) GetOk() *bool {
	return s.Ok
}

func (s *RunPolarClawCronJobResponseBody) GetRan() *bool {
	return s.Ran
}

func (s *RunPolarClawCronJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RunPolarClawCronJobResponseBody) SetApplicationId(v string) *RunPolarClawCronJobResponseBody {
	s.ApplicationId = &v
	return s
}

func (s *RunPolarClawCronJobResponseBody) SetCode(v int32) *RunPolarClawCronJobResponseBody {
	s.Code = &v
	return s
}

func (s *RunPolarClawCronJobResponseBody) SetJobId(v string) *RunPolarClawCronJobResponseBody {
	s.JobId = &v
	return s
}

func (s *RunPolarClawCronJobResponseBody) SetMessage(v string) *RunPolarClawCronJobResponseBody {
	s.Message = &v
	return s
}

func (s *RunPolarClawCronJobResponseBody) SetOk(v bool) *RunPolarClawCronJobResponseBody {
	s.Ok = &v
	return s
}

func (s *RunPolarClawCronJobResponseBody) SetRan(v bool) *RunPolarClawCronJobResponseBody {
	s.Ran = &v
	return s
}

func (s *RunPolarClawCronJobResponseBody) SetRequestId(v string) *RunPolarClawCronJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *RunPolarClawCronJobResponseBody) Validate() error {
	return dara.Validate(s)
}
