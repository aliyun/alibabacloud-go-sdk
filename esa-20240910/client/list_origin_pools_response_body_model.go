// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOriginPoolsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOriginPools(v []*ListOriginPoolsResponseBodyOriginPools) *ListOriginPoolsResponseBody
	GetOriginPools() []*ListOriginPoolsResponseBodyOriginPools
	SetPageNumber(v int32) *ListOriginPoolsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListOriginPoolsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListOriginPoolsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListOriginPoolsResponseBody
	GetTotalCount() *int32
	SetTotalPage(v int32) *ListOriginPoolsResponseBody
	GetTotalPage() *int32
}

type ListOriginPoolsResponseBody struct {
	// A list of origin pools.
	OriginPools []*ListOriginPoolsResponseBodyOriginPools `json:"OriginPools,omitempty" xml:"OriginPools,omitempty" type:"Repeated"`
	// The current page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of items returned per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 15C66C7B-671A-4297-9187-2C4477247A74
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of origin pools found.
	//
	// example:
	//
	// 16
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The total number of pages.
	//
	// example:
	//
	// 10
	TotalPage *int32 `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
}

func (s ListOriginPoolsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListOriginPoolsResponseBody) GoString() string {
	return s.String()
}

func (s *ListOriginPoolsResponseBody) GetOriginPools() []*ListOriginPoolsResponseBodyOriginPools {
	return s.OriginPools
}

func (s *ListOriginPoolsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListOriginPoolsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListOriginPoolsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListOriginPoolsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListOriginPoolsResponseBody) GetTotalPage() *int32 {
	return s.TotalPage
}

func (s *ListOriginPoolsResponseBody) SetOriginPools(v []*ListOriginPoolsResponseBodyOriginPools) *ListOriginPoolsResponseBody {
	s.OriginPools = v
	return s
}

func (s *ListOriginPoolsResponseBody) SetPageNumber(v int32) *ListOriginPoolsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListOriginPoolsResponseBody) SetPageSize(v int32) *ListOriginPoolsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListOriginPoolsResponseBody) SetRequestId(v string) *ListOriginPoolsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListOriginPoolsResponseBody) SetTotalCount(v int32) *ListOriginPoolsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListOriginPoolsResponseBody) SetTotalPage(v int32) *ListOriginPoolsResponseBody {
	s.TotalPage = &v
	return s
}

func (s *ListOriginPoolsResponseBody) Validate() error {
	if s.OriginPools != nil {
		for _, item := range s.OriginPools {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListOriginPoolsResponseBodyOriginPools struct {
	// Indicates whether the origin pool is enabled.
	//
	// - `true`: Enabled.
	//
	// - `false`: Disabled.
	//
	// example:
	//
	// false
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The ID of the origin pool.
	//
	// example:
	//
	// 1038520525196928
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the origin pool. The name must be unique within a site.
	//
	// example:
	//
	// pool1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The origins in the origin pool.
	Origins []*ListOriginPoolsResponseBodyOriginPoolsOrigins `json:"Origins,omitempty" xml:"Origins,omitempty" type:"Repeated"`
	// The domain name assigned to the origin pool. This can be used as the origin address for records within the site.
	//
	// example:
	//
	// pool1.example.com
	RecordName *string `json:"RecordName,omitempty" xml:"RecordName,omitempty"`
	// The number of load balancers that reference this origin pool.
	//
	// example:
	//
	// 5
	ReferenceLBCount *int32 `json:"ReferenceLBCount,omitempty" xml:"ReferenceLBCount,omitempty"`
	// The resources that reference this origin pool. An origin pool is considered referenced when it is configured as an origin for a load balancer or a record.
	References *ListOriginPoolsResponseBodyOriginPoolsReferences `json:"References,omitempty" xml:"References,omitempty" type:"Struct"`
	// The ID of the site to which the origin pool belongs.
	//
	// example:
	//
	// 216558609793952
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s ListOriginPoolsResponseBodyOriginPools) String() string {
	return dara.Prettify(s)
}

func (s ListOriginPoolsResponseBodyOriginPools) GoString() string {
	return s.String()
}

func (s *ListOriginPoolsResponseBodyOriginPools) GetEnabled() *bool {
	return s.Enabled
}

func (s *ListOriginPoolsResponseBodyOriginPools) GetId() *int64 {
	return s.Id
}

func (s *ListOriginPoolsResponseBodyOriginPools) GetName() *string {
	return s.Name
}

func (s *ListOriginPoolsResponseBodyOriginPools) GetOrigins() []*ListOriginPoolsResponseBodyOriginPoolsOrigins {
	return s.Origins
}

func (s *ListOriginPoolsResponseBodyOriginPools) GetRecordName() *string {
	return s.RecordName
}

func (s *ListOriginPoolsResponseBodyOriginPools) GetReferenceLBCount() *int32 {
	return s.ReferenceLBCount
}

func (s *ListOriginPoolsResponseBodyOriginPools) GetReferences() *ListOriginPoolsResponseBodyOriginPoolsReferences {
	return s.References
}

func (s *ListOriginPoolsResponseBodyOriginPools) GetSiteId() *int64 {
	return s.SiteId
}

func (s *ListOriginPoolsResponseBodyOriginPools) SetEnabled(v bool) *ListOriginPoolsResponseBodyOriginPools {
	s.Enabled = &v
	return s
}

func (s *ListOriginPoolsResponseBodyOriginPools) SetId(v int64) *ListOriginPoolsResponseBodyOriginPools {
	s.Id = &v
	return s
}

func (s *ListOriginPoolsResponseBodyOriginPools) SetName(v string) *ListOriginPoolsResponseBodyOriginPools {
	s.Name = &v
	return s
}

func (s *ListOriginPoolsResponseBodyOriginPools) SetOrigins(v []*ListOriginPoolsResponseBodyOriginPoolsOrigins) *ListOriginPoolsResponseBodyOriginPools {
	s.Origins = v
	return s
}

func (s *ListOriginPoolsResponseBodyOriginPools) SetRecordName(v string) *ListOriginPoolsResponseBodyOriginPools {
	s.RecordName = &v
	return s
}

func (s *ListOriginPoolsResponseBodyOriginPools) SetReferenceLBCount(v int32) *ListOriginPoolsResponseBodyOriginPools {
	s.ReferenceLBCount = &v
	return s
}

func (s *ListOriginPoolsResponseBodyOriginPools) SetReferences(v *ListOriginPoolsResponseBodyOriginPoolsReferences) *ListOriginPoolsResponseBodyOriginPools {
	s.References = v
	return s
}

func (s *ListOriginPoolsResponseBodyOriginPools) SetSiteId(v int64) *ListOriginPoolsResponseBodyOriginPools {
	s.SiteId = &v
	return s
}

func (s *ListOriginPoolsResponseBodyOriginPools) Validate() error {
	if s.Origins != nil {
		for _, item := range s.Origins {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.References != nil {
		if err := s.References.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListOriginPoolsResponseBodyOriginPoolsOrigins struct {
	// The address of the origin, such as `www.example.com`.
	//
	// example:
	//
	// www.example.com
	Address *string `json:"Address,omitempty" xml:"Address,omitempty"`
	// The authentication configuration. This parameter is required if the origin is an OSS or AWS S3 bucket that requires authentication.
	AuthConf *ListOriginPoolsResponseBodyOriginPoolsOriginsAuthConf `json:"AuthConf,omitempty" xml:"AuthConf,omitempty" type:"Struct"`
	// Indicates whether the origin is enabled.
	//
	// - `true`: Enabled.
	//
	// - `false`: Disabled.
	//
	// example:
	//
	// true
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The request header to include in origin fetches. Only the `Host` header is supported.
	//
	// example:
	//
	// {
	//
	//         "Host": [
	//
	//           "example.com"
	//
	//         ]
	//
	//       }
	Header interface{} `json:"Header,omitempty" xml:"Header,omitempty"`
	// The ID of the origin.
	//
	// example:
	//
	// 997502094872132
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The IP protocol version used for origin fetch requests. Valid values:
	//
	// - `round_robin`: Randomly selects an IPv4 or IPv6 origin.
	//
	// - `ipv4_first`: Prioritizes IPv4 origins.
	//
	// - `ipv6_first`: Prioritizes IPv6 origins.
	//
	// - `follow`: Follows the IP version used by the client.
	//
	// example:
	//
	// round_robin
	IpVersionPolicy *string `json:"IpVersionPolicy,omitempty" xml:"IpVersionPolicy,omitempty"`
	// The name of the origin.
	//
	// example:
	//
	// origin1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The type of the origin. Valid values:
	//
	// - `ip_domain`: An origin with an IP address or domain name.
	//
	// - `OSS`: An Alibaba Cloud Object Storage Service (OSS) origin.
	//
	// - `S3`: An AWS S3 origin.
	//
	// example:
	//
	// S3
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The weight of the origin, an integer from 0 to 100.
	//
	// example:
	//
	// 50
	Weight *int32 `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s ListOriginPoolsResponseBodyOriginPoolsOrigins) String() string {
	return dara.Prettify(s)
}

func (s ListOriginPoolsResponseBodyOriginPoolsOrigins) GoString() string {
	return s.String()
}

func (s *ListOriginPoolsResponseBodyOriginPoolsOrigins) GetAddress() *string {
	return s.Address
}

func (s *ListOriginPoolsResponseBodyOriginPoolsOrigins) GetAuthConf() *ListOriginPoolsResponseBodyOriginPoolsOriginsAuthConf {
	return s.AuthConf
}

func (s *ListOriginPoolsResponseBodyOriginPoolsOrigins) GetEnabled() *bool {
	return s.Enabled
}

func (s *ListOriginPoolsResponseBodyOriginPoolsOrigins) GetHeader() interface{} {
	return s.Header
}

func (s *ListOriginPoolsResponseBodyOriginPoolsOrigins) GetId() *int64 {
	return s.Id
}

func (s *ListOriginPoolsResponseBodyOriginPoolsOrigins) GetIpVersionPolicy() *string {
	return s.IpVersionPolicy
}

func (s *ListOriginPoolsResponseBodyOriginPoolsOrigins) GetName() *string {
	return s.Name
}

func (s *ListOriginPoolsResponseBodyOriginPoolsOrigins) GetType() *string {
	return s.Type
}

func (s *ListOriginPoolsResponseBodyOriginPoolsOrigins) GetWeight() *int32 {
	return s.Weight
}

func (s *ListOriginPoolsResponseBodyOriginPoolsOrigins) SetAddress(v string) *ListOriginPoolsResponseBodyOriginPoolsOrigins {
	s.Address = &v
	return s
}

func (s *ListOriginPoolsResponseBodyOriginPoolsOrigins) SetAuthConf(v *ListOriginPoolsResponseBodyOriginPoolsOriginsAuthConf) *ListOriginPoolsResponseBodyOriginPoolsOrigins {
	s.AuthConf = v
	return s
}

func (s *ListOriginPoolsResponseBodyOriginPoolsOrigins) SetEnabled(v bool) *ListOriginPoolsResponseBodyOriginPoolsOrigins {
	s.Enabled = &v
	return s
}

func (s *ListOriginPoolsResponseBodyOriginPoolsOrigins) SetHeader(v interface{}) *ListOriginPoolsResponseBodyOriginPoolsOrigins {
	s.Header = v
	return s
}

func (s *ListOriginPoolsResponseBodyOriginPoolsOrigins) SetId(v int64) *ListOriginPoolsResponseBodyOriginPoolsOrigins {
	s.Id = &v
	return s
}

func (s *ListOriginPoolsResponseBodyOriginPoolsOrigins) SetIpVersionPolicy(v string) *ListOriginPoolsResponseBodyOriginPoolsOrigins {
	s.IpVersionPolicy = &v
	return s
}

func (s *ListOriginPoolsResponseBodyOriginPoolsOrigins) SetName(v string) *ListOriginPoolsResponseBodyOriginPoolsOrigins {
	s.Name = &v
	return s
}

func (s *ListOriginPoolsResponseBodyOriginPoolsOrigins) SetType(v string) *ListOriginPoolsResponseBodyOriginPoolsOrigins {
	s.Type = &v
	return s
}

func (s *ListOriginPoolsResponseBodyOriginPoolsOrigins) SetWeight(v int32) *ListOriginPoolsResponseBodyOriginPoolsOrigins {
	s.Weight = &v
	return s
}

func (s *ListOriginPoolsResponseBodyOriginPoolsOrigins) Validate() error {
	if s.AuthConf != nil {
		if err := s.AuthConf.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListOriginPoolsResponseBodyOriginPoolsOriginsAuthConf struct {
	// The Access Key ID required for private authentication.
	//
	// example:
	//
	// yourAccessKeyID
	AccessKey *string `json:"AccessKey,omitempty" xml:"AccessKey,omitempty"`
	// The authentication type. Valid values:
	//
	// - `public`: Public read/write. Use for OSS or AWS S3 origins that have public access permissions.
	//
	// - `private_same_account`: Private, same account. Use for OSS origins that use private authentication within the same account.
	//
	// - `private_cross_account`: Private, cross-account. Use for OSS origins that use private authentication from a different account.
	//
	// - `private`: Private. Use for AWS S3 origins that require private authentication.
	//
	// example:
	//
	// public
	AuthType *string `json:"AuthType,omitempty" xml:"AuthType,omitempty"`
	// The region of the origin. Required for AWS S3 origins.
	//
	// example:
	//
	// us-east-1
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The Secret Access Key required for private authentication.
	//
	// example:
	//
	// yourAccessKeySecret
	SecretKey *string `json:"SecretKey,omitempty" xml:"SecretKey,omitempty"`
	// The signature version. Required for AWS S3 origins.
	//
	// example:
	//
	// v2
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s ListOriginPoolsResponseBodyOriginPoolsOriginsAuthConf) String() string {
	return dara.Prettify(s)
}

func (s ListOriginPoolsResponseBodyOriginPoolsOriginsAuthConf) GoString() string {
	return s.String()
}

func (s *ListOriginPoolsResponseBodyOriginPoolsOriginsAuthConf) GetAccessKey() *string {
	return s.AccessKey
}

func (s *ListOriginPoolsResponseBodyOriginPoolsOriginsAuthConf) GetAuthType() *string {
	return s.AuthType
}

func (s *ListOriginPoolsResponseBodyOriginPoolsOriginsAuthConf) GetRegion() *string {
	return s.Region
}

func (s *ListOriginPoolsResponseBodyOriginPoolsOriginsAuthConf) GetSecretKey() *string {
	return s.SecretKey
}

func (s *ListOriginPoolsResponseBodyOriginPoolsOriginsAuthConf) GetVersion() *string {
	return s.Version
}

func (s *ListOriginPoolsResponseBodyOriginPoolsOriginsAuthConf) SetAccessKey(v string) *ListOriginPoolsResponseBodyOriginPoolsOriginsAuthConf {
	s.AccessKey = &v
	return s
}

func (s *ListOriginPoolsResponseBodyOriginPoolsOriginsAuthConf) SetAuthType(v string) *ListOriginPoolsResponseBodyOriginPoolsOriginsAuthConf {
	s.AuthType = &v
	return s
}

func (s *ListOriginPoolsResponseBodyOriginPoolsOriginsAuthConf) SetRegion(v string) *ListOriginPoolsResponseBodyOriginPoolsOriginsAuthConf {
	s.Region = &v
	return s
}

func (s *ListOriginPoolsResponseBodyOriginPoolsOriginsAuthConf) SetSecretKey(v string) *ListOriginPoolsResponseBodyOriginPoolsOriginsAuthConf {
	s.SecretKey = &v
	return s
}

func (s *ListOriginPoolsResponseBodyOriginPoolsOriginsAuthConf) SetVersion(v string) *ListOriginPoolsResponseBodyOriginPoolsOriginsAuthConf {
	s.Version = &v
	return s
}

func (s *ListOriginPoolsResponseBodyOriginPoolsOriginsAuthConf) Validate() error {
	return dara.Validate(s)
}

type ListOriginPoolsResponseBodyOriginPoolsReferences struct {
	// A list of Layer 7 records that use this origin pool as an origin.
	DnsRecords []*ListOriginPoolsResponseBodyOriginPoolsReferencesDnsRecords `json:"DnsRecords,omitempty" xml:"DnsRecords,omitempty" type:"Repeated"`
	// A list of Layer 4 records that use this origin pool as an origin.
	IPARecords []*ListOriginPoolsResponseBodyOriginPoolsReferencesIPARecords `json:"IPARecords,omitempty" xml:"IPARecords,omitempty" type:"Repeated"`
	// A list of load balancers that use this origin pool.
	LoadBalancers []*ListOriginPoolsResponseBodyOriginPoolsReferencesLoadBalancers `json:"LoadBalancers,omitempty" xml:"LoadBalancers,omitempty" type:"Repeated"`
}

func (s ListOriginPoolsResponseBodyOriginPoolsReferences) String() string {
	return dara.Prettify(s)
}

func (s ListOriginPoolsResponseBodyOriginPoolsReferences) GoString() string {
	return s.String()
}

func (s *ListOriginPoolsResponseBodyOriginPoolsReferences) GetDnsRecords() []*ListOriginPoolsResponseBodyOriginPoolsReferencesDnsRecords {
	return s.DnsRecords
}

func (s *ListOriginPoolsResponseBodyOriginPoolsReferences) GetIPARecords() []*ListOriginPoolsResponseBodyOriginPoolsReferencesIPARecords {
	return s.IPARecords
}

func (s *ListOriginPoolsResponseBodyOriginPoolsReferences) GetLoadBalancers() []*ListOriginPoolsResponseBodyOriginPoolsReferencesLoadBalancers {
	return s.LoadBalancers
}

func (s *ListOriginPoolsResponseBodyOriginPoolsReferences) SetDnsRecords(v []*ListOriginPoolsResponseBodyOriginPoolsReferencesDnsRecords) *ListOriginPoolsResponseBodyOriginPoolsReferences {
	s.DnsRecords = v
	return s
}

func (s *ListOriginPoolsResponseBodyOriginPoolsReferences) SetIPARecords(v []*ListOriginPoolsResponseBodyOriginPoolsReferencesIPARecords) *ListOriginPoolsResponseBodyOriginPoolsReferences {
	s.IPARecords = v
	return s
}

func (s *ListOriginPoolsResponseBodyOriginPoolsReferences) SetLoadBalancers(v []*ListOriginPoolsResponseBodyOriginPoolsReferencesLoadBalancers) *ListOriginPoolsResponseBodyOriginPoolsReferences {
	s.LoadBalancers = v
	return s
}

func (s *ListOriginPoolsResponseBodyOriginPoolsReferences) Validate() error {
	if s.DnsRecords != nil {
		for _, item := range s.DnsRecords {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.IPARecords != nil {
		for _, item := range s.IPARecords {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.LoadBalancers != nil {
		for _, item := range s.LoadBalancers {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListOriginPoolsResponseBodyOriginPoolsReferencesDnsRecords struct {
	// The ID of the record.
	//
	// example:
	//
	// 1042852886352704
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the record.
	//
	// example:
	//
	// www.example.com
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ListOriginPoolsResponseBodyOriginPoolsReferencesDnsRecords) String() string {
	return dara.Prettify(s)
}

func (s ListOriginPoolsResponseBodyOriginPoolsReferencesDnsRecords) GoString() string {
	return s.String()
}

func (s *ListOriginPoolsResponseBodyOriginPoolsReferencesDnsRecords) GetId() *int64 {
	return s.Id
}

func (s *ListOriginPoolsResponseBodyOriginPoolsReferencesDnsRecords) GetName() *string {
	return s.Name
}

func (s *ListOriginPoolsResponseBodyOriginPoolsReferencesDnsRecords) SetId(v int64) *ListOriginPoolsResponseBodyOriginPoolsReferencesDnsRecords {
	s.Id = &v
	return s
}

func (s *ListOriginPoolsResponseBodyOriginPoolsReferencesDnsRecords) SetName(v string) *ListOriginPoolsResponseBodyOriginPoolsReferencesDnsRecords {
	s.Name = &v
	return s
}

func (s *ListOriginPoolsResponseBodyOriginPoolsReferencesDnsRecords) Validate() error {
	return dara.Validate(s)
}

type ListOriginPoolsResponseBodyOriginPoolsReferencesIPARecords struct {
	// The ID of the record.
	//
	// example:
	//
	// 1042852886352704
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the record.
	//
	// example:
	//
	// ipa.example.com
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ListOriginPoolsResponseBodyOriginPoolsReferencesIPARecords) String() string {
	return dara.Prettify(s)
}

func (s ListOriginPoolsResponseBodyOriginPoolsReferencesIPARecords) GoString() string {
	return s.String()
}

func (s *ListOriginPoolsResponseBodyOriginPoolsReferencesIPARecords) GetId() *int64 {
	return s.Id
}

func (s *ListOriginPoolsResponseBodyOriginPoolsReferencesIPARecords) GetName() *string {
	return s.Name
}

func (s *ListOriginPoolsResponseBodyOriginPoolsReferencesIPARecords) SetId(v int64) *ListOriginPoolsResponseBodyOriginPoolsReferencesIPARecords {
	s.Id = &v
	return s
}

func (s *ListOriginPoolsResponseBodyOriginPoolsReferencesIPARecords) SetName(v string) *ListOriginPoolsResponseBodyOriginPoolsReferencesIPARecords {
	s.Name = &v
	return s
}

func (s *ListOriginPoolsResponseBodyOriginPoolsReferencesIPARecords) Validate() error {
	return dara.Validate(s)
}

type ListOriginPoolsResponseBodyOriginPoolsReferencesLoadBalancers struct {
	// The ID of the load balancer.
	//
	// example:
	//
	// 998740660522624
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the load balancer.
	//
	// example:
	//
	// lb1.example.com
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ListOriginPoolsResponseBodyOriginPoolsReferencesLoadBalancers) String() string {
	return dara.Prettify(s)
}

func (s ListOriginPoolsResponseBodyOriginPoolsReferencesLoadBalancers) GoString() string {
	return s.String()
}

func (s *ListOriginPoolsResponseBodyOriginPoolsReferencesLoadBalancers) GetId() *int64 {
	return s.Id
}

func (s *ListOriginPoolsResponseBodyOriginPoolsReferencesLoadBalancers) GetName() *string {
	return s.Name
}

func (s *ListOriginPoolsResponseBodyOriginPoolsReferencesLoadBalancers) SetId(v int64) *ListOriginPoolsResponseBodyOriginPoolsReferencesLoadBalancers {
	s.Id = &v
	return s
}

func (s *ListOriginPoolsResponseBodyOriginPoolsReferencesLoadBalancers) SetName(v string) *ListOriginPoolsResponseBodyOriginPoolsReferencesLoadBalancers {
	s.Name = &v
	return s
}

func (s *ListOriginPoolsResponseBodyOriginPoolsReferencesLoadBalancers) Validate() error {
	return dara.Validate(s)
}
