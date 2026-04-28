// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPersonalSpaceInfo interface {
	dara.Model
	String() string
	GoString() string
	SetTotalSize(v int64) *PersonalSpaceInfo
	GetTotalSize() *int64
	SetUsedSize(v int64) *PersonalSpaceInfo
	GetUsedSize() *int64
}

type PersonalSpaceInfo struct {
	TotalSize *int64 `json:"total_size,omitempty" xml:"total_size,omitempty"`
	UsedSize  *int64 `json:"used_size,omitempty" xml:"used_size,omitempty"`
}

func (s PersonalSpaceInfo) String() string {
	return dara.Prettify(s)
}

func (s PersonalSpaceInfo) GoString() string {
	return s.String()
}

func (s *PersonalSpaceInfo) GetTotalSize() *int64 {
	return s.TotalSize
}

func (s *PersonalSpaceInfo) GetUsedSize() *int64 {
	return s.UsedSize
}

func (s *PersonalSpaceInfo) SetTotalSize(v int64) *PersonalSpaceInfo {
	s.TotalSize = &v
	return s
}

func (s *PersonalSpaceInfo) SetUsedSize(v int64) *PersonalSpaceInfo {
	s.UsedSize = &v
	return s
}

func (s *PersonalSpaceInfo) Validate() error {
	return dara.Validate(s)
}
