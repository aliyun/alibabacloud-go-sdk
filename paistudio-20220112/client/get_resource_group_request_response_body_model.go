// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetResourceGroupRequestResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestCPU(v int32) *GetResourceGroupRequestResponseBody
	GetRequestCPU() *int32
	SetRequestGPU(v int32) *GetResourceGroupRequestResponseBody
	GetRequestGPU() *int32
	SetRequestGPUInfos(v []*GPUInfo) *GetResourceGroupRequestResponseBody
	GetRequestGPUInfos() []*GPUInfo
	SetRequestMemory(v int32) *GetResourceGroupRequestResponseBody
	GetRequestMemory() *int32
}

type GetResourceGroupRequestResponseBody struct {
	// example:
	//
	// 1
	RequestCPU *int32 `json:"requestCPU,omitempty" xml:"requestCPU,omitempty"`
	// example:
	//
	// 8
	RequestGPU      *int32     `json:"requestGPU,omitempty" xml:"requestGPU,omitempty"`
	RequestGPUInfos []*GPUInfo `json:"requestGPUInfos,omitempty" xml:"requestGPUInfos,omitempty" type:"Repeated"`
	// example:
	//
	// 2
	RequestMemory *int32 `json:"requestMemory,omitempty" xml:"requestMemory,omitempty"`
}

func (s GetResourceGroupRequestResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetResourceGroupRequestResponseBody) GoString() string {
	return s.String()
}

func (s *GetResourceGroupRequestResponseBody) GetRequestCPU() *int32 {
	return s.RequestCPU
}

func (s *GetResourceGroupRequestResponseBody) GetRequestGPU() *int32 {
	return s.RequestGPU
}

func (s *GetResourceGroupRequestResponseBody) GetRequestGPUInfos() []*GPUInfo {
	return s.RequestGPUInfos
}

func (s *GetResourceGroupRequestResponseBody) GetRequestMemory() *int32 {
	return s.RequestMemory
}

func (s *GetResourceGroupRequestResponseBody) SetRequestCPU(v int32) *GetResourceGroupRequestResponseBody {
	s.RequestCPU = &v
	return s
}

func (s *GetResourceGroupRequestResponseBody) SetRequestGPU(v int32) *GetResourceGroupRequestResponseBody {
	s.RequestGPU = &v
	return s
}

func (s *GetResourceGroupRequestResponseBody) SetRequestGPUInfos(v []*GPUInfo) *GetResourceGroupRequestResponseBody {
	s.RequestGPUInfos = v
	return s
}

func (s *GetResourceGroupRequestResponseBody) SetRequestMemory(v int32) *GetResourceGroupRequestResponseBody {
	s.RequestMemory = &v
	return s
}

func (s *GetResourceGroupRequestResponseBody) Validate() error {
	return dara.Validate(s)
}
