// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeKmsAssociateResourcesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAssociateDBInstances(v []*DescribeKmsAssociateResourcesResponseBodyAssociateDBInstances) *DescribeKmsAssociateResourcesResponseBody
	GetAssociateDBInstances() []*DescribeKmsAssociateResourcesResponseBodyAssociateDBInstances
	SetAssociateStatus(v bool) *DescribeKmsAssociateResourcesResponseBody
	GetAssociateStatus() *bool
	SetRequestId(v string) *DescribeKmsAssociateResourcesResponseBody
	GetRequestId() *string
}

type DescribeKmsAssociateResourcesResponseBody struct {
	// The information about the associated ApsaraDB RDS instances.
	AssociateDBInstances []*DescribeKmsAssociateResourcesResponseBodyAssociateDBInstances `json:"AssociateDBInstances,omitempty" xml:"AssociateDBInstances,omitempty" type:"Repeated"`
	// Indicates whether an associated RDS instance exists.
	//
	// - **true**: Yes
	//
	// - **false**: No
	//
	// example:
	//
	// true
	AssociateStatus *bool `json:"AssociateStatus,omitempty" xml:"AssociateStatus,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1AD222E9-E606-4A42-BF6D-8A4442913CEF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeKmsAssociateResourcesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeKmsAssociateResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeKmsAssociateResourcesResponseBody) GetAssociateDBInstances() []*DescribeKmsAssociateResourcesResponseBodyAssociateDBInstances {
	return s.AssociateDBInstances
}

func (s *DescribeKmsAssociateResourcesResponseBody) GetAssociateStatus() *bool {
	return s.AssociateStatus
}

func (s *DescribeKmsAssociateResourcesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeKmsAssociateResourcesResponseBody) SetAssociateDBInstances(v []*DescribeKmsAssociateResourcesResponseBodyAssociateDBInstances) *DescribeKmsAssociateResourcesResponseBody {
	s.AssociateDBInstances = v
	return s
}

func (s *DescribeKmsAssociateResourcesResponseBody) SetAssociateStatus(v bool) *DescribeKmsAssociateResourcesResponseBody {
	s.AssociateStatus = &v
	return s
}

func (s *DescribeKmsAssociateResourcesResponseBody) SetRequestId(v string) *DescribeKmsAssociateResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeKmsAssociateResourcesResponseBody) Validate() error {
	if s.AssociateDBInstances != nil {
		for _, item := range s.AssociateDBInstances {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeKmsAssociateResourcesResponseBodyAssociateDBInstances struct {
	// The instance ID.
	//
	// example:
	//
	// pgm-bp16p6f68130****
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// The database engine. Valid values:
	//
	// 	- **MySQL**
	//
	// 	- **SQLServer**
	//
	// 	- **PostgreSQL**
	//
	// example:
	//
	// PostgreSQL
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// The purpose of the key. Valid values:
	//
	// 	- **DiskEncryption**: cloud disk encryption
	//
	// 	- **TDE**: transparent data encryption
	//
	// example:
	//
	// DiskEncryption
	KeyUsedBy *string `json:"KeyUsedBy,omitempty" xml:"KeyUsedBy,omitempty"`
	// The state of the instance. Valid values:
	//
	// 	- **CREATING**: The instance is being created.
	//
	// 	- **ACTIVATION**: The instance is running.
	//
	// 	- **DELETING**: The instance is being deleted.
	//
	// 	- **RESTARTING**: The instance is being restarted.
	//
	// 	- **INS_MAINTAINING**: The configuration of the instance is being changed.
	//
	// 	- **INS_MAINTAINING**: The instance is being maintained.
	//
	// 	- **BACKUP_RECOVERING**: The instance is being restored.
	//
	// 	- **NET_MODIFYING**: The network type of the instance is being changed.
	//
	// example:
	//
	// ACTIVATION
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeKmsAssociateResourcesResponseBodyAssociateDBInstances) String() string {
	return dara.Prettify(s)
}

func (s DescribeKmsAssociateResourcesResponseBodyAssociateDBInstances) GoString() string {
	return s.String()
}

func (s *DescribeKmsAssociateResourcesResponseBodyAssociateDBInstances) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *DescribeKmsAssociateResourcesResponseBodyAssociateDBInstances) GetEngine() *string {
	return s.Engine
}

func (s *DescribeKmsAssociateResourcesResponseBodyAssociateDBInstances) GetKeyUsedBy() *string {
	return s.KeyUsedBy
}

func (s *DescribeKmsAssociateResourcesResponseBodyAssociateDBInstances) GetStatus() *string {
	return s.Status
}

func (s *DescribeKmsAssociateResourcesResponseBodyAssociateDBInstances) SetDBInstanceName(v string) *DescribeKmsAssociateResourcesResponseBodyAssociateDBInstances {
	s.DBInstanceName = &v
	return s
}

func (s *DescribeKmsAssociateResourcesResponseBodyAssociateDBInstances) SetEngine(v string) *DescribeKmsAssociateResourcesResponseBodyAssociateDBInstances {
	s.Engine = &v
	return s
}

func (s *DescribeKmsAssociateResourcesResponseBodyAssociateDBInstances) SetKeyUsedBy(v string) *DescribeKmsAssociateResourcesResponseBodyAssociateDBInstances {
	s.KeyUsedBy = &v
	return s
}

func (s *DescribeKmsAssociateResourcesResponseBodyAssociateDBInstances) SetStatus(v string) *DescribeKmsAssociateResourcesResponseBodyAssociateDBInstances {
	s.Status = &v
	return s
}

func (s *DescribeKmsAssociateResourcesResponseBodyAssociateDBInstances) Validate() error {
	return dara.Validate(s)
}
