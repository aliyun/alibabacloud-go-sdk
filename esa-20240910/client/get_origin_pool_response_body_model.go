// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOriginPoolResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEnabled(v bool) *GetOriginPoolResponseBody
	GetEnabled() *bool
	SetId(v int64) *GetOriginPoolResponseBody
	GetId() *int64
	SetName(v string) *GetOriginPoolResponseBody
	GetName() *string
	SetOrigins(v []*GetOriginPoolResponseBodyOrigins) *GetOriginPoolResponseBody
	GetOrigins() []*GetOriginPoolResponseBodyOrigins
	SetRecordName(v string) *GetOriginPoolResponseBody
	GetRecordName() *string
	SetReferenceLBCount(v int32) *GetOriginPoolResponseBody
	GetReferenceLBCount() *int32
	SetReferences(v *GetOriginPoolResponseBodyReferences) *GetOriginPoolResponseBody
	GetReferences() *GetOriginPoolResponseBodyReferences
	SetRequestId(v string) *GetOriginPoolResponseBody
	GetRequestId() *string
	SetSiteId(v int64) *GetOriginPoolResponseBody
	GetSiteId() *int64
}

type GetOriginPoolResponseBody struct {
	// Specifies if the origin pool is enabled.
	//
	// - true: The origin pool is enabled.
	//
	// - false: The origin pool is disabled.
	//
	// example:
	//
	// true
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The origin pool ID.
	//
	// example:
	//
	// 103852052519****
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the origin pool. The name must be unique within a site.
	//
	// example:
	//
	// pool1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The origins in the origin pool.
	Origins []*GetOriginPoolResponseBodyOrigins `json:"Origins,omitempty" xml:"Origins,omitempty" type:"Repeated"`
	// The domain name assigned to the origin pool. It serves as the origin address for records within the site.
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
	// Resources that reference the origin pool. An origin pool is referenced if a load balancer or record uses it as an origin.
	References *GetOriginPoolResponseBodyReferences `json:"References,omitempty" xml:"References,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the site that contains the origin pool.
	//
	// example:
	//
	// 21655860979****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s GetOriginPoolResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetOriginPoolResponseBody) GoString() string {
	return s.String()
}

func (s *GetOriginPoolResponseBody) GetEnabled() *bool {
	return s.Enabled
}

func (s *GetOriginPoolResponseBody) GetId() *int64 {
	return s.Id
}

func (s *GetOriginPoolResponseBody) GetName() *string {
	return s.Name
}

func (s *GetOriginPoolResponseBody) GetOrigins() []*GetOriginPoolResponseBodyOrigins {
	return s.Origins
}

func (s *GetOriginPoolResponseBody) GetRecordName() *string {
	return s.RecordName
}

func (s *GetOriginPoolResponseBody) GetReferenceLBCount() *int32 {
	return s.ReferenceLBCount
}

func (s *GetOriginPoolResponseBody) GetReferences() *GetOriginPoolResponseBodyReferences {
	return s.References
}

func (s *GetOriginPoolResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetOriginPoolResponseBody) GetSiteId() *int64 {
	return s.SiteId
}

func (s *GetOriginPoolResponseBody) SetEnabled(v bool) *GetOriginPoolResponseBody {
	s.Enabled = &v
	return s
}

func (s *GetOriginPoolResponseBody) SetId(v int64) *GetOriginPoolResponseBody {
	s.Id = &v
	return s
}

func (s *GetOriginPoolResponseBody) SetName(v string) *GetOriginPoolResponseBody {
	s.Name = &v
	return s
}

func (s *GetOriginPoolResponseBody) SetOrigins(v []*GetOriginPoolResponseBodyOrigins) *GetOriginPoolResponseBody {
	s.Origins = v
	return s
}

func (s *GetOriginPoolResponseBody) SetRecordName(v string) *GetOriginPoolResponseBody {
	s.RecordName = &v
	return s
}

func (s *GetOriginPoolResponseBody) SetReferenceLBCount(v int32) *GetOriginPoolResponseBody {
	s.ReferenceLBCount = &v
	return s
}

func (s *GetOriginPoolResponseBody) SetReferences(v *GetOriginPoolResponseBodyReferences) *GetOriginPoolResponseBody {
	s.References = v
	return s
}

func (s *GetOriginPoolResponseBody) SetRequestId(v string) *GetOriginPoolResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOriginPoolResponseBody) SetSiteId(v int64) *GetOriginPoolResponseBody {
	s.SiteId = &v
	return s
}

func (s *GetOriginPoolResponseBody) Validate() error {
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

type GetOriginPoolResponseBodyOrigins struct {
	// The origin address. For example, www\\.example.com.
	//
	// example:
	//
	// www.example.com
	Address *string `json:"Address,omitempty" xml:"Address,omitempty"`
	// The authentication configuration. This parameter is required if the origin is an OSS or S3 bucket that requires authentication.
	AuthConf *GetOriginPoolResponseBodyOriginsAuthConf `json:"AuthConf,omitempty" xml:"AuthConf,omitempty" type:"Struct"`
	// Specifies if the origin is enabled.
	//
	// - true: The origin is enabled.
	//
	// - false: The origin is disabled.
	//
	// example:
	//
	// true
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The request header to include in origin requests. Only the Host header is supported.
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
	// The origin ID.
	//
	// example:
	//
	// 99750209487****
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The IP version policy for origin requests.
	//
	// - round_robin: Default policy. Routes requests to IPv4 and IPv6 origins on a round-robin basis.
	//
	// - ipv4_first: Prioritizes IPv4. Routes requests to IPv4 origins first.
	//
	// - ipv6_first: Prioritizes IPv6. Routes requests to IPv6 origins first.
	//
	// - follow: Matches the client\\"s IP version. The origin request uses the same IP version as the incoming request.
	//
	// example:
	//
	// round_robin
	IpVersionPolicy *string `json:"IpVersionPolicy,omitempty" xml:"IpVersionPolicy,omitempty"`
	// The origin name.
	//
	// example:
	//
	// origin1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The type of the origin.
	//
	// - ip_domain: An IP address or a domain name.
	//
	// - OSS: An OSS bucket.
	//
	// - S3: An AWS S3 bucket.
	//
	// example:
	//
	// ip_domain
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The weight of the origin. Must be an integer from 0 to 100.
	//
	// example:
	//
	// 50
	Weight *int32 `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s GetOriginPoolResponseBodyOrigins) String() string {
	return dara.Prettify(s)
}

func (s GetOriginPoolResponseBodyOrigins) GoString() string {
	return s.String()
}

func (s *GetOriginPoolResponseBodyOrigins) GetAddress() *string {
	return s.Address
}

func (s *GetOriginPoolResponseBodyOrigins) GetAuthConf() *GetOriginPoolResponseBodyOriginsAuthConf {
	return s.AuthConf
}

func (s *GetOriginPoolResponseBodyOrigins) GetEnabled() *bool {
	return s.Enabled
}

func (s *GetOriginPoolResponseBodyOrigins) GetHeader() interface{} {
	return s.Header
}

func (s *GetOriginPoolResponseBodyOrigins) GetId() *int64 {
	return s.Id
}

func (s *GetOriginPoolResponseBodyOrigins) GetIpVersionPolicy() *string {
	return s.IpVersionPolicy
}

func (s *GetOriginPoolResponseBodyOrigins) GetName() *string {
	return s.Name
}

func (s *GetOriginPoolResponseBodyOrigins) GetType() *string {
	return s.Type
}

func (s *GetOriginPoolResponseBodyOrigins) GetWeight() *int32 {
	return s.Weight
}

func (s *GetOriginPoolResponseBodyOrigins) SetAddress(v string) *GetOriginPoolResponseBodyOrigins {
	s.Address = &v
	return s
}

func (s *GetOriginPoolResponseBodyOrigins) SetAuthConf(v *GetOriginPoolResponseBodyOriginsAuthConf) *GetOriginPoolResponseBodyOrigins {
	s.AuthConf = v
	return s
}

func (s *GetOriginPoolResponseBodyOrigins) SetEnabled(v bool) *GetOriginPoolResponseBodyOrigins {
	s.Enabled = &v
	return s
}

func (s *GetOriginPoolResponseBodyOrigins) SetHeader(v interface{}) *GetOriginPoolResponseBodyOrigins {
	s.Header = v
	return s
}

func (s *GetOriginPoolResponseBodyOrigins) SetId(v int64) *GetOriginPoolResponseBodyOrigins {
	s.Id = &v
	return s
}

func (s *GetOriginPoolResponseBodyOrigins) SetIpVersionPolicy(v string) *GetOriginPoolResponseBodyOrigins {
	s.IpVersionPolicy = &v
	return s
}

func (s *GetOriginPoolResponseBodyOrigins) SetName(v string) *GetOriginPoolResponseBodyOrigins {
	s.Name = &v
	return s
}

func (s *GetOriginPoolResponseBodyOrigins) SetType(v string) *GetOriginPoolResponseBodyOrigins {
	s.Type = &v
	return s
}

func (s *GetOriginPoolResponseBodyOrigins) SetWeight(v int32) *GetOriginPoolResponseBodyOrigins {
	s.Weight = &v
	return s
}

func (s *GetOriginPoolResponseBodyOrigins) Validate() error {
	if s.AuthConf != nil {
		if err := s.AuthConf.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetOriginPoolResponseBodyOriginsAuthConf struct {
	// The AccessKey ID. This parameter is required if `AuthType` is set to `private_cross_account` or `private`.
	//
	// example:
	//
	// yourAccessKeyID
	AccessKey *string `json:"AccessKey,omitempty" xml:"AccessKey,omitempty"`
	// The authentication type.
	//
	// - public: Public access. For OSS or S3 origins with public read access.
	//
	// - private_same_account: Private, same account. For private OSS origins in the same account.
	//
	// - private_cross_account: Private, cross-account. For private OSS origins in a different account.
	//
	// - private: Private. For private S3 origins.
	//
	// example:
	//
	// public
	AuthType *string `json:"AuthType,omitempty" xml:"AuthType,omitempty"`
	// The AWS Region of the origin. Required if the origin is an AWS S3 bucket.
	//
	// example:
	//
	// us-east-1
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The AccessKey secret. This parameter is required if `AuthType` is set to `private_cross_account` or `private`.
	//
	// example:
	//
	// yourAccessKeySecret
	SecretKey *string `json:"SecretKey,omitempty" xml:"SecretKey,omitempty"`
	// The signature version. Required if the origin is an AWS S3 bucket.
	//
	// example:
	//
	// v4
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s GetOriginPoolResponseBodyOriginsAuthConf) String() string {
	return dara.Prettify(s)
}

func (s GetOriginPoolResponseBodyOriginsAuthConf) GoString() string {
	return s.String()
}

func (s *GetOriginPoolResponseBodyOriginsAuthConf) GetAccessKey() *string {
	return s.AccessKey
}

func (s *GetOriginPoolResponseBodyOriginsAuthConf) GetAuthType() *string {
	return s.AuthType
}

func (s *GetOriginPoolResponseBodyOriginsAuthConf) GetRegion() *string {
	return s.Region
}

func (s *GetOriginPoolResponseBodyOriginsAuthConf) GetSecretKey() *string {
	return s.SecretKey
}

func (s *GetOriginPoolResponseBodyOriginsAuthConf) GetVersion() *string {
	return s.Version
}

func (s *GetOriginPoolResponseBodyOriginsAuthConf) SetAccessKey(v string) *GetOriginPoolResponseBodyOriginsAuthConf {
	s.AccessKey = &v
	return s
}

func (s *GetOriginPoolResponseBodyOriginsAuthConf) SetAuthType(v string) *GetOriginPoolResponseBodyOriginsAuthConf {
	s.AuthType = &v
	return s
}

func (s *GetOriginPoolResponseBodyOriginsAuthConf) SetRegion(v string) *GetOriginPoolResponseBodyOriginsAuthConf {
	s.Region = &v
	return s
}

func (s *GetOriginPoolResponseBodyOriginsAuthConf) SetSecretKey(v string) *GetOriginPoolResponseBodyOriginsAuthConf {
	s.SecretKey = &v
	return s
}

func (s *GetOriginPoolResponseBodyOriginsAuthConf) SetVersion(v string) *GetOriginPoolResponseBodyOriginsAuthConf {
	s.Version = &v
	return s
}

func (s *GetOriginPoolResponseBodyOriginsAuthConf) Validate() error {
	return dara.Validate(s)
}

type GetOriginPoolResponseBodyReferences struct {
	// The Layer 7 records that use this origin pool as their origin.
	DnsRecords []*GetOriginPoolResponseBodyReferencesDnsRecords `json:"DnsRecords,omitempty" xml:"DnsRecords,omitempty" type:"Repeated"`
	// The Layer 4 records that use this origin pool as their origin.
	IPARecords []*GetOriginPoolResponseBodyReferencesIPARecords `json:"IPARecords,omitempty" xml:"IPARecords,omitempty" type:"Repeated"`
	// The load balancers that use this origin pool.
	LoadBalancers []*GetOriginPoolResponseBodyReferencesLoadBalancers `json:"LoadBalancers,omitempty" xml:"LoadBalancers,omitempty" type:"Repeated"`
}

func (s GetOriginPoolResponseBodyReferences) String() string {
	return dara.Prettify(s)
}

func (s GetOriginPoolResponseBodyReferences) GoString() string {
	return s.String()
}

func (s *GetOriginPoolResponseBodyReferences) GetDnsRecords() []*GetOriginPoolResponseBodyReferencesDnsRecords {
	return s.DnsRecords
}

func (s *GetOriginPoolResponseBodyReferences) GetIPARecords() []*GetOriginPoolResponseBodyReferencesIPARecords {
	return s.IPARecords
}

func (s *GetOriginPoolResponseBodyReferences) GetLoadBalancers() []*GetOriginPoolResponseBodyReferencesLoadBalancers {
	return s.LoadBalancers
}

func (s *GetOriginPoolResponseBodyReferences) SetDnsRecords(v []*GetOriginPoolResponseBodyReferencesDnsRecords) *GetOriginPoolResponseBodyReferences {
	s.DnsRecords = v
	return s
}

func (s *GetOriginPoolResponseBodyReferences) SetIPARecords(v []*GetOriginPoolResponseBodyReferencesIPARecords) *GetOriginPoolResponseBodyReferences {
	s.IPARecords = v
	return s
}

func (s *GetOriginPoolResponseBodyReferences) SetLoadBalancers(v []*GetOriginPoolResponseBodyReferencesLoadBalancers) *GetOriginPoolResponseBodyReferences {
	s.LoadBalancers = v
	return s
}

func (s *GetOriginPoolResponseBodyReferences) Validate() error {
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

type GetOriginPoolResponseBodyReferencesDnsRecords struct {
	// The ID of the record.
	//
	// example:
	//
	// 104285288635****
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the record.
	//
	// example:
	//
	// www.example.com
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetOriginPoolResponseBodyReferencesDnsRecords) String() string {
	return dara.Prettify(s)
}

func (s GetOriginPoolResponseBodyReferencesDnsRecords) GoString() string {
	return s.String()
}

func (s *GetOriginPoolResponseBodyReferencesDnsRecords) GetId() *int64 {
	return s.Id
}

func (s *GetOriginPoolResponseBodyReferencesDnsRecords) GetName() *string {
	return s.Name
}

func (s *GetOriginPoolResponseBodyReferencesDnsRecords) SetId(v int64) *GetOriginPoolResponseBodyReferencesDnsRecords {
	s.Id = &v
	return s
}

func (s *GetOriginPoolResponseBodyReferencesDnsRecords) SetName(v string) *GetOriginPoolResponseBodyReferencesDnsRecords {
	s.Name = &v
	return s
}

func (s *GetOriginPoolResponseBodyReferencesDnsRecords) Validate() error {
	return dara.Validate(s)
}

type GetOriginPoolResponseBodyReferencesIPARecords struct {
	// The ID of the record.
	//
	// example:
	//
	// 104285288635****
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the record.
	//
	// example:
	//
	// ipa.example.com
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetOriginPoolResponseBodyReferencesIPARecords) String() string {
	return dara.Prettify(s)
}

func (s GetOriginPoolResponseBodyReferencesIPARecords) GoString() string {
	return s.String()
}

func (s *GetOriginPoolResponseBodyReferencesIPARecords) GetId() *int64 {
	return s.Id
}

func (s *GetOriginPoolResponseBodyReferencesIPARecords) GetName() *string {
	return s.Name
}

func (s *GetOriginPoolResponseBodyReferencesIPARecords) SetId(v int64) *GetOriginPoolResponseBodyReferencesIPARecords {
	s.Id = &v
	return s
}

func (s *GetOriginPoolResponseBodyReferencesIPARecords) SetName(v string) *GetOriginPoolResponseBodyReferencesIPARecords {
	s.Name = &v
	return s
}

func (s *GetOriginPoolResponseBodyReferencesIPARecords) Validate() error {
	return dara.Validate(s)
}

type GetOriginPoolResponseBodyReferencesLoadBalancers struct {
	// The ID of the load balancer.
	//
	// example:
	//
	// 99874066052****
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the load balancer.
	//
	// example:
	//
	// lb1.example.com
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetOriginPoolResponseBodyReferencesLoadBalancers) String() string {
	return dara.Prettify(s)
}

func (s GetOriginPoolResponseBodyReferencesLoadBalancers) GoString() string {
	return s.String()
}

func (s *GetOriginPoolResponseBodyReferencesLoadBalancers) GetId() *int64 {
	return s.Id
}

func (s *GetOriginPoolResponseBodyReferencesLoadBalancers) GetName() *string {
	return s.Name
}

func (s *GetOriginPoolResponseBodyReferencesLoadBalancers) SetId(v int64) *GetOriginPoolResponseBodyReferencesLoadBalancers {
	s.Id = &v
	return s
}

func (s *GetOriginPoolResponseBodyReferencesLoadBalancers) SetName(v string) *GetOriginPoolResponseBodyReferencesLoadBalancers {
	s.Name = &v
	return s
}

func (s *GetOriginPoolResponseBodyReferencesLoadBalancers) Validate() error {
	return dara.Validate(s)
}
