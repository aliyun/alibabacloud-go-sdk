// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGtmInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessStrategyNum(v int32) *DescribeGtmInstanceResponseBody
	GetAccessStrategyNum() *int32
	SetAddressPoolNum(v int32) *DescribeGtmInstanceResponseBody
	GetAddressPoolNum() *int32
	SetAlertGroup(v string) *DescribeGtmInstanceResponseBody
	GetAlertGroup() *string
	SetCname(v string) *DescribeGtmInstanceResponseBody
	GetCname() *string
	SetCnameMode(v string) *DescribeGtmInstanceResponseBody
	GetCnameMode() *string
	SetCreateTime(v string) *DescribeGtmInstanceResponseBody
	GetCreateTime() *string
	SetCreateTimestamp(v int64) *DescribeGtmInstanceResponseBody
	GetCreateTimestamp() *int64
	SetExpireTime(v string) *DescribeGtmInstanceResponseBody
	GetExpireTime() *string
	SetExpireTimestamp(v int64) *DescribeGtmInstanceResponseBody
	GetExpireTimestamp() *int64
	SetInstanceId(v string) *DescribeGtmInstanceResponseBody
	GetInstanceId() *string
	SetInstanceName(v string) *DescribeGtmInstanceResponseBody
	GetInstanceName() *string
	SetLbaStrategy(v string) *DescribeGtmInstanceResponseBody
	GetLbaStrategy() *string
	SetRequestId(v string) *DescribeGtmInstanceResponseBody
	GetRequestId() *string
	SetResourceGroupId(v string) *DescribeGtmInstanceResponseBody
	GetResourceGroupId() *string
	SetTtl(v int32) *DescribeGtmInstanceResponseBody
	GetTtl() *int32
	SetUserDomainName(v string) *DescribeGtmInstanceResponseBody
	GetUserDomainName() *string
	SetVersionCode(v string) *DescribeGtmInstanceResponseBody
	GetVersionCode() *string
}

type DescribeGtmInstanceResponseBody struct {
	// The number of access policies.
	//
	// example:
	//
	// 5
	AccessStrategyNum *int32 `json:"AccessStrategyNum,omitempty" xml:"AccessStrategyNum,omitempty"`
	// The number of address pools.
	//
	// example:
	//
	// 5
	AddressPoolNum *int32 `json:"AddressPoolNum,omitempty" xml:"AddressPoolNum,omitempty"`
	// The alert contact group.
	//
	// example:
	//
	// [\\"日常测试-研发组\\"]
	AlertGroup *string `json:"AlertGroup,omitempty" xml:"AlertGroup,omitempty"`
	// The CNAME access domain name.
	//
	// example:
	//
	// gtm-cn-wwo3a3hbz**.dns-example.top
	Cname *string `json:"Cname,omitempty" xml:"Cname,omitempty"`
	// The CNAME access mode. Valid values:
	//
	// - **SYSTEM_ASSIGN**: The system assigns a domain name.
	//
	// - **CUSTOM**: You use a custom domain name.
	//
	// example:
	//
	// CUSTOM
	CnameMode *string `json:"CnameMode,omitempty" xml:"CnameMode,omitempty"`
	// The time when the instance was created.
	//
	// example:
	//
	// 2018-06-06T11:34Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The time when the instance was created. This is a UNIX timestamp.
	//
	// example:
	//
	// 1528284856000
	CreateTimestamp *int64 `json:"CreateTimestamp,omitempty" xml:"CreateTimestamp,omitempty"`
	// The time when the instance expires.
	//
	// example:
	//
	// 2018-06-06T11:34Z
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The expiration time of the instance. This is a UNIX timestamp.
	//
	// example:
	//
	// 1528284856000
	ExpireTimestamp *int64 `json:"ExpireTimestamp,omitempty" xml:"ExpireTimestamp,omitempty"`
	// The ID of the GTM instance.
	//
	// example:
	//
	// gtm-cn-wwo3a3hbz**
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the instance.
	//
	// example:
	//
	// test
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The load balancing policy.
	//
	// - **ALL_RR**: round-robin
	//
	// - **RATIO**: weighted round-robin
	//
	// example:
	//
	// RATIO
	LbaStrategy *string `json:"LbaStrategy,omitempty" xml:"LbaStrategy,omitempty"`
	// The unique ID of the request.
	//
	// example:
	//
	// E41AA251-F9BA-48C6-99B2-2B82B26A573A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-testgroupid
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The global TTL.
	//
	// example:
	//
	// 60
	Ttl *int32 `json:"Ttl,omitempty" xml:"Ttl,omitempty"`
	// The custom domain name.
	//
	// example:
	//
	// dns-example.top
	UserDomainName *string `json:"UserDomainName,omitempty" xml:"UserDomainName,omitempty"`
	// The version code.
	//
	// example:
	//
	// biaozhun
	VersionCode *string `json:"VersionCode,omitempty" xml:"VersionCode,omitempty"`
}

func (s DescribeGtmInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeGtmInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeGtmInstanceResponseBody) GetAccessStrategyNum() *int32 {
	return s.AccessStrategyNum
}

func (s *DescribeGtmInstanceResponseBody) GetAddressPoolNum() *int32 {
	return s.AddressPoolNum
}

func (s *DescribeGtmInstanceResponseBody) GetAlertGroup() *string {
	return s.AlertGroup
}

func (s *DescribeGtmInstanceResponseBody) GetCname() *string {
	return s.Cname
}

func (s *DescribeGtmInstanceResponseBody) GetCnameMode() *string {
	return s.CnameMode
}

func (s *DescribeGtmInstanceResponseBody) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeGtmInstanceResponseBody) GetCreateTimestamp() *int64 {
	return s.CreateTimestamp
}

func (s *DescribeGtmInstanceResponseBody) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *DescribeGtmInstanceResponseBody) GetExpireTimestamp() *int64 {
	return s.ExpireTimestamp
}

func (s *DescribeGtmInstanceResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeGtmInstanceResponseBody) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeGtmInstanceResponseBody) GetLbaStrategy() *string {
	return s.LbaStrategy
}

func (s *DescribeGtmInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeGtmInstanceResponseBody) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeGtmInstanceResponseBody) GetTtl() *int32 {
	return s.Ttl
}

func (s *DescribeGtmInstanceResponseBody) GetUserDomainName() *string {
	return s.UserDomainName
}

func (s *DescribeGtmInstanceResponseBody) GetVersionCode() *string {
	return s.VersionCode
}

func (s *DescribeGtmInstanceResponseBody) SetAccessStrategyNum(v int32) *DescribeGtmInstanceResponseBody {
	s.AccessStrategyNum = &v
	return s
}

func (s *DescribeGtmInstanceResponseBody) SetAddressPoolNum(v int32) *DescribeGtmInstanceResponseBody {
	s.AddressPoolNum = &v
	return s
}

func (s *DescribeGtmInstanceResponseBody) SetAlertGroup(v string) *DescribeGtmInstanceResponseBody {
	s.AlertGroup = &v
	return s
}

func (s *DescribeGtmInstanceResponseBody) SetCname(v string) *DescribeGtmInstanceResponseBody {
	s.Cname = &v
	return s
}

func (s *DescribeGtmInstanceResponseBody) SetCnameMode(v string) *DescribeGtmInstanceResponseBody {
	s.CnameMode = &v
	return s
}

func (s *DescribeGtmInstanceResponseBody) SetCreateTime(v string) *DescribeGtmInstanceResponseBody {
	s.CreateTime = &v
	return s
}

func (s *DescribeGtmInstanceResponseBody) SetCreateTimestamp(v int64) *DescribeGtmInstanceResponseBody {
	s.CreateTimestamp = &v
	return s
}

func (s *DescribeGtmInstanceResponseBody) SetExpireTime(v string) *DescribeGtmInstanceResponseBody {
	s.ExpireTime = &v
	return s
}

func (s *DescribeGtmInstanceResponseBody) SetExpireTimestamp(v int64) *DescribeGtmInstanceResponseBody {
	s.ExpireTimestamp = &v
	return s
}

func (s *DescribeGtmInstanceResponseBody) SetInstanceId(v string) *DescribeGtmInstanceResponseBody {
	s.InstanceId = &v
	return s
}

func (s *DescribeGtmInstanceResponseBody) SetInstanceName(v string) *DescribeGtmInstanceResponseBody {
	s.InstanceName = &v
	return s
}

func (s *DescribeGtmInstanceResponseBody) SetLbaStrategy(v string) *DescribeGtmInstanceResponseBody {
	s.LbaStrategy = &v
	return s
}

func (s *DescribeGtmInstanceResponseBody) SetRequestId(v string) *DescribeGtmInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeGtmInstanceResponseBody) SetResourceGroupId(v string) *DescribeGtmInstanceResponseBody {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeGtmInstanceResponseBody) SetTtl(v int32) *DescribeGtmInstanceResponseBody {
	s.Ttl = &v
	return s
}

func (s *DescribeGtmInstanceResponseBody) SetUserDomainName(v string) *DescribeGtmInstanceResponseBody {
	s.UserDomainName = &v
	return s
}

func (s *DescribeGtmInstanceResponseBody) SetVersionCode(v string) *DescribeGtmInstanceResponseBody {
	s.VersionCode = &v
	return s
}

func (s *DescribeGtmInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
