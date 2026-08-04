// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMemberApiKeyDTO interface {
	dara.Model
	String() string
	GoString() string
	SetClient(v *ClientDTO) *MemberApiKeyDTO
	GetClient() *ClientDTO
	SetClientId(v int64) *MemberApiKeyDTO
	GetClientId() *int64
	SetDeleteTag(v int32) *MemberApiKeyDTO
	GetDeleteTag() *int32
	SetExpireAt(v string) *MemberApiKeyDTO
	GetExpireAt() *string
	SetGmtCreate(v string) *MemberApiKeyDTO
	GetGmtCreate() *string
	SetGmtModified(v string) *MemberApiKeyDTO
	GetGmtModified() *string
	SetId(v int64) *MemberApiKeyDTO
	GetId() *int64
	SetKey(v string) *MemberApiKeyDTO
	GetKey() *string
	SetKeyPreview(v string) *MemberApiKeyDTO
	GetKeyPreview() *string
	SetMemberUserId(v int64) *MemberApiKeyDTO
	GetMemberUserId() *int64
	SetMemberUserName(v string) *MemberApiKeyDTO
	GetMemberUserName() *string
	SetName(v string) *MemberApiKeyDTO
	GetName() *string
	SetStatus(v string) *MemberApiKeyDTO
	GetStatus() *string
}

type MemberApiKeyDTO struct {
	Client *ClientDTO `json:"client,omitempty" xml:"client,omitempty"`
	// example:
	//
	// 438
	ClientId *int64 `json:"clientId,omitempty" xml:"clientId,omitempty"`
	// example:
	//
	// 0
	DeleteTag *int32 `json:"deleteTag,omitempty" xml:"deleteTag,omitempty"`
	// example:
	//
	// 2026-12-31T00:00:00Z
	ExpireAt *string `json:"expireAt,omitempty" xml:"expireAt,omitempty"`
	// example:
	//
	// 2026-08-03T18:41:40+08:00
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// example:
	//
	// 2026-08-03T18:41:40+08:00
	GmtModified *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// example:
	//
	// 502
	Id  *int64  `json:"id,omitempty" xml:"id,omitempty"`
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// example:
	//
	// sk-us****9876
	KeyPreview *string `json:"keyPreview,omitempty" xml:"keyPreview,omitempty"`
	// example:
	//
	// 304
	MemberUserId *int64 `json:"memberUserId,omitempty" xml:"memberUserId,omitempty"`
	// example:
	//
	// John
	MemberUserName *string `json:"memberUserName,omitempty" xml:"memberUserName,omitempty"`
	// example:
	//
	// John\\"s Key
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// active
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s MemberApiKeyDTO) String() string {
	return dara.Prettify(s)
}

func (s MemberApiKeyDTO) GoString() string {
	return s.String()
}

func (s *MemberApiKeyDTO) GetClient() *ClientDTO {
	return s.Client
}

func (s *MemberApiKeyDTO) GetClientId() *int64 {
	return s.ClientId
}

func (s *MemberApiKeyDTO) GetDeleteTag() *int32 {
	return s.DeleteTag
}

func (s *MemberApiKeyDTO) GetExpireAt() *string {
	return s.ExpireAt
}

func (s *MemberApiKeyDTO) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *MemberApiKeyDTO) GetGmtModified() *string {
	return s.GmtModified
}

func (s *MemberApiKeyDTO) GetId() *int64 {
	return s.Id
}

func (s *MemberApiKeyDTO) GetKey() *string {
	return s.Key
}

func (s *MemberApiKeyDTO) GetKeyPreview() *string {
	return s.KeyPreview
}

func (s *MemberApiKeyDTO) GetMemberUserId() *int64 {
	return s.MemberUserId
}

func (s *MemberApiKeyDTO) GetMemberUserName() *string {
	return s.MemberUserName
}

func (s *MemberApiKeyDTO) GetName() *string {
	return s.Name
}

func (s *MemberApiKeyDTO) GetStatus() *string {
	return s.Status
}

func (s *MemberApiKeyDTO) SetClient(v *ClientDTO) *MemberApiKeyDTO {
	s.Client = v
	return s
}

func (s *MemberApiKeyDTO) SetClientId(v int64) *MemberApiKeyDTO {
	s.ClientId = &v
	return s
}

func (s *MemberApiKeyDTO) SetDeleteTag(v int32) *MemberApiKeyDTO {
	s.DeleteTag = &v
	return s
}

func (s *MemberApiKeyDTO) SetExpireAt(v string) *MemberApiKeyDTO {
	s.ExpireAt = &v
	return s
}

func (s *MemberApiKeyDTO) SetGmtCreate(v string) *MemberApiKeyDTO {
	s.GmtCreate = &v
	return s
}

func (s *MemberApiKeyDTO) SetGmtModified(v string) *MemberApiKeyDTO {
	s.GmtModified = &v
	return s
}

func (s *MemberApiKeyDTO) SetId(v int64) *MemberApiKeyDTO {
	s.Id = &v
	return s
}

func (s *MemberApiKeyDTO) SetKey(v string) *MemberApiKeyDTO {
	s.Key = &v
	return s
}

func (s *MemberApiKeyDTO) SetKeyPreview(v string) *MemberApiKeyDTO {
	s.KeyPreview = &v
	return s
}

func (s *MemberApiKeyDTO) SetMemberUserId(v int64) *MemberApiKeyDTO {
	s.MemberUserId = &v
	return s
}

func (s *MemberApiKeyDTO) SetMemberUserName(v string) *MemberApiKeyDTO {
	s.MemberUserName = &v
	return s
}

func (s *MemberApiKeyDTO) SetName(v string) *MemberApiKeyDTO {
	s.Name = &v
	return s
}

func (s *MemberApiKeyDTO) SetStatus(v string) *MemberApiKeyDTO {
	s.Status = &v
	return s
}

func (s *MemberApiKeyDTO) Validate() error {
	if s.Client != nil {
		if err := s.Client.Validate(); err != nil {
			return err
		}
	}
	return nil
}
