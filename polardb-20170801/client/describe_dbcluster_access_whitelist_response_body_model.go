// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBClusterAccessWhitelistResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterSecurityGroups(v *DescribeDBClusterAccessWhitelistResponseBodyDBClusterSecurityGroups) *DescribeDBClusterAccessWhitelistResponseBody
	GetDBClusterSecurityGroups() *DescribeDBClusterAccessWhitelistResponseBodyDBClusterSecurityGroups
	SetItems(v *DescribeDBClusterAccessWhitelistResponseBodyItems) *DescribeDBClusterAccessWhitelistResponseBody
	GetItems() *DescribeDBClusterAccessWhitelistResponseBodyItems
	SetRequestId(v string) *DescribeDBClusterAccessWhitelistResponseBody
	GetRequestId() *string
}

type DescribeDBClusterAccessWhitelistResponseBody struct {
	// The Elastic Compute Service (ECS) security groups that are associated with the cluster.
	DBClusterSecurityGroups *DescribeDBClusterAccessWhitelistResponseBodyDBClusterSecurityGroups `json:"DBClusterSecurityGroups,omitempty" xml:"DBClusterSecurityGroups,omitempty" type:"Struct"`
	// The details about the cluster.
	Items *DescribeDBClusterAccessWhitelistResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 559E91A2-CDA3-4E9F-808B-29D738******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDBClusterAccessWhitelistResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterAccessWhitelistResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterAccessWhitelistResponseBody) GetDBClusterSecurityGroups() *DescribeDBClusterAccessWhitelistResponseBodyDBClusterSecurityGroups {
	return s.DBClusterSecurityGroups
}

func (s *DescribeDBClusterAccessWhitelistResponseBody) GetItems() *DescribeDBClusterAccessWhitelistResponseBodyItems {
	return s.Items
}

func (s *DescribeDBClusterAccessWhitelistResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDBClusterAccessWhitelistResponseBody) SetDBClusterSecurityGroups(v *DescribeDBClusterAccessWhitelistResponseBodyDBClusterSecurityGroups) *DescribeDBClusterAccessWhitelistResponseBody {
	s.DBClusterSecurityGroups = v
	return s
}

func (s *DescribeDBClusterAccessWhitelistResponseBody) SetItems(v *DescribeDBClusterAccessWhitelistResponseBodyItems) *DescribeDBClusterAccessWhitelistResponseBody {
	s.Items = v
	return s
}

func (s *DescribeDBClusterAccessWhitelistResponseBody) SetRequestId(v string) *DescribeDBClusterAccessWhitelistResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBClusterAccessWhitelistResponseBody) Validate() error {
	if s.DBClusterSecurityGroups != nil {
		if err := s.DBClusterSecurityGroups.Validate(); err != nil {
			return err
		}
	}
	if s.Items != nil {
		if err := s.Items.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDBClusterAccessWhitelistResponseBodyDBClusterSecurityGroups struct {
	DBClusterSecurityGroup []*DescribeDBClusterAccessWhitelistResponseBodyDBClusterSecurityGroupsDBClusterSecurityGroup `json:"DBClusterSecurityGroup,omitempty" xml:"DBClusterSecurityGroup,omitempty" type:"Repeated"`
}

func (s DescribeDBClusterAccessWhitelistResponseBodyDBClusterSecurityGroups) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterAccessWhitelistResponseBodyDBClusterSecurityGroups) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterAccessWhitelistResponseBodyDBClusterSecurityGroups) GetDBClusterSecurityGroup() []*DescribeDBClusterAccessWhitelistResponseBodyDBClusterSecurityGroupsDBClusterSecurityGroup {
	return s.DBClusterSecurityGroup
}

func (s *DescribeDBClusterAccessWhitelistResponseBodyDBClusterSecurityGroups) SetDBClusterSecurityGroup(v []*DescribeDBClusterAccessWhitelistResponseBodyDBClusterSecurityGroupsDBClusterSecurityGroup) *DescribeDBClusterAccessWhitelistResponseBodyDBClusterSecurityGroups {
	s.DBClusterSecurityGroup = v
	return s
}

func (s *DescribeDBClusterAccessWhitelistResponseBodyDBClusterSecurityGroups) Validate() error {
	if s.DBClusterSecurityGroup != nil {
		for _, item := range s.DBClusterSecurityGroup {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDBClusterAccessWhitelistResponseBodyDBClusterSecurityGroupsDBClusterSecurityGroup struct {
	// The ID of the ECS security group.
	//
	// example:
	//
	// sg-bp**************
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The name of the ECS security group.
	//
	// example:
	//
	// vpc-********************
	SecurityGroupName *string `json:"SecurityGroupName,omitempty" xml:"SecurityGroupName,omitempty"`
}

func (s DescribeDBClusterAccessWhitelistResponseBodyDBClusterSecurityGroupsDBClusterSecurityGroup) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterAccessWhitelistResponseBodyDBClusterSecurityGroupsDBClusterSecurityGroup) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterAccessWhitelistResponseBodyDBClusterSecurityGroupsDBClusterSecurityGroup) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *DescribeDBClusterAccessWhitelistResponseBodyDBClusterSecurityGroupsDBClusterSecurityGroup) GetSecurityGroupName() *string {
	return s.SecurityGroupName
}

func (s *DescribeDBClusterAccessWhitelistResponseBodyDBClusterSecurityGroupsDBClusterSecurityGroup) SetSecurityGroupId(v string) *DescribeDBClusterAccessWhitelistResponseBodyDBClusterSecurityGroupsDBClusterSecurityGroup {
	s.SecurityGroupId = &v
	return s
}

func (s *DescribeDBClusterAccessWhitelistResponseBodyDBClusterSecurityGroupsDBClusterSecurityGroup) SetSecurityGroupName(v string) *DescribeDBClusterAccessWhitelistResponseBodyDBClusterSecurityGroupsDBClusterSecurityGroup {
	s.SecurityGroupName = &v
	return s
}

func (s *DescribeDBClusterAccessWhitelistResponseBodyDBClusterSecurityGroupsDBClusterSecurityGroup) Validate() error {
	return dara.Validate(s)
}

type DescribeDBClusterAccessWhitelistResponseBodyItems struct {
	DBClusterIPArray []*DescribeDBClusterAccessWhitelistResponseBodyItemsDBClusterIPArray `json:"DBClusterIPArray,omitempty" xml:"DBClusterIPArray,omitempty" type:"Repeated"`
}

func (s DescribeDBClusterAccessWhitelistResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterAccessWhitelistResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterAccessWhitelistResponseBodyItems) GetDBClusterIPArray() []*DescribeDBClusterAccessWhitelistResponseBodyItemsDBClusterIPArray {
	return s.DBClusterIPArray
}

func (s *DescribeDBClusterAccessWhitelistResponseBodyItems) SetDBClusterIPArray(v []*DescribeDBClusterAccessWhitelistResponseBodyItemsDBClusterIPArray) *DescribeDBClusterAccessWhitelistResponseBodyItems {
	s.DBClusterIPArray = v
	return s
}

func (s *DescribeDBClusterAccessWhitelistResponseBodyItems) Validate() error {
	if s.DBClusterIPArray != nil {
		for _, item := range s.DBClusterIPArray {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDBClusterAccessWhitelistResponseBodyItemsDBClusterIPArray struct {
	// The attributes of the IP whitelist group. Set this parameter to **hidden*	- to hide the IP whitelist group in the console.
	//
	// > 	- The IP whitelist group that has appeared in the console cannot be hidden.
	//
	// > 	- This parameter can be specified only when the **WhiteListType*	- parameter is set to **IP**.
	//
	// example:
	//
	// hidden
	DBClusterIPArrayAttribute *string `json:"DBClusterIPArrayAttribute,omitempty" xml:"DBClusterIPArrayAttribute,omitempty"`
	// The name of the IP whitelist group. The group name must be 2 to 120 characters in length and consists of lowercase letters and digits. It must start with a letter, and end with a letter or a digit.
	//
	// 	- If the specified whitelist group name does not exist, the whitelist group is created.
	//
	// 	- If the specified whitelist group name exists, the whitelist group is modified.
	//
	// 	- If you do not specify this parameter, the default group is modified.
	//
	// > 	- You can create a maximum of 50 IP whitelist groups for a cluster.
	//
	// >	- This parameter can be specified only when the **WhiteListType*	- parameter is set to **IP**.
	//
	// example:
	//
	// default
	DBClusterIPArrayName *string `json:"DBClusterIPArrayName,omitempty" xml:"DBClusterIPArrayName,omitempty"`
	// The IP addresses or Classless Inter-Domain Routing (CIDR) blocks in the IP whitelist group. You can add 1,000 IP addresses or CIDR blocks to all the IP whitelist groups. Separate multiple IP addresses with commas (,). The following two formats are supported:
	//
	// 	- IP addresses. Example: 10.23.12.24.
	//
	// 	- CIDR blocks. Example: 10.23.12.24/24. 24 indicates the length of the prefix of the CIDR block. The length is the range of 1 to 32.
	//
	// >  This parameter can be specified only when the **WhiteListType*	- parameter is set to **IP**.
	//
	// example:
	//
	// 11.***.***.***,11.***.***.***
	SecurityIps *string `json:"SecurityIps,omitempty" xml:"SecurityIps,omitempty"`
}

func (s DescribeDBClusterAccessWhitelistResponseBodyItemsDBClusterIPArray) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterAccessWhitelistResponseBodyItemsDBClusterIPArray) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterAccessWhitelistResponseBodyItemsDBClusterIPArray) GetDBClusterIPArrayAttribute() *string {
	return s.DBClusterIPArrayAttribute
}

func (s *DescribeDBClusterAccessWhitelistResponseBodyItemsDBClusterIPArray) GetDBClusterIPArrayName() *string {
	return s.DBClusterIPArrayName
}

func (s *DescribeDBClusterAccessWhitelistResponseBodyItemsDBClusterIPArray) GetSecurityIps() *string {
	return s.SecurityIps
}

func (s *DescribeDBClusterAccessWhitelistResponseBodyItemsDBClusterIPArray) SetDBClusterIPArrayAttribute(v string) *DescribeDBClusterAccessWhitelistResponseBodyItemsDBClusterIPArray {
	s.DBClusterIPArrayAttribute = &v
	return s
}

func (s *DescribeDBClusterAccessWhitelistResponseBodyItemsDBClusterIPArray) SetDBClusterIPArrayName(v string) *DescribeDBClusterAccessWhitelistResponseBodyItemsDBClusterIPArray {
	s.DBClusterIPArrayName = &v
	return s
}

func (s *DescribeDBClusterAccessWhitelistResponseBodyItemsDBClusterIPArray) SetSecurityIps(v string) *DescribeDBClusterAccessWhitelistResponseBodyItemsDBClusterIPArray {
	s.SecurityIps = &v
	return s
}

func (s *DescribeDBClusterAccessWhitelistResponseBodyItemsDBClusterIPArray) Validate() error {
	return dara.Validate(s)
}
