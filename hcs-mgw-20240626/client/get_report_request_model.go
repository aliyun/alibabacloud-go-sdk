// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetReportRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRuntimeId(v int32) *GetReportRequest
	GetRuntimeId() *int32
	SetVersion(v string) *GetReportRequest
	GetVersion() *string
}

type GetReportRequest struct {
	// The execution ID of the migration task.
	//
	// example:
	//
	// 1
	RuntimeId *int32 `json:"runtimeId,omitempty" xml:"runtimeId,omitempty"`
	// The ID of the migration task.
	//
	// This parameter is required.
	//
	// example:
	//
	// test_job_id
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s GetReportRequest) String() string {
	return dara.Prettify(s)
}

func (s GetReportRequest) GoString() string {
	return s.String()
}

func (s *GetReportRequest) GetRuntimeId() *int32 {
	return s.RuntimeId
}

func (s *GetReportRequest) GetVersion() *string {
	return s.Version
}

func (s *GetReportRequest) SetRuntimeId(v int32) *GetReportRequest {
	s.RuntimeId = &v
	return s
}

func (s *GetReportRequest) SetVersion(v string) *GetReportRequest {
	s.Version = &v
	return s
}

func (s *GetReportRequest) Validate() error {
	return dara.Validate(s)
}
