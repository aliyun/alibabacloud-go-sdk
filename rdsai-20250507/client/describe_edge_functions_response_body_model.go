// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEdgeFunctionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEdgeFunctions(v []*DescribeEdgeFunctionsResponseBodyEdgeFunctions) *DescribeEdgeFunctionsResponseBody
	GetEdgeFunctions() []*DescribeEdgeFunctionsResponseBodyEdgeFunctions
	SetInstanceName(v string) *DescribeEdgeFunctionsResponseBody
	GetInstanceName() *string
	SetRequestId(v string) *DescribeEdgeFunctionsResponseBody
	GetRequestId() *string
}

type DescribeEdgeFunctionsResponseBody struct {
	EdgeFunctions []*DescribeEdgeFunctionsResponseBodyEdgeFunctions `json:"EdgeFunctions,omitempty" xml:"EdgeFunctions,omitempty" type:"Repeated"`
	// example:
	//
	// ra-supabase-8moov5lxba****
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// Id of the request
	//
	// example:
	//
	// FE9C65D7-930F-57A5-A207-8C396329****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeEdgeFunctionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeEdgeFunctionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEdgeFunctionsResponseBody) GetEdgeFunctions() []*DescribeEdgeFunctionsResponseBodyEdgeFunctions {
	return s.EdgeFunctions
}

func (s *DescribeEdgeFunctionsResponseBody) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeEdgeFunctionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeEdgeFunctionsResponseBody) SetEdgeFunctions(v []*DescribeEdgeFunctionsResponseBodyEdgeFunctions) *DescribeEdgeFunctionsResponseBody {
	s.EdgeFunctions = v
	return s
}

func (s *DescribeEdgeFunctionsResponseBody) SetInstanceName(v string) *DescribeEdgeFunctionsResponseBody {
	s.InstanceName = &v
	return s
}

func (s *DescribeEdgeFunctionsResponseBody) SetRequestId(v string) *DescribeEdgeFunctionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeEdgeFunctionsResponseBody) Validate() error {
	if s.EdgeFunctions != nil {
		for _, item := range s.EdgeFunctions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeEdgeFunctionsResponseBodyEdgeFunctions struct {
	// example:
	//
	// 1
	Cpu *string `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// example:
	//
	// 2021-11-12T21:35:03
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// example:
	//
	// ef-****
	EdgeFunctionName *string `json:"EdgeFunctionName,omitempty" xml:"EdgeFunctionName,omitempty"`
	// example:
	//
	// https://fcnext.console.aliyun.com/cn-beijing/functions/fc****
	FunctionUrl *string `json:"FunctionUrl,omitempty" xml:"FunctionUrl,omitempty"`
	// example:
	//
	// 512
	MemorySize *string `json:"MemorySize,omitempty" xml:"MemorySize,omitempty"`
	// example:
	//
	// 2025-05-25 10:22:54 +0800
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// example:
	//
	// custom.debian12-deno-2.5.6
	Runtime *string `json:"Runtime,omitempty" xml:"Runtime,omitempty"`
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// https://fc-bfvmoi****.cn-beijing.fcapp.run
	UrlInternet *string `json:"UrlInternet,omitempty" xml:"UrlInternet,omitempty"`
	// example:
	//
	// https://fc-bfvmoi****.cn-beijing-vpc.fcapp.run
	UrlIntranet *string `json:"UrlIntranet,omitempty" xml:"UrlIntranet,omitempty"`
}

func (s DescribeEdgeFunctionsResponseBodyEdgeFunctions) String() string {
	return dara.Prettify(s)
}

func (s DescribeEdgeFunctionsResponseBodyEdgeFunctions) GoString() string {
	return s.String()
}

func (s *DescribeEdgeFunctionsResponseBodyEdgeFunctions) GetCpu() *string {
	return s.Cpu
}

func (s *DescribeEdgeFunctionsResponseBodyEdgeFunctions) GetCreatedTime() *string {
	return s.CreatedTime
}

func (s *DescribeEdgeFunctionsResponseBodyEdgeFunctions) GetEdgeFunctionName() *string {
	return s.EdgeFunctionName
}

func (s *DescribeEdgeFunctionsResponseBodyEdgeFunctions) GetFunctionUrl() *string {
	return s.FunctionUrl
}

func (s *DescribeEdgeFunctionsResponseBodyEdgeFunctions) GetMemorySize() *string {
	return s.MemorySize
}

func (s *DescribeEdgeFunctionsResponseBodyEdgeFunctions) GetModifiedTime() *string {
	return s.ModifiedTime
}

func (s *DescribeEdgeFunctionsResponseBodyEdgeFunctions) GetRuntime() *string {
	return s.Runtime
}

func (s *DescribeEdgeFunctionsResponseBodyEdgeFunctions) GetStatus() *string {
	return s.Status
}

func (s *DescribeEdgeFunctionsResponseBodyEdgeFunctions) GetUrlInternet() *string {
	return s.UrlInternet
}

func (s *DescribeEdgeFunctionsResponseBodyEdgeFunctions) GetUrlIntranet() *string {
	return s.UrlIntranet
}

func (s *DescribeEdgeFunctionsResponseBodyEdgeFunctions) SetCpu(v string) *DescribeEdgeFunctionsResponseBodyEdgeFunctions {
	s.Cpu = &v
	return s
}

func (s *DescribeEdgeFunctionsResponseBodyEdgeFunctions) SetCreatedTime(v string) *DescribeEdgeFunctionsResponseBodyEdgeFunctions {
	s.CreatedTime = &v
	return s
}

func (s *DescribeEdgeFunctionsResponseBodyEdgeFunctions) SetEdgeFunctionName(v string) *DescribeEdgeFunctionsResponseBodyEdgeFunctions {
	s.EdgeFunctionName = &v
	return s
}

func (s *DescribeEdgeFunctionsResponseBodyEdgeFunctions) SetFunctionUrl(v string) *DescribeEdgeFunctionsResponseBodyEdgeFunctions {
	s.FunctionUrl = &v
	return s
}

func (s *DescribeEdgeFunctionsResponseBodyEdgeFunctions) SetMemorySize(v string) *DescribeEdgeFunctionsResponseBodyEdgeFunctions {
	s.MemorySize = &v
	return s
}

func (s *DescribeEdgeFunctionsResponseBodyEdgeFunctions) SetModifiedTime(v string) *DescribeEdgeFunctionsResponseBodyEdgeFunctions {
	s.ModifiedTime = &v
	return s
}

func (s *DescribeEdgeFunctionsResponseBodyEdgeFunctions) SetRuntime(v string) *DescribeEdgeFunctionsResponseBodyEdgeFunctions {
	s.Runtime = &v
	return s
}

func (s *DescribeEdgeFunctionsResponseBodyEdgeFunctions) SetStatus(v string) *DescribeEdgeFunctionsResponseBodyEdgeFunctions {
	s.Status = &v
	return s
}

func (s *DescribeEdgeFunctionsResponseBodyEdgeFunctions) SetUrlInternet(v string) *DescribeEdgeFunctionsResponseBodyEdgeFunctions {
	s.UrlInternet = &v
	return s
}

func (s *DescribeEdgeFunctionsResponseBodyEdgeFunctions) SetUrlIntranet(v string) *DescribeEdgeFunctionsResponseBodyEdgeFunctions {
	s.UrlIntranet = &v
	return s
}

func (s *DescribeEdgeFunctionsResponseBodyEdgeFunctions) Validate() error {
	return dara.Validate(s)
}
