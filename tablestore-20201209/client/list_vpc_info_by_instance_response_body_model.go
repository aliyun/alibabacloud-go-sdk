// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVpcInfoByInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceName(v string) *ListVpcInfoByInstanceResponseBody
	GetInstanceName() *string
	SetPageNum(v int64) *ListVpcInfoByInstanceResponseBody
	GetPageNum() *int64
	SetPageSize(v int64) *ListVpcInfoByInstanceResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *ListVpcInfoByInstanceResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListVpcInfoByInstanceResponseBody
	GetTotalCount() *int64
	SetVpcInfos(v []*ListVpcInfoByInstanceResponseBodyVpcInfos) *ListVpcInfoByInstanceResponseBody
	GetVpcInfos() []*ListVpcInfoByInstanceResponseBodyVpcInfos
}

type ListVpcInfoByInstanceResponseBody struct {
	// example:
	//
	// mkdd-test
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// example:
	//
	// 1
	PageNum *int64 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// example:
	//
	// 8
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 44FDC379-4443-560E-B652-9F7476D8854F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1
	TotalCount *int64                                       `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	VpcInfos   []*ListVpcInfoByInstanceResponseBodyVpcInfos `json:"VpcInfos,omitempty" xml:"VpcInfos,omitempty" type:"Repeated"`
}

func (s ListVpcInfoByInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListVpcInfoByInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *ListVpcInfoByInstanceResponseBody) GetInstanceName() *string {
	return s.InstanceName
}

func (s *ListVpcInfoByInstanceResponseBody) GetPageNum() *int64 {
	return s.PageNum
}

func (s *ListVpcInfoByInstanceResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListVpcInfoByInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListVpcInfoByInstanceResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListVpcInfoByInstanceResponseBody) GetVpcInfos() []*ListVpcInfoByInstanceResponseBodyVpcInfos {
	return s.VpcInfos
}

func (s *ListVpcInfoByInstanceResponseBody) SetInstanceName(v string) *ListVpcInfoByInstanceResponseBody {
	s.InstanceName = &v
	return s
}

func (s *ListVpcInfoByInstanceResponseBody) SetPageNum(v int64) *ListVpcInfoByInstanceResponseBody {
	s.PageNum = &v
	return s
}

func (s *ListVpcInfoByInstanceResponseBody) SetPageSize(v int64) *ListVpcInfoByInstanceResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListVpcInfoByInstanceResponseBody) SetRequestId(v string) *ListVpcInfoByInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListVpcInfoByInstanceResponseBody) SetTotalCount(v int64) *ListVpcInfoByInstanceResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListVpcInfoByInstanceResponseBody) SetVpcInfos(v []*ListVpcInfoByInstanceResponseBodyVpcInfos) *ListVpcInfoByInstanceResponseBody {
	s.VpcInfos = v
	return s
}

func (s *ListVpcInfoByInstanceResponseBody) Validate() error {
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

type ListVpcInfoByInstanceResponseBodyVpcInfos struct {
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
	// xu6666
	InstanceVpcName *string `json:"InstanceVpcName,omitempty" xml:"InstanceVpcName,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionNo *string `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
	// VPC ID
	//
	// example:
	//
	// vpc-2z***********n7w3dl
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s ListVpcInfoByInstanceResponseBodyVpcInfos) String() string {
	return dara.Prettify(s)
}

func (s ListVpcInfoByInstanceResponseBodyVpcInfos) GoString() string {
	return s.String()
}

func (s *ListVpcInfoByInstanceResponseBodyVpcInfos) GetDomain() *string {
	return s.Domain
}

func (s *ListVpcInfoByInstanceResponseBodyVpcInfos) GetEndpoint() *string {
	return s.Endpoint
}

func (s *ListVpcInfoByInstanceResponseBodyVpcInfos) GetInstanceVpcName() *string {
	return s.InstanceVpcName
}

func (s *ListVpcInfoByInstanceResponseBodyVpcInfos) GetRegionNo() *string {
	return s.RegionNo
}

func (s *ListVpcInfoByInstanceResponseBodyVpcInfos) GetVpcId() *string {
	return s.VpcId
}

func (s *ListVpcInfoByInstanceResponseBodyVpcInfos) SetDomain(v string) *ListVpcInfoByInstanceResponseBodyVpcInfos {
	s.Domain = &v
	return s
}

func (s *ListVpcInfoByInstanceResponseBodyVpcInfos) SetEndpoint(v string) *ListVpcInfoByInstanceResponseBodyVpcInfos {
	s.Endpoint = &v
	return s
}

func (s *ListVpcInfoByInstanceResponseBodyVpcInfos) SetInstanceVpcName(v string) *ListVpcInfoByInstanceResponseBodyVpcInfos {
	s.InstanceVpcName = &v
	return s
}

func (s *ListVpcInfoByInstanceResponseBodyVpcInfos) SetRegionNo(v string) *ListVpcInfoByInstanceResponseBodyVpcInfos {
	s.RegionNo = &v
	return s
}

func (s *ListVpcInfoByInstanceResponseBodyVpcInfos) SetVpcId(v string) *ListVpcInfoByInstanceResponseBodyVpcInfos {
	s.VpcId = &v
	return s
}

func (s *ListVpcInfoByInstanceResponseBodyVpcInfos) Validate() error {
	return dara.Validate(s)
}
