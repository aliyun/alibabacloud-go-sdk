// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResizeDiskNodeGroupParam interface {
	dara.Model
	String() string
	GoString() string
	SetDataDiskCapacity(v int64) *ResizeDiskNodeGroupParam
	GetDataDiskCapacity() *int64
	SetNodeGroupId(v string) *ResizeDiskNodeGroupParam
	GetNodeGroupId() *string
	SetRollingRestart(v bool) *ResizeDiskNodeGroupParam
	GetRollingRestart() *bool
}

type ResizeDiskNodeGroupParam struct {
	DataDiskCapacity *int64  `json:"DataDiskCapacity,omitempty" xml:"DataDiskCapacity,omitempty"`
	NodeGroupId      *string `json:"NodeGroupId,omitempty" xml:"NodeGroupId,omitempty"`
	RollingRestart   *bool   `json:"RollingRestart,omitempty" xml:"RollingRestart,omitempty"`
}

func (s ResizeDiskNodeGroupParam) String() string {
	return dara.Prettify(s)
}

func (s ResizeDiskNodeGroupParam) GoString() string {
	return s.String()
}

func (s *ResizeDiskNodeGroupParam) GetDataDiskCapacity() *int64 {
	return s.DataDiskCapacity
}

func (s *ResizeDiskNodeGroupParam) GetNodeGroupId() *string {
	return s.NodeGroupId
}

func (s *ResizeDiskNodeGroupParam) GetRollingRestart() *bool {
	return s.RollingRestart
}

func (s *ResizeDiskNodeGroupParam) SetDataDiskCapacity(v int64) *ResizeDiskNodeGroupParam {
	s.DataDiskCapacity = &v
	return s
}

func (s *ResizeDiskNodeGroupParam) SetNodeGroupId(v string) *ResizeDiskNodeGroupParam {
	s.NodeGroupId = &v
	return s
}

func (s *ResizeDiskNodeGroupParam) SetRollingRestart(v bool) *ResizeDiskNodeGroupParam {
	s.RollingRestart = &v
	return s
}

func (s *ResizeDiskNodeGroupParam) Validate() error {
	return dara.Validate(s)
}
