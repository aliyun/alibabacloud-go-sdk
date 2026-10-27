// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMemoryStoreResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateMemoryStoreResponseBody
	GetCode() *string
	SetData(v *UpdateMemoryStoreResponseBodyData) *UpdateMemoryStoreResponseBody
	GetData() *UpdateMemoryStoreResponseBodyData
	SetHttpStatusCode(v int32) *UpdateMemoryStoreResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateMemoryStoreResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateMemoryStoreResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateMemoryStoreResponseBody
	GetSuccess() *bool
}

type UpdateMemoryStoreResponseBody struct {
	// The business status code.
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The MemoryStore details.
	Data *UpdateMemoryStoreResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
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

func (s UpdateMemoryStoreResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateMemoryStoreResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateMemoryStoreResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateMemoryStoreResponseBody) GetData() *UpdateMemoryStoreResponseBodyData {
	return s.Data
}

func (s *UpdateMemoryStoreResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateMemoryStoreResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateMemoryStoreResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateMemoryStoreResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateMemoryStoreResponseBody) SetCode(v string) *UpdateMemoryStoreResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateMemoryStoreResponseBody) SetData(v *UpdateMemoryStoreResponseBodyData) *UpdateMemoryStoreResponseBody {
	s.Data = v
	return s
}

func (s *UpdateMemoryStoreResponseBody) SetHttpStatusCode(v int32) *UpdateMemoryStoreResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateMemoryStoreResponseBody) SetMessage(v string) *UpdateMemoryStoreResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateMemoryStoreResponseBody) SetRequestId(v string) *UpdateMemoryStoreResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateMemoryStoreResponseBody) SetSuccess(v bool) *UpdateMemoryStoreResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateMemoryStoreResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateMemoryStoreResponseBodyData struct {
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
	// Used to save long-term user preferences
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
	// The lifecycle status of the MemoryStore.
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

func (s UpdateMemoryStoreResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UpdateMemoryStoreResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateMemoryStoreResponseBodyData) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *UpdateMemoryStoreResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *UpdateMemoryStoreResponseBodyData) GetExtractInstructions() *string {
	return s.ExtractInstructions
}

func (s *UpdateMemoryStoreResponseBodyData) GetMemoryStoreId() *string {
	return s.MemoryStoreId
}

func (s *UpdateMemoryStoreResponseBodyData) GetMemoryStoreName() *string {
	return s.MemoryStoreName
}

func (s *UpdateMemoryStoreResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *UpdateMemoryStoreResponseBodyData) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *UpdateMemoryStoreResponseBodyData) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *UpdateMemoryStoreResponseBodyData) SetCreatedAt(v string) *UpdateMemoryStoreResponseBodyData {
	s.CreatedAt = &v
	return s
}

func (s *UpdateMemoryStoreResponseBodyData) SetDescription(v string) *UpdateMemoryStoreResponseBodyData {
	s.Description = &v
	return s
}

func (s *UpdateMemoryStoreResponseBodyData) SetExtractInstructions(v string) *UpdateMemoryStoreResponseBodyData {
	s.ExtractInstructions = &v
	return s
}

func (s *UpdateMemoryStoreResponseBodyData) SetMemoryStoreId(v string) *UpdateMemoryStoreResponseBodyData {
	s.MemoryStoreId = &v
	return s
}

func (s *UpdateMemoryStoreResponseBodyData) SetMemoryStoreName(v string) *UpdateMemoryStoreResponseBodyData {
	s.MemoryStoreName = &v
	return s
}

func (s *UpdateMemoryStoreResponseBodyData) SetStatus(v string) *UpdateMemoryStoreResponseBodyData {
	s.Status = &v
	return s
}

func (s *UpdateMemoryStoreResponseBodyData) SetUpdatedAt(v string) *UpdateMemoryStoreResponseBodyData {
	s.UpdatedAt = &v
	return s
}

func (s *UpdateMemoryStoreResponseBodyData) SetWorkspaceId(v string) *UpdateMemoryStoreResponseBodyData {
	s.WorkspaceId = &v
	return s
}

func (s *UpdateMemoryStoreResponseBodyData) Validate() error {
	return dara.Validate(s)
}
