// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePackIpListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribePackIpListResponseBody
	GetCode() *string
	SetIpList(v []*DescribePackIpListResponseBodyIpList) *DescribePackIpListResponseBody
	GetIpList() []*DescribePackIpListResponseBodyIpList
	SetRequestId(v string) *DescribePackIpListResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribePackIpListResponseBody
	GetSuccess() *bool
	SetTotal(v int32) *DescribePackIpListResponseBody
	GetTotal() *int32
}

type DescribePackIpListResponseBody struct {
	// The HTTP status code of the request.
	//
	// For more information about status codes, see [Common parameters](https://help.aliyun.com/document_detail/118841.html).
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The list of IP addresses that are protected by the Anti-DDoS Origin instance.
	IpList []*DescribePackIpListResponseBodyIpList `json:"IpList,omitempty" xml:"IpList,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 4FD1578A-BD77-50B7-A969-45A374A7ED22
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// - **true**: The request was successful.
	//
	// - **false**: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The number of protected IP addresses that are returned.
	//
	// example:
	//
	// 1
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribePackIpListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePackIpListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePackIpListResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribePackIpListResponseBody) GetIpList() []*DescribePackIpListResponseBodyIpList {
	return s.IpList
}

func (s *DescribePackIpListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePackIpListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribePackIpListResponseBody) GetTotal() *int32 {
	return s.Total
}

func (s *DescribePackIpListResponseBody) SetCode(v string) *DescribePackIpListResponseBody {
	s.Code = &v
	return s
}

func (s *DescribePackIpListResponseBody) SetIpList(v []*DescribePackIpListResponseBodyIpList) *DescribePackIpListResponseBody {
	s.IpList = v
	return s
}

func (s *DescribePackIpListResponseBody) SetRequestId(v string) *DescribePackIpListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePackIpListResponseBody) SetSuccess(v bool) *DescribePackIpListResponseBody {
	s.Success = &v
	return s
}

func (s *DescribePackIpListResponseBody) SetTotal(v int32) *DescribePackIpListResponseBody {
	s.Total = &v
	return s
}

func (s *DescribePackIpListResponseBody) Validate() error {
	if s.IpList != nil {
		for _, item := range s.IpList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribePackIpListResponseBodyIpList struct {
	// The IP address.
	//
	// example:
	//
	// 47.98.XX.XX
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The UID of the member account.
	//
	// example:
	//
	// 170858869679****
	MemberUid *string `json:"MemberUid,omitempty" xml:"MemberUid,omitempty"`
	// The end time of cross-border traffic blocking.
	//
	// example:
	//
	// 1715658000
	NsmExpireAt *int64 `json:"NsmExpireAt,omitempty" xml:"NsmExpireAt,omitempty"`
	// The start time of cross-border traffic blocking.
	//
	// example:
	//
	// 1715655000
	NsmStartAt *int64 `json:"NsmStartAt,omitempty" xml:"NsmStartAt,omitempty"`
	// The status of cross-border traffic blocking. Valid values:
	//
	// - **1**: Cross-border traffic is being blocked.
	//
	// - **0**: Cross-border traffic is not blocked.
	//
	// example:
	//
	// 0
	NsmStatus *int32 `json:"NsmStatus,omitempty" xml:"NsmStatus,omitempty"`
	// The type of the cloud asset to which the IP address belongs. Valid values:
	//
	// - **ECS**: an ECS instance.
	//
	// - **SLB**: a CLB instance.
	//
	// - **EIP**: an EIP instance. This includes the EIP used by an ALB instance.
	//
	// - **WAF**: a WAF instance.
	//
	// example:
	//
	// ECS
	Product *string `json:"Product,omitempty" xml:"Product,omitempty"`
	// The region where the protected IP address is deployed.
	//
	// > This parameter is not returned if the protected IP address is deployed in the same region as the instance.
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The remarks on the cloud asset to which the IP address belongs, such as an ECS instance or an SLB instance.
	//
	// > This parameter is not returned if no remarks are specified for the cloud asset.
	//
	// example:
	//
	// test
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The current status of the IP address. Valid values:
	//
	// - **normal**: The IP address is not under attack.
	//
	// - **hole_begin**: The IP address is in blackhole filtering status.
	//
	// example:
	//
	// normal
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribePackIpListResponseBodyIpList) String() string {
	return dara.Prettify(s)
}

func (s DescribePackIpListResponseBodyIpList) GoString() string {
	return s.String()
}

func (s *DescribePackIpListResponseBodyIpList) GetIp() *string {
	return s.Ip
}

func (s *DescribePackIpListResponseBodyIpList) GetMemberUid() *string {
	return s.MemberUid
}

func (s *DescribePackIpListResponseBodyIpList) GetNsmExpireAt() *int64 {
	return s.NsmExpireAt
}

func (s *DescribePackIpListResponseBodyIpList) GetNsmStartAt() *int64 {
	return s.NsmStartAt
}

func (s *DescribePackIpListResponseBodyIpList) GetNsmStatus() *int32 {
	return s.NsmStatus
}

func (s *DescribePackIpListResponseBodyIpList) GetProduct() *string {
	return s.Product
}

func (s *DescribePackIpListResponseBodyIpList) GetRegion() *string {
	return s.Region
}

func (s *DescribePackIpListResponseBodyIpList) GetRemark() *string {
	return s.Remark
}

func (s *DescribePackIpListResponseBodyIpList) GetStatus() *string {
	return s.Status
}

func (s *DescribePackIpListResponseBodyIpList) SetIp(v string) *DescribePackIpListResponseBodyIpList {
	s.Ip = &v
	return s
}

func (s *DescribePackIpListResponseBodyIpList) SetMemberUid(v string) *DescribePackIpListResponseBodyIpList {
	s.MemberUid = &v
	return s
}

func (s *DescribePackIpListResponseBodyIpList) SetNsmExpireAt(v int64) *DescribePackIpListResponseBodyIpList {
	s.NsmExpireAt = &v
	return s
}

func (s *DescribePackIpListResponseBodyIpList) SetNsmStartAt(v int64) *DescribePackIpListResponseBodyIpList {
	s.NsmStartAt = &v
	return s
}

func (s *DescribePackIpListResponseBodyIpList) SetNsmStatus(v int32) *DescribePackIpListResponseBodyIpList {
	s.NsmStatus = &v
	return s
}

func (s *DescribePackIpListResponseBodyIpList) SetProduct(v string) *DescribePackIpListResponseBodyIpList {
	s.Product = &v
	return s
}

func (s *DescribePackIpListResponseBodyIpList) SetRegion(v string) *DescribePackIpListResponseBodyIpList {
	s.Region = &v
	return s
}

func (s *DescribePackIpListResponseBodyIpList) SetRemark(v string) *DescribePackIpListResponseBodyIpList {
	s.Remark = &v
	return s
}

func (s *DescribePackIpListResponseBodyIpList) SetStatus(v string) *DescribePackIpListResponseBodyIpList {
	s.Status = &v
	return s
}

func (s *DescribePackIpListResponseBodyIpList) Validate() error {
	return dara.Validate(s)
}
