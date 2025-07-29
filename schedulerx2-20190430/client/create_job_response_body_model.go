// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *CreateJobResponseBody
	GetCode() *int32
	SetData(v *CreateJobResponseBodyData) *CreateJobResponseBody
	GetData() *CreateJobResponseBodyData
	SetMessage(v string) *CreateJobResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateJobResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateJobResponseBody
	GetSuccess() *bool
}

type CreateJobResponseBody struct {
	// The HTTP status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The details of the job.
	Data *CreateJobResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The additional information returned.
	//
	// example:
	//
	// message
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 39090022-1F3B-4797-8518-6B61095F1AF0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// If you set JobType to k8s, this parameter is required. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateJobResponseBody) GoString() string {
	return s.String()
}

func (s *CreateJobResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *CreateJobResponseBody) GetData() *CreateJobResponseBodyData {
	return s.Data
}

func (s *CreateJobResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateJobResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateJobResponseBody) SetCode(v int32) *CreateJobResponseBody {
	s.Code = &v
	return s
}

func (s *CreateJobResponseBody) SetData(v *CreateJobResponseBodyData) *CreateJobResponseBody {
	s.Data = v
	return s
}

func (s *CreateJobResponseBody) SetMessage(v string) *CreateJobResponseBody {
	s.Message = &v
	return s
}

func (s *CreateJobResponseBody) SetRequestId(v string) *CreateJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateJobResponseBody) SetSuccess(v bool) *CreateJobResponseBody {
	s.Success = &v
	return s
}

func (s *CreateJobResponseBody) Validate() error {
	return dara.Validate(s)
}

type CreateJobResponseBodyData struct {
	// The job ID.
	//
	// example:
	//
	// 92583
	JobId *int64 `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s CreateJobResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateJobResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateJobResponseBodyData) GetJobId() *int64 {
	return s.JobId
}

func (s *CreateJobResponseBodyData) SetJobId(v int64) *CreateJobResponseBodyData {
	s.JobId = &v
	return s
}

func (s *CreateJobResponseBodyData) Validate() error {
	return dara.Validate(s)
}
