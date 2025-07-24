// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDiskSize interface {
	dara.Model
	String() string
	GoString() string
	SetCategory(v string) *DiskSize
	GetCategory() *string
	SetSize(v int32) *DiskSize
	GetSize() *int32
}

type DiskSize struct {
	// 磁盘类型。
	//
	// This parameter is required.
	//
	// example:
	//
	// cloud_essd
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// 单位GB。
	//
	// example:
	//
	// 80
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s DiskSize) String() string {
	return dara.Prettify(s)
}

func (s DiskSize) GoString() string {
	return s.String()
}

func (s *DiskSize) GetCategory() *string {
	return s.Category
}

func (s *DiskSize) GetSize() *int32 {
	return s.Size
}

func (s *DiskSize) SetCategory(v string) *DiskSize {
	s.Category = &v
	return s
}

func (s *DiskSize) SetSize(v int32) *DiskSize {
	s.Size = &v
	return s
}

func (s *DiskSize) Validate() error {
	return dara.Validate(s)
}
