// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAINodeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAINodeNum(v int32) *DeleteAINodeRequest
	GetAINodeNum() *int32
	SetAINodePoolId(v string) *DeleteAINodeRequest
	GetAINodePoolId() *string
	SetDBInstanceId(v string) *DeleteAINodeRequest
	GetDBInstanceId() *string
	SetNodeNames(v []*string) *DeleteAINodeRequest
	GetNodeNames() []*string
}

type DeleteAINodeRequest struct {
	// The number of AINodes to be deleted.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	AINodeNum *int32 `json:"AINodeNum,omitempty" xml:"AINodeNum,omitempty"`
	// The ID of the AINode resource pool.
	//
	// example:
	//
	// aipool-xxxx
	AINodePoolId *string `json:"AINodePoolId,omitempty" xml:"AINodePoolId,omitempty"`
	// The ID of the instance.
	//
	// > You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/86911.html) operation to query the information about all AnalyticDB for PostgreSQL instances within a region, including instance IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// gp-xxxxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The list of AINodes to be deleted.
	NodeNames []*string `json:"NodeNames,omitempty" xml:"NodeNames,omitempty" type:"Repeated"`
}

func (s DeleteAINodeRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAINodeRequest) GoString() string {
	return s.String()
}

func (s *DeleteAINodeRequest) GetAINodeNum() *int32 {
	return s.AINodeNum
}

func (s *DeleteAINodeRequest) GetAINodePoolId() *string {
	return s.AINodePoolId
}

func (s *DeleteAINodeRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DeleteAINodeRequest) GetNodeNames() []*string {
	return s.NodeNames
}

func (s *DeleteAINodeRequest) SetAINodeNum(v int32) *DeleteAINodeRequest {
	s.AINodeNum = &v
	return s
}

func (s *DeleteAINodeRequest) SetAINodePoolId(v string) *DeleteAINodeRequest {
	s.AINodePoolId = &v
	return s
}

func (s *DeleteAINodeRequest) SetDBInstanceId(v string) *DeleteAINodeRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DeleteAINodeRequest) SetNodeNames(v []*string) *DeleteAINodeRequest {
	s.NodeNames = v
	return s
}

func (s *DeleteAINodeRequest) Validate() error {
	return dara.Validate(s)
}
