// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHandshakesForAccountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetHandshakes(v *ListHandshakesForAccountResponseBodyHandshakes) *ListHandshakesForAccountResponseBody
	GetHandshakes() *ListHandshakesForAccountResponseBodyHandshakes
	SetPageNumber(v int32) *ListHandshakesForAccountResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListHandshakesForAccountResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListHandshakesForAccountResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListHandshakesForAccountResponseBody
	GetTotalCount() *int32
}

type ListHandshakesForAccountResponseBody struct {
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
	return dara.Prettify(s)
}

func (s ListHandshakesForAccountResponseBody) GoString() string {
	return s.String()
}

func (s *ListHandshakesForAccountResponseBody) GetHandshakes() *ListHandshakesForAccountResponseBodyHandshakes {
	return s.Handshakes
}

func (s *ListHandshakesForAccountResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListHandshakesForAccountResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListHandshakesForAccountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListHandshakesForAccountResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
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

func (s *ListHandshakesForAccountResponseBody) Validate() error {
	if s.Handshakes != nil {
		if err := s.Handshakes.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListHandshakesForAccountResponseBodyHandshakes struct {
	Handshake []*ListHandshakesForAccountResponseBodyHandshakesHandshake `json:"Handshake,omitempty" xml:"Handshake,omitempty" type:"Repeated"`
}

func (s ListHandshakesForAccountResponseBodyHandshakes) String() string {
	return dara.Prettify(s)
}

func (s ListHandshakesForAccountResponseBodyHandshakes) GoString() string {
	return s.String()
}

func (s *ListHandshakesForAccountResponseBodyHandshakes) GetHandshake() []*ListHandshakesForAccountResponseBodyHandshakesHandshake {
	return s.Handshake
}

func (s *ListHandshakesForAccountResponseBodyHandshakes) SetHandshake(v []*ListHandshakesForAccountResponseBodyHandshakesHandshake) *ListHandshakesForAccountResponseBodyHandshakes {
	s.Handshake = v
	return s
}

func (s *ListHandshakesForAccountResponseBodyHandshakes) Validate() error {
	if s.Handshake != nil {
		for _, item := range s.Handshake {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListHandshakesForAccountResponseBodyHandshakesHandshake struct {
	CreateTime          *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	ExpireTime          *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	HandshakeId         *string `json:"HandshakeId,omitempty" xml:"HandshakeId,omitempty"`
	MasterAccountId     *string `json:"MasterAccountId,omitempty" xml:"MasterAccountId,omitempty"`
	MasterAccountName   *string `json:"MasterAccountName,omitempty" xml:"MasterAccountName,omitempty"`
	ModifyTime          *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	Note                *string `json:"Note,omitempty" xml:"Note,omitempty"`
	ResourceDirectoryId *string `json:"ResourceDirectoryId,omitempty" xml:"ResourceDirectoryId,omitempty"`
	Status              *string `json:"Status,omitempty" xml:"Status,omitempty"`
	TargetEntity        *string `json:"TargetEntity,omitempty" xml:"TargetEntity,omitempty"`
	TargetType          *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
}

func (s ListHandshakesForAccountResponseBodyHandshakesHandshake) String() string {
	return dara.Prettify(s)
}

func (s ListHandshakesForAccountResponseBodyHandshakesHandshake) GoString() string {
	return s.String()
}

func (s *ListHandshakesForAccountResponseBodyHandshakesHandshake) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListHandshakesForAccountResponseBodyHandshakesHandshake) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *ListHandshakesForAccountResponseBodyHandshakesHandshake) GetHandshakeId() *string {
	return s.HandshakeId
}

func (s *ListHandshakesForAccountResponseBodyHandshakesHandshake) GetMasterAccountId() *string {
	return s.MasterAccountId
}

func (s *ListHandshakesForAccountResponseBodyHandshakesHandshake) GetMasterAccountName() *string {
	return s.MasterAccountName
}

func (s *ListHandshakesForAccountResponseBodyHandshakesHandshake) GetModifyTime() *string {
	return s.ModifyTime
}

func (s *ListHandshakesForAccountResponseBodyHandshakesHandshake) GetNote() *string {
	return s.Note
}

func (s *ListHandshakesForAccountResponseBodyHandshakesHandshake) GetResourceDirectoryId() *string {
	return s.ResourceDirectoryId
}

func (s *ListHandshakesForAccountResponseBodyHandshakesHandshake) GetStatus() *string {
	return s.Status
}

func (s *ListHandshakesForAccountResponseBodyHandshakesHandshake) GetTargetEntity() *string {
	return s.TargetEntity
}

func (s *ListHandshakesForAccountResponseBodyHandshakesHandshake) GetTargetType() *string {
	return s.TargetType
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

func (s *ListHandshakesForAccountResponseBodyHandshakesHandshake) Validate() error {
	return dara.Validate(s)
}
