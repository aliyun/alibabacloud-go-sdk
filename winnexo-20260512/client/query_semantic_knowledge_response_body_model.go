// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySemanticKnowledgeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QuerySemanticKnowledgeResponseBody
	GetCode() *string
	SetMessage(v string) *QuerySemanticKnowledgeResponseBody
	GetMessage() *string
	SetRequestId(v string) *QuerySemanticKnowledgeResponseBody
	GetRequestId() *string
	SetSchemaKnowledge(v string) *QuerySemanticKnowledgeResponseBody
	GetSchemaKnowledge() *string
}

type QuerySemanticKnowledgeResponseBody struct {
	// The status code.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The status code description.
	//
	// example:
	//
	// successful
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The semantic knowledge text related to the query.
	//
	// This parameter is required.
	//
	// example:
	//
	// string_value
	SchemaKnowledge *string `json:"schemaKnowledge,omitempty" xml:"schemaKnowledge,omitempty"`
}

func (s QuerySemanticKnowledgeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QuerySemanticKnowledgeResponseBody) GoString() string {
	return s.String()
}

func (s *QuerySemanticKnowledgeResponseBody) GetCode() *string {
	return s.Code
}

func (s *QuerySemanticKnowledgeResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QuerySemanticKnowledgeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QuerySemanticKnowledgeResponseBody) GetSchemaKnowledge() *string {
	return s.SchemaKnowledge
}

func (s *QuerySemanticKnowledgeResponseBody) SetCode(v string) *QuerySemanticKnowledgeResponseBody {
	s.Code = &v
	return s
}

func (s *QuerySemanticKnowledgeResponseBody) SetMessage(v string) *QuerySemanticKnowledgeResponseBody {
	s.Message = &v
	return s
}

func (s *QuerySemanticKnowledgeResponseBody) SetRequestId(v string) *QuerySemanticKnowledgeResponseBody {
	s.RequestId = &v
	return s
}

func (s *QuerySemanticKnowledgeResponseBody) SetSchemaKnowledge(v string) *QuerySemanticKnowledgeResponseBody {
	s.SchemaKnowledge = &v
	return s
}

func (s *QuerySemanticKnowledgeResponseBody) Validate() error {
	return dara.Validate(s)
}
