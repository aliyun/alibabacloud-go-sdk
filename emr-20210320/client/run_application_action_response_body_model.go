// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunApplicationActionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAbnInstances(v []*RunApplicationActionResponseBodyAbnInstances) *RunApplicationActionResponseBody
	GetAbnInstances() []*RunApplicationActionResponseBodyAbnInstances
	SetOperationId(v string) *RunApplicationActionResponseBody
	GetOperationId() *string
	SetRequestId(v string) *RunApplicationActionResponseBody
	GetRequestId() *string
}

type RunApplicationActionResponseBody struct {
	// The abnormal nodes.
	AbnInstances []*RunApplicationActionResponseBodyAbnInstances `json:"AbnInstances,omitempty" xml:"AbnInstances,omitempty" type:"Repeated"`
	// The operation ID.
	//
	// example:
	//
	// op-13c37a77c505****
	OperationId *string `json:"OperationId,omitempty" xml:"OperationId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9E3A7161-EB7B-172B-8D18-FFB06BA3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RunApplicationActionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RunApplicationActionResponseBody) GoString() string {
	return s.String()
}

func (s *RunApplicationActionResponseBody) GetAbnInstances() []*RunApplicationActionResponseBodyAbnInstances {
	return s.AbnInstances
}

func (s *RunApplicationActionResponseBody) GetOperationId() *string {
	return s.OperationId
}

func (s *RunApplicationActionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RunApplicationActionResponseBody) SetAbnInstances(v []*RunApplicationActionResponseBodyAbnInstances) *RunApplicationActionResponseBody {
	s.AbnInstances = v
	return s
}

func (s *RunApplicationActionResponseBody) SetOperationId(v string) *RunApplicationActionResponseBody {
	s.OperationId = &v
	return s
}

func (s *RunApplicationActionResponseBody) SetRequestId(v string) *RunApplicationActionResponseBody {
	s.RequestId = &v
	return s
}

func (s *RunApplicationActionResponseBody) Validate() error {
	return dara.Validate(s)
}

type RunApplicationActionResponseBodyAbnInstances struct {
	// The ID of the abnormal node.
	//
	// example:
	//
	// i-bp1cudc25w2bfwl5****
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The name of the abnormal node.
	//
	// example:
	//
	// core1-1
	NodeName *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
}

func (s RunApplicationActionResponseBodyAbnInstances) String() string {
	return dara.Prettify(s)
}

func (s RunApplicationActionResponseBodyAbnInstances) GoString() string {
	return s.String()
}

func (s *RunApplicationActionResponseBodyAbnInstances) GetNodeId() *string {
	return s.NodeId
}

func (s *RunApplicationActionResponseBodyAbnInstances) GetNodeName() *string {
	return s.NodeName
}

func (s *RunApplicationActionResponseBodyAbnInstances) SetNodeId(v string) *RunApplicationActionResponseBodyAbnInstances {
	s.NodeId = &v
	return s
}

func (s *RunApplicationActionResponseBodyAbnInstances) SetNodeName(v string) *RunApplicationActionResponseBodyAbnInstances {
	s.NodeName = &v
	return s
}

func (s *RunApplicationActionResponseBodyAbnInstances) Validate() error {
	return dara.Validate(s)
}
