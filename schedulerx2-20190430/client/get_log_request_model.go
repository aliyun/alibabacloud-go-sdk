// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLogRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTimestamp(v int64) *GetLogRequest
	GetEndTimestamp() *int64
	SetGroupId(v string) *GetLogRequest
	GetGroupId() *string
	SetJobId(v string) *GetLogRequest
	GetJobId() *string
	SetJobInstanceId(v string) *GetLogRequest
	GetJobInstanceId() *string
	SetKeyword(v string) *GetLogRequest
	GetKeyword() *string
	SetLine(v int32) *GetLogRequest
	GetLine() *int32
	SetNamespace(v string) *GetLogRequest
	GetNamespace() *string
	SetNamespaceSource(v string) *GetLogRequest
	GetNamespaceSource() *string
	SetOffset(v int32) *GetLogRequest
	GetOffset() *int32
	SetRegionId(v string) *GetLogRequest
	GetRegionId() *string
	SetReverse(v bool) *GetLogRequest
	GetReverse() *bool
	SetStartTimestamp(v int64) *GetLogRequest
	GetStartTimestamp() *int64
}

type GetLogRequest struct {
	// The time when the job stops running. Specify a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1675739484000
	EndTimestamp *int64 `json:"EndTimestamp,omitempty" xml:"EndTimestamp,omitempty"`
	// The application group ID. You can obtain the application group ID on the Application Management page in the SchedulerX console.
	//
	// This parameter is required.
	//
	// example:
	//
	// testSchedulerx.defaultGroup
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The job ID.
	//
	// example:
	//
	// 123
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The job instance ID.
	//
	// example:
	//
	// 123456
	JobInstanceId *string `json:"JobInstanceId,omitempty" xml:"JobInstanceId,omitempty"`
	// The keyword.
	//
	// example:
	//
	// ERROR
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// The number of rows to return. The maximum number is 200.
	//
	// example:
	//
	// 50
	Line *int32 `json:"Line,omitempty" xml:"Line,omitempty"`
	// The namespace ID. You can obtain the namespace ID on the Namespace page in the SchedulerX console.
	//
	// This parameter is required.
	//
	// example:
	//
	// adcfc35d-e2fe-4fe9-bbaa-20e90ffc****
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The source of the namespace. This parameter is required only for a special third party.
	//
	// example:
	//
	// schedulerx
	NamespaceSource *string `json:"NamespaceSource,omitempty" xml:"NamespaceSource,omitempty"`
	// The number of offset rows. This parameter can be used for a paged query.
	//
	// example:
	//
	// 0
	Offset *int32 `json:"Offset,omitempty" xml:"Offset,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Specifies whether to reverse the order. By default, the order is reversed.
	//
	// example:
	//
	// true
	Reverse *bool `json:"Reverse,omitempty" xml:"Reverse,omitempty"`
	// The time when the job starts to run. Specify a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1675739364000
	StartTimestamp *int64 `json:"StartTimestamp,omitempty" xml:"StartTimestamp,omitempty"`
}

func (s GetLogRequest) String() string {
	return dara.Prettify(s)
}

func (s GetLogRequest) GoString() string {
	return s.String()
}

func (s *GetLogRequest) GetEndTimestamp() *int64 {
	return s.EndTimestamp
}

func (s *GetLogRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *GetLogRequest) GetJobId() *string {
	return s.JobId
}

func (s *GetLogRequest) GetJobInstanceId() *string {
	return s.JobInstanceId
}

func (s *GetLogRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *GetLogRequest) GetLine() *int32 {
	return s.Line
}

func (s *GetLogRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *GetLogRequest) GetNamespaceSource() *string {
	return s.NamespaceSource
}

func (s *GetLogRequest) GetOffset() *int32 {
	return s.Offset
}

func (s *GetLogRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetLogRequest) GetReverse() *bool {
	return s.Reverse
}

func (s *GetLogRequest) GetStartTimestamp() *int64 {
	return s.StartTimestamp
}

func (s *GetLogRequest) SetEndTimestamp(v int64) *GetLogRequest {
	s.EndTimestamp = &v
	return s
}

func (s *GetLogRequest) SetGroupId(v string) *GetLogRequest {
	s.GroupId = &v
	return s
}

func (s *GetLogRequest) SetJobId(v string) *GetLogRequest {
	s.JobId = &v
	return s
}

func (s *GetLogRequest) SetJobInstanceId(v string) *GetLogRequest {
	s.JobInstanceId = &v
	return s
}

func (s *GetLogRequest) SetKeyword(v string) *GetLogRequest {
	s.Keyword = &v
	return s
}

func (s *GetLogRequest) SetLine(v int32) *GetLogRequest {
	s.Line = &v
	return s
}

func (s *GetLogRequest) SetNamespace(v string) *GetLogRequest {
	s.Namespace = &v
	return s
}

func (s *GetLogRequest) SetNamespaceSource(v string) *GetLogRequest {
	s.NamespaceSource = &v
	return s
}

func (s *GetLogRequest) SetOffset(v int32) *GetLogRequest {
	s.Offset = &v
	return s
}

func (s *GetLogRequest) SetRegionId(v string) *GetLogRequest {
	s.RegionId = &v
	return s
}

func (s *GetLogRequest) SetReverse(v bool) *GetLogRequest {
	s.Reverse = &v
	return s
}

func (s *GetLogRequest) SetStartTimestamp(v int64) *GetLogRequest {
	s.StartTimestamp = &v
	return s
}

func (s *GetLogRequest) Validate() error {
	return dara.Validate(s)
}
