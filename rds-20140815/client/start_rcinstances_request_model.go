// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartRCInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBatchOptimization(v string) *StartRCInstancesRequest
	GetBatchOptimization() *string
	SetInstanceIds(v []*string) *StartRCInstancesRequest
	GetInstanceIds() []*string
	SetRegionId(v string) *StartRCInstancesRequest
	GetRegionId() *string
}

type StartRCInstancesRequest struct {
	// The batch operation mode. Set the value to **AllTogether**. In this mode, a success message is returned if all specified instances are started. If an instance fails the verification, none of the specified instances can be started and an error message is returned.
	//
	// example:
	//
	// AllTogether
	BatchOptimization *string `json:"BatchOptimization,omitempty" xml:"BatchOptimization,omitempty"`
	// The node IDs.
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	// The region ID of the instance. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/26243.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s StartRCInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s StartRCInstancesRequest) GoString() string {
	return s.String()
}

func (s *StartRCInstancesRequest) GetBatchOptimization() *string {
	return s.BatchOptimization
}

func (s *StartRCInstancesRequest) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *StartRCInstancesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *StartRCInstancesRequest) SetBatchOptimization(v string) *StartRCInstancesRequest {
	s.BatchOptimization = &v
	return s
}

func (s *StartRCInstancesRequest) SetInstanceIds(v []*string) *StartRCInstancesRequest {
	s.InstanceIds = v
	return s
}

func (s *StartRCInstancesRequest) SetRegionId(v string) *StartRCInstancesRequest {
	s.RegionId = &v
	return s
}

func (s *StartRCInstancesRequest) Validate() error {
	return dara.Validate(s)
}
