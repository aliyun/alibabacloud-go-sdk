// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyFileSystemRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *ModifyFileSystemRequest
	GetDescription() *string
	SetEnsRegionId(v string) *ModifyFileSystemRequest
	GetEnsRegionId() *string
	SetFileSystemId(v string) *ModifyFileSystemRequest
	GetFileSystemId() *string
}

type ModifyFileSystemRequest struct {
	// The description of the file system.
	//
	// This parameter is required.
	//
	// example:
	//
	// fileSystemTest
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the edge node.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing-cmcc
	EnsRegionId *string `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty"`
	// The ID of the file system.
	//
	// This parameter is required.
	//
	// example:
	//
	// c50f8*****
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
}

func (s ModifyFileSystemRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyFileSystemRequest) GoString() string {
	return s.String()
}

func (s *ModifyFileSystemRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifyFileSystemRequest) GetEnsRegionId() *string {
	return s.EnsRegionId
}

func (s *ModifyFileSystemRequest) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *ModifyFileSystemRequest) SetDescription(v string) *ModifyFileSystemRequest {
	s.Description = &v
	return s
}

func (s *ModifyFileSystemRequest) SetEnsRegionId(v string) *ModifyFileSystemRequest {
	s.EnsRegionId = &v
	return s
}

func (s *ModifyFileSystemRequest) SetFileSystemId(v string) *ModifyFileSystemRequest {
	s.FileSystemId = &v
	return s
}

func (s *ModifyFileSystemRequest) Validate() error {
	return dara.Validate(s)
}
