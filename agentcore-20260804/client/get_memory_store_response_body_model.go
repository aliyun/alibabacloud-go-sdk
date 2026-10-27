// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMemoryStoreResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetMemoryStoreResponseBody
	GetCode() *string
	SetData(v *GetMemoryStoreResponseBodyData) *GetMemoryStoreResponseBody
	GetData() *GetMemoryStoreResponseBodyData
	SetHttpStatusCode(v int32) *GetMemoryStoreResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetMemoryStoreResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetMemoryStoreResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetMemoryStoreResponseBody
	GetSuccess() *bool
}

type GetMemoryStoreResponseBody struct {
	// The business status code.
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The MemoryStore details.
	Data *GetMemoryStoreResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
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

func (s GetMemoryStoreResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetMemoryStoreResponseBody) GoString() string {
	return s.String()
}

func (s *GetMemoryStoreResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetMemoryStoreResponseBody) GetData() *GetMemoryStoreResponseBodyData {
	return s.Data
}

func (s *GetMemoryStoreResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetMemoryStoreResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetMemoryStoreResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetMemoryStoreResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetMemoryStoreResponseBody) SetCode(v string) *GetMemoryStoreResponseBody {
	s.Code = &v
	return s
}

func (s *GetMemoryStoreResponseBody) SetData(v *GetMemoryStoreResponseBodyData) *GetMemoryStoreResponseBody {
	s.Data = v
	return s
}

func (s *GetMemoryStoreResponseBody) SetHttpStatusCode(v int32) *GetMemoryStoreResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetMemoryStoreResponseBody) SetMessage(v string) *GetMemoryStoreResponseBody {
	s.Message = &v
	return s
}

func (s *GetMemoryStoreResponseBody) SetRequestId(v string) *GetMemoryStoreResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMemoryStoreResponseBody) SetSuccess(v bool) *GetMemoryStoreResponseBody {
	s.Success = &v
	return s
}

func (s *GetMemoryStoreResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetMemoryStoreResponseBodyData struct {
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
	// The last updated time.
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

func (s GetMemoryStoreResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetMemoryStoreResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetMemoryStoreResponseBodyData) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *GetMemoryStoreResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *GetMemoryStoreResponseBodyData) GetExtractInstructions() *string {
	return s.ExtractInstructions
}

func (s *GetMemoryStoreResponseBodyData) GetMemoryStoreId() *string {
	return s.MemoryStoreId
}

func (s *GetMemoryStoreResponseBodyData) GetMemoryStoreName() *string {
	return s.MemoryStoreName
}

func (s *GetMemoryStoreResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *GetMemoryStoreResponseBodyData) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *GetMemoryStoreResponseBodyData) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetMemoryStoreResponseBodyData) SetCreatedAt(v string) *GetMemoryStoreResponseBodyData {
	s.CreatedAt = &v
	return s
}

func (s *GetMemoryStoreResponseBodyData) SetDescription(v string) *GetMemoryStoreResponseBodyData {
	s.Description = &v
	return s
}

func (s *GetMemoryStoreResponseBodyData) SetExtractInstructions(v string) *GetMemoryStoreResponseBodyData {
	s.ExtractInstructions = &v
	return s
}

func (s *GetMemoryStoreResponseBodyData) SetMemoryStoreId(v string) *GetMemoryStoreResponseBodyData {
	s.MemoryStoreId = &v
	return s
}

func (s *GetMemoryStoreResponseBodyData) SetMemoryStoreName(v string) *GetMemoryStoreResponseBodyData {
	s.MemoryStoreName = &v
	return s
}

func (s *GetMemoryStoreResponseBodyData) SetStatus(v string) *GetMemoryStoreResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetMemoryStoreResponseBodyData) SetUpdatedAt(v string) *GetMemoryStoreResponseBodyData {
	s.UpdatedAt = &v
	return s
}

func (s *GetMemoryStoreResponseBodyData) SetWorkspaceId(v string) *GetMemoryStoreResponseBodyData {
	s.WorkspaceId = &v
	return s
}

func (s *GetMemoryStoreResponseBodyData) Validate() error {
	return dara.Validate(s)
}
