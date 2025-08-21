// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeARMServerInstancesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAICSpecsShrink(v string) *DescribeARMServerInstancesShrinkRequest
	GetAICSpecsShrink() *string
	SetDescribeAICInstances(v bool) *DescribeARMServerInstancesShrinkRequest
	GetDescribeAICInstances() *bool
	SetEnsRegionIdsShrink(v string) *DescribeARMServerInstancesShrinkRequest
	GetEnsRegionIdsShrink() *string
	SetMaxDate(v string) *DescribeARMServerInstancesShrinkRequest
	GetMaxDate() *string
	SetMinDate(v string) *DescribeARMServerInstancesShrinkRequest
	GetMinDate() *string
	SetName(v string) *DescribeARMServerInstancesShrinkRequest
	GetName() *string
	SetNamespace(v string) *DescribeARMServerInstancesShrinkRequest
	GetNamespace() *string
	SetOrderByParams(v string) *DescribeARMServerInstancesShrinkRequest
	GetOrderByParams() *string
	SetPageNumber(v int32) *DescribeARMServerInstancesShrinkRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeARMServerInstancesShrinkRequest
	GetPageSize() *int32
	SetServerIdsShrink(v string) *DescribeARMServerInstancesShrinkRequest
	GetServerIdsShrink() *string
	SetServerSpecsShrink(v string) *DescribeARMServerInstancesShrinkRequest
	GetServerSpecsShrink() *string
	SetStatesShrink(v string) *DescribeARMServerInstancesShrinkRequest
	GetStatesShrink() *string
}

type DescribeARMServerInstancesShrinkRequest struct {
	// The container specifications.
	AICSpecsShrink *string `json:"AICSpecs,omitempty" xml:"AICSpecs,omitempty"`
	// Spcifies whether the result contains the container information. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	DescribeAICInstances *bool `json:"DescribeAICInstances,omitempty" xml:"DescribeAICInstances,omitempty"`
	// The IDs of the Edge Node Service (ENS) nodes.
	EnsRegionIdsShrink *string `json:"EnsRegionIds,omitempty" xml:"EnsRegionIds,omitempty"`
	// The end of the time range to query. Specify the time in the 2006-01-02 format. By default, the time range to query is not restricted.
	//
	// example:
	//
	// 2006-01-03
	MaxDate *string `json:"MaxDate,omitempty" xml:"MaxDate,omitempty"`
	// The beginning of the time range to query. Specify the time in the 2006-01-02 format. By default, the time range to query is not restricted.
	//
	// example:
	//
	// 2006-01-02
	MinDate *string `json:"MinDate,omitempty" xml:"MinDate,omitempty"`
	// The name of the server.
	//
	// example:
	//
	// Server-Name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The namespace.
	//
	// example:
	//
	// default
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The sorting order of the results to return. Valid values: ServerIdSort, ServerNameSort, ExpireTimeSort, CreationTimeSort, and EnsRegionIdSort.
	//
	// asc: ascending order. desc: descending order.
	//
	// example:
	//
	// {"ServerIdSort":"desc","ServerNameSort":"desc","ExpireTimeSort":"asc","CreationTimeSort":"asc","EnsRegionIdSort":"asc"}
	OrderByParams *string `json:"OrderByParams,omitempty" xml:"OrderByParams,omitempty"`
	// The page number. Pages start from page **1**.
	//
	// Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Maximum value: **100**.
	//
	// Default value: **10**.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The IDs of the ARM servers.
	ServerIdsShrink *string `json:"ServerIds,omitempty" xml:"ServerIds,omitempty"`
	// The server specifications.
	ServerSpecsShrink *string `json:"ServerSpecs,omitempty" xml:"ServerSpecs,omitempty"`
	// The operation statuses.
	StatesShrink *string `json:"States,omitempty" xml:"States,omitempty"`
}

func (s DescribeARMServerInstancesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeARMServerInstancesShrinkRequest) GoString() string {
	return s.String()
}

func (s *DescribeARMServerInstancesShrinkRequest) GetAICSpecsShrink() *string {
	return s.AICSpecsShrink
}

func (s *DescribeARMServerInstancesShrinkRequest) GetDescribeAICInstances() *bool {
	return s.DescribeAICInstances
}

func (s *DescribeARMServerInstancesShrinkRequest) GetEnsRegionIdsShrink() *string {
	return s.EnsRegionIdsShrink
}

func (s *DescribeARMServerInstancesShrinkRequest) GetMaxDate() *string {
	return s.MaxDate
}

func (s *DescribeARMServerInstancesShrinkRequest) GetMinDate() *string {
	return s.MinDate
}

func (s *DescribeARMServerInstancesShrinkRequest) GetName() *string {
	return s.Name
}

func (s *DescribeARMServerInstancesShrinkRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *DescribeARMServerInstancesShrinkRequest) GetOrderByParams() *string {
	return s.OrderByParams
}

func (s *DescribeARMServerInstancesShrinkRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeARMServerInstancesShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeARMServerInstancesShrinkRequest) GetServerIdsShrink() *string {
	return s.ServerIdsShrink
}

func (s *DescribeARMServerInstancesShrinkRequest) GetServerSpecsShrink() *string {
	return s.ServerSpecsShrink
}

func (s *DescribeARMServerInstancesShrinkRequest) GetStatesShrink() *string {
	return s.StatesShrink
}

func (s *DescribeARMServerInstancesShrinkRequest) SetAICSpecsShrink(v string) *DescribeARMServerInstancesShrinkRequest {
	s.AICSpecsShrink = &v
	return s
}

func (s *DescribeARMServerInstancesShrinkRequest) SetDescribeAICInstances(v bool) *DescribeARMServerInstancesShrinkRequest {
	s.DescribeAICInstances = &v
	return s
}

func (s *DescribeARMServerInstancesShrinkRequest) SetEnsRegionIdsShrink(v string) *DescribeARMServerInstancesShrinkRequest {
	s.EnsRegionIdsShrink = &v
	return s
}

func (s *DescribeARMServerInstancesShrinkRequest) SetMaxDate(v string) *DescribeARMServerInstancesShrinkRequest {
	s.MaxDate = &v
	return s
}

func (s *DescribeARMServerInstancesShrinkRequest) SetMinDate(v string) *DescribeARMServerInstancesShrinkRequest {
	s.MinDate = &v
	return s
}

func (s *DescribeARMServerInstancesShrinkRequest) SetName(v string) *DescribeARMServerInstancesShrinkRequest {
	s.Name = &v
	return s
}

func (s *DescribeARMServerInstancesShrinkRequest) SetNamespace(v string) *DescribeARMServerInstancesShrinkRequest {
	s.Namespace = &v
	return s
}

func (s *DescribeARMServerInstancesShrinkRequest) SetOrderByParams(v string) *DescribeARMServerInstancesShrinkRequest {
	s.OrderByParams = &v
	return s
}

func (s *DescribeARMServerInstancesShrinkRequest) SetPageNumber(v int32) *DescribeARMServerInstancesShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeARMServerInstancesShrinkRequest) SetPageSize(v int32) *DescribeARMServerInstancesShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeARMServerInstancesShrinkRequest) SetServerIdsShrink(v string) *DescribeARMServerInstancesShrinkRequest {
	s.ServerIdsShrink = &v
	return s
}

func (s *DescribeARMServerInstancesShrinkRequest) SetServerSpecsShrink(v string) *DescribeARMServerInstancesShrinkRequest {
	s.ServerSpecsShrink = &v
	return s
}

func (s *DescribeARMServerInstancesShrinkRequest) SetStatesShrink(v string) *DescribeARMServerInstancesShrinkRequest {
	s.StatesShrink = &v
	return s
}

func (s *DescribeARMServerInstancesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
