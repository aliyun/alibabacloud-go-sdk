// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVmcoreDiagnosisTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDebuginfoCommonUrl(v string) *CreateVmcoreDiagnosisTaskRequest
	GetDebuginfoCommonUrl() *string
	SetDebuginfoUrl(v string) *CreateVmcoreDiagnosisTaskRequest
	GetDebuginfoUrl() *string
	SetDmesgUrl(v string) *CreateVmcoreDiagnosisTaskRequest
	GetDmesgUrl() *string
	SetTaskType(v string) *CreateVmcoreDiagnosisTaskRequest
	GetTaskType() *string
	SetVmcoreUrl(v string) *CreateVmcoreDiagnosisTaskRequest
	GetVmcoreUrl() *string
}

type CreateVmcoreDiagnosisTaskRequest struct {
	// example:
	//
	// https://bucket-cn-hangzhou.oss-cn-hangzhou.aliyuncs.com/debuginfo-common/file/path
	DebuginfoCommonUrl *string `json:"debuginfoCommonUrl,omitempty" xml:"debuginfoCommonUrl,omitempty"`
	// example:
	//
	// https://bucket-cn-hangzhou.oss-cn-hangzhou.aliyuncs.com/debuginfo/file/path
	DebuginfoUrl *string `json:"debuginfoUrl,omitempty" xml:"debuginfoUrl,omitempty"`
	// example:
	//
	// https://bucket-cn-hangzhou.oss-cn-hangzhou.aliyuncs.com/dmesg/file/path
	DmesgUrl *string `json:"dmesgUrl,omitempty" xml:"dmesgUrl,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// vmcore
	TaskType *string `json:"taskType,omitempty" xml:"taskType,omitempty"`
	// example:
	//
	// https://bucket-cn-hangzhou.oss-cn-hangzhou.aliyuncs.com/vmcore/file/path
	VmcoreUrl *string `json:"vmcoreUrl,omitempty" xml:"vmcoreUrl,omitempty"`
}

func (s CreateVmcoreDiagnosisTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateVmcoreDiagnosisTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateVmcoreDiagnosisTaskRequest) GetDebuginfoCommonUrl() *string {
	return s.DebuginfoCommonUrl
}

func (s *CreateVmcoreDiagnosisTaskRequest) GetDebuginfoUrl() *string {
	return s.DebuginfoUrl
}

func (s *CreateVmcoreDiagnosisTaskRequest) GetDmesgUrl() *string {
	return s.DmesgUrl
}

func (s *CreateVmcoreDiagnosisTaskRequest) GetTaskType() *string {
	return s.TaskType
}

func (s *CreateVmcoreDiagnosisTaskRequest) GetVmcoreUrl() *string {
	return s.VmcoreUrl
}

func (s *CreateVmcoreDiagnosisTaskRequest) SetDebuginfoCommonUrl(v string) *CreateVmcoreDiagnosisTaskRequest {
	s.DebuginfoCommonUrl = &v
	return s
}

func (s *CreateVmcoreDiagnosisTaskRequest) SetDebuginfoUrl(v string) *CreateVmcoreDiagnosisTaskRequest {
	s.DebuginfoUrl = &v
	return s
}

func (s *CreateVmcoreDiagnosisTaskRequest) SetDmesgUrl(v string) *CreateVmcoreDiagnosisTaskRequest {
	s.DmesgUrl = &v
	return s
}

func (s *CreateVmcoreDiagnosisTaskRequest) SetTaskType(v string) *CreateVmcoreDiagnosisTaskRequest {
	s.TaskType = &v
	return s
}

func (s *CreateVmcoreDiagnosisTaskRequest) SetVmcoreUrl(v string) *CreateVmcoreDiagnosisTaskRequest {
	s.VmcoreUrl = &v
	return s
}

func (s *CreateVmcoreDiagnosisTaskRequest) Validate() error {
	return dara.Validate(s)
}
