// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDataDisk interface {
	dara.Model
	String() string
	GoString() string
	SetSize(v int64) *DataDisk
	GetSize() *int64
}

type DataDisk struct {
	// example:
	//
	// 60
	Size *int64 `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s DataDisk) String() string {
	return dara.Prettify(s)
}

func (s DataDisk) GoString() string {
	return s.String()
}

func (s *DataDisk) GetSize() *int64 {
	return s.Size
}

func (s *DataDisk) SetSize(v int64) *DataDisk {
	s.Size = &v
	return s
}

func (s *DataDisk) Validate() error {
	return dara.Validate(s)
}
