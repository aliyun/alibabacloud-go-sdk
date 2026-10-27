// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMemoryStoresResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListMemoryStoresResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ListMemoryStoresResponseBody
	GetHttpStatusCode() *int32
	SetItems(v []*ListMemoryStoresResponseBodyItems) *ListMemoryStoresResponseBody
	GetItems() []*ListMemoryStoresResponseBodyItems
	SetMaxResults(v int32) *ListMemoryStoresResponseBody
	GetMaxResults() *int32
	SetMessage(v string) *ListMemoryStoresResponseBody
	GetMessage() *string
	SetNextToken(v string) *ListMemoryStoresResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListMemoryStoresResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListMemoryStoresResponseBody
	GetSuccess() *bool
	SetTotalCount(v int64) *ListMemoryStoresResponseBody
	GetTotalCount() *int64
}

type ListMemoryStoresResponseBody struct {
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
	// The list of MemoryStores.
	Items []*ListMemoryStoresResponseBodyItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// The maximum number of records per page.
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
	// eyJvZmZzZXQiOjEwfQ==
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 01A00000-0000-0000-0000-000000000001
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// The total number of records.
	//
	// example:
	//
	// 1
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListMemoryStoresResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListMemoryStoresResponseBody) GoString() string {
	return s.String()
}

func (s *ListMemoryStoresResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListMemoryStoresResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListMemoryStoresResponseBody) GetItems() []*ListMemoryStoresResponseBodyItems {
	return s.Items
}

func (s *ListMemoryStoresResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListMemoryStoresResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListMemoryStoresResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListMemoryStoresResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListMemoryStoresResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListMemoryStoresResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListMemoryStoresResponseBody) SetCode(v string) *ListMemoryStoresResponseBody {
	s.Code = &v
	return s
}

func (s *ListMemoryStoresResponseBody) SetHttpStatusCode(v int32) *ListMemoryStoresResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListMemoryStoresResponseBody) SetItems(v []*ListMemoryStoresResponseBodyItems) *ListMemoryStoresResponseBody {
	s.Items = v
	return s
}

func (s *ListMemoryStoresResponseBody) SetMaxResults(v int32) *ListMemoryStoresResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListMemoryStoresResponseBody) SetMessage(v string) *ListMemoryStoresResponseBody {
	s.Message = &v
	return s
}

func (s *ListMemoryStoresResponseBody) SetNextToken(v string) *ListMemoryStoresResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListMemoryStoresResponseBody) SetRequestId(v string) *ListMemoryStoresResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMemoryStoresResponseBody) SetSuccess(v bool) *ListMemoryStoresResponseBody {
	s.Success = &v
	return s
}

func (s *ListMemoryStoresResponseBody) SetTotalCount(v int64) *ListMemoryStoresResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListMemoryStoresResponseBody) Validate() error {
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

type ListMemoryStoresResponseBodyItems struct {
	// The creation time.
	//
	// example:
	//
	// 2026-09-23T10:00:00Z
	CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// The MemoryStore description.
	//
	// example:
	//
	// Used to store long-term user preferences
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The custom extraction instructions.
	//
	// example:
	//
	// Extract only long-term preferences explicitly expressed by the user
	ExtractInstructions *string `json:"extractInstructions,omitempty" xml:"extractInstructions,omitempty"`
	// The MemoryStore ID.
	//
	// example:
	//
	// ms-2f5f9c5a8d0e4c3b
	MemoryStoreId *string `json:"memoryStoreId,omitempty" xml:"memoryStoreId,omitempty"`
	// The MemoryStore name. It cannot be modified after creation.
	//
	// example:
	//
	// customer-memory
	MemoryStoreName *string `json:"memoryStoreName,omitempty" xml:"memoryStoreName,omitempty"`
	// The lifecycle state of the MemoryStore.
	//
	// example:
	//
	// Ready
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The last update time.
	//
	// example:
	//
	// 2026-09-23T10:05:00Z
	UpdatedAt *string `json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
	// The AgentCore workspace ID.
	//
	// example:
	//
	// ws-2f5f9c5a8d0e4c3b
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s ListMemoryStoresResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s ListMemoryStoresResponseBodyItems) GoString() string {
	return s.String()
}

func (s *ListMemoryStoresResponseBodyItems) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *ListMemoryStoresResponseBodyItems) GetDescription() *string {
	return s.Description
}

func (s *ListMemoryStoresResponseBodyItems) GetExtractInstructions() *string {
	return s.ExtractInstructions
}

func (s *ListMemoryStoresResponseBodyItems) GetMemoryStoreId() *string {
	return s.MemoryStoreId
}

func (s *ListMemoryStoresResponseBodyItems) GetMemoryStoreName() *string {
	return s.MemoryStoreName
}

func (s *ListMemoryStoresResponseBodyItems) GetStatus() *string {
	return s.Status
}

func (s *ListMemoryStoresResponseBodyItems) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *ListMemoryStoresResponseBodyItems) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListMemoryStoresResponseBodyItems) SetCreatedAt(v string) *ListMemoryStoresResponseBodyItems {
	s.CreatedAt = &v
	return s
}

func (s *ListMemoryStoresResponseBodyItems) SetDescription(v string) *ListMemoryStoresResponseBodyItems {
	s.Description = &v
	return s
}

func (s *ListMemoryStoresResponseBodyItems) SetExtractInstructions(v string) *ListMemoryStoresResponseBodyItems {
	s.ExtractInstructions = &v
	return s
}

func (s *ListMemoryStoresResponseBodyItems) SetMemoryStoreId(v string) *ListMemoryStoresResponseBodyItems {
	s.MemoryStoreId = &v
	return s
}

func (s *ListMemoryStoresResponseBodyItems) SetMemoryStoreName(v string) *ListMemoryStoresResponseBodyItems {
	s.MemoryStoreName = &v
	return s
}

func (s *ListMemoryStoresResponseBodyItems) SetStatus(v string) *ListMemoryStoresResponseBodyItems {
	s.Status = &v
	return s
}

func (s *ListMemoryStoresResponseBodyItems) SetUpdatedAt(v string) *ListMemoryStoresResponseBodyItems {
	s.UpdatedAt = &v
	return s
}

func (s *ListMemoryStoresResponseBodyItems) SetWorkspaceId(v string) *ListMemoryStoresResponseBodyItems {
	s.WorkspaceId = &v
	return s
}

func (s *ListMemoryStoresResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
