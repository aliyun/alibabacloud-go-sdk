// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateKnowledgeBaseSourceTagsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateKnowledgeBaseSourceTagsResponseBody
	GetCode() *string
	SetGmtModified(v string) *UpdateKnowledgeBaseSourceTagsResponseBody
	GetGmtModified() *string
	SetMessage(v string) *UpdateKnowledgeBaseSourceTagsResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateKnowledgeBaseSourceTagsResponseBody
	GetRequestId() *string
	SetSourceId(v string) *UpdateKnowledgeBaseSourceTagsResponseBody
	GetSourceId() *string
	SetSourceTags(v string) *UpdateKnowledgeBaseSourceTagsResponseBody
	GetSourceTags() *string
}

type UpdateKnowledgeBaseSourceTagsResponseBody struct {
	// The status code.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The last modified time in ISO 8601 format.
	//
	// example:
	//
	// string_value
	GmtModified *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// The description of the status code.
	//
	// example:
	//
	// ok
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The source ID.
	//
	// example:
	//
	// exampleSourceId
	SourceId *string `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	// The updated resource tags.
	//
	// example:
	//
	// string_value
	SourceTags *string `json:"sourceTags,omitempty" xml:"sourceTags,omitempty"`
}

func (s UpdateKnowledgeBaseSourceTagsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateKnowledgeBaseSourceTagsResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateKnowledgeBaseSourceTagsResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateKnowledgeBaseSourceTagsResponseBody) GetGmtModified() *string {
	return s.GmtModified
}

func (s *UpdateKnowledgeBaseSourceTagsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateKnowledgeBaseSourceTagsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateKnowledgeBaseSourceTagsResponseBody) GetSourceId() *string {
	return s.SourceId
}

func (s *UpdateKnowledgeBaseSourceTagsResponseBody) GetSourceTags() *string {
	return s.SourceTags
}

func (s *UpdateKnowledgeBaseSourceTagsResponseBody) SetCode(v string) *UpdateKnowledgeBaseSourceTagsResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateKnowledgeBaseSourceTagsResponseBody) SetGmtModified(v string) *UpdateKnowledgeBaseSourceTagsResponseBody {
	s.GmtModified = &v
	return s
}

func (s *UpdateKnowledgeBaseSourceTagsResponseBody) SetMessage(v string) *UpdateKnowledgeBaseSourceTagsResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateKnowledgeBaseSourceTagsResponseBody) SetRequestId(v string) *UpdateKnowledgeBaseSourceTagsResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateKnowledgeBaseSourceTagsResponseBody) SetSourceId(v string) *UpdateKnowledgeBaseSourceTagsResponseBody {
	s.SourceId = &v
	return s
}

func (s *UpdateKnowledgeBaseSourceTagsResponseBody) SetSourceTags(v string) *UpdateKnowledgeBaseSourceTagsResponseBody {
	s.SourceTags = &v
	return s
}

func (s *UpdateKnowledgeBaseSourceTagsResponseBody) Validate() error {
	return dara.Validate(s)
}
