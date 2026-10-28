// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPfsSqlSampleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v int64) *GetPfsSqlSampleRequest
	GetEndTime() *int64
	SetInstanceId(v string) *GetPfsSqlSampleRequest
	GetInstanceId() *string
	SetNodeId(v string) *GetPfsSqlSampleRequest
	GetNodeId() *string
	SetSqlId(v string) *GetPfsSqlSampleRequest
	GetSqlId() *string
	SetStartTime(v int64) *GetPfsSqlSampleRequest
	GetStartTime() *int64
}

type GetPfsSqlSampleRequest struct {
	// The end time of the query. The time is a UNIX timestamp in milliseconds.
	//
	// > The end time must be later than the start time. You can query data for up to seven days within the last month.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1678074351197
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The instance ID.
	//
	// > Only ApsaraDB RDS for MySQL and PolarDB for MySQL instances are supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-2ze1jdv45i7l6****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The node ID.
	//
	// > For an ApsaraDB RDS for MySQL Cluster Edition instance or a PolarDB for MySQL cluster, you must specify the node ID.
	//
	// example:
	//
	// r-x****-db-0
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The SQL ID.
	//
	// example:
	//
	// 651b56fe9418d48edb8fdf0980ec****
	SqlId *string `json:"SqlId,omitempty" xml:"SqlId,omitempty"`
	// The start time of the query. The time is a UNIX timestamp in milliseconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1676511134614
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s GetPfsSqlSampleRequest) String() string {
	return dara.Prettify(s)
}

func (s GetPfsSqlSampleRequest) GoString() string {
	return s.String()
}

func (s *GetPfsSqlSampleRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *GetPfsSqlSampleRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetPfsSqlSampleRequest) GetNodeId() *string {
	return s.NodeId
}

func (s *GetPfsSqlSampleRequest) GetSqlId() *string {
	return s.SqlId
}

func (s *GetPfsSqlSampleRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *GetPfsSqlSampleRequest) SetEndTime(v int64) *GetPfsSqlSampleRequest {
	s.EndTime = &v
	return s
}

func (s *GetPfsSqlSampleRequest) SetInstanceId(v string) *GetPfsSqlSampleRequest {
	s.InstanceId = &v
	return s
}

func (s *GetPfsSqlSampleRequest) SetNodeId(v string) *GetPfsSqlSampleRequest {
	s.NodeId = &v
	return s
}

func (s *GetPfsSqlSampleRequest) SetSqlId(v string) *GetPfsSqlSampleRequest {
	s.SqlId = &v
	return s
}

func (s *GetPfsSqlSampleRequest) SetStartTime(v int64) *GetPfsSqlSampleRequest {
	s.StartTime = &v
	return s
}

func (s *GetPfsSqlSampleRequest) Validate() error {
	return dara.Validate(s)
}
