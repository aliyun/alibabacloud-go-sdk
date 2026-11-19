// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMemoryStoreResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateMemoryStoreResponseBody
	GetCode() *string
	SetData(v *CreateMemoryStoreResponseBodyData) *CreateMemoryStoreResponseBody
	GetData() *CreateMemoryStoreResponseBodyData
	SetHttpStatusCode(v int32) *CreateMemoryStoreResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateMemoryStoreResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateMemoryStoreResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateMemoryStoreResponseBody
	GetSuccess() *bool
}

type CreateMemoryStoreResponseBody struct {
	// The business status code.
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The details of the memory store.
	Data *CreateMemoryStoreResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The response message.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 01A00000-0000-0000-0000-000000000001
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request is successful.
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CreateMemoryStoreResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateMemoryStoreResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMemoryStoreResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateMemoryStoreResponseBody) GetData() *CreateMemoryStoreResponseBodyData {
	return s.Data
}

func (s *CreateMemoryStoreResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateMemoryStoreResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateMemoryStoreResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateMemoryStoreResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateMemoryStoreResponseBody) SetCode(v string) *CreateMemoryStoreResponseBody {
	s.Code = &v
	return s
}

func (s *CreateMemoryStoreResponseBody) SetData(v *CreateMemoryStoreResponseBodyData) *CreateMemoryStoreResponseBody {
	s.Data = v
	return s
}

func (s *CreateMemoryStoreResponseBody) SetHttpStatusCode(v int32) *CreateMemoryStoreResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateMemoryStoreResponseBody) SetMessage(v string) *CreateMemoryStoreResponseBody {
	s.Message = &v
	return s
}

func (s *CreateMemoryStoreResponseBody) SetRequestId(v string) *CreateMemoryStoreResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateMemoryStoreResponseBody) SetSuccess(v bool) *CreateMemoryStoreResponseBody {
	s.Success = &v
	return s
}

func (s *CreateMemoryStoreResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateMemoryStoreResponseBodyData struct {
	// The creation time. The time is in RFC 3339 format. The Z at the end of the example indicates the UTC time zone.
	//
	// example:
	//
	// 2026-09-23T10:00:00Z
	CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// The description of the memory store.
	//
	// example:
	//
	// Used to store the long-term preferences of users
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The number of events. You can query this value by calling GetMemoryStore.
	//
	// example:
	//
	// 20
	EventCount *int64 `json:"eventCount,omitempty" xml:"eventCount,omitempty"`
	// The custom extraction instructions.
	//
	// example:
	//
	// Extract only long-term preferences explicitly expressed by users
	ExtractInstructions *string `json:"extractInstructions,omitempty" xml:"extractInstructions,omitempty"`
	// The number of valid memories. You can query this value by calling GetMemoryStore.
	//
	// example:
	//
	// 10
	MemoryCount *int64 `json:"memoryCount,omitempty" xml:"memoryCount,omitempty"`
	// The ID of the memory store.
	//
	// example:
	//
	// ms-2f5f9c5a8d0e4c3b
	MemoryStoreId *string `json:"memoryStoreId,omitempty" xml:"memoryStoreId,omitempty"`
	// The name of the memory store. The name cannot be modified after creation.
	//
	// example:
	//
	// customer-memory
	MemoryStoreName *string `json:"memoryStoreName,omitempty" xml:"memoryStoreName,omitempty"`
	// The lifecycle status of the memory store.
	//
	// example:
	//
	// Ready
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The last updated time. The time is in RFC 3339 format. The Z at the end of the example indicates the UTC time zone.
	//
	// example:
	//
	// 2026-09-23T10:05:00Z
	UpdatedAt *string `json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
	// The ID of the AgentCore workspace.
	//
	// example:
	//
	// ws-2f5f9c5a8d0e4c3b
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s CreateMemoryStoreResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateMemoryStoreResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateMemoryStoreResponseBodyData) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *CreateMemoryStoreResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *CreateMemoryStoreResponseBodyData) GetEventCount() *int64 {
	return s.EventCount
}

func (s *CreateMemoryStoreResponseBodyData) GetExtractInstructions() *string {
	return s.ExtractInstructions
}

func (s *CreateMemoryStoreResponseBodyData) GetMemoryCount() *int64 {
	return s.MemoryCount
}

func (s *CreateMemoryStoreResponseBodyData) GetMemoryStoreId() *string {
	return s.MemoryStoreId
}

func (s *CreateMemoryStoreResponseBodyData) GetMemoryStoreName() *string {
	return s.MemoryStoreName
}

func (s *CreateMemoryStoreResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *CreateMemoryStoreResponseBodyData) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *CreateMemoryStoreResponseBodyData) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *CreateMemoryStoreResponseBodyData) SetCreatedAt(v string) *CreateMemoryStoreResponseBodyData {
	s.CreatedAt = &v
	return s
}

func (s *CreateMemoryStoreResponseBodyData) SetDescription(v string) *CreateMemoryStoreResponseBodyData {
	s.Description = &v
	return s
}

func (s *CreateMemoryStoreResponseBodyData) SetEventCount(v int64) *CreateMemoryStoreResponseBodyData {
	s.EventCount = &v
	return s
}

func (s *CreateMemoryStoreResponseBodyData) SetExtractInstructions(v string) *CreateMemoryStoreResponseBodyData {
	s.ExtractInstructions = &v
	return s
}

func (s *CreateMemoryStoreResponseBodyData) SetMemoryCount(v int64) *CreateMemoryStoreResponseBodyData {
	s.MemoryCount = &v
	return s
}

func (s *CreateMemoryStoreResponseBodyData) SetMemoryStoreId(v string) *CreateMemoryStoreResponseBodyData {
	s.MemoryStoreId = &v
	return s
}

func (s *CreateMemoryStoreResponseBodyData) SetMemoryStoreName(v string) *CreateMemoryStoreResponseBodyData {
	s.MemoryStoreName = &v
	return s
}

func (s *CreateMemoryStoreResponseBodyData) SetStatus(v string) *CreateMemoryStoreResponseBodyData {
	s.Status = &v
	return s
}

func (s *CreateMemoryStoreResponseBodyData) SetUpdatedAt(v string) *CreateMemoryStoreResponseBodyData {
	s.UpdatedAt = &v
	return s
}

func (s *CreateMemoryStoreResponseBodyData) SetWorkspaceId(v string) *CreateMemoryStoreResponseBodyData {
	s.WorkspaceId = &v
	return s
}

func (s *CreateMemoryStoreResponseBodyData) Validate() error {
	return dara.Validate(s)
}
