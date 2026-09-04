// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateKnowledgeBaseAliDingDocResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateKnowledgeBaseAliDingDocResponseBody
	GetCode() *string
	SetDirectoryId(v string) *CreateKnowledgeBaseAliDingDocResponseBody
	GetDirectoryId() *string
	SetFilePublicUrl(v string) *CreateKnowledgeBaseAliDingDocResponseBody
	GetFilePublicUrl() *string
	SetGmtCreate(v string) *CreateKnowledgeBaseAliDingDocResponseBody
	GetGmtCreate() *string
	SetMessage(v string) *CreateKnowledgeBaseAliDingDocResponseBody
	GetMessage() *string
	SetName(v string) *CreateKnowledgeBaseAliDingDocResponseBody
	GetName() *string
	SetRequestId(v string) *CreateKnowledgeBaseAliDingDocResponseBody
	GetRequestId() *string
	SetScope(v string) *CreateKnowledgeBaseAliDingDocResponseBody
	GetScope() *string
	SetSourceId(v string) *CreateKnowledgeBaseAliDingDocResponseBody
	GetSourceId() *string
	SetStatus(v string) *CreateKnowledgeBaseAliDingDocResponseBody
	GetStatus() *string
}

type CreateKnowledgeBaseAliDingDocResponseBody struct {
	// The result code.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The directory ID.
	//
	// example:
	//
	// exampleDirectoryId
	DirectoryId *string `json:"directoryId,omitempty" xml:"directoryId,omitempty"`
	// The public URL of the document (echoes the input parameter).
	//
	// example:
	//
	// https://example.com/winnexo/resource
	FilePublicUrl *string `json:"filePublicUrl,omitempty" xml:"filePublicUrl,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2026-05-22 16:03:27
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// The error details.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The name.
	//
	// example:
	//
	// p-toolset-b8a1de80-e9f5-49f3-8a12-873d378889c6
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The request ID.
	//
	// example:
	//
	// F42FC60B-C54D-5DFB-A8EC-04625BFFF1F7
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The permission scope.
	//
	// example:
	//
	// repo
	Scope *string `json:"scope,omitempty" xml:"scope,omitempty"`
	// The unique identifier on the business system side, which is the business ID.
	//
	// example:
	//
	// 8
	SourceId *string `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	// The refund status. Query this field to confirm the refund status during processing. Valid values:
	//
	// - SUCCESS: All refunds are successful.
	//
	// - FAIL: The refund failed.
	//
	// - WAIT_PAY: Waiting for refund.
	//
	// - EXPIRE: The refund has expired.
	//
	// - PAYING: The refund is being processed.
	//
	// - TERMINATE: The refund is terminated.
	//
	// example:
	//
	// {\\"observedGeneration\\": 7, \\"servicesInstances\\": {}, \\"observedTime\\": \\"2025-12-17T11:57:07Z\\", \\"servicesWithPendingChanges\\": [], \\"latestEnvironmentDeploymentName\\": \\"manual-1765972627273-k7GZvr\\"}
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s CreateKnowledgeBaseAliDingDocResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateKnowledgeBaseAliDingDocResponseBody) GoString() string {
	return s.String()
}

func (s *CreateKnowledgeBaseAliDingDocResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateKnowledgeBaseAliDingDocResponseBody) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *CreateKnowledgeBaseAliDingDocResponseBody) GetFilePublicUrl() *string {
	return s.FilePublicUrl
}

func (s *CreateKnowledgeBaseAliDingDocResponseBody) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *CreateKnowledgeBaseAliDingDocResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateKnowledgeBaseAliDingDocResponseBody) GetName() *string {
	return s.Name
}

func (s *CreateKnowledgeBaseAliDingDocResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateKnowledgeBaseAliDingDocResponseBody) GetScope() *string {
	return s.Scope
}

func (s *CreateKnowledgeBaseAliDingDocResponseBody) GetSourceId() *string {
	return s.SourceId
}

func (s *CreateKnowledgeBaseAliDingDocResponseBody) GetStatus() *string {
	return s.Status
}

func (s *CreateKnowledgeBaseAliDingDocResponseBody) SetCode(v string) *CreateKnowledgeBaseAliDingDocResponseBody {
	s.Code = &v
	return s
}

func (s *CreateKnowledgeBaseAliDingDocResponseBody) SetDirectoryId(v string) *CreateKnowledgeBaseAliDingDocResponseBody {
	s.DirectoryId = &v
	return s
}

func (s *CreateKnowledgeBaseAliDingDocResponseBody) SetFilePublicUrl(v string) *CreateKnowledgeBaseAliDingDocResponseBody {
	s.FilePublicUrl = &v
	return s
}

func (s *CreateKnowledgeBaseAliDingDocResponseBody) SetGmtCreate(v string) *CreateKnowledgeBaseAliDingDocResponseBody {
	s.GmtCreate = &v
	return s
}

func (s *CreateKnowledgeBaseAliDingDocResponseBody) SetMessage(v string) *CreateKnowledgeBaseAliDingDocResponseBody {
	s.Message = &v
	return s
}

func (s *CreateKnowledgeBaseAliDingDocResponseBody) SetName(v string) *CreateKnowledgeBaseAliDingDocResponseBody {
	s.Name = &v
	return s
}

func (s *CreateKnowledgeBaseAliDingDocResponseBody) SetRequestId(v string) *CreateKnowledgeBaseAliDingDocResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateKnowledgeBaseAliDingDocResponseBody) SetScope(v string) *CreateKnowledgeBaseAliDingDocResponseBody {
	s.Scope = &v
	return s
}

func (s *CreateKnowledgeBaseAliDingDocResponseBody) SetSourceId(v string) *CreateKnowledgeBaseAliDingDocResponseBody {
	s.SourceId = &v
	return s
}

func (s *CreateKnowledgeBaseAliDingDocResponseBody) SetStatus(v string) *CreateKnowledgeBaseAliDingDocResponseBody {
	s.Status = &v
	return s
}

func (s *CreateKnowledgeBaseAliDingDocResponseBody) Validate() error {
	return dara.Validate(s)
}
