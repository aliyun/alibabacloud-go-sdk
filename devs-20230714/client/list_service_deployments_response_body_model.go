// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListServiceDeploymentsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*ServiceDeployment) *ListServiceDeploymentsResponseBody
	GetData() []*ServiceDeployment
	SetPageNumber(v int64) *ListServiceDeploymentsResponseBody
	GetPageNumber() *int64
	SetPageSize(v int64) *ListServiceDeploymentsResponseBody
	GetPageSize() *int64
	SetTotalCount(v int64) *ListServiceDeploymentsResponseBody
	GetTotalCount() *int64
}

type ListServiceDeploymentsResponseBody struct {
	Data []*ServiceDeployment `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// 50
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListServiceDeploymentsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListServiceDeploymentsResponseBody) GoString() string {
	return s.String()
}

func (s *ListServiceDeploymentsResponseBody) GetData() []*ServiceDeployment {
	return s.Data
}

func (s *ListServiceDeploymentsResponseBody) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListServiceDeploymentsResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListServiceDeploymentsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListServiceDeploymentsResponseBody) SetData(v []*ServiceDeployment) *ListServiceDeploymentsResponseBody {
	s.Data = v
	return s
}

func (s *ListServiceDeploymentsResponseBody) SetPageNumber(v int64) *ListServiceDeploymentsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListServiceDeploymentsResponseBody) SetPageSize(v int64) *ListServiceDeploymentsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListServiceDeploymentsResponseBody) SetTotalCount(v int64) *ListServiceDeploymentsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListServiceDeploymentsResponseBody) Validate() error {
	return dara.Validate(s)
}
