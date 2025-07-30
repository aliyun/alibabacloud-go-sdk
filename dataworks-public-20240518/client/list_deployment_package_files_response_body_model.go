// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDeploymentPackageFilesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPagingInfo(v *ListDeploymentPackageFilesResponseBodyPagingInfo) *ListDeploymentPackageFilesResponseBody
	GetPagingInfo() *ListDeploymentPackageFilesResponseBodyPagingInfo
	SetRequestId(v string) *ListDeploymentPackageFilesResponseBody
	GetRequestId() *string
}

type ListDeploymentPackageFilesResponseBody struct {
	// The pagination details.
	PagingInfo *ListDeploymentPackageFilesResponseBodyPagingInfo `json:"PagingInfo,omitempty" xml:"PagingInfo,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 0000-ABCD-EFG****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListDeploymentPackageFilesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDeploymentPackageFilesResponseBody) GoString() string {
	return s.String()
}

func (s *ListDeploymentPackageFilesResponseBody) GetPagingInfo() *ListDeploymentPackageFilesResponseBodyPagingInfo {
	return s.PagingInfo
}

func (s *ListDeploymentPackageFilesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDeploymentPackageFilesResponseBody) SetPagingInfo(v *ListDeploymentPackageFilesResponseBodyPagingInfo) *ListDeploymentPackageFilesResponseBody {
	s.PagingInfo = v
	return s
}

func (s *ListDeploymentPackageFilesResponseBody) SetRequestId(v string) *ListDeploymentPackageFilesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDeploymentPackageFilesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListDeploymentPackageFilesResponseBodyPagingInfo struct {
	// The list of files pending deployment.
	DeploymentPackageFiles []*ListDeploymentPackageFilesResponseBodyPagingInfoDeploymentPackageFiles `json:"DeploymentPackageFiles,omitempty" xml:"DeploymentPackageFiles,omitempty" type:"Repeated"`
	// The page number. Pages start from page 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListDeploymentPackageFilesResponseBodyPagingInfo) String() string {
	return dara.Prettify(s)
}

func (s ListDeploymentPackageFilesResponseBodyPagingInfo) GoString() string {
	return s.String()
}

func (s *ListDeploymentPackageFilesResponseBodyPagingInfo) GetDeploymentPackageFiles() []*ListDeploymentPackageFilesResponseBodyPagingInfoDeploymentPackageFiles {
	return s.DeploymentPackageFiles
}

func (s *ListDeploymentPackageFilesResponseBodyPagingInfo) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListDeploymentPackageFilesResponseBodyPagingInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDeploymentPackageFilesResponseBodyPagingInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListDeploymentPackageFilesResponseBodyPagingInfo) SetDeploymentPackageFiles(v []*ListDeploymentPackageFilesResponseBodyPagingInfoDeploymentPackageFiles) *ListDeploymentPackageFilesResponseBodyPagingInfo {
	s.DeploymentPackageFiles = v
	return s
}

func (s *ListDeploymentPackageFilesResponseBodyPagingInfo) SetPageNumber(v int32) *ListDeploymentPackageFilesResponseBodyPagingInfo {
	s.PageNumber = &v
	return s
}

func (s *ListDeploymentPackageFilesResponseBodyPagingInfo) SetPageSize(v int32) *ListDeploymentPackageFilesResponseBodyPagingInfo {
	s.PageSize = &v
	return s
}

func (s *ListDeploymentPackageFilesResponseBodyPagingInfo) SetTotalCount(v int32) *ListDeploymentPackageFilesResponseBodyPagingInfo {
	s.TotalCount = &v
	return s
}

func (s *ListDeploymentPackageFilesResponseBodyPagingInfo) Validate() error {
	return dara.Validate(s)
}

type ListDeploymentPackageFilesResponseBodyPagingInfoDeploymentPackageFiles struct {
	// The change type, which is an integer. Valid values:
	//
	// 	- 0: addition
	//
	// 	- 1: update
	//
	// 	- 2: deletion
	//
	// example:
	//
	// 0
	ChangeType *int32 `json:"ChangeType,omitempty" xml:"ChangeType,omitempty"`
	// The comment for committing.
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The time for committing.
	//
	// example:
	//
	// 2025-04-10 15:55:47
	CommitTime *string `json:"CommitTime,omitempty" xml:"CommitTime,omitempty"`
	// The ID of the Alibaba Cloud account used by the user who committed the file.
	//
	// example:
	//
	// 446***
	CommitUser *string `json:"CommitUser,omitempty" xml:"CommitUser,omitempty"`
	// The name of the Alibaba Cloud account used by the user who committed the file.
	//
	// example:
	//
	// user***
	CommitUserName *string `json:"CommitUserName,omitempty" xml:"CommitUserName,omitempty"`
	// The file ID.
	//
	// example:
	//
	// 520246913
	FileId *int64 `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// The name of the file of the current version.
	//
	// example:
	//
	// bak_part_basc_person_relation_all_da
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// The file type. The code for files varies based on the file type. For more information, see [DataWorks nodes](https://help.aliyun.com/document_detail/600169.html).
	//
	// example:
	//
	// 13
	FileType *int32 `json:"FileType,omitempty" xml:"FileType,omitempty"`
	// The file version.
	//
	// example:
	//
	// 34
	FileVersion *int64 `json:"FileVersion,omitempty" xml:"FileVersion,omitempty"`
	// The unique ID.
	//
	// example:
	//
	// 650433503
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// Indicates whether the version is a version in the production environment of the scheduling system.
	//
	// example:
	//
	// true
	IsSameAsProductionVersion *bool `json:"IsSameAsProductionVersion,omitempty" xml:"IsSameAsProductionVersion,omitempty"`
	// The scheduling property configurations of the node that corresponds to the file, which is a JSON string.
	//
	// example:
	//
	// {
	//
	// 	"tagList": [],
	//
	// 	"fileId": -1,
	//
	// 	"taskRerunTime": 0,
	//
	// 	"taskRerunInterval": 0,
	//
	// 	"reRunAble": 1,
	//
	// 	"nodeId": 125803000,
	//
	// 	"nodeName": "new",
	//
	// 	"nodeType": 0,
	//
	// 	"isStop": 0,
	//
	// 	"paraValue": "",
	//
	// 	"startEffectDate": "1970-01-01 00:00:00",
	//
	// 	"endEffectDate": "9999-01-01 00:00:00",
	//
	// 	"cronExpress": "00 26 00 	- 	- ?",
	//
	// 	"owner": "1107550004250000",
	//
	// 	"resgroupId": 6300000,
	//
	// 	"cu": "0.25",
	//
	// 	"appId": 170000,
	//
	// 	"tenantId": 524257424560000,
	//
	// 	"createTime": "2025-04-10 15:55:01",
	//
	// 	"createUser": "1107550004250000",
	//
	// 	"lastModifyTime": "2025-04-10 15:55:41",
	//
	// 	"cycleType": 0,
	//
	// 	"dependentType": 0,
	//
	// 	"dependentTypeList": [0],
	//
	// 	"lastModifyUser": "1107550004250000",
	//
	// 	"dependentDataNode": "",
	//
	// 	"input": "[{\\"regionId\\":\\"cn-hangzhou\\",\\"str\\":\\"root_input\\",\\"parseType\\":1}]",
	//
	// 	"output": "[{\\"str\\":\\"project_root.526586287_out\\",\\"parseType\\":2},{\\"str\\":\\"project_root.new\\",\\"parseType\\":1}]",
	//
	// 	"inputList": [{
	//
	// 		"regionId": "cn-hangzhou",
	//
	// 		"str": "root_input",
	//
	// 		"parseType": 1
	//
	// 	}],
	//
	// 	"outputList": [{
	//
	// 		"str": "project_root.526586287_out",
	//
	// 		"parseType": 2
	//
	// 	}, {
	//
	// 		"str": "project_root.new",
	//
	// 		"parseType": 1
	//
	// 	}],
	//
	// 	"isAutoParse": 1,
	//
	// 	"startRightNow": false,
	//
	// 	"extConfig": "{\\"openCustomCron\\":false,\\"formCron\\":\\"\\"}",
	//
	// 	"inputContextList": [],
	//
	// 	"outputContextList": []
	//
	// }
	NodeConfiguration *string `json:"NodeConfiguration,omitempty" xml:"NodeConfiguration,omitempty"`
	// The ID of the auto triggered node that corresponds to the file.
	//
	// example:
	//
	// 700005008419
	NodeId *int64 `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The workspace ID.
	//
	// example:
	//
	// 27595
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The test status in the development environment.
	SmokeTestStatus *string `json:"SmokeTestStatus,omitempty" xml:"SmokeTestStatus,omitempty"`
	// The status of the code file of the current version. Valid values:
	//
	// 	- 2: Commit check in progress.
	//
	// 	- 3: Commit check passed.
	//
	// 	- 4: Commit check failed.
	//
	// 	- 10: Committing.
	//
	// 	- 11: Committed.
	//
	// 	- 20: Approved.
	//
	// 	- 21: Rejected.
	//
	// 	- 22: Warning detected during checking.
	//
	// 	- 23: Under code review.
	//
	// 	- 24: Code review rejected.
	//
	// 	- 80: Deployment package created.
	//
	// 	- 100: Deploying.
	//
	// 	- 101: Deployed to the production environment.
	//
	// 	- 200: Cancelled.
	//
	// example:
	//
	// 100
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The DataWorks tenant ID.
	//
	// example:
	//
	// 639415964191360
	TenantId *int64 `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// The module to which the file belongs. Valid values:
	//
	// 	- NORMAL: The file is used for DataStudio.
	//
	// 	- MANUAL: The file is used for a manually triggered node.
	//
	// 	- MANUAL_BIZ: The file is used for a manually triggered workflow.
	//
	// 	- SKIP: The file is used for a dry-run node in DataStudio.
	//
	// 	- ADHOCQUERY: The file is used for an ad hoc query.
	//
	// 	- COMPONENT: The file is used for a script template.
	//
	// example:
	//
	// NORMAL
	UseType *string `json:"UseType,omitempty" xml:"UseType,omitempty"`
}

func (s ListDeploymentPackageFilesResponseBodyPagingInfoDeploymentPackageFiles) String() string {
	return dara.Prettify(s)
}

func (s ListDeploymentPackageFilesResponseBodyPagingInfoDeploymentPackageFiles) GoString() string {
	return s.String()
}

func (s *ListDeploymentPackageFilesResponseBodyPagingInfoDeploymentPackageFiles) GetChangeType() *int32 {
	return s.ChangeType
}

func (s *ListDeploymentPackageFilesResponseBodyPagingInfoDeploymentPackageFiles) GetComment() *string {
	return s.Comment
}

func (s *ListDeploymentPackageFilesResponseBodyPagingInfoDeploymentPackageFiles) GetCommitTime() *string {
	return s.CommitTime
}

func (s *ListDeploymentPackageFilesResponseBodyPagingInfoDeploymentPackageFiles) GetCommitUser() *string {
	return s.CommitUser
}

func (s *ListDeploymentPackageFilesResponseBodyPagingInfoDeploymentPackageFiles) GetCommitUserName() *string {
	return s.CommitUserName
}

func (s *ListDeploymentPackageFilesResponseBodyPagingInfoDeploymentPackageFiles) GetFileId() *int64 {
	return s.FileId
}

func (s *ListDeploymentPackageFilesResponseBodyPagingInfoDeploymentPackageFiles) GetFileName() *string {
	return s.FileName
}

func (s *ListDeploymentPackageFilesResponseBodyPagingInfoDeploymentPackageFiles) GetFileType() *int32 {
	return s.FileType
}

func (s *ListDeploymentPackageFilesResponseBodyPagingInfoDeploymentPackageFiles) GetFileVersion() *int64 {
	return s.FileVersion
}

func (s *ListDeploymentPackageFilesResponseBodyPagingInfoDeploymentPackageFiles) GetId() *int64 {
	return s.Id
}

func (s *ListDeploymentPackageFilesResponseBodyPagingInfoDeploymentPackageFiles) GetIsSameAsProductionVersion() *bool {
	return s.IsSameAsProductionVersion
}

func (s *ListDeploymentPackageFilesResponseBodyPagingInfoDeploymentPackageFiles) GetNodeConfiguration() *string {
	return s.NodeConfiguration
}

func (s *ListDeploymentPackageFilesResponseBodyPagingInfoDeploymentPackageFiles) GetNodeId() *int64 {
	return s.NodeId
}

func (s *ListDeploymentPackageFilesResponseBodyPagingInfoDeploymentPackageFiles) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListDeploymentPackageFilesResponseBodyPagingInfoDeploymentPackageFiles) GetSmokeTestStatus() *string {
	return s.SmokeTestStatus
}

func (s *ListDeploymentPackageFilesResponseBodyPagingInfoDeploymentPackageFiles) GetStatus() *int32 {
	return s.Status
}

func (s *ListDeploymentPackageFilesResponseBodyPagingInfoDeploymentPackageFiles) GetTenantId() *int64 {
	return s.TenantId
}

func (s *ListDeploymentPackageFilesResponseBodyPagingInfoDeploymentPackageFiles) GetUseType() *string {
	return s.UseType
}

func (s *ListDeploymentPackageFilesResponseBodyPagingInfoDeploymentPackageFiles) SetChangeType(v int32) *ListDeploymentPackageFilesResponseBodyPagingInfoDeploymentPackageFiles {
	s.ChangeType = &v
	return s
}

func (s *ListDeploymentPackageFilesResponseBodyPagingInfoDeploymentPackageFiles) SetComment(v string) *ListDeploymentPackageFilesResponseBodyPagingInfoDeploymentPackageFiles {
	s.Comment = &v
	return s
}

func (s *ListDeploymentPackageFilesResponseBodyPagingInfoDeploymentPackageFiles) SetCommitTime(v string) *ListDeploymentPackageFilesResponseBodyPagingInfoDeploymentPackageFiles {
	s.CommitTime = &v
	return s
}

func (s *ListDeploymentPackageFilesResponseBodyPagingInfoDeploymentPackageFiles) SetCommitUser(v string) *ListDeploymentPackageFilesResponseBodyPagingInfoDeploymentPackageFiles {
	s.CommitUser = &v
	return s
}

func (s *ListDeploymentPackageFilesResponseBodyPagingInfoDeploymentPackageFiles) SetCommitUserName(v string) *ListDeploymentPackageFilesResponseBodyPagingInfoDeploymentPackageFiles {
	s.CommitUserName = &v
	return s
}

func (s *ListDeploymentPackageFilesResponseBodyPagingInfoDeploymentPackageFiles) SetFileId(v int64) *ListDeploymentPackageFilesResponseBodyPagingInfoDeploymentPackageFiles {
	s.FileId = &v
	return s
}

func (s *ListDeploymentPackageFilesResponseBodyPagingInfoDeploymentPackageFiles) SetFileName(v string) *ListDeploymentPackageFilesResponseBodyPagingInfoDeploymentPackageFiles {
	s.FileName = &v
	return s
}

func (s *ListDeploymentPackageFilesResponseBodyPagingInfoDeploymentPackageFiles) SetFileType(v int32) *ListDeploymentPackageFilesResponseBodyPagingInfoDeploymentPackageFiles {
	s.FileType = &v
	return s
}

func (s *ListDeploymentPackageFilesResponseBodyPagingInfoDeploymentPackageFiles) SetFileVersion(v int64) *ListDeploymentPackageFilesResponseBodyPagingInfoDeploymentPackageFiles {
	s.FileVersion = &v
	return s
}

func (s *ListDeploymentPackageFilesResponseBodyPagingInfoDeploymentPackageFiles) SetId(v int64) *ListDeploymentPackageFilesResponseBodyPagingInfoDeploymentPackageFiles {
	s.Id = &v
	return s
}

func (s *ListDeploymentPackageFilesResponseBodyPagingInfoDeploymentPackageFiles) SetIsSameAsProductionVersion(v bool) *ListDeploymentPackageFilesResponseBodyPagingInfoDeploymentPackageFiles {
	s.IsSameAsProductionVersion = &v
	return s
}

func (s *ListDeploymentPackageFilesResponseBodyPagingInfoDeploymentPackageFiles) SetNodeConfiguration(v string) *ListDeploymentPackageFilesResponseBodyPagingInfoDeploymentPackageFiles {
	s.NodeConfiguration = &v
	return s
}

func (s *ListDeploymentPackageFilesResponseBodyPagingInfoDeploymentPackageFiles) SetNodeId(v int64) *ListDeploymentPackageFilesResponseBodyPagingInfoDeploymentPackageFiles {
	s.NodeId = &v
	return s
}

func (s *ListDeploymentPackageFilesResponseBodyPagingInfoDeploymentPackageFiles) SetProjectId(v int64) *ListDeploymentPackageFilesResponseBodyPagingInfoDeploymentPackageFiles {
	s.ProjectId = &v
	return s
}

func (s *ListDeploymentPackageFilesResponseBodyPagingInfoDeploymentPackageFiles) SetSmokeTestStatus(v string) *ListDeploymentPackageFilesResponseBodyPagingInfoDeploymentPackageFiles {
	s.SmokeTestStatus = &v
	return s
}

func (s *ListDeploymentPackageFilesResponseBodyPagingInfoDeploymentPackageFiles) SetStatus(v int32) *ListDeploymentPackageFilesResponseBodyPagingInfoDeploymentPackageFiles {
	s.Status = &v
	return s
}

func (s *ListDeploymentPackageFilesResponseBodyPagingInfoDeploymentPackageFiles) SetTenantId(v int64) *ListDeploymentPackageFilesResponseBodyPagingInfoDeploymentPackageFiles {
	s.TenantId = &v
	return s
}

func (s *ListDeploymentPackageFilesResponseBodyPagingInfoDeploymentPackageFiles) SetUseType(v string) *ListDeploymentPackageFilesResponseBodyPagingInfoDeploymentPackageFiles {
	s.UseType = &v
	return s
}

func (s *ListDeploymentPackageFilesResponseBodyPagingInfoDeploymentPackageFiles) Validate() error {
	return dara.Validate(s)
}
