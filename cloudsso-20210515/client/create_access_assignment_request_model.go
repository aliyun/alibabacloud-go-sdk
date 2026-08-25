// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAccessAssignmentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessConfigurationId(v string) *CreateAccessAssignmentRequest
	GetAccessConfigurationId() *string
	SetDirectoryId(v string) *CreateAccessAssignmentRequest
	GetDirectoryId() *string
	SetOriginTargetId(v string) *CreateAccessAssignmentRequest
	GetOriginTargetId() *string
	SetPrincipalId(v string) *CreateAccessAssignmentRequest
	GetPrincipalId() *string
	SetPrincipalType(v string) *CreateAccessAssignmentRequest
	GetPrincipalType() *string
	SetTargetId(v string) *CreateAccessAssignmentRequest
	GetTargetId() *string
	SetTargetType(v string) *CreateAccessAssignmentRequest
	GetTargetType() *string
}

type CreateAccessAssignmentRequest struct {
	// The ID of the access configuration.
	//
	// example:
	//
	// ac-00jhtfl8thteu6uj****
	AccessConfigurationId *string `json:"AccessConfigurationId,omitempty" xml:"AccessConfigurationId,omitempty"`
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

func (s CreateAccessAssignmentRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAccessAssignmentRequest) GoString() string {
	return s.String()
}

func (s *CreateAccessAssignmentRequest) GetAccessConfigurationId() *string {
	return s.AccessConfigurationId
}

func (s *CreateAccessAssignmentRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *CreateAccessAssignmentRequest) GetOriginTargetId() *string {
	return s.OriginTargetId
}

func (s *CreateAccessAssignmentRequest) GetPrincipalId() *string {
	return s.PrincipalId
}

func (s *CreateAccessAssignmentRequest) GetPrincipalType() *string {
	return s.PrincipalType
}

func (s *CreateAccessAssignmentRequest) GetTargetId() *string {
	return s.TargetId
}

func (s *CreateAccessAssignmentRequest) GetTargetType() *string {
	return s.TargetType
}

func (s *CreateAccessAssignmentRequest) SetAccessConfigurationId(v string) *CreateAccessAssignmentRequest {
	s.AccessConfigurationId = &v
	return s
}

func (s *CreateAccessAssignmentRequest) SetDirectoryId(v string) *CreateAccessAssignmentRequest {
	s.DirectoryId = &v
	return s
}

func (s *CreateAccessAssignmentRequest) SetOriginTargetId(v string) *CreateAccessAssignmentRequest {
	s.OriginTargetId = &v
	return s
}

func (s *CreateAccessAssignmentRequest) SetPrincipalId(v string) *CreateAccessAssignmentRequest {
	s.PrincipalId = &v
	return s
}

func (s *CreateAccessAssignmentRequest) SetPrincipalType(v string) *CreateAccessAssignmentRequest {
	s.PrincipalType = &v
	return s
}

func (s *CreateAccessAssignmentRequest) SetTargetId(v string) *CreateAccessAssignmentRequest {
	s.TargetId = &v
	return s
}

func (s *CreateAccessAssignmentRequest) SetTargetType(v string) *CreateAccessAssignmentRequest {
	s.TargetType = &v
	return s
}

func (s *CreateAccessAssignmentRequest) Validate() error {
	return dara.Validate(s)
}
