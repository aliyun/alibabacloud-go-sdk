// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportInsightWorkitemStatusJoinWorkitemDefectExtraResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetMaxResults(v int64) *ExportInsightWorkitemStatusJoinWorkitemDefectExtraResponseBody
  GetMaxResults() *int64 
  SetNextToken(v string) *ExportInsightWorkitemStatusJoinWorkitemDefectExtraResponseBody
  GetNextToken() *string 
  SetResult(v []*ExportInsightWorkitemStatusJoinWorkitemDefectExtraResponseBodyResult) *ExportInsightWorkitemStatusJoinWorkitemDefectExtraResponseBody
  GetResult() []*ExportInsightWorkitemStatusJoinWorkitemDefectExtraResponseBodyResult 
  SetTotalCount(v int64) *ExportInsightWorkitemStatusJoinWorkitemDefectExtraResponseBody
  GetTotalCount() *int64 
}

type ExportInsightWorkitemStatusJoinWorkitemDefectExtraResponseBody struct {
  // example:
  // 
  // 10
  MaxResults *int64 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
  // example:
  // 
  // 2
  NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
  Result []*ExportInsightWorkitemStatusJoinWorkitemDefectExtraResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
  // example:
  // 
  // 100
  TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ExportInsightWorkitemStatusJoinWorkitemDefectExtraResponseBody) String() string {
  return dara.Prettify(s)
}

func (s ExportInsightWorkitemStatusJoinWorkitemDefectExtraResponseBody) GoString() string {
  return s.String()
}

func (s *ExportInsightWorkitemStatusJoinWorkitemDefectExtraResponseBody) GetMaxResults() *int64  {
  return s.MaxResults
}

func (s *ExportInsightWorkitemStatusJoinWorkitemDefectExtraResponseBody) GetNextToken() *string  {
  return s.NextToken
}

func (s *ExportInsightWorkitemStatusJoinWorkitemDefectExtraResponseBody) GetResult() []*ExportInsightWorkitemStatusJoinWorkitemDefectExtraResponseBodyResult  {
  return s.Result
}

func (s *ExportInsightWorkitemStatusJoinWorkitemDefectExtraResponseBody) GetTotalCount() *int64  {
  return s.TotalCount
}

func (s *ExportInsightWorkitemStatusJoinWorkitemDefectExtraResponseBody) SetMaxResults(v int64) *ExportInsightWorkitemStatusJoinWorkitemDefectExtraResponseBody {
  s.MaxResults = &v
  return s
}

func (s *ExportInsightWorkitemStatusJoinWorkitemDefectExtraResponseBody) SetNextToken(v string) *ExportInsightWorkitemStatusJoinWorkitemDefectExtraResponseBody {
  s.NextToken = &v
  return s
}

func (s *ExportInsightWorkitemStatusJoinWorkitemDefectExtraResponseBody) SetResult(v []*ExportInsightWorkitemStatusJoinWorkitemDefectExtraResponseBodyResult) *ExportInsightWorkitemStatusJoinWorkitemDefectExtraResponseBody {
  s.Result = v
  return s
}

func (s *ExportInsightWorkitemStatusJoinWorkitemDefectExtraResponseBody) SetTotalCount(v int64) *ExportInsightWorkitemStatusJoinWorkitemDefectExtraResponseBody {
  s.TotalCount = &v
  return s
}

func (s *ExportInsightWorkitemStatusJoinWorkitemDefectExtraResponseBody) Validate() error {
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

type ExportInsightWorkitemStatusJoinWorkitemDefectExtraResponseBodyResult struct {
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
  // 123
  ExtraId *int64 `json:"extraId,omitempty" xml:"extraId,omitempty"`
  // example:
  // 
  // N
  ExtraIsDeleted *string `json:"extraIsDeleted,omitempty" xml:"extraIsDeleted,omitempty"`
  // example:
  // 
  // 10
  FoundPhase *int32 `json:"foundPhase,omitempty" xml:"foundPhase,omitempty"`
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
  GmtFixed *string `json:"gmtFixed,omitempty" xml:"gmtFixed,omitempty"`
  // example:
  // 
  // 1714755985000
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
  IsStupid *string `json:"isStupid,omitempty" xml:"isStupid,omitempty"`
  // example:
  // 
  // 61db9af2148974246bexxxx
  OrganizationId *string `json:"organizationId,omitempty" xml:"organizationId,omitempty"`
  // example:
  // 
  // a80a203a9078a7a1b1f2c6xxxx
  ParentId *string `json:"parentId,omitempty" xml:"parentId,omitempty"`
  // example:
  // 
  // 6135b21fb383ef39551cf02e,63466a385dc8531eebd764e9
  ParticipantIds *string `json:"participantIds,omitempty" xml:"participantIds,omitempty"`
  // example:
  // 
  // 10
  Phase *int32 `json:"phase,omitempty" xml:"phase,omitempty"`
  // example:
  // 
  // 2
  Priority *int32 `json:"priority,omitempty" xml:"priority,omitempty"`
  ProductId *string `json:"productId,omitempty" xml:"productId,omitempty"`
  // example:
  // 
  // 6732a29d846bf998dc09e7xxxx
  ProjectId *string `json:"projectId,omitempty" xml:"projectId,omitempty"`
  ReopenNum *int32 `json:"reopenNum,omitempty" xml:"reopenNum,omitempty"`
  SerialNumber *int32 `json:"serialNumber,omitempty" xml:"serialNumber,omitempty"`
  SeriousLevel *int32 `json:"seriousLevel,omitempty" xml:"seriousLevel,omitempty"`
  Solution *string `json:"solution,omitempty" xml:"solution,omitempty"`
  // example:
  // 
  // projex
  Source *string `json:"source,omitempty" xml:"source,omitempty"`
  // example:
  // 
  // 731c83a40bbf3c2f080e07xxxx
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
  // example:
  // 
  // 65e836b981d758be7a25xxxx
  VerifierId *string `json:"verifierId,omitempty" xml:"verifierId,omitempty"`
  VersionId *string `json:"versionId,omitempty" xml:"versionId,omitempty"`
  // example:
  // 
  // {6a8cdda167415bea1506c7262c}
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

func (s ExportInsightWorkitemStatusJoinWorkitemDefectExtraResponseBodyResult) String() string {
  return dara.Prettify(s)
}

func (s ExportInsightWorkitemStatusJoinWorkitemDefectExtraResponseBodyResult) GoString() string {
  return s.String()
}

func (s *ExportInsightWorkitemStatusJoinWorkitemDefectExtraResponseBodyResult) GetAssignedToId() *string  {
  return s.AssignedToId
}

func (s *ExportInsightWorkitemStatusJoinWorkitemDefectExtraResponseBodyResult) GetCreatorId() *string  {
  return s.CreatorId
}

func (s *ExportInsightWorkitemStatusJoinWorkitemDefectExtraResponseBodyResult) GetExpectedWorkTime() *int64  {
  return s.ExpectedWorkTime
}

func (s *ExportInsightWorkitemStatusJoinWorkitemDefectExtraResponseBodyResult) GetExtraId() *int64  {
  return s.ExtraId
}

func (s *ExportInsightWorkitemStatusJoinWorkitemDefectExtraResponseBodyResult) GetExtraIsDeleted() *string  {
  return s.ExtraIsDeleted
}

func (s *ExportInsightWorkitemStatusJoinWorkitemDefectExtraResponseBodyResult) GetFoundPhase() *int32  {
  return s.FoundPhase
}

func (s *ExportInsightWorkitemStatusJoinWorkitemDefectExtraResponseBodyResult) GetGmtClosed() *int64  {
  return s.GmtClosed
}

func (s *ExportInsightWorkitemStatusJoinWorkitemDefectExtraResponseBodyResult) GetGmtCreate() *int64  {
  return s.GmtCreate
}

func (s *ExportInsightWorkitemStatusJoinWorkitemDefectExtraResponseBodyResult) GetGmtDue() *int64  {
  return s.GmtDue
}

func (s *ExportInsightWorkitemStatusJoinWorkitemDefectExtraResponseBodyResult) GetGmtFixed() *string  {
  return s.GmtFixed
}

func (s *ExportInsightWorkitemStatusJoinWorkitemDefectExtraResponseBodyResult) GetGmtModified() *int64  {
  return s.GmtModified
}

func (s *ExportInsightWorkitemStatusJoinWorkitemDefectExtraResponseBodyResult) GetGmtStart() *int64  {
  return s.GmtStart
}

func (s *ExportInsightWorkitemStatusJoinWorkitemDefectExtraResponseBodyResult) GetGmtTodo() *int64  {
  return s.GmtTodo
}

func (s *ExportInsightWorkitemStatusJoinWorkitemDefectExtraResponseBodyResult) GetId() *int64  {
  return s.Id
}

func (s *ExportInsightWorkitemStatusJoinWorkitemDefectExtraResponseBodyResult) GetIsArchived() *string  {
  return s.IsArchived
}

func (s *ExportInsightWorkitemStatusJoinWorkitemDefectExtraResponseBodyResult) GetIsDeleted() *string  {
  return s.IsDeleted
}

func (s *ExportInsightWorkitemStatusJoinWorkitemDefectExtraResponseBodyResult) GetIsDone() *string  {
  return s.IsDone
}

func (s *ExportInsightWorkitemStatusJoinWorkitemDefectExtraResponseBodyResult) GetIsStupid() *string  {
  return s.IsStupid
}

func (s *ExportInsightWorkitemStatusJoinWorkitemDefectExtraResponseBodyResult) GetOrganizationId() *string  {
  return s.OrganizationId
}

func (s *ExportInsightWorkitemStatusJoinWorkitemDefectExtraResponseBodyResult) GetParentId() *string  {
  return s.ParentId
}

func (s *ExportInsightWorkitemStatusJoinWorkitemDefectExtraResponseBodyResult) GetParticipantIds() *string  {
  return s.ParticipantIds
}

func (s *ExportInsightWorkitemStatusJoinWorkitemDefectExtraResponseBodyResult) GetPhase() *int32  {
  return s.Phase
}

func (s *ExportInsightWorkitemStatusJoinWorkitemDefectExtraResponseBodyResult) GetPriority() *int32  {
  return s.Priority
}

func (s *ExportInsightWorkitemStatusJoinWorkitemDefectExtraResponseBodyResult) GetProductId() *string  {
  return s.ProductId
}

func (s *ExportInsightWorkitemStatusJoinWorkitemDefectExtraResponseBodyResult) GetProjectId() *string  {
  return s.ProjectId
}

func (s *ExportInsightWorkitemStatusJoinWorkitemDefectExtraResponseBodyResult) GetReopenNum() *int32  {
  return s.ReopenNum
}

func (s *ExportInsightWorkitemStatusJoinWorkitemDefectExtraResponseBodyResult) GetSerialNumber() *int32  {
  return s.SerialNumber
}

func (s *ExportInsightWorkitemStatusJoinWorkitemDefectExtraResponseBodyResult) GetSeriousLevel() *int32  {
  return s.SeriousLevel
}

func (s *ExportInsightWorkitemStatusJoinWorkitemDefectExtraResponseBodyResult) GetSolution() *string  {
  return s.Solution
}

func (s *ExportInsightWorkitemStatusJoinWorkitemDefectExtraResponseBodyResult) GetSource() *string  {
  return s.Source
}

func (s *ExportInsightWorkitemStatusJoinWorkitemDefectExtraResponseBodyResult) GetSprintId() *string  {
  return s.SprintId
}

func (s *ExportInsightWorkitemStatusJoinWorkitemDefectExtraResponseBodyResult) GetStage() *int32  {
  return s.Stage
}

func (s *ExportInsightWorkitemStatusJoinWorkitemDefectExtraResponseBodyResult) GetStatus() *string  {
  return s.Status
}

func (s *ExportInsightWorkitemStatusJoinWorkitemDefectExtraResponseBodyResult) GetStatusId() *string  {
  return s.StatusId
}

func (s *ExportInsightWorkitemStatusJoinWorkitemDefectExtraResponseBodyResult) GetStoryPoint() *float32  {
  return s.StoryPoint
}

func (s *ExportInsightWorkitemStatusJoinWorkitemDefectExtraResponseBodyResult) GetSubType() *string  {
  return s.SubType
}

func (s *ExportInsightWorkitemStatusJoinWorkitemDefectExtraResponseBodyResult) GetSubject() *string  {
  return s.Subject
}

func (s *ExportInsightWorkitemStatusJoinWorkitemDefectExtraResponseBodyResult) GetType() *int32  {
  return s.Type
}

func (s *ExportInsightWorkitemStatusJoinWorkitemDefectExtraResponseBodyResult) GetVerifierId() *string  {
  return s.VerifierId
}

func (s *ExportInsightWorkitemStatusJoinWorkitemDefectExtraResponseBodyResult) GetVersionId() *string  {
  return s.VersionId
}

func (s *ExportInsightWorkitemStatusJoinWorkitemDefectExtraResponseBodyResult) GetVersions() *string  {
  return s.Versions
}

func (s *ExportInsightWorkitemStatusJoinWorkitemDefectExtraResponseBodyResult) GetWorkTime() *int64  {
  return s.WorkTime
}

func (s *ExportInsightWorkitemStatusJoinWorkitemDefectExtraResponseBodyResult) GetWorkitemId() *string  {
  return s.WorkitemId
}

func (s *ExportInsightWorkitemStatusJoinWorkitemDefectExtraResponseBodyResult) SetAssignedToId(v string) *ExportInsightWorkitemStatusJoinWorkitemDefectExtraResponseBodyResult {
  s.AssignedToId = &v
  return s
}

func (s *ExportInsightWorkitemStatusJoinWorkitemDefectExtraResponseBodyResult) SetCreatorId(v string) *ExportInsightWorkitemStatusJoinWorkitemDefectExtraResponseBodyResult {
  s.CreatorId = &v
  return s
}

func (s *ExportInsightWorkitemStatusJoinWorkitemDefectExtraResponseBodyResult) SetExpectedWorkTime(v int64) *ExportInsightWorkitemStatusJoinWorkitemDefectExtraResponseBodyResult {
  s.ExpectedWorkTime = &v
  return s
}

func (s *ExportInsightWorkitemStatusJoinWorkitemDefectExtraResponseBodyResult) SetExtraId(v int64) *ExportInsightWorkitemStatusJoinWorkitemDefectExtraResponseBodyResult {
  s.ExtraId = &v
  return s
}

func (s *ExportInsightWorkitemStatusJoinWorkitemDefectExtraResponseBodyResult) SetExtraIsDeleted(v string) *ExportInsightWorkitemStatusJoinWorkitemDefectExtraResponseBodyResult {
  s.ExtraIsDeleted = &v
  return s
}

func (s *ExportInsightWorkitemStatusJoinWorkitemDefectExtraResponseBodyResult) SetFoundPhase(v int32) *ExportInsightWorkitemStatusJoinWorkitemDefectExtraResponseBodyResult {
  s.FoundPhase = &v
  return s
}

func (s *ExportInsightWorkitemStatusJoinWorkitemDefectExtraResponseBodyResult) SetGmtClosed(v int64) *ExportInsightWorkitemStatusJoinWorkitemDefectExtraResponseBodyResult {
  s.GmtClosed = &v
  return s
}

func (s *ExportInsightWorkitemStatusJoinWorkitemDefectExtraResponseBodyResult) SetGmtCreate(v int64) *ExportInsightWorkitemStatusJoinWorkitemDefectExtraResponseBodyResult {
  s.GmtCreate = &v
  return s
}

func (s *ExportInsightWorkitemStatusJoinWorkitemDefectExtraResponseBodyResult) SetGmtDue(v int64) *ExportInsightWorkitemStatusJoinWorkitemDefectExtraResponseBodyResult {
  s.GmtDue = &v
  return s
}

func (s *ExportInsightWorkitemStatusJoinWorkitemDefectExtraResponseBodyResult) SetGmtFixed(v string) *ExportInsightWorkitemStatusJoinWorkitemDefectExtraResponseBodyResult {
  s.GmtFixed = &v
  return s
}

func (s *ExportInsightWorkitemStatusJoinWorkitemDefectExtraResponseBodyResult) SetGmtModified(v int64) *ExportInsightWorkitemStatusJoinWorkitemDefectExtraResponseBodyResult {
  s.GmtModified = &v
  return s
}

func (s *ExportInsightWorkitemStatusJoinWorkitemDefectExtraResponseBodyResult) SetGmtStart(v int64) *ExportInsightWorkitemStatusJoinWorkitemDefectExtraResponseBodyResult {
  s.GmtStart = &v
  return s
}

func (s *ExportInsightWorkitemStatusJoinWorkitemDefectExtraResponseBodyResult) SetGmtTodo(v int64) *ExportInsightWorkitemStatusJoinWorkitemDefectExtraResponseBodyResult {
  s.GmtTodo = &v
  return s
}

func (s *ExportInsightWorkitemStatusJoinWorkitemDefectExtraResponseBodyResult) SetId(v int64) *ExportInsightWorkitemStatusJoinWorkitemDefectExtraResponseBodyResult {
  s.Id = &v
  return s
}

func (s *ExportInsightWorkitemStatusJoinWorkitemDefectExtraResponseBodyResult) SetIsArchived(v string) *ExportInsightWorkitemStatusJoinWorkitemDefectExtraResponseBodyResult {
  s.IsArchived = &v
  return s
}

func (s *ExportInsightWorkitemStatusJoinWorkitemDefectExtraResponseBodyResult) SetIsDeleted(v string) *ExportInsightWorkitemStatusJoinWorkitemDefectExtraResponseBodyResult {
  s.IsDeleted = &v
  return s
}

func (s *ExportInsightWorkitemStatusJoinWorkitemDefectExtraResponseBodyResult) SetIsDone(v string) *ExportInsightWorkitemStatusJoinWorkitemDefectExtraResponseBodyResult {
  s.IsDone = &v
  return s
}

func (s *ExportInsightWorkitemStatusJoinWorkitemDefectExtraResponseBodyResult) SetIsStupid(v string) *ExportInsightWorkitemStatusJoinWorkitemDefectExtraResponseBodyResult {
  s.IsStupid = &v
  return s
}

func (s *ExportInsightWorkitemStatusJoinWorkitemDefectExtraResponseBodyResult) SetOrganizationId(v string) *ExportInsightWorkitemStatusJoinWorkitemDefectExtraResponseBodyResult {
  s.OrganizationId = &v
  return s
}

func (s *ExportInsightWorkitemStatusJoinWorkitemDefectExtraResponseBodyResult) SetParentId(v string) *ExportInsightWorkitemStatusJoinWorkitemDefectExtraResponseBodyResult {
  s.ParentId = &v
  return s
}

func (s *ExportInsightWorkitemStatusJoinWorkitemDefectExtraResponseBodyResult) SetParticipantIds(v string) *ExportInsightWorkitemStatusJoinWorkitemDefectExtraResponseBodyResult {
  s.ParticipantIds = &v
  return s
}

func (s *ExportInsightWorkitemStatusJoinWorkitemDefectExtraResponseBodyResult) SetPhase(v int32) *ExportInsightWorkitemStatusJoinWorkitemDefectExtraResponseBodyResult {
  s.Phase = &v
  return s
}

func (s *ExportInsightWorkitemStatusJoinWorkitemDefectExtraResponseBodyResult) SetPriority(v int32) *ExportInsightWorkitemStatusJoinWorkitemDefectExtraResponseBodyResult {
  s.Priority = &v
  return s
}

func (s *ExportInsightWorkitemStatusJoinWorkitemDefectExtraResponseBodyResult) SetProductId(v string) *ExportInsightWorkitemStatusJoinWorkitemDefectExtraResponseBodyResult {
  s.ProductId = &v
  return s
}

func (s *ExportInsightWorkitemStatusJoinWorkitemDefectExtraResponseBodyResult) SetProjectId(v string) *ExportInsightWorkitemStatusJoinWorkitemDefectExtraResponseBodyResult {
  s.ProjectId = &v
  return s
}

func (s *ExportInsightWorkitemStatusJoinWorkitemDefectExtraResponseBodyResult) SetReopenNum(v int32) *ExportInsightWorkitemStatusJoinWorkitemDefectExtraResponseBodyResult {
  s.ReopenNum = &v
  return s
}

func (s *ExportInsightWorkitemStatusJoinWorkitemDefectExtraResponseBodyResult) SetSerialNumber(v int32) *ExportInsightWorkitemStatusJoinWorkitemDefectExtraResponseBodyResult {
  s.SerialNumber = &v
  return s
}

func (s *ExportInsightWorkitemStatusJoinWorkitemDefectExtraResponseBodyResult) SetSeriousLevel(v int32) *ExportInsightWorkitemStatusJoinWorkitemDefectExtraResponseBodyResult {
  s.SeriousLevel = &v
  return s
}

func (s *ExportInsightWorkitemStatusJoinWorkitemDefectExtraResponseBodyResult) SetSolution(v string) *ExportInsightWorkitemStatusJoinWorkitemDefectExtraResponseBodyResult {
  s.Solution = &v
  return s
}

func (s *ExportInsightWorkitemStatusJoinWorkitemDefectExtraResponseBodyResult) SetSource(v string) *ExportInsightWorkitemStatusJoinWorkitemDefectExtraResponseBodyResult {
  s.Source = &v
  return s
}

func (s *ExportInsightWorkitemStatusJoinWorkitemDefectExtraResponseBodyResult) SetSprintId(v string) *ExportInsightWorkitemStatusJoinWorkitemDefectExtraResponseBodyResult {
  s.SprintId = &v
  return s
}

func (s *ExportInsightWorkitemStatusJoinWorkitemDefectExtraResponseBodyResult) SetStage(v int32) *ExportInsightWorkitemStatusJoinWorkitemDefectExtraResponseBodyResult {
  s.Stage = &v
  return s
}

func (s *ExportInsightWorkitemStatusJoinWorkitemDefectExtraResponseBodyResult) SetStatus(v string) *ExportInsightWorkitemStatusJoinWorkitemDefectExtraResponseBodyResult {
  s.Status = &v
  return s
}

func (s *ExportInsightWorkitemStatusJoinWorkitemDefectExtraResponseBodyResult) SetStatusId(v string) *ExportInsightWorkitemStatusJoinWorkitemDefectExtraResponseBodyResult {
  s.StatusId = &v
  return s
}

func (s *ExportInsightWorkitemStatusJoinWorkitemDefectExtraResponseBodyResult) SetStoryPoint(v float32) *ExportInsightWorkitemStatusJoinWorkitemDefectExtraResponseBodyResult {
  s.StoryPoint = &v
  return s
}

func (s *ExportInsightWorkitemStatusJoinWorkitemDefectExtraResponseBodyResult) SetSubType(v string) *ExportInsightWorkitemStatusJoinWorkitemDefectExtraResponseBodyResult {
  s.SubType = &v
  return s
}

func (s *ExportInsightWorkitemStatusJoinWorkitemDefectExtraResponseBodyResult) SetSubject(v string) *ExportInsightWorkitemStatusJoinWorkitemDefectExtraResponseBodyResult {
  s.Subject = &v
  return s
}

func (s *ExportInsightWorkitemStatusJoinWorkitemDefectExtraResponseBodyResult) SetType(v int32) *ExportInsightWorkitemStatusJoinWorkitemDefectExtraResponseBodyResult {
  s.Type = &v
  return s
}

func (s *ExportInsightWorkitemStatusJoinWorkitemDefectExtraResponseBodyResult) SetVerifierId(v string) *ExportInsightWorkitemStatusJoinWorkitemDefectExtraResponseBodyResult {
  s.VerifierId = &v
  return s
}

func (s *ExportInsightWorkitemStatusJoinWorkitemDefectExtraResponseBodyResult) SetVersionId(v string) *ExportInsightWorkitemStatusJoinWorkitemDefectExtraResponseBodyResult {
  s.VersionId = &v
  return s
}

func (s *ExportInsightWorkitemStatusJoinWorkitemDefectExtraResponseBodyResult) SetVersions(v string) *ExportInsightWorkitemStatusJoinWorkitemDefectExtraResponseBodyResult {
  s.Versions = &v
  return s
}

func (s *ExportInsightWorkitemStatusJoinWorkitemDefectExtraResponseBodyResult) SetWorkTime(v int64) *ExportInsightWorkitemStatusJoinWorkitemDefectExtraResponseBodyResult {
  s.WorkTime = &v
  return s
}

func (s *ExportInsightWorkitemStatusJoinWorkitemDefectExtraResponseBodyResult) SetWorkitemId(v string) *ExportInsightWorkitemStatusJoinWorkitemDefectExtraResponseBodyResult {
  s.WorkitemId = &v
  return s
}

func (s *ExportInsightWorkitemStatusJoinWorkitemDefectExtraResponseBodyResult) Validate() error {
  return dara.Validate(s)
}

