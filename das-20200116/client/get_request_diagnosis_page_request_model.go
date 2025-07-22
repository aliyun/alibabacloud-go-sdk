// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRequestDiagnosisPageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v int64) *GetRequestDiagnosisPageRequest
	GetEndTime() *int64
	SetInstanceId(v string) *GetRequestDiagnosisPageRequest
	GetInstanceId() *string
	SetNodeId(v string) *GetRequestDiagnosisPageRequest
	GetNodeId() *string
	SetPageNo(v int32) *GetRequestDiagnosisPageRequest
	GetPageNo() *int32
	SetPageSize(v int32) *GetRequestDiagnosisPageRequest
	GetPageSize() *int32
	SetStartTime(v int64) *GetRequestDiagnosisPageRequest
	GetStartTime() *int64
}

type GetRequestDiagnosisPageRequest struct {
	// The end of the time range to query. Set this parameter to a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1634972640000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-0iwhhl8gx0ld6****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The node ID.
	//
	// >  You must specify the node ID if your database instance is a PolarDB for MySQL, PolarDB for PostgreSQL (Compatible with Oracle), or ApsaraDB for MongoDB instance.
	//
	// example:
	//
	// 202****
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The page number. The value must be a positive integer. Default value: 1.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page. The value must be a positive integer. Default value: 10.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The beginning of the time range to query. Set this parameter to a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1633071840000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s GetRequestDiagnosisPageRequest) String() string {
	return dara.Prettify(s)
}

func (s GetRequestDiagnosisPageRequest) GoString() string {
	return s.String()
}

func (s *GetRequestDiagnosisPageRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *GetRequestDiagnosisPageRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetRequestDiagnosisPageRequest) GetNodeId() *string {
	return s.NodeId
}

func (s *GetRequestDiagnosisPageRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *GetRequestDiagnosisPageRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetRequestDiagnosisPageRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *GetRequestDiagnosisPageRequest) SetEndTime(v int64) *GetRequestDiagnosisPageRequest {
	s.EndTime = &v
	return s
}

func (s *GetRequestDiagnosisPageRequest) SetInstanceId(v string) *GetRequestDiagnosisPageRequest {
	s.InstanceId = &v
	return s
}

func (s *GetRequestDiagnosisPageRequest) SetNodeId(v string) *GetRequestDiagnosisPageRequest {
	s.NodeId = &v
	return s
}

func (s *GetRequestDiagnosisPageRequest) SetPageNo(v int32) *GetRequestDiagnosisPageRequest {
	s.PageNo = &v
	return s
}

func (s *GetRequestDiagnosisPageRequest) SetPageSize(v int32) *GetRequestDiagnosisPageRequest {
	s.PageSize = &v
	return s
}

func (s *GetRequestDiagnosisPageRequest) SetStartTime(v int64) *GetRequestDiagnosisPageRequest {
	s.StartTime = &v
	return s
}

func (s *GetRequestDiagnosisPageRequest) Validate() error {
	return dara.Validate(s)
}
