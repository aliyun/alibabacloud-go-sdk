// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAINodesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *CreateAINodesRequest
	GetDBClusterId() *string
	SetDBNodes(v []*CreateAINodesRequestDBNodes) *CreateAINodesRequest
	GetDBNodes() []*CreateAINodesRequestDBNodes
}

type CreateAINodesRequest struct {
	// example:
	//
	// pm-xxxxxx
	DBClusterId *string                        `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	DBNodes     []*CreateAINodesRequestDBNodes `json:"DBNodes,omitempty" xml:"DBNodes,omitempty" type:"Repeated"`
}

func (s CreateAINodesRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAINodesRequest) GoString() string {
	return s.String()
}

func (s *CreateAINodesRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *CreateAINodesRequest) GetDBNodes() []*CreateAINodesRequestDBNodes {
	return s.DBNodes
}

func (s *CreateAINodesRequest) SetDBClusterId(v string) *CreateAINodesRequest {
	s.DBClusterId = &v
	return s
}

func (s *CreateAINodesRequest) SetDBNodes(v []*CreateAINodesRequestDBNodes) *CreateAINodesRequest {
	s.DBNodes = v
	return s
}

func (s *CreateAINodesRequest) Validate() error {
	if s.DBNodes != nil {
		for _, item := range s.DBNodes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateAINodesRequestDBNodes struct {
	// example:
	//
	// polar.mysql.g4.xlarge.gu10
	DBNodeClass *string `json:"DBNodeClass,omitempty" xml:"DBNodeClass,omitempty"`
}

func (s CreateAINodesRequestDBNodes) String() string {
	return dara.Prettify(s)
}

func (s CreateAINodesRequestDBNodes) GoString() string {
	return s.String()
}

func (s *CreateAINodesRequestDBNodes) GetDBNodeClass() *string {
	return s.DBNodeClass
}

func (s *CreateAINodesRequestDBNodes) SetDBNodeClass(v string) *CreateAINodesRequestDBNodes {
	s.DBNodeClass = &v
	return s
}

func (s *CreateAINodesRequestDBNodes) Validate() error {
	return dara.Validate(s)
}
