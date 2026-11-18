// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMediaFoldersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNo(v int32) *ListMediaFoldersRequest
	GetPageNo() *int32
	SetPageSize(v int32) *ListMediaFoldersRequest
	GetPageSize() *int32
	SetProductionId(v string) *ListMediaFoldersRequest
	GetProductionId() *string
}

type ListMediaFoldersRequest struct {
	// The page number. Pages start from 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page. Default value: 20. Maximum value: 50. If the value exceeds 50, it is truncated to 50.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The project ID. If this parameter is not specified, the default project of the current account is used. The value is in the `pd_{uid}` format.
	//
	// example:
	//
	// pd_1585***
	ProductionId *string `json:"ProductionId,omitempty" xml:"ProductionId,omitempty"`
}

func (s ListMediaFoldersRequest) String() string {
	return dara.Prettify(s)
}

func (s ListMediaFoldersRequest) GoString() string {
	return s.String()
}

func (s *ListMediaFoldersRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *ListMediaFoldersRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListMediaFoldersRequest) GetProductionId() *string {
	return s.ProductionId
}

func (s *ListMediaFoldersRequest) SetPageNo(v int32) *ListMediaFoldersRequest {
	s.PageNo = &v
	return s
}

func (s *ListMediaFoldersRequest) SetPageSize(v int32) *ListMediaFoldersRequest {
	s.PageSize = &v
	return s
}

func (s *ListMediaFoldersRequest) SetProductionId(v string) *ListMediaFoldersRequest {
	s.ProductionId = &v
	return s
}

func (s *ListMediaFoldersRequest) Validate() error {
	return dara.Validate(s)
}
