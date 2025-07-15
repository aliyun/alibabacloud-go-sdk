// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCdsFilesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCdsId(v string) *ListCdsFilesShrinkRequest
	GetCdsId() *string
	SetEndUserId(v string) *ListCdsFilesShrinkRequest
	GetEndUserId() *string
	SetFileIdsShrink(v string) *ListCdsFilesShrinkRequest
	GetFileIdsShrink() *string
	SetGroupId(v string) *ListCdsFilesShrinkRequest
	GetGroupId() *string
	SetMaxResults(v int32) *ListCdsFilesShrinkRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListCdsFilesShrinkRequest
	GetNextToken() *string
	SetOrderType(v string) *ListCdsFilesShrinkRequest
	GetOrderType() *string
	SetParentFileId(v string) *ListCdsFilesShrinkRequest
	GetParentFileId() *string
	SetRegionId(v string) *ListCdsFilesShrinkRequest
	GetRegionId() *string
	SetStatus(v string) *ListCdsFilesShrinkRequest
	GetStatus() *string
}

type ListCdsFilesShrinkRequest struct {
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
	FileIdsShrink *string `json:"FileIds,omitempty" xml:"FileIds,omitempty"`
	GroupId       *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
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

func (s ListCdsFilesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCdsFilesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListCdsFilesShrinkRequest) GetCdsId() *string {
	return s.CdsId
}

func (s *ListCdsFilesShrinkRequest) GetEndUserId() *string {
	return s.EndUserId
}

func (s *ListCdsFilesShrinkRequest) GetFileIdsShrink() *string {
	return s.FileIdsShrink
}

func (s *ListCdsFilesShrinkRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *ListCdsFilesShrinkRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListCdsFilesShrinkRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListCdsFilesShrinkRequest) GetOrderType() *string {
	return s.OrderType
}

func (s *ListCdsFilesShrinkRequest) GetParentFileId() *string {
	return s.ParentFileId
}

func (s *ListCdsFilesShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListCdsFilesShrinkRequest) GetStatus() *string {
	return s.Status
}

func (s *ListCdsFilesShrinkRequest) SetCdsId(v string) *ListCdsFilesShrinkRequest {
	s.CdsId = &v
	return s
}

func (s *ListCdsFilesShrinkRequest) SetEndUserId(v string) *ListCdsFilesShrinkRequest {
	s.EndUserId = &v
	return s
}

func (s *ListCdsFilesShrinkRequest) SetFileIdsShrink(v string) *ListCdsFilesShrinkRequest {
	s.FileIdsShrink = &v
	return s
}

func (s *ListCdsFilesShrinkRequest) SetGroupId(v string) *ListCdsFilesShrinkRequest {
	s.GroupId = &v
	return s
}

func (s *ListCdsFilesShrinkRequest) SetMaxResults(v int32) *ListCdsFilesShrinkRequest {
	s.MaxResults = &v
	return s
}

func (s *ListCdsFilesShrinkRequest) SetNextToken(v string) *ListCdsFilesShrinkRequest {
	s.NextToken = &v
	return s
}

func (s *ListCdsFilesShrinkRequest) SetOrderType(v string) *ListCdsFilesShrinkRequest {
	s.OrderType = &v
	return s
}

func (s *ListCdsFilesShrinkRequest) SetParentFileId(v string) *ListCdsFilesShrinkRequest {
	s.ParentFileId = &v
	return s
}

func (s *ListCdsFilesShrinkRequest) SetRegionId(v string) *ListCdsFilesShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *ListCdsFilesShrinkRequest) SetStatus(v string) *ListCdsFilesShrinkRequest {
	s.Status = &v
	return s
}

func (s *ListCdsFilesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
