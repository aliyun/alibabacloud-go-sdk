// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSnapshotsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCompleteTimeEnd(v int64) *ListSnapshotsRequest
	GetCompleteTimeEnd() *int64
	SetCompleteTimeStart(v int64) *ListSnapshotsRequest
	GetCompleteTimeStart() *int64
	SetMaxResults(v int32) *ListSnapshotsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListSnapshotsRequest
	GetNextToken() *string
	SetPlanId(v string) *ListSnapshotsRequest
	GetPlanId() *string
	SetProtectedResourceId(v string) *ListSnapshotsRequest
	GetProtectedResourceId() *string
	SetSkip(v int32) *ListSnapshotsRequest
	GetSkip() *int32
	SetSourceType(v string) *ListSnapshotsRequest
	GetSourceType() *string
	SetVaultId(v string) *ListSnapshotsRequest
	GetVaultId() *string
}

type ListSnapshotsRequest struct {
	// The end completion time. A UNIX timestamp, in seconds. Only backup points with a completion time earlier than this value are returned.
	//
	// example:
	//
	// 1776059382
	CompleteTimeEnd *int64 `json:"CompleteTimeEnd,omitempty" xml:"CompleteTimeEnd,omitempty"`
	// The start completion time. A UNIX timestamp, in seconds. Only backup points with a completion time greater than or equal to this value are returned.
	//
	// example:
	//
	// 1776059082
	CompleteTimeStart *int64 `json:"CompleteTimeStart,omitempty" xml:"CompleteTimeStart,omitempty"`
	// The maximum number of results to return in a paged query. Default value: 10. The value must be greater than 1.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token. You do not need to specify this parameter for the first request. For subsequent pages, pass in the **NextToken*	- value returned from the previous request. An empty **NextToken*	- indicates that no more pages are available. This parameter cannot be used together with **Skip**.
	//
	// example:
	//
	// eyJJbmRleCI6NjUwMX0=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The backup plan ID. This is an optional filter condition. If specified, only backup points under the specified backup plan are returned.
	//
	// example:
	//
	// po-0000****123
	PlanId *string `json:"PlanId,omitempty" xml:"PlanId,omitempty"`
	// The ID of the protected resource. The protected resource must exist, and its data source type must match the **SourceType*	- parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// pr-123***7890
	ProtectedResourceId *string `json:"ProtectedResourceId,omitempty" xml:"ProtectedResourceId,omitempty"`
	// The number of entries to skip. The value must be a non-negative integer and a multiple of **MaxResults**. This parameter cannot be used together with **NextToken**.
	//
	// example:
	//
	// 0
	Skip *int32 `json:"Skip,omitempty" xml:"Skip,omitempty"`
	// The data source type. The value must match the data source type of the protected resource (ProtectedResourceId). The supported data source types are listed in the valid values.
	//
	// This parameter is required.
	//
	// example:
	//
	// ECS_FILE
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// The vault ID. This is an optional filter condition. If specified, only backup points in the specified vault are returned.
	//
	// example:
	//
	// v-00030j3c******sn
	VaultId *string `json:"VaultId,omitempty" xml:"VaultId,omitempty"`
}

func (s ListSnapshotsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSnapshotsRequest) GoString() string {
	return s.String()
}

func (s *ListSnapshotsRequest) GetCompleteTimeEnd() *int64 {
	return s.CompleteTimeEnd
}

func (s *ListSnapshotsRequest) GetCompleteTimeStart() *int64 {
	return s.CompleteTimeStart
}

func (s *ListSnapshotsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListSnapshotsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListSnapshotsRequest) GetPlanId() *string {
	return s.PlanId
}

func (s *ListSnapshotsRequest) GetProtectedResourceId() *string {
	return s.ProtectedResourceId
}

func (s *ListSnapshotsRequest) GetSkip() *int32 {
	return s.Skip
}

func (s *ListSnapshotsRequest) GetSourceType() *string {
	return s.SourceType
}

func (s *ListSnapshotsRequest) GetVaultId() *string {
	return s.VaultId
}

func (s *ListSnapshotsRequest) SetCompleteTimeEnd(v int64) *ListSnapshotsRequest {
	s.CompleteTimeEnd = &v
	return s
}

func (s *ListSnapshotsRequest) SetCompleteTimeStart(v int64) *ListSnapshotsRequest {
	s.CompleteTimeStart = &v
	return s
}

func (s *ListSnapshotsRequest) SetMaxResults(v int32) *ListSnapshotsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListSnapshotsRequest) SetNextToken(v string) *ListSnapshotsRequest {
	s.NextToken = &v
	return s
}

func (s *ListSnapshotsRequest) SetPlanId(v string) *ListSnapshotsRequest {
	s.PlanId = &v
	return s
}

func (s *ListSnapshotsRequest) SetProtectedResourceId(v string) *ListSnapshotsRequest {
	s.ProtectedResourceId = &v
	return s
}

func (s *ListSnapshotsRequest) SetSkip(v int32) *ListSnapshotsRequest {
	s.Skip = &v
	return s
}

func (s *ListSnapshotsRequest) SetSourceType(v string) *ListSnapshotsRequest {
	s.SourceType = &v
	return s
}

func (s *ListSnapshotsRequest) SetVaultId(v string) *ListSnapshotsRequest {
	s.VaultId = &v
	return s
}

func (s *ListSnapshotsRequest) Validate() error {
	return dara.Validate(s)
}
