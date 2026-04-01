// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCurrentModifyOrderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetModifyOrder(v []*DescribeCurrentModifyOrderResponseBodyModifyOrder) *DescribeCurrentModifyOrderResponseBody
	GetModifyOrder() []*DescribeCurrentModifyOrderResponseBodyModifyOrder
	SetRequestId(v string) *DescribeCurrentModifyOrderResponseBody
	GetRequestId() *string
}

type DescribeCurrentModifyOrderResponseBody struct {
	// The specification change order.
	ModifyOrder []*DescribeCurrentModifyOrderResponseBodyModifyOrder `json:"ModifyOrder,omitempty" xml:"ModifyOrder,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// C87415BE-F5AB-55A4-A60E-A0A329EAF2A4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeCurrentModifyOrderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCurrentModifyOrderResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCurrentModifyOrderResponseBody) GetModifyOrder() []*DescribeCurrentModifyOrderResponseBodyModifyOrder {
	return s.ModifyOrder
}

func (s *DescribeCurrentModifyOrderResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCurrentModifyOrderResponseBody) SetModifyOrder(v []*DescribeCurrentModifyOrderResponseBodyModifyOrder) *DescribeCurrentModifyOrderResponseBody {
	s.ModifyOrder = v
	return s
}

func (s *DescribeCurrentModifyOrderResponseBody) SetRequestId(v string) *DescribeCurrentModifyOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCurrentModifyOrderResponseBody) Validate() error {
	if s.ModifyOrder != nil {
		for _, item := range s.ModifyOrder {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeCurrentModifyOrderResponseBodyModifyOrder struct {
	// The instance family of the instance.
	//
	// example:
	//
	// x
	ClassGroup *string `json:"ClassGroup,omitempty" xml:"ClassGroup,omitempty"`
	// The number of CPU cores that are supported by the instance type. Unit: cores.
	//
	// example:
	//
	// 8
	Cpu *string `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// rm-cn-nwy39qeys0003r
	DbInstanceId *string `json:"DbInstanceId,omitempty" xml:"DbInstanceId,omitempty"`
	// The effective time. Valid values:
	//
	// 	- **Immediate**: This is the default value.
	//
	// 	- **MaintainTime**: The effective time is within the maintenance window. For more information, see [ModifyDBInstanceMaintainTime](https://help.aliyun.com/document_detail/610402.html).
	//
	// example:
	//
	// MaintainTime
	EffectiveTime *string `json:"EffectiveTime,omitempty" xml:"EffectiveTime,omitempty"`
	// The description of the instance.
	//
	// example:
	//
	// eyJ2IjoibWV0YS5rOHMuaW8vdjEiLCJydiI6MTg2MjEwOTkwLCJzdGFydCI6InNob3BpZnktdXNlci1jb3JlXHUwMDAwIn0
	Mark *string `json:"Mark,omitempty" xml:"Mark,omitempty"`
	// The memory capacity that is supported by the instance type. Unit: GB.
	//
	// example:
	//
	// 1024
	MemoryClass *string `json:"MemoryClass,omitempty" xml:"MemoryClass,omitempty"`
	// The status of the task.
	//
	// example:
	//
	// Succeed,Scheduled,Running,Cancelling,Canceled,Waiting
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The storage capacity of the instance.
	//
	// example:
	//
	// 20
	Storage *string `json:"Storage,omitempty" xml:"Storage,omitempty"`
	// The new instance type of the instance. Valid values:
	//
	// example:
	//
	// mysql.x2.medium.2c
	TargetDBInstanceClass *string `json:"TargetDBInstanceClass,omitempty" xml:"TargetDBInstanceClass,omitempty"`
}

func (s DescribeCurrentModifyOrderResponseBodyModifyOrder) String() string {
	return dara.Prettify(s)
}

func (s DescribeCurrentModifyOrderResponseBodyModifyOrder) GoString() string {
	return s.String()
}

func (s *DescribeCurrentModifyOrderResponseBodyModifyOrder) GetClassGroup() *string {
	return s.ClassGroup
}

func (s *DescribeCurrentModifyOrderResponseBodyModifyOrder) GetCpu() *string {
	return s.Cpu
}

func (s *DescribeCurrentModifyOrderResponseBodyModifyOrder) GetDbInstanceId() *string {
	return s.DbInstanceId
}

func (s *DescribeCurrentModifyOrderResponseBodyModifyOrder) GetEffectiveTime() *string {
	return s.EffectiveTime
}

func (s *DescribeCurrentModifyOrderResponseBodyModifyOrder) GetMark() *string {
	return s.Mark
}

func (s *DescribeCurrentModifyOrderResponseBodyModifyOrder) GetMemoryClass() *string {
	return s.MemoryClass
}

func (s *DescribeCurrentModifyOrderResponseBodyModifyOrder) GetStatus() *string {
	return s.Status
}

func (s *DescribeCurrentModifyOrderResponseBodyModifyOrder) GetStorage() *string {
	return s.Storage
}

func (s *DescribeCurrentModifyOrderResponseBodyModifyOrder) GetTargetDBInstanceClass() *string {
	return s.TargetDBInstanceClass
}

func (s *DescribeCurrentModifyOrderResponseBodyModifyOrder) SetClassGroup(v string) *DescribeCurrentModifyOrderResponseBodyModifyOrder {
	s.ClassGroup = &v
	return s
}

func (s *DescribeCurrentModifyOrderResponseBodyModifyOrder) SetCpu(v string) *DescribeCurrentModifyOrderResponseBodyModifyOrder {
	s.Cpu = &v
	return s
}

func (s *DescribeCurrentModifyOrderResponseBodyModifyOrder) SetDbInstanceId(v string) *DescribeCurrentModifyOrderResponseBodyModifyOrder {
	s.DbInstanceId = &v
	return s
}

func (s *DescribeCurrentModifyOrderResponseBodyModifyOrder) SetEffectiveTime(v string) *DescribeCurrentModifyOrderResponseBodyModifyOrder {
	s.EffectiveTime = &v
	return s
}

func (s *DescribeCurrentModifyOrderResponseBodyModifyOrder) SetMark(v string) *DescribeCurrentModifyOrderResponseBodyModifyOrder {
	s.Mark = &v
	return s
}

func (s *DescribeCurrentModifyOrderResponseBodyModifyOrder) SetMemoryClass(v string) *DescribeCurrentModifyOrderResponseBodyModifyOrder {
	s.MemoryClass = &v
	return s
}

func (s *DescribeCurrentModifyOrderResponseBodyModifyOrder) SetStatus(v string) *DescribeCurrentModifyOrderResponseBodyModifyOrder {
	s.Status = &v
	return s
}

func (s *DescribeCurrentModifyOrderResponseBodyModifyOrder) SetStorage(v string) *DescribeCurrentModifyOrderResponseBodyModifyOrder {
	s.Storage = &v
	return s
}

func (s *DescribeCurrentModifyOrderResponseBodyModifyOrder) SetTargetDBInstanceClass(v string) *DescribeCurrentModifyOrderResponseBodyModifyOrder {
	s.TargetDBInstanceClass = &v
	return s
}

func (s *DescribeCurrentModifyOrderResponseBodyModifyOrder) Validate() error {
	return dara.Validate(s)
}
