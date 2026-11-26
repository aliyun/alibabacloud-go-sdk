// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReplacePersonalSourceObjectBindingsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ReplacePersonalSourceObjectBindingsResponseBody
	GetCode() *string
	SetMessage(v string) *ReplacePersonalSourceObjectBindingsResponseBody
	GetMessage() *string
	SetObjectBindings(v []*ReplacePersonalSourceObjectBindingsResponseBodyObjectBindings) *ReplacePersonalSourceObjectBindingsResponseBody
	GetObjectBindings() []*ReplacePersonalSourceObjectBindingsResponseBodyObjectBindings
	SetRequestId(v string) *ReplacePersonalSourceObjectBindingsResponseBody
	GetRequestId() *string
	SetSourceId(v string) *ReplacePersonalSourceObjectBindingsResponseBody
	GetSourceId() *string
}

type ReplacePersonalSourceObjectBindingsResponseBody struct {
	// The business status code. A value of 200 indicates success.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The description of the status code.
	//
	// example:
	//
	// ok
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The object bindings.
	ObjectBindings []*ReplacePersonalSourceObjectBindingsResponseBodyObjectBindings `json:"objectBindings,omitempty" xml:"objectBindings,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The ID of the data source.
	//
	// example:
	//
	// exampleSourceId
	SourceId *string `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
}

func (s ReplacePersonalSourceObjectBindingsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ReplacePersonalSourceObjectBindingsResponseBody) GoString() string {
	return s.String()
}

func (s *ReplacePersonalSourceObjectBindingsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ReplacePersonalSourceObjectBindingsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ReplacePersonalSourceObjectBindingsResponseBody) GetObjectBindings() []*ReplacePersonalSourceObjectBindingsResponseBodyObjectBindings {
	return s.ObjectBindings
}

func (s *ReplacePersonalSourceObjectBindingsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ReplacePersonalSourceObjectBindingsResponseBody) GetSourceId() *string {
	return s.SourceId
}

func (s *ReplacePersonalSourceObjectBindingsResponseBody) SetCode(v string) *ReplacePersonalSourceObjectBindingsResponseBody {
	s.Code = &v
	return s
}

func (s *ReplacePersonalSourceObjectBindingsResponseBody) SetMessage(v string) *ReplacePersonalSourceObjectBindingsResponseBody {
	s.Message = &v
	return s
}

func (s *ReplacePersonalSourceObjectBindingsResponseBody) SetObjectBindings(v []*ReplacePersonalSourceObjectBindingsResponseBodyObjectBindings) *ReplacePersonalSourceObjectBindingsResponseBody {
	s.ObjectBindings = v
	return s
}

func (s *ReplacePersonalSourceObjectBindingsResponseBody) SetRequestId(v string) *ReplacePersonalSourceObjectBindingsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReplacePersonalSourceObjectBindingsResponseBody) SetSourceId(v string) *ReplacePersonalSourceObjectBindingsResponseBody {
	s.SourceId = &v
	return s
}

func (s *ReplacePersonalSourceObjectBindingsResponseBody) Validate() error {
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

type ReplacePersonalSourceObjectBindingsResponseBodyObjectBindings struct {
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

func (s ReplacePersonalSourceObjectBindingsResponseBodyObjectBindings) String() string {
	return dara.Prettify(s)
}

func (s ReplacePersonalSourceObjectBindingsResponseBodyObjectBindings) GoString() string {
	return s.String()
}

func (s *ReplacePersonalSourceObjectBindingsResponseBodyObjectBindings) GetGraphName() *string {
	return s.GraphName
}

func (s *ReplacePersonalSourceObjectBindingsResponseBodyObjectBindings) GetObjectId() *string {
	return s.ObjectId
}

func (s *ReplacePersonalSourceObjectBindingsResponseBodyObjectBindings) GetObjectType() *string {
	return s.ObjectType
}

func (s *ReplacePersonalSourceObjectBindingsResponseBodyObjectBindings) SetGraphName(v string) *ReplacePersonalSourceObjectBindingsResponseBodyObjectBindings {
	s.GraphName = &v
	return s
}

func (s *ReplacePersonalSourceObjectBindingsResponseBodyObjectBindings) SetObjectId(v string) *ReplacePersonalSourceObjectBindingsResponseBodyObjectBindings {
	s.ObjectId = &v
	return s
}

func (s *ReplacePersonalSourceObjectBindingsResponseBodyObjectBindings) SetObjectType(v string) *ReplacePersonalSourceObjectBindingsResponseBodyObjectBindings {
	s.ObjectType = &v
	return s
}

func (s *ReplacePersonalSourceObjectBindingsResponseBodyObjectBindings) Validate() error {
	return dara.Validate(s)
}
