// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListServiceDeploymentsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKeyword(v string) *ListServiceDeploymentsRequest
	GetKeyword() *string
	SetLabelSelector(v []*string) *ListServiceDeploymentsRequest
	GetLabelSelector() []*string
	SetPageNumber(v int64) *ListServiceDeploymentsRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *ListServiceDeploymentsRequest
	GetPageSize() *int64
}

type ListServiceDeploymentsRequest struct {
	// example:
	//
	// demo
	Keyword       *string   `json:"keyword,omitempty" xml:"keyword,omitempty"`
	LabelSelector []*string `json:"labelSelector,omitempty" xml:"labelSelector,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s ListServiceDeploymentsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListServiceDeploymentsRequest) GoString() string {
	return s.String()
}

func (s *ListServiceDeploymentsRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListServiceDeploymentsRequest) GetLabelSelector() []*string {
	return s.LabelSelector
}

func (s *ListServiceDeploymentsRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListServiceDeploymentsRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListServiceDeploymentsRequest) SetKeyword(v string) *ListServiceDeploymentsRequest {
	s.Keyword = &v
	return s
}

func (s *ListServiceDeploymentsRequest) SetLabelSelector(v []*string) *ListServiceDeploymentsRequest {
	s.LabelSelector = v
	return s
}

func (s *ListServiceDeploymentsRequest) SetPageNumber(v int64) *ListServiceDeploymentsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListServiceDeploymentsRequest) SetPageSize(v int64) *ListServiceDeploymentsRequest {
	s.PageSize = &v
	return s
}

func (s *ListServiceDeploymentsRequest) Validate() error {
	return dara.Validate(s)
}
