// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMmsTimerLogsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ListMmsTimerLogsResponseBodyData) *ListMmsTimerLogsResponseBody
	GetData() *ListMmsTimerLogsResponseBodyData
	SetRequestId(v string) *ListMmsTimerLogsResponseBody
	GetRequestId() *string
}

type ListMmsTimerLogsResponseBody struct {
	Data *ListMmsTimerLogsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// 0b87b7e716665825896565060e87a4
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListMmsTimerLogsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListMmsTimerLogsResponseBody) GoString() string {
	return s.String()
}

func (s *ListMmsTimerLogsResponseBody) GetData() *ListMmsTimerLogsResponseBodyData {
	return s.Data
}

func (s *ListMmsTimerLogsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListMmsTimerLogsResponseBody) SetData(v *ListMmsTimerLogsResponseBodyData) *ListMmsTimerLogsResponseBody {
	s.Data = v
	return s
}

func (s *ListMmsTimerLogsResponseBody) SetRequestId(v string) *ListMmsTimerLogsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMmsTimerLogsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListMmsTimerLogsResponseBodyData struct {
	ObjectList []*ListMmsTimerLogsResponseBodyDataObjectList `json:"objectList,omitempty" xml:"objectList,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNum *int32 `json:"pageNum,omitempty" xml:"pageNum,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// 100
	Total *int32 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListMmsTimerLogsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListMmsTimerLogsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListMmsTimerLogsResponseBodyData) GetObjectList() []*ListMmsTimerLogsResponseBodyDataObjectList {
	return s.ObjectList
}

func (s *ListMmsTimerLogsResponseBodyData) GetPageNum() *int32 {
	return s.PageNum
}

func (s *ListMmsTimerLogsResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListMmsTimerLogsResponseBodyData) GetTotal() *int32 {
	return s.Total
}

func (s *ListMmsTimerLogsResponseBodyData) SetObjectList(v []*ListMmsTimerLogsResponseBodyDataObjectList) *ListMmsTimerLogsResponseBodyData {
	s.ObjectList = v
	return s
}

func (s *ListMmsTimerLogsResponseBodyData) SetPageNum(v int32) *ListMmsTimerLogsResponseBodyData {
	s.PageNum = &v
	return s
}

func (s *ListMmsTimerLogsResponseBodyData) SetPageSize(v int32) *ListMmsTimerLogsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListMmsTimerLogsResponseBodyData) SetTotal(v int32) *ListMmsTimerLogsResponseBodyData {
	s.Total = &v
	return s
}

func (s *ListMmsTimerLogsResponseBodyData) Validate() error {
	if s.ObjectList != nil {
		for _, item := range s.ObjectList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListMmsTimerLogsResponseBodyDataObjectList struct {
	// example:
	//
	// start job
	Action *string `json:"action,omitempty" xml:"action,omitempty"`
	// example:
	//
	// 2024-12-17 15:44:17
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// 1003476
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// start job success
	Msg *string `json:"msg,omitempty" xml:"msg,omitempty"`
	// example:
	//
	// 200018
	SourceId *int64 `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	// example:
	//
	// DOING
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s ListMmsTimerLogsResponseBodyDataObjectList) String() string {
	return dara.Prettify(s)
}

func (s ListMmsTimerLogsResponseBodyDataObjectList) GoString() string {
	return s.String()
}

func (s *ListMmsTimerLogsResponseBodyDataObjectList) GetAction() *string {
	return s.Action
}

func (s *ListMmsTimerLogsResponseBodyDataObjectList) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListMmsTimerLogsResponseBodyDataObjectList) GetId() *int64 {
	return s.Id
}

func (s *ListMmsTimerLogsResponseBodyDataObjectList) GetMsg() *string {
	return s.Msg
}

func (s *ListMmsTimerLogsResponseBodyDataObjectList) GetSourceId() *int64 {
	return s.SourceId
}

func (s *ListMmsTimerLogsResponseBodyDataObjectList) GetStatus() *string {
	return s.Status
}

func (s *ListMmsTimerLogsResponseBodyDataObjectList) SetAction(v string) *ListMmsTimerLogsResponseBodyDataObjectList {
	s.Action = &v
	return s
}

func (s *ListMmsTimerLogsResponseBodyDataObjectList) SetCreateTime(v string) *ListMmsTimerLogsResponseBodyDataObjectList {
	s.CreateTime = &v
	return s
}

func (s *ListMmsTimerLogsResponseBodyDataObjectList) SetId(v int64) *ListMmsTimerLogsResponseBodyDataObjectList {
	s.Id = &v
	return s
}

func (s *ListMmsTimerLogsResponseBodyDataObjectList) SetMsg(v string) *ListMmsTimerLogsResponseBodyDataObjectList {
	s.Msg = &v
	return s
}

func (s *ListMmsTimerLogsResponseBodyDataObjectList) SetSourceId(v int64) *ListMmsTimerLogsResponseBodyDataObjectList {
	s.SourceId = &v
	return s
}

func (s *ListMmsTimerLogsResponseBodyDataObjectList) SetStatus(v string) *ListMmsTimerLogsResponseBodyDataObjectList {
	s.Status = &v
	return s
}

func (s *ListMmsTimerLogsResponseBodyDataObjectList) Validate() error {
	return dara.Validate(s)
}
