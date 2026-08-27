// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMoveKnowledgeBaseResourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *MoveKnowledgeBaseResourceResponseBody
	GetCode() *string
	SetKnowledgeId(v string) *MoveKnowledgeBaseResourceResponseBody
	GetKnowledgeId() *string
	SetMessage(v string) *MoveKnowledgeBaseResourceResponseBody
	GetMessage() *string
	SetRequestId(v string) *MoveKnowledgeBaseResourceResponseBody
	GetRequestId() *string
	SetSourceDirectoryId(v string) *MoveKnowledgeBaseResourceResponseBody
	GetSourceDirectoryId() *string
	SetSourceId(v string) *MoveKnowledgeBaseResourceResponseBody
	GetSourceId() *string
	SetTargetDirectoryId(v string) *MoveKnowledgeBaseResourceResponseBody
	GetTargetDirectoryId() *string
}

type MoveKnowledgeBaseResourceResponseBody struct {
	// The response code.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The target knowledge base ID. This value is echoed from the request parameter.
	//
	// example:
	//
	// exampleKnowledgeId
	KnowledgeId *string `json:"knowledgeId,omitempty" xml:"knowledgeId,omitempty"`
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
	// 375701FC-2FB9-5782-BE8F-A3F5E2F2158C
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The source directory ID. This value is echoed from the request parameter.
	//
	// example:
	//
	// exampleSourceDirectoryId
	SourceDirectoryId *string `json:"sourceDirectoryId,omitempty" xml:"sourceDirectoryId,omitempty"`
	// The unique identifier on the business system side, that is, the business ID.
	//
	// example:
	//
	// 2000358
	SourceId *string `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	// The target directory ID. This value is echoed from the request parameter.
	//
	// example:
	//
	// exampleTargetDirectoryId
	TargetDirectoryId *string `json:"targetDirectoryId,omitempty" xml:"targetDirectoryId,omitempty"`
}

func (s MoveKnowledgeBaseResourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s MoveKnowledgeBaseResourceResponseBody) GoString() string {
	return s.String()
}

func (s *MoveKnowledgeBaseResourceResponseBody) GetCode() *string {
	return s.Code
}

func (s *MoveKnowledgeBaseResourceResponseBody) GetKnowledgeId() *string {
	return s.KnowledgeId
}

func (s *MoveKnowledgeBaseResourceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *MoveKnowledgeBaseResourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *MoveKnowledgeBaseResourceResponseBody) GetSourceDirectoryId() *string {
	return s.SourceDirectoryId
}

func (s *MoveKnowledgeBaseResourceResponseBody) GetSourceId() *string {
	return s.SourceId
}

func (s *MoveKnowledgeBaseResourceResponseBody) GetTargetDirectoryId() *string {
	return s.TargetDirectoryId
}

func (s *MoveKnowledgeBaseResourceResponseBody) SetCode(v string) *MoveKnowledgeBaseResourceResponseBody {
	s.Code = &v
	return s
}

func (s *MoveKnowledgeBaseResourceResponseBody) SetKnowledgeId(v string) *MoveKnowledgeBaseResourceResponseBody {
	s.KnowledgeId = &v
	return s
}

func (s *MoveKnowledgeBaseResourceResponseBody) SetMessage(v string) *MoveKnowledgeBaseResourceResponseBody {
	s.Message = &v
	return s
}

func (s *MoveKnowledgeBaseResourceResponseBody) SetRequestId(v string) *MoveKnowledgeBaseResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *MoveKnowledgeBaseResourceResponseBody) SetSourceDirectoryId(v string) *MoveKnowledgeBaseResourceResponseBody {
	s.SourceDirectoryId = &v
	return s
}

func (s *MoveKnowledgeBaseResourceResponseBody) SetSourceId(v string) *MoveKnowledgeBaseResourceResponseBody {
	s.SourceId = &v
	return s
}

func (s *MoveKnowledgeBaseResourceResponseBody) SetTargetDirectoryId(v string) *MoveKnowledgeBaseResourceResponseBody {
	s.TargetDirectoryId = &v
	return s
}

func (s *MoveKnowledgeBaseResourceResponseBody) Validate() error {
	return dara.Validate(s)
}
