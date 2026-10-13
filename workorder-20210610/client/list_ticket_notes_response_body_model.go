// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTicketNotesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ListTicketNotesResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v int32) *ListTicketNotesResponseBody
	GetCode() *int32
	SetData(v []*ListTicketNotesResponseBodyData) *ListTicketNotesResponseBody
	GetData() []*ListTicketNotesResponseBodyData
	SetMessage(v string) *ListTicketNotesResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListTicketNotesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListTicketNotesResponseBody
	GetSuccess() *bool
}

type ListTicketNotesResponseBody struct {
	// The details of the permission error. An empty string is returned if you have the required permissions.
	//
	// example:
	//
	// {\\"AuthAction\\":\\"ram:ListTicketNotes\\",\\"AuthPrincipalDisplayName\\":\\"2146216584788xxxxx\\",\\"AuthPrincipalOwnerId\\":\\"1135850448xxxxx\\",\\"AuthPrincipalType\\":\\"SubUser\\",\\"NoPermissionType\\":\\"ImplicitDeny\\",\\"PolicyType\\":\\"AccountLevelIdentityBasedPolicy\\"}
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The return code for the API request.
	//
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The return value, which is a list of communication records for the specified ticket.
	Data []*ListTicketNotesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The error message. This parameter is returned when \\`Success\\` is \\`false\\`.
	//
	// example:
	//
	// Invalid input parameter
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The unique ID of the request. Each request has a unique ID.
	//
	// example:
	//
	// AC0AB2EC-AFBC-44BA-AE77-132A5A1EC0AD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful. A value of \\`true\\` indicates a successful call.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListTicketNotesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTicketNotesResponseBody) GoString() string {
	return s.String()
}

func (s *ListTicketNotesResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ListTicketNotesResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListTicketNotesResponseBody) GetData() []*ListTicketNotesResponseBodyData {
	return s.Data
}

func (s *ListTicketNotesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListTicketNotesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTicketNotesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListTicketNotesResponseBody) SetAccessDeniedDetail(v string) *ListTicketNotesResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ListTicketNotesResponseBody) SetCode(v int32) *ListTicketNotesResponseBody {
	s.Code = &v
	return s
}

func (s *ListTicketNotesResponseBody) SetData(v []*ListTicketNotesResponseBodyData) *ListTicketNotesResponseBody {
	s.Data = v
	return s
}

func (s *ListTicketNotesResponseBody) SetMessage(v string) *ListTicketNotesResponseBody {
	s.Message = &v
	return s
}

func (s *ListTicketNotesResponseBody) SetRequestId(v string) *ListTicketNotesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTicketNotesResponseBody) SetSuccess(v bool) *ListTicketNotesResponseBody {
	s.Success = &v
	return s
}

func (s *ListTicketNotesResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListTicketNotesResponseBodyData struct {
	// The list of attachments.
	Attachments []*ListTicketNotesResponseBodyDataAttachments `json:"Attachments,omitempty" xml:"Attachments,omitempty" type:"Repeated"`
	// The UNIX timestamp when the communication message was created.
	//
	// example:
	//
	// 1623396736000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ticket communication record object.
	Dialog *ListTicketNotesResponseBodyDataDialog `json:"Dialog,omitempty" xml:"Dialog,omitempty" type:"Struct"`
	// The unique ID of the communication record.
	//
	// example:
	//
	// 9999
	DialogId *int64 `json:"DialogId,omitempty" xml:"DialogId,omitempty"`
	// The status of the communication message. Valid values: \\`1\\`: Unread. \\`2\\`: Read.
	//
	// example:
	//
	// 6
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// This parameter is not in use.
	//
	// example:
	//
	// null
	Tip *string `json:"Tip,omitempty" xml:"Tip,omitempty"`
	// The type of communication. Valid values: \\`1\\`: Card, which corresponds to \\`Schema\\`. \\`2\\`: Text, which corresponds to \\`Content\\`.
	//
	// example:
	//
	// 1
	Type *int32 `json:"Type,omitempty" xml:"Type,omitempty"`
	// The user in the communication.
	User *ListTicketNotesResponseBodyDataUser `json:"User,omitempty" xml:"User,omitempty" type:"Struct"`
}

func (s ListTicketNotesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListTicketNotesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListTicketNotesResponseBodyData) GetAttachments() []*ListTicketNotesResponseBodyDataAttachments {
	return s.Attachments
}

func (s *ListTicketNotesResponseBodyData) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListTicketNotesResponseBodyData) GetDialog() *ListTicketNotesResponseBodyDataDialog {
	return s.Dialog
}

func (s *ListTicketNotesResponseBodyData) GetDialogId() *int64 {
	return s.DialogId
}

func (s *ListTicketNotesResponseBodyData) GetStatus() *int32 {
	return s.Status
}

func (s *ListTicketNotesResponseBodyData) GetTip() *string {
	return s.Tip
}

func (s *ListTicketNotesResponseBodyData) GetType() *int32 {
	return s.Type
}

func (s *ListTicketNotesResponseBodyData) GetUser() *ListTicketNotesResponseBodyDataUser {
	return s.User
}

func (s *ListTicketNotesResponseBodyData) SetAttachments(v []*ListTicketNotesResponseBodyDataAttachments) *ListTicketNotesResponseBodyData {
	s.Attachments = v
	return s
}

func (s *ListTicketNotesResponseBodyData) SetCreateTime(v int64) *ListTicketNotesResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *ListTicketNotesResponseBodyData) SetDialog(v *ListTicketNotesResponseBodyDataDialog) *ListTicketNotesResponseBodyData {
	s.Dialog = v
	return s
}

func (s *ListTicketNotesResponseBodyData) SetDialogId(v int64) *ListTicketNotesResponseBodyData {
	s.DialogId = &v
	return s
}

func (s *ListTicketNotesResponseBodyData) SetStatus(v int32) *ListTicketNotesResponseBodyData {
	s.Status = &v
	return s
}

func (s *ListTicketNotesResponseBodyData) SetTip(v string) *ListTicketNotesResponseBodyData {
	s.Tip = &v
	return s
}

func (s *ListTicketNotesResponseBodyData) SetType(v int32) *ListTicketNotesResponseBodyData {
	s.Type = &v
	return s
}

func (s *ListTicketNotesResponseBodyData) SetUser(v *ListTicketNotesResponseBodyDataUser) *ListTicketNotesResponseBodyData {
	s.User = v
	return s
}

func (s *ListTicketNotesResponseBodyData) Validate() error {
	if s.Attachments != nil {
		for _, item := range s.Attachments {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Dialog != nil {
		if err := s.Dialog.Validate(); err != nil {
			return err
		}
	}
	if s.User != nil {
		if err := s.User.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListTicketNotesResponseBodyDataAttachments struct {
	// The name of the attachment.
	//
	// example:
	//
	// 003.jpg
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// A temporary URL to access the attachment.
	//
	// example:
	//
	// https://gts-workorder.oss-cn-beijing.aliyuncs.com/20221003/cbc00fb0-b612-4d89-a75b-8d535f750f9f/image.png
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s ListTicketNotesResponseBodyDataAttachments) String() string {
	return dara.Prettify(s)
}

func (s ListTicketNotesResponseBodyDataAttachments) GoString() string {
	return s.String()
}

func (s *ListTicketNotesResponseBodyDataAttachments) GetName() *string {
	return s.Name
}

func (s *ListTicketNotesResponseBodyDataAttachments) GetUrl() *string {
	return s.Url
}

func (s *ListTicketNotesResponseBodyDataAttachments) SetName(v string) *ListTicketNotesResponseBodyDataAttachments {
	s.Name = &v
	return s
}

func (s *ListTicketNotesResponseBodyDataAttachments) SetUrl(v string) *ListTicketNotesResponseBodyDataAttachments {
	s.Url = &v
	return s
}

func (s *ListTicketNotesResponseBodyDataAttachments) Validate() error {
	return dara.Validate(s)
}

type ListTicketNotesResponseBodyDataDialog struct {
	// The content of the ticket communication.
	//
	// example:
	//
	// Hello, my ECS renewal failed. What is the reason?
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The system card for the ticket communication record. This parameter is reserved for future use. Currently, you can use the \\`Content\\` parameter to retrieve the plain text.
	//
	// example:
	//
	// null
	Schema *string `json:"Schema,omitempty" xml:"Schema,omitempty"`
}

func (s ListTicketNotesResponseBodyDataDialog) String() string {
	return dara.Prettify(s)
}

func (s ListTicketNotesResponseBodyDataDialog) GoString() string {
	return s.String()
}

func (s *ListTicketNotesResponseBodyDataDialog) GetContent() *string {
	return s.Content
}

func (s *ListTicketNotesResponseBodyDataDialog) GetSchema() *string {
	return s.Schema
}

func (s *ListTicketNotesResponseBodyDataDialog) SetContent(v string) *ListTicketNotesResponseBodyDataDialog {
	s.Content = &v
	return s
}

func (s *ListTicketNotesResponseBodyDataDialog) SetSchema(v string) *ListTicketNotesResponseBodyDataDialog {
	s.Schema = &v
	return s
}

func (s *ListTicketNotesResponseBodyDataDialog) Validate() error {
	return dara.Validate(s)
}

type ListTicketNotesResponseBodyDataUser struct {
	// The name of the user.
	//
	// example:
	//
	// John
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The role of the user in the communication. This parameter distinguishes between customer service agents and customers. Valid values: \\`2\\`: Customer service agent. \\`3\\`: Customer.
	//
	// example:
	//
	// 3
	Role *int32 `json:"Role,omitempty" xml:"Role,omitempty"`
}

func (s ListTicketNotesResponseBodyDataUser) String() string {
	return dara.Prettify(s)
}

func (s ListTicketNotesResponseBodyDataUser) GoString() string {
	return s.String()
}

func (s *ListTicketNotesResponseBodyDataUser) GetName() *string {
	return s.Name
}

func (s *ListTicketNotesResponseBodyDataUser) GetRole() *int32 {
	return s.Role
}

func (s *ListTicketNotesResponseBodyDataUser) SetName(v string) *ListTicketNotesResponseBodyDataUser {
	s.Name = &v
	return s
}

func (s *ListTicketNotesResponseBodyDataUser) SetRole(v int32) *ListTicketNotesResponseBodyDataUser {
	s.Role = &v
	return s
}

func (s *ListTicketNotesResponseBodyDataUser) Validate() error {
	return dara.Validate(s)
}
