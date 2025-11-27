// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeResourceDetailsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBackupDataSize(v int64) *DescribeResourceDetailsResponseBody
	GetBackupDataSize() *int64
	SetBackupLogSize(v int64) *DescribeResourceDetailsResponseBody
	GetBackupLogSize() *int64
	SetBackupSize(v int64) *DescribeResourceDetailsResponseBody
	GetBackupSize() *int64
	SetDbInstanceStorage(v int64) *DescribeResourceDetailsResponseBody
	GetDbInstanceStorage() *int64
	SetDbProxyInstanceName(v string) *DescribeResourceDetailsResponseBody
	GetDbProxyInstanceName() *string
	SetDiskUsed(v int64) *DescribeResourceDetailsResponseBody
	GetDiskUsed() *int64
	SetInstanceStorageType(v string) *DescribeResourceDetailsResponseBody
	GetInstanceStorageType() *string
	SetRdsEcsSecurityGroupRel(v []*DescribeResourceDetailsResponseBodyRdsEcsSecurityGroupRel) *DescribeResourceDetailsResponseBody
	GetRdsEcsSecurityGroupRel() []*DescribeResourceDetailsResponseBodyRdsEcsSecurityGroupRel
	SetRegion(v string) *DescribeResourceDetailsResponseBody
	GetRegion() *string
	SetRequestId(v string) *DescribeResourceDetailsResponseBody
	GetRequestId() *string
	SetResourceGroupId(v string) *DescribeResourceDetailsResponseBody
	GetResourceGroupId() *string
	SetSecurityIPList(v string) *DescribeResourceDetailsResponseBody
	GetSecurityIPList() *string
	SetVSwitchId(v string) *DescribeResourceDetailsResponseBody
	GetVSwitchId() *string
	SetVpcId(v string) *DescribeResourceDetailsResponseBody
	GetVpcId() *string
}

type DescribeResourceDetailsResponseBody struct {
	// The storage that is occupied by data backup files, excluding archived backup files, on the instance. Unit: bytes.
	//
	// example:
	//
	// 8139046912
	BackupDataSize *int64 `json:"BackupDataSize,omitempty" xml:"BackupDataSize,omitempty"`
	// The size of the backup log. Unit: bytes.
	//
	// example:
	//
	// 21183797
	BackupLogSize *int64 `json:"BackupLogSize,omitempty" xml:"BackupLogSize,omitempty"`
	// The size of the backup data. Unit: MB.
	//
	// example:
	//
	// 53002759
	BackupSize *int64 `json:"BackupSize,omitempty" xml:"BackupSize,omitempty"`
	// The disk capacity of the instance.
	//
	// example:
	//
	// 200
	DbInstanceStorage *int64 `json:"DbInstanceStorage,omitempty" xml:"DbInstanceStorage,omitempty"`
	// The name of the proxy instance.
	//
	// example:
	//
	// mr-n1m1wjrylfolvrt67s
	DbProxyInstanceName *string `json:"DbProxyInstanceName,omitempty" xml:"DbProxyInstanceName,omitempty"`
	// The total storage used. The value is the sum of the DataSize and LogSize values. Unit: bytes. The value -1 indicates that no data files or log files are stored.
	//
	// example:
	//
	// 4871684096
	DiskUsed *int64 `json:"DiskUsed,omitempty" xml:"DiskUsed,omitempty"`
	// The storage type of the instance.
	//
	// example:
	//
	// cloud_essd
	InstanceStorageType *string `json:"InstanceStorageType,omitempty" xml:"InstanceStorageType,omitempty"`
	// The rule for the IP address whitelist of the instance.
	RdsEcsSecurityGroupRel []*DescribeResourceDetailsResponseBodyRdsEcsSecurityGroupRel `json:"RdsEcsSecurityGroupRel,omitempty" xml:"RdsEcsSecurityGroupRel,omitempty" type:"Repeated"`
	// The region ID.
	//
	// example:
	//
	// cn-beijing
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The request ID.
	//
	// example:
	//
	// EA815761-F7AC-5CFE-A1AC-709D6A00B58A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmv3h25bj7yhq
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The IP address whitelist of the instance. For more information, see [Configure IP address whitelists](https://help.aliyun.com/document_detail/43185.html). If the returned IP address whitelist contains more than one entry, these entries are separated with commas (,). Each entry is unique and up to 1,000 entries are returned. The entries in the IP address whitelist must be in one of the following formats:
	//
	// 	- IP addresses, such as 10.10.XX.XX.
	//
	// 	- CIDR blocks, such as 10.10.XX.XX/24. In this example, 24 indicates that the prefix of each IP address in the IP address whitelist is 24 bits in length. You can replace 24 with a value within the range of 1 to 32.
	//
	// If this parameter is not specified, the default IP address whitelist is used.
	//
	// example:
	//
	// 172.16.1.14,172.16.1.13,172.16.1.44,172.16.1.43,172.16.1.74,172.16.1.73
	SecurityIPList *string `json:"SecurityIPList,omitempty" xml:"SecurityIPList,omitempty"`
	// The vSwitch ID.
	//
	// >  The vSwitch must belong to the same zone as the instance.
	//
	// example:
	//
	// vsw-2zelwi1jd271p670lzl8h
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The ID of the virtual private cloud (VPC).
	//
	// example:
	//
	// vpc-wz9rbibex7v0lxbeyo6at
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DescribeResourceDetailsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeResourceDetailsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeResourceDetailsResponseBody) GetBackupDataSize() *int64 {
	return s.BackupDataSize
}

func (s *DescribeResourceDetailsResponseBody) GetBackupLogSize() *int64 {
	return s.BackupLogSize
}

func (s *DescribeResourceDetailsResponseBody) GetBackupSize() *int64 {
	return s.BackupSize
}

func (s *DescribeResourceDetailsResponseBody) GetDbInstanceStorage() *int64 {
	return s.DbInstanceStorage
}

func (s *DescribeResourceDetailsResponseBody) GetDbProxyInstanceName() *string {
	return s.DbProxyInstanceName
}

func (s *DescribeResourceDetailsResponseBody) GetDiskUsed() *int64 {
	return s.DiskUsed
}

func (s *DescribeResourceDetailsResponseBody) GetInstanceStorageType() *string {
	return s.InstanceStorageType
}

func (s *DescribeResourceDetailsResponseBody) GetRdsEcsSecurityGroupRel() []*DescribeResourceDetailsResponseBodyRdsEcsSecurityGroupRel {
	return s.RdsEcsSecurityGroupRel
}

func (s *DescribeResourceDetailsResponseBody) GetRegion() *string {
	return s.Region
}

func (s *DescribeResourceDetailsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeResourceDetailsResponseBody) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeResourceDetailsResponseBody) GetSecurityIPList() *string {
	return s.SecurityIPList
}

func (s *DescribeResourceDetailsResponseBody) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribeResourceDetailsResponseBody) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeResourceDetailsResponseBody) SetBackupDataSize(v int64) *DescribeResourceDetailsResponseBody {
	s.BackupDataSize = &v
	return s
}

func (s *DescribeResourceDetailsResponseBody) SetBackupLogSize(v int64) *DescribeResourceDetailsResponseBody {
	s.BackupLogSize = &v
	return s
}

func (s *DescribeResourceDetailsResponseBody) SetBackupSize(v int64) *DescribeResourceDetailsResponseBody {
	s.BackupSize = &v
	return s
}

func (s *DescribeResourceDetailsResponseBody) SetDbInstanceStorage(v int64) *DescribeResourceDetailsResponseBody {
	s.DbInstanceStorage = &v
	return s
}

func (s *DescribeResourceDetailsResponseBody) SetDbProxyInstanceName(v string) *DescribeResourceDetailsResponseBody {
	s.DbProxyInstanceName = &v
	return s
}

func (s *DescribeResourceDetailsResponseBody) SetDiskUsed(v int64) *DescribeResourceDetailsResponseBody {
	s.DiskUsed = &v
	return s
}

func (s *DescribeResourceDetailsResponseBody) SetInstanceStorageType(v string) *DescribeResourceDetailsResponseBody {
	s.InstanceStorageType = &v
	return s
}

func (s *DescribeResourceDetailsResponseBody) SetRdsEcsSecurityGroupRel(v []*DescribeResourceDetailsResponseBodyRdsEcsSecurityGroupRel) *DescribeResourceDetailsResponseBody {
	s.RdsEcsSecurityGroupRel = v
	return s
}

func (s *DescribeResourceDetailsResponseBody) SetRegion(v string) *DescribeResourceDetailsResponseBody {
	s.Region = &v
	return s
}

func (s *DescribeResourceDetailsResponseBody) SetRequestId(v string) *DescribeResourceDetailsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeResourceDetailsResponseBody) SetResourceGroupId(v string) *DescribeResourceDetailsResponseBody {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeResourceDetailsResponseBody) SetSecurityIPList(v string) *DescribeResourceDetailsResponseBody {
	s.SecurityIPList = &v
	return s
}

func (s *DescribeResourceDetailsResponseBody) SetVSwitchId(v string) *DescribeResourceDetailsResponseBody {
	s.VSwitchId = &v
	return s
}

func (s *DescribeResourceDetailsResponseBody) SetVpcId(v string) *DescribeResourceDetailsResponseBody {
	s.VpcId = &v
	return s
}

func (s *DescribeResourceDetailsResponseBody) Validate() error {
	if s.RdsEcsSecurityGroupRel != nil {
		for _, item := range s.RdsEcsSecurityGroupRel {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeResourceDetailsResponseBodyRdsEcsSecurityGroupRel struct {
	// The name of the security group.
	//
	// example:
	//
	// test_switch
	SecurityGroupName *string `json:"SecurityGroupName,omitempty" xml:"SecurityGroupName,omitempty"`
}

func (s DescribeResourceDetailsResponseBodyRdsEcsSecurityGroupRel) String() string {
	return dara.Prettify(s)
}

func (s DescribeResourceDetailsResponseBodyRdsEcsSecurityGroupRel) GoString() string {
	return s.String()
}

func (s *DescribeResourceDetailsResponseBodyRdsEcsSecurityGroupRel) GetSecurityGroupName() *string {
	return s.SecurityGroupName
}

func (s *DescribeResourceDetailsResponseBodyRdsEcsSecurityGroupRel) SetSecurityGroupName(v string) *DescribeResourceDetailsResponseBodyRdsEcsSecurityGroupRel {
	s.SecurityGroupName = &v
	return s
}

func (s *DescribeResourceDetailsResponseBodyRdsEcsSecurityGroupRel) Validate() error {
	return dara.Validate(s)
}
