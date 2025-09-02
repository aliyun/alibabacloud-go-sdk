// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEnabledExtensionsForProjectRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEventCode(v string) *ListEnabledExtensionsForProjectRequest
	GetEventCode() *string
	SetFileType(v string) *ListEnabledExtensionsForProjectRequest
	GetFileType() *string
	SetProjectId(v int64) *ListEnabledExtensionsForProjectRequest
	GetProjectId() *int64
}

type ListEnabledExtensionsForProjectRequest struct {
	// The code of the extension point event.
	//
	// This parameter is required.
	//
	// example:
	//
	// commit-file
	EventCode *string `json:"EventCode,omitempty" xml:"EventCode,omitempty"`
	// The type of the code for the file.
	//
	// Valid values: 6 (Shell), 10 (ODPS SQL), 11 (ODPS MR), 24 (ODPS Script), 99 (zero load), 221 (PyODPS 2), 225 (ODPS Spark), 227 (EMR Hive), 228 (EMR Spark), 229 (EMR Spark SQL), 230 (EMR MR), 239 (OSS object inspection), 257 (EMR Shell), 258 (EMR Spark Shell), 259 (EMR Presto), 260 (EMR Impala), 900 (real-time synchronization), 1089 (cross-tenant collaboration), 1091 (Hologres development), 1093 (Hologres SQL), 1100 (assignment), and 1221 (PyODPS 3).
	//
	// You can call the [ListFileType](https://help.aliyun.com/document_detail/212428.html) operation to query the type of the code for the file.
	//
	// example:
	//
	// 10
	FileType *string `json:"FileType,omitempty" xml:"FileType,omitempty"`
	// The DataWorks workspace ID. You can log on to the [DataWorks console](https://workbench.data.aliyun.com/console?spm=a2c4g.11186623.0.0.6b4d4941azHd2k) and go to the Workspace page to obtain the workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10000
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s ListEnabledExtensionsForProjectRequest) String() string {
	return dara.Prettify(s)
}

func (s ListEnabledExtensionsForProjectRequest) GoString() string {
	return s.String()
}

func (s *ListEnabledExtensionsForProjectRequest) GetEventCode() *string {
	return s.EventCode
}

func (s *ListEnabledExtensionsForProjectRequest) GetFileType() *string {
	return s.FileType
}

func (s *ListEnabledExtensionsForProjectRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListEnabledExtensionsForProjectRequest) SetEventCode(v string) *ListEnabledExtensionsForProjectRequest {
	s.EventCode = &v
	return s
}

func (s *ListEnabledExtensionsForProjectRequest) SetFileType(v string) *ListEnabledExtensionsForProjectRequest {
	s.FileType = &v
	return s
}

func (s *ListEnabledExtensionsForProjectRequest) SetProjectId(v int64) *ListEnabledExtensionsForProjectRequest {
	s.ProjectId = &v
	return s
}

func (s *ListEnabledExtensionsForProjectRequest) Validate() error {
	return dara.Validate(s)
}
