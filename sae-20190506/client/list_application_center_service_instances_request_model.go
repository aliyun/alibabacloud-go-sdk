// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApplicationCenterServiceInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNamespaceId(v string) *ListApplicationCenterServiceInstancesRequest
	GetNamespaceId() *string
	SetRegionId(v string) *ListApplicationCenterServiceInstancesRequest
	GetRegionId() *string
}

type ListApplicationCenterServiceInstancesRequest struct {
	// example:
	//
	// cn-beijing:test
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListApplicationCenterServiceInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationCenterServiceInstancesRequest) GoString() string {
	return s.String()
}

func (s *ListApplicationCenterServiceInstancesRequest) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *ListApplicationCenterServiceInstancesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListApplicationCenterServiceInstancesRequest) SetNamespaceId(v string) *ListApplicationCenterServiceInstancesRequest {
	s.NamespaceId = &v
	return s
}

func (s *ListApplicationCenterServiceInstancesRequest) SetRegionId(v string) *ListApplicationCenterServiceInstancesRequest {
	s.RegionId = &v
	return s
}

func (s *ListApplicationCenterServiceInstancesRequest) Validate() error {
	return dara.Validate(s)
}
