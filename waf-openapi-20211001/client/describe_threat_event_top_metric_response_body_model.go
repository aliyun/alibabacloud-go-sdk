// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeThreatEventTopMetricResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeThreatEventTopMetricResponseBody
	GetRequestId() *string
	SetTopMetrics(v []*DescribeThreatEventTopMetricResponseBodyTopMetrics) *DescribeThreatEventTopMetricResponseBody
	GetTopMetrics() []*DescribeThreatEventTopMetricResponseBodyTopMetrics
}

type DescribeThreatEventTopMetricResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 12EF3845-CCEB-4B84-AE60-2B49B*****EE5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of statistics.
	TopMetrics []*DescribeThreatEventTopMetricResponseBodyTopMetrics `json:"TopMetrics,omitempty" xml:"TopMetrics,omitempty" type:"Repeated"`
}

func (s DescribeThreatEventTopMetricResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeThreatEventTopMetricResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeThreatEventTopMetricResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeThreatEventTopMetricResponseBody) GetTopMetrics() []*DescribeThreatEventTopMetricResponseBodyTopMetrics {
	return s.TopMetrics
}

func (s *DescribeThreatEventTopMetricResponseBody) SetRequestId(v string) *DescribeThreatEventTopMetricResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeThreatEventTopMetricResponseBody) SetTopMetrics(v []*DescribeThreatEventTopMetricResponseBodyTopMetrics) *DescribeThreatEventTopMetricResponseBody {
	s.TopMetrics = v
	return s
}

func (s *DescribeThreatEventTopMetricResponseBody) Validate() error {
	if s.TopMetrics != nil {
		for _, item := range s.TopMetrics {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeThreatEventTopMetricResponseBodyTopMetrics struct {
	// The number of attacks.
	//
	// example:
	//
	// 20
	Cnt *int64 `json:"Cnt,omitempty" xml:"Cnt,omitempty"`
	// The country where the source IP address of the attack is located.
	//
	// > This parameter is returned only when **Metric*	- is set to **src**.
	//
	// example:
	//
	// CN
	Country *string `json:"Country,omitempty" xml:"Country,omitempty"`
	// The region where the source IP address of the attack is located.
	//
	// > This parameter is returned only when **Metric*	- is set to **src**.
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The attack value. The meaning of this parameter varies based on the value of **Metric**.
	//
	// - If **Metric*	- is set to **time**, this parameter indicates the attack time.
	//
	// - If **Metric*	- is set to **src**, this parameter indicates the source IP address of the attack.
	//
	// - If **Metric*	- is set to **target**, this parameter indicates the URL of the attack request.
	//
	// - If **Metric*	- is set to **type**, this parameter indicates the attack type. For example, **dirscan*	- indicates directory scan and **webscan*	- indicates web scan. For more information about other attack types, see the description of the **detectType*	- parameter for custom regular expression rules (**regular_custom**) in the [CreateDefenseRule](https://help.aliyun.com/document_detail/461421.html) operation.
	//
	// - If **Metric*	- is set to **tools**, this parameter indicates the attack tool.
	//
	// example:
	//
	// 115.28.209.212
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeThreatEventTopMetricResponseBodyTopMetrics) String() string {
	return dara.Prettify(s)
}

func (s DescribeThreatEventTopMetricResponseBodyTopMetrics) GoString() string {
	return s.String()
}

func (s *DescribeThreatEventTopMetricResponseBodyTopMetrics) GetCnt() *int64 {
	return s.Cnt
}

func (s *DescribeThreatEventTopMetricResponseBodyTopMetrics) GetCountry() *string {
	return s.Country
}

func (s *DescribeThreatEventTopMetricResponseBodyTopMetrics) GetRegion() *string {
	return s.Region
}

func (s *DescribeThreatEventTopMetricResponseBodyTopMetrics) GetValue() *string {
	return s.Value
}

func (s *DescribeThreatEventTopMetricResponseBodyTopMetrics) SetCnt(v int64) *DescribeThreatEventTopMetricResponseBodyTopMetrics {
	s.Cnt = &v
	return s
}

func (s *DescribeThreatEventTopMetricResponseBodyTopMetrics) SetCountry(v string) *DescribeThreatEventTopMetricResponseBodyTopMetrics {
	s.Country = &v
	return s
}

func (s *DescribeThreatEventTopMetricResponseBodyTopMetrics) SetRegion(v string) *DescribeThreatEventTopMetricResponseBodyTopMetrics {
	s.Region = &v
	return s
}

func (s *DescribeThreatEventTopMetricResponseBodyTopMetrics) SetValue(v string) *DescribeThreatEventTopMetricResponseBodyTopMetrics {
	s.Value = &v
	return s
}

func (s *DescribeThreatEventTopMetricResponseBodyTopMetrics) Validate() error {
	return dara.Validate(s)
}
