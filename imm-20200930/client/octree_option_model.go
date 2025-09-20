// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOctreeOption interface {
	dara.Model
	String() string
	GoString() string
	SetDoVoxelGridDownDownSampling(v bool) *OctreeOption
	GetDoVoxelGridDownDownSampling() *bool
	SetLibraryName(v string) *OctreeOption
	GetLibraryName() *string
	SetOctreeResolution(v float64) *OctreeOption
	GetOctreeResolution() *float64
	SetPointResolution(v float64) *OctreeOption
	GetPointResolution() *float64
}

type OctreeOption struct {
	// example:
	//
	// false
	DoVoxelGridDownDownSampling *bool `json:"DoVoxelGridDownDownSampling,omitempty" xml:"DoVoxelGridDownDownSampling,omitempty"`
	// example:
	//
	// pcl
	LibraryName *string `json:"LibraryName,omitempty" xml:"LibraryName,omitempty"`
	// example:
	//
	// 0.01
	OctreeResolution *float64 `json:"OctreeResolution,omitempty" xml:"OctreeResolution,omitempty"`
	// example:
	//
	// 0.01
	PointResolution *float64 `json:"PointResolution,omitempty" xml:"PointResolution,omitempty"`
}

func (s OctreeOption) String() string {
	return dara.Prettify(s)
}

func (s OctreeOption) GoString() string {
	return s.String()
}

func (s *OctreeOption) GetDoVoxelGridDownDownSampling() *bool {
	return s.DoVoxelGridDownDownSampling
}

func (s *OctreeOption) GetLibraryName() *string {
	return s.LibraryName
}

func (s *OctreeOption) GetOctreeResolution() *float64 {
	return s.OctreeResolution
}

func (s *OctreeOption) GetPointResolution() *float64 {
	return s.PointResolution
}

func (s *OctreeOption) SetDoVoxelGridDownDownSampling(v bool) *OctreeOption {
	s.DoVoxelGridDownDownSampling = &v
	return s
}

func (s *OctreeOption) SetLibraryName(v string) *OctreeOption {
	s.LibraryName = &v
	return s
}

func (s *OctreeOption) SetOctreeResolution(v float64) *OctreeOption {
	s.OctreeResolution = &v
	return s
}

func (s *OctreeOption) SetPointResolution(v float64) *OctreeOption {
	s.PointResolution = &v
	return s
}

func (s *OctreeOption) Validate() error {
	return dara.Validate(s)
}
