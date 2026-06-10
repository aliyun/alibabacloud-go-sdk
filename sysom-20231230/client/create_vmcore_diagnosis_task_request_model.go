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
	// Download link for the debuginfo-common file. This parameter is optional when the diagnosis type is vmcore.
	//
	// For CentOS or Alinux kernel diagnosis, the corresponding debuginfo-common file is automatically downloaded, so you do not need to provide this parameter. For kernels of other distributions, you must manually provide the download link for the debuginfo-common file that matches the kernel version.
	//
	// example:
	//
	// https://bucket-cn-hangzhou.oss-cn-hangzhou.aliyuncs.com/debuginfo-common/file/path
	DebuginfoCommonUrl *string `json:"debuginfoCommonUrl,omitempty" xml:"debuginfoCommonUrl,omitempty"`
	// The download link of the debuginfo file corresponding to the vmcore file. This parameter is optional when the diagnosis type is vmcore.
	//
	// For CentOS or Alinux kernel diagnosis, the corresponding debuginfo file is automatically downloaded, so you do not need to provide this parameter. For kernels from other distributions, you must manually provide the download link for the debuginfo file that matches the kernel version.
	//
	// example:
	//
	// https://bucket-cn-hangzhou.oss-cn-hangzhou.aliyuncs.com/debuginfo/file/path
	DebuginfoUrl *string `json:"debuginfoUrl,omitempty" xml:"debuginfoUrl,omitempty"`
	// Download link for the dmesg log file. This parameter is required when the diagnosis type is dmesg.
	//
	// example:
	//
	// https://bucket-cn-hangzhou.oss-cn-hangzhou.aliyuncs.com/dmesg/file/path
	DmesgUrl *string `json:"dmesgUrl,omitempty" xml:"dmesgUrl,omitempty"`
	// Task Type
	//
	// vmcore: vmcore file diagnosis task
	//
	// dmesg: dmesg log diagnosis task
	//
	// This parameter is required.
	//
	// example:
	//
	// vmcore
	TaskType *string `json:"taskType,omitempty" xml:"taskType,omitempty"`
	// The download link of the vmcore file. This parameter is required when the diagnosis type is vmcore.
	//
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
