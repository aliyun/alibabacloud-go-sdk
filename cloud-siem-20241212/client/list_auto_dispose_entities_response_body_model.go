// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAutoDisposeEntitiesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAutoDecisionEntities(v []*ListAutoDisposeEntitiesResponseBodyAutoDecisionEntities) *ListAutoDisposeEntitiesResponseBody
	GetAutoDecisionEntities() []*ListAutoDisposeEntitiesResponseBodyAutoDecisionEntities
	SetCurrentPage(v int32) *ListAutoDisposeEntitiesResponseBody
	GetCurrentPage() *int32
	SetMaxResults(v int32) *ListAutoDisposeEntitiesResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListAutoDisposeEntitiesResponseBody
	GetNextToken() *string
	SetPageSize(v int32) *ListAutoDisposeEntitiesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListAutoDisposeEntitiesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListAutoDisposeEntitiesResponseBody
	GetTotalCount() *int32
}

type ListAutoDisposeEntitiesResponseBody struct {
	// A list of AI-powered automated analysis entities.
	AutoDecisionEntities []*ListAutoDisposeEntitiesResponseBodyAutoDecisionEntities `json:"AutoDecisionEntities,omitempty" xml:"AutoDecisionEntities,omitempty" type:"Repeated"`
	// The current page number. The value must be greater than or equal to 1.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The maximum number of entries returned in this request.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token used to retrieve the next page of results. An empty value indicates that all results have been returned.
	//
	// example:
	//
	// AAAAAUqcj6VO4E3ECWIrFczs****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 6276D891-*****-55B2-87B9-74D413F7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total count of entries.
	//
	// example:
	//
	// 57
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListAutoDisposeEntitiesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAutoDisposeEntitiesResponseBody) GoString() string {
	return s.String()
}

func (s *ListAutoDisposeEntitiesResponseBody) GetAutoDecisionEntities() []*ListAutoDisposeEntitiesResponseBodyAutoDecisionEntities {
	return s.AutoDecisionEntities
}

func (s *ListAutoDisposeEntitiesResponseBody) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListAutoDisposeEntitiesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListAutoDisposeEntitiesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAutoDisposeEntitiesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAutoDisposeEntitiesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAutoDisposeEntitiesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListAutoDisposeEntitiesResponseBody) SetAutoDecisionEntities(v []*ListAutoDisposeEntitiesResponseBodyAutoDecisionEntities) *ListAutoDisposeEntitiesResponseBody {
	s.AutoDecisionEntities = v
	return s
}

func (s *ListAutoDisposeEntitiesResponseBody) SetCurrentPage(v int32) *ListAutoDisposeEntitiesResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *ListAutoDisposeEntitiesResponseBody) SetMaxResults(v int32) *ListAutoDisposeEntitiesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListAutoDisposeEntitiesResponseBody) SetNextToken(v string) *ListAutoDisposeEntitiesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListAutoDisposeEntitiesResponseBody) SetPageSize(v int32) *ListAutoDisposeEntitiesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListAutoDisposeEntitiesResponseBody) SetRequestId(v string) *ListAutoDisposeEntitiesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAutoDisposeEntitiesResponseBody) SetTotalCount(v int32) *ListAutoDisposeEntitiesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListAutoDisposeEntitiesResponseBody) Validate() error {
	if s.AutoDecisionEntities != nil {
		for _, item := range s.AutoDecisionEntities {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAutoDisposeEntitiesResponseBodyAutoDecisionEntities struct {
	// The alert rule ID.
	//
	// example:
	//
	// 20403189
	AlertId *string `json:"AlertId,omitempty" xml:"AlertId,omitempty"`
	// The disposal method. Valid values:
	//
	// - `ignore`: Ignore.
	//
	// - `dispose`: Playbook-based disposal.
	//
	// example:
	//
	// dispose
	DisposalMethod *string `json:"DisposalMethod,omitempty" xml:"DisposalMethod,omitempty"`
	// The disposal ID.
	//
	// example:
	//
	// 1ec121479b341a61886dbd2c4ccd*****
	DisposeRecordId *string `json:"DisposeRecordId,omitempty" xml:"DisposeRecordId,omitempty"`
	// The entity name.
	//
	// example:
	//
	// /apps/ext/ka****
	EntityName *string `json:"EntityName,omitempty" xml:"EntityName,omitempty"`
	// The entity type. Valid values:
	//
	// - `ip`: An IP address.
	//
	// - `domain`: A domain.
	//
	// - `process`: A process.
	//
	// - `file`: A file.
	//
	// - `host`: A host.
	//
	// example:
	//
	// process
	EntityType *string `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	// The entity UUID.
	//
	// example:
	//
	// 022ed6c601514a370cc9e3acd37a****
	EntityUuid *string `json:"EntityUuid,omitempty" xml:"EntityUuid,omitempty"`
	// The playbook UUID.
	//
	// example:
	//
	// 6fd2b143-e420-4c1b-a118-e764*****
	PlaybookUuid *string `json:"PlaybookUuid,omitempty" xml:"PlaybookUuid,omitempty"`
	// The unique identifier (UUID) of the entry.
	//
	// example:
	//
	// 0d23f133-22d7-4388-b738-ab******
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s ListAutoDisposeEntitiesResponseBodyAutoDecisionEntities) String() string {
	return dara.Prettify(s)
}

func (s ListAutoDisposeEntitiesResponseBodyAutoDecisionEntities) GoString() string {
	return s.String()
}

func (s *ListAutoDisposeEntitiesResponseBodyAutoDecisionEntities) GetAlertId() *string {
	return s.AlertId
}

func (s *ListAutoDisposeEntitiesResponseBodyAutoDecisionEntities) GetDisposalMethod() *string {
	return s.DisposalMethod
}

func (s *ListAutoDisposeEntitiesResponseBodyAutoDecisionEntities) GetDisposeRecordId() *string {
	return s.DisposeRecordId
}

func (s *ListAutoDisposeEntitiesResponseBodyAutoDecisionEntities) GetEntityName() *string {
	return s.EntityName
}

func (s *ListAutoDisposeEntitiesResponseBodyAutoDecisionEntities) GetEntityType() *string {
	return s.EntityType
}

func (s *ListAutoDisposeEntitiesResponseBodyAutoDecisionEntities) GetEntityUuid() *string {
	return s.EntityUuid
}

func (s *ListAutoDisposeEntitiesResponseBodyAutoDecisionEntities) GetPlaybookUuid() *string {
	return s.PlaybookUuid
}

func (s *ListAutoDisposeEntitiesResponseBodyAutoDecisionEntities) GetUuid() *string {
	return s.Uuid
}

func (s *ListAutoDisposeEntitiesResponseBodyAutoDecisionEntities) SetAlertId(v string) *ListAutoDisposeEntitiesResponseBodyAutoDecisionEntities {
	s.AlertId = &v
	return s
}

func (s *ListAutoDisposeEntitiesResponseBodyAutoDecisionEntities) SetDisposalMethod(v string) *ListAutoDisposeEntitiesResponseBodyAutoDecisionEntities {
	s.DisposalMethod = &v
	return s
}

func (s *ListAutoDisposeEntitiesResponseBodyAutoDecisionEntities) SetDisposeRecordId(v string) *ListAutoDisposeEntitiesResponseBodyAutoDecisionEntities {
	s.DisposeRecordId = &v
	return s
}

func (s *ListAutoDisposeEntitiesResponseBodyAutoDecisionEntities) SetEntityName(v string) *ListAutoDisposeEntitiesResponseBodyAutoDecisionEntities {
	s.EntityName = &v
	return s
}

func (s *ListAutoDisposeEntitiesResponseBodyAutoDecisionEntities) SetEntityType(v string) *ListAutoDisposeEntitiesResponseBodyAutoDecisionEntities {
	s.EntityType = &v
	return s
}

func (s *ListAutoDisposeEntitiesResponseBodyAutoDecisionEntities) SetEntityUuid(v string) *ListAutoDisposeEntitiesResponseBodyAutoDecisionEntities {
	s.EntityUuid = &v
	return s
}

func (s *ListAutoDisposeEntitiesResponseBodyAutoDecisionEntities) SetPlaybookUuid(v string) *ListAutoDisposeEntitiesResponseBodyAutoDecisionEntities {
	s.PlaybookUuid = &v
	return s
}

func (s *ListAutoDisposeEntitiesResponseBodyAutoDecisionEntities) SetUuid(v string) *ListAutoDisposeEntitiesResponseBodyAutoDecisionEntities {
	s.Uuid = &v
	return s
}

func (s *ListAutoDisposeEntitiesResponseBodyAutoDecisionEntities) Validate() error {
	return dara.Validate(s)
}
