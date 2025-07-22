// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAsyncErrorRequestStatResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDbName(v string) *GetAsyncErrorRequestStatResultRequest
	GetDbName() *string
	SetEnd(v int64) *GetAsyncErrorRequestStatResultRequest
	GetEnd() *int64
	SetInstanceId(v string) *GetAsyncErrorRequestStatResultRequest
	GetInstanceId() *string
	SetNodeId(v string) *GetAsyncErrorRequestStatResultRequest
	GetNodeId() *string
	SetSqlIdList(v string) *GetAsyncErrorRequestStatResultRequest
	GetSqlIdList() *string
	SetStart(v int64) *GetAsyncErrorRequestStatResultRequest
	GetStart() *int64
}

type GetAsyncErrorRequestStatResultRequest struct {
	// The name of the database.
	//
	// example:
	//
	// testdb01
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// The end of the time range to query. Set this parameter to a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// >  The end time must be later than the start time. The interval between the start time and the end time cannot exceed 24 hours.
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
	// >  This parameter must be specified for PolarDB for MySQL instances.
	//
	// example:
	//
	// pi-bp179lg03445l****
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The ID of the SQL template. Separate multiple SQL IDs with commas (,). You can call the [GetAsyncErrorRequestListByCode](https://help.aliyun.com/document_detail/410746.html) operation to query the ID of the SQL query for which MySQL error code is returned.
	//
	// example:
	//
	// ad78a4e7d3ce81590c9dc2d5f4bc****,0f92feacd92c048b06a16617a633****
	SqlIdList *string `json:"SqlIdList,omitempty" xml:"SqlIdList,omitempty"`
	// The beginning of the time range to query. Set this parameter to a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// >  The beginning of the time range to query must be within the storage duration of the database instance and can be up to 90 days earlier than the current time.
	//
	// example:
	//
	// 1642556990714
	Start *int64 `json:"Start,omitempty" xml:"Start,omitempty"`
}

func (s GetAsyncErrorRequestStatResultRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAsyncErrorRequestStatResultRequest) GoString() string {
	return s.String()
}

func (s *GetAsyncErrorRequestStatResultRequest) GetDbName() *string {
	return s.DbName
}

func (s *GetAsyncErrorRequestStatResultRequest) GetEnd() *int64 {
	return s.End
}

func (s *GetAsyncErrorRequestStatResultRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetAsyncErrorRequestStatResultRequest) GetNodeId() *string {
	return s.NodeId
}

func (s *GetAsyncErrorRequestStatResultRequest) GetSqlIdList() *string {
	return s.SqlIdList
}

func (s *GetAsyncErrorRequestStatResultRequest) GetStart() *int64 {
	return s.Start
}

func (s *GetAsyncErrorRequestStatResultRequest) SetDbName(v string) *GetAsyncErrorRequestStatResultRequest {
	s.DbName = &v
	return s
}

func (s *GetAsyncErrorRequestStatResultRequest) SetEnd(v int64) *GetAsyncErrorRequestStatResultRequest {
	s.End = &v
	return s
}

func (s *GetAsyncErrorRequestStatResultRequest) SetInstanceId(v string) *GetAsyncErrorRequestStatResultRequest {
	s.InstanceId = &v
	return s
}

func (s *GetAsyncErrorRequestStatResultRequest) SetNodeId(v string) *GetAsyncErrorRequestStatResultRequest {
	s.NodeId = &v
	return s
}

func (s *GetAsyncErrorRequestStatResultRequest) SetSqlIdList(v string) *GetAsyncErrorRequestStatResultRequest {
	s.SqlIdList = &v
	return s
}

func (s *GetAsyncErrorRequestStatResultRequest) SetStart(v int64) *GetAsyncErrorRequestStatResultRequest {
	s.Start = &v
	return s
}

func (s *GetAsyncErrorRequestStatResultRequest) Validate() error {
	return dara.Validate(s)
}
