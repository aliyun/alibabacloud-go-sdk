// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMountPointsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileSystemId(v string) *ListMountPointsRequest
	GetFileSystemId() *string
	SetInputRegionId(v string) *ListMountPointsRequest
	GetInputRegionId() *string
	SetLimit(v int32) *ListMountPointsRequest
	GetLimit() *int32
	SetNextToken(v string) *ListMountPointsRequest
	GetNextToken() *string
	SetOrderBy(v string) *ListMountPointsRequest
	GetOrderBy() *string
	SetOrderType(v string) *ListMountPointsRequest
	GetOrderType() *string
	SetStartOffset(v int32) *ListMountPointsRequest
	GetStartOffset() *int32
}

type ListMountPointsRequest struct {
	// This parameter is required.
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	InputRegionId *string `json:"InputRegionId,omitempty" xml:"InputRegionId,omitempty"`
	// example:
	//
	// 10
	Limit     *int32  `json:"Limit,omitempty" xml:"Limit,omitempty"`
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// CreateTime
	OrderBy *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	// example:
	//
	// ASC
	OrderType *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	// example:
	//
	// 10
	StartOffset *int32 `json:"StartOffset,omitempty" xml:"StartOffset,omitempty"`
}

func (s ListMountPointsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListMountPointsRequest) GoString() string {
	return s.String()
}

func (s *ListMountPointsRequest) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *ListMountPointsRequest) GetInputRegionId() *string {
	return s.InputRegionId
}

func (s *ListMountPointsRequest) GetLimit() *int32 {
	return s.Limit
}

func (s *ListMountPointsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListMountPointsRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *ListMountPointsRequest) GetOrderType() *string {
	return s.OrderType
}

func (s *ListMountPointsRequest) GetStartOffset() *int32 {
	return s.StartOffset
}

func (s *ListMountPointsRequest) SetFileSystemId(v string) *ListMountPointsRequest {
	s.FileSystemId = &v
	return s
}

func (s *ListMountPointsRequest) SetInputRegionId(v string) *ListMountPointsRequest {
	s.InputRegionId = &v
	return s
}

func (s *ListMountPointsRequest) SetLimit(v int32) *ListMountPointsRequest {
	s.Limit = &v
	return s
}

func (s *ListMountPointsRequest) SetNextToken(v string) *ListMountPointsRequest {
	s.NextToken = &v
	return s
}

func (s *ListMountPointsRequest) SetOrderBy(v string) *ListMountPointsRequest {
	s.OrderBy = &v
	return s
}

func (s *ListMountPointsRequest) SetOrderType(v string) *ListMountPointsRequest {
	s.OrderType = &v
	return s
}

func (s *ListMountPointsRequest) SetStartOffset(v int32) *ListMountPointsRequest {
	s.StartOffset = &v
	return s
}

func (s *ListMountPointsRequest) Validate() error {
	return dara.Validate(s)
}
