// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApisecProtectionGroupsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*DescribeApisecProtectionGroupsResponseBodyData) *DescribeApisecProtectionGroupsResponseBody
	GetData() []*DescribeApisecProtectionGroupsResponseBodyData
	SetRequestId(v string) *DescribeApisecProtectionGroupsResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *DescribeApisecProtectionGroupsResponseBody
	GetTotalCount() *int64
}

type DescribeApisecProtectionGroupsResponseBody struct {
	// The protected object groups.
	Data []*DescribeApisecProtectionGroupsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// D7861F61-5B61-46CE-A47C-6B19****5EB0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of protected object groups.
	//
	// example:
	//
	// 8
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeApisecProtectionGroupsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeApisecProtectionGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeApisecProtectionGroupsResponseBody) GetData() []*DescribeApisecProtectionGroupsResponseBodyData {
	return s.Data
}

func (s *DescribeApisecProtectionGroupsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeApisecProtectionGroupsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeApisecProtectionGroupsResponseBody) SetData(v []*DescribeApisecProtectionGroupsResponseBodyData) *DescribeApisecProtectionGroupsResponseBody {
	s.Data = v
	return s
}

func (s *DescribeApisecProtectionGroupsResponseBody) SetRequestId(v string) *DescribeApisecProtectionGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeApisecProtectionGroupsResponseBody) SetTotalCount(v int64) *DescribeApisecProtectionGroupsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeApisecProtectionGroupsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeApisecProtectionGroupsResponseBodyData struct {
	// The switch of the API security module.
	//
	// example:
	//
	// 1
	ApisecStatus *int64 `json:"ApisecStatus,omitempty" xml:"ApisecStatus,omitempty"`
	// The switch of the compliance check feature.
	//
	// example:
	//
	// 0
	ReportStatus *int64 `json:"ReportStatus,omitempty" xml:"ReportStatus,omitempty"`
	// The name of the protected object group.
	//
	// example:
	//
	// group1
	ResourceGroup *string `json:"ResourceGroup,omitempty" xml:"ResourceGroup,omitempty"`
	// The switch of the tracing and auditing feature.
	//
	// example:
	//
	// 0
	TraceStatus *int64 `json:"TraceStatus,omitempty" xml:"TraceStatus,omitempty"`
}

func (s DescribeApisecProtectionGroupsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeApisecProtectionGroupsResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeApisecProtectionGroupsResponseBodyData) GetApisecStatus() *int64 {
	return s.ApisecStatus
}

func (s *DescribeApisecProtectionGroupsResponseBodyData) GetReportStatus() *int64 {
	return s.ReportStatus
}

func (s *DescribeApisecProtectionGroupsResponseBodyData) GetResourceGroup() *string {
	return s.ResourceGroup
}

func (s *DescribeApisecProtectionGroupsResponseBodyData) GetTraceStatus() *int64 {
	return s.TraceStatus
}

func (s *DescribeApisecProtectionGroupsResponseBodyData) SetApisecStatus(v int64) *DescribeApisecProtectionGroupsResponseBodyData {
	s.ApisecStatus = &v
	return s
}

func (s *DescribeApisecProtectionGroupsResponseBodyData) SetReportStatus(v int64) *DescribeApisecProtectionGroupsResponseBodyData {
	s.ReportStatus = &v
	return s
}

func (s *DescribeApisecProtectionGroupsResponseBodyData) SetResourceGroup(v string) *DescribeApisecProtectionGroupsResponseBodyData {
	s.ResourceGroup = &v
	return s
}

func (s *DescribeApisecProtectionGroupsResponseBodyData) SetTraceStatus(v int64) *DescribeApisecProtectionGroupsResponseBodyData {
	s.TraceStatus = &v
	return s
}

func (s *DescribeApisecProtectionGroupsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
