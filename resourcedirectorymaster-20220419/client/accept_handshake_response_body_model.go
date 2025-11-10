// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAcceptHandshakeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetHandshake(v *AcceptHandshakeResponseBodyHandshake) *AcceptHandshakeResponseBody
	GetHandshake() *AcceptHandshakeResponseBodyHandshake
	SetRequestId(v string) *AcceptHandshakeResponseBody
	GetRequestId() *string
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
	return dara.Prettify(s)
}

func (s AcceptHandshakeResponseBody) GoString() string {
	return s.String()
}

func (s *AcceptHandshakeResponseBody) GetHandshake() *AcceptHandshakeResponseBodyHandshake {
	return s.Handshake
}

func (s *AcceptHandshakeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AcceptHandshakeResponseBody) SetHandshake(v *AcceptHandshakeResponseBodyHandshake) *AcceptHandshakeResponseBody {
	s.Handshake = v
	return s
}

func (s *AcceptHandshakeResponseBody) SetRequestId(v string) *AcceptHandshakeResponseBody {
	s.RequestId = &v
	return s
}

func (s *AcceptHandshakeResponseBody) Validate() error {
	if s.Handshake != nil {
		if err := s.Handshake.Validate(); err != nil {
			return err
		}
	}
	return nil
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
	return dara.Prettify(s)
}

func (s AcceptHandshakeResponseBodyHandshake) GoString() string {
	return s.String()
}

func (s *AcceptHandshakeResponseBodyHandshake) GetCreateTime() *string {
	return s.CreateTime
}

func (s *AcceptHandshakeResponseBodyHandshake) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *AcceptHandshakeResponseBodyHandshake) GetHandshakeId() *string {
	return s.HandshakeId
}

func (s *AcceptHandshakeResponseBodyHandshake) GetMasterAccountId() *string {
	return s.MasterAccountId
}

func (s *AcceptHandshakeResponseBodyHandshake) GetMasterAccountName() *string {
	return s.MasterAccountName
}

func (s *AcceptHandshakeResponseBodyHandshake) GetModifyTime() *string {
	return s.ModifyTime
}

func (s *AcceptHandshakeResponseBodyHandshake) GetNote() *string {
	return s.Note
}

func (s *AcceptHandshakeResponseBodyHandshake) GetResourceDirectoryId() *string {
	return s.ResourceDirectoryId
}

func (s *AcceptHandshakeResponseBodyHandshake) GetStatus() *string {
	return s.Status
}

func (s *AcceptHandshakeResponseBodyHandshake) GetTargetEntity() *string {
	return s.TargetEntity
}

func (s *AcceptHandshakeResponseBodyHandshake) GetTargetType() *string {
	return s.TargetType
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

func (s *AcceptHandshakeResponseBodyHandshake) Validate() error {
	return dara.Validate(s)
}
