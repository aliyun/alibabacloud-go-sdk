// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSensitiveRequestLogRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccount(v string) *DescribeSensitiveRequestLogRequest
	GetAccount() *string
	SetApiFormat(v string) *DescribeSensitiveRequestLogRequest
	GetApiFormat() *string
	SetClientIP(v string) *DescribeSensitiveRequestLogRequest
	GetClientIP() *string
	SetClusterId(v string) *DescribeSensitiveRequestLogRequest
	GetClusterId() *string
	SetEndTime(v int64) *DescribeSensitiveRequestLogRequest
	GetEndTime() *int64
	SetInstanceId(v string) *DescribeSensitiveRequestLogRequest
	GetInstanceId() *string
	SetMatchedHost(v string) *DescribeSensitiveRequestLogRequest
	GetMatchedHost() *string
	SetPageNumber(v int64) *DescribeSensitiveRequestLogRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribeSensitiveRequestLogRequest
	GetPageSize() *int64
	SetRegionId(v string) *DescribeSensitiveRequestLogRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *DescribeSensitiveRequestLogRequest
	GetResourceManagerResourceGroupId() *string
	SetSensitiveCode(v string) *DescribeSensitiveRequestLogRequest
	GetSensitiveCode() *string
	SetSensitiveData(v string) *DescribeSensitiveRequestLogRequest
	GetSensitiveData() *string
	SetStartTime(v int64) *DescribeSensitiveRequestLogRequest
	GetStartTime() *int64
}

type DescribeSensitiveRequestLogRequest struct {
	// example:
	//
	// admin
	Account *string `json:"Account,omitempty" xml:"Account,omitempty"`
	// The API.
	//
	// example:
	//
	// /api/users/login
	ApiFormat *string `json:"ApiFormat,omitempty" xml:"ApiFormat,omitempty"`
	// The IP address.
	//
	// example:
	//
	// 103.118.55.**
	ClientIP *string `json:"ClientIP,omitempty" xml:"ClientIP,omitempty"`
	// The ID of the hybrid cloud cluster.
	//
	// >For hybrid cloud scenarios only, you can call the [DescribeHybridCloudClusters](https://help.aliyun.com/document_detail/2849376.html) operation to query the hybrid cloud clusters.
	//
	// example:
	//
	// 433
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The end of the time range to query. The value is a UNIX timestamp displayed in UTC. Unit: seconds.
	//
	// example:
	//
	// 1726057800
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the WAF instance.
	//
	// >  You can call the [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) operation to query the ID of the WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_cdnsdf3****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The domain name of the API.
	//
	// example:
	//
	// a.***.com
	MatchedHost *string `json:"MatchedHost,omitempty" xml:"MatchedHost,omitempty"`
	// The page number. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: **10**.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the WAF instance. Valid values:
	//
	// 	- **cn-hangzhou**: Chinese mainland.
	//
	// 	- **ap-southeast-1**: outside the Chinese mainland.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the Alibaba Cloud resource group.
	//
	// example:
	//
	// rg-acfm***q
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
	// The type of the sensitive data.
	//
	// >  You can call the [DescribeApisecRules](https://help.aliyun.com/document_detail/2859155.html) operation to query the supported types of sensitive data.
	//
	// example:
	//
	// 1000,1001
	SensitiveCode *string `json:"SensitiveCode,omitempty" xml:"SensitiveCode,omitempty"`
	// The response sensitive data.
	//
	// example:
	//
	// user
	SensitiveData *string `json:"SensitiveData,omitempty" xml:"SensitiveData,omitempty"`
	// The beginning of the time range to query. The value is a UNIX timestamp displayed in UTC. Unit: seconds.
	//
	// example:
	//
	// 1723392000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeSensitiveRequestLogRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSensitiveRequestLogRequest) GoString() string {
	return s.String()
}

func (s *DescribeSensitiveRequestLogRequest) GetAccount() *string {
	return s.Account
}

func (s *DescribeSensitiveRequestLogRequest) GetApiFormat() *string {
	return s.ApiFormat
}

func (s *DescribeSensitiveRequestLogRequest) GetClientIP() *string {
	return s.ClientIP
}

func (s *DescribeSensitiveRequestLogRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeSensitiveRequestLogRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeSensitiveRequestLogRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeSensitiveRequestLogRequest) GetMatchedHost() *string {
	return s.MatchedHost
}

func (s *DescribeSensitiveRequestLogRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeSensitiveRequestLogRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeSensitiveRequestLogRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeSensitiveRequestLogRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *DescribeSensitiveRequestLogRequest) GetSensitiveCode() *string {
	return s.SensitiveCode
}

func (s *DescribeSensitiveRequestLogRequest) GetSensitiveData() *string {
	return s.SensitiveData
}

func (s *DescribeSensitiveRequestLogRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeSensitiveRequestLogRequest) SetAccount(v string) *DescribeSensitiveRequestLogRequest {
	s.Account = &v
	return s
}

func (s *DescribeSensitiveRequestLogRequest) SetApiFormat(v string) *DescribeSensitiveRequestLogRequest {
	s.ApiFormat = &v
	return s
}

func (s *DescribeSensitiveRequestLogRequest) SetClientIP(v string) *DescribeSensitiveRequestLogRequest {
	s.ClientIP = &v
	return s
}

func (s *DescribeSensitiveRequestLogRequest) SetClusterId(v string) *DescribeSensitiveRequestLogRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeSensitiveRequestLogRequest) SetEndTime(v int64) *DescribeSensitiveRequestLogRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeSensitiveRequestLogRequest) SetInstanceId(v string) *DescribeSensitiveRequestLogRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeSensitiveRequestLogRequest) SetMatchedHost(v string) *DescribeSensitiveRequestLogRequest {
	s.MatchedHost = &v
	return s
}

func (s *DescribeSensitiveRequestLogRequest) SetPageNumber(v int64) *DescribeSensitiveRequestLogRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeSensitiveRequestLogRequest) SetPageSize(v int64) *DescribeSensitiveRequestLogRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeSensitiveRequestLogRequest) SetRegionId(v string) *DescribeSensitiveRequestLogRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSensitiveRequestLogRequest) SetResourceManagerResourceGroupId(v string) *DescribeSensitiveRequestLogRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeSensitiveRequestLogRequest) SetSensitiveCode(v string) *DescribeSensitiveRequestLogRequest {
	s.SensitiveCode = &v
	return s
}

func (s *DescribeSensitiveRequestLogRequest) SetSensitiveData(v string) *DescribeSensitiveRequestLogRequest {
	s.SensitiveData = &v
	return s
}

func (s *DescribeSensitiveRequestLogRequest) SetStartTime(v int64) *DescribeSensitiveRequestLogRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeSensitiveRequestLogRequest) Validate() error {
	return dara.Validate(s)
}
