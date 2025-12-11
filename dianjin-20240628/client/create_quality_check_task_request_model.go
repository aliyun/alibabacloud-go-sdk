// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateQualityCheckTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConversationList(v *CreateQualityCheckTaskRequestConversationList) *CreateQualityCheckTaskRequest
	GetConversationList() *CreateQualityCheckTaskRequestConversationList
	SetGmtService(v string) *CreateQualityCheckTaskRequest
	GetGmtService() *string
	SetMetaData(v map[string]*string) *CreateQualityCheckTaskRequest
	GetMetaData() map[string]*string
	SetQualityGroup(v []*string) *CreateQualityCheckTaskRequest
	GetQualityGroup() []*string
	SetRequestId(v string) *CreateQualityCheckTaskRequest
	GetRequestId() *string
	SetSceneCode(v string) *CreateQualityCheckTaskRequest
	GetSceneCode() *string
	SetType(v string) *CreateQualityCheckTaskRequest
	GetType() *string
}

type CreateQualityCheckTaskRequest struct {
	// This parameter is required.
	ConversationList *CreateQualityCheckTaskRequestConversationList `json:"conversationList,omitempty" xml:"conversationList,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 2024-09-27 11:23:20
	GmtService   *string            `json:"gmtService,omitempty" xml:"gmtService,omitempty"`
	MetaData     map[string]*string `json:"metaData,omitempty" xml:"metaData,omitempty"`
	QualityGroup []*string          `json:"qualityGroup,omitempty" xml:"qualityGroup,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 0FC6636E-380A-5369-AE01-D1C15BB9B254
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	SceneCode *string `json:"sceneCode,omitempty" xml:"sceneCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s CreateQualityCheckTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateQualityCheckTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateQualityCheckTaskRequest) GetConversationList() *CreateQualityCheckTaskRequestConversationList {
	return s.ConversationList
}

func (s *CreateQualityCheckTaskRequest) GetGmtService() *string {
	return s.GmtService
}

func (s *CreateQualityCheckTaskRequest) GetMetaData() map[string]*string {
	return s.MetaData
}

func (s *CreateQualityCheckTaskRequest) GetQualityGroup() []*string {
	return s.QualityGroup
}

func (s *CreateQualityCheckTaskRequest) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateQualityCheckTaskRequest) GetSceneCode() *string {
	return s.SceneCode
}

func (s *CreateQualityCheckTaskRequest) GetType() *string {
	return s.Type
}

func (s *CreateQualityCheckTaskRequest) SetConversationList(v *CreateQualityCheckTaskRequestConversationList) *CreateQualityCheckTaskRequest {
	s.ConversationList = v
	return s
}

func (s *CreateQualityCheckTaskRequest) SetGmtService(v string) *CreateQualityCheckTaskRequest {
	s.GmtService = &v
	return s
}

func (s *CreateQualityCheckTaskRequest) SetMetaData(v map[string]*string) *CreateQualityCheckTaskRequest {
	s.MetaData = v
	return s
}

func (s *CreateQualityCheckTaskRequest) SetQualityGroup(v []*string) *CreateQualityCheckTaskRequest {
	s.QualityGroup = v
	return s
}

func (s *CreateQualityCheckTaskRequest) SetRequestId(v string) *CreateQualityCheckTaskRequest {
	s.RequestId = &v
	return s
}

func (s *CreateQualityCheckTaskRequest) SetSceneCode(v string) *CreateQualityCheckTaskRequest {
	s.SceneCode = &v
	return s
}

func (s *CreateQualityCheckTaskRequest) SetType(v string) *CreateQualityCheckTaskRequest {
	s.Type = &v
	return s
}

func (s *CreateQualityCheckTaskRequest) Validate() error {
	if s.ConversationList != nil {
		if err := s.ConversationList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateQualityCheckTaskRequestConversationList struct {
	// example:
	//
	// 1
	CallType *string `json:"callType,omitempty" xml:"callType,omitempty"`
	// example:
	//
	// 1
	CustomerId   *string `json:"customerId,omitempty" xml:"customerId,omitempty"`
	CustomerName *string `json:"customerName,omitempty" xml:"customerName,omitempty"`
	// example:
	//
	// xxx
	CustomerServiceId   *string `json:"customerServiceId,omitempty" xml:"customerServiceId,omitempty"`
	CustomerServiceName *string `json:"customerServiceName,omitempty" xml:"customerServiceName,omitempty"`
	// This parameter is required.
	DialogueList []*CreateQualityCheckTaskRequestConversationListDialogueList `json:"dialogueList,omitempty" xml:"dialogueList,omitempty" type:"Repeated"`
	// example:
	//
	// 2024-09-27 11:23:20
	GmtService *string `json:"gmtService,omitempty" xml:"gmtService,omitempty"`
}

func (s CreateQualityCheckTaskRequestConversationList) String() string {
	return dara.Prettify(s)
}

func (s CreateQualityCheckTaskRequestConversationList) GoString() string {
	return s.String()
}

func (s *CreateQualityCheckTaskRequestConversationList) GetCallType() *string {
	return s.CallType
}

func (s *CreateQualityCheckTaskRequestConversationList) GetCustomerId() *string {
	return s.CustomerId
}

func (s *CreateQualityCheckTaskRequestConversationList) GetCustomerName() *string {
	return s.CustomerName
}

func (s *CreateQualityCheckTaskRequestConversationList) GetCustomerServiceId() *string {
	return s.CustomerServiceId
}

func (s *CreateQualityCheckTaskRequestConversationList) GetCustomerServiceName() *string {
	return s.CustomerServiceName
}

func (s *CreateQualityCheckTaskRequestConversationList) GetDialogueList() []*CreateQualityCheckTaskRequestConversationListDialogueList {
	return s.DialogueList
}

func (s *CreateQualityCheckTaskRequestConversationList) GetGmtService() *string {
	return s.GmtService
}

func (s *CreateQualityCheckTaskRequestConversationList) SetCallType(v string) *CreateQualityCheckTaskRequestConversationList {
	s.CallType = &v
	return s
}

func (s *CreateQualityCheckTaskRequestConversationList) SetCustomerId(v string) *CreateQualityCheckTaskRequestConversationList {
	s.CustomerId = &v
	return s
}

func (s *CreateQualityCheckTaskRequestConversationList) SetCustomerName(v string) *CreateQualityCheckTaskRequestConversationList {
	s.CustomerName = &v
	return s
}

func (s *CreateQualityCheckTaskRequestConversationList) SetCustomerServiceId(v string) *CreateQualityCheckTaskRequestConversationList {
	s.CustomerServiceId = &v
	return s
}

func (s *CreateQualityCheckTaskRequestConversationList) SetCustomerServiceName(v string) *CreateQualityCheckTaskRequestConversationList {
	s.CustomerServiceName = &v
	return s
}

func (s *CreateQualityCheckTaskRequestConversationList) SetDialogueList(v []*CreateQualityCheckTaskRequestConversationListDialogueList) *CreateQualityCheckTaskRequestConversationList {
	s.DialogueList = v
	return s
}

func (s *CreateQualityCheckTaskRequestConversationList) SetGmtService(v string) *CreateQualityCheckTaskRequestConversationList {
	s.GmtService = &v
	return s
}

func (s *CreateQualityCheckTaskRequestConversationList) Validate() error {
	if s.DialogueList != nil {
		for _, item := range s.DialogueList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateQualityCheckTaskRequestConversationListDialogueList struct {
	// example:
	//
	// 0
	Begin *int32 `json:"begin,omitempty" xml:"begin,omitempty"`
	// example:
	//
	// 2024-05-23 14:57:50
	BeginTime *string `json:"beginTime,omitempty" xml:"beginTime,omitempty"`
	// This parameter is required.
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// example:
	//
	// 2348234
	CustomerId *string `json:"customerId,omitempty" xml:"customerId,omitempty"`
	// example:
	//
	// 23874627346
	CustomerServiceId *string `json:"customerServiceId,omitempty" xml:"customerServiceId,omitempty"`
	// example:
	//
	// 0
	CustomerServiceType *string `json:"customerServiceType,omitempty" xml:"customerServiceType,omitempty"`
	// example:
	//
	// 0
	End *int32 `json:"end,omitempty" xml:"end,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	Role *string `json:"role,omitempty" xml:"role,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// TEXT
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s CreateQualityCheckTaskRequestConversationListDialogueList) String() string {
	return dara.Prettify(s)
}

func (s CreateQualityCheckTaskRequestConversationListDialogueList) GoString() string {
	return s.String()
}

func (s *CreateQualityCheckTaskRequestConversationListDialogueList) GetBegin() *int32 {
	return s.Begin
}

func (s *CreateQualityCheckTaskRequestConversationListDialogueList) GetBeginTime() *string {
	return s.BeginTime
}

func (s *CreateQualityCheckTaskRequestConversationListDialogueList) GetContent() *string {
	return s.Content
}

func (s *CreateQualityCheckTaskRequestConversationListDialogueList) GetCustomerId() *string {
	return s.CustomerId
}

func (s *CreateQualityCheckTaskRequestConversationListDialogueList) GetCustomerServiceId() *string {
	return s.CustomerServiceId
}

func (s *CreateQualityCheckTaskRequestConversationListDialogueList) GetCustomerServiceType() *string {
	return s.CustomerServiceType
}

func (s *CreateQualityCheckTaskRequestConversationListDialogueList) GetEnd() *int32 {
	return s.End
}

func (s *CreateQualityCheckTaskRequestConversationListDialogueList) GetRole() *string {
	return s.Role
}

func (s *CreateQualityCheckTaskRequestConversationListDialogueList) GetType() *string {
	return s.Type
}

func (s *CreateQualityCheckTaskRequestConversationListDialogueList) SetBegin(v int32) *CreateQualityCheckTaskRequestConversationListDialogueList {
	s.Begin = &v
	return s
}

func (s *CreateQualityCheckTaskRequestConversationListDialogueList) SetBeginTime(v string) *CreateQualityCheckTaskRequestConversationListDialogueList {
	s.BeginTime = &v
	return s
}

func (s *CreateQualityCheckTaskRequestConversationListDialogueList) SetContent(v string) *CreateQualityCheckTaskRequestConversationListDialogueList {
	s.Content = &v
	return s
}

func (s *CreateQualityCheckTaskRequestConversationListDialogueList) SetCustomerId(v string) *CreateQualityCheckTaskRequestConversationListDialogueList {
	s.CustomerId = &v
	return s
}

func (s *CreateQualityCheckTaskRequestConversationListDialogueList) SetCustomerServiceId(v string) *CreateQualityCheckTaskRequestConversationListDialogueList {
	s.CustomerServiceId = &v
	return s
}

func (s *CreateQualityCheckTaskRequestConversationListDialogueList) SetCustomerServiceType(v string) *CreateQualityCheckTaskRequestConversationListDialogueList {
	s.CustomerServiceType = &v
	return s
}

func (s *CreateQualityCheckTaskRequestConversationListDialogueList) SetEnd(v int32) *CreateQualityCheckTaskRequestConversationListDialogueList {
	s.End = &v
	return s
}

func (s *CreateQualityCheckTaskRequestConversationListDialogueList) SetRole(v string) *CreateQualityCheckTaskRequestConversationListDialogueList {
	s.Role = &v
	return s
}

func (s *CreateQualityCheckTaskRequestConversationListDialogueList) SetType(v string) *CreateQualityCheckTaskRequestConversationListDialogueList {
	s.Type = &v
	return s
}

func (s *CreateQualityCheckTaskRequestConversationListDialogueList) Validate() error {
	return dara.Validate(s)
}
