// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDirectoriesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDirectoryId(v []*string) *DeleteDirectoriesRequest
	GetDirectoryId() []*string
	SetRegionId(v string) *DeleteDirectoriesRequest
	GetRegionId() *string
}

type DeleteDirectoriesRequest struct {
	// The directory IDs. You can specify one or more directory IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou+dir-gx2x1dhsmu52rd****
	DirectoryId []*string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty" type:"Repeated"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/196646.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteDirectoriesRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDirectoriesRequest) GoString() string {
	return s.String()
}

func (s *DeleteDirectoriesRequest) GetDirectoryId() []*string {
	return s.DirectoryId
}

func (s *DeleteDirectoriesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteDirectoriesRequest) SetDirectoryId(v []*string) *DeleteDirectoriesRequest {
	s.DirectoryId = v
	return s
}

func (s *DeleteDirectoriesRequest) SetRegionId(v string) *DeleteDirectoriesRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteDirectoriesRequest) Validate() error {
	return dara.Validate(s)
}
