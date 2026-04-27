// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudQueryAgentSkillResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *CloudQueryAgentSkillResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *CloudQueryAgentSkillResponseBody
	GetCode() *string
	SetData(v *CloudQueryAgentSkillResponseBodyData) *CloudQueryAgentSkillResponseBody
	GetData() *CloudQueryAgentSkillResponseBodyData
	SetMessage(v string) *CloudQueryAgentSkillResponseBody
	GetMessage() *string
	SetRequestId(v string) *CloudQueryAgentSkillResponseBody
	GetRequestId() *string
}

type CloudQueryAgentSkillResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *CloudQueryAgentSkillResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 5B0F201F-DCDA-45C2-AA92-1AE177F94991
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CloudQueryAgentSkillResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CloudQueryAgentSkillResponseBody) GoString() string {
	return s.String()
}

func (s *CloudQueryAgentSkillResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *CloudQueryAgentSkillResponseBody) GetCode() *string {
	return s.Code
}

func (s *CloudQueryAgentSkillResponseBody) GetData() *CloudQueryAgentSkillResponseBodyData {
	return s.Data
}

func (s *CloudQueryAgentSkillResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CloudQueryAgentSkillResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CloudQueryAgentSkillResponseBody) SetAccessDeniedDetail(v string) *CloudQueryAgentSkillResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CloudQueryAgentSkillResponseBody) SetCode(v string) *CloudQueryAgentSkillResponseBody {
	s.Code = &v
	return s
}

func (s *CloudQueryAgentSkillResponseBody) SetData(v *CloudQueryAgentSkillResponseBodyData) *CloudQueryAgentSkillResponseBody {
	s.Data = v
	return s
}

func (s *CloudQueryAgentSkillResponseBody) SetMessage(v string) *CloudQueryAgentSkillResponseBody {
	s.Message = &v
	return s
}

func (s *CloudQueryAgentSkillResponseBody) SetRequestId(v string) *CloudQueryAgentSkillResponseBody {
	s.RequestId = &v
	return s
}

func (s *CloudQueryAgentSkillResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CloudQueryAgentSkillResponseBodyData struct {
	// 座席技能列表
	List []*CloudQueryAgentSkillResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
}

func (s CloudQueryAgentSkillResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CloudQueryAgentSkillResponseBodyData) GoString() string {
	return s.String()
}

func (s *CloudQueryAgentSkillResponseBodyData) GetList() []*CloudQueryAgentSkillResponseBodyDataList {
	return s.List
}

func (s *CloudQueryAgentSkillResponseBodyData) SetList(v []*CloudQueryAgentSkillResponseBodyDataList) *CloudQueryAgentSkillResponseBodyData {
	s.List = v
	return s
}

func (s *CloudQueryAgentSkillResponseBodyData) Validate() error {
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

type CloudQueryAgentSkillResponseBodyDataList struct {
	// 座席id
	//
	// example:
	//
	// 64
	AgentId *int64 `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
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
	// 99
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// skill的id
	//
	// example:
	//
	// 79
	SkillId *int64 `json:"SkillId,omitempty" xml:"SkillId,omitempty"`
	// 技能值
	//
	// example:
	//
	// 7
	SkillLevel *int64 `json:"SkillLevel,omitempty" xml:"SkillLevel,omitempty"`
	// 技能名称
	//
	// example:
	//
	// skillname
	SkillName *string `json:"SkillName,omitempty" xml:"SkillName,omitempty"`
}

func (s CloudQueryAgentSkillResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s CloudQueryAgentSkillResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *CloudQueryAgentSkillResponseBodyDataList) GetAgentId() *int64 {
	return s.AgentId
}

func (s *CloudQueryAgentSkillResponseBodyDataList) GetCreateTime() *string {
	return s.CreateTime
}

func (s *CloudQueryAgentSkillResponseBodyDataList) GetEnterpriseId() *int64 {
	return s.EnterpriseId
}

func (s *CloudQueryAgentSkillResponseBodyDataList) GetId() *int64 {
	return s.Id
}

func (s *CloudQueryAgentSkillResponseBodyDataList) GetSkillId() *int64 {
	return s.SkillId
}

func (s *CloudQueryAgentSkillResponseBodyDataList) GetSkillLevel() *int64 {
	return s.SkillLevel
}

func (s *CloudQueryAgentSkillResponseBodyDataList) GetSkillName() *string {
	return s.SkillName
}

func (s *CloudQueryAgentSkillResponseBodyDataList) SetAgentId(v int64) *CloudQueryAgentSkillResponseBodyDataList {
	s.AgentId = &v
	return s
}

func (s *CloudQueryAgentSkillResponseBodyDataList) SetCreateTime(v string) *CloudQueryAgentSkillResponseBodyDataList {
	s.CreateTime = &v
	return s
}

func (s *CloudQueryAgentSkillResponseBodyDataList) SetEnterpriseId(v int64) *CloudQueryAgentSkillResponseBodyDataList {
	s.EnterpriseId = &v
	return s
}

func (s *CloudQueryAgentSkillResponseBodyDataList) SetId(v int64) *CloudQueryAgentSkillResponseBodyDataList {
	s.Id = &v
	return s
}

func (s *CloudQueryAgentSkillResponseBodyDataList) SetSkillId(v int64) *CloudQueryAgentSkillResponseBodyDataList {
	s.SkillId = &v
	return s
}

func (s *CloudQueryAgentSkillResponseBodyDataList) SetSkillLevel(v int64) *CloudQueryAgentSkillResponseBodyDataList {
	s.SkillLevel = &v
	return s
}

func (s *CloudQueryAgentSkillResponseBodyDataList) SetSkillName(v string) *CloudQueryAgentSkillResponseBodyDataList {
	s.SkillName = &v
	return s
}

func (s *CloudQueryAgentSkillResponseBodyDataList) Validate() error {
	return dara.Validate(s)
}
