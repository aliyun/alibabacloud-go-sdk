// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateYikeUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNickname(v string) *CreateYikeUserRequest
	GetNickname() *string
	SetPassword(v string) *CreateYikeUserRequest
	GetPassword() *string
	SetProductionIds(v string) *CreateYikeUserRequest
	GetProductionIds() *string
	SetUserNamePrefix(v string) *CreateYikeUserRequest
	GetUserNamePrefix() *string
	SetWorkspaceId(v string) *CreateYikeUserRequest
	GetWorkspaceId() *string
}

type CreateYikeUserRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// nick
	Nickname *string `json:"Nickname,omitempty" xml:"Nickname,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ***
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// id1,id2
	ProductionIds *string `json:"ProductionIds,omitempty" xml:"ProductionIds,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test
	UserNamePrefix *string `json:"UserNamePrefix,omitempty" xml:"UserNamePrefix,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// llm-az2gglkjauwnnhpq
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s CreateYikeUserRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateYikeUserRequest) GoString() string {
	return s.String()
}

func (s *CreateYikeUserRequest) GetNickname() *string {
	return s.Nickname
}

func (s *CreateYikeUserRequest) GetPassword() *string {
	return s.Password
}

func (s *CreateYikeUserRequest) GetProductionIds() *string {
	return s.ProductionIds
}

func (s *CreateYikeUserRequest) GetUserNamePrefix() *string {
	return s.UserNamePrefix
}

func (s *CreateYikeUserRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *CreateYikeUserRequest) SetNickname(v string) *CreateYikeUserRequest {
	s.Nickname = &v
	return s
}

func (s *CreateYikeUserRequest) SetPassword(v string) *CreateYikeUserRequest {
	s.Password = &v
	return s
}

func (s *CreateYikeUserRequest) SetProductionIds(v string) *CreateYikeUserRequest {
	s.ProductionIds = &v
	return s
}

func (s *CreateYikeUserRequest) SetUserNamePrefix(v string) *CreateYikeUserRequest {
	s.UserNamePrefix = &v
	return s
}

func (s *CreateYikeUserRequest) SetWorkspaceId(v string) *CreateYikeUserRequest {
	s.WorkspaceId = &v
	return s
}

func (s *CreateYikeUserRequest) Validate() error {
	return dara.Validate(s)
}
