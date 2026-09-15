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
	Code           *string                           `json:"code,omitempty" xml:"code,omitempty"`
	HttpStatusCode *int32                            `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	Items          []*ListSandboxesResponseBodyItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	MaxResults     *int32                            `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	Message        *string                           `json:"message,omitempty" xml:"message,omitempty"`
	NextToken      *string                           `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	RequestId      *string                           `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success        *bool                             `json:"success,omitempty" xml:"success,omitempty"`
	TotalCount     *int64                            `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
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
	ActiveSessionCount    *int32  `json:"activeSessionCount,omitempty" xml:"activeSessionCount,omitempty"`
	CreatedAt             *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	LastActiveAt          *string `json:"lastActiveAt,omitempty" xml:"lastActiveAt,omitempty"`
	LastHeartbeatAt       *string `json:"lastHeartbeatAt,omitempty" xml:"lastHeartbeatAt,omitempty"`
	MaxConcurrentSessions *int32  `json:"maxConcurrentSessions,omitempty" xml:"maxConcurrentSessions,omitempty"`
	Phase                 *string `json:"phase,omitempty" xml:"phase,omitempty"`
	SandboxId             *string `json:"sandboxId,omitempty" xml:"sandboxId,omitempty"`
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
