// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuickAddTaskShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentGroupId(v int64) *QuickAddTaskShrinkRequest
	GetAgentGroupId() *int64
	SetCallTimeListShrink(v string) *QuickAddTaskShrinkRequest
	GetCallTimeListShrink() *string
	SetName(v string) *QuickAddTaskShrinkRequest
	GetName() *string
	SetOwnerId(v int64) *QuickAddTaskShrinkRequest
	GetOwnerId() *int64
	SetReferenceTaskId(v int64) *QuickAddTaskShrinkRequest
	GetReferenceTaskId() *int64
	SetResourceOwnerAccount(v string) *QuickAddTaskShrinkRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *QuickAddTaskShrinkRequest
	GetResourceOwnerId() *int64
	SetSmsTemplateId(v int64) *QuickAddTaskShrinkRequest
	GetSmsTemplateId() *int64
	SetStartTime(v string) *QuickAddTaskShrinkRequest
	GetStartTime() *string
	SetTemplateId(v int64) *QuickAddTaskShrinkRequest
	GetTemplateId() *int64
	SetTemplateType(v int64) *QuickAddTaskShrinkRequest
	GetTemplateType() *int64
}

type QuickAddTaskShrinkRequest struct {
	// 坐席组ID
	//
	// example:
	//
	// 1
	AgentGroupId *int64 `json:"AgentGroupId,omitempty" xml:"AgentGroupId,omitempty"`
	// 外呼时间
	CallTimeListShrink *string `json:"CallTimeList,omitempty" xml:"CallTimeList,omitempty"`
	// 任务名称
	//
	// This parameter is required.
	//
	// example:
	//
	// a
	Name    *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// 被复制任务ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	ReferenceTaskId      *int64  `json:"ReferenceTaskId,omitempty" xml:"ReferenceTaskId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// 短信模板ID
	//
	// example:
	//
	// 1
	SmsTemplateId *int64 `json:"SmsTemplateId,omitempty" xml:"SmsTemplateId,omitempty"`
	// 任务启动日期
	//
	// example:
	//
	// 2019-04-30
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// 话术模板ID,如果不传则使用被复制任务的话术模板
	//
	// example:
	//
	// 1
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// 话术模板类型
	//
	// example:
	//
	// 1：单模版；2：智能话术策略模板。不填默认1。
	TemplateType *int64 `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
}

func (s QuickAddTaskShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s QuickAddTaskShrinkRequest) GoString() string {
	return s.String()
}

func (s *QuickAddTaskShrinkRequest) GetAgentGroupId() *int64 {
	return s.AgentGroupId
}

func (s *QuickAddTaskShrinkRequest) GetCallTimeListShrink() *string {
	return s.CallTimeListShrink
}

func (s *QuickAddTaskShrinkRequest) GetName() *string {
	return s.Name
}

func (s *QuickAddTaskShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *QuickAddTaskShrinkRequest) GetReferenceTaskId() *int64 {
	return s.ReferenceTaskId
}

func (s *QuickAddTaskShrinkRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *QuickAddTaskShrinkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *QuickAddTaskShrinkRequest) GetSmsTemplateId() *int64 {
	return s.SmsTemplateId
}

func (s *QuickAddTaskShrinkRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *QuickAddTaskShrinkRequest) GetTemplateId() *int64 {
	return s.TemplateId
}

func (s *QuickAddTaskShrinkRequest) GetTemplateType() *int64 {
	return s.TemplateType
}

func (s *QuickAddTaskShrinkRequest) SetAgentGroupId(v int64) *QuickAddTaskShrinkRequest {
	s.AgentGroupId = &v
	return s
}

func (s *QuickAddTaskShrinkRequest) SetCallTimeListShrink(v string) *QuickAddTaskShrinkRequest {
	s.CallTimeListShrink = &v
	return s
}

func (s *QuickAddTaskShrinkRequest) SetName(v string) *QuickAddTaskShrinkRequest {
	s.Name = &v
	return s
}

func (s *QuickAddTaskShrinkRequest) SetOwnerId(v int64) *QuickAddTaskShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *QuickAddTaskShrinkRequest) SetReferenceTaskId(v int64) *QuickAddTaskShrinkRequest {
	s.ReferenceTaskId = &v
	return s
}

func (s *QuickAddTaskShrinkRequest) SetResourceOwnerAccount(v string) *QuickAddTaskShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QuickAddTaskShrinkRequest) SetResourceOwnerId(v int64) *QuickAddTaskShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QuickAddTaskShrinkRequest) SetSmsTemplateId(v int64) *QuickAddTaskShrinkRequest {
	s.SmsTemplateId = &v
	return s
}

func (s *QuickAddTaskShrinkRequest) SetStartTime(v string) *QuickAddTaskShrinkRequest {
	s.StartTime = &v
	return s
}

func (s *QuickAddTaskShrinkRequest) SetTemplateId(v int64) *QuickAddTaskShrinkRequest {
	s.TemplateId = &v
	return s
}

func (s *QuickAddTaskShrinkRequest) SetTemplateType(v int64) *QuickAddTaskShrinkRequest {
	s.TemplateType = &v
	return s
}

func (s *QuickAddTaskShrinkRequest) Validate() error {
	return dara.Validate(s)
}
