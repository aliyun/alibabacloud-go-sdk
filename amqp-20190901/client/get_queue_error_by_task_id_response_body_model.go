// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetQueueErrorByTaskIdResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *GetQueueErrorByTaskIdResponseBody
	GetCode() *int32
	SetData(v *GetQueueErrorByTaskIdResponseBodyData) *GetQueueErrorByTaskIdResponseBody
	GetData() *GetQueueErrorByTaskIdResponseBodyData
	SetMessage(v string) *GetQueueErrorByTaskIdResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetQueueErrorByTaskIdResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetQueueErrorByTaskIdResponseBody
	GetSuccess() *bool
}

type GetQueueErrorByTaskIdResponseBody struct {
	Code      *int32                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GetQueueErrorByTaskIdResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                  `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetQueueErrorByTaskIdResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetQueueErrorByTaskIdResponseBody) GoString() string {
	return s.String()
}

func (s *GetQueueErrorByTaskIdResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetQueueErrorByTaskIdResponseBody) GetData() *GetQueueErrorByTaskIdResponseBodyData {
	return s.Data
}

func (s *GetQueueErrorByTaskIdResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetQueueErrorByTaskIdResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetQueueErrorByTaskIdResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetQueueErrorByTaskIdResponseBody) SetCode(v int32) *GetQueueErrorByTaskIdResponseBody {
	s.Code = &v
	return s
}

func (s *GetQueueErrorByTaskIdResponseBody) SetData(v *GetQueueErrorByTaskIdResponseBodyData) *GetQueueErrorByTaskIdResponseBody {
	s.Data = v
	return s
}

func (s *GetQueueErrorByTaskIdResponseBody) SetMessage(v string) *GetQueueErrorByTaskIdResponseBody {
	s.Message = &v
	return s
}

func (s *GetQueueErrorByTaskIdResponseBody) SetRequestId(v string) *GetQueueErrorByTaskIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetQueueErrorByTaskIdResponseBody) SetSuccess(v bool) *GetQueueErrorByTaskIdResponseBody {
	s.Success = &v
	return s
}

func (s *GetQueueErrorByTaskIdResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetQueueErrorByTaskIdResponseBodyData struct {
	CurrentPage *int32                                       `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize    *int32                                       `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount  *int64                                       `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	VoList      *GetQueueErrorByTaskIdResponseBodyDataVoList `json:"VoList,omitempty" xml:"VoList,omitempty" type:"Struct"`
}

func (s GetQueueErrorByTaskIdResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetQueueErrorByTaskIdResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetQueueErrorByTaskIdResponseBodyData) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *GetQueueErrorByTaskIdResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetQueueErrorByTaskIdResponseBodyData) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *GetQueueErrorByTaskIdResponseBodyData) GetVoList() *GetQueueErrorByTaskIdResponseBodyDataVoList {
	return s.VoList
}

func (s *GetQueueErrorByTaskIdResponseBodyData) SetCurrentPage(v int32) *GetQueueErrorByTaskIdResponseBodyData {
	s.CurrentPage = &v
	return s
}

func (s *GetQueueErrorByTaskIdResponseBodyData) SetPageSize(v int32) *GetQueueErrorByTaskIdResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *GetQueueErrorByTaskIdResponseBodyData) SetTotalCount(v int64) *GetQueueErrorByTaskIdResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *GetQueueErrorByTaskIdResponseBodyData) SetVoList(v *GetQueueErrorByTaskIdResponseBodyDataVoList) *GetQueueErrorByTaskIdResponseBodyData {
	s.VoList = v
	return s
}

func (s *GetQueueErrorByTaskIdResponseBodyData) Validate() error {
	if s.VoList != nil {
		if err := s.VoList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetQueueErrorByTaskIdResponseBodyDataVoList struct {
	QueueErrorDO []*GetQueueErrorByTaskIdResponseBodyDataVoListQueueErrorDO `json:"QueueErrorDO,omitempty" xml:"QueueErrorDO,omitempty" type:"Repeated"`
}

func (s GetQueueErrorByTaskIdResponseBodyDataVoList) String() string {
	return dara.Prettify(s)
}

func (s GetQueueErrorByTaskIdResponseBodyDataVoList) GoString() string {
	return s.String()
}

func (s *GetQueueErrorByTaskIdResponseBodyDataVoList) GetQueueErrorDO() []*GetQueueErrorByTaskIdResponseBodyDataVoListQueueErrorDO {
	return s.QueueErrorDO
}

func (s *GetQueueErrorByTaskIdResponseBodyDataVoList) SetQueueErrorDO(v []*GetQueueErrorByTaskIdResponseBodyDataVoListQueueErrorDO) *GetQueueErrorByTaskIdResponseBodyDataVoList {
	s.QueueErrorDO = v
	return s
}

func (s *GetQueueErrorByTaskIdResponseBodyDataVoList) Validate() error {
	if s.QueueErrorDO != nil {
		for _, item := range s.QueueErrorDO {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetQueueErrorByTaskIdResponseBodyDataVoListQueueErrorDO struct {
	AutoDelete   *bool   `json:"AutoDelete,omitempty" xml:"AutoDelete,omitempty"`
	Durable      *bool   `json:"Durable,omitempty" xml:"Durable,omitempty"`
	ErrorMessage *int32  `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	QueueName    *string `json:"QueueName,omitempty" xml:"QueueName,omitempty"`
	TaskId       *int64  `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	VhostName    *string `json:"VhostName,omitempty" xml:"VhostName,omitempty"`
}

func (s GetQueueErrorByTaskIdResponseBodyDataVoListQueueErrorDO) String() string {
	return dara.Prettify(s)
}

func (s GetQueueErrorByTaskIdResponseBodyDataVoListQueueErrorDO) GoString() string {
	return s.String()
}

func (s *GetQueueErrorByTaskIdResponseBodyDataVoListQueueErrorDO) GetAutoDelete() *bool {
	return s.AutoDelete
}

func (s *GetQueueErrorByTaskIdResponseBodyDataVoListQueueErrorDO) GetDurable() *bool {
	return s.Durable
}

func (s *GetQueueErrorByTaskIdResponseBodyDataVoListQueueErrorDO) GetErrorMessage() *int32 {
	return s.ErrorMessage
}

func (s *GetQueueErrorByTaskIdResponseBodyDataVoListQueueErrorDO) GetQueueName() *string {
	return s.QueueName
}

func (s *GetQueueErrorByTaskIdResponseBodyDataVoListQueueErrorDO) GetTaskId() *int64 {
	return s.TaskId
}

func (s *GetQueueErrorByTaskIdResponseBodyDataVoListQueueErrorDO) GetVhostName() *string {
	return s.VhostName
}

func (s *GetQueueErrorByTaskIdResponseBodyDataVoListQueueErrorDO) SetAutoDelete(v bool) *GetQueueErrorByTaskIdResponseBodyDataVoListQueueErrorDO {
	s.AutoDelete = &v
	return s
}

func (s *GetQueueErrorByTaskIdResponseBodyDataVoListQueueErrorDO) SetDurable(v bool) *GetQueueErrorByTaskIdResponseBodyDataVoListQueueErrorDO {
	s.Durable = &v
	return s
}

func (s *GetQueueErrorByTaskIdResponseBodyDataVoListQueueErrorDO) SetErrorMessage(v int32) *GetQueueErrorByTaskIdResponseBodyDataVoListQueueErrorDO {
	s.ErrorMessage = &v
	return s
}

func (s *GetQueueErrorByTaskIdResponseBodyDataVoListQueueErrorDO) SetQueueName(v string) *GetQueueErrorByTaskIdResponseBodyDataVoListQueueErrorDO {
	s.QueueName = &v
	return s
}

func (s *GetQueueErrorByTaskIdResponseBodyDataVoListQueueErrorDO) SetTaskId(v int64) *GetQueueErrorByTaskIdResponseBodyDataVoListQueueErrorDO {
	s.TaskId = &v
	return s
}

func (s *GetQueueErrorByTaskIdResponseBodyDataVoListQueueErrorDO) SetVhostName(v string) *GetQueueErrorByTaskIdResponseBodyDataVoListQueueErrorDO {
	s.VhostName = &v
	return s
}

func (s *GetQueueErrorByTaskIdResponseBodyDataVoListQueueErrorDO) Validate() error {
	return dara.Validate(s)
}
