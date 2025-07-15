// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCloudPhoneNodeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNewNodeName(v string) *ModifyCloudPhoneNodeRequest
	GetNewNodeName() *string
	SetNodeId(v string) *ModifyCloudPhoneNodeRequest
	GetNodeId() *string
	SetStreamMode(v int32) *ModifyCloudPhoneNodeRequest
	GetStreamMode() *int32
}

type ModifyCloudPhoneNodeRequest struct {
	// The name that you want to assign to the cloud phone matrix.
	//
	// example:
	//
	// node_name_new
	NewNodeName *string `json:"NewNodeName,omitempty" xml:"NewNodeName,omitempty"`
	// The ID of the cloud phone matrix.
	//
	// example:
	//
	// cpn-0ugbptfu473fy****
	NodeId     *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	StreamMode *int32  `json:"StreamMode,omitempty" xml:"StreamMode,omitempty"`
}

func (s ModifyCloudPhoneNodeRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyCloudPhoneNodeRequest) GoString() string {
	return s.String()
}

func (s *ModifyCloudPhoneNodeRequest) GetNewNodeName() *string {
	return s.NewNodeName
}

func (s *ModifyCloudPhoneNodeRequest) GetNodeId() *string {
	return s.NodeId
}

func (s *ModifyCloudPhoneNodeRequest) GetStreamMode() *int32 {
	return s.StreamMode
}

func (s *ModifyCloudPhoneNodeRequest) SetNewNodeName(v string) *ModifyCloudPhoneNodeRequest {
	s.NewNodeName = &v
	return s
}

func (s *ModifyCloudPhoneNodeRequest) SetNodeId(v string) *ModifyCloudPhoneNodeRequest {
	s.NodeId = &v
	return s
}

func (s *ModifyCloudPhoneNodeRequest) SetStreamMode(v int32) *ModifyCloudPhoneNodeRequest {
	s.StreamMode = &v
	return s
}

func (s *ModifyCloudPhoneNodeRequest) Validate() error {
	return dara.Validate(s)
}
