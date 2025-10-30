// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddAINodeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAINodePoolId(v string) *AddAINodeRequest
	GetAINodePoolId() *string
	SetAINodeSpecInfos(v []*AddAINodeRequestAINodeSpecInfos) *AddAINodeRequest
	GetAINodeSpecInfos() []*AddAINodeRequestAINodeSpecInfos
	SetDBInstanceId(v string) *AddAINodeRequest
	GetDBInstanceId() *string
}

type AddAINodeRequest struct {
	// example:
	//
	// aipool-xxxxx
	AINodePoolId *string `json:"AINodePoolId,omitempty" xml:"AINodePoolId,omitempty"`
	// This parameter is required.
	AINodeSpecInfos []*AddAINodeRequestAINodeSpecInfos `json:"AINodeSpecInfos,omitempty" xml:"AINodeSpecInfos,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// gp-xxxxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
}

func (s AddAINodeRequest) String() string {
	return dara.Prettify(s)
}

func (s AddAINodeRequest) GoString() string {
	return s.String()
}

func (s *AddAINodeRequest) GetAINodePoolId() *string {
	return s.AINodePoolId
}

func (s *AddAINodeRequest) GetAINodeSpecInfos() []*AddAINodeRequestAINodeSpecInfos {
	return s.AINodeSpecInfos
}

func (s *AddAINodeRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *AddAINodeRequest) SetAINodePoolId(v string) *AddAINodeRequest {
	s.AINodePoolId = &v
	return s
}

func (s *AddAINodeRequest) SetAINodeSpecInfos(v []*AddAINodeRequestAINodeSpecInfos) *AddAINodeRequest {
	s.AINodeSpecInfos = v
	return s
}

func (s *AddAINodeRequest) SetDBInstanceId(v string) *AddAINodeRequest {
	s.DBInstanceId = &v
	return s
}

func (s *AddAINodeRequest) Validate() error {
	if s.AINodeSpecInfos != nil {
		for _, item := range s.AINodeSpecInfos {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type AddAINodeRequestAINodeSpecInfos struct {
	// This parameter is required.
	//
	// example:
	//
	// 2
	NodeNum *string `json:"NodeNum,omitempty" xml:"NodeNum,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ADB.AIStandard.1
	NodeSpec *string `json:"NodeSpec,omitempty" xml:"NodeSpec,omitempty"`
}

func (s AddAINodeRequestAINodeSpecInfos) String() string {
	return dara.Prettify(s)
}

func (s AddAINodeRequestAINodeSpecInfos) GoString() string {
	return s.String()
}

func (s *AddAINodeRequestAINodeSpecInfos) GetNodeNum() *string {
	return s.NodeNum
}

func (s *AddAINodeRequestAINodeSpecInfos) GetNodeSpec() *string {
	return s.NodeSpec
}

func (s *AddAINodeRequestAINodeSpecInfos) SetNodeNum(v string) *AddAINodeRequestAINodeSpecInfos {
	s.NodeNum = &v
	return s
}

func (s *AddAINodeRequestAINodeSpecInfos) SetNodeSpec(v string) *AddAINodeRequestAINodeSpecInfos {
	s.NodeSpec = &v
	return s
}

func (s *AddAINodeRequestAINodeSpecInfos) Validate() error {
	return dara.Validate(s)
}
