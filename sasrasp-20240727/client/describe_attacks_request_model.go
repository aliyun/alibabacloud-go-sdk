// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAttacksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentType(v string) *DescribeAttacksRequest
	GetAgentType() *string
	SetApplicationId(v string) *DescribeAttacksRequest
	GetApplicationId() *string
	SetAttackHostId(v string) *DescribeAttacksRequest
	GetAttackHostId() *string
	SetAttackType(v string) *DescribeAttacksRequest
	GetAttackType() *string
	SetAttackUrl(v string) *DescribeAttacksRequest
	GetAttackUrl() *string
	SetEndTimestamp(v int64) *DescribeAttacksRequest
	GetEndTimestamp() *int64
	SetHandleStatus(v int32) *DescribeAttacksRequest
	GetHandleStatus() *int32
	SetHandlerType(v string) *DescribeAttacksRequest
	GetHandlerType() *string
	SetHostname(v string) *DescribeAttacksRequest
	GetHostname() *string
	SetIp(v string) *DescribeAttacksRequest
	GetIp() *string
	SetLang(v string) *DescribeAttacksRequest
	GetLang() *string
	SetPageNumber(v int64) *DescribeAttacksRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribeAttacksRequest
	GetPageSize() *int64
	SetPid(v string) *DescribeAttacksRequest
	GetPid() *string
	SetRaspType(v string) *DescribeAttacksRequest
	GetRaspType() *string
	SetRegion(v string) *DescribeAttacksRequest
	GetRegion() *string
	SetRemote(v string) *DescribeAttacksRequest
	GetRemote() *string
	SetSeverity(v string) *DescribeAttacksRequest
	GetSeverity() *string
	SetStartTimestamp(v int64) *DescribeAttacksRequest
	GetStartTimestamp() *int64
	SetUnionId(v string) *DescribeAttacksRequest
	GetUnionId() *string
}

type DescribeAttacksRequest struct {
	// example:
	//
	// sas
	AgentType *string `json:"AgentType,omitempty" xml:"AgentType,omitempty"`
	// example:
	//
	// 67e283ee866f097cf07d****
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// example:
	//
	// 127.0.0.1
	AttackHostId *string `json:"AttackHostId,omitempty" xml:"AttackHostId,omitempty"`
	// example:
	//
	// sql
	AttackType *string `json:"AttackType,omitempty" xml:"AttackType,omitempty"`
	// example:
	//
	// http://aliyun.com
	AttackUrl *string `json:"AttackUrl,omitempty" xml:"AttackUrl,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1737216000000
	EndTimestamp *int64 `json:"EndTimestamp,omitempty" xml:"EndTimestamp,omitempty"`
	HandleStatus *int32 `json:"HandleStatus,omitempty" xml:"HandleStatus,omitempty"`
	// example:
	//
	// block
	HandlerType *string `json:"HandlerType,omitempty" xml:"HandlerType,omitempty"`
	// example:
	//
	// lshm-sec-waf-new-38
	Hostname *string `json:"Hostname,omitempty" xml:"Hostname,omitempty"`
	// example:
	//
	// 127.0.0.1
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// 4
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Pid      *string `json:"Pid,omitempty" xml:"Pid,omitempty"`
	// example:
	//
	// java
	RaspType *string `json:"RaspType,omitempty" xml:"RaspType,omitempty"`
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	Remote *string `json:"Remote,omitempty" xml:"Remote,omitempty"`
	// example:
	//
	// high
	Severity *string `json:"Severity,omitempty" xml:"Severity,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1727281449756
	StartTimestamp *int64 `json:"StartTimestamp,omitempty" xml:"StartTimestamp,omitempty"`
	// example:
	//
	// 2d14556b77cf1bf7c696e010aaa*****
	UnionId *string `json:"UnionId,omitempty" xml:"UnionId,omitempty"`
}

func (s DescribeAttacksRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAttacksRequest) GoString() string {
	return s.String()
}

func (s *DescribeAttacksRequest) GetAgentType() *string {
	return s.AgentType
}

func (s *DescribeAttacksRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *DescribeAttacksRequest) GetAttackHostId() *string {
	return s.AttackHostId
}

func (s *DescribeAttacksRequest) GetAttackType() *string {
	return s.AttackType
}

func (s *DescribeAttacksRequest) GetAttackUrl() *string {
	return s.AttackUrl
}

func (s *DescribeAttacksRequest) GetEndTimestamp() *int64 {
	return s.EndTimestamp
}

func (s *DescribeAttacksRequest) GetHandleStatus() *int32 {
	return s.HandleStatus
}

func (s *DescribeAttacksRequest) GetHandlerType() *string {
	return s.HandlerType
}

func (s *DescribeAttacksRequest) GetHostname() *string {
	return s.Hostname
}

func (s *DescribeAttacksRequest) GetIp() *string {
	return s.Ip
}

func (s *DescribeAttacksRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeAttacksRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeAttacksRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeAttacksRequest) GetPid() *string {
	return s.Pid
}

func (s *DescribeAttacksRequest) GetRaspType() *string {
	return s.RaspType
}

func (s *DescribeAttacksRequest) GetRegion() *string {
	return s.Region
}

func (s *DescribeAttacksRequest) GetRemote() *string {
	return s.Remote
}

func (s *DescribeAttacksRequest) GetSeverity() *string {
	return s.Severity
}

func (s *DescribeAttacksRequest) GetStartTimestamp() *int64 {
	return s.StartTimestamp
}

func (s *DescribeAttacksRequest) GetUnionId() *string {
	return s.UnionId
}

func (s *DescribeAttacksRequest) SetAgentType(v string) *DescribeAttacksRequest {
	s.AgentType = &v
	return s
}

func (s *DescribeAttacksRequest) SetApplicationId(v string) *DescribeAttacksRequest {
	s.ApplicationId = &v
	return s
}

func (s *DescribeAttacksRequest) SetAttackHostId(v string) *DescribeAttacksRequest {
	s.AttackHostId = &v
	return s
}

func (s *DescribeAttacksRequest) SetAttackType(v string) *DescribeAttacksRequest {
	s.AttackType = &v
	return s
}

func (s *DescribeAttacksRequest) SetAttackUrl(v string) *DescribeAttacksRequest {
	s.AttackUrl = &v
	return s
}

func (s *DescribeAttacksRequest) SetEndTimestamp(v int64) *DescribeAttacksRequest {
	s.EndTimestamp = &v
	return s
}

func (s *DescribeAttacksRequest) SetHandleStatus(v int32) *DescribeAttacksRequest {
	s.HandleStatus = &v
	return s
}

func (s *DescribeAttacksRequest) SetHandlerType(v string) *DescribeAttacksRequest {
	s.HandlerType = &v
	return s
}

func (s *DescribeAttacksRequest) SetHostname(v string) *DescribeAttacksRequest {
	s.Hostname = &v
	return s
}

func (s *DescribeAttacksRequest) SetIp(v string) *DescribeAttacksRequest {
	s.Ip = &v
	return s
}

func (s *DescribeAttacksRequest) SetLang(v string) *DescribeAttacksRequest {
	s.Lang = &v
	return s
}

func (s *DescribeAttacksRequest) SetPageNumber(v int64) *DescribeAttacksRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeAttacksRequest) SetPageSize(v int64) *DescribeAttacksRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeAttacksRequest) SetPid(v string) *DescribeAttacksRequest {
	s.Pid = &v
	return s
}

func (s *DescribeAttacksRequest) SetRaspType(v string) *DescribeAttacksRequest {
	s.RaspType = &v
	return s
}

func (s *DescribeAttacksRequest) SetRegion(v string) *DescribeAttacksRequest {
	s.Region = &v
	return s
}

func (s *DescribeAttacksRequest) SetRemote(v string) *DescribeAttacksRequest {
	s.Remote = &v
	return s
}

func (s *DescribeAttacksRequest) SetSeverity(v string) *DescribeAttacksRequest {
	s.Severity = &v
	return s
}

func (s *DescribeAttacksRequest) SetStartTimestamp(v int64) *DescribeAttacksRequest {
	s.StartTimestamp = &v
	return s
}

func (s *DescribeAttacksRequest) SetUnionId(v string) *DescribeAttacksRequest {
	s.UnionId = &v
	return s
}

func (s *DescribeAttacksRequest) Validate() error {
	return dara.Validate(s)
}
