// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRebootRCInstancesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBatchOptimization(v string) *RebootRCInstancesShrinkRequest
	GetBatchOptimization() *string
	SetForceReboot(v bool) *RebootRCInstancesShrinkRequest
	GetForceReboot() *bool
	SetInstanceIdsShrink(v string) *RebootRCInstancesShrinkRequest
	GetInstanceIdsShrink() *string
	SetRegionId(v string) *RebootRCInstancesShrinkRequest
	GetRegionId() *string
}

type RebootRCInstancesShrinkRequest struct {
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
	InstanceIdsShrink *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
	// The region ID of the instance. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/26243.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s RebootRCInstancesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s RebootRCInstancesShrinkRequest) GoString() string {
	return s.String()
}

func (s *RebootRCInstancesShrinkRequest) GetBatchOptimization() *string {
	return s.BatchOptimization
}

func (s *RebootRCInstancesShrinkRequest) GetForceReboot() *bool {
	return s.ForceReboot
}

func (s *RebootRCInstancesShrinkRequest) GetInstanceIdsShrink() *string {
	return s.InstanceIdsShrink
}

func (s *RebootRCInstancesShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *RebootRCInstancesShrinkRequest) SetBatchOptimization(v string) *RebootRCInstancesShrinkRequest {
	s.BatchOptimization = &v
	return s
}

func (s *RebootRCInstancesShrinkRequest) SetForceReboot(v bool) *RebootRCInstancesShrinkRequest {
	s.ForceReboot = &v
	return s
}

func (s *RebootRCInstancesShrinkRequest) SetInstanceIdsShrink(v string) *RebootRCInstancesShrinkRequest {
	s.InstanceIdsShrink = &v
	return s
}

func (s *RebootRCInstancesShrinkRequest) SetRegionId(v string) *RebootRCInstancesShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *RebootRCInstancesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
