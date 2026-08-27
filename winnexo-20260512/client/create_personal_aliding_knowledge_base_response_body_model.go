// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePersonalAlidingKnowledgeBaseResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreatePersonalAlidingKnowledgeBaseResponseBody
	GetCode() *string
	SetDirectoryId(v string) *CreatePersonalAlidingKnowledgeBaseResponseBody
	GetDirectoryId() *string
	SetGmtCreate(v string) *CreatePersonalAlidingKnowledgeBaseResponseBody
	GetGmtCreate() *string
	SetKbUrl(v string) *CreatePersonalAlidingKnowledgeBaseResponseBody
	GetKbUrl() *string
	SetMessage(v string) *CreatePersonalAlidingKnowledgeBaseResponseBody
	GetMessage() *string
	SetName(v string) *CreatePersonalAlidingKnowledgeBaseResponseBody
	GetName() *string
	SetOperatingObjectName(v string) *CreatePersonalAlidingKnowledgeBaseResponseBody
	GetOperatingObjectName() *string
	SetRequestId(v string) *CreatePersonalAlidingKnowledgeBaseResponseBody
	GetRequestId() *string
	SetStatus(v string) *CreatePersonalAlidingKnowledgeBaseResponseBody
	GetStatus() *string
}

type CreatePersonalAlidingKnowledgeBaseResponseBody struct {
	// The status code.
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
	// The creation time in ISO 8601 format.
	//
	// example:
	//
	// 2025-11-14T02:18:27Z
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// The knowledge base URL (echoed from the request parameter for caller alignment).
	//
	// example:
	//
	// https://example.com/winnexo/resource
	KbUrl *string `json:"kbUrl,omitempty" xml:"kbUrl,omitempty"`
	// The response message.
	//
	// example:
	//
	// ok
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The name of the AI assistant.
	//
	// example:
	//
	// p-toolset-3dcef7ca-31b9-4d1c-8692-1ef03099cad3
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The name of the digital employee (operating object name, optional).
	//
	// example:
	//
	// string_value
	OperatingObjectName *string `json:"operatingObjectName,omitempty" xml:"operatingObjectName,omitempty"`
	// The request ID.
	//
	// example:
	//
	// E68654BD-F7BA-5837-8686-5645D739A47C
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The status. Valid values:
	//
	// - 200: Success.
	//
	// - 500: Failure.
	//
	// example:
	//
	// 200
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s CreatePersonalAlidingKnowledgeBaseResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreatePersonalAlidingKnowledgeBaseResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePersonalAlidingKnowledgeBaseResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreatePersonalAlidingKnowledgeBaseResponseBody) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *CreatePersonalAlidingKnowledgeBaseResponseBody) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *CreatePersonalAlidingKnowledgeBaseResponseBody) GetKbUrl() *string {
	return s.KbUrl
}

func (s *CreatePersonalAlidingKnowledgeBaseResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreatePersonalAlidingKnowledgeBaseResponseBody) GetName() *string {
	return s.Name
}

func (s *CreatePersonalAlidingKnowledgeBaseResponseBody) GetOperatingObjectName() *string {
	return s.OperatingObjectName
}

func (s *CreatePersonalAlidingKnowledgeBaseResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreatePersonalAlidingKnowledgeBaseResponseBody) GetStatus() *string {
	return s.Status
}

func (s *CreatePersonalAlidingKnowledgeBaseResponseBody) SetCode(v string) *CreatePersonalAlidingKnowledgeBaseResponseBody {
	s.Code = &v
	return s
}

func (s *CreatePersonalAlidingKnowledgeBaseResponseBody) SetDirectoryId(v string) *CreatePersonalAlidingKnowledgeBaseResponseBody {
	s.DirectoryId = &v
	return s
}

func (s *CreatePersonalAlidingKnowledgeBaseResponseBody) SetGmtCreate(v string) *CreatePersonalAlidingKnowledgeBaseResponseBody {
	s.GmtCreate = &v
	return s
}

func (s *CreatePersonalAlidingKnowledgeBaseResponseBody) SetKbUrl(v string) *CreatePersonalAlidingKnowledgeBaseResponseBody {
	s.KbUrl = &v
	return s
}

func (s *CreatePersonalAlidingKnowledgeBaseResponseBody) SetMessage(v string) *CreatePersonalAlidingKnowledgeBaseResponseBody {
	s.Message = &v
	return s
}

func (s *CreatePersonalAlidingKnowledgeBaseResponseBody) SetName(v string) *CreatePersonalAlidingKnowledgeBaseResponseBody {
	s.Name = &v
	return s
}

func (s *CreatePersonalAlidingKnowledgeBaseResponseBody) SetOperatingObjectName(v string) *CreatePersonalAlidingKnowledgeBaseResponseBody {
	s.OperatingObjectName = &v
	return s
}

func (s *CreatePersonalAlidingKnowledgeBaseResponseBody) SetRequestId(v string) *CreatePersonalAlidingKnowledgeBaseResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreatePersonalAlidingKnowledgeBaseResponseBody) SetStatus(v string) *CreatePersonalAlidingKnowledgeBaseResponseBody {
	s.Status = &v
	return s
}

func (s *CreatePersonalAlidingKnowledgeBaseResponseBody) Validate() error {
	return dara.Validate(s)
}
