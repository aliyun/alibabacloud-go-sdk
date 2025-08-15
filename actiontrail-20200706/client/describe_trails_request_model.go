// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTrailsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIncludeOrganizationTrail(v bool) *DescribeTrailsRequest
	GetIncludeOrganizationTrail() *bool
	SetIncludeShadowTrails(v bool) *DescribeTrailsRequest
	GetIncludeShadowTrails() *bool
	SetNameList(v string) *DescribeTrailsRequest
	GetNameList() *string
}

type DescribeTrailsRequest struct {
	// Specifies whether to query the information about multi-account trails. Valid values:
	//
	// 	- true
	//
	// 	- false (default)
	//
	// example:
	//
	// false
	IncludeOrganizationTrail *bool `json:"IncludeOrganizationTrail,omitempty" xml:"IncludeOrganizationTrail,omitempty"`
	// Specifies whether to return the information about shadow trails. Valid values:
	//
	// 	- false: Do not return the information about shadow trails. It is the default value.
	//
	// 	- true: Return the information about shadow trails.
	//
	// example:
	//
	// false
	IncludeShadowTrails *bool `json:"IncludeShadowTrails,omitempty" xml:"IncludeShadowTrails,omitempty"`
	// The names of the trails whose information you want to query. Separate multiple trail names with commas (,).
	//
	// example:
	//
	// abc,def
	NameList *string `json:"NameList,omitempty" xml:"NameList,omitempty"`
}

func (s DescribeTrailsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeTrailsRequest) GoString() string {
	return s.String()
}

func (s *DescribeTrailsRequest) GetIncludeOrganizationTrail() *bool {
	return s.IncludeOrganizationTrail
}

func (s *DescribeTrailsRequest) GetIncludeShadowTrails() *bool {
	return s.IncludeShadowTrails
}

func (s *DescribeTrailsRequest) GetNameList() *string {
	return s.NameList
}

func (s *DescribeTrailsRequest) SetIncludeOrganizationTrail(v bool) *DescribeTrailsRequest {
	s.IncludeOrganizationTrail = &v
	return s
}

func (s *DescribeTrailsRequest) SetIncludeShadowTrails(v bool) *DescribeTrailsRequest {
	s.IncludeShadowTrails = &v
	return s
}

func (s *DescribeTrailsRequest) SetNameList(v string) *DescribeTrailsRequest {
	s.NameList = &v
	return s
}

func (s *DescribeTrailsRequest) Validate() error {
	return dara.Validate(s)
}
