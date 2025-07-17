// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataLakeCatalogRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDataRegion(v string) *ListDataLakeCatalogRequest
	GetDataRegion() *string
	SetSearchKey(v string) *ListDataLakeCatalogRequest
	GetSearchKey() *string
	SetTid(v int64) *ListDataLakeCatalogRequest
	GetTid() *int64
	SetWorkspaceId(v int64) *ListDataLakeCatalogRequest
	GetWorkspaceId() *int64
}

type ListDataLakeCatalogRequest struct {
	// This parameter is required.
	//
	// if can be null:
	// false
	//
	// example:
	//
	// cn-hangzhou
	DataRegion *string `json:"DataRegion,omitempty" xml:"DataRegion,omitempty"`
	// if can be null:
	// false
	//
	// example:
	//
	// hive
	SearchKey *string `json:"SearchKey,omitempty" xml:"SearchKey,omitempty"`
	// example:
	//
	// 3
	Tid         *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
	WorkspaceId *int64 `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListDataLakeCatalogRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDataLakeCatalogRequest) GoString() string {
	return s.String()
}

func (s *ListDataLakeCatalogRequest) GetDataRegion() *string {
	return s.DataRegion
}

func (s *ListDataLakeCatalogRequest) GetSearchKey() *string {
	return s.SearchKey
}

func (s *ListDataLakeCatalogRequest) GetTid() *int64 {
	return s.Tid
}

func (s *ListDataLakeCatalogRequest) GetWorkspaceId() *int64 {
	return s.WorkspaceId
}

func (s *ListDataLakeCatalogRequest) SetDataRegion(v string) *ListDataLakeCatalogRequest {
	s.DataRegion = &v
	return s
}

func (s *ListDataLakeCatalogRequest) SetSearchKey(v string) *ListDataLakeCatalogRequest {
	s.SearchKey = &v
	return s
}

func (s *ListDataLakeCatalogRequest) SetTid(v int64) *ListDataLakeCatalogRequest {
	s.Tid = &v
	return s
}

func (s *ListDataLakeCatalogRequest) SetWorkspaceId(v int64) *ListDataLakeCatalogRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ListDataLakeCatalogRequest) Validate() error {
	return dara.Validate(s)
}
