// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetNASDefaultMountTargetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileSystemId(v string) *ResetNASDefaultMountTargetRequest
	GetFileSystemId() *string
	SetRegionId(v string) *ResetNASDefaultMountTargetRequest
	GetRegionId() *string
}

type ResetNASDefaultMountTargetRequest struct {
	// The ID of the NAS file system.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3a6ef4****
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ResetNASDefaultMountTargetRequest) String() string {
	return dara.Prettify(s)
}

func (s ResetNASDefaultMountTargetRequest) GoString() string {
	return s.String()
}

func (s *ResetNASDefaultMountTargetRequest) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *ResetNASDefaultMountTargetRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ResetNASDefaultMountTargetRequest) SetFileSystemId(v string) *ResetNASDefaultMountTargetRequest {
	s.FileSystemId = &v
	return s
}

func (s *ResetNASDefaultMountTargetRequest) SetRegionId(v string) *ResetNASDefaultMountTargetRequest {
	s.RegionId = &v
	return s
}

func (s *ResetNASDefaultMountTargetRequest) Validate() error {
	return dara.Validate(s)
}
