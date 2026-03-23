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
	AssociateDBInstances []*DescribeKmsAssociateResourcesResponseBodyAssociateDBInstances `json:"AssociateDBInstances,omitempty" xml:"AssociateDBInstances,omitempty" type:"Repeated"`
	AssociateStatus      *bool                                                            `json:"AssociateStatus,omitempty" xml:"AssociateStatus,omitempty"`
	RequestId            *string                                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	Engine         *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	KeyUsedBy      *string `json:"KeyUsedBy,omitempty" xml:"KeyUsedBy,omitempty"`
	Status         *string `json:"Status,omitempty" xml:"Status,omitempty"`
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
