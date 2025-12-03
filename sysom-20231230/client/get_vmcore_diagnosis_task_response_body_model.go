// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVmcoreDiagnosisTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetVmcoreDiagnosisTaskResponseBody
	GetCode() *string
	SetData(v *GetVmcoreDiagnosisTaskResponseBodyData) *GetVmcoreDiagnosisTaskResponseBody
	GetData() *GetVmcoreDiagnosisTaskResponseBodyData
	SetMessage(v string) *GetVmcoreDiagnosisTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetVmcoreDiagnosisTaskResponseBody
	GetRequestId() *string
}

type GetVmcoreDiagnosisTaskResponseBody struct {
	// example:
	//
	// Success
	Code *string                                 `json:"code,omitempty" xml:"code,omitempty"`
	Data *GetVmcoreDiagnosisTaskResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// SysomOpenAPIException: SysomOpenAPI.InvalidParameter Invalid params, should be json string or dict
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 2D693121-C925-5154-8DF6-C09A8B369822
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetVmcoreDiagnosisTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetVmcoreDiagnosisTaskResponseBody) GoString() string {
	return s.String()
}

func (s *GetVmcoreDiagnosisTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetVmcoreDiagnosisTaskResponseBody) GetData() *GetVmcoreDiagnosisTaskResponseBodyData {
	return s.Data
}

func (s *GetVmcoreDiagnosisTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetVmcoreDiagnosisTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetVmcoreDiagnosisTaskResponseBody) SetCode(v string) *GetVmcoreDiagnosisTaskResponseBody {
	s.Code = &v
	return s
}

func (s *GetVmcoreDiagnosisTaskResponseBody) SetData(v *GetVmcoreDiagnosisTaskResponseBodyData) *GetVmcoreDiagnosisTaskResponseBody {
	s.Data = v
	return s
}

func (s *GetVmcoreDiagnosisTaskResponseBody) SetMessage(v string) *GetVmcoreDiagnosisTaskResponseBody {
	s.Message = &v
	return s
}

func (s *GetVmcoreDiagnosisTaskResponseBody) SetRequestId(v string) *GetVmcoreDiagnosisTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetVmcoreDiagnosisTaskResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetVmcoreDiagnosisTaskResponseBodyData struct {
	// example:
	//
	// 2025-12-02T17:36:12
	CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// example:
	//
	// result
	DiagnoseResult *string `json:"diagnoseResult,omitempty" xml:"diagnoseResult,omitempty"`
	// example:
	//
	// error message
	ErrorMsg *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	// example:
	//
	// bbe94a98-4192-4172-b856-95777e0a55d7
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
	// example:
	//
	// running
	TaskStatus *string `json:"taskStatus,omitempty" xml:"taskStatus,omitempty"`
	// example:
	//
	// vmcore
	TaskType *string                                     `json:"taskType,omitempty" xml:"taskType,omitempty"`
	Urls     *GetVmcoreDiagnosisTaskResponseBodyDataUrls `json:"urls,omitempty" xml:"urls,omitempty" type:"Struct"`
}

func (s GetVmcoreDiagnosisTaskResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetVmcoreDiagnosisTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetVmcoreDiagnosisTaskResponseBodyData) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *GetVmcoreDiagnosisTaskResponseBodyData) GetDiagnoseResult() *string {
	return s.DiagnoseResult
}

func (s *GetVmcoreDiagnosisTaskResponseBodyData) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *GetVmcoreDiagnosisTaskResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *GetVmcoreDiagnosisTaskResponseBodyData) GetTaskStatus() *string {
	return s.TaskStatus
}

func (s *GetVmcoreDiagnosisTaskResponseBodyData) GetTaskType() *string {
	return s.TaskType
}

func (s *GetVmcoreDiagnosisTaskResponseBodyData) GetUrls() *GetVmcoreDiagnosisTaskResponseBodyDataUrls {
	return s.Urls
}

func (s *GetVmcoreDiagnosisTaskResponseBodyData) SetCreatedAt(v string) *GetVmcoreDiagnosisTaskResponseBodyData {
	s.CreatedAt = &v
	return s
}

func (s *GetVmcoreDiagnosisTaskResponseBodyData) SetDiagnoseResult(v string) *GetVmcoreDiagnosisTaskResponseBodyData {
	s.DiagnoseResult = &v
	return s
}

func (s *GetVmcoreDiagnosisTaskResponseBodyData) SetErrorMsg(v string) *GetVmcoreDiagnosisTaskResponseBodyData {
	s.ErrorMsg = &v
	return s
}

func (s *GetVmcoreDiagnosisTaskResponseBodyData) SetTaskId(v string) *GetVmcoreDiagnosisTaskResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *GetVmcoreDiagnosisTaskResponseBodyData) SetTaskStatus(v string) *GetVmcoreDiagnosisTaskResponseBodyData {
	s.TaskStatus = &v
	return s
}

func (s *GetVmcoreDiagnosisTaskResponseBodyData) SetTaskType(v string) *GetVmcoreDiagnosisTaskResponseBodyData {
	s.TaskType = &v
	return s
}

func (s *GetVmcoreDiagnosisTaskResponseBodyData) SetUrls(v *GetVmcoreDiagnosisTaskResponseBodyDataUrls) *GetVmcoreDiagnosisTaskResponseBodyData {
	s.Urls = v
	return s
}

func (s *GetVmcoreDiagnosisTaskResponseBodyData) Validate() error {
	if s.Urls != nil {
		if err := s.Urls.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetVmcoreDiagnosisTaskResponseBodyDataUrls struct {
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
	// example:
	//
	// https://bucket-cn-hangzhou.oss-cn-hangzhou.aliyuncs.com/vmcore/file/path
	VmcoreUrl *string `json:"vmcoreUrl,omitempty" xml:"vmcoreUrl,omitempty"`
}

func (s GetVmcoreDiagnosisTaskResponseBodyDataUrls) String() string {
	return dara.Prettify(s)
}

func (s GetVmcoreDiagnosisTaskResponseBodyDataUrls) GoString() string {
	return s.String()
}

func (s *GetVmcoreDiagnosisTaskResponseBodyDataUrls) GetDebuginfoCommonUrl() *string {
	return s.DebuginfoCommonUrl
}

func (s *GetVmcoreDiagnosisTaskResponseBodyDataUrls) GetDebuginfoUrl() *string {
	return s.DebuginfoUrl
}

func (s *GetVmcoreDiagnosisTaskResponseBodyDataUrls) GetDmesgUrl() *string {
	return s.DmesgUrl
}

func (s *GetVmcoreDiagnosisTaskResponseBodyDataUrls) GetVmcoreUrl() *string {
	return s.VmcoreUrl
}

func (s *GetVmcoreDiagnosisTaskResponseBodyDataUrls) SetDebuginfoCommonUrl(v string) *GetVmcoreDiagnosisTaskResponseBodyDataUrls {
	s.DebuginfoCommonUrl = &v
	return s
}

func (s *GetVmcoreDiagnosisTaskResponseBodyDataUrls) SetDebuginfoUrl(v string) *GetVmcoreDiagnosisTaskResponseBodyDataUrls {
	s.DebuginfoUrl = &v
	return s
}

func (s *GetVmcoreDiagnosisTaskResponseBodyDataUrls) SetDmesgUrl(v string) *GetVmcoreDiagnosisTaskResponseBodyDataUrls {
	s.DmesgUrl = &v
	return s
}

func (s *GetVmcoreDiagnosisTaskResponseBodyDataUrls) SetVmcoreUrl(v string) *GetVmcoreDiagnosisTaskResponseBodyDataUrls {
	s.VmcoreUrl = &v
	return s
}

func (s *GetVmcoreDiagnosisTaskResponseBodyDataUrls) Validate() error {
	return dara.Validate(s)
}
