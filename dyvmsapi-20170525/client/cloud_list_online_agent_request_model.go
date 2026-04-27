// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudListOnlineAgentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCnos(v string) *CloudListOnlineAgentRequest
	GetCnos() *string
	SetEnterpriseId(v int64) *CloudListOnlineAgentRequest
	GetEnterpriseId() *int64
	SetOwnerId(v int64) *CloudListOnlineAgentRequest
	GetOwnerId() *int64
	SetPauseDescription(v string) *CloudListOnlineAgentRequest
	GetPauseDescription() *string
	SetPauseType(v string) *CloudListOnlineAgentRequest
	GetPauseType() *string
	SetQnos(v string) *CloudListOnlineAgentRequest
	GetQnos() *string
	SetResourceOwnerAccount(v string) *CloudListOnlineAgentRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CloudListOnlineAgentRequest
	GetResourceOwnerId() *int64
	SetStateCode(v string) *CloudListOnlineAgentRequest
	GetStateCode() *string
}

type CloudListOnlineAgentRequest struct {
	// 坐席工号，指定座席工号则查询该工号座席的监控信息，可指定多个座席工号，指定多个座席工号使用英文逗号","分隔，不指定则查询所有队列监控信息，不指定则查询该账户下所有座席的监控信息
	//
	// example:
	//
	// 1111,1112
	Cnos *string `json:"Cnos,omitempty" xml:"Cnos,omitempty"`
	// 呼叫中心 id
	//
	// This parameter is required.
	//
	// example:
	//
	// 7000002
	EnterpriseId *int64 `json:"EnterpriseId,omitempty" xml:"EnterpriseId,omitempty"`
	OwnerId      *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// 置忙原因，根据座席置忙原因过滤监控数据，可指定多个置忙原因，指定多个置忙原因使用英文逗号","分隔，不指定则默认查询所有设置状态的座席
	//
	// example:
	//
	// 示例值
	PauseDescription *string `json:"PauseDescription,omitempty" xml:"PauseDescription,omitempty"`
	// 置忙类型，根据座席置忙类型过滤监控数据，可指定多个置忙类型，指定多个置忙类型使用英文逗号","分隔，不指定则默认查询所有设置状态的座席。<br>取值说明1普通，2休息，3 IM，4 强制
	//
	// example:
	//
	// 1
	PauseType *string `json:"PauseType,omitempty" xml:"PauseType,omitempty"`
	// 队列号，指定队列号则查询该队列号所对应队列的监控信息，支持同时查询多个队列号对应队列的监控信息，多个队列号使用英文逗号","分隔，不指定则查询所有队列监控信息
	//
	// example:
	//
	// 123,124
	Qnos                 *string `json:"Qnos,omitempty" xml:"Qnos,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// 坐席状态，s1:空闲，s2:置忙，s3:整理，s4:呼叫中，s5:响铃，s6通话   可传多个状态码，多个值之间以","分隔
	//
	// example:
	//
	// s1
	StateCode *string `json:"StateCode,omitempty" xml:"StateCode,omitempty"`
}

func (s CloudListOnlineAgentRequest) String() string {
	return dara.Prettify(s)
}

func (s CloudListOnlineAgentRequest) GoString() string {
	return s.String()
}

func (s *CloudListOnlineAgentRequest) GetCnos() *string {
	return s.Cnos
}

func (s *CloudListOnlineAgentRequest) GetEnterpriseId() *int64 {
	return s.EnterpriseId
}

func (s *CloudListOnlineAgentRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CloudListOnlineAgentRequest) GetPauseDescription() *string {
	return s.PauseDescription
}

func (s *CloudListOnlineAgentRequest) GetPauseType() *string {
	return s.PauseType
}

func (s *CloudListOnlineAgentRequest) GetQnos() *string {
	return s.Qnos
}

func (s *CloudListOnlineAgentRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CloudListOnlineAgentRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CloudListOnlineAgentRequest) GetStateCode() *string {
	return s.StateCode
}

func (s *CloudListOnlineAgentRequest) SetCnos(v string) *CloudListOnlineAgentRequest {
	s.Cnos = &v
	return s
}

func (s *CloudListOnlineAgentRequest) SetEnterpriseId(v int64) *CloudListOnlineAgentRequest {
	s.EnterpriseId = &v
	return s
}

func (s *CloudListOnlineAgentRequest) SetOwnerId(v int64) *CloudListOnlineAgentRequest {
	s.OwnerId = &v
	return s
}

func (s *CloudListOnlineAgentRequest) SetPauseDescription(v string) *CloudListOnlineAgentRequest {
	s.PauseDescription = &v
	return s
}

func (s *CloudListOnlineAgentRequest) SetPauseType(v string) *CloudListOnlineAgentRequest {
	s.PauseType = &v
	return s
}

func (s *CloudListOnlineAgentRequest) SetQnos(v string) *CloudListOnlineAgentRequest {
	s.Qnos = &v
	return s
}

func (s *CloudListOnlineAgentRequest) SetResourceOwnerAccount(v string) *CloudListOnlineAgentRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CloudListOnlineAgentRequest) SetResourceOwnerId(v int64) *CloudListOnlineAgentRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CloudListOnlineAgentRequest) SetStateCode(v string) *CloudListOnlineAgentRequest {
	s.StateCode = &v
	return s
}

func (s *CloudListOnlineAgentRequest) Validate() error {
	return dara.Validate(s)
}
