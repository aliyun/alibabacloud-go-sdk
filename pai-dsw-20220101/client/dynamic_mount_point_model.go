// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDynamicMountPoint interface {
	dara.Model
	String() string
	GoString() string
	SetOptions(v string) *DynamicMountPoint
	GetOptions() *string
	SetRootPath(v string) *DynamicMountPoint
	GetRootPath() *string
}

type DynamicMountPoint struct {
	Options *string `json:"Options,omitempty" xml:"Options,omitempty"`
	// This parameter is required.
	RootPath *string `json:"RootPath,omitempty" xml:"RootPath,omitempty"`
}

func (s DynamicMountPoint) String() string {
	return dara.Prettify(s)
}

func (s DynamicMountPoint) GoString() string {
	return s.String()
}

func (s *DynamicMountPoint) GetOptions() *string {
	return s.Options
}

func (s *DynamicMountPoint) GetRootPath() *string {
	return s.RootPath
}

func (s *DynamicMountPoint) SetOptions(v string) *DynamicMountPoint {
	s.Options = &v
	return s
}

func (s *DynamicMountPoint) SetRootPath(v string) *DynamicMountPoint {
	s.RootPath = &v
	return s
}

func (s *DynamicMountPoint) Validate() error {
	return dara.Validate(s)
}
