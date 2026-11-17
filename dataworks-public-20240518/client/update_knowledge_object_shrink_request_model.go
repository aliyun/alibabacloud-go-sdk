// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateKnowledgeObjectShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *UpdateKnowledgeObjectShrinkRequest
	GetClientToken() *string
	SetContentShrink(v string) *UpdateKnowledgeObjectShrinkRequest
	GetContentShrink() *string
	SetDescription(v string) *UpdateKnowledgeObjectShrinkRequest
	GetDescription() *string
	SetExpectedModifyTime(v int64) *UpdateKnowledgeObjectShrinkRequest
	GetExpectedModifyTime() *int64
	SetKnowledgeBaseId(v string) *UpdateKnowledgeObjectShrinkRequest
	GetKnowledgeBaseId() *string
	SetKnowledgeObjectId(v string) *UpdateKnowledgeObjectShrinkRequest
	GetKnowledgeObjectId() *string
	SetName(v string) *UpdateKnowledgeObjectShrinkRequest
	GetName() *string
	SetUploadId(v string) *UpdateKnowledgeObjectShrinkRequest
	GetUploadId() *string
	SetVersionNo(v int64) *UpdateKnowledgeObjectShrinkRequest
	GetVersionNo() *int64
}

type UpdateKnowledgeObjectShrinkRequest struct {
	// The client token. This parameter is supported only for DOCUMENT. Model updates do not guarantee exactly-once remote save notifications.
	//
	// example:
	//
	// edit-document-text-001
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The typed content update. Object properties are submitted as needed. Arrays are replaced in their entirety when submitted. Unspecified properties are retained.
	ContentShrink *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The new description. An empty string clears the description. If this parameter is not specified, the original value is retained.
	//
	// example:
	//
	// Actual paid amount
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The expected modification time. This parameter is supported only for DOCUMENT. It is required when modifying the body text and optional when modifying metadata.
	//
	// example:
	//
	// 1790200800000
	ExpectedModifyTime *int64 `json:"ExpectedModifyTime,omitempty" xml:"ExpectedModifyTime,omitempty"`
	// The ID of the knowledge base to which the object belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12345678-1234-4234-8234-123456789abc
	KnowledgeBaseId *string `json:"KnowledgeBaseId,omitempty" xml:"KnowledgeBaseId,omitempty"`
	// The stable knowledge object ID returned by the Get or List operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// obj-model-001
	KnowledgeObjectId *string `json:"KnowledgeObjectId,omitempty" xml:"KnowledgeObjectId,omitempty"`
	// The new name. This parameter cannot be empty. If this parameter is not specified, the original value is retained.
	//
	// example:
	//
	// orders
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The upload ID. This parameter is supported only for SEMANTIC_MODEL. It replaces the public model content with a complete YAML file and is mutually exclusive with Content.
	//
	// example:
	//
	// 12345678-1234-4234-8234-123456789abd
	UploadId *string `json:"UploadId,omitempty" xml:"UploadId,omitempty"`
	// The version number. This parameter is required when updating Content or UploadId for SEMANTIC_MODEL. Use the VersionNo returned by the Get operation. This is not an optimistic lock.
	//
	// example:
	//
	// 1
	VersionNo *int64 `json:"VersionNo,omitempty" xml:"VersionNo,omitempty"`
}

func (s UpdateKnowledgeObjectShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateKnowledgeObjectShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateKnowledgeObjectShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateKnowledgeObjectShrinkRequest) GetContentShrink() *string {
	return s.ContentShrink
}

func (s *UpdateKnowledgeObjectShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateKnowledgeObjectShrinkRequest) GetExpectedModifyTime() *int64 {
	return s.ExpectedModifyTime
}

func (s *UpdateKnowledgeObjectShrinkRequest) GetKnowledgeBaseId() *string {
	return s.KnowledgeBaseId
}

func (s *UpdateKnowledgeObjectShrinkRequest) GetKnowledgeObjectId() *string {
	return s.KnowledgeObjectId
}

func (s *UpdateKnowledgeObjectShrinkRequest) GetName() *string {
	return s.Name
}

func (s *UpdateKnowledgeObjectShrinkRequest) GetUploadId() *string {
	return s.UploadId
}

func (s *UpdateKnowledgeObjectShrinkRequest) GetVersionNo() *int64 {
	return s.VersionNo
}

func (s *UpdateKnowledgeObjectShrinkRequest) SetClientToken(v string) *UpdateKnowledgeObjectShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateKnowledgeObjectShrinkRequest) SetContentShrink(v string) *UpdateKnowledgeObjectShrinkRequest {
	s.ContentShrink = &v
	return s
}

func (s *UpdateKnowledgeObjectShrinkRequest) SetDescription(v string) *UpdateKnowledgeObjectShrinkRequest {
	s.Description = &v
	return s
}

func (s *UpdateKnowledgeObjectShrinkRequest) SetExpectedModifyTime(v int64) *UpdateKnowledgeObjectShrinkRequest {
	s.ExpectedModifyTime = &v
	return s
}

func (s *UpdateKnowledgeObjectShrinkRequest) SetKnowledgeBaseId(v string) *UpdateKnowledgeObjectShrinkRequest {
	s.KnowledgeBaseId = &v
	return s
}

func (s *UpdateKnowledgeObjectShrinkRequest) SetKnowledgeObjectId(v string) *UpdateKnowledgeObjectShrinkRequest {
	s.KnowledgeObjectId = &v
	return s
}

func (s *UpdateKnowledgeObjectShrinkRequest) SetName(v string) *UpdateKnowledgeObjectShrinkRequest {
	s.Name = &v
	return s
}

func (s *UpdateKnowledgeObjectShrinkRequest) SetUploadId(v string) *UpdateKnowledgeObjectShrinkRequest {
	s.UploadId = &v
	return s
}

func (s *UpdateKnowledgeObjectShrinkRequest) SetVersionNo(v int64) *UpdateKnowledgeObjectShrinkRequest {
	s.VersionNo = &v
	return s
}

func (s *UpdateKnowledgeObjectShrinkRequest) Validate() error {
	return dara.Validate(s)
}
