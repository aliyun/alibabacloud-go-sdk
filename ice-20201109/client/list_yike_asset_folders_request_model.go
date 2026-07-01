// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListYikeAssetFoldersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNo(v int32) *ListYikeAssetFoldersRequest
	GetPageNo() *int32
	SetPageSize(v int32) *ListYikeAssetFoldersRequest
	GetPageSize() *int32
	SetProductionId(v string) *ListYikeAssetFoldersRequest
	GetProductionId() *string
}

type ListYikeAssetFoldersRequest struct {
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page. Default value: 10. Maximum value: 50.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The project ID. If this parameter is omitted, folders in the default project are returned.
	//
	// example:
	//
	// ProductionId
	ProductionId *string `json:"ProductionId,omitempty" xml:"ProductionId,omitempty"`
}

func (s ListYikeAssetFoldersRequest) String() string {
	return dara.Prettify(s)
}

func (s ListYikeAssetFoldersRequest) GoString() string {
	return s.String()
}

func (s *ListYikeAssetFoldersRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *ListYikeAssetFoldersRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListYikeAssetFoldersRequest) GetProductionId() *string {
	return s.ProductionId
}

func (s *ListYikeAssetFoldersRequest) SetPageNo(v int32) *ListYikeAssetFoldersRequest {
	s.PageNo = &v
	return s
}

func (s *ListYikeAssetFoldersRequest) SetPageSize(v int32) *ListYikeAssetFoldersRequest {
	s.PageSize = &v
	return s
}

func (s *ListYikeAssetFoldersRequest) SetProductionId(v string) *ListYikeAssetFoldersRequest {
	s.ProductionId = &v
	return s
}

func (s *ListYikeAssetFoldersRequest) Validate() error {
	return dara.Validate(s)
}
