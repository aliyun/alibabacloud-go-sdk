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
	// Specifies whether to downsample the point cloud file. Valid values:
	//
	// 	- true: The point cloud file is downsampled, and the coordinates of the points in a voxel are replaced with the coordinates of the center point of the voxel. The average color of all points in the voxel is used as the color of the voxel. In this case, the PointResolution parameter does not take effect.
	//
	// 	- false: Specific coordinates and colors in a voxel are encoded by calculating the offsets from each point to the lower-left corner of the voxel. The offsets are divided by the PointResolution value to obtain the integer coordinates. The residual of the color for each point relative to the average color of all points in the voxel is encoded.
	//
	// example:
	//
	// false
	DoVoxelGridDownDownSampling *bool `json:"DoVoxelGridDownDownSampling,omitempty" xml:"DoVoxelGridDownDownSampling,omitempty"`
	// The library name. Set the value to pcl. Default value: pcl.
	//
	// example:
	//
	// pcl
	LibraryName *string `json:"LibraryName,omitempty" xml:"LibraryName,omitempty"`
	// The minimum block size when an octree is partitioned. The minimum block size indicates the edge length of a voxel. Default value: 0.01.
	//
	// example:
	//
	// 0.01
	OctreeResolution *float64 `json:"OctreeResolution,omitempty" xml:"OctreeResolution,omitempty"`
	// The point cloud resolution. This parameter determines the precision of the point coordinates during encoding. Default value: 0.01.
	//
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
