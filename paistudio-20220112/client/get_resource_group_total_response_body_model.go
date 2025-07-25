// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetResourceGroupTotalResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetTotalCPU(v int32) *GetResourceGroupTotalResponseBody
	GetTotalCPU() *int32
	SetTotalGPU(v int32) *GetResourceGroupTotalResponseBody
	GetTotalGPU() *int32
	SetTotalGPUInfos(v []*GPUInfo) *GetResourceGroupTotalResponseBody
	GetTotalGPUInfos() []*GPUInfo
	SetTotalMemory(v int32) *GetResourceGroupTotalResponseBody
	GetTotalMemory() *int32
}

type GetResourceGroupTotalResponseBody struct {
	// example:
	//
	// 100
	TotalCPU *int32 `json:"totalCPU,omitempty" xml:"totalCPU,omitempty"`
	// example:
	//
	// 24
	TotalGPU      *int32     `json:"totalGPU,omitempty" xml:"totalGPU,omitempty"`
	TotalGPUInfos []*GPUInfo `json:"totalGPUInfos,omitempty" xml:"totalGPUInfos,omitempty" type:"Repeated"`
	// example:
	//
	// 300
	TotalMemory *int32 `json:"totalMemory,omitempty" xml:"totalMemory,omitempty"`
}

func (s GetResourceGroupTotalResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetResourceGroupTotalResponseBody) GoString() string {
	return s.String()
}

func (s *GetResourceGroupTotalResponseBody) GetTotalCPU() *int32 {
	return s.TotalCPU
}

func (s *GetResourceGroupTotalResponseBody) GetTotalGPU() *int32 {
	return s.TotalGPU
}

func (s *GetResourceGroupTotalResponseBody) GetTotalGPUInfos() []*GPUInfo {
	return s.TotalGPUInfos
}

func (s *GetResourceGroupTotalResponseBody) GetTotalMemory() *int32 {
	return s.TotalMemory
}

func (s *GetResourceGroupTotalResponseBody) SetTotalCPU(v int32) *GetResourceGroupTotalResponseBody {
	s.TotalCPU = &v
	return s
}

func (s *GetResourceGroupTotalResponseBody) SetTotalGPU(v int32) *GetResourceGroupTotalResponseBody {
	s.TotalGPU = &v
	return s
}

func (s *GetResourceGroupTotalResponseBody) SetTotalGPUInfos(v []*GPUInfo) *GetResourceGroupTotalResponseBody {
	s.TotalGPUInfos = v
	return s
}

func (s *GetResourceGroupTotalResponseBody) SetTotalMemory(v int32) *GetResourceGroupTotalResponseBody {
	s.TotalMemory = &v
	return s
}

func (s *GetResourceGroupTotalResponseBody) Validate() error {
	return dara.Validate(s)
}
