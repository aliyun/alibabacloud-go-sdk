// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetQueueConsumersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *GetQueueConsumersResponseBody
	GetCode() *int32
	SetData(v *GetQueueConsumersResponseBodyData) *GetQueueConsumersResponseBody
	GetData() *GetQueueConsumersResponseBodyData
	SetMessage(v string) *GetQueueConsumersResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetQueueConsumersResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetQueueConsumersResponseBody
	GetSuccess() *bool
}

type GetQueueConsumersResponseBody struct {
	Code      *int32                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GetQueueConsumersResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                            `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                              `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetQueueConsumersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetQueueConsumersResponseBody) GoString() string {
	return s.String()
}

func (s *GetQueueConsumersResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetQueueConsumersResponseBody) GetData() *GetQueueConsumersResponseBodyData {
	return s.Data
}

func (s *GetQueueConsumersResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetQueueConsumersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetQueueConsumersResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetQueueConsumersResponseBody) SetCode(v int32) *GetQueueConsumersResponseBody {
	s.Code = &v
	return s
}

func (s *GetQueueConsumersResponseBody) SetData(v *GetQueueConsumersResponseBodyData) *GetQueueConsumersResponseBody {
	s.Data = v
	return s
}

func (s *GetQueueConsumersResponseBody) SetMessage(v string) *GetQueueConsumersResponseBody {
	s.Message = &v
	return s
}

func (s *GetQueueConsumersResponseBody) SetRequestId(v string) *GetQueueConsumersResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetQueueConsumersResponseBody) SetSuccess(v bool) *GetQueueConsumersResponseBody {
	s.Success = &v
	return s
}

func (s *GetQueueConsumersResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetQueueConsumersResponseBodyData struct {
	CurrentPage *int32                                     `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize    *int32                                     `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TaskId      *string                                    `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	TotalCount  *int32                                     `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	VoList      []*GetQueueConsumersResponseBodyDataVoList `json:"VoList,omitempty" xml:"VoList,omitempty" type:"Repeated"`
}

func (s GetQueueConsumersResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetQueueConsumersResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetQueueConsumersResponseBodyData) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *GetQueueConsumersResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetQueueConsumersResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *GetQueueConsumersResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *GetQueueConsumersResponseBodyData) GetVoList() []*GetQueueConsumersResponseBodyDataVoList {
	return s.VoList
}

func (s *GetQueueConsumersResponseBodyData) SetCurrentPage(v int32) *GetQueueConsumersResponseBodyData {
	s.CurrentPage = &v
	return s
}

func (s *GetQueueConsumersResponseBodyData) SetPageSize(v int32) *GetQueueConsumersResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *GetQueueConsumersResponseBodyData) SetTaskId(v string) *GetQueueConsumersResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *GetQueueConsumersResponseBodyData) SetTotalCount(v int32) *GetQueueConsumersResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *GetQueueConsumersResponseBodyData) SetVoList(v []*GetQueueConsumersResponseBodyDataVoList) *GetQueueConsumersResponseBodyData {
	s.VoList = v
	return s
}

func (s *GetQueueConsumersResponseBodyData) Validate() error {
	if s.VoList != nil {
		for _, item := range s.VoList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetQueueConsumersResponseBodyDataVoList struct {
	ClientAddress *string `json:"ClientAddress,omitempty" xml:"ClientAddress,omitempty"`
	ConsumerTag   *string `json:"ConsumerTag,omitempty" xml:"ConsumerTag,omitempty"`
	GmtCreate     *int64  `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
}

func (s GetQueueConsumersResponseBodyDataVoList) String() string {
	return dara.Prettify(s)
}

func (s GetQueueConsumersResponseBodyDataVoList) GoString() string {
	return s.String()
}

func (s *GetQueueConsumersResponseBodyDataVoList) GetClientAddress() *string {
	return s.ClientAddress
}

func (s *GetQueueConsumersResponseBodyDataVoList) GetConsumerTag() *string {
	return s.ConsumerTag
}

func (s *GetQueueConsumersResponseBodyDataVoList) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *GetQueueConsumersResponseBodyDataVoList) SetClientAddress(v string) *GetQueueConsumersResponseBodyDataVoList {
	s.ClientAddress = &v
	return s
}

func (s *GetQueueConsumersResponseBodyDataVoList) SetConsumerTag(v string) *GetQueueConsumersResponseBodyDataVoList {
	s.ConsumerTag = &v
	return s
}

func (s *GetQueueConsumersResponseBodyDataVoList) SetGmtCreate(v int64) *GetQueueConsumersResponseBodyDataVoList {
	s.GmtCreate = &v
	return s
}

func (s *GetQueueConsumersResponseBodyDataVoList) Validate() error {
	return dara.Validate(s)
}
