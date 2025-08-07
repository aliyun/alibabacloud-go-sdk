// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBClusterServerlessConfResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAgileScaleMax(v string) *DescribeDBClusterServerlessConfResponseBody
	GetAgileScaleMax() *string
	SetAllowShutDown(v string) *DescribeDBClusterServerlessConfResponseBody
	GetAllowShutDown() *string
	SetDBClusterId(v string) *DescribeDBClusterServerlessConfResponseBody
	GetDBClusterId() *string
	SetRequestId(v string) *DescribeDBClusterServerlessConfResponseBody
	GetRequestId() *string
	SetScaleApRoNumMax(v string) *DescribeDBClusterServerlessConfResponseBody
	GetScaleApRoNumMax() *string
	SetScaleApRoNumMin(v string) *DescribeDBClusterServerlessConfResponseBody
	GetScaleApRoNumMin() *string
	SetScaleMax(v string) *DescribeDBClusterServerlessConfResponseBody
	GetScaleMax() *string
	SetScaleMin(v string) *DescribeDBClusterServerlessConfResponseBody
	GetScaleMin() *string
	SetScaleRoNumMax(v string) *DescribeDBClusterServerlessConfResponseBody
	GetScaleRoNumMax() *string
	SetScaleRoNumMin(v string) *DescribeDBClusterServerlessConfResponseBody
	GetScaleRoNumMin() *string
	SetSecondsUntilAutoPause(v string) *DescribeDBClusterServerlessConfResponseBody
	GetSecondsUntilAutoPause() *string
	SetServerlessRuleCpuEnlargeThreshold(v string) *DescribeDBClusterServerlessConfResponseBody
	GetServerlessRuleCpuEnlargeThreshold() *string
	SetServerlessRuleCpuShrinkThreshold(v string) *DescribeDBClusterServerlessConfResponseBody
	GetServerlessRuleCpuShrinkThreshold() *string
	SetServerlessRuleMode(v string) *DescribeDBClusterServerlessConfResponseBody
	GetServerlessRuleMode() *string
	SetSwitchs(v string) *DescribeDBClusterServerlessConfResponseBody
	GetSwitchs() *string
	SetTraditionalScaleMaxThreshold(v string) *DescribeDBClusterServerlessConfResponseBody
	GetTraditionalScaleMaxThreshold() *string
}

type DescribeDBClusterServerlessConfResponseBody struct {
	AgileScaleMax *string `json:"AgileScaleMax,omitempty" xml:"AgileScaleMax,omitempty"`
	// Whether to enable idle shutdown. Values:
	//
	// - **true**: Enable
	//
	// - **false**: Disable (default)
	//
	// example:
	//
	// true
	AllowShutDown *string `json:"AllowShutDown,omitempty" xml:"AllowShutDown,omitempty"`
	// Serverless cluster ID.
	//
	// example:
	//
	// pc-bp10gr51qasnl****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// Request ID.
	//
	// example:
	//
	// 5E71541A-6007-4DCC-A38A-F872C31FEB45
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The maximum number of read-only column store nodes. Valid values: 0 to 15.
	//
	// example:
	//
	// 1
	ScaleApRoNumMax *string `json:"ScaleApRoNumMax,omitempty" xml:"ScaleApRoNumMax,omitempty"`
	// The minimum number of read-only column store nodes. Valid values: 0 to 15.
	//
	// example:
	//
	// 1
	ScaleApRoNumMin *string `json:"ScaleApRoNumMin,omitempty" xml:"ScaleApRoNumMin,omitempty"`
	// Maximum scaling limit for a single node. Range: 1 PCU~32 PCU.
	//
	// example:
	//
	// 3
	ScaleMax *string `json:"ScaleMax,omitempty" xml:"ScaleMax,omitempty"`
	// Minimum scaling limit for a single node. Range: 1 PCU~31 PCU.
	//
	// example:
	//
	// 1
	ScaleMin *string `json:"ScaleMin,omitempty" xml:"ScaleMin,omitempty"`
	// Maximum scaling limit for the number of read-only nodes. Range: 0~15.
	//
	// example:
	//
	// 4
	ScaleRoNumMax *string `json:"ScaleRoNumMax,omitempty" xml:"ScaleRoNumMax,omitempty"`
	// Minimum scaling limit for the number of read-only nodes. Range: 0~15.
	//
	// example:
	//
	// 2
	ScaleRoNumMin *string `json:"ScaleRoNumMin,omitempty" xml:"ScaleRoNumMin,omitempty"`
	// Detection duration for idle shutdown. Range: 300~86,400. Unit: seconds. The detection duration must be a multiple of 300 seconds.
	//
	// example:
	//
	// 10
	SecondsUntilAutoPause *string `json:"SecondsUntilAutoPause,omitempty" xml:"SecondsUntilAutoPause,omitempty"`
	// CPU upscale threshold.
	//
	// example:
	//
	// 60
	ServerlessRuleCpuEnlargeThreshold *string `json:"ServerlessRuleCpuEnlargeThreshold,omitempty" xml:"ServerlessRuleCpuEnlargeThreshold,omitempty"`
	// CPU downscale threshold.
	//
	// example:
	//
	// 30
	ServerlessRuleCpuShrinkThreshold *string `json:"ServerlessRuleCpuShrinkThreshold,omitempty" xml:"ServerlessRuleCpuShrinkThreshold,omitempty"`
	// Elasticity sensitivity. Values:
	//
	// - normal: Standard
	//
	// - flexible: Sensitive
	//
	// example:
	//
	// normal
	ServerlessRuleMode *string `json:"ServerlessRuleMode,omitempty" xml:"ServerlessRuleMode,omitempty"`
	// Whether steady state is enabled. Values:
	//
	// 1: Enabled
	//
	// 0: Disabled
	//
	// example:
	//
	// 1
	Switchs                      *string `json:"Switchs,omitempty" xml:"Switchs,omitempty"`
	TraditionalScaleMaxThreshold *string `json:"TraditionalScaleMaxThreshold,omitempty" xml:"TraditionalScaleMaxThreshold,omitempty"`
}

func (s DescribeDBClusterServerlessConfResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterServerlessConfResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterServerlessConfResponseBody) GetAgileScaleMax() *string {
	return s.AgileScaleMax
}

func (s *DescribeDBClusterServerlessConfResponseBody) GetAllowShutDown() *string {
	return s.AllowShutDown
}

func (s *DescribeDBClusterServerlessConfResponseBody) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeDBClusterServerlessConfResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDBClusterServerlessConfResponseBody) GetScaleApRoNumMax() *string {
	return s.ScaleApRoNumMax
}

func (s *DescribeDBClusterServerlessConfResponseBody) GetScaleApRoNumMin() *string {
	return s.ScaleApRoNumMin
}

func (s *DescribeDBClusterServerlessConfResponseBody) GetScaleMax() *string {
	return s.ScaleMax
}

func (s *DescribeDBClusterServerlessConfResponseBody) GetScaleMin() *string {
	return s.ScaleMin
}

func (s *DescribeDBClusterServerlessConfResponseBody) GetScaleRoNumMax() *string {
	return s.ScaleRoNumMax
}

func (s *DescribeDBClusterServerlessConfResponseBody) GetScaleRoNumMin() *string {
	return s.ScaleRoNumMin
}

func (s *DescribeDBClusterServerlessConfResponseBody) GetSecondsUntilAutoPause() *string {
	return s.SecondsUntilAutoPause
}

func (s *DescribeDBClusterServerlessConfResponseBody) GetServerlessRuleCpuEnlargeThreshold() *string {
	return s.ServerlessRuleCpuEnlargeThreshold
}

func (s *DescribeDBClusterServerlessConfResponseBody) GetServerlessRuleCpuShrinkThreshold() *string {
	return s.ServerlessRuleCpuShrinkThreshold
}

func (s *DescribeDBClusterServerlessConfResponseBody) GetServerlessRuleMode() *string {
	return s.ServerlessRuleMode
}

func (s *DescribeDBClusterServerlessConfResponseBody) GetSwitchs() *string {
	return s.Switchs
}

func (s *DescribeDBClusterServerlessConfResponseBody) GetTraditionalScaleMaxThreshold() *string {
	return s.TraditionalScaleMaxThreshold
}

func (s *DescribeDBClusterServerlessConfResponseBody) SetAgileScaleMax(v string) *DescribeDBClusterServerlessConfResponseBody {
	s.AgileScaleMax = &v
	return s
}

func (s *DescribeDBClusterServerlessConfResponseBody) SetAllowShutDown(v string) *DescribeDBClusterServerlessConfResponseBody {
	s.AllowShutDown = &v
	return s
}

func (s *DescribeDBClusterServerlessConfResponseBody) SetDBClusterId(v string) *DescribeDBClusterServerlessConfResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *DescribeDBClusterServerlessConfResponseBody) SetRequestId(v string) *DescribeDBClusterServerlessConfResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBClusterServerlessConfResponseBody) SetScaleApRoNumMax(v string) *DescribeDBClusterServerlessConfResponseBody {
	s.ScaleApRoNumMax = &v
	return s
}

func (s *DescribeDBClusterServerlessConfResponseBody) SetScaleApRoNumMin(v string) *DescribeDBClusterServerlessConfResponseBody {
	s.ScaleApRoNumMin = &v
	return s
}

func (s *DescribeDBClusterServerlessConfResponseBody) SetScaleMax(v string) *DescribeDBClusterServerlessConfResponseBody {
	s.ScaleMax = &v
	return s
}

func (s *DescribeDBClusterServerlessConfResponseBody) SetScaleMin(v string) *DescribeDBClusterServerlessConfResponseBody {
	s.ScaleMin = &v
	return s
}

func (s *DescribeDBClusterServerlessConfResponseBody) SetScaleRoNumMax(v string) *DescribeDBClusterServerlessConfResponseBody {
	s.ScaleRoNumMax = &v
	return s
}

func (s *DescribeDBClusterServerlessConfResponseBody) SetScaleRoNumMin(v string) *DescribeDBClusterServerlessConfResponseBody {
	s.ScaleRoNumMin = &v
	return s
}

func (s *DescribeDBClusterServerlessConfResponseBody) SetSecondsUntilAutoPause(v string) *DescribeDBClusterServerlessConfResponseBody {
	s.SecondsUntilAutoPause = &v
	return s
}

func (s *DescribeDBClusterServerlessConfResponseBody) SetServerlessRuleCpuEnlargeThreshold(v string) *DescribeDBClusterServerlessConfResponseBody {
	s.ServerlessRuleCpuEnlargeThreshold = &v
	return s
}

func (s *DescribeDBClusterServerlessConfResponseBody) SetServerlessRuleCpuShrinkThreshold(v string) *DescribeDBClusterServerlessConfResponseBody {
	s.ServerlessRuleCpuShrinkThreshold = &v
	return s
}

func (s *DescribeDBClusterServerlessConfResponseBody) SetServerlessRuleMode(v string) *DescribeDBClusterServerlessConfResponseBody {
	s.ServerlessRuleMode = &v
	return s
}

func (s *DescribeDBClusterServerlessConfResponseBody) SetSwitchs(v string) *DescribeDBClusterServerlessConfResponseBody {
	s.Switchs = &v
	return s
}

func (s *DescribeDBClusterServerlessConfResponseBody) SetTraditionalScaleMaxThreshold(v string) *DescribeDBClusterServerlessConfResponseBody {
	s.TraditionalScaleMaxThreshold = &v
	return s
}

func (s *DescribeDBClusterServerlessConfResponseBody) Validate() error {
	return dara.Validate(s)
}
