// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAsyncErrorRequestStatByCodeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDbName(v string) *GetAsyncErrorRequestStatByCodeRequest
	GetDbName() *string
	SetEnd(v int64) *GetAsyncErrorRequestStatByCodeRequest
	GetEnd() *int64
	SetInstanceId(v string) *GetAsyncErrorRequestStatByCodeRequest
	GetInstanceId() *string
	SetNodeId(v string) *GetAsyncErrorRequestStatByCodeRequest
	GetNodeId() *string
	SetStart(v int64) *GetAsyncErrorRequestStatByCodeRequest
	GetStart() *int64
}

type GetAsyncErrorRequestStatByCodeRequest struct {
	// The name of a database.
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
	// >  This parameter must be specified for PolarDB for MySQL clusters.
	//
	// example:
	//
	// pi-wz9s658475e58****
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The beginning of the time range to query. Set this parameter to a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// >  The start time must be within the storage duration of the SQL Explorer feature of the database instance and can be up to 90 days earlier than the current time.
	//
	// example:
	//
	// 1642556990714
	Start *int64 `json:"Start,omitempty" xml:"Start,omitempty"`
}

func (s GetAsyncErrorRequestStatByCodeRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAsyncErrorRequestStatByCodeRequest) GoString() string {
	return s.String()
}

func (s *GetAsyncErrorRequestStatByCodeRequest) GetDbName() *string {
	return s.DbName
}

func (s *GetAsyncErrorRequestStatByCodeRequest) GetEnd() *int64 {
	return s.End
}

func (s *GetAsyncErrorRequestStatByCodeRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetAsyncErrorRequestStatByCodeRequest) GetNodeId() *string {
	return s.NodeId
}

func (s *GetAsyncErrorRequestStatByCodeRequest) GetStart() *int64 {
	return s.Start
}

func (s *GetAsyncErrorRequestStatByCodeRequest) SetDbName(v string) *GetAsyncErrorRequestStatByCodeRequest {
	s.DbName = &v
	return s
}

func (s *GetAsyncErrorRequestStatByCodeRequest) SetEnd(v int64) *GetAsyncErrorRequestStatByCodeRequest {
	s.End = &v
	return s
}

func (s *GetAsyncErrorRequestStatByCodeRequest) SetInstanceId(v string) *GetAsyncErrorRequestStatByCodeRequest {
	s.InstanceId = &v
	return s
}

func (s *GetAsyncErrorRequestStatByCodeRequest) SetNodeId(v string) *GetAsyncErrorRequestStatByCodeRequest {
	s.NodeId = &v
	return s
}

func (s *GetAsyncErrorRequestStatByCodeRequest) SetStart(v int64) *GetAsyncErrorRequestStatByCodeRequest {
	s.Start = &v
	return s
}

func (s *GetAsyncErrorRequestStatByCodeRequest) Validate() error {
	return dara.Validate(s)
}
