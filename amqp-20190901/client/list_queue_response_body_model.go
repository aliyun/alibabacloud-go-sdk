// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListQueueResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ListQueueResponseBody
	GetCode() *int32
	SetData(v *ListQueueResponseBodyData) *ListQueueResponseBody
	GetData() *ListQueueResponseBodyData
	SetMessage(v string) *ListQueueResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListQueueResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListQueueResponseBody
	GetSuccess() *bool
}

type ListQueueResponseBody struct {
	Code      *int32                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *ListQueueResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                    `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                      `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListQueueResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListQueueResponseBody) GoString() string {
	return s.String()
}

func (s *ListQueueResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListQueueResponseBody) GetData() *ListQueueResponseBodyData {
	return s.Data
}

func (s *ListQueueResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListQueueResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListQueueResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListQueueResponseBody) SetCode(v int32) *ListQueueResponseBody {
	s.Code = &v
	return s
}

func (s *ListQueueResponseBody) SetData(v *ListQueueResponseBodyData) *ListQueueResponseBody {
	s.Data = v
	return s
}

func (s *ListQueueResponseBody) SetMessage(v string) *ListQueueResponseBody {
	s.Message = &v
	return s
}

func (s *ListQueueResponseBody) SetRequestId(v string) *ListQueueResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListQueueResponseBody) SetSuccess(v bool) *ListQueueResponseBody {
	s.Success = &v
	return s
}

func (s *ListQueueResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListQueueResponseBodyData struct {
	CurrentPage *int32                           `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize    *int32                           `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount  *int64                           `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	VoList      *ListQueueResponseBodyDataVoList `json:"VoList,omitempty" xml:"VoList,omitempty" type:"Struct"`
}

func (s ListQueueResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListQueueResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListQueueResponseBodyData) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListQueueResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListQueueResponseBodyData) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListQueueResponseBodyData) GetVoList() *ListQueueResponseBodyDataVoList {
	return s.VoList
}

func (s *ListQueueResponseBodyData) SetCurrentPage(v int32) *ListQueueResponseBodyData {
	s.CurrentPage = &v
	return s
}

func (s *ListQueueResponseBodyData) SetPageSize(v int32) *ListQueueResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListQueueResponseBodyData) SetTotalCount(v int64) *ListQueueResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListQueueResponseBodyData) SetVoList(v *ListQueueResponseBodyDataVoList) *ListQueueResponseBodyData {
	s.VoList = v
	return s
}

func (s *ListQueueResponseBodyData) Validate() error {
	if s.VoList != nil {
		if err := s.VoList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListQueueResponseBodyDataVoList struct {
	QueueVO []*ListQueueResponseBodyDataVoListQueueVO `json:"QueueVO,omitempty" xml:"QueueVO,omitempty" type:"Repeated"`
}

func (s ListQueueResponseBodyDataVoList) String() string {
	return dara.Prettify(s)
}

func (s ListQueueResponseBodyDataVoList) GoString() string {
	return s.String()
}

func (s *ListQueueResponseBodyDataVoList) GetQueueVO() []*ListQueueResponseBodyDataVoListQueueVO {
	return s.QueueVO
}

func (s *ListQueueResponseBodyDataVoList) SetQueueVO(v []*ListQueueResponseBodyDataVoListQueueVO) *ListQueueResponseBodyDataVoList {
	s.QueueVO = v
	return s
}

func (s *ListQueueResponseBodyDataVoList) Validate() error {
	if s.QueueVO != nil {
		for _, item := range s.QueueVO {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListQueueResponseBodyDataVoListQueueVO struct {
	AccumulationCount *int64                 `json:"AccumulationCount,omitempty" xml:"AccumulationCount,omitempty"`
	Attributes        map[string]interface{} `json:"Attributes,omitempty" xml:"Attributes,omitempty"`
	AutoDelete        *bool                  `json:"AutoDelete,omitempty" xml:"AutoDelete,omitempty"`
	CanDelete         *bool                  `json:"CanDelete,omitempty" xml:"CanDelete,omitempty"`
	CreateTime        *int64                 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Durable           *bool                  `json:"Durable,omitempty" xml:"Durable,omitempty"`
	Exclusive         *bool                  `json:"Exclusive,omitempty" xml:"Exclusive,omitempty"`
	LastConsumeTime   *int64                 `json:"LastConsumeTime,omitempty" xml:"LastConsumeTime,omitempty"`
	Name              *string                `json:"Name,omitempty" xml:"Name,omitempty"`
	VhostName         *string                `json:"VhostName,omitempty" xml:"VhostName,omitempty"`
}

func (s ListQueueResponseBodyDataVoListQueueVO) String() string {
	return dara.Prettify(s)
}

func (s ListQueueResponseBodyDataVoListQueueVO) GoString() string {
	return s.String()
}

func (s *ListQueueResponseBodyDataVoListQueueVO) GetAccumulationCount() *int64 {
	return s.AccumulationCount
}

func (s *ListQueueResponseBodyDataVoListQueueVO) GetAttributes() map[string]interface{} {
	return s.Attributes
}

func (s *ListQueueResponseBodyDataVoListQueueVO) GetAutoDelete() *bool {
	return s.AutoDelete
}

func (s *ListQueueResponseBodyDataVoListQueueVO) GetCanDelete() *bool {
	return s.CanDelete
}

func (s *ListQueueResponseBodyDataVoListQueueVO) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListQueueResponseBodyDataVoListQueueVO) GetDurable() *bool {
	return s.Durable
}

func (s *ListQueueResponseBodyDataVoListQueueVO) GetExclusive() *bool {
	return s.Exclusive
}

func (s *ListQueueResponseBodyDataVoListQueueVO) GetLastConsumeTime() *int64 {
	return s.LastConsumeTime
}

func (s *ListQueueResponseBodyDataVoListQueueVO) GetName() *string {
	return s.Name
}

func (s *ListQueueResponseBodyDataVoListQueueVO) GetVhostName() *string {
	return s.VhostName
}

func (s *ListQueueResponseBodyDataVoListQueueVO) SetAccumulationCount(v int64) *ListQueueResponseBodyDataVoListQueueVO {
	s.AccumulationCount = &v
	return s
}

func (s *ListQueueResponseBodyDataVoListQueueVO) SetAttributes(v map[string]interface{}) *ListQueueResponseBodyDataVoListQueueVO {
	s.Attributes = v
	return s
}

func (s *ListQueueResponseBodyDataVoListQueueVO) SetAutoDelete(v bool) *ListQueueResponseBodyDataVoListQueueVO {
	s.AutoDelete = &v
	return s
}

func (s *ListQueueResponseBodyDataVoListQueueVO) SetCanDelete(v bool) *ListQueueResponseBodyDataVoListQueueVO {
	s.CanDelete = &v
	return s
}

func (s *ListQueueResponseBodyDataVoListQueueVO) SetCreateTime(v int64) *ListQueueResponseBodyDataVoListQueueVO {
	s.CreateTime = &v
	return s
}

func (s *ListQueueResponseBodyDataVoListQueueVO) SetDurable(v bool) *ListQueueResponseBodyDataVoListQueueVO {
	s.Durable = &v
	return s
}

func (s *ListQueueResponseBodyDataVoListQueueVO) SetExclusive(v bool) *ListQueueResponseBodyDataVoListQueueVO {
	s.Exclusive = &v
	return s
}

func (s *ListQueueResponseBodyDataVoListQueueVO) SetLastConsumeTime(v int64) *ListQueueResponseBodyDataVoListQueueVO {
	s.LastConsumeTime = &v
	return s
}

func (s *ListQueueResponseBodyDataVoListQueueVO) SetName(v string) *ListQueueResponseBodyDataVoListQueueVO {
	s.Name = &v
	return s
}

func (s *ListQueueResponseBodyDataVoListQueueVO) SetVhostName(v string) *ListQueueResponseBodyDataVoListQueueVO {
	s.VhostName = &v
	return s
}

func (s *ListQueueResponseBodyDataVoListQueueVO) Validate() error {
	return dara.Validate(s)
}
