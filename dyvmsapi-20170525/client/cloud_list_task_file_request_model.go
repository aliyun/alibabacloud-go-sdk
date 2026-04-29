// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudListTaskFileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *CloudListTaskFileRequest
	GetEndTime() *string
	SetEnterpriseId(v int64) *CloudListTaskFileRequest
	GetEnterpriseId() *int64
	SetLimit(v int64) *CloudListTaskFileRequest
	GetLimit() *int64
	SetOwnerId(v int64) *CloudListTaskFileRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *CloudListTaskFileRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CloudListTaskFileRequest
	GetResourceOwnerId() *int64
	SetStart(v int64) *CloudListTaskFileRequest
	GetStart() *int64
	SetStartTime(v string) *CloudListTaskFileRequest
	GetStartTime() *string
	SetStatus(v int64) *CloudListTaskFileRequest
	GetStatus() *int64
	SetTaskId(v int64) *CloudListTaskFileRequest
	GetTaskId() *int64
}

type CloudListTaskFileRequest struct {
	// 说明：统计日期的结束时间，格式：yyyy-MM-dd HH:mm:ss, 最大查询范围1个月
	//
	// example:
	//
	// 2026-05-20 10:00:00
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// 呼叫中心 id
	//
	// This parameter is required.
	//
	// example:
	//
	// 7000002
	EnterpriseId *int64 `json:"EnterpriseId,omitempty" xml:"EnterpriseId,omitempty"`
	// 查询条数；取值：最大不能超过5000，默认10
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
	// 说明：统计日期的开始时间，格式：yyyy-MM-dd HH:mm:ss
	//
	// example:
	//
	// 2026-04-20 10:00:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// 批次状态；说明：0：未导入，1：导入完成 ，2：加载到缓存，3：呼叫完
	//
	// example:
	//
	// 0
	Status *int64 `json:"Status,omitempty" xml:"Status,omitempty"`
	// 任务Id
	//
	// This parameter is required.
	//
	// example:
	//
	// 27
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CloudListTaskFileRequest) String() string {
	return dara.Prettify(s)
}

func (s CloudListTaskFileRequest) GoString() string {
	return s.String()
}

func (s *CloudListTaskFileRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *CloudListTaskFileRequest) GetEnterpriseId() *int64 {
	return s.EnterpriseId
}

func (s *CloudListTaskFileRequest) GetLimit() *int64 {
	return s.Limit
}

func (s *CloudListTaskFileRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CloudListTaskFileRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CloudListTaskFileRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CloudListTaskFileRequest) GetStart() *int64 {
	return s.Start
}

func (s *CloudListTaskFileRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *CloudListTaskFileRequest) GetStatus() *int64 {
	return s.Status
}

func (s *CloudListTaskFileRequest) GetTaskId() *int64 {
	return s.TaskId
}

func (s *CloudListTaskFileRequest) SetEndTime(v string) *CloudListTaskFileRequest {
	s.EndTime = &v
	return s
}

func (s *CloudListTaskFileRequest) SetEnterpriseId(v int64) *CloudListTaskFileRequest {
	s.EnterpriseId = &v
	return s
}

func (s *CloudListTaskFileRequest) SetLimit(v int64) *CloudListTaskFileRequest {
	s.Limit = &v
	return s
}

func (s *CloudListTaskFileRequest) SetOwnerId(v int64) *CloudListTaskFileRequest {
	s.OwnerId = &v
	return s
}

func (s *CloudListTaskFileRequest) SetResourceOwnerAccount(v string) *CloudListTaskFileRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CloudListTaskFileRequest) SetResourceOwnerId(v int64) *CloudListTaskFileRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CloudListTaskFileRequest) SetStart(v int64) *CloudListTaskFileRequest {
	s.Start = &v
	return s
}

func (s *CloudListTaskFileRequest) SetStartTime(v string) *CloudListTaskFileRequest {
	s.StartTime = &v
	return s
}

func (s *CloudListTaskFileRequest) SetStatus(v int64) *CloudListTaskFileRequest {
	s.Status = &v
	return s
}

func (s *CloudListTaskFileRequest) SetTaskId(v int64) *CloudListTaskFileRequest {
	s.TaskId = &v
	return s
}

func (s *CloudListTaskFileRequest) Validate() error {
	return dara.Validate(s)
}
