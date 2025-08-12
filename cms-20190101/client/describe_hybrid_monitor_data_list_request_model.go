// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHybridMonitorDataListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnd(v int64) *DescribeHybridMonitorDataListRequest
	GetEnd() *int64
	SetNamespace(v string) *DescribeHybridMonitorDataListRequest
	GetNamespace() *string
	SetPeriod(v string) *DescribeHybridMonitorDataListRequest
	GetPeriod() *string
	SetPromSQL(v string) *DescribeHybridMonitorDataListRequest
	GetPromSQL() *string
	SetRegionId(v string) *DescribeHybridMonitorDataListRequest
	GetRegionId() *string
	SetStart(v int64) *DescribeHybridMonitorDataListRequest
	GetStart() *int64
}

type DescribeHybridMonitorDataListRequest struct {
	// The end of the time range to query.
	//
	// Unit: seconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1653805225
	End *int64 `json:"End,omitempty" xml:"End,omitempty"`
	// The name of the namespace.
	//
	// For more information about how to query the names of namespaces, see [DescribeHybridMonitorNamespaceList](https://help.aliyun.com/document_detail/428880.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// default-aliyun
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The statistical period of the monitoring data.
	//
	// Unit: seconds.
	//
	// example:
	//
	// 60
	Period *string `json:"Period,omitempty" xml:"Period,omitempty"`
	// The metric name.
	//
	// >  PromQL statements are supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// AliyunEcs_cpu_total
	PromSQL  *string `json:"PromSQL,omitempty" xml:"PromSQL,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The start of the time range to query.
	//
	// Unit: seconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1653804865
	Start *int64 `json:"Start,omitempty" xml:"Start,omitempty"`
}

func (s DescribeHybridMonitorDataListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeHybridMonitorDataListRequest) GoString() string {
	return s.String()
}

func (s *DescribeHybridMonitorDataListRequest) GetEnd() *int64 {
	return s.End
}

func (s *DescribeHybridMonitorDataListRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *DescribeHybridMonitorDataListRequest) GetPeriod() *string {
	return s.Period
}

func (s *DescribeHybridMonitorDataListRequest) GetPromSQL() *string {
	return s.PromSQL
}

func (s *DescribeHybridMonitorDataListRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeHybridMonitorDataListRequest) GetStart() *int64 {
	return s.Start
}

func (s *DescribeHybridMonitorDataListRequest) SetEnd(v int64) *DescribeHybridMonitorDataListRequest {
	s.End = &v
	return s
}

func (s *DescribeHybridMonitorDataListRequest) SetNamespace(v string) *DescribeHybridMonitorDataListRequest {
	s.Namespace = &v
	return s
}

func (s *DescribeHybridMonitorDataListRequest) SetPeriod(v string) *DescribeHybridMonitorDataListRequest {
	s.Period = &v
	return s
}

func (s *DescribeHybridMonitorDataListRequest) SetPromSQL(v string) *DescribeHybridMonitorDataListRequest {
	s.PromSQL = &v
	return s
}

func (s *DescribeHybridMonitorDataListRequest) SetRegionId(v string) *DescribeHybridMonitorDataListRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeHybridMonitorDataListRequest) SetStart(v int64) *DescribeHybridMonitorDataListRequest {
	s.Start = &v
	return s
}

func (s *DescribeHybridMonitorDataListRequest) Validate() error {
	return dara.Validate(s)
}
