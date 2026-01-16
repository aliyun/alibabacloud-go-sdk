// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteClusterNodesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *DeleteClusterNodesRequestBody) *DeleteClusterNodesRequest
	GetBody() *DeleteClusterNodesRequestBody
	SetClusterId(v string) *DeleteClusterNodesRequest
	GetClusterId() *string
}

type DeleteClusterNodesRequest struct {
	// This parameter is required.
	Body *DeleteClusterNodesRequestBody `json:"Body,omitempty" xml:"Body,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// eck-xxxxxxxx
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s DeleteClusterNodesRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteClusterNodesRequest) GoString() string {
	return s.String()
}

func (s *DeleteClusterNodesRequest) GetBody() *DeleteClusterNodesRequestBody {
	return s.Body
}

func (s *DeleteClusterNodesRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *DeleteClusterNodesRequest) SetBody(v *DeleteClusterNodesRequestBody) *DeleteClusterNodesRequest {
	s.Body = v
	return s
}

func (s *DeleteClusterNodesRequest) SetClusterId(v string) *DeleteClusterNodesRequest {
	s.ClusterId = &v
	return s
}

func (s *DeleteClusterNodesRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteClusterNodesRequestBody struct {
	// This parameter is required.
	Nodes []*string `json:"Nodes,omitempty" xml:"Nodes,omitempty" type:"Repeated"`
}

func (s DeleteClusterNodesRequestBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteClusterNodesRequestBody) GoString() string {
	return s.String()
}

func (s *DeleteClusterNodesRequestBody) GetNodes() []*string {
	return s.Nodes
}

func (s *DeleteClusterNodesRequestBody) SetNodes(v []*string) *DeleteClusterNodesRequestBody {
	s.Nodes = v
	return s
}

func (s *DeleteClusterNodesRequestBody) Validate() error {
	return dara.Validate(s)
}
