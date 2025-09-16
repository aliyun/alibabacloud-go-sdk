// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteConfigDirRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDirName(v string) *DeleteConfigDirRequest
	GetDirName() *string
	SetParentFullPath(v string) *DeleteConfigDirRequest
	GetParentFullPath() *string
}

type DeleteConfigDirRequest struct {
	// The directory name.
	//
	// This parameter is required.
	//
	// example:
	//
	// /clusters
	DirName *string `json:"dirName,omitempty" xml:"dirName,omitempty"`
	// The path of the parent directory.
	//
	// This parameter is required.
	//
	// example:
	//
	// /
	ParentFullPath *string `json:"parentFullPath,omitempty" xml:"parentFullPath,omitempty"`
}

func (s DeleteConfigDirRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteConfigDirRequest) GoString() string {
	return s.String()
}

func (s *DeleteConfigDirRequest) GetDirName() *string {
	return s.DirName
}

func (s *DeleteConfigDirRequest) GetParentFullPath() *string {
	return s.ParentFullPath
}

func (s *DeleteConfigDirRequest) SetDirName(v string) *DeleteConfigDirRequest {
	s.DirName = &v
	return s
}

func (s *DeleteConfigDirRequest) SetParentFullPath(v string) *DeleteConfigDirRequest {
	s.ParentFullPath = &v
	return s
}

func (s *DeleteConfigDirRequest) Validate() error {
	return dara.Validate(s)
}
