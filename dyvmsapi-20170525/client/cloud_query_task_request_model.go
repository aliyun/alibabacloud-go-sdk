// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudQueryTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoStart(v int64) *CloudQueryTaskRequest
	GetAutoStart() *int64
	SetAutoStop(v int64) *CloudQueryTaskRequest
	GetAutoStop() *int64
	SetEndTime(v string) *CloudQueryTaskRequest
	GetEndTime() *string
	SetEnterpriseId(v int64) *CloudQueryTaskRequest
	GetEnterpriseId() *int64
	SetLimit(v int64) *CloudQueryTaskRequest
	GetLimit() *int64
	SetName(v string) *CloudQueryTaskRequest
	GetName() *string
	SetOwnerId(v int64) *CloudQueryTaskRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *CloudQueryTaskRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CloudQueryTaskRequest
	GetResourceOwnerId() *int64
	SetStart(v int64) *CloudQueryTaskRequest
	GetStart() *int64
	SetStartTime(v string) *CloudQueryTaskRequest
	GetStartTime() *string
	SetStatus(v int64) *CloudQueryTaskRequest
	GetStatus() *int64
	SetTimeType(v int64) *CloudQueryTaskRequest
	GetTimeType() *int64
	SetType(v int64) *CloudQueryTaskRequest
	GetType() *int64
}

type CloudQueryTaskRequest struct {
	// 定时开始；0.关闭 1.开启
	//
	// example:
	//
	// 0
	AutoStart *int64 `json:"AutoStart,omitempty" xml:"AutoStart,omitempty"`
	// 定时结束；0.关闭 1.开启
	//
	// example:
	//
	// 0
	AutoStop *int64 `json:"AutoStop,omitempty" xml:"AutoStop,omitempty"`
	// 查询结束时间；格式说明："2019-10-11 23:59:59"
	//
	// example:
	//
	// 2026-04-20 20:00:00
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// 呼叫中心 id
	//
	// This parameter is required.
	//
	// example:
	//
	// 7000002
	EnterpriseId *int64 `json:"EnterpriseId,omitempty" xml:"EnterpriseId,omitempty"`
	// 需要取出的数据条数；大于0，最大不能超过100，默认10
	//
	// example:
	//
	// 10
	Limit *int64 `json:"Limit,omitempty" xml:"Limit,omitempty"`
	// 任务名称；需进行UTF-8格式的URLEncode编码
	//
	// example:
	//
	// name1
	Name                 *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// 从第几条开始获取；大于等于0，默认0
	//
	// example:
	//
	// 0
	Start *int64 `json:"Start,omitempty" xml:"Start,omitempty"`
	// 查询开始时间；格式说明："2019-10-11 00:00:00"
	//
	// example:
	//
	// 2026-04-20 10:00:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// 任务状态；0.初始 1.运行中 2.暂停 3.结束
	//
	// example:
	//
	// 3
	Status *int64 `json:"Status,omitempty" xml:"Status,omitempty"`
	// 时间过滤条件；1.任务启动时间 2.任务结束时间 3.任务创建时间
	//
	// example:
	//
	// 1
	TimeType *int64 `json:"TimeType,omitempty" xml:"TimeType,omitempty"`
	// 任务类型；1.预测外呼任务 2.自动外呼任务
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	Type *int64 `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CloudQueryTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CloudQueryTaskRequest) GoString() string {
	return s.String()
}

func (s *CloudQueryTaskRequest) GetAutoStart() *int64 {
	return s.AutoStart
}

func (s *CloudQueryTaskRequest) GetAutoStop() *int64 {
	return s.AutoStop
}

func (s *CloudQueryTaskRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *CloudQueryTaskRequest) GetEnterpriseId() *int64 {
	return s.EnterpriseId
}

func (s *CloudQueryTaskRequest) GetLimit() *int64 {
	return s.Limit
}

func (s *CloudQueryTaskRequest) GetName() *string {
	return s.Name
}

func (s *CloudQueryTaskRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CloudQueryTaskRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CloudQueryTaskRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CloudQueryTaskRequest) GetStart() *int64 {
	return s.Start
}

func (s *CloudQueryTaskRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *CloudQueryTaskRequest) GetStatus() *int64 {
	return s.Status
}

func (s *CloudQueryTaskRequest) GetTimeType() *int64 {
	return s.TimeType
}

func (s *CloudQueryTaskRequest) GetType() *int64 {
	return s.Type
}

func (s *CloudQueryTaskRequest) SetAutoStart(v int64) *CloudQueryTaskRequest {
	s.AutoStart = &v
	return s
}

func (s *CloudQueryTaskRequest) SetAutoStop(v int64) *CloudQueryTaskRequest {
	s.AutoStop = &v
	return s
}

func (s *CloudQueryTaskRequest) SetEndTime(v string) *CloudQueryTaskRequest {
	s.EndTime = &v
	return s
}

func (s *CloudQueryTaskRequest) SetEnterpriseId(v int64) *CloudQueryTaskRequest {
	s.EnterpriseId = &v
	return s
}

func (s *CloudQueryTaskRequest) SetLimit(v int64) *CloudQueryTaskRequest {
	s.Limit = &v
	return s
}

func (s *CloudQueryTaskRequest) SetName(v string) *CloudQueryTaskRequest {
	s.Name = &v
	return s
}

func (s *CloudQueryTaskRequest) SetOwnerId(v int64) *CloudQueryTaskRequest {
	s.OwnerId = &v
	return s
}

func (s *CloudQueryTaskRequest) SetResourceOwnerAccount(v string) *CloudQueryTaskRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CloudQueryTaskRequest) SetResourceOwnerId(v int64) *CloudQueryTaskRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CloudQueryTaskRequest) SetStart(v int64) *CloudQueryTaskRequest {
	s.Start = &v
	return s
}

func (s *CloudQueryTaskRequest) SetStartTime(v string) *CloudQueryTaskRequest {
	s.StartTime = &v
	return s
}

func (s *CloudQueryTaskRequest) SetStatus(v int64) *CloudQueryTaskRequest {
	s.Status = &v
	return s
}

func (s *CloudQueryTaskRequest) SetTimeType(v int64) *CloudQueryTaskRequest {
	s.TimeType = &v
	return s
}

func (s *CloudQueryTaskRequest) SetType(v int64) *CloudQueryTaskRequest {
	s.Type = &v
	return s
}

func (s *CloudQueryTaskRequest) Validate() error {
	return dara.Validate(s)
}
