// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTraceInfoNodeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNode(v *DescribeTraceInfoNodeResponseBodyNode) *DescribeTraceInfoNodeResponseBody
	GetNode() *DescribeTraceInfoNodeResponseBodyNode
	SetRequestId(v string) *DescribeTraceInfoNodeResponseBody
	GetRequestId() *string
}

type DescribeTraceInfoNodeResponseBody struct {
	// The details about the node.
	Node *DescribeTraceInfoNodeResponseBodyNode `json:"Node,omitempty" xml:"Node,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// F35F45B0-5D6B-4238-BE02-A62DXXXXXXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeTraceInfoNodeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeTraceInfoNodeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTraceInfoNodeResponseBody) GetNode() *DescribeTraceInfoNodeResponseBodyNode {
	return s.Node
}

func (s *DescribeTraceInfoNodeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeTraceInfoNodeResponseBody) SetNode(v *DescribeTraceInfoNodeResponseBodyNode) *DescribeTraceInfoNodeResponseBody {
	s.Node = v
	return s
}

func (s *DescribeTraceInfoNodeResponseBody) SetRequestId(v string) *DescribeTraceInfoNodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTraceInfoNodeResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeTraceInfoNodeResponseBodyNode struct {
	// The name of the node.
	//
	// example:
	//
	// login
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// An array that consists of the properties of the node.
	PropertyList []*DescribeTraceInfoNodeResponseBodyNodePropertyList `json:"PropertyList,omitempty" xml:"PropertyList,omitempty" type:"Repeated"`
	// The type of the node.
	//
	// example:
	//
	// Alert
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeTraceInfoNodeResponseBodyNode) String() string {
	return dara.Prettify(s)
}

func (s DescribeTraceInfoNodeResponseBodyNode) GoString() string {
	return s.String()
}

func (s *DescribeTraceInfoNodeResponseBodyNode) GetName() *string {
	return s.Name
}

func (s *DescribeTraceInfoNodeResponseBodyNode) GetPropertyList() []*DescribeTraceInfoNodeResponseBodyNodePropertyList {
	return s.PropertyList
}

func (s *DescribeTraceInfoNodeResponseBodyNode) GetType() *string {
	return s.Type
}

func (s *DescribeTraceInfoNodeResponseBodyNode) SetName(v string) *DescribeTraceInfoNodeResponseBodyNode {
	s.Name = &v
	return s
}

func (s *DescribeTraceInfoNodeResponseBodyNode) SetPropertyList(v []*DescribeTraceInfoNodeResponseBodyNodePropertyList) *DescribeTraceInfoNodeResponseBodyNode {
	s.PropertyList = v
	return s
}

func (s *DescribeTraceInfoNodeResponseBodyNode) SetType(v string) *DescribeTraceInfoNodeResponseBodyNode {
	s.Type = &v
	return s
}

func (s *DescribeTraceInfoNodeResponseBodyNode) Validate() error {
	return dara.Validate(s)
}

type DescribeTraceInfoNodeResponseBodyNodePropertyList struct {
	// The name of the property.
	//
	// example:
	//
	// Incident
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The value of the property.
	//
	// example:
	//
	// Alert
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeTraceInfoNodeResponseBodyNodePropertyList) String() string {
	return dara.Prettify(s)
}

func (s DescribeTraceInfoNodeResponseBodyNodePropertyList) GoString() string {
	return s.String()
}

func (s *DescribeTraceInfoNodeResponseBodyNodePropertyList) GetName() *string {
	return s.Name
}

func (s *DescribeTraceInfoNodeResponseBodyNodePropertyList) GetValue() *string {
	return s.Value
}

func (s *DescribeTraceInfoNodeResponseBodyNodePropertyList) SetName(v string) *DescribeTraceInfoNodeResponseBodyNodePropertyList {
	s.Name = &v
	return s
}

func (s *DescribeTraceInfoNodeResponseBodyNodePropertyList) SetValue(v string) *DescribeTraceInfoNodeResponseBodyNodePropertyList {
	s.Value = &v
	return s
}

func (s *DescribeTraceInfoNodeResponseBodyNodePropertyList) Validate() error {
	return dara.Validate(s)
}
