// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAffectedAssetsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrent(v string) *DescribeAffectedAssetsRequest
	GetCurrent() *string
	SetLevels(v string) *DescribeAffectedAssetsRequest
	GetLevels() *string
	SetPageSize(v string) *DescribeAffectedAssetsRequest
	GetPageSize() *string
}

type DescribeAffectedAssetsRequest struct {
	// The number of the page to return.
	//
	// example:
	//
	// 1
	Current *string `json:"Current,omitempty" xml:"Current,omitempty"`
	// The severity. Separate multiple severities with commas (,). Valid values:
	//
	// 	- serious
	//
	// 	- suspicious
	//
	// 	- remind
	//
	// example:
	//
	// serious,suspicious,remind
	Levels *string `json:"Levels,omitempty" xml:"Levels,omitempty"`
	// The number of entries to return on each page. Default value: 20. If you leave this parameter empty, 20 entries are returned on each page.
	//
	// > We recommend that you do not leave this parameter empty.
	//
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeAffectedAssetsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAffectedAssetsRequest) GoString() string {
	return s.String()
}

func (s *DescribeAffectedAssetsRequest) GetCurrent() *string {
	return s.Current
}

func (s *DescribeAffectedAssetsRequest) GetLevels() *string {
	return s.Levels
}

func (s *DescribeAffectedAssetsRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribeAffectedAssetsRequest) SetCurrent(v string) *DescribeAffectedAssetsRequest {
	s.Current = &v
	return s
}

func (s *DescribeAffectedAssetsRequest) SetLevels(v string) *DescribeAffectedAssetsRequest {
	s.Levels = &v
	return s
}

func (s *DescribeAffectedAssetsRequest) SetPageSize(v string) *DescribeAffectedAssetsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeAffectedAssetsRequest) Validate() error {
	return dara.Validate(s)
}
