// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCustinsResourceInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*DescribeCustinsResourceInfoResponseBodyData) *DescribeCustinsResourceInfoResponseBody
	GetData() []*DescribeCustinsResourceInfoResponseBodyData
	SetRequestId(v string) *DescribeCustinsResourceInfoResponseBody
	GetRequestId() *string
}

type DescribeCustinsResourceInfoResponseBody struct {
	// The returned data.
	Data []*DescribeCustinsResourceInfoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// D4D4BE8A-DD46-440A-BFCD-EE31DA81C9DD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeCustinsResourceInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCustinsResourceInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCustinsResourceInfoResponseBody) GetData() []*DescribeCustinsResourceInfoResponseBodyData {
	return s.Data
}

func (s *DescribeCustinsResourceInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCustinsResourceInfoResponseBody) SetData(v []*DescribeCustinsResourceInfoResponseBodyData) *DescribeCustinsResourceInfoResponseBody {
	s.Data = v
	return s
}

func (s *DescribeCustinsResourceInfoResponseBody) SetRequestId(v string) *DescribeCustinsResourceInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCustinsResourceInfoResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeCustinsResourceInfoResponseBodyData struct {
	// The deadline for the CPU adjustment.
	//
	// example:
	//
	// 2023-10-25
	CpuAdjustDeadline *string `json:"CpuAdjustDeadline,omitempty" xml:"CpuAdjustDeadline,omitempty"`
	// The maximum percentage of the system CPU resources that the instance can use.
	//
	// example:
	//
	// 30
	CpuAdjustableMaxRatio *string `json:"CpuAdjustableMaxRatio,omitempty" xml:"CpuAdjustableMaxRatio,omitempty"`
	// The maximum CPU utilization.
	//
	// example:
	//
	// 60
	CpuAdjustableMaxValue *string `json:"CpuAdjustableMaxValue,omitempty" xml:"CpuAdjustableMaxValue,omitempty"`
	// The CPU utilization.
	//
	// example:
	//
	// 10
	CpuIncreaseRatio *string `json:"CpuIncreaseRatio,omitempty" xml:"CpuIncreaseRatio,omitempty"`
	// The CPU utilization. Unit: percentage.
	//
	// example:
	//
	// 20
	CpuIncreaseRatioValue *string `json:"CpuIncreaseRatioValue,omitempty" xml:"CpuIncreaseRatioValue,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// rm-wz92gn1ll9fe5d3a4
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The maximum IOPS.
	//
	// example:
	//
	// 20
	IopsAdjustableMaxValue *string `json:"IopsAdjustableMaxValue,omitempty" xml:"IopsAdjustableMaxValue,omitempty"`
	// The deadline for the adjustment of the maximum number of connections.
	//
	// example:
	//
	// 2023-10-25
	MaxConnAdjustDeadline *string `json:"MaxConnAdjustDeadline,omitempty" xml:"MaxConnAdjustDeadline,omitempty"`
	// The maximum number of concurrent connections.
	//
	// example:
	//
	// 100
	MaxConnAdjustableMaxValue *string `json:"MaxConnAdjustableMaxValue,omitempty" xml:"MaxConnAdjustableMaxValue,omitempty"`
	// The maximum number of concurrent connections.
	//
	// example:
	//
	// 100
	MaxConnIncreaseRatio *string `json:"MaxConnIncreaseRatio,omitempty" xml:"MaxConnIncreaseRatio,omitempty"`
	// The maximum number of concurrent connections.
	//
	// example:
	//
	// 20
	MaxConnIncreaseRatioValue *string `json:"MaxConnIncreaseRatioValue,omitempty" xml:"MaxConnIncreaseRatioValue,omitempty"`
	// The deadline for the adjustment of the maximum IOPS.
	//
	// example:
	//
	// 2023-10-25
	MaxIopsAdjustDeadline *string `json:"MaxIopsAdjustDeadline,omitempty" xml:"MaxIopsAdjustDeadline,omitempty"`
	// The maximum IOPS.
	//
	// example:
	//
	// 100
	MaxIopsIncreaseRatio *string `json:"MaxIopsIncreaseRatio,omitempty" xml:"MaxIopsIncreaseRatio,omitempty"`
	// The maximum IOPS that can be supported by the instance.
	//
	// example:
	//
	// 20
	MaxIopsIncreaseRatioValue *string `json:"MaxIopsIncreaseRatioValue,omitempty" xml:"MaxIopsIncreaseRatioValue,omitempty"`
	// The maximum percentage of the system memory that the instance can use.
	//
	// example:
	//
	// 60
	MemAdjustableMaxRatio *string `json:"MemAdjustableMaxRatio,omitempty" xml:"MemAdjustableMaxRatio,omitempty"`
	// The maximum value of the resources to be evaluated.
	//
	// example:
	//
	// 200
	MemAdjustableMaxValue *string `json:"MemAdjustableMaxValue,omitempty" xml:"MemAdjustableMaxValue,omitempty"`
	// The deadline for the memory adjustment.
	//
	// example:
	//
	// 2023-10-25
	MemoryAdjustDeadline *string `json:"MemoryAdjustDeadline,omitempty" xml:"MemoryAdjustDeadline,omitempty"`
	// The memory increase percentage.
	//
	// example:
	//
	// 2023-10-25
	MemoryIncreaseRatio *string `json:"MemoryIncreaseRatio,omitempty" xml:"MemoryIncreaseRatio,omitempty"`
	// The memory usage. Unit: MB.
	//
	// example:
	//
	// 200
	MemoryIncreaseRatioValue *string `json:"MemoryIncreaseRatioValue,omitempty" xml:"MemoryIncreaseRatioValue,omitempty"`
	// The number of CPUs of the instance.
	//
	// example:
	//
	// 2
	OriginCpu *string `json:"OriginCpu,omitempty" xml:"OriginCpu,omitempty"`
	// The maximum number of concurrent connections.
	//
	// example:
	//
	// 30
	OriginMaxConn *string `json:"OriginMaxConn,omitempty" xml:"OriginMaxConn,omitempty"`
	// The maximum IOPS.
	//
	// example:
	//
	// 20
	OriginMaxIops *string `json:"OriginMaxIops,omitempty" xml:"OriginMaxIops,omitempty"`
	// The actual memory used. Unit: MB.
	//
	// example:
	//
	// 20
	OriginMemory *string `json:"OriginMemory,omitempty" xml:"OriginMemory,omitempty"`
}

func (s DescribeCustinsResourceInfoResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeCustinsResourceInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeCustinsResourceInfoResponseBodyData) GetCpuAdjustDeadline() *string {
	return s.CpuAdjustDeadline
}

func (s *DescribeCustinsResourceInfoResponseBodyData) GetCpuAdjustableMaxRatio() *string {
	return s.CpuAdjustableMaxRatio
}

func (s *DescribeCustinsResourceInfoResponseBodyData) GetCpuAdjustableMaxValue() *string {
	return s.CpuAdjustableMaxValue
}

func (s *DescribeCustinsResourceInfoResponseBodyData) GetCpuIncreaseRatio() *string {
	return s.CpuIncreaseRatio
}

func (s *DescribeCustinsResourceInfoResponseBodyData) GetCpuIncreaseRatioValue() *string {
	return s.CpuIncreaseRatioValue
}

func (s *DescribeCustinsResourceInfoResponseBodyData) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeCustinsResourceInfoResponseBodyData) GetIopsAdjustableMaxValue() *string {
	return s.IopsAdjustableMaxValue
}

func (s *DescribeCustinsResourceInfoResponseBodyData) GetMaxConnAdjustDeadline() *string {
	return s.MaxConnAdjustDeadline
}

func (s *DescribeCustinsResourceInfoResponseBodyData) GetMaxConnAdjustableMaxValue() *string {
	return s.MaxConnAdjustableMaxValue
}

func (s *DescribeCustinsResourceInfoResponseBodyData) GetMaxConnIncreaseRatio() *string {
	return s.MaxConnIncreaseRatio
}

func (s *DescribeCustinsResourceInfoResponseBodyData) GetMaxConnIncreaseRatioValue() *string {
	return s.MaxConnIncreaseRatioValue
}

func (s *DescribeCustinsResourceInfoResponseBodyData) GetMaxIopsAdjustDeadline() *string {
	return s.MaxIopsAdjustDeadline
}

func (s *DescribeCustinsResourceInfoResponseBodyData) GetMaxIopsIncreaseRatio() *string {
	return s.MaxIopsIncreaseRatio
}

func (s *DescribeCustinsResourceInfoResponseBodyData) GetMaxIopsIncreaseRatioValue() *string {
	return s.MaxIopsIncreaseRatioValue
}

func (s *DescribeCustinsResourceInfoResponseBodyData) GetMemAdjustableMaxRatio() *string {
	return s.MemAdjustableMaxRatio
}

func (s *DescribeCustinsResourceInfoResponseBodyData) GetMemAdjustableMaxValue() *string {
	return s.MemAdjustableMaxValue
}

func (s *DescribeCustinsResourceInfoResponseBodyData) GetMemoryAdjustDeadline() *string {
	return s.MemoryAdjustDeadline
}

func (s *DescribeCustinsResourceInfoResponseBodyData) GetMemoryIncreaseRatio() *string {
	return s.MemoryIncreaseRatio
}

func (s *DescribeCustinsResourceInfoResponseBodyData) GetMemoryIncreaseRatioValue() *string {
	return s.MemoryIncreaseRatioValue
}

func (s *DescribeCustinsResourceInfoResponseBodyData) GetOriginCpu() *string {
	return s.OriginCpu
}

func (s *DescribeCustinsResourceInfoResponseBodyData) GetOriginMaxConn() *string {
	return s.OriginMaxConn
}

func (s *DescribeCustinsResourceInfoResponseBodyData) GetOriginMaxIops() *string {
	return s.OriginMaxIops
}

func (s *DescribeCustinsResourceInfoResponseBodyData) GetOriginMemory() *string {
	return s.OriginMemory
}

func (s *DescribeCustinsResourceInfoResponseBodyData) SetCpuAdjustDeadline(v string) *DescribeCustinsResourceInfoResponseBodyData {
	s.CpuAdjustDeadline = &v
	return s
}

func (s *DescribeCustinsResourceInfoResponseBodyData) SetCpuAdjustableMaxRatio(v string) *DescribeCustinsResourceInfoResponseBodyData {
	s.CpuAdjustableMaxRatio = &v
	return s
}

func (s *DescribeCustinsResourceInfoResponseBodyData) SetCpuAdjustableMaxValue(v string) *DescribeCustinsResourceInfoResponseBodyData {
	s.CpuAdjustableMaxValue = &v
	return s
}

func (s *DescribeCustinsResourceInfoResponseBodyData) SetCpuIncreaseRatio(v string) *DescribeCustinsResourceInfoResponseBodyData {
	s.CpuIncreaseRatio = &v
	return s
}

func (s *DescribeCustinsResourceInfoResponseBodyData) SetCpuIncreaseRatioValue(v string) *DescribeCustinsResourceInfoResponseBodyData {
	s.CpuIncreaseRatioValue = &v
	return s
}

func (s *DescribeCustinsResourceInfoResponseBodyData) SetDBInstanceId(v string) *DescribeCustinsResourceInfoResponseBodyData {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeCustinsResourceInfoResponseBodyData) SetIopsAdjustableMaxValue(v string) *DescribeCustinsResourceInfoResponseBodyData {
	s.IopsAdjustableMaxValue = &v
	return s
}

func (s *DescribeCustinsResourceInfoResponseBodyData) SetMaxConnAdjustDeadline(v string) *DescribeCustinsResourceInfoResponseBodyData {
	s.MaxConnAdjustDeadline = &v
	return s
}

func (s *DescribeCustinsResourceInfoResponseBodyData) SetMaxConnAdjustableMaxValue(v string) *DescribeCustinsResourceInfoResponseBodyData {
	s.MaxConnAdjustableMaxValue = &v
	return s
}

func (s *DescribeCustinsResourceInfoResponseBodyData) SetMaxConnIncreaseRatio(v string) *DescribeCustinsResourceInfoResponseBodyData {
	s.MaxConnIncreaseRatio = &v
	return s
}

func (s *DescribeCustinsResourceInfoResponseBodyData) SetMaxConnIncreaseRatioValue(v string) *DescribeCustinsResourceInfoResponseBodyData {
	s.MaxConnIncreaseRatioValue = &v
	return s
}

func (s *DescribeCustinsResourceInfoResponseBodyData) SetMaxIopsAdjustDeadline(v string) *DescribeCustinsResourceInfoResponseBodyData {
	s.MaxIopsAdjustDeadline = &v
	return s
}

func (s *DescribeCustinsResourceInfoResponseBodyData) SetMaxIopsIncreaseRatio(v string) *DescribeCustinsResourceInfoResponseBodyData {
	s.MaxIopsIncreaseRatio = &v
	return s
}

func (s *DescribeCustinsResourceInfoResponseBodyData) SetMaxIopsIncreaseRatioValue(v string) *DescribeCustinsResourceInfoResponseBodyData {
	s.MaxIopsIncreaseRatioValue = &v
	return s
}

func (s *DescribeCustinsResourceInfoResponseBodyData) SetMemAdjustableMaxRatio(v string) *DescribeCustinsResourceInfoResponseBodyData {
	s.MemAdjustableMaxRatio = &v
	return s
}

func (s *DescribeCustinsResourceInfoResponseBodyData) SetMemAdjustableMaxValue(v string) *DescribeCustinsResourceInfoResponseBodyData {
	s.MemAdjustableMaxValue = &v
	return s
}

func (s *DescribeCustinsResourceInfoResponseBodyData) SetMemoryAdjustDeadline(v string) *DescribeCustinsResourceInfoResponseBodyData {
	s.MemoryAdjustDeadline = &v
	return s
}

func (s *DescribeCustinsResourceInfoResponseBodyData) SetMemoryIncreaseRatio(v string) *DescribeCustinsResourceInfoResponseBodyData {
	s.MemoryIncreaseRatio = &v
	return s
}

func (s *DescribeCustinsResourceInfoResponseBodyData) SetMemoryIncreaseRatioValue(v string) *DescribeCustinsResourceInfoResponseBodyData {
	s.MemoryIncreaseRatioValue = &v
	return s
}

func (s *DescribeCustinsResourceInfoResponseBodyData) SetOriginCpu(v string) *DescribeCustinsResourceInfoResponseBodyData {
	s.OriginCpu = &v
	return s
}

func (s *DescribeCustinsResourceInfoResponseBodyData) SetOriginMaxConn(v string) *DescribeCustinsResourceInfoResponseBodyData {
	s.OriginMaxConn = &v
	return s
}

func (s *DescribeCustinsResourceInfoResponseBodyData) SetOriginMaxIops(v string) *DescribeCustinsResourceInfoResponseBodyData {
	s.OriginMaxIops = &v
	return s
}

func (s *DescribeCustinsResourceInfoResponseBodyData) SetOriginMemory(v string) *DescribeCustinsResourceInfoResponseBodyData {
	s.OriginMemory = &v
	return s
}

func (s *DescribeCustinsResourceInfoResponseBodyData) Validate() error {
	return dara.Validate(s)
}
