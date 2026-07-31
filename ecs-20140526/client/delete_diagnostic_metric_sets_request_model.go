// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDiagnosticMetricSetsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMetricSetIds(v []*string) *DeleteDiagnosticMetricSetsRequest
	GetMetricSetIds() []*string
	SetRegionId(v string) *DeleteDiagnosticMetricSetsRequest
	GetRegionId() *string
}

type DeleteDiagnosticMetricSetsRequest struct {
	// The list of diagnostic metric set IDs. You can specify up to 10 IDs.
	//
	// This parameter is required.
	MetricSetIds []*string `json:"MetricSetIds,omitempty" xml:"MetricSetIds,omitempty" type:"Repeated"`
	// The region ID. You can call [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteDiagnosticMetricSetsRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDiagnosticMetricSetsRequest) GoString() string {
	return s.String()
}

func (s *DeleteDiagnosticMetricSetsRequest) GetMetricSetIds() []*string {
	return s.MetricSetIds
}

func (s *DeleteDiagnosticMetricSetsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteDiagnosticMetricSetsRequest) SetMetricSetIds(v []*string) *DeleteDiagnosticMetricSetsRequest {
	s.MetricSetIds = v
	return s
}

func (s *DeleteDiagnosticMetricSetsRequest) SetRegionId(v string) *DeleteDiagnosticMetricSetsRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteDiagnosticMetricSetsRequest) Validate() error {
	return dara.Validate(s)
}
