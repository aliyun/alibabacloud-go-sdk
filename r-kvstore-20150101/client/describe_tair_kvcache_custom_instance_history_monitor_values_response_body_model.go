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
	Datapoints *string `json:"Datapoints,omitempty" xml:"Datapoints,omitempty"`
	NextToken  *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	Period     *string `json:"Period,omitempty" xml:"Period,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
