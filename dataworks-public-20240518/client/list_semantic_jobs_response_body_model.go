// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSemanticJobsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ListSemanticJobsResponseBodyData) *ListSemanticJobsResponseBody
	GetData() *ListSemanticJobsResponseBodyData
	SetRequestId(v string) *ListSemanticJobsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListSemanticJobsResponseBody
	GetSuccess() *bool
}

type ListSemanticJobsResponseBody struct {
	// The paginated result of job definitions. Use the Name field of a list item to run, delete, query run records, or download results. Use the ProjectId field to query run details, logs, or stop a run.
	Data *ListSemanticJobsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID. Used for locating logs and troubleshooting issues.
	//
	// example:
	//
	// 676271D6-53B4-57BE-89FA-72F7AE1418DF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListSemanticJobsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSemanticJobsResponseBody) GoString() string {
	return s.String()
}

func (s *ListSemanticJobsResponseBody) GetData() *ListSemanticJobsResponseBodyData {
	return s.Data
}

func (s *ListSemanticJobsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSemanticJobsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListSemanticJobsResponseBody) SetData(v *ListSemanticJobsResponseBodyData) *ListSemanticJobsResponseBody {
	s.Data = v
	return s
}

func (s *ListSemanticJobsResponseBody) SetRequestId(v string) *ListSemanticJobsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSemanticJobsResponseBody) SetSuccess(v bool) *ListSemanticJobsResponseBody {
	s.Success = &v
	return s
}

func (s *ListSemanticJobsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListSemanticJobsResponseBodyData struct {
	// The page number returned, starting from 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of job definitions per page returned.
	//
	// example:
	//
	// 50
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The list of job definitions.
	SemanticJobs []*ListSemanticJobsResponseBodyDataSemanticJobs `json:"SemanticJobs,omitempty" xml:"SemanticJobs,omitempty" type:"Repeated"`
	// The total number of job definitions that meet the conditions in the current tenant.
	//
	// example:
	//
	// 1
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListSemanticJobsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListSemanticJobsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListSemanticJobsResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListSemanticJobsResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListSemanticJobsResponseBodyData) GetSemanticJobs() []*ListSemanticJobsResponseBodyDataSemanticJobs {
	return s.SemanticJobs
}

func (s *ListSemanticJobsResponseBodyData) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListSemanticJobsResponseBodyData) SetPageNumber(v int32) *ListSemanticJobsResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListSemanticJobsResponseBodyData) SetPageSize(v int32) *ListSemanticJobsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListSemanticJobsResponseBodyData) SetSemanticJobs(v []*ListSemanticJobsResponseBodyDataSemanticJobs) *ListSemanticJobsResponseBodyData {
	s.SemanticJobs = v
	return s
}

func (s *ListSemanticJobsResponseBodyData) SetTotalCount(v int64) *ListSemanticJobsResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListSemanticJobsResponseBodyData) Validate() error {
	if s.SemanticJobs != nil {
		for _, item := range s.SemanticJobs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListSemanticJobsResponseBodyDataSemanticJobs struct {
	// The user identifier of the semantic job creator.
	//
	// example:
	//
	// user-demo
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// The creation time, expressed as a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1700000000000
	GmtCreate *int64 `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The last modification time, expressed as a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1700000000000
	GmtModified *int64 `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The internal unique ID of the job definition.
	//
	// example:
	//
	// 1
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The job name. Used for RunSemanticJob, DeleteSemanticJob, ListSemanticJobRuns, and DownloadSemanticResults.
	//
	// example:
	//
	// semantic-job-demo
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The workspace ID to which the job belongs. Used for GetSemanticJobDetail, GetSemanticJobLog, and KillSemanticJob.
	//
	// example:
	//
	// 100
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The list of uploaded file IDs associated with the job.
	ReferenceFileIds []*string `json:"ReferenceFileIds,omitempty" xml:"ReferenceFileIds,omitempty" type:"Repeated"`
	// The list of external reference file URIs associated with the job.
	ReferenceFileUris []*string `json:"ReferenceFileUris,omitempty" xml:"ReferenceFileUris,omitempty" type:"Repeated"`
	// The resource group identifier used when running this job.
	//
	// example:
	//
	// rg-demo
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The input datasource config saved with the node. This configuration determines the analysis scope at runtime.
	Source map[string]interface{} `json:"Source,omitempty" xml:"Source,omitempty"`
	// The Source.type data source type saved with the job.
	//
	// example:
	//
	// maxcompute
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The user identifier of the semantic job creator.
	//
	// example:
	//
	// user-demo
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ListSemanticJobsResponseBodyDataSemanticJobs) String() string {
	return dara.Prettify(s)
}

func (s ListSemanticJobsResponseBodyDataSemanticJobs) GoString() string {
	return s.String()
}

func (s *ListSemanticJobsResponseBodyDataSemanticJobs) GetCreator() *string {
	return s.Creator
}

func (s *ListSemanticJobsResponseBodyDataSemanticJobs) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *ListSemanticJobsResponseBodyDataSemanticJobs) GetGmtModified() *int64 {
	return s.GmtModified
}

func (s *ListSemanticJobsResponseBodyDataSemanticJobs) GetId() *int64 {
	return s.Id
}

func (s *ListSemanticJobsResponseBodyDataSemanticJobs) GetName() *string {
	return s.Name
}

func (s *ListSemanticJobsResponseBodyDataSemanticJobs) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListSemanticJobsResponseBodyDataSemanticJobs) GetReferenceFileIds() []*string {
	return s.ReferenceFileIds
}

func (s *ListSemanticJobsResponseBodyDataSemanticJobs) GetReferenceFileUris() []*string {
	return s.ReferenceFileUris
}

func (s *ListSemanticJobsResponseBodyDataSemanticJobs) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListSemanticJobsResponseBodyDataSemanticJobs) GetSource() map[string]interface{} {
	return s.Source
}

func (s *ListSemanticJobsResponseBodyDataSemanticJobs) GetType() *string {
	return s.Type
}

func (s *ListSemanticJobsResponseBodyDataSemanticJobs) GetUserId() *string {
	return s.UserId
}

func (s *ListSemanticJobsResponseBodyDataSemanticJobs) SetCreator(v string) *ListSemanticJobsResponseBodyDataSemanticJobs {
	s.Creator = &v
	return s
}

func (s *ListSemanticJobsResponseBodyDataSemanticJobs) SetGmtCreate(v int64) *ListSemanticJobsResponseBodyDataSemanticJobs {
	s.GmtCreate = &v
	return s
}

func (s *ListSemanticJobsResponseBodyDataSemanticJobs) SetGmtModified(v int64) *ListSemanticJobsResponseBodyDataSemanticJobs {
	s.GmtModified = &v
	return s
}

func (s *ListSemanticJobsResponseBodyDataSemanticJobs) SetId(v int64) *ListSemanticJobsResponseBodyDataSemanticJobs {
	s.Id = &v
	return s
}

func (s *ListSemanticJobsResponseBodyDataSemanticJobs) SetName(v string) *ListSemanticJobsResponseBodyDataSemanticJobs {
	s.Name = &v
	return s
}

func (s *ListSemanticJobsResponseBodyDataSemanticJobs) SetProjectId(v int64) *ListSemanticJobsResponseBodyDataSemanticJobs {
	s.ProjectId = &v
	return s
}

func (s *ListSemanticJobsResponseBodyDataSemanticJobs) SetReferenceFileIds(v []*string) *ListSemanticJobsResponseBodyDataSemanticJobs {
	s.ReferenceFileIds = v
	return s
}

func (s *ListSemanticJobsResponseBodyDataSemanticJobs) SetReferenceFileUris(v []*string) *ListSemanticJobsResponseBodyDataSemanticJobs {
	s.ReferenceFileUris = v
	return s
}

func (s *ListSemanticJobsResponseBodyDataSemanticJobs) SetResourceGroupId(v string) *ListSemanticJobsResponseBodyDataSemanticJobs {
	s.ResourceGroupId = &v
	return s
}

func (s *ListSemanticJobsResponseBodyDataSemanticJobs) SetSource(v map[string]interface{}) *ListSemanticJobsResponseBodyDataSemanticJobs {
	s.Source = v
	return s
}

func (s *ListSemanticJobsResponseBodyDataSemanticJobs) SetType(v string) *ListSemanticJobsResponseBodyDataSemanticJobs {
	s.Type = &v
	return s
}

func (s *ListSemanticJobsResponseBodyDataSemanticJobs) SetUserId(v string) *ListSemanticJobsResponseBodyDataSemanticJobs {
	s.UserId = &v
	return s
}

func (s *ListSemanticJobsResponseBodyDataSemanticJobs) Validate() error {
	return dara.Validate(s)
}
