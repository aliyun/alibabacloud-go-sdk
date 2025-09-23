// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstances(v []*DescribeInstancesResponseBodyInstances) *DescribeInstancesResponseBody
	GetInstances() []*DescribeInstancesResponseBodyInstances
	SetRequestId(v string) *DescribeInstancesResponseBody
	GetRequestId() *string
	SetTotal(v int64) *DescribeInstancesResponseBody
	GetTotal() *int64
}

type DescribeInstancesResponseBody struct {
	Instances []*DescribeInstancesResponseBodyInstances `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Repeated"`
	// example:
	//
	// CF33B4C3-196E-4015-AADD-5CAD00057B80
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 10
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBody) GetInstances() []*DescribeInstancesResponseBodyInstances {
	return s.Instances
}

func (s *DescribeInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeInstancesResponseBody) GetTotal() *int64 {
	return s.Total
}

func (s *DescribeInstancesResponseBody) SetInstances(v []*DescribeInstancesResponseBodyInstances) *DescribeInstancesResponseBody {
	s.Instances = v
	return s
}

func (s *DescribeInstancesResponseBody) SetRequestId(v string) *DescribeInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstancesResponseBody) SetTotal(v int64) *DescribeInstancesResponseBody {
	s.Total = &v
	return s
}

func (s *DescribeInstancesResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeInstancesResponseBodyInstances struct {
	// example:
	//
	// 0
	DebtStatus *int32 `json:"DebtStatus,omitempty" xml:"DebtStatus,omitempty"`
	// example:
	//
	// 9
	Edition *int32 `json:"Edition,omitempty" xml:"Edition,omitempty"`
	// example:
	//
	// 1
	Enabled *int32 `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// example:
	//
	// 2308402384
	ExpireTime *int64 `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// example:
	//
	// 2308402384
	GmtCreate *int64 `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// example:
	//
	// ddoscoo-cn-XXXXX
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// testRemark
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeInstancesResponseBodyInstances) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstancesResponseBodyInstances) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBodyInstances) GetDebtStatus() *int32 {
	return s.DebtStatus
}

func (s *DescribeInstancesResponseBodyInstances) GetEdition() *int32 {
	return s.Edition
}

func (s *DescribeInstancesResponseBodyInstances) GetEnabled() *int32 {
	return s.Enabled
}

func (s *DescribeInstancesResponseBodyInstances) GetExpireTime() *int64 {
	return s.ExpireTime
}

func (s *DescribeInstancesResponseBodyInstances) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *DescribeInstancesResponseBodyInstances) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeInstancesResponseBodyInstances) GetRemark() *string {
	return s.Remark
}

func (s *DescribeInstancesResponseBodyInstances) GetStatus() *int32 {
	return s.Status
}

func (s *DescribeInstancesResponseBodyInstances) SetDebtStatus(v int32) *DescribeInstancesResponseBodyInstances {
	s.DebtStatus = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetEdition(v int32) *DescribeInstancesResponseBodyInstances {
	s.Edition = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetEnabled(v int32) *DescribeInstancesResponseBodyInstances {
	s.Enabled = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetExpireTime(v int64) *DescribeInstancesResponseBodyInstances {
	s.ExpireTime = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetGmtCreate(v int64) *DescribeInstancesResponseBodyInstances {
	s.GmtCreate = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetInstanceId(v string) *DescribeInstancesResponseBodyInstances {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetRemark(v string) *DescribeInstancesResponseBodyInstances {
	s.Remark = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetStatus(v int32) *DescribeInstancesResponseBodyInstances {
	s.Status = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) Validate() error {
	return dara.Validate(s)
}
