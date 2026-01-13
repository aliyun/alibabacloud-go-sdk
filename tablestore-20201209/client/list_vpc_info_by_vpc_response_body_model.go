// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVpcInfoByVpcResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNum(v int64) *ListVpcInfoByVpcResponseBody
	GetPageNum() *int64
	SetPageSize(v int64) *ListVpcInfoByVpcResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *ListVpcInfoByVpcResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListVpcInfoByVpcResponseBody
	GetTotalCount() *int64
	SetVpcId(v string) *ListVpcInfoByVpcResponseBody
	GetVpcId() *string
	SetVpcInfos(v []*ListVpcInfoByVpcResponseBodyVpcInfos) *ListVpcInfoByVpcResponseBody
	GetVpcInfos() []*ListVpcInfoByVpcResponseBodyVpcInfos
}

type ListVpcInfoByVpcResponseBody struct {
	// example:
	//
	// 1
	PageNum *int64 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 4141784E-91FF-56E3-9371-FD011FF876F4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// VPC ID
	//
	// example:
	//
	// vpc-bp1***********0mh8
	VpcId    *string                                 `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	VpcInfos []*ListVpcInfoByVpcResponseBodyVpcInfos `json:"VpcInfos,omitempty" xml:"VpcInfos,omitempty" type:"Repeated"`
}

func (s ListVpcInfoByVpcResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListVpcInfoByVpcResponseBody) GoString() string {
	return s.String()
}

func (s *ListVpcInfoByVpcResponseBody) GetPageNum() *int64 {
	return s.PageNum
}

func (s *ListVpcInfoByVpcResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListVpcInfoByVpcResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListVpcInfoByVpcResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListVpcInfoByVpcResponseBody) GetVpcId() *string {
	return s.VpcId
}

func (s *ListVpcInfoByVpcResponseBody) GetVpcInfos() []*ListVpcInfoByVpcResponseBodyVpcInfos {
	return s.VpcInfos
}

func (s *ListVpcInfoByVpcResponseBody) SetPageNum(v int64) *ListVpcInfoByVpcResponseBody {
	s.PageNum = &v
	return s
}

func (s *ListVpcInfoByVpcResponseBody) SetPageSize(v int64) *ListVpcInfoByVpcResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListVpcInfoByVpcResponseBody) SetRequestId(v string) *ListVpcInfoByVpcResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListVpcInfoByVpcResponseBody) SetTotalCount(v int64) *ListVpcInfoByVpcResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListVpcInfoByVpcResponseBody) SetVpcId(v string) *ListVpcInfoByVpcResponseBody {
	s.VpcId = &v
	return s
}

func (s *ListVpcInfoByVpcResponseBody) SetVpcInfos(v []*ListVpcInfoByVpcResponseBodyVpcInfos) *ListVpcInfoByVpcResponseBody {
	s.VpcInfos = v
	return s
}

func (s *ListVpcInfoByVpcResponseBody) Validate() error {
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

type ListVpcInfoByVpcResponseBodyVpcInfos struct {
	// example:
	//
	// http://xu6666-mkdd-test.cn-hangzhou.vpc.ots.aliyuncs.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// example:
	//
	// http://172.**.***.34:80/
	Endpoint *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	// example:
	//
	// mkdd-test
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// example:
	//
	// xu6666
	InstanceVpcName *string `json:"InstanceVpcName,omitempty" xml:"InstanceVpcName,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionNo *string `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
}

func (s ListVpcInfoByVpcResponseBodyVpcInfos) String() string {
	return dara.Prettify(s)
}

func (s ListVpcInfoByVpcResponseBodyVpcInfos) GoString() string {
	return s.String()
}

func (s *ListVpcInfoByVpcResponseBodyVpcInfos) GetDomain() *string {
	return s.Domain
}

func (s *ListVpcInfoByVpcResponseBodyVpcInfos) GetEndpoint() *string {
	return s.Endpoint
}

func (s *ListVpcInfoByVpcResponseBodyVpcInfos) GetInstanceName() *string {
	return s.InstanceName
}

func (s *ListVpcInfoByVpcResponseBodyVpcInfos) GetInstanceVpcName() *string {
	return s.InstanceVpcName
}

func (s *ListVpcInfoByVpcResponseBodyVpcInfos) GetRegionNo() *string {
	return s.RegionNo
}

func (s *ListVpcInfoByVpcResponseBodyVpcInfos) SetDomain(v string) *ListVpcInfoByVpcResponseBodyVpcInfos {
	s.Domain = &v
	return s
}

func (s *ListVpcInfoByVpcResponseBodyVpcInfos) SetEndpoint(v string) *ListVpcInfoByVpcResponseBodyVpcInfos {
	s.Endpoint = &v
	return s
}

func (s *ListVpcInfoByVpcResponseBodyVpcInfos) SetInstanceName(v string) *ListVpcInfoByVpcResponseBodyVpcInfos {
	s.InstanceName = &v
	return s
}

func (s *ListVpcInfoByVpcResponseBodyVpcInfos) SetInstanceVpcName(v string) *ListVpcInfoByVpcResponseBodyVpcInfos {
	s.InstanceVpcName = &v
	return s
}

func (s *ListVpcInfoByVpcResponseBodyVpcInfos) SetRegionNo(v string) *ListVpcInfoByVpcResponseBodyVpcInfos {
	s.RegionNo = &v
	return s
}

func (s *ListVpcInfoByVpcResponseBodyVpcInfos) Validate() error {
	return dara.Validate(s)
}
