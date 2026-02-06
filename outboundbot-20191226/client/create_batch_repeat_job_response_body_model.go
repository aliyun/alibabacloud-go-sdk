// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBatchRepeatJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateBatchRepeatJobResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *CreateBatchRepeatJobResponseBody
	GetHttpStatusCode() *int32
	SetJobGroup(v *CreateBatchRepeatJobResponseBodyJobGroup) *CreateBatchRepeatJobResponseBody
	GetJobGroup() *CreateBatchRepeatJobResponseBodyJobGroup
	SetMessage(v string) *CreateBatchRepeatJobResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateBatchRepeatJobResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateBatchRepeatJobResponseBody
	GetSuccess() *bool
}

type CreateBatchRepeatJobResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// {}
	JobGroup *CreateBatchRepeatJobResponseBodyJobGroup `json:"JobGroup,omitempty" xml:"JobGroup,omitempty" type:"Struct"`
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 254EB995-DEDF-48A4-9101-9CA5B72FFBCC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateBatchRepeatJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateBatchRepeatJobResponseBody) GoString() string {
	return s.String()
}

func (s *CreateBatchRepeatJobResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateBatchRepeatJobResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateBatchRepeatJobResponseBody) GetJobGroup() *CreateBatchRepeatJobResponseBodyJobGroup {
	return s.JobGroup
}

func (s *CreateBatchRepeatJobResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateBatchRepeatJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateBatchRepeatJobResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateBatchRepeatJobResponseBody) SetCode(v string) *CreateBatchRepeatJobResponseBody {
	s.Code = &v
	return s
}

func (s *CreateBatchRepeatJobResponseBody) SetHttpStatusCode(v int32) *CreateBatchRepeatJobResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateBatchRepeatJobResponseBody) SetJobGroup(v *CreateBatchRepeatJobResponseBodyJobGroup) *CreateBatchRepeatJobResponseBody {
	s.JobGroup = v
	return s
}

func (s *CreateBatchRepeatJobResponseBody) SetMessage(v string) *CreateBatchRepeatJobResponseBody {
	s.Message = &v
	return s
}

func (s *CreateBatchRepeatJobResponseBody) SetRequestId(v string) *CreateBatchRepeatJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateBatchRepeatJobResponseBody) SetSuccess(v bool) *CreateBatchRepeatJobResponseBody {
	s.Success = &v
	return s
}

func (s *CreateBatchRepeatJobResponseBody) Validate() error {
	if s.JobGroup != nil {
		if err := s.JobGroup.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateBatchRepeatJobResponseBodyJobGroup struct {
	// example:
	//
	// 254EB995-DEDF-48A4-9101-9CA5B72FFBCC
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 1
	MinConcurrency *int64 `json:"MinConcurrency,omitempty" xml:"MinConcurrency,omitempty"`
	// example:
	//
	// 3
	Priority *string `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// example:
	//
	// 30
	RingingDuration *int64 `json:"RingingDuration,omitempty" xml:"RingingDuration,omitempty"`
}

func (s CreateBatchRepeatJobResponseBodyJobGroup) String() string {
	return dara.Prettify(s)
}

func (s CreateBatchRepeatJobResponseBodyJobGroup) GoString() string {
	return s.String()
}

func (s *CreateBatchRepeatJobResponseBodyJobGroup) GetId() *string {
	return s.Id
}

func (s *CreateBatchRepeatJobResponseBodyJobGroup) GetMinConcurrency() *int64 {
	return s.MinConcurrency
}

func (s *CreateBatchRepeatJobResponseBodyJobGroup) GetPriority() *string {
	return s.Priority
}

func (s *CreateBatchRepeatJobResponseBodyJobGroup) GetRingingDuration() *int64 {
	return s.RingingDuration
}

func (s *CreateBatchRepeatJobResponseBodyJobGroup) SetId(v string) *CreateBatchRepeatJobResponseBodyJobGroup {
	s.Id = &v
	return s
}

func (s *CreateBatchRepeatJobResponseBodyJobGroup) SetMinConcurrency(v int64) *CreateBatchRepeatJobResponseBodyJobGroup {
	s.MinConcurrency = &v
	return s
}

func (s *CreateBatchRepeatJobResponseBodyJobGroup) SetPriority(v string) *CreateBatchRepeatJobResponseBodyJobGroup {
	s.Priority = &v
	return s
}

func (s *CreateBatchRepeatJobResponseBodyJobGroup) SetRingingDuration(v int64) *CreateBatchRepeatJobResponseBodyJobGroup {
	s.RingingDuration = &v
	return s
}

func (s *CreateBatchRepeatJobResponseBodyJobGroup) Validate() error {
	return dara.Validate(s)
}
