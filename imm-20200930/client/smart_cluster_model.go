// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSmartCluster interface {
	dara.Model
	String() string
	GoString() string
	SetCreateTime(v string) *SmartCluster
	GetCreateTime() *string
	SetDatasetName(v string) *SmartCluster
	GetDatasetName() *string
	SetDescription(v string) *SmartCluster
	GetDescription() *string
	SetName(v string) *SmartCluster
	GetName() *string
	SetObjectId(v string) *SmartCluster
	GetObjectId() *string
	SetObjectStatus(v string) *SmartCluster
	GetObjectStatus() *string
	SetObjectType(v string) *SmartCluster
	GetObjectType() *string
	SetOwnerId(v string) *SmartCluster
	GetOwnerId() *string
	SetProjectName(v string) *SmartCluster
	GetProjectName() *string
	SetRule(v *SmartClusterRule) *SmartCluster
	GetRule() *SmartClusterRule
	SetUpdateTime(v string) *SmartCluster
	GetUpdateTime() *string
}

type SmartCluster struct {
	// The time when the cluster was created.
	//
	// example:
	//
	// 2022-01-14T10:17:18.102700407+08:00
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The name of the dataset.
	//
	// example:
	//
	// MyDataset
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	// The description of the cluster.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the cluster.
	//
	// example:
	//
	// MySmartCluster1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the cluster.
	//
	// example:
	//
	// SmartCluster-12cd1645-deae-4b5e-9434-613747b75f6d
	ObjectId *string `json:"ObjectId,omitempty" xml:"ObjectId,omitempty"`
	// The status of the cluster.
	//
	// example:
	//
	// Succeeded
	ObjectStatus *string `json:"ObjectStatus,omitempty" xml:"ObjectStatus,omitempty"`
	// The type of the cluster.
	//
	// example:
	//
	// smart-cluster
	ObjectType *string `json:"ObjectType,omitempty" xml:"ObjectType,omitempty"`
	// The user ID.
	//
	// example:
	//
	// 1023210024677934
	OwnerId *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The name of the project.
	//
	// example:
	//
	// MyProject
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The clustering rule.
	Rule *SmartClusterRule `json:"Rule,omitempty" xml:"Rule,omitempty"`
	// The time when the cluster was updated.
	//
	// example:
	//
	// 2022-01-14T10:17:18.102700407+08:00
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s SmartCluster) String() string {
	return dara.Prettify(s)
}

func (s SmartCluster) GoString() string {
	return s.String()
}

func (s *SmartCluster) GetCreateTime() *string {
	return s.CreateTime
}

func (s *SmartCluster) GetDatasetName() *string {
	return s.DatasetName
}

func (s *SmartCluster) GetDescription() *string {
	return s.Description
}

func (s *SmartCluster) GetName() *string {
	return s.Name
}

func (s *SmartCluster) GetObjectId() *string {
	return s.ObjectId
}

func (s *SmartCluster) GetObjectStatus() *string {
	return s.ObjectStatus
}

func (s *SmartCluster) GetObjectType() *string {
	return s.ObjectType
}

func (s *SmartCluster) GetOwnerId() *string {
	return s.OwnerId
}

func (s *SmartCluster) GetProjectName() *string {
	return s.ProjectName
}

func (s *SmartCluster) GetRule() *SmartClusterRule {
	return s.Rule
}

func (s *SmartCluster) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *SmartCluster) SetCreateTime(v string) *SmartCluster {
	s.CreateTime = &v
	return s
}

func (s *SmartCluster) SetDatasetName(v string) *SmartCluster {
	s.DatasetName = &v
	return s
}

func (s *SmartCluster) SetDescription(v string) *SmartCluster {
	s.Description = &v
	return s
}

func (s *SmartCluster) SetName(v string) *SmartCluster {
	s.Name = &v
	return s
}

func (s *SmartCluster) SetObjectId(v string) *SmartCluster {
	s.ObjectId = &v
	return s
}

func (s *SmartCluster) SetObjectStatus(v string) *SmartCluster {
	s.ObjectStatus = &v
	return s
}

func (s *SmartCluster) SetObjectType(v string) *SmartCluster {
	s.ObjectType = &v
	return s
}

func (s *SmartCluster) SetOwnerId(v string) *SmartCluster {
	s.OwnerId = &v
	return s
}

func (s *SmartCluster) SetProjectName(v string) *SmartCluster {
	s.ProjectName = &v
	return s
}

func (s *SmartCluster) SetRule(v *SmartClusterRule) *SmartCluster {
	s.Rule = v
	return s
}

func (s *SmartCluster) SetUpdateTime(v string) *SmartCluster {
	s.UpdateTime = &v
	return s
}

func (s *SmartCluster) Validate() error {
	if s.Rule != nil {
		if err := s.Rule.Validate(); err != nil {
			return err
		}
	}
	return nil
}
