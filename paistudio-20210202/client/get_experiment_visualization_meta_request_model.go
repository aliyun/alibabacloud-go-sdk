// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetExperimentVisualizationMetaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNodeIds(v string) *GetExperimentVisualizationMetaRequest
	GetNodeIds() *string
}

type GetExperimentVisualizationMetaRequest struct {
	// example:
	//
	// node_id1,node_id2
	NodeIds *string `json:"NodeIds,omitempty" xml:"NodeIds,omitempty"`
}

func (s GetExperimentVisualizationMetaRequest) String() string {
	return dara.Prettify(s)
}

func (s GetExperimentVisualizationMetaRequest) GoString() string {
	return s.String()
}

func (s *GetExperimentVisualizationMetaRequest) GetNodeIds() *string {
	return s.NodeIds
}

func (s *GetExperimentVisualizationMetaRequest) SetNodeIds(v string) *GetExperimentVisualizationMetaRequest {
	s.NodeIds = &v
	return s
}

func (s *GetExperimentVisualizationMetaRequest) Validate() error {
	return dara.Validate(s)
}
