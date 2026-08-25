// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDirectoryId(v string) *GetGroupRequest
	GetDirectoryId() *string
	SetGroupId(v string) *GetGroupRequest
	GetGroupId() *string
}

type GetGroupRequest struct {
	// The ID of the directory.
	//
	// example:
	//
	// d-00fc2p61****
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	// The ID of the group.
	//
	// example:
	//
	// g-00jqzghi2n3o5hkh****
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
}

func (s GetGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s GetGroupRequest) GoString() string {
	return s.String()
}

func (s *GetGroupRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *GetGroupRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *GetGroupRequest) SetDirectoryId(v string) *GetGroupRequest {
	s.DirectoryId = &v
	return s
}

func (s *GetGroupRequest) SetGroupId(v string) *GetGroupRequest {
	s.GroupId = &v
	return s
}

func (s *GetGroupRequest) Validate() error {
	return dara.Validate(s)
}
