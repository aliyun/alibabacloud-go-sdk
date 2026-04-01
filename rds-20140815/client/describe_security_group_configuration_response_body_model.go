// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSecurityGroupConfigurationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceName(v string) *DescribeSecurityGroupConfigurationResponseBody
	GetDBInstanceName() *string
	SetItems(v *DescribeSecurityGroupConfigurationResponseBodyItems) *DescribeSecurityGroupConfigurationResponseBody
	GetItems() *DescribeSecurityGroupConfigurationResponseBodyItems
	SetRequestId(v string) *DescribeSecurityGroupConfigurationResponseBody
	GetRequestId() *string
}

type DescribeSecurityGroupConfigurationResponseBody struct {
	// The instance ID.
	//
	// example:
	//
	// rm-uf6wjk5xxxxxx
	DBInstanceName *string                                              `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	Items          *DescribeSecurityGroupConfigurationResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 87BDAE8C-xxxx-4A26-BBCC-7D1DD31D630D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeSecurityGroupConfigurationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSecurityGroupConfigurationResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSecurityGroupConfigurationResponseBody) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *DescribeSecurityGroupConfigurationResponseBody) GetItems() *DescribeSecurityGroupConfigurationResponseBodyItems {
	return s.Items
}

func (s *DescribeSecurityGroupConfigurationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSecurityGroupConfigurationResponseBody) SetDBInstanceName(v string) *DescribeSecurityGroupConfigurationResponseBody {
	s.DBInstanceName = &v
	return s
}

func (s *DescribeSecurityGroupConfigurationResponseBody) SetItems(v *DescribeSecurityGroupConfigurationResponseBodyItems) *DescribeSecurityGroupConfigurationResponseBody {
	s.Items = v
	return s
}

func (s *DescribeSecurityGroupConfigurationResponseBody) SetRequestId(v string) *DescribeSecurityGroupConfigurationResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSecurityGroupConfigurationResponseBody) Validate() error {
	if s.Items != nil {
		if err := s.Items.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeSecurityGroupConfigurationResponseBodyItems struct {
	EcsSecurityGroupRelation []*DescribeSecurityGroupConfigurationResponseBodyItemsEcsSecurityGroupRelation `json:"EcsSecurityGroupRelation,omitempty" xml:"EcsSecurityGroupRelation,omitempty" type:"Repeated"`
}

func (s DescribeSecurityGroupConfigurationResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeSecurityGroupConfigurationResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeSecurityGroupConfigurationResponseBodyItems) GetEcsSecurityGroupRelation() []*DescribeSecurityGroupConfigurationResponseBodyItemsEcsSecurityGroupRelation {
	return s.EcsSecurityGroupRelation
}

func (s *DescribeSecurityGroupConfigurationResponseBodyItems) SetEcsSecurityGroupRelation(v []*DescribeSecurityGroupConfigurationResponseBodyItemsEcsSecurityGroupRelation) *DescribeSecurityGroupConfigurationResponseBodyItems {
	s.EcsSecurityGroupRelation = v
	return s
}

func (s *DescribeSecurityGroupConfigurationResponseBodyItems) Validate() error {
	if s.EcsSecurityGroupRelation != nil {
		for _, item := range s.EcsSecurityGroupRelation {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeSecurityGroupConfigurationResponseBodyItemsEcsSecurityGroupRelation struct {
	NetworkType       *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	RegionId          *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SecurityGroupId   *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	SecurityGroupName *string `json:"SecurityGroupName,omitempty" xml:"SecurityGroupName,omitempty"`
}

func (s DescribeSecurityGroupConfigurationResponseBodyItemsEcsSecurityGroupRelation) String() string {
	return dara.Prettify(s)
}

func (s DescribeSecurityGroupConfigurationResponseBodyItemsEcsSecurityGroupRelation) GoString() string {
	return s.String()
}

func (s *DescribeSecurityGroupConfigurationResponseBodyItemsEcsSecurityGroupRelation) GetNetworkType() *string {
	return s.NetworkType
}

func (s *DescribeSecurityGroupConfigurationResponseBodyItemsEcsSecurityGroupRelation) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeSecurityGroupConfigurationResponseBodyItemsEcsSecurityGroupRelation) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *DescribeSecurityGroupConfigurationResponseBodyItemsEcsSecurityGroupRelation) GetSecurityGroupName() *string {
	return s.SecurityGroupName
}

func (s *DescribeSecurityGroupConfigurationResponseBodyItemsEcsSecurityGroupRelation) SetNetworkType(v string) *DescribeSecurityGroupConfigurationResponseBodyItemsEcsSecurityGroupRelation {
	s.NetworkType = &v
	return s
}

func (s *DescribeSecurityGroupConfigurationResponseBodyItemsEcsSecurityGroupRelation) SetRegionId(v string) *DescribeSecurityGroupConfigurationResponseBodyItemsEcsSecurityGroupRelation {
	s.RegionId = &v
	return s
}

func (s *DescribeSecurityGroupConfigurationResponseBodyItemsEcsSecurityGroupRelation) SetSecurityGroupId(v string) *DescribeSecurityGroupConfigurationResponseBodyItemsEcsSecurityGroupRelation {
	s.SecurityGroupId = &v
	return s
}

func (s *DescribeSecurityGroupConfigurationResponseBodyItemsEcsSecurityGroupRelation) SetSecurityGroupName(v string) *DescribeSecurityGroupConfigurationResponseBodyItemsEcsSecurityGroupRelation {
	s.SecurityGroupName = &v
	return s
}

func (s *DescribeSecurityGroupConfigurationResponseBodyItemsEcsSecurityGroupRelation) Validate() error {
	return dara.Validate(s)
}
