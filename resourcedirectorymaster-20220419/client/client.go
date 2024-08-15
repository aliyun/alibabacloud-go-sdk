// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type AcceptHandshakeRequest struct {
	// The ID of the invitation.
	//
	// You can call the [ListHandshakesForAccount](~~ListHandshakesForAccount~~) operation to obtain the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// h-Ih8IuPfvV0t0****
	HandshakeId *string `json:"HandshakeId,omitempty" xml:"HandshakeId,omitempty"`
}

func (s AcceptHandshakeRequest) String() string {
	return tea.Prettify(s)
}

func (s AcceptHandshakeRequest) GoString() string {
	return s.String()
}

func (s *AcceptHandshakeRequest) SetHandshakeId(v string) *AcceptHandshakeRequest {
	s.HandshakeId = &v
	return s
}

type AcceptHandshakeResponseBody struct {
	// The information of the invitation.
	Handshake *AcceptHandshakeResponseBodyHandshake `json:"Handshake,omitempty" xml:"Handshake,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 5828C836-3286-49A6-9006-15231BB19342
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AcceptHandshakeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AcceptHandshakeResponseBody) GoString() string {
	return s.String()
}

func (s *AcceptHandshakeResponseBody) SetHandshake(v *AcceptHandshakeResponseBodyHandshake) *AcceptHandshakeResponseBody {
	s.Handshake = v
	return s
}

func (s *AcceptHandshakeResponseBody) SetRequestId(v string) *AcceptHandshakeResponseBody {
	s.RequestId = &v
	return s
}

type AcceptHandshakeResponseBodyHandshake struct {
	// The time when the invitation was created. The time is displayed in UTC.
	//
	// example:
	//
	// 2021-01-06T02:15:40Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The time when the invitation expires. The time is displayed in UTC.
	//
	// example:
	//
	// 2021-01-20T02:15:40Z
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The ID of the invitation.
	//
	// example:
	//
	// h-Ih8IuPfvV0t0****
	HandshakeId *string `json:"HandshakeId,omitempty" xml:"HandshakeId,omitempty"`
	// The ID of the management account of the resource directory.
	//
	// example:
	//
	// 151266687691****
	MasterAccountId *string `json:"MasterAccountId,omitempty" xml:"MasterAccountId,omitempty"`
	// The name of the management account of the resource directory.
	//
	// example:
	//
	// CompanyA
	MasterAccountName *string `json:"MasterAccountName,omitempty" xml:"MasterAccountName,omitempty"`
	// The time when the invitation was modified. The time is displayed in UTC.
	//
	// example:
	//
	// 2021-01-06T02:16:40Z
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The description of the invitation.
	//
	// example:
	//
	// Welcome
	Note *string `json:"Note,omitempty" xml:"Note,omitempty"`
	// The ID of the resource directory.
	//
	// example:
	//
	// rd-3G****
	ResourceDirectoryId *string `json:"ResourceDirectoryId,omitempty" xml:"ResourceDirectoryId,omitempty"`
	// The status of the invitation. Valid values:
	//
	// 	- Pending: The invitation is waiting for confirmation.
	//
	// 	- Accepted: The invitation is accepted.
	//
	// 	- Cancelled: The invitation is canceled.
	//
	// 	- Declined: The invitation is rejected.
	//
	// 	- Expired: The invitation expires.
	//
	// example:
	//
	// Accepted
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID or logon email address of the invited Alibaba Cloud account.
	//
	// example:
	//
	// 177242285274****
	TargetEntity *string `json:"TargetEntity,omitempty" xml:"TargetEntity,omitempty"`
	// The type of the invited Alibaba Cloud account. Valid values:
	//
	// 	- Account: indicates the ID of the Alibaba Cloud account.
	//
	// 	- Email: indicates the logon email address of the Alibaba Cloud account.
	//
	// example:
	//
	// Account
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
}

func (s AcceptHandshakeResponseBodyHandshake) String() string {
	return tea.Prettify(s)
}

func (s AcceptHandshakeResponseBodyHandshake) GoString() string {
	return s.String()
}

func (s *AcceptHandshakeResponseBodyHandshake) SetCreateTime(v string) *AcceptHandshakeResponseBodyHandshake {
	s.CreateTime = &v
	return s
}

func (s *AcceptHandshakeResponseBodyHandshake) SetExpireTime(v string) *AcceptHandshakeResponseBodyHandshake {
	s.ExpireTime = &v
	return s
}

func (s *AcceptHandshakeResponseBodyHandshake) SetHandshakeId(v string) *AcceptHandshakeResponseBodyHandshake {
	s.HandshakeId = &v
	return s
}

func (s *AcceptHandshakeResponseBodyHandshake) SetMasterAccountId(v string) *AcceptHandshakeResponseBodyHandshake {
	s.MasterAccountId = &v
	return s
}

func (s *AcceptHandshakeResponseBodyHandshake) SetMasterAccountName(v string) *AcceptHandshakeResponseBodyHandshake {
	s.MasterAccountName = &v
	return s
}

func (s *AcceptHandshakeResponseBodyHandshake) SetModifyTime(v string) *AcceptHandshakeResponseBodyHandshake {
	s.ModifyTime = &v
	return s
}

func (s *AcceptHandshakeResponseBodyHandshake) SetNote(v string) *AcceptHandshakeResponseBodyHandshake {
	s.Note = &v
	return s
}

func (s *AcceptHandshakeResponseBodyHandshake) SetResourceDirectoryId(v string) *AcceptHandshakeResponseBodyHandshake {
	s.ResourceDirectoryId = &v
	return s
}

func (s *AcceptHandshakeResponseBodyHandshake) SetStatus(v string) *AcceptHandshakeResponseBodyHandshake {
	s.Status = &v
	return s
}

func (s *AcceptHandshakeResponseBodyHandshake) SetTargetEntity(v string) *AcceptHandshakeResponseBodyHandshake {
	s.TargetEntity = &v
	return s
}

func (s *AcceptHandshakeResponseBodyHandshake) SetTargetType(v string) *AcceptHandshakeResponseBodyHandshake {
	s.TargetType = &v
	return s
}

type AcceptHandshakeResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AcceptHandshakeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AcceptHandshakeResponse) String() string {
	return tea.Prettify(s)
}

func (s AcceptHandshakeResponse) GoString() string {
	return s.String()
}

func (s *AcceptHandshakeResponse) SetHeaders(v map[string]*string) *AcceptHandshakeResponse {
	s.Headers = v
	return s
}

func (s *AcceptHandshakeResponse) SetStatusCode(v int32) *AcceptHandshakeResponse {
	s.StatusCode = &v
	return s
}

func (s *AcceptHandshakeResponse) SetBody(v *AcceptHandshakeResponseBody) *AcceptHandshakeResponse {
	s.Body = v
	return s
}

type AddMessageContactRequest struct {
	// The email address of the contact.
	//
	// After you specify an email address, you need to call [SendEmailVerificationForMessageContact](~~SendEmailVerificationForMessageContact~~) to send verification information to the email address. After the verification is passed, the email address takes effect.
	//
	// This parameter is required.
	//
	// example:
	//
	// someone***@example.com
	EmailAddress *string `json:"EmailAddress,omitempty" xml:"EmailAddress,omitempty"`
	// The types of messages received by the contact.
	//
	// This parameter is required.
	MessageTypes []*string `json:"MessageTypes,omitempty" xml:"MessageTypes,omitempty" type:"Repeated"`
	// The name of the contact.
	//
	// The name must be unique in your resource directory.
	//
	// The name must be 2 to 12 characters in length and can contain only letters.
	//
	// This parameter is required.
	//
	// example:
	//
	// tom
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The mobile phone number of the contact.
	//
	// Specify the mobile phone number in the `<Country code>-<Mobile phone number>` format.
	//
	// > Only mobile phone numbers in the `86-<Mobile phone number>` format in the Chinese mainland are supported.
	//
	// After you specify a mobile phone number, you need to call [SendPhoneVerificationForMessageContact](~~SendPhoneVerificationForMessageContact~~) to send verification information to the mobile phone number. After the verification is passed, the mobile phone number takes effect.
	//
	// example:
	//
	// 86-139****1234
	PhoneNumber *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	// The job title of the contact.
	//
	// Valid values:
	//
	// 	- FinanceDirector
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- TechnicalDirector
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- MaintenanceDirector
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- CEO
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- ProjectDirector
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- Other
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// This parameter is required.
	//
	// example:
	//
	// TechnicalDirector
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s AddMessageContactRequest) String() string {
	return tea.Prettify(s)
}

func (s AddMessageContactRequest) GoString() string {
	return s.String()
}

func (s *AddMessageContactRequest) SetEmailAddress(v string) *AddMessageContactRequest {
	s.EmailAddress = &v
	return s
}

func (s *AddMessageContactRequest) SetMessageTypes(v []*string) *AddMessageContactRequest {
	s.MessageTypes = v
	return s
}

func (s *AddMessageContactRequest) SetName(v string) *AddMessageContactRequest {
	s.Name = &v
	return s
}

func (s *AddMessageContactRequest) SetPhoneNumber(v string) *AddMessageContactRequest {
	s.PhoneNumber = &v
	return s
}

func (s *AddMessageContactRequest) SetTitle(v string) *AddMessageContactRequest {
	s.Title = &v
	return s
}

type AddMessageContactResponseBody struct {
	// The information about the contact.
	Contact *AddMessageContactResponseBodyContact `json:"Contact,omitempty" xml:"Contact,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 2DFCE4C9-04A9-4C83-BB14-FE791275EC53
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddMessageContactResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddMessageContactResponseBody) GoString() string {
	return s.String()
}

func (s *AddMessageContactResponseBody) SetContact(v *AddMessageContactResponseBodyContact) *AddMessageContactResponseBody {
	s.Contact = v
	return s
}

func (s *AddMessageContactResponseBody) SetRequestId(v string) *AddMessageContactResponseBody {
	s.RequestId = &v
	return s
}

type AddMessageContactResponseBodyContact struct {
	// The ID of the contact.
	//
	// example:
	//
	// c-qL4HqKONzOM7****
	ContactId *string `json:"ContactId,omitempty" xml:"ContactId,omitempty"`
	// The time when the contact was created.
	//
	// example:
	//
	// 2023-03-27 17:19:21
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
}

func (s AddMessageContactResponseBodyContact) String() string {
	return tea.Prettify(s)
}

func (s AddMessageContactResponseBodyContact) GoString() string {
	return s.String()
}

func (s *AddMessageContactResponseBodyContact) SetContactId(v string) *AddMessageContactResponseBodyContact {
	s.ContactId = &v
	return s
}

func (s *AddMessageContactResponseBodyContact) SetCreateDate(v string) *AddMessageContactResponseBodyContact {
	s.CreateDate = &v
	return s
}

type AddMessageContactResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddMessageContactResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddMessageContactResponse) String() string {
	return tea.Prettify(s)
}

func (s AddMessageContactResponse) GoString() string {
	return s.String()
}

func (s *AddMessageContactResponse) SetHeaders(v map[string]*string) *AddMessageContactResponse {
	s.Headers = v
	return s
}

func (s *AddMessageContactResponse) SetStatusCode(v int32) *AddMessageContactResponse {
	s.StatusCode = &v
	return s
}

func (s *AddMessageContactResponse) SetBody(v *AddMessageContactResponseBody) *AddMessageContactResponse {
	s.Body = v
	return s
}

type AssociateMembersRequest struct {
	// The ID of the contact.
	//
	// example:
	//
	// c-qL4HqKONzOM7****
	ContactId *string `json:"ContactId,omitempty" xml:"ContactId,omitempty"`
	// The IDs of objects to which you want to bind the contact.
	Members []*string `json:"Members,omitempty" xml:"Members,omitempty" type:"Repeated"`
}

func (s AssociateMembersRequest) String() string {
	return tea.Prettify(s)
}

func (s AssociateMembersRequest) GoString() string {
	return s.String()
}

func (s *AssociateMembersRequest) SetContactId(v string) *AssociateMembersRequest {
	s.ContactId = &v
	return s
}

func (s *AssociateMembersRequest) SetMembers(v []*string) *AssociateMembersRequest {
	s.Members = v
	return s
}

type AssociateMembersResponseBody struct {
	// The time when the contact was bound to the object.
	Members []*AssociateMembersResponseBodyMembers `json:"Members,omitempty" xml:"Members,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 95060F1D-6990-4645-8920-A81D1BBFE992
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AssociateMembersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AssociateMembersResponseBody) GoString() string {
	return s.String()
}

func (s *AssociateMembersResponseBody) SetMembers(v []*AssociateMembersResponseBodyMembers) *AssociateMembersResponseBody {
	s.Members = v
	return s
}

func (s *AssociateMembersResponseBody) SetRequestId(v string) *AssociateMembersResponseBody {
	s.RequestId = &v
	return s
}

type AssociateMembersResponseBodyMembers struct {
	// The ID of the contact.
	//
	// example:
	//
	// c-qL4HqKONzOM7****
	ContactId *string `json:"ContactId,omitempty" xml:"ContactId,omitempty"`
	// The ID of the object. Valid values:
	//
	// - ID of the resource directory
	//
	// - ID of the folder
	//
	// - ID of the member
	//
	// example:
	//
	// fd-ZDNPiT****
	MemberId *string `json:"MemberId,omitempty" xml:"MemberId,omitempty"`
	// The time when the contact was bound to the object.
	//
	// example:
	//
	// 2023-03-27 17:19:21
	ModifyDate *string `json:"ModifyDate,omitempty" xml:"ModifyDate,omitempty"`
}

func (s AssociateMembersResponseBodyMembers) String() string {
	return tea.Prettify(s)
}

func (s AssociateMembersResponseBodyMembers) GoString() string {
	return s.String()
}

func (s *AssociateMembersResponseBodyMembers) SetContactId(v string) *AssociateMembersResponseBodyMembers {
	s.ContactId = &v
	return s
}

func (s *AssociateMembersResponseBodyMembers) SetMemberId(v string) *AssociateMembersResponseBodyMembers {
	s.MemberId = &v
	return s
}

func (s *AssociateMembersResponseBodyMembers) SetModifyDate(v string) *AssociateMembersResponseBodyMembers {
	s.ModifyDate = &v
	return s
}

type AssociateMembersResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AssociateMembersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AssociateMembersResponse) String() string {
	return tea.Prettify(s)
}

func (s AssociateMembersResponse) GoString() string {
	return s.String()
}

func (s *AssociateMembersResponse) SetHeaders(v map[string]*string) *AssociateMembersResponse {
	s.Headers = v
	return s
}

func (s *AssociateMembersResponse) SetStatusCode(v int32) *AssociateMembersResponse {
	s.StatusCode = &v
	return s
}

func (s *AssociateMembersResponse) SetBody(v *AssociateMembersResponseBody) *AssociateMembersResponse {
	s.Body = v
	return s
}

type AttachControlPolicyRequest struct {
	// The ID of the access control policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// cp-jExXAqIYkwHN****
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// The ID of the object to which you want to attach the access control policy. Access control policies can be attached to the following objects:
	//
	// 	- Root folder
	//
	// 	- Subfolders of the Root folder
	//
	// 	- Members
	//
	// This parameter is required.
	//
	// example:
	//
	// fd-ZDNPiT****
	TargetId *string `json:"TargetId,omitempty" xml:"TargetId,omitempty"`
}

func (s AttachControlPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s AttachControlPolicyRequest) GoString() string {
	return s.String()
}

func (s *AttachControlPolicyRequest) SetPolicyId(v string) *AttachControlPolicyRequest {
	s.PolicyId = &v
	return s
}

func (s *AttachControlPolicyRequest) SetTargetId(v string) *AttachControlPolicyRequest {
	s.TargetId = &v
	return s
}

type AttachControlPolicyResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 95060F1D-6990-4645-8920-A81D1BBFE992
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AttachControlPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AttachControlPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *AttachControlPolicyResponseBody) SetRequestId(v string) *AttachControlPolicyResponseBody {
	s.RequestId = &v
	return s
}

type AttachControlPolicyResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AttachControlPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AttachControlPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s AttachControlPolicyResponse) GoString() string {
	return s.String()
}

func (s *AttachControlPolicyResponse) SetHeaders(v map[string]*string) *AttachControlPolicyResponse {
	s.Headers = v
	return s
}

func (s *AttachControlPolicyResponse) SetStatusCode(v int32) *AttachControlPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *AttachControlPolicyResponse) SetBody(v *AttachControlPolicyResponseBody) *AttachControlPolicyResponse {
	s.Body = v
	return s
}

type BindSecureMobilePhoneRequest struct {
	// The Alibaba Cloud account ID of the member.
	//
	// This parameter is required.
	//
	// example:
	//
	// 138660628348****
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The mobile phone number that you want to bind to the member for security purposes.
	//
	// The mobile phone number you specify must be the same as the mobile phone number that you specify when you call the [SendVerificationCodeForBindSecureMobilePhone](~~SendVerificationCodeForBindSecureMobilePhone~~) operation to obtain a verification code.
	//
	// Specify the mobile phone number in the \\<Country code>-\\<Mobile phone number> format.
	//
	// > Mobile phone numbers in the `86-<Mobile phone number>` format in the Chinese mainland are not supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// xx-13900001234
	SecureMobilePhone *string `json:"SecureMobilePhone,omitempty" xml:"SecureMobilePhone,omitempty"`
	// The verification code.
	//
	// You can call the [SendVerificationCodeForBindSecureMobilePhone](~~SendVerificationCodeForBindSecureMobilePhone~~) operation to obtain the verification code.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456
	VerificationCode *string `json:"VerificationCode,omitempty" xml:"VerificationCode,omitempty"`
}

func (s BindSecureMobilePhoneRequest) String() string {
	return tea.Prettify(s)
}

func (s BindSecureMobilePhoneRequest) GoString() string {
	return s.String()
}

func (s *BindSecureMobilePhoneRequest) SetAccountId(v string) *BindSecureMobilePhoneRequest {
	s.AccountId = &v
	return s
}

func (s *BindSecureMobilePhoneRequest) SetSecureMobilePhone(v string) *BindSecureMobilePhoneRequest {
	s.SecureMobilePhone = &v
	return s
}

func (s *BindSecureMobilePhoneRequest) SetVerificationCode(v string) *BindSecureMobilePhoneRequest {
	s.VerificationCode = &v
	return s
}

type BindSecureMobilePhoneResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 0217AFEB-5318-56D4-B167-1933D83EDF3F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BindSecureMobilePhoneResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BindSecureMobilePhoneResponseBody) GoString() string {
	return s.String()
}

func (s *BindSecureMobilePhoneResponseBody) SetRequestId(v string) *BindSecureMobilePhoneResponseBody {
	s.RequestId = &v
	return s
}

type BindSecureMobilePhoneResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BindSecureMobilePhoneResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BindSecureMobilePhoneResponse) String() string {
	return tea.Prettify(s)
}

func (s BindSecureMobilePhoneResponse) GoString() string {
	return s.String()
}

func (s *BindSecureMobilePhoneResponse) SetHeaders(v map[string]*string) *BindSecureMobilePhoneResponse {
	s.Headers = v
	return s
}

func (s *BindSecureMobilePhoneResponse) SetStatusCode(v int32) *BindSecureMobilePhoneResponse {
	s.StatusCode = &v
	return s
}

func (s *BindSecureMobilePhoneResponse) SetBody(v *BindSecureMobilePhoneResponseBody) *BindSecureMobilePhoneResponse {
	s.Body = v
	return s
}

type CancelChangeAccountEmailRequest struct {
	// The Alibaba Cloud account ID of the member.
	//
	// This parameter is required.
	//
	// example:
	//
	// 181761095690****
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
}

func (s CancelChangeAccountEmailRequest) String() string {
	return tea.Prettify(s)
}

func (s CancelChangeAccountEmailRequest) GoString() string {
	return s.String()
}

func (s *CancelChangeAccountEmailRequest) SetAccountId(v string) *CancelChangeAccountEmailRequest {
	s.AccountId = &v
	return s
}

type CancelChangeAccountEmailResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 9B34724D-54B0-4A51-B34D-4512372FE1BE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CancelChangeAccountEmailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CancelChangeAccountEmailResponseBody) GoString() string {
	return s.String()
}

func (s *CancelChangeAccountEmailResponseBody) SetRequestId(v string) *CancelChangeAccountEmailResponseBody {
	s.RequestId = &v
	return s
}

type CancelChangeAccountEmailResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CancelChangeAccountEmailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CancelChangeAccountEmailResponse) String() string {
	return tea.Prettify(s)
}

func (s CancelChangeAccountEmailResponse) GoString() string {
	return s.String()
}

func (s *CancelChangeAccountEmailResponse) SetHeaders(v map[string]*string) *CancelChangeAccountEmailResponse {
	s.Headers = v
	return s
}

func (s *CancelChangeAccountEmailResponse) SetStatusCode(v int32) *CancelChangeAccountEmailResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelChangeAccountEmailResponse) SetBody(v *CancelChangeAccountEmailResponseBody) *CancelChangeAccountEmailResponse {
	s.Body = v
	return s
}

type CancelHandshakeRequest struct {
	// The ID of the invitation.
	//
	// This parameter is required.
	//
	// example:
	//
	// h-ycm4rp****
	HandshakeId *string `json:"HandshakeId,omitempty" xml:"HandshakeId,omitempty"`
}

func (s CancelHandshakeRequest) String() string {
	return tea.Prettify(s)
}

func (s CancelHandshakeRequest) GoString() string {
	return s.String()
}

func (s *CancelHandshakeRequest) SetHandshakeId(v string) *CancelHandshakeRequest {
	s.HandshakeId = &v
	return s
}

type CancelHandshakeResponseBody struct {
	// The information of the invitation.
	Handshake *CancelHandshakeResponseBodyHandshake `json:"Handshake,omitempty" xml:"Handshake,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 9B34724D-54B0-4A51-B34D-4512372FE1BE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CancelHandshakeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CancelHandshakeResponseBody) GoString() string {
	return s.String()
}

func (s *CancelHandshakeResponseBody) SetHandshake(v *CancelHandshakeResponseBodyHandshake) *CancelHandshakeResponseBody {
	s.Handshake = v
	return s
}

func (s *CancelHandshakeResponseBody) SetRequestId(v string) *CancelHandshakeResponseBody {
	s.RequestId = &v
	return s
}

type CancelHandshakeResponseBodyHandshake struct {
	// The time when the invitation was created. The time is displayed in UTC.
	//
	// example:
	//
	// 2018-08-10T09:55:41Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The time when the invitation expires. The time is displayed in UTC.
	//
	// example:
	//
	// 2018-08-24T09:55:41Z
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The ID of the invitation.
	//
	// example:
	//
	// h-ycm4rp****
	HandshakeId *string `json:"HandshakeId,omitempty" xml:"HandshakeId,omitempty"`
	// The ID of the management account of the resource directory.
	//
	// example:
	//
	// 172841235500****
	MasterAccountId *string `json:"MasterAccountId,omitempty" xml:"MasterAccountId,omitempty"`
	// The name of the management account of the resource directory.
	//
	// example:
	//
	// Alice
	MasterAccountName *string `json:"MasterAccountName,omitempty" xml:"MasterAccountName,omitempty"`
	// The time when the invitation was modified. The time is displayed in UTC.
	//
	// example:
	//
	// 2018-08-10T09:55:41Z
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The description of the invitation.
	//
	// example:
	//
	// Welcome
	Note *string `json:"Note,omitempty" xml:"Note,omitempty"`
	// The ID of the resource directory.
	//
	// example:
	//
	// h-ycm4rp****
	ResourceDirectoryId *string `json:"ResourceDirectoryId,omitempty" xml:"ResourceDirectoryId,omitempty"`
	// The status of the invitation. Valid values:
	//
	// 	- Pending: The invitation is waiting for confirmation.
	//
	// 	- Accepted: The invitation is accepted.
	//
	// 	- Cancelled: The invitation is canceled.
	//
	// 	- Declined: The invitation is rejected.
	//
	// 	- Expired: The invitation expires.
	//
	// example:
	//
	// Cancelled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID or logon email address of the invited account.
	//
	// example:
	//
	// someone@example.com
	TargetEntity *string `json:"TargetEntity,omitempty" xml:"TargetEntity,omitempty"`
	// The type of the invited account. Valid values:
	//
	// 	- Account: indicates the ID of the account.
	//
	// 	- Email: indicates the logon email address of the account.
	//
	// example:
	//
	// Email
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
}

func (s CancelHandshakeResponseBodyHandshake) String() string {
	return tea.Prettify(s)
}

func (s CancelHandshakeResponseBodyHandshake) GoString() string {
	return s.String()
}

func (s *CancelHandshakeResponseBodyHandshake) SetCreateTime(v string) *CancelHandshakeResponseBodyHandshake {
	s.CreateTime = &v
	return s
}

func (s *CancelHandshakeResponseBodyHandshake) SetExpireTime(v string) *CancelHandshakeResponseBodyHandshake {
	s.ExpireTime = &v
	return s
}

func (s *CancelHandshakeResponseBodyHandshake) SetHandshakeId(v string) *CancelHandshakeResponseBodyHandshake {
	s.HandshakeId = &v
	return s
}

func (s *CancelHandshakeResponseBodyHandshake) SetMasterAccountId(v string) *CancelHandshakeResponseBodyHandshake {
	s.MasterAccountId = &v
	return s
}

func (s *CancelHandshakeResponseBodyHandshake) SetMasterAccountName(v string) *CancelHandshakeResponseBodyHandshake {
	s.MasterAccountName = &v
	return s
}

func (s *CancelHandshakeResponseBodyHandshake) SetModifyTime(v string) *CancelHandshakeResponseBodyHandshake {
	s.ModifyTime = &v
	return s
}

func (s *CancelHandshakeResponseBodyHandshake) SetNote(v string) *CancelHandshakeResponseBodyHandshake {
	s.Note = &v
	return s
}

func (s *CancelHandshakeResponseBodyHandshake) SetResourceDirectoryId(v string) *CancelHandshakeResponseBodyHandshake {
	s.ResourceDirectoryId = &v
	return s
}

func (s *CancelHandshakeResponseBodyHandshake) SetStatus(v string) *CancelHandshakeResponseBodyHandshake {
	s.Status = &v
	return s
}

func (s *CancelHandshakeResponseBodyHandshake) SetTargetEntity(v string) *CancelHandshakeResponseBodyHandshake {
	s.TargetEntity = &v
	return s
}

func (s *CancelHandshakeResponseBodyHandshake) SetTargetType(v string) *CancelHandshakeResponseBodyHandshake {
	s.TargetType = &v
	return s
}

type CancelHandshakeResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CancelHandshakeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CancelHandshakeResponse) String() string {
	return tea.Prettify(s)
}

func (s CancelHandshakeResponse) GoString() string {
	return s.String()
}

func (s *CancelHandshakeResponse) SetHeaders(v map[string]*string) *CancelHandshakeResponse {
	s.Headers = v
	return s
}

func (s *CancelHandshakeResponse) SetStatusCode(v int32) *CancelHandshakeResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelHandshakeResponse) SetBody(v *CancelHandshakeResponseBody) *CancelHandshakeResponse {
	s.Body = v
	return s
}

type CancelMessageContactUpdateRequest struct {
	// The ID of the contact.
	//
	// example:
	//
	// c-qL4HqKONzOM7****
	ContactId *string `json:"ContactId,omitempty" xml:"ContactId,omitempty"`
	// The email address of the contact.
	//
	// example:
	//
	// someone***@example.com
	EmailAddress *string `json:"EmailAddress,omitempty" xml:"EmailAddress,omitempty"`
	// The mobile phone number of the contact.
	//
	// Specify the mobile phone number in the `<Country code>-<Mobile phone number>` format.
	//
	// example:
	//
	// 86-139****1234
	PhoneNumber *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
}

func (s CancelMessageContactUpdateRequest) String() string {
	return tea.Prettify(s)
}

func (s CancelMessageContactUpdateRequest) GoString() string {
	return s.String()
}

func (s *CancelMessageContactUpdateRequest) SetContactId(v string) *CancelMessageContactUpdateRequest {
	s.ContactId = &v
	return s
}

func (s *CancelMessageContactUpdateRequest) SetEmailAddress(v string) *CancelMessageContactUpdateRequest {
	s.EmailAddress = &v
	return s
}

func (s *CancelMessageContactUpdateRequest) SetPhoneNumber(v string) *CancelMessageContactUpdateRequest {
	s.PhoneNumber = &v
	return s
}

type CancelMessageContactUpdateResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 9B34724D-54B0-4A51-B34D-4512372FE1BE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CancelMessageContactUpdateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CancelMessageContactUpdateResponseBody) GoString() string {
	return s.String()
}

func (s *CancelMessageContactUpdateResponseBody) SetRequestId(v string) *CancelMessageContactUpdateResponseBody {
	s.RequestId = &v
	return s
}

type CancelMessageContactUpdateResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CancelMessageContactUpdateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CancelMessageContactUpdateResponse) String() string {
	return tea.Prettify(s)
}

func (s CancelMessageContactUpdateResponse) GoString() string {
	return s.String()
}

func (s *CancelMessageContactUpdateResponse) SetHeaders(v map[string]*string) *CancelMessageContactUpdateResponse {
	s.Headers = v
	return s
}

func (s *CancelMessageContactUpdateResponse) SetStatusCode(v int32) *CancelMessageContactUpdateResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelMessageContactUpdateResponse) SetBody(v *CancelMessageContactUpdateResponseBody) *CancelMessageContactUpdateResponse {
	s.Body = v
	return s
}

type ChangeAccountEmailRequest struct {
	// The Alibaba Cloud account ID of the member.
	//
	// This parameter is required.
	//
	// example:
	//
	// 181761095690****
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The email address to be bound to the member.
	//
	// > The system automatically sends a verification email to the email address. After the verification is passed, the email address takes effect, and the system changes both the logon email address and secure email address of the member.
	//
	// This parameter is required.
	//
	// example:
	//
	// someone@example.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
}

func (s ChangeAccountEmailRequest) String() string {
	return tea.Prettify(s)
}

func (s ChangeAccountEmailRequest) GoString() string {
	return s.String()
}

func (s *ChangeAccountEmailRequest) SetAccountId(v string) *ChangeAccountEmailRequest {
	s.AccountId = &v
	return s
}

func (s *ChangeAccountEmailRequest) SetEmail(v string) *ChangeAccountEmailRequest {
	s.Email = &v
	return s
}

type ChangeAccountEmailResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 9B34724D-54B0-4A51-B34D-4512372FE1BE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ChangeAccountEmailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ChangeAccountEmailResponseBody) GoString() string {
	return s.String()
}

func (s *ChangeAccountEmailResponseBody) SetRequestId(v string) *ChangeAccountEmailResponseBody {
	s.RequestId = &v
	return s
}

type ChangeAccountEmailResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ChangeAccountEmailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ChangeAccountEmailResponse) String() string {
	return tea.Prettify(s)
}

func (s ChangeAccountEmailResponse) GoString() string {
	return s.String()
}

func (s *ChangeAccountEmailResponse) SetHeaders(v map[string]*string) *ChangeAccountEmailResponse {
	s.Headers = v
	return s
}

func (s *ChangeAccountEmailResponse) SetStatusCode(v int32) *ChangeAccountEmailResponse {
	s.StatusCode = &v
	return s
}

func (s *ChangeAccountEmailResponse) SetBody(v *ChangeAccountEmailResponseBody) *ChangeAccountEmailResponse {
	s.Body = v
	return s
}

type CheckAccountDeleteRequest struct {
	// The Alibaba Cloud account ID of the member that you want to delete.
	//
	// This parameter is required.
	//
	// example:
	//
	// 179855839641****
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
}

func (s CheckAccountDeleteRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckAccountDeleteRequest) GoString() string {
	return s.String()
}

func (s *CheckAccountDeleteRequest) SetAccountId(v string) *CheckAccountDeleteRequest {
	s.AccountId = &v
	return s
}

type CheckAccountDeleteResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 7CDDDCEF-CDFD-0825-B7D7-217BE0897B22
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CheckAccountDeleteResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CheckAccountDeleteResponseBody) GoString() string {
	return s.String()
}

func (s *CheckAccountDeleteResponseBody) SetRequestId(v string) *CheckAccountDeleteResponseBody {
	s.RequestId = &v
	return s
}

type CheckAccountDeleteResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckAccountDeleteResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckAccountDeleteResponse) String() string {
	return tea.Prettify(s)
}

func (s CheckAccountDeleteResponse) GoString() string {
	return s.String()
}

func (s *CheckAccountDeleteResponse) SetHeaders(v map[string]*string) *CheckAccountDeleteResponse {
	s.Headers = v
	return s
}

func (s *CheckAccountDeleteResponse) SetStatusCode(v int32) *CheckAccountDeleteResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckAccountDeleteResponse) SetBody(v *CheckAccountDeleteResponseBody) *CheckAccountDeleteResponse {
	s.Body = v
	return s
}

type CreateControlPolicyRequest struct {
	// The description of the access control policy.
	//
	// The description must be 1 to 1,024 characters in length. The description can contain letters, digits, underscores (_), and hyphens (-) and must start with a letter.
	//
	// example:
	//
	// ExampleControlPolicy
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The effective scope of the access control policy.
	//
	// The value RAM indicates that the access control policy takes effect only for RAM users and RAM roles.
	//
	// This parameter is required.
	//
	// example:
	//
	// RAM
	EffectScope *string `json:"EffectScope,omitempty" xml:"EffectScope,omitempty"`
	// The document of the access control policy.
	//
	// The document can be a maximum of 4,096 characters in length.
	//
	// For more information about the languages of access control policies, see [Languages of access control policies](https://help.aliyun.com/document_detail/179096.html).
	//
	// For more information about the examples of access control policies, see [Examples of custom access control policies](https://help.aliyun.com/document_detail/181474.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// {"Version":"1","Statement":[{"Effect":"Deny","Action":["ram:UpdateRole","ram:DeleteRole","ram:AttachPolicyToRole","ram:DetachPolicyFromRole"],"Resource":"acs:ram:*:*:role/ResourceDirectoryAccountAccessRole"}]}
	PolicyDocument *string `json:"PolicyDocument,omitempty" xml:"PolicyDocument,omitempty"`
	// The name of the access control policy.
	//
	// The name must be 1 to 128 characters in length. The name can contain letters, digits, and hyphens (-) and must start with a letter.
	//
	// This parameter is required.
	//
	// example:
	//
	// ExampleControlPolicy
	PolicyName *string                          `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	Tag        []*CreateControlPolicyRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s CreateControlPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateControlPolicyRequest) GoString() string {
	return s.String()
}

func (s *CreateControlPolicyRequest) SetDescription(v string) *CreateControlPolicyRequest {
	s.Description = &v
	return s
}

func (s *CreateControlPolicyRequest) SetEffectScope(v string) *CreateControlPolicyRequest {
	s.EffectScope = &v
	return s
}

func (s *CreateControlPolicyRequest) SetPolicyDocument(v string) *CreateControlPolicyRequest {
	s.PolicyDocument = &v
	return s
}

func (s *CreateControlPolicyRequest) SetPolicyName(v string) *CreateControlPolicyRequest {
	s.PolicyName = &v
	return s
}

func (s *CreateControlPolicyRequest) SetTag(v []*CreateControlPolicyRequestTag) *CreateControlPolicyRequest {
	s.Tag = v
	return s
}

type CreateControlPolicyRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateControlPolicyRequestTag) String() string {
	return tea.Prettify(s)
}

func (s CreateControlPolicyRequestTag) GoString() string {
	return s.String()
}

func (s *CreateControlPolicyRequestTag) SetKey(v string) *CreateControlPolicyRequestTag {
	s.Key = &v
	return s
}

func (s *CreateControlPolicyRequestTag) SetValue(v string) *CreateControlPolicyRequestTag {
	s.Value = &v
	return s
}

type CreateControlPolicyResponseBody struct {
	// The details of the access control policy.
	ControlPolicy *CreateControlPolicyResponseBodyControlPolicy `json:"ControlPolicy,omitempty" xml:"ControlPolicy,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 776B05B3-A0B0-464B-A191-F4E1119A94B2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateControlPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateControlPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *CreateControlPolicyResponseBody) SetControlPolicy(v *CreateControlPolicyResponseBodyControlPolicy) *CreateControlPolicyResponseBody {
	s.ControlPolicy = v
	return s
}

func (s *CreateControlPolicyResponseBody) SetRequestId(v string) *CreateControlPolicyResponseBody {
	s.RequestId = &v
	return s
}

type CreateControlPolicyResponseBodyControlPolicy struct {
	// The number of times that the access control policy is referenced.
	//
	// example:
	//
	// 0
	AttachmentCount *string `json:"AttachmentCount,omitempty" xml:"AttachmentCount,omitempty"`
	// The time when the access control policy was created.
	//
	// example:
	//
	// 2021-03-18T09:24:19Z
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	// The description of the access control policy.
	//
	// example:
	//
	// ExampleControlPolicy
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The effective scope of the access control policy.
	//
	// The value RAM indicates that the access control policy takes effect only for RAM users and RAM roles.
	//
	// example:
	//
	// RAM
	EffectScope *string `json:"EffectScope,omitempty" xml:"EffectScope,omitempty"`
	// The ID of the access control policy.
	//
	// example:
	//
	// cp-jExXAqIYkwHN****
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// The name of the access control policy.
	//
	// example:
	//
	// ExampleControlPolicy
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	// The type of the access control policy. Valid values:
	//
	// 	- System: system access control policy
	//
	// 	- Custom: custom access control policy
	//
	// example:
	//
	// Custom
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
	// The time when the access control policy was updated.
	//
	// example:
	//
	// 2021-03-18T09:24:19Z
	UpdateDate *string `json:"UpdateDate,omitempty" xml:"UpdateDate,omitempty"`
}

func (s CreateControlPolicyResponseBodyControlPolicy) String() string {
	return tea.Prettify(s)
}

func (s CreateControlPolicyResponseBodyControlPolicy) GoString() string {
	return s.String()
}

func (s *CreateControlPolicyResponseBodyControlPolicy) SetAttachmentCount(v string) *CreateControlPolicyResponseBodyControlPolicy {
	s.AttachmentCount = &v
	return s
}

func (s *CreateControlPolicyResponseBodyControlPolicy) SetCreateDate(v string) *CreateControlPolicyResponseBodyControlPolicy {
	s.CreateDate = &v
	return s
}

func (s *CreateControlPolicyResponseBodyControlPolicy) SetDescription(v string) *CreateControlPolicyResponseBodyControlPolicy {
	s.Description = &v
	return s
}

func (s *CreateControlPolicyResponseBodyControlPolicy) SetEffectScope(v string) *CreateControlPolicyResponseBodyControlPolicy {
	s.EffectScope = &v
	return s
}

func (s *CreateControlPolicyResponseBodyControlPolicy) SetPolicyId(v string) *CreateControlPolicyResponseBodyControlPolicy {
	s.PolicyId = &v
	return s
}

func (s *CreateControlPolicyResponseBodyControlPolicy) SetPolicyName(v string) *CreateControlPolicyResponseBodyControlPolicy {
	s.PolicyName = &v
	return s
}

func (s *CreateControlPolicyResponseBodyControlPolicy) SetPolicyType(v string) *CreateControlPolicyResponseBodyControlPolicy {
	s.PolicyType = &v
	return s
}

func (s *CreateControlPolicyResponseBodyControlPolicy) SetUpdateDate(v string) *CreateControlPolicyResponseBodyControlPolicy {
	s.UpdateDate = &v
	return s
}

type CreateControlPolicyResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateControlPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateControlPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateControlPolicyResponse) GoString() string {
	return s.String()
}

func (s *CreateControlPolicyResponse) SetHeaders(v map[string]*string) *CreateControlPolicyResponse {
	s.Headers = v
	return s
}

func (s *CreateControlPolicyResponse) SetStatusCode(v int32) *CreateControlPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateControlPolicyResponse) SetBody(v *CreateControlPolicyResponseBody) *CreateControlPolicyResponse {
	s.Body = v
	return s
}

type CreateFolderRequest struct {
	// The name of the folder.
	//
	// The name must be 1 to 24 characters in length and can contain letters, digits, underscores (_), periods (.),and hyphens (-).
	//
	// This parameter is required.
	//
	// example:
	//
	// rdFolder
	FolderName *string `json:"FolderName,omitempty" xml:"FolderName,omitempty"`
	// The ID of the parent folder.
	//
	// example:
	//
	// r-b1****
	ParentFolderId *string                   `json:"ParentFolderId,omitempty" xml:"ParentFolderId,omitempty"`
	Tag            []*CreateFolderRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s CreateFolderRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateFolderRequest) GoString() string {
	return s.String()
}

func (s *CreateFolderRequest) SetFolderName(v string) *CreateFolderRequest {
	s.FolderName = &v
	return s
}

func (s *CreateFolderRequest) SetParentFolderId(v string) *CreateFolderRequest {
	s.ParentFolderId = &v
	return s
}

func (s *CreateFolderRequest) SetTag(v []*CreateFolderRequestTag) *CreateFolderRequest {
	s.Tag = v
	return s
}

type CreateFolderRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateFolderRequestTag) String() string {
	return tea.Prettify(s)
}

func (s CreateFolderRequestTag) GoString() string {
	return s.String()
}

func (s *CreateFolderRequestTag) SetKey(v string) *CreateFolderRequestTag {
	s.Key = &v
	return s
}

func (s *CreateFolderRequestTag) SetValue(v string) *CreateFolderRequestTag {
	s.Value = &v
	return s
}

type CreateFolderResponseBody struct {
	// The information about the folder.
	Folder *CreateFolderResponseBodyFolder `json:"Folder,omitempty" xml:"Folder,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// C2CBCA30-C8DD-423E-B4AD-4FB694C9180C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateFolderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateFolderResponseBody) GoString() string {
	return s.String()
}

func (s *CreateFolderResponseBody) SetFolder(v *CreateFolderResponseBodyFolder) *CreateFolderResponseBody {
	s.Folder = v
	return s
}

func (s *CreateFolderResponseBody) SetRequestId(v string) *CreateFolderResponseBody {
	s.RequestId = &v
	return s
}

type CreateFolderResponseBodyFolder struct {
	// The time when the folder was created.
	//
	// example:
	//
	// 2019-02-19T09:34:50.757Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the folder.
	//
	// example:
	//
	// fd-u8B321****
	FolderId *string `json:"FolderId,omitempty" xml:"FolderId,omitempty"`
	// The name of the folder.
	//
	// example:
	//
	// rdFolder
	FolderName *string `json:"FolderName,omitempty" xml:"FolderName,omitempty"`
	// The ID of the parent folder.
	//
	// example:
	//
	// r-b1****
	ParentFolderId *string `json:"ParentFolderId,omitempty" xml:"ParentFolderId,omitempty"`
}

func (s CreateFolderResponseBodyFolder) String() string {
	return tea.Prettify(s)
}

func (s CreateFolderResponseBodyFolder) GoString() string {
	return s.String()
}

func (s *CreateFolderResponseBodyFolder) SetCreateTime(v string) *CreateFolderResponseBodyFolder {
	s.CreateTime = &v
	return s
}

func (s *CreateFolderResponseBodyFolder) SetFolderId(v string) *CreateFolderResponseBodyFolder {
	s.FolderId = &v
	return s
}

func (s *CreateFolderResponseBodyFolder) SetFolderName(v string) *CreateFolderResponseBodyFolder {
	s.FolderName = &v
	return s
}

func (s *CreateFolderResponseBodyFolder) SetParentFolderId(v string) *CreateFolderResponseBodyFolder {
	s.ParentFolderId = &v
	return s
}

type CreateFolderResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateFolderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateFolderResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateFolderResponse) GoString() string {
	return s.String()
}

func (s *CreateFolderResponse) SetHeaders(v map[string]*string) *CreateFolderResponse {
	s.Headers = v
	return s
}

func (s *CreateFolderResponse) SetStatusCode(v int32) *CreateFolderResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateFolderResponse) SetBody(v *CreateFolderResponseBody) *CreateFolderResponse {
	s.Body = v
	return s
}

type CreateResourceAccountRequest struct {
	// The prefix for the Alibaba Cloud account name of the member. If you leave this parameter empty, the system randomly generates a prefix.
	//
	// The prefix must be 2 to 37 characters in length.
	//
	// The prefix can contain letters, digits, and special characters but cannot contain consecutive special characters. The prefix must start with a letter or digit and end with a letter or digit. Valid special characters include underscores (`_`), periods (.), and hyphens (`-`).
	//
	// The complete Alibaba Cloud account name of a member in a resource directory is in the @.aliyunid.com format, such as `alice@rd-3G****.aliyunid.com`.
	//
	// Each name must be unique in the resource directory.
	//
	// example:
	//
	// alice
	AccountNamePrefix *string `json:"AccountNamePrefix,omitempty" xml:"AccountNamePrefix,omitempty"`
	// The display name of the member.
	//
	// The name must be 2 to 50 characters in length.
	//
	// The name can contain letters, digits, underscores (_), periods (.), hyphens (-), and spaces.
	//
	// The name must be unique in the resource directory.
	//
	// This parameter is required.
	//
	// example:
	//
	// Dev
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// 	- true: performs only a dry run. The system checks whether an identity type can be specified for the member. If the request does not pass the dry run, an error code is returned.
	//
	// 	- false (default): performs a dry run and performs the actual request.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ID of the parent folder.
	//
	// example:
	//
	// fd-r23M55****
	ParentFolderId *string `json:"ParentFolderId,omitempty" xml:"ParentFolderId,omitempty"`
	// The ID of the billing account. If you leave this parameter empty, the newly created member is used as its billing account.
	//
	// example:
	//
	// 12323344****
	PayerAccountId *string `json:"PayerAccountId,omitempty" xml:"PayerAccountId,omitempty"`
	// The identity type of the member. Valid values:
	//
	// 	- resell: The member is an account for a reseller. This is the default value. A relationship is automatically established between the member and the reseller. The management account of the resource directory must be used as the billing account of the member.
	//
	// 	- non_resell: The member is not an account for a reseller. The member is an account that is not associated with a reseller. You can directly use the account to purchase Alibaba Cloud resources. The member is used as its own billing account.
	//
	// > This parameter is available only for resellers at the international site (alibabacloud.com).
	//
	// example:
	//
	// resell
	ResellAccountType *string `json:"ResellAccountType,omitempty" xml:"ResellAccountType,omitempty"`
	// The tag of the member.
	Tag []*CreateResourceAccountRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s CreateResourceAccountRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateResourceAccountRequest) GoString() string {
	return s.String()
}

func (s *CreateResourceAccountRequest) SetAccountNamePrefix(v string) *CreateResourceAccountRequest {
	s.AccountNamePrefix = &v
	return s
}

func (s *CreateResourceAccountRequest) SetDisplayName(v string) *CreateResourceAccountRequest {
	s.DisplayName = &v
	return s
}

func (s *CreateResourceAccountRequest) SetDryRun(v bool) *CreateResourceAccountRequest {
	s.DryRun = &v
	return s
}

func (s *CreateResourceAccountRequest) SetParentFolderId(v string) *CreateResourceAccountRequest {
	s.ParentFolderId = &v
	return s
}

func (s *CreateResourceAccountRequest) SetPayerAccountId(v string) *CreateResourceAccountRequest {
	s.PayerAccountId = &v
	return s
}

func (s *CreateResourceAccountRequest) SetResellAccountType(v string) *CreateResourceAccountRequest {
	s.ResellAccountType = &v
	return s
}

func (s *CreateResourceAccountRequest) SetTag(v []*CreateResourceAccountRequestTag) *CreateResourceAccountRequest {
	s.Tag = v
	return s
}

type CreateResourceAccountRequestTag struct {
	// The key of the tag.
	//
	// example:
	//
	// k1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag.
	//
	// example:
	//
	// v1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateResourceAccountRequestTag) String() string {
	return tea.Prettify(s)
}

func (s CreateResourceAccountRequestTag) GoString() string {
	return s.String()
}

func (s *CreateResourceAccountRequestTag) SetKey(v string) *CreateResourceAccountRequestTag {
	s.Key = &v
	return s
}

func (s *CreateResourceAccountRequestTag) SetValue(v string) *CreateResourceAccountRequestTag {
	s.Value = &v
	return s
}

type CreateResourceAccountResponseBody struct {
	// The information of the member.
	Account *CreateResourceAccountResponseBodyAccount `json:"Account,omitempty" xml:"Account,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// B356A415-D860-43E5-865A-E2193D62BBD6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateResourceAccountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateResourceAccountResponseBody) GoString() string {
	return s.String()
}

func (s *CreateResourceAccountResponseBody) SetAccount(v *CreateResourceAccountResponseBodyAccount) *CreateResourceAccountResponseBody {
	s.Account = v
	return s
}

func (s *CreateResourceAccountResponseBody) SetRequestId(v string) *CreateResourceAccountResponseBody {
	s.RequestId = &v
	return s
}

type CreateResourceAccountResponseBodyAccount struct {
	// The Alibaba Cloud account ID of the member.
	//
	// example:
	//
	// 112730938585****
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The Alibaba Cloud account name of the member.
	//
	// example:
	//
	// alice@rd-3g****.aliyunid.com
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The display name of the member.
	//
	// example:
	//
	// Dev
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The ID of the folder.
	//
	// example:
	//
	// fd-r23M55****
	FolderId *string `json:"FolderId,omitempty" xml:"FolderId,omitempty"`
	// The way in which the member joins the resource directory. Valid values:
	//
	// 	- invited: The member is invited to join the resource directory.
	//
	// 	- created: The member is directly created in the resource directory.
	//
	// example:
	//
	// created
	JoinMethod *string `json:"JoinMethod,omitempty" xml:"JoinMethod,omitempty"`
	// The time when the member joined the resource directory. The time is displayed in UTC.
	//
	// example:
	//
	// 2020-12-31T03:37:39.456Z
	JoinTime *string `json:"JoinTime,omitempty" xml:"JoinTime,omitempty"`
	// The time when the member was modified. The time is displayed in UTC.
	//
	// example:
	//
	// 2020-12-31T03:37:39.456Z
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The ID of the resource directory.
	//
	// example:
	//
	// rd-3G****
	ResourceDirectoryId *string `json:"ResourceDirectoryId,omitempty" xml:"ResourceDirectoryId,omitempty"`
	// The status of the member. The value CreateSuccess indicates that the member is created.
	//
	// example:
	//
	// CreateSuccess
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The type of the member. The value ResourceAccount indicates that the member is a resource account.
	//
	// example:
	//
	// ResourceAccount
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateResourceAccountResponseBodyAccount) String() string {
	return tea.Prettify(s)
}

func (s CreateResourceAccountResponseBodyAccount) GoString() string {
	return s.String()
}

func (s *CreateResourceAccountResponseBodyAccount) SetAccountId(v string) *CreateResourceAccountResponseBodyAccount {
	s.AccountId = &v
	return s
}

func (s *CreateResourceAccountResponseBodyAccount) SetAccountName(v string) *CreateResourceAccountResponseBodyAccount {
	s.AccountName = &v
	return s
}

func (s *CreateResourceAccountResponseBodyAccount) SetDisplayName(v string) *CreateResourceAccountResponseBodyAccount {
	s.DisplayName = &v
	return s
}

func (s *CreateResourceAccountResponseBodyAccount) SetFolderId(v string) *CreateResourceAccountResponseBodyAccount {
	s.FolderId = &v
	return s
}

func (s *CreateResourceAccountResponseBodyAccount) SetJoinMethod(v string) *CreateResourceAccountResponseBodyAccount {
	s.JoinMethod = &v
	return s
}

func (s *CreateResourceAccountResponseBodyAccount) SetJoinTime(v string) *CreateResourceAccountResponseBodyAccount {
	s.JoinTime = &v
	return s
}

func (s *CreateResourceAccountResponseBodyAccount) SetModifyTime(v string) *CreateResourceAccountResponseBodyAccount {
	s.ModifyTime = &v
	return s
}

func (s *CreateResourceAccountResponseBodyAccount) SetResourceDirectoryId(v string) *CreateResourceAccountResponseBodyAccount {
	s.ResourceDirectoryId = &v
	return s
}

func (s *CreateResourceAccountResponseBodyAccount) SetStatus(v string) *CreateResourceAccountResponseBodyAccount {
	s.Status = &v
	return s
}

func (s *CreateResourceAccountResponseBodyAccount) SetType(v string) *CreateResourceAccountResponseBodyAccount {
	s.Type = &v
	return s
}

type CreateResourceAccountResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateResourceAccountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateResourceAccountResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateResourceAccountResponse) GoString() string {
	return s.String()
}

func (s *CreateResourceAccountResponse) SetHeaders(v map[string]*string) *CreateResourceAccountResponse {
	s.Headers = v
	return s
}

func (s *CreateResourceAccountResponse) SetStatusCode(v int32) *CreateResourceAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateResourceAccountResponse) SetBody(v *CreateResourceAccountResponseBody) *CreateResourceAccountResponse {
	s.Body = v
	return s
}

type DeclineHandshakeRequest struct {
	// The ID of the invitation.
	//
	// This parameter is required.
	//
	// example:
	//
	// h-ycm4rp****
	HandshakeId *string `json:"HandshakeId,omitempty" xml:"HandshakeId,omitempty"`
}

func (s DeclineHandshakeRequest) String() string {
	return tea.Prettify(s)
}

func (s DeclineHandshakeRequest) GoString() string {
	return s.String()
}

func (s *DeclineHandshakeRequest) SetHandshakeId(v string) *DeclineHandshakeRequest {
	s.HandshakeId = &v
	return s
}

type DeclineHandshakeResponseBody struct {
	// The information of the invitation.
	Handshake *DeclineHandshakeResponseBodyHandshake `json:"Handshake,omitempty" xml:"Handshake,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 9B34724D-54B0-4A51-B34D-4512372FE1BE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeclineHandshakeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeclineHandshakeResponseBody) GoString() string {
	return s.String()
}

func (s *DeclineHandshakeResponseBody) SetHandshake(v *DeclineHandshakeResponseBodyHandshake) *DeclineHandshakeResponseBody {
	s.Handshake = v
	return s
}

func (s *DeclineHandshakeResponseBody) SetRequestId(v string) *DeclineHandshakeResponseBody {
	s.RequestId = &v
	return s
}

type DeclineHandshakeResponseBodyHandshake struct {
	// The time when the invitation was created.
	//
	// example:
	//
	// 2018-08-10T09:55:41Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The time when the invitation expires.
	//
	// example:
	//
	// 2018-08-10T09:55:41Z
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The ID of the invitation.
	//
	// example:
	//
	// h-ycm4rp****
	HandshakeId *string `json:"HandshakeId,omitempty" xml:"HandshakeId,omitempty"`
	// The ID of the management account of the resource directory.
	//
	// example:
	//
	// 172841235500****
	MasterAccountId *string `json:"MasterAccountId,omitempty" xml:"MasterAccountId,omitempty"`
	// The name of the management account of the resource directory.
	//
	// example:
	//
	// Alice
	MasterAccountName *string `json:"MasterAccountName,omitempty" xml:"MasterAccountName,omitempty"`
	// The time when the invitation was modified.
	//
	// example:
	//
	// 2018-08-10T09:55:41Z
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The description of the invitation.
	//
	// example:
	//
	// Welcome
	Note *string `json:"Note,omitempty" xml:"Note,omitempty"`
	// The ID of the resource directory.
	//
	// example:
	//
	// rd-abcdef****
	ResourceDirectoryId *string `json:"ResourceDirectoryId,omitempty" xml:"ResourceDirectoryId,omitempty"`
	// The status of the invitation. Valid values:
	//
	// 	- Pending: The invitation is waiting for confirmation.
	//
	// 	- Accepted: The invitation is accepted.
	//
	// 	- Cancelled: The invitation is canceled.
	//
	// 	- Declined: The invitation is rejected.
	//
	// 	- Expired: The invitation expires.
	//
	// example:
	//
	// Declined
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID or logon email address of the invited account.
	//
	// example:
	//
	// someone@example.com
	TargetEntity *string `json:"TargetEntity,omitempty" xml:"TargetEntity,omitempty"`
	// The type of the invited account. Valid values:
	//
	// 	- Account: indicates the ID of the account.
	//
	// 	- Email: indicates the logon email address of the account.
	//
	// example:
	//
	// Email
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
}

func (s DeclineHandshakeResponseBodyHandshake) String() string {
	return tea.Prettify(s)
}

func (s DeclineHandshakeResponseBodyHandshake) GoString() string {
	return s.String()
}

func (s *DeclineHandshakeResponseBodyHandshake) SetCreateTime(v string) *DeclineHandshakeResponseBodyHandshake {
	s.CreateTime = &v
	return s
}

func (s *DeclineHandshakeResponseBodyHandshake) SetExpireTime(v string) *DeclineHandshakeResponseBodyHandshake {
	s.ExpireTime = &v
	return s
}

func (s *DeclineHandshakeResponseBodyHandshake) SetHandshakeId(v string) *DeclineHandshakeResponseBodyHandshake {
	s.HandshakeId = &v
	return s
}

func (s *DeclineHandshakeResponseBodyHandshake) SetMasterAccountId(v string) *DeclineHandshakeResponseBodyHandshake {
	s.MasterAccountId = &v
	return s
}

func (s *DeclineHandshakeResponseBodyHandshake) SetMasterAccountName(v string) *DeclineHandshakeResponseBodyHandshake {
	s.MasterAccountName = &v
	return s
}

func (s *DeclineHandshakeResponseBodyHandshake) SetModifyTime(v string) *DeclineHandshakeResponseBodyHandshake {
	s.ModifyTime = &v
	return s
}

func (s *DeclineHandshakeResponseBodyHandshake) SetNote(v string) *DeclineHandshakeResponseBodyHandshake {
	s.Note = &v
	return s
}

func (s *DeclineHandshakeResponseBodyHandshake) SetResourceDirectoryId(v string) *DeclineHandshakeResponseBodyHandshake {
	s.ResourceDirectoryId = &v
	return s
}

func (s *DeclineHandshakeResponseBodyHandshake) SetStatus(v string) *DeclineHandshakeResponseBodyHandshake {
	s.Status = &v
	return s
}

func (s *DeclineHandshakeResponseBodyHandshake) SetTargetEntity(v string) *DeclineHandshakeResponseBodyHandshake {
	s.TargetEntity = &v
	return s
}

func (s *DeclineHandshakeResponseBodyHandshake) SetTargetType(v string) *DeclineHandshakeResponseBodyHandshake {
	s.TargetType = &v
	return s
}

type DeclineHandshakeResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeclineHandshakeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeclineHandshakeResponse) String() string {
	return tea.Prettify(s)
}

func (s DeclineHandshakeResponse) GoString() string {
	return s.String()
}

func (s *DeclineHandshakeResponse) SetHeaders(v map[string]*string) *DeclineHandshakeResponse {
	s.Headers = v
	return s
}

func (s *DeclineHandshakeResponse) SetStatusCode(v int32) *DeclineHandshakeResponse {
	s.StatusCode = &v
	return s
}

func (s *DeclineHandshakeResponse) SetBody(v *DeclineHandshakeResponseBody) *DeclineHandshakeResponse {
	s.Body = v
	return s
}

type DeleteAccountRequest struct {
	// The IDs of the check items that you can choose to ignore for the member deletion.
	//
	// You can obtain the IDs from the response of the [GetAccountDeletionCheckResult](~~GetAccountDeletionCheckResult~~) operation.
	AbandonableCheckId []*string `json:"AbandonableCheckId,omitempty" xml:"AbandonableCheckId,omitempty" type:"Repeated"`
	// The Alibaba Cloud account ID of the member that you want to delete.
	//
	// This parameter is required.
	//
	// example:
	//
	// 169946124551****
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
}

func (s DeleteAccountRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteAccountRequest) GoString() string {
	return s.String()
}

func (s *DeleteAccountRequest) SetAbandonableCheckId(v []*string) *DeleteAccountRequest {
	s.AbandonableCheckId = v
	return s
}

func (s *DeleteAccountRequest) SetAccountId(v string) *DeleteAccountRequest {
	s.AccountId = &v
	return s
}

type DeleteAccountShrinkRequest struct {
	// The IDs of the check items that you can choose to ignore for the member deletion.
	//
	// You can obtain the IDs from the response of the [GetAccountDeletionCheckResult](~~GetAccountDeletionCheckResult~~) operation.
	AbandonableCheckIdShrink *string `json:"AbandonableCheckId,omitempty" xml:"AbandonableCheckId,omitempty"`
	// The Alibaba Cloud account ID of the member that you want to delete.
	//
	// This parameter is required.
	//
	// example:
	//
	// 169946124551****
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
}

func (s DeleteAccountShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteAccountShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeleteAccountShrinkRequest) SetAbandonableCheckIdShrink(v string) *DeleteAccountShrinkRequest {
	s.AbandonableCheckIdShrink = &v
	return s
}

func (s *DeleteAccountShrinkRequest) SetAccountId(v string) *DeleteAccountShrinkRequest {
	s.AccountId = &v
	return s
}

type DeleteAccountResponseBody struct {
	// The type of the deletion. Valid values:
	//
	// 	- 0: direct deletion. If the member does not have pay-as-you-go resources that are purchased within the previous 30 days, the system directly deletes the member.
	//
	// 	- 1: deletion with a silence period. If the member has pay-as-you-go resources that are purchased within the previous 30 days, the member enters a silence period. The system starts to delete the member until the silence period ends. For more information about the silence period, see [What is the silence period for member deletion?](https://help.aliyun.com/document_detail/446079.html)
	//
	// example:
	//
	// 0
	DeletionType *string `json:"DeletionType,omitempty" xml:"DeletionType,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 009429F8-C1C0-5872-B674-A6C2333B9647
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteAccountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteAccountResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAccountResponseBody) SetDeletionType(v string) *DeleteAccountResponseBody {
	s.DeletionType = &v
	return s
}

func (s *DeleteAccountResponseBody) SetRequestId(v string) *DeleteAccountResponseBody {
	s.RequestId = &v
	return s
}

type DeleteAccountResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAccountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAccountResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteAccountResponse) GoString() string {
	return s.String()
}

func (s *DeleteAccountResponse) SetHeaders(v map[string]*string) *DeleteAccountResponse {
	s.Headers = v
	return s
}

func (s *DeleteAccountResponse) SetStatusCode(v int32) *DeleteAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAccountResponse) SetBody(v *DeleteAccountResponseBody) *DeleteAccountResponse {
	s.Body = v
	return s
}

type DeleteControlPolicyRequest struct {
	// The ID of the access control policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// cp-SImPt8GCEwiq****
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
}

func (s DeleteControlPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteControlPolicyRequest) GoString() string {
	return s.String()
}

func (s *DeleteControlPolicyRequest) SetPolicyId(v string) *DeleteControlPolicyRequest {
	s.PolicyId = &v
	return s
}

type DeleteControlPolicyResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// C8541E06-B207-46BF-92C9-DC8DE4609D75
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteControlPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteControlPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteControlPolicyResponseBody) SetRequestId(v string) *DeleteControlPolicyResponseBody {
	s.RequestId = &v
	return s
}

type DeleteControlPolicyResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteControlPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteControlPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteControlPolicyResponse) GoString() string {
	return s.String()
}

func (s *DeleteControlPolicyResponse) SetHeaders(v map[string]*string) *DeleteControlPolicyResponse {
	s.Headers = v
	return s
}

func (s *DeleteControlPolicyResponse) SetStatusCode(v int32) *DeleteControlPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteControlPolicyResponse) SetBody(v *DeleteControlPolicyResponseBody) *DeleteControlPolicyResponse {
	s.Body = v
	return s
}

type DeleteFolderRequest struct {
	// The ID of the folder.
	//
	// This parameter is required.
	//
	// example:
	//
	// fd-ae1in7****
	FolderId *string `json:"FolderId,omitempty" xml:"FolderId,omitempty"`
}

func (s DeleteFolderRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteFolderRequest) GoString() string {
	return s.String()
}

func (s *DeleteFolderRequest) SetFolderId(v string) *DeleteFolderRequest {
	s.FolderId = &v
	return s
}

type DeleteFolderResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 9B34724D-54B0-4A51-B34D-4512372FE1BE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteFolderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteFolderResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteFolderResponseBody) SetRequestId(v string) *DeleteFolderResponseBody {
	s.RequestId = &v
	return s
}

type DeleteFolderResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteFolderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteFolderResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteFolderResponse) GoString() string {
	return s.String()
}

func (s *DeleteFolderResponse) SetHeaders(v map[string]*string) *DeleteFolderResponse {
	s.Headers = v
	return s
}

func (s *DeleteFolderResponse) SetStatusCode(v int32) *DeleteFolderResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteFolderResponse) SetBody(v *DeleteFolderResponseBody) *DeleteFolderResponse {
	s.Body = v
	return s
}

type DeleteMessageContactRequest struct {
	// The ID of the contact.
	//
	// This parameter is required.
	//
	// example:
	//
	// c-qL4HqKONzOM7****
	ContactId *string `json:"ContactId,omitempty" xml:"ContactId,omitempty"`
	// Specifies whether to retain the contact for members. Valid values:
	//
	// 	- true (default): retains the contact for members. In this case, the contact can still receive messages for the members.
	//
	// 	- false: does not retain the contact for members. In this case, the contact can no longer receive messages for the members. If you set this parameter to false, the response is asynchronously returned. You can call [GetMessageContactDeletionStatus](~~GetMessageContactDeletionStatus~~) to obtain the deletion result.
	//
	// example:
	//
	// true
	RetainContactInMembers *bool `json:"RetainContactInMembers,omitempty" xml:"RetainContactInMembers,omitempty"`
}

func (s DeleteMessageContactRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteMessageContactRequest) GoString() string {
	return s.String()
}

func (s *DeleteMessageContactRequest) SetContactId(v string) *DeleteMessageContactRequest {
	s.ContactId = &v
	return s
}

func (s *DeleteMessageContactRequest) SetRetainContactInMembers(v bool) *DeleteMessageContactRequest {
	s.RetainContactInMembers = &v
	return s
}

type DeleteMessageContactResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 9B34724D-54B0-4A51-B34D-4512372FE1BE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The deletion status of the contact. Valid values:
	//
	// 	- Deleting
	//
	// 	- Deleted
	//
	// example:
	//
	// Deleting
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DeleteMessageContactResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteMessageContactResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteMessageContactResponseBody) SetRequestId(v string) *DeleteMessageContactResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteMessageContactResponseBody) SetStatus(v string) *DeleteMessageContactResponseBody {
	s.Status = &v
	return s
}

type DeleteMessageContactResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteMessageContactResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteMessageContactResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteMessageContactResponse) GoString() string {
	return s.String()
}

func (s *DeleteMessageContactResponse) SetHeaders(v map[string]*string) *DeleteMessageContactResponse {
	s.Headers = v
	return s
}

func (s *DeleteMessageContactResponse) SetStatusCode(v int32) *DeleteMessageContactResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteMessageContactResponse) SetBody(v *DeleteMessageContactResponseBody) *DeleteMessageContactResponse {
	s.Body = v
	return s
}

type DeregisterDelegatedAdministratorRequest struct {
	// The Alibaba Cloud account ID of the member in the resource directory.
	//
	// This parameter is required.
	//
	// example:
	//
	// 181761095690****
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The identifier of the trusted service.
	//
	// This parameter is required.
	//
	// example:
	//
	// cloudfw.aliyuncs.com
	ServicePrincipal *string `json:"ServicePrincipal,omitempty" xml:"ServicePrincipal,omitempty"`
}

func (s DeregisterDelegatedAdministratorRequest) String() string {
	return tea.Prettify(s)
}

func (s DeregisterDelegatedAdministratorRequest) GoString() string {
	return s.String()
}

func (s *DeregisterDelegatedAdministratorRequest) SetAccountId(v string) *DeregisterDelegatedAdministratorRequest {
	s.AccountId = &v
	return s
}

func (s *DeregisterDelegatedAdministratorRequest) SetServicePrincipal(v string) *DeregisterDelegatedAdministratorRequest {
	s.ServicePrincipal = &v
	return s
}

type DeregisterDelegatedAdministratorResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// DF5D5C52-7BD0-40E7-94C6-23A1505038A2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeregisterDelegatedAdministratorResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeregisterDelegatedAdministratorResponseBody) GoString() string {
	return s.String()
}

func (s *DeregisterDelegatedAdministratorResponseBody) SetRequestId(v string) *DeregisterDelegatedAdministratorResponseBody {
	s.RequestId = &v
	return s
}

type DeregisterDelegatedAdministratorResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeregisterDelegatedAdministratorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeregisterDelegatedAdministratorResponse) String() string {
	return tea.Prettify(s)
}

func (s DeregisterDelegatedAdministratorResponse) GoString() string {
	return s.String()
}

func (s *DeregisterDelegatedAdministratorResponse) SetHeaders(v map[string]*string) *DeregisterDelegatedAdministratorResponse {
	s.Headers = v
	return s
}

func (s *DeregisterDelegatedAdministratorResponse) SetStatusCode(v int32) *DeregisterDelegatedAdministratorResponse {
	s.StatusCode = &v
	return s
}

func (s *DeregisterDelegatedAdministratorResponse) SetBody(v *DeregisterDelegatedAdministratorResponseBody) *DeregisterDelegatedAdministratorResponse {
	s.Body = v
	return s
}

type DestroyResourceDirectoryResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 9B34724D-54B0-4A51-B34D-4512372FE1BE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DestroyResourceDirectoryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DestroyResourceDirectoryResponseBody) GoString() string {
	return s.String()
}

func (s *DestroyResourceDirectoryResponseBody) SetRequestId(v string) *DestroyResourceDirectoryResponseBody {
	s.RequestId = &v
	return s
}

type DestroyResourceDirectoryResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DestroyResourceDirectoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DestroyResourceDirectoryResponse) String() string {
	return tea.Prettify(s)
}

func (s DestroyResourceDirectoryResponse) GoString() string {
	return s.String()
}

func (s *DestroyResourceDirectoryResponse) SetHeaders(v map[string]*string) *DestroyResourceDirectoryResponse {
	s.Headers = v
	return s
}

func (s *DestroyResourceDirectoryResponse) SetStatusCode(v int32) *DestroyResourceDirectoryResponse {
	s.StatusCode = &v
	return s
}

func (s *DestroyResourceDirectoryResponse) SetBody(v *DestroyResourceDirectoryResponseBody) *DestroyResourceDirectoryResponse {
	s.Body = v
	return s
}

type DetachControlPolicyRequest struct {
	// The ID of the access control policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// cp-jExXAqIYkwHN****
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// The ID of the object from which you want to detach the access control policy. Access control policies can be attached to the following objects:
	//
	// 	- Root folder
	//
	// 	- Subfolders of the Root folder
	//
	// 	- Members
	//
	// This parameter is required.
	//
	// example:
	//
	// fd-ZDNPiT****
	TargetId *string `json:"TargetId,omitempty" xml:"TargetId,omitempty"`
}

func (s DetachControlPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s DetachControlPolicyRequest) GoString() string {
	return s.String()
}

func (s *DetachControlPolicyRequest) SetPolicyId(v string) *DetachControlPolicyRequest {
	s.PolicyId = &v
	return s
}

func (s *DetachControlPolicyRequest) SetTargetId(v string) *DetachControlPolicyRequest {
	s.TargetId = &v
	return s
}

type DetachControlPolicyResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 9EA4F962-1A2E-4AFE-BE0C-B14736FC46CC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DetachControlPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DetachControlPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *DetachControlPolicyResponseBody) SetRequestId(v string) *DetachControlPolicyResponseBody {
	s.RequestId = &v
	return s
}

type DetachControlPolicyResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DetachControlPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DetachControlPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s DetachControlPolicyResponse) GoString() string {
	return s.String()
}

func (s *DetachControlPolicyResponse) SetHeaders(v map[string]*string) *DetachControlPolicyResponse {
	s.Headers = v
	return s
}

func (s *DetachControlPolicyResponse) SetStatusCode(v int32) *DetachControlPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *DetachControlPolicyResponse) SetBody(v *DetachControlPolicyResponseBody) *DetachControlPolicyResponse {
	s.Body = v
	return s
}

type DisableControlPolicyResponseBody struct {
	// The status of the Control Policy feature. Valid values:
	//
	// 	- Enabled: The feature is enabled.
	//
	// 	- PendingEnable: The feature is being enabled.
	//
	// 	- Disabled: The feature is disabled.
	//
	// 	- PendingDisable: The feature is being disabled.
	//
	// example:
	//
	// PendingDisable
	EnablementStatus *string `json:"EnablementStatus,omitempty" xml:"EnablementStatus,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 7C709979-451D-4C92-835D-7DDCCAA562E9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisableControlPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DisableControlPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *DisableControlPolicyResponseBody) SetEnablementStatus(v string) *DisableControlPolicyResponseBody {
	s.EnablementStatus = &v
	return s
}

func (s *DisableControlPolicyResponseBody) SetRequestId(v string) *DisableControlPolicyResponseBody {
	s.RequestId = &v
	return s
}

type DisableControlPolicyResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisableControlPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisableControlPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s DisableControlPolicyResponse) GoString() string {
	return s.String()
}

func (s *DisableControlPolicyResponse) SetHeaders(v map[string]*string) *DisableControlPolicyResponse {
	s.Headers = v
	return s
}

func (s *DisableControlPolicyResponse) SetStatusCode(v int32) *DisableControlPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableControlPolicyResponse) SetBody(v *DisableControlPolicyResponseBody) *DisableControlPolicyResponse {
	s.Body = v
	return s
}

type DisassociateMembersRequest struct {
	// The ID of the contact.
	//
	// This parameter is required.
	//
	// example:
	//
	// c-qL4HqKONzOM7****
	ContactId *string `json:"ContactId,omitempty" xml:"ContactId,omitempty"`
	// The IDs of objects from which you want to unbind the contact.
	//
	// This parameter is required.
	Members []*string `json:"Members,omitempty" xml:"Members,omitempty" type:"Repeated"`
}

func (s DisassociateMembersRequest) String() string {
	return tea.Prettify(s)
}

func (s DisassociateMembersRequest) GoString() string {
	return s.String()
}

func (s *DisassociateMembersRequest) SetContactId(v string) *DisassociateMembersRequest {
	s.ContactId = &v
	return s
}

func (s *DisassociateMembersRequest) SetMembers(v []*string) *DisassociateMembersRequest {
	s.Members = v
	return s
}

type DisassociateMembersResponseBody struct {
	// The time when the contact was unbound from the object.
	Members []*DisassociateMembersResponseBodyMembers `json:"Members,omitempty" xml:"Members,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 95060F1D-6990-4645-8920-A81D1BBFE992
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisassociateMembersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DisassociateMembersResponseBody) GoString() string {
	return s.String()
}

func (s *DisassociateMembersResponseBody) SetMembers(v []*DisassociateMembersResponseBodyMembers) *DisassociateMembersResponseBody {
	s.Members = v
	return s
}

func (s *DisassociateMembersResponseBody) SetRequestId(v string) *DisassociateMembersResponseBody {
	s.RequestId = &v
	return s
}

type DisassociateMembersResponseBodyMembers struct {
	// The ID of the contact.
	//
	// example:
	//
	// c-qL4HqKONzOM7****
	ContactId *string `json:"ContactId,omitempty" xml:"ContactId,omitempty"`
	// The ID of the object. Valid values:
	//
	// - ID of the resource directory
	//
	// - ID of the folder
	//
	// - ID of the member
	//
	// example:
	//
	// fd-ZDNPiT****
	MemberId *string `json:"MemberId,omitempty" xml:"MemberId,omitempty"`
	// The time when the contact was unbound from the object.
	//
	// example:
	//
	// 2023-03-27 17:19:21
	ModifyDate *string `json:"ModifyDate,omitempty" xml:"ModifyDate,omitempty"`
}

func (s DisassociateMembersResponseBodyMembers) String() string {
	return tea.Prettify(s)
}

func (s DisassociateMembersResponseBodyMembers) GoString() string {
	return s.String()
}

func (s *DisassociateMembersResponseBodyMembers) SetContactId(v string) *DisassociateMembersResponseBodyMembers {
	s.ContactId = &v
	return s
}

func (s *DisassociateMembersResponseBodyMembers) SetMemberId(v string) *DisassociateMembersResponseBodyMembers {
	s.MemberId = &v
	return s
}

func (s *DisassociateMembersResponseBodyMembers) SetModifyDate(v string) *DisassociateMembersResponseBodyMembers {
	s.ModifyDate = &v
	return s
}

type DisassociateMembersResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisassociateMembersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisassociateMembersResponse) String() string {
	return tea.Prettify(s)
}

func (s DisassociateMembersResponse) GoString() string {
	return s.String()
}

func (s *DisassociateMembersResponse) SetHeaders(v map[string]*string) *DisassociateMembersResponse {
	s.Headers = v
	return s
}

func (s *DisassociateMembersResponse) SetStatusCode(v int32) *DisassociateMembersResponse {
	s.StatusCode = &v
	return s
}

func (s *DisassociateMembersResponse) SetBody(v *DisassociateMembersResponseBody) *DisassociateMembersResponse {
	s.Body = v
	return s
}

type EnableControlPolicyResponseBody struct {
	// The status of the Control Policy feature. Valid values:
	//
	// 	- Enabled: The feature is enabled.
	//
	// 	- PendingEnable: The feature is being enabled.
	//
	// 	- Disabled: The feature is disabled.
	//
	// 	- PendingDisable: The feature is being disabled.
	//
	// example:
	//
	// PendingEnable
	EnablementStatus *string `json:"EnablementStatus,omitempty" xml:"EnablementStatus,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 8CE7BD95-EFFA-4911-A1E0-BD4412697FEB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnableControlPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EnableControlPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *EnableControlPolicyResponseBody) SetEnablementStatus(v string) *EnableControlPolicyResponseBody {
	s.EnablementStatus = &v
	return s
}

func (s *EnableControlPolicyResponseBody) SetRequestId(v string) *EnableControlPolicyResponseBody {
	s.RequestId = &v
	return s
}

type EnableControlPolicyResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *EnableControlPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableControlPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s EnableControlPolicyResponse) GoString() string {
	return s.String()
}

func (s *EnableControlPolicyResponse) SetHeaders(v map[string]*string) *EnableControlPolicyResponse {
	s.Headers = v
	return s
}

func (s *EnableControlPolicyResponse) SetStatusCode(v int32) *EnableControlPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *EnableControlPolicyResponse) SetBody(v *EnableControlPolicyResponseBody) *EnableControlPolicyResponse {
	s.Body = v
	return s
}

type EnableResourceDirectoryRequest struct {
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// 	- **true**: performs only a dry run.
	//
	// 	- **false*	- (default): performs a dry run and performs the actual request.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The mode in which you enable a resource directory. Valid values:
	//
	// 	- CurrentAccount: The current account is used to enable a resource directory.
	//
	// 	- NewManagementAccount: A newly created account is used to enable a resource directory. If you select this mode, you must configure the `MAName`, `MASecureMobilePhone`, and `VerificationCode` parameters.
	//
	// This parameter is required.
	//
	// example:
	//
	// CurrentAccount
	EnableMode *string `json:"EnableMode,omitempty" xml:"EnableMode,omitempty"`
	// The name of the newly created account.
	//
	// Specify the name in the `<Prefix>@rdadmin.aliyunid.com` format. The prefix can contain letters, digits, and special characters but cannot contain consecutive special characters. The prefix must start and end with a letter or digit. Valid special characters include underscores (`_`), periods (.), and hyphens (-). The prefix must be 2 to 50 characters in length.
	//
	// example:
	//
	// user01@rdadmin.aliyunid.com
	MAName *string `json:"MAName,omitempty" xml:"MAName,omitempty"`
	// The mobile phone number that is bound to the newly created account.
	//
	// If you leave this parameter empty, the mobile phone number that is bound to the current account is used. The mobile phone number you specify must be the same as the mobile phone number that you specify when you call the [SendVerificationCodeForEnableRD](~~SendVerificationCodeForEnableRD~~) operation to obtain a verification code.
	//
	// Specify the mobile phone number in the `<Country code>-<Mobile phone number>` format.
	//
	// > Mobile phone numbers in the `86-<Mobile phone number>` format in the Chinese mainland are not supported.
	//
	// example:
	//
	// xx-13900001234
	MASecureMobilePhone *string `json:"MASecureMobilePhone,omitempty" xml:"MASecureMobilePhone,omitempty"`
	// The verification code.
	//
	// You can call the [SendVerificationCodeForEnableRD](~~SendVerificationCodeForEnableRD~~) operation to obtain the verification code.
	//
	// example:
	//
	// 123456
	VerificationCode *string `json:"VerificationCode,omitempty" xml:"VerificationCode,omitempty"`
}

func (s EnableResourceDirectoryRequest) String() string {
	return tea.Prettify(s)
}

func (s EnableResourceDirectoryRequest) GoString() string {
	return s.String()
}

func (s *EnableResourceDirectoryRequest) SetDryRun(v bool) *EnableResourceDirectoryRequest {
	s.DryRun = &v
	return s
}

func (s *EnableResourceDirectoryRequest) SetEnableMode(v string) *EnableResourceDirectoryRequest {
	s.EnableMode = &v
	return s
}

func (s *EnableResourceDirectoryRequest) SetMAName(v string) *EnableResourceDirectoryRequest {
	s.MAName = &v
	return s
}

func (s *EnableResourceDirectoryRequest) SetMASecureMobilePhone(v string) *EnableResourceDirectoryRequest {
	s.MASecureMobilePhone = &v
	return s
}

func (s *EnableResourceDirectoryRequest) SetVerificationCode(v string) *EnableResourceDirectoryRequest {
	s.VerificationCode = &v
	return s
}

type EnableResourceDirectoryResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// EC2FE94D-A4A2-51A1-A493-5C273A36C46A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the resource directory.
	ResourceDirectory *EnableResourceDirectoryResponseBodyResourceDirectory `json:"ResourceDirectory,omitempty" xml:"ResourceDirectory,omitempty" type:"Struct"`
}

func (s EnableResourceDirectoryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EnableResourceDirectoryResponseBody) GoString() string {
	return s.String()
}

func (s *EnableResourceDirectoryResponseBody) SetRequestId(v string) *EnableResourceDirectoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *EnableResourceDirectoryResponseBody) SetResourceDirectory(v *EnableResourceDirectoryResponseBodyResourceDirectory) *EnableResourceDirectoryResponseBody {
	s.ResourceDirectory = v
	return s
}

type EnableResourceDirectoryResponseBodyResourceDirectory struct {
	// The time when the resource directory was enabled.
	//
	// example:
	//
	// 2021-12-08T02:15:31.744Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the management account.
	//
	// example:
	//
	// 507408460615****
	MasterAccountId *string `json:"MasterAccountId,omitempty" xml:"MasterAccountId,omitempty"`
	// The name of the management account.
	//
	// example:
	//
	// alice@example.com
	MasterAccountName *string `json:"MasterAccountName,omitempty" xml:"MasterAccountName,omitempty"`
	// The ID of the resource directory.
	//
	// example:
	//
	// rd-54****
	ResourceDirectoryId *string `json:"ResourceDirectoryId,omitempty" xml:"ResourceDirectoryId,omitempty"`
	// The ID of the Root folder.
	//
	// example:
	//
	// r-G9****
	RootFolderId *string `json:"RootFolderId,omitempty" xml:"RootFolderId,omitempty"`
}

func (s EnableResourceDirectoryResponseBodyResourceDirectory) String() string {
	return tea.Prettify(s)
}

func (s EnableResourceDirectoryResponseBodyResourceDirectory) GoString() string {
	return s.String()
}

func (s *EnableResourceDirectoryResponseBodyResourceDirectory) SetCreateTime(v string) *EnableResourceDirectoryResponseBodyResourceDirectory {
	s.CreateTime = &v
	return s
}

func (s *EnableResourceDirectoryResponseBodyResourceDirectory) SetMasterAccountId(v string) *EnableResourceDirectoryResponseBodyResourceDirectory {
	s.MasterAccountId = &v
	return s
}

func (s *EnableResourceDirectoryResponseBodyResourceDirectory) SetMasterAccountName(v string) *EnableResourceDirectoryResponseBodyResourceDirectory {
	s.MasterAccountName = &v
	return s
}

func (s *EnableResourceDirectoryResponseBodyResourceDirectory) SetResourceDirectoryId(v string) *EnableResourceDirectoryResponseBodyResourceDirectory {
	s.ResourceDirectoryId = &v
	return s
}

func (s *EnableResourceDirectoryResponseBodyResourceDirectory) SetRootFolderId(v string) *EnableResourceDirectoryResponseBodyResourceDirectory {
	s.RootFolderId = &v
	return s
}

type EnableResourceDirectoryResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *EnableResourceDirectoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableResourceDirectoryResponse) String() string {
	return tea.Prettify(s)
}

func (s EnableResourceDirectoryResponse) GoString() string {
	return s.String()
}

func (s *EnableResourceDirectoryResponse) SetHeaders(v map[string]*string) *EnableResourceDirectoryResponse {
	s.Headers = v
	return s
}

func (s *EnableResourceDirectoryResponse) SetStatusCode(v int32) *EnableResourceDirectoryResponse {
	s.StatusCode = &v
	return s
}

func (s *EnableResourceDirectoryResponse) SetBody(v *EnableResourceDirectoryResponseBody) *EnableResourceDirectoryResponse {
	s.Body = v
	return s
}

type GetAccountRequest struct {
	// The Alibaba Cloud account ID of the member.
	//
	// This parameter is required.
	//
	// example:
	//
	// 181761095690****
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// Specifies whether to return the information of tags. Valid values:
	//
	// 	- false (default value)
	//
	// 	- true
	//
	// example:
	//
	// true
	IncludeTags *bool `json:"IncludeTags,omitempty" xml:"IncludeTags,omitempty"`
}

func (s GetAccountRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAccountRequest) GoString() string {
	return s.String()
}

func (s *GetAccountRequest) SetAccountId(v string) *GetAccountRequest {
	s.AccountId = &v
	return s
}

func (s *GetAccountRequest) SetIncludeTags(v bool) *GetAccountRequest {
	s.IncludeTags = &v
	return s
}

type GetAccountResponseBody struct {
	// The information about the member.
	Account *GetAccountResponseBodyAccount `json:"Account,omitempty" xml:"Account,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 9B34724D-54B0-4A51-B34D-4512372FE1BE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAccountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAccountResponseBody) GoString() string {
	return s.String()
}

func (s *GetAccountResponseBody) SetAccount(v *GetAccountResponseBodyAccount) *GetAccountResponseBody {
	s.Account = v
	return s
}

func (s *GetAccountResponseBody) SetRequestId(v string) *GetAccountResponseBody {
	s.RequestId = &v
	return s
}

type GetAccountResponseBodyAccount struct {
	// The Alibaba Cloud account ID of the member.
	//
	// example:
	//
	// 181761095690****
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The Alibaba Cloud account name of the member.
	//
	// example:
	//
	// someone@example.com
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The display name of the member.
	//
	// example:
	//
	// admin
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The status of the modification for the email address bound to the member. Valid values:
	//
	// 	- If the value of this parameter is empty, no modification is performed for the email address.
	//
	// 	- WAIT_MODIFY: The modification is being performed.
	//
	// 	- CANCELLED: The modification is canceled.
	//
	// 	- EXPIRED: The modification expires.
	//
	// example:
	//
	// WAIT_MODIFY
	EmailStatus *string `json:"EmailStatus,omitempty" xml:"EmailStatus,omitempty"`
	// The ID of the folder.
	//
	// example:
	//
	// fd-bVaRIG****
	FolderId *string `json:"FolderId,omitempty" xml:"FolderId,omitempty"`
	// Indicates whether a mobile phone number is bound to the member for security purposes. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	HasSecureMobilePhone *bool `json:"HasSecureMobilePhone,omitempty" xml:"HasSecureMobilePhone,omitempty"`
	// The real-name verification information.
	//
	// example:
	//
	// aliyun-admin
	IdentityInformation *string `json:"IdentityInformation,omitempty" xml:"IdentityInformation,omitempty"`
	// The way in which the member joins the resource directory. Valid values:
	//
	// 	- invited: The member is invited to join the resource directory.
	//
	// 	- created: The member is directly created in the resource directory.
	//
	// example:
	//
	// created
	JoinMethod *string `json:"JoinMethod,omitempty" xml:"JoinMethod,omitempty"`
	// The time when the member joined the resource directory.
	//
	// example:
	//
	// 2015-01-23T12:33:18Z
	JoinTime *string `json:"JoinTime,omitempty" xml:"JoinTime,omitempty"`
	// The location of the member in the resource directory.
	Location *string `json:"Location,omitempty" xml:"Location,omitempty"`
	// The time when the member was modified.
	//
	// example:
	//
	// 2015-01-23T12:33:18Z
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The ID of the resource directory.
	//
	// example:
	//
	// rd-k3****
	ResourceDirectoryId *string `json:"ResourceDirectoryId,omitempty" xml:"ResourceDirectoryId,omitempty"`
	// The path of the member in the resource directory.
	ResourceDirectoryPath *string `json:"ResourceDirectoryPath,omitempty" xml:"ResourceDirectoryPath,omitempty"`
	// The status of the member. Valid values:
	//
	// 	- CreateSuccess: The member is created.
	//
	// 	- PromoteVerifying: The upgrade of the member is being confirmed.
	//
	// 	- PromoteFailed: The upgrade of the member fails.
	//
	// 	- PromoteExpired: The upgrade of the member expires.
	//
	// 	- PromoteCancelled: The upgrade of the member is canceled.
	//
	// 	- PromoteSuccess: The member is upgraded.
	//
	// 	- InviteSuccess: The member accepts the invitation.
	//
	// example:
	//
	// CreateSuccess
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tags of the member.
	Tags []*GetAccountResponseBodyAccountTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The type of the member. Valid values:
	//
	// 	- CloudAccount: cloud account
	//
	// 	- ResourceAccount: resource account
	//
	// example:
	//
	// ResourceAccount
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetAccountResponseBodyAccount) String() string {
	return tea.Prettify(s)
}

func (s GetAccountResponseBodyAccount) GoString() string {
	return s.String()
}

func (s *GetAccountResponseBodyAccount) SetAccountId(v string) *GetAccountResponseBodyAccount {
	s.AccountId = &v
	return s
}

func (s *GetAccountResponseBodyAccount) SetAccountName(v string) *GetAccountResponseBodyAccount {
	s.AccountName = &v
	return s
}

func (s *GetAccountResponseBodyAccount) SetDisplayName(v string) *GetAccountResponseBodyAccount {
	s.DisplayName = &v
	return s
}

func (s *GetAccountResponseBodyAccount) SetEmailStatus(v string) *GetAccountResponseBodyAccount {
	s.EmailStatus = &v
	return s
}

func (s *GetAccountResponseBodyAccount) SetFolderId(v string) *GetAccountResponseBodyAccount {
	s.FolderId = &v
	return s
}

func (s *GetAccountResponseBodyAccount) SetHasSecureMobilePhone(v bool) *GetAccountResponseBodyAccount {
	s.HasSecureMobilePhone = &v
	return s
}

func (s *GetAccountResponseBodyAccount) SetIdentityInformation(v string) *GetAccountResponseBodyAccount {
	s.IdentityInformation = &v
	return s
}

func (s *GetAccountResponseBodyAccount) SetJoinMethod(v string) *GetAccountResponseBodyAccount {
	s.JoinMethod = &v
	return s
}

func (s *GetAccountResponseBodyAccount) SetJoinTime(v string) *GetAccountResponseBodyAccount {
	s.JoinTime = &v
	return s
}

func (s *GetAccountResponseBodyAccount) SetLocation(v string) *GetAccountResponseBodyAccount {
	s.Location = &v
	return s
}

func (s *GetAccountResponseBodyAccount) SetModifyTime(v string) *GetAccountResponseBodyAccount {
	s.ModifyTime = &v
	return s
}

func (s *GetAccountResponseBodyAccount) SetResourceDirectoryId(v string) *GetAccountResponseBodyAccount {
	s.ResourceDirectoryId = &v
	return s
}

func (s *GetAccountResponseBodyAccount) SetResourceDirectoryPath(v string) *GetAccountResponseBodyAccount {
	s.ResourceDirectoryPath = &v
	return s
}

func (s *GetAccountResponseBodyAccount) SetStatus(v string) *GetAccountResponseBodyAccount {
	s.Status = &v
	return s
}

func (s *GetAccountResponseBodyAccount) SetTags(v []*GetAccountResponseBodyAccountTags) *GetAccountResponseBodyAccount {
	s.Tags = v
	return s
}

func (s *GetAccountResponseBodyAccount) SetType(v string) *GetAccountResponseBodyAccount {
	s.Type = &v
	return s
}

type GetAccountResponseBodyAccountTags struct {
	// The tag key.
	//
	// example:
	//
	// tag_key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// tag_value
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetAccountResponseBodyAccountTags) String() string {
	return tea.Prettify(s)
}

func (s GetAccountResponseBodyAccountTags) GoString() string {
	return s.String()
}

func (s *GetAccountResponseBodyAccountTags) SetKey(v string) *GetAccountResponseBodyAccountTags {
	s.Key = &v
	return s
}

func (s *GetAccountResponseBodyAccountTags) SetValue(v string) *GetAccountResponseBodyAccountTags {
	s.Value = &v
	return s
}

type GetAccountResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAccountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAccountResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAccountResponse) GoString() string {
	return s.String()
}

func (s *GetAccountResponse) SetHeaders(v map[string]*string) *GetAccountResponse {
	s.Headers = v
	return s
}

func (s *GetAccountResponse) SetStatusCode(v int32) *GetAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAccountResponse) SetBody(v *GetAccountResponseBody) *GetAccountResponse {
	s.Body = v
	return s
}

type GetAccountDeletionCheckResultRequest struct {
	// The Alibaba Cloud account ID of the member that you want to delete.
	//
	// This parameter is required.
	//
	// example:
	//
	// 179855839641****
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
}

func (s GetAccountDeletionCheckResultRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAccountDeletionCheckResultRequest) GoString() string {
	return s.String()
}

func (s *GetAccountDeletionCheckResultRequest) SetAccountId(v string) *GetAccountDeletionCheckResultRequest {
	s.AccountId = &v
	return s
}

type GetAccountDeletionCheckResultResponseBody struct {
	// The result of the deletion check for the member.
	AccountDeletionCheckResultInfo *GetAccountDeletionCheckResultResponseBodyAccountDeletionCheckResultInfo `json:"AccountDeletionCheckResultInfo,omitempty" xml:"AccountDeletionCheckResultInfo,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 54AC391D-4F7F-5F08-B8D3-0AECDE6EC5BD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAccountDeletionCheckResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAccountDeletionCheckResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetAccountDeletionCheckResultResponseBody) SetAccountDeletionCheckResultInfo(v *GetAccountDeletionCheckResultResponseBodyAccountDeletionCheckResultInfo) *GetAccountDeletionCheckResultResponseBody {
	s.AccountDeletionCheckResultInfo = v
	return s
}

func (s *GetAccountDeletionCheckResultResponseBody) SetRequestId(v string) *GetAccountDeletionCheckResultResponseBody {
	s.RequestId = &v
	return s
}

type GetAccountDeletionCheckResultResponseBodyAccountDeletionCheckResultInfo struct {
	// The check items that you can choose to ignore for the member deletion.
	//
	// > This parameter may be returned if the value of AllowDelete is true.
	AbandonableChecks []*GetAccountDeletionCheckResultResponseBodyAccountDeletionCheckResultInfoAbandonableChecks `json:"AbandonableChecks,omitempty" xml:"AbandonableChecks,omitempty" type:"Repeated"`
	// Indicates whether the member can be deleted. Valid values:
	//
	// 	- true: The member can be deleted.
	//
	// 	- false: The member cannot be deleted.
	//
	// example:
	//
	// false
	AllowDelete *string `json:"AllowDelete,omitempty" xml:"AllowDelete,omitempty"`
	// The reasons why the member cannot be deleted.
	//
	// > This parameter is returned only if the value of AllowDelete is false.
	NotAllowReason []*GetAccountDeletionCheckResultResponseBodyAccountDeletionCheckResultInfoNotAllowReason `json:"NotAllowReason,omitempty" xml:"NotAllowReason,omitempty" type:"Repeated"`
	// The status of the check. Valid values:
	//
	// 	- PreCheckComplete: The check is complete.
	//
	// 	- PreChecking: The check is in progress.
	//
	// example:
	//
	// PreCheckComplete
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetAccountDeletionCheckResultResponseBodyAccountDeletionCheckResultInfo) String() string {
	return tea.Prettify(s)
}

func (s GetAccountDeletionCheckResultResponseBodyAccountDeletionCheckResultInfo) GoString() string {
	return s.String()
}

func (s *GetAccountDeletionCheckResultResponseBodyAccountDeletionCheckResultInfo) SetAbandonableChecks(v []*GetAccountDeletionCheckResultResponseBodyAccountDeletionCheckResultInfoAbandonableChecks) *GetAccountDeletionCheckResultResponseBodyAccountDeletionCheckResultInfo {
	s.AbandonableChecks = v
	return s
}

func (s *GetAccountDeletionCheckResultResponseBodyAccountDeletionCheckResultInfo) SetAllowDelete(v string) *GetAccountDeletionCheckResultResponseBodyAccountDeletionCheckResultInfo {
	s.AllowDelete = &v
	return s
}

func (s *GetAccountDeletionCheckResultResponseBodyAccountDeletionCheckResultInfo) SetNotAllowReason(v []*GetAccountDeletionCheckResultResponseBodyAccountDeletionCheckResultInfoNotAllowReason) *GetAccountDeletionCheckResultResponseBodyAccountDeletionCheckResultInfo {
	s.NotAllowReason = v
	return s
}

func (s *GetAccountDeletionCheckResultResponseBodyAccountDeletionCheckResultInfo) SetStatus(v string) *GetAccountDeletionCheckResultResponseBodyAccountDeletionCheckResultInfo {
	s.Status = &v
	return s
}

type GetAccountDeletionCheckResultResponseBodyAccountDeletionCheckResultInfoAbandonableChecks struct {
	// The ID of the check item.
	//
	// example:
	//
	// NON_SP_cs
	CheckId *string `json:"CheckId,omitempty" xml:"CheckId,omitempty"`
	// The name of the cloud service to which the check item belongs.
	//
	// example:
	//
	// Container Service for Kubernetes
	CheckName *string `json:"CheckName,omitempty" xml:"CheckName,omitempty"`
	// The description of the check item.
	//
	// example:
	//
	// An instance of a cloud service is running within the member. Submit a ticket to contact Alibaba Cloud technical support.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
}

func (s GetAccountDeletionCheckResultResponseBodyAccountDeletionCheckResultInfoAbandonableChecks) String() string {
	return tea.Prettify(s)
}

func (s GetAccountDeletionCheckResultResponseBodyAccountDeletionCheckResultInfoAbandonableChecks) GoString() string {
	return s.String()
}

func (s *GetAccountDeletionCheckResultResponseBodyAccountDeletionCheckResultInfoAbandonableChecks) SetCheckId(v string) *GetAccountDeletionCheckResultResponseBodyAccountDeletionCheckResultInfoAbandonableChecks {
	s.CheckId = &v
	return s
}

func (s *GetAccountDeletionCheckResultResponseBodyAccountDeletionCheckResultInfoAbandonableChecks) SetCheckName(v string) *GetAccountDeletionCheckResultResponseBodyAccountDeletionCheckResultInfoAbandonableChecks {
	s.CheckName = &v
	return s
}

func (s *GetAccountDeletionCheckResultResponseBodyAccountDeletionCheckResultInfoAbandonableChecks) SetDescription(v string) *GetAccountDeletionCheckResultResponseBodyAccountDeletionCheckResultInfoAbandonableChecks {
	s.Description = &v
	return s
}

type GetAccountDeletionCheckResultResponseBodyAccountDeletionCheckResultInfoNotAllowReason struct {
	// The ID of the check item.
	//
	// example:
	//
	// NON_SP_efc
	CheckId *string `json:"CheckId,omitempty" xml:"CheckId,omitempty"`
	// The name of the cloud service to which the check item belongs.
	//
	// example:
	//
	// Enterprise finance
	CheckName *string `json:"CheckName,omitempty" xml:"CheckName,omitempty"`
	// The description of the check item.
	//
	// example:
	//
	// This account is an Enterprise Finance associated account. Please remove the financial association of this account before deleting it.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
}

func (s GetAccountDeletionCheckResultResponseBodyAccountDeletionCheckResultInfoNotAllowReason) String() string {
	return tea.Prettify(s)
}

func (s GetAccountDeletionCheckResultResponseBodyAccountDeletionCheckResultInfoNotAllowReason) GoString() string {
	return s.String()
}

func (s *GetAccountDeletionCheckResultResponseBodyAccountDeletionCheckResultInfoNotAllowReason) SetCheckId(v string) *GetAccountDeletionCheckResultResponseBodyAccountDeletionCheckResultInfoNotAllowReason {
	s.CheckId = &v
	return s
}

func (s *GetAccountDeletionCheckResultResponseBodyAccountDeletionCheckResultInfoNotAllowReason) SetCheckName(v string) *GetAccountDeletionCheckResultResponseBodyAccountDeletionCheckResultInfoNotAllowReason {
	s.CheckName = &v
	return s
}

func (s *GetAccountDeletionCheckResultResponseBodyAccountDeletionCheckResultInfoNotAllowReason) SetDescription(v string) *GetAccountDeletionCheckResultResponseBodyAccountDeletionCheckResultInfoNotAllowReason {
	s.Description = &v
	return s
}

type GetAccountDeletionCheckResultResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAccountDeletionCheckResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAccountDeletionCheckResultResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAccountDeletionCheckResultResponse) GoString() string {
	return s.String()
}

func (s *GetAccountDeletionCheckResultResponse) SetHeaders(v map[string]*string) *GetAccountDeletionCheckResultResponse {
	s.Headers = v
	return s
}

func (s *GetAccountDeletionCheckResultResponse) SetStatusCode(v int32) *GetAccountDeletionCheckResultResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAccountDeletionCheckResultResponse) SetBody(v *GetAccountDeletionCheckResultResponseBody) *GetAccountDeletionCheckResultResponse {
	s.Body = v
	return s
}

type GetAccountDeletionStatusRequest struct {
	// The Alibaba Cloud account ID of the member.
	//
	// This parameter is required.
	//
	// example:
	//
	// 169946124551****
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
}

func (s GetAccountDeletionStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAccountDeletionStatusRequest) GoString() string {
	return s.String()
}

func (s *GetAccountDeletionStatusRequest) SetAccountId(v string) *GetAccountDeletionStatusRequest {
	s.AccountId = &v
	return s
}

type GetAccountDeletionStatusResponseBody struct {
	// The deletion status of the member.
	RdAccountDeletionStatus *GetAccountDeletionStatusResponseBodyRdAccountDeletionStatus `json:"RdAccountDeletionStatus,omitempty" xml:"RdAccountDeletionStatus,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 8AA43293-7C8F-5730-8F2D-7F864EC092C5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAccountDeletionStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAccountDeletionStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetAccountDeletionStatusResponseBody) SetRdAccountDeletionStatus(v *GetAccountDeletionStatusResponseBodyRdAccountDeletionStatus) *GetAccountDeletionStatusResponseBody {
	s.RdAccountDeletionStatus = v
	return s
}

func (s *GetAccountDeletionStatusResponseBody) SetRequestId(v string) *GetAccountDeletionStatusResponseBody {
	s.RequestId = &v
	return s
}

type GetAccountDeletionStatusResponseBodyRdAccountDeletionStatus struct {
	// The Alibaba Cloud account ID of the member.
	//
	// example:
	//
	// 169946124551****
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The start time of the deletion.
	//
	// example:
	//
	// 2022-08-23T17:05:30+08:00
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The end time of the deletion.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2022-08-23T17:06:01+08:00
	DeletionTime *string `json:"DeletionTime,omitempty" xml:"DeletionTime,omitempty"`
	// The type of the deletion. Valid values:
	//
	// 	- 0: direct deletion. If the member does not have pay-as-you-go resources that are purchased within the previous 30 days, the system directly deletes the member.
	//
	// 	- 1: deletion with a silence period. If the member has pay-as-you-go resources that are purchased within the previous 30 days, the member enters a silence period of 45 days. The system starts to delete the member until the silence period ends. For more information about the silence period, see [What is the silence period for member deletion?](https://help.aliyun.com/document_detail/446079.html)
	//
	// example:
	//
	// 0
	DeletionType *string `json:"DeletionType,omitempty" xml:"DeletionType,omitempty"`
	// The reasons why the member fails to be deleted.
	FailReasonList []*GetAccountDeletionStatusResponseBodyRdAccountDeletionStatusFailReasonList `json:"FailReasonList,omitempty" xml:"FailReasonList,omitempty" type:"Repeated"`
	// The status. Valid values:
	//
	// 	- Success: The member is deleted.
	//
	// 	- Checking: A deletion check is being performed for the member.
	//
	// 	- Deleting: The member is being deleted.
	//
	// 	- CheckFailed: The deletion check for the member fails.
	//
	// 	- DeleteFailed: The member fails to be deleted.
	//
	// example:
	//
	// Success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetAccountDeletionStatusResponseBodyRdAccountDeletionStatus) String() string {
	return tea.Prettify(s)
}

func (s GetAccountDeletionStatusResponseBodyRdAccountDeletionStatus) GoString() string {
	return s.String()
}

func (s *GetAccountDeletionStatusResponseBodyRdAccountDeletionStatus) SetAccountId(v string) *GetAccountDeletionStatusResponseBodyRdAccountDeletionStatus {
	s.AccountId = &v
	return s
}

func (s *GetAccountDeletionStatusResponseBodyRdAccountDeletionStatus) SetCreateTime(v string) *GetAccountDeletionStatusResponseBodyRdAccountDeletionStatus {
	s.CreateTime = &v
	return s
}

func (s *GetAccountDeletionStatusResponseBodyRdAccountDeletionStatus) SetDeletionTime(v string) *GetAccountDeletionStatusResponseBodyRdAccountDeletionStatus {
	s.DeletionTime = &v
	return s
}

func (s *GetAccountDeletionStatusResponseBodyRdAccountDeletionStatus) SetDeletionType(v string) *GetAccountDeletionStatusResponseBodyRdAccountDeletionStatus {
	s.DeletionType = &v
	return s
}

func (s *GetAccountDeletionStatusResponseBodyRdAccountDeletionStatus) SetFailReasonList(v []*GetAccountDeletionStatusResponseBodyRdAccountDeletionStatusFailReasonList) *GetAccountDeletionStatusResponseBodyRdAccountDeletionStatus {
	s.FailReasonList = v
	return s
}

func (s *GetAccountDeletionStatusResponseBodyRdAccountDeletionStatus) SetStatus(v string) *GetAccountDeletionStatusResponseBodyRdAccountDeletionStatus {
	s.Status = &v
	return s
}

type GetAccountDeletionStatusResponseBodyRdAccountDeletionStatusFailReasonList struct {
	// The description of the check item.
	//
	// example:
	//
	// This account has a payer account. Please release the financial relationship of this account first.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the cloud service to which the check item belongs.
	//
	// example:
	//
	// Others
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetAccountDeletionStatusResponseBodyRdAccountDeletionStatusFailReasonList) String() string {
	return tea.Prettify(s)
}

func (s GetAccountDeletionStatusResponseBodyRdAccountDeletionStatusFailReasonList) GoString() string {
	return s.String()
}

func (s *GetAccountDeletionStatusResponseBodyRdAccountDeletionStatusFailReasonList) SetDescription(v string) *GetAccountDeletionStatusResponseBodyRdAccountDeletionStatusFailReasonList {
	s.Description = &v
	return s
}

func (s *GetAccountDeletionStatusResponseBodyRdAccountDeletionStatusFailReasonList) SetName(v string) *GetAccountDeletionStatusResponseBodyRdAccountDeletionStatusFailReasonList {
	s.Name = &v
	return s
}

type GetAccountDeletionStatusResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAccountDeletionStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAccountDeletionStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAccountDeletionStatusResponse) GoString() string {
	return s.String()
}

func (s *GetAccountDeletionStatusResponse) SetHeaders(v map[string]*string) *GetAccountDeletionStatusResponse {
	s.Headers = v
	return s
}

func (s *GetAccountDeletionStatusResponse) SetStatusCode(v int32) *GetAccountDeletionStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAccountDeletionStatusResponse) SetBody(v *GetAccountDeletionStatusResponseBody) *GetAccountDeletionStatusResponse {
	s.Body = v
	return s
}

type GetControlPolicyRequest struct {
	// The language in which you want to return the description of the access control policy. Valid values:
	//
	// 	- zh-CN (default value): Chinese
	//
	// 	- en: English
	//
	// 	- ja: Japanese
	//
	// > This parameter is valid only for system access control policies.
	//
	// example:
	//
	// zh-CN
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// The ID of the access control policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// cp-SImPt8GCEwiq****
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
}

func (s GetControlPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s GetControlPolicyRequest) GoString() string {
	return s.String()
}

func (s *GetControlPolicyRequest) SetLanguage(v string) *GetControlPolicyRequest {
	s.Language = &v
	return s
}

func (s *GetControlPolicyRequest) SetPolicyId(v string) *GetControlPolicyRequest {
	s.PolicyId = &v
	return s
}

type GetControlPolicyResponseBody struct {
	// The details of the access control policy.
	ControlPolicy *GetControlPolicyResponseBodyControlPolicy `json:"ControlPolicy,omitempty" xml:"ControlPolicy,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// AB769936-CDFA-4D52-8CE2-A3581800044A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetControlPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetControlPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *GetControlPolicyResponseBody) SetControlPolicy(v *GetControlPolicyResponseBodyControlPolicy) *GetControlPolicyResponseBody {
	s.ControlPolicy = v
	return s
}

func (s *GetControlPolicyResponseBody) SetRequestId(v string) *GetControlPolicyResponseBody {
	s.RequestId = &v
	return s
}

type GetControlPolicyResponseBodyControlPolicy struct {
	// The number of times that the access control policy is referenced.
	//
	// example:
	//
	// 0
	AttachmentCount *string `json:"AttachmentCount,omitempty" xml:"AttachmentCount,omitempty"`
	// The time when the access control policy was created.
	//
	// example:
	//
	// 2021-03-18T08:51:33Z
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	// The description of the access control policy.
	//
	// example:
	//
	// ExampleControlPolicy
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The effective scope of the access control policy. Valid values:
	//
	// 	- All: The access control policy is in effect for Alibaba Cloud accounts, RAM users, and RAM roles.
	//
	// 	- RAM: The access control policy is in effect only for RAM users and RAM roles.
	//
	// example:
	//
	// RAM
	EffectScope *string `json:"EffectScope,omitempty" xml:"EffectScope,omitempty"`
	// The document of the access control policy.
	//
	// example:
	//
	// {\\"Version\\":\\"1\\",\\"Statement\\":[{\\"Effect\\":\\"Deny\\",\\"Action\\":[\\"ram:UpdateRole\\",\\"ram:DeleteRole\\",\\"ram:AttachPolicyToRole\\",\\"ram:DetachPolicyFromRole\\"],\\"Resource\\":\\"acs:ram:*:*:role/ResourceDirectoryAccountAccessRole\\"}]}
	PolicyDocument *string `json:"PolicyDocument,omitempty" xml:"PolicyDocument,omitempty"`
	// The ID of the access control policy.
	//
	// example:
	//
	// cp-SImPt8GCEwiq****
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// The name of the access control policy.
	//
	// example:
	//
	// test
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	// The type of the access control policy. Valid values:
	//
	// 	- System: system access control policy
	//
	// 	- Custom: custom access control policy
	//
	// example:
	//
	// Custom
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
	// The time when the access control policy was updated.
	//
	// example:
	//
	// 2021-03-18T08:51:33Z
	UpdateDate *string `json:"UpdateDate,omitempty" xml:"UpdateDate,omitempty"`
}

func (s GetControlPolicyResponseBodyControlPolicy) String() string {
	return tea.Prettify(s)
}

func (s GetControlPolicyResponseBodyControlPolicy) GoString() string {
	return s.String()
}

func (s *GetControlPolicyResponseBodyControlPolicy) SetAttachmentCount(v string) *GetControlPolicyResponseBodyControlPolicy {
	s.AttachmentCount = &v
	return s
}

func (s *GetControlPolicyResponseBodyControlPolicy) SetCreateDate(v string) *GetControlPolicyResponseBodyControlPolicy {
	s.CreateDate = &v
	return s
}

func (s *GetControlPolicyResponseBodyControlPolicy) SetDescription(v string) *GetControlPolicyResponseBodyControlPolicy {
	s.Description = &v
	return s
}

func (s *GetControlPolicyResponseBodyControlPolicy) SetEffectScope(v string) *GetControlPolicyResponseBodyControlPolicy {
	s.EffectScope = &v
	return s
}

func (s *GetControlPolicyResponseBodyControlPolicy) SetPolicyDocument(v string) *GetControlPolicyResponseBodyControlPolicy {
	s.PolicyDocument = &v
	return s
}

func (s *GetControlPolicyResponseBodyControlPolicy) SetPolicyId(v string) *GetControlPolicyResponseBodyControlPolicy {
	s.PolicyId = &v
	return s
}

func (s *GetControlPolicyResponseBodyControlPolicy) SetPolicyName(v string) *GetControlPolicyResponseBodyControlPolicy {
	s.PolicyName = &v
	return s
}

func (s *GetControlPolicyResponseBodyControlPolicy) SetPolicyType(v string) *GetControlPolicyResponseBodyControlPolicy {
	s.PolicyType = &v
	return s
}

func (s *GetControlPolicyResponseBodyControlPolicy) SetUpdateDate(v string) *GetControlPolicyResponseBodyControlPolicy {
	s.UpdateDate = &v
	return s
}

type GetControlPolicyResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetControlPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetControlPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s GetControlPolicyResponse) GoString() string {
	return s.String()
}

func (s *GetControlPolicyResponse) SetHeaders(v map[string]*string) *GetControlPolicyResponse {
	s.Headers = v
	return s
}

func (s *GetControlPolicyResponse) SetStatusCode(v int32) *GetControlPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *GetControlPolicyResponse) SetBody(v *GetControlPolicyResponseBody) *GetControlPolicyResponse {
	s.Body = v
	return s
}

type GetControlPolicyEnablementStatusResponseBody struct {
	// The status of the Control Policy feature. Valid values:
	//
	// 	- Enabled: The feature is enabled.
	//
	// 	- PendingEnable: The feature is being enabled.
	//
	// 	- Disabled: The feature is disabled.
	//
	// 	- PendingDisable: The feature is being disabled.
	//
	// example:
	//
	// Disabled
	EnablementStatus *string `json:"EnablementStatus,omitempty" xml:"EnablementStatus,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 1DC39A4E-3B52-4EFE-9F93-4897D7FFA0C4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetControlPolicyEnablementStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetControlPolicyEnablementStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetControlPolicyEnablementStatusResponseBody) SetEnablementStatus(v string) *GetControlPolicyEnablementStatusResponseBody {
	s.EnablementStatus = &v
	return s
}

func (s *GetControlPolicyEnablementStatusResponseBody) SetRequestId(v string) *GetControlPolicyEnablementStatusResponseBody {
	s.RequestId = &v
	return s
}

type GetControlPolicyEnablementStatusResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetControlPolicyEnablementStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetControlPolicyEnablementStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s GetControlPolicyEnablementStatusResponse) GoString() string {
	return s.String()
}

func (s *GetControlPolicyEnablementStatusResponse) SetHeaders(v map[string]*string) *GetControlPolicyEnablementStatusResponse {
	s.Headers = v
	return s
}

func (s *GetControlPolicyEnablementStatusResponse) SetStatusCode(v int32) *GetControlPolicyEnablementStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetControlPolicyEnablementStatusResponse) SetBody(v *GetControlPolicyEnablementStatusResponseBody) *GetControlPolicyEnablementStatusResponse {
	s.Body = v
	return s
}

type GetFolderRequest struct {
	// The ID of the folder.
	//
	// This parameter is required.
	//
	// example:
	//
	// fd-Jyl5U7****
	FolderId *string `json:"FolderId,omitempty" xml:"FolderId,omitempty"`
}

func (s GetFolderRequest) String() string {
	return tea.Prettify(s)
}

func (s GetFolderRequest) GoString() string {
	return s.String()
}

func (s *GetFolderRequest) SetFolderId(v string) *GetFolderRequest {
	s.FolderId = &v
	return s
}

type GetFolderResponseBody struct {
	// The information about the folder.
	Folder *GetFolderResponseBodyFolder `json:"Folder,omitempty" xml:"Folder,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// C2CBCA30-C8DD-423E-B4AD-4FB694C9180C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetFolderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetFolderResponseBody) GoString() string {
	return s.String()
}

func (s *GetFolderResponseBody) SetFolder(v *GetFolderResponseBodyFolder) *GetFolderResponseBody {
	s.Folder = v
	return s
}

func (s *GetFolderResponseBody) SetRequestId(v string) *GetFolderResponseBody {
	s.RequestId = &v
	return s
}

type GetFolderResponseBodyFolder struct {
	// The time when the folder was created.
	//
	// example:
	//
	// 2021-06-15T06:39:08.521Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the folder.
	//
	// example:
	//
	// fd-Jyl5U7****
	FolderId *string `json:"FolderId,omitempty" xml:"FolderId,omitempty"`
	// The name of the folder.
	//
	// example:
	//
	// Applications
	FolderName *string `json:"FolderName,omitempty" xml:"FolderName,omitempty"`
	// The ID of the parent folder.
	//
	// example:
	//
	// r-Wm****
	ParentFolderId *string `json:"ParentFolderId,omitempty" xml:"ParentFolderId,omitempty"`
	// The path of the folder in the resource directory.
	ResourceDirectoryPath *string `json:"ResourceDirectoryPath,omitempty" xml:"ResourceDirectoryPath,omitempty"`
}

func (s GetFolderResponseBodyFolder) String() string {
	return tea.Prettify(s)
}

func (s GetFolderResponseBodyFolder) GoString() string {
	return s.String()
}

func (s *GetFolderResponseBodyFolder) SetCreateTime(v string) *GetFolderResponseBodyFolder {
	s.CreateTime = &v
	return s
}

func (s *GetFolderResponseBodyFolder) SetFolderId(v string) *GetFolderResponseBodyFolder {
	s.FolderId = &v
	return s
}

func (s *GetFolderResponseBodyFolder) SetFolderName(v string) *GetFolderResponseBodyFolder {
	s.FolderName = &v
	return s
}

func (s *GetFolderResponseBodyFolder) SetParentFolderId(v string) *GetFolderResponseBodyFolder {
	s.ParentFolderId = &v
	return s
}

func (s *GetFolderResponseBodyFolder) SetResourceDirectoryPath(v string) *GetFolderResponseBodyFolder {
	s.ResourceDirectoryPath = &v
	return s
}

type GetFolderResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetFolderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetFolderResponse) String() string {
	return tea.Prettify(s)
}

func (s GetFolderResponse) GoString() string {
	return s.String()
}

func (s *GetFolderResponse) SetHeaders(v map[string]*string) *GetFolderResponse {
	s.Headers = v
	return s
}

func (s *GetFolderResponse) SetStatusCode(v int32) *GetFolderResponse {
	s.StatusCode = &v
	return s
}

func (s *GetFolderResponse) SetBody(v *GetFolderResponseBody) *GetFolderResponse {
	s.Body = v
	return s
}

type GetHandshakeRequest struct {
	// The ID of the invitation.
	//
	// This parameter is required.
	//
	// example:
	//
	// h-ycm4rp****
	HandshakeId *string `json:"HandshakeId,omitempty" xml:"HandshakeId,omitempty"`
}

func (s GetHandshakeRequest) String() string {
	return tea.Prettify(s)
}

func (s GetHandshakeRequest) GoString() string {
	return s.String()
}

func (s *GetHandshakeRequest) SetHandshakeId(v string) *GetHandshakeRequest {
	s.HandshakeId = &v
	return s
}

type GetHandshakeResponseBody struct {
	// The information of the invitation.
	Handshake *GetHandshakeResponseBodyHandshake `json:"Handshake,omitempty" xml:"Handshake,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 9B34724D-54B0-4A51-B34D-4512372FE1BE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetHandshakeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetHandshakeResponseBody) GoString() string {
	return s.String()
}

func (s *GetHandshakeResponseBody) SetHandshake(v *GetHandshakeResponseBodyHandshake) *GetHandshakeResponseBody {
	s.Handshake = v
	return s
}

func (s *GetHandshakeResponseBody) SetRequestId(v string) *GetHandshakeResponseBody {
	s.RequestId = &v
	return s
}

type GetHandshakeResponseBodyHandshake struct {
	// The time when the invitation was created. The time is displayed in UTC.
	//
	// example:
	//
	// 2018-08-10T09:55:41Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The time when the invitation expires. The time is displayed in UTC.
	//
	// example:
	//
	// 2018-08-24T09:55:41Z
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The ID of the invitation.
	//
	// example:
	//
	// h-ycm4rp****
	HandshakeId *string `json:"HandshakeId,omitempty" xml:"HandshakeId,omitempty"`
	// The real-name verification information of the invitee.
	//
	// > This parameter is available only when an invitee calls this operation.
	//
	// example:
	//
	// Alice
	InvitedAccountRealName *string `json:"InvitedAccountRealName,omitempty" xml:"InvitedAccountRealName,omitempty"`
	// The ID of the management account of the resource directory.
	//
	// example:
	//
	// 172841235500****
	MasterAccountId *string `json:"MasterAccountId,omitempty" xml:"MasterAccountId,omitempty"`
	// The name of the management account of the resource directory.
	//
	// example:
	//
	// company@example.com
	MasterAccountName *string `json:"MasterAccountName,omitempty" xml:"MasterAccountName,omitempty"`
	// The real-name verification information of the management account of the resource directory.
	//
	// > This parameter is available only when an invitee calls this operation.
	//
	// example:
	//
	// company
	MasterAccountRealName *string `json:"MasterAccountRealName,omitempty" xml:"MasterAccountRealName,omitempty"`
	// The time when the invitation was modified. The time is displayed in UTC.
	//
	// example:
	//
	// 2018-08-10T09:55:41Z
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The description of the invitation.
	//
	// example:
	//
	// Welcome
	Note *string `json:"Note,omitempty" xml:"Note,omitempty"`
	// The ID of the resource directory.
	//
	// example:
	//
	// rd-abcdef****
	ResourceDirectoryId *string `json:"ResourceDirectoryId,omitempty" xml:"ResourceDirectoryId,omitempty"`
	// The status of the invitation. Valid values:
	//
	// 	- Pending: The invitation is waiting for confirmation.
	//
	// 	- Accepted: The invitation is accepted.
	//
	// 	- Cancelled: The invitation is canceled.
	//
	// 	- Declined: The invitation is rejected.
	//
	// 	- Expired: The invitation expires.
	//
	// example:
	//
	// Pending
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID or logon email address of the invited account.
	//
	// example:
	//
	// someone@example.com
	TargetEntity *string `json:"TargetEntity,omitempty" xml:"TargetEntity,omitempty"`
	// The type of the invited account. Valid values:
	//
	// 	- Account: indicates the ID of the account.
	//
	// 	- Email: indicates the logon email address of the account.
	//
	// example:
	//
	// Email
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
}

func (s GetHandshakeResponseBodyHandshake) String() string {
	return tea.Prettify(s)
}

func (s GetHandshakeResponseBodyHandshake) GoString() string {
	return s.String()
}

func (s *GetHandshakeResponseBodyHandshake) SetCreateTime(v string) *GetHandshakeResponseBodyHandshake {
	s.CreateTime = &v
	return s
}

func (s *GetHandshakeResponseBodyHandshake) SetExpireTime(v string) *GetHandshakeResponseBodyHandshake {
	s.ExpireTime = &v
	return s
}

func (s *GetHandshakeResponseBodyHandshake) SetHandshakeId(v string) *GetHandshakeResponseBodyHandshake {
	s.HandshakeId = &v
	return s
}

func (s *GetHandshakeResponseBodyHandshake) SetInvitedAccountRealName(v string) *GetHandshakeResponseBodyHandshake {
	s.InvitedAccountRealName = &v
	return s
}

func (s *GetHandshakeResponseBodyHandshake) SetMasterAccountId(v string) *GetHandshakeResponseBodyHandshake {
	s.MasterAccountId = &v
	return s
}

func (s *GetHandshakeResponseBodyHandshake) SetMasterAccountName(v string) *GetHandshakeResponseBodyHandshake {
	s.MasterAccountName = &v
	return s
}

func (s *GetHandshakeResponseBodyHandshake) SetMasterAccountRealName(v string) *GetHandshakeResponseBodyHandshake {
	s.MasterAccountRealName = &v
	return s
}

func (s *GetHandshakeResponseBodyHandshake) SetModifyTime(v string) *GetHandshakeResponseBodyHandshake {
	s.ModifyTime = &v
	return s
}

func (s *GetHandshakeResponseBodyHandshake) SetNote(v string) *GetHandshakeResponseBodyHandshake {
	s.Note = &v
	return s
}

func (s *GetHandshakeResponseBodyHandshake) SetResourceDirectoryId(v string) *GetHandshakeResponseBodyHandshake {
	s.ResourceDirectoryId = &v
	return s
}

func (s *GetHandshakeResponseBodyHandshake) SetStatus(v string) *GetHandshakeResponseBodyHandshake {
	s.Status = &v
	return s
}

func (s *GetHandshakeResponseBodyHandshake) SetTargetEntity(v string) *GetHandshakeResponseBodyHandshake {
	s.TargetEntity = &v
	return s
}

func (s *GetHandshakeResponseBodyHandshake) SetTargetType(v string) *GetHandshakeResponseBodyHandshake {
	s.TargetType = &v
	return s
}

type GetHandshakeResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetHandshakeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetHandshakeResponse) String() string {
	return tea.Prettify(s)
}

func (s GetHandshakeResponse) GoString() string {
	return s.String()
}

func (s *GetHandshakeResponse) SetHeaders(v map[string]*string) *GetHandshakeResponse {
	s.Headers = v
	return s
}

func (s *GetHandshakeResponse) SetStatusCode(v int32) *GetHandshakeResponse {
	s.StatusCode = &v
	return s
}

func (s *GetHandshakeResponse) SetBody(v *GetHandshakeResponseBody) *GetHandshakeResponse {
	s.Body = v
	return s
}

type GetMessageContactRequest struct {
	// The ID of the contact.
	//
	// This parameter is required.
	//
	// example:
	//
	// c-qL4HqKONzOM7****
	ContactId *string `json:"ContactId,omitempty" xml:"ContactId,omitempty"`
}

func (s GetMessageContactRequest) String() string {
	return tea.Prettify(s)
}

func (s GetMessageContactRequest) GoString() string {
	return s.String()
}

func (s *GetMessageContactRequest) SetContactId(v string) *GetMessageContactRequest {
	s.ContactId = &v
	return s
}

type GetMessageContactResponseBody struct {
	// The information about the contact.
	Contact *GetMessageContactResponseBodyContact `json:"Contact,omitempty" xml:"Contact,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 9B34724D-54B0-4A51-B34D-4512372FE1BE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetMessageContactResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetMessageContactResponseBody) GoString() string {
	return s.String()
}

func (s *GetMessageContactResponseBody) SetContact(v *GetMessageContactResponseBodyContact) *GetMessageContactResponseBody {
	s.Contact = v
	return s
}

func (s *GetMessageContactResponseBody) SetRequestId(v string) *GetMessageContactResponseBody {
	s.RequestId = &v
	return s
}

type GetMessageContactResponseBodyContact struct {
	// The ID of the contact.
	//
	// example:
	//
	// c-qL4HqKONzOM7****
	ContactId *string `json:"ContactId,omitempty" xml:"ContactId,omitempty"`
	// The time when the contact was created.
	//
	// example:
	//
	// 2023-03-27 17:19:21
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	// The email address of the contact.
	//
	// example:
	//
	// someone***@example.com
	EmailAddress *string `json:"EmailAddress,omitempty" xml:"EmailAddress,omitempty"`
	// The IDs of objects to which the contact is bound.
	Members []*string `json:"Members,omitempty" xml:"Members,omitempty" type:"Repeated"`
	// The types of messages received by the contact.
	MessageTypes []*string `json:"MessageTypes,omitempty" xml:"MessageTypes,omitempty" type:"Repeated"`
	// The name of the contact.
	//
	// example:
	//
	// tom
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The mobile phone number of the contact.
	//
	// example:
	//
	// 86-139****1234
	PhoneNumber *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	// The status of the contact. Valid values:
	//
	// 	- Verifying
	//
	// 	- Active
	//
	// 	- Deleting
	//
	// example:
	//
	// Active
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The job title of the contact.
	//
	// example:
	//
	// TechnicalDirector
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s GetMessageContactResponseBodyContact) String() string {
	return tea.Prettify(s)
}

func (s GetMessageContactResponseBodyContact) GoString() string {
	return s.String()
}

func (s *GetMessageContactResponseBodyContact) SetContactId(v string) *GetMessageContactResponseBodyContact {
	s.ContactId = &v
	return s
}

func (s *GetMessageContactResponseBodyContact) SetCreateDate(v string) *GetMessageContactResponseBodyContact {
	s.CreateDate = &v
	return s
}

func (s *GetMessageContactResponseBodyContact) SetEmailAddress(v string) *GetMessageContactResponseBodyContact {
	s.EmailAddress = &v
	return s
}

func (s *GetMessageContactResponseBodyContact) SetMembers(v []*string) *GetMessageContactResponseBodyContact {
	s.Members = v
	return s
}

func (s *GetMessageContactResponseBodyContact) SetMessageTypes(v []*string) *GetMessageContactResponseBodyContact {
	s.MessageTypes = v
	return s
}

func (s *GetMessageContactResponseBodyContact) SetName(v string) *GetMessageContactResponseBodyContact {
	s.Name = &v
	return s
}

func (s *GetMessageContactResponseBodyContact) SetPhoneNumber(v string) *GetMessageContactResponseBodyContact {
	s.PhoneNumber = &v
	return s
}

func (s *GetMessageContactResponseBodyContact) SetStatus(v string) *GetMessageContactResponseBodyContact {
	s.Status = &v
	return s
}

func (s *GetMessageContactResponseBodyContact) SetTitle(v string) *GetMessageContactResponseBodyContact {
	s.Title = &v
	return s
}

type GetMessageContactResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMessageContactResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMessageContactResponse) String() string {
	return tea.Prettify(s)
}

func (s GetMessageContactResponse) GoString() string {
	return s.String()
}

func (s *GetMessageContactResponse) SetHeaders(v map[string]*string) *GetMessageContactResponse {
	s.Headers = v
	return s
}

func (s *GetMessageContactResponse) SetStatusCode(v int32) *GetMessageContactResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMessageContactResponse) SetBody(v *GetMessageContactResponseBody) *GetMessageContactResponse {
	s.Body = v
	return s
}

type GetMessageContactDeletionStatusRequest struct {
	// The ID of the contact.
	//
	// example:
	//
	// c-qL4HqKONzOM7****
	ContactId *string `json:"ContactId,omitempty" xml:"ContactId,omitempty"`
}

func (s GetMessageContactDeletionStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s GetMessageContactDeletionStatusRequest) GoString() string {
	return s.String()
}

func (s *GetMessageContactDeletionStatusRequest) SetContactId(v string) *GetMessageContactDeletionStatusRequest {
	s.ContactId = &v
	return s
}

type GetMessageContactDeletionStatusResponseBody struct {
	// The deletion information of the contact.
	ContactDeletionStatus *GetMessageContactDeletionStatusResponseBodyContactDeletionStatus `json:"ContactDeletionStatus,omitempty" xml:"ContactDeletionStatus,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 95060F1D-6990-4645-8920-A81D1BBFE992
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetMessageContactDeletionStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetMessageContactDeletionStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetMessageContactDeletionStatusResponseBody) SetContactDeletionStatus(v *GetMessageContactDeletionStatusResponseBodyContactDeletionStatus) *GetMessageContactDeletionStatusResponseBody {
	s.ContactDeletionStatus = v
	return s
}

func (s *GetMessageContactDeletionStatusResponseBody) SetRequestId(v string) *GetMessageContactDeletionStatusResponseBody {
	s.RequestId = &v
	return s
}

type GetMessageContactDeletionStatusResponseBodyContactDeletionStatus struct {
	// The ID of the contact.
	//
	// example:
	//
	// c-qL4HqKONzOM7****
	ContactId *string `json:"ContactId,omitempty" xml:"ContactId,omitempty"`
	// The types of messages received by the contact.
	FailReasonList []*GetMessageContactDeletionStatusResponseBodyContactDeletionStatusFailReasonList `json:"FailReasonList,omitempty" xml:"FailReasonList,omitempty" type:"Repeated"`
	// The deletion status of the contact. Valid values:
	//
	// 	- Deleting
	//
	// 	- Failed
	//
	// example:
	//
	// Deleting
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetMessageContactDeletionStatusResponseBodyContactDeletionStatus) String() string {
	return tea.Prettify(s)
}

func (s GetMessageContactDeletionStatusResponseBodyContactDeletionStatus) GoString() string {
	return s.String()
}

func (s *GetMessageContactDeletionStatusResponseBodyContactDeletionStatus) SetContactId(v string) *GetMessageContactDeletionStatusResponseBodyContactDeletionStatus {
	s.ContactId = &v
	return s
}

func (s *GetMessageContactDeletionStatusResponseBodyContactDeletionStatus) SetFailReasonList(v []*GetMessageContactDeletionStatusResponseBodyContactDeletionStatusFailReasonList) *GetMessageContactDeletionStatusResponseBodyContactDeletionStatus {
	s.FailReasonList = v
	return s
}

func (s *GetMessageContactDeletionStatusResponseBodyContactDeletionStatus) SetStatus(v string) *GetMessageContactDeletionStatusResponseBodyContactDeletionStatus {
	s.Status = &v
	return s
}

type GetMessageContactDeletionStatusResponseBodyContactDeletionStatusFailReasonList struct {
	// The Alibaba Cloud account ID of the member.
	//
	// example:
	//
	// 199796839435****
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The types of messages received by the contact.
	MessageTypes []*string `json:"MessageTypes,omitempty" xml:"MessageTypes,omitempty" type:"Repeated"`
}

func (s GetMessageContactDeletionStatusResponseBodyContactDeletionStatusFailReasonList) String() string {
	return tea.Prettify(s)
}

func (s GetMessageContactDeletionStatusResponseBodyContactDeletionStatusFailReasonList) GoString() string {
	return s.String()
}

func (s *GetMessageContactDeletionStatusResponseBodyContactDeletionStatusFailReasonList) SetAccountId(v string) *GetMessageContactDeletionStatusResponseBodyContactDeletionStatusFailReasonList {
	s.AccountId = &v
	return s
}

func (s *GetMessageContactDeletionStatusResponseBodyContactDeletionStatusFailReasonList) SetMessageTypes(v []*string) *GetMessageContactDeletionStatusResponseBodyContactDeletionStatusFailReasonList {
	s.MessageTypes = v
	return s
}

type GetMessageContactDeletionStatusResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMessageContactDeletionStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMessageContactDeletionStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s GetMessageContactDeletionStatusResponse) GoString() string {
	return s.String()
}

func (s *GetMessageContactDeletionStatusResponse) SetHeaders(v map[string]*string) *GetMessageContactDeletionStatusResponse {
	s.Headers = v
	return s
}

func (s *GetMessageContactDeletionStatusResponse) SetStatusCode(v int32) *GetMessageContactDeletionStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMessageContactDeletionStatusResponse) SetBody(v *GetMessageContactDeletionStatusResponseBody) *GetMessageContactDeletionStatusResponse {
	s.Body = v
	return s
}

type GetPayerForAccountRequest struct {
	// The ID of the billing account.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12323344****
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
}

func (s GetPayerForAccountRequest) String() string {
	return tea.Prettify(s)
}

func (s GetPayerForAccountRequest) GoString() string {
	return s.String()
}

func (s *GetPayerForAccountRequest) SetAccountId(v string) *GetPayerForAccountRequest {
	s.AccountId = &v
	return s
}

type GetPayerForAccountResponseBody struct {
	// The ID of the billing account.
	//
	// example:
	//
	// 172841235500****
	PayerAccountId *string `json:"PayerAccountId,omitempty" xml:"PayerAccountId,omitempty"`
	// The name of the billing account.
	//
	// example:
	//
	// Alice
	PayerAccountName *string `json:"PayerAccountName,omitempty" xml:"PayerAccountName,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 9B34724D-54B0-4A51-B34D-4512372FE1BE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetPayerForAccountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetPayerForAccountResponseBody) GoString() string {
	return s.String()
}

func (s *GetPayerForAccountResponseBody) SetPayerAccountId(v string) *GetPayerForAccountResponseBody {
	s.PayerAccountId = &v
	return s
}

func (s *GetPayerForAccountResponseBody) SetPayerAccountName(v string) *GetPayerForAccountResponseBody {
	s.PayerAccountName = &v
	return s
}

func (s *GetPayerForAccountResponseBody) SetRequestId(v string) *GetPayerForAccountResponseBody {
	s.RequestId = &v
	return s
}

type GetPayerForAccountResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPayerForAccountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPayerForAccountResponse) String() string {
	return tea.Prettify(s)
}

func (s GetPayerForAccountResponse) GoString() string {
	return s.String()
}

func (s *GetPayerForAccountResponse) SetHeaders(v map[string]*string) *GetPayerForAccountResponse {
	s.Headers = v
	return s
}

func (s *GetPayerForAccountResponse) SetStatusCode(v int32) *GetPayerForAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPayerForAccountResponse) SetBody(v *GetPayerForAccountResponseBody) *GetPayerForAccountResponse {
	s.Body = v
	return s
}

type GetResourceDirectoryResponseBody struct {
	// example:
	//
	// CD76D376-2517-4924-92C5-DBC52262F93A
	RequestId         *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResourceDirectory *GetResourceDirectoryResponseBodyResourceDirectory `json:"ResourceDirectory,omitempty" xml:"ResourceDirectory,omitempty" type:"Struct"`
}

func (s GetResourceDirectoryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetResourceDirectoryResponseBody) GoString() string {
	return s.String()
}

func (s *GetResourceDirectoryResponseBody) SetRequestId(v string) *GetResourceDirectoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetResourceDirectoryResponseBody) SetResourceDirectory(v *GetResourceDirectoryResponseBodyResourceDirectory) *GetResourceDirectoryResponseBody {
	s.ResourceDirectory = v
	return s
}

type GetResourceDirectoryResponseBodyResourceDirectory struct {
	// example:
	//
	// Enabled
	ControlPolicyStatus *string `json:"ControlPolicyStatus,omitempty" xml:"ControlPolicyStatus,omitempty"`
	// example:
	//
	// 2019-02-18T15:32:10.473Z
	CreateTime          *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	IdentityInformation *string `json:"IdentityInformation,omitempty" xml:"IdentityInformation,omitempty"`
	// example:
	//
	// 172845045600****
	MasterAccountId *string `json:"MasterAccountId,omitempty" xml:"MasterAccountId,omitempty"`
	// example:
	//
	// aliyun-admin
	MasterAccountName *string `json:"MasterAccountName,omitempty" xml:"MasterAccountName,omitempty"`
	// example:
	//
	// Enabled
	MemberDeletionStatus *string `json:"MemberDeletionStatus,omitempty" xml:"MemberDeletionStatus,omitempty"`
	// example:
	//
	// rd-St****
	ResourceDirectoryId *string `json:"ResourceDirectoryId,omitempty" xml:"ResourceDirectoryId,omitempty"`
	// example:
	//
	// r-Zo****
	RootFolderId *string `json:"RootFolderId,omitempty" xml:"RootFolderId,omitempty"`
}

func (s GetResourceDirectoryResponseBodyResourceDirectory) String() string {
	return tea.Prettify(s)
}

func (s GetResourceDirectoryResponseBodyResourceDirectory) GoString() string {
	return s.String()
}

func (s *GetResourceDirectoryResponseBodyResourceDirectory) SetControlPolicyStatus(v string) *GetResourceDirectoryResponseBodyResourceDirectory {
	s.ControlPolicyStatus = &v
	return s
}

func (s *GetResourceDirectoryResponseBodyResourceDirectory) SetCreateTime(v string) *GetResourceDirectoryResponseBodyResourceDirectory {
	s.CreateTime = &v
	return s
}

func (s *GetResourceDirectoryResponseBodyResourceDirectory) SetIdentityInformation(v string) *GetResourceDirectoryResponseBodyResourceDirectory {
	s.IdentityInformation = &v
	return s
}

func (s *GetResourceDirectoryResponseBodyResourceDirectory) SetMasterAccountId(v string) *GetResourceDirectoryResponseBodyResourceDirectory {
	s.MasterAccountId = &v
	return s
}

func (s *GetResourceDirectoryResponseBodyResourceDirectory) SetMasterAccountName(v string) *GetResourceDirectoryResponseBodyResourceDirectory {
	s.MasterAccountName = &v
	return s
}

func (s *GetResourceDirectoryResponseBodyResourceDirectory) SetMemberDeletionStatus(v string) *GetResourceDirectoryResponseBodyResourceDirectory {
	s.MemberDeletionStatus = &v
	return s
}

func (s *GetResourceDirectoryResponseBodyResourceDirectory) SetResourceDirectoryId(v string) *GetResourceDirectoryResponseBodyResourceDirectory {
	s.ResourceDirectoryId = &v
	return s
}

func (s *GetResourceDirectoryResponseBodyResourceDirectory) SetRootFolderId(v string) *GetResourceDirectoryResponseBodyResourceDirectory {
	s.RootFolderId = &v
	return s
}

type GetResourceDirectoryResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetResourceDirectoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetResourceDirectoryResponse) String() string {
	return tea.Prettify(s)
}

func (s GetResourceDirectoryResponse) GoString() string {
	return s.String()
}

func (s *GetResourceDirectoryResponse) SetHeaders(v map[string]*string) *GetResourceDirectoryResponse {
	s.Headers = v
	return s
}

func (s *GetResourceDirectoryResponse) SetStatusCode(v int32) *GetResourceDirectoryResponse {
	s.StatusCode = &v
	return s
}

func (s *GetResourceDirectoryResponse) SetBody(v *GetResourceDirectoryResponseBody) *GetResourceDirectoryResponse {
	s.Body = v
	return s
}

type InviteAccountToResourceDirectoryRequest struct {
	// The description of the invitation.
	//
	// The description can be up to 1,024 characters in length.
	//
	// example:
	//
	// Welcome
	Note *string `json:"Note,omitempty" xml:"Note,omitempty"`
	// The ID of the parent folder.
	//
	// example:
	//
	// r-b1****
	ParentFolderId *string `json:"ParentFolderId,omitempty" xml:"ParentFolderId,omitempty"`
	// The tags.
	Tag []*InviteAccountToResourceDirectoryRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The ID or logon email address of the account that you want to invite.
	//
	// This parameter is required.
	//
	// example:
	//
	// someone@example.com
	TargetEntity *string `json:"TargetEntity,omitempty" xml:"TargetEntity,omitempty"`
	// The type of the invited account. Valid values:
	//
	// 	- Account: indicates the ID of the account.
	//
	// 	- Email: indicates the logon email address of the account.
	//
	// This parameter is required.
	//
	// example:
	//
	// Email
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
}

func (s InviteAccountToResourceDirectoryRequest) String() string {
	return tea.Prettify(s)
}

func (s InviteAccountToResourceDirectoryRequest) GoString() string {
	return s.String()
}

func (s *InviteAccountToResourceDirectoryRequest) SetNote(v string) *InviteAccountToResourceDirectoryRequest {
	s.Note = &v
	return s
}

func (s *InviteAccountToResourceDirectoryRequest) SetParentFolderId(v string) *InviteAccountToResourceDirectoryRequest {
	s.ParentFolderId = &v
	return s
}

func (s *InviteAccountToResourceDirectoryRequest) SetTag(v []*InviteAccountToResourceDirectoryRequestTag) *InviteAccountToResourceDirectoryRequest {
	s.Tag = v
	return s
}

func (s *InviteAccountToResourceDirectoryRequest) SetTargetEntity(v string) *InviteAccountToResourceDirectoryRequest {
	s.TargetEntity = &v
	return s
}

func (s *InviteAccountToResourceDirectoryRequest) SetTargetType(v string) *InviteAccountToResourceDirectoryRequest {
	s.TargetType = &v
	return s
}

type InviteAccountToResourceDirectoryRequestTag struct {
	// The tag key.
	//
	// example:
	//
	// k1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// v1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s InviteAccountToResourceDirectoryRequestTag) String() string {
	return tea.Prettify(s)
}

func (s InviteAccountToResourceDirectoryRequestTag) GoString() string {
	return s.String()
}

func (s *InviteAccountToResourceDirectoryRequestTag) SetKey(v string) *InviteAccountToResourceDirectoryRequestTag {
	s.Key = &v
	return s
}

func (s *InviteAccountToResourceDirectoryRequestTag) SetValue(v string) *InviteAccountToResourceDirectoryRequestTag {
	s.Value = &v
	return s
}

type InviteAccountToResourceDirectoryResponseBody struct {
	// The information about the invitation.
	Handshake *InviteAccountToResourceDirectoryResponseBodyHandshake `json:"Handshake,omitempty" xml:"Handshake,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 9B34724D-54B0-4A51-B34D-4512372FE1BE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s InviteAccountToResourceDirectoryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s InviteAccountToResourceDirectoryResponseBody) GoString() string {
	return s.String()
}

func (s *InviteAccountToResourceDirectoryResponseBody) SetHandshake(v *InviteAccountToResourceDirectoryResponseBodyHandshake) *InviteAccountToResourceDirectoryResponseBody {
	s.Handshake = v
	return s
}

func (s *InviteAccountToResourceDirectoryResponseBody) SetRequestId(v string) *InviteAccountToResourceDirectoryResponseBody {
	s.RequestId = &v
	return s
}

type InviteAccountToResourceDirectoryResponseBodyHandshake struct {
	// The time when the invitation was created. The time is displayed in UTC.
	//
	// example:
	//
	// 2018-08-10T09:55:41Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The time when the invitation expires. The time is displayed in UTC.
	//
	// example:
	//
	// 2018-08-24T09:55:41Z
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The ID of the invitation.
	//
	// example:
	//
	// h-ycm4rp****
	HandshakeId *string `json:"HandshakeId,omitempty" xml:"HandshakeId,omitempty"`
	// The ID of the management account of the resource directory.
	//
	// example:
	//
	// 172841235500****
	MasterAccountId *string `json:"MasterAccountId,omitempty" xml:"MasterAccountId,omitempty"`
	// The name of the management account of the resource directory.
	//
	// example:
	//
	// Alice
	MasterAccountName *string `json:"MasterAccountName,omitempty" xml:"MasterAccountName,omitempty"`
	// The time when the invitation was modified. The time is displayed in UTC.
	//
	// example:
	//
	// 2018-08-10T09:55:41Z
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The description of the invitation.
	//
	// example:
	//
	// Welcome
	Note *string `json:"Note,omitempty" xml:"Note,omitempty"`
	// The ID of the resource directory.
	//
	// example:
	//
	// rd-abcdef****
	ResourceDirectoryId *string `json:"ResourceDirectoryId,omitempty" xml:"ResourceDirectoryId,omitempty"`
	// The status of the invitation. Valid values:
	//
	// 	- Pending
	//
	// 	- Accepted
	//
	// 	- Cancelled
	//
	// 	- Declined
	//
	// 	- Expired
	//
	// example:
	//
	// Pending
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID or logon email address of the invited account.
	//
	// example:
	//
	// someone@example.com
	TargetEntity *string `json:"TargetEntity,omitempty" xml:"TargetEntity,omitempty"`
	// The type of the invited account. Valid values:
	//
	// 	- Account: indicates the ID of the account.
	//
	// 	- Email: indicates the logon email address of the account.
	//
	// example:
	//
	// Email
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
}

func (s InviteAccountToResourceDirectoryResponseBodyHandshake) String() string {
	return tea.Prettify(s)
}

func (s InviteAccountToResourceDirectoryResponseBodyHandshake) GoString() string {
	return s.String()
}

func (s *InviteAccountToResourceDirectoryResponseBodyHandshake) SetCreateTime(v string) *InviteAccountToResourceDirectoryResponseBodyHandshake {
	s.CreateTime = &v
	return s
}

func (s *InviteAccountToResourceDirectoryResponseBodyHandshake) SetExpireTime(v string) *InviteAccountToResourceDirectoryResponseBodyHandshake {
	s.ExpireTime = &v
	return s
}

func (s *InviteAccountToResourceDirectoryResponseBodyHandshake) SetHandshakeId(v string) *InviteAccountToResourceDirectoryResponseBodyHandshake {
	s.HandshakeId = &v
	return s
}

func (s *InviteAccountToResourceDirectoryResponseBodyHandshake) SetMasterAccountId(v string) *InviteAccountToResourceDirectoryResponseBodyHandshake {
	s.MasterAccountId = &v
	return s
}

func (s *InviteAccountToResourceDirectoryResponseBodyHandshake) SetMasterAccountName(v string) *InviteAccountToResourceDirectoryResponseBodyHandshake {
	s.MasterAccountName = &v
	return s
}

func (s *InviteAccountToResourceDirectoryResponseBodyHandshake) SetModifyTime(v string) *InviteAccountToResourceDirectoryResponseBodyHandshake {
	s.ModifyTime = &v
	return s
}

func (s *InviteAccountToResourceDirectoryResponseBodyHandshake) SetNote(v string) *InviteAccountToResourceDirectoryResponseBodyHandshake {
	s.Note = &v
	return s
}

func (s *InviteAccountToResourceDirectoryResponseBodyHandshake) SetResourceDirectoryId(v string) *InviteAccountToResourceDirectoryResponseBodyHandshake {
	s.ResourceDirectoryId = &v
	return s
}

func (s *InviteAccountToResourceDirectoryResponseBodyHandshake) SetStatus(v string) *InviteAccountToResourceDirectoryResponseBodyHandshake {
	s.Status = &v
	return s
}

func (s *InviteAccountToResourceDirectoryResponseBodyHandshake) SetTargetEntity(v string) *InviteAccountToResourceDirectoryResponseBodyHandshake {
	s.TargetEntity = &v
	return s
}

func (s *InviteAccountToResourceDirectoryResponseBodyHandshake) SetTargetType(v string) *InviteAccountToResourceDirectoryResponseBodyHandshake {
	s.TargetType = &v
	return s
}

type InviteAccountToResourceDirectoryResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InviteAccountToResourceDirectoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InviteAccountToResourceDirectoryResponse) String() string {
	return tea.Prettify(s)
}

func (s InviteAccountToResourceDirectoryResponse) GoString() string {
	return s.String()
}

func (s *InviteAccountToResourceDirectoryResponse) SetHeaders(v map[string]*string) *InviteAccountToResourceDirectoryResponse {
	s.Headers = v
	return s
}

func (s *InviteAccountToResourceDirectoryResponse) SetStatusCode(v int32) *InviteAccountToResourceDirectoryResponse {
	s.StatusCode = &v
	return s
}

func (s *InviteAccountToResourceDirectoryResponse) SetBody(v *InviteAccountToResourceDirectoryResponseBody) *InviteAccountToResourceDirectoryResponse {
	s.Body = v
	return s
}

type ListAccountsRequest struct {
	// Specifies whether to return information about tags. Valid values:
	//
	// 	- false (default value)
	//
	// 	- true
	//
	// example:
	//
	// true
	IncludeTags *bool `json:"IncludeTags,omitempty" xml:"IncludeTags,omitempty"`
	// The number of the page to return.
	//
	// Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page.
	//
	// Valid values: 1 to 100. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The keyword used for the query, such as the display name of a member.
	//
	// Fuzzy match is supported.
	//
	// example:
	//
	// Advance
	QueryKeyword *string `json:"QueryKeyword,omitempty" xml:"QueryKeyword,omitempty"`
	// The tags. This parameter specifies a filter condition.
	Tag []*ListAccountsRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListAccountsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAccountsRequest) GoString() string {
	return s.String()
}

func (s *ListAccountsRequest) SetIncludeTags(v bool) *ListAccountsRequest {
	s.IncludeTags = &v
	return s
}

func (s *ListAccountsRequest) SetPageNumber(v int32) *ListAccountsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListAccountsRequest) SetPageSize(v int32) *ListAccountsRequest {
	s.PageSize = &v
	return s
}

func (s *ListAccountsRequest) SetQueryKeyword(v string) *ListAccountsRequest {
	s.QueryKeyword = &v
	return s
}

func (s *ListAccountsRequest) SetTag(v []*ListAccountsRequestTag) *ListAccountsRequest {
	s.Tag = v
	return s
}

type ListAccountsRequestTag struct {
	// The tag key.
	//
	// example:
	//
	// tag_key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// tag_value
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListAccountsRequestTag) String() string {
	return tea.Prettify(s)
}

func (s ListAccountsRequestTag) GoString() string {
	return s.String()
}

func (s *ListAccountsRequestTag) SetKey(v string) *ListAccountsRequestTag {
	s.Key = &v
	return s
}

func (s *ListAccountsRequestTag) SetValue(v string) *ListAccountsRequestTag {
	s.Value = &v
	return s
}

type ListAccountsResponseBody struct {
	// The information about the members.
	Accounts *ListAccountsResponseBodyAccounts `json:"Accounts,omitempty" xml:"Accounts,omitempty" type:"Struct"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 5
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 7B8A4E7D-6CFF-471D-84DF-195A7A241ECB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 2
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListAccountsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAccountsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAccountsResponseBody) SetAccounts(v *ListAccountsResponseBodyAccounts) *ListAccountsResponseBody {
	s.Accounts = v
	return s
}

func (s *ListAccountsResponseBody) SetPageNumber(v int32) *ListAccountsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListAccountsResponseBody) SetPageSize(v int32) *ListAccountsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListAccountsResponseBody) SetRequestId(v string) *ListAccountsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAccountsResponseBody) SetTotalCount(v int32) *ListAccountsResponseBody {
	s.TotalCount = &v
	return s
}

type ListAccountsResponseBodyAccounts struct {
	Account []*ListAccountsResponseBodyAccountsAccount `json:"Account,omitempty" xml:"Account,omitempty" type:"Repeated"`
}

func (s ListAccountsResponseBodyAccounts) String() string {
	return tea.Prettify(s)
}

func (s ListAccountsResponseBodyAccounts) GoString() string {
	return s.String()
}

func (s *ListAccountsResponseBodyAccounts) SetAccount(v []*ListAccountsResponseBodyAccountsAccount) *ListAccountsResponseBodyAccounts {
	s.Account = v
	return s
}

type ListAccountsResponseBodyAccountsAccount struct {
	// The Alibaba Cloud account ID of the member.
	//
	// example:
	//
	// 181761095690****
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The Alibaba Cloud account name of the member.
	//
	// example:
	//
	// oxy01414357@alibaba-inc.com
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The deletion status of the member. Valid values:
	//
	// 	- Checking: A deletion check is being performed for the member.
	//
	// 	- Deleting: The member is being deleted.
	//
	// 	- CheckFailed: The deletion check for the member fails.
	//
	// 	- DeleteFailed: The member fails to be deleted.
	//
	// >  If deletion is not performed for the member, the value of this parameter is empty.
	//
	// example:
	//
	// Checking
	DeletionStatus *string `json:"DeletionStatus,omitempty" xml:"DeletionStatus,omitempty"`
	// The display name of the member.
	//
	// example:
	//
	// test
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The ID of the folder.
	//
	// example:
	//
	// fd-QRzuim****
	FolderId *string `json:"FolderId,omitempty" xml:"FolderId,omitempty"`
	// The way in which the member joins the resource directory. Valid values:
	//
	// 	- invited: The member is invited to join the resource directory.
	//
	// 	- created: The member is directly created in the resource directory.
	//
	// example:
	//
	// created
	JoinMethod *string `json:"JoinMethod,omitempty" xml:"JoinMethod,omitempty"`
	// The time when the member joined the resource directory. The time is displayed in UTC.
	//
	// example:
	//
	// 2021-01-18T08:01:50.522Z
	JoinTime *string `json:"JoinTime,omitempty" xml:"JoinTime,omitempty"`
	// The time when the member was modified. The time is displayed in UTC.
	//
	// example:
	//
	// 2021-01-18T08:04:37.668Z
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The ID of the resource directory.
	//
	// example:
	//
	// rd-3G****
	ResourceDirectoryId *string `json:"ResourceDirectoryId,omitempty" xml:"ResourceDirectoryId,omitempty"`
	// The RDPath of the member.
	ResourceDirectoryPath *string `json:"ResourceDirectoryPath,omitempty" xml:"ResourceDirectoryPath,omitempty"`
	// The status of the member. Valid values:
	//
	// 	- CreateSuccess: The member is created.
	//
	// 	- PromoteVerifying: The upgrade of the member is under confirmation.
	//
	// 	- PromoteFailed: The upgrade of the member fails.
	//
	// 	- PromoteExpired: The upgrade of the member expires.
	//
	// 	- PromoteCancelled: The upgrade of the member is canceled.
	//
	// 	- PromoteSuccess: The member is upgraded.
	//
	// 	- InviteSuccess: The member accepts the invitation.
	//
	// example:
	//
	// CreateSuccess
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tags that are added to the member.
	Tags *ListAccountsResponseBodyAccountsAccountTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	// The type of the member. Valid values:
	//
	// 	- CloudAccount: cloud account
	//
	// 	- ResourceAccount: resource account
	//
	// example:
	//
	// ResourceAccount
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListAccountsResponseBodyAccountsAccount) String() string {
	return tea.Prettify(s)
}

func (s ListAccountsResponseBodyAccountsAccount) GoString() string {
	return s.String()
}

func (s *ListAccountsResponseBodyAccountsAccount) SetAccountId(v string) *ListAccountsResponseBodyAccountsAccount {
	s.AccountId = &v
	return s
}

func (s *ListAccountsResponseBodyAccountsAccount) SetAccountName(v string) *ListAccountsResponseBodyAccountsAccount {
	s.AccountName = &v
	return s
}

func (s *ListAccountsResponseBodyAccountsAccount) SetDeletionStatus(v string) *ListAccountsResponseBodyAccountsAccount {
	s.DeletionStatus = &v
	return s
}

func (s *ListAccountsResponseBodyAccountsAccount) SetDisplayName(v string) *ListAccountsResponseBodyAccountsAccount {
	s.DisplayName = &v
	return s
}

func (s *ListAccountsResponseBodyAccountsAccount) SetFolderId(v string) *ListAccountsResponseBodyAccountsAccount {
	s.FolderId = &v
	return s
}

func (s *ListAccountsResponseBodyAccountsAccount) SetJoinMethod(v string) *ListAccountsResponseBodyAccountsAccount {
	s.JoinMethod = &v
	return s
}

func (s *ListAccountsResponseBodyAccountsAccount) SetJoinTime(v string) *ListAccountsResponseBodyAccountsAccount {
	s.JoinTime = &v
	return s
}

func (s *ListAccountsResponseBodyAccountsAccount) SetModifyTime(v string) *ListAccountsResponseBodyAccountsAccount {
	s.ModifyTime = &v
	return s
}

func (s *ListAccountsResponseBodyAccountsAccount) SetResourceDirectoryId(v string) *ListAccountsResponseBodyAccountsAccount {
	s.ResourceDirectoryId = &v
	return s
}

func (s *ListAccountsResponseBodyAccountsAccount) SetResourceDirectoryPath(v string) *ListAccountsResponseBodyAccountsAccount {
	s.ResourceDirectoryPath = &v
	return s
}

func (s *ListAccountsResponseBodyAccountsAccount) SetStatus(v string) *ListAccountsResponseBodyAccountsAccount {
	s.Status = &v
	return s
}

func (s *ListAccountsResponseBodyAccountsAccount) SetTags(v *ListAccountsResponseBodyAccountsAccountTags) *ListAccountsResponseBodyAccountsAccount {
	s.Tags = v
	return s
}

func (s *ListAccountsResponseBodyAccountsAccount) SetType(v string) *ListAccountsResponseBodyAccountsAccount {
	s.Type = &v
	return s
}

type ListAccountsResponseBodyAccountsAccountTags struct {
	Tag []*ListAccountsResponseBodyAccountsAccountTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListAccountsResponseBodyAccountsAccountTags) String() string {
	return tea.Prettify(s)
}

func (s ListAccountsResponseBodyAccountsAccountTags) GoString() string {
	return s.String()
}

func (s *ListAccountsResponseBodyAccountsAccountTags) SetTag(v []*ListAccountsResponseBodyAccountsAccountTagsTag) *ListAccountsResponseBodyAccountsAccountTags {
	s.Tag = v
	return s
}

type ListAccountsResponseBodyAccountsAccountTagsTag struct {
	// The key of the tag.
	//
	// example:
	//
	// tag_key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag.
	//
	// example:
	//
	// tag_value
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListAccountsResponseBodyAccountsAccountTagsTag) String() string {
	return tea.Prettify(s)
}

func (s ListAccountsResponseBodyAccountsAccountTagsTag) GoString() string {
	return s.String()
}

func (s *ListAccountsResponseBodyAccountsAccountTagsTag) SetKey(v string) *ListAccountsResponseBodyAccountsAccountTagsTag {
	s.Key = &v
	return s
}

func (s *ListAccountsResponseBodyAccountsAccountTagsTag) SetValue(v string) *ListAccountsResponseBodyAccountsAccountTagsTag {
	s.Value = &v
	return s
}

type ListAccountsResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAccountsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAccountsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAccountsResponse) GoString() string {
	return s.String()
}

func (s *ListAccountsResponse) SetHeaders(v map[string]*string) *ListAccountsResponse {
	s.Headers = v
	return s
}

func (s *ListAccountsResponse) SetStatusCode(v int32) *ListAccountsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAccountsResponse) SetBody(v *ListAccountsResponseBody) *ListAccountsResponse {
	s.Body = v
	return s
}

type ListAccountsForParentRequest struct {
	// Specifies whether to return information about tags. Valid values:
	//
	// 	- false (default value)
	//
	// 	- true
	//
	// example:
	//
	// true
	IncludeTags *bool `json:"IncludeTags,omitempty" xml:"IncludeTags,omitempty"`
	// The number of the page to return.
	//
	// Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page.
	//
	// Valid values: 1 to 100. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the folder.
	//
	// example:
	//
	// fd-bVaRIG****
	ParentFolderId *string `json:"ParentFolderId,omitempty" xml:"ParentFolderId,omitempty"`
	// The keyword used for the query, such as the display name of a member.
	//
	// Fuzzy match is supported.
	//
	// example:
	//
	// admin
	QueryKeyword *string `json:"QueryKeyword,omitempty" xml:"QueryKeyword,omitempty"`
	// The tags. This parameter specifies a filter condition.
	Tag []*ListAccountsForParentRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListAccountsForParentRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAccountsForParentRequest) GoString() string {
	return s.String()
}

func (s *ListAccountsForParentRequest) SetIncludeTags(v bool) *ListAccountsForParentRequest {
	s.IncludeTags = &v
	return s
}

func (s *ListAccountsForParentRequest) SetPageNumber(v int32) *ListAccountsForParentRequest {
	s.PageNumber = &v
	return s
}

func (s *ListAccountsForParentRequest) SetPageSize(v int32) *ListAccountsForParentRequest {
	s.PageSize = &v
	return s
}

func (s *ListAccountsForParentRequest) SetParentFolderId(v string) *ListAccountsForParentRequest {
	s.ParentFolderId = &v
	return s
}

func (s *ListAccountsForParentRequest) SetQueryKeyword(v string) *ListAccountsForParentRequest {
	s.QueryKeyword = &v
	return s
}

func (s *ListAccountsForParentRequest) SetTag(v []*ListAccountsForParentRequestTag) *ListAccountsForParentRequest {
	s.Tag = v
	return s
}

type ListAccountsForParentRequestTag struct {
	// The key of the tag.
	//
	// example:
	//
	// tag_key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag.
	//
	// example:
	//
	// tag_value
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListAccountsForParentRequestTag) String() string {
	return tea.Prettify(s)
}

func (s ListAccountsForParentRequestTag) GoString() string {
	return s.String()
}

func (s *ListAccountsForParentRequestTag) SetKey(v string) *ListAccountsForParentRequestTag {
	s.Key = &v
	return s
}

func (s *ListAccountsForParentRequestTag) SetValue(v string) *ListAccountsForParentRequestTag {
	s.Value = &v
	return s
}

type ListAccountsForParentResponseBody struct {
	// The information about the members.
	Accounts *ListAccountsForParentResponseBodyAccounts `json:"Accounts,omitempty" xml:"Accounts,omitempty" type:"Struct"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 5
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 7B8A4E7D-6CFF-471D-84DF-195A7A241ECB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 2
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListAccountsForParentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAccountsForParentResponseBody) GoString() string {
	return s.String()
}

func (s *ListAccountsForParentResponseBody) SetAccounts(v *ListAccountsForParentResponseBodyAccounts) *ListAccountsForParentResponseBody {
	s.Accounts = v
	return s
}

func (s *ListAccountsForParentResponseBody) SetPageNumber(v int32) *ListAccountsForParentResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListAccountsForParentResponseBody) SetPageSize(v int32) *ListAccountsForParentResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListAccountsForParentResponseBody) SetRequestId(v string) *ListAccountsForParentResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAccountsForParentResponseBody) SetTotalCount(v int32) *ListAccountsForParentResponseBody {
	s.TotalCount = &v
	return s
}

type ListAccountsForParentResponseBodyAccounts struct {
	Account []*ListAccountsForParentResponseBodyAccountsAccount `json:"Account,omitempty" xml:"Account,omitempty" type:"Repeated"`
}

func (s ListAccountsForParentResponseBodyAccounts) String() string {
	return tea.Prettify(s)
}

func (s ListAccountsForParentResponseBodyAccounts) GoString() string {
	return s.String()
}

func (s *ListAccountsForParentResponseBodyAccounts) SetAccount(v []*ListAccountsForParentResponseBodyAccountsAccount) *ListAccountsForParentResponseBodyAccounts {
	s.Account = v
	return s
}

type ListAccountsForParentResponseBodyAccountsAccount struct {
	// The Alibaba Cloud account ID of the member.
	//
	// example:
	//
	// 184311716100****
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The Alibaba Cloud account name of the member.
	//
	// example:
	//
	// hdd01429358@alibaba-inc.com
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The deletion status of the member. Valid values:
	//
	// 	- Checking: A deletion check is being performed for the member.
	//
	// 	- Deleting: The member is being deleted.
	//
	// 	- CheckFailed: The deletion check for the member fails.
	//
	// 	- DeleteFailed: The member fails to be deleted.
	//
	// >  If deletion is not performed for the member, the value of this parameter is empty.
	//
	// example:
	//
	// Checking
	DeletionStatus *string `json:"DeletionStatus,omitempty" xml:"DeletionStatus,omitempty"`
	// The display name of the member.
	//
	// example:
	//
	// admin
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The ID of the folder.
	//
	// example:
	//
	// fd-bVaRIG****
	FolderId *string `json:"FolderId,omitempty" xml:"FolderId,omitempty"`
	// The way in which the member joins the resource directory.
	//
	// 	- invited: The member is invited to join the resource directory.
	//
	// 	- created: The member is directly created in the resource directory.
	//
	// example:
	//
	// created
	JoinMethod *string `json:"JoinMethod,omitempty" xml:"JoinMethod,omitempty"`
	// The time when the member joined the resource directory. The time is displayed in UTC.
	//
	// example:
	//
	// 2015-01-23T12:33:18Z
	JoinTime *string `json:"JoinTime,omitempty" xml:"JoinTime,omitempty"`
	// The time when the member was modified. The time is displayed in UTC.
	//
	// example:
	//
	// 2015-01-23T12:33:18Z
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The ID of the resource directory.
	//
	// example:
	//
	// rd-k4****
	ResourceDirectoryId *string `json:"ResourceDirectoryId,omitempty" xml:"ResourceDirectoryId,omitempty"`
	// The status of the member. Valid values:
	//
	// 	- CreateSuccess: The member is created.
	//
	// 	- PromoteVerifying: The upgrade of the member is under confirmation.
	//
	// 	- PromoteFailed: The upgrade of the member fails.
	//
	// 	- PromoteExpired: The upgrade of the member expires.
	//
	// 	- PromoteCancelled: The upgrade of the member is canceled.
	//
	// 	- PromoteSuccess: The member is upgraded.
	//
	// 	- InviteSuccess: The member accepts the invitation.
	//
	// example:
	//
	// CreateSuccess
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tags that are added to the member.
	Tags *ListAccountsForParentResponseBodyAccountsAccountTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	// The type of the member. Valid values:
	//
	// 	- CloudAccount: cloud account
	//
	// 	- ResourceAccount: resource account
	//
	// example:
	//
	// ResourceAccount
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListAccountsForParentResponseBodyAccountsAccount) String() string {
	return tea.Prettify(s)
}

func (s ListAccountsForParentResponseBodyAccountsAccount) GoString() string {
	return s.String()
}

func (s *ListAccountsForParentResponseBodyAccountsAccount) SetAccountId(v string) *ListAccountsForParentResponseBodyAccountsAccount {
	s.AccountId = &v
	return s
}

func (s *ListAccountsForParentResponseBodyAccountsAccount) SetAccountName(v string) *ListAccountsForParentResponseBodyAccountsAccount {
	s.AccountName = &v
	return s
}

func (s *ListAccountsForParentResponseBodyAccountsAccount) SetDeletionStatus(v string) *ListAccountsForParentResponseBodyAccountsAccount {
	s.DeletionStatus = &v
	return s
}

func (s *ListAccountsForParentResponseBodyAccountsAccount) SetDisplayName(v string) *ListAccountsForParentResponseBodyAccountsAccount {
	s.DisplayName = &v
	return s
}

func (s *ListAccountsForParentResponseBodyAccountsAccount) SetFolderId(v string) *ListAccountsForParentResponseBodyAccountsAccount {
	s.FolderId = &v
	return s
}

func (s *ListAccountsForParentResponseBodyAccountsAccount) SetJoinMethod(v string) *ListAccountsForParentResponseBodyAccountsAccount {
	s.JoinMethod = &v
	return s
}

func (s *ListAccountsForParentResponseBodyAccountsAccount) SetJoinTime(v string) *ListAccountsForParentResponseBodyAccountsAccount {
	s.JoinTime = &v
	return s
}

func (s *ListAccountsForParentResponseBodyAccountsAccount) SetModifyTime(v string) *ListAccountsForParentResponseBodyAccountsAccount {
	s.ModifyTime = &v
	return s
}

func (s *ListAccountsForParentResponseBodyAccountsAccount) SetResourceDirectoryId(v string) *ListAccountsForParentResponseBodyAccountsAccount {
	s.ResourceDirectoryId = &v
	return s
}

func (s *ListAccountsForParentResponseBodyAccountsAccount) SetStatus(v string) *ListAccountsForParentResponseBodyAccountsAccount {
	s.Status = &v
	return s
}

func (s *ListAccountsForParentResponseBodyAccountsAccount) SetTags(v *ListAccountsForParentResponseBodyAccountsAccountTags) *ListAccountsForParentResponseBodyAccountsAccount {
	s.Tags = v
	return s
}

func (s *ListAccountsForParentResponseBodyAccountsAccount) SetType(v string) *ListAccountsForParentResponseBodyAccountsAccount {
	s.Type = &v
	return s
}

type ListAccountsForParentResponseBodyAccountsAccountTags struct {
	Tag []*ListAccountsForParentResponseBodyAccountsAccountTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListAccountsForParentResponseBodyAccountsAccountTags) String() string {
	return tea.Prettify(s)
}

func (s ListAccountsForParentResponseBodyAccountsAccountTags) GoString() string {
	return s.String()
}

func (s *ListAccountsForParentResponseBodyAccountsAccountTags) SetTag(v []*ListAccountsForParentResponseBodyAccountsAccountTagsTag) *ListAccountsForParentResponseBodyAccountsAccountTags {
	s.Tag = v
	return s
}

type ListAccountsForParentResponseBodyAccountsAccountTagsTag struct {
	// The key of the tag.
	//
	// example:
	//
	// tag_key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag.
	//
	// example:
	//
	// tag_value
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListAccountsForParentResponseBodyAccountsAccountTagsTag) String() string {
	return tea.Prettify(s)
}

func (s ListAccountsForParentResponseBodyAccountsAccountTagsTag) GoString() string {
	return s.String()
}

func (s *ListAccountsForParentResponseBodyAccountsAccountTagsTag) SetKey(v string) *ListAccountsForParentResponseBodyAccountsAccountTagsTag {
	s.Key = &v
	return s
}

func (s *ListAccountsForParentResponseBodyAccountsAccountTagsTag) SetValue(v string) *ListAccountsForParentResponseBodyAccountsAccountTagsTag {
	s.Value = &v
	return s
}

type ListAccountsForParentResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAccountsForParentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAccountsForParentResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAccountsForParentResponse) GoString() string {
	return s.String()
}

func (s *ListAccountsForParentResponse) SetHeaders(v map[string]*string) *ListAccountsForParentResponse {
	s.Headers = v
	return s
}

func (s *ListAccountsForParentResponse) SetStatusCode(v int32) *ListAccountsForParentResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAccountsForParentResponse) SetBody(v *ListAccountsForParentResponseBody) *ListAccountsForParentResponse {
	s.Body = v
	return s
}

type ListAncestorsRequest struct {
	// The ID of the subfolder.
	//
	// This parameter is required.
	//
	// example:
	//
	// fd-i1c9nr****
	ChildId *string `json:"ChildId,omitempty" xml:"ChildId,omitempty"`
}

func (s ListAncestorsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAncestorsRequest) GoString() string {
	return s.String()
}

func (s *ListAncestorsRequest) SetChildId(v string) *ListAncestorsRequest {
	s.ChildId = &v
	return s
}

type ListAncestorsResponseBody struct {
	// The information of the folders.
	Folders *ListAncestorsResponseBodyFolders `json:"Folders,omitempty" xml:"Folders,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 83AFBEB6-DC03-406E-9686-867461FF6698
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListAncestorsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAncestorsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAncestorsResponseBody) SetFolders(v *ListAncestorsResponseBodyFolders) *ListAncestorsResponseBody {
	s.Folders = v
	return s
}

func (s *ListAncestorsResponseBody) SetRequestId(v string) *ListAncestorsResponseBody {
	s.RequestId = &v
	return s
}

type ListAncestorsResponseBodyFolders struct {
	Folder []*ListAncestorsResponseBodyFoldersFolder `json:"Folder,omitempty" xml:"Folder,omitempty" type:"Repeated"`
}

func (s ListAncestorsResponseBodyFolders) String() string {
	return tea.Prettify(s)
}

func (s ListAncestorsResponseBodyFolders) GoString() string {
	return s.String()
}

func (s *ListAncestorsResponseBodyFolders) SetFolder(v []*ListAncestorsResponseBodyFoldersFolder) *ListAncestorsResponseBodyFolders {
	s.Folder = v
	return s
}

type ListAncestorsResponseBodyFoldersFolder struct {
	// The time when the folder was created.
	//
	// example:
	//
	// 2019-01-18T10:03:35.217Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the folder.
	//
	// example:
	//
	// r-b1****
	FolderId *string `json:"FolderId,omitempty" xml:"FolderId,omitempty"`
	// The name of the folder.
	//
	// example:
	//
	// root
	FolderName *string `json:"FolderName,omitempty" xml:"FolderName,omitempty"`
}

func (s ListAncestorsResponseBodyFoldersFolder) String() string {
	return tea.Prettify(s)
}

func (s ListAncestorsResponseBodyFoldersFolder) GoString() string {
	return s.String()
}

func (s *ListAncestorsResponseBodyFoldersFolder) SetCreateTime(v string) *ListAncestorsResponseBodyFoldersFolder {
	s.CreateTime = &v
	return s
}

func (s *ListAncestorsResponseBodyFoldersFolder) SetFolderId(v string) *ListAncestorsResponseBodyFoldersFolder {
	s.FolderId = &v
	return s
}

func (s *ListAncestorsResponseBodyFoldersFolder) SetFolderName(v string) *ListAncestorsResponseBodyFoldersFolder {
	s.FolderName = &v
	return s
}

type ListAncestorsResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAncestorsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAncestorsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAncestorsResponse) GoString() string {
	return s.String()
}

func (s *ListAncestorsResponse) SetHeaders(v map[string]*string) *ListAncestorsResponse {
	s.Headers = v
	return s
}

func (s *ListAncestorsResponse) SetStatusCode(v int32) *ListAncestorsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAncestorsResponse) SetBody(v *ListAncestorsResponseBody) *ListAncestorsResponse {
	s.Body = v
	return s
}

type ListControlPoliciesRequest struct {
	// The language in which you want to return the descriptions of the access control policies. Valid values:
	//
	// 	- zh-CN (default value): Chinese
	//
	// 	- en: English
	//
	// 	- ja: Japanese
	//
	// > This parameter is available only for system access control policies.
	//
	// example:
	//
	// zh-CN
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// The number of the page to return.
	//
	// Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page.
	//
	// Valid values: 1 to 100. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The type of the access control policies. Valid values:
	//
	// 	- System: system access control policy
	//
	// 	- Custom: custom access control policy
	//
	// example:
	//
	// System
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s ListControlPoliciesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListControlPoliciesRequest) GoString() string {
	return s.String()
}

func (s *ListControlPoliciesRequest) SetLanguage(v string) *ListControlPoliciesRequest {
	s.Language = &v
	return s
}

func (s *ListControlPoliciesRequest) SetPageNumber(v int32) *ListControlPoliciesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListControlPoliciesRequest) SetPageSize(v int32) *ListControlPoliciesRequest {
	s.PageSize = &v
	return s
}

func (s *ListControlPoliciesRequest) SetPolicyType(v string) *ListControlPoliciesRequest {
	s.PolicyType = &v
	return s
}

type ListControlPoliciesResponseBody struct {
	// The information of the access control policies.
	ControlPolicies *ListControlPoliciesResponseBodyControlPolicies `json:"ControlPolicies,omitempty" xml:"ControlPolicies,omitempty" type:"Struct"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 9424A34C-3471-45AD-B6AB-924BBDFE42F9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of access control policies.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListControlPoliciesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListControlPoliciesResponseBody) GoString() string {
	return s.String()
}

func (s *ListControlPoliciesResponseBody) SetControlPolicies(v *ListControlPoliciesResponseBodyControlPolicies) *ListControlPoliciesResponseBody {
	s.ControlPolicies = v
	return s
}

func (s *ListControlPoliciesResponseBody) SetPageNumber(v int32) *ListControlPoliciesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListControlPoliciesResponseBody) SetPageSize(v int32) *ListControlPoliciesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListControlPoliciesResponseBody) SetRequestId(v string) *ListControlPoliciesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListControlPoliciesResponseBody) SetTotalCount(v int32) *ListControlPoliciesResponseBody {
	s.TotalCount = &v
	return s
}

type ListControlPoliciesResponseBodyControlPolicies struct {
	ControlPolicy []*ListControlPoliciesResponseBodyControlPoliciesControlPolicy `json:"ControlPolicy,omitempty" xml:"ControlPolicy,omitempty" type:"Repeated"`
}

func (s ListControlPoliciesResponseBodyControlPolicies) String() string {
	return tea.Prettify(s)
}

func (s ListControlPoliciesResponseBodyControlPolicies) GoString() string {
	return s.String()
}

func (s *ListControlPoliciesResponseBodyControlPolicies) SetControlPolicy(v []*ListControlPoliciesResponseBodyControlPoliciesControlPolicy) *ListControlPoliciesResponseBodyControlPolicies {
	s.ControlPolicy = v
	return s
}

type ListControlPoliciesResponseBodyControlPoliciesControlPolicy struct {
	// The number of times that the access control policy is referenced.
	//
	// example:
	//
	// 44
	AttachmentCount *string `json:"AttachmentCount,omitempty" xml:"AttachmentCount,omitempty"`
	// The time when the access control policy was created.
	//
	// example:
	//
	// 2020-08-05T06:32:24Z
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	// The description of the access control policy.
	//
	// example:
	//
	// System access control policy available for all operations on the cloud
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The effective scope of the access control policy. Valid values:
	//
	// 	- All: The access control policy is in effect for Alibaba Cloud accounts, RAM users, and RAM roles.
	//
	// 	- RAM: The access control policy is in effect only for RAM users and RAM roles.
	//
	// example:
	//
	// All
	EffectScope *string `json:"EffectScope,omitempty" xml:"EffectScope,omitempty"`
	// The ID of the access control policy.
	//
	// example:
	//
	// cp-FullAliyunAccess
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// The name of the access control policy.
	//
	// example:
	//
	// FullAliyunAccess
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	// The type of the access control policy. Valid values:
	//
	// 	- System: system access control policy
	//
	// 	- Custom: custom access control policy
	//
	// example:
	//
	// System
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
	// The time when the access control policy was updated.
	//
	// example:
	//
	// 2020-08-05T06:32:24Z
	UpdateDate *string `json:"UpdateDate,omitempty" xml:"UpdateDate,omitempty"`
}

func (s ListControlPoliciesResponseBodyControlPoliciesControlPolicy) String() string {
	return tea.Prettify(s)
}

func (s ListControlPoliciesResponseBodyControlPoliciesControlPolicy) GoString() string {
	return s.String()
}

func (s *ListControlPoliciesResponseBodyControlPoliciesControlPolicy) SetAttachmentCount(v string) *ListControlPoliciesResponseBodyControlPoliciesControlPolicy {
	s.AttachmentCount = &v
	return s
}

func (s *ListControlPoliciesResponseBodyControlPoliciesControlPolicy) SetCreateDate(v string) *ListControlPoliciesResponseBodyControlPoliciesControlPolicy {
	s.CreateDate = &v
	return s
}

func (s *ListControlPoliciesResponseBodyControlPoliciesControlPolicy) SetDescription(v string) *ListControlPoliciesResponseBodyControlPoliciesControlPolicy {
	s.Description = &v
	return s
}

func (s *ListControlPoliciesResponseBodyControlPoliciesControlPolicy) SetEffectScope(v string) *ListControlPoliciesResponseBodyControlPoliciesControlPolicy {
	s.EffectScope = &v
	return s
}

func (s *ListControlPoliciesResponseBodyControlPoliciesControlPolicy) SetPolicyId(v string) *ListControlPoliciesResponseBodyControlPoliciesControlPolicy {
	s.PolicyId = &v
	return s
}

func (s *ListControlPoliciesResponseBodyControlPoliciesControlPolicy) SetPolicyName(v string) *ListControlPoliciesResponseBodyControlPoliciesControlPolicy {
	s.PolicyName = &v
	return s
}

func (s *ListControlPoliciesResponseBodyControlPoliciesControlPolicy) SetPolicyType(v string) *ListControlPoliciesResponseBodyControlPoliciesControlPolicy {
	s.PolicyType = &v
	return s
}

func (s *ListControlPoliciesResponseBodyControlPoliciesControlPolicy) SetUpdateDate(v string) *ListControlPoliciesResponseBodyControlPoliciesControlPolicy {
	s.UpdateDate = &v
	return s
}

type ListControlPoliciesResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListControlPoliciesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListControlPoliciesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListControlPoliciesResponse) GoString() string {
	return s.String()
}

func (s *ListControlPoliciesResponse) SetHeaders(v map[string]*string) *ListControlPoliciesResponse {
	s.Headers = v
	return s
}

func (s *ListControlPoliciesResponse) SetStatusCode(v int32) *ListControlPoliciesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListControlPoliciesResponse) SetBody(v *ListControlPoliciesResponseBody) *ListControlPoliciesResponse {
	s.Body = v
	return s
}

type ListControlPolicyAttachmentsForTargetRequest struct {
	// The language in which you want to return the descriptions of the access control policies. Valid values:
	//
	// 	- zh-CN (default value): Chinese
	//
	// 	- en: English
	//
	// 	- ja: Japanese
	//
	// > This parameter is valid only for system access control policies.
	//
	// example:
	//
	// zh-CN
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// The ID of the object whose access control policies you want to query. Access control policies can be attached to the following objects:
	//
	// 	- Root folder
	//
	// 	- Subfolders of the Root folder
	//
	// 	- Members
	//
	// This parameter is required.
	//
	// example:
	//
	// fd-ZDNPiT****
	TargetId *string `json:"TargetId,omitempty" xml:"TargetId,omitempty"`
}

func (s ListControlPolicyAttachmentsForTargetRequest) String() string {
	return tea.Prettify(s)
}

func (s ListControlPolicyAttachmentsForTargetRequest) GoString() string {
	return s.String()
}

func (s *ListControlPolicyAttachmentsForTargetRequest) SetLanguage(v string) *ListControlPolicyAttachmentsForTargetRequest {
	s.Language = &v
	return s
}

func (s *ListControlPolicyAttachmentsForTargetRequest) SetTargetId(v string) *ListControlPolicyAttachmentsForTargetRequest {
	s.TargetId = &v
	return s
}

type ListControlPolicyAttachmentsForTargetResponseBody struct {
	// The information about the attached access control policies.
	ControlPolicyAttachments *ListControlPolicyAttachmentsForTargetResponseBodyControlPolicyAttachments `json:"ControlPolicyAttachments,omitempty" xml:"ControlPolicyAttachments,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// C276B600-7B7A-49E8-938C-E16CFA955A82
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListControlPolicyAttachmentsForTargetResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListControlPolicyAttachmentsForTargetResponseBody) GoString() string {
	return s.String()
}

func (s *ListControlPolicyAttachmentsForTargetResponseBody) SetControlPolicyAttachments(v *ListControlPolicyAttachmentsForTargetResponseBodyControlPolicyAttachments) *ListControlPolicyAttachmentsForTargetResponseBody {
	s.ControlPolicyAttachments = v
	return s
}

func (s *ListControlPolicyAttachmentsForTargetResponseBody) SetRequestId(v string) *ListControlPolicyAttachmentsForTargetResponseBody {
	s.RequestId = &v
	return s
}

type ListControlPolicyAttachmentsForTargetResponseBodyControlPolicyAttachments struct {
	ControlPolicyAttachment []*ListControlPolicyAttachmentsForTargetResponseBodyControlPolicyAttachmentsControlPolicyAttachment `json:"ControlPolicyAttachment,omitempty" xml:"ControlPolicyAttachment,omitempty" type:"Repeated"`
}

func (s ListControlPolicyAttachmentsForTargetResponseBodyControlPolicyAttachments) String() string {
	return tea.Prettify(s)
}

func (s ListControlPolicyAttachmentsForTargetResponseBodyControlPolicyAttachments) GoString() string {
	return s.String()
}

func (s *ListControlPolicyAttachmentsForTargetResponseBodyControlPolicyAttachments) SetControlPolicyAttachment(v []*ListControlPolicyAttachmentsForTargetResponseBodyControlPolicyAttachmentsControlPolicyAttachment) *ListControlPolicyAttachmentsForTargetResponseBodyControlPolicyAttachments {
	s.ControlPolicyAttachment = v
	return s
}

type ListControlPolicyAttachmentsForTargetResponseBodyControlPolicyAttachmentsControlPolicyAttachment struct {
	// The time when the access control policy was attached.
	//
	// example:
	//
	// 2021-03-19T02:56:24Z
	AttachDate *string `json:"AttachDate,omitempty" xml:"AttachDate,omitempty"`
	// The description of the access control policy.
	//
	// example:
	//
	// ExampleControlPolicy
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The effective scope of the access control policy. Valid values:
	//
	// 	- All: The access control policy is in effect for Alibaba Cloud accounts, RAM users, and RAM roles.
	//
	// 	- RAM: The access control policy is in effect only for RAM users and RAM roles.
	//
	// example:
	//
	// RAM
	EffectScope *string `json:"EffectScope,omitempty" xml:"EffectScope,omitempty"`
	// The ID of the access control policy.
	//
	// example:
	//
	// cp-jExXAqIYkwHN****
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// The name of the access control policy.
	//
	// example:
	//
	// ExampleControlPolicy
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	// The type of the access control policy. Valid values:
	//
	// 	- System: system access control policy
	//
	// 	- Custom: custom access control policy
	//
	// example:
	//
	// Custom
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s ListControlPolicyAttachmentsForTargetResponseBodyControlPolicyAttachmentsControlPolicyAttachment) String() string {
	return tea.Prettify(s)
}

func (s ListControlPolicyAttachmentsForTargetResponseBodyControlPolicyAttachmentsControlPolicyAttachment) GoString() string {
	return s.String()
}

func (s *ListControlPolicyAttachmentsForTargetResponseBodyControlPolicyAttachmentsControlPolicyAttachment) SetAttachDate(v string) *ListControlPolicyAttachmentsForTargetResponseBodyControlPolicyAttachmentsControlPolicyAttachment {
	s.AttachDate = &v
	return s
}

func (s *ListControlPolicyAttachmentsForTargetResponseBodyControlPolicyAttachmentsControlPolicyAttachment) SetDescription(v string) *ListControlPolicyAttachmentsForTargetResponseBodyControlPolicyAttachmentsControlPolicyAttachment {
	s.Description = &v
	return s
}

func (s *ListControlPolicyAttachmentsForTargetResponseBodyControlPolicyAttachmentsControlPolicyAttachment) SetEffectScope(v string) *ListControlPolicyAttachmentsForTargetResponseBodyControlPolicyAttachmentsControlPolicyAttachment {
	s.EffectScope = &v
	return s
}

func (s *ListControlPolicyAttachmentsForTargetResponseBodyControlPolicyAttachmentsControlPolicyAttachment) SetPolicyId(v string) *ListControlPolicyAttachmentsForTargetResponseBodyControlPolicyAttachmentsControlPolicyAttachment {
	s.PolicyId = &v
	return s
}

func (s *ListControlPolicyAttachmentsForTargetResponseBodyControlPolicyAttachmentsControlPolicyAttachment) SetPolicyName(v string) *ListControlPolicyAttachmentsForTargetResponseBodyControlPolicyAttachmentsControlPolicyAttachment {
	s.PolicyName = &v
	return s
}

func (s *ListControlPolicyAttachmentsForTargetResponseBodyControlPolicyAttachmentsControlPolicyAttachment) SetPolicyType(v string) *ListControlPolicyAttachmentsForTargetResponseBodyControlPolicyAttachmentsControlPolicyAttachment {
	s.PolicyType = &v
	return s
}

type ListControlPolicyAttachmentsForTargetResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListControlPolicyAttachmentsForTargetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListControlPolicyAttachmentsForTargetResponse) String() string {
	return tea.Prettify(s)
}

func (s ListControlPolicyAttachmentsForTargetResponse) GoString() string {
	return s.String()
}

func (s *ListControlPolicyAttachmentsForTargetResponse) SetHeaders(v map[string]*string) *ListControlPolicyAttachmentsForTargetResponse {
	s.Headers = v
	return s
}

func (s *ListControlPolicyAttachmentsForTargetResponse) SetStatusCode(v int32) *ListControlPolicyAttachmentsForTargetResponse {
	s.StatusCode = &v
	return s
}

func (s *ListControlPolicyAttachmentsForTargetResponse) SetBody(v *ListControlPolicyAttachmentsForTargetResponseBody) *ListControlPolicyAttachmentsForTargetResponse {
	s.Body = v
	return s
}

type ListDelegatedAdministratorsRequest struct {
	// The number of the page to return.
	//
	// Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page.
	//
	// Valid values: 1 to 100. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The identifier of the trusted service.
	//
	// For more information, see the `Trusted service identifier` column in [Supported trusted services](https://help.aliyun.com/document_detail/208133.html).
	//
	// example:
	//
	// cloudfw.aliyuncs.com
	ServicePrincipal *string `json:"ServicePrincipal,omitempty" xml:"ServicePrincipal,omitempty"`
}

func (s ListDelegatedAdministratorsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDelegatedAdministratorsRequest) GoString() string {
	return s.String()
}

func (s *ListDelegatedAdministratorsRequest) SetPageNumber(v int64) *ListDelegatedAdministratorsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListDelegatedAdministratorsRequest) SetPageSize(v int64) *ListDelegatedAdministratorsRequest {
	s.PageSize = &v
	return s
}

func (s *ListDelegatedAdministratorsRequest) SetServicePrincipal(v string) *ListDelegatedAdministratorsRequest {
	s.ServicePrincipal = &v
	return s
}

type ListDelegatedAdministratorsResponseBody struct {
	// The information about the delegated administrator accounts.
	Accounts *ListDelegatedAdministratorsResponseBodyAccounts `json:"Accounts,omitempty" xml:"Accounts,omitempty" type:"Struct"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 009C3A02-7D4B-416C-9CE7-548C91508F1E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 2
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListDelegatedAdministratorsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDelegatedAdministratorsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDelegatedAdministratorsResponseBody) SetAccounts(v *ListDelegatedAdministratorsResponseBodyAccounts) *ListDelegatedAdministratorsResponseBody {
	s.Accounts = v
	return s
}

func (s *ListDelegatedAdministratorsResponseBody) SetPageNumber(v int64) *ListDelegatedAdministratorsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListDelegatedAdministratorsResponseBody) SetPageSize(v int64) *ListDelegatedAdministratorsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListDelegatedAdministratorsResponseBody) SetRequestId(v string) *ListDelegatedAdministratorsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDelegatedAdministratorsResponseBody) SetTotalCount(v int64) *ListDelegatedAdministratorsResponseBody {
	s.TotalCount = &v
	return s
}

type ListDelegatedAdministratorsResponseBodyAccounts struct {
	Account []*ListDelegatedAdministratorsResponseBodyAccountsAccount `json:"Account,omitempty" xml:"Account,omitempty" type:"Repeated"`
}

func (s ListDelegatedAdministratorsResponseBodyAccounts) String() string {
	return tea.Prettify(s)
}

func (s ListDelegatedAdministratorsResponseBodyAccounts) GoString() string {
	return s.String()
}

func (s *ListDelegatedAdministratorsResponseBodyAccounts) SetAccount(v []*ListDelegatedAdministratorsResponseBodyAccountsAccount) *ListDelegatedAdministratorsResponseBodyAccounts {
	s.Account = v
	return s
}

type ListDelegatedAdministratorsResponseBodyAccountsAccount struct {
	// The Alibaba Cloud account ID of the member.
	//
	// example:
	//
	// 138660628348****
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The time when the member was specified as a delegated administrator account.
	//
	// example:
	//
	// 1616652684164
	DelegationEnabledTime *string `json:"DelegationEnabledTime,omitempty" xml:"DelegationEnabledTime,omitempty"`
	// The display name of the member.
	//
	// example:
	//
	// abc
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The way in which the member joins the resource directory. Valid values:
	//
	// 	- invited: The member is invited to join the resource directory.
	//
	// 	- created: The member is directly created in the resource directory.
	//
	// example:
	//
	// created
	JoinMethod *string `json:"JoinMethod,omitempty" xml:"JoinMethod,omitempty"`
	// The identifier of the trusted service.
	//
	// example:
	//
	// cloudfw.aliyuncs.com
	ServicePrincipal *string `json:"ServicePrincipal,omitempty" xml:"ServicePrincipal,omitempty"`
}

func (s ListDelegatedAdministratorsResponseBodyAccountsAccount) String() string {
	return tea.Prettify(s)
}

func (s ListDelegatedAdministratorsResponseBodyAccountsAccount) GoString() string {
	return s.String()
}

func (s *ListDelegatedAdministratorsResponseBodyAccountsAccount) SetAccountId(v string) *ListDelegatedAdministratorsResponseBodyAccountsAccount {
	s.AccountId = &v
	return s
}

func (s *ListDelegatedAdministratorsResponseBodyAccountsAccount) SetDelegationEnabledTime(v string) *ListDelegatedAdministratorsResponseBodyAccountsAccount {
	s.DelegationEnabledTime = &v
	return s
}

func (s *ListDelegatedAdministratorsResponseBodyAccountsAccount) SetDisplayName(v string) *ListDelegatedAdministratorsResponseBodyAccountsAccount {
	s.DisplayName = &v
	return s
}

func (s *ListDelegatedAdministratorsResponseBodyAccountsAccount) SetJoinMethod(v string) *ListDelegatedAdministratorsResponseBodyAccountsAccount {
	s.JoinMethod = &v
	return s
}

func (s *ListDelegatedAdministratorsResponseBodyAccountsAccount) SetServicePrincipal(v string) *ListDelegatedAdministratorsResponseBodyAccountsAccount {
	s.ServicePrincipal = &v
	return s
}

type ListDelegatedAdministratorsResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDelegatedAdministratorsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDelegatedAdministratorsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDelegatedAdministratorsResponse) GoString() string {
	return s.String()
}

func (s *ListDelegatedAdministratorsResponse) SetHeaders(v map[string]*string) *ListDelegatedAdministratorsResponse {
	s.Headers = v
	return s
}

func (s *ListDelegatedAdministratorsResponse) SetStatusCode(v int32) *ListDelegatedAdministratorsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDelegatedAdministratorsResponse) SetBody(v *ListDelegatedAdministratorsResponseBody) *ListDelegatedAdministratorsResponse {
	s.Body = v
	return s
}

type ListDelegatedServicesForAccountRequest struct {
	// The Alibaba Cloud account ID of the member.
	//
	// This parameter is required.
	//
	// example:
	//
	// 138660628348****
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
}

func (s ListDelegatedServicesForAccountRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDelegatedServicesForAccountRequest) GoString() string {
	return s.String()
}

func (s *ListDelegatedServicesForAccountRequest) SetAccountId(v string) *ListDelegatedServicesForAccountRequest {
	s.AccountId = &v
	return s
}

type ListDelegatedServicesForAccountResponseBody struct {
	// The information about the trusted services.
	//
	// > If the value of this parameter is empty, the member is not specified as a delegated administrator account.
	DelegatedServices *ListDelegatedServicesForAccountResponseBodyDelegatedServices `json:"DelegatedServices,omitempty" xml:"DelegatedServices,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// D9C03B94-9396-4794-A74B-13DC437556A6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListDelegatedServicesForAccountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDelegatedServicesForAccountResponseBody) GoString() string {
	return s.String()
}

func (s *ListDelegatedServicesForAccountResponseBody) SetDelegatedServices(v *ListDelegatedServicesForAccountResponseBodyDelegatedServices) *ListDelegatedServicesForAccountResponseBody {
	s.DelegatedServices = v
	return s
}

func (s *ListDelegatedServicesForAccountResponseBody) SetRequestId(v string) *ListDelegatedServicesForAccountResponseBody {
	s.RequestId = &v
	return s
}

type ListDelegatedServicesForAccountResponseBodyDelegatedServices struct {
	DelegatedService []*ListDelegatedServicesForAccountResponseBodyDelegatedServicesDelegatedService `json:"DelegatedService,omitempty" xml:"DelegatedService,omitempty" type:"Repeated"`
}

func (s ListDelegatedServicesForAccountResponseBodyDelegatedServices) String() string {
	return tea.Prettify(s)
}

func (s ListDelegatedServicesForAccountResponseBodyDelegatedServices) GoString() string {
	return s.String()
}

func (s *ListDelegatedServicesForAccountResponseBodyDelegatedServices) SetDelegatedService(v []*ListDelegatedServicesForAccountResponseBodyDelegatedServicesDelegatedService) *ListDelegatedServicesForAccountResponseBodyDelegatedServices {
	s.DelegatedService = v
	return s
}

type ListDelegatedServicesForAccountResponseBodyDelegatedServicesDelegatedService struct {
	// The time when the member was specified as a delegated administrator account.
	//
	// example:
	//
	// 1616652684164
	DelegationEnabledTime *string `json:"DelegationEnabledTime,omitempty" xml:"DelegationEnabledTime,omitempty"`
	// The identifier of the trusted service.
	//
	// example:
	//
	// cloudfw.aliyuncs.com
	ServicePrincipal *string `json:"ServicePrincipal,omitempty" xml:"ServicePrincipal,omitempty"`
	// The status of the trusted service. Valid values:
	//
	// 	- ENABLED: enabled
	//
	// 	- DISABLED: disabled
	//
	// example:
	//
	// ENABLED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListDelegatedServicesForAccountResponseBodyDelegatedServicesDelegatedService) String() string {
	return tea.Prettify(s)
}

func (s ListDelegatedServicesForAccountResponseBodyDelegatedServicesDelegatedService) GoString() string {
	return s.String()
}

func (s *ListDelegatedServicesForAccountResponseBodyDelegatedServicesDelegatedService) SetDelegationEnabledTime(v string) *ListDelegatedServicesForAccountResponseBodyDelegatedServicesDelegatedService {
	s.DelegationEnabledTime = &v
	return s
}

func (s *ListDelegatedServicesForAccountResponseBodyDelegatedServicesDelegatedService) SetServicePrincipal(v string) *ListDelegatedServicesForAccountResponseBodyDelegatedServicesDelegatedService {
	s.ServicePrincipal = &v
	return s
}

func (s *ListDelegatedServicesForAccountResponseBodyDelegatedServicesDelegatedService) SetStatus(v string) *ListDelegatedServicesForAccountResponseBodyDelegatedServicesDelegatedService {
	s.Status = &v
	return s
}

type ListDelegatedServicesForAccountResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDelegatedServicesForAccountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDelegatedServicesForAccountResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDelegatedServicesForAccountResponse) GoString() string {
	return s.String()
}

func (s *ListDelegatedServicesForAccountResponse) SetHeaders(v map[string]*string) *ListDelegatedServicesForAccountResponse {
	s.Headers = v
	return s
}

func (s *ListDelegatedServicesForAccountResponse) SetStatusCode(v int32) *ListDelegatedServicesForAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDelegatedServicesForAccountResponse) SetBody(v *ListDelegatedServicesForAccountResponseBody) *ListDelegatedServicesForAccountResponse {
	s.Body = v
	return s
}

type ListFoldersForParentRequest struct {
	// The number of the page to return.
	//
	// Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page.
	//
	// Valid values: 1 to 100. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the parent folder.
	//
	// If you leave this parameter empty, the information of the first-level subfolders of the Root folder is queried.
	//
	// example:
	//
	// r-b1****
	ParentFolderId *string `json:"ParentFolderId,omitempty" xml:"ParentFolderId,omitempty"`
	// The keyword used for the query, such as a folder name.
	//
	// Fuzzy match is supported.
	//
	// example:
	//
	// rdFolder
	QueryKeyword *string `json:"QueryKeyword,omitempty" xml:"QueryKeyword,omitempty"`
}

func (s ListFoldersForParentRequest) String() string {
	return tea.Prettify(s)
}

func (s ListFoldersForParentRequest) GoString() string {
	return s.String()
}

func (s *ListFoldersForParentRequest) SetPageNumber(v int32) *ListFoldersForParentRequest {
	s.PageNumber = &v
	return s
}

func (s *ListFoldersForParentRequest) SetPageSize(v int32) *ListFoldersForParentRequest {
	s.PageSize = &v
	return s
}

func (s *ListFoldersForParentRequest) SetParentFolderId(v string) *ListFoldersForParentRequest {
	s.ParentFolderId = &v
	return s
}

func (s *ListFoldersForParentRequest) SetQueryKeyword(v string) *ListFoldersForParentRequest {
	s.QueryKeyword = &v
	return s
}

type ListFoldersForParentResponseBody struct {
	// The information of the folders.
	Folders *ListFoldersForParentResponseBodyFolders `json:"Folders,omitempty" xml:"Folders,omitempty" type:"Struct"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 5
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 7B8A4E7D-6CFF-471D-84DF-195A7A241ECB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 2
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListFoldersForParentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListFoldersForParentResponseBody) GoString() string {
	return s.String()
}

func (s *ListFoldersForParentResponseBody) SetFolders(v *ListFoldersForParentResponseBodyFolders) *ListFoldersForParentResponseBody {
	s.Folders = v
	return s
}

func (s *ListFoldersForParentResponseBody) SetPageNumber(v int32) *ListFoldersForParentResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListFoldersForParentResponseBody) SetPageSize(v int32) *ListFoldersForParentResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListFoldersForParentResponseBody) SetRequestId(v string) *ListFoldersForParentResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListFoldersForParentResponseBody) SetTotalCount(v int32) *ListFoldersForParentResponseBody {
	s.TotalCount = &v
	return s
}

type ListFoldersForParentResponseBodyFolders struct {
	Folder []*ListFoldersForParentResponseBodyFoldersFolder `json:"Folder,omitempty" xml:"Folder,omitempty" type:"Repeated"`
}

func (s ListFoldersForParentResponseBodyFolders) String() string {
	return tea.Prettify(s)
}

func (s ListFoldersForParentResponseBodyFolders) GoString() string {
	return s.String()
}

func (s *ListFoldersForParentResponseBodyFolders) SetFolder(v []*ListFoldersForParentResponseBodyFoldersFolder) *ListFoldersForParentResponseBodyFolders {
	s.Folder = v
	return s
}

type ListFoldersForParentResponseBodyFoldersFolder struct {
	// The time when the folder was created.
	//
	// example:
	//
	// 2015-01-23T12:33:18Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the folder.
	//
	// example:
	//
	// rd-evic31****
	FolderId *string `json:"FolderId,omitempty" xml:"FolderId,omitempty"`
	// The name of the folder.
	//
	// example:
	//
	// project-1
	FolderName *string `json:"FolderName,omitempty" xml:"FolderName,omitempty"`
}

func (s ListFoldersForParentResponseBodyFoldersFolder) String() string {
	return tea.Prettify(s)
}

func (s ListFoldersForParentResponseBodyFoldersFolder) GoString() string {
	return s.String()
}

func (s *ListFoldersForParentResponseBodyFoldersFolder) SetCreateTime(v string) *ListFoldersForParentResponseBodyFoldersFolder {
	s.CreateTime = &v
	return s
}

func (s *ListFoldersForParentResponseBodyFoldersFolder) SetFolderId(v string) *ListFoldersForParentResponseBodyFoldersFolder {
	s.FolderId = &v
	return s
}

func (s *ListFoldersForParentResponseBodyFoldersFolder) SetFolderName(v string) *ListFoldersForParentResponseBodyFoldersFolder {
	s.FolderName = &v
	return s
}

type ListFoldersForParentResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListFoldersForParentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListFoldersForParentResponse) String() string {
	return tea.Prettify(s)
}

func (s ListFoldersForParentResponse) GoString() string {
	return s.String()
}

func (s *ListFoldersForParentResponse) SetHeaders(v map[string]*string) *ListFoldersForParentResponse {
	s.Headers = v
	return s
}

func (s *ListFoldersForParentResponse) SetStatusCode(v int32) *ListFoldersForParentResponse {
	s.StatusCode = &v
	return s
}

func (s *ListFoldersForParentResponse) SetBody(v *ListFoldersForParentResponseBody) *ListFoldersForParentResponse {
	s.Body = v
	return s
}

type ListHandshakesForAccountRequest struct {
	// The number of the page to return.
	//
	// Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page.
	//
	// Valid values: 1 to 100. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListHandshakesForAccountRequest) String() string {
	return tea.Prettify(s)
}

func (s ListHandshakesForAccountRequest) GoString() string {
	return s.String()
}

func (s *ListHandshakesForAccountRequest) SetPageNumber(v int32) *ListHandshakesForAccountRequest {
	s.PageNumber = &v
	return s
}

func (s *ListHandshakesForAccountRequest) SetPageSize(v int32) *ListHandshakesForAccountRequest {
	s.PageSize = &v
	return s
}

type ListHandshakesForAccountResponseBody struct {
	// The information of the invitations.
	Handshakes *ListHandshakesForAccountResponseBodyHandshakes `json:"Handshakes,omitempty" xml:"Handshakes,omitempty" type:"Struct"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 9B34724D-54B0-4A51-B34D-4512372FE1BE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of invitations.
	//
	// example:
	//
	// 2
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListHandshakesForAccountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListHandshakesForAccountResponseBody) GoString() string {
	return s.String()
}

func (s *ListHandshakesForAccountResponseBody) SetHandshakes(v *ListHandshakesForAccountResponseBodyHandshakes) *ListHandshakesForAccountResponseBody {
	s.Handshakes = v
	return s
}

func (s *ListHandshakesForAccountResponseBody) SetPageNumber(v int32) *ListHandshakesForAccountResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListHandshakesForAccountResponseBody) SetPageSize(v int32) *ListHandshakesForAccountResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListHandshakesForAccountResponseBody) SetRequestId(v string) *ListHandshakesForAccountResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListHandshakesForAccountResponseBody) SetTotalCount(v int32) *ListHandshakesForAccountResponseBody {
	s.TotalCount = &v
	return s
}

type ListHandshakesForAccountResponseBodyHandshakes struct {
	Handshake []*ListHandshakesForAccountResponseBodyHandshakesHandshake `json:"Handshake,omitempty" xml:"Handshake,omitempty" type:"Repeated"`
}

func (s ListHandshakesForAccountResponseBodyHandshakes) String() string {
	return tea.Prettify(s)
}

func (s ListHandshakesForAccountResponseBodyHandshakes) GoString() string {
	return s.String()
}

func (s *ListHandshakesForAccountResponseBodyHandshakes) SetHandshake(v []*ListHandshakesForAccountResponseBodyHandshakesHandshake) *ListHandshakesForAccountResponseBodyHandshakes {
	s.Handshake = v
	return s
}

type ListHandshakesForAccountResponseBodyHandshakesHandshake struct {
	// The time when the invitation was created. The time is displayed in UTC.
	//
	// example:
	//
	// 2018-08-10T09:55:41Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The time when the invitation expires. The time is displayed in UTC.
	//
	// example:
	//
	// 2018-08-24T09:55:41Z
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The ID of the invitation.
	//
	// example:
	//
	// h-4N57QZzCTtES****
	HandshakeId *string `json:"HandshakeId,omitempty" xml:"HandshakeId,omitempty"`
	// The ID of the management account of the resource directory.
	//
	// example:
	//
	// 172841235500****
	MasterAccountId *string `json:"MasterAccountId,omitempty" xml:"MasterAccountId,omitempty"`
	// The name of the management account of the resource directory.
	//
	// example:
	//
	// CompanyA
	MasterAccountName *string `json:"MasterAccountName,omitempty" xml:"MasterAccountName,omitempty"`
	// The time when the invitation was modified. The time is displayed in UTC.
	//
	// example:
	//
	// 2018-08-10T09:55:41Z
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The description of the invitation.
	//
	// example:
	//
	// Welcome
	Note *string `json:"Note,omitempty" xml:"Note,omitempty"`
	// The ID of the resource directory.
	//
	// example:
	//
	// rd-abcdef****
	ResourceDirectoryId *string `json:"ResourceDirectoryId,omitempty" xml:"ResourceDirectoryId,omitempty"`
	// The status of the invitation. Valid values:
	//
	// 	- Pending: The invitation is waiting for confirmation.
	//
	// 	- Accepted: The invitation is accepted.
	//
	// 	- Cancelled: The invitation is canceled.
	//
	// 	- Declined: The invitation is rejected.
	//
	// 	- Expired: The invitation expires.
	//
	// example:
	//
	// Pending
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID or logon email address of the invited Alibaba Cloud account.
	//
	// example:
	//
	// someone@example.com
	TargetEntity *string `json:"TargetEntity,omitempty" xml:"TargetEntity,omitempty"`
	// The type of the invited Alibaba Cloud account. Valid values:
	//
	// 	- Account: indicates the ID of the account.
	//
	// 	- Email: indicates the logon email address of the account.
	//
	// example:
	//
	// Email
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
}

func (s ListHandshakesForAccountResponseBodyHandshakesHandshake) String() string {
	return tea.Prettify(s)
}

func (s ListHandshakesForAccountResponseBodyHandshakesHandshake) GoString() string {
	return s.String()
}

func (s *ListHandshakesForAccountResponseBodyHandshakesHandshake) SetCreateTime(v string) *ListHandshakesForAccountResponseBodyHandshakesHandshake {
	s.CreateTime = &v
	return s
}

func (s *ListHandshakesForAccountResponseBodyHandshakesHandshake) SetExpireTime(v string) *ListHandshakesForAccountResponseBodyHandshakesHandshake {
	s.ExpireTime = &v
	return s
}

func (s *ListHandshakesForAccountResponseBodyHandshakesHandshake) SetHandshakeId(v string) *ListHandshakesForAccountResponseBodyHandshakesHandshake {
	s.HandshakeId = &v
	return s
}

func (s *ListHandshakesForAccountResponseBodyHandshakesHandshake) SetMasterAccountId(v string) *ListHandshakesForAccountResponseBodyHandshakesHandshake {
	s.MasterAccountId = &v
	return s
}

func (s *ListHandshakesForAccountResponseBodyHandshakesHandshake) SetMasterAccountName(v string) *ListHandshakesForAccountResponseBodyHandshakesHandshake {
	s.MasterAccountName = &v
	return s
}

func (s *ListHandshakesForAccountResponseBodyHandshakesHandshake) SetModifyTime(v string) *ListHandshakesForAccountResponseBodyHandshakesHandshake {
	s.ModifyTime = &v
	return s
}

func (s *ListHandshakesForAccountResponseBodyHandshakesHandshake) SetNote(v string) *ListHandshakesForAccountResponseBodyHandshakesHandshake {
	s.Note = &v
	return s
}

func (s *ListHandshakesForAccountResponseBodyHandshakesHandshake) SetResourceDirectoryId(v string) *ListHandshakesForAccountResponseBodyHandshakesHandshake {
	s.ResourceDirectoryId = &v
	return s
}

func (s *ListHandshakesForAccountResponseBodyHandshakesHandshake) SetStatus(v string) *ListHandshakesForAccountResponseBodyHandshakesHandshake {
	s.Status = &v
	return s
}

func (s *ListHandshakesForAccountResponseBodyHandshakesHandshake) SetTargetEntity(v string) *ListHandshakesForAccountResponseBodyHandshakesHandshake {
	s.TargetEntity = &v
	return s
}

func (s *ListHandshakesForAccountResponseBodyHandshakesHandshake) SetTargetType(v string) *ListHandshakesForAccountResponseBodyHandshakesHandshake {
	s.TargetType = &v
	return s
}

type ListHandshakesForAccountResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListHandshakesForAccountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListHandshakesForAccountResponse) String() string {
	return tea.Prettify(s)
}

func (s ListHandshakesForAccountResponse) GoString() string {
	return s.String()
}

func (s *ListHandshakesForAccountResponse) SetHeaders(v map[string]*string) *ListHandshakesForAccountResponse {
	s.Headers = v
	return s
}

func (s *ListHandshakesForAccountResponse) SetStatusCode(v int32) *ListHandshakesForAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *ListHandshakesForAccountResponse) SetBody(v *ListHandshakesForAccountResponseBody) *ListHandshakesForAccountResponse {
	s.Body = v
	return s
}

type ListHandshakesForResourceDirectoryRequest struct {
	// The number of the page to return.
	//
	// Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page.
	//
	// Valid values: 1 to 100. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListHandshakesForResourceDirectoryRequest) String() string {
	return tea.Prettify(s)
}

func (s ListHandshakesForResourceDirectoryRequest) GoString() string {
	return s.String()
}

func (s *ListHandshakesForResourceDirectoryRequest) SetPageNumber(v int32) *ListHandshakesForResourceDirectoryRequest {
	s.PageNumber = &v
	return s
}

func (s *ListHandshakesForResourceDirectoryRequest) SetPageSize(v int32) *ListHandshakesForResourceDirectoryRequest {
	s.PageSize = &v
	return s
}

type ListHandshakesForResourceDirectoryResponseBody struct {
	// The information of the invitations.
	Handshakes *ListHandshakesForResourceDirectoryResponseBodyHandshakes `json:"Handshakes,omitempty" xml:"Handshakes,omitempty" type:"Struct"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 9B34724D-54B0-4A51-B34D-4512372FE1BE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 2
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListHandshakesForResourceDirectoryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListHandshakesForResourceDirectoryResponseBody) GoString() string {
	return s.String()
}

func (s *ListHandshakesForResourceDirectoryResponseBody) SetHandshakes(v *ListHandshakesForResourceDirectoryResponseBodyHandshakes) *ListHandshakesForResourceDirectoryResponseBody {
	s.Handshakes = v
	return s
}

func (s *ListHandshakesForResourceDirectoryResponseBody) SetPageNumber(v int32) *ListHandshakesForResourceDirectoryResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListHandshakesForResourceDirectoryResponseBody) SetPageSize(v int32) *ListHandshakesForResourceDirectoryResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListHandshakesForResourceDirectoryResponseBody) SetRequestId(v string) *ListHandshakesForResourceDirectoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListHandshakesForResourceDirectoryResponseBody) SetTotalCount(v int32) *ListHandshakesForResourceDirectoryResponseBody {
	s.TotalCount = &v
	return s
}

type ListHandshakesForResourceDirectoryResponseBodyHandshakes struct {
	Handshake []*ListHandshakesForResourceDirectoryResponseBodyHandshakesHandshake `json:"Handshake,omitempty" xml:"Handshake,omitempty" type:"Repeated"`
}

func (s ListHandshakesForResourceDirectoryResponseBodyHandshakes) String() string {
	return tea.Prettify(s)
}

func (s ListHandshakesForResourceDirectoryResponseBodyHandshakes) GoString() string {
	return s.String()
}

func (s *ListHandshakesForResourceDirectoryResponseBodyHandshakes) SetHandshake(v []*ListHandshakesForResourceDirectoryResponseBodyHandshakesHandshake) *ListHandshakesForResourceDirectoryResponseBodyHandshakes {
	s.Handshake = v
	return s
}

type ListHandshakesForResourceDirectoryResponseBodyHandshakesHandshake struct {
	// The time when the invitation was created. The time is displayed in UTC.
	//
	// example:
	//
	// 2018-08-10T09:55:41Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The time when the invitation expires. The time is displayed in UTC.
	//
	// example:
	//
	// 2018-08-24T09:55:41Z
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The ID of the invitation.
	//
	// example:
	//
	// h-ycm4rp****
	HandshakeId *string `json:"HandshakeId,omitempty" xml:"HandshakeId,omitempty"`
	// The ID of the management account of the resource directory.
	//
	// example:
	//
	// 172841235500****
	MasterAccountId *string `json:"MasterAccountId,omitempty" xml:"MasterAccountId,omitempty"`
	// The name of the management account of the resource directory.
	//
	// example:
	//
	// Alice
	MasterAccountName *string `json:"MasterAccountName,omitempty" xml:"MasterAccountName,omitempty"`
	// The time when the invitation was modified. The time is displayed in UTC.
	//
	// example:
	//
	// 2018-08-10T09:55:41Z
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The description of the invitation.
	//
	// example:
	//
	// Welcome
	Note *string `json:"Note,omitempty" xml:"Note,omitempty"`
	// The ID of the resource directory.
	//
	// example:
	//
	// rd-abcdef****
	ResourceDirectoryId *string `json:"ResourceDirectoryId,omitempty" xml:"ResourceDirectoryId,omitempty"`
	// The status of the invitation. Valid values:
	//
	// 	- Pending: The invitation is waiting for confirmation.
	//
	// 	- Accepted: The invitation is accepted.
	//
	// 	- Cancelled: The invitation is canceled.
	//
	// 	- Declined: The invitation is rejected.
	//
	// 	- Expired: The invitation expires.
	//
	// example:
	//
	// Pending
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID or logon email address of the invited account.
	//
	// example:
	//
	// someone@example.com
	TargetEntity *string `json:"TargetEntity,omitempty" xml:"TargetEntity,omitempty"`
	// The type of the invited account. Valid values:
	//
	// 	- Account: indicates the ID of the account.
	//
	// 	- Email: indicates the logon email address of the account.
	//
	// example:
	//
	// Email
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
}

func (s ListHandshakesForResourceDirectoryResponseBodyHandshakesHandshake) String() string {
	return tea.Prettify(s)
}

func (s ListHandshakesForResourceDirectoryResponseBodyHandshakesHandshake) GoString() string {
	return s.String()
}

func (s *ListHandshakesForResourceDirectoryResponseBodyHandshakesHandshake) SetCreateTime(v string) *ListHandshakesForResourceDirectoryResponseBodyHandshakesHandshake {
	s.CreateTime = &v
	return s
}

func (s *ListHandshakesForResourceDirectoryResponseBodyHandshakesHandshake) SetExpireTime(v string) *ListHandshakesForResourceDirectoryResponseBodyHandshakesHandshake {
	s.ExpireTime = &v
	return s
}

func (s *ListHandshakesForResourceDirectoryResponseBodyHandshakesHandshake) SetHandshakeId(v string) *ListHandshakesForResourceDirectoryResponseBodyHandshakesHandshake {
	s.HandshakeId = &v
	return s
}

func (s *ListHandshakesForResourceDirectoryResponseBodyHandshakesHandshake) SetMasterAccountId(v string) *ListHandshakesForResourceDirectoryResponseBodyHandshakesHandshake {
	s.MasterAccountId = &v
	return s
}

func (s *ListHandshakesForResourceDirectoryResponseBodyHandshakesHandshake) SetMasterAccountName(v string) *ListHandshakesForResourceDirectoryResponseBodyHandshakesHandshake {
	s.MasterAccountName = &v
	return s
}

func (s *ListHandshakesForResourceDirectoryResponseBodyHandshakesHandshake) SetModifyTime(v string) *ListHandshakesForResourceDirectoryResponseBodyHandshakesHandshake {
	s.ModifyTime = &v
	return s
}

func (s *ListHandshakesForResourceDirectoryResponseBodyHandshakesHandshake) SetNote(v string) *ListHandshakesForResourceDirectoryResponseBodyHandshakesHandshake {
	s.Note = &v
	return s
}

func (s *ListHandshakesForResourceDirectoryResponseBodyHandshakesHandshake) SetResourceDirectoryId(v string) *ListHandshakesForResourceDirectoryResponseBodyHandshakesHandshake {
	s.ResourceDirectoryId = &v
	return s
}

func (s *ListHandshakesForResourceDirectoryResponseBodyHandshakesHandshake) SetStatus(v string) *ListHandshakesForResourceDirectoryResponseBodyHandshakesHandshake {
	s.Status = &v
	return s
}

func (s *ListHandshakesForResourceDirectoryResponseBodyHandshakesHandshake) SetTargetEntity(v string) *ListHandshakesForResourceDirectoryResponseBodyHandshakesHandshake {
	s.TargetEntity = &v
	return s
}

func (s *ListHandshakesForResourceDirectoryResponseBodyHandshakesHandshake) SetTargetType(v string) *ListHandshakesForResourceDirectoryResponseBodyHandshakesHandshake {
	s.TargetType = &v
	return s
}

type ListHandshakesForResourceDirectoryResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListHandshakesForResourceDirectoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListHandshakesForResourceDirectoryResponse) String() string {
	return tea.Prettify(s)
}

func (s ListHandshakesForResourceDirectoryResponse) GoString() string {
	return s.String()
}

func (s *ListHandshakesForResourceDirectoryResponse) SetHeaders(v map[string]*string) *ListHandshakesForResourceDirectoryResponse {
	s.Headers = v
	return s
}

func (s *ListHandshakesForResourceDirectoryResponse) SetStatusCode(v int32) *ListHandshakesForResourceDirectoryResponse {
	s.StatusCode = &v
	return s
}

func (s *ListHandshakesForResourceDirectoryResponse) SetBody(v *ListHandshakesForResourceDirectoryResponseBody) *ListHandshakesForResourceDirectoryResponse {
	s.Body = v
	return s
}

type ListMessageContactVerificationsRequest struct {
	// The ID of the contact.
	//
	// example:
	//
	// c-qL4HqKONzOM7****
	ContactId *string `json:"ContactId,omitempty" xml:"ContactId,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListMessageContactVerificationsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListMessageContactVerificationsRequest) GoString() string {
	return s.String()
}

func (s *ListMessageContactVerificationsRequest) SetContactId(v string) *ListMessageContactVerificationsRequest {
	s.ContactId = &v
	return s
}

func (s *ListMessageContactVerificationsRequest) SetPageNumber(v int32) *ListMessageContactVerificationsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListMessageContactVerificationsRequest) SetPageSize(v int32) *ListMessageContactVerificationsRequest {
	s.PageSize = &v
	return s
}

type ListMessageContactVerificationsResponseBody struct {
	// The record for the contact to be verified.
	ContactVerifications []*ListMessageContactVerificationsResponseBodyContactVerifications `json:"ContactVerifications,omitempty" xml:"ContactVerifications,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// CD76D376-2517-4924-92C5-DBC52262F93A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 48
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListMessageContactVerificationsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListMessageContactVerificationsResponseBody) GoString() string {
	return s.String()
}

func (s *ListMessageContactVerificationsResponseBody) SetContactVerifications(v []*ListMessageContactVerificationsResponseBodyContactVerifications) *ListMessageContactVerificationsResponseBody {
	s.ContactVerifications = v
	return s
}

func (s *ListMessageContactVerificationsResponseBody) SetPageNumber(v int32) *ListMessageContactVerificationsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListMessageContactVerificationsResponseBody) SetPageSize(v int32) *ListMessageContactVerificationsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListMessageContactVerificationsResponseBody) SetRequestId(v string) *ListMessageContactVerificationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMessageContactVerificationsResponseBody) SetTotalCount(v int32) *ListMessageContactVerificationsResponseBody {
	s.TotalCount = &v
	return s
}

type ListMessageContactVerificationsResponseBodyContactVerifications struct {
	// The ID of the contact.
	//
	// example:
	//
	// c-qL4HqKONzOM7****
	ContactId *string `json:"ContactId,omitempty" xml:"ContactId,omitempty"`
	// The object that is used for verification. Valid values:
	//
	// - Mobile phone number
	//
	// - Email address
	//
	// example:
	//
	// someone***@example.com
	Target *string `json:"Target,omitempty" xml:"Target,omitempty"`
}

func (s ListMessageContactVerificationsResponseBodyContactVerifications) String() string {
	return tea.Prettify(s)
}

func (s ListMessageContactVerificationsResponseBodyContactVerifications) GoString() string {
	return s.String()
}

func (s *ListMessageContactVerificationsResponseBodyContactVerifications) SetContactId(v string) *ListMessageContactVerificationsResponseBodyContactVerifications {
	s.ContactId = &v
	return s
}

func (s *ListMessageContactVerificationsResponseBodyContactVerifications) SetTarget(v string) *ListMessageContactVerificationsResponseBodyContactVerifications {
	s.Target = &v
	return s
}

type ListMessageContactVerificationsResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListMessageContactVerificationsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMessageContactVerificationsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListMessageContactVerificationsResponse) GoString() string {
	return s.String()
}

func (s *ListMessageContactVerificationsResponse) SetHeaders(v map[string]*string) *ListMessageContactVerificationsResponse {
	s.Headers = v
	return s
}

func (s *ListMessageContactVerificationsResponse) SetStatusCode(v int32) *ListMessageContactVerificationsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMessageContactVerificationsResponse) SetBody(v *ListMessageContactVerificationsResponseBody) *ListMessageContactVerificationsResponse {
	s.Body = v
	return s
}

type ListMessageContactsRequest struct {
	// The ID of the contact.
	//
	// example:
	//
	// c-qL4HqKONzOM7****
	ContactId *string `json:"ContactId,omitempty" xml:"ContactId,omitempty"`
	// The ID of the object to which the contact is bound. Valid values:
	//
	// 	- ID of the resource directory
	//
	// 	- ID of the folder
	//
	// 	- ID of the member
	//
	// example:
	//
	// fd-ZDNPiT****
	Member *string `json:"Member,omitempty" xml:"Member,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListMessageContactsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListMessageContactsRequest) GoString() string {
	return s.String()
}

func (s *ListMessageContactsRequest) SetContactId(v string) *ListMessageContactsRequest {
	s.ContactId = &v
	return s
}

func (s *ListMessageContactsRequest) SetMember(v string) *ListMessageContactsRequest {
	s.Member = &v
	return s
}

func (s *ListMessageContactsRequest) SetPageNumber(v int32) *ListMessageContactsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListMessageContactsRequest) SetPageSize(v int32) *ListMessageContactsRequest {
	s.PageSize = &v
	return s
}

type ListMessageContactsResponseBody struct {
	// The time when the contact was bound to the objects.
	Contacts []*ListMessageContactsResponseBodyContacts `json:"Contacts,omitempty" xml:"Contacts,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 95060F1D-6990-4645-8920-A81D1BBFE992
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 20
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListMessageContactsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListMessageContactsResponseBody) GoString() string {
	return s.String()
}

func (s *ListMessageContactsResponseBody) SetContacts(v []*ListMessageContactsResponseBodyContacts) *ListMessageContactsResponseBody {
	s.Contacts = v
	return s
}

func (s *ListMessageContactsResponseBody) SetPageNumber(v int32) *ListMessageContactsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListMessageContactsResponseBody) SetPageSize(v int32) *ListMessageContactsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListMessageContactsResponseBody) SetRequestId(v string) *ListMessageContactsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMessageContactsResponseBody) SetTotalCount(v int32) *ListMessageContactsResponseBody {
	s.TotalCount = &v
	return s
}

type ListMessageContactsResponseBodyContacts struct {
	// The time when the contact was bound to the objects.
	//
	// example:
	//
	// 2023-03-27 17:19:21
	AssociatedDate *string `json:"AssociatedDate,omitempty" xml:"AssociatedDate,omitempty"`
	// The ID of the contact.
	//
	// example:
	//
	// c-qL4HqKONzOM7****
	ContactId *string `json:"ContactId,omitempty" xml:"ContactId,omitempty"`
	// The time when the contact was added.
	//
	// example:
	//
	// 2023-03-27 17:19:21
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	// The email address of the contact.
	//
	// example:
	//
	// someone***@example.com
	EmailAddress *string `json:"EmailAddress,omitempty" xml:"EmailAddress,omitempty"`
	// The IDs of objects to which the contact is bound.
	Members []*string `json:"Members,omitempty" xml:"Members,omitempty" type:"Repeated"`
	// The types of messages received by the contact.
	MessageTypes []*string `json:"MessageTypes,omitempty" xml:"MessageTypes,omitempty" type:"Repeated"`
	// The name of the contact.
	//
	// example:
	//
	// tom
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The mobile phone number of the contact.
	//
	// example:
	//
	// 86-139****1234
	PhoneNumber *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	// The status of the contact. Valid values:
	//
	// - Verifying
	//
	// - Active
	//
	// - Deleting
	//
	// example:
	//
	// Active
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The job title of the contact.
	//
	// example:
	//
	// TechnicalDirector
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s ListMessageContactsResponseBodyContacts) String() string {
	return tea.Prettify(s)
}

func (s ListMessageContactsResponseBodyContacts) GoString() string {
	return s.String()
}

func (s *ListMessageContactsResponseBodyContacts) SetAssociatedDate(v string) *ListMessageContactsResponseBodyContacts {
	s.AssociatedDate = &v
	return s
}

func (s *ListMessageContactsResponseBodyContacts) SetContactId(v string) *ListMessageContactsResponseBodyContacts {
	s.ContactId = &v
	return s
}

func (s *ListMessageContactsResponseBodyContacts) SetCreateDate(v string) *ListMessageContactsResponseBodyContacts {
	s.CreateDate = &v
	return s
}

func (s *ListMessageContactsResponseBodyContacts) SetEmailAddress(v string) *ListMessageContactsResponseBodyContacts {
	s.EmailAddress = &v
	return s
}

func (s *ListMessageContactsResponseBodyContacts) SetMembers(v []*string) *ListMessageContactsResponseBodyContacts {
	s.Members = v
	return s
}

func (s *ListMessageContactsResponseBodyContacts) SetMessageTypes(v []*string) *ListMessageContactsResponseBodyContacts {
	s.MessageTypes = v
	return s
}

func (s *ListMessageContactsResponseBodyContacts) SetName(v string) *ListMessageContactsResponseBodyContacts {
	s.Name = &v
	return s
}

func (s *ListMessageContactsResponseBodyContacts) SetPhoneNumber(v string) *ListMessageContactsResponseBodyContacts {
	s.PhoneNumber = &v
	return s
}

func (s *ListMessageContactsResponseBodyContacts) SetStatus(v string) *ListMessageContactsResponseBodyContacts {
	s.Status = &v
	return s
}

func (s *ListMessageContactsResponseBodyContacts) SetTitle(v string) *ListMessageContactsResponseBodyContacts {
	s.Title = &v
	return s
}

type ListMessageContactsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListMessageContactsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMessageContactsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListMessageContactsResponse) GoString() string {
	return s.String()
}

func (s *ListMessageContactsResponse) SetHeaders(v map[string]*string) *ListMessageContactsResponse {
	s.Headers = v
	return s
}

func (s *ListMessageContactsResponse) SetStatusCode(v int32) *ListMessageContactsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMessageContactsResponse) SetBody(v *ListMessageContactsResponseBody) *ListMessageContactsResponse {
	s.Body = v
	return s
}

type ListTagKeysRequest struct {
	// The tag key for a fuzzy query.
	//
	// example:
	//
	// team
	KeyFilter *string `json:"KeyFilter,omitempty" xml:"KeyFilter,omitempty"`
	// The maximum number of entries to return for a single request.
	//
	// Valid values: 1 to 100. Default value: 10.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. You do not need to specify this parameter for the first request.
	//
	// example:
	//
	// TGlzdFJlc291cm****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The resource type.
	//
	// The value Account indicates the members of the resource directory.
	//
	// This parameter is required.
	//
	// example:
	//
	// Account
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s ListTagKeysRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTagKeysRequest) GoString() string {
	return s.String()
}

func (s *ListTagKeysRequest) SetKeyFilter(v string) *ListTagKeysRequest {
	s.KeyFilter = &v
	return s
}

func (s *ListTagKeysRequest) SetMaxResults(v int32) *ListTagKeysRequest {
	s.MaxResults = &v
	return s
}

func (s *ListTagKeysRequest) SetNextToken(v string) *ListTagKeysRequest {
	s.NextToken = &v
	return s
}

func (s *ListTagKeysRequest) SetResourceType(v string) *ListTagKeysRequest {
	s.ResourceType = &v
	return s
}

type ListTagKeysResponseBody struct {
	// Indicates whether the next query is required.
	//
	// 	- If the value of this parameter is empty (`"NextToken": ""`), all results are returned, and the next query is not required.
	//
	// 	- If the value of this parameter is not empty, the next query is required, and the value is the token used to start the next query.
	//
	// example:
	//
	// TGlzdFJlc291cm****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// DC09A6AA-2713-4E10-A2E9-E6C5C43A8842
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the tag keys.
	Tags []*ListTagKeysResponseBodyTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s ListTagKeysResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTagKeysResponseBody) GoString() string {
	return s.String()
}

func (s *ListTagKeysResponseBody) SetNextToken(v string) *ListTagKeysResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTagKeysResponseBody) SetRequestId(v string) *ListTagKeysResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTagKeysResponseBody) SetTags(v []*ListTagKeysResponseBodyTags) *ListTagKeysResponseBody {
	s.Tags = v
	return s
}

type ListTagKeysResponseBodyTags struct {
	// The tag key.
	//
	// example:
	//
	// team
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
}

func (s ListTagKeysResponseBodyTags) String() string {
	return tea.Prettify(s)
}

func (s ListTagKeysResponseBodyTags) GoString() string {
	return s.String()
}

func (s *ListTagKeysResponseBodyTags) SetKey(v string) *ListTagKeysResponseBodyTags {
	s.Key = &v
	return s
}

type ListTagKeysResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTagKeysResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTagKeysResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTagKeysResponse) GoString() string {
	return s.String()
}

func (s *ListTagKeysResponse) SetHeaders(v map[string]*string) *ListTagKeysResponse {
	s.Headers = v
	return s
}

func (s *ListTagKeysResponse) SetStatusCode(v int32) *ListTagKeysResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTagKeysResponse) SetBody(v *ListTagKeysResponseBody) *ListTagKeysResponse {
	s.Body = v
	return s
}

type ListTagResourcesRequest struct {
	// The number of entries to return on each page.
	//
	// Valid values: 1 to 100. Default value: 10.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that is used to start the next query.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The Alibaba Cloud account IDs of the members. This parameter specifies a filter condition for the query.
	//
	// > If you want to query the tags that are added to the members in a resource directory, you must configure both the `ResourceId` and `ResourceType` parameters and set the `ResourceType` parameter to `Account` in your request.
	ResourceId []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	// The type of the objects whose tags you want to query. This parameter specifies a filter condition for the query. Valid values:
	//
	// 	- Account: member
	//
	// This parameter is required.
	//
	// example:
	//
	// Account
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The tags. This parameter specifies a filter condition for the query.
	//
	// You can specify a maximum of 20 tags.
	Tag []*ListTagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListTagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesRequest) GoString() string {
	return s.String()
}

func (s *ListTagResourcesRequest) SetMaxResults(v int32) *ListTagResourcesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListTagResourcesRequest) SetNextToken(v string) *ListTagResourcesRequest {
	s.NextToken = &v
	return s
}

func (s *ListTagResourcesRequest) SetResourceId(v []*string) *ListTagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *ListTagResourcesRequest) SetResourceType(v string) *ListTagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesRequest) SetTag(v []*ListTagResourcesRequestTag) *ListTagResourcesRequest {
	s.Tag = v
	return s
}

type ListTagResourcesRequestTag struct {
	// The key of the tag.
	//
	// example:
	//
	// k1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag.
	//
	// example:
	//
	// v1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListTagResourcesRequestTag) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesRequestTag) GoString() string {
	return s.String()
}

func (s *ListTagResourcesRequestTag) SetKey(v string) *ListTagResourcesRequestTag {
	s.Key = &v
	return s
}

func (s *ListTagResourcesRequestTag) SetValue(v string) *ListTagResourcesRequestTag {
	s.Value = &v
	return s
}

type ListTagResourcesResponseBody struct {
	// Indicates whether the next query is required.```` Valid values:
	//
	// 	- If the value of this parameter is empty (`"NextToken": ""`), all results are returned, and the `next query` is not required.
	//
	// 	- If the value of this parameter is not empty, the next query is required, and the value is the token used to start the next query.````
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 8054B059-6B36-53BF-AA45-B8C9A0ED05AB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The tags.
	TagResources []*ListTagResourcesResponseBodyTagResources `json:"TagResources,omitempty" xml:"TagResources,omitempty" type:"Repeated"`
}

func (s ListTagResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBody) SetNextToken(v string) *ListTagResourcesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTagResourcesResponseBody) SetRequestId(v string) *ListTagResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTagResourcesResponseBody) SetTagResources(v []*ListTagResourcesResponseBodyTagResources) *ListTagResourcesResponseBody {
	s.TagResources = v
	return s
}

type ListTagResourcesResponseBodyTagResources struct {
	// The Alibaba Cloud account ID of the member.
	//
	// example:
	//
	// 1098***
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The type of the object whose tags are queried. Valid values:
	//
	// 	- Account: member
	//
	// example:
	//
	// Account
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The key of the tag.
	//
	// example:
	//
	// k1
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The value of the tag.
	//
	// example:
	//
	// k1
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s ListTagResourcesResponseBodyTagResources) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseBodyTagResources) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBodyTagResources) SetResourceId(v string) *ListTagResourcesResponseBodyTagResources {
	s.ResourceId = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResources) SetResourceType(v string) *ListTagResourcesResponseBodyTagResources {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResources) SetTagKey(v string) *ListTagResourcesResponseBodyTagResources {
	s.TagKey = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResources) SetTagValue(v string) *ListTagResourcesResponseBodyTagResources {
	s.TagValue = &v
	return s
}

type ListTagResourcesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponse) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponse) SetHeaders(v map[string]*string) *ListTagResourcesResponse {
	s.Headers = v
	return s
}

func (s *ListTagResourcesResponse) SetStatusCode(v int32) *ListTagResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTagResourcesResponse) SetBody(v *ListTagResourcesResponseBody) *ListTagResourcesResponse {
	s.Body = v
	return s
}

type ListTagValuesRequest struct {
	// The maximum number of entries to return for a single request.
	//
	// Valid values: 1 to 100. Default value: 10.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. You do not need to specify this parameter for the first request.
	//
	// example:
	//
	// TGlzdFJlc291cm****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The resource type.
	//
	// The value Account indicates the members of the resource directory.
	//
	// This parameter is required.
	//
	// example:
	//
	// Account
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The tag key. This parameter specifies a filter condition for the query.
	//
	// This parameter is required.
	//
	// example:
	//
	// k1
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The tag value for a fuzzy query.
	//
	// example:
	//
	// v1
	ValueFilter *string `json:"ValueFilter,omitempty" xml:"ValueFilter,omitempty"`
}

func (s ListTagValuesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTagValuesRequest) GoString() string {
	return s.String()
}

func (s *ListTagValuesRequest) SetMaxResults(v int32) *ListTagValuesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListTagValuesRequest) SetNextToken(v string) *ListTagValuesRequest {
	s.NextToken = &v
	return s
}

func (s *ListTagValuesRequest) SetResourceType(v string) *ListTagValuesRequest {
	s.ResourceType = &v
	return s
}

func (s *ListTagValuesRequest) SetTagKey(v string) *ListTagValuesRequest {
	s.TagKey = &v
	return s
}

func (s *ListTagValuesRequest) SetValueFilter(v string) *ListTagValuesRequest {
	s.ValueFilter = &v
	return s
}

type ListTagValuesResponseBody struct {
	// Indicates whether the next query is required.
	//
	// 	- If the value of this parameter is empty (`"NextToken": ""`), all results are returned, and the next query is not required.
	//
	// 	- If the value of this parameter is not empty, the next query is required, and the value is the token used to start the next query.
	//
	// example:
	//
	// TGlzdFJlc291cm****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// DC09A6AA-2713-4E10-A2E9-E6C5C43A8842
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the tag values.
	Tags []*ListTagValuesResponseBodyTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s ListTagValuesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTagValuesResponseBody) GoString() string {
	return s.String()
}

func (s *ListTagValuesResponseBody) SetNextToken(v string) *ListTagValuesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTagValuesResponseBody) SetRequestId(v string) *ListTagValuesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTagValuesResponseBody) SetTags(v []*ListTagValuesResponseBodyTags) *ListTagValuesResponseBody {
	s.Tags = v
	return s
}

type ListTagValuesResponseBodyTags struct {
	// The tag value.
	//
	// example:
	//
	// v1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListTagValuesResponseBodyTags) String() string {
	return tea.Prettify(s)
}

func (s ListTagValuesResponseBodyTags) GoString() string {
	return s.String()
}

func (s *ListTagValuesResponseBodyTags) SetValue(v string) *ListTagValuesResponseBodyTags {
	s.Value = &v
	return s
}

type ListTagValuesResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTagValuesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTagValuesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTagValuesResponse) GoString() string {
	return s.String()
}

func (s *ListTagValuesResponse) SetHeaders(v map[string]*string) *ListTagValuesResponse {
	s.Headers = v
	return s
}

func (s *ListTagValuesResponse) SetStatusCode(v int32) *ListTagValuesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTagValuesResponse) SetBody(v *ListTagValuesResponseBody) *ListTagValuesResponse {
	s.Body = v
	return s
}

type ListTargetAttachmentsForControlPolicyRequest struct {
	// The number of the page to return.
	//
	// Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page.
	//
	// Valid values: 1 to 100. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the access control policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// cp-jExXAqIYkwHN****
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
}

func (s ListTargetAttachmentsForControlPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTargetAttachmentsForControlPolicyRequest) GoString() string {
	return s.String()
}

func (s *ListTargetAttachmentsForControlPolicyRequest) SetPageNumber(v int32) *ListTargetAttachmentsForControlPolicyRequest {
	s.PageNumber = &v
	return s
}

func (s *ListTargetAttachmentsForControlPolicyRequest) SetPageSize(v int32) *ListTargetAttachmentsForControlPolicyRequest {
	s.PageSize = &v
	return s
}

func (s *ListTargetAttachmentsForControlPolicyRequest) SetPolicyId(v string) *ListTargetAttachmentsForControlPolicyRequest {
	s.PolicyId = &v
	return s
}

type ListTargetAttachmentsForControlPolicyResponseBody struct {
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// B32BD3D6-1089-41F3-8E70-E0079BC7D760
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the objects to which the access control policy is attached.
	TargetAttachments *ListTargetAttachmentsForControlPolicyResponseBodyTargetAttachments `json:"TargetAttachments,omitempty" xml:"TargetAttachments,omitempty" type:"Struct"`
	// The total number of objects to which the access control policy is attached.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListTargetAttachmentsForControlPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTargetAttachmentsForControlPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *ListTargetAttachmentsForControlPolicyResponseBody) SetPageNumber(v int32) *ListTargetAttachmentsForControlPolicyResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListTargetAttachmentsForControlPolicyResponseBody) SetPageSize(v int32) *ListTargetAttachmentsForControlPolicyResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListTargetAttachmentsForControlPolicyResponseBody) SetRequestId(v string) *ListTargetAttachmentsForControlPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTargetAttachmentsForControlPolicyResponseBody) SetTargetAttachments(v *ListTargetAttachmentsForControlPolicyResponseBodyTargetAttachments) *ListTargetAttachmentsForControlPolicyResponseBody {
	s.TargetAttachments = v
	return s
}

func (s *ListTargetAttachmentsForControlPolicyResponseBody) SetTotalCount(v int32) *ListTargetAttachmentsForControlPolicyResponseBody {
	s.TotalCount = &v
	return s
}

type ListTargetAttachmentsForControlPolicyResponseBodyTargetAttachments struct {
	TargetAttachment []*ListTargetAttachmentsForControlPolicyResponseBodyTargetAttachmentsTargetAttachment `json:"TargetAttachment,omitempty" xml:"TargetAttachment,omitempty" type:"Repeated"`
}

func (s ListTargetAttachmentsForControlPolicyResponseBodyTargetAttachments) String() string {
	return tea.Prettify(s)
}

func (s ListTargetAttachmentsForControlPolicyResponseBodyTargetAttachments) GoString() string {
	return s.String()
}

func (s *ListTargetAttachmentsForControlPolicyResponseBodyTargetAttachments) SetTargetAttachment(v []*ListTargetAttachmentsForControlPolicyResponseBodyTargetAttachmentsTargetAttachment) *ListTargetAttachmentsForControlPolicyResponseBodyTargetAttachments {
	s.TargetAttachment = v
	return s
}

type ListTargetAttachmentsForControlPolicyResponseBodyTargetAttachmentsTargetAttachment struct {
	// The time when the access control policy was attached to the object.
	//
	// example:
	//
	// 2021-03-19T02:56:24Z
	AttachDate *string `json:"AttachDate,omitempty" xml:"AttachDate,omitempty"`
	// The ID of the object.
	//
	// example:
	//
	// fd-ZDNPiT****
	TargetId *string `json:"TargetId,omitempty" xml:"TargetId,omitempty"`
	// The name of the object.
	//
	// example:
	//
	// Dev_Department
	TargetName *string `json:"TargetName,omitempty" xml:"TargetName,omitempty"`
	// The type of the object. Valid values:
	//
	// 	- Root: Root folder
	//
	// 	- Folder: subfolder of the Root folder
	//
	// 	- Account: member
	//
	// example:
	//
	// Folder
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
}

func (s ListTargetAttachmentsForControlPolicyResponseBodyTargetAttachmentsTargetAttachment) String() string {
	return tea.Prettify(s)
}

func (s ListTargetAttachmentsForControlPolicyResponseBodyTargetAttachmentsTargetAttachment) GoString() string {
	return s.String()
}

func (s *ListTargetAttachmentsForControlPolicyResponseBodyTargetAttachmentsTargetAttachment) SetAttachDate(v string) *ListTargetAttachmentsForControlPolicyResponseBodyTargetAttachmentsTargetAttachment {
	s.AttachDate = &v
	return s
}

func (s *ListTargetAttachmentsForControlPolicyResponseBodyTargetAttachmentsTargetAttachment) SetTargetId(v string) *ListTargetAttachmentsForControlPolicyResponseBodyTargetAttachmentsTargetAttachment {
	s.TargetId = &v
	return s
}

func (s *ListTargetAttachmentsForControlPolicyResponseBodyTargetAttachmentsTargetAttachment) SetTargetName(v string) *ListTargetAttachmentsForControlPolicyResponseBodyTargetAttachmentsTargetAttachment {
	s.TargetName = &v
	return s
}

func (s *ListTargetAttachmentsForControlPolicyResponseBodyTargetAttachmentsTargetAttachment) SetTargetType(v string) *ListTargetAttachmentsForControlPolicyResponseBodyTargetAttachmentsTargetAttachment {
	s.TargetType = &v
	return s
}

type ListTargetAttachmentsForControlPolicyResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTargetAttachmentsForControlPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTargetAttachmentsForControlPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTargetAttachmentsForControlPolicyResponse) GoString() string {
	return s.String()
}

func (s *ListTargetAttachmentsForControlPolicyResponse) SetHeaders(v map[string]*string) *ListTargetAttachmentsForControlPolicyResponse {
	s.Headers = v
	return s
}

func (s *ListTargetAttachmentsForControlPolicyResponse) SetStatusCode(v int32) *ListTargetAttachmentsForControlPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTargetAttachmentsForControlPolicyResponse) SetBody(v *ListTargetAttachmentsForControlPolicyResponseBody) *ListTargetAttachmentsForControlPolicyResponse {
	s.Body = v
	return s
}

type ListTrustedServiceStatusRequest struct {
	// The ID of the management account or delegated administrator account.
	//
	// 	- If you set this parameter to the ID of a management account, the trusted services that are enabled within the management account are queried. The default value of this parameter is the ID of an management account.
	//
	// 	- If you set this parameter to the ID of a delegated administrator account, the trusted services that are enabled within the delegated administrator account are queried.
	//
	// For more information about trusted services and delegated administrator accounts, see [Overview of trusted services](https://help.aliyun.com/document_detail/208133.html) and [Delegated administrator accounts](https://help.aliyun.com/document_detail/208117.html).
	//
	// example:
	//
	// 177242285274****
	AdminAccountId *string `json:"AdminAccountId,omitempty" xml:"AdminAccountId,omitempty"`
	// The number of the page to return.
	//
	// Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page.
	//
	// Valid values: 1 to 100. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListTrustedServiceStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTrustedServiceStatusRequest) GoString() string {
	return s.String()
}

func (s *ListTrustedServiceStatusRequest) SetAdminAccountId(v string) *ListTrustedServiceStatusRequest {
	s.AdminAccountId = &v
	return s
}

func (s *ListTrustedServiceStatusRequest) SetPageNumber(v int32) *ListTrustedServiceStatusRequest {
	s.PageNumber = &v
	return s
}

func (s *ListTrustedServiceStatusRequest) SetPageSize(v int32) *ListTrustedServiceStatusRequest {
	s.PageSize = &v
	return s
}

type ListTrustedServiceStatusResponseBody struct {
	// The information about the trusted services that are enabled.
	EnabledServicePrincipals *ListTrustedServiceStatusResponseBodyEnabledServicePrincipals `json:"EnabledServicePrincipals,omitempty" xml:"EnabledServicePrincipals,omitempty" type:"Struct"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// CD76D376-2517-4924-92C5-DBC52262F93A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 2
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListTrustedServiceStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTrustedServiceStatusResponseBody) GoString() string {
	return s.String()
}

func (s *ListTrustedServiceStatusResponseBody) SetEnabledServicePrincipals(v *ListTrustedServiceStatusResponseBodyEnabledServicePrincipals) *ListTrustedServiceStatusResponseBody {
	s.EnabledServicePrincipals = v
	return s
}

func (s *ListTrustedServiceStatusResponseBody) SetPageNumber(v int32) *ListTrustedServiceStatusResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListTrustedServiceStatusResponseBody) SetPageSize(v int32) *ListTrustedServiceStatusResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListTrustedServiceStatusResponseBody) SetRequestId(v string) *ListTrustedServiceStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTrustedServiceStatusResponseBody) SetTotalCount(v int32) *ListTrustedServiceStatusResponseBody {
	s.TotalCount = &v
	return s
}

type ListTrustedServiceStatusResponseBodyEnabledServicePrincipals struct {
	EnabledServicePrincipal []*ListTrustedServiceStatusResponseBodyEnabledServicePrincipalsEnabledServicePrincipal `json:"EnabledServicePrincipal,omitempty" xml:"EnabledServicePrincipal,omitempty" type:"Repeated"`
}

func (s ListTrustedServiceStatusResponseBodyEnabledServicePrincipals) String() string {
	return tea.Prettify(s)
}

func (s ListTrustedServiceStatusResponseBodyEnabledServicePrincipals) GoString() string {
	return s.String()
}

func (s *ListTrustedServiceStatusResponseBodyEnabledServicePrincipals) SetEnabledServicePrincipal(v []*ListTrustedServiceStatusResponseBodyEnabledServicePrincipalsEnabledServicePrincipal) *ListTrustedServiceStatusResponseBodyEnabledServicePrincipals {
	s.EnabledServicePrincipal = v
	return s
}

type ListTrustedServiceStatusResponseBodyEnabledServicePrincipalsEnabledServicePrincipal struct {
	// The time when the trusted service was enabled.
	//
	// example:
	//
	// 2019-02-18T15:32:10.473Z
	EnableTime *string `json:"EnableTime,omitempty" xml:"EnableTime,omitempty"`
	// The identifier of the trusted service.
	//
	// example:
	//
	// config.aliyuncs.com
	ServicePrincipal *string `json:"ServicePrincipal,omitempty" xml:"ServicePrincipal,omitempty"`
}

func (s ListTrustedServiceStatusResponseBodyEnabledServicePrincipalsEnabledServicePrincipal) String() string {
	return tea.Prettify(s)
}

func (s ListTrustedServiceStatusResponseBodyEnabledServicePrincipalsEnabledServicePrincipal) GoString() string {
	return s.String()
}

func (s *ListTrustedServiceStatusResponseBodyEnabledServicePrincipalsEnabledServicePrincipal) SetEnableTime(v string) *ListTrustedServiceStatusResponseBodyEnabledServicePrincipalsEnabledServicePrincipal {
	s.EnableTime = &v
	return s
}

func (s *ListTrustedServiceStatusResponseBodyEnabledServicePrincipalsEnabledServicePrincipal) SetServicePrincipal(v string) *ListTrustedServiceStatusResponseBodyEnabledServicePrincipalsEnabledServicePrincipal {
	s.ServicePrincipal = &v
	return s
}

type ListTrustedServiceStatusResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTrustedServiceStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTrustedServiceStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTrustedServiceStatusResponse) GoString() string {
	return s.String()
}

func (s *ListTrustedServiceStatusResponse) SetHeaders(v map[string]*string) *ListTrustedServiceStatusResponse {
	s.Headers = v
	return s
}

func (s *ListTrustedServiceStatusResponse) SetStatusCode(v int32) *ListTrustedServiceStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTrustedServiceStatusResponse) SetBody(v *ListTrustedServiceStatusResponseBody) *ListTrustedServiceStatusResponse {
	s.Body = v
	return s
}

type MoveAccountRequest struct {
	// The Alibaba Cloud account ID of the member that you want to move.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12323344****
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The ID of the destination folder.
	//
	// This parameter is required.
	//
	// example:
	//
	// fd-bVaRIG****
	DestinationFolderId *string `json:"DestinationFolderId,omitempty" xml:"DestinationFolderId,omitempty"`
}

func (s MoveAccountRequest) String() string {
	return tea.Prettify(s)
}

func (s MoveAccountRequest) GoString() string {
	return s.String()
}

func (s *MoveAccountRequest) SetAccountId(v string) *MoveAccountRequest {
	s.AccountId = &v
	return s
}

func (s *MoveAccountRequest) SetDestinationFolderId(v string) *MoveAccountRequest {
	s.DestinationFolderId = &v
	return s
}

type MoveAccountResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 9B34724D-54B0-4A51-B34D-4512372FE1BE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s MoveAccountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s MoveAccountResponseBody) GoString() string {
	return s.String()
}

func (s *MoveAccountResponseBody) SetRequestId(v string) *MoveAccountResponseBody {
	s.RequestId = &v
	return s
}

type MoveAccountResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MoveAccountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s MoveAccountResponse) String() string {
	return tea.Prettify(s)
}

func (s MoveAccountResponse) GoString() string {
	return s.String()
}

func (s *MoveAccountResponse) SetHeaders(v map[string]*string) *MoveAccountResponse {
	s.Headers = v
	return s
}

func (s *MoveAccountResponse) SetStatusCode(v int32) *MoveAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *MoveAccountResponse) SetBody(v *MoveAccountResponseBody) *MoveAccountResponse {
	s.Body = v
	return s
}

type PrecheckForConsolidatedBillingAccountRequest struct {
	// The ID of the management account or member to be used as a main financial account.
	//
	// This parameter is required.
	//
	// example:
	//
	// 111***089
	BillingAccountId *string `json:"BillingAccountId,omitempty" xml:"BillingAccountId,omitempty"`
}

func (s PrecheckForConsolidatedBillingAccountRequest) String() string {
	return tea.Prettify(s)
}

func (s PrecheckForConsolidatedBillingAccountRequest) GoString() string {
	return s.String()
}

func (s *PrecheckForConsolidatedBillingAccountRequest) SetBillingAccountId(v string) *PrecheckForConsolidatedBillingAccountRequest {
	s.BillingAccountId = &v
	return s
}

type PrecheckForConsolidatedBillingAccountResponseBody struct {
	// The cause of the check failure.
	Reasons []*PrecheckForConsolidatedBillingAccountResponseBodyReasons `json:"Reasons,omitempty" xml:"Reasons,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 9E6B6CA8-9E7A-521F-A743-AA582714727E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the check was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// false
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s PrecheckForConsolidatedBillingAccountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PrecheckForConsolidatedBillingAccountResponseBody) GoString() string {
	return s.String()
}

func (s *PrecheckForConsolidatedBillingAccountResponseBody) SetReasons(v []*PrecheckForConsolidatedBillingAccountResponseBodyReasons) *PrecheckForConsolidatedBillingAccountResponseBody {
	s.Reasons = v
	return s
}

func (s *PrecheckForConsolidatedBillingAccountResponseBody) SetRequestId(v string) *PrecheckForConsolidatedBillingAccountResponseBody {
	s.RequestId = &v
	return s
}

func (s *PrecheckForConsolidatedBillingAccountResponseBody) SetResult(v bool) *PrecheckForConsolidatedBillingAccountResponseBody {
	s.Result = &v
	return s
}

type PrecheckForConsolidatedBillingAccountResponseBodyReasons struct {
	// The error code.
	//
	// example:
	//
	// PaymentAccountEnterpriseInvoiceError
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message.
	//
	// example:
	//
	// No enterprise invoice header information is set for the payment account.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s PrecheckForConsolidatedBillingAccountResponseBodyReasons) String() string {
	return tea.Prettify(s)
}

func (s PrecheckForConsolidatedBillingAccountResponseBodyReasons) GoString() string {
	return s.String()
}

func (s *PrecheckForConsolidatedBillingAccountResponseBodyReasons) SetCode(v string) *PrecheckForConsolidatedBillingAccountResponseBodyReasons {
	s.Code = &v
	return s
}

func (s *PrecheckForConsolidatedBillingAccountResponseBodyReasons) SetMessage(v string) *PrecheckForConsolidatedBillingAccountResponseBodyReasons {
	s.Message = &v
	return s
}

type PrecheckForConsolidatedBillingAccountResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PrecheckForConsolidatedBillingAccountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PrecheckForConsolidatedBillingAccountResponse) String() string {
	return tea.Prettify(s)
}

func (s PrecheckForConsolidatedBillingAccountResponse) GoString() string {
	return s.String()
}

func (s *PrecheckForConsolidatedBillingAccountResponse) SetHeaders(v map[string]*string) *PrecheckForConsolidatedBillingAccountResponse {
	s.Headers = v
	return s
}

func (s *PrecheckForConsolidatedBillingAccountResponse) SetStatusCode(v int32) *PrecheckForConsolidatedBillingAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *PrecheckForConsolidatedBillingAccountResponse) SetBody(v *PrecheckForConsolidatedBillingAccountResponseBody) *PrecheckForConsolidatedBillingAccountResponse {
	s.Body = v
	return s
}

type RegisterDelegatedAdministratorRequest struct {
	// The Alibaba Cloud account ID of the member in the resource directory.
	//
	// This parameter is required.
	//
	// example:
	//
	// 181761095690****
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The identifier of the trusted service.
	//
	// For more information, see the `Trusted service identifier` column in [Supported trusted services](https://help.aliyun.com/document_detail/208133.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// cloudfw.aliyuncs.com
	ServicePrincipal *string `json:"ServicePrincipal,omitempty" xml:"ServicePrincipal,omitempty"`
}

func (s RegisterDelegatedAdministratorRequest) String() string {
	return tea.Prettify(s)
}

func (s RegisterDelegatedAdministratorRequest) GoString() string {
	return s.String()
}

func (s *RegisterDelegatedAdministratorRequest) SetAccountId(v string) *RegisterDelegatedAdministratorRequest {
	s.AccountId = &v
	return s
}

func (s *RegisterDelegatedAdministratorRequest) SetServicePrincipal(v string) *RegisterDelegatedAdministratorRequest {
	s.ServicePrincipal = &v
	return s
}

type RegisterDelegatedAdministratorResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 0A45FC8F-54D2-4A65-8338-25E5DEBDA304
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RegisterDelegatedAdministratorResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RegisterDelegatedAdministratorResponseBody) GoString() string {
	return s.String()
}

func (s *RegisterDelegatedAdministratorResponseBody) SetRequestId(v string) *RegisterDelegatedAdministratorResponseBody {
	s.RequestId = &v
	return s
}

type RegisterDelegatedAdministratorResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RegisterDelegatedAdministratorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RegisterDelegatedAdministratorResponse) String() string {
	return tea.Prettify(s)
}

func (s RegisterDelegatedAdministratorResponse) GoString() string {
	return s.String()
}

func (s *RegisterDelegatedAdministratorResponse) SetHeaders(v map[string]*string) *RegisterDelegatedAdministratorResponse {
	s.Headers = v
	return s
}

func (s *RegisterDelegatedAdministratorResponse) SetStatusCode(v int32) *RegisterDelegatedAdministratorResponse {
	s.StatusCode = &v
	return s
}

func (s *RegisterDelegatedAdministratorResponse) SetBody(v *RegisterDelegatedAdministratorResponseBody) *RegisterDelegatedAdministratorResponse {
	s.Body = v
	return s
}

type RemoveCloudAccountRequest struct {
	// The Alibaba Cloud account ID of the member.
	//
	// This parameter is required.
	//
	// example:
	//
	// 177242285274****
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
}

func (s RemoveCloudAccountRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveCloudAccountRequest) GoString() string {
	return s.String()
}

func (s *RemoveCloudAccountRequest) SetAccountId(v string) *RemoveCloudAccountRequest {
	s.AccountId = &v
	return s
}

type RemoveCloudAccountResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 9B34724D-54B0-4A51-B34D-4512372FE1BE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RemoveCloudAccountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveCloudAccountResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveCloudAccountResponseBody) SetRequestId(v string) *RemoveCloudAccountResponseBody {
	s.RequestId = &v
	return s
}

type RemoveCloudAccountResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveCloudAccountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveCloudAccountResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveCloudAccountResponse) GoString() string {
	return s.String()
}

func (s *RemoveCloudAccountResponse) SetHeaders(v map[string]*string) *RemoveCloudAccountResponse {
	s.Headers = v
	return s
}

func (s *RemoveCloudAccountResponse) SetStatusCode(v int32) *RemoveCloudAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveCloudAccountResponse) SetBody(v *RemoveCloudAccountResponseBody) *RemoveCloudAccountResponse {
	s.Body = v
	return s
}

type RetryChangeAccountEmailRequest struct {
	// The Alibaba Cloud account ID of the member.
	//
	// This parameter is required.
	//
	// example:
	//
	// 181761095690****
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
}

func (s RetryChangeAccountEmailRequest) String() string {
	return tea.Prettify(s)
}

func (s RetryChangeAccountEmailRequest) GoString() string {
	return s.String()
}

func (s *RetryChangeAccountEmailRequest) SetAccountId(v string) *RetryChangeAccountEmailRequest {
	s.AccountId = &v
	return s
}

type RetryChangeAccountEmailResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 9B34724D-54B0-4A51-B34D-4512372FE1BE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RetryChangeAccountEmailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RetryChangeAccountEmailResponseBody) GoString() string {
	return s.String()
}

func (s *RetryChangeAccountEmailResponseBody) SetRequestId(v string) *RetryChangeAccountEmailResponseBody {
	s.RequestId = &v
	return s
}

type RetryChangeAccountEmailResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RetryChangeAccountEmailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RetryChangeAccountEmailResponse) String() string {
	return tea.Prettify(s)
}

func (s RetryChangeAccountEmailResponse) GoString() string {
	return s.String()
}

func (s *RetryChangeAccountEmailResponse) SetHeaders(v map[string]*string) *RetryChangeAccountEmailResponse {
	s.Headers = v
	return s
}

func (s *RetryChangeAccountEmailResponse) SetStatusCode(v int32) *RetryChangeAccountEmailResponse {
	s.StatusCode = &v
	return s
}

func (s *RetryChangeAccountEmailResponse) SetBody(v *RetryChangeAccountEmailResponseBody) *RetryChangeAccountEmailResponse {
	s.Body = v
	return s
}

type SendEmailVerificationForMessageContactRequest struct {
	// The ID of the contact.
	//
	// example:
	//
	// c-5gsZAGt***PGduDF
	ContactId *string `json:"ContactId,omitempty" xml:"ContactId,omitempty"`
	// The email address of the contact.
	//
	// The specified email address must be the one you specify when you call [AddMessageContact](~~AddMessageContact~~).
	//
	// example:
	//
	// someone***@example.com
	EmailAddress *string `json:"EmailAddress,omitempty" xml:"EmailAddress,omitempty"`
}

func (s SendEmailVerificationForMessageContactRequest) String() string {
	return tea.Prettify(s)
}

func (s SendEmailVerificationForMessageContactRequest) GoString() string {
	return s.String()
}

func (s *SendEmailVerificationForMessageContactRequest) SetContactId(v string) *SendEmailVerificationForMessageContactRequest {
	s.ContactId = &v
	return s
}

func (s *SendEmailVerificationForMessageContactRequest) SetEmailAddress(v string) *SendEmailVerificationForMessageContactRequest {
	s.EmailAddress = &v
	return s
}

type SendEmailVerificationForMessageContactResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 7B8A4E7D-6CFF-471D-84DF-195A7A241ECB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SendEmailVerificationForMessageContactResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SendEmailVerificationForMessageContactResponseBody) GoString() string {
	return s.String()
}

func (s *SendEmailVerificationForMessageContactResponseBody) SetRequestId(v string) *SendEmailVerificationForMessageContactResponseBody {
	s.RequestId = &v
	return s
}

type SendEmailVerificationForMessageContactResponse struct {
	Headers    map[string]*string                                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SendEmailVerificationForMessageContactResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SendEmailVerificationForMessageContactResponse) String() string {
	return tea.Prettify(s)
}

func (s SendEmailVerificationForMessageContactResponse) GoString() string {
	return s.String()
}

func (s *SendEmailVerificationForMessageContactResponse) SetHeaders(v map[string]*string) *SendEmailVerificationForMessageContactResponse {
	s.Headers = v
	return s
}

func (s *SendEmailVerificationForMessageContactResponse) SetStatusCode(v int32) *SendEmailVerificationForMessageContactResponse {
	s.StatusCode = &v
	return s
}

func (s *SendEmailVerificationForMessageContactResponse) SetBody(v *SendEmailVerificationForMessageContactResponseBody) *SendEmailVerificationForMessageContactResponse {
	s.Body = v
	return s
}

type SendPhoneVerificationForMessageContactRequest struct {
	// The ID of the contact.
	//
	// example:
	//
	// c-qL4HqKONzOM7****
	ContactId *string `json:"ContactId,omitempty" xml:"ContactId,omitempty"`
	// The mobile phone number of the contact.
	//
	// The value must be in the `<Country code>-<Mobile phone number>` format.
	//
	// The specified mobile phone number must be the one you specify when you call the [AddMessageContact](~~AddMessageContact~~) operation.
	//
	// example:
	//
	// 86-139****1234
	PhoneNumber *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
}

func (s SendPhoneVerificationForMessageContactRequest) String() string {
	return tea.Prettify(s)
}

func (s SendPhoneVerificationForMessageContactRequest) GoString() string {
	return s.String()
}

func (s *SendPhoneVerificationForMessageContactRequest) SetContactId(v string) *SendPhoneVerificationForMessageContactRequest {
	s.ContactId = &v
	return s
}

func (s *SendPhoneVerificationForMessageContactRequest) SetPhoneNumber(v string) *SendPhoneVerificationForMessageContactRequest {
	s.PhoneNumber = &v
	return s
}

type SendPhoneVerificationForMessageContactResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// CD76D376-2517-4924-92C5-DBC52262F93A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SendPhoneVerificationForMessageContactResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SendPhoneVerificationForMessageContactResponseBody) GoString() string {
	return s.String()
}

func (s *SendPhoneVerificationForMessageContactResponseBody) SetRequestId(v string) *SendPhoneVerificationForMessageContactResponseBody {
	s.RequestId = &v
	return s
}

type SendPhoneVerificationForMessageContactResponse struct {
	Headers    map[string]*string                                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SendPhoneVerificationForMessageContactResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SendPhoneVerificationForMessageContactResponse) String() string {
	return tea.Prettify(s)
}

func (s SendPhoneVerificationForMessageContactResponse) GoString() string {
	return s.String()
}

func (s *SendPhoneVerificationForMessageContactResponse) SetHeaders(v map[string]*string) *SendPhoneVerificationForMessageContactResponse {
	s.Headers = v
	return s
}

func (s *SendPhoneVerificationForMessageContactResponse) SetStatusCode(v int32) *SendPhoneVerificationForMessageContactResponse {
	s.StatusCode = &v
	return s
}

func (s *SendPhoneVerificationForMessageContactResponse) SetBody(v *SendPhoneVerificationForMessageContactResponseBody) *SendPhoneVerificationForMessageContactResponse {
	s.Body = v
	return s
}

type SendVerificationCodeForBindSecureMobilePhoneRequest struct {
	// The Alibaba Cloud account ID of the member.
	//
	// This parameter is required.
	//
	// example:
	//
	// 138660628348****
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The mobile phone number that you want to bind to the member for security purposes.
	//
	// Specify the mobile phone number in the \\<Country code>-\\<Mobile phone number> format.
	//
	// > Mobile phone numbers in the `86-<Mobile phone number>` format in the Chinese mainland are not supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// xx-13900001234
	SecureMobilePhone *string `json:"SecureMobilePhone,omitempty" xml:"SecureMobilePhone,omitempty"`
}

func (s SendVerificationCodeForBindSecureMobilePhoneRequest) String() string {
	return tea.Prettify(s)
}

func (s SendVerificationCodeForBindSecureMobilePhoneRequest) GoString() string {
	return s.String()
}

func (s *SendVerificationCodeForBindSecureMobilePhoneRequest) SetAccountId(v string) *SendVerificationCodeForBindSecureMobilePhoneRequest {
	s.AccountId = &v
	return s
}

func (s *SendVerificationCodeForBindSecureMobilePhoneRequest) SetSecureMobilePhone(v string) *SendVerificationCodeForBindSecureMobilePhoneRequest {
	s.SecureMobilePhone = &v
	return s
}

type SendVerificationCodeForBindSecureMobilePhoneResponseBody struct {
	// The time when the verification code expires.
	//
	// example:
	//
	// 2021-12-17T07:38:41.747Z
	ExpirationDate *string `json:"ExpirationDate,omitempty" xml:"ExpirationDate,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// DCD43660-75DD-5D15-B342-1B83FCA5B913
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SendVerificationCodeForBindSecureMobilePhoneResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SendVerificationCodeForBindSecureMobilePhoneResponseBody) GoString() string {
	return s.String()
}

func (s *SendVerificationCodeForBindSecureMobilePhoneResponseBody) SetExpirationDate(v string) *SendVerificationCodeForBindSecureMobilePhoneResponseBody {
	s.ExpirationDate = &v
	return s
}

func (s *SendVerificationCodeForBindSecureMobilePhoneResponseBody) SetRequestId(v string) *SendVerificationCodeForBindSecureMobilePhoneResponseBody {
	s.RequestId = &v
	return s
}

type SendVerificationCodeForBindSecureMobilePhoneResponse struct {
	Headers    map[string]*string                                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SendVerificationCodeForBindSecureMobilePhoneResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SendVerificationCodeForBindSecureMobilePhoneResponse) String() string {
	return tea.Prettify(s)
}

func (s SendVerificationCodeForBindSecureMobilePhoneResponse) GoString() string {
	return s.String()
}

func (s *SendVerificationCodeForBindSecureMobilePhoneResponse) SetHeaders(v map[string]*string) *SendVerificationCodeForBindSecureMobilePhoneResponse {
	s.Headers = v
	return s
}

func (s *SendVerificationCodeForBindSecureMobilePhoneResponse) SetStatusCode(v int32) *SendVerificationCodeForBindSecureMobilePhoneResponse {
	s.StatusCode = &v
	return s
}

func (s *SendVerificationCodeForBindSecureMobilePhoneResponse) SetBody(v *SendVerificationCodeForBindSecureMobilePhoneResponseBody) *SendVerificationCodeForBindSecureMobilePhoneResponse {
	s.Body = v
	return s
}

type SendVerificationCodeForEnableRDRequest struct {
	// The mobile phone number that is bound to the newly created account. If you leave this parameter empty, the mobile phone number that is bound to the current account is used.
	//
	// Specify the mobile phone number in the `<Country code>-<Mobile phone number>` format.
	//
	// > Mobile phone numbers in the `86-<Mobile phone number>` format in the Chinese mainland are not supported.
	//
	// example:
	//
	// xx-13900001234
	SecureMobilePhone *string `json:"SecureMobilePhone,omitempty" xml:"SecureMobilePhone,omitempty"`
}

func (s SendVerificationCodeForEnableRDRequest) String() string {
	return tea.Prettify(s)
}

func (s SendVerificationCodeForEnableRDRequest) GoString() string {
	return s.String()
}

func (s *SendVerificationCodeForEnableRDRequest) SetSecureMobilePhone(v string) *SendVerificationCodeForEnableRDRequest {
	s.SecureMobilePhone = &v
	return s
}

type SendVerificationCodeForEnableRDResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// EC2FE94D-A4A2-51A1-A493-5C273A36C46A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SendVerificationCodeForEnableRDResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SendVerificationCodeForEnableRDResponseBody) GoString() string {
	return s.String()
}

func (s *SendVerificationCodeForEnableRDResponseBody) SetRequestId(v string) *SendVerificationCodeForEnableRDResponseBody {
	s.RequestId = &v
	return s
}

type SendVerificationCodeForEnableRDResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SendVerificationCodeForEnableRDResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SendVerificationCodeForEnableRDResponse) String() string {
	return tea.Prettify(s)
}

func (s SendVerificationCodeForEnableRDResponse) GoString() string {
	return s.String()
}

func (s *SendVerificationCodeForEnableRDResponse) SetHeaders(v map[string]*string) *SendVerificationCodeForEnableRDResponse {
	s.Headers = v
	return s
}

func (s *SendVerificationCodeForEnableRDResponse) SetStatusCode(v int32) *SendVerificationCodeForEnableRDResponse {
	s.StatusCode = &v
	return s
}

func (s *SendVerificationCodeForEnableRDResponse) SetBody(v *SendVerificationCodeForEnableRDResponseBody) *SendVerificationCodeForEnableRDResponse {
	s.Body = v
	return s
}

type SetMemberDeletionPermissionRequest struct {
	// Specifies whether to enable the member deletion feature. Valid values:
	//
	// 	- Enabled: enables the member deletion feature.
	//
	// 	- Disabled: disables the member deletion feature.
	//
	// This parameter is required.
	//
	// example:
	//
	// Enabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s SetMemberDeletionPermissionRequest) String() string {
	return tea.Prettify(s)
}

func (s SetMemberDeletionPermissionRequest) GoString() string {
	return s.String()
}

func (s *SetMemberDeletionPermissionRequest) SetStatus(v string) *SetMemberDeletionPermissionRequest {
	s.Status = &v
	return s
}

type SetMemberDeletionPermissionResponseBody struct {
	// The ID of the management account of the resource directory.
	//
	// example:
	//
	// 151266687691****
	ManagementAccountId *string `json:"ManagementAccountId,omitempty" xml:"ManagementAccountId,omitempty"`
	// The status of the member deletion feature. Valid values:
	//
	// 	- Enabled: The feature is enabled.
	//
	// 	- Disabled: The feature is disabled.
	//
	// example:
	//
	// Enabled
	MemberDeletionStatus *string `json:"MemberDeletionStatus,omitempty" xml:"MemberDeletionStatus,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// C55A4CAA-9039-1DDF-91CE-FCC134513D29
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the resource directory.
	//
	// example:
	//
	// rd-3G****
	ResourceDirectoryId *string `json:"ResourceDirectoryId,omitempty" xml:"ResourceDirectoryId,omitempty"`
}

func (s SetMemberDeletionPermissionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetMemberDeletionPermissionResponseBody) GoString() string {
	return s.String()
}

func (s *SetMemberDeletionPermissionResponseBody) SetManagementAccountId(v string) *SetMemberDeletionPermissionResponseBody {
	s.ManagementAccountId = &v
	return s
}

func (s *SetMemberDeletionPermissionResponseBody) SetMemberDeletionStatus(v string) *SetMemberDeletionPermissionResponseBody {
	s.MemberDeletionStatus = &v
	return s
}

func (s *SetMemberDeletionPermissionResponseBody) SetRequestId(v string) *SetMemberDeletionPermissionResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetMemberDeletionPermissionResponseBody) SetResourceDirectoryId(v string) *SetMemberDeletionPermissionResponseBody {
	s.ResourceDirectoryId = &v
	return s
}

type SetMemberDeletionPermissionResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetMemberDeletionPermissionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetMemberDeletionPermissionResponse) String() string {
	return tea.Prettify(s)
}

func (s SetMemberDeletionPermissionResponse) GoString() string {
	return s.String()
}

func (s *SetMemberDeletionPermissionResponse) SetHeaders(v map[string]*string) *SetMemberDeletionPermissionResponse {
	s.Headers = v
	return s
}

func (s *SetMemberDeletionPermissionResponse) SetStatusCode(v int32) *SetMemberDeletionPermissionResponse {
	s.StatusCode = &v
	return s
}

func (s *SetMemberDeletionPermissionResponse) SetBody(v *SetMemberDeletionPermissionResponseBody) *SetMemberDeletionPermissionResponse {
	s.Body = v
	return s
}

type TagResourcesRequest struct {
	// The Alibaba Cloud account IDs of the members.
	//
	// You can specify a maximum of 50 IDs.
	//
	// This parameter is required.
	ResourceId []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	// The type of the objects to which you want to add tags. Valid values:
	//
	// 	- Account: member
	//
	// This parameter is required.
	//
	// example:
	//
	// Account
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The tags.
	//
	// You can specify a maximum of 20 tags.
	//
	// This parameter is required.
	Tag []*TagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s TagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesRequest) GoString() string {
	return s.String()
}

func (s *TagResourcesRequest) SetResourceId(v []*string) *TagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *TagResourcesRequest) SetResourceType(v string) *TagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *TagResourcesRequest) SetTag(v []*TagResourcesRequestTag) *TagResourcesRequest {
	s.Tag = v
	return s
}

type TagResourcesRequestTag struct {
	// The key of the tag.
	//
	// A tag key can be a maximum of 128 characters in length. It cannot contain `http://` or `https://` and cannot start with `acs:` or `aliyun`.
	//
	// example:
	//
	// k1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag.
	//
	// A tag value can be a maximum of 128 characters in length. It cannot contain `http://` or `https://` and cannot start with `acs:`.
	//
	// example:
	//
	// v1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s TagResourcesRequestTag) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesRequestTag) GoString() string {
	return s.String()
}

func (s *TagResourcesRequestTag) SetKey(v string) *TagResourcesRequestTag {
	s.Key = &v
	return s
}

func (s *TagResourcesRequestTag) SetValue(v string) *TagResourcesRequestTag {
	s.Value = &v
	return s
}

type TagResourcesResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// E7747EDF-EDDC-5B38-9B6A-6392B9C92B1C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s TagResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *TagResourcesResponseBody) SetRequestId(v string) *TagResourcesResponseBody {
	s.RequestId = &v
	return s
}

type TagResourcesResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesResponse) GoString() string {
	return s.String()
}

func (s *TagResourcesResponse) SetHeaders(v map[string]*string) *TagResourcesResponse {
	s.Headers = v
	return s
}

func (s *TagResourcesResponse) SetStatusCode(v int32) *TagResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *TagResourcesResponse) SetBody(v *TagResourcesResponseBody) *TagResourcesResponse {
	s.Body = v
	return s
}

type UntagResourcesRequest struct {
	// Specifies whether to remove all tags from the specified members. Valid values:
	//
	// 	- false (default value)
	//
	// 	- true
	//
	// example:
	//
	// false
	All *bool `json:"All,omitempty" xml:"All,omitempty"`
	// The Alibaba Cloud account IDs of the members.
	//
	// You can specify a maximum of 50 IDs.
	//
	// This parameter is required.
	ResourceId []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	// The type of the objects from which you want to remove tags. Valid values:
	//
	// 	- Account: member
	//
	// This parameter is required.
	//
	// example:
	//
	// Account
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The tag keys.
	//
	// You can specify a maximum of 20 tag keys.
	//
	// > If you set the `All` parameter to `true`, you do not need to specify tag keys.
	TagKey []*string `json:"TagKey,omitempty" xml:"TagKey,omitempty" type:"Repeated"`
}

func (s UntagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesRequest) GoString() string {
	return s.String()
}

func (s *UntagResourcesRequest) SetAll(v bool) *UntagResourcesRequest {
	s.All = &v
	return s
}

func (s *UntagResourcesRequest) SetResourceId(v []*string) *UntagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *UntagResourcesRequest) SetResourceType(v string) *UntagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *UntagResourcesRequest) SetTagKey(v []*string) *UntagResourcesRequest {
	s.TagKey = v
	return s
}

type UntagResourcesResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// E7747EDF-EDDC-5B38-9B6A-6392B9C92B1C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UntagResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *UntagResourcesResponseBody) SetRequestId(v string) *UntagResourcesResponseBody {
	s.RequestId = &v
	return s
}

type UntagResourcesResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UntagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UntagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesResponse) GoString() string {
	return s.String()
}

func (s *UntagResourcesResponse) SetHeaders(v map[string]*string) *UntagResourcesResponse {
	s.Headers = v
	return s
}

func (s *UntagResourcesResponse) SetStatusCode(v int32) *UntagResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *UntagResourcesResponse) SetBody(v *UntagResourcesResponseBody) *UntagResourcesResponse {
	s.Body = v
	return s
}

type UpdateAccountRequest struct {
	// The Alibaba Cloud account ID of the member.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12323344****
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// 	- true: performs only a dry run. The system checks items such as whether the member status can be modified and whether security information is configured for the member. If the request does not pass the dry run, an error code is returned.
	//
	// 	- false (default): performs a dry run and performs the actual request.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The new type of the member. Valid values:
	//
	// 	- ResourceAccount: resource account
	//
	// 	- CloudAccount: cloud account
	//
	// > You can specify either `NewDisplayName` or `NewAccountType`.
	//
	// example:
	//
	// ResourceAccount
	NewAccountType *string `json:"NewAccountType,omitempty" xml:"NewAccountType,omitempty"`
	// The new display name of the member.
	//
	// > You can specify either `NewDisplayName` or `NewAccountType`.
	//
	// example:
	//
	// admin
	NewDisplayName *string `json:"NewDisplayName,omitempty" xml:"NewDisplayName,omitempty"`
}

func (s UpdateAccountRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateAccountRequest) GoString() string {
	return s.String()
}

func (s *UpdateAccountRequest) SetAccountId(v string) *UpdateAccountRequest {
	s.AccountId = &v
	return s
}

func (s *UpdateAccountRequest) SetDryRun(v bool) *UpdateAccountRequest {
	s.DryRun = &v
	return s
}

func (s *UpdateAccountRequest) SetNewAccountType(v string) *UpdateAccountRequest {
	s.NewAccountType = &v
	return s
}

func (s *UpdateAccountRequest) SetNewDisplayName(v string) *UpdateAccountRequest {
	s.NewDisplayName = &v
	return s
}

type UpdateAccountResponseBody struct {
	// The information of the member.
	Account *UpdateAccountResponseBodyAccount `json:"Account,omitempty" xml:"Account,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 9B34724D-54B0-4A51-B34D-4512372FE1BE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateAccountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateAccountResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAccountResponseBody) SetAccount(v *UpdateAccountResponseBodyAccount) *UpdateAccountResponseBody {
	s.Account = v
	return s
}

func (s *UpdateAccountResponseBody) SetRequestId(v string) *UpdateAccountResponseBody {
	s.RequestId = &v
	return s
}

type UpdateAccountResponseBodyAccount struct {
	// The Alibaba Cloud account ID of the member.
	//
	// example:
	//
	// 12323344****
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The Alibaba Cloud account name of the member.
	//
	// example:
	//
	// ecs-manager@aliyun.com
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The display name of the member.
	//
	// example:
	//
	// admin
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The ID of the folder.
	//
	// example:
	//
	// fd-bVaRIG****
	FolderId *string `json:"FolderId,omitempty" xml:"FolderId,omitempty"`
	// The way in which the member joins the resource directory. Valid values:
	//
	// 	- invited: The member is invited to join the resource directory.
	//
	// 	- created: The member is directly created in the resource directory.
	//
	// example:
	//
	// created
	JoinMethod *string `json:"JoinMethod,omitempty" xml:"JoinMethod,omitempty"`
	// The time when the member joined the resource directory. The time is displayed in UTC.
	//
	// example:
	//
	// 2015-01-23T12:33:18Z
	JoinTime *string `json:"JoinTime,omitempty" xml:"JoinTime,omitempty"`
	// The time when the member was modified. The time is displayed in UTC.
	//
	// example:
	//
	// 2015-01-23T12:33:18Z
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The ID of the resource directory.
	//
	// example:
	//
	// rd-k3****
	ResourceDirectoryId *string `json:"ResourceDirectoryId,omitempty" xml:"ResourceDirectoryId,omitempty"`
	// The status of the member. Valid values:
	//
	// 	- CreateSuccess: The member is created.
	//
	// 	- InviteSuccess: The member accepts the invitation.
	//
	// 	- Removed: The member is removed.
	//
	// 	- SwitchSuccess: The type of the member is switched.
	//
	// example:
	//
	// CreateSuccess
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The type of the member. Valid values:
	//
	// 	- CloudAccount: cloud account
	//
	// 	- ResourceAccount: resource account
	//
	// example:
	//
	// ResourceAccount
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s UpdateAccountResponseBodyAccount) String() string {
	return tea.Prettify(s)
}

func (s UpdateAccountResponseBodyAccount) GoString() string {
	return s.String()
}

func (s *UpdateAccountResponseBodyAccount) SetAccountId(v string) *UpdateAccountResponseBodyAccount {
	s.AccountId = &v
	return s
}

func (s *UpdateAccountResponseBodyAccount) SetAccountName(v string) *UpdateAccountResponseBodyAccount {
	s.AccountName = &v
	return s
}

func (s *UpdateAccountResponseBodyAccount) SetDisplayName(v string) *UpdateAccountResponseBodyAccount {
	s.DisplayName = &v
	return s
}

func (s *UpdateAccountResponseBodyAccount) SetFolderId(v string) *UpdateAccountResponseBodyAccount {
	s.FolderId = &v
	return s
}

func (s *UpdateAccountResponseBodyAccount) SetJoinMethod(v string) *UpdateAccountResponseBodyAccount {
	s.JoinMethod = &v
	return s
}

func (s *UpdateAccountResponseBodyAccount) SetJoinTime(v string) *UpdateAccountResponseBodyAccount {
	s.JoinTime = &v
	return s
}

func (s *UpdateAccountResponseBodyAccount) SetModifyTime(v string) *UpdateAccountResponseBodyAccount {
	s.ModifyTime = &v
	return s
}

func (s *UpdateAccountResponseBodyAccount) SetResourceDirectoryId(v string) *UpdateAccountResponseBodyAccount {
	s.ResourceDirectoryId = &v
	return s
}

func (s *UpdateAccountResponseBodyAccount) SetStatus(v string) *UpdateAccountResponseBodyAccount {
	s.Status = &v
	return s
}

func (s *UpdateAccountResponseBodyAccount) SetType(v string) *UpdateAccountResponseBodyAccount {
	s.Type = &v
	return s
}

type UpdateAccountResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAccountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAccountResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateAccountResponse) GoString() string {
	return s.String()
}

func (s *UpdateAccountResponse) SetHeaders(v map[string]*string) *UpdateAccountResponse {
	s.Headers = v
	return s
}

func (s *UpdateAccountResponse) SetStatusCode(v int32) *UpdateAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAccountResponse) SetBody(v *UpdateAccountResponseBody) *UpdateAccountResponse {
	s.Body = v
	return s
}

type UpdateControlPolicyRequest struct {
	// The new description of the access control policy.
	//
	// The description must be 1 to 1,024 characters in length. The description can contain letters, digits, underscores (_), and hyphens (-) and must start with a letter.
	//
	// example:
	//
	// ExampleControlPolicy
	NewDescription *string `json:"NewDescription,omitempty" xml:"NewDescription,omitempty"`
	// The new document of the access control policy.
	//
	// The document can be a maximum of 4,096 characters in length.
	//
	// For more information about the languages of access control policies, see [Languages of access control policies](https://help.aliyun.com/document_detail/179096.html).
	//
	// For more information about the examples of access control policies, see [Examples of custom access control policies](https://help.aliyun.com/document_detail/181474.html).
	//
	// example:
	//
	// {"Version":"1","Statement":[{"Effect":"Deny","Action":["ram:UpdateRole","ram:DeleteRole","ram:AttachPolicyToRole","ram:DetachPolicyFromRole"],"Resource":"acs:ram:*:*:role/ResourceDirectoryAccountAccessRole"}]}
	NewPolicyDocument *string `json:"NewPolicyDocument,omitempty" xml:"NewPolicyDocument,omitempty"`
	// The new name of the access control policy.
	//
	// The name must be 1 to 128 characters in length. The name can contain letters, digits, and hyphens (-) and must start with a letter.
	//
	// example:
	//
	// NewControlPolicy
	NewPolicyName *string `json:"NewPolicyName,omitempty" xml:"NewPolicyName,omitempty"`
	// The ID of the access control policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// cp-jExXAqIYkwHN****
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
}

func (s UpdateControlPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateControlPolicyRequest) GoString() string {
	return s.String()
}

func (s *UpdateControlPolicyRequest) SetNewDescription(v string) *UpdateControlPolicyRequest {
	s.NewDescription = &v
	return s
}

func (s *UpdateControlPolicyRequest) SetNewPolicyDocument(v string) *UpdateControlPolicyRequest {
	s.NewPolicyDocument = &v
	return s
}

func (s *UpdateControlPolicyRequest) SetNewPolicyName(v string) *UpdateControlPolicyRequest {
	s.NewPolicyName = &v
	return s
}

func (s *UpdateControlPolicyRequest) SetPolicyId(v string) *UpdateControlPolicyRequest {
	s.PolicyId = &v
	return s
}

type UpdateControlPolicyResponseBody struct {
	// The details of the access control policy.
	ControlPolicy *UpdateControlPolicyResponseBodyControlPolicy `json:"ControlPolicy,omitempty" xml:"ControlPolicy,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 2DFCE4C9-04A9-4C83-BB14-FE791275EC53
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateControlPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateControlPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateControlPolicyResponseBody) SetControlPolicy(v *UpdateControlPolicyResponseBodyControlPolicy) *UpdateControlPolicyResponseBody {
	s.ControlPolicy = v
	return s
}

func (s *UpdateControlPolicyResponseBody) SetRequestId(v string) *UpdateControlPolicyResponseBody {
	s.RequestId = &v
	return s
}

type UpdateControlPolicyResponseBodyControlPolicy struct {
	// The number of times that the access control policy is referenced.
	//
	// example:
	//
	// 0
	AttachmentCount *string `json:"AttachmentCount,omitempty" xml:"AttachmentCount,omitempty"`
	// The time when the access control policy was created.
	//
	// example:
	//
	// 2021-03-18T09:24:19Z
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	// The description of the access control policy.
	//
	// example:
	//
	// ExampleControlPolicy
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The effective scope of the access control policy. Valid values:
	//
	// 	- All: The access control policy is in effect for Alibaba Cloud accounts, RAM users, and RAM roles.
	//
	// 	- RAM: The access control policy is in effect only for RAM users and RAM roles.
	//
	// example:
	//
	// RAM
	EffectScope *string `json:"EffectScope,omitempty" xml:"EffectScope,omitempty"`
	// The ID of the access control policy.
	//
	// example:
	//
	// cp-jExXAqIYkwHN****
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// The name of the access control policy.
	//
	// example:
	//
	// NewControlPolicy
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	// The type of the access control policy. Valid values:
	//
	// 	- System: system access control policy
	//
	// 	- Custom: custom access control policy
	//
	// example:
	//
	// Custom
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
	// The time when the access control policy was updated.
	//
	// example:
	//
	// 2021-03-18T10:04:55Z
	UpdateDate *string `json:"UpdateDate,omitempty" xml:"UpdateDate,omitempty"`
}

func (s UpdateControlPolicyResponseBodyControlPolicy) String() string {
	return tea.Prettify(s)
}

func (s UpdateControlPolicyResponseBodyControlPolicy) GoString() string {
	return s.String()
}

func (s *UpdateControlPolicyResponseBodyControlPolicy) SetAttachmentCount(v string) *UpdateControlPolicyResponseBodyControlPolicy {
	s.AttachmentCount = &v
	return s
}

func (s *UpdateControlPolicyResponseBodyControlPolicy) SetCreateDate(v string) *UpdateControlPolicyResponseBodyControlPolicy {
	s.CreateDate = &v
	return s
}

func (s *UpdateControlPolicyResponseBodyControlPolicy) SetDescription(v string) *UpdateControlPolicyResponseBodyControlPolicy {
	s.Description = &v
	return s
}

func (s *UpdateControlPolicyResponseBodyControlPolicy) SetEffectScope(v string) *UpdateControlPolicyResponseBodyControlPolicy {
	s.EffectScope = &v
	return s
}

func (s *UpdateControlPolicyResponseBodyControlPolicy) SetPolicyId(v string) *UpdateControlPolicyResponseBodyControlPolicy {
	s.PolicyId = &v
	return s
}

func (s *UpdateControlPolicyResponseBodyControlPolicy) SetPolicyName(v string) *UpdateControlPolicyResponseBodyControlPolicy {
	s.PolicyName = &v
	return s
}

func (s *UpdateControlPolicyResponseBodyControlPolicy) SetPolicyType(v string) *UpdateControlPolicyResponseBodyControlPolicy {
	s.PolicyType = &v
	return s
}

func (s *UpdateControlPolicyResponseBodyControlPolicy) SetUpdateDate(v string) *UpdateControlPolicyResponseBodyControlPolicy {
	s.UpdateDate = &v
	return s
}

type UpdateControlPolicyResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateControlPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateControlPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateControlPolicyResponse) GoString() string {
	return s.String()
}

func (s *UpdateControlPolicyResponse) SetHeaders(v map[string]*string) *UpdateControlPolicyResponse {
	s.Headers = v
	return s
}

func (s *UpdateControlPolicyResponse) SetStatusCode(v int32) *UpdateControlPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateControlPolicyResponse) SetBody(v *UpdateControlPolicyResponseBody) *UpdateControlPolicyResponse {
	s.Body = v
	return s
}

type UpdateFolderRequest struct {
	// The ID of the folder.
	//
	// This parameter is required.
	//
	// example:
	//
	// fd-u8B321****
	FolderId *string `json:"FolderId,omitempty" xml:"FolderId,omitempty"`
	// The new name of the folder.
	//
	// The name must be 1 to 24 characters in length and can contain letters, digits, underscores (_), periods (.), and hyphens (-).
	//
	// This parameter is required.
	//
	// example:
	//
	// rdFolder
	NewFolderName *string `json:"NewFolderName,omitempty" xml:"NewFolderName,omitempty"`
}

func (s UpdateFolderRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateFolderRequest) GoString() string {
	return s.String()
}

func (s *UpdateFolderRequest) SetFolderId(v string) *UpdateFolderRequest {
	s.FolderId = &v
	return s
}

func (s *UpdateFolderRequest) SetNewFolderName(v string) *UpdateFolderRequest {
	s.NewFolderName = &v
	return s
}

type UpdateFolderResponseBody struct {
	// The information about the folder.
	Folder *UpdateFolderResponseBodyFolder `json:"Folder,omitempty" xml:"Folder,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// C2CBCA30-C8DD-423E-B4AD-4FB694C9180C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateFolderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateFolderResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateFolderResponseBody) SetFolder(v *UpdateFolderResponseBodyFolder) *UpdateFolderResponseBody {
	s.Folder = v
	return s
}

func (s *UpdateFolderResponseBody) SetRequestId(v string) *UpdateFolderResponseBody {
	s.RequestId = &v
	return s
}

type UpdateFolderResponseBodyFolder struct {
	// The time when the folder was created.
	//
	// example:
	//
	// 2019-02-19T09:34:50.757Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the folder.
	//
	// example:
	//
	// fd-u8B321****
	FolderId *string `json:"FolderId,omitempty" xml:"FolderId,omitempty"`
	// The name of the folder.
	//
	// example:
	//
	// rdFolder
	FolderName *string `json:"FolderName,omitempty" xml:"FolderName,omitempty"`
	// The ID of the parent folder.
	//
	// example:
	//
	// r-b1****
	ParentFolderId *string `json:"ParentFolderId,omitempty" xml:"ParentFolderId,omitempty"`
}

func (s UpdateFolderResponseBodyFolder) String() string {
	return tea.Prettify(s)
}

func (s UpdateFolderResponseBodyFolder) GoString() string {
	return s.String()
}

func (s *UpdateFolderResponseBodyFolder) SetCreateTime(v string) *UpdateFolderResponseBodyFolder {
	s.CreateTime = &v
	return s
}

func (s *UpdateFolderResponseBodyFolder) SetFolderId(v string) *UpdateFolderResponseBodyFolder {
	s.FolderId = &v
	return s
}

func (s *UpdateFolderResponseBodyFolder) SetFolderName(v string) *UpdateFolderResponseBodyFolder {
	s.FolderName = &v
	return s
}

func (s *UpdateFolderResponseBodyFolder) SetParentFolderId(v string) *UpdateFolderResponseBodyFolder {
	s.ParentFolderId = &v
	return s
}

type UpdateFolderResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateFolderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateFolderResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateFolderResponse) GoString() string {
	return s.String()
}

func (s *UpdateFolderResponse) SetHeaders(v map[string]*string) *UpdateFolderResponse {
	s.Headers = v
	return s
}

func (s *UpdateFolderResponse) SetStatusCode(v int32) *UpdateFolderResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateFolderResponse) SetBody(v *UpdateFolderResponseBody) *UpdateFolderResponse {
	s.Body = v
	return s
}

type UpdateMessageContactRequest struct {
	// The ID of the contact.
	//
	// example:
	//
	// c-qL4HqKONzOM7****
	ContactId *string `json:"ContactId,omitempty" xml:"ContactId,omitempty"`
	// The email address of the contact.
	//
	// After you specify an email address, you need to call [SendEmailVerificationForMessageContact](~~SendEmailVerificationForMessageContact~~) to send verification information to the email address. After the verification is passed, the email address takes effect.
	//
	// example:
	//
	// someone***@example.com
	EmailAddress *string `json:"EmailAddress,omitempty" xml:"EmailAddress,omitempty"`
	// The types of messages received by the contact.
	MessageTypes []*string `json:"MessageTypes,omitempty" xml:"MessageTypes,omitempty" type:"Repeated"`
	// The name of the contact.
	//
	// example:
	//
	// tom
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The mobile phone number of the contact.
	//
	// Specify the mobile phone number in the `<Country code>-<Mobile phone number>` format.
	//
	// After you specify a mobile phone number, you need to call [SendPhoneVerificationForMessageContact](~~SendPhoneVerificationForMessageContact~~) to send verification information to the mobile phone number. After the verification is passed, the mobile phone number takes effect.
	//
	// example:
	//
	// 86-139****1234
	PhoneNumber *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	// The job title of the contact.
	//
	// Valid values:
	//
	// 	- FinanceDirector
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- TechnicalDirector
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- MaintenanceDirector
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- CEO
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- ProjectDirector
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- Other
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// TechnicalDirector
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s UpdateMessageContactRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateMessageContactRequest) GoString() string {
	return s.String()
}

func (s *UpdateMessageContactRequest) SetContactId(v string) *UpdateMessageContactRequest {
	s.ContactId = &v
	return s
}

func (s *UpdateMessageContactRequest) SetEmailAddress(v string) *UpdateMessageContactRequest {
	s.EmailAddress = &v
	return s
}

func (s *UpdateMessageContactRequest) SetMessageTypes(v []*string) *UpdateMessageContactRequest {
	s.MessageTypes = v
	return s
}

func (s *UpdateMessageContactRequest) SetName(v string) *UpdateMessageContactRequest {
	s.Name = &v
	return s
}

func (s *UpdateMessageContactRequest) SetPhoneNumber(v string) *UpdateMessageContactRequest {
	s.PhoneNumber = &v
	return s
}

func (s *UpdateMessageContactRequest) SetTitle(v string) *UpdateMessageContactRequest {
	s.Title = &v
	return s
}

type UpdateMessageContactResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 9B34724D-54B0-4A51-B34D-4512372FE1BE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateMessageContactResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateMessageContactResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateMessageContactResponseBody) SetRequestId(v string) *UpdateMessageContactResponseBody {
	s.RequestId = &v
	return s
}

type UpdateMessageContactResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateMessageContactResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateMessageContactResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateMessageContactResponse) GoString() string {
	return s.String()
}

func (s *UpdateMessageContactResponse) SetHeaders(v map[string]*string) *UpdateMessageContactResponse {
	s.Headers = v
	return s
}

func (s *UpdateMessageContactResponse) SetStatusCode(v int32) *UpdateMessageContactResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateMessageContactResponse) SetBody(v *UpdateMessageContactResponseBody) *UpdateMessageContactResponse {
	s.Body = v
	return s
}

type Client struct {
	openapi.Client
}

func NewClient(config *openapi.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapi.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.SignatureAlgorithm = tea.String("v2")
	client.EndpointRule = tea.String("")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("resourcedirectorymaster"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !tea.BoolValue(util.Empty(endpoint)) {
		_result = endpoint
		return _result, _err
	}

	if !tea.BoolValue(util.IsUnset(endpointMap)) && !tea.BoolValue(util.Empty(endpointMap[tea.StringValue(regionId)])) {
		_result = endpointMap[tea.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := endpointutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Accepts an invitation.
//
// Description:
//
// After an invited Alibaba Cloud account joins a resource directory, it becomes a member of the resource directory. By default, the name of the invited Alibaba Cloud account is used as the display name of the account in the resource directory.
//
// @param request - AcceptHandshakeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AcceptHandshakeResponse
func (client *Client) AcceptHandshakeWithOptions(request *AcceptHandshakeRequest, runtime *util.RuntimeOptions) (_result *AcceptHandshakeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.HandshakeId)) {
		query["HandshakeId"] = request.HandshakeId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AcceptHandshake"),
		Version:     tea.String("2022-04-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AcceptHandshakeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Accepts an invitation.
//
// Description:
//
// After an invited Alibaba Cloud account joins a resource directory, it becomes a member of the resource directory. By default, the name of the invited Alibaba Cloud account is used as the display name of the account in the resource directory.
//
// @param request - AcceptHandshakeRequest
//
// @return AcceptHandshakeResponse
func (client *Client) AcceptHandshake(request *AcceptHandshakeRequest) (_result *AcceptHandshakeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AcceptHandshakeResponse{}
	_body, _err := client.AcceptHandshakeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds a contact.
//
// @param request - AddMessageContactRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddMessageContactResponse
func (client *Client) AddMessageContactWithOptions(request *AddMessageContactRequest, runtime *util.RuntimeOptions) (_result *AddMessageContactResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EmailAddress)) {
		query["EmailAddress"] = request.EmailAddress
	}

	if !tea.BoolValue(util.IsUnset(request.MessageTypes)) {
		query["MessageTypes"] = request.MessageTypes
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.PhoneNumber)) {
		query["PhoneNumber"] = request.PhoneNumber
	}

	if !tea.BoolValue(util.IsUnset(request.Title)) {
		query["Title"] = request.Title
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AddMessageContact"),
		Version:     tea.String("2022-04-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddMessageContactResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds a contact.
//
// @param request - AddMessageContactRequest
//
// @return AddMessageContactResponse
func (client *Client) AddMessageContact(request *AddMessageContactRequest) (_result *AddMessageContactResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddMessageContactResponse{}
	_body, _err := client.AddMessageContactWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Binds a contact to a resource directory, folder, or member.
//
// @param request - AssociateMembersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AssociateMembersResponse
func (client *Client) AssociateMembersWithOptions(request *AssociateMembersRequest, runtime *util.RuntimeOptions) (_result *AssociateMembersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ContactId)) {
		query["ContactId"] = request.ContactId
	}

	if !tea.BoolValue(util.IsUnset(request.Members)) {
		query["Members"] = request.Members
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AssociateMembers"),
		Version:     tea.String("2022-04-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AssociateMembersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Binds a contact to a resource directory, folder, or member.
//
// @param request - AssociateMembersRequest
//
// @return AssociateMembersResponse
func (client *Client) AssociateMembers(request *AssociateMembersRequest) (_result *AssociateMembersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AssociateMembersResponse{}
	_body, _err := client.AssociateMembersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Attaches an access control policy.
//
// Description:
//
// After you attach a custom access control policy, the operations performed on resources by using members are limited by the policy. Make sure that the attached policy meets your expectations. Otherwise, your business may be affected.
//
// By default, the system access control policy FullAliyunAccess is attached to each folder and member.
//
// The access control policy that is attached to a folder also applies to all its subfolders and all members in the subfolders.
//
// A maximum of 10 access control policies can be attached to a folder or member.
//
// @param request - AttachControlPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AttachControlPolicyResponse
func (client *Client) AttachControlPolicyWithOptions(request *AttachControlPolicyRequest, runtime *util.RuntimeOptions) (_result *AttachControlPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PolicyId)) {
		query["PolicyId"] = request.PolicyId
	}

	if !tea.BoolValue(util.IsUnset(request.TargetId)) {
		query["TargetId"] = request.TargetId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AttachControlPolicy"),
		Version:     tea.String("2022-04-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AttachControlPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Attaches an access control policy.
//
// Description:
//
// After you attach a custom access control policy, the operations performed on resources by using members are limited by the policy. Make sure that the attached policy meets your expectations. Otherwise, your business may be affected.
//
// By default, the system access control policy FullAliyunAccess is attached to each folder and member.
//
// The access control policy that is attached to a folder also applies to all its subfolders and all members in the subfolders.
//
// A maximum of 10 access control policies can be attached to a folder or member.
//
// @param request - AttachControlPolicyRequest
//
// @return AttachControlPolicyResponse
func (client *Client) AttachControlPolicy(request *AttachControlPolicyRequest) (_result *AttachControlPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AttachControlPolicyResponse{}
	_body, _err := client.AttachControlPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Binds a mobile phone number to a member of the resource account type in a resource directory for security purposes.
//
// Description:
//
// You can call this API operation only to bind a mobile phone number to a member of the resource account type. You cannot call this API operation to change the mobile phone number that is bound to a member of the resource account type.
//
// To ensure that the system can record the operators of management operations, you must use a RAM user or RAM role to which the AliyunResourceDirectoryFullAccess policy is attached within the management account of your resource directory to call this API operation.
//
// @param request - BindSecureMobilePhoneRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BindSecureMobilePhoneResponse
func (client *Client) BindSecureMobilePhoneWithOptions(request *BindSecureMobilePhoneRequest, runtime *util.RuntimeOptions) (_result *BindSecureMobilePhoneResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountId)) {
		query["AccountId"] = request.AccountId
	}

	if !tea.BoolValue(util.IsUnset(request.SecureMobilePhone)) {
		query["SecureMobilePhone"] = request.SecureMobilePhone
	}

	if !tea.BoolValue(util.IsUnset(request.VerificationCode)) {
		query["VerificationCode"] = request.VerificationCode
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("BindSecureMobilePhone"),
		Version:     tea.String("2022-04-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &BindSecureMobilePhoneResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Binds a mobile phone number to a member of the resource account type in a resource directory for security purposes.
//
// Description:
//
// You can call this API operation only to bind a mobile phone number to a member of the resource account type. You cannot call this API operation to change the mobile phone number that is bound to a member of the resource account type.
//
// To ensure that the system can record the operators of management operations, you must use a RAM user or RAM role to which the AliyunResourceDirectoryFullAccess policy is attached within the management account of your resource directory to call this API operation.
//
// @param request - BindSecureMobilePhoneRequest
//
// @return BindSecureMobilePhoneResponse
func (client *Client) BindSecureMobilePhone(request *BindSecureMobilePhoneRequest) (_result *BindSecureMobilePhoneResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BindSecureMobilePhoneResponse{}
	_body, _err := client.BindSecureMobilePhoneWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Cancels the email address change of a member.
//
// @param request - CancelChangeAccountEmailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelChangeAccountEmailResponse
func (client *Client) CancelChangeAccountEmailWithOptions(request *CancelChangeAccountEmailRequest, runtime *util.RuntimeOptions) (_result *CancelChangeAccountEmailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountId)) {
		query["AccountId"] = request.AccountId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CancelChangeAccountEmail"),
		Version:     tea.String("2022-04-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CancelChangeAccountEmailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Cancels the email address change of a member.
//
// @param request - CancelChangeAccountEmailRequest
//
// @return CancelChangeAccountEmailResponse
func (client *Client) CancelChangeAccountEmail(request *CancelChangeAccountEmailRequest) (_result *CancelChangeAccountEmailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CancelChangeAccountEmailResponse{}
	_body, _err := client.CancelChangeAccountEmailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Cancels an invitation.
//
// @param request - CancelHandshakeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelHandshakeResponse
func (client *Client) CancelHandshakeWithOptions(request *CancelHandshakeRequest, runtime *util.RuntimeOptions) (_result *CancelHandshakeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.HandshakeId)) {
		query["HandshakeId"] = request.HandshakeId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CancelHandshake"),
		Version:     tea.String("2022-04-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CancelHandshakeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Cancels an invitation.
//
// @param request - CancelHandshakeRequest
//
// @return CancelHandshakeResponse
func (client *Client) CancelHandshake(request *CancelHandshakeRequest) (_result *CancelHandshakeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CancelHandshakeResponse{}
	_body, _err := client.CancelHandshakeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Cancels the update of the mobile phone number or email address of a contact.
//
// @param request - CancelMessageContactUpdateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelMessageContactUpdateResponse
func (client *Client) CancelMessageContactUpdateWithOptions(request *CancelMessageContactUpdateRequest, runtime *util.RuntimeOptions) (_result *CancelMessageContactUpdateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ContactId)) {
		query["ContactId"] = request.ContactId
	}

	if !tea.BoolValue(util.IsUnset(request.EmailAddress)) {
		query["EmailAddress"] = request.EmailAddress
	}

	if !tea.BoolValue(util.IsUnset(request.PhoneNumber)) {
		query["PhoneNumber"] = request.PhoneNumber
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CancelMessageContactUpdate"),
		Version:     tea.String("2022-04-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CancelMessageContactUpdateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Cancels the update of the mobile phone number or email address of a contact.
//
// @param request - CancelMessageContactUpdateRequest
//
// @return CancelMessageContactUpdateResponse
func (client *Client) CancelMessageContactUpdate(request *CancelMessageContactUpdateRequest) (_result *CancelMessageContactUpdateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CancelMessageContactUpdateResponse{}
	_body, _err := client.CancelMessageContactUpdateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Changes the email address that is bound to a member.
//
// @param request - ChangeAccountEmailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChangeAccountEmailResponse
func (client *Client) ChangeAccountEmailWithOptions(request *ChangeAccountEmailRequest, runtime *util.RuntimeOptions) (_result *ChangeAccountEmailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountId)) {
		query["AccountId"] = request.AccountId
	}

	if !tea.BoolValue(util.IsUnset(request.Email)) {
		query["Email"] = request.Email
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ChangeAccountEmail"),
		Version:     tea.String("2022-04-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ChangeAccountEmailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Changes the email address that is bound to a member.
//
// @param request - ChangeAccountEmailRequest
//
// @return ChangeAccountEmailResponse
func (client *Client) ChangeAccountEmail(request *ChangeAccountEmailRequest) (_result *ChangeAccountEmailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ChangeAccountEmailResponse{}
	_body, _err := client.ChangeAccountEmailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Performs a member deletion check.
//
// Description:
//
// Before you delete a member, you must call this API operation to check whether the member can be deleted.
//
// @param request - CheckAccountDeleteRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckAccountDeleteResponse
func (client *Client) CheckAccountDeleteWithOptions(request *CheckAccountDeleteRequest, runtime *util.RuntimeOptions) (_result *CheckAccountDeleteResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountId)) {
		query["AccountId"] = request.AccountId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CheckAccountDelete"),
		Version:     tea.String("2022-04-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CheckAccountDeleteResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Performs a member deletion check.
//
// Description:
//
// Before you delete a member, you must call this API operation to check whether the member can be deleted.
//
// @param request - CheckAccountDeleteRequest
//
// @return CheckAccountDeleteResponse
func (client *Client) CheckAccountDelete(request *CheckAccountDeleteRequest) (_result *CheckAccountDeleteResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CheckAccountDeleteResponse{}
	_body, _err := client.CheckAccountDeleteWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a custom access control policy.
//
// @param request - CreateControlPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateControlPolicyResponse
func (client *Client) CreateControlPolicyWithOptions(request *CreateControlPolicyRequest, runtime *util.RuntimeOptions) (_result *CreateControlPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.EffectScope)) {
		query["EffectScope"] = request.EffectScope
	}

	if !tea.BoolValue(util.IsUnset(request.PolicyDocument)) {
		query["PolicyDocument"] = request.PolicyDocument
	}

	if !tea.BoolValue(util.IsUnset(request.PolicyName)) {
		query["PolicyName"] = request.PolicyName
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateControlPolicy"),
		Version:     tea.String("2022-04-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateControlPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a custom access control policy.
//
// @param request - CreateControlPolicyRequest
//
// @return CreateControlPolicyResponse
func (client *Client) CreateControlPolicy(request *CreateControlPolicyRequest) (_result *CreateControlPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateControlPolicyResponse{}
	_body, _err := client.CreateControlPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a folder.
//
// Description:
//
// A maximum of five levels of folders can be created under the Root folder.
//
// @param request - CreateFolderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateFolderResponse
func (client *Client) CreateFolderWithOptions(request *CreateFolderRequest, runtime *util.RuntimeOptions) (_result *CreateFolderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FolderName)) {
		query["FolderName"] = request.FolderName
	}

	if !tea.BoolValue(util.IsUnset(request.ParentFolderId)) {
		query["ParentFolderId"] = request.ParentFolderId
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateFolder"),
		Version:     tea.String("2022-04-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateFolderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a folder.
//
// Description:
//
// A maximum of five levels of folders can be created under the Root folder.
//
// @param request - CreateFolderRequest
//
// @return CreateFolderResponse
func (client *Client) CreateFolder(request *CreateFolderRequest) (_result *CreateFolderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateFolderResponse{}
	_body, _err := client.CreateFolderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a member of the resource account type.
//
// Description:
//
// A member serves as a container for resources and is also an organizational unit in a resource directory. A member indicates a project or application. The resources of different members are isolated.
//
// This topic provides an example on how to call the API operation to create a member in the `fd-r23M55****` folder. The display name of the member is `Dev`, and the prefix for the Alibaba Cloud account name of the member is `alice`.
//
// @param request - CreateResourceAccountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateResourceAccountResponse
func (client *Client) CreateResourceAccountWithOptions(request *CreateResourceAccountRequest, runtime *util.RuntimeOptions) (_result *CreateResourceAccountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountNamePrefix)) {
		query["AccountNamePrefix"] = request.AccountNamePrefix
	}

	if !tea.BoolValue(util.IsUnset(request.DisplayName)) {
		query["DisplayName"] = request.DisplayName
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.ParentFolderId)) {
		query["ParentFolderId"] = request.ParentFolderId
	}

	if !tea.BoolValue(util.IsUnset(request.PayerAccountId)) {
		query["PayerAccountId"] = request.PayerAccountId
	}

	if !tea.BoolValue(util.IsUnset(request.ResellAccountType)) {
		query["ResellAccountType"] = request.ResellAccountType
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateResourceAccount"),
		Version:     tea.String("2022-04-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateResourceAccountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a member of the resource account type.
//
// Description:
//
// A member serves as a container for resources and is also an organizational unit in a resource directory. A member indicates a project or application. The resources of different members are isolated.
//
// This topic provides an example on how to call the API operation to create a member in the `fd-r23M55****` folder. The display name of the member is `Dev`, and the prefix for the Alibaba Cloud account name of the member is `alice`.
//
// @param request - CreateResourceAccountRequest
//
// @return CreateResourceAccountResponse
func (client *Client) CreateResourceAccount(request *CreateResourceAccountRequest) (_result *CreateResourceAccountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateResourceAccountResponse{}
	_body, _err := client.CreateResourceAccountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Rejects an invitation.
//
// @param request - DeclineHandshakeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeclineHandshakeResponse
func (client *Client) DeclineHandshakeWithOptions(request *DeclineHandshakeRequest, runtime *util.RuntimeOptions) (_result *DeclineHandshakeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.HandshakeId)) {
		query["HandshakeId"] = request.HandshakeId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeclineHandshake"),
		Version:     tea.String("2022-04-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeclineHandshakeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Rejects an invitation.
//
// @param request - DeclineHandshakeRequest
//
// @return DeclineHandshakeResponse
func (client *Client) DeclineHandshake(request *DeclineHandshakeRequest) (_result *DeclineHandshakeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeclineHandshakeResponse{}
	_body, _err := client.DeclineHandshakeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a member of the resource account type.
//
// Description:
//
// Before you delete a member, we recommend that you call the [CheckAccountDelete](~~CheckAccountDelete~~) and [GetAccountDeletionCheckResult](~~GetAccountDeletionCheckResult~~) operations to check whether the member meets deletion requirements. You can call the DeleteAccount operation to delete only members that meet the deletion requirements.
//
// After you submit a deletion request for a member, you can call the [GetAccountDeletionStatus](~~GetAccountDeletionStatus~~) operation to query the deletion status of the member. After a member is deleted, the resources and data within the member are deleted, and you can no longer use the member to log on to the Alibaba Cloud Management Console. In addition, the member cannot be recovered. Proceed with caution. For more information about how to delete a member, see [Delete a member of the resource account type](https://help.aliyun.com/document_detail/446078.html).
//
// @param tmpReq - DeleteAccountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAccountResponse
func (client *Client) DeleteAccountWithOptions(tmpReq *DeleteAccountRequest, runtime *util.RuntimeOptions) (_result *DeleteAccountResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &DeleteAccountShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.AbandonableCheckId)) {
		request.AbandonableCheckIdShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AbandonableCheckId, tea.String("AbandonableCheckId"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AbandonableCheckIdShrink)) {
		query["AbandonableCheckId"] = request.AbandonableCheckIdShrink
	}

	if !tea.BoolValue(util.IsUnset(request.AccountId)) {
		query["AccountId"] = request.AccountId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteAccount"),
		Version:     tea.String("2022-04-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteAccountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a member of the resource account type.
//
// Description:
//
// Before you delete a member, we recommend that you call the [CheckAccountDelete](~~CheckAccountDelete~~) and [GetAccountDeletionCheckResult](~~GetAccountDeletionCheckResult~~) operations to check whether the member meets deletion requirements. You can call the DeleteAccount operation to delete only members that meet the deletion requirements.
//
// After you submit a deletion request for a member, you can call the [GetAccountDeletionStatus](~~GetAccountDeletionStatus~~) operation to query the deletion status of the member. After a member is deleted, the resources and data within the member are deleted, and you can no longer use the member to log on to the Alibaba Cloud Management Console. In addition, the member cannot be recovered. Proceed with caution. For more information about how to delete a member, see [Delete a member of the resource account type](https://help.aliyun.com/document_detail/446078.html).
//
// @param request - DeleteAccountRequest
//
// @return DeleteAccountResponse
func (client *Client) DeleteAccount(request *DeleteAccountRequest) (_result *DeleteAccountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteAccountResponse{}
	_body, _err := client.DeleteAccountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a custom access control policy.
//
// Description:
//
// If you want to delete a custom access control policy that is attached to folders or members, you must call the [DetachControlPolicy](~~DetachControlPolicy~~) operation to detach the policy before you delete it.
//
// @param request - DeleteControlPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteControlPolicyResponse
func (client *Client) DeleteControlPolicyWithOptions(request *DeleteControlPolicyRequest, runtime *util.RuntimeOptions) (_result *DeleteControlPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PolicyId)) {
		query["PolicyId"] = request.PolicyId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteControlPolicy"),
		Version:     tea.String("2022-04-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteControlPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a custom access control policy.
//
// Description:
//
// If you want to delete a custom access control policy that is attached to folders or members, you must call the [DetachControlPolicy](~~DetachControlPolicy~~) operation to detach the policy before you delete it.
//
// @param request - DeleteControlPolicyRequest
//
// @return DeleteControlPolicyResponse
func (client *Client) DeleteControlPolicy(request *DeleteControlPolicyRequest) (_result *DeleteControlPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteControlPolicyResponse{}
	_body, _err := client.DeleteControlPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a folder.
//
// Description:
//
// Before you delete a folder, you must make sure that the folder does not contain members or subfolders.
//
// @param request - DeleteFolderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteFolderResponse
func (client *Client) DeleteFolderWithOptions(request *DeleteFolderRequest, runtime *util.RuntimeOptions) (_result *DeleteFolderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FolderId)) {
		query["FolderId"] = request.FolderId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteFolder"),
		Version:     tea.String("2022-04-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteFolderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a folder.
//
// Description:
//
// Before you delete a folder, you must make sure that the folder does not contain members or subfolders.
//
// @param request - DeleteFolderRequest
//
// @return DeleteFolderResponse
func (client *Client) DeleteFolder(request *DeleteFolderRequest) (_result *DeleteFolderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteFolderResponse{}
	_body, _err := client.DeleteFolderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a contact.
//
// @param request - DeleteMessageContactRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteMessageContactResponse
func (client *Client) DeleteMessageContactWithOptions(request *DeleteMessageContactRequest, runtime *util.RuntimeOptions) (_result *DeleteMessageContactResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ContactId)) {
		query["ContactId"] = request.ContactId
	}

	if !tea.BoolValue(util.IsUnset(request.RetainContactInMembers)) {
		query["RetainContactInMembers"] = request.RetainContactInMembers
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteMessageContact"),
		Version:     tea.String("2022-04-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteMessageContactResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a contact.
//
// @param request - DeleteMessageContactRequest
//
// @return DeleteMessageContactResponse
func (client *Client) DeleteMessageContact(request *DeleteMessageContactRequest) (_result *DeleteMessageContactResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteMessageContactResponse{}
	_body, _err := client.DeleteMessageContactWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Removes a delegated administrator account for a trusted service.
//
// Description:
//
// If the delegated administrator account that you want to remove has historical management tasks in the related trusted service, the trusted service may be affected after the delegated administrator account is removed. Therefore, proceed with caution.
//
// @param request - DeregisterDelegatedAdministratorRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeregisterDelegatedAdministratorResponse
func (client *Client) DeregisterDelegatedAdministratorWithOptions(request *DeregisterDelegatedAdministratorRequest, runtime *util.RuntimeOptions) (_result *DeregisterDelegatedAdministratorResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountId)) {
		query["AccountId"] = request.AccountId
	}

	if !tea.BoolValue(util.IsUnset(request.ServicePrincipal)) {
		query["ServicePrincipal"] = request.ServicePrincipal
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeregisterDelegatedAdministrator"),
		Version:     tea.String("2022-04-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeregisterDelegatedAdministratorResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes a delegated administrator account for a trusted service.
//
// Description:
//
// If the delegated administrator account that you want to remove has historical management tasks in the related trusted service, the trusted service may be affected after the delegated administrator account is removed. Therefore, proceed with caution.
//
// @param request - DeregisterDelegatedAdministratorRequest
//
// @return DeregisterDelegatedAdministratorResponse
func (client *Client) DeregisterDelegatedAdministrator(request *DeregisterDelegatedAdministratorRequest) (_result *DeregisterDelegatedAdministratorResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeregisterDelegatedAdministratorResponse{}
	_body, _err := client.DeregisterDelegatedAdministratorWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Disables a resource directory. This operation cannot be undone. Therefore, proceed with caution.
//
// Description:
//
// Before you disable a resource directory, you must make sure that the following requirements are met:
//
// 	- All members of the cloud account type in the resource directory are removed. You can call the [RemoveCloudAccount](~~RemoveCloudAccount~~) operation to remove a member of the cloud account type.
//
// 	- All folders except the Root folder are deleted from the resource directory. You can call the [DeleteFolder](~~DeleteFolder~~) operation to delete a folder.
//
// @param request - DestroyResourceDirectoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DestroyResourceDirectoryResponse
func (client *Client) DestroyResourceDirectoryWithOptions(runtime *util.RuntimeOptions) (_result *DestroyResourceDirectoryResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("DestroyResourceDirectory"),
		Version:     tea.String("2022-04-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DestroyResourceDirectoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disables a resource directory. This operation cannot be undone. Therefore, proceed with caution.
//
// Description:
//
// Before you disable a resource directory, you must make sure that the following requirements are met:
//
// 	- All members of the cloud account type in the resource directory are removed. You can call the [RemoveCloudAccount](~~RemoveCloudAccount~~) operation to remove a member of the cloud account type.
//
// 	- All folders except the Root folder are deleted from the resource directory. You can call the [DeleteFolder](~~DeleteFolder~~) operation to delete a folder.
//
// @return DestroyResourceDirectoryResponse
func (client *Client) DestroyResourceDirectory() (_result *DestroyResourceDirectoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DestroyResourceDirectoryResponse{}
	_body, _err := client.DestroyResourceDirectoryWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Detaches an access control policy.
//
// Description:
//
// After you detach an access control policy, the operations performed on resources by using members are not limited by the policy. Make sure that the detached policy meets your expectations. Otherwise, your business may be affected.
//
// Both the system and custom access control policies can be detached. If an object has only one access control policy attached, the policy cannot be detached.
//
// @param request - DetachControlPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DetachControlPolicyResponse
func (client *Client) DetachControlPolicyWithOptions(request *DetachControlPolicyRequest, runtime *util.RuntimeOptions) (_result *DetachControlPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PolicyId)) {
		query["PolicyId"] = request.PolicyId
	}

	if !tea.BoolValue(util.IsUnset(request.TargetId)) {
		query["TargetId"] = request.TargetId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DetachControlPolicy"),
		Version:     tea.String("2022-04-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DetachControlPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Detaches an access control policy.
//
// Description:
//
// After you detach an access control policy, the operations performed on resources by using members are not limited by the policy. Make sure that the detached policy meets your expectations. Otherwise, your business may be affected.
//
// Both the system and custom access control policies can be detached. If an object has only one access control policy attached, the policy cannot be detached.
//
// @param request - DetachControlPolicyRequest
//
// @return DetachControlPolicyResponse
func (client *Client) DetachControlPolicy(request *DetachControlPolicyRequest) (_result *DetachControlPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DetachControlPolicyResponse{}
	_body, _err := client.DetachControlPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Disables the Control Policy feature.
//
// Description:
//
// After you disable the Control Policy feature, the system automatically detaches all access control policies that are attached to folders and members. The system does not delete these access control policies, but you cannot attach them to folders or members again.
//
// > If you disable the Control Policy feature, the permissions of all folders and members in your resource directory are affected. Therefore, proceed with caution.
//
// @param request - DisableControlPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableControlPolicyResponse
func (client *Client) DisableControlPolicyWithOptions(runtime *util.RuntimeOptions) (_result *DisableControlPolicyResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("DisableControlPolicy"),
		Version:     tea.String("2022-04-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DisableControlPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disables the Control Policy feature.
//
// Description:
//
// After you disable the Control Policy feature, the system automatically detaches all access control policies that are attached to folders and members. The system does not delete these access control policies, but you cannot attach them to folders or members again.
//
// > If you disable the Control Policy feature, the permissions of all folders and members in your resource directory are affected. Therefore, proceed with caution.
//
// @return DisableControlPolicyResponse
func (client *Client) DisableControlPolicy() (_result *DisableControlPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DisableControlPolicyResponse{}
	_body, _err := client.DisableControlPolicyWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Unbinds a contact from a resource directory, folder, or member.
//
// @param request - DisassociateMembersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisassociateMembersResponse
func (client *Client) DisassociateMembersWithOptions(request *DisassociateMembersRequest, runtime *util.RuntimeOptions) (_result *DisassociateMembersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ContactId)) {
		query["ContactId"] = request.ContactId
	}

	if !tea.BoolValue(util.IsUnset(request.Members)) {
		query["Members"] = request.Members
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DisassociateMembers"),
		Version:     tea.String("2022-04-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DisassociateMembersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Unbinds a contact from a resource directory, folder, or member.
//
// @param request - DisassociateMembersRequest
//
// @return DisassociateMembersResponse
func (client *Client) DisassociateMembers(request *DisassociateMembersRequest) (_result *DisassociateMembersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DisassociateMembersResponse{}
	_body, _err := client.DisassociateMembersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Enables the Control Policy feature.
//
// Description:
//
// The Control Policy feature provided by the Resource Directory service allows you to manage the permission boundaries of the folders or members in your resource directory in a centralized manner. This feature is implemented based on the resource directory. You can use this feature to develop common or dedicated rules for access control. The Control Policy feature does not grant permissions but only defines permission boundaries. A member in a resource directory can be used to access resources only after it is granted the required permissions by using the Resource Access Management (RAM) service. For more information, see [Overview of the Control Policy feature](https://help.aliyun.com/document_detail/178671.html).
//
// @param request - EnableControlPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableControlPolicyResponse
func (client *Client) EnableControlPolicyWithOptions(runtime *util.RuntimeOptions) (_result *EnableControlPolicyResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("EnableControlPolicy"),
		Version:     tea.String("2022-04-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &EnableControlPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables the Control Policy feature.
//
// Description:
//
// The Control Policy feature provided by the Resource Directory service allows you to manage the permission boundaries of the folders or members in your resource directory in a centralized manner. This feature is implemented based on the resource directory. You can use this feature to develop common or dedicated rules for access control. The Control Policy feature does not grant permissions but only defines permission boundaries. A member in a resource directory can be used to access resources only after it is granted the required permissions by using the Resource Access Management (RAM) service. For more information, see [Overview of the Control Policy feature](https://help.aliyun.com/document_detail/178671.html).
//
// @return EnableControlPolicyResponse
func (client *Client) EnableControlPolicy() (_result *EnableControlPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EnableControlPolicyResponse{}
	_body, _err := client.EnableControlPolicyWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Enables a resource directory.
//
// Description:
//
// You can use the current account or a newly created account to enable a resource directory. For more information, see [Enable a resource directory](https://help.aliyun.com/document_detail/111215.html).
//
// @param request - EnableResourceDirectoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableResourceDirectoryResponse
func (client *Client) EnableResourceDirectoryWithOptions(request *EnableResourceDirectoryRequest, runtime *util.RuntimeOptions) (_result *EnableResourceDirectoryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.EnableMode)) {
		query["EnableMode"] = request.EnableMode
	}

	if !tea.BoolValue(util.IsUnset(request.MAName)) {
		query["MAName"] = request.MAName
	}

	if !tea.BoolValue(util.IsUnset(request.MASecureMobilePhone)) {
		query["MASecureMobilePhone"] = request.MASecureMobilePhone
	}

	if !tea.BoolValue(util.IsUnset(request.VerificationCode)) {
		query["VerificationCode"] = request.VerificationCode
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("EnableResourceDirectory"),
		Version:     tea.String("2022-04-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &EnableResourceDirectoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables a resource directory.
//
// Description:
//
// You can use the current account or a newly created account to enable a resource directory. For more information, see [Enable a resource directory](https://help.aliyun.com/document_detail/111215.html).
//
// @param request - EnableResourceDirectoryRequest
//
// @return EnableResourceDirectoryResponse
func (client *Client) EnableResourceDirectory(request *EnableResourceDirectoryRequest) (_result *EnableResourceDirectoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EnableResourceDirectoryResponse{}
	_body, _err := client.EnableResourceDirectoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information of a member.
//
// @param request - GetAccountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAccountResponse
func (client *Client) GetAccountWithOptions(request *GetAccountRequest, runtime *util.RuntimeOptions) (_result *GetAccountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountId)) {
		query["AccountId"] = request.AccountId
	}

	if !tea.BoolValue(util.IsUnset(request.IncludeTags)) {
		query["IncludeTags"] = request.IncludeTags
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetAccount"),
		Version:     tea.String("2022-04-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAccountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information of a member.
//
// @param request - GetAccountRequest
//
// @return GetAccountResponse
func (client *Client) GetAccount(request *GetAccountRequest) (_result *GetAccountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAccountResponse{}
	_body, _err := client.GetAccountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the result of a member deletion check.
//
// Description:
//
// After you call the [CheckAccountDelete](~~CheckAccountDelete~~) operation to perform a member deletion check, you can call the [GetAccountDeletionCheckResult](~~GetAccountDeletionCheckResult~~) operation to query the check result. If the check result shows that the member meets deletion requirements, you can delete the member. Otherwise, you need to first modify the items that do not meet requirements.
//
// @param request - GetAccountDeletionCheckResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAccountDeletionCheckResultResponse
func (client *Client) GetAccountDeletionCheckResultWithOptions(request *GetAccountDeletionCheckResultRequest, runtime *util.RuntimeOptions) (_result *GetAccountDeletionCheckResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountId)) {
		query["AccountId"] = request.AccountId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetAccountDeletionCheckResult"),
		Version:     tea.String("2022-04-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAccountDeletionCheckResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the result of a member deletion check.
//
// Description:
//
// After you call the [CheckAccountDelete](~~CheckAccountDelete~~) operation to perform a member deletion check, you can call the [GetAccountDeletionCheckResult](~~GetAccountDeletionCheckResult~~) operation to query the check result. If the check result shows that the member meets deletion requirements, you can delete the member. Otherwise, you need to first modify the items that do not meet requirements.
//
// @param request - GetAccountDeletionCheckResultRequest
//
// @return GetAccountDeletionCheckResultResponse
func (client *Client) GetAccountDeletionCheckResult(request *GetAccountDeletionCheckResultRequest) (_result *GetAccountDeletionCheckResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAccountDeletionCheckResultResponse{}
	_body, _err := client.GetAccountDeletionCheckResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the deletion status of a member.
//
// @param request - GetAccountDeletionStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAccountDeletionStatusResponse
func (client *Client) GetAccountDeletionStatusWithOptions(request *GetAccountDeletionStatusRequest, runtime *util.RuntimeOptions) (_result *GetAccountDeletionStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountId)) {
		query["AccountId"] = request.AccountId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetAccountDeletionStatus"),
		Version:     tea.String("2022-04-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAccountDeletionStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the deletion status of a member.
//
// @param request - GetAccountDeletionStatusRequest
//
// @return GetAccountDeletionStatusResponse
func (client *Client) GetAccountDeletionStatus(request *GetAccountDeletionStatusRequest) (_result *GetAccountDeletionStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAccountDeletionStatusResponse{}
	_body, _err := client.GetAccountDeletionStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of an access control policy.
//
// @param request - GetControlPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetControlPolicyResponse
func (client *Client) GetControlPolicyWithOptions(request *GetControlPolicyRequest, runtime *util.RuntimeOptions) (_result *GetControlPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Language)) {
		query["Language"] = request.Language
	}

	if !tea.BoolValue(util.IsUnset(request.PolicyId)) {
		query["PolicyId"] = request.PolicyId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetControlPolicy"),
		Version:     tea.String("2022-04-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetControlPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of an access control policy.
//
// @param request - GetControlPolicyRequest
//
// @return GetControlPolicyResponse
func (client *Client) GetControlPolicy(request *GetControlPolicyRequest) (_result *GetControlPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetControlPolicyResponse{}
	_body, _err := client.GetControlPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the status of the Control Policy feature.
//
// @param request - GetControlPolicyEnablementStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetControlPolicyEnablementStatusResponse
func (client *Client) GetControlPolicyEnablementStatusWithOptions(runtime *util.RuntimeOptions) (_result *GetControlPolicyEnablementStatusResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("GetControlPolicyEnablementStatus"),
		Version:     tea.String("2022-04-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetControlPolicyEnablementStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the status of the Control Policy feature.
//
// @return GetControlPolicyEnablementStatusResponse
func (client *Client) GetControlPolicyEnablementStatus() (_result *GetControlPolicyEnablementStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetControlPolicyEnablementStatusResponse{}
	_body, _err := client.GetControlPolicyEnablementStatusWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about a folder.
//
// @param request - GetFolderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetFolderResponse
func (client *Client) GetFolderWithOptions(request *GetFolderRequest, runtime *util.RuntimeOptions) (_result *GetFolderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FolderId)) {
		query["FolderId"] = request.FolderId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetFolder"),
		Version:     tea.String("2022-04-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetFolderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a folder.
//
// @param request - GetFolderRequest
//
// @return GetFolderResponse
func (client *Client) GetFolder(request *GetFolderRequest) (_result *GetFolderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetFolderResponse{}
	_body, _err := client.GetFolderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information of an invitation.
//
// @param request - GetHandshakeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetHandshakeResponse
func (client *Client) GetHandshakeWithOptions(request *GetHandshakeRequest, runtime *util.RuntimeOptions) (_result *GetHandshakeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.HandshakeId)) {
		query["HandshakeId"] = request.HandshakeId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetHandshake"),
		Version:     tea.String("2022-04-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetHandshakeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information of an invitation.
//
// @param request - GetHandshakeRequest
//
// @return GetHandshakeResponse
func (client *Client) GetHandshake(request *GetHandshakeRequest) (_result *GetHandshakeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetHandshakeResponse{}
	_body, _err := client.GetHandshakeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about a contact.
//
// @param request - GetMessageContactRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMessageContactResponse
func (client *Client) GetMessageContactWithOptions(request *GetMessageContactRequest, runtime *util.RuntimeOptions) (_result *GetMessageContactResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ContactId)) {
		query["ContactId"] = request.ContactId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetMessageContact"),
		Version:     tea.String("2022-04-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetMessageContactResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a contact.
//
// @param request - GetMessageContactRequest
//
// @return GetMessageContactResponse
func (client *Client) GetMessageContact(request *GetMessageContactRequest) (_result *GetMessageContactResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetMessageContactResponse{}
	_body, _err := client.GetMessageContactWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the deletion status of a contact.
//
// @param request - GetMessageContactDeletionStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMessageContactDeletionStatusResponse
func (client *Client) GetMessageContactDeletionStatusWithOptions(request *GetMessageContactDeletionStatusRequest, runtime *util.RuntimeOptions) (_result *GetMessageContactDeletionStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ContactId)) {
		query["ContactId"] = request.ContactId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetMessageContactDeletionStatus"),
		Version:     tea.String("2022-04-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetMessageContactDeletionStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the deletion status of a contact.
//
// @param request - GetMessageContactDeletionStatusRequest
//
// @return GetMessageContactDeletionStatusResponse
func (client *Client) GetMessageContactDeletionStatus(request *GetMessageContactDeletionStatusRequest) (_result *GetMessageContactDeletionStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetMessageContactDeletionStatusResponse{}
	_body, _err := client.GetMessageContactDeletionStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information of a billing account.
//
// @param request - GetPayerForAccountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPayerForAccountResponse
func (client *Client) GetPayerForAccountWithOptions(request *GetPayerForAccountRequest, runtime *util.RuntimeOptions) (_result *GetPayerForAccountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountId)) {
		query["AccountId"] = request.AccountId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetPayerForAccount"),
		Version:     tea.String("2022-04-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetPayerForAccountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information of a billing account.
//
// @param request - GetPayerForAccountRequest
//
// @return GetPayerForAccountResponse
func (client *Client) GetPayerForAccount(request *GetPayerForAccountRequest) (_result *GetPayerForAccountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetPayerForAccountResponse{}
	_body, _err := client.GetPayerForAccountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// \\*\\*\\	- Co., Ltd.
//
// @param request - GetResourceDirectoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetResourceDirectoryResponse
func (client *Client) GetResourceDirectoryWithOptions(runtime *util.RuntimeOptions) (_result *GetResourceDirectoryResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("GetResourceDirectory"),
		Version:     tea.String("2022-04-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetResourceDirectoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// \\*\\*\\	- Co., Ltd.
//
// @return GetResourceDirectoryResponse
func (client *Client) GetResourceDirectory() (_result *GetResourceDirectoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetResourceDirectoryResponse{}
	_body, _err := client.GetResourceDirectoryWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Invites an account to join a resource directory.
//
// @param request - InviteAccountToResourceDirectoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InviteAccountToResourceDirectoryResponse
func (client *Client) InviteAccountToResourceDirectoryWithOptions(request *InviteAccountToResourceDirectoryRequest, runtime *util.RuntimeOptions) (_result *InviteAccountToResourceDirectoryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Note)) {
		query["Note"] = request.Note
	}

	if !tea.BoolValue(util.IsUnset(request.ParentFolderId)) {
		query["ParentFolderId"] = request.ParentFolderId
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	if !tea.BoolValue(util.IsUnset(request.TargetEntity)) {
		query["TargetEntity"] = request.TargetEntity
	}

	if !tea.BoolValue(util.IsUnset(request.TargetType)) {
		query["TargetType"] = request.TargetType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("InviteAccountToResourceDirectory"),
		Version:     tea.String("2022-04-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &InviteAccountToResourceDirectoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Invites an account to join a resource directory.
//
// @param request - InviteAccountToResourceDirectoryRequest
//
// @return InviteAccountToResourceDirectoryResponse
func (client *Client) InviteAccountToResourceDirectory(request *InviteAccountToResourceDirectoryRequest) (_result *InviteAccountToResourceDirectoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &InviteAccountToResourceDirectoryResponse{}
	_body, _err := client.InviteAccountToResourceDirectoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries all the members in a resource directory.
//
// Description:
//
// You can use only the management account of a resource directory or a delegated administrator account of a trusted service to call this operation.
//
// @param request - ListAccountsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAccountsResponse
func (client *Client) ListAccountsWithOptions(request *ListAccountsRequest, runtime *util.RuntimeOptions) (_result *ListAccountsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IncludeTags)) {
		query["IncludeTags"] = request.IncludeTags
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.QueryKeyword)) {
		query["QueryKeyword"] = request.QueryKeyword
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListAccounts"),
		Version:     tea.String("2022-04-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListAccountsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries all the members in a resource directory.
//
// Description:
//
// You can use only the management account of a resource directory or a delegated administrator account of a trusted service to call this operation.
//
// @param request - ListAccountsRequest
//
// @return ListAccountsResponse
func (client *Client) ListAccounts(request *ListAccountsRequest) (_result *ListAccountsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListAccountsResponse{}
	_body, _err := client.ListAccountsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information of members in a folder.
//
// @param request - ListAccountsForParentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAccountsForParentResponse
func (client *Client) ListAccountsForParentWithOptions(request *ListAccountsForParentRequest, runtime *util.RuntimeOptions) (_result *ListAccountsForParentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IncludeTags)) {
		query["IncludeTags"] = request.IncludeTags
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ParentFolderId)) {
		query["ParentFolderId"] = request.ParentFolderId
	}

	if !tea.BoolValue(util.IsUnset(request.QueryKeyword)) {
		query["QueryKeyword"] = request.QueryKeyword
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListAccountsForParent"),
		Version:     tea.String("2022-04-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListAccountsForParentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information of members in a folder.
//
// @param request - ListAccountsForParentRequest
//
// @return ListAccountsForParentResponse
func (client *Client) ListAccountsForParent(request *ListAccountsForParentRequest) (_result *ListAccountsForParentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListAccountsForParentResponse{}
	_body, _err := client.ListAccountsForParentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information of all the parent folders of a specified folder.
//
// @param request - ListAncestorsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAncestorsResponse
func (client *Client) ListAncestorsWithOptions(request *ListAncestorsRequest, runtime *util.RuntimeOptions) (_result *ListAncestorsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ChildId)) {
		query["ChildId"] = request.ChildId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListAncestors"),
		Version:     tea.String("2022-04-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListAncestorsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information of all the parent folders of a specified folder.
//
// @param request - ListAncestorsRequest
//
// @return ListAncestorsResponse
func (client *Client) ListAncestors(request *ListAncestorsRequest) (_result *ListAncestorsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListAncestorsResponse{}
	_body, _err := client.ListAncestorsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries access control policies.
//
// @param request - ListControlPoliciesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListControlPoliciesResponse
func (client *Client) ListControlPoliciesWithOptions(request *ListControlPoliciesRequest, runtime *util.RuntimeOptions) (_result *ListControlPoliciesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Language)) {
		query["Language"] = request.Language
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.PolicyType)) {
		query["PolicyType"] = request.PolicyType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListControlPolicies"),
		Version:     tea.String("2022-04-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListControlPoliciesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries access control policies.
//
// @param request - ListControlPoliciesRequest
//
// @return ListControlPoliciesResponse
func (client *Client) ListControlPolicies(request *ListControlPoliciesRequest) (_result *ListControlPoliciesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListControlPoliciesResponse{}
	_body, _err := client.ListControlPoliciesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the access control policies that are attached to a folder or member.
//
// @param request - ListControlPolicyAttachmentsForTargetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListControlPolicyAttachmentsForTargetResponse
func (client *Client) ListControlPolicyAttachmentsForTargetWithOptions(request *ListControlPolicyAttachmentsForTargetRequest, runtime *util.RuntimeOptions) (_result *ListControlPolicyAttachmentsForTargetResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Language)) {
		query["Language"] = request.Language
	}

	if !tea.BoolValue(util.IsUnset(request.TargetId)) {
		query["TargetId"] = request.TargetId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListControlPolicyAttachmentsForTarget"),
		Version:     tea.String("2022-04-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListControlPolicyAttachmentsForTargetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the access control policies that are attached to a folder or member.
//
// @param request - ListControlPolicyAttachmentsForTargetRequest
//
// @return ListControlPolicyAttachmentsForTargetResponse
func (client *Client) ListControlPolicyAttachmentsForTarget(request *ListControlPolicyAttachmentsForTargetRequest) (_result *ListControlPolicyAttachmentsForTargetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListControlPolicyAttachmentsForTargetResponse{}
	_body, _err := client.ListControlPolicyAttachmentsForTargetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries delegated administrator accounts.
//
// @param request - ListDelegatedAdministratorsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDelegatedAdministratorsResponse
func (client *Client) ListDelegatedAdministratorsWithOptions(request *ListDelegatedAdministratorsRequest, runtime *util.RuntimeOptions) (_result *ListDelegatedAdministratorsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ServicePrincipal)) {
		query["ServicePrincipal"] = request.ServicePrincipal
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListDelegatedAdministrators"),
		Version:     tea.String("2022-04-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListDelegatedAdministratorsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries delegated administrator accounts.
//
// @param request - ListDelegatedAdministratorsRequest
//
// @return ListDelegatedAdministratorsResponse
func (client *Client) ListDelegatedAdministrators(request *ListDelegatedAdministratorsRequest) (_result *ListDelegatedAdministratorsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListDelegatedAdministratorsResponse{}
	_body, _err := client.ListDelegatedAdministratorsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the trusted services for which a member is specified as a delegated administrator account.
//
// @param request - ListDelegatedServicesForAccountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDelegatedServicesForAccountResponse
func (client *Client) ListDelegatedServicesForAccountWithOptions(request *ListDelegatedServicesForAccountRequest, runtime *util.RuntimeOptions) (_result *ListDelegatedServicesForAccountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountId)) {
		query["AccountId"] = request.AccountId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListDelegatedServicesForAccount"),
		Version:     tea.String("2022-04-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListDelegatedServicesForAccountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the trusted services for which a member is specified as a delegated administrator account.
//
// @param request - ListDelegatedServicesForAccountRequest
//
// @return ListDelegatedServicesForAccountResponse
func (client *Client) ListDelegatedServicesForAccount(request *ListDelegatedServicesForAccountRequest) (_result *ListDelegatedServicesForAccountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListDelegatedServicesForAccountResponse{}
	_body, _err := client.ListDelegatedServicesForAccountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information of all subfolders of a folder.
//
// Description:
//
// You can call this API operation to view the information of only the first-level subfolders of a folder.
//
// @param request - ListFoldersForParentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListFoldersForParentResponse
func (client *Client) ListFoldersForParentWithOptions(request *ListFoldersForParentRequest, runtime *util.RuntimeOptions) (_result *ListFoldersForParentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ParentFolderId)) {
		query["ParentFolderId"] = request.ParentFolderId
	}

	if !tea.BoolValue(util.IsUnset(request.QueryKeyword)) {
		query["QueryKeyword"] = request.QueryKeyword
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListFoldersForParent"),
		Version:     tea.String("2022-04-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListFoldersForParentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information of all subfolders of a folder.
//
// Description:
//
// You can call this API operation to view the information of only the first-level subfolders of a folder.
//
// @param request - ListFoldersForParentRequest
//
// @return ListFoldersForParentResponse
func (client *Client) ListFoldersForParent(request *ListFoldersForParentRequest) (_result *ListFoldersForParentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListFoldersForParentResponse{}
	_body, _err := client.ListFoldersForParentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the invitations that are associated with an account.
//
// @param request - ListHandshakesForAccountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListHandshakesForAccountResponse
func (client *Client) ListHandshakesForAccountWithOptions(request *ListHandshakesForAccountRequest, runtime *util.RuntimeOptions) (_result *ListHandshakesForAccountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListHandshakesForAccount"),
		Version:     tea.String("2022-04-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListHandshakesForAccountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the invitations that are associated with an account.
//
// @param request - ListHandshakesForAccountRequest
//
// @return ListHandshakesForAccountResponse
func (client *Client) ListHandshakesForAccount(request *ListHandshakesForAccountRequest) (_result *ListHandshakesForAccountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListHandshakesForAccountResponse{}
	_body, _err := client.ListHandshakesForAccountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries invitations in a resource directory.
//
// @param request - ListHandshakesForResourceDirectoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListHandshakesForResourceDirectoryResponse
func (client *Client) ListHandshakesForResourceDirectoryWithOptions(request *ListHandshakesForResourceDirectoryRequest, runtime *util.RuntimeOptions) (_result *ListHandshakesForResourceDirectoryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListHandshakesForResourceDirectory"),
		Version:     tea.String("2022-04-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListHandshakesForResourceDirectoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries invitations in a resource directory.
//
// @param request - ListHandshakesForResourceDirectoryRequest
//
// @return ListHandshakesForResourceDirectoryResponse
func (client *Client) ListHandshakesForResourceDirectory(request *ListHandshakesForResourceDirectoryRequest) (_result *ListHandshakesForResourceDirectoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListHandshakesForResourceDirectoryResponse{}
	_body, _err := client.ListHandshakesForResourceDirectoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the mobile phone number or email address to be verified for a contact.
//
// @param request - ListMessageContactVerificationsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMessageContactVerificationsResponse
func (client *Client) ListMessageContactVerificationsWithOptions(request *ListMessageContactVerificationsRequest, runtime *util.RuntimeOptions) (_result *ListMessageContactVerificationsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ContactId)) {
		query["ContactId"] = request.ContactId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListMessageContactVerifications"),
		Version:     tea.String("2022-04-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListMessageContactVerificationsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the mobile phone number or email address to be verified for a contact.
//
// @param request - ListMessageContactVerificationsRequest
//
// @return ListMessageContactVerificationsResponse
func (client *Client) ListMessageContactVerifications(request *ListMessageContactVerificationsRequest) (_result *ListMessageContactVerificationsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListMessageContactVerificationsResponse{}
	_body, _err := client.ListMessageContactVerificationsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries contacts.
//
// @param request - ListMessageContactsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMessageContactsResponse
func (client *Client) ListMessageContactsWithOptions(request *ListMessageContactsRequest, runtime *util.RuntimeOptions) (_result *ListMessageContactsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ContactId)) {
		query["ContactId"] = request.ContactId
	}

	if !tea.BoolValue(util.IsUnset(request.Member)) {
		query["Member"] = request.Member
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListMessageContacts"),
		Version:     tea.String("2022-04-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListMessageContactsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries contacts.
//
// @param request - ListMessageContactsRequest
//
// @return ListMessageContactsResponse
func (client *Client) ListMessageContacts(request *ListMessageContactsRequest) (_result *ListMessageContactsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListMessageContactsResponse{}
	_body, _err := client.ListMessageContactsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries tag keys.
//
// @param request - ListTagKeysRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTagKeysResponse
func (client *Client) ListTagKeysWithOptions(request *ListTagKeysRequest, runtime *util.RuntimeOptions) (_result *ListTagKeysResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.KeyFilter)) {
		query["KeyFilter"] = request.KeyFilter
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTagKeys"),
		Version:     tea.String("2022-04-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTagKeysResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries tag keys.
//
// @param request - ListTagKeysRequest
//
// @return ListTagKeysResponse
func (client *Client) ListTagKeys(request *ListTagKeysRequest) (_result *ListTagKeysResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTagKeysResponse{}
	_body, _err := client.ListTagKeysWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the tags that are added to the members in a resource directory.
//
// @param request - ListTagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTagResourcesResponse
func (client *Client) ListTagResourcesWithOptions(request *ListTagResourcesRequest, runtime *util.RuntimeOptions) (_result *ListTagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTagResources"),
		Version:     tea.String("2022-04-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTagResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the tags that are added to the members in a resource directory.
//
// @param request - ListTagResourcesRequest
//
// @return ListTagResourcesResponse
func (client *Client) ListTagResources(request *ListTagResourcesRequest) (_result *ListTagResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTagResourcesResponse{}
	_body, _err := client.ListTagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the tag values of a tag key.
//
// @param request - ListTagValuesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTagValuesResponse
func (client *Client) ListTagValuesWithOptions(request *ListTagValuesRequest, runtime *util.RuntimeOptions) (_result *ListTagValuesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.TagKey)) {
		query["TagKey"] = request.TagKey
	}

	if !tea.BoolValue(util.IsUnset(request.ValueFilter)) {
		query["ValueFilter"] = request.ValueFilter
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTagValues"),
		Version:     tea.String("2022-04-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTagValuesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the tag values of a tag key.
//
// @param request - ListTagValuesRequest
//
// @return ListTagValuesResponse
func (client *Client) ListTagValues(request *ListTagValuesRequest) (_result *ListTagValuesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTagValuesResponse{}
	_body, _err := client.ListTagValuesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the objects to which an access control policy is attached.
//
// @param request - ListTargetAttachmentsForControlPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTargetAttachmentsForControlPolicyResponse
func (client *Client) ListTargetAttachmentsForControlPolicyWithOptions(request *ListTargetAttachmentsForControlPolicyRequest, runtime *util.RuntimeOptions) (_result *ListTargetAttachmentsForControlPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.PolicyId)) {
		query["PolicyId"] = request.PolicyId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTargetAttachmentsForControlPolicy"),
		Version:     tea.String("2022-04-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTargetAttachmentsForControlPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the objects to which an access control policy is attached.
//
// @param request - ListTargetAttachmentsForControlPolicyRequest
//
// @return ListTargetAttachmentsForControlPolicyResponse
func (client *Client) ListTargetAttachmentsForControlPolicy(request *ListTargetAttachmentsForControlPolicyRequest) (_result *ListTargetAttachmentsForControlPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTargetAttachmentsForControlPolicyResponse{}
	_body, _err := client.ListTargetAttachmentsForControlPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the trusted services that are enabled within a management account or delegated administrator account.
//
// Description:
//
// Only a management account or delegated administrator account can be used to call this operation.
//
// @param request - ListTrustedServiceStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTrustedServiceStatusResponse
func (client *Client) ListTrustedServiceStatusWithOptions(request *ListTrustedServiceStatusRequest, runtime *util.RuntimeOptions) (_result *ListTrustedServiceStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AdminAccountId)) {
		query["AdminAccountId"] = request.AdminAccountId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTrustedServiceStatus"),
		Version:     tea.String("2022-04-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTrustedServiceStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the trusted services that are enabled within a management account or delegated administrator account.
//
// Description:
//
// Only a management account or delegated administrator account can be used to call this operation.
//
// @param request - ListTrustedServiceStatusRequest
//
// @return ListTrustedServiceStatusResponse
func (client *Client) ListTrustedServiceStatus(request *ListTrustedServiceStatusRequest) (_result *ListTrustedServiceStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTrustedServiceStatusResponse{}
	_body, _err := client.ListTrustedServiceStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Moves a member from a folder to another.
//
// @param request - MoveAccountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return MoveAccountResponse
func (client *Client) MoveAccountWithOptions(request *MoveAccountRequest, runtime *util.RuntimeOptions) (_result *MoveAccountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountId)) {
		query["AccountId"] = request.AccountId
	}

	if !tea.BoolValue(util.IsUnset(request.DestinationFolderId)) {
		query["DestinationFolderId"] = request.DestinationFolderId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("MoveAccount"),
		Version:     tea.String("2022-04-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &MoveAccountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Moves a member from a folder to another.
//
// @param request - MoveAccountRequest
//
// @return MoveAccountResponse
func (client *Client) MoveAccount(request *MoveAccountRequest) (_result *MoveAccountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &MoveAccountResponse{}
	_body, _err := client.MoveAccountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Checks whether a management account or member can be used as a main financial account.
//
// @param request - PrecheckForConsolidatedBillingAccountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PrecheckForConsolidatedBillingAccountResponse
func (client *Client) PrecheckForConsolidatedBillingAccountWithOptions(request *PrecheckForConsolidatedBillingAccountRequest, runtime *util.RuntimeOptions) (_result *PrecheckForConsolidatedBillingAccountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BillingAccountId)) {
		query["BillingAccountId"] = request.BillingAccountId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("PrecheckForConsolidatedBillingAccount"),
		Version:     tea.String("2022-04-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PrecheckForConsolidatedBillingAccountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Checks whether a management account or member can be used as a main financial account.
//
// @param request - PrecheckForConsolidatedBillingAccountRequest
//
// @return PrecheckForConsolidatedBillingAccountResponse
func (client *Client) PrecheckForConsolidatedBillingAccount(request *PrecheckForConsolidatedBillingAccountRequest) (_result *PrecheckForConsolidatedBillingAccountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PrecheckForConsolidatedBillingAccountResponse{}
	_body, _err := client.PrecheckForConsolidatedBillingAccountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Specifies a member in a resource directory as a delegated administrator account of a trusted service.
//
// Description:
//
// The delegated administrator account can be used to access the information of the resource directory and view the structure and members of the resource directory. The delegated administrator account can also be used to perform service-related management operations in the trusted service on behalf of the management account of the resource directory. When you call this operation, you must take note of the following limits:
//
// 	- Only some trusted services support delegated administrator accounts. For more information, see [Supported trusted services](https://help.aliyun.com/document_detail/208133.html).
//
// 	- Only the management account of a resource directory or an authorized RAM user or RAM role of the management account can be used to call this operation.
//
// 	- The number of delegated administrator accounts that are allowed for a trusted service is defined by the trusted service.
//
// @param request - RegisterDelegatedAdministratorRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RegisterDelegatedAdministratorResponse
func (client *Client) RegisterDelegatedAdministratorWithOptions(request *RegisterDelegatedAdministratorRequest, runtime *util.RuntimeOptions) (_result *RegisterDelegatedAdministratorResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountId)) {
		query["AccountId"] = request.AccountId
	}

	if !tea.BoolValue(util.IsUnset(request.ServicePrincipal)) {
		query["ServicePrincipal"] = request.ServicePrincipal
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RegisterDelegatedAdministrator"),
		Version:     tea.String("2022-04-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RegisterDelegatedAdministratorResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Specifies a member in a resource directory as a delegated administrator account of a trusted service.
//
// Description:
//
// The delegated administrator account can be used to access the information of the resource directory and view the structure and members of the resource directory. The delegated administrator account can also be used to perform service-related management operations in the trusted service on behalf of the management account of the resource directory. When you call this operation, you must take note of the following limits:
//
// 	- Only some trusted services support delegated administrator accounts. For more information, see [Supported trusted services](https://help.aliyun.com/document_detail/208133.html).
//
// 	- Only the management account of a resource directory or an authorized RAM user or RAM role of the management account can be used to call this operation.
//
// 	- The number of delegated administrator accounts that are allowed for a trusted service is defined by the trusted service.
//
// @param request - RegisterDelegatedAdministratorRequest
//
// @return RegisterDelegatedAdministratorResponse
func (client *Client) RegisterDelegatedAdministrator(request *RegisterDelegatedAdministratorRequest) (_result *RegisterDelegatedAdministratorResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RegisterDelegatedAdministratorResponse{}
	_body, _err := client.RegisterDelegatedAdministratorWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Removes a member of the cloud account type.
//
// @param request - RemoveCloudAccountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveCloudAccountResponse
func (client *Client) RemoveCloudAccountWithOptions(request *RemoveCloudAccountRequest, runtime *util.RuntimeOptions) (_result *RemoveCloudAccountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountId)) {
		query["AccountId"] = request.AccountId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RemoveCloudAccount"),
		Version:     tea.String("2022-04-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RemoveCloudAccountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes a member of the cloud account type.
//
// @param request - RemoveCloudAccountRequest
//
// @return RemoveCloudAccountResponse
func (client *Client) RemoveCloudAccount(request *RemoveCloudAccountRequest) (_result *RemoveCloudAccountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RemoveCloudAccountResponse{}
	_body, _err := client.RemoveCloudAccountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Resends a verification email for the email address change of a member.
//
// @param request - RetryChangeAccountEmailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RetryChangeAccountEmailResponse
func (client *Client) RetryChangeAccountEmailWithOptions(request *RetryChangeAccountEmailRequest, runtime *util.RuntimeOptions) (_result *RetryChangeAccountEmailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountId)) {
		query["AccountId"] = request.AccountId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RetryChangeAccountEmail"),
		Version:     tea.String("2022-04-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RetryChangeAccountEmailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Resends a verification email for the email address change of a member.
//
// @param request - RetryChangeAccountEmailRequest
//
// @return RetryChangeAccountEmailResponse
func (client *Client) RetryChangeAccountEmail(request *RetryChangeAccountEmailRequest) (_result *RetryChangeAccountEmailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RetryChangeAccountEmailResponse{}
	_body, _err := client.RetryChangeAccountEmailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Sends verification information to the email address of a contact.
//
// @param request - SendEmailVerificationForMessageContactRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SendEmailVerificationForMessageContactResponse
func (client *Client) SendEmailVerificationForMessageContactWithOptions(request *SendEmailVerificationForMessageContactRequest, runtime *util.RuntimeOptions) (_result *SendEmailVerificationForMessageContactResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ContactId)) {
		query["ContactId"] = request.ContactId
	}

	if !tea.BoolValue(util.IsUnset(request.EmailAddress)) {
		query["EmailAddress"] = request.EmailAddress
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SendEmailVerificationForMessageContact"),
		Version:     tea.String("2022-04-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SendEmailVerificationForMessageContactResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Sends verification information to the email address of a contact.
//
// @param request - SendEmailVerificationForMessageContactRequest
//
// @return SendEmailVerificationForMessageContactResponse
func (client *Client) SendEmailVerificationForMessageContact(request *SendEmailVerificationForMessageContactRequest) (_result *SendEmailVerificationForMessageContactResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SendEmailVerificationForMessageContactResponse{}
	_body, _err := client.SendEmailVerificationForMessageContactWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Sends verification information to the mobile phone number of a contact.
//
// @param request - SendPhoneVerificationForMessageContactRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SendPhoneVerificationForMessageContactResponse
func (client *Client) SendPhoneVerificationForMessageContactWithOptions(request *SendPhoneVerificationForMessageContactRequest, runtime *util.RuntimeOptions) (_result *SendPhoneVerificationForMessageContactResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ContactId)) {
		query["ContactId"] = request.ContactId
	}

	if !tea.BoolValue(util.IsUnset(request.PhoneNumber)) {
		query["PhoneNumber"] = request.PhoneNumber
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SendPhoneVerificationForMessageContact"),
		Version:     tea.String("2022-04-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SendPhoneVerificationForMessageContactResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Sends verification information to the mobile phone number of a contact.
//
// @param request - SendPhoneVerificationForMessageContactRequest
//
// @return SendPhoneVerificationForMessageContactResponse
func (client *Client) SendPhoneVerificationForMessageContact(request *SendPhoneVerificationForMessageContactRequest) (_result *SendPhoneVerificationForMessageContactResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SendPhoneVerificationForMessageContactResponse{}
	_body, _err := client.SendPhoneVerificationForMessageContactWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Sends a verification code to the mobile phone number that you want to bind to a member of the resource account type in a resource directory for security purposes.
//
// Description:
//
// To ensure that the system can record the operators of management operations, you must use a RAM user or RAM role to which the AliyunResourceDirectoryFullAccess policy is attached within the management account of your resource directory to call this API operation.
//
// @param request - SendVerificationCodeForBindSecureMobilePhoneRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SendVerificationCodeForBindSecureMobilePhoneResponse
func (client *Client) SendVerificationCodeForBindSecureMobilePhoneWithOptions(request *SendVerificationCodeForBindSecureMobilePhoneRequest, runtime *util.RuntimeOptions) (_result *SendVerificationCodeForBindSecureMobilePhoneResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountId)) {
		query["AccountId"] = request.AccountId
	}

	if !tea.BoolValue(util.IsUnset(request.SecureMobilePhone)) {
		query["SecureMobilePhone"] = request.SecureMobilePhone
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SendVerificationCodeForBindSecureMobilePhone"),
		Version:     tea.String("2022-04-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SendVerificationCodeForBindSecureMobilePhoneResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Sends a verification code to the mobile phone number that you want to bind to a member of the resource account type in a resource directory for security purposes.
//
// Description:
//
// To ensure that the system can record the operators of management operations, you must use a RAM user or RAM role to which the AliyunResourceDirectoryFullAccess policy is attached within the management account of your resource directory to call this API operation.
//
// @param request - SendVerificationCodeForBindSecureMobilePhoneRequest
//
// @return SendVerificationCodeForBindSecureMobilePhoneResponse
func (client *Client) SendVerificationCodeForBindSecureMobilePhone(request *SendVerificationCodeForBindSecureMobilePhoneRequest) (_result *SendVerificationCodeForBindSecureMobilePhoneResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SendVerificationCodeForBindSecureMobilePhoneResponse{}
	_body, _err := client.SendVerificationCodeForBindSecureMobilePhoneWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Sends a verification code to the mobile phone number bound to a newly created account when you use the account to enable a resource directory.
//
// Description:
//
// Each Alibaba Cloud account can be used to send a maximum of 100 verification codes per day.
//
// @param request - SendVerificationCodeForEnableRDRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SendVerificationCodeForEnableRDResponse
func (client *Client) SendVerificationCodeForEnableRDWithOptions(request *SendVerificationCodeForEnableRDRequest, runtime *util.RuntimeOptions) (_result *SendVerificationCodeForEnableRDResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SecureMobilePhone)) {
		query["SecureMobilePhone"] = request.SecureMobilePhone
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SendVerificationCodeForEnableRD"),
		Version:     tea.String("2022-04-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SendVerificationCodeForEnableRDResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Sends a verification code to the mobile phone number bound to a newly created account when you use the account to enable a resource directory.
//
// Description:
//
// Each Alibaba Cloud account can be used to send a maximum of 100 verification codes per day.
//
// @param request - SendVerificationCodeForEnableRDRequest
//
// @return SendVerificationCodeForEnableRDResponse
func (client *Client) SendVerificationCodeForEnableRD(request *SendVerificationCodeForEnableRDRequest) (_result *SendVerificationCodeForEnableRDResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SendVerificationCodeForEnableRDResponse{}
	_body, _err := client.SendVerificationCodeForEnableRDWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Enables or disables the member deletion feature.
//
// Description:
//
// Members of the resource account type can be deleted only after the member deletion feature is enabled.
//
// @param request - SetMemberDeletionPermissionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetMemberDeletionPermissionResponse
func (client *Client) SetMemberDeletionPermissionWithOptions(request *SetMemberDeletionPermissionRequest, runtime *util.RuntimeOptions) (_result *SetMemberDeletionPermissionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SetMemberDeletionPermission"),
		Version:     tea.String("2022-04-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SetMemberDeletionPermissionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables or disables the member deletion feature.
//
// Description:
//
// Members of the resource account type can be deleted only after the member deletion feature is enabled.
//
// @param request - SetMemberDeletionPermissionRequest
//
// @return SetMemberDeletionPermissionResponse
func (client *Client) SetMemberDeletionPermission(request *SetMemberDeletionPermissionRequest) (_result *SetMemberDeletionPermissionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetMemberDeletionPermissionResponse{}
	_body, _err := client.SetMemberDeletionPermissionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds tags to the members in a resource directory.
//
// @param request - TagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TagResourcesResponse
func (client *Client) TagResourcesWithOptions(request *TagResourcesRequest, runtime *util.RuntimeOptions) (_result *TagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("TagResources"),
		Version:     tea.String("2022-04-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &TagResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds tags to the members in a resource directory.
//
// @param request - TagResourcesRequest
//
// @return TagResourcesResponse
func (client *Client) TagResources(request *TagResourcesRequest) (_result *TagResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &TagResourcesResponse{}
	_body, _err := client.TagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Removes tags from the members in a resource directory.
//
// @param request - UntagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UntagResourcesResponse
func (client *Client) UntagResourcesWithOptions(request *UntagResourcesRequest, runtime *util.RuntimeOptions) (_result *UntagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.All)) {
		query["All"] = request.All
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.TagKey)) {
		query["TagKey"] = request.TagKey
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UntagResources"),
		Version:     tea.String("2022-04-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UntagResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes tags from the members in a resource directory.
//
// @param request - UntagResourcesRequest
//
// @return UntagResourcesResponse
func (client *Client) UntagResources(request *UntagResourcesRequest) (_result *UntagResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UntagResourcesResponse{}
	_body, _err := client.UntagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Changes the display name of a member, or switches the type of a member.
//
// Description:
//
//   To ensure that the system can record the operators of management operations, you must use a RAM user or RAM role to which the AliyunResourceDirectoryFullAccess policy is attached within the management account of your resource directory to call this operation.
//
// 	- Before you switch the type of a member from resource account to cloud account, make sure that specific conditions are met. For more information about the conditions, see [Switch a resource account to a cloud account](https://help.aliyun.com/document_detail/111233.html).
//
// 	- Before you switch the type of a member from cloud account to resource account, make sure that specific conditions are met. For more information about the conditions, see [Switch a cloud account to a resource account](https://help.aliyun.com/document_detail/209980.html).
//
// @param request - UpdateAccountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAccountResponse
func (client *Client) UpdateAccountWithOptions(request *UpdateAccountRequest, runtime *util.RuntimeOptions) (_result *UpdateAccountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountId)) {
		query["AccountId"] = request.AccountId
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.NewAccountType)) {
		query["NewAccountType"] = request.NewAccountType
	}

	if !tea.BoolValue(util.IsUnset(request.NewDisplayName)) {
		query["NewDisplayName"] = request.NewDisplayName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateAccount"),
		Version:     tea.String("2022-04-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateAccountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Changes the display name of a member, or switches the type of a member.
//
// Description:
//
//   To ensure that the system can record the operators of management operations, you must use a RAM user or RAM role to which the AliyunResourceDirectoryFullAccess policy is attached within the management account of your resource directory to call this operation.
//
// 	- Before you switch the type of a member from resource account to cloud account, make sure that specific conditions are met. For more information about the conditions, see [Switch a resource account to a cloud account](https://help.aliyun.com/document_detail/111233.html).
//
// 	- Before you switch the type of a member from cloud account to resource account, make sure that specific conditions are met. For more information about the conditions, see [Switch a cloud account to a resource account](https://help.aliyun.com/document_detail/209980.html).
//
// @param request - UpdateAccountRequest
//
// @return UpdateAccountResponse
func (client *Client) UpdateAccount(request *UpdateAccountRequest) (_result *UpdateAccountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateAccountResponse{}
	_body, _err := client.UpdateAccountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates a custom access control policy.
//
// @param request - UpdateControlPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateControlPolicyResponse
func (client *Client) UpdateControlPolicyWithOptions(request *UpdateControlPolicyRequest, runtime *util.RuntimeOptions) (_result *UpdateControlPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NewDescription)) {
		query["NewDescription"] = request.NewDescription
	}

	if !tea.BoolValue(util.IsUnset(request.NewPolicyDocument)) {
		query["NewPolicyDocument"] = request.NewPolicyDocument
	}

	if !tea.BoolValue(util.IsUnset(request.NewPolicyName)) {
		query["NewPolicyName"] = request.NewPolicyName
	}

	if !tea.BoolValue(util.IsUnset(request.PolicyId)) {
		query["PolicyId"] = request.PolicyId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateControlPolicy"),
		Version:     tea.String("2022-04-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateControlPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a custom access control policy.
//
// @param request - UpdateControlPolicyRequest
//
// @return UpdateControlPolicyResponse
func (client *Client) UpdateControlPolicy(request *UpdateControlPolicyRequest) (_result *UpdateControlPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateControlPolicyResponse{}
	_body, _err := client.UpdateControlPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Changes the name of a folder.
//
// @param request - UpdateFolderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateFolderResponse
func (client *Client) UpdateFolderWithOptions(request *UpdateFolderRequest, runtime *util.RuntimeOptions) (_result *UpdateFolderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FolderId)) {
		query["FolderId"] = request.FolderId
	}

	if !tea.BoolValue(util.IsUnset(request.NewFolderName)) {
		query["NewFolderName"] = request.NewFolderName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateFolder"),
		Version:     tea.String("2022-04-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateFolderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Changes the name of a folder.
//
// @param request - UpdateFolderRequest
//
// @return UpdateFolderResponse
func (client *Client) UpdateFolder(request *UpdateFolderRequest) (_result *UpdateFolderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateFolderResponse{}
	_body, _err := client.UpdateFolderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates a contact.
//
// @param request - UpdateMessageContactRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateMessageContactResponse
func (client *Client) UpdateMessageContactWithOptions(request *UpdateMessageContactRequest, runtime *util.RuntimeOptions) (_result *UpdateMessageContactResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ContactId)) {
		query["ContactId"] = request.ContactId
	}

	if !tea.BoolValue(util.IsUnset(request.EmailAddress)) {
		query["EmailAddress"] = request.EmailAddress
	}

	if !tea.BoolValue(util.IsUnset(request.MessageTypes)) {
		query["MessageTypes"] = request.MessageTypes
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.PhoneNumber)) {
		query["PhoneNumber"] = request.PhoneNumber
	}

	if !tea.BoolValue(util.IsUnset(request.Title)) {
		query["Title"] = request.Title
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateMessageContact"),
		Version:     tea.String("2022-04-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateMessageContactResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a contact.
//
// @param request - UpdateMessageContactRequest
//
// @return UpdateMessageContactResponse
func (client *Client) UpdateMessageContact(request *UpdateMessageContactRequest) (_result *UpdateMessageContactResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateMessageContactResponse{}
	_body, _err := client.UpdateMessageContactWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
