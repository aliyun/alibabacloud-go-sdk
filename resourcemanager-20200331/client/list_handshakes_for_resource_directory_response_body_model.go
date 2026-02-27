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
	if s.Handshakes != nil {
		if err := s.Handshakes.Validate(); err != nil {
			return err
		}
	}
	return nil
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

type ListHandshakesForResourceDirectoryResponseBodyHandshakesHandshake struct {
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
