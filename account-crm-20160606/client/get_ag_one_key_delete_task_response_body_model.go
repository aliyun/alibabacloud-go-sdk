// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAgOneKeyDeleteTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetAgOneKeyDeleteTaskResponseBody
	GetCode() *string
	SetMessage(v string) *GetAgOneKeyDeleteTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetAgOneKeyDeleteTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetAgOneKeyDeleteTaskResponseBody
	GetSuccess() *bool
	SetTaskDto(v *GetAgOneKeyDeleteTaskResponseBodyTaskDto) *GetAgOneKeyDeleteTaskResponseBody
	GetTaskDto() *GetAgOneKeyDeleteTaskResponseBodyTaskDto
}

type GetAgOneKeyDeleteTaskResponseBody struct {
	Code      *string                                   `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                                   `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                     `json:"Success,omitempty" xml:"Success,omitempty"`
	TaskDto   *GetAgOneKeyDeleteTaskResponseBodyTaskDto `json:"TaskDto,omitempty" xml:"TaskDto,omitempty" type:"Struct"`
}

func (s GetAgOneKeyDeleteTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAgOneKeyDeleteTaskResponseBody) GoString() string {
	return s.String()
}

func (s *GetAgOneKeyDeleteTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetAgOneKeyDeleteTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetAgOneKeyDeleteTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAgOneKeyDeleteTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetAgOneKeyDeleteTaskResponseBody) GetTaskDto() *GetAgOneKeyDeleteTaskResponseBodyTaskDto {
	return s.TaskDto
}

func (s *GetAgOneKeyDeleteTaskResponseBody) SetCode(v string) *GetAgOneKeyDeleteTaskResponseBody {
	s.Code = &v
	return s
}

func (s *GetAgOneKeyDeleteTaskResponseBody) SetMessage(v string) *GetAgOneKeyDeleteTaskResponseBody {
	s.Message = &v
	return s
}

func (s *GetAgOneKeyDeleteTaskResponseBody) SetRequestId(v string) *GetAgOneKeyDeleteTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAgOneKeyDeleteTaskResponseBody) SetSuccess(v bool) *GetAgOneKeyDeleteTaskResponseBody {
	s.Success = &v
	return s
}

func (s *GetAgOneKeyDeleteTaskResponseBody) SetTaskDto(v *GetAgOneKeyDeleteTaskResponseBodyTaskDto) *GetAgOneKeyDeleteTaskResponseBody {
	s.TaskDto = v
	return s
}

func (s *GetAgOneKeyDeleteTaskResponseBody) Validate() error {
	if s.TaskDto != nil {
		if err := s.TaskDto.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAgOneKeyDeleteTaskResponseBodyTaskDto struct {
	DeleteStatus       *string `json:"DeleteStatus,omitempty" xml:"DeleteStatus,omitempty"`
	ExistQuietPeriod   *bool   `json:"ExistQuietPeriod,omitempty" xml:"ExistQuietPeriod,omitempty"`
	QuietPeriodEndTime *string `json:"QuietPeriodEndTime,omitempty" xml:"QuietPeriodEndTime,omitempty"`
}

func (s GetAgOneKeyDeleteTaskResponseBodyTaskDto) String() string {
	return dara.Prettify(s)
}

func (s GetAgOneKeyDeleteTaskResponseBodyTaskDto) GoString() string {
	return s.String()
}

func (s *GetAgOneKeyDeleteTaskResponseBodyTaskDto) GetDeleteStatus() *string {
	return s.DeleteStatus
}

func (s *GetAgOneKeyDeleteTaskResponseBodyTaskDto) GetExistQuietPeriod() *bool {
	return s.ExistQuietPeriod
}

func (s *GetAgOneKeyDeleteTaskResponseBodyTaskDto) GetQuietPeriodEndTime() *string {
	return s.QuietPeriodEndTime
}

func (s *GetAgOneKeyDeleteTaskResponseBodyTaskDto) SetDeleteStatus(v string) *GetAgOneKeyDeleteTaskResponseBodyTaskDto {
	s.DeleteStatus = &v
	return s
}

func (s *GetAgOneKeyDeleteTaskResponseBodyTaskDto) SetExistQuietPeriod(v bool) *GetAgOneKeyDeleteTaskResponseBodyTaskDto {
	s.ExistQuietPeriod = &v
	return s
}

func (s *GetAgOneKeyDeleteTaskResponseBodyTaskDto) SetQuietPeriodEndTime(v string) *GetAgOneKeyDeleteTaskResponseBodyTaskDto {
	s.QuietPeriodEndTime = &v
	return s
}

func (s *GetAgOneKeyDeleteTaskResponseBodyTaskDto) Validate() error {
	return dara.Validate(s)
}
