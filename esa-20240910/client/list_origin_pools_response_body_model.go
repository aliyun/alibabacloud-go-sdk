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
	// The list of origin address pools.
	OriginPools []*ListOriginPoolsResponseBodyOriginPools `json:"OriginPools,omitempty" xml:"OriginPools,omitempty" type:"Repeated"`
	// The current page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
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
	// The total number of entries.
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
	// Indicates whether the origin address pool is enabled. Valid values:
	//
	// - true: enabled.
	//
	// - false: disabled.
	//
	// example:
	//
	// false
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The origin address pool ID.
	//
	// example:
	//
	// 1038520525196928
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the origin address pool. The name must be unique within a site.
	//
	// example:
	//
	// pool1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The origin server information added to the origin address pool.
	Origins []*ListOriginPoolsResponseBodyOriginPoolsOrigins `json:"Origins,omitempty" xml:"Origins,omitempty" type:"Repeated"`
	// The domain name allocated to the origin IPAM pool. It can be used as the addresses of the origin server for records under the site.
	//
	// example:
	//
	// pool1.example.com
	RecordName *string `json:"RecordName,omitempty" xml:"RecordName,omitempty"`
	// The number of load balancing instances that reference this origin IPAM pool.
	//
	// example:
	//
	// 5
	ReferenceLBCount *int32 `json:"ReferenceLBCount,omitempty" xml:"ReferenceLBCount,omitempty"`
	// The reference information of the origin IPAM pool. An origin IPAM pool is considered referenced when it is configured for load balancing or configured as an origin server in a record.
	References *ListOriginPoolsResponseBodyOriginPoolsReferences `json:"References,omitempty" xml:"References,omitempty" type:"Struct"`
	// The ID of the site to which the origin address pool belongs.
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
	// The addresses of the origin server, such as www.example.com.
	//
	// example:
	//
	// www.example.com
	Address *string `json:"Address,omitempty" xml:"Address,omitempty"`
	// The authentication information. If the origin server is an OSS or S3 origin server that requires authentication, you must provide the authentication configuration information.
	AuthConf *ListOriginPoolsResponseBodyOriginPoolsOriginsAuthConf `json:"AuthConf,omitempty" xml:"AuthConf,omitempty" type:"Struct"`
	// Indicates whether the origin server is enabled. Valid values:
	//
	// - true: enabled.
	//
	// - false: disabled.
	//
	// example:
	//
	// true
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The request header included in back-to-origin requests. Only the Host header is supported.
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
	// The origin server ID.
	//
	// example:
	//
	// 997502094872132
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The IP protocol version used for back-to-origin requests. Valid values:
	//
	// - round_robin: the default policy. Randomly polls IPv4 or IPv6 origin servers.
	//
	// - ipv4_first: prioritizes IPv4 origin servers.
	//
	// - ipv6_first: prioritizes IPv6 origin servers.
	//
	// - follow: prioritizes the IP version used by the client.
	//
	// example:
	//
	// round_robin
	IpVersionPolicy *string `json:"IpVersionPolicy,omitempty" xml:"IpVersionPolicy,omitempty"`
	// The origin server name.
	//
	// example:
	//
	// origin1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The origin server type. Valid values:
	//
	// - ip_domain: an origin server that uses an IP address or a domain name.
	//
	// - OSS: an OSS origin server.
	//
	// - S3: an AWS S3 origin server.
	//
	// example:
	//
	// S3
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The weight. The value must be an integer from 0 to 100.
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
	// The AccessKey required for private authentication.
	//
	// example:
	//
	// yourAccessKeyID
	AccessKey *string `json:"AccessKey,omitempty" xml:"AccessKey,omitempty"`
	// The authentication type. Valid values:
	//
	// - public: public read/write. Used when the origin server is an OSS or S3 origin server with public read/write permissions.
	//
	// - private_same_account: private within the same account. Used when the origin server is an OSS origin server and the authentication type is private authentication within the same account.
	//
	// - private_cross_account: private across accounts. Used when the origin server is an OSS origin server and the authentication type is private authentication across accounts.
	//
	// - private: private. Used when the origin server is an S3 origin server and the authentication type is private.
	//
	// example:
	//
	// public
	AuthType *string `json:"AuthType,omitempty" xml:"AuthType,omitempty"`
	// The region of the origin server required when the origin server is an AWS S3 origin server.
	//
	// example:
	//
	// us-east-1
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The SecretKey required for private authentication.
	//
	// example:
	//
	// yourAccessKeySecret
	SecretKey *string `json:"SecretKey,omitempty" xml:"SecretKey,omitempty"`
	// The signature version required when the origin server is an AWS S3 origin server.
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
	// The list of Layer 7 records that use this origin address pool as the origin server.
	DnsRecords []*ListOriginPoolsResponseBodyOriginPoolsReferencesDnsRecords `json:"DnsRecords,omitempty" xml:"DnsRecords,omitempty" type:"Repeated"`
	// The list of Layer 4 records that use this origin address pool as the origin server.
	IPARecords []*ListOriginPoolsResponseBodyOriginPoolsReferencesIPARecords `json:"IPARecords,omitempty" xml:"IPARecords,omitempty" type:"Repeated"`
	// The list of load balancing instances that use this origin IPAM pool.
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
	// The record ID.
	//
	// example:
	//
	// 1042852886352704
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The record name.
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
	// The record ID.
	//
	// example:
	//
	// 1042852886352704
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The record name.
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
	// The load balancing instance ID.
	//
	// example:
	//
	// 998740660522624
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the load balancing instance.
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
