// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVpcInfoByAgentStorageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAgentStorageName(v string) *ListVpcInfoByAgentStorageResponseBody
	GetAgentStorageName() *string
	SetPageNum(v int64) *ListVpcInfoByAgentStorageResponseBody
	GetPageNum() *int64
	SetPageSize(v int64) *ListVpcInfoByAgentStorageResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *ListVpcInfoByAgentStorageResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListVpcInfoByAgentStorageResponseBody
	GetTotalCount() *int64
	SetVpcInfos(v []*ListVpcInfoByAgentStorageResponseBodyVpcInfos) *ListVpcInfoByAgentStorageResponseBody
	GetVpcInfos() []*ListVpcInfoByAgentStorageResponseBodyVpcInfos
}

type ListVpcInfoByAgentStorageResponseBody struct {
	// The agent storage name.
	//
	// example:
	//
	// agent-test
	AgentStorageName *string `json:"AgentStorageName,omitempty" xml:"AgentStorageName,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNum *int64 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 8
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID, which can be used for troubleshooting.
	//
	// example:
	//
	// 39871ED2-62C0-578F-A32E-B88072D5582F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of VPCs.
	//
	// example:
	//
	// 1
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The VPC details.
	VpcInfos []*ListVpcInfoByAgentStorageResponseBodyVpcInfos `json:"VpcInfos,omitempty" xml:"VpcInfos,omitempty" type:"Repeated"`
}

func (s ListVpcInfoByAgentStorageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListVpcInfoByAgentStorageResponseBody) GoString() string {
	return s.String()
}

func (s *ListVpcInfoByAgentStorageResponseBody) GetAgentStorageName() *string {
	return s.AgentStorageName
}

func (s *ListVpcInfoByAgentStorageResponseBody) GetPageNum() *int64 {
	return s.PageNum
}

func (s *ListVpcInfoByAgentStorageResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListVpcInfoByAgentStorageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListVpcInfoByAgentStorageResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListVpcInfoByAgentStorageResponseBody) GetVpcInfos() []*ListVpcInfoByAgentStorageResponseBodyVpcInfos {
	return s.VpcInfos
}

func (s *ListVpcInfoByAgentStorageResponseBody) SetAgentStorageName(v string) *ListVpcInfoByAgentStorageResponseBody {
	s.AgentStorageName = &v
	return s
}

func (s *ListVpcInfoByAgentStorageResponseBody) SetPageNum(v int64) *ListVpcInfoByAgentStorageResponseBody {
	s.PageNum = &v
	return s
}

func (s *ListVpcInfoByAgentStorageResponseBody) SetPageSize(v int64) *ListVpcInfoByAgentStorageResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListVpcInfoByAgentStorageResponseBody) SetRequestId(v string) *ListVpcInfoByAgentStorageResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListVpcInfoByAgentStorageResponseBody) SetTotalCount(v int64) *ListVpcInfoByAgentStorageResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListVpcInfoByAgentStorageResponseBody) SetVpcInfos(v []*ListVpcInfoByAgentStorageResponseBodyVpcInfos) *ListVpcInfoByAgentStorageResponseBody {
	s.VpcInfos = v
	return s
}

func (s *ListVpcInfoByAgentStorageResponseBody) Validate() error {
	if s.VpcInfos != nil {
		for _, item := range s.VpcInfos {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListVpcInfoByAgentStorageResponseBodyVpcInfos struct {
	// The VPC name.
	//
	// example:
	//
	// remua
	AgentStorageVpcName *string `json:"AgentStorageVpcName,omitempty" xml:"AgentStorageVpcName,omitempty"`
	// The VPC access address.
	//
	// example:
	//
	// http://remua-agent-test.cn-beijing.vpc.ots.aliyuncs.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The endpoint of the agent storage.
	//
	// example:
	//
	// http://172.**.***.34:80/
	Endpoint *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-beijing
	RegionNo *string `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
	// VPC ID
	//
	// example:
	//
	// vpc-2z***********n7w3dl
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s ListVpcInfoByAgentStorageResponseBodyVpcInfos) String() string {
	return dara.Prettify(s)
}

func (s ListVpcInfoByAgentStorageResponseBodyVpcInfos) GoString() string {
	return s.String()
}

func (s *ListVpcInfoByAgentStorageResponseBodyVpcInfos) GetAgentStorageVpcName() *string {
	return s.AgentStorageVpcName
}

func (s *ListVpcInfoByAgentStorageResponseBodyVpcInfos) GetDomain() *string {
	return s.Domain
}

func (s *ListVpcInfoByAgentStorageResponseBodyVpcInfos) GetEndpoint() *string {
	return s.Endpoint
}

func (s *ListVpcInfoByAgentStorageResponseBodyVpcInfos) GetRegionNo() *string {
	return s.RegionNo
}

func (s *ListVpcInfoByAgentStorageResponseBodyVpcInfos) GetVpcId() *string {
	return s.VpcId
}

func (s *ListVpcInfoByAgentStorageResponseBodyVpcInfos) SetAgentStorageVpcName(v string) *ListVpcInfoByAgentStorageResponseBodyVpcInfos {
	s.AgentStorageVpcName = &v
	return s
}

func (s *ListVpcInfoByAgentStorageResponseBodyVpcInfos) SetDomain(v string) *ListVpcInfoByAgentStorageResponseBodyVpcInfos {
	s.Domain = &v
	return s
}

func (s *ListVpcInfoByAgentStorageResponseBodyVpcInfos) SetEndpoint(v string) *ListVpcInfoByAgentStorageResponseBodyVpcInfos {
	s.Endpoint = &v
	return s
}

func (s *ListVpcInfoByAgentStorageResponseBodyVpcInfos) SetRegionNo(v string) *ListVpcInfoByAgentStorageResponseBodyVpcInfos {
	s.RegionNo = &v
	return s
}

func (s *ListVpcInfoByAgentStorageResponseBodyVpcInfos) SetVpcId(v string) *ListVpcInfoByAgentStorageResponseBodyVpcInfos {
	s.VpcId = &v
	return s
}

func (s *ListVpcInfoByAgentStorageResponseBodyVpcInfos) Validate() error {
	return dara.Validate(s)
}
