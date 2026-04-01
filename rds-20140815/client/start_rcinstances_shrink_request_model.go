// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartRCInstancesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBatchOptimization(v string) *StartRCInstancesShrinkRequest
	GetBatchOptimization() *string
	SetInstanceIdsShrink(v string) *StartRCInstancesShrinkRequest
	GetInstanceIdsShrink() *string
	SetRegionId(v string) *StartRCInstancesShrinkRequest
	GetRegionId() *string
}

type StartRCInstancesShrinkRequest struct {
	// The batch operation mode. Set the value to **AllTogether**. In this mode, a success message is returned if all specified instances are started. If an instance fails the verification, none of the specified instances can be started and an error message is returned.
	//
	// example:
	//
	// AllTogether
	BatchOptimization *string `json:"BatchOptimization,omitempty" xml:"BatchOptimization,omitempty"`
	// The node IDs.
	InstanceIdsShrink *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
	// The region ID of the instance. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/26243.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s StartRCInstancesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s StartRCInstancesShrinkRequest) GoString() string {
	return s.String()
}

func (s *StartRCInstancesShrinkRequest) GetBatchOptimization() *string {
	return s.BatchOptimization
}

func (s *StartRCInstancesShrinkRequest) GetInstanceIdsShrink() *string {
	return s.InstanceIdsShrink
}

func (s *StartRCInstancesShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *StartRCInstancesShrinkRequest) SetBatchOptimization(v string) *StartRCInstancesShrinkRequest {
	s.BatchOptimization = &v
	return s
}

func (s *StartRCInstancesShrinkRequest) SetInstanceIdsShrink(v string) *StartRCInstancesShrinkRequest {
	s.InstanceIdsShrink = &v
	return s
}

func (s *StartRCInstancesShrinkRequest) SetRegionId(v string) *StartRCInstancesShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *StartRCInstancesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
