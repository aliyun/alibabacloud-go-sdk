// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDBNodesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *CreateDBNodesResponseBody
	GetDBClusterId() *string
	SetDBNodeIds(v *CreateDBNodesResponseBodyDBNodeIds) *CreateDBNodesResponseBody
	GetDBNodeIds() *CreateDBNodesResponseBodyDBNodeIds
	SetOrderId(v string) *CreateDBNodesResponseBody
	GetOrderId() *string
	SetRequestId(v string) *CreateDBNodesResponseBody
	GetRequestId() *string
}

type CreateDBNodesResponseBody struct {
	// The ID of the cluster.
	//
	// example:
	//
	// pc-***************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// Details about the nodes.
	DBNodeIds *CreateDBNodesResponseBodyDBNodeIds `json:"DBNodeIds,omitempty" xml:"DBNodeIds,omitempty" type:"Struct"`
	// The ID of the order.
	//
	// example:
	//
	// 2148126708*****
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// E053E730-5755-4AA9-AA9D-A47867******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDBNodesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDBNodesResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDBNodesResponseBody) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *CreateDBNodesResponseBody) GetDBNodeIds() *CreateDBNodesResponseBodyDBNodeIds {
	return s.DBNodeIds
}

func (s *CreateDBNodesResponseBody) GetOrderId() *string {
	return s.OrderId
}

func (s *CreateDBNodesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDBNodesResponseBody) SetDBClusterId(v string) *CreateDBNodesResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *CreateDBNodesResponseBody) SetDBNodeIds(v *CreateDBNodesResponseBodyDBNodeIds) *CreateDBNodesResponseBody {
	s.DBNodeIds = v
	return s
}

func (s *CreateDBNodesResponseBody) SetOrderId(v string) *CreateDBNodesResponseBody {
	s.OrderId = &v
	return s
}

func (s *CreateDBNodesResponseBody) SetRequestId(v string) *CreateDBNodesResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDBNodesResponseBody) Validate() error {
	return dara.Validate(s)
}

type CreateDBNodesResponseBodyDBNodeIds struct {
	DBNodeId []*string `json:"DBNodeId,omitempty" xml:"DBNodeId,omitempty" type:"Repeated"`
}

func (s CreateDBNodesResponseBodyDBNodeIds) String() string {
	return dara.Prettify(s)
}

func (s CreateDBNodesResponseBodyDBNodeIds) GoString() string {
	return s.String()
}

func (s *CreateDBNodesResponseBodyDBNodeIds) GetDBNodeId() []*string {
	return s.DBNodeId
}

func (s *CreateDBNodesResponseBodyDBNodeIds) SetDBNodeId(v []*string) *CreateDBNodesResponseBodyDBNodeIds {
	s.DBNodeId = v
	return s
}

func (s *CreateDBNodesResponseBodyDBNodeIds) Validate() error {
	return dara.Validate(s)
}
