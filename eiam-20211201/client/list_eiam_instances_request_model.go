// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEiamInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceIds(v []*string) *ListEiamInstancesRequest
	GetInstanceIds() []*string
	SetInstanceRegionId(v string) *ListEiamInstancesRequest
	GetInstanceRegionId() *string
}

type ListEiamInstancesRequest struct {
	// The instance ID list.
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	// The region in which the instance resides.
	//
	// example:
	//
	// cn-hangzhou
	InstanceRegionId *string `json:"InstanceRegionId,omitempty" xml:"InstanceRegionId,omitempty"`
}

func (s ListEiamInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListEiamInstancesRequest) GoString() string {
	return s.String()
}

func (s *ListEiamInstancesRequest) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *ListEiamInstancesRequest) GetInstanceRegionId() *string {
	return s.InstanceRegionId
}

func (s *ListEiamInstancesRequest) SetInstanceIds(v []*string) *ListEiamInstancesRequest {
	s.InstanceIds = v
	return s
}

func (s *ListEiamInstancesRequest) SetInstanceRegionId(v string) *ListEiamInstancesRequest {
	s.InstanceRegionId = &v
	return s
}

func (s *ListEiamInstancesRequest) Validate() error {
	return dara.Validate(s)
}
