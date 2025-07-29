// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeClusterVulsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetVulRecords(v []*DescribeClusterVulsResponseBodyVulRecords) *DescribeClusterVulsResponseBody
	GetVulRecords() []*DescribeClusterVulsResponseBodyVulRecords
}

type DescribeClusterVulsResponseBody struct {
	// The list of vulnerabilities.
	VulRecords []*DescribeClusterVulsResponseBodyVulRecords `json:"vul_records,omitempty" xml:"vul_records,omitempty" type:"Repeated"`
}

func (s DescribeClusterVulsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterVulsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeClusterVulsResponseBody) GetVulRecords() []*DescribeClusterVulsResponseBodyVulRecords {
	return s.VulRecords
}

func (s *DescribeClusterVulsResponseBody) SetVulRecords(v []*DescribeClusterVulsResponseBodyVulRecords) *DescribeClusterVulsResponseBody {
	s.VulRecords = v
	return s
}

func (s *DescribeClusterVulsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeClusterVulsResponseBodyVulRecords struct {
	// The CVE list.
	CveList []*string `json:"cve_list,omitempty" xml:"cve_list,omitempty" type:"Repeated"`
	// The severity level of the vulnerability.
	//
	// Valid values:
	//
	// 	- nntf: low
	//
	// 	- later: medium
	//
	// 	- asap: high
	//
	// example:
	//
	// asap
	Necessity *string `json:"necessity,omitempty" xml:"necessity,omitempty"`
	// The number of nodes that have the vulnerability.
	//
	// example:
	//
	// 1
	NodeCount *int32 `json:"node_count,omitempty" xml:"node_count,omitempty"`
	// The node pool ID.
	//
	// example:
	//
	// np0156da1082b54fa987e32618dd45f5d3
	NodepoolId *string `json:"nodepool_id,omitempty" xml:"nodepool_id,omitempty"`
	// The name of the node pool.
	//
	// example:
	//
	// test
	NodepoolName *string `json:"nodepool_name,omitempty" xml:"nodepool_name,omitempty"`
	// The alias of the vulnerability.
	//
	// example:
	//
	// CVE-2022-xxxx:rsync Security vulnerabilities
	VulAliasName *string `json:"vul_alias_name,omitempty" xml:"vul_alias_name,omitempty"`
	// The name of the vulnerability.
	//
	// example:
	//
	// oval:com.redhat.rhsa:def:xxxxxxx
	VulName *string `json:"vul_name,omitempty" xml:"vul_name,omitempty"`
	// The type of vulnerability.
	//
	// Valid values:
	//
	// 	- app: application vulnerabilities
	//
	// 	- sca: application vulnerabilities (software component analysis)
	//
	// 	- cve: Linux vulnerabilities
	//
	// 	- cms: Web-CMS vulnerabilities
	//
	// 	- sys: Windows vulnerabilities
	//
	// 	- emg:  emergency vulnerabilities
	//
	// example:
	//
	// cve
	VulType *string `json:"vul_type,omitempty" xml:"vul_type,omitempty"`
}

func (s DescribeClusterVulsResponseBodyVulRecords) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterVulsResponseBodyVulRecords) GoString() string {
	return s.String()
}

func (s *DescribeClusterVulsResponseBodyVulRecords) GetCveList() []*string {
	return s.CveList
}

func (s *DescribeClusterVulsResponseBodyVulRecords) GetNecessity() *string {
	return s.Necessity
}

func (s *DescribeClusterVulsResponseBodyVulRecords) GetNodeCount() *int32 {
	return s.NodeCount
}

func (s *DescribeClusterVulsResponseBodyVulRecords) GetNodepoolId() *string {
	return s.NodepoolId
}

func (s *DescribeClusterVulsResponseBodyVulRecords) GetNodepoolName() *string {
	return s.NodepoolName
}

func (s *DescribeClusterVulsResponseBodyVulRecords) GetVulAliasName() *string {
	return s.VulAliasName
}

func (s *DescribeClusterVulsResponseBodyVulRecords) GetVulName() *string {
	return s.VulName
}

func (s *DescribeClusterVulsResponseBodyVulRecords) GetVulType() *string {
	return s.VulType
}

func (s *DescribeClusterVulsResponseBodyVulRecords) SetCveList(v []*string) *DescribeClusterVulsResponseBodyVulRecords {
	s.CveList = v
	return s
}

func (s *DescribeClusterVulsResponseBodyVulRecords) SetNecessity(v string) *DescribeClusterVulsResponseBodyVulRecords {
	s.Necessity = &v
	return s
}

func (s *DescribeClusterVulsResponseBodyVulRecords) SetNodeCount(v int32) *DescribeClusterVulsResponseBodyVulRecords {
	s.NodeCount = &v
	return s
}

func (s *DescribeClusterVulsResponseBodyVulRecords) SetNodepoolId(v string) *DescribeClusterVulsResponseBodyVulRecords {
	s.NodepoolId = &v
	return s
}

func (s *DescribeClusterVulsResponseBodyVulRecords) SetNodepoolName(v string) *DescribeClusterVulsResponseBodyVulRecords {
	s.NodepoolName = &v
	return s
}

func (s *DescribeClusterVulsResponseBodyVulRecords) SetVulAliasName(v string) *DescribeClusterVulsResponseBodyVulRecords {
	s.VulAliasName = &v
	return s
}

func (s *DescribeClusterVulsResponseBodyVulRecords) SetVulName(v string) *DescribeClusterVulsResponseBodyVulRecords {
	s.VulName = &v
	return s
}

func (s *DescribeClusterVulsResponseBodyVulRecords) SetVulType(v string) *DescribeClusterVulsResponseBodyVulRecords {
	s.VulType = &v
	return s
}

func (s *DescribeClusterVulsResponseBodyVulRecords) Validate() error {
	return dara.Validate(s)
}
