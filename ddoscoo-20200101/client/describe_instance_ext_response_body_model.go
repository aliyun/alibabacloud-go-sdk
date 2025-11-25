// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceExtResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceExtSpecs(v []*DescribeInstanceExtResponseBodyInstanceExtSpecs) *DescribeInstanceExtResponseBody
	GetInstanceExtSpecs() []*DescribeInstanceExtResponseBodyInstanceExtSpecs
	SetRequestId(v string) *DescribeInstanceExtResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *DescribeInstanceExtResponseBody
	GetTotalCount() *int64
}

type DescribeInstanceExtResponseBody struct {
	// The extended information about the Anti-DDoS Proxy instance.
	InstanceExtSpecs []*DescribeInstanceExtResponseBodyInstanceExtSpecs `json:"InstanceExtSpecs,omitempty" xml:"InstanceExtSpecs,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// CF33B4C3-196E-4015-AADD-5CAD00057B80
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of queried instances.
	//
	// example:
	//
	// 10
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeInstanceExtResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceExtResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceExtResponseBody) GetInstanceExtSpecs() []*DescribeInstanceExtResponseBodyInstanceExtSpecs {
	return s.InstanceExtSpecs
}

func (s *DescribeInstanceExtResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeInstanceExtResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeInstanceExtResponseBody) SetInstanceExtSpecs(v []*DescribeInstanceExtResponseBodyInstanceExtSpecs) *DescribeInstanceExtResponseBody {
	s.InstanceExtSpecs = v
	return s
}

func (s *DescribeInstanceExtResponseBody) SetRequestId(v string) *DescribeInstanceExtResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstanceExtResponseBody) SetTotalCount(v int64) *DescribeInstanceExtResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeInstanceExtResponseBody) Validate() error {
	if s.InstanceExtSpecs != nil {
		for _, item := range s.InstanceExtSpecs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeInstanceExtResponseBodyInstanceExtSpecs struct {
	// The function plan. Valid values:
	//
	// 	- **0**: Standard
	//
	// 	- **1**: Enhanced
	//
	// example:
	//
	// 0
	FunctionVersion *int64 `json:"FunctionVersion,omitempty" xml:"FunctionVersion,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// ddoscoo-cn-i7m25564****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The clean bandwidth. Unit: Mbit/s.
	//
	// example:
	//
	// 100
	NormalBandwidth *int64 `json:"NormalBandwidth,omitempty" xml:"NormalBandwidth,omitempty"`
	// The type of the instance. Valid values:
	//
	// 	- **0**: Anti-DDoS Proxy (Outside Chinese Mainland) instance of the Insurance mitigation plan
	//
	// 	- **1**: Anti-DDoS Proxy (Outside Chinese Mainland) instance of the Unlimited mitigation plan
	//
	// 	- **2**: Anti-DDoS Proxy (Outside Chinese Mainland) instance of the Chinese Mainland Acceleration (CMA) mitigation plan
	//
	// 	- **3**: Anti-DDoS Proxy (Outside Chinese Mainland) instance of the Secure Chinese Mainland Acceleration (Sec-CMA) mitigation plan
	//
	// 	- **9**: Anti-DDoS Proxy (Chinese Mainland) instance of the Profession mitigation plan
	//
	// example:
	//
	// 0
	ProductPlan *int64 `json:"ProductPlan,omitempty" xml:"ProductPlan,omitempty"`
	// The Internet service provider (ISP) line of the Anti-DDoS Proxy (Chinese Mainland) instance.
	//
	// example:
	//
	// coop-line-001
	ServicePartner *string `json:"ServicePartner,omitempty" xml:"ServicePartner,omitempty"`
}

func (s DescribeInstanceExtResponseBodyInstanceExtSpecs) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceExtResponseBodyInstanceExtSpecs) GoString() string {
	return s.String()
}

func (s *DescribeInstanceExtResponseBodyInstanceExtSpecs) GetFunctionVersion() *int64 {
	return s.FunctionVersion
}

func (s *DescribeInstanceExtResponseBodyInstanceExtSpecs) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeInstanceExtResponseBodyInstanceExtSpecs) GetNormalBandwidth() *int64 {
	return s.NormalBandwidth
}

func (s *DescribeInstanceExtResponseBodyInstanceExtSpecs) GetProductPlan() *int64 {
	return s.ProductPlan
}

func (s *DescribeInstanceExtResponseBodyInstanceExtSpecs) GetServicePartner() *string {
	return s.ServicePartner
}

func (s *DescribeInstanceExtResponseBodyInstanceExtSpecs) SetFunctionVersion(v int64) *DescribeInstanceExtResponseBodyInstanceExtSpecs {
	s.FunctionVersion = &v
	return s
}

func (s *DescribeInstanceExtResponseBodyInstanceExtSpecs) SetInstanceId(v string) *DescribeInstanceExtResponseBodyInstanceExtSpecs {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstanceExtResponseBodyInstanceExtSpecs) SetNormalBandwidth(v int64) *DescribeInstanceExtResponseBodyInstanceExtSpecs {
	s.NormalBandwidth = &v
	return s
}

func (s *DescribeInstanceExtResponseBodyInstanceExtSpecs) SetProductPlan(v int64) *DescribeInstanceExtResponseBodyInstanceExtSpecs {
	s.ProductPlan = &v
	return s
}

func (s *DescribeInstanceExtResponseBodyInstanceExtSpecs) SetServicePartner(v string) *DescribeInstanceExtResponseBodyInstanceExtSpecs {
	s.ServicePartner = &v
	return s
}

func (s *DescribeInstanceExtResponseBodyInstanceExtSpecs) Validate() error {
	return dara.Validate(s)
}
