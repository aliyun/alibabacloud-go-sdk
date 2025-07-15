// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCloudPhoneNodesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNodeIds(v []*string) *DeleteCloudPhoneNodesRequest
	GetNodeIds() []*string
}

type DeleteCloudPhoneNodesRequest struct {
	// The cloud phone matrix IDs.
	NodeIds []*string `json:"NodeIds,omitempty" xml:"NodeIds,omitempty" type:"Repeated"`
}

func (s DeleteCloudPhoneNodesRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteCloudPhoneNodesRequest) GoString() string {
	return s.String()
}

func (s *DeleteCloudPhoneNodesRequest) GetNodeIds() []*string {
	return s.NodeIds
}

func (s *DeleteCloudPhoneNodesRequest) SetNodeIds(v []*string) *DeleteCloudPhoneNodesRequest {
	s.NodeIds = v
	return s
}

func (s *DeleteCloudPhoneNodesRequest) Validate() error {
	return dara.Validate(s)
}
