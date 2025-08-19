// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHandshakeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetHandshake(v *GetHandshakeResponseBodyHandshake) *GetHandshakeResponseBody
	GetHandshake() *GetHandshakeResponseBodyHandshake
	SetRequestId(v string) *GetHandshakeResponseBody
	GetRequestId() *string
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
	return dara.Prettify(s)
}

func (s GetHandshakeResponseBody) GoString() string {
	return s.String()
}

func (s *GetHandshakeResponseBody) GetHandshake() *GetHandshakeResponseBodyHandshake {
	return s.Handshake
}

func (s *GetHandshakeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetHandshakeResponseBody) SetHandshake(v *GetHandshakeResponseBodyHandshake) *GetHandshakeResponseBody {
	s.Handshake = v
	return s
}

func (s *GetHandshakeResponseBody) SetRequestId(v string) *GetHandshakeResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetHandshakeResponseBody) Validate() error {
	return dara.Validate(s)
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
	return dara.Prettify(s)
}

func (s GetHandshakeResponseBodyHandshake) GoString() string {
	return s.String()
}

func (s *GetHandshakeResponseBodyHandshake) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetHandshakeResponseBodyHandshake) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *GetHandshakeResponseBodyHandshake) GetHandshakeId() *string {
	return s.HandshakeId
}

func (s *GetHandshakeResponseBodyHandshake) GetInvitedAccountRealName() *string {
	return s.InvitedAccountRealName
}

func (s *GetHandshakeResponseBodyHandshake) GetMasterAccountId() *string {
	return s.MasterAccountId
}

func (s *GetHandshakeResponseBodyHandshake) GetMasterAccountName() *string {
	return s.MasterAccountName
}

func (s *GetHandshakeResponseBodyHandshake) GetMasterAccountRealName() *string {
	return s.MasterAccountRealName
}

func (s *GetHandshakeResponseBodyHandshake) GetModifyTime() *string {
	return s.ModifyTime
}

func (s *GetHandshakeResponseBodyHandshake) GetNote() *string {
	return s.Note
}

func (s *GetHandshakeResponseBodyHandshake) GetResourceDirectoryId() *string {
	return s.ResourceDirectoryId
}

func (s *GetHandshakeResponseBodyHandshake) GetStatus() *string {
	return s.Status
}

func (s *GetHandshakeResponseBodyHandshake) GetTargetEntity() *string {
	return s.TargetEntity
}

func (s *GetHandshakeResponseBodyHandshake) GetTargetType() *string {
	return s.TargetType
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

func (s *GetHandshakeResponseBodyHandshake) Validate() error {
	return dara.Validate(s)
}
