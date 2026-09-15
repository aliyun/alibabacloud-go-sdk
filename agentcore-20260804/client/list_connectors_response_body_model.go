// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListConnectorsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListConnectorsResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ListConnectorsResponseBody
	GetHttpStatusCode() *int32
	SetItems(v []*ListConnectorsResponseBodyItems) *ListConnectorsResponseBody
	GetItems() []*ListConnectorsResponseBodyItems
	SetMaxResults(v int32) *ListConnectorsResponseBody
	GetMaxResults() *int32
	SetMessage(v string) *ListConnectorsResponseBody
	GetMessage() *string
	SetNextToken(v string) *ListConnectorsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListConnectorsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListConnectorsResponseBody
	GetSuccess() *bool
	SetTotalCount(v int64) *ListConnectorsResponseBody
	GetTotalCount() *int64
}

type ListConnectorsResponseBody struct {
	// The business status code.
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The list of connectors.
	Items []*ListConnectorsResponseBodyItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// The number of entries returned in this request.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The response message.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The pagination token for the next page.
	//
	// example:
	//
	// dGVzdA==
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// request-123456
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// The total number of connectors.
	//
	// example:
	//
	// 10
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListConnectorsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListConnectorsResponseBody) GoString() string {
	return s.String()
}

func (s *ListConnectorsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListConnectorsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListConnectorsResponseBody) GetItems() []*ListConnectorsResponseBodyItems {
	return s.Items
}

func (s *ListConnectorsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListConnectorsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListConnectorsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListConnectorsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListConnectorsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListConnectorsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListConnectorsResponseBody) SetCode(v string) *ListConnectorsResponseBody {
	s.Code = &v
	return s
}

func (s *ListConnectorsResponseBody) SetHttpStatusCode(v int32) *ListConnectorsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListConnectorsResponseBody) SetItems(v []*ListConnectorsResponseBodyItems) *ListConnectorsResponseBody {
	s.Items = v
	return s
}

func (s *ListConnectorsResponseBody) SetMaxResults(v int32) *ListConnectorsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListConnectorsResponseBody) SetMessage(v string) *ListConnectorsResponseBody {
	s.Message = &v
	return s
}

func (s *ListConnectorsResponseBody) SetNextToken(v string) *ListConnectorsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListConnectorsResponseBody) SetRequestId(v string) *ListConnectorsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListConnectorsResponseBody) SetSuccess(v bool) *ListConnectorsResponseBody {
	s.Success = &v
	return s
}

func (s *ListConnectorsResponseBody) SetTotalCount(v int64) *ListConnectorsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListConnectorsResponseBody) Validate() error {
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

type ListConnectorsResponseBodyItems struct {
	// The number of agents bound to the connector.
	//
	// example:
	//
	// 3
	BoundAgentCount *int64 `json:"boundAgentCount,omitempty" xml:"boundAgentCount,omitempty"`
	// The time when the connector was enabled.
	//
	// example:
	//
	// 2026-09-01T08:00:00Z
	EnabledAt *string `json:"enabledAt,omitempty" xml:"enabledAt,omitempty"`
	// A JSON string. qodercli: {"site":"global|cn","organizationId":"...","apiKey":"...","serviceAccountKeys":[{"id":"ckey-xxx","name":"default","serviceAccountKey":"..."}]}. This field is absent when the connector is not enabled.
	//
	// example:
	//
	// {"site":"global","organizationId":"org-xxxx"}
	Metadata *string `json:"metadata,omitempty" xml:"metadata,omitempty"`
	// The connector name.
	//
	// example:
	//
	// qodercli
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The connector status.
	//
	// example:
	//
	// ENABLED
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s ListConnectorsResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s ListConnectorsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *ListConnectorsResponseBodyItems) GetBoundAgentCount() *int64 {
	return s.BoundAgentCount
}

func (s *ListConnectorsResponseBodyItems) GetEnabledAt() *string {
	return s.EnabledAt
}

func (s *ListConnectorsResponseBodyItems) GetMetadata() *string {
	return s.Metadata
}

func (s *ListConnectorsResponseBodyItems) GetName() *string {
	return s.Name
}

func (s *ListConnectorsResponseBodyItems) GetStatus() *string {
	return s.Status
}

func (s *ListConnectorsResponseBodyItems) SetBoundAgentCount(v int64) *ListConnectorsResponseBodyItems {
	s.BoundAgentCount = &v
	return s
}

func (s *ListConnectorsResponseBodyItems) SetEnabledAt(v string) *ListConnectorsResponseBodyItems {
	s.EnabledAt = &v
	return s
}

func (s *ListConnectorsResponseBodyItems) SetMetadata(v string) *ListConnectorsResponseBodyItems {
	s.Metadata = &v
	return s
}

func (s *ListConnectorsResponseBodyItems) SetName(v string) *ListConnectorsResponseBodyItems {
	s.Name = &v
	return s
}

func (s *ListConnectorsResponseBodyItems) SetStatus(v string) *ListConnectorsResponseBodyItems {
	s.Status = &v
	return s
}

func (s *ListConnectorsResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
