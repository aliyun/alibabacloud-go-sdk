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
	SetUserId(v string) *GetFullRequestSampleByInstanceIdRequest
	GetUserId() *string
}

type GetFullRequestSampleByInstanceIdRequest struct {
	// The end of the time range to query. Set this parameter to a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// >  The end time must be later than the start time. The interval between the start time and the end time must be equal to or greater than 1 hour.
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
	// The role of the PolarDB-X 2.0 node. Valid values:
	//
	// 	- **polarx_cn**: compute node.
	//
	// 	- **polarx_en**: data node.
	//
	// example:
	//
	// polarx_cn
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// The SQL statement ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 651b56fe9418d48edb8fdf0980ec****
	SqlId *string `json:"SqlId,omitempty" xml:"SqlId,omitempty"`
	// The beginning of the time range to query. Set this parameter to a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// >  The start time must be within the storage duration of the SQL Explorer feature of the database instance, and can be up to 90 days earlier than the current time.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1660097421000
	Start *int64 `json:"Start,omitempty" xml:"Start,omitempty"`
	// The ID of the Alibaba Cloud account that is used to create the database instance.
	//
	// >  This parameter is optional. The system can automatically obtain the account ID based on the value of InstanceId when you call this operation.
	//
	// example:
	//
	// 196278346919****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
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

func (s *GetFullRequestSampleByInstanceIdRequest) GetUserId() *string {
	return s.UserId
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

func (s *GetFullRequestSampleByInstanceIdRequest) SetUserId(v string) *GetFullRequestSampleByInstanceIdRequest {
	s.UserId = &v
	return s
}

func (s *GetFullRequestSampleByInstanceIdRequest) Validate() error {
	return dara.Validate(s)
}
