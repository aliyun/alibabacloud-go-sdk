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
	// The current page number.
	//
	// example:
	//
	// 1
	Current *string `json:"Current,omitempty" xml:"Current,omitempty"`
	// The severity level. Separate multiple values with commas (,). Valid values:
	//
	// - serious: urgent
	//
	// - suspicious: suspicious
	//
	// - remind: reminder.
	//
	// example:
	//
	// serious,suspicious,remind
	Levels *string `json:"Levels,omitempty" xml:"Levels,omitempty"`
	// The maximum number of entries per page in a paginated query. Default value: 20. If this parameter is left empty, 20 entries are returned.
	//
	// >Do not leave PageSize empty.
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
