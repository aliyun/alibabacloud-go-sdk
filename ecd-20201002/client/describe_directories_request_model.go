// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDirectoriesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientId(v string) *DescribeDirectoriesRequest
	GetClientId() *string
	SetDirectoryId(v []*string) *DescribeDirectoriesRequest
	GetDirectoryId() []*string
	SetRegionId(v string) *DescribeDirectoriesRequest
	GetRegionId() *string
}

type DescribeDirectoriesRequest struct {
	// The client ID. The system generates a unique ID for each client.
	//
	// This parameter is required.
	//
	// example:
	//
	// 54c17e1d-2d72-4b87-aa33-25f3b3f2****
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// The directory IDs.
	DirectoryId []*string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty" type:"Repeated"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeDirectoriesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDirectoriesRequest) GoString() string {
	return s.String()
}

func (s *DescribeDirectoriesRequest) GetClientId() *string {
	return s.ClientId
}

func (s *DescribeDirectoriesRequest) GetDirectoryId() []*string {
	return s.DirectoryId
}

func (s *DescribeDirectoriesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDirectoriesRequest) SetClientId(v string) *DescribeDirectoriesRequest {
	s.ClientId = &v
	return s
}

func (s *DescribeDirectoriesRequest) SetDirectoryId(v []*string) *DescribeDirectoriesRequest {
	s.DirectoryId = v
	return s
}

func (s *DescribeDirectoriesRequest) SetRegionId(v string) *DescribeDirectoriesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDirectoriesRequest) Validate() error {
	return dara.Validate(s)
}
