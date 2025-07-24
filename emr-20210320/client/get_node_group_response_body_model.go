// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNodeGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNodeGroup(v *NodeGroup) *GetNodeGroupResponseBody
	GetNodeGroup() *NodeGroup
	SetRequestId(v string) *GetNodeGroupResponseBody
	GetRequestId() *string
}

type GetNodeGroupResponseBody struct {
	// The node group.
	NodeGroup *NodeGroup `json:"NodeGroup,omitempty" xml:"NodeGroup,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetNodeGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetNodeGroupResponseBody) GoString() string {
	return s.String()
}

func (s *GetNodeGroupResponseBody) GetNodeGroup() *NodeGroup {
	return s.NodeGroup
}

func (s *GetNodeGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetNodeGroupResponseBody) SetNodeGroup(v *NodeGroup) *GetNodeGroupResponseBody {
	s.NodeGroup = v
	return s
}

func (s *GetNodeGroupResponseBody) SetRequestId(v string) *GetNodeGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetNodeGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
