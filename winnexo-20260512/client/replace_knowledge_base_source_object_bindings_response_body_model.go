// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReplaceKnowledgeBaseSourceObjectBindingsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ReplaceKnowledgeBaseSourceObjectBindingsResponseBody
	GetCode() *string
	SetMessage(v string) *ReplaceKnowledgeBaseSourceObjectBindingsResponseBody
	GetMessage() *string
	SetObjectBindings(v []*ReplaceKnowledgeBaseSourceObjectBindingsResponseBodyObjectBindings) *ReplaceKnowledgeBaseSourceObjectBindingsResponseBody
	GetObjectBindings() []*ReplaceKnowledgeBaseSourceObjectBindingsResponseBodyObjectBindings
	SetRequestId(v string) *ReplaceKnowledgeBaseSourceObjectBindingsResponseBody
	GetRequestId() *string
	SetSourceId(v string) *ReplaceKnowledgeBaseSourceObjectBindingsResponseBody
	GetSourceId() *string
}

type ReplaceKnowledgeBaseSourceObjectBindingsResponseBody struct {
	// The error code.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The status code description.
	//
	// example:
	//
	// ok
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The object bindings.
	ObjectBindings []*ReplaceKnowledgeBaseSourceObjectBindingsResponseBodyObjectBindings `json:"objectBindings,omitempty" xml:"objectBindings,omitempty" type:"Repeated"`
	// The request tracking ID.
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The ID of the personal FILE data source to be replaced. The ID is unique within the tenant.
	//
	// example:
	//
	// exampleSourceId
	SourceId *string `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
}

func (s ReplaceKnowledgeBaseSourceObjectBindingsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ReplaceKnowledgeBaseSourceObjectBindingsResponseBody) GoString() string {
	return s.String()
}

func (s *ReplaceKnowledgeBaseSourceObjectBindingsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ReplaceKnowledgeBaseSourceObjectBindingsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ReplaceKnowledgeBaseSourceObjectBindingsResponseBody) GetObjectBindings() []*ReplaceKnowledgeBaseSourceObjectBindingsResponseBodyObjectBindings {
	return s.ObjectBindings
}

func (s *ReplaceKnowledgeBaseSourceObjectBindingsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ReplaceKnowledgeBaseSourceObjectBindingsResponseBody) GetSourceId() *string {
	return s.SourceId
}

func (s *ReplaceKnowledgeBaseSourceObjectBindingsResponseBody) SetCode(v string) *ReplaceKnowledgeBaseSourceObjectBindingsResponseBody {
	s.Code = &v
	return s
}

func (s *ReplaceKnowledgeBaseSourceObjectBindingsResponseBody) SetMessage(v string) *ReplaceKnowledgeBaseSourceObjectBindingsResponseBody {
	s.Message = &v
	return s
}

func (s *ReplaceKnowledgeBaseSourceObjectBindingsResponseBody) SetObjectBindings(v []*ReplaceKnowledgeBaseSourceObjectBindingsResponseBodyObjectBindings) *ReplaceKnowledgeBaseSourceObjectBindingsResponseBody {
	s.ObjectBindings = v
	return s
}

func (s *ReplaceKnowledgeBaseSourceObjectBindingsResponseBody) SetRequestId(v string) *ReplaceKnowledgeBaseSourceObjectBindingsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReplaceKnowledgeBaseSourceObjectBindingsResponseBody) SetSourceId(v string) *ReplaceKnowledgeBaseSourceObjectBindingsResponseBody {
	s.SourceId = &v
	return s
}

func (s *ReplaceKnowledgeBaseSourceObjectBindingsResponseBody) Validate() error {
	if s.ObjectBindings != nil {
		for _, item := range s.ObjectBindings {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ReplaceKnowledgeBaseSourceObjectBindingsResponseBodyObjectBindings struct {
	// The name of the semantic graph to which the bound object belongs. The object ID is unique within this graph. This parameter is required.
	//
	// example:
	//
	// string_value
	GraphName *string `json:"graphName,omitempty" xml:"graphName,omitempty"`
	// The ID of the bound object.
	//
	// example:
	//
	// exampleObjectId
	ObjectId *string `json:"objectId,omitempty" xml:"objectId,omitempty"`
	// The object type to attach, such as customer or project.
	//
	// example:
	//
	// string_value
	ObjectType *string `json:"objectType,omitempty" xml:"objectType,omitempty"`
}

func (s ReplaceKnowledgeBaseSourceObjectBindingsResponseBodyObjectBindings) String() string {
	return dara.Prettify(s)
}

func (s ReplaceKnowledgeBaseSourceObjectBindingsResponseBodyObjectBindings) GoString() string {
	return s.String()
}

func (s *ReplaceKnowledgeBaseSourceObjectBindingsResponseBodyObjectBindings) GetGraphName() *string {
	return s.GraphName
}

func (s *ReplaceKnowledgeBaseSourceObjectBindingsResponseBodyObjectBindings) GetObjectId() *string {
	return s.ObjectId
}

func (s *ReplaceKnowledgeBaseSourceObjectBindingsResponseBodyObjectBindings) GetObjectType() *string {
	return s.ObjectType
}

func (s *ReplaceKnowledgeBaseSourceObjectBindingsResponseBodyObjectBindings) SetGraphName(v string) *ReplaceKnowledgeBaseSourceObjectBindingsResponseBodyObjectBindings {
	s.GraphName = &v
	return s
}

func (s *ReplaceKnowledgeBaseSourceObjectBindingsResponseBodyObjectBindings) SetObjectId(v string) *ReplaceKnowledgeBaseSourceObjectBindingsResponseBodyObjectBindings {
	s.ObjectId = &v
	return s
}

func (s *ReplaceKnowledgeBaseSourceObjectBindingsResponseBodyObjectBindings) SetObjectType(v string) *ReplaceKnowledgeBaseSourceObjectBindingsResponseBodyObjectBindings {
	s.ObjectType = &v
	return s
}

func (s *ReplaceKnowledgeBaseSourceObjectBindingsResponseBodyObjectBindings) Validate() error {
	return dara.Validate(s)
}
