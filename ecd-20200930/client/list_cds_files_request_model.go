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
	FileIds []*string `json:"FileIds,omitempty" xml:"FileIds,omitempty" type:"Repeated"`
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
