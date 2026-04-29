// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudListQueueSkillResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *CloudListQueueSkillResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *CloudListQueueSkillResponseBody
	GetCode() *string
	SetData(v *CloudListQueueSkillResponseBodyData) *CloudListQueueSkillResponseBody
	GetData() *CloudListQueueSkillResponseBodyData
	SetMessage(v string) *CloudListQueueSkillResponseBody
	GetMessage() *string
	SetRequestId(v string) *CloudListQueueSkillResponseBody
	GetRequestId() *string
}

type CloudListQueueSkillResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *CloudListQueueSkillResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 7BF47617-7851-48F7-A3A1-2021342A78E2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CloudListQueueSkillResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CloudListQueueSkillResponseBody) GoString() string {
	return s.String()
}

func (s *CloudListQueueSkillResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *CloudListQueueSkillResponseBody) GetCode() *string {
	return s.Code
}

func (s *CloudListQueueSkillResponseBody) GetData() *CloudListQueueSkillResponseBodyData {
	return s.Data
}

func (s *CloudListQueueSkillResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CloudListQueueSkillResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CloudListQueueSkillResponseBody) SetAccessDeniedDetail(v string) *CloudListQueueSkillResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CloudListQueueSkillResponseBody) SetCode(v string) *CloudListQueueSkillResponseBody {
	s.Code = &v
	return s
}

func (s *CloudListQueueSkillResponseBody) SetData(v *CloudListQueueSkillResponseBodyData) *CloudListQueueSkillResponseBody {
	s.Data = v
	return s
}

func (s *CloudListQueueSkillResponseBody) SetMessage(v string) *CloudListQueueSkillResponseBody {
	s.Message = &v
	return s
}

func (s *CloudListQueueSkillResponseBody) SetRequestId(v string) *CloudListQueueSkillResponseBody {
	s.RequestId = &v
	return s
}

func (s *CloudListQueueSkillResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CloudListQueueSkillResponseBodyData struct {
	// 返回数据
	List []*CloudListQueueSkillResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
}

func (s CloudListQueueSkillResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CloudListQueueSkillResponseBodyData) GoString() string {
	return s.String()
}

func (s *CloudListQueueSkillResponseBodyData) GetList() []*CloudListQueueSkillResponseBodyDataList {
	return s.List
}

func (s *CloudListQueueSkillResponseBodyData) SetList(v []*CloudListQueueSkillResponseBodyDataList) *CloudListQueueSkillResponseBodyData {
	s.List = v
	return s
}

func (s *CloudListQueueSkillResponseBodyData) Validate() error {
	if s.List != nil {
		for _, item := range s.List {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CloudListQueueSkillResponseBodyDataList struct {
	// 创建时间，格式: yyyy-MM-dd HH:mm:ss
	//
	// example:
	//
	// 2026-04-20 10:00:00
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 企业编号
	//
	// example:
	//
	// 7000002
	EnterpriseId *int64 `json:"EnterpriseId,omitempty" xml:"EnterpriseId,omitempty"`
	// queueSkill关系表中id
	//
	// example:
	//
	// 93
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// 队列号
	//
	// example:
	//
	// 26
	Qno *int64 `json:"Qno,omitempty" xml:"Qno,omitempty"`
	// 队列id
	//
	// example:
	//
	// 33
	QueueId *int64 `json:"QueueId,omitempty" xml:"QueueId,omitempty"`
	// skill的id
	//
	// example:
	//
	// 44
	SkillId *int64 `json:"SkillId,omitempty" xml:"SkillId,omitempty"`
	// 技能值
	//
	// example:
	//
	// 5
	SkillLevel *int64 `json:"SkillLevel,omitempty" xml:"SkillLevel,omitempty"`
}

func (s CloudListQueueSkillResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s CloudListQueueSkillResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *CloudListQueueSkillResponseBodyDataList) GetCreateTime() *string {
	return s.CreateTime
}

func (s *CloudListQueueSkillResponseBodyDataList) GetEnterpriseId() *int64 {
	return s.EnterpriseId
}

func (s *CloudListQueueSkillResponseBodyDataList) GetId() *int64 {
	return s.Id
}

func (s *CloudListQueueSkillResponseBodyDataList) GetQno() *int64 {
	return s.Qno
}

func (s *CloudListQueueSkillResponseBodyDataList) GetQueueId() *int64 {
	return s.QueueId
}

func (s *CloudListQueueSkillResponseBodyDataList) GetSkillId() *int64 {
	return s.SkillId
}

func (s *CloudListQueueSkillResponseBodyDataList) GetSkillLevel() *int64 {
	return s.SkillLevel
}

func (s *CloudListQueueSkillResponseBodyDataList) SetCreateTime(v string) *CloudListQueueSkillResponseBodyDataList {
	s.CreateTime = &v
	return s
}

func (s *CloudListQueueSkillResponseBodyDataList) SetEnterpriseId(v int64) *CloudListQueueSkillResponseBodyDataList {
	s.EnterpriseId = &v
	return s
}

func (s *CloudListQueueSkillResponseBodyDataList) SetId(v int64) *CloudListQueueSkillResponseBodyDataList {
	s.Id = &v
	return s
}

func (s *CloudListQueueSkillResponseBodyDataList) SetQno(v int64) *CloudListQueueSkillResponseBodyDataList {
	s.Qno = &v
	return s
}

func (s *CloudListQueueSkillResponseBodyDataList) SetQueueId(v int64) *CloudListQueueSkillResponseBodyDataList {
	s.QueueId = &v
	return s
}

func (s *CloudListQueueSkillResponseBodyDataList) SetSkillId(v int64) *CloudListQueueSkillResponseBodyDataList {
	s.SkillId = &v
	return s
}

func (s *CloudListQueueSkillResponseBodyDataList) SetSkillLevel(v int64) *CloudListQueueSkillResponseBodyDataList {
	s.SkillLevel = &v
	return s
}

func (s *CloudListQueueSkillResponseBodyDataList) Validate() error {
	return dara.Validate(s)
}
