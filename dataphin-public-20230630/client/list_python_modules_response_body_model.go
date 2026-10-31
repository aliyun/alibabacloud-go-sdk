// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPythonModulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListPythonModulesResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ListPythonModulesResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListPythonModulesResponseBody
	GetMessage() *string
	SetPageResult(v *ListPythonModulesResponseBodyPageResult) *ListPythonModulesResponseBody
	GetPageResult() *ListPythonModulesResponseBodyPageResult
	SetRequestId(v string) *ListPythonModulesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListPythonModulesResponseBody
	GetSuccess() *bool
}

type ListPythonModulesResponseBody struct {
	// The backend response code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The backend exception details.
	//
	// example:
	//
	// internal error
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The paged query result with paging details.
	PageResult *ListPythonModulesResponseBodyPageResult `json:"PageResult,omitempty" xml:"PageResult,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 82E78D6B-AA8F-1FEF-8AA3-5C9DA2A79140
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListPythonModulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListPythonModulesResponseBody) GoString() string {
	return s.String()
}

func (s *ListPythonModulesResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListPythonModulesResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListPythonModulesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListPythonModulesResponseBody) GetPageResult() *ListPythonModulesResponseBodyPageResult {
	return s.PageResult
}

func (s *ListPythonModulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListPythonModulesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListPythonModulesResponseBody) SetCode(v string) *ListPythonModulesResponseBody {
	s.Code = &v
	return s
}

func (s *ListPythonModulesResponseBody) SetHttpStatusCode(v int32) *ListPythonModulesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListPythonModulesResponseBody) SetMessage(v string) *ListPythonModulesResponseBody {
	s.Message = &v
	return s
}

func (s *ListPythonModulesResponseBody) SetPageResult(v *ListPythonModulesResponseBodyPageResult) *ListPythonModulesResponseBody {
	s.PageResult = v
	return s
}

func (s *ListPythonModulesResponseBody) SetRequestId(v string) *ListPythonModulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPythonModulesResponseBody) SetSuccess(v bool) *ListPythonModulesResponseBody {
	s.Success = &v
	return s
}

func (s *ListPythonModulesResponseBody) Validate() error {
	if s.PageResult != nil {
		if err := s.PageResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListPythonModulesResponseBodyPageResult struct {
	// The paginated list of Python third-party packages.
	PythonModuleList []*ListPythonModulesResponseBodyPageResultPythonModuleList `json:"PythonModuleList,omitempty" xml:"PythonModuleList,omitempty" type:"Repeated"`
	// The total number of entries.
	//
	// example:
	//
	// 100
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListPythonModulesResponseBodyPageResult) String() string {
	return dara.Prettify(s)
}

func (s ListPythonModulesResponseBodyPageResult) GoString() string {
	return s.String()
}

func (s *ListPythonModulesResponseBodyPageResult) GetPythonModuleList() []*ListPythonModulesResponseBodyPageResultPythonModuleList {
	return s.PythonModuleList
}

func (s *ListPythonModulesResponseBodyPageResult) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListPythonModulesResponseBodyPageResult) SetPythonModuleList(v []*ListPythonModulesResponseBodyPageResultPythonModuleList) *ListPythonModulesResponseBodyPageResult {
	s.PythonModuleList = v
	return s
}

func (s *ListPythonModulesResponseBodyPageResult) SetTotalCount(v int64) *ListPythonModulesResponseBodyPageResult {
	s.TotalCount = &v
	return s
}

func (s *ListPythonModulesResponseBodyPageResult) Validate() error {
	if s.PythonModuleList != nil {
		for _, item := range s.PythonModuleList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListPythonModulesResponseBodyPageResultPythonModuleList struct {
	// The creation time.
	//
	// example:
	//
	// 2026-01-01 10:00:00
	GmtCreated *string `json:"GmtCreated,omitempty" xml:"GmtCreated,omitempty"`
	// The last update time.
	//
	// example:
	//
	// 2026-01-01 10:00:00
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The ID of the Python third-party package.
	//
	// example:
	//
	// 1001
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The installation time.
	//
	// example:
	//
	// 2026-01-01 10:00:00
	InstallTime *string `json:"InstallTime,omitempty" xml:"InstallTime,omitempty"`
	// The installation method. Valid values:
	//
	// - UPLOAD_INSTALL: upload installation package
	//
	// - ONLINE_INSTALL: online installation
	//
	// example:
	//
	// UPLOAD_INSTALL
	InstallType *string `json:"InstallType,omitempty" xml:"InstallType,omitempty"`
	// The installer.
	//
	// example:
	//
	// 30001011
	Installer *string `json:"Installer,omitempty" xml:"Installer,omitempty"`
	// The name of the Python third-party package.
	//
	// example:
	//
	// numpy
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The Python version. Valid values: PYTHON_2_7, PYTHON_3_7, and PYTHON_3_11.
	//
	// example:
	//
	// PYTHON_3_7
	PythonVersion *string `json:"PythonVersion,omitempty" xml:"PythonVersion,omitempty"`
	// The status of the third-party package. Valid values:
	//
	// - INSTALL_SUCCESS: installation succeeded
	//
	// - INSTALL_FAILED: installation failed
	//
	// - INSTALLING: installing
	//
	// - UPGRADE_SUCCESS: upgrade succeeded
	//
	// - UPGRADE_FAILED: upgrade failed
	//
	// - UPGRADING: upgrading
	//
	// example:
	//
	// INSTALL_SUCCESS
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The version of the third-party package.
	//
	// example:
	//
	// 1.0.1
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s ListPythonModulesResponseBodyPageResultPythonModuleList) String() string {
	return dara.Prettify(s)
}

func (s ListPythonModulesResponseBodyPageResultPythonModuleList) GoString() string {
	return s.String()
}

func (s *ListPythonModulesResponseBodyPageResultPythonModuleList) GetGmtCreated() *string {
	return s.GmtCreated
}

func (s *ListPythonModulesResponseBodyPageResultPythonModuleList) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ListPythonModulesResponseBodyPageResultPythonModuleList) GetId() *int64 {
	return s.Id
}

func (s *ListPythonModulesResponseBodyPageResultPythonModuleList) GetInstallTime() *string {
	return s.InstallTime
}

func (s *ListPythonModulesResponseBodyPageResultPythonModuleList) GetInstallType() *string {
	return s.InstallType
}

func (s *ListPythonModulesResponseBodyPageResultPythonModuleList) GetInstaller() *string {
	return s.Installer
}

func (s *ListPythonModulesResponseBodyPageResultPythonModuleList) GetName() *string {
	return s.Name
}

func (s *ListPythonModulesResponseBodyPageResultPythonModuleList) GetPythonVersion() *string {
	return s.PythonVersion
}

func (s *ListPythonModulesResponseBodyPageResultPythonModuleList) GetStatus() *string {
	return s.Status
}

func (s *ListPythonModulesResponseBodyPageResultPythonModuleList) GetVersion() *string {
	return s.Version
}

func (s *ListPythonModulesResponseBodyPageResultPythonModuleList) SetGmtCreated(v string) *ListPythonModulesResponseBodyPageResultPythonModuleList {
	s.GmtCreated = &v
	return s
}

func (s *ListPythonModulesResponseBodyPageResultPythonModuleList) SetGmtModified(v string) *ListPythonModulesResponseBodyPageResultPythonModuleList {
	s.GmtModified = &v
	return s
}

func (s *ListPythonModulesResponseBodyPageResultPythonModuleList) SetId(v int64) *ListPythonModulesResponseBodyPageResultPythonModuleList {
	s.Id = &v
	return s
}

func (s *ListPythonModulesResponseBodyPageResultPythonModuleList) SetInstallTime(v string) *ListPythonModulesResponseBodyPageResultPythonModuleList {
	s.InstallTime = &v
	return s
}

func (s *ListPythonModulesResponseBodyPageResultPythonModuleList) SetInstallType(v string) *ListPythonModulesResponseBodyPageResultPythonModuleList {
	s.InstallType = &v
	return s
}

func (s *ListPythonModulesResponseBodyPageResultPythonModuleList) SetInstaller(v string) *ListPythonModulesResponseBodyPageResultPythonModuleList {
	s.Installer = &v
	return s
}

func (s *ListPythonModulesResponseBodyPageResultPythonModuleList) SetName(v string) *ListPythonModulesResponseBodyPageResultPythonModuleList {
	s.Name = &v
	return s
}

func (s *ListPythonModulesResponseBodyPageResultPythonModuleList) SetPythonVersion(v string) *ListPythonModulesResponseBodyPageResultPythonModuleList {
	s.PythonVersion = &v
	return s
}

func (s *ListPythonModulesResponseBodyPageResultPythonModuleList) SetStatus(v string) *ListPythonModulesResponseBodyPageResultPythonModuleList {
	s.Status = &v
	return s
}

func (s *ListPythonModulesResponseBodyPageResultPythonModuleList) SetVersion(v string) *ListPythonModulesResponseBodyPageResultPythonModuleList {
	s.Version = &v
	return s
}

func (s *ListPythonModulesResponseBodyPageResultPythonModuleList) Validate() error {
	return dara.Validate(s)
}
