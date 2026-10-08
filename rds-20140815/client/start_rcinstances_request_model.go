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
	BatchOptimization *string   `json:"BatchOptimization,omitempty" xml:"BatchOptimization,omitempty"`
	InstanceIds       []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	RegionId          *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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
