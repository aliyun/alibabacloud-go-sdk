// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFullRequestSampleByInstanceIdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnd(v int64) *GetFullRequestSampleByInstanceIdRequest
	GetEnd() *int64
	SetInstanceId(v string) *GetFullRequestSampleByInstanceIdRequest
	GetInstanceId() *string
	SetRole(v string) *GetFullRequestSampleByInstanceIdRequest
	GetRole() *string
	SetSqlId(v string) *GetFullRequestSampleByInstanceIdRequest
	GetSqlId() *string
	SetStart(v int64) *GetFullRequestSampleByInstanceIdRequest
	GetStart() *int64
}

type GetFullRequestSampleByInstanceIdRequest struct {
	// The end of the time range to query. Specify a UNIX timestamp in milliseconds.
	//
	// > The end time must be later than the start time, and the interval between the start time and end time cannot be less than 1 hour.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1660104621000
	End *int64 `json:"End,omitempty" xml:"End,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-2ze8g2am97624****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The node information of a PolarDB-X 2.0 database instance.
	//
	// - **polarx_cn**: compute node.
	//
	// - **polarx_en**: data node.
	//
	// example:
	//
	// polarx_cn
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// SQL ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// 651b56fe9418d48edb8fdf0980ec****
	SqlId *string `json:"SqlId,omitempty" xml:"SqlId,omitempty"`
	// The beginning of the time range to query. Specify a UNIX timestamp in milliseconds.
	//
	// > The start time must be within the storage duration of SQL Explorer for the database instance and cannot be earlier than 90 days before the current time.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1660097421000
	Start *int64 `json:"Start,omitempty" xml:"Start,omitempty"`
}

func (s GetFullRequestSampleByInstanceIdRequest) String() string {
	return dara.Prettify(s)
}

func (s GetFullRequestSampleByInstanceIdRequest) GoString() string {
	return s.String()
}

func (s *GetFullRequestSampleByInstanceIdRequest) GetEnd() *int64 {
	return s.End
}

func (s *GetFullRequestSampleByInstanceIdRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetFullRequestSampleByInstanceIdRequest) GetRole() *string {
	return s.Role
}

func (s *GetFullRequestSampleByInstanceIdRequest) GetSqlId() *string {
	return s.SqlId
}

func (s *GetFullRequestSampleByInstanceIdRequest) GetStart() *int64 {
	return s.Start
}

func (s *GetFullRequestSampleByInstanceIdRequest) SetEnd(v int64) *GetFullRequestSampleByInstanceIdRequest {
	s.End = &v
	return s
}

func (s *GetFullRequestSampleByInstanceIdRequest) SetInstanceId(v string) *GetFullRequestSampleByInstanceIdRequest {
	s.InstanceId = &v
	return s
}

func (s *GetFullRequestSampleByInstanceIdRequest) SetRole(v string) *GetFullRequestSampleByInstanceIdRequest {
	s.Role = &v
	return s
}

func (s *GetFullRequestSampleByInstanceIdRequest) SetSqlId(v string) *GetFullRequestSampleByInstanceIdRequest {
	s.SqlId = &v
	return s
}

func (s *GetFullRequestSampleByInstanceIdRequest) SetStart(v int64) *GetFullRequestSampleByInstanceIdRequest {
	s.Start = &v
	return s
}

func (s *GetFullRequestSampleByInstanceIdRequest) Validate() error {
	return dara.Validate(s)
}
