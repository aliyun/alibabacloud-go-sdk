// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDirectoriesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDirectoryId(v []*string) *DescribeDirectoriesRequest
	GetDirectoryId() []*string
	SetDirectoryStatus(v string) *DescribeDirectoriesRequest
	GetDirectoryStatus() *string
	SetDirectoryType(v string) *DescribeDirectoriesRequest
	GetDirectoryType() *string
	SetMaxResults(v int32) *DescribeDirectoriesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeDirectoriesRequest
	GetNextToken() *string
	SetRegionId(v string) *DescribeDirectoriesRequest
	GetRegionId() *string
	SetStatus(v string) *DescribeDirectoriesRequest
	GetStatus() *string
}

type DescribeDirectoriesRequest struct {
	// Details of directory IDs. You can specify one or more directory IDs.
	//
	// example:
	//
	// cn-hangzhou+dir-gx2x1dhsmu52rd****
	DirectoryId []*string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty" type:"Repeated"`
	// The directory status. This parameter is equivalent to `Status`.
	//
	// example:
	//
	// REGISTERED
	DirectoryStatus *string `json:"DirectoryStatus,omitempty" xml:"DirectoryStatus,omitempty"`
	// The directory type.
	//
	// Valid value:
	//
	// 	- SIMPLE: the convenience directory.
	//
	// 	- AD_CONNECTOR: the Active Directory (AD) directory.
	//
	// example:
	//
	// RAM
	DirectoryType *string `json:"DirectoryType,omitempty" xml:"DirectoryType,omitempty"`
	// The number of entries to return on each page.
	//
	// Maximum value: 100.
	//
	// Default value: 10.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that determines the start point of the next query. If this parameter is empty, all results are returned.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a4883
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/196646.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The directory status.
	//
	// Valid values:
	//
	// 	- REGISTERING: The directory is being registered.
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- DEREGISTERING: The directory is being deregistered.
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- REGISTERED: The directory is registered.
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- NEEDCONFIGTRUST: A trust relationship needs to be configured for the directory.
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- CONFIGTRUSTFAILED: A trust relationship fails to be configured for the directory.
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- DEREGISTERED: The directory is deregistered.
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- ERROR: One or more configurations of the directory are invalid.
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- CONFIGTRUSTING: A trust relationship is being configured.
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- NEEDCONFIGUSER: Users need to be configured for the directory.
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// REGISTERED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeDirectoriesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDirectoriesRequest) GoString() string {
	return s.String()
}

func (s *DescribeDirectoriesRequest) GetDirectoryId() []*string {
	return s.DirectoryId
}

func (s *DescribeDirectoriesRequest) GetDirectoryStatus() *string {
	return s.DirectoryStatus
}

func (s *DescribeDirectoriesRequest) GetDirectoryType() *string {
	return s.DirectoryType
}

func (s *DescribeDirectoriesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeDirectoriesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeDirectoriesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDirectoriesRequest) GetStatus() *string {
	return s.Status
}

func (s *DescribeDirectoriesRequest) SetDirectoryId(v []*string) *DescribeDirectoriesRequest {
	s.DirectoryId = v
	return s
}

func (s *DescribeDirectoriesRequest) SetDirectoryStatus(v string) *DescribeDirectoriesRequest {
	s.DirectoryStatus = &v
	return s
}

func (s *DescribeDirectoriesRequest) SetDirectoryType(v string) *DescribeDirectoriesRequest {
	s.DirectoryType = &v
	return s
}

func (s *DescribeDirectoriesRequest) SetMaxResults(v int32) *DescribeDirectoriesRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeDirectoriesRequest) SetNextToken(v string) *DescribeDirectoriesRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeDirectoriesRequest) SetRegionId(v string) *DescribeDirectoriesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDirectoriesRequest) SetStatus(v string) *DescribeDirectoriesRequest {
	s.Status = &v
	return s
}

func (s *DescribeDirectoriesRequest) Validate() error {
	return dara.Validate(s)
}
