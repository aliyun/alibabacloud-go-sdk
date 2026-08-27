// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateKnowledgeBaseDirectoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateKnowledgeBaseDirectoryResponseBody
	GetCode() *string
	SetDirectoryId(v string) *UpdateKnowledgeBaseDirectoryResponseBody
	GetDirectoryId() *string
	SetMessage(v string) *UpdateKnowledgeBaseDirectoryResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateKnowledgeBaseDirectoryResponseBody
	GetRequestId() *string
}

type UpdateKnowledgeBaseDirectoryResponseBody struct {
	// The status code.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The category ID (echoed from the input parameter).
	//
	// example:
	//
	// exampleDirectoryId
	DirectoryId *string `json:"directoryId,omitempty" xml:"directoryId,omitempty"`
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
}

func (s UpdateKnowledgeBaseDirectoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateKnowledgeBaseDirectoryResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateKnowledgeBaseDirectoryResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateKnowledgeBaseDirectoryResponseBody) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *UpdateKnowledgeBaseDirectoryResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateKnowledgeBaseDirectoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateKnowledgeBaseDirectoryResponseBody) SetCode(v string) *UpdateKnowledgeBaseDirectoryResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateKnowledgeBaseDirectoryResponseBody) SetDirectoryId(v string) *UpdateKnowledgeBaseDirectoryResponseBody {
	s.DirectoryId = &v
	return s
}

func (s *UpdateKnowledgeBaseDirectoryResponseBody) SetMessage(v string) *UpdateKnowledgeBaseDirectoryResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateKnowledgeBaseDirectoryResponseBody) SetRequestId(v string) *UpdateKnowledgeBaseDirectoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateKnowledgeBaseDirectoryResponseBody) Validate() error {
	return dara.Validate(s)
}
