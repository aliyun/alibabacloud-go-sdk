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
	// The enterprise cloud drive ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou+cds-320357****
	CdsId *string `json:"CdsId,omitempty" xml:"CdsId,omitempty"`
	// The ID of the user to whom the cloud drive is assigned.
	//
	// example:
	//
	// alice
	EndUserId *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	// The list of file IDs to query.
	FileIdsShrink *string `json:"FileIds,omitempty" xml:"FileIds,omitempty"`
	// The team space ID.
	//
	// example:
	//
	// cg-i1ruuudp92qpj****
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The maximum number of entries per page in a paging query. Default value: 100.
	//
	// example:
	//
	// 100
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token. Set this parameter to the NextToken value returned in the previous call. You do not need to set this parameter for the first request.
	//
	// example:
	//
	// aGN4YzAxQGNuLWhhbmd6aG91LjExNzU5NTMyNjgzMTQ1****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The sort order of the file list.
	//
	// example:
	//
	// CreateTimeDesc
	OrderType *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	// The parent file ID. You can obtain this value from the FileId response parameter of this operation.
	//
	// example:
	//
	// 63636837e47e5a24a8a940218bef395c210e****
	ParentFileId *string `json:"ParentFileId,omitempty" xml:"ParentFileId,omitempty"`
	// The region ID. You can call [DescribeRegions](https://help.aliyun.com/document_detail/196646.html) to query the regions supported by Elastic Desktop Service.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The file status.
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
