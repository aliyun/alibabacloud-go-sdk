// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelHandshakeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetHandshake(v *CancelHandshakeResponseBodyHandshake) *CancelHandshakeResponseBody
	GetHandshake() *CancelHandshakeResponseBodyHandshake
	SetRequestId(v string) *CancelHandshakeResponseBody
	GetRequestId() *string
}

type CancelHandshakeResponseBody struct {
	// The handshake information.
	Handshake *CancelHandshakeResponseBodyHandshake `json:"Handshake,omitempty" xml:"Handshake,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 9B34724D-54B0-4A51-B34D-4512372FE1BE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CancelHandshakeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CancelHandshakeResponseBody) GoString() string {
	return s.String()
}

func (s *CancelHandshakeResponseBody) GetHandshake() *CancelHandshakeResponseBodyHandshake {
	return s.Handshake
}

func (s *CancelHandshakeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CancelHandshakeResponseBody) SetHandshake(v *CancelHandshakeResponseBodyHandshake) *CancelHandshakeResponseBody {
	s.Handshake = v
	return s
}

func (s *CancelHandshakeResponseBody) SetRequestId(v string) *CancelHandshakeResponseBody {
	s.RequestId = &v
	return s
}

func (s *CancelHandshakeResponseBody) Validate() error {
	if s.Handshake != nil {
		if err := s.Handshake.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CancelHandshakeResponseBodyHandshake struct {
	// The creation time of the handshake (UTC).
	//
	// example:
	//
	// 2018-08-10T09:55:41Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The expiration time of the handshake (UTC).
	//
	// example:
	//
	// 2018-08-24T09:55:41Z
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The handshake ID.
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
	// The modification time of the handshake (UTC).
	//
	// example:
	//
	// 2018-08-10T09:55:41Z
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The remarks.
	//
	// example:
	//
	// Welcome to join us
	Note *string `json:"Note,omitempty" xml:"Note,omitempty"`
	// The resource directory ID.
	//
	// example:
	//
	// h-ycm4rp****
	ResourceDirectoryId *string `json:"ResourceDirectoryId,omitempty" xml:"ResourceDirectoryId,omitempty"`
	// The status of the handshake. Valid values:
	//
	// - Pending: Waiting for confirmation.
	//
	// - Accepted: Accepted.
	//
	// - Cancelled: Cancelled.
	//
	// - Declined: Rejected.
	//
	// - Expired: Expired.
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
	// - Account: account ID.
	//
	// - Email: account logon email address.
	//
	// example:
	//
	// Email
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
}

func (s CancelHandshakeResponseBodyHandshake) String() string {
	return dara.Prettify(s)
}

func (s CancelHandshakeResponseBodyHandshake) GoString() string {
	return s.String()
}

func (s *CancelHandshakeResponseBodyHandshake) GetCreateTime() *string {
	return s.CreateTime
}

func (s *CancelHandshakeResponseBodyHandshake) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *CancelHandshakeResponseBodyHandshake) GetHandshakeId() *string {
	return s.HandshakeId
}

func (s *CancelHandshakeResponseBodyHandshake) GetMasterAccountId() *string {
	return s.MasterAccountId
}

func (s *CancelHandshakeResponseBodyHandshake) GetMasterAccountName() *string {
	return s.MasterAccountName
}

func (s *CancelHandshakeResponseBodyHandshake) GetModifyTime() *string {
	return s.ModifyTime
}

func (s *CancelHandshakeResponseBodyHandshake) GetNote() *string {
	return s.Note
}

func (s *CancelHandshakeResponseBodyHandshake) GetResourceDirectoryId() *string {
	return s.ResourceDirectoryId
}

func (s *CancelHandshakeResponseBodyHandshake) GetStatus() *string {
	return s.Status
}

func (s *CancelHandshakeResponseBodyHandshake) GetTargetEntity() *string {
	return s.TargetEntity
}

func (s *CancelHandshakeResponseBodyHandshake) GetTargetType() *string {
	return s.TargetType
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

func (s *CancelHandshakeResponseBodyHandshake) Validate() error {
	return dara.Validate(s)
}
