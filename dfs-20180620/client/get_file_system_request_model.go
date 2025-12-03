// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFileSystemRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileSystemId(v string) *GetFileSystemRequest
	GetFileSystemId() *string
	SetInputRegionId(v string) *GetFileSystemRequest
	GetInputRegionId() *string
}

type GetFileSystemRequest struct {
	// This parameter is required.
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	InputRegionId *string `json:"InputRegionId,omitempty" xml:"InputRegionId,omitempty"`
}

func (s GetFileSystemRequest) String() string {
	return dara.Prettify(s)
}

func (s GetFileSystemRequest) GoString() string {
	return s.String()
}

func (s *GetFileSystemRequest) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *GetFileSystemRequest) GetInputRegionId() *string {
	return s.InputRegionId
}

func (s *GetFileSystemRequest) SetFileSystemId(v string) *GetFileSystemRequest {
	s.FileSystemId = &v
	return s
}

func (s *GetFileSystemRequest) SetInputRegionId(v string) *GetFileSystemRequest {
	s.InputRegionId = &v
	return s
}

func (s *GetFileSystemRequest) Validate() error {
	return dara.Validate(s)
}
