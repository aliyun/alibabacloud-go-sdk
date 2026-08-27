// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateKnowledgeBaseFileResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateKnowledgeBaseFileResponseBody
	GetCode() *string
	SetDirectoryId(v string) *CreateKnowledgeBaseFileResponseBody
	GetDirectoryId() *string
	SetGmtCreate(v string) *CreateKnowledgeBaseFileResponseBody
	GetGmtCreate() *string
	SetMessage(v string) *CreateKnowledgeBaseFileResponseBody
	GetMessage() *string
	SetName(v string) *CreateKnowledgeBaseFileResponseBody
	GetName() *string
	SetRequestId(v string) *CreateKnowledgeBaseFileResponseBody
	GetRequestId() *string
	SetScope(v string) *CreateKnowledgeBaseFileResponseBody
	GetScope() *string
	SetSourceId(v string) *CreateKnowledgeBaseFileResponseBody
	GetSourceId() *string
	SetStatus(v string) *CreateKnowledgeBaseFileResponseBody
	GetStatus() *string
}

type CreateKnowledgeBaseFileResponseBody struct {
	// The response status code.
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
	// 2025-11-12T03:08:56Z
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// The response message.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The image name.
	//
	// example:
	//
	// oklabs_tongyici
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 911656E1-9A09-5C77-BAAD-915EB4958D68
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The export scope. Valid values:
	//
	// - ALL: all.
	//
	// - SELECT: selected rows.
	//
	// example:
	//
	// user_info projects pull_requests hook gists emails
	Scope *string `json:"scope,omitempty" xml:"scope,omitempty"`
	// The unique identifier on the business system side, which is the business ID.
	//
	// example:
	//
	// 2001549
	SourceId *string `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	// The store status.
	//
	// example:
	//
	// 200
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s CreateKnowledgeBaseFileResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateKnowledgeBaseFileResponseBody) GoString() string {
	return s.String()
}

func (s *CreateKnowledgeBaseFileResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateKnowledgeBaseFileResponseBody) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *CreateKnowledgeBaseFileResponseBody) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *CreateKnowledgeBaseFileResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateKnowledgeBaseFileResponseBody) GetName() *string {
	return s.Name
}

func (s *CreateKnowledgeBaseFileResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateKnowledgeBaseFileResponseBody) GetScope() *string {
	return s.Scope
}

func (s *CreateKnowledgeBaseFileResponseBody) GetSourceId() *string {
	return s.SourceId
}

func (s *CreateKnowledgeBaseFileResponseBody) GetStatus() *string {
	return s.Status
}

func (s *CreateKnowledgeBaseFileResponseBody) SetCode(v string) *CreateKnowledgeBaseFileResponseBody {
	s.Code = &v
	return s
}

func (s *CreateKnowledgeBaseFileResponseBody) SetDirectoryId(v string) *CreateKnowledgeBaseFileResponseBody {
	s.DirectoryId = &v
	return s
}

func (s *CreateKnowledgeBaseFileResponseBody) SetGmtCreate(v string) *CreateKnowledgeBaseFileResponseBody {
	s.GmtCreate = &v
	return s
}

func (s *CreateKnowledgeBaseFileResponseBody) SetMessage(v string) *CreateKnowledgeBaseFileResponseBody {
	s.Message = &v
	return s
}

func (s *CreateKnowledgeBaseFileResponseBody) SetName(v string) *CreateKnowledgeBaseFileResponseBody {
	s.Name = &v
	return s
}

func (s *CreateKnowledgeBaseFileResponseBody) SetRequestId(v string) *CreateKnowledgeBaseFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateKnowledgeBaseFileResponseBody) SetScope(v string) *CreateKnowledgeBaseFileResponseBody {
	s.Scope = &v
	return s
}

func (s *CreateKnowledgeBaseFileResponseBody) SetSourceId(v string) *CreateKnowledgeBaseFileResponseBody {
	s.SourceId = &v
	return s
}

func (s *CreateKnowledgeBaseFileResponseBody) SetStatus(v string) *CreateKnowledgeBaseFileResponseBody {
	s.Status = &v
	return s
}

func (s *CreateKnowledgeBaseFileResponseBody) Validate() error {
	return dara.Validate(s)
}
