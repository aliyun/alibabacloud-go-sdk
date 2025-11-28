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
	// The ID of the resource pool to which the AI node belongs.
	//
	// example:
	//
	// aipool-xxxxx
	AINodePoolId *string `json:"AINodePoolId,omitempty" xml:"AINodePoolId,omitempty"`
	// The AI node specifications.
	//
	// This parameter is required.
	AINodeSpecInfos []*AddAINodeRequestAINodeSpecInfos `json:"AINodeSpecInfos,omitempty" xml:"AINodeSpecInfos,omitempty" type:"Repeated"`
	// The cluster ID.
	//
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
	// The number of AI nodes.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2
	NodeNum *string `json:"NodeNum,omitempty" xml:"NodeNum,omitempty"`
	// The AI node specifications.
	//
	//     ADB.AIStandard.1
	//
	//     ADB.AIMedium.1
	//
	//     ADB.AILarge.1
	//
	//     ADB.AIStandard.2
	//
	//     ADB.AIMedium.2
	//
	//     ADB.AILarge.2
	//
	//     ADB.AIXLarge.2
	//
	//     ADB.AIStandard.6
	//
	//     ADB.AIMedium.6
	//
	//     ADB.AILarge.6
	//
	//     ADB.AIXLarge.6
	//
	//     ADB.AIStandard.3
	//
	//     ADB.AIMedium.3
	//
	//     ADB.AILarge.3
	//
	//     ADB.AIXLarge.3
	//
	//     ADB.AIStandard.4
	//
	//     ADB.AIMedium.4
	//
	//     ADB.AILarge.4
	//
	//     ADB.AIXLarge.4
	//
	//     ADB.AIStandard.5
	//
	//     ADB.AIMedium.5
	//
	//     ADB.AILarge.5
	//
	//     ADB.AIXLarge.5
	//
	//     ADB.AIStandard.8
	//
	//     ADB.AIMedium.8
	//
	//     ADB.AILarge.8
	//
	//     ADB.AIXLarge.8
	//
	//     ADB.AI2XLarge.8
	//
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
