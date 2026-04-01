// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGadInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetGadInstances(v []*DescribeGadInstancesResponseBodyGadInstances) *DescribeGadInstancesResponseBody
	GetGadInstances() []*DescribeGadInstancesResponseBodyGadInstances
	SetRequestId(v string) *DescribeGadInstancesResponseBody
	GetRequestId() *string
}

type DescribeGadInstancesResponseBody struct {
	// The details about the global active database cluster.
	GadInstances []*DescribeGadInstancesResponseBodyGadInstances `json:"GadInstances,omitempty" xml:"GadInstances,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 76AF0609-4195-5DFC-BC78-3AD76FF872BB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeGadInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeGadInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeGadInstancesResponseBody) GetGadInstances() []*DescribeGadInstancesResponseBodyGadInstances {
	return s.GadInstances
}

func (s *DescribeGadInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeGadInstancesResponseBody) SetGadInstances(v []*DescribeGadInstancesResponseBodyGadInstances) *DescribeGadInstancesResponseBody {
	s.GadInstances = v
	return s
}

func (s *DescribeGadInstancesResponseBody) SetRequestId(v string) *DescribeGadInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeGadInstancesResponseBody) Validate() error {
	if s.GadInstances != nil {
		for _, item := range s.GadInstances {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeGadInstancesResponseBodyGadInstances struct {
	// The time when the global active database cluster was created. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2021-10-21T02:57:08Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The name of the cluster.
	//
	// example:
	//
	// GadTest
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The information about each node in the cluster.
	GadInstanceMembers []*DescribeGadInstancesResponseBodyGadInstancesGadInstanceMembers `json:"GadInstanceMembers,omitempty" xml:"GadInstanceMembers,omitempty" type:"Repeated"`
	// The ID of the global active database cluster.
	//
	// example:
	//
	// gad-rm-bp1npi2j8********
	GadInstanceName *string `json:"GadInstanceName,omitempty" xml:"GadInstanceName,omitempty"`
	// The time when the most recent modification was made to the global active database cluster. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2021-10-21T03:01:20Z
	ModificationTime *string `json:"ModificationTime,omitempty" xml:"ModificationTime,omitempty"`
	// The database engine that is run by the global active database cluster.
	//
	// >  The value of this parameter is fixed as **mysql**.
	//
	// example:
	//
	// mysql
	Service *string `json:"Service,omitempty" xml:"Service,omitempty"`
	// The status of the cluster. Valid values:
	//
	// 	- **activation**: The cluster is running.
	//
	// 	- **creating**: The cluster is being created.
	//
	// 	- **replica_adding**: Nodes are being added to the cluster.
	//
	// example:
	//
	// activation
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeGadInstancesResponseBodyGadInstances) String() string {
	return dara.Prettify(s)
}

func (s DescribeGadInstancesResponseBodyGadInstances) GoString() string {
	return s.String()
}

func (s *DescribeGadInstancesResponseBodyGadInstances) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeGadInstancesResponseBodyGadInstances) GetDescription() *string {
	return s.Description
}

func (s *DescribeGadInstancesResponseBodyGadInstances) GetGadInstanceMembers() []*DescribeGadInstancesResponseBodyGadInstancesGadInstanceMembers {
	return s.GadInstanceMembers
}

func (s *DescribeGadInstancesResponseBodyGadInstances) GetGadInstanceName() *string {
	return s.GadInstanceName
}

func (s *DescribeGadInstancesResponseBodyGadInstances) GetModificationTime() *string {
	return s.ModificationTime
}

func (s *DescribeGadInstancesResponseBodyGadInstances) GetService() *string {
	return s.Service
}

func (s *DescribeGadInstancesResponseBodyGadInstances) GetStatus() *string {
	return s.Status
}

func (s *DescribeGadInstancesResponseBodyGadInstances) SetCreationTime(v string) *DescribeGadInstancesResponseBodyGadInstances {
	s.CreationTime = &v
	return s
}

func (s *DescribeGadInstancesResponseBodyGadInstances) SetDescription(v string) *DescribeGadInstancesResponseBodyGadInstances {
	s.Description = &v
	return s
}

func (s *DescribeGadInstancesResponseBodyGadInstances) SetGadInstanceMembers(v []*DescribeGadInstancesResponseBodyGadInstancesGadInstanceMembers) *DescribeGadInstancesResponseBodyGadInstances {
	s.GadInstanceMembers = v
	return s
}

func (s *DescribeGadInstancesResponseBodyGadInstances) SetGadInstanceName(v string) *DescribeGadInstancesResponseBodyGadInstances {
	s.GadInstanceName = &v
	return s
}

func (s *DescribeGadInstancesResponseBodyGadInstances) SetModificationTime(v string) *DescribeGadInstancesResponseBodyGadInstances {
	s.ModificationTime = &v
	return s
}

func (s *DescribeGadInstancesResponseBodyGadInstances) SetService(v string) *DescribeGadInstancesResponseBodyGadInstances {
	s.Service = &v
	return s
}

func (s *DescribeGadInstancesResponseBodyGadInstances) SetStatus(v string) *DescribeGadInstancesResponseBodyGadInstances {
	s.Status = &v
	return s
}

func (s *DescribeGadInstancesResponseBodyGadInstances) Validate() error {
	if s.GadInstanceMembers != nil {
		for _, item := range s.GadInstanceMembers {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeGadInstancesResponseBodyGadInstancesGadInstanceMembers struct {
	// The ID of the node.
	//
	// example:
	//
	// rm-bp1npi2j8********
	DBInstanceID *string `json:"DBInstanceID,omitempty" xml:"DBInstanceID,omitempty"`
	// A JSON array that consists of the details about the Data Transmission Service (DTS) synchronization task.
	//
	// >  Each unit node (secondary node) synchronizes data from the central node (primary node) by using DTS. This parameter contains the synchronization link ID and request ID of DTS.
	//
	// example:
	//
	// {\\"dtsInstanceId\\":\\"dtsm9t107c********\\",\\"dtsRequestId\\":\\"190F0C6C-4BE6-5676-989B-DBDE6D34CD9C\\"}
	DtsInstance *string `json:"DtsInstance,omitempty" xml:"DtsInstance,omitempty"`
	// The database engine that is run by the node.
	//
	// >  The value of this parameter is fixed as **mysql**.
	//
	// example:
	//
	// mysql
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// The database engine version that is run by the node.
	//
	// example:
	//
	// 8.0
	EngineVersion *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	// The ID of the region where the node resides.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmy*****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The type of the node. Valid values:
	//
	// 	- **CENTRAL**: The node is a central node. Each global active database cluster has only one central node. All unit nodes synchronize data from the central node.
	//
	// 	- **UNIT**: The node is a unit node. Each global active database cluster can have up to 10 unit nodes. All unit nodes synchronize data from the central node.
	//
	// example:
	//
	// CENTRAL
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// The node status. Valid values:
	//
	// 	- **activation**: The node is running.
	//
	// 	- **creating**: The node is being created.
	//
	// example:
	//
	// activation
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeGadInstancesResponseBodyGadInstancesGadInstanceMembers) String() string {
	return dara.Prettify(s)
}

func (s DescribeGadInstancesResponseBodyGadInstancesGadInstanceMembers) GoString() string {
	return s.String()
}

func (s *DescribeGadInstancesResponseBodyGadInstancesGadInstanceMembers) GetDBInstanceID() *string {
	return s.DBInstanceID
}

func (s *DescribeGadInstancesResponseBodyGadInstancesGadInstanceMembers) GetDtsInstance() *string {
	return s.DtsInstance
}

func (s *DescribeGadInstancesResponseBodyGadInstancesGadInstanceMembers) GetEngine() *string {
	return s.Engine
}

func (s *DescribeGadInstancesResponseBodyGadInstancesGadInstanceMembers) GetEngineVersion() *string {
	return s.EngineVersion
}

func (s *DescribeGadInstancesResponseBodyGadInstancesGadInstanceMembers) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeGadInstancesResponseBodyGadInstancesGadInstanceMembers) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeGadInstancesResponseBodyGadInstancesGadInstanceMembers) GetRole() *string {
	return s.Role
}

func (s *DescribeGadInstancesResponseBodyGadInstancesGadInstanceMembers) GetStatus() *string {
	return s.Status
}

func (s *DescribeGadInstancesResponseBodyGadInstancesGadInstanceMembers) SetDBInstanceID(v string) *DescribeGadInstancesResponseBodyGadInstancesGadInstanceMembers {
	s.DBInstanceID = &v
	return s
}

func (s *DescribeGadInstancesResponseBodyGadInstancesGadInstanceMembers) SetDtsInstance(v string) *DescribeGadInstancesResponseBodyGadInstancesGadInstanceMembers {
	s.DtsInstance = &v
	return s
}

func (s *DescribeGadInstancesResponseBodyGadInstancesGadInstanceMembers) SetEngine(v string) *DescribeGadInstancesResponseBodyGadInstancesGadInstanceMembers {
	s.Engine = &v
	return s
}

func (s *DescribeGadInstancesResponseBodyGadInstancesGadInstanceMembers) SetEngineVersion(v string) *DescribeGadInstancesResponseBodyGadInstancesGadInstanceMembers {
	s.EngineVersion = &v
	return s
}

func (s *DescribeGadInstancesResponseBodyGadInstancesGadInstanceMembers) SetRegionId(v string) *DescribeGadInstancesResponseBodyGadInstancesGadInstanceMembers {
	s.RegionId = &v
	return s
}

func (s *DescribeGadInstancesResponseBodyGadInstancesGadInstanceMembers) SetResourceGroupId(v string) *DescribeGadInstancesResponseBodyGadInstancesGadInstanceMembers {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeGadInstancesResponseBodyGadInstancesGadInstanceMembers) SetRole(v string) *DescribeGadInstancesResponseBodyGadInstancesGadInstanceMembers {
	s.Role = &v
	return s
}

func (s *DescribeGadInstancesResponseBodyGadInstancesGadInstanceMembers) SetStatus(v string) *DescribeGadInstancesResponseBodyGadInstancesGadInstanceMembers {
	s.Status = &v
	return s
}

func (s *DescribeGadInstancesResponseBodyGadInstancesGadInstanceMembers) Validate() error {
	return dara.Validate(s)
}
