// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudCreateQueueSkillResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *CloudCreateQueueSkillResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *CloudCreateQueueSkillResponseBody
	GetCode() *string
	SetData(v *CloudCreateQueueSkillResponseBodyData) *CloudCreateQueueSkillResponseBody
	GetData() *CloudCreateQueueSkillResponseBodyData
	SetMessage(v string) *CloudCreateQueueSkillResponseBody
	GetMessage() *string
	SetRequestId(v string) *CloudCreateQueueSkillResponseBody
	GetRequestId() *string
}

type CloudCreateQueueSkillResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *CloudCreateQueueSkillResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 9927CD06-C33A-50CC-A9BB-A3427967A66F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CloudCreateQueueSkillResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CloudCreateQueueSkillResponseBody) GoString() string {
	return s.String()
}

func (s *CloudCreateQueueSkillResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *CloudCreateQueueSkillResponseBody) GetCode() *string {
	return s.Code
}

func (s *CloudCreateQueueSkillResponseBody) GetData() *CloudCreateQueueSkillResponseBodyData {
	return s.Data
}

func (s *CloudCreateQueueSkillResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CloudCreateQueueSkillResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CloudCreateQueueSkillResponseBody) SetAccessDeniedDetail(v string) *CloudCreateQueueSkillResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CloudCreateQueueSkillResponseBody) SetCode(v string) *CloudCreateQueueSkillResponseBody {
	s.Code = &v
	return s
}

func (s *CloudCreateQueueSkillResponseBody) SetData(v *CloudCreateQueueSkillResponseBodyData) *CloudCreateQueueSkillResponseBody {
	s.Data = v
	return s
}

func (s *CloudCreateQueueSkillResponseBody) SetMessage(v string) *CloudCreateQueueSkillResponseBody {
	s.Message = &v
	return s
}

func (s *CloudCreateQueueSkillResponseBody) SetRequestId(v string) *CloudCreateQueueSkillResponseBody {
	s.RequestId = &v
	return s
}

func (s *CloudCreateQueueSkillResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CloudCreateQueueSkillResponseBodyData struct {
	// 创建时间，格式: yyyy-MM-dd HH:mm:ss
	//
	// example:
	//
	// 2026-03-30 06:12:34
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 企业编号
	//
	// example:
	//
	// 7122600
	EnterpriseId *int64 `json:"EnterpriseId,omitempty" xml:"EnterpriseId,omitempty"`
	// id
	//
	// example:
	//
	// 54475
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// 队列号
	//
	// example:
	//
	// 1000
	Qno *int64 `json:"Qno,omitempty" xml:"Qno,omitempty"`
	// 队列id
	//
	// example:
	//
	// 51937
	QueueId *int64 `json:"QueueId,omitempty" xml:"QueueId,omitempty"`
	// skill的id
	//
	// example:
	//
	// 48735
	SkillId *int64 `json:"SkillId,omitempty" xml:"SkillId,omitempty"`
	// 技能值
	//
	// example:
	//
	// 10
	SkillLevel *int64 `json:"SkillLevel,omitempty" xml:"SkillLevel,omitempty"`
}

func (s CloudCreateQueueSkillResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CloudCreateQueueSkillResponseBodyData) GoString() string {
	return s.String()
}

func (s *CloudCreateQueueSkillResponseBodyData) GetCreateTime() *string {
	return s.CreateTime
}

func (s *CloudCreateQueueSkillResponseBodyData) GetEnterpriseId() *int64 {
	return s.EnterpriseId
}

func (s *CloudCreateQueueSkillResponseBodyData) GetId() *int64 {
	return s.Id
}

func (s *CloudCreateQueueSkillResponseBodyData) GetQno() *int64 {
	return s.Qno
}

func (s *CloudCreateQueueSkillResponseBodyData) GetQueueId() *int64 {
	return s.QueueId
}

func (s *CloudCreateQueueSkillResponseBodyData) GetSkillId() *int64 {
	return s.SkillId
}

func (s *CloudCreateQueueSkillResponseBodyData) GetSkillLevel() *int64 {
	return s.SkillLevel
}

func (s *CloudCreateQueueSkillResponseBodyData) SetCreateTime(v string) *CloudCreateQueueSkillResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *CloudCreateQueueSkillResponseBodyData) SetEnterpriseId(v int64) *CloudCreateQueueSkillResponseBodyData {
	s.EnterpriseId = &v
	return s
}

func (s *CloudCreateQueueSkillResponseBodyData) SetId(v int64) *CloudCreateQueueSkillResponseBodyData {
	s.Id = &v
	return s
}

func (s *CloudCreateQueueSkillResponseBodyData) SetQno(v int64) *CloudCreateQueueSkillResponseBodyData {
	s.Qno = &v
	return s
}

func (s *CloudCreateQueueSkillResponseBodyData) SetQueueId(v int64) *CloudCreateQueueSkillResponseBodyData {
	s.QueueId = &v
	return s
}

func (s *CloudCreateQueueSkillResponseBodyData) SetSkillId(v int64) *CloudCreateQueueSkillResponseBodyData {
	s.SkillId = &v
	return s
}

func (s *CloudCreateQueueSkillResponseBodyData) SetSkillLevel(v int64) *CloudCreateQueueSkillResponseBodyData {
	s.SkillLevel = &v
	return s
}

func (s *CloudCreateQueueSkillResponseBodyData) Validate() error {
	return dara.Validate(s)
}
