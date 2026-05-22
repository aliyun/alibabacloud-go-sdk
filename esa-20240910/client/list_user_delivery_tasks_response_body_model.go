// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserDeliveryTasksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListUserDeliveryTasksResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListUserDeliveryTasksResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListUserDeliveryTasksResponseBody
	GetRequestId() *string
	SetTasks(v []*ListUserDeliveryTasksResponseBodyTasks) *ListUserDeliveryTasksResponseBody
	GetTasks() []*ListUserDeliveryTasksResponseBodyTasks
	SetTotalCount(v int32) *ListUserDeliveryTasksResponseBody
	GetTotalCount() *int32
}

type ListUserDeliveryTasksResponseBody struct {
	PageNumber *int32                                    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Tasks      []*ListUserDeliveryTasksResponseBodyTasks `json:"Tasks,omitempty" xml:"Tasks,omitempty" type:"Repeated"`
	TotalCount *int32                                    `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListUserDeliveryTasksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListUserDeliveryTasksResponseBody) GoString() string {
	return s.String()
}

func (s *ListUserDeliveryTasksResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListUserDeliveryTasksResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListUserDeliveryTasksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListUserDeliveryTasksResponseBody) GetTasks() []*ListUserDeliveryTasksResponseBodyTasks {
	return s.Tasks
}

func (s *ListUserDeliveryTasksResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListUserDeliveryTasksResponseBody) SetPageNumber(v int32) *ListUserDeliveryTasksResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListUserDeliveryTasksResponseBody) SetPageSize(v int32) *ListUserDeliveryTasksResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListUserDeliveryTasksResponseBody) SetRequestId(v string) *ListUserDeliveryTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListUserDeliveryTasksResponseBody) SetTasks(v []*ListUserDeliveryTasksResponseBodyTasks) *ListUserDeliveryTasksResponseBody {
	s.Tasks = v
	return s
}

func (s *ListUserDeliveryTasksResponseBody) SetTotalCount(v int32) *ListUserDeliveryTasksResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListUserDeliveryTasksResponseBody) Validate() error {
	if s.Tasks != nil {
		for _, item := range s.Tasks {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListUserDeliveryTasksResponseBodyTasks struct {
	BusinessType *string `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	DataCenter   *string `json:"DataCenter,omitempty" xml:"DataCenter,omitempty"`
	DeliveryType *string `json:"DeliveryType,omitempty" xml:"DeliveryType,omitempty"`
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
	TaskName     *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
}

func (s ListUserDeliveryTasksResponseBodyTasks) String() string {
	return dara.Prettify(s)
}

func (s ListUserDeliveryTasksResponseBodyTasks) GoString() string {
	return s.String()
}

func (s *ListUserDeliveryTasksResponseBodyTasks) GetBusinessType() *string {
	return s.BusinessType
}

func (s *ListUserDeliveryTasksResponseBodyTasks) GetDataCenter() *string {
	return s.DataCenter
}

func (s *ListUserDeliveryTasksResponseBodyTasks) GetDeliveryType() *string {
	return s.DeliveryType
}

func (s *ListUserDeliveryTasksResponseBodyTasks) GetStatus() *string {
	return s.Status
}

func (s *ListUserDeliveryTasksResponseBodyTasks) GetTaskName() *string {
	return s.TaskName
}

func (s *ListUserDeliveryTasksResponseBodyTasks) SetBusinessType(v string) *ListUserDeliveryTasksResponseBodyTasks {
	s.BusinessType = &v
	return s
}

func (s *ListUserDeliveryTasksResponseBodyTasks) SetDataCenter(v string) *ListUserDeliveryTasksResponseBodyTasks {
	s.DataCenter = &v
	return s
}

func (s *ListUserDeliveryTasksResponseBodyTasks) SetDeliveryType(v string) *ListUserDeliveryTasksResponseBodyTasks {
	s.DeliveryType = &v
	return s
}

func (s *ListUserDeliveryTasksResponseBodyTasks) SetStatus(v string) *ListUserDeliveryTasksResponseBodyTasks {
	s.Status = &v
	return s
}

func (s *ListUserDeliveryTasksResponseBodyTasks) SetTaskName(v string) *ListUserDeliveryTasksResponseBodyTasks {
	s.TaskName = &v
	return s
}

func (s *ListUserDeliveryTasksResponseBodyTasks) Validate() error {
	return dara.Validate(s)
}
