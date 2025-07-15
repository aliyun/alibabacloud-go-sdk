// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCdsFilesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCdsId(v string) *ListCdsFilesRequest
	GetCdsId() *string
	SetEndUserId(v string) *ListCdsFilesRequest
	GetEndUserId() *string
	SetFileIds(v []*string) *ListCdsFilesRequest
	GetFileIds() []*string
	SetGroupId(v string) *ListCdsFilesRequest
	GetGroupId() *string
	SetMaxResults(v int32) *ListCdsFilesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListCdsFilesRequest
	GetNextToken() *string
	SetOrderType(v string) *ListCdsFilesRequest
	GetOrderType() *string
	SetParentFileId(v string) *ListCdsFilesRequest
	GetParentFileId() *string
	SetRegionId(v string) *ListCdsFilesRequest
	GetRegionId() *string
	SetStatus(v string) *ListCdsFilesRequest
	GetStatus() *string
}

type ListCdsFilesRequest struct {
	// The ID of the cloud disk.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou+cds-320357****
	CdsId *string `json:"CdsId,omitempty" xml:"CdsId,omitempty"`
	// The ID of the user to whom the cloud disk is allocated.
	//
	// example:
	//
	// testUser
	EndUserId *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	// The IDs of the files to be queried.
	FileIds []*string `json:"FileIds,omitempty" xml:"FileIds,omitempty" type:"Repeated"`
	GroupId *string   `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The number of entries to return on each page. Default value: 100.
	//
	// example:
	//
	// 100
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token used for the next query. If this parameter is empty, all results are returned.
	//
	// example:
	//
	// WyI2Mzg4MjAwMzFhNGQwZWVmN2I3MjRkZjZhZjAyMWU4YzY1MmRjZmUyIiwibiIsIm4iLDEsLTEsMTY2OTg2NTQ3NTMxMiwiNjM4ODIwMDNlNTU0YmZiZjFkYTk0MmEyYTZhMjEyZDkxODdjMjAy****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The sorting method of the files.
	//
	// Valid values:
	//
	// 	- CreateTimeDesc
	//
	//     <!-- -->
	//
	//     :
	//
	//     <!-- -->
	//
	//     sorts files in descending order based on the time when they are created.
	//
	//     <!-- -->
	//
	// 	- ModifiedTimeAsc
	//
	//     <!-- -->
	//
	//     :
	//
	//     <!-- -->
	//
	//     sorts files in ascending order based on the time when they are modified.
	//
	//     <!-- -->
	//
	// 	- NameDesc
	//
	//     <!-- -->
	//
	//     :
	//
	//     <!-- -->
	//
	//     sorts files in descending order based on their names.
	//
	//     <!-- -->
	//
	// 	- SizeAsc
	//
	//     <!-- -->
	//
	//     :
	//
	//     <!-- -->
	//
	//     sorts files in ascending order based on their sizes.
	//
	//     <!-- -->
	//
	// 	- ModifiedTimeDesc
	//
	//     <!-- -->
	//
	//     :
	//
	//     <!-- -->
	//
	//     sorts files in descending order based on the time when they are modified.
	//
	//     <!-- -->
	//
	// 	- CreateTimeAsc
	//
	//     <!-- -->
	//
	//     :
	//
	//     <!-- -->
	//
	//     sorts files in ascending order based on the time when they are created.
	//
	//     <!-- -->
	//
	// 	- SizeDesc
	//
	//     <!-- -->
	//
	//     :
	//
	//     <!-- -->
	//
	//     sorts files in descending order based on their sizes.
	//
	//     <!-- -->
	//
	// 	- NameAsc
	//
	//     <!-- -->
	//
	//     :
	//
	//     <!-- -->
	//
	//     sorts files in ascending order based on their names.
	//
	//     <!-- -->
	//
	// example:
	//
	// CreateTimeDesc
	OrderType *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	// The ID of the parent file.
	//
	// example:
	//
	// 63636837e47e5a24a8a940218bef395c210e****
	ParentFileId *string `json:"ParentFileId,omitempty" xml:"ParentFileId,omitempty"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/196646.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The file status.
	//
	// Valid values:
	//
	// 	- available
	//
	//     <!-- -->
	//
	//     :
	//
	//     <!-- -->
	//
	//     returns only normal files.
	//
	//     <!-- -->
	//
	// 	- uploading
	//
	//     <!-- -->
	//
	//     :
	//
	//     <!-- -->
	//
	//     returns only the files that are being uploaded.
	//
	//     <!-- -->
	//
	// example:
	//
	// available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListCdsFilesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCdsFilesRequest) GoString() string {
	return s.String()
}

func (s *ListCdsFilesRequest) GetCdsId() *string {
	return s.CdsId
}

func (s *ListCdsFilesRequest) GetEndUserId() *string {
	return s.EndUserId
}

func (s *ListCdsFilesRequest) GetFileIds() []*string {
	return s.FileIds
}

func (s *ListCdsFilesRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *ListCdsFilesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListCdsFilesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListCdsFilesRequest) GetOrderType() *string {
	return s.OrderType
}

func (s *ListCdsFilesRequest) GetParentFileId() *string {
	return s.ParentFileId
}

func (s *ListCdsFilesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListCdsFilesRequest) GetStatus() *string {
	return s.Status
}

func (s *ListCdsFilesRequest) SetCdsId(v string) *ListCdsFilesRequest {
	s.CdsId = &v
	return s
}

func (s *ListCdsFilesRequest) SetEndUserId(v string) *ListCdsFilesRequest {
	s.EndUserId = &v
	return s
}

func (s *ListCdsFilesRequest) SetFileIds(v []*string) *ListCdsFilesRequest {
	s.FileIds = v
	return s
}

func (s *ListCdsFilesRequest) SetGroupId(v string) *ListCdsFilesRequest {
	s.GroupId = &v
	return s
}

func (s *ListCdsFilesRequest) SetMaxResults(v int32) *ListCdsFilesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListCdsFilesRequest) SetNextToken(v string) *ListCdsFilesRequest {
	s.NextToken = &v
	return s
}

func (s *ListCdsFilesRequest) SetOrderType(v string) *ListCdsFilesRequest {
	s.OrderType = &v
	return s
}

func (s *ListCdsFilesRequest) SetParentFileId(v string) *ListCdsFilesRequest {
	s.ParentFileId = &v
	return s
}

func (s *ListCdsFilesRequest) SetRegionId(v string) *ListCdsFilesRequest {
	s.RegionId = &v
	return s
}

func (s *ListCdsFilesRequest) SetStatus(v string) *ListCdsFilesRequest {
	s.Status = &v
	return s
}

func (s *ListCdsFilesRequest) Validate() error {
	return dara.Validate(s)
}
