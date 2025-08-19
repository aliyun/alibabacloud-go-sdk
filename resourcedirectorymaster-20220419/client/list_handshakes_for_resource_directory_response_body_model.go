// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHandshakesForResourceDirectoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetHandshakes(v *ListHandshakesForResourceDirectoryResponseBodyHandshakes) *ListHandshakesForResourceDirectoryResponseBody
	GetHandshakes() *ListHandshakesForResourceDirectoryResponseBodyHandshakes
	SetPageNumber(v int32) *ListHandshakesForResourceDirectoryResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListHandshakesForResourceDirectoryResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListHandshakesForResourceDirectoryResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListHandshakesForResourceDirectoryResponseBody
	GetTotalCount() *int32
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
	return dara.Prettify(s)
}

func (s ListHandshakesForResourceDirectoryResponseBody) GoString() string {
	return s.String()
}

func (s *ListHandshakesForResourceDirectoryResponseBody) GetHandshakes() *ListHandshakesForResourceDirectoryResponseBodyHandshakes {
	return s.Handshakes
}

func (s *ListHandshakesForResourceDirectoryResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListHandshakesForResourceDirectoryResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListHandshakesForResourceDirectoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListHandshakesForResourceDirectoryResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
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

func (s *ListHandshakesForResourceDirectoryResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListHandshakesForResourceDirectoryResponseBodyHandshakes struct {
	Handshake []*ListHandshakesForResourceDirectoryResponseBodyHandshakesHandshake `json:"Handshake,omitempty" xml:"Handshake,omitempty" type:"Repeated"`
}

func (s ListHandshakesForResourceDirectoryResponseBodyHandshakes) String() string {
	return dara.Prettify(s)
}

func (s ListHandshakesForResourceDirectoryResponseBodyHandshakes) GoString() string {
	return s.String()
}

func (s *ListHandshakesForResourceDirectoryResponseBodyHandshakes) GetHandshake() []*ListHandshakesForResourceDirectoryResponseBodyHandshakesHandshake {
	return s.Handshake
}

func (s *ListHandshakesForResourceDirectoryResponseBodyHandshakes) SetHandshake(v []*ListHandshakesForResourceDirectoryResponseBodyHandshakesHandshake) *ListHandshakesForResourceDirectoryResponseBodyHandshakes {
	s.Handshake = v
	return s
}

func (s *ListHandshakesForResourceDirectoryResponseBodyHandshakes) Validate() error {
	return dara.Validate(s)
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
	return dara.Prettify(s)
}

func (s ListHandshakesForResourceDirectoryResponseBodyHandshakesHandshake) GoString() string {
	return s.String()
}

func (s *ListHandshakesForResourceDirectoryResponseBodyHandshakesHandshake) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListHandshakesForResourceDirectoryResponseBodyHandshakesHandshake) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *ListHandshakesForResourceDirectoryResponseBodyHandshakesHandshake) GetHandshakeId() *string {
	return s.HandshakeId
}

func (s *ListHandshakesForResourceDirectoryResponseBodyHandshakesHandshake) GetMasterAccountId() *string {
	return s.MasterAccountId
}

func (s *ListHandshakesForResourceDirectoryResponseBodyHandshakesHandshake) GetMasterAccountName() *string {
	return s.MasterAccountName
}

func (s *ListHandshakesForResourceDirectoryResponseBodyHandshakesHandshake) GetModifyTime() *string {
	return s.ModifyTime
}

func (s *ListHandshakesForResourceDirectoryResponseBodyHandshakesHandshake) GetNote() *string {
	return s.Note
}

func (s *ListHandshakesForResourceDirectoryResponseBodyHandshakesHandshake) GetResourceDirectoryId() *string {
	return s.ResourceDirectoryId
}

func (s *ListHandshakesForResourceDirectoryResponseBodyHandshakesHandshake) GetStatus() *string {
	return s.Status
}

func (s *ListHandshakesForResourceDirectoryResponseBodyHandshakesHandshake) GetTargetEntity() *string {
	return s.TargetEntity
}

func (s *ListHandshakesForResourceDirectoryResponseBodyHandshakesHandshake) GetTargetType() *string {
	return s.TargetType
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

func (s *ListHandshakesForResourceDirectoryResponseBodyHandshakesHandshake) Validate() error {
	return dara.Validate(s)
}
