// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSemanticJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *CreateSemanticJobResponseBodyData) *CreateSemanticJobResponseBody
	GetData() *CreateSemanticJobResponseBodyData
	SetRequestId(v string) *CreateSemanticJobResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateSemanticJobResponseBody
	GetSuccess() *bool
}

type CreateSemanticJobResponseBody struct {
	// The saved semantic task definition. Use Data.Name to call RunSemanticJob, DeleteSemanticJob, ListSemanticJobRuns, and DownloadSemanticResults.
	Data *CreateSemanticJobResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID. Used for locating logs and troubleshooting issues.
	//
	// example:
	//
	// 676271D6-53B4-57BE-89FA-72F7AE1418DF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateSemanticJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateSemanticJobResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSemanticJobResponseBody) GetData() *CreateSemanticJobResponseBodyData {
	return s.Data
}

func (s *CreateSemanticJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateSemanticJobResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateSemanticJobResponseBody) SetData(v *CreateSemanticJobResponseBodyData) *CreateSemanticJobResponseBody {
	s.Data = v
	return s
}

func (s *CreateSemanticJobResponseBody) SetRequestId(v string) *CreateSemanticJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSemanticJobResponseBody) SetSuccess(v bool) *CreateSemanticJobResponseBody {
	s.Success = &v
	return s
}

func (s *CreateSemanticJobResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateSemanticJobResponseBodyData struct {
	// The task creator identifier, equivalent to UserId, used to display creation ownership.
	//
	// example:
	//
	// user-demo
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// The creation time of the task definition, as a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1700000000000
	GmtCreate *int64 `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The last modification time of the task definition, as a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1700000000000
	GmtModified *int64 `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The internal unique ID of the task definition, which identifies the task created by this call.
	//
	// example:
	//
	// 1
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The saved task name. Use this value for subsequent run, delete, list runs, and download results operations.
	//
	// example:
	//
	// semantic-job-demo
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The DataWorks workspace ID to which the task belongs. Use this value as the ProjectId for GetSemanticJobDetail, GetSemanticJobLog, and KillSemanticJob.
	//
	// example:
	//
	// 100
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The list of associated uploaded file IDs. For singleTableFile, the single file in this list is read at runtime.
	ReferenceFileIds []*string `json:"ReferenceFileIds,omitempty" xml:"ReferenceFileIds,omitempty" type:"Repeated"`
	// The list of associated external reference file URIs. For singleTableFile, the single file in this list is read at runtime.
	ReferenceFileUris []*string `json:"ReferenceFileUris,omitempty" xml:"ReferenceFileUris,omitempty" type:"Repeated"`
	// The ID of the resource group that will be used when running this task.
	//
	// example:
	//
	// rg-demo
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The saved input datasource config, corresponding to the Source in the creation request. The data scope to be analyzed is determined based on this configuration at runtime.
	Source map[string]interface{} `json:"Source,omitempty" xml:"Source,omitempty"`
	// The saved Source.type data source type, used to quickly identify the task input type.
	//
	// example:
	//
	// maxcompute
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The identifier of the user who created the task.
	//
	// example:
	//
	// user-demo
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s CreateSemanticJobResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateSemanticJobResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateSemanticJobResponseBodyData) GetCreator() *string {
	return s.Creator
}

func (s *CreateSemanticJobResponseBodyData) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *CreateSemanticJobResponseBodyData) GetGmtModified() *int64 {
	return s.GmtModified
}

func (s *CreateSemanticJobResponseBodyData) GetId() *int64 {
	return s.Id
}

func (s *CreateSemanticJobResponseBodyData) GetName() *string {
	return s.Name
}

func (s *CreateSemanticJobResponseBodyData) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *CreateSemanticJobResponseBodyData) GetReferenceFileIds() []*string {
	return s.ReferenceFileIds
}

func (s *CreateSemanticJobResponseBodyData) GetReferenceFileUris() []*string {
	return s.ReferenceFileUris
}

func (s *CreateSemanticJobResponseBodyData) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateSemanticJobResponseBodyData) GetSource() map[string]interface{} {
	return s.Source
}

func (s *CreateSemanticJobResponseBodyData) GetType() *string {
	return s.Type
}

func (s *CreateSemanticJobResponseBodyData) GetUserId() *string {
	return s.UserId
}

func (s *CreateSemanticJobResponseBodyData) SetCreator(v string) *CreateSemanticJobResponseBodyData {
	s.Creator = &v
	return s
}

func (s *CreateSemanticJobResponseBodyData) SetGmtCreate(v int64) *CreateSemanticJobResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *CreateSemanticJobResponseBodyData) SetGmtModified(v int64) *CreateSemanticJobResponseBodyData {
	s.GmtModified = &v
	return s
}

func (s *CreateSemanticJobResponseBodyData) SetId(v int64) *CreateSemanticJobResponseBodyData {
	s.Id = &v
	return s
}

func (s *CreateSemanticJobResponseBodyData) SetName(v string) *CreateSemanticJobResponseBodyData {
	s.Name = &v
	return s
}

func (s *CreateSemanticJobResponseBodyData) SetProjectId(v int64) *CreateSemanticJobResponseBodyData {
	s.ProjectId = &v
	return s
}

func (s *CreateSemanticJobResponseBodyData) SetReferenceFileIds(v []*string) *CreateSemanticJobResponseBodyData {
	s.ReferenceFileIds = v
	return s
}

func (s *CreateSemanticJobResponseBodyData) SetReferenceFileUris(v []*string) *CreateSemanticJobResponseBodyData {
	s.ReferenceFileUris = v
	return s
}

func (s *CreateSemanticJobResponseBodyData) SetResourceGroupId(v string) *CreateSemanticJobResponseBodyData {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateSemanticJobResponseBodyData) SetSource(v map[string]interface{}) *CreateSemanticJobResponseBodyData {
	s.Source = v
	return s
}

func (s *CreateSemanticJobResponseBodyData) SetType(v string) *CreateSemanticJobResponseBodyData {
	s.Type = &v
	return s
}

func (s *CreateSemanticJobResponseBodyData) SetUserId(v string) *CreateSemanticJobResponseBodyData {
	s.UserId = &v
	return s
}

func (s *CreateSemanticJobResponseBodyData) Validate() error {
	return dara.Validate(s)
}
