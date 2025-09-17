// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySourceServerAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *ModifySourceServerAttributeRequest
	GetDescription() *string
	SetName(v string) *ModifySourceServerAttributeRequest
	GetName() *string
	SetOwnerId(v int64) *ModifySourceServerAttributeRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ModifySourceServerAttributeRequest
	GetResourceOwnerAccount() *string
	SetSourceId(v string) *ModifySourceServerAttributeRequest
	GetSourceId() *string
}

type ModifySourceServerAttributeRequest struct {
	// The description of the migration source. The description can be up to 256 characters in length and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// This is a source server.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the migration source. The name must be 2 to 128 characters in length. The name must start with a letter and cannot start with `http://` or `https://`. It can contain letters, digits, colons (:), underscores (_), and hyphens (-).
	//
	// example:
	//
	// testSourceServerName
	Name                 *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	// The migration source ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// s-bp17m1vi6x20c6g6****
	SourceId *string `json:"SourceId,omitempty" xml:"SourceId,omitempty"`
}

func (s ModifySourceServerAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifySourceServerAttributeRequest) GoString() string {
	return s.String()
}

func (s *ModifySourceServerAttributeRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifySourceServerAttributeRequest) GetName() *string {
	return s.Name
}

func (s *ModifySourceServerAttributeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifySourceServerAttributeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifySourceServerAttributeRequest) GetSourceId() *string {
	return s.SourceId
}

func (s *ModifySourceServerAttributeRequest) SetDescription(v string) *ModifySourceServerAttributeRequest {
	s.Description = &v
	return s
}

func (s *ModifySourceServerAttributeRequest) SetName(v string) *ModifySourceServerAttributeRequest {
	s.Name = &v
	return s
}

func (s *ModifySourceServerAttributeRequest) SetOwnerId(v int64) *ModifySourceServerAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifySourceServerAttributeRequest) SetResourceOwnerAccount(v string) *ModifySourceServerAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifySourceServerAttributeRequest) SetSourceId(v string) *ModifySourceServerAttributeRequest {
	s.SourceId = &v
	return s
}

func (s *ModifySourceServerAttributeRequest) Validate() error {
	return dara.Validate(s)
}
