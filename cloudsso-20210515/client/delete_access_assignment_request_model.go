// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAccessAssignmentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessConfigurationId(v string) *DeleteAccessAssignmentRequest
	GetAccessConfigurationId() *string
	SetDeprovisionStrategy(v string) *DeleteAccessAssignmentRequest
	GetDeprovisionStrategy() *string
	SetDirectoryId(v string) *DeleteAccessAssignmentRequest
	GetDirectoryId() *string
	SetOriginTargetId(v string) *DeleteAccessAssignmentRequest
	GetOriginTargetId() *string
	SetPrincipalId(v string) *DeleteAccessAssignmentRequest
	GetPrincipalId() *string
	SetPrincipalType(v string) *DeleteAccessAssignmentRequest
	GetPrincipalType() *string
	SetTargetId(v string) *DeleteAccessAssignmentRequest
	GetTargetId() *string
	SetTargetType(v string) *DeleteAccessAssignmentRequest
	GetTargetType() *string
}

type DeleteAccessAssignmentRequest struct {
	// The ID of the access configuration.
	//
	// example:
	//
	// ac-00jhtfl8thteu6uj****
	AccessConfigurationId *string `json:"AccessConfigurationId,omitempty" xml:"AccessConfigurationId,omitempty"`
	// Specifies whether to de-provision the access configuration when you remove the access permissions from the CloudSSO identity. The access configuration is used to assign the access permissions, and the identity is the only one that uses the access configuration and is associated with the account. Valid values:
	//
	// - DeprovisionForLastAccessAssignmentOnAccount: de-provisions the access configuration.
	//
	// - None: does not de-provision the access configuration. This is the default value.
	//
	// example:
	//
	// None
	DeprovisionStrategy *string `json:"DeprovisionStrategy,omitempty" xml:"DeprovisionStrategy,omitempty"`
	// The ID of the directory.
	//
	// example:
	//
	// d-00fc2p61****
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	// example:
	//
	// 114240524784****
	OriginTargetId *string `json:"OriginTargetId,omitempty" xml:"OriginTargetId,omitempty"`
	// The ID of the CloudSSO identity.
	//
	// - If you set `PrincipalType` to `User`, set `PrincipalId` to the ID of the CloudSSO user.
	//
	// - If you set `PrincipalType` to `Group`, set `PrincipalId` to the ID of the CloudSSO group.
	//
	// example:
	//
	// u-00q8wbq42wiltcrk****
	PrincipalId *string `json:"PrincipalId,omitempty" xml:"PrincipalId,omitempty"`
	// The type of the CloudSSO identity. Valid values:
	//
	// - User
	//
	// - Group
	//
	// example:
	//
	// User
	PrincipalType *string `json:"PrincipalType,omitempty" xml:"PrincipalType,omitempty"`
	// The ID of the task object.
	//
	// example:
	//
	// 114240524784****
	TargetId *string `json:"TargetId,omitempty" xml:"TargetId,omitempty"`
	// The type of the task object. Set the value to RD-Account, which specifies the accounts in the resource directory.
	//
	// example:
	//
	// RD-Account
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
}

func (s DeleteAccessAssignmentRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAccessAssignmentRequest) GoString() string {
	return s.String()
}

func (s *DeleteAccessAssignmentRequest) GetAccessConfigurationId() *string {
	return s.AccessConfigurationId
}

func (s *DeleteAccessAssignmentRequest) GetDeprovisionStrategy() *string {
	return s.DeprovisionStrategy
}

func (s *DeleteAccessAssignmentRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *DeleteAccessAssignmentRequest) GetOriginTargetId() *string {
	return s.OriginTargetId
}

func (s *DeleteAccessAssignmentRequest) GetPrincipalId() *string {
	return s.PrincipalId
}

func (s *DeleteAccessAssignmentRequest) GetPrincipalType() *string {
	return s.PrincipalType
}

func (s *DeleteAccessAssignmentRequest) GetTargetId() *string {
	return s.TargetId
}

func (s *DeleteAccessAssignmentRequest) GetTargetType() *string {
	return s.TargetType
}

func (s *DeleteAccessAssignmentRequest) SetAccessConfigurationId(v string) *DeleteAccessAssignmentRequest {
	s.AccessConfigurationId = &v
	return s
}

func (s *DeleteAccessAssignmentRequest) SetDeprovisionStrategy(v string) *DeleteAccessAssignmentRequest {
	s.DeprovisionStrategy = &v
	return s
}

func (s *DeleteAccessAssignmentRequest) SetDirectoryId(v string) *DeleteAccessAssignmentRequest {
	s.DirectoryId = &v
	return s
}

func (s *DeleteAccessAssignmentRequest) SetOriginTargetId(v string) *DeleteAccessAssignmentRequest {
	s.OriginTargetId = &v
	return s
}

func (s *DeleteAccessAssignmentRequest) SetPrincipalId(v string) *DeleteAccessAssignmentRequest {
	s.PrincipalId = &v
	return s
}

func (s *DeleteAccessAssignmentRequest) SetPrincipalType(v string) *DeleteAccessAssignmentRequest {
	s.PrincipalType = &v
	return s
}

func (s *DeleteAccessAssignmentRequest) SetTargetId(v string) *DeleteAccessAssignmentRequest {
	s.TargetId = &v
	return s
}

func (s *DeleteAccessAssignmentRequest) SetTargetType(v string) *DeleteAccessAssignmentRequest {
	s.TargetType = &v
	return s
}

func (s *DeleteAccessAssignmentRequest) Validate() error {
	return dara.Validate(s)
}
