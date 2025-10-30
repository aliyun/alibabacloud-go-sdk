// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNodePoolVulsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetVulRecords(v []*DescribeNodePoolVulsResponseBodyVulRecords) *DescribeNodePoolVulsResponseBody
	GetVulRecords() []*DescribeNodePoolVulsResponseBodyVulRecords
	SetVulsFixServicePurchased(v bool) *DescribeNodePoolVulsResponseBody
	GetVulsFixServicePurchased() *bool
}

type DescribeNodePoolVulsResponseBody struct {
	// The vulnerability list of all node pools.
	VulRecords []*DescribeNodePoolVulsResponseBodyVulRecords `json:"vul_records,omitempty" xml:"vul_records,omitempty" type:"Repeated"`
	// Indicates whether the CVE vulnerability patching service provided by Security Center is purchased.
	//
	// 	- true: yes
	//
	// 	- false: no
	//
	// example:
	//
	// false
	VulsFixServicePurchased *bool `json:"vuls_fix_service_purchased,omitempty" xml:"vuls_fix_service_purchased,omitempty"`
}

func (s DescribeNodePoolVulsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeNodePoolVulsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeNodePoolVulsResponseBody) GetVulRecords() []*DescribeNodePoolVulsResponseBodyVulRecords {
	return s.VulRecords
}

func (s *DescribeNodePoolVulsResponseBody) GetVulsFixServicePurchased() *bool {
	return s.VulsFixServicePurchased
}

func (s *DescribeNodePoolVulsResponseBody) SetVulRecords(v []*DescribeNodePoolVulsResponseBodyVulRecords) *DescribeNodePoolVulsResponseBody {
	s.VulRecords = v
	return s
}

func (s *DescribeNodePoolVulsResponseBody) SetVulsFixServicePurchased(v bool) *DescribeNodePoolVulsResponseBody {
	s.VulsFixServicePurchased = &v
	return s
}

func (s *DescribeNodePoolVulsResponseBody) Validate() error {
	if s.VulRecords != nil {
		for _, item := range s.VulRecords {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeNodePoolVulsResponseBodyVulRecords struct {
	// The ID of the node.
	//
	// example:
	//
	// i-t4n2qolb0wtzt0pz****
	InstanceId *string `json:"instance_id,omitempty" xml:"instance_id,omitempty"`
	// The node name. This name is the identifier of the node in the cluster.
	//
	// example:
	//
	// cn-hangzhou.192.168.x.x
	NodeName *string `json:"node_name,omitempty" xml:"node_name,omitempty"`
	// The list of vulnerabilities.
	VulList []*DescribeNodePoolVulsResponseBodyVulRecordsVulList `json:"vul_list,omitempty" xml:"vul_list,omitempty" type:"Repeated"`
}

func (s DescribeNodePoolVulsResponseBodyVulRecords) String() string {
	return dara.Prettify(s)
}

func (s DescribeNodePoolVulsResponseBodyVulRecords) GoString() string {
	return s.String()
}

func (s *DescribeNodePoolVulsResponseBodyVulRecords) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeNodePoolVulsResponseBodyVulRecords) GetNodeName() *string {
	return s.NodeName
}

func (s *DescribeNodePoolVulsResponseBodyVulRecords) GetVulList() []*DescribeNodePoolVulsResponseBodyVulRecordsVulList {
	return s.VulList
}

func (s *DescribeNodePoolVulsResponseBodyVulRecords) SetInstanceId(v string) *DescribeNodePoolVulsResponseBodyVulRecords {
	s.InstanceId = &v
	return s
}

func (s *DescribeNodePoolVulsResponseBodyVulRecords) SetNodeName(v string) *DescribeNodePoolVulsResponseBodyVulRecords {
	s.NodeName = &v
	return s
}

func (s *DescribeNodePoolVulsResponseBodyVulRecords) SetVulList(v []*DescribeNodePoolVulsResponseBodyVulRecordsVulList) *DescribeNodePoolVulsResponseBodyVulRecords {
	s.VulList = v
	return s
}

func (s *DescribeNodePoolVulsResponseBodyVulRecords) Validate() error {
	if s.VulList != nil {
		for _, item := range s.VulList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeNodePoolVulsResponseBodyVulRecordsVulList struct {
	// The alias of the vulnerability.
	//
	// example:
	//
	// RHSA-2019:3197-Important: sudo security update
	AliasName *string `json:"alias_name,omitempty" xml:"alias_name,omitempty"`
	// A list of CVE names corresponding to the vulnerabilities.
	CveList []*string `json:"cve_list,omitempty" xml:"cve_list,omitempty" type:"Repeated"`
	// The name of the vulnerability.
	//
	// example:
	//
	// oval:com.redhat.rhsa:def:20193197
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The severity level of the vulnerability.
	//
	// Valid values:
	//
	// 	- nntf: You can ignore the vulnerability.
	//
	// 	- later: You can fix the vulnerability later.
	//
	// 	- asap: You need to fix the vulnerability at the earliest opportunity.
	//
	// example:
	//
	// asap
	Necessity *string `json:"necessity,omitempty" xml:"necessity,omitempty"`
	// Indicates whether a restart is required.
	//
	// example:
	//
	// false
	NeedReboot  *bool                                                           `json:"need_reboot,omitempty" xml:"need_reboot,omitempty"`
	PackageList []*DescribeNodePoolVulsResponseBodyVulRecordsVulListPackageList `json:"package_list,omitempty" xml:"package_list,omitempty" type:"Repeated"`
}

func (s DescribeNodePoolVulsResponseBodyVulRecordsVulList) String() string {
	return dara.Prettify(s)
}

func (s DescribeNodePoolVulsResponseBodyVulRecordsVulList) GoString() string {
	return s.String()
}

func (s *DescribeNodePoolVulsResponseBodyVulRecordsVulList) GetAliasName() *string {
	return s.AliasName
}

func (s *DescribeNodePoolVulsResponseBodyVulRecordsVulList) GetCveList() []*string {
	return s.CveList
}

func (s *DescribeNodePoolVulsResponseBodyVulRecordsVulList) GetName() *string {
	return s.Name
}

func (s *DescribeNodePoolVulsResponseBodyVulRecordsVulList) GetNecessity() *string {
	return s.Necessity
}

func (s *DescribeNodePoolVulsResponseBodyVulRecordsVulList) GetNeedReboot() *bool {
	return s.NeedReboot
}

func (s *DescribeNodePoolVulsResponseBodyVulRecordsVulList) GetPackageList() []*DescribeNodePoolVulsResponseBodyVulRecordsVulListPackageList {
	return s.PackageList
}

func (s *DescribeNodePoolVulsResponseBodyVulRecordsVulList) SetAliasName(v string) *DescribeNodePoolVulsResponseBodyVulRecordsVulList {
	s.AliasName = &v
	return s
}

func (s *DescribeNodePoolVulsResponseBodyVulRecordsVulList) SetCveList(v []*string) *DescribeNodePoolVulsResponseBodyVulRecordsVulList {
	s.CveList = v
	return s
}

func (s *DescribeNodePoolVulsResponseBodyVulRecordsVulList) SetName(v string) *DescribeNodePoolVulsResponseBodyVulRecordsVulList {
	s.Name = &v
	return s
}

func (s *DescribeNodePoolVulsResponseBodyVulRecordsVulList) SetNecessity(v string) *DescribeNodePoolVulsResponseBodyVulRecordsVulList {
	s.Necessity = &v
	return s
}

func (s *DescribeNodePoolVulsResponseBodyVulRecordsVulList) SetNeedReboot(v bool) *DescribeNodePoolVulsResponseBodyVulRecordsVulList {
	s.NeedReboot = &v
	return s
}

func (s *DescribeNodePoolVulsResponseBodyVulRecordsVulList) SetPackageList(v []*DescribeNodePoolVulsResponseBodyVulRecordsVulListPackageList) *DescribeNodePoolVulsResponseBodyVulRecordsVulList {
	s.PackageList = v
	return s
}

func (s *DescribeNodePoolVulsResponseBodyVulRecordsVulList) Validate() error {
	if s.PackageList != nil {
		for _, item := range s.PackageList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeNodePoolVulsResponseBodyVulRecordsVulListPackageList struct {
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s DescribeNodePoolVulsResponseBodyVulRecordsVulListPackageList) String() string {
	return dara.Prettify(s)
}

func (s DescribeNodePoolVulsResponseBodyVulRecordsVulListPackageList) GoString() string {
	return s.String()
}

func (s *DescribeNodePoolVulsResponseBodyVulRecordsVulListPackageList) GetName() *string {
	return s.Name
}

func (s *DescribeNodePoolVulsResponseBodyVulRecordsVulListPackageList) SetName(v string) *DescribeNodePoolVulsResponseBodyVulRecordsVulListPackageList {
	s.Name = &v
	return s
}

func (s *DescribeNodePoolVulsResponseBodyVulRecordsVulListPackageList) Validate() error {
	return dara.Validate(s)
}
