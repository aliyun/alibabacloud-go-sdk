// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeIpsPrivateAssocResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetIpsPrivateAssoc(v []*DescribeIpsPrivateAssocResponseBodyIpsPrivateAssoc) *DescribeIpsPrivateAssocResponseBody
	GetIpsPrivateAssoc() []*DescribeIpsPrivateAssocResponseBodyIpsPrivateAssoc
	SetRequestId(v string) *DescribeIpsPrivateAssocResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *DescribeIpsPrivateAssocResponseBody
	GetTotalCount() *int64
	SetTotalOpenCount(v int64) *DescribeIpsPrivateAssocResponseBody
	GetTotalOpenCount() *int64
}

type DescribeIpsPrivateAssocResponseBody struct {
	IpsPrivateAssoc []*DescribeIpsPrivateAssocResponseBodyIpsPrivateAssoc `json:"IpsPrivateAssoc,omitempty" xml:"IpsPrivateAssoc,omitempty" type:"Repeated"`
	// example:
	//
	// B2841452-CB8D-4F7D-B247-38E1CF7334F8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 6
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// example:
	//
	// 0
	TotalOpenCount *int64 `json:"TotalOpenCount,omitempty" xml:"TotalOpenCount,omitempty"`
}

func (s DescribeIpsPrivateAssocResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeIpsPrivateAssocResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeIpsPrivateAssocResponseBody) GetIpsPrivateAssoc() []*DescribeIpsPrivateAssocResponseBodyIpsPrivateAssoc {
	return s.IpsPrivateAssoc
}

func (s *DescribeIpsPrivateAssocResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeIpsPrivateAssocResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeIpsPrivateAssocResponseBody) GetTotalOpenCount() *int64 {
	return s.TotalOpenCount
}

func (s *DescribeIpsPrivateAssocResponseBody) SetIpsPrivateAssoc(v []*DescribeIpsPrivateAssocResponseBodyIpsPrivateAssoc) *DescribeIpsPrivateAssocResponseBody {
	s.IpsPrivateAssoc = v
	return s
}

func (s *DescribeIpsPrivateAssocResponseBody) SetRequestId(v string) *DescribeIpsPrivateAssocResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeIpsPrivateAssocResponseBody) SetTotalCount(v int64) *DescribeIpsPrivateAssocResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeIpsPrivateAssocResponseBody) SetTotalOpenCount(v int64) *DescribeIpsPrivateAssocResponseBody {
	s.TotalOpenCount = &v
	return s
}

func (s *DescribeIpsPrivateAssocResponseBody) Validate() error {
	if s.IpsPrivateAssoc != nil {
		for _, item := range s.IpsPrivateAssoc {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeIpsPrivateAssocResponseBodyIpsPrivateAssoc struct {
	// example:
	//
	// close
	AssocInfoStatus *string `json:"AssocInfoStatus,omitempty" xml:"AssocInfoStatus,omitempty"`
	// example:
	//
	// aliuid:1096080848305847 assumeOssRole not exist,serviceName:aliyunesarealtimelogpushossrole
	ErrorMsg *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// example:
	//
	// 1507956024994407
	MemberUid       *int64    `json:"MemberUid,omitempty" xml:"MemberUid,omitempty"`
	ProtectedIpList []*string `json:"ProtectedIpList,omitempty" xml:"ProtectedIpList,omitempty" type:"Repeated"`
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// cba48ec510bb640559c6f5161cde58014
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// example:
	//
	// job-0000000061279FB000001BBB31F9D673
	ResourceName *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	// example:
	//
	// closed
	Status            *string   `json:"Status,omitempty" xml:"Status,omitempty"`
	UnprotectedIpList []*string `json:"UnprotectedIpList,omitempty" xml:"UnprotectedIpList,omitempty" type:"Repeated"`
	// example:
	//
	// vpc-wz92dxepli2pgnut796tf
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// example:
	//
	// vpc-uf62vdtifj7kffpxrydqd
	VpcName *string `json:"VpcName,omitempty" xml:"VpcName,omitempty"`
}

func (s DescribeIpsPrivateAssocResponseBodyIpsPrivateAssoc) String() string {
	return dara.Prettify(s)
}

func (s DescribeIpsPrivateAssocResponseBodyIpsPrivateAssoc) GoString() string {
	return s.String()
}

func (s *DescribeIpsPrivateAssocResponseBodyIpsPrivateAssoc) GetAssocInfoStatus() *string {
	return s.AssocInfoStatus
}

func (s *DescribeIpsPrivateAssocResponseBodyIpsPrivateAssoc) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *DescribeIpsPrivateAssocResponseBodyIpsPrivateAssoc) GetMemberUid() *int64 {
	return s.MemberUid
}

func (s *DescribeIpsPrivateAssocResponseBodyIpsPrivateAssoc) GetProtectedIpList() []*string {
	return s.ProtectedIpList
}

func (s *DescribeIpsPrivateAssocResponseBodyIpsPrivateAssoc) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeIpsPrivateAssocResponseBodyIpsPrivateAssoc) GetResourceId() *string {
	return s.ResourceId
}

func (s *DescribeIpsPrivateAssocResponseBodyIpsPrivateAssoc) GetResourceName() *string {
	return s.ResourceName
}

func (s *DescribeIpsPrivateAssocResponseBodyIpsPrivateAssoc) GetStatus() *string {
	return s.Status
}

func (s *DescribeIpsPrivateAssocResponseBodyIpsPrivateAssoc) GetUnprotectedIpList() []*string {
	return s.UnprotectedIpList
}

func (s *DescribeIpsPrivateAssocResponseBodyIpsPrivateAssoc) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeIpsPrivateAssocResponseBodyIpsPrivateAssoc) GetVpcName() *string {
	return s.VpcName
}

func (s *DescribeIpsPrivateAssocResponseBodyIpsPrivateAssoc) SetAssocInfoStatus(v string) *DescribeIpsPrivateAssocResponseBodyIpsPrivateAssoc {
	s.AssocInfoStatus = &v
	return s
}

func (s *DescribeIpsPrivateAssocResponseBodyIpsPrivateAssoc) SetErrorMsg(v string) *DescribeIpsPrivateAssocResponseBodyIpsPrivateAssoc {
	s.ErrorMsg = &v
	return s
}

func (s *DescribeIpsPrivateAssocResponseBodyIpsPrivateAssoc) SetMemberUid(v int64) *DescribeIpsPrivateAssocResponseBodyIpsPrivateAssoc {
	s.MemberUid = &v
	return s
}

func (s *DescribeIpsPrivateAssocResponseBodyIpsPrivateAssoc) SetProtectedIpList(v []*string) *DescribeIpsPrivateAssocResponseBodyIpsPrivateAssoc {
	s.ProtectedIpList = v
	return s
}

func (s *DescribeIpsPrivateAssocResponseBodyIpsPrivateAssoc) SetRegionId(v string) *DescribeIpsPrivateAssocResponseBodyIpsPrivateAssoc {
	s.RegionId = &v
	return s
}

func (s *DescribeIpsPrivateAssocResponseBodyIpsPrivateAssoc) SetResourceId(v string) *DescribeIpsPrivateAssocResponseBodyIpsPrivateAssoc {
	s.ResourceId = &v
	return s
}

func (s *DescribeIpsPrivateAssocResponseBodyIpsPrivateAssoc) SetResourceName(v string) *DescribeIpsPrivateAssocResponseBodyIpsPrivateAssoc {
	s.ResourceName = &v
	return s
}

func (s *DescribeIpsPrivateAssocResponseBodyIpsPrivateAssoc) SetStatus(v string) *DescribeIpsPrivateAssocResponseBodyIpsPrivateAssoc {
	s.Status = &v
	return s
}

func (s *DescribeIpsPrivateAssocResponseBodyIpsPrivateAssoc) SetUnprotectedIpList(v []*string) *DescribeIpsPrivateAssocResponseBodyIpsPrivateAssoc {
	s.UnprotectedIpList = v
	return s
}

func (s *DescribeIpsPrivateAssocResponseBodyIpsPrivateAssoc) SetVpcId(v string) *DescribeIpsPrivateAssocResponseBodyIpsPrivateAssoc {
	s.VpcId = &v
	return s
}

func (s *DescribeIpsPrivateAssocResponseBodyIpsPrivateAssoc) SetVpcName(v string) *DescribeIpsPrivateAssocResponseBodyIpsPrivateAssoc {
	s.VpcName = &v
	return s
}

func (s *DescribeIpsPrivateAssocResponseBodyIpsPrivateAssoc) Validate() error {
	return dara.Validate(s)
}
