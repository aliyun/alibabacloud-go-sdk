// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSaasDataUnitUsageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *GetSaasDataUnitUsageRequest
	GetEndTime() *string
	SetRegionId(v string) *GetSaasDataUnitUsageRequest
	GetRegionId() *string
	SetServiceId(v string) *GetSaasDataUnitUsageRequest
	GetServiceId() *string
	SetStartTime(v string) *GetSaasDataUnitUsageRequest
	GetStartTime() *string
	SetTimeUnit(v string) *GetSaasDataUnitUsageRequest
	GetTimeUnit() *string
}

type GetSaasDataUnitUsageRequest struct {
	// UTC时间，格式yyyy-MM-ddTHH:mm:ssZ，秒必须为00。必须晚于StartTime且不晚于请求时刻，不自动裁剪查询范围。
	//
	// This parameter is required.
	//
	// example:
	//
	// 2026-10-20T02:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// 未删除的SaaS服务实例ID，不接受普通DBInstanceId替代。
	//
	// This parameter is required.
	//
	// example:
	//
	// agdb-example
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// UTC时间，格式yyyy-MM-ddTHH:mm:ssZ，秒必须为00。不得早于请求时刻前30天；与EndTime组成左闭右开区间，跨度最多7天。
	//
	// This parameter is required.
	//
	// example:
	//
	// 2026-10-20T00:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// 支持hour（UTC小时）或day（UTC日），未传默认hour。仅决定趋势汇总方式，不在响应中回显。
	//
	// example:
	//
	// hour
	TimeUnit *string `json:"TimeUnit,omitempty" xml:"TimeUnit,omitempty"`
}

func (s GetSaasDataUnitUsageRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSaasDataUnitUsageRequest) GoString() string {
	return s.String()
}

func (s *GetSaasDataUnitUsageRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *GetSaasDataUnitUsageRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetSaasDataUnitUsageRequest) GetServiceId() *string {
	return s.ServiceId
}

func (s *GetSaasDataUnitUsageRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *GetSaasDataUnitUsageRequest) GetTimeUnit() *string {
	return s.TimeUnit
}

func (s *GetSaasDataUnitUsageRequest) SetEndTime(v string) *GetSaasDataUnitUsageRequest {
	s.EndTime = &v
	return s
}

func (s *GetSaasDataUnitUsageRequest) SetRegionId(v string) *GetSaasDataUnitUsageRequest {
	s.RegionId = &v
	return s
}

func (s *GetSaasDataUnitUsageRequest) SetServiceId(v string) *GetSaasDataUnitUsageRequest {
	s.ServiceId = &v
	return s
}

func (s *GetSaasDataUnitUsageRequest) SetStartTime(v string) *GetSaasDataUnitUsageRequest {
	s.StartTime = &v
	return s
}

func (s *GetSaasDataUnitUsageRequest) SetTimeUnit(v string) *GetSaasDataUnitUsageRequest {
	s.TimeUnit = &v
	return s
}

func (s *GetSaasDataUnitUsageRequest) Validate() error {
	return dara.Validate(s)
}
