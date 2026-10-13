// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMemoryStoreResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteMemoryStoreResponseBody
	GetCode() *string
	SetData(v *DeleteMemoryStoreResponseBodyData) *DeleteMemoryStoreResponseBody
	GetData() *DeleteMemoryStoreResponseBodyData
	SetHttpStatusCode(v int32) *DeleteMemoryStoreResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteMemoryStoreResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteMemoryStoreResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteMemoryStoreResponseBody
	GetSuccess() *bool
}

type DeleteMemoryStoreResponseBody struct {
	// The business status code.
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The MemoryStore details.
	Data *DeleteMemoryStoreResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
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

func (s DeleteMemoryStoreResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteMemoryStoreResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteMemoryStoreResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteMemoryStoreResponseBody) GetData() *DeleteMemoryStoreResponseBodyData {
	return s.Data
}

func (s *DeleteMemoryStoreResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteMemoryStoreResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteMemoryStoreResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteMemoryStoreResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteMemoryStoreResponseBody) SetCode(v string) *DeleteMemoryStoreResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteMemoryStoreResponseBody) SetData(v *DeleteMemoryStoreResponseBodyData) *DeleteMemoryStoreResponseBody {
	s.Data = v
	return s
}

func (s *DeleteMemoryStoreResponseBody) SetHttpStatusCode(v int32) *DeleteMemoryStoreResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteMemoryStoreResponseBody) SetMessage(v string) *DeleteMemoryStoreResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteMemoryStoreResponseBody) SetRequestId(v string) *DeleteMemoryStoreResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteMemoryStoreResponseBody) SetSuccess(v bool) *DeleteMemoryStoreResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteMemoryStoreResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteMemoryStoreResponseBodyData struct {
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

func (s DeleteMemoryStoreResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DeleteMemoryStoreResponseBodyData) GoString() string {
	return s.String()
}

func (s *DeleteMemoryStoreResponseBodyData) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *DeleteMemoryStoreResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *DeleteMemoryStoreResponseBodyData) GetExtractInstructions() *string {
	return s.ExtractInstructions
}

func (s *DeleteMemoryStoreResponseBodyData) GetMemoryStoreId() *string {
	return s.MemoryStoreId
}

func (s *DeleteMemoryStoreResponseBodyData) GetMemoryStoreName() *string {
	return s.MemoryStoreName
}

func (s *DeleteMemoryStoreResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *DeleteMemoryStoreResponseBodyData) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *DeleteMemoryStoreResponseBodyData) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *DeleteMemoryStoreResponseBodyData) SetCreatedAt(v string) *DeleteMemoryStoreResponseBodyData {
	s.CreatedAt = &v
	return s
}

func (s *DeleteMemoryStoreResponseBodyData) SetDescription(v string) *DeleteMemoryStoreResponseBodyData {
	s.Description = &v
	return s
}

func (s *DeleteMemoryStoreResponseBodyData) SetExtractInstructions(v string) *DeleteMemoryStoreResponseBodyData {
	s.ExtractInstructions = &v
	return s
}

func (s *DeleteMemoryStoreResponseBodyData) SetMemoryStoreId(v string) *DeleteMemoryStoreResponseBodyData {
	s.MemoryStoreId = &v
	return s
}

func (s *DeleteMemoryStoreResponseBodyData) SetMemoryStoreName(v string) *DeleteMemoryStoreResponseBodyData {
	s.MemoryStoreName = &v
	return s
}

func (s *DeleteMemoryStoreResponseBodyData) SetStatus(v string) *DeleteMemoryStoreResponseBodyData {
	s.Status = &v
	return s
}

func (s *DeleteMemoryStoreResponseBodyData) SetUpdatedAt(v string) *DeleteMemoryStoreResponseBodyData {
	s.UpdatedAt = &v
	return s
}

func (s *DeleteMemoryStoreResponseBodyData) SetWorkspaceId(v string) *DeleteMemoryStoreResponseBodyData {
	s.WorkspaceId = &v
	return s
}

func (s *DeleteMemoryStoreResponseBodyData) Validate() error {
	return dara.Validate(s)
}
