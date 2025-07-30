// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTairKVCacheCustomInstanceHistoryMonitorValuesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDatapoints(v string) *DescribeTairKVCacheCustomInstanceHistoryMonitorValuesResponseBody
	GetDatapoints() *string
	SetNextToken(v string) *DescribeTairKVCacheCustomInstanceHistoryMonitorValuesResponseBody
	GetNextToken() *string
	SetPeriod(v string) *DescribeTairKVCacheCustomInstanceHistoryMonitorValuesResponseBody
	GetPeriod() *string
	SetRequestId(v string) *DescribeTairKVCacheCustomInstanceHistoryMonitorValuesResponseBody
	GetRequestId() *string
}

type DescribeTairKVCacheCustomInstanceHistoryMonitorValuesResponseBody struct {
	// example:
	//
	// { “timestamp”: 1490164200000, “Maximum”: 100, “userId”: “1234567898765432”, “Minimum”: 4.55, “instanceId”: “i-bp18abl200xk9599ck7c”, “Average”: 93.84 }
	Datapoints *string `json:"Datapoints,omitempty" xml:"Datapoints,omitempty"`
	// example:
	//
	// 212db86sca4384811e0b5e8707ec2****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 60
	Period *string `json:"Period,omitempty" xml:"Period,omitempty"`
	// example:
	//
	// F3F44BE3-5419-4B61-9BAC-E66E295A****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeTairKVCacheCustomInstanceHistoryMonitorValuesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeTairKVCacheCustomInstanceHistoryMonitorValuesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTairKVCacheCustomInstanceHistoryMonitorValuesResponseBody) GetDatapoints() *string {
	return s.Datapoints
}

func (s *DescribeTairKVCacheCustomInstanceHistoryMonitorValuesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeTairKVCacheCustomInstanceHistoryMonitorValuesResponseBody) GetPeriod() *string {
	return s.Period
}

func (s *DescribeTairKVCacheCustomInstanceHistoryMonitorValuesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeTairKVCacheCustomInstanceHistoryMonitorValuesResponseBody) SetDatapoints(v string) *DescribeTairKVCacheCustomInstanceHistoryMonitorValuesResponseBody {
	s.Datapoints = &v
	return s
}

func (s *DescribeTairKVCacheCustomInstanceHistoryMonitorValuesResponseBody) SetNextToken(v string) *DescribeTairKVCacheCustomInstanceHistoryMonitorValuesResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeTairKVCacheCustomInstanceHistoryMonitorValuesResponseBody) SetPeriod(v string) *DescribeTairKVCacheCustomInstanceHistoryMonitorValuesResponseBody {
	s.Period = &v
	return s
}

func (s *DescribeTairKVCacheCustomInstanceHistoryMonitorValuesResponseBody) SetRequestId(v string) *DescribeTairKVCacheCustomInstanceHistoryMonitorValuesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTairKVCacheCustomInstanceHistoryMonitorValuesResponseBody) Validate() error {
	return dara.Validate(s)
}
