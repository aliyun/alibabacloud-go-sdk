// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEgressProxyIpPoolsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEgressProxyIpPools(v []*ListEgressProxyIpPoolsResponseBodyEgressProxyIpPools) *ListEgressProxyIpPoolsResponseBody
	GetEgressProxyIpPools() []*ListEgressProxyIpPoolsResponseBodyEgressProxyIpPools
	SetMaxResults(v int32) *ListEgressProxyIpPoolsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListEgressProxyIpPoolsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListEgressProxyIpPoolsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListEgressProxyIpPoolsResponseBody
	GetTotalCount() *int32
}

type ListEgressProxyIpPoolsResponseBody struct {
	EgressProxyIpPools []*ListEgressProxyIpPoolsResponseBodyEgressProxyIpPools `json:"EgressProxyIpPools,omitempty" xml:"EgressProxyIpPools,omitempty" type:"Repeated"`
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// FFydY8mmse+9oiDWdR7XiKw/GLUmEZbYNsGL0evBky1i+FT5KS4q2RSGRQDNmuSWlq6KV1PiloZen3lB3ZXwB0vA==
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// C79118A9-CA6D-5359-B53E-111BE8A03DCF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 3
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListEgressProxyIpPoolsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListEgressProxyIpPoolsResponseBody) GoString() string {
	return s.String()
}

func (s *ListEgressProxyIpPoolsResponseBody) GetEgressProxyIpPools() []*ListEgressProxyIpPoolsResponseBodyEgressProxyIpPools {
	return s.EgressProxyIpPools
}

func (s *ListEgressProxyIpPoolsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListEgressProxyIpPoolsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListEgressProxyIpPoolsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListEgressProxyIpPoolsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListEgressProxyIpPoolsResponseBody) SetEgressProxyIpPools(v []*ListEgressProxyIpPoolsResponseBodyEgressProxyIpPools) *ListEgressProxyIpPoolsResponseBody {
	s.EgressProxyIpPools = v
	return s
}

func (s *ListEgressProxyIpPoolsResponseBody) SetMaxResults(v int32) *ListEgressProxyIpPoolsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListEgressProxyIpPoolsResponseBody) SetNextToken(v string) *ListEgressProxyIpPoolsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListEgressProxyIpPoolsResponseBody) SetRequestId(v string) *ListEgressProxyIpPoolsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListEgressProxyIpPoolsResponseBody) SetTotalCount(v int32) *ListEgressProxyIpPoolsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListEgressProxyIpPoolsResponseBody) Validate() error {
	if s.EgressProxyIpPools != nil {
		for _, item := range s.EgressProxyIpPools {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListEgressProxyIpPoolsResponseBodyEgressProxyIpPools struct {
	// example:
	//
	// Intranet
	AddressType *string `json:"AddressType,omitempty" xml:"AddressType,omitempty"`
	// example:
	//
	// 2026-03-11T11:18:35
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// {\\"Type\\": \\"text\\", \\"Value\\": \\"Ensures that assets are protected by Cloud Firewall. Requirement: Enable the Internet Firewall feature.\\"}
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// epg-gw80bb6a94f835db****
	EgressProxyGatewayId *string `json:"EgressProxyGatewayId,omitempty" xml:"EgressProxyGatewayId,omitempty"`
	// example:
	//
	// eppool-gw80d96a96ca0b00****
	EgressProxyIpPoolId *string `json:"EgressProxyIpPoolId,omitempty" xml:"EgressProxyIpPoolId,omitempty"`
	// example:
	//
	// testEpg
	EgressProxyIpPoolName *string `json:"EgressProxyIpPoolName,omitempty" xml:"EgressProxyIpPoolName,omitempty"`
	// example:
	//
	// SourceIP
	HashFactor *string `json:"HashFactor,omitempty" xml:"HashFactor,omitempty"`
	// example:
	//
	// X-Forwarded-For
	HashHeaderName *string `json:"HashHeaderName,omitempty" xml:"HashHeaderName,omitempty"`
	// example:
	//
	// Random
	Scheduler *string `json:"Scheduler,omitempty" xml:"Scheduler,omitempty"`
	// example:
	//
	// Active
	Status *string                                                     `json:"Status,omitempty" xml:"Status,omitempty"`
	Tags   []*ListEgressProxyIpPoolsResponseBodyEgressProxyIpPoolsTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s ListEgressProxyIpPoolsResponseBodyEgressProxyIpPools) String() string {
	return dara.Prettify(s)
}

func (s ListEgressProxyIpPoolsResponseBodyEgressProxyIpPools) GoString() string {
	return s.String()
}

func (s *ListEgressProxyIpPoolsResponseBodyEgressProxyIpPools) GetAddressType() *string {
	return s.AddressType
}

func (s *ListEgressProxyIpPoolsResponseBodyEgressProxyIpPools) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListEgressProxyIpPoolsResponseBodyEgressProxyIpPools) GetDescription() *string {
	return s.Description
}

func (s *ListEgressProxyIpPoolsResponseBodyEgressProxyIpPools) GetEgressProxyGatewayId() *string {
	return s.EgressProxyGatewayId
}

func (s *ListEgressProxyIpPoolsResponseBodyEgressProxyIpPools) GetEgressProxyIpPoolId() *string {
	return s.EgressProxyIpPoolId
}

func (s *ListEgressProxyIpPoolsResponseBodyEgressProxyIpPools) GetEgressProxyIpPoolName() *string {
	return s.EgressProxyIpPoolName
}

func (s *ListEgressProxyIpPoolsResponseBodyEgressProxyIpPools) GetHashFactor() *string {
	return s.HashFactor
}

func (s *ListEgressProxyIpPoolsResponseBodyEgressProxyIpPools) GetHashHeaderName() *string {
	return s.HashHeaderName
}

func (s *ListEgressProxyIpPoolsResponseBodyEgressProxyIpPools) GetScheduler() *string {
	return s.Scheduler
}

func (s *ListEgressProxyIpPoolsResponseBodyEgressProxyIpPools) GetStatus() *string {
	return s.Status
}

func (s *ListEgressProxyIpPoolsResponseBodyEgressProxyIpPools) GetTags() []*ListEgressProxyIpPoolsResponseBodyEgressProxyIpPoolsTags {
	return s.Tags
}

func (s *ListEgressProxyIpPoolsResponseBodyEgressProxyIpPools) SetAddressType(v string) *ListEgressProxyIpPoolsResponseBodyEgressProxyIpPools {
	s.AddressType = &v
	return s
}

func (s *ListEgressProxyIpPoolsResponseBodyEgressProxyIpPools) SetCreateTime(v string) *ListEgressProxyIpPoolsResponseBodyEgressProxyIpPools {
	s.CreateTime = &v
	return s
}

func (s *ListEgressProxyIpPoolsResponseBodyEgressProxyIpPools) SetDescription(v string) *ListEgressProxyIpPoolsResponseBodyEgressProxyIpPools {
	s.Description = &v
	return s
}

func (s *ListEgressProxyIpPoolsResponseBodyEgressProxyIpPools) SetEgressProxyGatewayId(v string) *ListEgressProxyIpPoolsResponseBodyEgressProxyIpPools {
	s.EgressProxyGatewayId = &v
	return s
}

func (s *ListEgressProxyIpPoolsResponseBodyEgressProxyIpPools) SetEgressProxyIpPoolId(v string) *ListEgressProxyIpPoolsResponseBodyEgressProxyIpPools {
	s.EgressProxyIpPoolId = &v
	return s
}

func (s *ListEgressProxyIpPoolsResponseBodyEgressProxyIpPools) SetEgressProxyIpPoolName(v string) *ListEgressProxyIpPoolsResponseBodyEgressProxyIpPools {
	s.EgressProxyIpPoolName = &v
	return s
}

func (s *ListEgressProxyIpPoolsResponseBodyEgressProxyIpPools) SetHashFactor(v string) *ListEgressProxyIpPoolsResponseBodyEgressProxyIpPools {
	s.HashFactor = &v
	return s
}

func (s *ListEgressProxyIpPoolsResponseBodyEgressProxyIpPools) SetHashHeaderName(v string) *ListEgressProxyIpPoolsResponseBodyEgressProxyIpPools {
	s.HashHeaderName = &v
	return s
}

func (s *ListEgressProxyIpPoolsResponseBodyEgressProxyIpPools) SetScheduler(v string) *ListEgressProxyIpPoolsResponseBodyEgressProxyIpPools {
	s.Scheduler = &v
	return s
}

func (s *ListEgressProxyIpPoolsResponseBodyEgressProxyIpPools) SetStatus(v string) *ListEgressProxyIpPoolsResponseBodyEgressProxyIpPools {
	s.Status = &v
	return s
}

func (s *ListEgressProxyIpPoolsResponseBodyEgressProxyIpPools) SetTags(v []*ListEgressProxyIpPoolsResponseBodyEgressProxyIpPoolsTags) *ListEgressProxyIpPoolsResponseBodyEgressProxyIpPools {
	s.Tags = v
	return s
}

func (s *ListEgressProxyIpPoolsResponseBodyEgressProxyIpPools) Validate() error {
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListEgressProxyIpPoolsResponseBodyEgressProxyIpPoolsTags struct {
	// example:
	//
	// Physical_QPS,Thread_Detail,Logic_RT,Physical_RT,Slow_RequestCount
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// 1770104466
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListEgressProxyIpPoolsResponseBodyEgressProxyIpPoolsTags) String() string {
	return dara.Prettify(s)
}

func (s ListEgressProxyIpPoolsResponseBodyEgressProxyIpPoolsTags) GoString() string {
	return s.String()
}

func (s *ListEgressProxyIpPoolsResponseBodyEgressProxyIpPoolsTags) GetKey() *string {
	return s.Key
}

func (s *ListEgressProxyIpPoolsResponseBodyEgressProxyIpPoolsTags) GetValue() *string {
	return s.Value
}

func (s *ListEgressProxyIpPoolsResponseBodyEgressProxyIpPoolsTags) SetKey(v string) *ListEgressProxyIpPoolsResponseBodyEgressProxyIpPoolsTags {
	s.Key = &v
	return s
}

func (s *ListEgressProxyIpPoolsResponseBodyEgressProxyIpPoolsTags) SetValue(v string) *ListEgressProxyIpPoolsResponseBodyEgressProxyIpPoolsTags {
	s.Value = &v
	return s
}

func (s *ListEgressProxyIpPoolsResponseBodyEgressProxyIpPoolsTags) Validate() error {
	return dara.Validate(s)
}
