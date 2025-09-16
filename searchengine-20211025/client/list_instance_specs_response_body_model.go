// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstanceSpecsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListInstanceSpecsResponseBody
	GetRequestId() *string
	SetResult(v []*ListInstanceSpecsResponseBodyResult) *ListInstanceSpecsResponseBody
	GetResult() []*ListInstanceSpecsResponseBodyResult
}

type ListInstanceSpecsResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 10D5E615-69F7-5F49-B850-00169ADE513C
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The instance types.
	Result []*ListInstanceSpecsResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s ListInstanceSpecsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceSpecsResponseBody) GoString() string {
	return s.String()
}

func (s *ListInstanceSpecsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListInstanceSpecsResponseBody) GetResult() []*ListInstanceSpecsResponseBodyResult {
	return s.Result
}

func (s *ListInstanceSpecsResponseBody) SetRequestId(v string) *ListInstanceSpecsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInstanceSpecsResponseBody) SetResult(v []*ListInstanceSpecsResponseBodyResult) *ListInstanceSpecsResponseBody {
	s.Result = v
	return s
}

func (s *ListInstanceSpecsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListInstanceSpecsResponseBodyResult struct {
	// The number of vCPUs.
	//
	// example:
	//
	// 1
	Cpu *int32 `json:"cpu,omitempty" xml:"cpu,omitempty"`
	// The maximum storage of a single data node. Unit: GB.
	//
	// example:
	//
	// 600
	MaxDisk *int32 `json:"maxDisk,omitempty" xml:"maxDisk,omitempty"`
	// The memory of the instance. Unit: GB.
	//
	// example:
	//
	// 4
	Mem *int32 `json:"mem,omitempty" xml:"mem,omitempty"`
	// The minimum storage of a single data node. Unit: GB.
	//
	// example:
	//
	// 100
	MinDisk *int32 `json:"minDisk,omitempty" xml:"minDisk,omitempty"`
}

func (s ListInstanceSpecsResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceSpecsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListInstanceSpecsResponseBodyResult) GetCpu() *int32 {
	return s.Cpu
}

func (s *ListInstanceSpecsResponseBodyResult) GetMaxDisk() *int32 {
	return s.MaxDisk
}

func (s *ListInstanceSpecsResponseBodyResult) GetMem() *int32 {
	return s.Mem
}

func (s *ListInstanceSpecsResponseBodyResult) GetMinDisk() *int32 {
	return s.MinDisk
}

func (s *ListInstanceSpecsResponseBodyResult) SetCpu(v int32) *ListInstanceSpecsResponseBodyResult {
	s.Cpu = &v
	return s
}

func (s *ListInstanceSpecsResponseBodyResult) SetMaxDisk(v int32) *ListInstanceSpecsResponseBodyResult {
	s.MaxDisk = &v
	return s
}

func (s *ListInstanceSpecsResponseBodyResult) SetMem(v int32) *ListInstanceSpecsResponseBodyResult {
	s.Mem = &v
	return s
}

func (s *ListInstanceSpecsResponseBodyResult) SetMinDisk(v int32) *ListInstanceSpecsResponseBodyResult {
	s.MinDisk = &v
	return s
}

func (s *ListInstanceSpecsResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
