// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVfwIPSConfigListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNo(v int32) *DescribeVfwIPSConfigListResponseBody
	GetPageNo() *int32
	SetPageSize(v int32) *DescribeVfwIPSConfigListResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeVfwIPSConfigListResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeVfwIPSConfigListResponseBody
	GetTotalCount() *int32
	SetVfwIpsSwitchConfigList(v []*DescribeVfwIPSConfigListResponseBodyVfwIpsSwitchConfigList) *DescribeVfwIPSConfigListResponseBody
	GetVfwIpsSwitchConfigList() []*DescribeVfwIPSConfigListResponseBodyVfwIpsSwitchConfigList
}

type DescribeVfwIPSConfigListResponseBody struct {
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 15FCCC52-1E23-57AE-B5EF-3E00A3******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 2
	TotalCount             *int32                                                        `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	VfwIpsSwitchConfigList []*DescribeVfwIPSConfigListResponseBodyVfwIpsSwitchConfigList `json:"VfwIpsSwitchConfigList,omitempty" xml:"VfwIpsSwitchConfigList,omitempty" type:"Repeated"`
}

func (s DescribeVfwIPSConfigListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVfwIPSConfigListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVfwIPSConfigListResponseBody) GetPageNo() *int32 {
	return s.PageNo
}

func (s *DescribeVfwIPSConfigListResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeVfwIPSConfigListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVfwIPSConfigListResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeVfwIPSConfigListResponseBody) GetVfwIpsSwitchConfigList() []*DescribeVfwIPSConfigListResponseBodyVfwIpsSwitchConfigList {
	return s.VfwIpsSwitchConfigList
}

func (s *DescribeVfwIPSConfigListResponseBody) SetPageNo(v int32) *DescribeVfwIPSConfigListResponseBody {
	s.PageNo = &v
	return s
}

func (s *DescribeVfwIPSConfigListResponseBody) SetPageSize(v int32) *DescribeVfwIPSConfigListResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeVfwIPSConfigListResponseBody) SetRequestId(v string) *DescribeVfwIPSConfigListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVfwIPSConfigListResponseBody) SetTotalCount(v int32) *DescribeVfwIPSConfigListResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeVfwIPSConfigListResponseBody) SetVfwIpsSwitchConfigList(v []*DescribeVfwIPSConfigListResponseBodyVfwIpsSwitchConfigList) *DescribeVfwIPSConfigListResponseBody {
	s.VfwIpsSwitchConfigList = v
	return s
}

func (s *DescribeVfwIPSConfigListResponseBody) Validate() error {
	if s.VfwIpsSwitchConfigList != nil {
		for _, item := range s.VfwIpsSwitchConfigList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeVfwIPSConfigListResponseBodyVfwIpsSwitchConfigList struct {
	// example:
	//
	// 1
	BasicRules *int32 `json:"BasicRules,omitempty" xml:"BasicRules,omitempty"`
	// example:
	//
	// 134646920647****
	MemberUid *string `json:"MemberUid,omitempty" xml:"MemberUid,omitempty"`
	// example:
	//
	// 1
	PatchRules *int32 `json:"PatchRules,omitempty" xml:"PatchRules,omitempty"`
	// example:
	//
	// 1
	RuleClass *int32 `json:"RuleClass,omitempty" xml:"RuleClass,omitempty"`
	// example:
	//
	// 1
	RunMode *int32 `json:"RunMode,omitempty" xml:"RunMode,omitempty"`
	// example:
	//
	// cen-h678sl4wv3yd5v****
	VpcFirewallId     *string   `json:"VpcFirewallId,omitempty" xml:"VpcFirewallId,omitempty"`
	VpcFirewallIdList []*string `json:"VpcFirewallIdList,omitempty" xml:"VpcFirewallIdList,omitempty" type:"Repeated"`
	// example:
	//
	// vpc-test
	VpcFirewallName *string `json:"VpcFirewallName,omitempty" xml:"VpcFirewallName,omitempty"`
}

func (s DescribeVfwIPSConfigListResponseBodyVfwIpsSwitchConfigList) String() string {
	return dara.Prettify(s)
}

func (s DescribeVfwIPSConfigListResponseBodyVfwIpsSwitchConfigList) GoString() string {
	return s.String()
}

func (s *DescribeVfwIPSConfigListResponseBodyVfwIpsSwitchConfigList) GetBasicRules() *int32 {
	return s.BasicRules
}

func (s *DescribeVfwIPSConfigListResponseBodyVfwIpsSwitchConfigList) GetMemberUid() *string {
	return s.MemberUid
}

func (s *DescribeVfwIPSConfigListResponseBodyVfwIpsSwitchConfigList) GetPatchRules() *int32 {
	return s.PatchRules
}

func (s *DescribeVfwIPSConfigListResponseBodyVfwIpsSwitchConfigList) GetRuleClass() *int32 {
	return s.RuleClass
}

func (s *DescribeVfwIPSConfigListResponseBodyVfwIpsSwitchConfigList) GetRunMode() *int32 {
	return s.RunMode
}

func (s *DescribeVfwIPSConfigListResponseBodyVfwIpsSwitchConfigList) GetVpcFirewallId() *string {
	return s.VpcFirewallId
}

func (s *DescribeVfwIPSConfigListResponseBodyVfwIpsSwitchConfigList) GetVpcFirewallIdList() []*string {
	return s.VpcFirewallIdList
}

func (s *DescribeVfwIPSConfigListResponseBodyVfwIpsSwitchConfigList) GetVpcFirewallName() *string {
	return s.VpcFirewallName
}

func (s *DescribeVfwIPSConfigListResponseBodyVfwIpsSwitchConfigList) SetBasicRules(v int32) *DescribeVfwIPSConfigListResponseBodyVfwIpsSwitchConfigList {
	s.BasicRules = &v
	return s
}

func (s *DescribeVfwIPSConfigListResponseBodyVfwIpsSwitchConfigList) SetMemberUid(v string) *DescribeVfwIPSConfigListResponseBodyVfwIpsSwitchConfigList {
	s.MemberUid = &v
	return s
}

func (s *DescribeVfwIPSConfigListResponseBodyVfwIpsSwitchConfigList) SetPatchRules(v int32) *DescribeVfwIPSConfigListResponseBodyVfwIpsSwitchConfigList {
	s.PatchRules = &v
	return s
}

func (s *DescribeVfwIPSConfigListResponseBodyVfwIpsSwitchConfigList) SetRuleClass(v int32) *DescribeVfwIPSConfigListResponseBodyVfwIpsSwitchConfigList {
	s.RuleClass = &v
	return s
}

func (s *DescribeVfwIPSConfigListResponseBodyVfwIpsSwitchConfigList) SetRunMode(v int32) *DescribeVfwIPSConfigListResponseBodyVfwIpsSwitchConfigList {
	s.RunMode = &v
	return s
}

func (s *DescribeVfwIPSConfigListResponseBodyVfwIpsSwitchConfigList) SetVpcFirewallId(v string) *DescribeVfwIPSConfigListResponseBodyVfwIpsSwitchConfigList {
	s.VpcFirewallId = &v
	return s
}

func (s *DescribeVfwIPSConfigListResponseBodyVfwIpsSwitchConfigList) SetVpcFirewallIdList(v []*string) *DescribeVfwIPSConfigListResponseBodyVfwIpsSwitchConfigList {
	s.VpcFirewallIdList = v
	return s
}

func (s *DescribeVfwIPSConfigListResponseBodyVfwIpsSwitchConfigList) SetVpcFirewallName(v string) *DescribeVfwIPSConfigListResponseBodyVfwIpsSwitchConfigList {
	s.VpcFirewallName = &v
	return s
}

func (s *DescribeVfwIPSConfigListResponseBodyVfwIpsSwitchConfigList) Validate() error {
	return dara.Validate(s)
}
