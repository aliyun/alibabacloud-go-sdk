// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGtmInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetGtmInstances(v *DescribeGtmInstancesResponseBodyGtmInstances) *DescribeGtmInstancesResponseBody
	GetGtmInstances() *DescribeGtmInstancesResponseBodyGtmInstances
	SetPageNumber(v int32) *DescribeGtmInstancesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeGtmInstancesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeGtmInstancesResponseBody
	GetRequestId() *string
	SetTotalItems(v int32) *DescribeGtmInstancesResponseBody
	GetTotalItems() *int32
	SetTotalPages(v int32) *DescribeGtmInstancesResponseBody
	GetTotalPages() *int32
}

type DescribeGtmInstancesResponseBody struct {
	// The list of queried instances.
	GtmInstances *DescribeGtmInstancesResponseBodyGtmInstances `json:"GtmInstances,omitempty" xml:"GtmInstances,omitempty" type:"Struct"`
	// The returned page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 193B0163-7F93-42DF-AB05-ACEEB7D22707
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalItems *int32 `json:"TotalItems,omitempty" xml:"TotalItems,omitempty"`
	// The total number of pages returned.
	//
	// example:
	//
	// 1
	TotalPages *int32 `json:"TotalPages,omitempty" xml:"TotalPages,omitempty"`
}

func (s DescribeGtmInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeGtmInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeGtmInstancesResponseBody) GetGtmInstances() *DescribeGtmInstancesResponseBodyGtmInstances {
	return s.GtmInstances
}

func (s *DescribeGtmInstancesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeGtmInstancesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeGtmInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeGtmInstancesResponseBody) GetTotalItems() *int32 {
	return s.TotalItems
}

func (s *DescribeGtmInstancesResponseBody) GetTotalPages() *int32 {
	return s.TotalPages
}

func (s *DescribeGtmInstancesResponseBody) SetGtmInstances(v *DescribeGtmInstancesResponseBodyGtmInstances) *DescribeGtmInstancesResponseBody {
	s.GtmInstances = v
	return s
}

func (s *DescribeGtmInstancesResponseBody) SetPageNumber(v int32) *DescribeGtmInstancesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeGtmInstancesResponseBody) SetPageSize(v int32) *DescribeGtmInstancesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeGtmInstancesResponseBody) SetRequestId(v string) *DescribeGtmInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeGtmInstancesResponseBody) SetTotalItems(v int32) *DescribeGtmInstancesResponseBody {
	s.TotalItems = &v
	return s
}

func (s *DescribeGtmInstancesResponseBody) SetTotalPages(v int32) *DescribeGtmInstancesResponseBody {
	s.TotalPages = &v
	return s
}

func (s *DescribeGtmInstancesResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeGtmInstancesResponseBodyGtmInstances struct {
	GtmInstance []*DescribeGtmInstancesResponseBodyGtmInstancesGtmInstance `json:"GtmInstance,omitempty" xml:"GtmInstance,omitempty" type:"Repeated"`
}

func (s DescribeGtmInstancesResponseBodyGtmInstances) String() string {
	return dara.Prettify(s)
}

func (s DescribeGtmInstancesResponseBodyGtmInstances) GoString() string {
	return s.String()
}

func (s *DescribeGtmInstancesResponseBodyGtmInstances) GetGtmInstance() []*DescribeGtmInstancesResponseBodyGtmInstancesGtmInstance {
	return s.GtmInstance
}

func (s *DescribeGtmInstancesResponseBodyGtmInstances) SetGtmInstance(v []*DescribeGtmInstancesResponseBodyGtmInstancesGtmInstance) *DescribeGtmInstancesResponseBodyGtmInstances {
	s.GtmInstance = v
	return s
}

func (s *DescribeGtmInstancesResponseBodyGtmInstances) Validate() error {
	return dara.Validate(s)
}

type DescribeGtmInstancesResponseBodyGtmInstancesGtmInstance struct {
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
	// The name of the alert group.
	//
	// example:
	//
	// [\\\\"R\\&D group\\\\"]
	AlertGroup *string `json:"AlertGroup,omitempty" xml:"AlertGroup,omitempty"`
	// The CNAME domain name that is used to access the instance.
	//
	// example:
	//
	// instance1.14.com
	Cname *string `json:"Cname,omitempty" xml:"Cname,omitempty"`
	// The CNAME domain name used to access the instance. Valid values:
	//
	// 	- **SYSTEM_ASSIGN**: A CNAME domain name assigned by the system is used.
	//
	// 	- **CUSTOM**: A custom CNAME domain name is used.
	//
	// example:
	//
	// SYSTEM_ASSIGN
	CnameMode *string `json:"CnameMode,omitempty" xml:"CnameMode,omitempty"`
	// The time when the instance was created.
	//
	// example:
	//
	// 2018-06-06T11:34Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The UNIX timestamp that indicates when the instance was created.
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
	// The UNIX timestamp that indicates when the instance expires.
	//
	// example:
	//
	// 1528284856000
	ExpireTimestamp *int64 `json:"ExpireTimestamp,omitempty" xml:"ExpireTimestamp,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// instance1
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the instance.
	//
	// example:
	//
	// test
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The load balancing policy that is used. Valid values:
	//
	// 	- **ALL_RR**: Load balancing
	//
	// 	- **RATIO**: Weighted round-robin
	//
	// example:
	//
	// RATIO
	LbaStrategy *string `json:"LbaStrategy,omitempty" xml:"LbaStrategy,omitempty"`
	// The ID of the resource group to which the instance belongs.
	//
	// example:
	//
	// rg-acfm2q2jqpjh***
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The global time to live (TTL).
	//
	// example:
	//
	// 60
	Ttl *int32 `json:"Ttl,omitempty" xml:"Ttl,omitempty"`
	// The domain name of the user.
	//
	// example:
	//
	// www.example.com
	UserDomainName *string `json:"UserDomainName,omitempty" xml:"UserDomainName,omitempty"`
	// The version code of the instance.
	//
	// example:
	//
	// biaozhun
	VersionCode *string `json:"VersionCode,omitempty" xml:"VersionCode,omitempty"`
}

func (s DescribeGtmInstancesResponseBodyGtmInstancesGtmInstance) String() string {
	return dara.Prettify(s)
}

func (s DescribeGtmInstancesResponseBodyGtmInstancesGtmInstance) GoString() string {
	return s.String()
}

func (s *DescribeGtmInstancesResponseBodyGtmInstancesGtmInstance) GetAccessStrategyNum() *int32 {
	return s.AccessStrategyNum
}

func (s *DescribeGtmInstancesResponseBodyGtmInstancesGtmInstance) GetAddressPoolNum() *int32 {
	return s.AddressPoolNum
}

func (s *DescribeGtmInstancesResponseBodyGtmInstancesGtmInstance) GetAlertGroup() *string {
	return s.AlertGroup
}

func (s *DescribeGtmInstancesResponseBodyGtmInstancesGtmInstance) GetCname() *string {
	return s.Cname
}

func (s *DescribeGtmInstancesResponseBodyGtmInstancesGtmInstance) GetCnameMode() *string {
	return s.CnameMode
}

func (s *DescribeGtmInstancesResponseBodyGtmInstancesGtmInstance) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeGtmInstancesResponseBodyGtmInstancesGtmInstance) GetCreateTimestamp() *int64 {
	return s.CreateTimestamp
}

func (s *DescribeGtmInstancesResponseBodyGtmInstancesGtmInstance) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *DescribeGtmInstancesResponseBodyGtmInstancesGtmInstance) GetExpireTimestamp() *int64 {
	return s.ExpireTimestamp
}

func (s *DescribeGtmInstancesResponseBodyGtmInstancesGtmInstance) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeGtmInstancesResponseBodyGtmInstancesGtmInstance) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeGtmInstancesResponseBodyGtmInstancesGtmInstance) GetLbaStrategy() *string {
	return s.LbaStrategy
}

func (s *DescribeGtmInstancesResponseBodyGtmInstancesGtmInstance) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeGtmInstancesResponseBodyGtmInstancesGtmInstance) GetTtl() *int32 {
	return s.Ttl
}

func (s *DescribeGtmInstancesResponseBodyGtmInstancesGtmInstance) GetUserDomainName() *string {
	return s.UserDomainName
}

func (s *DescribeGtmInstancesResponseBodyGtmInstancesGtmInstance) GetVersionCode() *string {
	return s.VersionCode
}

func (s *DescribeGtmInstancesResponseBodyGtmInstancesGtmInstance) SetAccessStrategyNum(v int32) *DescribeGtmInstancesResponseBodyGtmInstancesGtmInstance {
	s.AccessStrategyNum = &v
	return s
}

func (s *DescribeGtmInstancesResponseBodyGtmInstancesGtmInstance) SetAddressPoolNum(v int32) *DescribeGtmInstancesResponseBodyGtmInstancesGtmInstance {
	s.AddressPoolNum = &v
	return s
}

func (s *DescribeGtmInstancesResponseBodyGtmInstancesGtmInstance) SetAlertGroup(v string) *DescribeGtmInstancesResponseBodyGtmInstancesGtmInstance {
	s.AlertGroup = &v
	return s
}

func (s *DescribeGtmInstancesResponseBodyGtmInstancesGtmInstance) SetCname(v string) *DescribeGtmInstancesResponseBodyGtmInstancesGtmInstance {
	s.Cname = &v
	return s
}

func (s *DescribeGtmInstancesResponseBodyGtmInstancesGtmInstance) SetCnameMode(v string) *DescribeGtmInstancesResponseBodyGtmInstancesGtmInstance {
	s.CnameMode = &v
	return s
}

func (s *DescribeGtmInstancesResponseBodyGtmInstancesGtmInstance) SetCreateTime(v string) *DescribeGtmInstancesResponseBodyGtmInstancesGtmInstance {
	s.CreateTime = &v
	return s
}

func (s *DescribeGtmInstancesResponseBodyGtmInstancesGtmInstance) SetCreateTimestamp(v int64) *DescribeGtmInstancesResponseBodyGtmInstancesGtmInstance {
	s.CreateTimestamp = &v
	return s
}

func (s *DescribeGtmInstancesResponseBodyGtmInstancesGtmInstance) SetExpireTime(v string) *DescribeGtmInstancesResponseBodyGtmInstancesGtmInstance {
	s.ExpireTime = &v
	return s
}

func (s *DescribeGtmInstancesResponseBodyGtmInstancesGtmInstance) SetExpireTimestamp(v int64) *DescribeGtmInstancesResponseBodyGtmInstancesGtmInstance {
	s.ExpireTimestamp = &v
	return s
}

func (s *DescribeGtmInstancesResponseBodyGtmInstancesGtmInstance) SetInstanceId(v string) *DescribeGtmInstancesResponseBodyGtmInstancesGtmInstance {
	s.InstanceId = &v
	return s
}

func (s *DescribeGtmInstancesResponseBodyGtmInstancesGtmInstance) SetInstanceName(v string) *DescribeGtmInstancesResponseBodyGtmInstancesGtmInstance {
	s.InstanceName = &v
	return s
}

func (s *DescribeGtmInstancesResponseBodyGtmInstancesGtmInstance) SetLbaStrategy(v string) *DescribeGtmInstancesResponseBodyGtmInstancesGtmInstance {
	s.LbaStrategy = &v
	return s
}

func (s *DescribeGtmInstancesResponseBodyGtmInstancesGtmInstance) SetResourceGroupId(v string) *DescribeGtmInstancesResponseBodyGtmInstancesGtmInstance {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeGtmInstancesResponseBodyGtmInstancesGtmInstance) SetTtl(v int32) *DescribeGtmInstancesResponseBodyGtmInstancesGtmInstance {
	s.Ttl = &v
	return s
}

func (s *DescribeGtmInstancesResponseBodyGtmInstancesGtmInstance) SetUserDomainName(v string) *DescribeGtmInstancesResponseBodyGtmInstancesGtmInstance {
	s.UserDomainName = &v
	return s
}

func (s *DescribeGtmInstancesResponseBodyGtmInstancesGtmInstance) SetVersionCode(v string) *DescribeGtmInstancesResponseBodyGtmInstancesGtmInstance {
	s.VersionCode = &v
	return s
}

func (s *DescribeGtmInstancesResponseBodyGtmInstancesGtmInstance) Validate() error {
	return dara.Validate(s)
}
