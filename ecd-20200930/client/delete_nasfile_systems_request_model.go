// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteNASFileSystemsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileSystemId(v []*string) *DeleteNASFileSystemsRequest
	GetFileSystemId() []*string
	SetRegionId(v string) *DeleteNASFileSystemsRequest
	GetRegionId() *string
}

type DeleteNASFileSystemsRequest struct {
	// The IDs of the NAS file systems that you want to delete.
	//
	// This parameter is required.
	//
	// example:
	//
	// 04f314****
	FileSystemId []*string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty" type:"Repeated"`
	// The region ID of the NAS file system that you want to delete.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteNASFileSystemsRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteNASFileSystemsRequest) GoString() string {
	return s.String()
}

func (s *DeleteNASFileSystemsRequest) GetFileSystemId() []*string {
	return s.FileSystemId
}

func (s *DeleteNASFileSystemsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteNASFileSystemsRequest) SetFileSystemId(v []*string) *DeleteNASFileSystemsRequest {
	s.FileSystemId = v
	return s
}

func (s *DeleteNASFileSystemsRequest) SetRegionId(v string) *DeleteNASFileSystemsRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteNASFileSystemsRequest) Validate() error {
	return dara.Validate(s)
}
