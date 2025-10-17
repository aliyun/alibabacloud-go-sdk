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
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// -
	Data *CreateJobResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// Parameter error: content is null.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 3835AA29-2298-5434-BC53-9CC377CDFD2C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateJobResponseBodyData struct {
	// example:
	//
	// 74
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
