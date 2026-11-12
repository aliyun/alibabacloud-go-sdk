// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSandboxesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListSandboxesResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ListSandboxesResponseBody
	GetHttpStatusCode() *int32
	SetItems(v []*ListSandboxesResponseBodyItems) *ListSandboxesResponseBody
	GetItems() []*ListSandboxesResponseBodyItems
	SetMaxResults(v int32) *ListSandboxesResponseBody
	GetMaxResults() *int32
	SetMessage(v string) *ListSandboxesResponseBody
	GetMessage() *string
	SetNextToken(v string) *ListSandboxesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListSandboxesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListSandboxesResponseBody
	GetSuccess() *bool
	SetTotalCount(v int64) *ListSandboxesResponseBody
	GetTotalCount() *int64
}

type ListSandboxesResponseBody struct {
	// The business status code. The value is SUCCESS if the request is successful.
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The HTTP status code. The value is 200 if the request is successful.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The list of sandboxes that meet the filter conditions.
	Items []*ListSandboxesResponseBodyItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// The maximum number of records per page for this query.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The response message. The value is success if the request is successful.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The pagination token for the next page. An empty value indicates that no more results are available.
	//
	// example:
	//
	// next-token-1
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The request ID, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 1a2b3c4d-xxxx-xxxx-xxxx-xxxxxxxxxxxx
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request is successful. A value of true indicates success, and a value of false indicates failure.
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// The total number of records that meet the query conditions.
	//
	// example:
	//
	// 27
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListSandboxesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSandboxesResponseBody) GoString() string {
	return s.String()
}

func (s *ListSandboxesResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListSandboxesResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListSandboxesResponseBody) GetItems() []*ListSandboxesResponseBodyItems {
	return s.Items
}

func (s *ListSandboxesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListSandboxesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListSandboxesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListSandboxesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSandboxesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListSandboxesResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListSandboxesResponseBody) SetCode(v string) *ListSandboxesResponseBody {
	s.Code = &v
	return s
}

func (s *ListSandboxesResponseBody) SetHttpStatusCode(v int32) *ListSandboxesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListSandboxesResponseBody) SetItems(v []*ListSandboxesResponseBodyItems) *ListSandboxesResponseBody {
	s.Items = v
	return s
}

func (s *ListSandboxesResponseBody) SetMaxResults(v int32) *ListSandboxesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListSandboxesResponseBody) SetMessage(v string) *ListSandboxesResponseBody {
	s.Message = &v
	return s
}

func (s *ListSandboxesResponseBody) SetNextToken(v string) *ListSandboxesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListSandboxesResponseBody) SetRequestId(v string) *ListSandboxesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSandboxesResponseBody) SetSuccess(v bool) *ListSandboxesResponseBody {
	s.Success = &v
	return s
}

func (s *ListSandboxesResponseBody) SetTotalCount(v int64) *ListSandboxesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListSandboxesResponseBody) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListSandboxesResponseBodyItems struct {
	// The number of active sessions in the sandbox.
	//
	// example:
	//
	// 1
	ActiveSessionCount *int32 `json:"activeSessionCount,omitempty" xml:"activeSessionCount,omitempty"`
	// The time when the sandbox was created. The time is in RFC 3339 UTC format.
	//
	// example:
	//
	// 2026-08-29T00:00:00Z
	CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// The tag of the actual observed image of the currently returned instance. This parameter is empty if the tag cannot be parsed.
	//
	// example:
	//
	// 202611031754
	HarnessVersion *string `json:"harnessVersion,omitempty" xml:"harnessVersion,omitempty"`
	// The time of the last activity of the sandbox. The time is in RFC 3339 UTC format.
	//
	// example:
	//
	// 2026-08-29T00:02:00Z
	LastActiveAt *string `json:"lastActiveAt,omitempty" xml:"lastActiveAt,omitempty"`
	// The time of the last heartbeat of the sandbox. The time is in RFC 3339 UTC format.
	//
	// example:
	//
	// 2026-08-29T00:01:00Z
	LastHeartbeatAt *string `json:"lastHeartbeatAt,omitempty" xml:"lastHeartbeatAt,omitempty"`
	// Used to continue calling GetSandboxRebuildOperation to query the overall result after the page is refreshed. This parameter is empty if there is no historical rebuild. Do not use the ID of a subsequently created task instead.
	//
	// example:
	//
	// 01a19f98-906d-7e82-96aa-8b7ea1d1a29d
	LastRebuildOperationId *string `json:"lastRebuildOperationId,omitempty" xml:"lastRebuildOperationId,omitempty"`
	// The maximum number of concurrent sessions allowed for the sandbox, which is obtained from the effective auto-scaling configuration at runtime. This parameter is empty if auto-scaling is disabled or the configuration is unavailable.
	//
	// example:
	//
	// 7
	MaxConcurrentSessions *int32 `json:"maxConcurrentSessions,omitempty" xml:"maxConcurrentSessions,omitempty"`
	// The current running phase of the sandbox.
	//
	// example:
	//
	// RUNNING
	Phase *string `json:"phase,omitempty" xml:"phase,omitempty"`
	// The sandbox ID. You can query it by calling ListSandboxes.
	//
	// example:
	//
	// sbx-2
	SandboxId *string `json:"sandboxId,omitempty" xml:"sandboxId,omitempty"`
	// The server compares the actual image tag of the instance with the target image tag of the pool. The value is true if the tags are different during the RUNNING, HIBERNATED, or FAILED phase. The value is false if the tags are the same or during other known phases. This parameter is empty if the tag or phase cannot be confirmed, and it cannot be treated as false. The version sequence is not determined by the tag size.
	//
	// example:
	//
	// true
	UpdateAvailable *bool `json:"updateAvailable,omitempty" xml:"updateAvailable,omitempty"`
	// A value of true indicates that the most recent rebuild corresponding to the current instance generation failed as a whole. A value of false indicates that there is no historical rebuild or the rebuild succeeded as a whole. This parameter is empty if the rebuild is in progress, canceled, replaced, the generation does not match, or the result cannot be confirmed.
	//
	// example:
	//
	// false
	UpdateFailed *bool `json:"updateFailed,omitempty" xml:"updateFailed,omitempty"`
}

func (s ListSandboxesResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s ListSandboxesResponseBodyItems) GoString() string {
	return s.String()
}

func (s *ListSandboxesResponseBodyItems) GetActiveSessionCount() *int32 {
	return s.ActiveSessionCount
}

func (s *ListSandboxesResponseBodyItems) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *ListSandboxesResponseBodyItems) GetHarnessVersion() *string {
	return s.HarnessVersion
}

func (s *ListSandboxesResponseBodyItems) GetLastActiveAt() *string {
	return s.LastActiveAt
}

func (s *ListSandboxesResponseBodyItems) GetLastHeartbeatAt() *string {
	return s.LastHeartbeatAt
}

func (s *ListSandboxesResponseBodyItems) GetLastRebuildOperationId() *string {
	return s.LastRebuildOperationId
}

func (s *ListSandboxesResponseBodyItems) GetMaxConcurrentSessions() *int32 {
	return s.MaxConcurrentSessions
}

func (s *ListSandboxesResponseBodyItems) GetPhase() *string {
	return s.Phase
}

func (s *ListSandboxesResponseBodyItems) GetSandboxId() *string {
	return s.SandboxId
}

func (s *ListSandboxesResponseBodyItems) GetUpdateAvailable() *bool {
	return s.UpdateAvailable
}

func (s *ListSandboxesResponseBodyItems) GetUpdateFailed() *bool {
	return s.UpdateFailed
}

func (s *ListSandboxesResponseBodyItems) SetActiveSessionCount(v int32) *ListSandboxesResponseBodyItems {
	s.ActiveSessionCount = &v
	return s
}

func (s *ListSandboxesResponseBodyItems) SetCreatedAt(v string) *ListSandboxesResponseBodyItems {
	s.CreatedAt = &v
	return s
}

func (s *ListSandboxesResponseBodyItems) SetHarnessVersion(v string) *ListSandboxesResponseBodyItems {
	s.HarnessVersion = &v
	return s
}

func (s *ListSandboxesResponseBodyItems) SetLastActiveAt(v string) *ListSandboxesResponseBodyItems {
	s.LastActiveAt = &v
	return s
}

func (s *ListSandboxesResponseBodyItems) SetLastHeartbeatAt(v string) *ListSandboxesResponseBodyItems {
	s.LastHeartbeatAt = &v
	return s
}

func (s *ListSandboxesResponseBodyItems) SetLastRebuildOperationId(v string) *ListSandboxesResponseBodyItems {
	s.LastRebuildOperationId = &v
	return s
}

func (s *ListSandboxesResponseBodyItems) SetMaxConcurrentSessions(v int32) *ListSandboxesResponseBodyItems {
	s.MaxConcurrentSessions = &v
	return s
}

func (s *ListSandboxesResponseBodyItems) SetPhase(v string) *ListSandboxesResponseBodyItems {
	s.Phase = &v
	return s
}

func (s *ListSandboxesResponseBodyItems) SetSandboxId(v string) *ListSandboxesResponseBodyItems {
	s.SandboxId = &v
	return s
}

func (s *ListSandboxesResponseBodyItems) SetUpdateAvailable(v bool) *ListSandboxesResponseBodyItems {
	s.UpdateAvailable = &v
	return s
}

func (s *ListSandboxesResponseBodyItems) SetUpdateFailed(v bool) *ListSandboxesResponseBodyItems {
	s.UpdateFailed = &v
	return s
}

func (s *ListSandboxesResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
