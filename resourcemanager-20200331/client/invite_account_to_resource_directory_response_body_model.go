// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInviteAccountToResourceDirectoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetHandshake(v *InviteAccountToResourceDirectoryResponseBodyHandshake) *InviteAccountToResourceDirectoryResponseBody
	GetHandshake() *InviteAccountToResourceDirectoryResponseBodyHandshake
	SetRequestId(v string) *InviteAccountToResourceDirectoryResponseBody
	GetRequestId() *string
}

type InviteAccountToResourceDirectoryResponseBody struct {
	// The information of the invitation.
	Handshake *InviteAccountToResourceDirectoryResponseBodyHandshake `json:"Handshake,omitempty" xml:"Handshake,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 9B34724D-54B0-4A51-B34D-4512372FE1BE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s InviteAccountToResourceDirectoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s InviteAccountToResourceDirectoryResponseBody) GoString() string {
	return s.String()
}

func (s *InviteAccountToResourceDirectoryResponseBody) GetHandshake() *InviteAccountToResourceDirectoryResponseBodyHandshake {
	return s.Handshake
}

func (s *InviteAccountToResourceDirectoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *InviteAccountToResourceDirectoryResponseBody) SetHandshake(v *InviteAccountToResourceDirectoryResponseBodyHandshake) *InviteAccountToResourceDirectoryResponseBody {
	s.Handshake = v
	return s
}

func (s *InviteAccountToResourceDirectoryResponseBody) SetRequestId(v string) *InviteAccountToResourceDirectoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *InviteAccountToResourceDirectoryResponseBody) Validate() error {
	if s.Handshake != nil {
		if err := s.Handshake.Validate(); err != nil {
			return err
		}
	}
	return nil
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
	// The comment on the invitation.
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
	// The status of the invitation.
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

func (s InviteAccountToResourceDirectoryResponseBodyHandshake) String() string {
	return dara.Prettify(s)
}

func (s InviteAccountToResourceDirectoryResponseBodyHandshake) GoString() string {
	return s.String()
}

func (s *InviteAccountToResourceDirectoryResponseBodyHandshake) GetCreateTime() *string {
	return s.CreateTime
}

func (s *InviteAccountToResourceDirectoryResponseBodyHandshake) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *InviteAccountToResourceDirectoryResponseBodyHandshake) GetHandshakeId() *string {
	return s.HandshakeId
}

func (s *InviteAccountToResourceDirectoryResponseBodyHandshake) GetMasterAccountId() *string {
	return s.MasterAccountId
}

func (s *InviteAccountToResourceDirectoryResponseBodyHandshake) GetMasterAccountName() *string {
	return s.MasterAccountName
}

func (s *InviteAccountToResourceDirectoryResponseBodyHandshake) GetModifyTime() *string {
	return s.ModifyTime
}

func (s *InviteAccountToResourceDirectoryResponseBodyHandshake) GetNote() *string {
	return s.Note
}

func (s *InviteAccountToResourceDirectoryResponseBodyHandshake) GetResourceDirectoryId() *string {
	return s.ResourceDirectoryId
}

func (s *InviteAccountToResourceDirectoryResponseBodyHandshake) GetStatus() *string {
	return s.Status
}

func (s *InviteAccountToResourceDirectoryResponseBodyHandshake) GetTargetEntity() *string {
	return s.TargetEntity
}

func (s *InviteAccountToResourceDirectoryResponseBodyHandshake) GetTargetType() *string {
	return s.TargetType
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

func (s *InviteAccountToResourceDirectoryResponseBodyHandshake) Validate() error {
	return dara.Validate(s)
}
