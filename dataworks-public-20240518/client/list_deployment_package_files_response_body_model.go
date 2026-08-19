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
	// The pagination information.
	PagingInfo *ListDeploymentPackageFilesResponseBodyPagingInfo `json:"PagingInfo,omitempty" xml:"PagingInfo,omitempty" type:"Struct"`
	// The request ID. You can use this ID to troubleshoot issues.
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
	if s.PagingInfo != nil {
		if err := s.PagingInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListDeploymentPackageFilesResponseBodyPagingInfo struct {
	// The list of file versions pending deployment.
	DeploymentPackageFiles []*ListDeploymentPackageFilesResponseBodyPagingInfoDeploymentPackageFiles `json:"DeploymentPackageFiles,omitempty" xml:"DeploymentPackageFiles,omitempty" type:"Repeated"`
	// The page number, starting from 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The page size. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries that meet the conditions.
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
	if s.DeploymentPackageFiles != nil {
		for _, item := range s.DeploymentPackageFiles {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListDeploymentPackageFilesResponseBodyPagingInfoDeploymentPackageFiles struct {
	// The change type. Valid values:
	//
	// - 0: added.
	//
	// - 1: updated.
	//
	// - 2: deleted.
	//
	// example:
	//
	// 0
	ChangeType *int32 `json:"ChangeType,omitempty" xml:"ChangeType,omitempty"`
	// The comment provided at the time of commit.
	//
	// example:
	//
	// Test commit
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The commit time.
	//
	// The format is `yyyy-MM-dd HH:mm:ss`, for example, `2025-04-10 15:55:47`. This example does not include a time zone identifier.
	//
	// example:
	//
	// 2025-04-10 15:55:47
	CommitTime *string `json:"CommitTime,omitempty" xml:"CommitTime,omitempty"`
	// The Alibaba Cloud account ID of the committer.
	//
	// example:
	//
	// 446***
	CommitUser *string `json:"CommitUser,omitempty" xml:"CommitUser,omitempty"`
	// The Alibaba Cloud account name of the committer.
	//
	// example:
	//
	// user***
	CommitUserName *string `json:"CommitUserName,omitempty" xml:"CommitUserName,omitempty"`
	// The ID of the file.
	//
	// example:
	//
	// 520246913
	FileId *int64 `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// The name of the file that generated this file version.
	//
	// example:
	//
	// bak_part_basc_person_relation_all_da
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// The file type. Different file types have different codes. For more information, see [DataWorks nodes](https://help.aliyun.com/document_detail/600169.html).
	//
	// example:
	//
	// 13
	FileType *int32 `json:"FileType,omitempty" xml:"FileType,omitempty"`
	// The version number of the file.
	//
	// example:
	//
	// 34
	FileVersion *int64 `json:"FileVersion,omitempty" xml:"FileVersion,omitempty"`
	// The unique identifier.
	//
	// example:
	//
	// 650433503
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// Indicates whether this version is the same as the current production version in scheduling.
	//
	// example:
	//
	// true
	IsSameAsProductionVersion *bool `json:"IsSameAsProductionVersion,omitempty" xml:"IsSameAsProductionVersion,omitempty"`
	// The scheduling property configuration of the scheduling node to which this file belongs, stored as a JSON string.
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
	// The node ID in scheduling that corresponds to this file.
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
	// The testing status in the development environment.
	//
	// example:
	//
	// Not tested
	SmokeTestStatus *string `json:"SmokeTestStatus,omitempty" xml:"SmokeTestStatus,omitempty"`
	// The status of the code file for this version. Valid values:
	//
	// - 2: commit check in progress.
	//
	// - 3: commit check succeeded.
	//
	// - 4: commit check rejected.
	//
	// - 10: committing.
	//
	// - 11: committed to the scheduling development environment.
	//
	// - 20: review approved.
	//
	// - 21: review failed.
	//
	// - 22: check has warnings.
	//
	// - 23: code review in progress.
	//
	// - 24: code review rejected.
	//
	// - 80: deployment package created.
	//
	// - 100: deploying.
	//
	// - 101: deployed to production.
	//
	// - 200: canceled.
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
	// The functional module to which the file belongs. Valid values:
	//
	// - NORMAL: data development.
	//
	// - MANUAL: manual task.
	//
	// - MANUAL_BIZ: manual workflow.
	//
	// - SKIP: dry-run scheduling in data development.
	//
	// - ADHOCQUERY: ad hoc query.
	//
	// - COMPONENT: component management.
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
