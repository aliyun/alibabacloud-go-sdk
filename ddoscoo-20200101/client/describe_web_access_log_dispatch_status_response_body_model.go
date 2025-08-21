// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeWebAccessLogDispatchStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeWebAccessLogDispatchStatusResponseBody
	GetRequestId() *string
	SetSlsConfigStatus(v []*DescribeWebAccessLogDispatchStatusResponseBodySlsConfigStatus) *DescribeWebAccessLogDispatchStatusResponseBody
	GetSlsConfigStatus() []*DescribeWebAccessLogDispatchStatusResponseBodySlsConfigStatus
	SetTotalCount(v int32) *DescribeWebAccessLogDispatchStatusResponseBody
	GetTotalCount() *int32
}

type DescribeWebAccessLogDispatchStatusResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// CF33B4C3-196E-4015-AADD-5CAD00057B80
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the log analysis feature is enabled for domain names.
	SlsConfigStatus []*DescribeWebAccessLogDispatchStatusResponseBodySlsConfigStatus `json:"SlsConfigStatus,omitempty" xml:"SlsConfigStatus,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeWebAccessLogDispatchStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeWebAccessLogDispatchStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeWebAccessLogDispatchStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeWebAccessLogDispatchStatusResponseBody) GetSlsConfigStatus() []*DescribeWebAccessLogDispatchStatusResponseBodySlsConfigStatus {
	return s.SlsConfigStatus
}

func (s *DescribeWebAccessLogDispatchStatusResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeWebAccessLogDispatchStatusResponseBody) SetRequestId(v string) *DescribeWebAccessLogDispatchStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeWebAccessLogDispatchStatusResponseBody) SetSlsConfigStatus(v []*DescribeWebAccessLogDispatchStatusResponseBodySlsConfigStatus) *DescribeWebAccessLogDispatchStatusResponseBody {
	s.SlsConfigStatus = v
	return s
}

func (s *DescribeWebAccessLogDispatchStatusResponseBody) SetTotalCount(v int32) *DescribeWebAccessLogDispatchStatusResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeWebAccessLogDispatchStatusResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeWebAccessLogDispatchStatusResponseBodySlsConfigStatus struct {
	// The domain name.
	//
	// example:
	//
	// www.aliyundoc.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// Indicates whether the log analysis feature is enabled. Valid values:
	//
	// 	- **true**: yes
	//
	// 	- **false**: no
	//
	// example:
	//
	// true
	Enable *bool `json:"Enable,omitempty" xml:"Enable,omitempty"`
}

func (s DescribeWebAccessLogDispatchStatusResponseBodySlsConfigStatus) String() string {
	return dara.Prettify(s)
}

func (s DescribeWebAccessLogDispatchStatusResponseBodySlsConfigStatus) GoString() string {
	return s.String()
}

func (s *DescribeWebAccessLogDispatchStatusResponseBodySlsConfigStatus) GetDomain() *string {
	return s.Domain
}

func (s *DescribeWebAccessLogDispatchStatusResponseBodySlsConfigStatus) GetEnable() *bool {
	return s.Enable
}

func (s *DescribeWebAccessLogDispatchStatusResponseBodySlsConfigStatus) SetDomain(v string) *DescribeWebAccessLogDispatchStatusResponseBodySlsConfigStatus {
	s.Domain = &v
	return s
}

func (s *DescribeWebAccessLogDispatchStatusResponseBodySlsConfigStatus) SetEnable(v bool) *DescribeWebAccessLogDispatchStatusResponseBodySlsConfigStatus {
	s.Enable = &v
	return s
}

func (s *DescribeWebAccessLogDispatchStatusResponseBodySlsConfigStatus) Validate() error {
	return dara.Validate(s)
}
