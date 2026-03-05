// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRechargeBillsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListRechargeBillsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListRechargeBillsRequest
	GetPageSize() *int32
	SetSceneType(v string) *ListRechargeBillsRequest
	GetSceneType() *string
}

type ListRechargeBillsRequest struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// sales_pick
	SceneType *string `json:"SceneType,omitempty" xml:"SceneType,omitempty"`
}

func (s ListRechargeBillsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListRechargeBillsRequest) GoString() string {
	return s.String()
}

func (s *ListRechargeBillsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListRechargeBillsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListRechargeBillsRequest) GetSceneType() *string {
	return s.SceneType
}

func (s *ListRechargeBillsRequest) SetPageNumber(v int32) *ListRechargeBillsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListRechargeBillsRequest) SetPageSize(v int32) *ListRechargeBillsRequest {
	s.PageSize = &v
	return s
}

func (s *ListRechargeBillsRequest) SetSceneType(v string) *ListRechargeBillsRequest {
	s.SceneType = &v
	return s
}

func (s *ListRechargeBillsRequest) Validate() error {
	return dara.Validate(s)
}
