// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeclineHandshakeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetHandshake(v *DeclineHandshakeResponseBodyHandshake) *DeclineHandshakeResponseBody
	GetHandshake() *DeclineHandshakeResponseBodyHandshake
	SetRequestId(v string) *DeclineHandshakeResponseBody
	GetRequestId() *string
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
	return dara.Prettify(s)
}

func (s DeclineHandshakeResponseBody) GoString() string {
	return s.String()
}

func (s *DeclineHandshakeResponseBody) GetHandshake() *DeclineHandshakeResponseBodyHandshake {
	return s.Handshake
}

func (s *DeclineHandshakeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeclineHandshakeResponseBody) SetHandshake(v *DeclineHandshakeResponseBodyHandshake) *DeclineHandshakeResponseBody {
	s.Handshake = v
	return s
}

func (s *DeclineHandshakeResponseBody) SetRequestId(v string) *DeclineHandshakeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeclineHandshakeResponseBody) Validate() error {
	if s.Handshake != nil {
		if err := s.Handshake.Validate(); err != nil {
			return err
		}
	}
	return nil
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
	return dara.Prettify(s)
}

func (s DeclineHandshakeResponseBodyHandshake) GoString() string {
	return s.String()
}

func (s *DeclineHandshakeResponseBodyHandshake) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DeclineHandshakeResponseBodyHandshake) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *DeclineHandshakeResponseBodyHandshake) GetHandshakeId() *string {
	return s.HandshakeId
}

func (s *DeclineHandshakeResponseBodyHandshake) GetMasterAccountId() *string {
	return s.MasterAccountId
}

func (s *DeclineHandshakeResponseBodyHandshake) GetMasterAccountName() *string {
	return s.MasterAccountName
}

func (s *DeclineHandshakeResponseBodyHandshake) GetModifyTime() *string {
	return s.ModifyTime
}

func (s *DeclineHandshakeResponseBodyHandshake) GetNote() *string {
	return s.Note
}

func (s *DeclineHandshakeResponseBodyHandshake) GetResourceDirectoryId() *string {
	return s.ResourceDirectoryId
}

func (s *DeclineHandshakeResponseBodyHandshake) GetStatus() *string {
	return s.Status
}

func (s *DeclineHandshakeResponseBodyHandshake) GetTargetEntity() *string {
	return s.TargetEntity
}

func (s *DeclineHandshakeResponseBodyHandshake) GetTargetType() *string {
	return s.TargetType
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

func (s *DeclineHandshakeResponseBodyHandshake) Validate() error {
	return dara.Validate(s)
}
