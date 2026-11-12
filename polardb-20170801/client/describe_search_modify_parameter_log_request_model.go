// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSearchModifyParameterLogRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeSearchModifyParameterLogRequest
	GetDBClusterId() *string
	SetEndTime(v string) *DescribeSearchModifyParameterLogRequest
	GetEndTime() *string
	SetStartTime(v string) *DescribeSearchModifyParameterLogRequest
	GetStartTime() *string
}

type DescribeSearchModifyParameterLogRequest struct {
	// The ID of the PolarDB for MySQL primary cluster, which is used to query the parameter modification history of the associated Search cluster. Before you call this operation, ensure that PolarSearch is enabled for the target cluster. You can add a search node on the basic information page of the cluster in the console. This operation is not applicable to clusters for which PolarSearch is not enabled.
	//
	// This parameter is required.
	//
	// example:
	//
	// pc-xxxxxxxx
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The end time of the query. Specify the time in the yyyy-MM-ddTHH:mmZ format. The time must be in UTC. Only records modified before this time are returned. The end time must not be earlier than the start time.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2026-10-14T02:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The start time of the query. Specify the time in the yyyy-MM-ddTHH:mmZ format. The time must be in UTC. Only records modified after this time are returned.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2026-10-14T01:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeSearchModifyParameterLogRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSearchModifyParameterLogRequest) GoString() string {
	return s.String()
}

func (s *DescribeSearchModifyParameterLogRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeSearchModifyParameterLogRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeSearchModifyParameterLogRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeSearchModifyParameterLogRequest) SetDBClusterId(v string) *DescribeSearchModifyParameterLogRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeSearchModifyParameterLogRequest) SetEndTime(v string) *DescribeSearchModifyParameterLogRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeSearchModifyParameterLogRequest) SetStartTime(v string) *DescribeSearchModifyParameterLogRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeSearchModifyParameterLogRequest) Validate() error {
	return dara.Validate(s)
}
