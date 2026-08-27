// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateKnowledgeBaseTextResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateKnowledgeBaseTextResponseBody
	GetCode() *string
	SetDirectoryId(v string) *CreateKnowledgeBaseTextResponseBody
	GetDirectoryId() *string
	SetGmtCreate(v string) *CreateKnowledgeBaseTextResponseBody
	GetGmtCreate() *string
	SetMessage(v string) *CreateKnowledgeBaseTextResponseBody
	GetMessage() *string
	SetName(v string) *CreateKnowledgeBaseTextResponseBody
	GetName() *string
	SetRequestId(v string) *CreateKnowledgeBaseTextResponseBody
	GetRequestId() *string
	SetScope(v string) *CreateKnowledgeBaseTextResponseBody
	GetScope() *string
	SetSourceId(v string) *CreateKnowledgeBaseTextResponseBody
	GetSourceId() *string
	SetStatus(v string) *CreateKnowledgeBaseTextResponseBody
	GetStatus() *string
}

type CreateKnowledgeBaseTextResponseBody struct {
	// The status code.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The folder ID.
	//
	// example:
	//
	// exampleDirectoryId
	DirectoryId *string `json:"directoryId,omitempty" xml:"directoryId,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2026-04-22T07:10:40.000+00:00
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// The error message returned when the request fails.
	//
	// example:
	//
	// The current zone list is illegal.
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The image name.
	//
	// example:
	//
	// p-default-af484147-e026-487b-a9eb-bd25464f0667
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9005F6D0-748F-559D-ABDA-F4F31B983316
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The permission scope.
	//
	// example:
	//
	// read:user,read:repo,write:repo,read:org,read:group
	Scope *string `json:"scope,omitempty" xml:"scope,omitempty"`
	// The data source ID.
	//
	// example:
	//
	// 8
	SourceId *string `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	// The status.
	//
	// example:
	//
	// 200
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s CreateKnowledgeBaseTextResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateKnowledgeBaseTextResponseBody) GoString() string {
	return s.String()
}

func (s *CreateKnowledgeBaseTextResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateKnowledgeBaseTextResponseBody) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *CreateKnowledgeBaseTextResponseBody) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *CreateKnowledgeBaseTextResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateKnowledgeBaseTextResponseBody) GetName() *string {
	return s.Name
}

func (s *CreateKnowledgeBaseTextResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateKnowledgeBaseTextResponseBody) GetScope() *string {
	return s.Scope
}

func (s *CreateKnowledgeBaseTextResponseBody) GetSourceId() *string {
	return s.SourceId
}

func (s *CreateKnowledgeBaseTextResponseBody) GetStatus() *string {
	return s.Status
}

func (s *CreateKnowledgeBaseTextResponseBody) SetCode(v string) *CreateKnowledgeBaseTextResponseBody {
	s.Code = &v
	return s
}

func (s *CreateKnowledgeBaseTextResponseBody) SetDirectoryId(v string) *CreateKnowledgeBaseTextResponseBody {
	s.DirectoryId = &v
	return s
}

func (s *CreateKnowledgeBaseTextResponseBody) SetGmtCreate(v string) *CreateKnowledgeBaseTextResponseBody {
	s.GmtCreate = &v
	return s
}

func (s *CreateKnowledgeBaseTextResponseBody) SetMessage(v string) *CreateKnowledgeBaseTextResponseBody {
	s.Message = &v
	return s
}

func (s *CreateKnowledgeBaseTextResponseBody) SetName(v string) *CreateKnowledgeBaseTextResponseBody {
	s.Name = &v
	return s
}

func (s *CreateKnowledgeBaseTextResponseBody) SetRequestId(v string) *CreateKnowledgeBaseTextResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateKnowledgeBaseTextResponseBody) SetScope(v string) *CreateKnowledgeBaseTextResponseBody {
	s.Scope = &v
	return s
}

func (s *CreateKnowledgeBaseTextResponseBody) SetSourceId(v string) *CreateKnowledgeBaseTextResponseBody {
	s.SourceId = &v
	return s
}

func (s *CreateKnowledgeBaseTextResponseBody) SetStatus(v string) *CreateKnowledgeBaseTextResponseBody {
	s.Status = &v
	return s
}

func (s *CreateKnowledgeBaseTextResponseBody) Validate() error {
	return dara.Validate(s)
}
