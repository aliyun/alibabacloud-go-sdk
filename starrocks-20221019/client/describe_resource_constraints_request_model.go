// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeResourceConstraintsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetArchitecture(v string) *DescribeResourceConstraintsRequest
	GetArchitecture() *string
	SetPackageType(v string) *DescribeResourceConstraintsRequest
	GetPackageType() *string
	SetRunMode(v string) *DescribeResourceConstraintsRequest
	GetRunMode() *string
}

type DescribeResourceConstraintsRequest struct {
	// example:
	//
	// onEcs
	Architecture *string `json:"Architecture,omitempty" xml:"Architecture,omitempty"`
	// example:
	//
	// trial
	PackageType *string `json:"PackageType,omitempty" xml:"PackageType,omitempty"`
	// example:
	//
	// shared_data
	RunMode *string `json:"RunMode,omitempty" xml:"RunMode,omitempty"`
}

func (s DescribeResourceConstraintsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeResourceConstraintsRequest) GoString() string {
	return s.String()
}

func (s *DescribeResourceConstraintsRequest) GetArchitecture() *string {
	return s.Architecture
}

func (s *DescribeResourceConstraintsRequest) GetPackageType() *string {
	return s.PackageType
}

func (s *DescribeResourceConstraintsRequest) GetRunMode() *string {
	return s.RunMode
}

func (s *DescribeResourceConstraintsRequest) SetArchitecture(v string) *DescribeResourceConstraintsRequest {
	s.Architecture = &v
	return s
}

func (s *DescribeResourceConstraintsRequest) SetPackageType(v string) *DescribeResourceConstraintsRequest {
	s.PackageType = &v
	return s
}

func (s *DescribeResourceConstraintsRequest) SetRunMode(v string) *DescribeResourceConstraintsRequest {
	s.RunMode = &v
	return s
}

func (s *DescribeResourceConstraintsRequest) Validate() error {
	return dara.Validate(s)
}
