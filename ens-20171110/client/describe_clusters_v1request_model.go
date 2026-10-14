// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeClustersV1Request interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *DescribeClustersV1Request
	GetClusterId() *string
	SetClusterType(v string) *DescribeClustersV1Request
	GetClusterType() *string
	SetEnsRegionId(v string) *DescribeClustersV1Request
	GetEnsRegionId() *string
	SetName(v string) *DescribeClustersV1Request
	GetName() *string
	SetProfile(v string) *DescribeClustersV1Request
	GetProfile() *string
}

type DescribeClustersV1Request struct {
	// The cluster ID.
	//
	// example:
	//
	// eck-xxxxxxxx
	ClusterId   *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	ClusterType *string `json:"ClusterType,omitempty" xml:"ClusterType,omitempty"`
	// The ENS region ID where the cluster resides.
	//
	// example:
	//
	// cn-hangzhou-58
	EnsRegionId *string `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty"`
	// The cluster name.
	//
	// example:
	//
	// test-eck-name
	Name    *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Profile *string `json:"Profile,omitempty" xml:"Profile,omitempty"`
}

func (s DescribeClustersV1Request) String() string {
	return dara.Prettify(s)
}

func (s DescribeClustersV1Request) GoString() string {
	return s.String()
}

func (s *DescribeClustersV1Request) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeClustersV1Request) GetClusterType() *string {
	return s.ClusterType
}

func (s *DescribeClustersV1Request) GetEnsRegionId() *string {
	return s.EnsRegionId
}

func (s *DescribeClustersV1Request) GetName() *string {
	return s.Name
}

func (s *DescribeClustersV1Request) GetProfile() *string {
	return s.Profile
}

func (s *DescribeClustersV1Request) SetClusterId(v string) *DescribeClustersV1Request {
	s.ClusterId = &v
	return s
}

func (s *DescribeClustersV1Request) SetClusterType(v string) *DescribeClustersV1Request {
	s.ClusterType = &v
	return s
}

func (s *DescribeClustersV1Request) SetEnsRegionId(v string) *DescribeClustersV1Request {
	s.EnsRegionId = &v
	return s
}

func (s *DescribeClustersV1Request) SetName(v string) *DescribeClustersV1Request {
	s.Name = &v
	return s
}

func (s *DescribeClustersV1Request) SetProfile(v string) *DescribeClustersV1Request {
	s.Profile = &v
	return s
}

func (s *DescribeClustersV1Request) Validate() error {
	return dara.Validate(s)
}
