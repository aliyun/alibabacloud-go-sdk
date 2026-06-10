// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAIStaffChatEventsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ListAIStaffChatEventsResponseBody
	GetAccessDeniedDetail() *string
	SetAllowRetry(v bool) *ListAIStaffChatEventsResponseBody
	GetAllowRetry() *bool
	SetAppName(v string) *ListAIStaffChatEventsResponseBody
	GetAppName() *string
	SetDynamicCode(v string) *ListAIStaffChatEventsResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *ListAIStaffChatEventsResponseBody
	GetDynamicMessage() *string
	SetErrorArgs(v []interface{}) *ListAIStaffChatEventsResponseBody
	GetErrorArgs() []interface{}
	SetModule(v *ListAIStaffChatEventsResponseBodyModule) *ListAIStaffChatEventsResponseBody
	GetModule() *ListAIStaffChatEventsResponseBodyModule
	SetRequestId(v string) *ListAIStaffChatEventsResponseBody
	GetRequestId() *string
	SetRootErrorCode(v string) *ListAIStaffChatEventsResponseBody
	GetRootErrorCode() *string
	SetRootErrorMsg(v string) *ListAIStaffChatEventsResponseBody
	GetRootErrorMsg() *string
	SetSynchro(v bool) *ListAIStaffChatEventsResponseBody
	GetSynchro() *bool
}

type ListAIStaffChatEventsResponseBody struct {
	// Detailed reason for access denial.
	//
	// example:
	//
	// {}
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// is retry allowed
	//
	// example:
	//
	// False
	AllowRetry *bool `json:"AllowRetry,omitempty" xml:"AllowRetry,omitempty"`
	// App Name.
	//
	// example:
	//
	// spring-cloud-b
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// dynamic error code.
	//
	// example:
	//
	// ERROR-oo1
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// dynamic error message, used to replace `%s` in the **ErrMessage*	- error message.
	//
	// > If **ErrMessage*	- returns **The Value of Input Parameter %s is not valid*	- and **DynamicMessage*	- returns **DtsJobId**, it indicates that the request parameter **DtsJobId*	- is invalid.
	//
	// example:
	//
	// SYSTEM_ERROR
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// faulty parameters
	ErrorArgs []interface{} `json:"ErrorArgs,omitempty" xml:"ErrorArgs,omitempty" type:"Repeated"`
	// returned object.
	Module *ListAIStaffChatEventsResponseBodyModule `json:"Module,omitempty" xml:"Module,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 6C6B99AC-39EC-5350-874C-204128C905E6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// error code
	//
	// example:
	//
	// SYSTEM.ERROR
	RootErrorCode *string `json:"RootErrorCode,omitempty" xml:"RootErrorCode,omitempty"`
	// abnormal message
	//
	// example:
	//
	// 系统异常
	RootErrorMsg *string `json:"RootErrorMsg,omitempty" xml:"RootErrorMsg,omitempty"`
	// is processed synchronously
	//
	// example:
	//
	// True
	Synchro *bool `json:"Synchro,omitempty" xml:"Synchro,omitempty"`
}

func (s ListAIStaffChatEventsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAIStaffChatEventsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAIStaffChatEventsResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ListAIStaffChatEventsResponseBody) GetAllowRetry() *bool {
	return s.AllowRetry
}

func (s *ListAIStaffChatEventsResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *ListAIStaffChatEventsResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *ListAIStaffChatEventsResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *ListAIStaffChatEventsResponseBody) GetErrorArgs() []interface{} {
	return s.ErrorArgs
}

func (s *ListAIStaffChatEventsResponseBody) GetModule() *ListAIStaffChatEventsResponseBodyModule {
	return s.Module
}

func (s *ListAIStaffChatEventsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAIStaffChatEventsResponseBody) GetRootErrorCode() *string {
	return s.RootErrorCode
}

func (s *ListAIStaffChatEventsResponseBody) GetRootErrorMsg() *string {
	return s.RootErrorMsg
}

func (s *ListAIStaffChatEventsResponseBody) GetSynchro() *bool {
	return s.Synchro
}

func (s *ListAIStaffChatEventsResponseBody) SetAccessDeniedDetail(v string) *ListAIStaffChatEventsResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ListAIStaffChatEventsResponseBody) SetAllowRetry(v bool) *ListAIStaffChatEventsResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *ListAIStaffChatEventsResponseBody) SetAppName(v string) *ListAIStaffChatEventsResponseBody {
	s.AppName = &v
	return s
}

func (s *ListAIStaffChatEventsResponseBody) SetDynamicCode(v string) *ListAIStaffChatEventsResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *ListAIStaffChatEventsResponseBody) SetDynamicMessage(v string) *ListAIStaffChatEventsResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *ListAIStaffChatEventsResponseBody) SetErrorArgs(v []interface{}) *ListAIStaffChatEventsResponseBody {
	s.ErrorArgs = v
	return s
}

func (s *ListAIStaffChatEventsResponseBody) SetModule(v *ListAIStaffChatEventsResponseBodyModule) *ListAIStaffChatEventsResponseBody {
	s.Module = v
	return s
}

func (s *ListAIStaffChatEventsResponseBody) SetRequestId(v string) *ListAIStaffChatEventsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAIStaffChatEventsResponseBody) SetRootErrorCode(v string) *ListAIStaffChatEventsResponseBody {
	s.RootErrorCode = &v
	return s
}

func (s *ListAIStaffChatEventsResponseBody) SetRootErrorMsg(v string) *ListAIStaffChatEventsResponseBody {
	s.RootErrorMsg = &v
	return s
}

func (s *ListAIStaffChatEventsResponseBody) SetSynchro(v bool) *ListAIStaffChatEventsResponseBody {
	s.Synchro = &v
	return s
}

func (s *ListAIStaffChatEventsResponseBody) Validate() error {
	if s.Module != nil {
		if err := s.Module.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListAIStaffChatEventsResponseBodyModule struct {
	// Unique ID of the sentence
	//
	// example:
	//
	// chat-xrz3etcl2bsygwlx8g
	ChatId *string `json:"ChatId,omitempty" xml:"ChatId,omitempty"`
	// session ID
	//
	// example:
	//
	// 872be9bc-3097-433d-b462-596202455102
	ConversationId *string `json:"ConversationId,omitempty" xml:"ConversationId,omitempty"`
	// object ID
	Events []*ListAIStaffChatEventsResponseBodyModuleEvents `json:"Events,omitempty" xml:"Events,omitempty" type:"Repeated"`
	// ID of the last SSE event
	//
	// example:
	//
	// event-21dd6124760a4a259ae33bbd878f6e20
	LastEventId *int32 `json:"LastEventId,omitempty" xml:"LastEventId,omitempty"`
}

func (s ListAIStaffChatEventsResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s ListAIStaffChatEventsResponseBodyModule) GoString() string {
	return s.String()
}

func (s *ListAIStaffChatEventsResponseBodyModule) GetChatId() *string {
	return s.ChatId
}

func (s *ListAIStaffChatEventsResponseBodyModule) GetConversationId() *string {
	return s.ConversationId
}

func (s *ListAIStaffChatEventsResponseBodyModule) GetEvents() []*ListAIStaffChatEventsResponseBodyModuleEvents {
	return s.Events
}

func (s *ListAIStaffChatEventsResponseBodyModule) GetLastEventId() *int32 {
	return s.LastEventId
}

func (s *ListAIStaffChatEventsResponseBodyModule) SetChatId(v string) *ListAIStaffChatEventsResponseBodyModule {
	s.ChatId = &v
	return s
}

func (s *ListAIStaffChatEventsResponseBodyModule) SetConversationId(v string) *ListAIStaffChatEventsResponseBodyModule {
	s.ConversationId = &v
	return s
}

func (s *ListAIStaffChatEventsResponseBodyModule) SetEvents(v []*ListAIStaffChatEventsResponseBodyModuleEvents) *ListAIStaffChatEventsResponseBodyModule {
	s.Events = v
	return s
}

func (s *ListAIStaffChatEventsResponseBodyModule) SetLastEventId(v int32) *ListAIStaffChatEventsResponseBodyModule {
	s.LastEventId = &v
	return s
}

func (s *ListAIStaffChatEventsResponseBodyModule) Validate() error {
	if s.Events != nil {
		for _, item := range s.Events {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAIStaffChatEventsResponseBodyModuleEvents struct {
	// error message.
	//
	// example:
	//
	// {\\"Recorded\\": False}
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// primary key
	//
	// example:
	//
	// 10426
	Id *int32 `json:"Id,omitempty" xml:"Id,omitempty"`
	// Website Name
	//
	// example:
	//
	// 文章素材2026050704
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ListAIStaffChatEventsResponseBodyModuleEvents) String() string {
	return dara.Prettify(s)
}

func (s ListAIStaffChatEventsResponseBodyModuleEvents) GoString() string {
	return s.String()
}

func (s *ListAIStaffChatEventsResponseBodyModuleEvents) GetData() *string {
	return s.Data
}

func (s *ListAIStaffChatEventsResponseBodyModuleEvents) GetId() *int32 {
	return s.Id
}

func (s *ListAIStaffChatEventsResponseBodyModuleEvents) GetName() *string {
	return s.Name
}

func (s *ListAIStaffChatEventsResponseBodyModuleEvents) SetData(v string) *ListAIStaffChatEventsResponseBodyModuleEvents {
	s.Data = &v
	return s
}

func (s *ListAIStaffChatEventsResponseBodyModuleEvents) SetId(v int32) *ListAIStaffChatEventsResponseBodyModuleEvents {
	s.Id = &v
	return s
}

func (s *ListAIStaffChatEventsResponseBodyModuleEvents) SetName(v string) *ListAIStaffChatEventsResponseBodyModuleEvents {
	s.Name = &v
	return s
}

func (s *ListAIStaffChatEventsResponseBodyModuleEvents) Validate() error {
	return dara.Validate(s)
}
