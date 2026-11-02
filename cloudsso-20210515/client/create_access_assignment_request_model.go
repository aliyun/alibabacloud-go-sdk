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
	// The access configuration ID.
	//
	// example:
	//
	// ac-00jhtfl8thteu6uj****
	AccessConfigurationId *string `json:"AccessConfigurationId,omitempty" xml:"AccessConfigurationId,omitempty"`
	// The directory ID.
	//
	// example:
	//
	// d-00fc2p61****
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	// The Alibaba Cloud account ID of the authorization target.
	//
	// example:
	//
	// 114240524784****
	OriginTargetId *string `json:"OriginTargetId,omitempty" xml:"OriginTargetId,omitempty"`
	// The CloudSSO identity ID. Details:
	//
	// - If PrincipalType is set to User, PrincipalId is the CloudSSO user ID.
	//
	// - If PrincipalType is set to Group, PrincipalId is the CloudSSO group ID.
	//
	// example:
	//
	// u-00q8wbq42wiltcrk****
	PrincipalId *string `json:"PrincipalId,omitempty" xml:"PrincipalId,omitempty"`
	// The CloudSSO identity type. Valid values:
	//
	// - User: user
	//
	// - Group: group
	//
	// example:
	//
	// User
	PrincipalType *string `json:"PrincipalType,omitempty" xml:"PrincipalType,omitempty"`
	// The ID of the deployment target.
	//
	// example:
	//
	// 114240524784****
	TargetId *string `json:"TargetId,omitempty" xml:"TargetId,omitempty"`
	// The target type of the task. Valid value: RD-Account, which indicates that the task target is an account in a resource directory.
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
