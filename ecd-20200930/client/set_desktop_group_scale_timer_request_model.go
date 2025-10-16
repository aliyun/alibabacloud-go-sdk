// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetDesktopGroupScaleTimerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDesktopGroupId(v string) *SetDesktopGroupScaleTimerRequest
	GetDesktopGroupId() *string
	SetRegionId(v string) *SetDesktopGroupScaleTimerRequest
	GetRegionId() *string
	SetScaleTimerInfos(v []*SetDesktopGroupScaleTimerRequestScaleTimerInfos) *SetDesktopGroupScaleTimerRequest
	GetScaleTimerInfos() []*SetDesktopGroupScaleTimerRequestScaleTimerInfos
}

type SetDesktopGroupScaleTimerRequest struct {
	// The ID of the cloud computer pool.
	//
	// This parameter is required.
	//
	// example:
	//
	// dg-2i8qxpv6t1a03****
	DesktopGroupId *string `json:"DesktopGroupId,omitempty" xml:"DesktopGroupId,omitempty"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/196646.html) operation to query the regions supported by Elastic Desktop Service.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The information about the scheduled auto scaling task.
	ScaleTimerInfos []*SetDesktopGroupScaleTimerRequestScaleTimerInfos `json:"ScaleTimerInfos,omitempty" xml:"ScaleTimerInfos,omitempty" type:"Repeated"`
}

func (s SetDesktopGroupScaleTimerRequest) String() string {
	return dara.Prettify(s)
}

func (s SetDesktopGroupScaleTimerRequest) GoString() string {
	return s.String()
}

func (s *SetDesktopGroupScaleTimerRequest) GetDesktopGroupId() *string {
	return s.DesktopGroupId
}

func (s *SetDesktopGroupScaleTimerRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *SetDesktopGroupScaleTimerRequest) GetScaleTimerInfos() []*SetDesktopGroupScaleTimerRequestScaleTimerInfos {
	return s.ScaleTimerInfos
}

func (s *SetDesktopGroupScaleTimerRequest) SetDesktopGroupId(v string) *SetDesktopGroupScaleTimerRequest {
	s.DesktopGroupId = &v
	return s
}

func (s *SetDesktopGroupScaleTimerRequest) SetRegionId(v string) *SetDesktopGroupScaleTimerRequest {
	s.RegionId = &v
	return s
}

func (s *SetDesktopGroupScaleTimerRequest) SetScaleTimerInfos(v []*SetDesktopGroupScaleTimerRequestScaleTimerInfos) *SetDesktopGroupScaleTimerRequest {
	s.ScaleTimerInfos = v
	return s
}

func (s *SetDesktopGroupScaleTimerRequest) Validate() error {
	if s.ScaleTimerInfos != nil {
		for _, item := range s.ScaleTimerInfos {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SetDesktopGroupScaleTimerRequestScaleTimerInfos struct {
	// One option for the auto scaling policy. This option specifies the number of cloud computers that you want to create in the cloud computer pool. Valid values: 0 to 200.
	//
	// example:
	//
	// 5
	BuyResAmount *int32 `json:"BuyResAmount,omitempty" xml:"BuyResAmount,omitempty"`
	// The cron expression of the trigger time.
	//
	// example:
	//
	// 0 0 12 ? 	- 1
	Cron *string `json:"Cron,omitempty" xml:"Cron,omitempty"`
	// The keep-alive duration of a session after the session is disconnected. Unit: milliseconds. Valid values: 180000 (3 minutes) to 345600000 (4 days). A value of 0 indicates that the session always keeps alive.
	//
	// If a session is disconnected by the end user or accidentally due to a factor and the end user does not re-establish a connection with the session within the keep-alive duration, the session expires and unsaved data is deleted. If the end user successfully re-establishes a connection with the session within the keep-alive duration, the end user returns to the session and can still access the original data.
	//
	// example:
	//
	// 1000
	KeepDuration *int64 `json:"KeepDuration,omitempty" xml:"KeepDuration,omitempty"`
	// The load balancing policy for the multi-session cloud computer pool.
	//
	// Valid values:
	//
	// 	- 0: depth-first
	//
	// 	- 1: breadth first.
	//
	// example:
	//
	// 0
	LoadPolicy *int32 `json:"LoadPolicy,omitempty" xml:"LoadPolicy,omitempty"`
	// One option for the auto scaling policy. This option specifies the maximum number of cloud computers that you can create in the cloud computer pool. Valid values: 0 to 200.
	//
	// example:
	//
	// 100
	MaxResAmount *int32 `json:"MaxResAmount,omitempty" xml:"MaxResAmount,omitempty"`
	// One option for the auto scaling policy. This option specifies the minimum number of cloud computers that you must create in the cloud computer pool. Valid values: 0 to 200.
	//
	// example:
	//
	// 5
	MinResAmount *int32 `json:"MinResAmount,omitempty" xml:"MinResAmount,omitempty"`
	// The threshold for the ratio of connected sessions. This parameter is the condition that triggers auto scaling in a multi-session cloud computer pool. Formula:
	//
	// `Ratio of connected sessions = Number of connected sessions/(Total number of cloud computers × Maximum number of sessions allowed for each cloud computer) × 100%`.
	//
	// When the specified threshold is reached, new cloud computers are automatically created. When the specified threshold is not reached, idle cloud computers are released.
	//
	// example:
	//
	// 0.9
	RatioThreshold *float32 `json:"RatioThreshold,omitempty" xml:"RatioThreshold,omitempty"`
	// The type of the auto scaling policy.
	//
	// Valid values:
	//
	// 	- drop
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- normal
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- peak
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- rise
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// rise
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s SetDesktopGroupScaleTimerRequestScaleTimerInfos) String() string {
	return dara.Prettify(s)
}

func (s SetDesktopGroupScaleTimerRequestScaleTimerInfos) GoString() string {
	return s.String()
}

func (s *SetDesktopGroupScaleTimerRequestScaleTimerInfos) GetBuyResAmount() *int32 {
	return s.BuyResAmount
}

func (s *SetDesktopGroupScaleTimerRequestScaleTimerInfos) GetCron() *string {
	return s.Cron
}

func (s *SetDesktopGroupScaleTimerRequestScaleTimerInfos) GetKeepDuration() *int64 {
	return s.KeepDuration
}

func (s *SetDesktopGroupScaleTimerRequestScaleTimerInfos) GetLoadPolicy() *int32 {
	return s.LoadPolicy
}

func (s *SetDesktopGroupScaleTimerRequestScaleTimerInfos) GetMaxResAmount() *int32 {
	return s.MaxResAmount
}

func (s *SetDesktopGroupScaleTimerRequestScaleTimerInfos) GetMinResAmount() *int32 {
	return s.MinResAmount
}

func (s *SetDesktopGroupScaleTimerRequestScaleTimerInfos) GetRatioThreshold() *float32 {
	return s.RatioThreshold
}

func (s *SetDesktopGroupScaleTimerRequestScaleTimerInfos) GetType() *string {
	return s.Type
}

func (s *SetDesktopGroupScaleTimerRequestScaleTimerInfos) SetBuyResAmount(v int32) *SetDesktopGroupScaleTimerRequestScaleTimerInfos {
	s.BuyResAmount = &v
	return s
}

func (s *SetDesktopGroupScaleTimerRequestScaleTimerInfos) SetCron(v string) *SetDesktopGroupScaleTimerRequestScaleTimerInfos {
	s.Cron = &v
	return s
}

func (s *SetDesktopGroupScaleTimerRequestScaleTimerInfos) SetKeepDuration(v int64) *SetDesktopGroupScaleTimerRequestScaleTimerInfos {
	s.KeepDuration = &v
	return s
}

func (s *SetDesktopGroupScaleTimerRequestScaleTimerInfos) SetLoadPolicy(v int32) *SetDesktopGroupScaleTimerRequestScaleTimerInfos {
	s.LoadPolicy = &v
	return s
}

func (s *SetDesktopGroupScaleTimerRequestScaleTimerInfos) SetMaxResAmount(v int32) *SetDesktopGroupScaleTimerRequestScaleTimerInfos {
	s.MaxResAmount = &v
	return s
}

func (s *SetDesktopGroupScaleTimerRequestScaleTimerInfos) SetMinResAmount(v int32) *SetDesktopGroupScaleTimerRequestScaleTimerInfos {
	s.MinResAmount = &v
	return s
}

func (s *SetDesktopGroupScaleTimerRequestScaleTimerInfos) SetRatioThreshold(v float32) *SetDesktopGroupScaleTimerRequestScaleTimerInfos {
	s.RatioThreshold = &v
	return s
}

func (s *SetDesktopGroupScaleTimerRequestScaleTimerInfos) SetType(v string) *SetDesktopGroupScaleTimerRequestScaleTimerInfos {
	s.Type = &v
	return s
}

func (s *SetDesktopGroupScaleTimerRequestScaleTimerInfos) Validate() error {
	return dara.Validate(s)
}
