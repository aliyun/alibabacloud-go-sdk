// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCmsInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *ListCmsInstancesRequest
	GetClusterId() *string
	SetRegionId(v string) *ListCmsInstancesRequest
	GetRegionId() *string
	SetTypeFilter(v string) *ListCmsInstancesRequest
	GetTypeFilter() *string
}

type ListCmsInstancesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// n9p9o9o3se
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The type of the cloud service integration. Valid values:
	//
	// 	- direct: self-monitoring
	//
	// 	- cms: Hybrid Cloud Monitoring
	//
	// example:
	//
	// direct
	TypeFilter *string `json:"TypeFilter,omitempty" xml:"TypeFilter,omitempty"`
}

func (s ListCmsInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCmsInstancesRequest) GoString() string {
	return s.String()
}

func (s *ListCmsInstancesRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *ListCmsInstancesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListCmsInstancesRequest) GetTypeFilter() *string {
	return s.TypeFilter
}

func (s *ListCmsInstancesRequest) SetClusterId(v string) *ListCmsInstancesRequest {
	s.ClusterId = &v
	return s
}

func (s *ListCmsInstancesRequest) SetRegionId(v string) *ListCmsInstancesRequest {
	s.RegionId = &v
	return s
}

func (s *ListCmsInstancesRequest) SetTypeFilter(v string) *ListCmsInstancesRequest {
	s.TypeFilter = &v
	return s
}

func (s *ListCmsInstancesRequest) Validate() error {
	return dara.Validate(s)
}
