// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMmsTimersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ListMmsTimersResponseBodyData) *ListMmsTimersResponseBody
	GetData() *ListMmsTimersResponseBodyData
	SetRequestId(v string) *ListMmsTimersResponseBody
	GetRequestId() *string
}

type ListMmsTimersResponseBody struct {
	Data *ListMmsTimersResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// 0b87b7e716665825896565060e87a4
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListMmsTimersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListMmsTimersResponseBody) GoString() string {
	return s.String()
}

func (s *ListMmsTimersResponseBody) GetData() *ListMmsTimersResponseBodyData {
	return s.Data
}

func (s *ListMmsTimersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListMmsTimersResponseBody) SetData(v *ListMmsTimersResponseBodyData) *ListMmsTimersResponseBody {
	s.Data = v
	return s
}

func (s *ListMmsTimersResponseBody) SetRequestId(v string) *ListMmsTimersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMmsTimersResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListMmsTimersResponseBodyData struct {
	ObjectList []*ListMmsTimersResponseBodyDataObjectList `json:"objectList,omitempty" xml:"objectList,omitempty" type:"Repeated"`
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
	// 13
	Total *int32 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListMmsTimersResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListMmsTimersResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListMmsTimersResponseBodyData) GetObjectList() []*ListMmsTimersResponseBodyDataObjectList {
	return s.ObjectList
}

func (s *ListMmsTimersResponseBodyData) GetPageNum() *int32 {
	return s.PageNum
}

func (s *ListMmsTimersResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListMmsTimersResponseBodyData) GetTotal() *int32 {
	return s.Total
}

func (s *ListMmsTimersResponseBodyData) SetObjectList(v []*ListMmsTimersResponseBodyDataObjectList) *ListMmsTimersResponseBodyData {
	s.ObjectList = v
	return s
}

func (s *ListMmsTimersResponseBodyData) SetPageNum(v int32) *ListMmsTimersResponseBodyData {
	s.PageNum = &v
	return s
}

func (s *ListMmsTimersResponseBodyData) SetPageSize(v int32) *ListMmsTimersResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListMmsTimersResponseBodyData) SetTotal(v int32) *ListMmsTimersResponseBodyData {
	s.Total = &v
	return s
}

func (s *ListMmsTimersResponseBodyData) Validate() error {
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

type ListMmsTimersResponseBodyDataObjectList struct {
	// example:
	//
	// 2024-12-17 09:29:58
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// 196
	DbId *int64 `json:"dbId,omitempty" xml:"dbId,omitempty"`
	// example:
	//
	// 18
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// migrate_db_1
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// Daily
	ScheduleType *string `json:"scheduleType,omitempty" xml:"scheduleType,omitempty"`
	// example:
	//
	// 2000015
	SourceId *int64 `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	// example:
	//
	// db_1
	SrcDbName *string `json:"srcDbName,omitempty" xml:"srcDbName,omitempty"`
	// example:
	//
	// false
	Stopped *bool `json:"stopped,omitempty" xml:"stopped,omitempty"`
	// example:
	//
	// TABLES
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// example:
	//
	// 00:00
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s ListMmsTimersResponseBodyDataObjectList) String() string {
	return dara.Prettify(s)
}

func (s ListMmsTimersResponseBodyDataObjectList) GoString() string {
	return s.String()
}

func (s *ListMmsTimersResponseBodyDataObjectList) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListMmsTimersResponseBodyDataObjectList) GetDbId() *int64 {
	return s.DbId
}

func (s *ListMmsTimersResponseBodyDataObjectList) GetId() *int64 {
	return s.Id
}

func (s *ListMmsTimersResponseBodyDataObjectList) GetName() *string {
	return s.Name
}

func (s *ListMmsTimersResponseBodyDataObjectList) GetScheduleType() *string {
	return s.ScheduleType
}

func (s *ListMmsTimersResponseBodyDataObjectList) GetSourceId() *int64 {
	return s.SourceId
}

func (s *ListMmsTimersResponseBodyDataObjectList) GetSrcDbName() *string {
	return s.SrcDbName
}

func (s *ListMmsTimersResponseBodyDataObjectList) GetStopped() *bool {
	return s.Stopped
}

func (s *ListMmsTimersResponseBodyDataObjectList) GetType() *string {
	return s.Type
}

func (s *ListMmsTimersResponseBodyDataObjectList) GetValue() *string {
	return s.Value
}

func (s *ListMmsTimersResponseBodyDataObjectList) SetCreateTime(v string) *ListMmsTimersResponseBodyDataObjectList {
	s.CreateTime = &v
	return s
}

func (s *ListMmsTimersResponseBodyDataObjectList) SetDbId(v int64) *ListMmsTimersResponseBodyDataObjectList {
	s.DbId = &v
	return s
}

func (s *ListMmsTimersResponseBodyDataObjectList) SetId(v int64) *ListMmsTimersResponseBodyDataObjectList {
	s.Id = &v
	return s
}

func (s *ListMmsTimersResponseBodyDataObjectList) SetName(v string) *ListMmsTimersResponseBodyDataObjectList {
	s.Name = &v
	return s
}

func (s *ListMmsTimersResponseBodyDataObjectList) SetScheduleType(v string) *ListMmsTimersResponseBodyDataObjectList {
	s.ScheduleType = &v
	return s
}

func (s *ListMmsTimersResponseBodyDataObjectList) SetSourceId(v int64) *ListMmsTimersResponseBodyDataObjectList {
	s.SourceId = &v
	return s
}

func (s *ListMmsTimersResponseBodyDataObjectList) SetSrcDbName(v string) *ListMmsTimersResponseBodyDataObjectList {
	s.SrcDbName = &v
	return s
}

func (s *ListMmsTimersResponseBodyDataObjectList) SetStopped(v bool) *ListMmsTimersResponseBodyDataObjectList {
	s.Stopped = &v
	return s
}

func (s *ListMmsTimersResponseBodyDataObjectList) SetType(v string) *ListMmsTimersResponseBodyDataObjectList {
	s.Type = &v
	return s
}

func (s *ListMmsTimersResponseBodyDataObjectList) SetValue(v string) *ListMmsTimersResponseBodyDataObjectList {
	s.Value = &v
	return s
}

func (s *ListMmsTimersResponseBodyDataObjectList) Validate() error {
	return dara.Validate(s)
}
