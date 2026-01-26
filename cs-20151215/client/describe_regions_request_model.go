// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRegionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *DescribeRegionsRequest
	GetAcceptLanguage() *string
	SetClusterType(v string) *DescribeRegionsRequest
	GetClusterType() *string
	SetProfile(v string) *DescribeRegionsRequest
	GetProfile() *string
}

type DescribeRegionsRequest struct {
	// example:
	//
	// zh-CN
	AcceptLanguage *string `json:"acceptLanguage,omitempty" xml:"acceptLanguage,omitempty"`
	// example:
	//
	// ManagedKubernetes
	ClusterType *string `json:"clusterType,omitempty" xml:"clusterType,omitempty"`
	// example:
	//
	// Default
	Profile *string `json:"profile,omitempty" xml:"profile,omitempty"`
}

func (s DescribeRegionsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRegionsRequest) GoString() string {
	return s.String()
}

func (s *DescribeRegionsRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *DescribeRegionsRequest) GetClusterType() *string {
	return s.ClusterType
}

func (s *DescribeRegionsRequest) GetProfile() *string {
	return s.Profile
}

func (s *DescribeRegionsRequest) SetAcceptLanguage(v string) *DescribeRegionsRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *DescribeRegionsRequest) SetClusterType(v string) *DescribeRegionsRequest {
	s.ClusterType = &v
	return s
}

func (s *DescribeRegionsRequest) SetProfile(v string) *DescribeRegionsRequest {
	s.Profile = &v
	return s
}

func (s *DescribeRegionsRequest) Validate() error {
	return dara.Validate(s)
}
