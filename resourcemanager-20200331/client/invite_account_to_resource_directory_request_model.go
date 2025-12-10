// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInviteAccountToResourceDirectoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNote(v string) *InviteAccountToResourceDirectoryRequest
	GetNote() *string
	SetTag(v []*InviteAccountToResourceDirectoryRequestTag) *InviteAccountToResourceDirectoryRequest
	GetTag() []*InviteAccountToResourceDirectoryRequestTag
	SetTargetEntity(v string) *InviteAccountToResourceDirectoryRequest
	GetTargetEntity() *string
	SetTargetType(v string) *InviteAccountToResourceDirectoryRequest
	GetTargetType() *string
}

type InviteAccountToResourceDirectoryRequest struct {
	// The comment on the invitation.
	//
	// The comment can be up to 1,024 characters in length.
	//
	// example:
	//
	// Welcome
	Note *string `json:"Note,omitempty" xml:"Note,omitempty"`
	// The tag key and value.
	Tag []*InviteAccountToResourceDirectoryRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The ID or logon email address of the account that you want to invite.
	//
	// This parameter is required.
	//
	// example:
	//
	// someone@example.com
	TargetEntity *string `json:"TargetEntity,omitempty" xml:"TargetEntity,omitempty"`
	// The type of the account. Valid values:
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
	return dara.Prettify(s)
}

func (s InviteAccountToResourceDirectoryRequest) GoString() string {
	return s.String()
}

func (s *InviteAccountToResourceDirectoryRequest) GetNote() *string {
	return s.Note
}

func (s *InviteAccountToResourceDirectoryRequest) GetTag() []*InviteAccountToResourceDirectoryRequestTag {
	return s.Tag
}

func (s *InviteAccountToResourceDirectoryRequest) GetTargetEntity() *string {
	return s.TargetEntity
}

func (s *InviteAccountToResourceDirectoryRequest) GetTargetType() *string {
	return s.TargetType
}

func (s *InviteAccountToResourceDirectoryRequest) SetNote(v string) *InviteAccountToResourceDirectoryRequest {
	s.Note = &v
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

func (s *InviteAccountToResourceDirectoryRequest) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type InviteAccountToResourceDirectoryRequestTag struct {
	// A tag key.
	//
	// example:
	//
	// k1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// A tag value.
	//
	// example:
	//
	// v1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s InviteAccountToResourceDirectoryRequestTag) String() string {
	return dara.Prettify(s)
}

func (s InviteAccountToResourceDirectoryRequestTag) GoString() string {
	return s.String()
}

func (s *InviteAccountToResourceDirectoryRequestTag) GetKey() *string {
	return s.Key
}

func (s *InviteAccountToResourceDirectoryRequestTag) GetValue() *string {
	return s.Value
}

func (s *InviteAccountToResourceDirectoryRequestTag) SetKey(v string) *InviteAccountToResourceDirectoryRequestTag {
	s.Key = &v
	return s
}

func (s *InviteAccountToResourceDirectoryRequestTag) SetValue(v string) *InviteAccountToResourceDirectoryRequestTag {
	s.Value = &v
	return s
}

func (s *InviteAccountToResourceDirectoryRequestTag) Validate() error {
	return dara.Validate(s)
}
