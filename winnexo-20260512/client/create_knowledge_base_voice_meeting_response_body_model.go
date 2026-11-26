// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateKnowledgeBaseVoiceMeetingResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateKnowledgeBaseVoiceMeetingResponseBody
	GetCode() *string
	SetDirectoryId(v string) *CreateKnowledgeBaseVoiceMeetingResponseBody
	GetDirectoryId() *string
	SetGmtCreate(v string) *CreateKnowledgeBaseVoiceMeetingResponseBody
	GetGmtCreate() *string
	SetMessage(v string) *CreateKnowledgeBaseVoiceMeetingResponseBody
	GetMessage() *string
	SetName(v string) *CreateKnowledgeBaseVoiceMeetingResponseBody
	GetName() *string
	SetRequestId(v string) *CreateKnowledgeBaseVoiceMeetingResponseBody
	GetRequestId() *string
	SetScope(v string) *CreateKnowledgeBaseVoiceMeetingResponseBody
	GetScope() *string
	SetSourceId(v string) *CreateKnowledgeBaseVoiceMeetingResponseBody
	GetSourceId() *string
	SetStatus(v string) *CreateKnowledgeBaseVoiceMeetingResponseBody
	GetStatus() *string
}

type CreateKnowledgeBaseVoiceMeetingResponseBody struct {
	// The status code.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The ID of the enterprise knowledge base directory that is actually bound.
	//
	// example:
	//
	// exampleDirectoryId
	DirectoryId *string `json:"directoryId,omitempty" xml:"directoryId,omitempty"`
	// The creation time in ISO 8601 format.
	//
	// example:
	//
	// 2026-11-18T09:00:00+08:00
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// The description of the status code.
	//
	// example:
	//
	// ok
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The file name.
	//
	// example:
	//
	// Project Review Recording
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The scope of the resource, which is fixed to TENANT.
	//
	// example:
	//
	// TENANT
	Scope *string `json:"scope,omitempty" xml:"scope,omitempty"`
	// The unique identifier on the business system side, which is the business ID.
	//
	// example:
	//
	// exampleSourceId
	SourceId *string `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	// The actual status of the resource.
	//
	// example:
	//
	// RUNNING
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s CreateKnowledgeBaseVoiceMeetingResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateKnowledgeBaseVoiceMeetingResponseBody) GoString() string {
	return s.String()
}

func (s *CreateKnowledgeBaseVoiceMeetingResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateKnowledgeBaseVoiceMeetingResponseBody) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *CreateKnowledgeBaseVoiceMeetingResponseBody) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *CreateKnowledgeBaseVoiceMeetingResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateKnowledgeBaseVoiceMeetingResponseBody) GetName() *string {
	return s.Name
}

func (s *CreateKnowledgeBaseVoiceMeetingResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateKnowledgeBaseVoiceMeetingResponseBody) GetScope() *string {
	return s.Scope
}

func (s *CreateKnowledgeBaseVoiceMeetingResponseBody) GetSourceId() *string {
	return s.SourceId
}

func (s *CreateKnowledgeBaseVoiceMeetingResponseBody) GetStatus() *string {
	return s.Status
}

func (s *CreateKnowledgeBaseVoiceMeetingResponseBody) SetCode(v string) *CreateKnowledgeBaseVoiceMeetingResponseBody {
	s.Code = &v
	return s
}

func (s *CreateKnowledgeBaseVoiceMeetingResponseBody) SetDirectoryId(v string) *CreateKnowledgeBaseVoiceMeetingResponseBody {
	s.DirectoryId = &v
	return s
}

func (s *CreateKnowledgeBaseVoiceMeetingResponseBody) SetGmtCreate(v string) *CreateKnowledgeBaseVoiceMeetingResponseBody {
	s.GmtCreate = &v
	return s
}

func (s *CreateKnowledgeBaseVoiceMeetingResponseBody) SetMessage(v string) *CreateKnowledgeBaseVoiceMeetingResponseBody {
	s.Message = &v
	return s
}

func (s *CreateKnowledgeBaseVoiceMeetingResponseBody) SetName(v string) *CreateKnowledgeBaseVoiceMeetingResponseBody {
	s.Name = &v
	return s
}

func (s *CreateKnowledgeBaseVoiceMeetingResponseBody) SetRequestId(v string) *CreateKnowledgeBaseVoiceMeetingResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateKnowledgeBaseVoiceMeetingResponseBody) SetScope(v string) *CreateKnowledgeBaseVoiceMeetingResponseBody {
	s.Scope = &v
	return s
}

func (s *CreateKnowledgeBaseVoiceMeetingResponseBody) SetSourceId(v string) *CreateKnowledgeBaseVoiceMeetingResponseBody {
	s.SourceId = &v
	return s
}

func (s *CreateKnowledgeBaseVoiceMeetingResponseBody) SetStatus(v string) *CreateKnowledgeBaseVoiceMeetingResponseBody {
	s.Status = &v
	return s
}

func (s *CreateKnowledgeBaseVoiceMeetingResponseBody) Validate() error {
	return dara.Validate(s)
}
