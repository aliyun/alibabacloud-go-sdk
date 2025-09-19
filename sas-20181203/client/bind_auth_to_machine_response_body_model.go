// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindAuthToMachineResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBindCount(v int32) *BindAuthToMachineResponseBody
	GetBindCount() *int32
	SetInsufficientCoreCount(v int32) *BindAuthToMachineResponseBody
	GetInsufficientCoreCount() *int32
	SetInsufficientEcsCount(v int32) *BindAuthToMachineResponseBody
	GetInsufficientEcsCount() *int32
	SetRequestId(v string) *BindAuthToMachineResponseBody
	GetRequestId() *string
	SetResultCode(v int32) *BindAuthToMachineResponseBody
	GetResultCode() *int32
	SetUnBindCount(v int32) *BindAuthToMachineResponseBody
	GetUnBindCount() *int32
}

type BindAuthToMachineResponseBody struct {
	// The number of bound servers.
	//
	// example:
	//
	// 1
	BindCount *int32 `json:"BindCount,omitempty" xml:"BindCount,omitempty"`
	// The shortage in the quota for cores of servers that can be protected.
	//
	// example:
	//
	// 1
	InsufficientCoreCount *int32 `json:"InsufficientCoreCount,omitempty" xml:"InsufficientCoreCount,omitempty"`
	// The shortage in the quota for servers that can be protected.
	//
	// example:
	//
	// 1
	InsufficientEcsCount *int32 `json:"InsufficientEcsCount,omitempty" xml:"InsufficientEcsCount,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 23AD0BD2-8771-5647-819E-6BA51E212F80
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status code that indicates the result. Valid values:
	//
	// 	- **0**: The servers are bound to or unbound from Security Center.
	//
	// 	- **1**: The values that you specified for the parameters are invalid.
	//
	// 	- **2**: The quota for servers that can be protected is insufficient.
	//
	// 	- **3**: The quota for cores of servers that can be protected is insufficient.
	//
	// example:
	//
	// 2
	ResultCode *int32 `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	// The number of unbound servers.
	//
	// example:
	//
	// 1
	UnBindCount *int32 `json:"UnBindCount,omitempty" xml:"UnBindCount,omitempty"`
}

func (s BindAuthToMachineResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BindAuthToMachineResponseBody) GoString() string {
	return s.String()
}

func (s *BindAuthToMachineResponseBody) GetBindCount() *int32 {
	return s.BindCount
}

func (s *BindAuthToMachineResponseBody) GetInsufficientCoreCount() *int32 {
	return s.InsufficientCoreCount
}

func (s *BindAuthToMachineResponseBody) GetInsufficientEcsCount() *int32 {
	return s.InsufficientEcsCount
}

func (s *BindAuthToMachineResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BindAuthToMachineResponseBody) GetResultCode() *int32 {
	return s.ResultCode
}

func (s *BindAuthToMachineResponseBody) GetUnBindCount() *int32 {
	return s.UnBindCount
}

func (s *BindAuthToMachineResponseBody) SetBindCount(v int32) *BindAuthToMachineResponseBody {
	s.BindCount = &v
	return s
}

func (s *BindAuthToMachineResponseBody) SetInsufficientCoreCount(v int32) *BindAuthToMachineResponseBody {
	s.InsufficientCoreCount = &v
	return s
}

func (s *BindAuthToMachineResponseBody) SetInsufficientEcsCount(v int32) *BindAuthToMachineResponseBody {
	s.InsufficientEcsCount = &v
	return s
}

func (s *BindAuthToMachineResponseBody) SetRequestId(v string) *BindAuthToMachineResponseBody {
	s.RequestId = &v
	return s
}

func (s *BindAuthToMachineResponseBody) SetResultCode(v int32) *BindAuthToMachineResponseBody {
	s.ResultCode = &v
	return s
}

func (s *BindAuthToMachineResponseBody) SetUnBindCount(v int32) *BindAuthToMachineResponseBody {
	s.UnBindCount = &v
	return s
}

func (s *BindAuthToMachineResponseBody) Validate() error {
	return dara.Validate(s)
}
