// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuickAddTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentGroupId(v int64) *QuickAddTaskRequest
	GetAgentGroupId() *int64
	SetCallTimeList(v []*QuickAddTaskRequestCallTimeList) *QuickAddTaskRequest
	GetCallTimeList() []*QuickAddTaskRequestCallTimeList
	SetName(v string) *QuickAddTaskRequest
	GetName() *string
	SetOwnerId(v int64) *QuickAddTaskRequest
	GetOwnerId() *int64
	SetReferenceTaskId(v int64) *QuickAddTaskRequest
	GetReferenceTaskId() *int64
	SetResourceOwnerAccount(v string) *QuickAddTaskRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *QuickAddTaskRequest
	GetResourceOwnerId() *int64
	SetSmsTemplateId(v int64) *QuickAddTaskRequest
	GetSmsTemplateId() *int64
	SetStartTime(v string) *QuickAddTaskRequest
	GetStartTime() *string
	SetTemplateId(v int64) *QuickAddTaskRequest
	GetTemplateId() *int64
	SetTemplateType(v int64) *QuickAddTaskRequest
	GetTemplateType() *int64
}

type QuickAddTaskRequest struct {
	// 坐席组ID
	//
	// example:
	//
	// 1
	AgentGroupId *int64 `json:"AgentGroupId,omitempty" xml:"AgentGroupId,omitempty"`
	// 外呼时间
	CallTimeList []*QuickAddTaskRequestCallTimeList `json:"CallTimeList,omitempty" xml:"CallTimeList,omitempty" type:"Repeated"`
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

func (s QuickAddTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s QuickAddTaskRequest) GoString() string {
	return s.String()
}

func (s *QuickAddTaskRequest) GetAgentGroupId() *int64 {
	return s.AgentGroupId
}

func (s *QuickAddTaskRequest) GetCallTimeList() []*QuickAddTaskRequestCallTimeList {
	return s.CallTimeList
}

func (s *QuickAddTaskRequest) GetName() *string {
	return s.Name
}

func (s *QuickAddTaskRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *QuickAddTaskRequest) GetReferenceTaskId() *int64 {
	return s.ReferenceTaskId
}

func (s *QuickAddTaskRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *QuickAddTaskRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *QuickAddTaskRequest) GetSmsTemplateId() *int64 {
	return s.SmsTemplateId
}

func (s *QuickAddTaskRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *QuickAddTaskRequest) GetTemplateId() *int64 {
	return s.TemplateId
}

func (s *QuickAddTaskRequest) GetTemplateType() *int64 {
	return s.TemplateType
}

func (s *QuickAddTaskRequest) SetAgentGroupId(v int64) *QuickAddTaskRequest {
	s.AgentGroupId = &v
	return s
}

func (s *QuickAddTaskRequest) SetCallTimeList(v []*QuickAddTaskRequestCallTimeList) *QuickAddTaskRequest {
	s.CallTimeList = v
	return s
}

func (s *QuickAddTaskRequest) SetName(v string) *QuickAddTaskRequest {
	s.Name = &v
	return s
}

func (s *QuickAddTaskRequest) SetOwnerId(v int64) *QuickAddTaskRequest {
	s.OwnerId = &v
	return s
}

func (s *QuickAddTaskRequest) SetReferenceTaskId(v int64) *QuickAddTaskRequest {
	s.ReferenceTaskId = &v
	return s
}

func (s *QuickAddTaskRequest) SetResourceOwnerAccount(v string) *QuickAddTaskRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QuickAddTaskRequest) SetResourceOwnerId(v int64) *QuickAddTaskRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QuickAddTaskRequest) SetSmsTemplateId(v int64) *QuickAddTaskRequest {
	s.SmsTemplateId = &v
	return s
}

func (s *QuickAddTaskRequest) SetStartTime(v string) *QuickAddTaskRequest {
	s.StartTime = &v
	return s
}

func (s *QuickAddTaskRequest) SetTemplateId(v int64) *QuickAddTaskRequest {
	s.TemplateId = &v
	return s
}

func (s *QuickAddTaskRequest) SetTemplateType(v int64) *QuickAddTaskRequest {
	s.TemplateType = &v
	return s
}

func (s *QuickAddTaskRequest) Validate() error {
	if s.CallTimeList != nil {
		for _, item := range s.CallTimeList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QuickAddTaskRequestCallTimeList struct {
	// 外呼时间
	CallTime []*string `json:"CallTime,omitempty" xml:"CallTime,omitempty" type:"Repeated"`
}

func (s QuickAddTaskRequestCallTimeList) String() string {
	return dara.Prettify(s)
}

func (s QuickAddTaskRequestCallTimeList) GoString() string {
	return s.String()
}

func (s *QuickAddTaskRequestCallTimeList) GetCallTime() []*string {
	return s.CallTime
}

func (s *QuickAddTaskRequestCallTimeList) SetCallTime(v []*string) *QuickAddTaskRequestCallTimeList {
	s.CallTime = v
	return s
}

func (s *QuickAddTaskRequestCallTimeList) Validate() error {
	return dara.Validate(s)
}
