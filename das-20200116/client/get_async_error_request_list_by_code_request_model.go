// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAsyncErrorRequestListByCodeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnd(v int64) *GetAsyncErrorRequestListByCodeRequest
	GetEnd() *int64
	SetErrorCode(v string) *GetAsyncErrorRequestListByCodeRequest
	GetErrorCode() *string
	SetInstanceId(v string) *GetAsyncErrorRequestListByCodeRequest
	GetInstanceId() *string
	SetNodeId(v string) *GetAsyncErrorRequestListByCodeRequest
	GetNodeId() *string
	SetStart(v int64) *GetAsyncErrorRequestListByCodeRequest
	GetStart() *int64
}

type GetAsyncErrorRequestListByCodeRequest struct {
	// The end of the time range to query. Set this parameter to a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// >  The end time must be later than the start time. The interval between the start time and the end time cannot exceed 24 hours.
	//
	// example:
	//
	// 1642566830000
	End *int64 `json:"End,omitempty" xml:"End,omitempty"`
	// The error code. You can call the [GetAsyncErrorRequestStatByCode](https://help.aliyun.com/document_detail/409804.html) operation to query the MySQL error codes that may be generated in the SQL Explorer results of an instance.
	//
	// example:
	//
	// 1064
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
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
	// >  This parameter must be specified if the database instance is a PolarDB for MySQL cluster.
	//
	// example:
	//
	// pi-wz9s658475e58****
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The beginning of the time range to query. Set this parameter to a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// >  The start time must be within the storage duration of the SQL Explorer feature of the database instance, and can be up to 90 days earlier than the current time.
	//
	// example:
	//
	// 1642556990714
	Start *int64 `json:"Start,omitempty" xml:"Start,omitempty"`
}

func (s GetAsyncErrorRequestListByCodeRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAsyncErrorRequestListByCodeRequest) GoString() string {
	return s.String()
}

func (s *GetAsyncErrorRequestListByCodeRequest) GetEnd() *int64 {
	return s.End
}

func (s *GetAsyncErrorRequestListByCodeRequest) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetAsyncErrorRequestListByCodeRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetAsyncErrorRequestListByCodeRequest) GetNodeId() *string {
	return s.NodeId
}

func (s *GetAsyncErrorRequestListByCodeRequest) GetStart() *int64 {
	return s.Start
}

func (s *GetAsyncErrorRequestListByCodeRequest) SetEnd(v int64) *GetAsyncErrorRequestListByCodeRequest {
	s.End = &v
	return s
}

func (s *GetAsyncErrorRequestListByCodeRequest) SetErrorCode(v string) *GetAsyncErrorRequestListByCodeRequest {
	s.ErrorCode = &v
	return s
}

func (s *GetAsyncErrorRequestListByCodeRequest) SetInstanceId(v string) *GetAsyncErrorRequestListByCodeRequest {
	s.InstanceId = &v
	return s
}

func (s *GetAsyncErrorRequestListByCodeRequest) SetNodeId(v string) *GetAsyncErrorRequestListByCodeRequest {
	s.NodeId = &v
	return s
}

func (s *GetAsyncErrorRequestListByCodeRequest) SetStart(v int64) *GetAsyncErrorRequestListByCodeRequest {
	s.Start = &v
	return s
}

func (s *GetAsyncErrorRequestListByCodeRequest) Validate() error {
	return dara.Validate(s)
}
