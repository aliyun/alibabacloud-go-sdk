// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetErrorRequestSampleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDbName(v string) *GetErrorRequestSampleRequest
	GetDbName() *string
	SetEnd(v int64) *GetErrorRequestSampleRequest
	GetEnd() *int64
	SetInstanceId(v string) *GetErrorRequestSampleRequest
	GetInstanceId() *string
	SetNodeId(v string) *GetErrorRequestSampleRequest
	GetNodeId() *string
	SetSqlId(v string) *GetErrorRequestSampleRequest
	GetSqlId() *string
	SetStart(v int64) *GetErrorRequestSampleRequest
	GetStart() *int64
}

type GetErrorRequestSampleRequest struct {
	// The name of the database.
	//
	// example:
	//
	// testdb01
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// The end of the time range to query. Set this parameter to a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// >  The end time must be later than the start time. The interval cannot exceed 24 hours.
	//
	// example:
	//
	// 1642566830000
	End *int64 `json:"End,omitempty" xml:"End,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-2ze8g2am97624****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The node ID.
	//
	// >  You must specify the node ID if your database instance is a PolarDB for MySQL cluster.
	//
	// example:
	//
	// pi-bp179lg03445l****
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The SQL query ID. You can call the [GetAsyncErrorRequestListByCode](https://help.aliyun.com/document_detail/410746.html) operation to query the ID of the SQL query for which MySQL error code is returned.
	//
	// example:
	//
	// 2cd4432556c3dab9d825ba363637****
	SqlId *string `json:"SqlId,omitempty" xml:"SqlId,omitempty"`
	// The beginning of the time range to query. Set this parameter to a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// >  The start time must be within the storage duration of the SQL Explorer feature of the database instance, and can be up to 90 days earlier than the current time.
	//
	// example:
	//
	// 1642556990714
	Start *int64 `json:"Start,omitempty" xml:"Start,omitempty"`
}

func (s GetErrorRequestSampleRequest) String() string {
	return dara.Prettify(s)
}

func (s GetErrorRequestSampleRequest) GoString() string {
	return s.String()
}

func (s *GetErrorRequestSampleRequest) GetDbName() *string {
	return s.DbName
}

func (s *GetErrorRequestSampleRequest) GetEnd() *int64 {
	return s.End
}

func (s *GetErrorRequestSampleRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetErrorRequestSampleRequest) GetNodeId() *string {
	return s.NodeId
}

func (s *GetErrorRequestSampleRequest) GetSqlId() *string {
	return s.SqlId
}

func (s *GetErrorRequestSampleRequest) GetStart() *int64 {
	return s.Start
}

func (s *GetErrorRequestSampleRequest) SetDbName(v string) *GetErrorRequestSampleRequest {
	s.DbName = &v
	return s
}

func (s *GetErrorRequestSampleRequest) SetEnd(v int64) *GetErrorRequestSampleRequest {
	s.End = &v
	return s
}

func (s *GetErrorRequestSampleRequest) SetInstanceId(v string) *GetErrorRequestSampleRequest {
	s.InstanceId = &v
	return s
}

func (s *GetErrorRequestSampleRequest) SetNodeId(v string) *GetErrorRequestSampleRequest {
	s.NodeId = &v
	return s
}

func (s *GetErrorRequestSampleRequest) SetSqlId(v string) *GetErrorRequestSampleRequest {
	s.SqlId = &v
	return s
}

func (s *GetErrorRequestSampleRequest) SetStart(v int64) *GetErrorRequestSampleRequest {
	s.Start = &v
	return s
}

func (s *GetErrorRequestSampleRequest) Validate() error {
	return dara.Validate(s)
}
