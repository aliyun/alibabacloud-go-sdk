// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeModelingProjectDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *DescribeModelingProjectDetailResponseBody
	GetCode() *int64
	SetRequestId(v string) *DescribeModelingProjectDetailResponseBody
	GetRequestId() *string
	SetResultObject(v *DescribeModelingProjectDetailResponseBodyResultObject) *DescribeModelingProjectDetailResponseBody
	GetResultObject() *DescribeModelingProjectDetailResponseBodyResultObject
	SetSuccess(v bool) *DescribeModelingProjectDetailResponseBody
	GetSuccess() *bool
}

type DescribeModelingProjectDetailResponseBody struct {
	// Status code. A return value of 200 indicates success.
	//
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// Request ID.
	//
	// example:
	//
	// 4A91D2D1-AEC9-1658-ABCE-5A39DE66A5C2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Returned result.
	ResultObject *DescribeModelingProjectDetailResponseBodyResultObject `json:"ResultObject,omitempty" xml:"ResultObject,omitempty" type:"Struct"`
	// Whether the call was successful.
	//
	// - **true**: Call succeeded.
	//
	// - **false**: Call failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeModelingProjectDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeModelingProjectDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeModelingProjectDetailResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *DescribeModelingProjectDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeModelingProjectDetailResponseBody) GetResultObject() *DescribeModelingProjectDetailResponseBodyResultObject {
	return s.ResultObject
}

func (s *DescribeModelingProjectDetailResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeModelingProjectDetailResponseBody) SetCode(v int64) *DescribeModelingProjectDetailResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeModelingProjectDetailResponseBody) SetRequestId(v string) *DescribeModelingProjectDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeModelingProjectDetailResponseBody) SetResultObject(v *DescribeModelingProjectDetailResponseBodyResultObject) *DescribeModelingProjectDetailResponseBody {
	s.ResultObject = v
	return s
}

func (s *DescribeModelingProjectDetailResponseBody) SetSuccess(v bool) *DescribeModelingProjectDetailResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeModelingProjectDetailResponseBody) Validate() error {
	if s.ResultObject != nil {
		if err := s.ResultObject.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeModelingProjectDetailResponseBodyResultObject struct {
	// End time of the security modeling project.
	//
	// example:
	//
	// 2026-02-07T02:22:30Z
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// Security environment status.
	//
	// example:
	//
	// ACTIVE
	EnvStatus *string `json:"EnvStatus,omitempty" xml:"EnvStatus,omitempty"`
	// Login account.
	//
	// example:
	//
	// xxx
	LoginAccount *string `json:"LoginAccount,omitempty" xml:"LoginAccount,omitempty"`
	// Machine specification.
	//
	// example:
	//
	// SECURE_ENV_LITE
	MachineSpec *string `json:"MachineSpec,omitempty" xml:"MachineSpec,omitempty"`
	// Manual operation phase.
	//
	// example:
	//
	// INIT
	ManualPhase *string `json:"ManualPhase,omitempty" xml:"ManualPhase,omitempty"`
	// Modeling project status.
	//
	// example:
	//
	// PREPARING
	ModelingStatus *string `json:"ModelingStatus,omitempty" xml:"ModelingStatus,omitempty"`
	// Associated POC tasks.
	PocTasks []*DescribeModelingProjectDetailResponseBodyResultObjectPocTasks `json:"PocTasks,omitempty" xml:"PocTasks,omitempty" type:"Repeated"`
	// Project ID
	//
	// example:
	//
	// 01
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// Project name.
	//
	// example:
	//
	// project_01
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// Remark information.
	//
	// example:
	//
	// remark
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// Start time of the security modeling project.
	//
	// example:
	//
	// 2025-08-13T01:09:00Z
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// Synchronized files.
	SyncedFiles []*DescribeModelingProjectDetailResponseBodyResultObjectSyncedFiles `json:"SyncedFiles,omitempty" xml:"SyncedFiles,omitempty" type:"Repeated"`
}

func (s DescribeModelingProjectDetailResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s DescribeModelingProjectDetailResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *DescribeModelingProjectDetailResponseBodyResultObject) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeModelingProjectDetailResponseBodyResultObject) GetEnvStatus() *string {
	return s.EnvStatus
}

func (s *DescribeModelingProjectDetailResponseBodyResultObject) GetLoginAccount() *string {
	return s.LoginAccount
}

func (s *DescribeModelingProjectDetailResponseBodyResultObject) GetMachineSpec() *string {
	return s.MachineSpec
}

func (s *DescribeModelingProjectDetailResponseBodyResultObject) GetManualPhase() *string {
	return s.ManualPhase
}

func (s *DescribeModelingProjectDetailResponseBodyResultObject) GetModelingStatus() *string {
	return s.ModelingStatus
}

func (s *DescribeModelingProjectDetailResponseBodyResultObject) GetPocTasks() []*DescribeModelingProjectDetailResponseBodyResultObjectPocTasks {
	return s.PocTasks
}

func (s *DescribeModelingProjectDetailResponseBodyResultObject) GetProjectId() *string {
	return s.ProjectId
}

func (s *DescribeModelingProjectDetailResponseBodyResultObject) GetProjectName() *string {
	return s.ProjectName
}

func (s *DescribeModelingProjectDetailResponseBodyResultObject) GetRemark() *string {
	return s.Remark
}

func (s *DescribeModelingProjectDetailResponseBodyResultObject) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeModelingProjectDetailResponseBodyResultObject) GetSyncedFiles() []*DescribeModelingProjectDetailResponseBodyResultObjectSyncedFiles {
	return s.SyncedFiles
}

func (s *DescribeModelingProjectDetailResponseBodyResultObject) SetEndTime(v int64) *DescribeModelingProjectDetailResponseBodyResultObject {
	s.EndTime = &v
	return s
}

func (s *DescribeModelingProjectDetailResponseBodyResultObject) SetEnvStatus(v string) *DescribeModelingProjectDetailResponseBodyResultObject {
	s.EnvStatus = &v
	return s
}

func (s *DescribeModelingProjectDetailResponseBodyResultObject) SetLoginAccount(v string) *DescribeModelingProjectDetailResponseBodyResultObject {
	s.LoginAccount = &v
	return s
}

func (s *DescribeModelingProjectDetailResponseBodyResultObject) SetMachineSpec(v string) *DescribeModelingProjectDetailResponseBodyResultObject {
	s.MachineSpec = &v
	return s
}

func (s *DescribeModelingProjectDetailResponseBodyResultObject) SetManualPhase(v string) *DescribeModelingProjectDetailResponseBodyResultObject {
	s.ManualPhase = &v
	return s
}

func (s *DescribeModelingProjectDetailResponseBodyResultObject) SetModelingStatus(v string) *DescribeModelingProjectDetailResponseBodyResultObject {
	s.ModelingStatus = &v
	return s
}

func (s *DescribeModelingProjectDetailResponseBodyResultObject) SetPocTasks(v []*DescribeModelingProjectDetailResponseBodyResultObjectPocTasks) *DescribeModelingProjectDetailResponseBodyResultObject {
	s.PocTasks = v
	return s
}

func (s *DescribeModelingProjectDetailResponseBodyResultObject) SetProjectId(v string) *DescribeModelingProjectDetailResponseBodyResultObject {
	s.ProjectId = &v
	return s
}

func (s *DescribeModelingProjectDetailResponseBodyResultObject) SetProjectName(v string) *DescribeModelingProjectDetailResponseBodyResultObject {
	s.ProjectName = &v
	return s
}

func (s *DescribeModelingProjectDetailResponseBodyResultObject) SetRemark(v string) *DescribeModelingProjectDetailResponseBodyResultObject {
	s.Remark = &v
	return s
}

func (s *DescribeModelingProjectDetailResponseBodyResultObject) SetStartTime(v int64) *DescribeModelingProjectDetailResponseBodyResultObject {
	s.StartTime = &v
	return s
}

func (s *DescribeModelingProjectDetailResponseBodyResultObject) SetSyncedFiles(v []*DescribeModelingProjectDetailResponseBodyResultObjectSyncedFiles) *DescribeModelingProjectDetailResponseBodyResultObject {
	s.SyncedFiles = v
	return s
}

func (s *DescribeModelingProjectDetailResponseBodyResultObject) Validate() error {
	if s.PocTasks != nil {
		for _, item := range s.PocTasks {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.SyncedFiles != nil {
		for _, item := range s.SyncedFiles {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeModelingProjectDetailResponseBodyResultObjectPocTasks struct {
	// Retrospective file name.
	//
	// example:
	//
	// xxx.csv
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// Service code
	//
	// example:
	//
	// xxx
	ServiceCode *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
	// File synchronization status.
	//
	// example:
	//
	// 1
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Retrospective task name.
	//
	// example:
	//
	// xxx
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
}

func (s DescribeModelingProjectDetailResponseBodyResultObjectPocTasks) String() string {
	return dara.Prettify(s)
}

func (s DescribeModelingProjectDetailResponseBodyResultObjectPocTasks) GoString() string {
	return s.String()
}

func (s *DescribeModelingProjectDetailResponseBodyResultObjectPocTasks) GetFileName() *string {
	return s.FileName
}

func (s *DescribeModelingProjectDetailResponseBodyResultObjectPocTasks) GetServiceCode() *string {
	return s.ServiceCode
}

func (s *DescribeModelingProjectDetailResponseBodyResultObjectPocTasks) GetStatus() *string {
	return s.Status
}

func (s *DescribeModelingProjectDetailResponseBodyResultObjectPocTasks) GetTaskName() *string {
	return s.TaskName
}

func (s *DescribeModelingProjectDetailResponseBodyResultObjectPocTasks) SetFileName(v string) *DescribeModelingProjectDetailResponseBodyResultObjectPocTasks {
	s.FileName = &v
	return s
}

func (s *DescribeModelingProjectDetailResponseBodyResultObjectPocTasks) SetServiceCode(v string) *DescribeModelingProjectDetailResponseBodyResultObjectPocTasks {
	s.ServiceCode = &v
	return s
}

func (s *DescribeModelingProjectDetailResponseBodyResultObjectPocTasks) SetStatus(v string) *DescribeModelingProjectDetailResponseBodyResultObjectPocTasks {
	s.Status = &v
	return s
}

func (s *DescribeModelingProjectDetailResponseBodyResultObjectPocTasks) SetTaskName(v string) *DescribeModelingProjectDetailResponseBodyResultObjectPocTasks {
	s.TaskName = &v
	return s
}

func (s *DescribeModelingProjectDetailResponseBodyResultObjectPocTasks) Validate() error {
	return dara.Validate(s)
}

type DescribeModelingProjectDetailResponseBodyResultObjectSyncedFiles struct {
	// Whether deployment is supported.
	//
	// example:
	//
	// true
	Deployable *bool `json:"Deployable,omitempty" xml:"Deployable,omitempty"`
	// Files generated by modeling.
	Files []*DescribeModelingProjectDetailResponseBodyResultObjectSyncedFilesFiles `json:"Files,omitempty" xml:"Files,omitempty" type:"Repeated"`
	// Application group ID.
	//
	// example:
	//
	// 56790766
	GroupId *int64 `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// File group name.
	//
	// example:
	//
	// am
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// Synchronization time.
	//
	// example:
	//
	// 1770607862000
	SyncedTime *string `json:"SyncedTime,omitempty" xml:"SyncedTime,omitempty"`
}

func (s DescribeModelingProjectDetailResponseBodyResultObjectSyncedFiles) String() string {
	return dara.Prettify(s)
}

func (s DescribeModelingProjectDetailResponseBodyResultObjectSyncedFiles) GoString() string {
	return s.String()
}

func (s *DescribeModelingProjectDetailResponseBodyResultObjectSyncedFiles) GetDeployable() *bool {
	return s.Deployable
}

func (s *DescribeModelingProjectDetailResponseBodyResultObjectSyncedFiles) GetFiles() []*DescribeModelingProjectDetailResponseBodyResultObjectSyncedFilesFiles {
	return s.Files
}

func (s *DescribeModelingProjectDetailResponseBodyResultObjectSyncedFiles) GetGroupId() *int64 {
	return s.GroupId
}

func (s *DescribeModelingProjectDetailResponseBodyResultObjectSyncedFiles) GetGroupName() *string {
	return s.GroupName
}

func (s *DescribeModelingProjectDetailResponseBodyResultObjectSyncedFiles) GetSyncedTime() *string {
	return s.SyncedTime
}

func (s *DescribeModelingProjectDetailResponseBodyResultObjectSyncedFiles) SetDeployable(v bool) *DescribeModelingProjectDetailResponseBodyResultObjectSyncedFiles {
	s.Deployable = &v
	return s
}

func (s *DescribeModelingProjectDetailResponseBodyResultObjectSyncedFiles) SetFiles(v []*DescribeModelingProjectDetailResponseBodyResultObjectSyncedFilesFiles) *DescribeModelingProjectDetailResponseBodyResultObjectSyncedFiles {
	s.Files = v
	return s
}

func (s *DescribeModelingProjectDetailResponseBodyResultObjectSyncedFiles) SetGroupId(v int64) *DescribeModelingProjectDetailResponseBodyResultObjectSyncedFiles {
	s.GroupId = &v
	return s
}

func (s *DescribeModelingProjectDetailResponseBodyResultObjectSyncedFiles) SetGroupName(v string) *DescribeModelingProjectDetailResponseBodyResultObjectSyncedFiles {
	s.GroupName = &v
	return s
}

func (s *DescribeModelingProjectDetailResponseBodyResultObjectSyncedFiles) SetSyncedTime(v string) *DescribeModelingProjectDetailResponseBodyResultObjectSyncedFiles {
	s.SyncedTime = &v
	return s
}

func (s *DescribeModelingProjectDetailResponseBodyResultObjectSyncedFiles) Validate() error {
	if s.Files != nil {
		for _, item := range s.Files {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeModelingProjectDetailResponseBodyResultObjectSyncedFilesFiles struct {
	// Whether it is downloadable.
	Downloadable *bool `json:"Downloadable,omitempty" xml:"Downloadable,omitempty"`
	// The ID of the file.
	//
	// example:
	//
	// 376920
	FileId *int64 `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// The name of the file.
	//
	// example:
	//
	// 融山-个贷-精催-演示-0625.json
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// Visibility:
	//
	// **true**: Visible
	//
	// **false**: Not visible
	//
	// example:
	//
	// true
	Visible *bool `json:"Visible,omitempty" xml:"Visible,omitempty"`
}

func (s DescribeModelingProjectDetailResponseBodyResultObjectSyncedFilesFiles) String() string {
	return dara.Prettify(s)
}

func (s DescribeModelingProjectDetailResponseBodyResultObjectSyncedFilesFiles) GoString() string {
	return s.String()
}

func (s *DescribeModelingProjectDetailResponseBodyResultObjectSyncedFilesFiles) GetDownloadable() *bool {
	return s.Downloadable
}

func (s *DescribeModelingProjectDetailResponseBodyResultObjectSyncedFilesFiles) GetFileId() *int64 {
	return s.FileId
}

func (s *DescribeModelingProjectDetailResponseBodyResultObjectSyncedFilesFiles) GetFileName() *string {
	return s.FileName
}

func (s *DescribeModelingProjectDetailResponseBodyResultObjectSyncedFilesFiles) GetVisible() *bool {
	return s.Visible
}

func (s *DescribeModelingProjectDetailResponseBodyResultObjectSyncedFilesFiles) SetDownloadable(v bool) *DescribeModelingProjectDetailResponseBodyResultObjectSyncedFilesFiles {
	s.Downloadable = &v
	return s
}

func (s *DescribeModelingProjectDetailResponseBodyResultObjectSyncedFilesFiles) SetFileId(v int64) *DescribeModelingProjectDetailResponseBodyResultObjectSyncedFilesFiles {
	s.FileId = &v
	return s
}

func (s *DescribeModelingProjectDetailResponseBodyResultObjectSyncedFilesFiles) SetFileName(v string) *DescribeModelingProjectDetailResponseBodyResultObjectSyncedFilesFiles {
	s.FileName = &v
	return s
}

func (s *DescribeModelingProjectDetailResponseBodyResultObjectSyncedFilesFiles) SetVisible(v bool) *DescribeModelingProjectDetailResponseBodyResultObjectSyncedFilesFiles {
	s.Visible = &v
	return s
}

func (s *DescribeModelingProjectDetailResponseBodyResultObjectSyncedFilesFiles) Validate() error {
	return dara.Validate(s)
}
