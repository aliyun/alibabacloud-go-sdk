// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportInsightWorkitemStatusResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetMaxResults(v int64) *ExportInsightWorkitemStatusResponseBody
  GetMaxResults() *int64 
  SetNextToken(v string) *ExportInsightWorkitemStatusResponseBody
  GetNextToken() *string 
  SetResult(v []*ExportInsightWorkitemStatusResponseBodyResult) *ExportInsightWorkitemStatusResponseBody
  GetResult() []*ExportInsightWorkitemStatusResponseBodyResult 
  SetTotalCount(v int64) *ExportInsightWorkitemStatusResponseBody
  GetTotalCount() *int64 
}

type ExportInsightWorkitemStatusResponseBody struct {
  // example:
  // 
  // 10
  MaxResults *int64 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
  // example:
  // 
  // 2
  NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
  Result []*ExportInsightWorkitemStatusResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
  // example:
  // 
  // 100
  TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ExportInsightWorkitemStatusResponseBody) String() string {
  return dara.Prettify(s)
}

func (s ExportInsightWorkitemStatusResponseBody) GoString() string {
  return s.String()
}

func (s *ExportInsightWorkitemStatusResponseBody) GetMaxResults() *int64  {
  return s.MaxResults
}

func (s *ExportInsightWorkitemStatusResponseBody) GetNextToken() *string  {
  return s.NextToken
}

func (s *ExportInsightWorkitemStatusResponseBody) GetResult() []*ExportInsightWorkitemStatusResponseBodyResult  {
  return s.Result
}

func (s *ExportInsightWorkitemStatusResponseBody) GetTotalCount() *int64  {
  return s.TotalCount
}

func (s *ExportInsightWorkitemStatusResponseBody) SetMaxResults(v int64) *ExportInsightWorkitemStatusResponseBody {
  s.MaxResults = &v
  return s
}

func (s *ExportInsightWorkitemStatusResponseBody) SetNextToken(v string) *ExportInsightWorkitemStatusResponseBody {
  s.NextToken = &v
  return s
}

func (s *ExportInsightWorkitemStatusResponseBody) SetResult(v []*ExportInsightWorkitemStatusResponseBodyResult) *ExportInsightWorkitemStatusResponseBody {
  s.Result = v
  return s
}

func (s *ExportInsightWorkitemStatusResponseBody) SetTotalCount(v int64) *ExportInsightWorkitemStatusResponseBody {
  s.TotalCount = &v
  return s
}

func (s *ExportInsightWorkitemStatusResponseBody) Validate() error {
  if s.Result != nil {
    for _, item := range s.Result {
      if item != nil {
        if err := item.Validate(); err != nil {
          return err
        }
      }
    }
  }
  return nil
}

type ExportInsightWorkitemStatusResponseBodyResult struct {
  // example:
  // 
  // 65e836b981d758be7a25xxxx
  AssignedToId *string `json:"assignedToId,omitempty" xml:"assignedToId,omitempty"`
  // example:
  // 
  // 65e836b981d758be7a25xxxx
  CreatorId *string `json:"creatorId,omitempty" xml:"creatorId,omitempty"`
  // example:
  // 
  // 10
  ExpectedWorkTime *int64 `json:"expectedWorkTime,omitempty" xml:"expectedWorkTime,omitempty"`
  // example:
  // 
  // 33166339200000
  GmtClosed *int64 `json:"gmtClosed,omitempty" xml:"gmtClosed,omitempty"`
  // example:
  // 
  // 1713430241000
  GmtCreate *int64 `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
  // example:
  // 
  // 33166339200000
  GmtDue *int64 `json:"gmtDue,omitempty" xml:"gmtDue,omitempty"`
  // example:
  // 
  // 1713430241000
  GmtModified *int64 `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
  // example:
  // 
  // 33166339200000
  GmtStart *int64 `json:"gmtStart,omitempty" xml:"gmtStart,omitempty"`
  // example:
  // 
  // 33166339200000
  GmtTodo *int64 `json:"gmtTodo,omitempty" xml:"gmtTodo,omitempty"`
  // example:
  // 
  // 701615370
  Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
  // example:
  // 
  // N
  IsArchived *string `json:"isArchived,omitempty" xml:"isArchived,omitempty"`
  // example:
  // 
  // N
  IsDeleted *string `json:"isDeleted,omitempty" xml:"isDeleted,omitempty"`
  // example:
  // 
  // Y
  IsDone *string `json:"isDone,omitempty" xml:"isDone,omitempty"`
  // example:
  // 
  // 61db9af2148974246bexxxx
  OrganizationId *string `json:"organizationId,omitempty" xml:"organizationId,omitempty"`
  // example:
  // 
  // c3640ab6233fcc10a7e3aaxxxx
  ParentId *string `json:"parentId,omitempty" xml:"parentId,omitempty"`
  // example:
  // 
  // 6135b21fb383ef39551cxxxx,63466a385dc8531eebd7xxxx
  ParticipantIds *string `json:"participantIds,omitempty" xml:"participantIds,omitempty"`
  // example:
  // 
  // 10
  Phase *int32 `json:"phase,omitempty" xml:"phase,omitempty"`
  // example:
  // 
  // 0
  Priority *int32 `json:"priority,omitempty" xml:"priority,omitempty"`
  ProductId *string `json:"productId,omitempty" xml:"productId,omitempty"`
  // example:
  // 
  // 505ac6433dfbda8df0b08bxxxx
  ProjectId *string `json:"projectId,omitempty" xml:"projectId,omitempty"`
  SerialNumber *int32 `json:"serialNumber,omitempty" xml:"serialNumber,omitempty"`
  // example:
  // 
  // projex
  Source *string `json:"source,omitempty" xml:"source,omitempty"`
  // example:
  // 
  // 505ac6433dfbda8df0b08bxxxx
  SprintId *string `json:"sprintId,omitempty" xml:"sprintId,omitempty"`
  // example:
  // 
  // 10
  Stage *int32 `json:"stage,omitempty" xml:"stage,omitempty"`
  Status *string `json:"status,omitempty" xml:"status,omitempty"`
  // example:
  // 
  // 100005
  StatusId *string `json:"statusId,omitempty" xml:"statusId,omitempty"`
  // example:
  // 
  // 10.0
  StoryPoint *float32 `json:"storyPoint,omitempty" xml:"storyPoint,omitempty"`
  // example:
  // 
  // 9uy29901re573f561d69xxxx
  SubType *string `json:"subType,omitempty" xml:"subType,omitempty"`
  Subject *string `json:"subject,omitempty" xml:"subject,omitempty"`
  // example:
  // 
  // 1
  Type *int32 `json:"type,omitempty" xml:"type,omitempty"`
  VersionId *string `json:"versionId,omitempty" xml:"versionId,omitempty"`
  // example:
  // 
  // [6a8cdda167415bea1506c7262c]
  Versions *string `json:"versions,omitempty" xml:"versions,omitempty"`
  // example:
  // 
  // 10
  WorkTime *int64 `json:"workTime,omitempty" xml:"workTime,omitempty"`
  // example:
  // 
  // 636f661a612a945bbcdb4cxxxx
  WorkitemId *string `json:"workitemId,omitempty" xml:"workitemId,omitempty"`
}

func (s ExportInsightWorkitemStatusResponseBodyResult) String() string {
  return dara.Prettify(s)
}

func (s ExportInsightWorkitemStatusResponseBodyResult) GoString() string {
  return s.String()
}

func (s *ExportInsightWorkitemStatusResponseBodyResult) GetAssignedToId() *string  {
  return s.AssignedToId
}

func (s *ExportInsightWorkitemStatusResponseBodyResult) GetCreatorId() *string  {
  return s.CreatorId
}

func (s *ExportInsightWorkitemStatusResponseBodyResult) GetExpectedWorkTime() *int64  {
  return s.ExpectedWorkTime
}

func (s *ExportInsightWorkitemStatusResponseBodyResult) GetGmtClosed() *int64  {
  return s.GmtClosed
}

func (s *ExportInsightWorkitemStatusResponseBodyResult) GetGmtCreate() *int64  {
  return s.GmtCreate
}

func (s *ExportInsightWorkitemStatusResponseBodyResult) GetGmtDue() *int64  {
  return s.GmtDue
}

func (s *ExportInsightWorkitemStatusResponseBodyResult) GetGmtModified() *int64  {
  return s.GmtModified
}

func (s *ExportInsightWorkitemStatusResponseBodyResult) GetGmtStart() *int64  {
  return s.GmtStart
}

func (s *ExportInsightWorkitemStatusResponseBodyResult) GetGmtTodo() *int64  {
  return s.GmtTodo
}

func (s *ExportInsightWorkitemStatusResponseBodyResult) GetId() *int64  {
  return s.Id
}

func (s *ExportInsightWorkitemStatusResponseBodyResult) GetIsArchived() *string  {
  return s.IsArchived
}

func (s *ExportInsightWorkitemStatusResponseBodyResult) GetIsDeleted() *string  {
  return s.IsDeleted
}

func (s *ExportInsightWorkitemStatusResponseBodyResult) GetIsDone() *string  {
  return s.IsDone
}

func (s *ExportInsightWorkitemStatusResponseBodyResult) GetOrganizationId() *string  {
  return s.OrganizationId
}

func (s *ExportInsightWorkitemStatusResponseBodyResult) GetParentId() *string  {
  return s.ParentId
}

func (s *ExportInsightWorkitemStatusResponseBodyResult) GetParticipantIds() *string  {
  return s.ParticipantIds
}

func (s *ExportInsightWorkitemStatusResponseBodyResult) GetPhase() *int32  {
  return s.Phase
}

func (s *ExportInsightWorkitemStatusResponseBodyResult) GetPriority() *int32  {
  return s.Priority
}

func (s *ExportInsightWorkitemStatusResponseBodyResult) GetProductId() *string  {
  return s.ProductId
}

func (s *ExportInsightWorkitemStatusResponseBodyResult) GetProjectId() *string  {
  return s.ProjectId
}

func (s *ExportInsightWorkitemStatusResponseBodyResult) GetSerialNumber() *int32  {
  return s.SerialNumber
}

func (s *ExportInsightWorkitemStatusResponseBodyResult) GetSource() *string  {
  return s.Source
}

func (s *ExportInsightWorkitemStatusResponseBodyResult) GetSprintId() *string  {
  return s.SprintId
}

func (s *ExportInsightWorkitemStatusResponseBodyResult) GetStage() *int32  {
  return s.Stage
}

func (s *ExportInsightWorkitemStatusResponseBodyResult) GetStatus() *string  {
  return s.Status
}

func (s *ExportInsightWorkitemStatusResponseBodyResult) GetStatusId() *string  {
  return s.StatusId
}

func (s *ExportInsightWorkitemStatusResponseBodyResult) GetStoryPoint() *float32  {
  return s.StoryPoint
}

func (s *ExportInsightWorkitemStatusResponseBodyResult) GetSubType() *string  {
  return s.SubType
}

func (s *ExportInsightWorkitemStatusResponseBodyResult) GetSubject() *string  {
  return s.Subject
}

func (s *ExportInsightWorkitemStatusResponseBodyResult) GetType() *int32  {
  return s.Type
}

func (s *ExportInsightWorkitemStatusResponseBodyResult) GetVersionId() *string  {
  return s.VersionId
}

func (s *ExportInsightWorkitemStatusResponseBodyResult) GetVersions() *string  {
  return s.Versions
}

func (s *ExportInsightWorkitemStatusResponseBodyResult) GetWorkTime() *int64  {
  return s.WorkTime
}

func (s *ExportInsightWorkitemStatusResponseBodyResult) GetWorkitemId() *string  {
  return s.WorkitemId
}

func (s *ExportInsightWorkitemStatusResponseBodyResult) SetAssignedToId(v string) *ExportInsightWorkitemStatusResponseBodyResult {
  s.AssignedToId = &v
  return s
}

func (s *ExportInsightWorkitemStatusResponseBodyResult) SetCreatorId(v string) *ExportInsightWorkitemStatusResponseBodyResult {
  s.CreatorId = &v
  return s
}

func (s *ExportInsightWorkitemStatusResponseBodyResult) SetExpectedWorkTime(v int64) *ExportInsightWorkitemStatusResponseBodyResult {
  s.ExpectedWorkTime = &v
  return s
}

func (s *ExportInsightWorkitemStatusResponseBodyResult) SetGmtClosed(v int64) *ExportInsightWorkitemStatusResponseBodyResult {
  s.GmtClosed = &v
  return s
}

func (s *ExportInsightWorkitemStatusResponseBodyResult) SetGmtCreate(v int64) *ExportInsightWorkitemStatusResponseBodyResult {
  s.GmtCreate = &v
  return s
}

func (s *ExportInsightWorkitemStatusResponseBodyResult) SetGmtDue(v int64) *ExportInsightWorkitemStatusResponseBodyResult {
  s.GmtDue = &v
  return s
}

func (s *ExportInsightWorkitemStatusResponseBodyResult) SetGmtModified(v int64) *ExportInsightWorkitemStatusResponseBodyResult {
  s.GmtModified = &v
  return s
}

func (s *ExportInsightWorkitemStatusResponseBodyResult) SetGmtStart(v int64) *ExportInsightWorkitemStatusResponseBodyResult {
  s.GmtStart = &v
  return s
}

func (s *ExportInsightWorkitemStatusResponseBodyResult) SetGmtTodo(v int64) *ExportInsightWorkitemStatusResponseBodyResult {
  s.GmtTodo = &v
  return s
}

func (s *ExportInsightWorkitemStatusResponseBodyResult) SetId(v int64) *ExportInsightWorkitemStatusResponseBodyResult {
  s.Id = &v
  return s
}

func (s *ExportInsightWorkitemStatusResponseBodyResult) SetIsArchived(v string) *ExportInsightWorkitemStatusResponseBodyResult {
  s.IsArchived = &v
  return s
}

func (s *ExportInsightWorkitemStatusResponseBodyResult) SetIsDeleted(v string) *ExportInsightWorkitemStatusResponseBodyResult {
  s.IsDeleted = &v
  return s
}

func (s *ExportInsightWorkitemStatusResponseBodyResult) SetIsDone(v string) *ExportInsightWorkitemStatusResponseBodyResult {
  s.IsDone = &v
  return s
}

func (s *ExportInsightWorkitemStatusResponseBodyResult) SetOrganizationId(v string) *ExportInsightWorkitemStatusResponseBodyResult {
  s.OrganizationId = &v
  return s
}

func (s *ExportInsightWorkitemStatusResponseBodyResult) SetParentId(v string) *ExportInsightWorkitemStatusResponseBodyResult {
  s.ParentId = &v
  return s
}

func (s *ExportInsightWorkitemStatusResponseBodyResult) SetParticipantIds(v string) *ExportInsightWorkitemStatusResponseBodyResult {
  s.ParticipantIds = &v
  return s
}

func (s *ExportInsightWorkitemStatusResponseBodyResult) SetPhase(v int32) *ExportInsightWorkitemStatusResponseBodyResult {
  s.Phase = &v
  return s
}

func (s *ExportInsightWorkitemStatusResponseBodyResult) SetPriority(v int32) *ExportInsightWorkitemStatusResponseBodyResult {
  s.Priority = &v
  return s
}

func (s *ExportInsightWorkitemStatusResponseBodyResult) SetProductId(v string) *ExportInsightWorkitemStatusResponseBodyResult {
  s.ProductId = &v
  return s
}

func (s *ExportInsightWorkitemStatusResponseBodyResult) SetProjectId(v string) *ExportInsightWorkitemStatusResponseBodyResult {
  s.ProjectId = &v
  return s
}

func (s *ExportInsightWorkitemStatusResponseBodyResult) SetSerialNumber(v int32) *ExportInsightWorkitemStatusResponseBodyResult {
  s.SerialNumber = &v
  return s
}

func (s *ExportInsightWorkitemStatusResponseBodyResult) SetSource(v string) *ExportInsightWorkitemStatusResponseBodyResult {
  s.Source = &v
  return s
}

func (s *ExportInsightWorkitemStatusResponseBodyResult) SetSprintId(v string) *ExportInsightWorkitemStatusResponseBodyResult {
  s.SprintId = &v
  return s
}

func (s *ExportInsightWorkitemStatusResponseBodyResult) SetStage(v int32) *ExportInsightWorkitemStatusResponseBodyResult {
  s.Stage = &v
  return s
}

func (s *ExportInsightWorkitemStatusResponseBodyResult) SetStatus(v string) *ExportInsightWorkitemStatusResponseBodyResult {
  s.Status = &v
  return s
}

func (s *ExportInsightWorkitemStatusResponseBodyResult) SetStatusId(v string) *ExportInsightWorkitemStatusResponseBodyResult {
  s.StatusId = &v
  return s
}

func (s *ExportInsightWorkitemStatusResponseBodyResult) SetStoryPoint(v float32) *ExportInsightWorkitemStatusResponseBodyResult {
  s.StoryPoint = &v
  return s
}

func (s *ExportInsightWorkitemStatusResponseBodyResult) SetSubType(v string) *ExportInsightWorkitemStatusResponseBodyResult {
  s.SubType = &v
  return s
}

func (s *ExportInsightWorkitemStatusResponseBodyResult) SetSubject(v string) *ExportInsightWorkitemStatusResponseBodyResult {
  s.Subject = &v
  return s
}

func (s *ExportInsightWorkitemStatusResponseBodyResult) SetType(v int32) *ExportInsightWorkitemStatusResponseBodyResult {
  s.Type = &v
  return s
}

func (s *ExportInsightWorkitemStatusResponseBodyResult) SetVersionId(v string) *ExportInsightWorkitemStatusResponseBodyResult {
  s.VersionId = &v
  return s
}

func (s *ExportInsightWorkitemStatusResponseBodyResult) SetVersions(v string) *ExportInsightWorkitemStatusResponseBodyResult {
  s.Versions = &v
  return s
}

func (s *ExportInsightWorkitemStatusResponseBodyResult) SetWorkTime(v int64) *ExportInsightWorkitemStatusResponseBodyResult {
  s.WorkTime = &v
  return s
}

func (s *ExportInsightWorkitemStatusResponseBodyResult) SetWorkitemId(v string) *ExportInsightWorkitemStatusResponseBodyResult {
  s.WorkitemId = &v
  return s
}

func (s *ExportInsightWorkitemStatusResponseBodyResult) Validate() error {
  return dara.Validate(s)
}

