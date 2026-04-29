// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudAgentWorkloadReportRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCnos(v string) *CloudAgentWorkloadReportRequest
	GetCnos() *string
	SetEndTime(v string) *CloudAgentWorkloadReportRequest
	GetEndTime() *string
	SetEnterpriseId(v int64) *CloudAgentWorkloadReportRequest
	GetEnterpriseId() *int64
	SetGnos(v string) *CloudAgentWorkloadReportRequest
	GetGnos() *string
	SetLimit(v int64) *CloudAgentWorkloadReportRequest
	GetLimit() *int64
	SetOwnerId(v int64) *CloudAgentWorkloadReportRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *CloudAgentWorkloadReportRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CloudAgentWorkloadReportRequest
	GetResourceOwnerId() *int64
	SetStart(v int64) *CloudAgentWorkloadReportRequest
	GetStart() *int64
	SetStartTime(v string) *CloudAgentWorkloadReportRequest
	GetStartTime() *string
	SetStatisticMethod(v int64) *CloudAgentWorkloadReportRequest
	GetStatisticMethod() *int64
	SetTimeRangeType(v int64) *CloudAgentWorkloadReportRequest
	GetTimeRangeType() *int64
}

type CloudAgentWorkloadReportRequest struct {
	// 座席号；说明：根据座席工号查询指定座席的工作量统计支持按照多个座席工号进行查询，多个座席工号之间使用英文逗号","分隔，默认查询账户下所有座席的工作量统计
	//
	// example:
	//
	// 1111,2222
	Cnos *string `json:"Cnos,omitempty" xml:"Cnos,omitempty"`
	// 说明：统计日期的结束时间，格式：yyyy-MM-dd
	//
	// This parameter is required.
	//
	// example:
	//
	// 2025-06-15
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// 呼叫中心 id
	//
	// This parameter is required.
	//
	// example:
	//
	// 7000002
	EnterpriseId *int64 `json:"EnterpriseId,omitempty" xml:"EnterpriseId,omitempty"`
	// 说明：根据外呼编号查询指定座席的工作量统计,默认查询账户下所有座席的工作量统计
	//
	// example:
	//
	// WH12
	Gnos *string `json:"Gnos,omitempty" xml:"Gnos,omitempty"`
	// 查询条数；取值：最大不能超过1000，默认10
	//
	// example:
	//
	// 10
	Limit                *int64  `json:"Limit,omitempty" xml:"Limit,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// 查询起始位置；取值：大于等于0，默认0
	//
	// example:
	//
	// 0
	Start *int64 `json:"Start,omitempty" xml:"Start,omitempty"`
	// 说明：统计日期的开始时间，格式：yyyy-MM-dd
	//
	// This parameter is required.
	//
	// example:
	//
	// 2025-06-13
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// 统计方法；说明：0：分时1：分日2：汇总 10：队列汇总
	//
	// This parameter is required.
	//
	// example:
	//
	// 2
	StatisticMethod *int64 `json:"StatisticMethod,omitempty" xml:"StatisticMethod,omitempty"`
	// 说明：统计报表类型，1：日报表2：周报表3：月报表4：自定义时间
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	TimeRangeType *int64 `json:"TimeRangeType,omitempty" xml:"TimeRangeType,omitempty"`
}

func (s CloudAgentWorkloadReportRequest) String() string {
	return dara.Prettify(s)
}

func (s CloudAgentWorkloadReportRequest) GoString() string {
	return s.String()
}

func (s *CloudAgentWorkloadReportRequest) GetCnos() *string {
	return s.Cnos
}

func (s *CloudAgentWorkloadReportRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *CloudAgentWorkloadReportRequest) GetEnterpriseId() *int64 {
	return s.EnterpriseId
}

func (s *CloudAgentWorkloadReportRequest) GetGnos() *string {
	return s.Gnos
}

func (s *CloudAgentWorkloadReportRequest) GetLimit() *int64 {
	return s.Limit
}

func (s *CloudAgentWorkloadReportRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CloudAgentWorkloadReportRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CloudAgentWorkloadReportRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CloudAgentWorkloadReportRequest) GetStart() *int64 {
	return s.Start
}

func (s *CloudAgentWorkloadReportRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *CloudAgentWorkloadReportRequest) GetStatisticMethod() *int64 {
	return s.StatisticMethod
}

func (s *CloudAgentWorkloadReportRequest) GetTimeRangeType() *int64 {
	return s.TimeRangeType
}

func (s *CloudAgentWorkloadReportRequest) SetCnos(v string) *CloudAgentWorkloadReportRequest {
	s.Cnos = &v
	return s
}

func (s *CloudAgentWorkloadReportRequest) SetEndTime(v string) *CloudAgentWorkloadReportRequest {
	s.EndTime = &v
	return s
}

func (s *CloudAgentWorkloadReportRequest) SetEnterpriseId(v int64) *CloudAgentWorkloadReportRequest {
	s.EnterpriseId = &v
	return s
}

func (s *CloudAgentWorkloadReportRequest) SetGnos(v string) *CloudAgentWorkloadReportRequest {
	s.Gnos = &v
	return s
}

func (s *CloudAgentWorkloadReportRequest) SetLimit(v int64) *CloudAgentWorkloadReportRequest {
	s.Limit = &v
	return s
}

func (s *CloudAgentWorkloadReportRequest) SetOwnerId(v int64) *CloudAgentWorkloadReportRequest {
	s.OwnerId = &v
	return s
}

func (s *CloudAgentWorkloadReportRequest) SetResourceOwnerAccount(v string) *CloudAgentWorkloadReportRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CloudAgentWorkloadReportRequest) SetResourceOwnerId(v int64) *CloudAgentWorkloadReportRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CloudAgentWorkloadReportRequest) SetStart(v int64) *CloudAgentWorkloadReportRequest {
	s.Start = &v
	return s
}

func (s *CloudAgentWorkloadReportRequest) SetStartTime(v string) *CloudAgentWorkloadReportRequest {
	s.StartTime = &v
	return s
}

func (s *CloudAgentWorkloadReportRequest) SetStatisticMethod(v int64) *CloudAgentWorkloadReportRequest {
	s.StatisticMethod = &v
	return s
}

func (s *CloudAgentWorkloadReportRequest) SetTimeRangeType(v int64) *CloudAgentWorkloadReportRequest {
	s.TimeRangeType = &v
	return s
}

func (s *CloudAgentWorkloadReportRequest) Validate() error {
	return dara.Validate(s)
}
