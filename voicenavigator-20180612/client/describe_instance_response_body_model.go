// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAbilityType(v string) *DescribeInstanceResponseBody
	GetAbilityType() *string
	SetApplicableOperations(v []*string) *DescribeInstanceResponseBody
	GetApplicableOperations() []*string
	SetConcurrency(v int64) *DescribeInstanceResponseBody
	GetConcurrency() *int64
	SetDescription(v string) *DescribeInstanceResponseBody
	GetDescription() *string
	SetInstanceId(v string) *DescribeInstanceResponseBody
	GetInstanceId() *string
	SetModifyTime(v int64) *DescribeInstanceResponseBody
	GetModifyTime() *int64
	SetModifyUserName(v string) *DescribeInstanceResponseBody
	GetModifyUserName() *string
	SetName(v string) *DescribeInstanceResponseBody
	GetName() *string
	SetNluServiceParamsJson(v string) *DescribeInstanceResponseBody
	GetNluServiceParamsJson() *string
	SetRequestId(v string) *DescribeInstanceResponseBody
	GetRequestId() *string
	SetStatus(v string) *DescribeInstanceResponseBody
	GetStatus() *string
	SetUnionInstanceId(v string) *DescribeInstanceResponseBody
	GetUnionInstanceId() *string
	SetUnionSource(v string) *DescribeInstanceResponseBody
	GetUnionSource() *string
}

type DescribeInstanceResponseBody struct {
	AbilityType          *string   `json:"AbilityType,omitempty" xml:"AbilityType,omitempty"`
	ApplicableOperations []*string `json:"ApplicableOperations,omitempty" xml:"ApplicableOperations,omitempty" type:"Repeated"`
	// example:
	//
	// 10
	Concurrency *int64  `json:"Concurrency,omitempty" xml:"Concurrency,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// test1_instanceId
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 1532436395329
	ModifyTime           *int64  `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	ModifyUserName       *string `json:"ModifyUserName,omitempty" xml:"ModifyUserName,omitempty"`
	Name                 *string `json:"Name,omitempty" xml:"Name,omitempty"`
	NluServiceParamsJson *string `json:"NluServiceParamsJson,omitempty" xml:"NluServiceParamsJson,omitempty"`
	// example:
	//
	// 14C39896-AE6D-4643-9C9A-E0566B2C2DDD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// Drafted
	Status          *string `json:"Status,omitempty" xml:"Status,omitempty"`
	UnionInstanceId *string `json:"UnionInstanceId,omitempty" xml:"UnionInstanceId,omitempty"`
	UnionSource     *string `json:"UnionSource,omitempty" xml:"UnionSource,omitempty"`
}

func (s DescribeInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceResponseBody) GetAbilityType() *string {
	return s.AbilityType
}

func (s *DescribeInstanceResponseBody) GetApplicableOperations() []*string {
	return s.ApplicableOperations
}

func (s *DescribeInstanceResponseBody) GetConcurrency() *int64 {
	return s.Concurrency
}

func (s *DescribeInstanceResponseBody) GetDescription() *string {
	return s.Description
}

func (s *DescribeInstanceResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeInstanceResponseBody) GetModifyTime() *int64 {
	return s.ModifyTime
}

func (s *DescribeInstanceResponseBody) GetModifyUserName() *string {
	return s.ModifyUserName
}

func (s *DescribeInstanceResponseBody) GetName() *string {
	return s.Name
}

func (s *DescribeInstanceResponseBody) GetNluServiceParamsJson() *string {
	return s.NluServiceParamsJson
}

func (s *DescribeInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeInstanceResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DescribeInstanceResponseBody) GetUnionInstanceId() *string {
	return s.UnionInstanceId
}

func (s *DescribeInstanceResponseBody) GetUnionSource() *string {
	return s.UnionSource
}

func (s *DescribeInstanceResponseBody) SetAbilityType(v string) *DescribeInstanceResponseBody {
	s.AbilityType = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetApplicableOperations(v []*string) *DescribeInstanceResponseBody {
	s.ApplicableOperations = v
	return s
}

func (s *DescribeInstanceResponseBody) SetConcurrency(v int64) *DescribeInstanceResponseBody {
	s.Concurrency = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetDescription(v string) *DescribeInstanceResponseBody {
	s.Description = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetInstanceId(v string) *DescribeInstanceResponseBody {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetModifyTime(v int64) *DescribeInstanceResponseBody {
	s.ModifyTime = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetModifyUserName(v string) *DescribeInstanceResponseBody {
	s.ModifyUserName = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetName(v string) *DescribeInstanceResponseBody {
	s.Name = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetNluServiceParamsJson(v string) *DescribeInstanceResponseBody {
	s.NluServiceParamsJson = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetRequestId(v string) *DescribeInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetStatus(v string) *DescribeInstanceResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetUnionInstanceId(v string) *DescribeInstanceResponseBody {
	s.UnionInstanceId = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetUnionSource(v string) *DescribeInstanceResponseBody {
	s.UnionSource = &v
	return s
}

func (s *DescribeInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
