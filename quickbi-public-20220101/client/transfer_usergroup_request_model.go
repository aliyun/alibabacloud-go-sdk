// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTransferUsergroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetParentUserGroupId(v string) *TransferUsergroupRequest
	GetParentUserGroupId() *string
	SetUserGroupId(v string) *TransferUsergroupRequest
	GetUserGroupId() *string
}

type TransferUsergroupRequest struct {
	// The ID of the parent user group.
	//
	// This parameter is required.
	//
	// example:
	//
	// asdasva***123124asdasd
	ParentUserGroupId *string `json:"ParentUserGroupId,omitempty" xml:"ParentUserGroupId,omitempty"`
	// The ID of the user group to migrate.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12qwda****sdada
	UserGroupId *string `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
}

func (s TransferUsergroupRequest) String() string {
	return dara.Prettify(s)
}

func (s TransferUsergroupRequest) GoString() string {
	return s.String()
}

func (s *TransferUsergroupRequest) GetParentUserGroupId() *string {
	return s.ParentUserGroupId
}

func (s *TransferUsergroupRequest) GetUserGroupId() *string {
	return s.UserGroupId
}

func (s *TransferUsergroupRequest) SetParentUserGroupId(v string) *TransferUsergroupRequest {
	s.ParentUserGroupId = &v
	return s
}

func (s *TransferUsergroupRequest) SetUserGroupId(v string) *TransferUsergroupRequest {
	s.UserGroupId = &v
	return s
}

func (s *TransferUsergroupRequest) Validate() error {
	return dara.Validate(s)
}
