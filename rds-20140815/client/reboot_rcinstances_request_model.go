// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRebootRCInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBatchOptimization(v string) *RebootRCInstancesRequest
	GetBatchOptimization() *string
	SetForceReboot(v bool) *RebootRCInstancesRequest
	GetForceReboot() *bool
	SetInstanceIds(v []*string) *RebootRCInstancesRequest
	GetInstanceIds() []*string
	SetRegionId(v string) *RebootRCInstancesRequest
	GetRegionId() *string
}

type RebootRCInstancesRequest struct {
	// The batch operation mode. Set the value to **AllTogether**. In this mode, if all specified instances are restarted, a success message is returned. If an instance fails the verification, none of the specified instances can be restarted and an error message is returned.
	//
	// example:
	//
	// AllTogether
	BatchOptimization *string `json:"BatchOptimization,omitempty" xml:"BatchOptimization,omitempty"`
	// Specifies whether to forcefully restart the instance. Valid values:
	//
	// 	- **true**
	//
	// 	- **false*	- (default)
	//
	// example:
	//
	// false
	ForceReboot *bool `json:"ForceReboot,omitempty" xml:"ForceReboot,omitempty"`
	// The node IDs.
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	// The region ID of the instance. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/26243.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s RebootRCInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s RebootRCInstancesRequest) GoString() string {
	return s.String()
}

func (s *RebootRCInstancesRequest) GetBatchOptimization() *string {
	return s.BatchOptimization
}

func (s *RebootRCInstancesRequest) GetForceReboot() *bool {
	return s.ForceReboot
}

func (s *RebootRCInstancesRequest) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *RebootRCInstancesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *RebootRCInstancesRequest) SetBatchOptimization(v string) *RebootRCInstancesRequest {
	s.BatchOptimization = &v
	return s
}

func (s *RebootRCInstancesRequest) SetForceReboot(v bool) *RebootRCInstancesRequest {
	s.ForceReboot = &v
	return s
}

func (s *RebootRCInstancesRequest) SetInstanceIds(v []*string) *RebootRCInstancesRequest {
	s.InstanceIds = v
	return s
}

func (s *RebootRCInstancesRequest) SetRegionId(v string) *RebootRCInstancesRequest {
	s.RegionId = &v
	return s
}

func (s *RebootRCInstancesRequest) Validate() error {
	return dara.Validate(s)
}
