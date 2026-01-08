// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetConversationalAutomationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *GetConversationalAutomationResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *GetConversationalAutomationResponseBody
	GetCode() *string
	SetData(v *GetConversationalAutomationResponseBodyData) *GetConversationalAutomationResponseBody
	GetData() *GetConversationalAutomationResponseBodyData
	SetMessage(v string) *GetConversationalAutomationResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetConversationalAutomationResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetConversationalAutomationResponseBody
	GetSuccess() *bool
}

type GetConversationalAutomationResponseBody struct {
	// The details about the access denial.
	//
	// example:
	//
	// None
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The response code.
	//
	// 	- The value OK indicates that the request was successful.
	//
	// 	- For more information about other response codes, see [Error codes](https://help.aliyun.com/document_detail/196974.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *GetConversationalAutomationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message.
	//
	// example:
	//
	// None
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 90E63D28-E31D-1EB2-8939-A94866411B2O
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetConversationalAutomationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetConversationalAutomationResponseBody) GoString() string {
	return s.String()
}

func (s *GetConversationalAutomationResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *GetConversationalAutomationResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetConversationalAutomationResponseBody) GetData() *GetConversationalAutomationResponseBodyData {
	return s.Data
}

func (s *GetConversationalAutomationResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetConversationalAutomationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetConversationalAutomationResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetConversationalAutomationResponseBody) SetAccessDeniedDetail(v string) *GetConversationalAutomationResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *GetConversationalAutomationResponseBody) SetCode(v string) *GetConversationalAutomationResponseBody {
	s.Code = &v
	return s
}

func (s *GetConversationalAutomationResponseBody) SetData(v *GetConversationalAutomationResponseBodyData) *GetConversationalAutomationResponseBody {
	s.Data = v
	return s
}

func (s *GetConversationalAutomationResponseBody) SetMessage(v string) *GetConversationalAutomationResponseBody {
	s.Message = &v
	return s
}

func (s *GetConversationalAutomationResponseBody) SetRequestId(v string) *GetConversationalAutomationResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetConversationalAutomationResponseBody) SetSuccess(v bool) *GetConversationalAutomationResponseBody {
	s.Success = &v
	return s
}

func (s *GetConversationalAutomationResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetConversationalAutomationResponseBodyData struct {
	// The commands.
	Commands []*GetConversationalAutomationResponseBodyDataCommands `json:"Commands,omitempty" xml:"Commands,omitempty" type:"Repeated"`
	// Indicates whether the welcoming message is enabled.
	//
	// example:
	//
	// true
	EnableWelcomeMessage *bool `json:"EnableWelcomeMessage,omitempty" xml:"EnableWelcomeMessage,omitempty"`
	// The phone number of the enterprise.
	//
	// example:
	//
	// 86138****
	PhoneNumber *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	// The opening remarks.
	Prompts []*string `json:"Prompts,omitempty" xml:"Prompts,omitempty" type:"Repeated"`
}

func (s GetConversationalAutomationResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetConversationalAutomationResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetConversationalAutomationResponseBodyData) GetCommands() []*GetConversationalAutomationResponseBodyDataCommands {
	return s.Commands
}

func (s *GetConversationalAutomationResponseBodyData) GetEnableWelcomeMessage() *bool {
	return s.EnableWelcomeMessage
}

func (s *GetConversationalAutomationResponseBodyData) GetPhoneNumber() *string {
	return s.PhoneNumber
}

func (s *GetConversationalAutomationResponseBodyData) GetPrompts() []*string {
	return s.Prompts
}

func (s *GetConversationalAutomationResponseBodyData) SetCommands(v []*GetConversationalAutomationResponseBodyDataCommands) *GetConversationalAutomationResponseBodyData {
	s.Commands = v
	return s
}

func (s *GetConversationalAutomationResponseBodyData) SetEnableWelcomeMessage(v bool) *GetConversationalAutomationResponseBodyData {
	s.EnableWelcomeMessage = &v
	return s
}

func (s *GetConversationalAutomationResponseBodyData) SetPhoneNumber(v string) *GetConversationalAutomationResponseBodyData {
	s.PhoneNumber = &v
	return s
}

func (s *GetConversationalAutomationResponseBodyData) SetPrompts(v []*string) *GetConversationalAutomationResponseBodyData {
	s.Prompts = v
	return s
}

func (s *GetConversationalAutomationResponseBodyData) Validate() error {
	if s.Commands != nil {
		for _, item := range s.Commands {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetConversationalAutomationResponseBodyDataCommands struct {
	// The description of the command.
	//
	// example:
	//
	// description
	CommandDescription *string `json:"CommandDescription,omitempty" xml:"CommandDescription,omitempty"`
	// The name of the command.
	//
	// example:
	//
	// common1
	CommandName *string `json:"CommandName,omitempty" xml:"CommandName,omitempty"`
}

func (s GetConversationalAutomationResponseBodyDataCommands) String() string {
	return dara.Prettify(s)
}

func (s GetConversationalAutomationResponseBodyDataCommands) GoString() string {
	return s.String()
}

func (s *GetConversationalAutomationResponseBodyDataCommands) GetCommandDescription() *string {
	return s.CommandDescription
}

func (s *GetConversationalAutomationResponseBodyDataCommands) GetCommandName() *string {
	return s.CommandName
}

func (s *GetConversationalAutomationResponseBodyDataCommands) SetCommandDescription(v string) *GetConversationalAutomationResponseBodyDataCommands {
	s.CommandDescription = &v
	return s
}

func (s *GetConversationalAutomationResponseBodyDataCommands) SetCommandName(v string) *GetConversationalAutomationResponseBodyDataCommands {
	s.CommandName = &v
	return s
}

func (s *GetConversationalAutomationResponseBodyDataCommands) Validate() error {
	return dara.Validate(s)
}
