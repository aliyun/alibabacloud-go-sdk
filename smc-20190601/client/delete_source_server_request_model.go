// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSourceServerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetForce(v bool) *DeleteSourceServerRequest
	GetForce() *bool
	SetOwnerId(v int64) *DeleteSourceServerRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DeleteSourceServerRequest
	GetResourceOwnerAccount() *string
	SetSourceId(v string) *DeleteSourceServerRequest
	GetSourceId() *string
}

type DeleteSourceServerRequest struct {
	// Specifies whether to forcibly delete the migration source. Valid values:
	//
	// 	- true: forcibly deletes the migration source and the migration job created for the migration source, and releases the intermediate resources of the migration job.
	//
	// 	- false: does not delete the migration source if a migration job is created for the migration source.
	//
	// example:
	//
	// true
	Force                *bool   `json:"Force,omitempty" xml:"Force,omitempty"`
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

func (s DeleteSourceServerRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteSourceServerRequest) GoString() string {
	return s.String()
}

func (s *DeleteSourceServerRequest) GetForce() *bool {
	return s.Force
}

func (s *DeleteSourceServerRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteSourceServerRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteSourceServerRequest) GetSourceId() *string {
	return s.SourceId
}

func (s *DeleteSourceServerRequest) SetForce(v bool) *DeleteSourceServerRequest {
	s.Force = &v
	return s
}

func (s *DeleteSourceServerRequest) SetOwnerId(v int64) *DeleteSourceServerRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteSourceServerRequest) SetResourceOwnerAccount(v string) *DeleteSourceServerRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteSourceServerRequest) SetSourceId(v string) *DeleteSourceServerRequest {
	s.SourceId = &v
	return s
}

func (s *DeleteSourceServerRequest) Validate() error {
	return dara.Validate(s)
}
