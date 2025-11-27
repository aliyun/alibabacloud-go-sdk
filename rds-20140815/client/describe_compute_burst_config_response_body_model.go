// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeComputeBurstConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetComputeBurstConfig(v map[string]interface{}) *DescribeComputeBurstConfigResponseBody
	GetComputeBurstConfig() map[string]interface{}
	SetComputeBurstEnabled(v bool) *DescribeComputeBurstConfigResponseBody
	GetComputeBurstEnabled() *bool
	SetRequestId(v string) *DescribeComputeBurstConfigResponseBody
	GetRequestId() *string
}

type DescribeComputeBurstConfigResponseBody struct {
	// The detailed configurations of the assured serverless feature.
	//
	// example:
	//
	// {
	//
	//     "cpuEnlargeThreshold": "60",
	//
	//     "memoryEnlargeThreshold": "60",
	//
	//     "scaleMaxMemory": "4",
	//
	//     "memoryShrinkThreshold": "50",
	//
	//     "scaleMaxCpus": "2",
	//
	//     "cpuShrinkThreshold": "30"
	//
	//   }
	ComputeBurstConfig map[string]interface{} `json:"ComputeBurstConfig,omitempty" xml:"ComputeBurstConfig,omitempty"`
	// Indicates whether the assured serverless feature is enabled. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	ComputeBurstEnabled *bool `json:"ComputeBurstEnabled,omitempty" xml:"ComputeBurstEnabled,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 847BA085-B377-4BFA-8267-F82345ECE1D2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeComputeBurstConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeComputeBurstConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeComputeBurstConfigResponseBody) GetComputeBurstConfig() map[string]interface{} {
	return s.ComputeBurstConfig
}

func (s *DescribeComputeBurstConfigResponseBody) GetComputeBurstEnabled() *bool {
	return s.ComputeBurstEnabled
}

func (s *DescribeComputeBurstConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeComputeBurstConfigResponseBody) SetComputeBurstConfig(v map[string]interface{}) *DescribeComputeBurstConfigResponseBody {
	s.ComputeBurstConfig = v
	return s
}

func (s *DescribeComputeBurstConfigResponseBody) SetComputeBurstEnabled(v bool) *DescribeComputeBurstConfigResponseBody {
	s.ComputeBurstEnabled = &v
	return s
}

func (s *DescribeComputeBurstConfigResponseBody) SetRequestId(v string) *DescribeComputeBurstConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeComputeBurstConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
