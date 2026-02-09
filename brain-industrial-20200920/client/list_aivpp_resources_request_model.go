// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAivppResourcesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *ListAivppResourcesRequest
	GetCurrentPage() *int32
	SetInstanceType(v string) *ListAivppResourcesRequest
	GetInstanceType() *string
	SetMaxResults(v int32) *ListAivppResourcesRequest
	GetMaxResults() *int32
}

type ListAivppResourcesRequest struct {
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// example:
	//
	// DATA
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
}

func (s ListAivppResourcesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAivppResourcesRequest) GoString() string {
	return s.String()
}

func (s *ListAivppResourcesRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListAivppResourcesRequest) GetInstanceType() *string {
	return s.InstanceType
}

func (s *ListAivppResourcesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListAivppResourcesRequest) SetCurrentPage(v int32) *ListAivppResourcesRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListAivppResourcesRequest) SetInstanceType(v string) *ListAivppResourcesRequest {
	s.InstanceType = &v
	return s
}

func (s *ListAivppResourcesRequest) SetMaxResults(v int32) *ListAivppResourcesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListAivppResourcesRequest) Validate() error {
	return dara.Validate(s)
}
