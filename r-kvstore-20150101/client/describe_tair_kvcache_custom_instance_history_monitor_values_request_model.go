// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTairKVCacheCustomInstanceHistoryMonitorValuesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *DescribeTairKVCacheCustomInstanceHistoryMonitorValuesRequest
	GetEndTime() *string
	SetExpress(v string) *DescribeTairKVCacheCustomInstanceHistoryMonitorValuesRequest
	GetExpress() *string
	SetInstanceId(v string) *DescribeTairKVCacheCustomInstanceHistoryMonitorValuesRequest
	GetInstanceId() *string
	SetLength(v string) *DescribeTairKVCacheCustomInstanceHistoryMonitorValuesRequest
	GetLength() *string
	SetMetricName(v string) *DescribeTairKVCacheCustomInstanceHistoryMonitorValuesRequest
	GetMetricName() *string
	SetOwnerAccount(v string) *DescribeTairKVCacheCustomInstanceHistoryMonitorValuesRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeTairKVCacheCustomInstanceHistoryMonitorValuesRequest
	GetOwnerId() *int64
	SetPeriod(v string) *DescribeTairKVCacheCustomInstanceHistoryMonitorValuesRequest
	GetPeriod() *string
	SetResourceOwnerAccount(v string) *DescribeTairKVCacheCustomInstanceHistoryMonitorValuesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeTairKVCacheCustomInstanceHistoryMonitorValuesRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *DescribeTairKVCacheCustomInstanceHistoryMonitorValuesRequest
	GetSecurityToken() *string
	SetStartTime(v string) *DescribeTairKVCacheCustomInstanceHistoryMonitorValuesRequest
	GetStartTime() *string
}

type DescribeTairKVCacheCustomInstanceHistoryMonitorValuesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 2024-09-20T00:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// {\\"extend\\":{\\"workers\\":\\"avg_dispatchers\\"}}
	Express *string `json:"Express,omitempty" xml:"Express,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// tc-bp1zxszhcgatnx****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 1000
	Length *string `json:"Length,omitempty" xml:"Length,omitempty"`
	// example:
	//
	// CPUUtilization
	MetricName   *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// 60
	Period               *string `json:"Period,omitempty" xml:"Period,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2024-09-05T08:49:27Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeTairKVCacheCustomInstanceHistoryMonitorValuesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeTairKVCacheCustomInstanceHistoryMonitorValuesRequest) GoString() string {
	return s.String()
}

func (s *DescribeTairKVCacheCustomInstanceHistoryMonitorValuesRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeTairKVCacheCustomInstanceHistoryMonitorValuesRequest) GetExpress() *string {
	return s.Express
}

func (s *DescribeTairKVCacheCustomInstanceHistoryMonitorValuesRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeTairKVCacheCustomInstanceHistoryMonitorValuesRequest) GetLength() *string {
	return s.Length
}

func (s *DescribeTairKVCacheCustomInstanceHistoryMonitorValuesRequest) GetMetricName() *string {
	return s.MetricName
}

func (s *DescribeTairKVCacheCustomInstanceHistoryMonitorValuesRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeTairKVCacheCustomInstanceHistoryMonitorValuesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeTairKVCacheCustomInstanceHistoryMonitorValuesRequest) GetPeriod() *string {
	return s.Period
}

func (s *DescribeTairKVCacheCustomInstanceHistoryMonitorValuesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeTairKVCacheCustomInstanceHistoryMonitorValuesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeTairKVCacheCustomInstanceHistoryMonitorValuesRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeTairKVCacheCustomInstanceHistoryMonitorValuesRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeTairKVCacheCustomInstanceHistoryMonitorValuesRequest) SetEndTime(v string) *DescribeTairKVCacheCustomInstanceHistoryMonitorValuesRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeTairKVCacheCustomInstanceHistoryMonitorValuesRequest) SetExpress(v string) *DescribeTairKVCacheCustomInstanceHistoryMonitorValuesRequest {
	s.Express = &v
	return s
}

func (s *DescribeTairKVCacheCustomInstanceHistoryMonitorValuesRequest) SetInstanceId(v string) *DescribeTairKVCacheCustomInstanceHistoryMonitorValuesRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeTairKVCacheCustomInstanceHistoryMonitorValuesRequest) SetLength(v string) *DescribeTairKVCacheCustomInstanceHistoryMonitorValuesRequest {
	s.Length = &v
	return s
}

func (s *DescribeTairKVCacheCustomInstanceHistoryMonitorValuesRequest) SetMetricName(v string) *DescribeTairKVCacheCustomInstanceHistoryMonitorValuesRequest {
	s.MetricName = &v
	return s
}

func (s *DescribeTairKVCacheCustomInstanceHistoryMonitorValuesRequest) SetOwnerAccount(v string) *DescribeTairKVCacheCustomInstanceHistoryMonitorValuesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeTairKVCacheCustomInstanceHistoryMonitorValuesRequest) SetOwnerId(v int64) *DescribeTairKVCacheCustomInstanceHistoryMonitorValuesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeTairKVCacheCustomInstanceHistoryMonitorValuesRequest) SetPeriod(v string) *DescribeTairKVCacheCustomInstanceHistoryMonitorValuesRequest {
	s.Period = &v
	return s
}

func (s *DescribeTairKVCacheCustomInstanceHistoryMonitorValuesRequest) SetResourceOwnerAccount(v string) *DescribeTairKVCacheCustomInstanceHistoryMonitorValuesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeTairKVCacheCustomInstanceHistoryMonitorValuesRequest) SetResourceOwnerId(v int64) *DescribeTairKVCacheCustomInstanceHistoryMonitorValuesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeTairKVCacheCustomInstanceHistoryMonitorValuesRequest) SetSecurityToken(v string) *DescribeTairKVCacheCustomInstanceHistoryMonitorValuesRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeTairKVCacheCustomInstanceHistoryMonitorValuesRequest) SetStartTime(v string) *DescribeTairKVCacheCustomInstanceHistoryMonitorValuesRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeTairKVCacheCustomInstanceHistoryMonitorValuesRequest) Validate() error {
	return dara.Validate(s)
}
