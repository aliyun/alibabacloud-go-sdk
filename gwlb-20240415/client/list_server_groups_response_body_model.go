// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListServerGroupsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListServerGroupsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListServerGroupsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListServerGroupsResponseBody
	GetRequestId() *string
	SetServerGroups(v []*ListServerGroupsResponseBodyServerGroups) *ListServerGroupsResponseBody
	GetServerGroups() []*ListServerGroupsResponseBodyServerGroups
	SetTotalCount(v int32) *ListServerGroupsResponseBody
	GetTotalCount() *int32
}

type ListServerGroupsResponseBody struct {
	// The number of entries per page.
	//
	// Valid values: 1 to 1000.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// A pagination token. It can be used in the next request to retrieve a new page of results. Valid values:
	//
	// 	- If **NextToken*	- is empty, no next page exists.
	//
	// 	- If a value of **NextToken*	- is returned, the value indicates the token that is used for the next query.
	//
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 54B48E3D-DF70-471B-AA93-08E683A1B45
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The configurations of the server group.
	ServerGroups []*ListServerGroupsResponseBodyServerGroups `json:"ServerGroups,omitempty" xml:"ServerGroups,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListServerGroupsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListServerGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *ListServerGroupsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListServerGroupsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListServerGroupsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListServerGroupsResponseBody) GetServerGroups() []*ListServerGroupsResponseBodyServerGroups {
	return s.ServerGroups
}

func (s *ListServerGroupsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListServerGroupsResponseBody) SetMaxResults(v int32) *ListServerGroupsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListServerGroupsResponseBody) SetNextToken(v string) *ListServerGroupsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListServerGroupsResponseBody) SetRequestId(v string) *ListServerGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListServerGroupsResponseBody) SetServerGroups(v []*ListServerGroupsResponseBodyServerGroups) *ListServerGroupsResponseBody {
	s.ServerGroups = v
	return s
}

func (s *ListServerGroupsResponseBody) SetTotalCount(v int32) *ListServerGroupsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListServerGroupsResponseBody) Validate() error {
	if s.ServerGroups != nil {
		for _, item := range s.ServerGroups {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListServerGroupsResponseBodyServerGroups struct {
	// The configurations of connection draining.
	ConnectionDrainConfig *ListServerGroupsResponseBodyServerGroupsConnectionDrainConfig `json:"ConnectionDrainConfig,omitempty" xml:"ConnectionDrainConfig,omitempty" type:"Struct"`
	// The time when the resource was created. The time follows the ISO 8601 standard in the **yyyy-MM-ddTHH:mm:ssZ*	- format. The time is displayed in UTC.
	//
	// example:
	//
	// 2024-08-05T18:24:07Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The configuration of health checks.
	HealthCheckConfig *ListServerGroupsResponseBodyServerGroupsHealthCheckConfig `json:"HealthCheckConfig,omitempty" xml:"HealthCheckConfig,omitempty" type:"Struct"`
	// The backend protocol. Valid values:
	//
	// 	- **GENEVE**.
	//
	// example:
	//
	// GENEVE
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// The IDs of the GWLB instances that are associated with the server group.
	RelatedLoadBalancerIds []*string `json:"RelatedLoadBalancerIds,omitempty" xml:"RelatedLoadBalancerIds,omitempty" type:"Repeated"`
	// The resource group ID.
	//
	// example:
	//
	// rg-atstuj3rtop****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The scheduling algorithm. Valid values:
	//
	// 	- **5TCH**: indicates consistent hashing that is based on the following factors: source IP address, destination IP address, source port, protocol, and destination port. Requests that contain the same information based on the preceding factors are forwarded to the same backend server.
	//
	// 	- **3TCH**: indicates consistent hashing that is based on the following factors: source IP address, destination IP address, and protocol. Requests that contain the same information based on the preceding factors are forwarded to the same backend server.
	//
	// 	- **2TCH**: indicates consistent hashing that is based on the following factors: source IP address and destination IP address. Requests that contain the same information based on the preceding factors are forwarded to the same backend server.
	//
	// example:
	//
	// 5TCH
	Scheduler *string `json:"Scheduler,omitempty" xml:"Scheduler,omitempty"`
	// The number of server groups.
	//
	// example:
	//
	// 2
	ServerCount *int32 `json:"ServerCount,omitempty" xml:"ServerCount,omitempty"`
	// Specifies how GWLB processes requests over existing connections when a backend server is not running as expected. Valid values:
	//
	// 	- **NoRebalance**: GWLB continues to forward requests over existing connections to the unhealthy backend server.
	//
	// 	- **Rebalance**: GWLB forwards requests over existing connections to the remaining healthy backend servers.
	//
	// example:
	//
	// NoRebalance
	ServerFailoverMode *string `json:"ServerFailoverMode,omitempty" xml:"ServerFailoverMode,omitempty"`
	// The server group ID.
	//
	// example:
	//
	// sgp-atstuj3rtoptyui****
	ServerGroupId *string `json:"ServerGroupId,omitempty" xml:"ServerGroupId,omitempty"`
	// The server group name.
	//
	// example:
	//
	// testServerGroupName
	ServerGroupName *string `json:"ServerGroupName,omitempty" xml:"ServerGroupName,omitempty"`
	// The status of the server group. Valid values:
	//
	// 	- **Creating**
	//
	// 	- **Available**
	//
	// 	- **Configuring**
	//
	// example:
	//
	// Available
	ServerGroupStatus *string `json:"ServerGroupStatus,omitempty" xml:"ServerGroupStatus,omitempty"`
	// The server group type. Valid values:
	//
	// 	- **Instance**: allows you to specify servers of the **Ecs**, **Eni**, or **Eci*	- type.
	//
	// 	- **Ip**: allows you to add servers of by specifying IP addresses.
	//
	// example:
	//
	// Instance
	ServerGroupType *string `json:"ServerGroupType,omitempty" xml:"ServerGroupType,omitempty"`
	// The tags.
	Tags []*ListServerGroupsResponseBodyServerGroupsTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The VPC ID.
	//
	// example:
	//
	// vpc-bp15zckdt37pq72zv****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s ListServerGroupsResponseBodyServerGroups) String() string {
	return dara.Prettify(s)
}

func (s ListServerGroupsResponseBodyServerGroups) GoString() string {
	return s.String()
}

func (s *ListServerGroupsResponseBodyServerGroups) GetConnectionDrainConfig() *ListServerGroupsResponseBodyServerGroupsConnectionDrainConfig {
	return s.ConnectionDrainConfig
}

func (s *ListServerGroupsResponseBodyServerGroups) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListServerGroupsResponseBodyServerGroups) GetHealthCheckConfig() *ListServerGroupsResponseBodyServerGroupsHealthCheckConfig {
	return s.HealthCheckConfig
}

func (s *ListServerGroupsResponseBodyServerGroups) GetProtocol() *string {
	return s.Protocol
}

func (s *ListServerGroupsResponseBodyServerGroups) GetRelatedLoadBalancerIds() []*string {
	return s.RelatedLoadBalancerIds
}

func (s *ListServerGroupsResponseBodyServerGroups) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListServerGroupsResponseBodyServerGroups) GetScheduler() *string {
	return s.Scheduler
}

func (s *ListServerGroupsResponseBodyServerGroups) GetServerCount() *int32 {
	return s.ServerCount
}

func (s *ListServerGroupsResponseBodyServerGroups) GetServerFailoverMode() *string {
	return s.ServerFailoverMode
}

func (s *ListServerGroupsResponseBodyServerGroups) GetServerGroupId() *string {
	return s.ServerGroupId
}

func (s *ListServerGroupsResponseBodyServerGroups) GetServerGroupName() *string {
	return s.ServerGroupName
}

func (s *ListServerGroupsResponseBodyServerGroups) GetServerGroupStatus() *string {
	return s.ServerGroupStatus
}

func (s *ListServerGroupsResponseBodyServerGroups) GetServerGroupType() *string {
	return s.ServerGroupType
}

func (s *ListServerGroupsResponseBodyServerGroups) GetTags() []*ListServerGroupsResponseBodyServerGroupsTags {
	return s.Tags
}

func (s *ListServerGroupsResponseBodyServerGroups) GetVpcId() *string {
	return s.VpcId
}

func (s *ListServerGroupsResponseBodyServerGroups) SetConnectionDrainConfig(v *ListServerGroupsResponseBodyServerGroupsConnectionDrainConfig) *ListServerGroupsResponseBodyServerGroups {
	s.ConnectionDrainConfig = v
	return s
}

func (s *ListServerGroupsResponseBodyServerGroups) SetCreateTime(v string) *ListServerGroupsResponseBodyServerGroups {
	s.CreateTime = &v
	return s
}

func (s *ListServerGroupsResponseBodyServerGroups) SetHealthCheckConfig(v *ListServerGroupsResponseBodyServerGroupsHealthCheckConfig) *ListServerGroupsResponseBodyServerGroups {
	s.HealthCheckConfig = v
	return s
}

func (s *ListServerGroupsResponseBodyServerGroups) SetProtocol(v string) *ListServerGroupsResponseBodyServerGroups {
	s.Protocol = &v
	return s
}

func (s *ListServerGroupsResponseBodyServerGroups) SetRelatedLoadBalancerIds(v []*string) *ListServerGroupsResponseBodyServerGroups {
	s.RelatedLoadBalancerIds = v
	return s
}

func (s *ListServerGroupsResponseBodyServerGroups) SetResourceGroupId(v string) *ListServerGroupsResponseBodyServerGroups {
	s.ResourceGroupId = &v
	return s
}

func (s *ListServerGroupsResponseBodyServerGroups) SetScheduler(v string) *ListServerGroupsResponseBodyServerGroups {
	s.Scheduler = &v
	return s
}

func (s *ListServerGroupsResponseBodyServerGroups) SetServerCount(v int32) *ListServerGroupsResponseBodyServerGroups {
	s.ServerCount = &v
	return s
}

func (s *ListServerGroupsResponseBodyServerGroups) SetServerFailoverMode(v string) *ListServerGroupsResponseBodyServerGroups {
	s.ServerFailoverMode = &v
	return s
}

func (s *ListServerGroupsResponseBodyServerGroups) SetServerGroupId(v string) *ListServerGroupsResponseBodyServerGroups {
	s.ServerGroupId = &v
	return s
}

func (s *ListServerGroupsResponseBodyServerGroups) SetServerGroupName(v string) *ListServerGroupsResponseBodyServerGroups {
	s.ServerGroupName = &v
	return s
}

func (s *ListServerGroupsResponseBodyServerGroups) SetServerGroupStatus(v string) *ListServerGroupsResponseBodyServerGroups {
	s.ServerGroupStatus = &v
	return s
}

func (s *ListServerGroupsResponseBodyServerGroups) SetServerGroupType(v string) *ListServerGroupsResponseBodyServerGroups {
	s.ServerGroupType = &v
	return s
}

func (s *ListServerGroupsResponseBodyServerGroups) SetTags(v []*ListServerGroupsResponseBodyServerGroupsTags) *ListServerGroupsResponseBodyServerGroups {
	s.Tags = v
	return s
}

func (s *ListServerGroupsResponseBodyServerGroups) SetVpcId(v string) *ListServerGroupsResponseBodyServerGroups {
	s.VpcId = &v
	return s
}

func (s *ListServerGroupsResponseBodyServerGroups) Validate() error {
	if s.ConnectionDrainConfig != nil {
		if err := s.ConnectionDrainConfig.Validate(); err != nil {
			return err
		}
	}
	if s.HealthCheckConfig != nil {
		if err := s.HealthCheckConfig.Validate(); err != nil {
			return err
		}
	}
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListServerGroupsResponseBodyServerGroupsConnectionDrainConfig struct {
	// Indicates whether connection draining is enabled. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// false
	ConnectionDrainEnabled *bool `json:"ConnectionDrainEnabled,omitempty" xml:"ConnectionDrainEnabled,omitempty"`
	// The timeout period of connection draining.
	//
	// Unit: seconds
	//
	// Valid values: 1 to 3600.
	//
	// example:
	//
	// 300
	ConnectionDrainTimeout *int32 `json:"ConnectionDrainTimeout,omitempty" xml:"ConnectionDrainTimeout,omitempty"`
}

func (s ListServerGroupsResponseBodyServerGroupsConnectionDrainConfig) String() string {
	return dara.Prettify(s)
}

func (s ListServerGroupsResponseBodyServerGroupsConnectionDrainConfig) GoString() string {
	return s.String()
}

func (s *ListServerGroupsResponseBodyServerGroupsConnectionDrainConfig) GetConnectionDrainEnabled() *bool {
	return s.ConnectionDrainEnabled
}

func (s *ListServerGroupsResponseBodyServerGroupsConnectionDrainConfig) GetConnectionDrainTimeout() *int32 {
	return s.ConnectionDrainTimeout
}

func (s *ListServerGroupsResponseBodyServerGroupsConnectionDrainConfig) SetConnectionDrainEnabled(v bool) *ListServerGroupsResponseBodyServerGroupsConnectionDrainConfig {
	s.ConnectionDrainEnabled = &v
	return s
}

func (s *ListServerGroupsResponseBodyServerGroupsConnectionDrainConfig) SetConnectionDrainTimeout(v int32) *ListServerGroupsResponseBodyServerGroupsConnectionDrainConfig {
	s.ConnectionDrainTimeout = &v
	return s
}

func (s *ListServerGroupsResponseBodyServerGroupsConnectionDrainConfig) Validate() error {
	return dara.Validate(s)
}

type ListServerGroupsResponseBodyServerGroupsHealthCheckConfig struct {
	// The backend server port that is used for health checks.
	//
	// Valid values: **1*	- to **65535**.
	//
	// example:
	//
	// 80
	HealthCheckConnectPort *int32 `json:"HealthCheckConnectPort,omitempty" xml:"HealthCheckConnectPort,omitempty"`
	// The maximum timeout period of a health check.
	//
	// Unit: seconds
	//
	// Valid values: **1*	- to **300**.
	//
	// example:
	//
	// 5
	HealthCheckConnectTimeout *int32 `json:"HealthCheckConnectTimeout,omitempty" xml:"HealthCheckConnectTimeout,omitempty"`
	// The domain name that is used for health checks. Valid values:
	//
	// 	- **$SERVER_IP**: the internal IP address of a backend server.
	//
	// 	- **domain**: a domain name. The domain name must be 1 to 80 characters in length, and can contain letters, digits, hyphens (-), and periods (.).
	//
	// > This parameter takes effect only if you set **HealthCheckProtocol*	- to **HTTP**.
	//
	// example:
	//
	// $SERVER_IP
	HealthCheckDomain *string `json:"HealthCheckDomain,omitempty" xml:"HealthCheckDomain,omitempty"`
	// Indicates whether the health check feature is enabled. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	HealthCheckEnabled *bool   `json:"HealthCheckEnabled,omitempty" xml:"HealthCheckEnabled,omitempty"`
	HealthCheckExp     *string `json:"HealthCheckExp,omitempty" xml:"HealthCheckExp,omitempty"`
	// The HTTP status codes that the system returns for health checks.
	HealthCheckHttpCode []*string `json:"HealthCheckHttpCode,omitempty" xml:"HealthCheckHttpCode,omitempty" type:"Repeated"`
	// The interval at which health checks are performed.
	//
	// Unit: seconds
	//
	// Valid values: **1*	- to **50**.
	//
	// example:
	//
	// 10
	HealthCheckInterval *int32 `json:"HealthCheckInterval,omitempty" xml:"HealthCheckInterval,omitempty"`
	// The URL that is used for health checks.
	//
	// The URL must be 1 to 80 characters in length, and can contain letters, digits, and the following special characters: ` - / . % ? # &  `The URL must start with a forward slash (/).
	//
	// > This parameter takes effect only if you set **HealthCheckProtocol*	- to **HTTP**.
	//
	// example:
	//
	// /test/index.html
	HealthCheckPath *string `json:"HealthCheckPath,omitempty" xml:"HealthCheckPath,omitempty"`
	// The protocol that is used for health checks. Valid values:
	//
	// 	- **TCP**: TCP health checks send TCP SYN packets to a backend server to check whether the port of the backend server is reachable.
	//
	// 	- **HTTP**: HTTP health checks simulate a process that uses a web browser to access resources by sending HEAD or GET requests to an instance. These requests are used to check whether the instance is healthy.
	//
	// example:
	//
	// TCP
	HealthCheckProtocol *string `json:"HealthCheckProtocol,omitempty" xml:"HealthCheckProtocol,omitempty"`
	HealthCheckReq      *string `json:"HealthCheckReq,omitempty" xml:"HealthCheckReq,omitempty"`
	// The number of times that an unhealthy backend server must consecutively pass health checks before it is declared healthy. In this case, the health status changes from **fail*	- to **success**.
	//
	// Valid values: **2*	- to **10**.
	//
	// example:
	//
	// 2
	HealthyThreshold *int32 `json:"HealthyThreshold,omitempty" xml:"HealthyThreshold,omitempty"`
	// The number of times that a healthy backend server must consecutively fail health checks before it is declared unhealthy. In this case, the health status changes from **success*	- to **fail**.
	//
	// Valid values: **2*	- to **10**.
	//
	// example:
	//
	// 2
	UnhealthyThreshold *int32 `json:"UnhealthyThreshold,omitempty" xml:"UnhealthyThreshold,omitempty"`
}

func (s ListServerGroupsResponseBodyServerGroupsHealthCheckConfig) String() string {
	return dara.Prettify(s)
}

func (s ListServerGroupsResponseBodyServerGroupsHealthCheckConfig) GoString() string {
	return s.String()
}

func (s *ListServerGroupsResponseBodyServerGroupsHealthCheckConfig) GetHealthCheckConnectPort() *int32 {
	return s.HealthCheckConnectPort
}

func (s *ListServerGroupsResponseBodyServerGroupsHealthCheckConfig) GetHealthCheckConnectTimeout() *int32 {
	return s.HealthCheckConnectTimeout
}

func (s *ListServerGroupsResponseBodyServerGroupsHealthCheckConfig) GetHealthCheckDomain() *string {
	return s.HealthCheckDomain
}

func (s *ListServerGroupsResponseBodyServerGroupsHealthCheckConfig) GetHealthCheckEnabled() *bool {
	return s.HealthCheckEnabled
}

func (s *ListServerGroupsResponseBodyServerGroupsHealthCheckConfig) GetHealthCheckExp() *string {
	return s.HealthCheckExp
}

func (s *ListServerGroupsResponseBodyServerGroupsHealthCheckConfig) GetHealthCheckHttpCode() []*string {
	return s.HealthCheckHttpCode
}

func (s *ListServerGroupsResponseBodyServerGroupsHealthCheckConfig) GetHealthCheckInterval() *int32 {
	return s.HealthCheckInterval
}

func (s *ListServerGroupsResponseBodyServerGroupsHealthCheckConfig) GetHealthCheckPath() *string {
	return s.HealthCheckPath
}

func (s *ListServerGroupsResponseBodyServerGroupsHealthCheckConfig) GetHealthCheckProtocol() *string {
	return s.HealthCheckProtocol
}

func (s *ListServerGroupsResponseBodyServerGroupsHealthCheckConfig) GetHealthCheckReq() *string {
	return s.HealthCheckReq
}

func (s *ListServerGroupsResponseBodyServerGroupsHealthCheckConfig) GetHealthyThreshold() *int32 {
	return s.HealthyThreshold
}

func (s *ListServerGroupsResponseBodyServerGroupsHealthCheckConfig) GetUnhealthyThreshold() *int32 {
	return s.UnhealthyThreshold
}

func (s *ListServerGroupsResponseBodyServerGroupsHealthCheckConfig) SetHealthCheckConnectPort(v int32) *ListServerGroupsResponseBodyServerGroupsHealthCheckConfig {
	s.HealthCheckConnectPort = &v
	return s
}

func (s *ListServerGroupsResponseBodyServerGroupsHealthCheckConfig) SetHealthCheckConnectTimeout(v int32) *ListServerGroupsResponseBodyServerGroupsHealthCheckConfig {
	s.HealthCheckConnectTimeout = &v
	return s
}

func (s *ListServerGroupsResponseBodyServerGroupsHealthCheckConfig) SetHealthCheckDomain(v string) *ListServerGroupsResponseBodyServerGroupsHealthCheckConfig {
	s.HealthCheckDomain = &v
	return s
}

func (s *ListServerGroupsResponseBodyServerGroupsHealthCheckConfig) SetHealthCheckEnabled(v bool) *ListServerGroupsResponseBodyServerGroupsHealthCheckConfig {
	s.HealthCheckEnabled = &v
	return s
}

func (s *ListServerGroupsResponseBodyServerGroupsHealthCheckConfig) SetHealthCheckExp(v string) *ListServerGroupsResponseBodyServerGroupsHealthCheckConfig {
	s.HealthCheckExp = &v
	return s
}

func (s *ListServerGroupsResponseBodyServerGroupsHealthCheckConfig) SetHealthCheckHttpCode(v []*string) *ListServerGroupsResponseBodyServerGroupsHealthCheckConfig {
	s.HealthCheckHttpCode = v
	return s
}

func (s *ListServerGroupsResponseBodyServerGroupsHealthCheckConfig) SetHealthCheckInterval(v int32) *ListServerGroupsResponseBodyServerGroupsHealthCheckConfig {
	s.HealthCheckInterval = &v
	return s
}

func (s *ListServerGroupsResponseBodyServerGroupsHealthCheckConfig) SetHealthCheckPath(v string) *ListServerGroupsResponseBodyServerGroupsHealthCheckConfig {
	s.HealthCheckPath = &v
	return s
}

func (s *ListServerGroupsResponseBodyServerGroupsHealthCheckConfig) SetHealthCheckProtocol(v string) *ListServerGroupsResponseBodyServerGroupsHealthCheckConfig {
	s.HealthCheckProtocol = &v
	return s
}

func (s *ListServerGroupsResponseBodyServerGroupsHealthCheckConfig) SetHealthCheckReq(v string) *ListServerGroupsResponseBodyServerGroupsHealthCheckConfig {
	s.HealthCheckReq = &v
	return s
}

func (s *ListServerGroupsResponseBodyServerGroupsHealthCheckConfig) SetHealthyThreshold(v int32) *ListServerGroupsResponseBodyServerGroupsHealthCheckConfig {
	s.HealthyThreshold = &v
	return s
}

func (s *ListServerGroupsResponseBodyServerGroupsHealthCheckConfig) SetUnhealthyThreshold(v int32) *ListServerGroupsResponseBodyServerGroupsHealthCheckConfig {
	s.UnhealthyThreshold = &v
	return s
}

func (s *ListServerGroupsResponseBodyServerGroupsHealthCheckConfig) Validate() error {
	return dara.Validate(s)
}

type ListServerGroupsResponseBodyServerGroupsTags struct {
	// The tag key. The tag key cannot be an empty string. The tag key can be up to 128 characters in length, and cannot start with `acs:` or `aliyun`. The tag key cannot contain `http://` or `https://`.
	//
	// example:
	//
	// testTagKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value. The tag value can be up to 256 characters in length and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// testTagValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListServerGroupsResponseBodyServerGroupsTags) String() string {
	return dara.Prettify(s)
}

func (s ListServerGroupsResponseBodyServerGroupsTags) GoString() string {
	return s.String()
}

func (s *ListServerGroupsResponseBodyServerGroupsTags) GetKey() *string {
	return s.Key
}

func (s *ListServerGroupsResponseBodyServerGroupsTags) GetValue() *string {
	return s.Value
}

func (s *ListServerGroupsResponseBodyServerGroupsTags) SetKey(v string) *ListServerGroupsResponseBodyServerGroupsTags {
	s.Key = &v
	return s
}

func (s *ListServerGroupsResponseBodyServerGroupsTags) SetValue(v string) *ListServerGroupsResponseBodyServerGroupsTags {
	s.Value = &v
	return s
}

func (s *ListServerGroupsResponseBodyServerGroupsTags) Validate() error {
	return dara.Validate(s)
}
