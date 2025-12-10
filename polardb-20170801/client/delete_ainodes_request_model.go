// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAINodesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DeleteAINodesRequest
	GetDBClusterId() *string
	SetDBNodeId(v []*string) *DeleteAINodesRequest
	GetDBNodeId() []*string
}

type DeleteAINodesRequest struct {
	// example:
	//
	// pm-xxxxxx
	DBClusterId *string   `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	DBNodeId    []*string `json:"DBNodeId,omitempty" xml:"DBNodeId,omitempty" type:"Repeated"`
}

func (s DeleteAINodesRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAINodesRequest) GoString() string {
	return s.String()
}

func (s *DeleteAINodesRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DeleteAINodesRequest) GetDBNodeId() []*string {
	return s.DBNodeId
}

func (s *DeleteAINodesRequest) SetDBClusterId(v string) *DeleteAINodesRequest {
	s.DBClusterId = &v
	return s
}

func (s *DeleteAINodesRequest) SetDBNodeId(v []*string) *DeleteAINodesRequest {
	s.DBNodeId = v
	return s
}

func (s *DeleteAINodesRequest) Validate() error {
	return dara.Validate(s)
}
