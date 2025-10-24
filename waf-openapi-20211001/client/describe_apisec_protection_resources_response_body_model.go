// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApisecProtectionResourcesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*DescribeApisecProtectionResourcesResponseBodyData) *DescribeApisecProtectionResourcesResponseBody
	GetData() []*DescribeApisecProtectionResourcesResponseBodyData
	SetRequestId(v string) *DescribeApisecProtectionResourcesResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *DescribeApisecProtectionResourcesResponseBody
	GetTotalCount() *int64
}

type DescribeApisecProtectionResourcesResponseBody struct {
	// The protected objects.
	Data []*DescribeApisecProtectionResourcesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 2EFCFE18-78F8-5079-B312-07***48B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 5
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeApisecProtectionResourcesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeApisecProtectionResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeApisecProtectionResourcesResponseBody) GetData() []*DescribeApisecProtectionResourcesResponseBodyData {
	return s.Data
}

func (s *DescribeApisecProtectionResourcesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeApisecProtectionResourcesResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeApisecProtectionResourcesResponseBody) SetData(v []*DescribeApisecProtectionResourcesResponseBodyData) *DescribeApisecProtectionResourcesResponseBody {
	s.Data = v
	return s
}

func (s *DescribeApisecProtectionResourcesResponseBody) SetRequestId(v string) *DescribeApisecProtectionResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeApisecProtectionResourcesResponseBody) SetTotalCount(v int64) *DescribeApisecProtectionResourcesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeApisecProtectionResourcesResponseBody) Validate() error {
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

type DescribeApisecProtectionResourcesResponseBodyData struct {
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
	// 1
	ReportStatus *int64 `json:"ReportStatus,omitempty" xml:"ReportStatus,omitempty"`
	// The protected object.
	//
	// example:
	//
	// cwaf-***-waf
	Resource *string `json:"Resource,omitempty" xml:"Resource,omitempty"`
	// The switch of the tracing and auditing feature.
	//
	// example:
	//
	// 0
	TraceStatus *int64 `json:"TraceStatus,omitempty" xml:"TraceStatus,omitempty"`
}

func (s DescribeApisecProtectionResourcesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeApisecProtectionResourcesResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeApisecProtectionResourcesResponseBodyData) GetApisecStatus() *int64 {
	return s.ApisecStatus
}

func (s *DescribeApisecProtectionResourcesResponseBodyData) GetReportStatus() *int64 {
	return s.ReportStatus
}

func (s *DescribeApisecProtectionResourcesResponseBodyData) GetResource() *string {
	return s.Resource
}

func (s *DescribeApisecProtectionResourcesResponseBodyData) GetTraceStatus() *int64 {
	return s.TraceStatus
}

func (s *DescribeApisecProtectionResourcesResponseBodyData) SetApisecStatus(v int64) *DescribeApisecProtectionResourcesResponseBodyData {
	s.ApisecStatus = &v
	return s
}

func (s *DescribeApisecProtectionResourcesResponseBodyData) SetReportStatus(v int64) *DescribeApisecProtectionResourcesResponseBodyData {
	s.ReportStatus = &v
	return s
}

func (s *DescribeApisecProtectionResourcesResponseBodyData) SetResource(v string) *DescribeApisecProtectionResourcesResponseBodyData {
	s.Resource = &v
	return s
}

func (s *DescribeApisecProtectionResourcesResponseBodyData) SetTraceStatus(v int64) *DescribeApisecProtectionResourcesResponseBodyData {
	s.TraceStatus = &v
	return s
}

func (s *DescribeApisecProtectionResourcesResponseBodyData) Validate() error {
	return dara.Validate(s)
}
