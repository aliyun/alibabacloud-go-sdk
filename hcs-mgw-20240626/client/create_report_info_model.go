// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateReportInfo interface {
	dara.Model
	String() string
	GoString() string
	SetJobName(v string) *CreateReportInfo
	GetJobName() *string
	SetRuntimeId(v int32) *CreateReportInfo
	GetRuntimeId() *int32
	SetVersion(v string) *CreateReportInfo
	GetVersion() *string
}

type CreateReportInfo struct {
	// example:
	//
	// <your-job-name>
	JobName *string `json:"JobName,omitempty" xml:"JobName,omitempty"`
	// example:
	//
	// 1
	RuntimeId *int32 `json:"RuntimeId,omitempty" xml:"RuntimeId,omitempty"`
	// example:
	//
	// <your-job-version>
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s CreateReportInfo) String() string {
	return dara.Prettify(s)
}

func (s CreateReportInfo) GoString() string {
	return s.String()
}

func (s *CreateReportInfo) GetJobName() *string {
	return s.JobName
}

func (s *CreateReportInfo) GetRuntimeId() *int32 {
	return s.RuntimeId
}

func (s *CreateReportInfo) GetVersion() *string {
	return s.Version
}

func (s *CreateReportInfo) SetJobName(v string) *CreateReportInfo {
	s.JobName = &v
	return s
}

func (s *CreateReportInfo) SetRuntimeId(v int32) *CreateReportInfo {
	s.RuntimeId = &v
	return s
}

func (s *CreateReportInfo) SetVersion(v string) *CreateReportInfo {
	s.Version = &v
	return s
}

func (s *CreateReportInfo) Validate() error {
	return dara.Validate(s)
}
