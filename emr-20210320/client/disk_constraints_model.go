// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDiskConstraints interface {
	dara.Model
	String() string
	GoString() string
	SetCategories(v []*string) *DiskConstraints
	GetCategories() []*string
	SetCountConstraint(v *ValueConstraints) *DiskConstraints
	GetCountConstraint() *ValueConstraints
	SetSizeConstraint(v *ValueConstraints) *DiskConstraints
	GetSizeConstraint() *ValueConstraints
}

type DiskConstraints struct {
	// 支持的磁盘类型。
	//
	// example:
	//
	// ["cloud_efficiency","cloud_ssd","cloud_essd","local_disk"]
	Categories []*string `json:"Categories,omitempty" xml:"Categories,omitempty" type:"Repeated"`
	// 磁盘数量最小值。
	CountConstraint *ValueConstraints `json:"CountConstraint,omitempty" xml:"CountConstraint,omitempty"`
	// 磁盘容量限制。
	SizeConstraint *ValueConstraints `json:"SizeConstraint,omitempty" xml:"SizeConstraint,omitempty"`
}

func (s DiskConstraints) String() string {
	return dara.Prettify(s)
}

func (s DiskConstraints) GoString() string {
	return s.String()
}

func (s *DiskConstraints) GetCategories() []*string {
	return s.Categories
}

func (s *DiskConstraints) GetCountConstraint() *ValueConstraints {
	return s.CountConstraint
}

func (s *DiskConstraints) GetSizeConstraint() *ValueConstraints {
	return s.SizeConstraint
}

func (s *DiskConstraints) SetCategories(v []*string) *DiskConstraints {
	s.Categories = v
	return s
}

func (s *DiskConstraints) SetCountConstraint(v *ValueConstraints) *DiskConstraints {
	s.CountConstraint = v
	return s
}

func (s *DiskConstraints) SetSizeConstraint(v *ValueConstraints) *DiskConstraints {
	s.SizeConstraint = v
	return s
}

func (s *DiskConstraints) Validate() error {
	if s.CountConstraint != nil {
		if err := s.CountConstraint.Validate(); err != nil {
			return err
		}
	}
	if s.SizeConstraint != nil {
		if err := s.SizeConstraint.Validate(); err != nil {
			return err
		}
	}
	return nil
}
