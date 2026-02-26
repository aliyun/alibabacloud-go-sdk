// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySecurityGroupConfigurationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceName(v string) *ModifySecurityGroupConfigurationResponseBody
	GetDBInstanceName() *string
	SetItems(v *ModifySecurityGroupConfigurationResponseBodyItems) *ModifySecurityGroupConfigurationResponseBody
	GetItems() *ModifySecurityGroupConfigurationResponseBodyItems
	SetRequestId(v string) *ModifySecurityGroupConfigurationResponseBody
	GetRequestId() *string
}

type ModifySecurityGroupConfigurationResponseBody struct {
	// The instance ID.
	//
	// example:
	//
	// rm-uf6wjk5xxxxxx
	DBInstanceName *string                                            `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	Items          *ModifySecurityGroupConfigurationResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 8585861B-8F0D-4D17-9460-C42255EB10C0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifySecurityGroupConfigurationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifySecurityGroupConfigurationResponseBody) GoString() string {
	return s.String()
}

func (s *ModifySecurityGroupConfigurationResponseBody) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *ModifySecurityGroupConfigurationResponseBody) GetItems() *ModifySecurityGroupConfigurationResponseBodyItems {
	return s.Items
}

func (s *ModifySecurityGroupConfigurationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifySecurityGroupConfigurationResponseBody) SetDBInstanceName(v string) *ModifySecurityGroupConfigurationResponseBody {
	s.DBInstanceName = &v
	return s
}

func (s *ModifySecurityGroupConfigurationResponseBody) SetItems(v *ModifySecurityGroupConfigurationResponseBodyItems) *ModifySecurityGroupConfigurationResponseBody {
	s.Items = v
	return s
}

func (s *ModifySecurityGroupConfigurationResponseBody) SetRequestId(v string) *ModifySecurityGroupConfigurationResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifySecurityGroupConfigurationResponseBody) Validate() error {
	if s.Items != nil {
		if err := s.Items.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModifySecurityGroupConfigurationResponseBodyItems struct {
	EcsSecurityGroupRelation []*ModifySecurityGroupConfigurationResponseBodyItemsEcsSecurityGroupRelation `json:"EcsSecurityGroupRelation,omitempty" xml:"EcsSecurityGroupRelation,omitempty" type:"Repeated"`
}

func (s ModifySecurityGroupConfigurationResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s ModifySecurityGroupConfigurationResponseBodyItems) GoString() string {
	return s.String()
}

func (s *ModifySecurityGroupConfigurationResponseBodyItems) GetEcsSecurityGroupRelation() []*ModifySecurityGroupConfigurationResponseBodyItemsEcsSecurityGroupRelation {
	return s.EcsSecurityGroupRelation
}

func (s *ModifySecurityGroupConfigurationResponseBodyItems) SetEcsSecurityGroupRelation(v []*ModifySecurityGroupConfigurationResponseBodyItemsEcsSecurityGroupRelation) *ModifySecurityGroupConfigurationResponseBodyItems {
	s.EcsSecurityGroupRelation = v
	return s
}

func (s *ModifySecurityGroupConfigurationResponseBodyItems) Validate() error {
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

type ModifySecurityGroupConfigurationResponseBodyItemsEcsSecurityGroupRelation struct {
	NetworkType     *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
}

func (s ModifySecurityGroupConfigurationResponseBodyItemsEcsSecurityGroupRelation) String() string {
	return dara.Prettify(s)
}

func (s ModifySecurityGroupConfigurationResponseBodyItemsEcsSecurityGroupRelation) GoString() string {
	return s.String()
}

func (s *ModifySecurityGroupConfigurationResponseBodyItemsEcsSecurityGroupRelation) GetNetworkType() *string {
	return s.NetworkType
}

func (s *ModifySecurityGroupConfigurationResponseBodyItemsEcsSecurityGroupRelation) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifySecurityGroupConfigurationResponseBodyItemsEcsSecurityGroupRelation) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *ModifySecurityGroupConfigurationResponseBodyItemsEcsSecurityGroupRelation) SetNetworkType(v string) *ModifySecurityGroupConfigurationResponseBodyItemsEcsSecurityGroupRelation {
	s.NetworkType = &v
	return s
}

func (s *ModifySecurityGroupConfigurationResponseBodyItemsEcsSecurityGroupRelation) SetRegionId(v string) *ModifySecurityGroupConfigurationResponseBodyItemsEcsSecurityGroupRelation {
	s.RegionId = &v
	return s
}

func (s *ModifySecurityGroupConfigurationResponseBodyItemsEcsSecurityGroupRelation) SetSecurityGroupId(v string) *ModifySecurityGroupConfigurationResponseBodyItemsEcsSecurityGroupRelation {
	s.SecurityGroupId = &v
	return s
}

func (s *ModifySecurityGroupConfigurationResponseBodyItemsEcsSecurityGroupRelation) Validate() error {
	return dara.Validate(s)
}
