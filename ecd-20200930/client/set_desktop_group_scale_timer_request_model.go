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
	// The cloud computer pool ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// dg-2i8qxpv6t1a03****
	DesktopGroupId *string `json:"DesktopGroupId,omitempty" xml:"DesktopGroupId,omitempty"`
	// The region ID. You can call [DescribeRegions](https://help.aliyun.com/document_detail/196646.html) to query the regions supported by Elastic Desktop Service.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The list of scheduled task information for automatic scaling.
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
	// The number of cloud computers to purchase. This is one of the scaling policy parameters. Valid values: 0 to 200.
	//
	// example:
	//
	// 5
	BuyResAmount *int32 `json:"BuyResAmount,omitempty" xml:"BuyResAmount,omitempty"`
	// The cron expression for the trigger time.
	//
	// example:
	//
	// 0 0 12 ? 	- 1
	Cron *string `json:"Cron,omitempty" xml:"Cron,omitempty"`
	// The duration for which a session is retained after disconnection. Unit: milliseconds. Valid values: 180000 (3 minutes) to 345600000 (4 days). A value of 0 indicates that the session is always retained.
	//
	// When a session is disconnected because the user actively disconnects or because of other unexpected factors, the retention period starts from the time of disconnection. If the user does not reconnect to the session within the retention period, the session is logged off and all unsaved data is destroyed. If the user reconnects within the retention period, the user can still access the original session and the data that existed before the disconnection.
	//
	// example:
	//
	// 180000
	KeepDuration *int64 `json:"KeepDuration,omitempty" xml:"KeepDuration,omitempty"`
	// The load balancing policy for the multi-session cloud computer pool.
	//
	// example:
	//
	// 0
	LoadPolicy *int32 `json:"LoadPolicy,omitempty" xml:"LoadPolicy,omitempty"`
	// The maximum number of cloud computers. This is one of the scaling policy parameters. Valid values: 0 to 200.
	//
	// example:
	//
	// 100
	MaxResAmount *int32 `json:"MaxResAmount,omitempty" xml:"MaxResAmount,omitempty"`
	// The minimum number of cloud computers. This is one of the scaling policy parameters. Valid values: 0 to 200.
	//
	// example:
	//
	// 5
	MinResAmount *int32 `json:"MinResAmount,omitempty" xml:"MinResAmount,omitempty"`
	// The session occupancy threshold, which is used as the trigger condition for automatic scaling of the multi-session cloud computer pool. The session occupancy is calculated by using the following formula:
	//
	// ```Session occupancy = Number of attached sessions / (Total number of cloud computer resources × Maximum number of sessions supported per cloud computer) × 100%```
	//
	// When the session occupancy reaches this threshold, new cloud computers are created. When the session occupancy does not reach this threshold, excess cloud computers are deleted.
	//
	// example:
	//
	// 0.85
	RatioThreshold *float32 `json:"RatioThreshold,omitempty" xml:"RatioThreshold,omitempty"`
	// The policy type.
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
