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
	SetEnsRegionId(v string) *DescribeClustersV1Request
	GetEnsRegionId() *string
	SetName(v string) *DescribeClustersV1Request
	GetName() *string
}

type DescribeClustersV1Request struct {
	// example:
	//
	// eck-xxxxxxxx
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// example:
	//
	// cn-hangzhou-58
	EnsRegionId *string `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty"`
	// example:
	//
	// test-eck-name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
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

func (s *DescribeClustersV1Request) GetEnsRegionId() *string {
	return s.EnsRegionId
}

func (s *DescribeClustersV1Request) GetName() *string {
	return s.Name
}

func (s *DescribeClustersV1Request) SetClusterId(v string) *DescribeClustersV1Request {
	s.ClusterId = &v
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

func (s *DescribeClustersV1Request) Validate() error {
	return dara.Validate(s)
}
