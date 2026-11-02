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
	// The access configuration ID.
	//
	// example:
	//
	// ac-00jhtfl8thteu6uj****
	AccessConfigurationId *string `json:"AccessConfigurationId,omitempty" xml:"AccessConfigurationId,omitempty"`
	// Specifies whether to remove the access configuration deployment when you remove the last authorization that uses the access configuration on an account in a resource directory. Valid values:
	//
	// - DeprovisionForLastAccessAssignmentOnAccount: Removes the access configuration deployment.
	//
	// - None: Does not remove the access configuration deployment. This is the default value.
	//
	// example:
	//
	// None
	DeprovisionStrategy *string `json:"DeprovisionStrategy,omitempty" xml:"DeprovisionStrategy,omitempty"`
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
	// The CloudSSO identity ID. The value varies based on the principal type:
	//
	// - If `PrincipalType` is `User`, `PrincipalId` is the CloudSSO user ID.
	//
	// - If `PrincipalType` is `Group`, `PrincipalId` is the CloudSSO group ID.
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
	// The target ID of the task.
	//
	// example:
	//
	// 114240524784****
	TargetId *string `json:"TargetId,omitempty" xml:"TargetId,omitempty"`
	// The target type of the task. Valid values:
	//
	// - RD-Account: The task target is an account in a resource directory.
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
