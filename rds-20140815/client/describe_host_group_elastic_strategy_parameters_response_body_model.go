// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHostGroupElasticStrategyParametersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCpuShar(v int32) *DescribeHostGroupElasticStrategyParametersResponseBody
	GetCpuShar() *int32
	SetCpuZoom(v int32) *DescribeHostGroupElasticStrategyParametersResponseBody
	GetCpuZoom() *int32
	SetIopsZoom(v int32) *DescribeHostGroupElasticStrategyParametersResponseBody
	GetIopsZoom() *int32
	SetMaxConnZoom(v int32) *DescribeHostGroupElasticStrategyParametersResponseBody
	GetMaxConnZoom() *int32
	SetMemoryZoom(v int32) *DescribeHostGroupElasticStrategyParametersResponseBody
	GetMemoryZoom() *int32
	SetRequestId(v string) *DescribeHostGroupElasticStrategyParametersResponseBody
	GetRequestId() *string
}

type DescribeHostGroupElasticStrategyParametersResponseBody struct {
	// The CPU utilization of the instance. Unit: percentage.
	//
	// example:
	//
	// 20
	CpuShar *int32 `json:"CpuShar,omitempty" xml:"CpuShar,omitempty"`
	// The number of CPU cores used by the instance. Unit: cores.
	//
	// example:
	//
	// 4
	CpuZoom *int32 `json:"CpuZoom,omitempty" xml:"CpuZoom,omitempty"`
	// The number of I/O requests.
	//
	// example:
	//
	// 10
	IopsZoom *int32 `json:"IopsZoom,omitempty" xml:"IopsZoom,omitempty"`
	// The maximum number of concurrent connections supported by the instance type.
	//
	// example:
	//
	// 30
	MaxConnZoom *int32 `json:"MaxConnZoom,omitempty" xml:"MaxConnZoom,omitempty"`
	// The total memory size of the instance in the dedicated cluster. Unit: MB.
	//
	// example:
	//
	// 1000
	MemoryZoom *int32 `json:"MemoryZoom,omitempty" xml:"MemoryZoom,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 60F9A12A-16B8-4728-B099-4CA38D32C31C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeHostGroupElasticStrategyParametersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeHostGroupElasticStrategyParametersResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeHostGroupElasticStrategyParametersResponseBody) GetCpuShar() *int32 {
	return s.CpuShar
}

func (s *DescribeHostGroupElasticStrategyParametersResponseBody) GetCpuZoom() *int32 {
	return s.CpuZoom
}

func (s *DescribeHostGroupElasticStrategyParametersResponseBody) GetIopsZoom() *int32 {
	return s.IopsZoom
}

func (s *DescribeHostGroupElasticStrategyParametersResponseBody) GetMaxConnZoom() *int32 {
	return s.MaxConnZoom
}

func (s *DescribeHostGroupElasticStrategyParametersResponseBody) GetMemoryZoom() *int32 {
	return s.MemoryZoom
}

func (s *DescribeHostGroupElasticStrategyParametersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeHostGroupElasticStrategyParametersResponseBody) SetCpuShar(v int32) *DescribeHostGroupElasticStrategyParametersResponseBody {
	s.CpuShar = &v
	return s
}

func (s *DescribeHostGroupElasticStrategyParametersResponseBody) SetCpuZoom(v int32) *DescribeHostGroupElasticStrategyParametersResponseBody {
	s.CpuZoom = &v
	return s
}

func (s *DescribeHostGroupElasticStrategyParametersResponseBody) SetIopsZoom(v int32) *DescribeHostGroupElasticStrategyParametersResponseBody {
	s.IopsZoom = &v
	return s
}

func (s *DescribeHostGroupElasticStrategyParametersResponseBody) SetMaxConnZoom(v int32) *DescribeHostGroupElasticStrategyParametersResponseBody {
	s.MaxConnZoom = &v
	return s
}

func (s *DescribeHostGroupElasticStrategyParametersResponseBody) SetMemoryZoom(v int32) *DescribeHostGroupElasticStrategyParametersResponseBody {
	s.MemoryZoom = &v
	return s
}

func (s *DescribeHostGroupElasticStrategyParametersResponseBody) SetRequestId(v string) *DescribeHostGroupElasticStrategyParametersResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeHostGroupElasticStrategyParametersResponseBody) Validate() error {
	return dara.Validate(s)
}
