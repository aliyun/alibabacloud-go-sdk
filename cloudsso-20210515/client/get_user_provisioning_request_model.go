// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserProvisioningRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDirectoryId(v string) *GetUserProvisioningRequest
	GetDirectoryId() *string
	SetUserProvisioningId(v string) *GetUserProvisioningRequest
	GetUserProvisioningId() *string
}

type GetUserProvisioningRequest struct {
	// The ID of the resource directory.
	//
	// example:
	//
	// d-003qew84****
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	// The ID of the RAM user provisioning.
	//
	// example:
	//
	// up-002axzhapcbz6e63****
	UserProvisioningId *string `json:"UserProvisioningId,omitempty" xml:"UserProvisioningId,omitempty"`
}

func (s GetUserProvisioningRequest) String() string {
	return dara.Prettify(s)
}

func (s GetUserProvisioningRequest) GoString() string {
	return s.String()
}

func (s *GetUserProvisioningRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *GetUserProvisioningRequest) GetUserProvisioningId() *string {
	return s.UserProvisioningId
}

func (s *GetUserProvisioningRequest) SetDirectoryId(v string) *GetUserProvisioningRequest {
	s.DirectoryId = &v
	return s
}

func (s *GetUserProvisioningRequest) SetUserProvisioningId(v string) *GetUserProvisioningRequest {
	s.UserProvisioningId = &v
	return s
}

func (s *GetUserProvisioningRequest) Validate() error {
	return dara.Validate(s)
}
