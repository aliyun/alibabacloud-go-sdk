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
	// The business status code. The value SUCCESS is returned if the request was successful.
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The HTTP status code. The value 200 is returned if the request was successful.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The list of sandboxes that match the filter conditions.
	Items []*ListSandboxesResponseBodyItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// The maximum number of records per page for this query.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The response message. The value success is returned if the request was successful.
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
	// Indicates whether the request was successful. Valid values:
	//
	// - true: The request was successful.
	//
	// - false: The request failed.
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// The total number of records that match the query conditions.
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
	// The number of active sessions for this sandbox.
	//
	// example:
	//
	// 1
	ActiveSessionCount *int32 `json:"activeSessionCount,omitempty" xml:"activeSessionCount,omitempty"`
	// The time when the sandbox was created, in RFC 3339 UTC format.
	//
	// example:
	//
	// 2026-08-29T00:00:00Z
	CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// The time of the last activity on the sandbox, in RFC 3339 UTC format.
	//
	// example:
	//
	// 2026-08-29T00:02:00Z
	LastActiveAt *string `json:"lastActiveAt,omitempty" xml:"lastActiveAt,omitempty"`
	// The time of the last heartbeat from the sandbox, in RFC 3339 UTC format.
	//
	// example:
	//
	// 2026-08-29T00:01:00Z
	LastHeartbeatAt *string `json:"lastHeartbeatAt,omitempty" xml:"lastHeartbeatAt,omitempty"`
	// The maximum number of concurrent sessions allowed for this sandbox, derived from the auto scaling configuration in effect at runtime. This value is empty if auto scaling is not enabled or the configuration is unavailable.
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
	// The sandbox ID. You can call the ListSandboxes operation to query sandbox IDs.
	//
	// example:
	//
	// sbx-2
	SandboxId *string `json:"sandboxId,omitempty" xml:"sandboxId,omitempty"`
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

func (s *ListSandboxesResponseBodyItems) GetLastActiveAt() *string {
	return s.LastActiveAt
}

func (s *ListSandboxesResponseBodyItems) GetLastHeartbeatAt() *string {
	return s.LastHeartbeatAt
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

func (s *ListSandboxesResponseBodyItems) SetActiveSessionCount(v int32) *ListSandboxesResponseBodyItems {
	s.ActiveSessionCount = &v
	return s
}

func (s *ListSandboxesResponseBodyItems) SetCreatedAt(v string) *ListSandboxesResponseBodyItems {
	s.CreatedAt = &v
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

func (s *ListSandboxesResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
