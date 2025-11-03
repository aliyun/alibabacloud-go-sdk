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
	// The ID of the enterprise drive.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou+cds-320357****
	CdsId *string `json:"CdsId,omitempty" xml:"CdsId,omitempty"`
	// The ID of the user to which the network disk is assigned.
	//
	// example:
	//
	// testUser
	EndUserId *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	// The IDs of the files to be queried.
	FileIdsShrink *string `json:"FileIds,omitempty" xml:"FileIds,omitempty"`
	// The ID of the team space.
	//
	// example:
	//
	// cg-i1ruuudp92qpj****
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The number of entries to return on each page. Default value: 100.
	//
	// example:
	//
	// 100
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The query token. Set the value to the value of the `NextToken` parameter returned in the last call to the operation. You do not need to set this parameter when you call the operation for the first time.
	//
	// example:
	//
	// WyI2Mzg4MjAwMzFhNGQwZWVmN2I3MjRkZjZhZjAyMWU4YzY1MmRjZmUyIiwibiIsIm4iLDEsLTEsMTY2OTg2NTQ3NTMxMiwiNjM4ODIwMDNlNTU0YmZiZjFkYTk0MmEyYTZhMjEyZDkxODdjMjAy****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The sorting method of the files.
	//
	// Valid values:
	//
	// 	- CreateTimeDesc: sorts the by creation time in descending order.
	//
	// 	- ModifiedTimeAsc: sort the by modification time in ascending order.
	//
	// 	- NameDesc: sorts the by file name in descending order.
	//
	// 	- SizeAsc: sorts by file size in ascending order.
	//
	// 	- ModifiedTimeDesc: sort the by modification time in descending order.
	//
	// 	- CreateTimeAsc: sorts the by creation time in ascending order.
	//
	// 	- SizeDesc: sorts by file size in descending order.
	//
	// 	- NameAsc: sorts by file name in ascending order.
	//
	// example:
	//
	// CreateTimeDesc
	OrderType *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	// The parent folder ID. You can obtain the value by using the response parameter `FileId` of this operation.
	//
	// example:
	//
	// 63636837e47e5a24a8a940218bef395c210e****
	ParentFileId *string `json:"ParentFileId,omitempty" xml:"ParentFileId,omitempty"`
	// The ID of the logon region. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/196646.html) operation to obtain the list of regions supported by cloud computers.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The file status.
	//
	// Valid values:
	//
	// 	- available: returns only normal file.
	//
	// 	- uploading: returns only the of objects that are being uploaded.
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
