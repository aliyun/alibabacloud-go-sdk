// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDirectoryId(v string) *DeleteGroupRequest
	GetDirectoryId() *string
	SetGroupId(v string) *DeleteGroupRequest
	GetGroupId() *string
}

type DeleteGroupRequest struct {
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

func (s DeleteGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteGroupRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *DeleteGroupRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *DeleteGroupRequest) SetDirectoryId(v string) *DeleteGroupRequest {
	s.DirectoryId = &v
	return s
}

func (s *DeleteGroupRequest) SetGroupId(v string) *DeleteGroupRequest {
	s.GroupId = &v
	return s
}

func (s *DeleteGroupRequest) Validate() error {
	return dara.Validate(s)
}
