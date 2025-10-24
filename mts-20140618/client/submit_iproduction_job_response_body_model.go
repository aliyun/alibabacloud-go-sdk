// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitIProductionJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *SubmitIProductionJobResponseBody
	GetJobId() *string
	SetRequestId(v string) *SubmitIProductionJobResponseBody
	GetRequestId() *string
	SetResult(v string) *SubmitIProductionJobResponseBody
	GetResult() *string
}

type SubmitIProductionJobResponseBody struct {
	// example:
	//
	// 39f8e0bc005e4f309379701645f4****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// example:
	//
	// 5210DBB0-E327-4D45-ADBC-0B83C8796E26
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// { "Code":"Success", "Details":[], "FunctionName":"ImageCartoonize", "JobId":"39f8e0bc005e4f309379701645f4****", "Message":"success", "State":"Success", "Type":"IProduction" }
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s SubmitIProductionJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitIProductionJobResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitIProductionJobResponseBody) GetJobId() *string {
	return s.JobId
}

func (s *SubmitIProductionJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitIProductionJobResponseBody) GetResult() *string {
	return s.Result
}

func (s *SubmitIProductionJobResponseBody) SetJobId(v string) *SubmitIProductionJobResponseBody {
	s.JobId = &v
	return s
}

func (s *SubmitIProductionJobResponseBody) SetRequestId(v string) *SubmitIProductionJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitIProductionJobResponseBody) SetResult(v string) *SubmitIProductionJobResponseBody {
	s.Result = &v
	return s
}

func (s *SubmitIProductionJobResponseBody) Validate() error {
	return dara.Validate(s)
}
