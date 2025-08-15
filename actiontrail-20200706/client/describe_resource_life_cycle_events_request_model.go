// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeResourceLifeCycleEventsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetResourceType(v string) *DescribeResourceLifeCycleEventsRequest
	GetResourceType() *string
	SetServiceName(v string) *DescribeResourceLifeCycleEventsRequest
	GetServiceName() *string
}

type DescribeResourceLifeCycleEventsRequest struct {
	// example:
	//
	// ACS::ECS::Instance
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// example:
	//
	// ECS
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
}

func (s DescribeResourceLifeCycleEventsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeResourceLifeCycleEventsRequest) GoString() string {
	return s.String()
}

func (s *DescribeResourceLifeCycleEventsRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *DescribeResourceLifeCycleEventsRequest) GetServiceName() *string {
	return s.ServiceName
}

func (s *DescribeResourceLifeCycleEventsRequest) SetResourceType(v string) *DescribeResourceLifeCycleEventsRequest {
	s.ResourceType = &v
	return s
}

func (s *DescribeResourceLifeCycleEventsRequest) SetServiceName(v string) *DescribeResourceLifeCycleEventsRequest {
	s.ServiceName = &v
	return s
}

func (s *DescribeResourceLifeCycleEventsRequest) Validate() error {
	return dara.Validate(s)
}
