// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateConfigDirRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDirName(v string) *CreateConfigDirRequest
	GetDirName() *string
	SetParentFullPath(v string) *CreateConfigDirRequest
	GetParentFullPath() *string
}

type CreateConfigDirRequest struct {
	// The directory name.
	//
	// example:
	//
	// /clusters
	DirName *string `json:"dirName,omitempty" xml:"dirName,omitempty"`
	// The path of the parent directory.
	//
	// example:
	//
	// /
	ParentFullPath *string `json:"parentFullPath,omitempty" xml:"parentFullPath,omitempty"`
}

func (s CreateConfigDirRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateConfigDirRequest) GoString() string {
	return s.String()
}

func (s *CreateConfigDirRequest) GetDirName() *string {
	return s.DirName
}

func (s *CreateConfigDirRequest) GetParentFullPath() *string {
	return s.ParentFullPath
}

func (s *CreateConfigDirRequest) SetDirName(v string) *CreateConfigDirRequest {
	s.DirName = &v
	return s
}

func (s *CreateConfigDirRequest) SetParentFullPath(v string) *CreateConfigDirRequest {
	s.ParentFullPath = &v
	return s
}

func (s *CreateConfigDirRequest) Validate() error {
	return dara.Validate(s)
}
